package goctp

import (
	"fmt"
	"testing"
)

func TestTrade(test *testing.T) {
	t := NewTrade()
	t.OnFrontConnected = func() {
		fmt.Println("trade connected")
		f := CThostFtdcReqAuthenticateField{}
		copy(f.BrokerID[:], "9999")
		copy(f.UserID[:], "008107")
		copy(f.AppID[:], "simnow_client_test")
		copy(f.AuthCode[:], "0000000000000000")
		t.ReqAuthenticate(&f, 1)
	}

	t.OnRspAuthenticate = func(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("看穿式: ", pRspInfo.ErrorMsg)
		if pRspInfo.ErrorID == 0 {
			f := CThostFtdcReqUserLoginField{}
			copy(f.BrokerID[:], "9999")
			copy(f.UserID[:], "008107")
			copy(f.Password[:], "1")
			copy(f.UserProductInfo[:], "@HF")
			t.ReqUserLogin(&f, 2)
		}
	}

	t.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("登录: ", pRspInfo.ErrorMsg)
		if pRspInfo.ErrorID == 0 {
			f := CThostFtdcSettlementInfoConfirmField{}
			copy(f.AccountID[:], "008107")
			copy(f.BrokerID[:], "9999")
			copy(f.InvestorID[:], "008107")
			t.ReqSettlementInfoConfirm(&f, 3)
		}
	}

	t.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("确认结算: ", pRspInfo.ErrorMsg)
		f := CThostFtdcQryInvestorField{}
		copy(f.BrokerID[:], "9999")
		copy(f.InvestorID[:], "008107")
		t.ReqQryInvestor(&f, 4)
	}

	t.OnRtnInstrumentStatus = func(pInstrumentStatus *CThostFtdcInstrumentStatusField) {}

	t.OnRspQryInvestor = func(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("姓名: ", pInvestor.InvestorName)
	}
	// var tradeFront = "tcp://180.168.146.187:10202"
	var tradeFront = "tcp://180.168.146.187:10130" // 7*24
	t.RegisterFront(tradeFront)
	t.SubscribePrivateTopic(THOST_TERT_RESTART)
	t.SubscribePublicTopic(THOST_TERT_RESTART)
	t.Init()
	select {}
}
