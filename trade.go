package goctp

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"unsafe"

	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

// HFTrade 交易接口
type HFTrade struct {
	InvestorID string // 帐号
	BrokerID   string // 经纪商
	TradingDay string // 交易日
	passWord   string
	SessionID  int // 判断是否自己的委托用

	Instruments       sync.Map      // 合约列表 (key: InstrumentID, value: *InstrumentField)
	InstrumentStatuss sync.Map      // 合约状态 (key: InstrumentID, value: *InstrumentStatus)
	posiDetail        sync.Map      // 原始持仓
	Positions         sync.Map      // 合成后的持仓 (key: instrument_long/short value: *ctp.CThostFtdcInvestorPositionField)
	Orders            sync.Map      // 委托 (key: sessionID_OrderRef, value: *OrderField)
	Trades            sync.Map      // 成交 (key: TradeID_buy/sell, value: &TradeField)
	sysID4Order       sync.Map      // key:OrderSysID,value: *OrderField
	Account           *AccountField // 帐户权益

	IsLogin bool   // 登录成功
	Version string // 版本号,如 v6.5.1_20200908 10:25:08

	qryTicker *time.Ticker   // 循环查询
	waitGroup sync.WaitGroup // 登录信号

	reqID    int // requestid
	cntOrder int // 计算order数量

	onFrontConnected      OnFrontConnectedType // 事件
	onFrontDisConnected   OnFrontDisConnectedType
	onRspUserLogin        OnRspUserLoginType
	onRtnOrder            OnRtnOrderType
	onRtnCancel           OnRtnOrderType
	onErrRtnOrder         OnRtnErrOrderType
	onErrAction           OnRtnErrActionType
	onRtnTrade            OnRtnTradeType
	onRtnInstrumentStatus OnRtnInstrumentStatusType
	onRtnBankToFuture     OnRtnFromBankToFutureByFuture
	onRtnFutureToBank     OnRtnFromFutureToBankByFuture

	// 继承类要实现的函数
	ReqConnect                  ReqConnectType
	ReleaseAPI                  ReleaseAPIType
	ReqUserLogin                ReqUserLoginType
	ReqAuthenticate             ReqAuthenticateType
	ReqSettlementInfoConfirm    ReqSettlementInfoConfirmType
	ReqQryInstrument            ReqQryInstrumentType
	ReqQryClassifiedInstrument  ReqQryClassifiedInstrumentType
	ReqQryTradingAccount        ReqQryTradingAccountType
	ReqQryInvestorPosition      ReqQryInvestorPositionType
	ReqOrder                    ReqOrderInsertType
	ReqAction                   ReqOrderActionType
	ReqFromBankToFutureByFuture ReqTransferType
	ReqFromFutureToBankByFuture ReqTransferType
	GetVersion                  GetVersionType
}
type ReqAuthenticateType func(*ctp.CThostFtdcReqAuthenticateField, int)
type ReqUserLoginType func(*ctp.CThostFtdcReqUserLoginField, int)
type ReqSettlementInfoConfirmType func(*ctp.CThostFtdcSettlementInfoConfirmField, int)
type ReqQryInstrumentType func(*ctp.CThostFtdcQryInstrumentField, int)
type ReqQryClassifiedInstrumentType func(*ctp.CThostFtdcQryClassifiedInstrumentField, int)
type ReqQryTradingAccountType func(*ctp.CThostFtdcQryTradingAccountField, int)
type ReqQryInvestorPositionType func(*ctp.CThostFtdcQryInvestorPositionField, int)
type ReqOrderInsertType func(*ctp.CThostFtdcInputOrderField, int)
type ReqOrderActionType = func(*ctp.CThostFtdcInputOrderActionField, int)
type ReqTransferType = func(*ctp.CThostFtdcReqTransferField, int)
type ReqConnectType = func(string)
type ReleaseAPIType func()
type GetVersionType func() string

func (t *HFTrade) Init() {
	for _, r := range []interface{}{t.ReqAuthenticate, t.ReqUserLogin, t.ReqSettlementInfoConfirm, t.ReqQryInstrument, t.ReqQryClassifiedInstrument, t.ReqQryTradingAccount, t.ReqQryInvestorPosition, t.ReqOrder, t.ReqAction, t.GetVersion} {
		if r == nil {
			panic("缺少继承函数")
		}
	}
	t.Version = t.GetVersion()
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
	t.waitGroup = sync.WaitGroup{}
	t.Account = new(AccountField)
}

func (t *HFTrade) Release() {
	t.qryTicker.Stop()
	t.IsLogin = false
	t.ReleaseAPI()
	t.FrontDisConnected(0) // 需手动触发
}

func (t *HFTrade) getReqID() int {
	t.reqID++
	return t.reqID
}

