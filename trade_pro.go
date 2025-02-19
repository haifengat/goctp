package goctp

import (
	"errors"
	"fmt"
	"math"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// TradePro 业务逻辑封装
type TradePro struct {
	*TradeExt
	IsLogin      bool                               // 是否登录成功
	OnOrder      func(pOrder *CThostFtdcOrderField) // 委托响应
	OnTrade      func(pTrade *CThostFtdcTradeField) // 成交响应
	onOrderLocal func(pOrder *CThostFtdcOrderField) // 本地委托响应(用于 reqOrderInsert)

	// 合约 key: InstrumentID
	Instruments map[string]CThostFtdcInstrumentField
	// 委托 key: OrderLocalID
	Orders map[string]CThostFtdcOrderField
	// 成交 key: OrderLocalID values: []TradeField
	Trades map[string][]CThostFtdcTradeField
	// 投资者 key:InvestorID
	Investors map[string]CThostFtdcInvestorField
	// 银行开户信息
	AccountRegisters map[string]CThostFtdcAccountregisterField

	// 用于判断是否此连接的委托
	sessionID TThostFtdcSessionIDType
}

// NewTradePro
//
//	@return *TradePro 简易封装 CTP
func NewTradePro() *TradePro {
	trd := TradePro{}
	trd.TradeExt = NewTradeExt()

	// 登录过程中查询的信息
	trd.Instruments = make(map[string]CThostFtdcInstrumentField)
	trd.Investors = make(map[string]CThostFtdcInvestorField)
	trd.Orders = make(map[string]CThostFtdcOrderField)
	trd.Trades = make(map[string][]CThostFtdcTradeField)
	trd.AccountRegisters = make(map[string]CThostFtdcAccountregisterField)

	trd.Trade.OnRtnOrder = func(pOrder *CThostFtdcOrderField) {
		_, ok := trd.Orders[pOrder.OrderLocalID.String()]
		trd.Orders[pOrder.OrderLocalID.String()] = *pOrder
		if !ok { // 首次响应
			if pOrder.SessionID == trd.sessionID { // 此连接的委托
				if trd.onOrderLocal != nil {
					trd.onOrderLocal(pOrder)
				}
			}
		}
		if trd.OnOrder != nil {
			trd.OnOrder(pOrder)
		}
	}
	// 成交
	trd.Trade.OnRtnTrade = func(pTrade *CThostFtdcTradeField) {
		if _, ok := trd.Trades[pTrade.OrderLocalID.String()]; ok {
			trd.Trades[pTrade.OrderLocalID.String()] = append(trd.Trades[pTrade.OrderLocalID.String()], *pTrade)
		} else {
			trd.Trades[pTrade.OrderLocalID.String()] = []CThostFtdcTradeField{*pTrade}
		}
		if trd.OnTrade != nil {
			trd.OnTrade(pTrade)
		}
	}

	// 错误响应
	trd.OnRspError = func(pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println(*pRspInfo)
		// trd.errorChan <- *pRspInfo
	}
	return &trd
}

// LoginConfig 登录配置(行情不需要 AppID AuthCode)
type LoginConfig struct {
	Front, Broker, UserID, Password, AppID, AuthCode string
}

func (trd *TradePro) Release() {
	trd.TradeExt.Release()
	trd.TradeExt = nil
	trd.IsLogin = false
}

// Start 接口启动/登录/查询客户基础信息/查询委托/成交/权益
//
//	@receiver trd TradePro
//	@param cfg 登录配置
//	@return loginInfo 登录响应
//	@return rsp 错误响应
func (trd *TradePro) Start(cfg LoginConfig) (loginInfo CThostFtdcRspUserLoginField, rsp CThostFtdcRspInfoField) {
	var done = make(chan struct{})          // 登录完成信号
	timer := time.NewTimer(time.Second * 3) // 登录超时
	first := true                           // 首次登录

	trd.Trade.OnFrontConnected = func() {
		trd.ReqAuthenticate(cfg.Broker, cfg.UserID, cfg.AppID, cfg.AuthCode) // 认证
	}
	trd.Trade.OnFrontDisconnected = func(nReason int) {
		trd.IsLogin = false
		fmt.Println("断开连接: ", nReason)
	}
	trd.Trade.OnRspAuthenticate = func(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			if pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
			} else {
				fmt.Printf("认证: %+v\n", *pRspAuthenticateField)
				trd.TradeExt.ReqUserLogin(cfg.Password) // 登录
			}
		}
	}
	trd.Trade.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			if pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
			} else {
				trd.IsLogin = true
				if first {
					first = false
					trd.sessionID = pRspUserLogin.SessionID
					loginInfo = *pRspUserLogin
					timer.Reset(time.Minute * 3)  // 重置超时
					trd.TradeExt.ReqQryInvestor() // 查用户
				}
			}
		}
	}
	trd.Trade.OnRspQryInvestor = func(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestor != nil {
			trd.Investors[pInvestor.InvestorID.String()] = *pInvestor
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
			} else {
				time.Sleep(time.Millisecond * 1100)
				trd.TradeExt.ReqQryClassifiedInstrument() // 查合约
			}
		}
	}
	trd.Trade.OnRspQryClassifiedInstrument = func(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInstrument != nil {
			trd.Instruments[pInstrument.InstrumentID.String()] = *pInstrument
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
			} else {
				if _, exists := trd.Investors[trd.UserID]; !exists { // 交易员登录: 跳过查询过程
					time.Sleep(time.Millisecond * 1100)
					trd.TradeExt.ReqQryAccountregister() // 查银期签约
				} else {
					trd.TradeExt.ReqSettlementInfoConfirm() // 确认结算
				}
			}
		}
	}
	trd.Trade.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			// 交易员无确认结算权限,此处忽略
			if pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
				fmt.Printf("确认结算错误: %+v\n", *pRspInfo)
			} //else {
			time.Sleep(time.Millisecond * 1100)
			trd.TradeExt.ReqQryOrder() // 查委托
			// }
		}
	}
	trd.Trade.OnRspQryOrder = func(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			rsp = *pRspInfo
		} else if pOrder != nil {
			trd.Orders[pOrder.OrderLocalID.String()] = *pOrder
		}
		if bIsLast {
			time.Sleep(time.Millisecond * 1100)
			trd.TradeExt.ReqQryTrade() // 查成交
		}
	}
	trd.Trade.OnRspQryTrade = func(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pTrade != nil {
			if _, ok := trd.Trades[pTrade.OrderLocalID.String()]; ok {
				trd.Trades[pTrade.OrderLocalID.String()] = append(trd.Trades[pTrade.OrderLocalID.String()], *pTrade)
			} else {
				trd.Trades[pTrade.OrderLocalID.String()] = []CThostFtdcTradeField{*pTrade}
			}
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
			} else {
				time.Sleep(time.Millisecond * 1100)
				trd.TradeExt.ReqQryAccountregister() // 查银期签约
			}
		}
	}
	// 银期
	trd.OnRspQryAccountregister = func(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pAccountregister != nil {
			trd.AccountRegisters[pAccountregister.BankAccount.String()] = *pAccountregister
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				rsp = *pRspInfo
			} else {
				bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("正确"))
				copy(rsp.ErrorMsg[:], bs)
			}
			done <- struct{}{}
		}
	}
	// 登录
	trd.TradeExt.RegisterFront(cfg.Front)
	trd.TradeExt.SubscribePrivateTopic(THOST_TERT_QUICK)
	trd.TradeExt.SubscribePublicTopic(THOST_TERT_RESTART)
	trd.TradeExt.Init()

	select {
	case <-done: // 登录成功
	case c := <-timer.C: // 登录超时
		if rsp.ErrorID == 0 { // 登录过程中无错误
			str, _ := simplifiedchinese.GB18030.NewEncoder().String("登录超时 " + c.String())
			rsp.ErrorID = -1
			copy(rsp.ErrorMsg[:], str)
		}
	}
	return
}

