package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"gitee.com/haifengat/goctp/v2"
)

func main() {
	// 测试多帐号
	go trdTest("008105", "1")
	time.Sleep(time.Second * 20)
	go trdTest("008107", "1")
	select {}
}

func trdTest(user, pwd string) {
	trd := goctp.NewTradePro()
	var lastPrice float64
	trd.OnOrder = func(pOrder *goctp.CThostFtdcOrderField) {
		fmt.Printf("%s,OnOrder\n", trd.InvestorID)
	}
	trd.OnTrade = func(pTrade *goctp.CThostFtdcTradeField) {
		fmt.Printf("%s,OnTrade\n", trd.InvestorID)
	}
	trd.OnRtnInstrumentStatus = func(pInstrumentStatus *goctp.CThostFtdcInstrumentStatusField) {}

	trd.OnRspOrderAction = func(pInputOrderAction *goctp.CThostFtdcInputOrderActionField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspOrderAction %+v\n", pRspInfo)
	}
	trd.OnErrRtnOrderAction = func(pOrderAction *goctp.CThostFtdcOrderActionField, pRspInfo *goctp.CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnOrderAction %+v\n", pRspInfo)
	}
	trd.OnRspQryDepthMarketData = func(pDepthMarketData *goctp.CThostFtdcDepthMarketDataField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pDepthMarketData != nil && string(bytes.TrimRight(pDepthMarketData.InstrumentID[:], "\x00")) == "rb2310" {
			lastPrice = float64(pDepthMarketData.AskPrice1)
		}
		// fmt.Println(pDepthMarketData.InstrumentID, "::", lastPrice)
	}
	// 银转相关 -----------------
	var regInfo goctp.CThostFtdcAccountregisterField
	var errorChan = make(chan goctp.CThostFtdcRspInfoField)
	trd.OnRspQryAccountregister = func(pAccountregister *goctp.CThostFtdcAccountregisterField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pAccountregister != nil {
			regInfo = *pAccountregister
		}
		if bIsLast {
			fmt.Println(regInfo)
		}
	}
	trd.OnRspQryTransferBank = func(pTransferBank *goctp.CThostFtdcTransferBankField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTransferBank != nil {
			fmt.Printf("%+v\n", pTransferBank)
		}
		if bIsLast {
		}
	}
	trd.OnErrRtnBankToFutureByFuture = func(pReqTransfer *goctp.CThostFtdcReqTransferField, pRspInfo *goctp.CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnBankToFutureByFuture %+v\n", pRspInfo)
	}
	trd.OnRspFromBankToFutureByFuture = func(pReqTransfer *goctp.CThostFtdcReqTransferField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspFromBankToFutureByFuture %+v\n", pRspInfo)
	}
	trd.OnRtnFromBankToFutureByFuture = func(pRspTransfer *goctp.CThostFtdcRspTransferField) {
		fmt.Printf("%+v\n", pRspTransfer)
	}

	info, rsp := trd.Start(goctp.LoginConfig{
		Front:    "tcp://180.168.146.187:10130",
		Broker:   "9999",
		UserID:   user,
		Password: pwd,
		AppID:    "",
		AuthCode: "",
	})
	if rsp.ErrorID != 0 {
		fmt.Printf("%+v\n", rsp)
		os.Exit(-1)
	}
	fmt.Printf("%+v\n", info)

	// testAction
	trd.ReqQryDepthMarketData("SHFE", "rb2310")
	time.Sleep(1 * time.Second)
	if lastPrice == 0 {
		fmt.Println("行情未订阅")
		return
	}
	id, rsp := trd.ReqOrderInsertLimit("rb2310", goctp.THOST_FTDC_D_Sell, goctp.THOST_FTDC_OF_Open, lastPrice+20, 3)
	if rsp.ErrorID != 0 {
		fmt.Printf("%+v\n", rsp)
		return
	}
	time.Sleep(1 * time.Second)
	trd.ReqOrderAction(id)
	//testIn
	// trd.ReqQryAccountregister()
	// time.Sleep(1 * time.Second)
	// trd.ReqFromBankToFutureByFuture(regInfo.BankAccount.String(), "bankPwd", "", 100)
	select {}
}
