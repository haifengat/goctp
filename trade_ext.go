package goctp

import (
	"fmt"
)

type TradeExt struct {
	*Trade
	Broker, UserID, InvestorID, pwd string
	id                              int
}

// NewTradeExt 接口实例
func NewTradeExt() *TradeExt {
	ext := &TradeExt{}
	ext.Trade = NewTrade()
	return ext
}

func (t *TradeExt) getReqID() int {
	t.id++
	return t.id
}

// ReqAuthenticate 认证
func (t *TradeExt) ReqAuthenticate(broker, user, appID, authCode string) {
	f := CThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], []byte(broker))
	copy(f.UserID[:], []byte(user))
	copy(f.AppID[:], []byte(appID))
	copy(f.AuthCode[:], []byte(authCode))
	t.Broker = broker
	t.UserID = user
	t.InvestorID = user
	t.Trade.ReqAuthenticate(&f, 1)
}

// ReqUserLogin 登录
func (t *TradeExt) ReqUserLogin(pwd string) {
	t.pwd = pwd
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
	// copy(f.InvestorID[:], []byte(t.UserID))
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
func (t *TradeExt) ReqQryInvestor() int {
	f := CThostFtdcQryInvestorField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	// copy(f.InvestorID[:], []byte(t.UserID))
	return t.Trade.ReqQryInvestor(&f, t.getReqID())
}

// ReqQryInstrument 查合约
func (t *TradeExt) ReqQryInstrument() int {
	return t.Trade.ReqQryInstrument(&CThostFtdcQryInstrumentField{}, t.getReqID())
}

// ReqQryClassifiedInstrument 查可交易合约
func (t *TradeExt) ReqQryClassifiedInstrument() int {
	return t.Trade.ReqQryClassifiedInstrument(&CThostFtdcQryClassifiedInstrumentField{
		TradingType: THOST_FTDC_TD_TRADE,
		ClassType:   THOST_FTDC_INS_ALL}, t.getReqID())
}

// ReqQryOrder 查委托
func (t *TradeExt) ReqQryOrder() int {
	return t.Trade.ReqQryOrder(&CThostFtdcQryOrderField{}, t.getReqID())
}

// ReqQryTrade 查成交
func (t *TradeExt) ReqQryTrade() int {
	return t.Trade.ReqQryTrade(&CThostFtdcQryTradeField{}, t.getReqID())
}

// ReqQryInvestorPosition 查持仓
func (t *TradeExt) ReqQryInvestorPosition() int {
	return t.Trade.ReqQryInvestorPosition(&CThostFtdcQryInvestorPositionField{}, t.getReqID())
}

// ReqQryInvestorPositionDetail 查持仓明细
func (t *TradeExt) ReqQryInvestorPositionDetail() int {
	return t.Trade.ReqQryInvestorPositionDetail(&CThostFtdcQryInvestorPositionDetailField{}, t.getReqID())
}

// ReqQryTradingAccount 查资金帐号
func (t *TradeExt) ReqQryTradingAccount() int {
	return t.Trade.ReqQryTradingAccount(&CThostFtdcQryTradingAccountField{}, t.getReqID())
}

// ReqOrderInsert 委托
func (t *TradeExt) ReqOrderInsert(buySell TThostFtdcDirectionType, openClose TThostFtdcOffsetFlagType, instrument, exchange string, price float64, volume int, investor string, priceType TThostFtdcOrderPriceTypeType, timeType TThostFtdcTimeConditionType, volumeType TThostFtdcVolumeConditionType, contingentType TThostFtdcContingentConditionType) int {
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
	f.LimitPrice = TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = TThostFtdcVolumeType(volume)
	// 不同类型的Order
	f.OrderPriceType = priceType
	f.TimeCondition = timeType
	f.VolumeCondition = volumeType
	f.ContingentCondition = contingentType

	return t.Trade.ReqOrderInsert(&f, id)
}

// ReqTradingAccountPasswordUpdate 修改投资者密码
func (t *TradeExt) ReqTradingAccountPasswordUpdate(oldPwd, newPwd string) int {
	f := CThostFtdcTradingAccountPasswordUpdateField{}
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.OldPassword[:], []byte(oldPwd))
	copy(f.NewPassword[:], []byte(newPwd))
	return t.Trade.ReqTradingAccountPasswordUpdate(&f, t.getReqID())
}

// ReqUserPasswordUpdate 修改用户密码
func (t *TradeExt) ReqUserPasswordUpdate(oldPwd, newPwd string) int {
	f := CThostFtdcUserPasswordUpdateField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.OldPassword[:], []byte(oldPwd))
	copy(f.NewPassword[:], []byte(newPwd))
	return t.Trade.ReqUserPasswordUpdate(&f, t.getReqID())
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
}) int {
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
	return t.Trade.ReqOrderAction(&f, i)
}

