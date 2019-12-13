package go_ctp_lnx

// #include <stdlib.h>
import "C"
import (
	. "github.com/mayiweb/goctp" // 发布时使用，编码时可注释掉
	. "hf_go_ctp/common"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"unsafe"
)

type Quote struct {
	api              CThostFtdcMdApi
	onFrontConnected OnFrontConnectedType
	onRspUserLogin   OnRspUserLoginType
	onTick           OnTickType
	reqID            int
}

func NewQuote() *Quote {
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	logPath := path.Join(filepath.Dir(curFile), "log")
	_, err := os.Stat(logPath)
	if err != nil {
		_ = os.Mkdir(logPath, os.ModePerm)
	}
	q := new(Quote)
	//q.spi = NewCThostFtdcMdSpi()
	q.api = CThostFtdcMdApiCreateFtdcMdApi("./log/")
	q.api.RegisterSpi(NewDirectorCThostFtdcMdSpi(q))
	return q
}

func (q *Quote) getReqID() int {
	q.reqID++
	return q.reqID
}

func (q *Quote) Release() {
	q.api.Release()
}

func (q *Quote) ReqConnect(addr string) {
	q.api.RegisterFront(addr)
	q.api.Init()
}

func (q *Quote) ReqLogin(investor, pwd, broker string) {
	f := NewCThostFtdcReqUserLoginField()
	f.SetBrokerID(broker)
	f.SetUserID(investor)
	f.SetPassword(pwd)
	f.SetUserProductInfo("@haifeng")
	q.api.ReqUserLogin(f, q.getReqID())
}

func (q *Quote) ReqSubscript(instrument []string) {
	args := make([]*C.char, 0)
	for _, v := range instrument {
		char := C.CString(v)
		defer C.free(unsafe.Pointer(char))
		args = append(args, char)
	}
	q.api.SubscribeMarketData((*string)(unsafe.Pointer(&args[0])), 1)
}

func (q *Quote) RegOnFrontConnected(on OnFrontConnectedType) {
	q.onFrontConnected = on
}
func (q *Quote) RegOnRspUserLogin(on OnRspUserLoginType) {
	q.onRspUserLogin = on
}
func (q *Quote) RegOnTick(on OnTickType) {
	q.onTick = on
}

func (q *Quote) OnRtnDepthMarketData(dataField CThostFtdcDepthMarketDataField) {
	//fmt.Printf("%#v", q.onTick) // linux无响应
	if q.onTick == nil {
		return
	}
	tick := TickField{
		TradingDay:      dataField.GetTradingDay(),
		InstrumentID:    dataField.GetInstrumentID(),
		ExchangeID:      dataField.GetExchangeID(),
		LastPrice:       dataField.GetLastPrice(),
		OpenPrice:       dataField.GetOpenPrice(),
		HighestPrice:    dataField.GetHighestPrice(),
		LowestPrice:     dataField.GetLowestPrice(),
		Volume:          dataField.GetVolume(),
		Turnover:        dataField.GetTurnover(),
		OpenInterest:    dataField.GetOpenInterest(),
		SettlementPrice: dataField.GetSettlementPrice(),
		UpperLimitPrice: dataField.GetUpperLimitPrice(),
		LowerLimitPrice: dataField.GetLowerLimitPrice(),
		CurrDelta:       dataField.GetCurrDelta(),
		UpdateTime:      dataField.GetUpdateTime(),
		UpdateMillisec:  dataField.GetUpdateMillisec(),
		BidPrice1:       dataField.GetBidPrice1(),
		BidVolume1:      dataField.GetBidVolume1(),
		AskPrice1:       dataField.GetAskPrice1(),
		AskVolume1:      dataField.GetAskVolume1(),
		BidPrice2:       dataField.GetBidPrice2(),
		BidVolume2:      dataField.GetBidVolume2(),
		AskPrice2:       dataField.GetAskPrice2(),
		AskVolume2:      dataField.GetAskVolume2(),
		BidPrice3:       dataField.GetBidPrice3(),
		BidVolume3:      dataField.GetBidVolume3(),
		AskPrice3:       dataField.GetAskPrice3(),
		AskVolume3:      dataField.GetAskVolume3(),
		BidPrice4:       dataField.GetBidPrice4(),
		BidVolume4:      dataField.GetBidVolume4(),
		AskPrice4:       dataField.GetAskPrice4(),
		AskVolume4:      dataField.GetAskVolume4(),
		BidPrice5:       dataField.GetBidPrice5(),
		BidVolume5:      dataField.GetBidVolume5(),
		AskPrice5:       dataField.GetAskPrice5(),
		AskVolume5:      dataField.GetAskVolume5(),
		AveragePrice:    dataField.GetAskPrice5(),
		ActionDay:       dataField.GetActionDay(),
	}
	q.onTick(&tick)
}

func (q *Quote) OnRspUserLogin(loginField CThostFtdcRspUserLoginField, infoField CThostFtdcRspInfoField, i int, b bool) {
	if q.onRspUserLogin == nil {
		return
	}
	q.onRspUserLogin(&RspUserLoginField{
		TradingDay:  loginField.GetTradingDay(),
		LoginTime:   loginField.GetLoginTime(),
		BrokerID:    loginField.GetBrokerID(),
		UserID:      loginField.GetUserID(),
		FrontID:     loginField.GetFrontID(),
		SessionID:   loginField.GetSessionID(),
		MaxOrderRef: loginField.GetMaxOrderRef(),
	}, &RspInfoField{
		ErrorID:  infoField.GetErrorID(),
		ErrorMsg: infoField.GetErrorMsg(),
	})
}

func (q *Quote) OnFrontConnected() {
	if q.onFrontConnected == nil {
		return
	}
	q.onFrontConnected()
}
