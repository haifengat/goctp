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
		if pRspInfo.ErrorID != 0 {
			fmt.Printf("%+v\n", pRspInfo.ErrorID)
		} else {
			fmt.Println("ReqUserLogin")
			trd.ReqUserLogin("1")
		}
	}
	trd.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			fmt.Printf("%+v\n", pRspInfo.ErrorID)
		} else {
			fmt.Printf("%+v\n", pRspUserLogin)
			fmt.Printf("%+v\n", pRspInfo)
			trd.ReqSettlementInfoConfirm()
		}
	}
	trd.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			fmt.Printf("%+v\n", pRspInfo.ErrorID)
		} else {
			fmt.Println("ReqQryInvestor")
			trd.ReqQryInvestor()
		}
	}
	trd.OnRspQryInvestor = func(pInvestor *def.CThostFtdcInvestorField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pInvestor)
		if bIsLast {
			time.Sleep(time.Millisecond * 1100)
			fmt.Println("ReqQryClassifiedInstrument")
			trd.ReqQryClassifiedInstrument()
		}
	}
	trd.OnRspQryClassifiedInstrument = func(pInstrument *def.CThostFtdcInstrumentField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pInstrument)
		if bIsLast {
			time.Sleep(time.Millisecond * 1100)
			fmt.Println("ReqQryOrder")
			trd.ReqQryOrder()
		}
	}
	trd.OnRspQryOrder = func(pOrder *def.CThostFtdcOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pOrder != nil {
			fmt.Printf("%+v\n", pOrder)
		}
		if bIsLast {
			time.Sleep(time.Millisecond * 1100)
			fmt.Println("ReqQryTrade")
			trd.ReqQryTrade()
		}
	}
	trd.OnRspQryTrade = func(pTrade *def.CThostFtdcTradeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pTrade)
		if bIsLast {
			time.Sleep(time.Millisecond * 1100)
			fmt.Println("ReqQryPosition")
			trd.ReqQryPosition()
		}
	}
	trd.OnRspQryInvestorPosition = func(pInvestorPosition *def.CThostFtdcInvestorPositionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pInvestorPosition)
		if bIsLast {
			fmt.Println("onrsp qry position")
			time.Sleep(time.Millisecond * 1100)
			fmt.Println("ReqQryPositionDetail")
			trd.ReqQryPositionDetail()
		}
	}
	trd.OnRspQryInvestorPositionDetail = func(pInvestorPositionDetail *def.CThostFtdcInvestorPositionDetailField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pInvestorPositionDetail)
		if bIsLast {
			fmt.Println("ReqQryPositionDetail")
			time.Sleep(time.Millisecond * 1100)
			fmt.Println("ReqQryAccount")
			trd.ReqQryAccount()
		}
	}
	trd.OnRspQryTradingAccount = func(pTradingAccount *def.CThostFtdcTradingAccountField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("%+v\n", pTradingAccount)
		if bIsLast {
			fmt.Println("OnRspQryTradingAccount")
			time.Sleep(time.Millisecond * 1100)
			trd.ReqOrderInsert("rb2305", "SHFE", def.THOST_FTDC_D_Buy, def.THOST_FTDC_OF_Open, 2300, 3, trd.UserID)
		}
	}
	trd.OnRspOrderInsert = func(pInputOrder *def.CThostFtdcInputOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspOrderInsert %+v\n", pRspInfo)
	}
	trd.OnErrRtnOrderInsert = func(pInputOrder *def.CThostFtdcInputOrderField, pRspInfo *def.CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnOrderInsert %+v\n", pRspInfo)
	}
	trd.OnRtnOrder = func(pOrder *def.CThostFtdcOrderField) {
		fmt.Printf("%+v\n", pOrder)
		fmt.Println(pOrder.RequestID, "::", pOrder.StatusMsg)
	}
	trd.Subscribe(def.THOST_TERT_QUICK, def.THOST_TERT_RESTART)
	trd.ReqConnect("tcp://180.168.146.187:10130")

	select {}
}