// --------- 银转相关 -------------

// ReqQryAccountregister 查询银期签约关系
func (t *TradeExt) ReqQryAccountregister() int {
	f := CThostFtdcQryAccountregisterField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	// copy(f.AccountID[:], []byte(t.InvestorID))
	return t.Trade.ReqQryAccountregister(&f, t.getReqID())
}

// ReqQryTransferBank 查询转帐银行
func (t *TradeExt) ReqQryTransferBank() int {
	f := CThostFtdcQryTransferBankField{}
	return t.Trade.ReqQryTransferBank(&f, t.getReqID())
}

// ReqFromBankToFutureByFuture 入金
func (t *TradeExt) ReqFromBankToFutureByFuture(regInfo CThostFtdcAccountregisterField, bankPwd, accountPwd string, amount float64) int {
	f := CThostFtdcReqTransferField{}
	copy(f.BrokerID[:], regInfo.BrokerID[:])
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.Password[:], []byte(accountPwd))

	// f.InstallID = 1
	// f.FutureSerial=0
	f.LastFragment = THOST_FTDC_LF_Yes          // 最后分片:是
	f.SecuPwdFlag = THOST_FTDC_BPWDF_BlankCheck // 资金密码核对标志
	f.CustType = regInfo.CustType
	f.IdCardType = regInfo.IdCardType
	f.VerifyCertNoFlag = THOST_FTDC_YNI_No
	f.RequestID = TThostFtdcRequestIDType(t.getReqID())
	f.TID = regInfo.TID
	f.BankAccType = regInfo.BankAccType

	copy(f.AccountID[:], regInfo.AccountID[:])
	copy(f.CurrencyID[:], regInfo.CurrencyID[:])
	copy(f.IdentifiedCardNo[:], regInfo.IdentifiedCardNo[:])
	copy(f.BankAccount[:], regInfo.BankAccount[:]) // 卡号
	copy(f.BankPassWord[:], []byte(bankPwd))
	copy(f.BankID[:], regInfo.BankID[:])
	copy(f.BankBranchID[:], regInfo.BankBranchID[:])
	copy(f.BrokerIDByBank[:], regInfo.BrokerBranchID[:])

	copy(f.TradeCode[:], []byte(THOST_FTDC_VTC_FutureBankToFuture))
	f.TradeAmount = TThostFtdcTradeAmountType(amount)

	return t.Trade.ReqFromBankToFutureByFuture(&f, int(f.RequestID))
}

// ReqFromFutureToBankByFuture 出金
// accountPwd 资金密码
func (t *TradeExt) ReqFromFutureToBankByFuture(regInfo CThostFtdcAccountregisterField, accountPwd string, amount float64) int {
	f := CThostFtdcReqTransferField{}
	copy(f.BrokerID[:], regInfo.BrokerID[:])
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.Password[:], []byte(accountPwd))

	f.LastFragment = THOST_FTDC_LF_Yes          // 最后分片:是
	f.SecuPwdFlag = THOST_FTDC_BPWDF_BlankCheck // 资金密码核对标志
	f.CustType = regInfo.CustType
	f.IdCardType = regInfo.IdCardType
	f.VerifyCertNoFlag = THOST_FTDC_YNI_No
	f.RequestID = TThostFtdcRequestIDType(t.getReqID())
	f.TID = regInfo.TID
	f.BankAccType = regInfo.BankAccType

	copy(f.AccountID[:], regInfo.AccountID[:])
	copy(f.CurrencyID[:], regInfo.CurrencyID[:])
	copy(f.IdentifiedCardNo[:], regInfo.IdentifiedCardNo[:])
	copy(f.BankAccount[:], regInfo.BankAccount[:]) // 卡号
	// copy(f.BankPassWord[:], []byte(bankPwd))
	copy(f.BankID[:], regInfo.BankID[:])
	copy(f.BankBranchID[:], regInfo.BankBranchID[:])
	copy(f.BrokerIDByBank[:], regInfo.BrokerBranchID[:])

	copy(f.TradeCode[:], []byte(THOST_FTDC_VTC_FutureFutureToBank))
	f.TradeAmount = TThostFtdcTradeAmountType(amount)

	return t.Trade.ReqFromFutureToBankByFuture(&f, t.getReqID())
}

// --------------- 行情

// ReqQryDepthMarketData 查最后一笔行情(会同时返回对应的期权合约的行情)
func (t *TradeExt) ReqQryDepthMarketData(exchange, instrument string) int {
	f := CThostFtdcQryDepthMarketDataField{}
	copy(f.ExchangeID[:], []byte(exchange))
	copy(f.InstrumentID[:], instrument)
	return t.Trade.ReqQryDepthMarketData(&f, t.getReqID())
}