// ReqOrderInsertLimit 限价单
//
//	@receiver trd TradePro
//	@param buySell 买卖
//	@param openClose 开平
//	@param instrument 合约
//	@param price 价格
//	@param volume 手数
//	@return localID 成功返回本地编号
//	@return rsp 错误信息
func (trd *TradePro) ReqOrderInsertLimit(instrument string, buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	return trd.reqOrderInsert(instrument, buySell, openClose, price, volume, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_GFD, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)
}

// ReqOrderInsertFAK 部成全撤
//
//	@receiver trd TradePro
//	@param buySell 买卖
//	@param openClose 开平
//	@param instrument 合约
//	@param price 价格
//	@param volume 手数
//	@return localID 成功返回本地编号
//	@return rsp 错误信息
func (trd *TradePro) ReqOrderInsertFAK(instrument string, buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	return trd.reqOrderInsert(instrument, buySell, openClose, price, volume, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_IOC, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)
}

// ReqOrderInsertFOK 全成or撤单
//
//	@receiver trd TradePro
//	@param buySell 买卖
//	@param openClose 开平
//	@param instrument 合约
//	@param price 价格
//	@param volume 手数
//	@return localID 成功返回本地编号
//	@return rsp 错误信息
func (trd *TradePro) ReqOrderInsertFOK(instrument string, buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, price float64, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	return trd.reqOrderInsert(instrument, buySell, openClose, price, volume, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_IOC, THOST_FTDC_VC_CV, THOST_FTDC_CC_Immediately)
}

