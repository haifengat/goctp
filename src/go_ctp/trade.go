package go_ctp

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type OnRtnOrderType func(field *OrderField)
type OnRtnErrOrderType func(field *OrderField, info *RspInfoField)
type OnRtnTradeType func(field *TradeField)

type Trade struct {
	t *trade
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

	t.t = newTrade()
	t.t.regOnFrontConnected(t.onConnected)
	t.t.regOnRspAuthenticate(t.onAuth)
	t.t.regOnRspSettlementInfoConfirm(t.onSettlement)
	t.t.regOnRspUserLogin(t.onLogin)
	t.t.regOnRspQryInstrument(t.onInstrument)
	t.t.regOnRspQryTradingAccount(t.onAccount)
	t.t.regOnRspQryInvestorPosition(t.onPosition)
	t.t.regOnRtnOrder(t.onOrder)
	t.t.regOnErrRtnOrderInsert(t.onErrOrder)
	t.t.regOnRtnTrade(t.onTrade)
	t.t.regOnRtnInstrumentStatus(t.onStatus)
	return t
}

// ********************** 主调函数 ************************

// 接口销毁处理
func (t *Trade) Release() {
	t.IsLogin = false
	t.t.Release()
}

// 连接
func (t *Trade) ReqConnect(addr string) {
	t.t.RegisterFront(addr)
	t.t.SubscribePrivateTopic(THOST_TERT_RESTART)
	t.t.SubscribePublicTopic(THOST_TERT_RESTART)
	t.t.Init()
}

// 登录
func (t *Trade) ReqLogin(investor, pwd, broker, appID, authCode string) {
	t.InvestorID = investor
	t.passWord = pwd
	t.BrokerID = broker
	f := tCThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.AppID[:], appID)
	copy(f.AuthCode[:], authCode)
	t.t.ReqAuthenticate(f)
}

// 限价委托
// 返回委托的ID
func (t *Trade) ReqOrderInsert(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := tCThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	copy(f.ExchangeID[:], t.Instruments[instrument].ExchangeID)
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = tTThostFtdcBoolType(0)
	f.IsSwapOrder = tTThostFtdcBoolType(0)
	f.ForceCloseReason = TThostFtdcForceCloseReasonType_NotForceClose
	// 参数赋值
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = tTThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = TThostFtdcOrderPriceTypeType_LimitPrice
	f.TimeCondition = TThostFtdcTimeConditionType_IOC
	f.VolumeCondition = TThostFtdcVolumeConditionType_AV
	f.ContingentCondition = TThostFtdcContingentConditionType_Immediately
	f.LimitPrice = tTThostFtdcPriceType(price)
	f.VolumeTotalOriginal = tTThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, f.OrderRef)
}

// 市价委托
func (t *Trade) ReqOrderInsertMarket(instrument string, buySell DirectionType, openClose OffsetFlagType, volume int) string {
	f := tCThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	copy(f.ExchangeID[:], t.Instruments[instrument].ExchangeID)
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = tTThostFtdcBoolType(0)
	f.IsSwapOrder = tTThostFtdcBoolType(0)
	f.ForceCloseReason = TThostFtdcForceCloseReasonType_NotForceClose
	// 参数赋值
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = tTThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = TThostFtdcOrderPriceTypeType_AnyPrice
	f.TimeCondition = TThostFtdcTimeConditionType_IOC
	f.VolumeCondition = TThostFtdcVolumeConditionType_AV
	f.ContingentCondition = TThostFtdcContingentConditionType_Immediately
	f.LimitPrice = tTThostFtdcPriceType(0)
	f.VolumeTotalOriginal = tTThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, f.OrderRef)
}

// FOK委托[部成撤单]
func (t *Trade) ReqOrderInsertFOK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := tCThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	copy(f.ExchangeID[:], t.Instruments[instrument].ExchangeID)
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = tTThostFtdcBoolType(0)
	f.IsSwapOrder = tTThostFtdcBoolType(0)
	f.ForceCloseReason = TThostFtdcForceCloseReasonType_NotForceClose
	// 参数赋值
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = tTThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = TThostFtdcOrderPriceTypeType_LimitPrice
	f.TimeCondition = TThostFtdcTimeConditionType_GFD
	f.VolumeCondition = TThostFtdcVolumeConditionType_CV
	f.ContingentCondition = TThostFtdcContingentConditionType_Immediately
	f.LimitPrice = tTThostFtdcPriceType(price)
	f.VolumeTotalOriginal = tTThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, f.OrderRef)
}

