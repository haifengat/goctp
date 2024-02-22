package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"gitee.com/haifengat/goctp/v2"
)

func main() {
	var trdFront, broker, user, pwd, appid, code, inst string
	var price float64
	var lot, period, times int

	flag.StringVar(&trdFront, "trdFront", "180.168.146.187:10202", "交易前置")
	flag.StringVar(&broker, "broker", "9999", "broker")
	flag.StringVar(&user, "user", "008107", "帐号")
	flag.StringVar(&pwd, "pwd", "1", "密码")
	flag.StringVar(&appid, "appID", "simnow_client_test", "appID")
	flag.StringVar(&code, "code", "0000000000000000", "authCode")
	flag.StringVar(&inst, "inst", "rb2410", "合约")
	flag.Float64Var(&price, "price", 0.0, "价格")
	flag.IntVar(&lot, "lot", 1, "手数")
	flag.IntVar(&period, "period", 40, "间隔(ms)")
	flag.IntVar(&times, "times", 1, "发单次数")
	flag.Parse()

	// 测试多帐号
	go trdTest(trdFront, broker, user, pwd, appid, code, inst, price, lot, period, times)
	// time.Sleep(time.Second * 20)
	// go trdTest("008107", "1")
	select {}
}

func trdTest(trdFront, broker, user, pwd, appid, code, instrument string, price float64, lot, period, times int) {
	trd := goctp.NewTradePro()
	trd.OnOrder = func(pOrder *goctp.CThostFtdcOrderField) {
		fmt.Printf("%s,OnOrder: %+v\n", trd.InvestorID, pOrder)
	}
	trd.OnTrade = func(pTrade *goctp.CThostFtdcTradeField) {
		fmt.Printf("%s,OnTrade: %+v\n", trd.InvestorID, pTrade)
	}
	trd.OnRtnInstrumentStatus = func(pInstrumentStatus *goctp.CThostFtdcInstrumentStatusField) {}

	trd.OnRspOrderAction = func(pInputOrderAction *goctp.CThostFtdcInputOrderActionField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspOrderAction %+v\n", pRspInfo)
	}
	trd.OnErrRtnOrderAction = func(pOrderAction *goctp.CThostFtdcOrderActionField, pRspInfo *goctp.CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnOrderAction %+v\n", pRspInfo)
	}
	trd.OnRspQryDepthMarketData = func(pDepthMarketData *goctp.CThostFtdcDepthMarketDataField, pRspInfo *goctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pDepthMarketData != nil && pDepthMarketData.InstrumentID.String() == "rb2310" {
			price = float64(pDepthMarketData.AskPrice1)
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
		// if bIsLast {
		// }
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
		Front:    "tcp://" + trdFront,
		Broker:   broker,
		UserID:   user,
		Password: pwd,
		AppID:    appid,
		AuthCode: code,
	})
	if rsp.ErrorID != 0 {
		fmt.Printf("%+v\n", rsp)
		os.Exit(-1)
	}
	fmt.Printf("%+v\n", info)

	if len(instrument) > 0 {
		for i := 0; i < times; i++ {
			_, rsp := trd.ReqOrderInsertLimit(instrument, goctp.THOST_FTDC_D_Sell, goctp.THOST_FTDC_OF_Open, price, lot)
			if rsp.ErrorID != 0 {
				fmt.Printf("%+v\n", rsp)
				return
			}
			time.Sleep(time.Millisecond * time.Duration(period))
		}
	}
	// trd.ReqOrderAction(id)
	//testIn
	// trd.ReqQryAccountregister()
	// time.Sleep(1 * time.Second)
	// trd.ReqFromBankToFutureByFuture(regInfo.BankAccount.String(), "bankPwd", "", 100)
	// select {}
}
