package go_ctp_swig_lnx

import (
	"fmt"
	. "github.com/mayiweb/goctp" // 发布时使用，编码时可注释掉
	. "hf_go_ctp"
	"strings"
	"sync"
	"time"
)

type Trade struct {
	api CThostFtdcTraderApi
	// 帐号
	InvestorID string
	// 经纪商
	BrokerID string
	// 交易日
	TradingDay string
	// 密码
	passWord string
	// 判断是否自己的委托用
	sessionID int
	// 合约列表
	Instruments map[string]InstrumentField
	// 合约状态
	InstrumentStatuss map[string]InstrumentStatusType
	// 持仓列表
	Positions map[string]*PositionField
	// 委托
	Orders map[string]*OrderField
	// 成交
	Trades map[string]*TradeField
	// 帐户权益
	Account *AccountField
	// 登录成功
	IsLogin bool
	// 循环查询
	qryTicker        *time.Ticker
	onFrontConnected OnFrontConnectedType
	onRspUserLogin   OnRspUserLoginType
	onRtnOrder       OnRtnOrderType
	onRtnCancel      OnRtnOrderType
	onErrRtnOrder    OnRtnErrOrderType
	onErrAction      OnRtnOrderType
	onRtnTrade       OnRtnTradeType
	// chan 登录信号
	waitGroup sync.WaitGroup
	// orderSysID 对应的 Order
	sysID4Order map[string]*OrderField
	reqID       int
}

func (t *Trade) getReqID() int {
	t.reqID++
	return t.reqID
}

func NewTrade() *Trade {
	t := new(Trade)
	// 初始化变量
	t.waitGroup = sync.WaitGroup{}
	t.IsLogin = false
	t.Instruments = make(map[string]InstrumentField)
	t.Positions = make(map[string]*PositionField)
	t.Orders = make(map[string]*OrderField)
	t.Trades = make(map[string]*TradeField)
	t.Account = new(AccountField)
	t.InstrumentStatuss = make(map[string]InstrumentStatusType)
	t.sysID4Order = make(map[string]*OrderField)

	t.api = CThostFtdcTraderApiCreateFtdcTraderApi("./log/")
	t.api.RegisterSpi(NewDirectorCThostFtdcTraderSpi(t))
	return t
}

// ********************** 主调函数 ************************

// 接口销毁处理
func (t *Trade) Release() {
	t.IsLogin = false
	t.api.Release()
}

// 连接
func (t *Trade) ReqConnect(addr string) {
	t.api.RegisterFront(addr)
	t.api.SubscribePrivateTopic(THOST_TERT_RESTART)
	t.api.SubscribePublicTopic(THOST_TERT_RESTART)
	t.api.Init()
}

// 登录
func (t *Trade) ReqLogin(investor, pwd, broker, appID, authCode string) {
	t.InvestorID = investor
	t.passWord = pwd
	t.BrokerID = broker
	f := NewCThostFtdcReqAuthenticateField()
	f.SetBrokerID(broker)
	f.SetUserID(investor)
	f.SetAppID(appID)
	f.SetAuthCode(authCode)
	t.api.ReqAuthenticate(f, t.getReqID())
}

// 限价委托
// 返回委托的ID
func (t *Trade) ReqOrderInsert(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := NewCThostFtdcInputOrderField()
	f.SetBrokerID(t.BrokerID)
	f.SetExchangeID(t.Instruments[instrument].ExchangeID)
	f.SetUserID(t.InvestorID)
	f.SetInvestorID(t.InvestorID)
	f.SetAccountID(t.InvestorID)
	f.SetIsAutoSuspend(0)
	f.SetIsSwapOrder(0)
	f.SetForceCloseReason(THOST_FTDC_FCC_NotForceClose)
	f.SetContingentCondition(THOST_FTDC_CC_Immediately)
	// 参数赋值
	f.SetOrderRef(fmt.Sprintf("%012d", t.getReqID()))
	f.SetInstrumentID(instrument)
	f.SetDirection(byte(buySell))
	f.SetCombOffsetFlag(string(openClose))
	f.SetCombHedgeFlag(string(HedgeFlagSpeculation))
	f.SetVolumeTotalOriginal(volume)
	// 不同类型的Order
	f.SetOrderPriceType(THOST_FTDC_OPT_LimitPrice)
	f.SetTimeCondition(THOST_FTDC_TC_IOC)
	f.SetVolumeCondition(THOST_FTDC_VC_AV)
	f.SetLimitPrice(price)
	t.api.ReqOrderInsert(f, t.getReqID())
	return fmt.Sprintf("%d_%s", t.sessionID, f.GetOrderRef())
}

