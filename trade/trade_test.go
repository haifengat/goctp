package trade

import (
	"fmt"
	"testing"

	"gitee.com/haifengat/goctp/def"
)

func TestTrade(test *testing.T) {
	t := NewTrade()
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
		fmt.Println("看穿式: ", pRspInfo.ErrorMsg)
		if pRspInfo.ErrorID == 0 {
			f := def.CThostFtdcReqUserLoginField{}
			copy(f.BrokerID[:], "9999")
			copy(f.UserID[:], "008107")
			copy(f.Password[:], "1")
			copy(f.UserProductInfo[:], "@HF")
			t.ReqUserLogin(&f, 2)
		}
	}

	t.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("登录: ", pRspInfo.ErrorMsg)
		if pRspInfo.ErrorID == 0 {
			f := def.CThostFtdcSettlementInfoConfirmField{}
			copy(f.AccountID[:], "008107")
			copy(f.BrokerID[:], "9999")
			copy(f.InvestorID[:], "008107")
			t.ReqSettlementInfoConfirm(&f, 3)
		}
	}

	t.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("确认结算: ", pRspInfo.ErrorMsg)
		f := def.CThostFtdcQryInvestorField{}
		copy(f.BrokerID[:], "9999")
		copy(f.InvestorID[:], "008107")
		t.ReqQryInvestor(&f, 4)
	}

	t.OnRtnInstrumentStatus = func(pInstrumentStatus *def.CThostFtdcInstrumentStatusField) {}

	t.OnRspQryInvestor = func(pInvestor *def.CThostFtdcInvestorField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println("姓名: ", pInvestor.InvestorName)
	}
	// var tradeFront = "tcp://180.168.146.187:10202"
	var tradeFront = "tcp://180.168.146.187:10130" // 7*24
	t.RegisterFront(tradeFront)
	t.SubscribePrivateTopic(def.THOST_TERT_RESTART)
	t.SubscribePublicTopic(def.THOST_TERT_RESTART)
	t.Init()
	select {}
}
