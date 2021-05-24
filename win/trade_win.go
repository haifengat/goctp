package win

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

// Trade 交易接口
type Trade struct {
	t *trade
	// 密码
	passWord string
	// 判断是否自己的委托用
	sessionID         int
	InvestorID        string   // 帐号
	BrokerID          string   // 经纪商
	TradingDay        string   // 交易日
	Instruments       sync.Map // 合约列表
	InstrumentStatuss sync.Map // 合约状态
	Positions         sync.Map // 持仓列表
	Orders            sync.Map // 委托
	Trades            sync.Map // 成交
	sysID4Order       sync.Map // orderSysID 对应的 Order
	// 帐户权益
	Account *goctp.AccountField
	// 登录成功
	IsLogin bool
	// 循环查询
	qryTicker             *time.Ticker
	onFrontConnected      goctp.OnFrontConnectedType
	onRspUserLogin        goctp.OnRspUserLoginType
	onRtnOrder            goctp.OnRtnOrderType
	onRtnCancel           goctp.OnRtnOrderType
	onErrRtnOrder         goctp.OnRtnErrOrderType
	onErrAction           goctp.OnRtnErrActionType
	onRtnTrade            goctp.OnRtnTradeType
	onRtnInstrumentStatus goctp.OnRtnInstrumentStatusType
	// chan 登录信号
	waitGroup sync.WaitGroup

	cntOrder int
}

// NewTrade 交易接口实例
func NewTrade() *Trade {
	t := new(Trade)
	// 初始化变量
	t.waitGroup = sync.WaitGroup{}
	t.IsLogin = false
	t.Account = &goctp.AccountField{}

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
	t.t.regOnErrRtnOrderAction(t.onErrRtnOrderAction)
	return t
}

// ********************** 主调函数 ************************

// Release 接口销毁处理
func (t *Trade) Release() {
	t.IsLogin = false
	t.t.Release()
}

// ReqConnect 连接
func (t *Trade) ReqConnect(addr string) {
	t.t.RegisterFront(addr)
	t.t.SubscribePrivateTopic(ctp.THOST_TERT_RESTART)
	t.t.SubscribePublicTopic(ctp.THOST_TERT_RESTART)
	t.t.Init()
}

// ReqLogin 登录
func (t *Trade) ReqLogin(investor, pwd, broker, appID, authCode string) {
	t.InvestorID = investor
	t.passWord = pwd
	t.BrokerID = broker
	f := ctp.CThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.AppID[:], appID)
	copy(f.AuthCode[:], authCode)
	t.t.ReqAuthenticate(f)
}