// 市价委托
func (t *Trade) ReqOrderInsertMarket(instrument string, buySell DirectionType, openClose OffsetFlagType, volume int) string {
	f := NewCThostFtdcInputOrderField()
	f.SetBrokerID(t.BrokerID)
	f.SetExchangeID(t.Instruments[instrument].ExchangeID)
	f.SetUserID(t.InvestorID)
	f.SetInvestorID(t.InvestorID)
	f.SetAccountID(t.InvestorID)
	f.SetIsAutoSuspend(0)
	f.SetIsSwapOrder(0)
	f.SetForceCloseReason(THOST_FTDC_FCC_NotForceClose)
	f.SetContingentCondition(THOST_FTDC_CC_Immediately)
	// 参数赋值
	f.SetOrderRef(fmt.Sprintf("%012d", t.getReqID()))
	f.SetInstrumentID(instrument)
	f.SetDirection(byte(buySell))
	f.SetCombOffsetFlag(string(openClose))
	f.SetCombHedgeFlag(string(HedgeFlagSpeculation))
	f.SetVolumeTotalOriginal(volume)
	// 不同类型的Order
	f.SetOrderPriceType(THOST_FTDC_OPT_AnyPrice)
	f.SetTimeCondition(THOST_FTDC_TC_IOC)
	f.SetVolumeCondition(THOST_FTDC_VC_AV)
	f.SetLimitPrice(0)
	t.api.ReqOrderInsert(f, t.getReqID())
	return fmt.Sprintf("%d_%s", t.sessionID, f.GetOrderRef())
}

// FOK委托[部成撤单]
func (t *Trade) ReqOrderInsertFOK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := NewCThostFtdcInputOrderField()
	f.SetBrokerID(t.BrokerID)
	f.SetExchangeID(t.Instruments[instrument].ExchangeID)
	f.SetUserID(t.InvestorID)
	f.SetInvestorID(t.InvestorID)
	f.SetAccountID(t.InvestorID)
	f.SetIsAutoSuspend(0)
	f.SetIsSwapOrder(0)
	f.SetForceCloseReason(THOST_FTDC_FCC_NotForceClose)
	f.SetContingentCondition(THOST_FTDC_CC_Immediately)
	// 参数赋值
	f.SetOrderRef(fmt.Sprintf("%012d", t.getReqID()))
	f.SetInstrumentID(instrument)
	f.SetDirection(byte(buySell))
	f.SetCombOffsetFlag(string(openClose))
	f.SetCombHedgeFlag(string(HedgeFlagSpeculation))
	f.SetVolumeTotalOriginal(volume)
	// 不同类型的Order
	f.SetOrderPriceType(THOST_FTDC_OPT_LimitPrice)
	f.SetTimeCondition(THOST_FTDC_TC_GFD)
	f.SetVolumeCondition(THOST_FTDC_VC_CV)
	f.SetLimitPrice(price)
	t.api.ReqOrderInsert(f, t.getReqID())
	return fmt.Sprintf("%d_%s", t.sessionID, f.GetOrderRef())
}