// ReqOrderInsertMarket 市价单(不是所有交易所都支持)
//
//	@receiver trd TradePro
//	@param buySell 买卖
//	@param openClose 开平
//	@param instrument 合约
//	@param price 价格
//	@param volume 手数
//	@return localID 成功返回本地编号
//	@return rsp 错误信息
func (trd *TradePro) ReqOrderInsertMarket(instrument string, buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, volume int) (localID string, rsp CThostFtdcRspInfoField) {
	return trd.reqOrderInsert(instrument, buySell, openClose, 0, volume, THOST_FTDC_OPT_AnyPrice, THOST_FTDC_TC_IOC, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)
}

// reqOrderInsert ReqOrderInsertMarket 市价单(不是所有交易所都支持)
//
//	@param instrument 合约
//	@param buySell 买卖
//	@param openClose 开平
//	@param price 价格
//	@param volume 手数
//	@param priceType TThostFtdcOrderPriceTypeType
//	@param timeType TThostFtdcTimeConditionType
//	@param volumeType TThostFtdcVolumeConditionType
//	@param contingentType TThostFtdcContingentConditionType
//	@return localID 成功返回本地编号
//	@return rsp 错误信息
func (trd *TradePro) reqOrderInsert(instrument string, buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, price float64, volume int, priceType TThostFtdcOrderPriceTypeType, timeType TThostFtdcTimeConditionType, volumeType TThostFtdcVolumeConditionType, contingentType TThostFtdcContingentConditionType) (localID string, rsp CThostFtdcRspInfoField) {
	done := make(chan struct{})
	defer func() {
		trd.Trade.OnRspOrderInsert = nil
		trd.onOrderLocal = nil
		close(done)
	}()

	// 委托
	trd.Trade.OnRspOrderInsert = func(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		rsp = *pRspInfo
		done <- struct{}{}
	}
	trd.onOrderLocal = func(pOrder *CThostFtdcOrderField) {
		localID = pOrder.OrderLocalID.String()
		done <- struct{}{}
	}

	inst, exists := trd.Instruments[instrument]
	if !exists {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此合约:" + instrument))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	exchange := inst.ExchangeID.String()

	if priceType == THOST_FTDC_OPT_LimitPrice { // 限价单 最小变动的倍数
		price = math.Round(price/float64(inst.PriceTick)) * float64(inst.PriceTick)
	}

	if rtn := trd.TradeExt.ReqOrderInsert(instrument, buySell, openClose, price, volume, priceType, timeType, volumeType, contingentType, exchange, trd.InvestorID); rtn != 0 { // 流控
		rsp.ErrorID = -2
		str, _ := simplifiedchinese.GB18030.NewEncoder().String(fmt.Sprintf("流控: %d", rtn))
		copy(rsp.ErrorMsg[:], str)
		return
	}

	// f.RequestID = TThostFtdcRequestIDType(id)
	// copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	localID = fmt.Sprintf("%d-%s", trd.sessionID, fmt.Sprintf("%012d", trd.id))
	select {
	case <-done:
	case c := <-time.NewTimer(2 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout: "+c.String())
	}
	return
}

// ReqOrderAction 撤单
//
//	@receiver trd TradePro
//	@param localID 本地报单编号
//	@return int -9:未查到localID对应的委托
func (trd *TradePro) ReqOrderAction(localID string) int {
	of, exists := trd.Orders[localID]
	if exists {
		return trd.TradeExt.ReqOrderAction(of.InvestorID.String(), of.ExchangeID.String(), of.InstrumentID.String(), of.OrderRef.String(), int(of.SessionID), int(of.FrontID))
	}
	return -9
}

// ReqFromBankToFutureByFuture 入金
//
//	@receiver trd TradePro
//	@param bankAccount 银行帐号
//	@param bankPwd 银行密码
//	@param accountPwd 出入金密码
//	@param amount 出入金金额
//	@return rsp 错误响应
func (trd *TradePro) ReqFromBankToFutureByFuture(bankAccount, bankPwd, accountPwd string, amount float64) (rsp CThostFtdcRspInfoField) {
	done := make(chan struct{})
	defer func() {
		trd.Trade.OnRspFromBankToFutureByFuture = nil
		trd.Trade.OnRtnFromBankToFutureByFuture = nil
		close(done)
	}()

	// 银转:入金
	trd.Trade.OnRspFromBankToFutureByFuture = func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast { // 有主动请求时才响应
			rsp = *pRspInfo
			done <- struct{}{}
		}
	}
	trd.Trade.OnRtnFromBankToFutureByFuture = func(pRspTransfer *CThostFtdcRspTransferField) {
		rsp.ErrorID = pRspTransfer.ErrorID
		copy(rsp.ErrorMsg[:], pRspTransfer.ErrorMsg[:])
		done <- struct{}{}
	}

	regInfo, ok := trd.AccountRegisters[bankAccount]
	if !ok {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此帐号:" + bankAccount))
		copy(rsp.ErrorMsg[:], bs)
		return
	}

	if rtn := trd.TradeExt.ReqFromBankToFutureByFuture(regInfo, bankPwd, accountPwd, amount); rtn != 0 { // 流控
		rsp.ErrorID = -2
		str, _ := simplifiedchinese.GB18030.NewEncoder().String(fmt.Sprintf("流控: %d", rtn))
		copy(rsp.ErrorMsg[:], str)
		return
	}
	select {
	case <-done:
	case c := <-time.NewTimer(3 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout "+c.String())
	}
	return
}

// ReqFromFutureToBankByFuture 出金
//
//	@receiver trd TradePro
//	@param bankAccount 银行帐号
//	@param bankPwd 银行密码
//	@param accountPwd 出入金密码
//	@param amount 出入金金额
//	@return rsp 错误响应
func (trd *TradePro) ReqFromFutureToBankByFuture(bankAccount, accountPwd string, amount float64) (rsp CThostFtdcRspInfoField) {
	done := make(chan struct{})
	defer func() {
		trd.Trade.OnRspFromFutureToBankByFuture = nil
		trd.Trade.OnRtnFromFutureToBankByFuture = nil
		close(done)
	}()
	// 银转:出金
	trd.Trade.OnRspFromFutureToBankByFuture = func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			rsp = *pRspInfo
			done <- struct{}{}
		}
	}
	trd.Trade.OnRtnFromFutureToBankByFuture = func(pRspTransfer *CThostFtdcRspTransferField) {
		rsp.ErrorID = pRspTransfer.ErrorID
		copy(rsp.ErrorMsg[:], pRspTransfer.ErrorMsg[:])
		done <- struct{}{}
	}

	regInfo, ok := trd.AccountRegisters[bankAccount]
	if !ok {
		rsp.ErrorID = -1
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("无此帐号:" + bankAccount))
		copy(rsp.ErrorMsg[:], bs)
		return
	}
	if rtn := trd.TradeExt.ReqFromFutureToBankByFuture(regInfo, accountPwd, amount); rtn != 0 { // 流控
		rsp.ErrorID = -2
		str, _ := simplifiedchinese.GB18030.NewEncoder().String(fmt.Sprintf("流控: %d", rtn))
		copy(rsp.ErrorMsg[:], str)
		return
	}
	select {
	case <-done:
	case c := <-time.NewTimer(3 * time.Second).C:
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], "timeout "+c.String())
	}
	return
}

