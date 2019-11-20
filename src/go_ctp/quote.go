package go_ctp

type OnFrontConnectedType func()
type OnRspUserLoginType func(data RspUserLoginField, info RspInfoField)
type OnTick func(tick TickField)

type Quote struct {
	q                *quote
	onFrontConnected OnFrontConnectedType
	onRspUserLogin   OnRspUserLoginType
	onTick           OnTick
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
	f := CThostFtdcReqUserLoginField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@haifeng")
	q.q.ReqUserLogin(f)
}

func (q *Quote) ReqSubscript(instrument string) {
	q.q.SubscribeMarketData([1][]byte{[]byte(instrument)}, 1)
}

func (q *Quote) RegOnFrontConnected(on OnFrontConnectedType) {
	q.onFrontConnected = on
}
func (q *Quote) RegOnRspUserLogin(on OnRspUserLoginType) {
	q.onRspUserLogin = on
}
func (q *Quote) RegOnTick(on OnTick) {
	q.onTick = on
}

func (q *Quote) onDepthMarketData(dataField *CThostFtdcDepthMarketDataField) uintptr {
	if q.onTick != nil {
		tick := TickField{
			bytes2String(dataField.TradingDay[:]),
			bytes2String(dataField.InstrumentID[:]),
			bytes2String(dataField.ExchangeID[:]),
			float64(dataField.LastPrice),
			float64(dataField.OpenPrice),
			float64(dataField.HighestPrice),
			float64(dataField.LowestPrice),
			int(dataField.Volume),
			float64(dataField.Turnover),
			float64(dataField.OpenInterest),
			float64(dataField.SettlementPrice),
			float64(dataField.UpperLimitPrice),
			float64(dataField.LowerLimitPrice),
			float64(dataField.CurrDelta),
			string(dataField.UpdateTime[:]),
			int(dataField.UpdateMillisec),
			float64(dataField.BidPrice1),
			int(dataField.BidVolume1),
			float64(dataField.AskPrice1),
			int(dataField.AskVolume1),
			float64(dataField.BidPrice2),
			int(dataField.BidVolume2),
			float64(dataField.AskPrice2),
			int(dataField.AskVolume2),
			float64(dataField.BidPrice3),
			int(dataField.BidVolume3),
			float64(dataField.AskPrice3),
			int(dataField.AskVolume3),
			float64(dataField.BidPrice4),
			int(dataField.BidVolume4),
			float64(dataField.AskPrice4),
			int(dataField.AskVolume4),
			float64(dataField.BidPrice5),
			int(dataField.BidVolume5),
			float64(dataField.AskPrice5),
			int(dataField.AskVolume5),
			float64(dataField.AskPrice5),
			bytes2String(dataField.ActionDay[:]),
		}
		q.onTick(tick)
	}
	return 0
}

func (q *Quote) onUserLogin(loginField *CThostFtdcRspUserLoginField, infoField *CThostFtdcRspInfoField, i int, b bool) uintptr {
	if q.onRspUserLogin != nil {
		f := RspUserLoginField{
			string(loginField.TradingDay[:]),
			string(loginField.LoginTime[:]),
			string(loginField.BrokerID[:]),
			string(loginField.UserID[:]),
			//string(loginField.SystemName[:]),
			int(loginField.FrontID),
			int(loginField.SessionID),
			string(loginField.MaxOrderRef[:]),
			//string(loginField.SHFETime[:]),
			//string(loginField.DCETime[:]),
			//string(loginField.CZCETime[:]),
			//string(loginField.FFEXTime[:]),
			//string(loginField.INETime[:]),
		}
		i := RspInfoField{
			int(infoField.ErrorID),
			string(infoField.ErrorMsg[:]),
		}
		q.onRspUserLogin(f, i)
	}
	return 0
}

func (q *Quote) onConnected() uintptr {
	if q.onFrontConnected != nil {
		q.onFrontConnected()
	}
	return 0
}