// FAK委托[全成or撤单]
func (t *Trade) ReqOrderInsertFAK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := NewCThostFtdcInputOrderField()
	f.SetBrokerID(t.BrokerID)
	f.SetExchangeID(t.Instruments[instrument].ExchangeID)
	f.SetUserID(t.InvestorID)
	f.SetInvestorID(t.InvestorID)
	f.SetAccountID(t.InvestorID)
	f.SetIsAutoSuspend(0)
	f.SetIsSwapOrder(0)
	f.SetForceCloseReason(THOST_FTDC_FCC_NotForceClose)
	f.SetContingentCondition(THOST_FTDC_CC_Immediately)
	// 参数赋值
	f.SetOrderRef(fmt.Sprintf("%012d", t.getReqID()))
	f.SetInstrumentID(instrument)
	f.SetDirection(byte(buySell))
	f.SetCombOffsetFlag(string(openClose))
	f.SetCombHedgeFlag(string(HedgeFlagSpeculation))
	f.SetVolumeTotalOriginal(volume)
	// 不同类型的Order
	f.SetOrderPriceType(THOST_FTDC_OPT_LimitPrice)
	f.SetTimeCondition(THOST_FTDC_TC_IOC)
	f.SetVolumeCondition(THOST_FTDC_VC_AV)
	f.SetLimitPrice(price)
	t.api.ReqOrderInsert(f, t.getReqID())
	return fmt.Sprintf("%d_%s", t.sessionID, f.GetOrderRef())
}

// ********************** 注册客户响应 ************************

// 注册连接响应
func (t *Trade) RegOnFrontConnected(on OnFrontConnectedType) {
	t.onFrontConnected = on
}

// 注册登陆响应
func (t *Trade) RegOnRspUserLogin(on OnRspUserLoginType) {
	t.onRspUserLogin = on
}

// 注册委托响应
func (t *Trade) RegOnRtnOrder(on OnRtnOrderType) {
	t.onRtnOrder = on
}

// 注册委托响应
func (t *Trade) RegOnErrRtnOrder(on OnRtnErrOrderType) {
	t.onErrRtnOrder = on
}

// 注册撤单响应
func (t *Trade) RegOnErrAction(on OnRtnOrderType) {
	t.onErrAction = on
}

// 注册撤单响应
func (t *Trade) RegOnRtnCancel(on OnRtnOrderType) {
	t.onRtnCancel = on
}

// 注册成交响应
func (t *Trade) RegOnRtnTrade(on OnRtnTradeType) {
	t.onRtnTrade = on
}

// ********************** 底层接口响应处理 **********************************

// 合约状态响应
func (t *Trade) OnRtnInstrumentStatus(statusField CThostFtdcInstrumentStatusField) {
	t.InstrumentStatuss[statusField.GetInstrumentID()] = InstrumentStatusType(statusField.GetInstrumentStatus())
}

// 成交响应
func (t *Trade) OnRtnTrade(tradeField CThostFtdcTradeField) {
	key := fmt.Sprintf("%s_%c", tradeField.GetTradeID(), tradeField.GetDirection())
	f, ok := t.Trades[key]
	if !ok {
		f = new(TradeField)
		t.Trades[key] = f
		f.Direction = DirectionType(tradeField.GetDirection())
		f.HedgeFlag = HedgeFlagType(tradeField.GetHedgeFlag())
		f.InstrumentID = tradeField.GetInstrumentID()
		f.ExchangeID = tradeField.GetExchangeID()
		f.TradingDay = tradeField.GetTradingDay()
		f.Volume = tradeField.GetVolume()
		f.OffsetFlag = OffsetFlagType(tradeField.GetOffsetFlag())
		f.OrderSysID = tradeField.GetOrderSysID()
		f.Price = tradeField.GetPrice()
		f.TradeDate = tradeField.GetTradeDate()
		f.TradeTime = tradeField.GetTradeTime()
		f.TradeID = key
	}
	// 更新持仓
	if f.OffsetFlag == OffsetFlagOpen {
		var key string
		if f.Direction == DirectionBuy {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, PosiDirectionLong, f.HedgeFlag)
		} else {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, PosiDirectionShort, f.HedgeFlag)
		}
		p, ok := t.Positions[key]
		if !ok {
			p = &PositionField{
				InstrumentID:      f.InstrumentID,
				PositionDirection: PosiDirectionLong,
				HedgeFlag:         f.HedgeFlag,
				ExchangeID:        f.ExchangeID,
			}
			t.Positions[key] = p
		}
		p.OpenVolume += f.Volume
		p.OpenAmount += f.Price * float64(f.Volume)
		p.OpenCost += f.Price * float64(f.Volume) * float64(t.Instruments[f.InstrumentID].VolumeMultiple)
		p.Position += f.Volume
		p.TodayPosition += f.Volume
	} else {
		var key string
		if f.Direction == DirectionBuy {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, PosiDirectionShort, f.HedgeFlag)
		} else {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, PosiDirectionLong, f.HedgeFlag)
		}
		p, _ := t.Positions[key]
		p.OpenVolume -= f.Volume
		p.OpenAmount -= f.Price * float64(f.Volume)
		p.OpenCost -= f.Price * float64(f.Volume) * float64(t.Instruments[f.InstrumentID].VolumeMultiple)
		p.Position -= f.Volume
		if f.OffsetFlag == OffsetFlagCloseToday {
			p.TodayPosition -= f.Volume
		} else {
			p.YdPosition -= f.Volume
		}
	}
	// 处理对应的Order
	o, ok := t.sysID4Order[f.OrderSysID]
	if ok {
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
		if t.onRtnOrder != nil {
			t.onRtnOrder(o)
		}
	}
	// 客户端响应
	if t.onRtnTrade != nil {
		t.onRtnTrade(f)
	}
}

