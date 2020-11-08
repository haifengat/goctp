package win

import (
	"github.com/haifengat/goctp"
	ctp "github.com/haifengat/goctp/ctpdefine"
)

// Quote md class
type Quote struct {
	q                *quote
	onFrontConnected goctp.OnFrontConnectedType
	onRspUserLogin   goctp.OnRspUserLoginType
	onTick           goctp.OnTickType
}

// NewQuote new md api instanse
func NewQuote() *Quote {
	q := new(Quote)
	q.q = newQuote()
	q.q.regOnFrontConnected(q.onConnected)
	q.q.regOnRspUserLogin(q.onUserLogin)
	q.q.regOnRtnDepthMarketData(q.onDepthMarketData)
	return q
}

// Release release
func (q *Quote) Release() {
	q.q.Release()
}

// ReqConnect
func (q *Quote) ReqConnect(addr string) {
	q.q.RegisterFront(addr)
	q.q.Init()
}

func (q *Quote) ReqLogin(investor, pwd, broker string) {
	f := ctp.CThostFtdcReqUserLoginField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@haifeng")
	q.q.ReqUserLogin(f)
}

func (q *Quote) ReqSubscript(instrument string) {
	q.q.SubscribeMarketData([1][]byte{[]byte(instrument)}, 1)
}

func (q *Quote) RegOnFrontConnected(on goctp.OnFrontConnectedType) {
	q.onFrontConnected = on
}
func (q *Quote) RegOnRspUserLogin(on goctp.OnRspUserLoginType) {
	q.onRspUserLogin = on
}
func (q *Quote) RegOnTick(on goctp.OnTickType) {
	q.onTick = on
}

func (q *Quote) onDepthMarketData(dataField *ctp.CThostFtdcDepthMarketDataField) uintptr {
	if q.onTick == nil {
		return 0
	}
	tick := goctp.TickField{
		TradingDay:      goctp.Bytes2String(dataField.TradingDay[:]),
		InstrumentID:    goctp.Bytes2String(dataField.InstrumentID[:]),
		ExchangeID:      goctp.Bytes2String(dataField.ExchangeID[:]),
		LastPrice:       float64(dataField.LastPrice),
		OpenPrice:       float64(dataField.OpenPrice),
		HighestPrice:    float64(dataField.HighestPrice),
		LowestPrice:     float64(dataField.LowestPrice),
		Volume:          int(dataField.Volume),
		Turnover:        float64(dataField.Turnover),
		OpenInterest:    float64(dataField.OpenInterest),
		SettlementPrice: float64(dataField.SettlementPrice),
		UpperLimitPrice: float64(dataField.UpperLimitPrice),
		LowerLimitPrice: float64(dataField.LowerLimitPrice),
		CurrDelta:       float64(dataField.CurrDelta),
		UpdateTime:      goctp.Bytes2String(dataField.UpdateTime[:]),
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
		AveragePrice:    float64(dataField.AskPrice5),
		ActionDay:       goctp.Bytes2String(dataField.ActionDay[:]),
	}
	q.onTick(&tick)
	return 0
}

func (q *Quote) onUserLogin(loginField *ctp.CThostFtdcRspUserLoginField, infoField *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	if q.onRspUserLogin == nil {
		return 0
	}
	q.onRspUserLogin(&goctp.RspUserLoginField{
		TradingDay:  string(loginField.TradingDay[:]),
		LoginTime:   string(loginField.LoginTime[:]),
		BrokerID:    string(loginField.BrokerID[:]),
		UserID:      string(loginField.UserID[:]),
		FrontID:     int(loginField.FrontID),
		SessionID:   int(loginField.SessionID),
		MaxOrderRef: string(loginField.MaxOrderRef[:]),
	}, &goctp.RspInfoField{
		ErrorID:  int(infoField.ErrorID),
		ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:]),
	})
	return 0
}

func (q *Quote) onConnected() uintptr {
	if q.onFrontConnected == nil {
		return 0
	}
	q.onFrontConnected()
	return 0
}
