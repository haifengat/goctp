package trade

import (
	"fmt"
	"goctp/def"
	"time"
)

type TradeExt struct {
	*Trade
	Broker, UserID string
	id             int
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
func (t *Trade) Subscribe(private, public def.THOST_TE_RESUME_TYPE) {
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
	f := def.CThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], []byte(broker))
	copy(f.UserID[:], []byte(user))
	copy(f.AppID[:], []byte(appID))
	copy(f.AuthCode[:], []byte(authCode))
	t.Broker = broker
	t.UserID = user
	t.ReqAuthenticate(&f, 1)
}

// ReqUserLogin 登录
func (t *TradeExt) ReqUserLogin(pwd string) {
	f := def.CThostFtdcReqUserLoginField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.UserID[:], []byte(t.UserID))
	copy(f.Password[:], []byte(pwd))
	t.Trade.ReqUserLogin(&f, t.getReqID())
}

// ReqSettlementInfoConfirm 确认结算结果
func (t *TradeExt) ReqSettlementInfoConfirm() {
	f := def.CThostFtdcSettlementInfoConfirmField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.InvestorID[:], []byte(t.UserID))
	t.Trade.ReqSettlementInfoConfirm(&f, t.getReqID())
}

// Release 消毁接口
func (t *TradeExt) Release() {
	t.RegisterSpi(nil)
	t.Trade.Release()
	t.Trade = nil
}

// ReqQryInvestor 查 userid 管理的用户
func (t *TradeExt) ReqQryInvestor() {
	f := def.CThostFtdcQryInvestorField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.InvestorID[:], []byte(t.UserID))
	t.Trade.ReqQryInvestor(&f, t.getReqID())
}

// ReqQryInstrument 查合约
func (t *TradeExt) ReqQryInstrument() {
	t.Trade.ReqQryInstrument(&def.CThostFtdcQryInstrumentField{}, t.getReqID())
}

// ReqQryClassifiedInstrument 查可交易合约
func (t *TradeExt) ReqQryClassifiedInstrument() {
	t.Trade.ReqQryClassifiedInstrument(&def.CThostFtdcQryClassifiedInstrumentField{}, t.getReqID())
}

// ReqQryOrder 查委托
func (t *TradeExt) ReqQryOrder() {
	t.Trade.ReqQryOrder(&def.CThostFtdcQryOrderField{}, t.getReqID())
}

// ReqQryTrade 查成交
func (t *TradeExt) ReqQryTrade() {
	t.Trade.ReqQryTrade(&def.CThostFtdcQryTradeField{}, t.getReqID())
}

// ReqQryPosition 查持仓
func (t *TradeExt) ReqQryPosition() {
	t.Trade.ReqQryInvestorPosition(&def.CThostFtdcQryInvestorPositionField{}, t.getReqID())
}

// ReqQryPositionDetail 查持仓明细
func (t *TradeExt) ReqQryPositionDetail() {
	t.Trade.ReqQryInvestorPositionDetail(&def.CThostFtdcQryInvestorPositionDetailField{}, t.getReqID())
}

// ReqQryDeposit 查权益
func (t *TradeExt) ReqQryAccount() {
	t.Trade.ReqQryTradingAccount(&def.CThostFtdcQryTradingAccountField{}, t.getReqID())
}

// ReqOrderInsert 限价委托
func (t *TradeExt) ReqOrderInsert(instrument, exchange string, buySell def.TThostFtdcDirectionType, openClose def.TThostFtdcOffsetFlagType, price float64, volume int, investor string) {
	f := def.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], []byte(t.Broker))
	copy(f.UserID[:], t.UserID)
	copy(f.InvestorID[:], []byte(investor))
	copy(f.AccountID[:], []byte(investor))
	f.IsAutoSuspend = def.TThostFtdcBoolType(0)
	f.IsSwapOrder = def.TThostFtdcBoolType(0)
	f.ForceCloseReason = def.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	f.RequestID = def.TThostFtdcRequestIDType(id)
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	copy(f.ExchangeID[:], []byte(exchange))
	f.Direction = buySell
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(def.THOST_FTDC_HF_Speculation)
	// 不同类型的Order
	f.OrderPriceType = def.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = def.THOST_FTDC_TC_GFD
	f.VolumeCondition = def.THOST_FTDC_VC_AV
	f.ContingentCondition = def.THOST_FTDC_CC_Immediately
	f.LimitPrice = def.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = def.TThostFtdcVolumeType(volume)

	t.Trade.ReqOrderInsert(&f, id)
}