// 委托响应
func (t *Trade) OnRtnOrder(orderField CThostFtdcOrderField) {
	key := fmt.Sprintf("%d_%s", orderField.GetSessionID(), orderField.GetOrderRef())
	o, ok := t.Orders[key]
	if !ok {
		o = new(OrderField)
		t.Orders[key] = o
		o.InstrumentID = orderField.GetInstrumentID()
		o.SessionID = orderField.GetSessionID()
		o.FrontID = orderField.GetFrontID()
		o.OrderRef = orderField.GetOrderRef()
		o.Direction = DirectionType(orderField.GetDirection())
		o.OffsetFlag = OffsetFlagType(orderField.GetCombOffsetFlag()[0])
		o.HedgeFlag = HedgeFlagType(orderField.GetCombHedgeFlag()[0])
		o.LimitPrice = orderField.GetLimitPrice()
		o.VolumeTotalOriginal = orderField.GetVolumeTotalOriginal()
		o.VolumeLeft = o.VolumeTotalOriginal
		o.ExchangeID = orderField.GetExchangeID()
		o.InsertDate = orderField.GetInsertDate()
		o.InsertTime = orderField.GetInsertTime()
		o.OrderStatus = OrderStatusNoTradeQueueing // OrderStatusType(orderField.OrderStatus)
		o.StatusMsg = "委托已提交"                      // bytes2GBKbytes2GBKString(orderField.StatusMsg[:])
		o.IsLocal = orderField.GetSessionID() == t.sessionID
		if t.onRtnOrder != nil {
			t.onRtnOrder(o)
		}
	} else if o.OrderStatus == OrderStatusCanceled {
		o.CancelTime = orderField.GetCancelTime()
		// 错单
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
	} else {
		o.OrderSysID = orderField.GetOrderSysID()
		if len(o.OrderSysID) > 0 {
			t.sysID4Order[o.OrderSysID] = o
		}

		// 有成交的响应由onTrade触发
		//if o.ExchangeID == "CZCE" {
		//	o.VolumeTraded = int(orderField.ZCETotalTradedVolume)
		//} else {
		//	o.VolumeTraded = int(orderField.VolumeTraded)
		//}
		//o.VolumeLeft = int(orderField.VolumeTotal)
		//o.OrderStatus = OrderStatusType(orderField.OrderStatus)
		//// 避免出现"全部成交XXX"这种
		//if OrderStatusAllTraded == o.OrderStatus {
		//	o.StatusMsg = "全部成交"
		//} else {
		//	o.StatusMsg = bytes2GBKbytes2GBKString(orderField.StatusMsg[:])
		//}

		//else if t.onRtnOrder != nil {
		//	t.onRtnOrder(o)
		//}
	}
}

