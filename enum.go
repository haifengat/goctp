package goctp

// 产品类型类型
type ProductClassType byte

const (
	// 期货
	ProductClassFutures ProductClassType = '1'
	// 期货期权
	ProductClassOptions ProductClassType = '2'
	// 组合
	ProductClassCombinationProductClassType = '3'
	// 即期
	ProductClassSpot ProductClassType = '4'
	// 期转现
	ProductClassEFP ProductClassType = '5'
	// 现货期权
	ProductClassSpotOption ProductClassType = '6'
)

// 持仓类型类型
type PositionTypeType byte

const (
	// 净持仓
	PositionTypeNet PositionTypeType = '1'
	// 综合持仓
	PositionTypeGross PositionTypeType = '2'
)

// 期权类型
type OptionsTypeType byte

const (
	// 看涨
	OptionsTypeCallOptions OptionsTypeType = '1'
	// 看跌
	OptionsTypePutOptions OptionsTypeType = '2'
)

// 组合类型
type CombinationTypeType byte

const (
	// 期货组合
	CombinationTypeFuture CombinationTypeType = '0'
	// 垂直价差BUL
	CombinationTypeBUL CombinationTypeType = '1'
	// 垂直价差BER
	CombinationTypeBER CombinationTypeType = '2'
	// 跨式组合
	CombinationTypeSTD CombinationTypeType = '3'
	// 宽跨式组合
	CombinationTypeSTG CombinationTypeType = '4'
	// 备兑组合
	CombinationTypePRT CombinationTypeType = '5'
	// 时间价差组合
	CombinationTypeCLD CombinationTypeType = '6'
)

// 持仓多空方向类型
type PosiDirectionType byte

const (
	// 净
	PosiDirectionNet PosiDirectionType = '1'
	// 多头
	PosiDirectionLong PosiDirectionType = '2'
	// 空头
	PosiDirectionShort PosiDirectionType = '3'
)

// 投机套保标志类型
type HedgeFlagType byte

const (
	// 投机
	HedgeFlagSpeculation HedgeFlagType = '1'
	// 套利
	HedgeFlagArbitrage HedgeFlagType = '2'
	// 套保
	HedgeFlagHedge HedgeFlagType = '3'
	// 做市商
	HedgeFlagMarketMaker HedgeFlagType = '5'
	// 第一腿投机第二腿套保 大商所专用
	HedgeFlagSpecHedge HedgeFlagType = '6'
	// 第一腿套保第二腿投机  大商所专用
	HedgeFlagHedgeSpec HedgeFlagType = '7'
)

// 买卖方向类型
type DirectionType byte

const (
	// 买
	DirectionBuy DirectionType = '0'
	// 卖
	DirectionSell DirectionType = '1'
)

// 开平标志类型
type OffsetFlagType byte

const (
	// 开仓
	OffsetFlagOpen OffsetFlagType = '0'
	// 平仓
	OffsetFlagClose OffsetFlagType = '1'
	// 强平
	OffsetFlagForceClose OffsetFlagType = '2'
	// 平今
	OffsetFlagCloseToday OffsetFlagType = '3'
	// 平昨
	OffsetFlagCloseYesterday OffsetFlagType = '4'
	// 强减
	OffsetFlagForceOff OffsetFlagType = '5'
	// 本地强平
	OffsetFlagLocalForceClose OffsetFlagType = '6'
)

// 报单状态类型
type OrderStatusType byte

const (
	// 全部成交
	OrderStatusAllTraded OrderStatusType = '0'
	// 部分成交还在队列中
	OrderStatusPartTradedQueueing OrderStatusType = '1'
	// 部分成交不在队列中
	OrderStatusPartTradedNotQueueing OrderStatusType = '2'
	// 未成交还在队列中
	OrderStatusNoTradeQueueing OrderStatusType = '3'
	// 未成交不在队列中
	OrderStatusNoTradeNotQueueing OrderStatusType = '4'
	// 撤单
	OrderStatusCanceled OrderStatusType = '5'
	// 未知
	OrderStatusUnknown OrderStatusType = 'a'
	// 尚未触发
	OrderStatusNotTouched OrderStatusType = 'b'
	// 已触发
	OrderStatusTouched OrderStatusType = 'c'
)

// 合约交易状态类型
type InstrumentStatusType byte

const (
	// 开盘前
	InstrumentStatusBeforeTrading InstrumentStatusType = '0'
	// 非交易
	InstrumentStatusNoTrading InstrumentStatusType = '1'
	// 连续交易
	InstrumentStatusContinous InstrumentStatusType = '2'
	// 集合竞价报单
	InstrumentStatusAuctionOrdering InstrumentStatusType = '3'
	// 集合竞价价格平衡
	InstrumentStatusAuctionBalance InstrumentStatusType = '4'
	// 集合竞价撮合
	InstrumentStatusAuctionMatch InstrumentStatusType = '5'
	// 收盘
	InstrumentStatusClosed InstrumentStatusType = '6'
)
