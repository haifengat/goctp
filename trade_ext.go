package goctp

import (
	"fmt"
	"time"
)

type TradeExt struct {
	*Trade
	Broker, UserID, InvestorID string
	id                         int
}

// NewTradeExt 接口实例
func NewTradeExt() *TradeExt {
	qExt := &TradeExt{}
	qExt.Trade = NewTrade()
	// 从1970年到现在的天数作为高端
	qExt.id = int(time.Since(time.Unix(0, 0)).Seconds()/24/3600) * 100000
	return qExt
}

func (t *TradeExt) getReqID() int {
	t.id++
	return t.id
}

// Subscribe 订阅流(私有,公有)
func (t *Trade) Subscribe(private, public THOST_TE_RESUME_TYPE) {
	t.SubscribePrivateTopic(private)
	t.SubscribePublicTopic(public)
}

// ReqConnect 连接前置
func (t *TradeExt) ReqConnect(front string) {
	t.RegisterFront(front)
	t.Init()
}

// ReqAuthenticateField 认证
func (t *TradeExt) ReqAuthenticateField(broker, user, appID, authCode string) {
	f := CThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], []byte(broker))
	copy(f.UserID[:], []byte(user))
	copy(f.AppID[:], []byte(appID))
	copy(f.AuthCode[:], []byte(authCode))
	t.Broker = broker
	t.UserID = user
	t.InvestorID = user
	t.ReqAuthenticate(&f, 1)
}

// ReqUserLogin 登录
func (t *TradeExt) ReqUserLogin(pwd string) {
	f := CThostFtdcReqUserLoginField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.Password[:], []byte(pwd))
	t.Trade.ReqUserLogin(&f, t.getReqID())
}

// ReqSettlementInfoConfirm 确认结算结果
func (t *TradeExt) ReqSettlementInfoConfirm() {
	f := CThostFtdcSettlementInfoConfirmField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.InvestorID[:], []byte(t.UserID))
	t.Trade.ReqSettlementInfoConfirm(&f, t.getReqID())
}

// Release 消毁接口
// 6.6.1说明中提到,会导致程序崩溃
// 若柜台前置开着,而交易核心服务关闭报错:
// DesignError:pthread_mutex_unlock in line 116 of file ../../source/event/Mutex.h
func (t *TradeExt) Release() {
	t.RegisterSpi(nil)
	t.Trade.Release()
	t.Trade = nil
}

// ReqQryInvestor 查 userid 管理的用户
func (t *TradeExt) ReqQryInvestor() {
	f := CThostFtdcQryInvestorField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.InvestorID[:], []byte(t.UserID))
	t.Trade.ReqQryInvestor(&f, t.getReqID())
}

// ReqQryInstrument 查合约
func (t *TradeExt) ReqQryInstrument() {
	t.Trade.ReqQryInstrument(&CThostFtdcQryInstrumentField{}, t.getReqID())
}

// ReqQryClassifiedInstrument 查可交易合约
func (t *TradeExt) ReqQryClassifiedInstrument() {
	t.Trade.ReqQryClassifiedInstrument(&CThostFtdcQryClassifiedInstrumentField{}, t.getReqID())
}

// ReqQryOrder 查委托
func (t *TradeExt) ReqQryOrder() {
	t.Trade.ReqQryOrder(&CThostFtdcQryOrderField{}, t.getReqID())
}

// ReqQryTrade 查成交
func (t *TradeExt) ReqQryTrade() {
	t.Trade.ReqQryTrade(&CThostFtdcQryTradeField{}, t.getReqID())
}

// ReqQryPosition 查持仓
func (t *TradeExt) ReqQryPosition() {
	t.Trade.ReqQryInvestorPosition(&CThostFtdcQryInvestorPositionField{}, t.getReqID())
}

// ReqQryPositionDetail 查持仓明细
func (t *TradeExt) ReqQryPositionDetail() {
	t.Trade.ReqQryInvestorPositionDetail(&CThostFtdcQryInvestorPositionDetailField{}, t.getReqID())
}

// ReqQryTradingAccount 查资金帐号
func (t *TradeExt) ReqQryTradingAccount() {
	t.Trade.ReqQryTradingAccount(&CThostFtdcQryTradingAccountField{}, t.getReqID())
}

// ReqOrderInsert 限价委托
func (t *TradeExt) ReqOrderInsert(instrument, exchange string, buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, price float64, volume int, investor string) {
	f := CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.UserID[:], t.UserID)
	copy(f.InvestorID[:], []byte(investor))
	copy(f.AccountID[:], []byte(investor))
	f.IsAutoSuspend = TThostFtdcBoolType(0)
	f.IsSwapOrder = TThostFtdcBoolType(0)
	f.ForceCloseReason = THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	f.RequestID = TThostFtdcRequestIDType(id)
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	copy(f.ExchangeID[:], []byte(exchange))
	f.Direction = buySell
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(THOST_FTDC_HF_Speculation)
	// 不同类型的Order
	f.OrderPriceType = THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = THOST_FTDC_TC_GFD
	f.VolumeCondition = THOST_FTDC_VC_AV
	f.ContingentCondition = THOST_FTDC_CC_Immediately
	f.LimitPrice = TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = TThostFtdcVolumeType(volume)

	t.Trade.ReqOrderInsert(&f, id)
}