// 委托错误响应
func (t *Trade) OnErrRtnOrderInsert(orderField CThostFtdcInputOrderField, infoField CThostFtdcRspInfoField) {
	key := fmt.Sprintf("%d_%s", t.sessionID, orderField.GetOrderRef())
	o, ok := t.Orders[key]
	if !ok {
		o = new(OrderField)
		t.Orders[key] = o
		o.InstrumentID = orderField.GetInstrumentID()
		o.SessionID = t.sessionID
		o.FrontID = 0
		o.OrderRef = orderField.GetOrderRef()
		o.Direction = DirectionType(orderField.GetDirection())
		o.OffsetFlag = OffsetFlagType(orderField.GetCombOffsetFlag()[0])
		o.HedgeFlag = HedgeFlagType(orderField.GetCombHedgeFlag()[0])
		o.LimitPrice = orderField.GetLimitPrice()
		o.VolumeTotalOriginal = orderField.GetVolumeTotalOriginal()
		o.VolumeLeft = o.VolumeTotalOriginal
		o.ExchangeID = orderField.GetExchangeID()
		o.IsLocal = true
	}
	o.OrderStatus = OrderStatusCanceled
	if t.onErrRtnOrder != nil {
		t.onErrRtnOrder(o, &RspInfoField{ErrorID: infoField.GetErrorID(), ErrorMsg: infoField.GetErrorMsg()})
	}
}

// 持仓查询响应
func (t *Trade) OnRspQryInvestorPosition(positionField CThostFtdcInvestorPositionField, infoField CThostFtdcRspInfoField, i int, b bool) {
	if strings.Compare(positionField.GetInstrumentID(), "") != 0 {
		key := fmt.Sprintf("%s_%c_%c", positionField.GetInstrumentID(), positionField.GetPosiDirection(), positionField.GetHedgeFlag())
		p, ok := t.Positions[key]
		if !ok {
			p = new(PositionField)
			t.Positions[key] = p
			p.InstrumentID = positionField.GetInstrumentID()
			p.PositionDirection = PosiDirectionType(positionField.GetPosiDirection())
			p.HedgeFlag = HedgeFlagType(positionField.GetHedgeFlag())
			p.ExchangeID = positionField.GetExchangeID()
		}
		p.YdPosition = positionField.GetYdPosition()
		p.Position = positionField.GetPosition()
		p.LongFrozen = positionField.GetLongFrozen()
		p.ShortFrozen = positionField.GetShortFrozen()
		p.LongFrozenAmount = positionField.GetLongFrozenAmount()
		p.ShortFrozenAmount = positionField.GetShortFrozenAmount()
		p.OpenVolume = positionField.GetOpenVolume()
		p.CloseVolume = positionField.GetCloseVolume()
		p.OpenAmount = positionField.GetOpenAmount()
		p.CloseAmount = positionField.GetCloseAmount()
		p.PositionCost = positionField.GetPositionCost()
		p.PreMargin = positionField.GetPreMargin()
		p.UseMargin = positionField.GetUseMargin()
		p.FrozenMargin = positionField.GetFrozenMargin()
		p.FrozenCash = positionField.GetFrozenCash()
		p.FrozenCommission = positionField.GetFrozenCommission()
		p.CashIn = positionField.GetCashIn()
		p.Commission = positionField.GetCommission()
		p.CloseProfit = positionField.GetCloseProfit()
		p.PositionProfit = positionField.GetPositionProfit()
		p.PreSettlementPrice = positionField.GetPreSettlementPrice()
		p.SettlementPrice = positionField.GetSettlementPrice()
		p.OpenCost = positionField.GetOpenCost()
		p.ExchangeMargin = positionField.GetExchangeMargin()
		p.CombPosition = positionField.GetCombPosition()
		p.CombLongFrozen = positionField.GetCombLongFrozen()
		p.CombShortFrozen = positionField.GetCombShortFrozen()
		p.CloseProfitByDate = positionField.GetCloseProfitByDate()
		p.CloseProfitByTrade = positionField.GetCloseProfitByTrade()
		p.TodayPosition = positionField.GetTodayPosition()
		p.StrikeFrozen = positionField.GetStrikeFrozen()
		p.StrikeFrozenAmount = positionField.GetStrikeFrozenAmount()
		p.AbandonFrozen = positionField.GetAbandonFrozen()
		p.YdStrikeFrozen = positionField.GetYdStrikeFrozen()
		p.PositionCostOffset = positionField.GetPositionCostOffset()
	}
}

