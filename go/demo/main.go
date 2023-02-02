package main

import "C"

import (
	"fmt"
	"goctp"
	"goctp/def"
	"goctp/quote"
)

func main() {
	q := quote.NewQuote()
	q.OnFrontConnected = func() {
		fmt.Println("quote connected")

		f := def.CThostFtdcReqUserLoginField{}
		copy(f.BrokerID[:], "9999")
		copy(f.UserID[:], "008105")
		copy(f.Password[:], "1")
		q.ReqUserLogin(&f, 1)
	}
	q.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", goctp.ToGBK(pRspInfo.ErrorMsg[:]))
		q.SubscribeMarketData([]string{"rb2302", "ru2302", "au2302", "rb2302"}, 1)
	}
	q.OnRtnDepthMarketData = func(pDepthMarketData *def.CThostFtdcDepthMarketDataField) {
		fmt.Printf("%s, %s, %s\n", pDepthMarketData.UpdateTime, pDepthMarketData.InstrumentID, goctp.FormatFloat(float64(pDepthMarketData.LastPrice), 3))
	}
	q.RegisterFront("tcp://180.168.146.187:10131")

	q.Init()

	select {}
}