// ReqQryPosition 查持仓
//
//	@receiver trd TradePro
//	@return []CThostFtdcInvestorPositionField 返回 nil 时注意流控
func (trd *TradePro) ReqQryPosition() (positions []CThostFtdcInvestorPositionField, errRtn error) {
	positions = make([]CThostFtdcInvestorPositionField, 0)
	done := make(chan struct{})
	defer func() { // 恢复空事件
		trd.Trade.OnRspQryInvestorPosition = nil
		close(done)
	}()

	trd.Trade.OnRspQryInvestorPosition = func(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestorPosition != nil {
			positions = append(positions, *pInvestorPosition)
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				errRtn = errors.New(pRspInfo.ErrorMsg.String())
			} else {
				done <- struct{}{}
			}
		}
	}

	go func() {
		var i int
		for i = 0; i < 3; i++ { // 3 次流控
			if n := trd.TradeExt.ReqQryInvestorPosition(); n == 0 {
				break
			} else {
				fmt.Println("ReqQryInvestorPosition: ", n)
			}
			time.Sleep(time.Second * 2)
		}
		if i == 3 {
			fmt.Println("被流控 3 次, 查询失败")
		}
	}()

	select {
	case <-done:
	case c := <-time.NewTimer(time.Second * 8).C:
		errRtn = errors.New("ReqQryInvestorPosition 超时: " + c.String())
	}
	return
}