// 账户资金响应
func (t *Trade) OnRspQryTradingAccount(accountField CThostFtdcTradingAccountField, infoField CThostFtdcRspInfoField, i int, b bool) {
	t.Account.PreMortgage = accountField.GetPreMortgage()
	t.Account.PreDeposit = accountField.GetPreDeposit()
	t.Account.PreBalance = accountField.GetPreBalance()
	t.Account.PreMargin = accountField.GetPreMargin()
	t.Account.InterestBase = accountField.GetInterestBase()
	t.Account.Interest = accountField.GetInterest()
	t.Account.Deposit = accountField.GetDeposit()
	t.Account.Withdraw = accountField.GetWithdraw()
	t.Account.FrozenMargin = accountField.GetFrozenMargin()
	t.Account.FrozenCash = accountField.GetFrozenCash()
	t.Account.FrozenCommission = accountField.GetFrozenCommission()
	t.Account.CurrMargin = accountField.GetCurrMargin()
	t.Account.CashIn = accountField.GetCashIn()
	t.Account.Commission = accountField.GetCommission()
	t.Account.CloseProfit = accountField.GetCloseProfit()
	t.Account.PositionProfit = accountField.GetPositionProfit()
	t.Account.Balance = accountField.GetBalance()
	t.Account.Available = accountField.GetAvailable()
	t.Account.WithdrawQuota = accountField.GetWithdrawQuota()
	t.Account.Reserve = accountField.GetReserve()
	t.Account.Credit = accountField.GetCredit()
	t.Account.Mortgage = accountField.GetMortgage()
	t.Account.ExchangeMargin = accountField.GetExchangeMargin()
	t.Account.DeliveryMargin = accountField.GetDeliveryMargin()
	t.Account.ExchangeDeliveryMargin = accountField.GetExchangeDeliveryMargin()
	t.Account.ReserveBalance = accountField.GetReserveBalance()
	t.Account.CurrencyID = accountField.GetCurrencyID()
	t.Account.PreFundMortgageIn = accountField.GetPreFundMortgageIn()
	t.Account.PreFundMortgageOut = accountField.GetPreFundMortgageOut()
	t.Account.FundMortgageIn = accountField.GetFundMortgageIn()
	t.Account.FundMortgageOut = accountField.GetFundMortgageOut()
	t.Account.FundMortgageAvailable = accountField.GetFundMortgageAvailable()
	t.Account.MortgageableFund = accountField.GetMortgageableFund()
}

// 合约查询响应
func (t *Trade) OnRspQryInstrument(instrumentField CThostFtdcInstrumentField, infoField CThostFtdcRspInfoField, i int, b bool) {
	if instrumentField != nil {
		t.Instruments[instrumentField.GetInstrumentID()] = InstrumentField{
			InstrumentID:              instrumentField.GetInstrumentID(),
			ExchangeID:                instrumentField.GetExchangeID(),
			ProductID:                 instrumentField.GetProductID(),
			ProductClass:              ProductClassType(instrumentField.GetProductClass()),
			MaxMarketOrderVolume:      instrumentField.GetMaxMarketOrderVolume(),
			MinMarketOrderVolume:      instrumentField.GetMinMarketOrderVolume(),
			MaxLimitOrderVolume:       instrumentField.GetMaxLimitOrderVolume(),
			MinLimitOrderVolume:       instrumentField.GetMinLimitOrderVolume(),
			VolumeMultiple:            instrumentField.GetVolumeMultiple(),
			PriceTick:                 instrumentField.GetPriceTick(),
			PositionType:              PositionTypeType(instrumentField.GetPositionType()),
			UseMaxMarginSideAlgorithm: instrumentField.GetMaxMarginSideAlgorithm() == '1',
			UnderlyingInstrID:         instrumentField.GetUnderlyingInstrID(),
			StrikePrice:               instrumentField.GetStrikePrice(),
			OptionsType:               OptionsTypeType(instrumentField.GetOptionsType()),
			UnderlyingMultiple:        instrumentField.GetUnderlyingMultiple(),
			CombinationType:           CombinationTypeType(instrumentField.GetCombinationType()),
		}
	}
	if b {
		fmt.Printf("收到合约: %d\n", len(t.Instruments))
		// 登录成功响应
		t.IsLogin = true

		// 循环查询持仓与权益
		go t.qry()
	}
}

