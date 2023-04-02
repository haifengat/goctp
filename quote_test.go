package goctp

import (
	"fmt"
	"testing"
)

func TestQuote(t *testing.T) {
	// var quoteFront = "tcp://180.168.146.187:10212"
	var quoteFront = "tcp://180.168.146.187:10131" // 7*24

	q := NewQuote()
	q.OnFrontConnected = func() {
		fmt.Println("quote connected")

		f := CThostFtdcReqUserLoginField{}
		copy(f.BrokerID[:], "9999")
		copy(f.UserID[:], "008105")
		copy(f.Password[:], "1")
		q.ReqUserLogin(&f, 1)
	}
	q.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pRspInfo.ErrorMsg)
		q.SubscribeMarketData([]string{"rb2303", "ru2304", "au2305", "rb2306"}, 0)

		fmt.Printf("%T", THOST_FTDC_D_Buy)
	}
	q.OnRtnDepthMarketData = func(pDepthMarketData *CThostFtdcDepthMarketDataField) {
		fmt.Printf("%s, %s, %s\n", pDepthMarketData.UpdateTime, pDepthMarketData.InstrumentID, pDepthMarketData.LastPrice)
	}
	q.RegisterFront(quoteFront)

	q.Init()

	select {}
}