// FAK委托[全成or撤单]
func (t *Trade) ReqOrderInsertFAK(instrument string, buySell DirectionType, openClose OffsetFlagType, price float64, volume int) string {
	f := tCThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	copy(f.ExchangeID[:], t.Instruments[instrument].ExchangeID)
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = tTThostFtdcBoolType(0)
	f.IsSwapOrder = tTThostFtdcBoolType(0)
	f.ForceCloseReason = TThostFtdcForceCloseReasonType_NotForceClose
	// 参数赋值
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = tTThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = TThostFtdcOrderPriceTypeType_LimitPrice
	f.TimeCondition = TThostFtdcTimeConditionType_IOC
	f.VolumeCondition = TThostFtdcVolumeConditionType_AV
	f.ContingentCondition = TThostFtdcContingentConditionType_Immediately
	f.LimitPrice = tTThostFtdcPriceType(price)
	f.VolumeTotalOriginal = tTThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, f.OrderRef)
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
func (t *Trade) onStatus(statusField *tCThostFtdcInstrumentStatusField) uintptr {
	t.InstrumentStatuss[bytes2String(statusField.InstrumentID[:])] = InstrumentStatusType(statusField.InstrumentStatus)
	return 0
}

// 成交响应
func (t *Trade) onTrade(tradeField *tCThostFtdcTradeField) uintptr {
	key := fmt.Sprintf("%s_%c", tradeField.TradeID, tradeField.Direction)
	f, ok := t.Trades[key]
	if !ok {
		f = new(TradeField)
		t.Trades[key] = f
		f.Direction = DirectionType(tradeField.Direction)
		f.HedgeFlag = HedgeFlagType(tradeField.HedgeFlag)
		f.InstrumentID = bytes2String(tradeField.InstrumentID[:])
		f.ExchangeID = bytes2String(tradeField.ExchangeID[:])
		f.TradingDay = bytes2String(tradeField.TradingDay[:])
		f.Volume = int(tradeField.Volume)
		f.OffsetFlag = OffsetFlagType(tradeField.OffsetFlag)
		f.OrderSysID = bytes2String(tradeField.OrderSysID[:])
		f.Price = float64(tradeField.Price)
		f.TradeDate = bytes2String(tradeField.TradeDate[:])
		f.TradeTime = bytes2String(tradeField.TradeTime[:])
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
	return 0
}

// 委托响应
func (t *Trade) onOrder(orderField *tCThostFtdcOrderField) uintptr {
	key := fmt.Sprintf("%d_%s", orderField.SessionID, orderField.OrderRef)
	o, ok := t.Orders[key]
	if !ok {
		o = new(OrderField)
		t.Orders[key] = o
		o.InstrumentID = bytes2String(orderField.InstrumentID[:])
		o.SessionID = int(orderField.SessionID)
		o.FrontID = int(orderField.FrontID)
		o.OrderRef = bytes2String(orderField.OrderRef[:])
		o.Direction = DirectionType(orderField.Direction)
		o.OffsetFlag = OffsetFlagType(orderField.CombOffsetFlag[0])
		o.HedgeFlag = HedgeFlagType(orderField.CombHedgeFlag[0])
		o.LimitPrice = float64(orderField.LimitPrice)
		o.VolumeTotalOriginal = int(orderField.VolumeTotalOriginal)
		o.VolumeLeft = o.VolumeTotalOriginal
		o.ExchangeID = bytes2String(orderField.ExchangeID[:])
		o.InsertDate = bytes2String(orderField.InsertDate[:])
		o.InsertTime = bytes2String(orderField.InsertTime[:])
		o.OrderStatus = OrderStatusNoTradeQueueing // OrderStatusType(orderField.OrderStatus)
		o.StatusMsg = "委托已提交"                      // bytes2GBKbytes2GBKString(orderField.StatusMsg[:])
		o.IsLocal = int(orderField.SessionID) == t.sessionID
		if t.onRtnOrder != nil {
			t.onRtnOrder(o)
		}
	} else if o.OrderStatus == OrderStatusCanceled {
		o.CancelTime = bytes2String(orderField.CancelTime[:])
		// 错单
		if strings.Contains(o.StatusMsg, "被拒绝") {
			if t.onErrRtnOrder != nil {
				t.onErrRtnOrder(o, &RspInfoField{-1, o.StatusMsg})
			}
		} else if t.onRtnCancel != nil {
			t.onRtnCancel(o)
		}
	} else {
		o.OrderSysID = bytes2String(orderField.OrderSysID[:])
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
	return 0
}

// 委托错误响应
func (t *Trade) onErrOrder(orderField *tCThostFtdcInputOrderField, infoField *tCThostFtdcRspInfoField) uintptr {
	key := fmt.Sprintf("%d_%s", t.sessionID, orderField.OrderRef)
	o, ok := t.Orders[key]
	if !ok {
		o = new(OrderField)
		t.Orders[key] = o
		o.InstrumentID = bytes2String(orderField.InstrumentID[:])
		o.SessionID = t.sessionID
		o.FrontID = 0
		o.OrderRef = bytes2String(orderField.OrderRef[:])
		o.Direction = DirectionType(orderField.Direction)
		o.OffsetFlag = OffsetFlagType(orderField.CombOffsetFlag[0])
		o.HedgeFlag = HedgeFlagType(orderField.CombHedgeFlag[0])
		o.LimitPrice = float64(orderField.LimitPrice)
		o.VolumeTotalOriginal = int(orderField.VolumeTotalOriginal)
		o.VolumeLeft = o.VolumeTotalOriginal
		o.ExchangeID = bytes2String(orderField.ExchangeID[:])
		o.IsLocal = true
	}
	o.OrderStatus = OrderStatusCanceled
	if t.onErrRtnOrder != nil {
		t.onErrRtnOrder(o, &RspInfoField{int(infoField.ErrorID), bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

// 持仓查询响应
func (t *Trade) onPosition(positionField *tCThostFtdcInvestorPositionField, infoField *tCThostFtdcRspInfoField, i int, b bool) uintptr {
	if strings.Compare(bytes2String(positionField.InstrumentID[:]), "") != 0 {
		key := fmt.Sprintf("%s_%c_%c", positionField.InstrumentID, positionField.PosiDirection, positionField.HedgeFlag)
		p, ok := t.Positions[key]
		if !ok {
			p = new(PositionField)
			t.Positions[key] = p
			p.InstrumentID = bytes2String(positionField.InstrumentID[:])
			p.PositionDirection = PosiDirectionType(positionField.PosiDirection)
			p.HedgeFlag = HedgeFlagType(positionField.HedgeFlag)
			p.ExchangeID = bytes2String(positionField.ExchangeID[:])
		}
		p.YdPosition = int(positionField.YdPosition)
		p.Position = int(positionField.Position)
		p.LongFrozen = int(positionField.LongFrozen)
		p.ShortFrozen = int(positionField.ShortFrozen)
		p.LongFrozenAmount = float64(positionField.LongFrozenAmount)
		p.ShortFrozenAmount = float64(positionField.ShortFrozenAmount)
		p.OpenVolume = int(positionField.OpenVolume)
		p.CloseVolume = int(positionField.CloseVolume)
		p.OpenAmount = float64(positionField.OpenAmount)
		p.CloseAmount = float64(positionField.CloseAmount)
		p.PositionCost = float64(positionField.PositionCost)
		p.PreMargin = float64(positionField.PreMargin)
		p.UseMargin = float64(positionField.UseMargin)
		p.FrozenMargin = float64(positionField.FrozenMargin)
		p.FrozenCash = float64(positionField.FrozenCash)
		p.FrozenCommission = float64(positionField.FrozenCommission)
		p.CashIn = float64(positionField.CashIn)
		p.Commission = float64(positionField.Commission)
		p.CloseProfit = float64(positionField.CloseProfit)
		p.PositionProfit = float64(positionField.PositionProfit)
		p.PreSettlementPrice = float64(positionField.PreSettlementPrice)
		p.SettlementPrice = float64(positionField.SettlementPrice)
		p.OpenCost = float64(positionField.OpenCost)
		p.ExchangeMargin = float64(positionField.ExchangeMargin)
		p.CombPosition = int(positionField.CombPosition)
		p.CombLongFrozen = int(positionField.CombLongFrozen)
		p.CombShortFrozen = int(positionField.CombShortFrozen)
		p.CloseProfitByDate = float64(positionField.CloseProfitByDate)
		p.CloseProfitByTrade = float64(positionField.CloseProfitByTrade)
		p.TodayPosition = int(positionField.TodayPosition)
		p.StrikeFrozen = int(positionField.StrikeFrozen)
		p.StrikeFrozenAmount = float64(positionField.StrikeFrozenAmount)
		p.AbandonFrozen = int(positionField.AbandonFrozen)
		p.YdStrikeFrozen = int(positionField.YdStrikeFrozen)
		p.PositionCostOffset = float64(positionField.PositionCostOffset)
	}
	return 0
}

// 账户资金响应
func (t *Trade) onAccount(accountField *tCThostFtdcTradingAccountField, infoField *tCThostFtdcRspInfoField, i int, b bool) uintptr {
	t.Account.PreMortgage = float64(accountField.PreMortgage)
	t.Account.PreDeposit = float64(accountField.PreDeposit)
	t.Account.PreBalance = float64(accountField.PreBalance)
	t.Account.PreMargin = float64(accountField.PreMargin)
	t.Account.InterestBase = float64(accountField.InterestBase)
	t.Account.Interest = float64(accountField.Interest)
	t.Account.Deposit = float64(accountField.Deposit)
	t.Account.Withdraw = float64(accountField.Withdraw)
	t.Account.FrozenMargin = float64(accountField.FrozenMargin)
	t.Account.FrozenCash = float64(accountField.FrozenCash)
	t.Account.FrozenCommission = float64(accountField.FrozenCommission)
	t.Account.CurrMargin = float64(accountField.CurrMargin)
	t.Account.CashIn = float64(accountField.CashIn)
	t.Account.Commission = float64(accountField.Commission)
	t.Account.CloseProfit = float64(accountField.CloseProfit)
	t.Account.PositionProfit = float64(accountField.PositionProfit)
	t.Account.Balance = float64(accountField.Balance)
	t.Account.Available = float64(accountField.Available)
	t.Account.WithdrawQuota = float64(accountField.WithdrawQuota)
	t.Account.Reserve = float64(accountField.Reserve)
	t.Account.Credit = float64(accountField.Credit)
	t.Account.Mortgage = float64(accountField.Mortgage)
	t.Account.ExchangeMargin = float64(accountField.ExchangeMargin)
	t.Account.DeliveryMargin = float64(accountField.DeliveryMargin)
	t.Account.ExchangeDeliveryMargin = float64(accountField.ExchangeDeliveryMargin)
	t.Account.ReserveBalance = float64(accountField.ReserveBalance)
	t.Account.CurrencyID = bytes2String(accountField.CurrencyID[:])
	t.Account.PreFundMortgageIn = float64(accountField.PreFundMortgageIn)
	t.Account.PreFundMortgageOut = float64(accountField.PreFundMortgageOut)
	t.Account.FundMortgageIn = float64(accountField.FundMortgageIn)
	t.Account.FundMortgageOut = float64(accountField.FundMortgageOut)
	t.Account.FundMortgageAvailable = float64(accountField.FundMortgageAvailable)
	t.Account.MortgageableFund = float64(accountField.MortgageableFund)
	return 0
}

// 合约查询响应
func (t *Trade) onInstrument(instrumentField *tCThostFtdcInstrumentField, infoField *tCThostFtdcRspInfoField, i int, b bool) uintptr {
	if instrumentField != nil {
		t.Instruments[bytes2String(instrumentField.InstrumentID[:])] = InstrumentField{
			bytes2String(instrumentField.InstrumentID[:]),
			bytes2String(instrumentField.ExchangeID[:]),
			bytes2String(instrumentField.ProductID[:]),
			ProductClassType(instrumentField.ProductClass),
			int(instrumentField.MaxMarketOrderVolume),
			int(instrumentField.MinMarketOrderVolume),
			int(instrumentField.MaxLimitOrderVolume),
			int(instrumentField.MinLimitOrderVolume),
			int(instrumentField.VolumeMultiple),
			float64(instrumentField.PriceTick),
			PositionTypeType(instrumentField.PositionType),
			instrumentField.MaxMarginSideAlgorithm == '1',
			bytes2String(instrumentField.UnderlyingInstrID[:]),
			float64(instrumentField.StrikePrice),
			OptionsTypeType(instrumentField.OptionsType),
			int(instrumentField.UnderlyingMultiple),
			CombinationTypeType(instrumentField.CombinationType),
		}
	}
	if b {
		fmt.Printf("收到合约: %d\n", len(t.Instruments))
		// 登录成功响应
		t.IsLogin = true

		// 循环查询持仓与权益
		go t.qry()
	}
	return 0
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
	qryAccount := tCThostFtdcQryTradingAccountField{}
	copy(qryAccount.InvestorID[:], t.InvestorID)
	copy(qryAccount.BrokerID[:], t.BrokerID)
	qryPosition := tCThostFtdcQryInvestorPositionField{}
	copy(qryPosition.InvestorID[:], t.InvestorID)
	copy(qryPosition.BrokerID[:], t.BrokerID)
	// 启动查询
	bQryAccount := false
	for range t.qryTicker.C {
		if bQryAccount {
			t.t.ReqQryTradingAccount(qryAccount)
		} else {
			t.t.ReqQryInvestorPosition(qryPosition)
		}
		bQryAccount = !bQryAccount
		if !t.IsLogin {
			break
		}
	}
	t.qryTicker.Stop()
}

// 确认结算相应
func (t *Trade) onSettlement(confirmField *tCThostFtdcSettlementInfoConfirmField, infoField *tCThostFtdcRspInfoField, i int, b bool) uintptr {
	t.t.ReqQryInstrument(tCThostFtdcQryInstrumentField{})
	return 0
}

// 登陆响应
func (t *Trade) onLogin(loginField *tCThostFtdcRspUserLoginField, infoField *tCThostFtdcRspInfoField, i int, b bool) uintptr {
	if infoField.ErrorID == 0 {
		t.sessionID = int(loginField.SessionID)
		t.TradingDay = bytes2String(loginField.TradingDay[:])
		f := tCThostFtdcSettlementInfoConfirmField{}
		copy(f.InvestorID[:], t.InvestorID)
		copy(f.AccountID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		t.t.ReqSettlementInfoConfirm(f)
		if t.onRspUserLogin != nil {
			t.waitGroup.Add(1)
			go func(field *RspUserLoginField) {
				t.waitGroup.Wait()
				t.onRspUserLogin(field, &RspInfoField{ErrorID: 0, ErrorMsg: "成功"})
			}(&RspUserLoginField{
				t.TradingDay,
				bytes2String(loginField.LoginTime[:]),
				t.BrokerID,
				t.InvestorID,
				int(loginField.FrontID),
				t.sessionID,
				bytes2String(loginField.MaxOrderRef[:]),
			})
		}
	} else {
		t.onRspUserLogin(&RspUserLoginField{}, &RspInfoField{int(infoField.ErrorID), bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

// 看穿式验证响应
func (t *Trade) onAuth(auth *tCThostFtdcRspAuthenticateField, info *tCThostFtdcRspInfoField, i int, b bool) uintptr {
	if info.ErrorID == 0 {
		f := tCThostFtdcReqUserLoginField{}
		copy(f.UserID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.Password[:], t.passWord)
		copy(f.UserProductInfo[:], "@HF")
		t.t.ReqUserLogin(f)
	} else if t.onRspUserLogin != nil {
		t.onRspUserLogin(&RspUserLoginField{}, &RspInfoField{int(info.ErrorID), bytes2String(info.ErrorMsg[:])})
	}
	return 0
}

// 连接前置响应
func (t *Trade) onConnected() uintptr {
	if t.onFrontConnected != nil {
		t.onFrontConnected()
	}
	return 0
}
