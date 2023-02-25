package trade

import (
	"fmt"
	"goctp/def"
	"testing"
	"time"
)

func TestTradeExt(t *testing.T) {
	trd := NewTradeExt()
	trd.OnFrontConnected = func() {
		trd.ReqAuthenticateField("9999", "008107", "simnow_client_test", "0000000000000000")
	}
	trd.OnRspAuthenticate = func(pRspAuthenticateField *def.CThostFtdcRspAuthenticateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		trd.ReqUserLogin("1")
	}
	trd.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Println(pRspInfo.ErrorMsg)
		trd.ReqQryInvestor()
	}
	trd.OnRspQryInvestor = func(pInvestor *def.CThostFtdcInvestorField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%s, %s, %s, %s\n", pInvestor.InvestorID, pInvestor.InvestorName, pInvestor.Address, pInvestor.Mobile)
		if bIsLast {
			time.Sleep(time.Millisecond * 1100)
			trd.ReqQryOrder()
		}
	}
	trd.OnRspQryOrder = func(pOrder *def.CThostFtdcOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pOrder != nil {
			fmt.Println(pOrder.OrderRef)
		}
	}
	trd.Subscribe(def.THOST_TERT_QUICK, def.THOST_TERT_RESTART)
	trd.ReqConnect("tcp://180.168.146.187:10130")

	select {}
}
