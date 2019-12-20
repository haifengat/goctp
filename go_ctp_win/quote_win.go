package go_ctp_win

import (
	"hf_go_ctp"
	"hf_go_ctp/go_ctp"
)

type Quote struct {
	q                *quote
	onFrontConnected hf_go_ctp.OnFrontConnectedType
	onRspUserLogin   hf_go_ctp.OnRspUserLoginType
	onTick           hf_go_ctp.OnTickType
}

func NewQuote() *Quote {
	q := new(Quote)
	q.q = newQuote()
	q.q.regOnFrontConnected(q.onConnected)
	q.q.regOnRspUserLogin(q.onUserLogin)
	q.q.regOnRtnDepthMarketData(q.onDepthMarketData)
	return q
}

func (q *Quote) Release() {
	q.q.Release()
}

func (q *Quote) ReqConnect(addr string) {
	q.q.RegisterFront(addr)
	q.q.Init()
}

func (q *Quote) ReqLogin(investor, pwd, broker string) {
	f := go_ctp.CThostFtdcReqUserLoginField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@haifeng")
	q.q.ReqUserLogin(f)
}

func (q *Quote) ReqSubscript(instrument string) {
	q.q.SubscribeMarketData([1][]byte{[]byte(instrument)}, 1)
}

func (q *Quote) RegOnFrontConnected(on hf_go_ctp.OnFrontConnectedType) {
	q.onFrontConnected = on
}
func (q *Quote) RegOnRspUserLogin(on hf_go_ctp.OnRspUserLoginType) {
	q.onRspUserLogin = on
}
func (q *Quote) RegOnTick(on hf_go_ctp.OnTickType) {
	q.onTick = on
}

func (q *Quote) onDepthMarketData(dataField *go_ctp.CThostFtdcDepthMarketDataField) uintptr {
	if q.onTick == nil {
		return 0
	}
	tick := hf_go_ctp.TickField{
		TradingDay:      hf_go_ctp.Bytes2String(dataField.TradingDay[:]),
		InstrumentID:    hf_go_ctp.Bytes2String(dataField.InstrumentID[:]),
		ExchangeID:      hf_go_ctp.Bytes2String(dataField.ExchangeID[:]),
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
		UpdateTime:      hf_go_ctp.Bytes2String(dataField.UpdateTime[:]),
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
		ActionDay:       hf_go_ctp.Bytes2String(dataField.ActionDay[:]),
	}
	q.onTick(&tick)
	return 0
}

func (q *Quote) onUserLogin(loginField *go_ctp.CThostFtdcRspUserLoginField, infoField *go_ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	if q.onRspUserLogin == nil {
		return 0
	}
	q.onRspUserLogin(&hf_go_ctp.RspUserLoginField{
		TradingDay:  string(loginField.TradingDay[:]),
		LoginTime:   string(loginField.LoginTime[:]),
		BrokerID:    string(loginField.BrokerID[:]),
		UserID:      string(loginField.UserID[:]),
		FrontID:     int(loginField.FrontID),
		SessionID:   int(loginField.SessionID),
		MaxOrderRef: string(loginField.MaxOrderRef[:]),
	}, &hf_go_ctp.RspInfoField{
		ErrorID:  int(infoField.ErrorID),
		ErrorMsg: hf_go_ctp.Bytes2String(infoField.ErrorMsg[:]),
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
