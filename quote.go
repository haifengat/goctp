package goctp

import (
	"os"
	"sync"

	"gitee.com/haifengat/goctp/ctpdefine"
)

// HFQuote 行情接口
type HFQuote struct {
	IsLogin    bool
	InvestorID string
	BrokerID   string

	ReqConnect   ReqConnectType
	ReleaseAPI   ReleaseAPIType
	ReqUserLogin ReqUserLoginType
	ReqSubscript ReqSubscriptType

	onFrontConnected    OnFrontConnectedType
	onFrontDisConnected OnFrontDisConnectedType
	onRspUserLogin      OnRspUserLoginType
	onTick              OnTickType

	Ticks sync.Map // 合约:TickField
}

type ReqSubscriptType func(...string)

func (q *HFQuote) Init() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
}

func (q *HFQuote) Release() {
	q.IsLogin = false
	q.ReleaseAPI()
	q.FrontDisConnected(0) // 需手动触发
}

// ReqLogin 登录
func (q *HFQuote) ReqLogin(investor, pwd, broker string) {
	q.InvestorID = investor
	q.BrokerID = broker
	f := ctpdefine.CThostFtdcReqUserLoginField{}
	copy(f.UserID[:], q.InvestorID)
	copy(f.BrokerID[:], q.BrokerID)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@HF")
	q.ReqUserLogin(&f, 1)
}

// RegOnFrontConnected 注册前置响应
func (q *HFQuote) RegOnFrontConnected(on OnFrontConnectedType) {
	q.onFrontConnected = on
}

// RegOnFrontDisConnected 注册连接响应
func (q *HFQuote) RegOnFrontDisConnected(on OnFrontDisConnectedType) {
	q.onFrontDisConnected = on
}

// RegOnRspUserLogin 注册登录响应
func (q *HFQuote) RegOnRspUserLogin(on OnRspUserLoginType) {
	q.onRspUserLogin = on
}

// RegOnTick 注册行情响应
func (q *HFQuote) RegOnTick(on OnTickType) {
	q.onTick = on
}

func (q *HFQuote) RtnDepthMarketData(dataField *ctpdefine.CThostFtdcDepthMarketDataField) {
	tick := TickField{
		TradingDay:      Bytes2String(dataField.TradingDay[:]),
		InstrumentID:    Bytes2String(dataField.InstrumentID[:]),
		ExchangeID:      Bytes2String(dataField.ExchangeID[:]),
		LastPrice:       float64(dataField.LastPrice),
		OpenPrice:       float64(dataField.OpenPrice),
		HighestPrice:    float64(dataField.HighestPrice),
		LowestPrice:     float64(dataField.LowestPrice),
		Volume:          int(dataField.Volume),
		Turnover:        float64(dataField.Turnover),
		OpenInterest:    float64(dataField.OpenInterest),
		ClosePrice:      float64(dataField.ClosePrice),
		SettlementPrice: float64(dataField.SettlementPrice),
		UpperLimitPrice: float64(dataField.UpperLimitPrice),
		LowerLimitPrice: float64(dataField.LowerLimitPrice),
		CurrDelta:       float64(dataField.CurrDelta),
		UpdateTime:      Bytes2String(dataField.UpdateTime[:]),
		UpdateMillisec:  int(dataField.UpdateMillisec),
		BidPrice1:       float64(dataField.BidPrice1),
		BidVolume1:      int(dataField.BidVolume1),
		AskPrice1:       float64(dataField.AskPrice1),
		AskVolume1:      int(dataField.AskVolume1),
		BidPrice2:       float64(dataField.BidPrice2),
		BidVolume2:      int(dataField.BidVolume2),
		AskPrice2:       float64(dataField.AskPrice2),
		AskVolume2:      int(dataField.AskVolume2),
		BidPrice3:       float64(dataField.BidPrice3),
		BidVolume3:      int(dataField.BidVolume3),
		AskPrice3:       float64(dataField.AskPrice3),
		AskVolume3:      int(dataField.AskVolume3),
		BidPrice4:       float64(dataField.BidPrice4),
		BidVolume4:      int(dataField.BidVolume4),
		AskPrice4:       float64(dataField.AskPrice4),
		AskVolume4:      int(dataField.AskVolume4),
		BidPrice5:       float64(dataField.BidPrice5),
		BidVolume5:      int(dataField.BidVolume5),
		AskPrice5:       float64(dataField.AskPrice5),
		AskVolume5:      int(dataField.AskVolume5),
		AveragePrice:    float64(dataField.AveragePrice),
		ActionDay:       Bytes2String(dataField.ActionDay[:]),
	}
	q.Ticks.Store(tick.InstrumentID, &tick)
	if q.onTick == nil {
		return
	}
	q.onTick(&tick)
}

func (q *HFQuote) RspUserLogin(loginField *ctpdefine.CThostFtdcRspUserLoginField, infoField *ctpdefine.CThostFtdcRspInfoField) {
	q.IsLogin = infoField.ErrorID == 0

	if q.onRspUserLogin == nil {
		return
	}
	q.onRspUserLogin(&RspUserLoginField{
		TradingDay:  string(loginField.TradingDay[:]),
		LoginTime:   string(loginField.LoginTime[:]),
		BrokerID:    string(loginField.BrokerID[:]),
		UserID:      string(loginField.UserID[:]),
		FrontID:     int(loginField.FrontID),
		SessionID:   int(loginField.SessionID),
		MaxOrderRef: string(loginField.MaxOrderRef[:]),
	}, &RspInfoField{
		ErrorID:  int(infoField.ErrorID),
		ErrorMsg: Bytes2String(infoField.ErrorMsg[:]),
	})
}

func (q *HFQuote) FrontConnected() {
	if q.onFrontConnected != nil {
		q.onFrontConnected()
	}
}

func (q *HFQuote) FrontDisConnected(reason int) {
	if q.onFrontDisConnected != nil {
		q.onFrontDisConnected(reason)
	}
}
