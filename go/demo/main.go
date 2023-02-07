package main

import "C"

import (
	"fmt"
	"goctp"
	"goctp/def"
	"goctp/quote"
	"goctp/trade"
)

func main() {
	testT()
}

func testT() {
	t := trade.NewTrade()
	t.OnFrontConnected = func() {
		fmt.Println("trade connected")
		f := def.CThostFtdcReqAuthenticateField{}
		copy(f.BrokerID[:], "9999")
		copy(f.UserID[:], "008107")
		copy(f.AppID[:], "simnow_client_test")
		copy(f.AuthCode[:], "0000000000000000")
		t.ReqAuthenticate(&f, 1)
	}

	t.OnRspAuthenticate = func(pRspAuthenticateField *def.CThostFtdcRspAuthenticateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("看穿式: ") //, goctp.ToGBK(pRspInfo.ErrorMsg[:]))
		if pRspInfo.ErrorID == 0 {
			f := def.CThostFtdcReqUserLoginField{}
			copy(f.BrokerID[:], "9999")
			copy(f.UserID[:], "008107")
			copy(f.Password[:], "1")
			t.ReqUserLogin(&f, 2)
		}
	}

	t.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("登录: ", goctp.ToGBK(pRspInfo.ErrorMsg[:]))
		if pRspInfo.ErrorID == 0 {
			f := def.CThostFtdcSettlementInfoConfirmField{}
			copy(f.AccountID[:], "008107")
			copy(f.BrokerID[:], "9999")
			copy(f.InvestorID[:], "008107")
			t.ReqSettlementInfoConfirm(&f, 3)
		}
	}

	t.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("确认结算: ", goctp.ToGBK(pRspInfo.ErrorMsg[:]))
		f := def.CThostFtdcQryInvestorField{}
		copy(f.BrokerID[:], "9999")
		copy(f.InvestorID[:], "008107")
		t.ReqQryInvestor(&f, 4)
	}

	t.OnRspQryInvestor = func(pInvestor *def.CThostFtdcInvestorField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("姓名: ", pInvestor.InvestorName)
	}

	// t.RegisterFront("tcp://180.168.146.187:10130") // 不提供结算
	t.RegisterFront("tcp://180.168.146.187:10202")
	// t.RegisterFront("tcp://180.168.146.187:10201")
	t.SubscribePrivateTopic(def.THOST_TERT_RESTART)
	t.SubscribePublicTopic(def.THOST_TERT_RESTART)
	t.Init()
	select {}

	testQ()
}

func testQ() {
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
