package goctp

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestTradeExt(t *testing.T) {
	trd := NewTradeExt()
	eventChan := make(chan string)
	errorChan := make(chan CThostFtdcRspInfoField)
	var sysID string
	var lastPrice float64

	trd.OnFrontConnected = func() {
		eventChan <- "OnFrontConnected"
	}
	trd.OnRspAuthenticate = func(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			eventChan <- "OnRspAuthenticate"
		}
	}
	trd.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			fmt.Printf("%+v\n", pRspUserLogin)
			fmt.Printf("%+v\n", pRspInfo)
			eventChan <- "OnRspUserLogin"
		}
	}
	trd.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			eventChan <- "OnRspSettlementInfoConfirm"
		}
	}
	trd.OnRspQryInvestor = func(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInvestor != nil {
			fmt.Printf("OnRspQryInvestor %+v\n", pInvestor)
		}
		if bIsLast {
			eventChan <- "OnRspQryInvestor"
		}
	}
	trd.OnRspQryClassifiedInstrument = func(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInstrument != nil {
			fmt.Printf("OnRspQryClassifiedInstrument %+v\n", pInstrument)
		}
		if bIsLast {
			eventChan <- "OnRspQryClassifiedInstrument"
		}
	}
	trd.OnRspQryOrder = func(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pOrder != nil {
			fmt.Printf("OnRspQryOrder %+v\n", pOrder)
		}
		if bIsLast {
			eventChan <- "OnRspQryOrder"
		}
	}
	trd.OnRspQryTrade = func(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTrade != nil {
			fmt.Printf("OnRspQryTrade %+v\n", pTrade)
		}
		if bIsLast {
			eventChan <- "OnRspQryTrade"
		}
	}
	trd.OnRspQryInvestorPosition = func(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInvestorPosition != nil {
			fmt.Printf("OnRspQryInvestorPosition %+v\n", pInvestorPosition)
		}
		if bIsLast {
			eventChan <- "OnRspQryInvestorPosition"
		}
	}
	trd.OnRspQryInvestorPositionDetail = func(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInvestorPositionDetail != nil {
			fmt.Printf("OnRspQryInvestorPositionDetail %+v\n", pInvestorPositionDetail)
		}
		if bIsLast {
			eventChan <- "OnRspQryInvestorPositionDetail"
		}
	}
	trd.OnRspQryTradingAccount = func(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTradingAccount != nil {
			fmt.Printf("OnRspQryTradingAccount %+v\n", pTradingAccount)
		}
		if bIsLast {
			eventChan <- "OnRspQryTradingAccount"
		}
	}
	trd.OnRspOrderInsert = func(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInputOrder != nil {
			fmt.Printf("OnRspQryTradingAccount %+v\n", pInputOrder)
		}
		if bIsLast {
			eventChan <- "OnRspOrderInsert"
		}
	}
	trd.OnErrRtnOrderInsert = func(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			fmt.Printf("OnErrRtnOrderInsert %+v\n", pRspInfo)
		}
	}
	trd.OnRtnOrder = func(pOrder *CThostFtdcOrderField) {
		fmt.Printf("OnRtnOrder %s %+v\n", pOrder.InstrumentID, pOrder.LimitPrice)
		sysID = pOrder.OrderSysID.String()
		fmt.Println(pOrder.RequestID, "::", sysID, "::", pOrder.StatusMsg)
	}
	trd.OnRspOrderAction = func(pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspOrderAction %+v\n", pRspInfo)
	}
	trd.OnErrRtnOrderAction = func(pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnOrderAction %+v\n", pRspInfo)
	}
	trd.OnRspQryDepthMarketData = func(pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if string(bytes.TrimRight(pDepthMarketData.InstrumentID[:], "\x00")) == "rb2305" {
			lastPrice = float64(pDepthMarketData.AskPrice1)
		}
		// fmt.Println(pDepthMarketData.InstrumentID, "::", lastPrice)
	}
	// 银转相关 -----------------
	var regInfo CThostFtdcAccountregisterField
	trd.OnRspQryAccountregister = func(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pAccountregister != nil {
			regInfo = *pAccountregister
		}
		if bIsLast {
			eventChan <- "OnRspQryAccountregister"
		}
	}
	trd.OnRspQryTransferBank = func(pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTransferBank != nil {
			fmt.Printf("%+v\n", pTransferBank)
		}
		if bIsLast {
			eventChan <- "OnRspQryTransferBank"
		}
	}
	trd.OnErrRtnBankToFutureByFuture = func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnBankToFutureByFuture %+v\n", pRspInfo)
	}
	trd.OnRspFromBankToFutureByFuture = func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspFromBankToFutureByFuture %+v\n", pRspInfo)
	}
	trd.OnRtnFromBankToFutureByFuture = func(pRspTransfer *CThostFtdcRspTransferField) {
		fmt.Printf("%+v\n", pRspTransfer)
	}

	trd.SubscribePrivateTopic(THOST_TERT_QUICK)
	trd.SubscribePublicTopic(THOST_TERT_RESTART)
	trd.RegisterFront("tcp://180.168.146.187:10130")
	trd.Init()

	var testAction = func() {
		trd.ReqQryDepthMarketData("SHFE", "rb2305")
		time.Sleep(1 * time.Second)
		if lastPrice == 0 {
			return
		}
		// trd.ReqOrderInsert("rb2305", "SHFE", THOST_FTDC_D_Buy, THOST_FTDC_OF_Open, lastPrice, 3, trd.InvestorID)
		trd.ReqOrderInsert(THOST_FTDC_D_Sell, THOST_FTDC_OF_CloseToday, "rb2305", "SHFE", 3200+20, 3, trd.InvestorID, THOST_FTDC_OPT_LimitPrice, THOST_FTDC_TC_GFD, THOST_FTDC_VC_AV, THOST_FTDC_CC_Immediately)
		time.Sleep(1 * time.Second)
		trd.ReqOrderAction(struct {
			ExchangeID   string
			InstrumentID string
			InvestUnitID string
			OrderSysID   string
			OrderRef     string
			SessionID    int
			FrontID      int
		}{
			ExchangeID:   "SHFE",
			InstrumentID: "rb2305",
			OrderSysID:   sysID,
		})
	}

	var testIn = func() {
		trd.ReqQryAccountregister()
		time.Sleep(1 * time.Second)
		trd.ReqFromBankToFutureByFuture(regInfo, "bankpwd", "accountPwd", 100)
	}

	for {
		select {
		case cb := <-eventChan:
			switch cb {
			case "OnFrontConnected": // 连接
				trd.ReqAuthenticate("9999", "008107", "simnow_client_test", "0000000000000000")
			case "OnRspAuthenticate": // 认证
				trd.ReqUserLogin("1")
			case "OnRspUserLogin": // 登录
				trd.ReqSettlementInfoConfirm()
			case "OnRspSettlementInfoConfirm": // 确认结算
				trd.ReqQryInvestor()
			case "OnRspQryInvestor": // 查用户
				time.Sleep(time.Millisecond * 1100)
				// trd.ReqQryClassifiedInstrument()
				fmt.Println("登录过程完成")
				testAction() // 测试撤单
				// testIn() // 测试入金
			case "OnRspQryClassifiedInstrument": // 查合约
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryOrder()
			case "OnRspQryOrder": // 查委托
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryTrade()
			case "OnRspQryTrade": // 查成交
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryPosition()
			case "OnRspQryInvestorPosition": //查持仓
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryPositionDetail()
			case "OnRspQryInvestorPositionDetail": // 查持仓明细
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryTradingAccount()
			case "OnRspQryTradingAccount": // 查持仓后,发送委托
				fmt.Println("登录过程完成")
				testAction() // 测试撤单
				testIn()
			case "OnRspQryTransferBank":
			case "OnRspQryAccountregister":

			default:
				fmt.Println("未处理标识", cb)
			}
		case rspInfo := <-errorChan:
			fmt.Printf("%+v\n", rspInfo)
		}
	}
}
