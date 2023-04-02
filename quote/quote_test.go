package quote

import (
	"fmt"
	"testing"

	"gitee.com/haifengat/goctp/def"
)

func TestQuote(t *testing.T) {
	// var quoteFront = "tcp://180.168.146.187:10212"
	var quoteFront = "tcp://180.168.146.187:10131" // 7*24

	q := NewQuote()
	q.OnFrontConnected = func() {
		fmt.Println("quote connected")

		f := def.CThostFtdcReqUserLoginField{}
		copy(f.BrokerID[:], "9999")
		copy(f.UserID[:], "008105")
		copy(f.Password[:], "1")
		q.ReqUserLogin(&f, 1)
	}
	q.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pRspInfo.ErrorMsg)
		q.SubscribeMarketData([]string{"rb2303", "ru2304", "au2305", "rb2306"}, 0)

		fmt.Printf("%T", def.THOST_FTDC_D_Buy)
	}
	q.OnRtnDepthMarketData = func(pDepthMarketData *def.CThostFtdcDepthMarketDataField) {
		fmt.Printf("%s, %s, %s\n", pDepthMarketData.UpdateTime, pDepthMarketData.InstrumentID, pDepthMarketData.LastPrice)
	}
	q.RegisterFront(quoteFront)

	q.Init()

	select {}
}
