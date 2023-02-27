package trade

import "goctp/def"

type TradeExt struct {
	*Trade
	Broker, UserID string
	id             int
}

// NewTradeExt 接口实例
func NewTradeExt() *TradeExt {
	qExt := &TradeExt{}
	qExt.Trade = NewTrade()
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
