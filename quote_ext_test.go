package goctp

import (
	"fmt"
	"testing"
	"time"
)

func TestQuoteExt(t *testing.T) {
	for { // 测试多次连接
		q := NewQuoteExt()
		q.OnFrontConnected = func() {
			fmt.Println("quote connected")
			q.ReqUserLogin("9999", "008105", "1")
		}
		q.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
			fmt.Println(pRspInfo.ErrorMsg)
			q.ReqSubscript("rb2303", "ru2305")
		}
		q.OnRtnDepthMarketData = func(pDepthMarketData *CThostFtdcDepthMarketDataField) {
			fmt.Println(pDepthMarketData.InstrumentID, pDepthMarketData.LastPrice)
		}
		q.OnFrontDisconnected = func(nReason int) {
			fmt.Println("quote disconnect: ", nReason)
		}

		// q.RegisterFront("tcp://180.168.146.187:10212")
		q.RegisterFront("tcp://180.168.146.187:10131")
		q.Init()

		time.Sleep(6 * time.Second)
		q.Release()
	}
}