// ReqLogin 登录
func (t *HFTrade) ReqLogin(investor, pwd, broker, appID, authCode string) {
	t.InvestorID = investor
	t.passWord = pwd
	t.BrokerID = broker
	f := ctp.CThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.AppID[:], appID)
	copy(f.AuthCode[:], authCode)
	t.ReqAuthenticate(&f, t.getReqID())
}

//------------------- 函数封装 ----------------------

// ReqOrderInsert 限价委托
func (t *HFTrade) ReqOrderInsert(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_GFD
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.ReqOrder(&f, id)
	return fmt.Sprintf("%d_%s", t.SessionID, Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertMarket 市价委托
func (t *HFTrade) ReqOrderInsertMarket(instrument string, buySell DirectionType, openClose OffsetFlagType, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_AnyPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_IOC
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(0)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.ReqOrder(&f, id)
	return fmt.Sprintf("%d_%s", t.SessionID, Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertFOK FOK委托[部成撤单]
func (t *HFTrade) ReqOrderInsertFOK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_GFD
	f.VolumeCondition = ctp.THOST_FTDC_VC_CV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.ReqOrder(&f, id)
	return fmt.Sprintf("%d_%s", t.SessionID, Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertFAK FAK委托[全成or撤单]
func (t *HFTrade) ReqOrderInsertFAK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_IOC
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.ReqOrder(&f, id)
	return fmt.Sprintf("%d_%s", t.SessionID, Bytes2String(f.OrderRef[:]))
}

// ReqOrderAction 撤单
func (t *HFTrade) ReqOrderAction(orderID string) int {
	if o, ok := t.Orders.Load(orderID); ok {
		var order = o.(*OrderField)
		f := ctp.CThostFtdcInputOrderActionField{}
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.UserID[:], t.InvestorID)
		copy(f.InstrumentID[:], order.InstrumentID)
		copy(f.ExchangeID[:], order.ExchangeID)
		copy(f.OrderRef[:], order.OrderRef)
		f.ActionFlag = ctp.THOST_FTDC_AF_Delete
		f.FrontID = ctp.TThostFtdcFrontIDType(order.FrontID)
		f.SessionID = ctp.TThostFtdcSessionIDType(order.SessionID)
		t.ReqAction(&f, t.getReqID())
		return 0
	}
	return -1
}

// ReqBankToFuture 银行转期货
func (t *HFTrade) ReqBankToFuture(bankID, bankAccount, bankPwd string, amount float64) {
	f := ctp.CThostFtdcReqTransferField{}
	copy(f.TradeCode[:], "202001")
	copy(f.BankBranchID[:], "0000")
	copy(f.BrokerID[:], []byte(t.BrokerID))
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.Password[:], []byte(t.passWord))
	copy(f.CurrencyID[:], "CNY")
	f.LastFragment = ctp.THOST_FTDC_LF_Yes
	f.IdCardType = ctp.THOST_FTDC_ICT_IDCard
	f.CustType = ctp.THOST_FTDC_CUSTT_Person
	f.InstallID = 1
	f.FutureSerial = 0
	f.VerifyCertNoFlag = ctp.THOST_FTDC_YNI_No
	f.FutureFetchAmount = 0
	f.CustFee = 0
	f.BrokerFee = 0
	f.SecuPwdFlag = ctp.THOST_FTDC_BPWDF_BlankCheck
	f.RequestID = ctp.TThostFtdcRequestIDType(t.getReqID())
	f.TID = 0
	copy(f.BankID[:], []byte(bankID))
	copy(f.BankAccount[:], []byte(bankAccount))
	copy(f.BankPassWord[:], []byte(bankPwd))
	f.TradeAmount = ctp.TThostFtdcTradeAmountType(amount)
	t.ReqFromBankToFutureByFuture(&f, t.getReqID())
}

// ReqFutureToBank 期货转银行
func (t *HFTrade) ReqFutureToBank(bankID, bankAccount string, amount float64) {
	f := ctp.CThostFtdcReqTransferField{}
	copy(f.TradeCode[:], "202002")
	copy(f.BankBranchID[:], "0000")
	copy(f.BrokerID[:], []byte(t.BrokerID))
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.Password[:], []byte(t.passWord))
	copy(f.CurrencyID[:], "CNY")
	f.LastFragment = ctp.THOST_FTDC_LF_Yes
	f.IdCardType = ctp.THOST_FTDC_ICT_IDCard
	f.CustType = ctp.THOST_FTDC_CUSTT_Person
	f.InstallID = 1
	f.FutureSerial = 0
	f.VerifyCertNoFlag = ctp.THOST_FTDC_YNI_No
	f.FutureFetchAmount = 0
	f.CustFee = 0
	f.BrokerFee = 0
	f.SecuPwdFlag = ctp.THOST_FTDC_BPWDF_BlankCheck
	f.RequestID = ctp.TThostFtdcRequestIDType(t.getReqID())
	f.TID = 0
	copy(f.BankID[:], []byte(bankID))
	copy(f.BankAccount[:], []byte(bankAccount))
	// copy(f.BankPassWord[:], []byte(bankPwd))
	f.TradeAmount = ctp.TThostFtdcTradeAmountType(amount)
	t.ReqFromFutureToBankByFuture(&f, t.getReqID())
}

//-------------------- 响应封装 -----------------------

// RegOnFrontConnected 注册连接响应
func (t *HFTrade) RegOnFrontConnected(on OnFrontConnectedType) {
	t.onFrontConnected = on
}

// RegOnFrontDisConnected 注册连接响应
func (t *HFTrade) RegOnFrontDisConnected(on OnFrontDisConnectedType) {
	t.onFrontDisConnected = on
}

// RegOnRspUserLogin 注册登陆响应
func (t *HFTrade) RegOnRspUserLogin(on OnRspUserLoginType) {
	t.onRspUserLogin = on
}

// RegOnRtnOrder 注册委托响应
func (t *HFTrade) RegOnRtnOrder(on OnRtnOrderType) {
	t.onRtnOrder = on
}

// RegOnErrRtnOrder 注册委托响应
func (t *HFTrade) RegOnErrRtnOrder(on OnRtnErrOrderType) {
	t.onErrRtnOrder = on
}

// RegOnErrAction 注册撤单响应
func (t *HFTrade) RegOnErrAction(on OnRtnErrActionType) {
	t.onErrAction = on
}

// RegOnRtnCancel 注册撤单响应
func (t *HFTrade) RegOnRtnCancel(on OnRtnOrderType) {
	t.onRtnCancel = on
}

// RegOnRtnTrade 注册成交响应
func (t *HFTrade) RegOnRtnTrade(on OnRtnTradeType) {
	t.onRtnTrade = on
}

// RegOnRtnInstrumentStatus 注册合约状态变化
func (t *HFTrade) RegOnRtnInstrumentStatus(on OnRtnInstrumentStatusType) {
	t.onRtnInstrumentStatus = on
}

// RegOnRtnFromBankToFuture 注册银行转期货
func (t *HFTrade) RegOnRtnFromBankToFuture(on OnRtnFromBankToFutureByFuture) {
	t.onRtnBankToFuture = on
}

// RegOnRtnFromFutureToBank 注册期货转银行
func (t *HFTrade) RegOnRtnFromFutureToBank(on OnRtnFromFutureToBankByFuture) {
	t.onRtnFutureToBank = on
}

// RtnFromBankToFutureByFuture 银行转期货-期货端
func (t *HFTrade) RtnFromBankToFutureByFuture(field *ctp.CThostFtdcRspTransferField) {
	if t.onRtnBankToFuture != nil {
		f := TransferField{
			Amout:      float64(field.TradeAmount),
			CurrencyID: Bytes2String(field.CurrencyID[:]),
			ErrorID:    int(field.ErrorID),
			ErrorMsg:   Bytes2String(field.ErrorMsg[:]),
		}
		t.onRtnBankToFuture(&f)
	}
}

// RtnFromFutureToBankByFuture // 期货转银行-期货端
func (t *HFTrade) RtnFromFutureToBankByFuture(field *ctp.CThostFtdcRspTransferField) {
	if t.onRtnFutureToBank != nil {
		f := TransferField{
			Amout:      float64(field.TradeAmount),
			CurrencyID: Bytes2String(field.CurrencyID[:]),
			ErrorID:    int(field.ErrorID),
			ErrorMsg:   Bytes2String(field.ErrorMsg[:]),
		}
		t.onRtnFutureToBank(&f)
	}
}

// RtnInstrumentStatus 合约/品种/交易所 状态响应
func (t *HFTrade) RtnInstrumentStatus(field *ctp.CThostFtdcInstrumentStatusField) {
	status, loaded := t.InstrumentStatuss.LoadOrStore(Bytes2String(field.InstrumentID[:]), &InstrumentStatus{
		ExchangeID:       Bytes2String(field.ExchangeID[:]),
		InstrumentID:     Bytes2String(field.InstrumentID[:]),
		InstrumentStatus: InstrumentStatusType(field.InstrumentStatus),
		EnterTime:        Bytes2String(field.EnterTime[:]),
	})
	if loaded {
		status.(*InstrumentStatus).InstrumentStatus = InstrumentStatusType(field.InstrumentStatus)
		status.(*InstrumentStatus).EnterTime = Bytes2String(field.EnterTime[:])
	}

	if t.onRtnInstrumentStatus != nil {
		t.onRtnInstrumentStatus(status.(*InstrumentStatus))
	}
}

// RtnTrade 成交响应
func (t *HFTrade) RtnTrade(field *ctp.CThostFtdcTradeField) {
	var key string
	tradeid := Bytes2String(field.TradeID[:])
	if field.Direction == ctp.THOST_FTDC_D_Buy {
		key = fmt.Sprintf("%s_buy", tradeid)
	} else if field.Direction == ctp.THOST_FTDC_D_Sell {
		key = fmt.Sprintf("%s_sell", tradeid)
	} else {
		key = "error"
	}
	tf, _ := t.Trades.LoadOrStore(key, &TradeField{
		Direction:    DirectionType(field.Direction),
		HedgeFlag:    HedgeFlagType(field.HedgeFlag),
		InstrumentID: Bytes2String(field.InstrumentID[:]),
		ExchangeID:   Bytes2String(field.ExchangeID[:]),
		TradingDay:   Bytes2String(field.TradingDay[:]),
		Volume:       int(field.Volume),
		OffsetFlag:   OffsetFlagType(field.OffsetFlag),
		OrderSysID:   Bytes2String(field.OrderSysID[:]),
		Price:        float64(field.Price),
		TradeDate:    Bytes2String(field.TradeDate[:]),
		TradeTime:    Bytes2String(field.TradeTime[:]),
		TradeID:      tradeid,
	})
	var f = tf.(*TradeField)
	if t.IsLogin { // 登录后：更新持仓
		if f.OffsetFlag == OffsetFlagOpen { // 开仓
			var key string
			if f.Direction == DirectionBuy {
				key = fmt.Sprintf("%s_long", f.InstrumentID)
			} else {
				key = fmt.Sprintf("%s_short", f.InstrumentID)
			}
			pf, _ := t.Positions.LoadOrStore(key, &PositionField{
				InstrumentID:      f.InstrumentID,
				PositionDirection: PosiDirectionLong,
				HedgeFlag:         f.HedgeFlag,
				ExchangeID:        f.ExchangeID,
			})
			var p = pf.(*PositionField)
			p.OpenVolume += f.Volume
			p.OpenAmount += f.Price * float64(f.Volume)
			if info, ok := t.Instruments.Load(f.InstrumentID); ok {
				p.OpenCost += f.Price * float64(f.Volume) * float64(info.(*InstrumentField).VolumeMultiple)
			}
			p.Position += f.Volume
			p.TodayPosition += f.Volume
		} else {
			var key string
			if f.Direction == DirectionBuy {
				key = fmt.Sprintf("%s_short", f.InstrumentID)
			} else {
				key = fmt.Sprintf("%s_long", f.InstrumentID)
			}
			if posi, ok := t.Positions.Load(key); ok {
				var p = posi.(*PositionField)
				p.OpenVolume -= f.Volume
				p.OpenAmount -= f.Price * float64(f.Volume)
				if info, ok := t.Instruments.Load(f.InstrumentID); ok {
					p.OpenCost -= f.Price * float64(f.Volume) * float64(info.(*InstrumentField).VolumeMultiple)
				}
				p.Position -= f.Volume
				if f.OffsetFlag == OffsetFlagCloseToday {
					p.TodayPosition -= f.Volume
				} else {
					p.YdPosition -= f.Volume
				}
			}
		}
	}
	// 处理对应的Order
	if ord, ok := t.sysID4Order.Load(f.OrderSysID); ok {
		var o = ord.(*OrderField)
		if o.TradePrice == 0 { // 在 VolumeLeft 前计算
			o.TradePrice = f.Price
		} else { // 计算均价
			o.TradePrice = (o.TradePrice*float64(o.VolumeTotalOriginal-o.VolumeLeft) + f.Price*float64(f.Volume)) / float64(o.VolumeTotalOriginal-o.VolumeLeft+f.Volume)
		}
		o.LastTradeTime = f.TradeTime
		o.VolumeTraded = f.Volume
		o.VolumeLeft -= f.Volume
		if o.VolumeLeft == 0 {
			o.OrderStatus = OrderStatusAllTraded
			o.StatusMsg = "全部成交"
		} else {
			o.OrderStatus = OrderStatusPartTradedQueueing
			o.StatusMsg = "部分成交"
		}
		if t.IsLogin && t.onRtnOrder != nil {
			t.onRtnOrder(o)
		}
	}
	// 客户端响应
	if t.IsLogin && t.onRtnTrade != nil {
		t.onRtnTrade(f)
	}
}

// RtnOrder 委托响应
func (t *HFTrade) RtnOrder(field *ctp.CThostFtdcOrderField) {
	t.cntOrder++
	key := fmt.Sprintf("%d_%s", field.SessionID, Bytes2String(field.OrderRef[:]))
	if of, exists := t.Orders.LoadOrStore(key, &OrderField{
		InstrumentID:        Bytes2String(field.InstrumentID[:]),
		SessionID:           int(field.SessionID),
		FrontID:             int(field.FrontID),
		OrderRef:            Bytes2String(field.OrderRef[:]),
		Direction:           DirectionType(field.Direction),
		OffsetFlag:          OffsetFlagType(field.CombOffsetFlag[0]),
		HedgeFlag:           HedgeFlagType(field.CombHedgeFlag[0]),
		LimitPrice:          float64(field.LimitPrice),
		VolumeTotalOriginal: int(field.VolumeTotalOriginal),
		VolumeLeft:          int(field.VolumeTotalOriginal),
		ExchangeID:          Bytes2String(field.ExchangeID[:]),
		InsertDate:          Bytes2String(field.InsertDate[:]),
		InsertTime:          Bytes2String(field.InsertTime[:]),
		OrderStatus:         OrderStatusNoTradeQueueing, // OrderStatusType(orderField.OrderStatus)
		StatusMsg:           "委托已提交",                    // bytes2GBKbytes2GBKString(orderField.StatusMsg[:])
		IsLocal:             int(field.SessionID) == t.SessionID,
	}); !exists { // 新添加
		if t.IsLogin && t.onRtnOrder != nil {
			// 平仓指令, 冻结持仓(随后的持仓查询会进行修正),冻结持仓恢复会滞后 <=2s
			order := of.(*OrderField)
			if order.OffsetFlag != OffsetFlagOpen {
				if order.Direction == DirectionBuy { // 冻结空头
					key := fmt.Sprintf("%s_short", order.InstrumentID)
					if posiField, ok := t.Positions.Load(key); ok {
						posiField.(*PositionField).LongFrozen += order.VolumeTotalOriginal
					}
				} else {
					key := fmt.Sprintf("%s_long", order.InstrumentID)
					if posiField, ok := t.Positions.Load(key); ok { // 冻结多头
						posiField.(*PositionField).ShortFrozen += order.VolumeTotalOriginal
					}
				}
			}
			t.onRtnOrder(order)
		}
	} else {
		var o = of.(*OrderField)
		if OrderStatusType(field.OrderStatus) == OrderStatusCanceled { // 处理撤单
			o.OrderStatus = OrderStatusCanceled
			o.StatusMsg = Bytes2String(field.StatusMsg[:])
			o.CancelTime = Bytes2String(field.CancelTime[:])
			// 错单
			if t.IsLogin { // 登录前不响应
				if strings.Contains(o.StatusMsg, "被拒绝") {
					if t.onErrRtnOrder != nil {
						t.onErrRtnOrder(o, &RspInfoField{
							ErrorID:  -1,
							ErrorMsg: o.StatusMsg,
						})
					}
				} else if t.onRtnCancel != nil {
					t.onRtnCancel(o)
				}
			}
		} else {
			o.OrderSysID = Bytes2String(field.OrderSysID[:])
			if len(o.OrderSysID) > 0 {
				t.sysID4Order.Store(o.OrderSysID, o)
			}
		}
	}
}

// ErrRtnOrderAction 撤单错误
func (t *HFTrade) ErrRtnOrderAction(field *ctp.CThostFtdcOrderActionField, info *ctp.CThostFtdcRspInfoField) {
	if t.IsLogin && t.onErrAction != nil {
		t.onErrAction(fmt.Sprintf("%d_%s", field.SessionID, field.OrderRef), &RspInfoField{
			ErrorID:  int(info.ErrorID),
			ErrorMsg: Bytes2String(info.ErrorMsg[:]),
		})
	}
}

// ErrRtnOrderInsert 委托错误
func (t *HFTrade) ErrRtnOrderInsert(field *ctp.CThostFtdcInputOrderField, info *ctp.CThostFtdcRspInfoField) {
	if !t.IsLogin { // 过滤当日以前登录时的错误委托
		return
	}
	key := fmt.Sprintf("%d_%s", t.SessionID, Bytes2String(field.OrderRef[:]))
	of, _ := t.Orders.LoadOrStore(key, &OrderField{
		InstrumentID:        Bytes2String(field.InstrumentID[:]),
		SessionID:           t.SessionID,
		FrontID:             0,
		OrderRef:            Bytes2String(field.OrderRef[:]),
		Direction:           DirectionType(field.Direction),
		OffsetFlag:          OffsetFlagType(field.CombOffsetFlag[0]),
		HedgeFlag:           HedgeFlagType(field.CombHedgeFlag[0]),
		LimitPrice:          float64(field.LimitPrice),
		VolumeTotalOriginal: int(field.VolumeTotalOriginal),
		VolumeLeft:          int(field.VolumeTotalOriginal),
		ExchangeID:          Bytes2String(field.ExchangeID[:]),
		IsLocal:             true,
	})
	var o = of.(*OrderField)
	o.OrderStatus = OrderStatusCanceled
	if t.onErrRtnOrder != nil {
		t.onErrRtnOrder(o, &RspInfoField{ErrorID: int(info.ErrorID), ErrorMsg: Bytes2String(info.ErrorMsg[:])})
	}
}

// RspQryInvestorPosition 持仓
func (t *HFTrade) RspQryInvestorPosition(field *ctp.CThostFtdcInvestorPositionField, b bool) {
	// 复制接口中的数据
	var buf bytes.Buffer
	var p = new(ctp.CThostFtdcInvestorPositionField)
	gob.NewEncoder(&buf).Encode(field)
	gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(p)
	instrumentID := Bytes2String(p.InstrumentID[:])
	if len(instrumentID) > 0 { // 偶尔出现NULL的数据导致数据转换错误
		if _, ok := t.Instruments.Load(instrumentID); ok { // 解决交易所自主合成某些不可交易的套利合约的问题如 SPC y2005&p2001
			var key string
			if p.PosiDirection == ctp.THOST_FTDC_PD_Long {
				key = fmt.Sprintf("%s_long", Bytes2String(p.InstrumentID[:]))
			} else if p.PosiDirection == ctp.THOST_FTDC_PD_Short {
				key = fmt.Sprintf("%s_short", Bytes2String(p.InstrumentID[:]))
			} else {
				key = fmt.Sprintf("%s_net", Bytes2String(p.InstrumentID[:]))
			}
			// key := fmt.Sprintf("%s_%c", Bytes2String(p.InstrumentID[:]), PosiDirectionType(p.PosiDirection))
			ps, _ := t.posiDetail.LoadOrStore(key, make([]*ctp.CThostFtdcInvestorPositionField, 0))
			ps = append(ps.([]*ctp.CThostFtdcInvestorPositionField), p)
			t.posiDetail.Store(key, ps) // append后指针有变化,需重新赋值
		}
	}
	if b {
		t.posiDetail.Range(func(key, ps interface{}) bool {
			pFinal := PositionField{}
			for _, p := range ps.([]*ctp.CThostFtdcInvestorPositionField) {
				pFinal.InstrumentID = Bytes2String(p.InstrumentID[:])
				pFinal.PositionDirection = PosiDirectionType(p.PosiDirection)
				pFinal.HedgeFlag = HedgeFlagType(p.HedgeFlag)
				pFinal.ExchangeID = Bytes2String(p.ExchangeID[:])
				pFinal.PreSettlementPrice = float64(p.PreSettlementPrice)
				pFinal.SettlementPrice = float64(p.SettlementPrice)

				pFinal.Position += int(p.Position)
				pFinal.TodayPosition += int(p.TodayPosition)
				// pFinal.YdPosition += int(p.YdPosition) // 直接取值还需要减去当日平仓
				pFinal.YdPosition = pFinal.Position - pFinal.TodayPosition
				pFinal.LongFrozen += int(p.LongFrozen)
				pFinal.ShortFrozen += int(p.ShortFrozen)
				pFinal.LongFrozenAmount += float64(p.LongFrozenAmount)
				pFinal.ShortFrozenAmount += float64(p.ShortFrozenAmount)
				pFinal.OpenVolume += int(p.OpenVolume)
				pFinal.CloseVolume += int(p.CloseVolume)
				pFinal.OpenAmount += float64(p.OpenAmount)
				pFinal.CloseAmount += float64(p.CloseAmount)
				pFinal.PositionCost += float64(p.PositionCost)
				pFinal.PreMargin += float64(p.PreMargin)
				pFinal.UseMargin += float64(p.UseMargin)
				pFinal.FrozenMargin += float64(p.FrozenMargin)
				pFinal.FrozenCash += float64(p.FrozenCash)
				pFinal.FrozenCommission += float64(p.FrozenCommission)
				pFinal.CashIn += float64(p.CashIn)
				pFinal.Commission += float64(p.Commission)
				pFinal.CloseProfit += float64(p.CloseProfit)
				pFinal.PositionProfit += float64(p.PositionProfit)
				pFinal.OpenCost += float64(p.OpenCost)
				pFinal.ExchangeMargin += float64(p.ExchangeMargin)
				pFinal.CombPosition += int(p.CombPosition)
				pFinal.CombLongFrozen += int(p.CombLongFrozen)
				pFinal.CombShortFrozen += int(p.CombShortFrozen)
				pFinal.CloseProfitByDate += float64(p.CloseProfitByDate)
				pFinal.CloseProfitByTrade += float64(p.CloseProfitByTrade)
				pFinal.StrikeFrozen += int(p.StrikeFrozen)
				pFinal.StrikeFrozenAmount += float64(p.StrikeFrozenAmount)
				pFinal.AbandonFrozen += int(p.AbandonFrozen)
				pFinal.YdStrikeFrozen += int(p.YdStrikeFrozen)
				pFinal.PositionCostOffset += float64(p.PositionCostOffset)
			}
			t.Positions.Store(key, &pFinal)
			return true
		})
	}
}

// RspQryTradingAccount 权益
func (t *HFTrade) RspQryTradingAccount(field *ctp.CThostFtdcTradingAccountField) {
	//infoField := (* ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.Account.PreMortgage = float64(field.PreMortgage)
	t.Account.PreDeposit = float64(field.PreDeposit)
	t.Account.PreBalance = float64(field.PreBalance)
	t.Account.PreMargin = float64(field.PreMargin)
	t.Account.InterestBase = float64(field.InterestBase)
	t.Account.Interest = float64(field.Interest)
	t.Account.Deposit = float64(field.Deposit)
	t.Account.Withdraw = float64(field.Withdraw)
	t.Account.FrozenMargin = float64(field.FrozenMargin)
	t.Account.FrozenCash = float64(field.FrozenCash)
	t.Account.FrozenCommission = float64(field.FrozenCommission)
	t.Account.CurrMargin = float64(field.CurrMargin)
	t.Account.CashIn = float64(field.CashIn)
	t.Account.Commission = float64(field.Commission)
	t.Account.CloseProfit = float64(field.CloseProfit)
	t.Account.PositionProfit = float64(field.PositionProfit)
	t.Account.Balance = float64(field.Balance)
	t.Account.Available = float64(field.Available)
	t.Account.WithdrawQuota = float64(field.WithdrawQuota)
	t.Account.Reserve = float64(field.Reserve)
	t.Account.Credit = float64(field.Credit)
	t.Account.Mortgage = float64(field.Mortgage)
	t.Account.ExchangeMargin = float64(field.ExchangeMargin)
	t.Account.DeliveryMargin = float64(field.DeliveryMargin)
	t.Account.ExchangeDeliveryMargin = float64(field.ExchangeDeliveryMargin)
	t.Account.ReserveBalance = float64(field.ReserveBalance)
	t.Account.CurrencyID = Bytes2String(field.CurrencyID[:])
	t.Account.PreFundMortgageIn = float64(field.PreFundMortgageIn)
	t.Account.PreFundMortgageOut = float64(field.PreFundMortgageOut)
	t.Account.FundMortgageIn = float64(field.FundMortgageIn)
	t.Account.FundMortgageOut = float64(field.FundMortgageOut)
	t.Account.FundMortgageAvailable = float64(field.FundMortgageAvailable)
	t.Account.MortgageableFund = float64(field.MortgageableFund)
}

// RspQryInstrument 合约
func (t *HFTrade) RspQryInstrument(field *ctp.CThostFtdcInstrumentField, b bool) {
	if field != nil {
		t.Instruments.Store(Bytes2String(field.InstrumentID[:]), &InstrumentField{
			InstrumentID:              Bytes2String(field.InstrumentID[:]),
			ExchangeID:                Bytes2String(field.ExchangeID[:]),
			ProductID:                 Bytes2String(field.ProductID[:]),
			ProductClass:              ProductClassType(field.ProductClass),
			MaxMarketOrderVolume:      int(field.MaxMarketOrderVolume),
			MinMarketOrderVolume:      int(field.MinMarketOrderVolume),
			MaxLimitOrderVolume:       int(field.MaxLimitOrderVolume),
			MinLimitOrderVolume:       int(field.MinLimitOrderVolume),
			VolumeMultiple:            int(field.VolumeMultiple),
			PriceTick:                 float64(field.PriceTick),
			PositionType:              PositionTypeType(field.PositionType),
			UseMaxMarginSideAlgorithm: field.MaxMarginSideAlgorithm == '1',
			UnderlyingInstrID:         Bytes2String(field.UnderlyingInstrID[:]),
			StrikePrice:               float64(field.StrikePrice),
			OptionsType:               OptionsTypeType(field.OptionsType),
			UnderlyingMultiple:        float64(field.UnderlyingMultiple),
			CombinationType:           CombinationTypeType(field.CombinationType),
		})
	}
	if b && !t.IsLogin {
		go t.qry()
	}
}

// 循环查询持仓&资金
func (t *HFTrade) qry() {
	t.qryTicker = time.NewTicker(1100 * time.Millisecond)
	// 等待之前的Order响应完再发送登录通知
	ordCnt := t.cntOrder
	for range t.qryTicker.C {
		if ordCnt == t.cntOrder {
			break
		}
		ordCnt = t.cntOrder
	}

	faccount := ctp.CThostFtdcQryTradingAccountField{}
	copy(faccount.InvestorID[:], t.InvestorID)
	copy(faccount.BrokerID[:], t.BrokerID)
	fposition := ctp.CThostFtdcQryInvestorPositionField{}
	copy(fposition.InvestorID[:], t.InvestorID)
	copy(fposition.BrokerID[:], t.BrokerID)

	// 启动查询
	t.ReqQryTradingAccount(&faccount, t.getReqID())
	time.Sleep(1100 * time.Millisecond)
	bQryAccount := false
	for range t.qryTicker.C { // tick 每秒执行一次
		if bQryAccount {
			t.ReqQryTradingAccount(&faccount, t.getReqID())
		} else {
			t.posiDetail = sync.Map{} // 清空原始持仓数据
			t.ReqQryInvestorPosition(&fposition, t.getReqID())
		}
		bQryAccount = !bQryAccount
		if !t.IsLogin && !bQryAccount { // account position 都查一遍后再通知
			// 登录成功响应
			t.IsLogin = true
			t.waitGroup.Done() // 通知:登录响应可以发了
		}
	}
}

// RspSettlementInfoConfirm 确认结算
func (t *HFTrade) RspSettlementInfoConfirm() {
	if strings.Compare(t.Version, "v6.5.1") < 0 {
		t.ReqQryInstrument(&ctp.CThostFtdcQryInstrumentField{}, t.getReqID())
	} else {
		f := ctp.CThostFtdcQryClassifiedInstrumentField{
			TradingType: ctp.THOST_FTDC_TD_TRADE,
			ClassType:   ctp.THOST_FTDC_INS_ALL,
		}
		t.ReqQryClassifiedInstrument(&f, t.getReqID())
	}
}

// RspUserLogin 登录
func (t *HFTrade) RspUserLogin(loginField *ctp.CThostFtdcRspUserLoginField, infoField *ctp.CThostFtdcRspInfoField) {
	if infoField.ErrorID == 0 {
		t.SessionID = int(loginField.SessionID)
		t.TradingDay = Bytes2String(loginField.TradingDay[:])
		f := ctp.CThostFtdcSettlementInfoConfirmField{}
		copy(f.InvestorID[:], t.InvestorID)
		copy(f.AccountID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)

		// 用waitgroup控制登录消息发送信号
		if t.onRspUserLogin != nil {
			t.waitGroup.Add(1)
			go func(field *RspUserLoginField) {
				f := ctp.CThostFtdcSettlementInfoConfirmField{}
				copy(f.InvestorID[:], t.InvestorID)
				copy(f.AccountID[:], t.InvestorID)
				copy(f.BrokerID[:], t.BrokerID)
				t.ReqSettlementInfoConfirm(&f, t.getReqID())

				t.waitGroup.Wait()
				t.onRspUserLogin(field, &RspInfoField{ErrorID: 0, ErrorMsg: "成功"})
			}(&RspUserLoginField{
				TradingDay:  t.TradingDay,
				LoginTime:   Bytes2String(loginField.LoginTime[:]),
				BrokerID:    t.BrokerID,
				UserID:      t.InvestorID,
				FrontID:     int(loginField.FrontID),
				SessionID:   t.SessionID,
				MaxOrderRef: Bytes2String(loginField.MaxOrderRef[:]),
			})
		}
	} else {
		t.onRspUserLogin(&RspUserLoginField{}, &RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: Bytes2String(infoField.ErrorMsg[:])})
	}
}

// RspAuthenticate 认证
func (t *HFTrade) RspAuthenticate(info *ctp.CThostFtdcRspInfoField) {
	if info.ErrorID == 0 {
		f := ctp.CThostFtdcReqUserLoginField{}
		copy(f.UserID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.Password[:], t.passWord)
		copy(f.UserProductInfo[:], "@HF")
		t.ReqUserLogin(&f, t.getReqID())
	} else if t.onRspUserLogin != nil {
		infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
		t.onRspUserLogin(&RspUserLoginField{}, &RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: Bytes2String(infoField.ErrorMsg[:])})
	}
}

// FrontDisConnected 断开响应
func (t *HFTrade) FrontDisConnected(reason int) {
	if t.onFrontDisConnected != nil {
		t.onFrontDisConnected(reason)
	}
}

// FrontConnected 连接
func (t *HFTrade) FrontConnected() {
	if t.onFrontConnected != nil {
		t.onFrontConnected()
	}
}
