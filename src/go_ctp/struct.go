package go_ctp

type Direction uint8

const (
	Buy  Direction = 0
	Sell Direction = 1
)

// 响应信息
type RspInfoField struct {
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
}

// 用户登录应答
type RspUserLoginField struct {
	// 交易日
	TradingDay string
	// 登录成功时间
	LoginTime string
	// 经纪公司代码
	BrokerID string
	// 用户代码
	UserID string
	// 交易系统名称
	//SystemName string
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 最大报单引用
	MaxOrderRef string
	// 上期所时间
	//SHFETime string
	//// 大商所时间
	//DCETime string
	//// 郑商所时间
	//CZCETime string
	//// 中金所时间
	//FFEXTime string
	//// 能源中心时间
	//INETime string
}

// 行情响应
type TickField struct {
	// 交易日
	TradingDay string `json:TradingDay`
	// 合约代码
	InstrumentID string `json:InstrumentID`
	// 交易所代码
	ExchangeID string
	// 合约在交易所的代码
	//ExchangeInstID string
	// 最新价
	LastPrice float64
	// 上次结算价
	//PreSettlementPrice float64
	// 昨收盘
	//PreClosePrice float64
	// 昨持仓量
	//PreOpenInterest float64
	// 今开盘
	OpenPrice float64
	// 最高价
	HighestPrice float64
	// 最低价
	LowestPrice float64
	// 数量
	Volume int
	// 成交金额
	Turnover float64
	// 持仓量
	OpenInterest float64
	// 今收盘
	//ClosePrice float64
	// 本次结算价
	SettlementPrice float64
	// 涨停板价
	UpperLimitPrice float64
	// 跌停板价
	LowerLimitPrice float64
	// 昨虚实度
	//PreDelta float64
	// 今虚实度
	CurrDelta float64
	// 最后修改时间
	UpdateTime string
	// 最后修改毫秒
	UpdateMillisec int
	// 申买价一
	BidPrice1 float64
	// 申买量一
	BidVolume1 int
	// 申卖价一
	AskPrice1 float64
	// 申卖量一
	AskVolume1 int
	// 申买价二
	BidPrice2 float64
	// 申买量二
	BidVolume2 int
	// 申卖价二
	AskPrice2 float64
	// 申卖量二
	AskVolume2 int
	// 申买价三
	BidPrice3 float64
	// 申买量三
	BidVolume3 int
	// 申卖价三
	AskPrice3 float64
	// 申卖量三
	AskVolume3 int
	// 申买价四
	BidPrice4 float64
	// 申买量四
	BidVolume4 int
	// 申卖价四
	AskPrice4 float64
	// 申卖量四
	AskVolume4 int
	// 申买价五
	BidPrice5 float64
	// 申买量五
	BidVolume5 int
	// 申卖价五
	AskPrice5 float64
	// 申卖量五
	AskVolume5 int
	// 当日均价
	AveragePrice float64
	// 业务日期
	ActionDay string `json:ActionDay`
}