// ReqUpdateInvestorPwd 修改投资者密码
func (t *TradeExt) ReqUpdateInvestorPwd(oldPwd, newPwd string) {
	f := CThostFtdcTradingAccountPasswordUpdateField{}
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.OldPassword[:], []byte(oldPwd))
	copy(f.NewPassword[:], []byte(newPwd))
	t.Trade.ReqTradingAccountPasswordUpdate(&f, t.getReqID())
}

// ReqUpdateUserPwd 修改用户密码
func (t *TradeExt) ReqUpdateUserPwd(oldPwd, newPwd string) {
	f := CThostFtdcUserPasswordUpdateField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.OldPassword[:], []byte(oldPwd))
	copy(f.NewPassword[:], []byte(newPwd))
	t.Trade.ReqUserPasswordUpdate(&f, t.getReqID())
}

// ReqOrderAction 撤单
func (t *TradeExt) ReqOrderAction(order struct {
	ExchangeID   string // 交易所
	InstrumentID string // 合约
	InvestUnitID string // 投资单元代码(中介的子帐户)
	OrderSysID   string // 报单编号(交易所)
	OrderRef     string // 报单引用(客户端)
	SessionID    int
	FrontID      int
}) {
	f := CThostFtdcInputOrderActionField{}
	f.ActionFlag = THOST_FTDC_AF_Delete
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.InvestorID[:], []byte(t.InvestorID))
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.ExchangeID[:], []byte(order.ExchangeID))
	copy(f.InstrumentID[:], []byte(order.InstrumentID))
	copy(f.InvestUnitID[:], []byte(order.InvestUnitID))
	copy(f.OrderSysID[:], []byte(order.OrderSysID))
	copy(f.OrderRef[:], []byte(order.OrderRef))
	f.SessionID = TThostFtdcSessionIDType(order.SessionID)
	f.FrontID = TThostFtdcFrontIDType(order.FrontID)
	i := t.getReqID()
	f.OrderActionRef = TThostFtdcOrderActionRefType(i)
	f.RequestID = TThostFtdcRequestIDType(i)
	t.Trade.ReqOrderAction(&f, i)
}

// --------- 银转相关 -------------

// ReqQryAccountregister 查询银期签约关系
func (t *TradeExt) ReqQryAccountregister() {
	f := CThostFtdcQryAccountregisterField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.AccountID[:], []byte(t.InvestorID))
	t.Trade.ReqQryAccountregister(&f, t.getReqID())
}

// ReqQryTransferBank 查询转帐银行
func (t *TradeExt) ReqQryTransferBank() {
	f := CThostFtdcQryTransferBankField{}
	t.Trade.ReqQryTransferBank(&f, t.getReqID())
}

// ReqFromBankToFutureByFuture 入金
func (t *TradeExt) ReqFromBankToFutureByFuture(regInfo CThostFtdcAccountregisterField, amount float64) {
	f := CThostFtdcReqTransferField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.UserID[:], []byte(t.UserID))
	f.LastFragment = THOST_FTDC_LF_Yes          // 最后分片:是
	f.SecuPwdFlag = THOST_FTDC_BPWDF_BlankCheck // 资金密码核对标志

	f.CustType = regInfo.CustType
	f.IdCardType = regInfo.IdCardType
	copy(f.CurrencyID[:], regInfo.CurrencyID[:])
	copy(f.IdentifiedCardNo[:], regInfo.IdentifiedCardNo[:])
	copy(f.BankAccount[:], regInfo.BankAccount[:]) // 卡号
	copy(f.BankID[:], regInfo.BankID[:])
	copy(f.BankBranchID[:], regInfo.BankBranchID[:])

	copy(f.TradeCode[:], []byte(THOST_FTDC_VTC_FutureFutureToBank))
	f.TradeAmount = TThostFtdcTradeAmountType(amount)

	t.Trade.ReqFromBankToFutureByFuture(&f, t.getReqID())
}

// ReqFromFutureToBankByFuture 出金
func (t *TradeExt) ReqFromFutureToBankByFuture(regInfo CThostFtdcAccountregisterField, amount float64) {
	f := CThostFtdcReqTransferField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.UserID[:], []byte(t.UserID))
	f.LastFragment = THOST_FTDC_LF_Yes          // 最后分片:是
	f.SecuPwdFlag = THOST_FTDC_BPWDF_BlankCheck // 资金密码核对标志

	f.CustType = regInfo.CustType
	f.IdCardType = regInfo.IdCardType
	copy(f.CurrencyID[:], regInfo.CurrencyID[:])
	copy(f.IdentifiedCardNo[:], regInfo.IdentifiedCardNo[:])
	copy(f.BankAccount[:], regInfo.BankAccount[:]) // 卡号
	copy(f.BankID[:], regInfo.BankID[:])
	copy(f.BankBranchID[:], regInfo.BankBranchID[:])

	copy(f.TradeCode[:], []byte(THOST_FTDC_VTC_FutureBankToFuture))
	f.TradeAmount = TThostFtdcTradeAmountType(amount)

	t.Trade.ReqFromFutureToBankByFuture(&f, t.getReqID())
}

// --------------- 行情

// ReqQryDepthMarketData 查最后一笔行情(会同时返回对应的期权合约的行情)
func (t *TradeExt) ReqQryDepthMarketData(exchange, instrument string) {
	f := CThostFtdcQryDepthMarketDataField{}
	copy(f.ExchangeID[:], []byte(exchange))
	copy(f.InstrumentID[:], instrument)
	t.Trade.ReqQryDepthMarketData(&f, t.getReqID())
}