// ReqQryPositionDetail 查持仓明细
//
//	@receiver trd TradePro
//	@return []CThostFtdcInvestorPositionDetailField 持仓明细, 返回 nil 时注意流控
func (trd *TradePro) ReqQryPositionDetail() (positionDetails []CThostFtdcInvestorPositionDetailField, errRtn error) {
	positionDetails = make([]CThostFtdcInvestorPositionDetailField, 0)
	done := make(chan struct{})
	defer func() { // 恢复空事件
		trd.Trade.OnRspQryInvestorPositionDetail = nil
		close(done)
	}()

	// 持仓明细
	trd.Trade.OnRspQryInvestorPositionDetail = func(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestorPositionDetail != nil {
			positionDetails = append(positionDetails, *pInvestorPositionDetail)
		}
		if bIsLast {
			if pRspInfo != nil && pRspInfo.ErrorID != 0 {
				errRtn = errors.New(pRspInfo.ErrorMsg.String())
			} else {
				done <- struct{}{}
			}
		}
	}

	if n := trd.TradeExt.ReqQryInvestorPositionDetail(); n != 0 {
		errRtn = fmt.Errorf("权益查询失败: %d", n)
		return
	}

	select {
	case <-done:
	case c := <-time.NewTimer(time.Second * 8).C:
		errRtn = errors.New("ReqQryInvestorPositionDetail 超时: " + c.String())
	}
	return
}

// ReqQryTradingAccount 查帐户权益
//
//	@receiver trd TradePro
//	@return map 投资者帐号:权益
func (trd *TradePro) ReqQryTradingAccount() (accounts map[string]CThostFtdcTradingAccountField, errRtn error) {
	// 权益 查询时返回
	accounts = make(map[string]CThostFtdcTradingAccountField)
	done := make(chan struct{})
	defer func() { // 恢复空事件
		trd.Trade.OnRspQryTradingAccount = nil
		close(done)
	}()

	// 权益
	trd.Trade.OnRspQryTradingAccount = func(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pTradingAccount != nil {
			accounts[pTradingAccount.AccountID.String()] = *pTradingAccount
		}
		if bIsLast {
			done <- struct{}{}
		}
	}

	if n := trd.TradeExt.ReqQryTradingAccount(); n != 0 {
		errRtn = fmt.Errorf("权益查询失败: %d", n)
		return
	}

	select {
	case <-done:
	case c := <-time.NewTimer(time.Second * 8).C:
		errRtn = errors.New("ReqQryTradingAccount 超时: " + c.String())
	}
	return
}