// 循环查询持仓&资金
func (t *Trade) qry() {
	t.qryTicker = time.NewTicker(1100 * time.Millisecond)
	// 等待之前的Order响应完再发送登录通知
	ordCnt := len(t.Orders)
	for range t.qryTicker.C {
		if ordCnt == len(t.Orders) {
			break
		}
		ordCnt = len(t.Orders)
	}
	// 通知:登录响应可以发了
	t.waitGroup.Done()
	qryAccount := NewCThostFtdcQryTradingAccountField()
	qryAccount.SetInvestorID(t.InvestorID)
	qryAccount.SetBrokerID(t.BrokerID)

	qryPosition := NewCThostFtdcQryInvestorPositionField()
	qryPosition.SetInvestorID(t.InvestorID)
	qryPosition.SetBrokerID(t.BrokerID)
	// 启动查询
	bQryAccount := false
	for range t.qryTicker.C {
		if bQryAccount {
			t.api.ReqQryTradingAccount(qryAccount, t.getReqID())
		} else {
			t.api.ReqQryInvestorPosition(qryPosition, t.getReqID())
		}
		bQryAccount = !bQryAccount
		if !t.IsLogin {
			break
		}
	}
	t.qryTicker.Stop()
}

// 确认结算相应
func (t *Trade) OnRspSettlementInfoConfirm(confirmField CThostFtdcSettlementInfoConfirmField, infoField CThostFtdcRspInfoField, i int, b bool) {
	t.api.ReqQryInstrument(NewCThostFtdcQryInstrumentField(), t.getReqID())
}

// 登陆响应
func (t *Trade) OnRspUserLogin(loginField CThostFtdcRspUserLoginField, infoField CThostFtdcRspInfoField, i int, b bool) {
	if infoField.GetErrorID() == 0 {
		t.sessionID = loginField.GetSessionID()
		t.TradingDay = loginField.GetTradingDay()
		f := NewCThostFtdcSettlementInfoConfirmField()
		f.SetInvestorID(t.InvestorID)
		f.SetAccountID(t.InvestorID)
		f.SetBrokerID(t.BrokerID)
		t.api.ReqSettlementInfoConfirm(f, t.getReqID())
		if t.onRspUserLogin != nil {
			t.waitGroup.Add(1)
			go func(field *RspUserLoginField) {
				t.waitGroup.Wait()
				t.onRspUserLogin(field, &RspInfoField{ErrorID: 0, ErrorMsg: "成功"})
			}(&RspUserLoginField{
				TradingDay:  t.TradingDay,
				LoginTime:   loginField.GetLoginTime(),
				BrokerID:    t.BrokerID,
				UserID:      t.InvestorID,
				FrontID:     int(loginField.GetFrontID()),
				SessionID:   t.sessionID,
				MaxOrderRef: loginField.GetMaxOrderRef(),
			})
		}
	} else if t.onRspUserLogin != nil {
		t.onRspUserLogin(&RspUserLoginField{}, &RspInfoField{ErrorID: infoField.GetErrorID(), ErrorMsg: infoField.GetErrorMsg()})
	}
}

// 看穿式验证响应
func (t *Trade) OnRspAuthenticate(auth CThostFtdcRspAuthenticateField, info CThostFtdcRspInfoField, i int, b bool) {
	if info.GetErrorID() == 0 {
		f := NewCThostFtdcReqUserLoginField()
		f.SetUserID(t.InvestorID)
		f.SetBrokerID(t.BrokerID)
		f.SetPassword(t.passWord)
		f.SetUserProductInfo("@HaiFeng")
		t.api.ReqUserLogin(f, t.getReqID())
	} else if t.onRspUserLogin != nil {
		t.onRspUserLogin(&RspUserLoginField{}, &RspInfoField{ErrorID: info.GetErrorID(), ErrorMsg: info.GetErrorMsg()})
	}
}

// 连接前置响应
func (t *Trade) OnFrontConnected() {
	if t.onFrontConnected != nil {
		t.onFrontConnected()
	}
}