// ReqOrderInsert 限价委托
func (t *Trade) ReqOrderInsert(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	t.t.nRequestID++
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_GFD
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertMarket 市价委托
func (t *Trade) ReqOrderInsertMarket(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	t.t.nRequestID++
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_AnyPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_IOC
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(0)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertFOK FOK委托[部成撤单]
func (t *Trade) ReqOrderInsertFOK(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	t.t.nRequestID++
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_GFD
	f.VolumeCondition = ctp.THOST_FTDC_VC_CV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertFAK FAK委托[全成or撤单]
func (t *Trade) ReqOrderInsertFAK(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	t.t.nRequestID++
	copy(f.OrderRef[:], fmt.Sprintf("%012d", t.t.nRequestID))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_IOC
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	t.t.ReqOrderInsert(f)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderAction 撤单
func (t *Trade) ReqOrderAction(orderID string) uintptr {
	if o, ok := t.Orders.Load(orderID); ok {
		var order = o.(*goctp.OrderField)
		f := ctp.CThostFtdcInputOrderActionField{}
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.UserID[:], t.InvestorID)
		copy(f.InstrumentID[:], order.InstrumentID)
		copy(f.ExchangeID[:], order.ExchangeID)
		copy(f.OrderRef[:], order.OrderRef)
		f.FrontID = ctp.TThostFtdcFrontIDType(order.FrontID)
		f.SessionID = ctp.TThostFtdcSessionIDType(order.SessionID)
		t.t.ReqOrderAction(f)
		return 0
	}
	return 0
}

// ********************** 注册客户响应 ************************

// RegOnFrontConnected 注册连接响应
func (t *Trade) RegOnFrontConnected(on goctp.OnFrontConnectedType) {
	t.onFrontConnected = on
}

// RegOnRspUserLogin 注册登陆响应
func (t *Trade) RegOnRspUserLogin(on goctp.OnRspUserLoginType) {
	t.onRspUserLogin = on
}

// RegOnRtnOrder 注册委托响应
func (t *Trade) RegOnRtnOrder(on goctp.OnRtnOrderType) {
	t.onRtnOrder = on
}

// RegOnErrRtnOrder 注册委托响应
func (t *Trade) RegOnErrRtnOrder(on goctp.OnRtnErrOrderType) {
	t.onErrRtnOrder = on
}

// RegOnErrAction 注册撤单响应
func (t *Trade) RegOnErrAction(on goctp.OnRtnErrActionType) {
	t.onErrAction = on
}

// RegOnRtnCancel 注册撤单响应
func (t *Trade) RegOnRtnCancel(on goctp.OnRtnOrderType) {
	t.onRtnCancel = on
}

// RegOnRtnTrade 注册成交响应
func (t *Trade) RegOnRtnTrade(on goctp.OnRtnTradeType) {
	t.onRtnTrade = on
}

func (t *Trade) RegOnRtnInstrumentStatus(on goctp.OnRtnInstrumentStatusType) {
	t.onRtnInstrumentStatus = on
}

// ********************** 底层接口响应处理 **********************************

// 合约状态响应
func (t *Trade) onStatus(statusField *ctp.CThostFtdcInstrumentStatusField) uintptr {
	status, loaded := t.InstrumentStatuss.LoadOrStore(goctp.Bytes2String(statusField.InstrumentID[:]), &goctp.InstrumentStatus{
		ExchangeID:       goctp.Bytes2String(statusField.ExchangeID[:]),
		InstrumentID:     goctp.Bytes2String(statusField.InstrumentID[:]),
		InstrumentStatus: goctp.InstrumentStatusType(statusField.InstrumentStatus),
		EnterTime:        goctp.Bytes2String(statusField.EnterTime[:]),
	})
	if loaded {
		status.(*goctp.InstrumentStatus).InstrumentStatus = goctp.InstrumentStatusType(statusField.InstrumentStatus)
		status.(*goctp.InstrumentStatus).EnterTime = goctp.Bytes2String(statusField.EnterTime[:])
	}

	if t.onRtnInstrumentStatus != nil {
		t.onRtnInstrumentStatus(status.(*goctp.InstrumentStatus))
	}
	return 0
}

// 成交响应
func (t *Trade) onTrade(tradeField *ctp.CThostFtdcTradeField) uintptr {
	key := fmt.Sprintf("%s_%c", tradeField.TradeID, tradeField.Direction)
	tf, _ := t.Trades.LoadOrStore(key, &goctp.TradeField{
		Direction:    goctp.DirectionType(tradeField.Direction),
		HedgeFlag:    goctp.HedgeFlagType(tradeField.HedgeFlag),
		InstrumentID: goctp.Bytes2String(tradeField.InstrumentID[:]),
		ExchangeID:   goctp.Bytes2String(tradeField.ExchangeID[:]),
		TradingDay:   goctp.Bytes2String(tradeField.TradingDay[:]),
		Volume:       int(tradeField.Volume),
		OffsetFlag:   goctp.OffsetFlagType(tradeField.OffsetFlag),
		OrderSysID:   goctp.Bytes2String(tradeField.OrderSysID[:]),
		Price:        float64(tradeField.Price),
		TradeDate:    goctp.Bytes2String(tradeField.TradeDate[:]),
		TradeTime:    goctp.Bytes2String(tradeField.TradeTime[:]),
		TradeID:      key,
	})
	var f = tf.(*goctp.TradeField)
	// 更新持仓
	if f.OffsetFlag == goctp.OffsetFlagOpen {
		var key string
		if f.Direction == goctp.DirectionBuy {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionLong, f.HedgeFlag)
		} else {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionShort, f.HedgeFlag)
		}
		pf, _ := t.Positions.LoadOrStore(key, &goctp.PositionField{
			InstrumentID:      f.InstrumentID,
			PositionDirection: goctp.PosiDirectionLong,
			HedgeFlag:         f.HedgeFlag,
			ExchangeID:        f.ExchangeID,
		})
		var p = pf.(*goctp.PositionField)
		p.OpenVolume += f.Volume
		p.OpenAmount += f.Price * float64(f.Volume)
		if info, ok := t.Instruments.Load(f.InstrumentID); ok {
			p.OpenCost += f.Price * float64(f.Volume) * float64(info.(*goctp.InstrumentField).VolumeMultiple)
		}
		p.Position += f.Volume
		p.TodayPosition += f.Volume
	} else {
		var key string
		if f.Direction == goctp.DirectionBuy {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionShort, f.HedgeFlag)
		} else {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionLong, f.HedgeFlag)
		}
		if posi, ok := t.Positions.Load(key); ok {
			var p = posi.(*goctp.PositionField)
			p.OpenVolume -= f.Volume
			p.OpenAmount -= f.Price * float64(f.Volume)
			if info, ok := t.Instruments.Load(f.InstrumentID); ok {
				p.OpenCost -= f.Price * float64(f.Volume) * float64(info.(*goctp.InstrumentField).VolumeMultiple)
			}
			p.Position -= f.Volume
			if f.OffsetFlag == goctp.OffsetFlagCloseToday {
				p.TodayPosition -= f.Volume
			} else {
				p.YdPosition -= f.Volume
			}
		}
	}
	// 处理对应的Order
	if ord, ok := t.sysID4Order.Load(f.OrderSysID); ok {
		var o = ord.(*goctp.OrderField)
		o.LastTradeTime = f.TradeTime
		o.VolumeTraded = f.Volume
		o.VolumeLeft -= f.Volume
		if o.VolumeLeft == 0 {
			o.OrderStatus = goctp.OrderStatusAllTraded
			o.StatusMsg = "全部成交"
		} else {
			o.OrderStatus = goctp.OrderStatusPartTradedQueueing
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
func (t *Trade) onOrder(orderField *ctp.CThostFtdcOrderField) uintptr {
	key := fmt.Sprintf("%d_%s", orderField.SessionID, orderField.OrderRef)
	if of, ok := t.Orders.LoadOrStore(key, &goctp.OrderField{
		InstrumentID:        goctp.Bytes2String(orderField.InstrumentID[:]),
		SessionID:           int(orderField.SessionID),
		FrontID:             int(orderField.FrontID),
		OrderRef:            goctp.Bytes2String(orderField.OrderRef[:]),
		Direction:           goctp.DirectionType(orderField.Direction),
		OffsetFlag:          goctp.OffsetFlagType(orderField.CombOffsetFlag[0]),
		HedgeFlag:           goctp.HedgeFlagType(orderField.CombHedgeFlag[0]),
		LimitPrice:          float64(orderField.LimitPrice),
		VolumeTotalOriginal: int(orderField.VolumeTotalOriginal),
		VolumeLeft:          int(orderField.VolumeTotalOriginal),
		ExchangeID:          goctp.Bytes2String(orderField.ExchangeID[:]),
		InsertDate:          goctp.Bytes2String(orderField.InsertDate[:]),
		InsertTime:          goctp.Bytes2String(orderField.InsertTime[:]),
		OrderStatus:         goctp.OrderStatusNoTradeQueueing, // OrderStatusType(orderField.OrderStatus)
		StatusMsg:           "委托已提交",                          // bytes2GBKbytes2GBKString(orderField.StatusMsg[:])
		IsLocal:             int(orderField.SessionID) == t.sessionID,
	}); !ok { // 新添加
		if t.onRtnOrder != nil {
			t.onRtnOrder(of.(*goctp.OrderField))
		}
	} else {
		var o = of.(*goctp.OrderField)
		if o.OrderStatus == goctp.OrderStatusCanceled {
			o.CancelTime = goctp.Bytes2String(orderField.CancelTime[:])
			// 错单
			if strings.Contains(o.StatusMsg, "被拒绝") {
				if t.onErrRtnOrder != nil {
					t.onErrRtnOrder(o, &goctp.RspInfoField{
						ErrorID:  -1,
						ErrorMsg: o.StatusMsg,
					})
				}
			} else if t.onRtnCancel != nil {
				t.onRtnCancel(o)
			}
		} else {
			o.OrderSysID = goctp.Bytes2String(orderField.OrderSysID[:])
			if len(o.OrderSysID) > 0 {
				t.sysID4Order.Store(o.OrderSysID, o)
			}
		}
	}
	return 0
}

// 委托错误响应
func (t *Trade) onErrOrder(orderField *ctp.CThostFtdcInputOrderField, infoField *ctp.CThostFtdcRspInfoField) uintptr {
	key := fmt.Sprintf("%d_%s", t.sessionID, orderField.OrderRef)
	of, _ := t.Orders.LoadOrStore(key, &goctp.OrderField{
		InstrumentID:        goctp.Bytes2String(orderField.InstrumentID[:]),
		SessionID:           t.sessionID,
		FrontID:             0,
		OrderRef:            goctp.Bytes2String(orderField.OrderRef[:]),
		Direction:           goctp.DirectionType(orderField.Direction),
		OffsetFlag:          goctp.OffsetFlagType(orderField.CombOffsetFlag[0]),
		HedgeFlag:           goctp.HedgeFlagType(orderField.CombHedgeFlag[0]),
		LimitPrice:          float64(orderField.LimitPrice),
		VolumeTotalOriginal: int(orderField.VolumeTotalOriginal),
		VolumeLeft:          int(orderField.VolumeTotalOriginal),
		ExchangeID:          goctp.Bytes2String(orderField.ExchangeID[:]),
		IsLocal:             true,
	})
	var o = of.(*goctp.OrderField)
	o.OrderStatus = goctp.OrderStatusCanceled
	if t.onErrRtnOrder != nil {
		t.onErrRtnOrder(o, &goctp.RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

// 撤单错误
func (t *Trade) onErrRtnOrderAction(field *ctp.CThostFtdcOrderActionField, infoField *ctp.CThostFtdcRspInfoField) uintptr {
	if t.onErrAction != nil {
		t.onErrAction(fmt.Sprintf("%d_%s", field.SessionID, field.OrderRef), &goctp.RspInfoField{
			ErrorID:  int(infoField.ErrorID),
			ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:]),
		})
	}
	return 0
}

// 持仓查询响应
func (t *Trade) onPosition(positionField *ctp.CThostFtdcInvestorPositionField, infoField *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	if strings.Compare(goctp.Bytes2String(positionField.InstrumentID[:]), "") != 0 {
		key := fmt.Sprintf("%s_%c_%c", goctp.Bytes2String(positionField.InstrumentID[:]), goctp.PosiDirectionType(positionField.PosiDirection), goctp.HedgeFlagType(positionField.HedgeFlag))
		pf, _ := t.Positions.LoadOrStore(key, &goctp.PositionField{
			InstrumentID:      goctp.Bytes2String(positionField.InstrumentID[:]),
			PositionDirection: goctp.PosiDirectionType(positionField.PosiDirection),
			HedgeFlag:         goctp.HedgeFlagType(positionField.HedgeFlag),
			ExchangeID:        goctp.Bytes2String(positionField.ExchangeID[:]),
		})
		var p = pf.(*goctp.PositionField)
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
func (t *Trade) onAccount(accountField *ctp.CThostFtdcTradingAccountField, infoField *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
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
	t.Account.CurrencyID = goctp.Bytes2String(accountField.CurrencyID[:])
	t.Account.PreFundMortgageIn = float64(accountField.PreFundMortgageIn)
	t.Account.PreFundMortgageOut = float64(accountField.PreFundMortgageOut)
	t.Account.FundMortgageIn = float64(accountField.FundMortgageIn)
	t.Account.FundMortgageOut = float64(accountField.FundMortgageOut)
	t.Account.FundMortgageAvailable = float64(accountField.FundMortgageAvailable)
	t.Account.MortgageableFund = float64(accountField.MortgageableFund)
	return 0
}

// 合约查询响应
func (t *Trade) onInstrument(instrumentField *ctp.CThostFtdcInstrumentField, infoField *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	if instrumentField != nil {
		t.Instruments.Store(goctp.Bytes2String(instrumentField.InstrumentID[:]), &goctp.InstrumentField{
			InstrumentID:              goctp.Bytes2String(instrumentField.InstrumentID[:]),
			ExchangeID:                goctp.Bytes2String(instrumentField.ExchangeID[:]),
			ProductID:                 goctp.Bytes2String(instrumentField.ProductID[:]),
			ProductClass:              goctp.ProductClassType(instrumentField.ProductClass),
			MaxMarketOrderVolume:      int(instrumentField.MaxMarketOrderVolume),
			MinMarketOrderVolume:      int(instrumentField.MinMarketOrderVolume),
			MaxLimitOrderVolume:       int(instrumentField.MaxLimitOrderVolume),
			MinLimitOrderVolume:       int(instrumentField.MinLimitOrderVolume),
			VolumeMultiple:            int(instrumentField.VolumeMultiple),
			PriceTick:                 float64(instrumentField.PriceTick),
			PositionType:              goctp.PositionTypeType(instrumentField.PositionType),
			UseMaxMarginSideAlgorithm: instrumentField.MaxMarginSideAlgorithm == '1',
			UnderlyingInstrID:         goctp.Bytes2String(instrumentField.UnderlyingInstrID[:]),
			StrikePrice:               float64(instrumentField.StrikePrice),
			OptionsType:               goctp.OptionsTypeType(instrumentField.OptionsType),
			UnderlyingMultiple:        float64(instrumentField.UnderlyingMultiple),
			CombinationType:           goctp.CombinationTypeType(instrumentField.CombinationType),
		})
	}
	if b {
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
	ordCnt := t.cntOrder
	for range t.qryTicker.C {
		if ordCnt == t.cntOrder {
			break
		}
		ordCnt = t.cntOrder
	}
	qryAccount := ctp.CThostFtdcQryTradingAccountField{}
	copy(qryAccount.InvestorID[:], t.InvestorID)
	copy(qryAccount.BrokerID[:], t.BrokerID)
	qryPosition := ctp.CThostFtdcQryInvestorPositionField{}
	copy(qryPosition.InvestorID[:], t.InvestorID)
	copy(qryPosition.BrokerID[:], t.BrokerID)
	// 启动查询
	bQryAccount := false
	sended := false
	for range t.qryTicker.C {
		if bQryAccount {
			t.t.ReqQryTradingAccount(qryAccount)
		} else {
			t.t.ReqQryInvestorPosition(qryPosition)
		}
		bQryAccount = !bQryAccount
		if !bQryAccount && !sended {
			sended = true
			// 通知:登录响应可以发了
			t.waitGroup.Done()
		}
		if !t.IsLogin {
			break
		}
	}
	t.qryTicker.Stop()
}

// 确认结算相应
func (t *Trade) onSettlement(confirmField *ctp.CThostFtdcSettlementInfoConfirmField, infoField *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.t.ReqQryInstrument(ctp.CThostFtdcQryInstrumentField{})
	return 0
}

// 登陆响应
func (t *Trade) onLogin(loginField *ctp.CThostFtdcRspUserLoginField, infoField *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	if infoField.ErrorID == 0 {
		t.sessionID = int(loginField.SessionID)
		t.TradingDay = goctp.Bytes2String(loginField.TradingDay[:])
		f := ctp.CThostFtdcSettlementInfoConfirmField{}
		copy(f.InvestorID[:], t.InvestorID)
		copy(f.AccountID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		t.t.ReqSettlementInfoConfirm(f)
		if t.onRspUserLogin != nil {
			t.waitGroup.Add(1)
			go func(field *goctp.RspUserLoginField) {
				t.waitGroup.Wait()
				t.onRspUserLogin(field, &goctp.RspInfoField{ErrorID: 0, ErrorMsg: "成功"})
			}(&goctp.RspUserLoginField{
				TradingDay:  t.TradingDay,
				LoginTime:   goctp.Bytes2String(loginField.LoginTime[:]),
				BrokerID:    t.BrokerID,
				UserID:      t.InvestorID,
				FrontID:     int(loginField.FrontID),
				SessionID:   t.sessionID,
				MaxOrderRef: goctp.Bytes2String(loginField.MaxOrderRef[:]),
			})
		}
	} else {
		t.onRspUserLogin(&goctp.RspUserLoginField{}, &goctp.RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

// 看穿式验证响应
func (t *Trade) onAuth(auth *ctp.CThostFtdcRspAuthenticateField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	if info.ErrorID == 0 {
		f := ctp.CThostFtdcReqUserLoginField{}
		copy(f.UserID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.Password[:], t.passWord)
		copy(f.UserProductInfo[:], "@HF")
		t.t.ReqUserLogin(f)
	} else if t.onRspUserLogin != nil {
		t.onRspUserLogin(&goctp.RspUserLoginField{}, &goctp.RspInfoField{ErrorID: int(info.ErrorID), ErrorMsg: goctp.Bytes2String(info.ErrorMsg[:])})
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
