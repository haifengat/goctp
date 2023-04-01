package trade

import (
	"bytes"
	"fmt"
	"goctp/def"
	"testing"
	"time"
)

func TestTradeExt(t *testing.T) {
	trd := NewTradeExt()
	eventChan := make(chan string)
	errorChan := make(chan def.CThostFtdcRspInfoField)
	var sysID string
	var lastPrice float64
	// var custType def.TThostFtdcCustTypeType

	trd.OnFrontConnected = func() {
		eventChan <- "OnFrontConnected"
	}
	trd.OnRspAuthenticate = func(pRspAuthenticateField *def.CThostFtdcRspAuthenticateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			eventChan <- "OnRspAuthenticate"
		}
	}
	trd.OnRspUserLogin = func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			fmt.Printf("%+v\n", pRspUserLogin)
			fmt.Printf("%+v\n", pRspInfo)
			eventChan <- "OnRspUserLogin"
		}
	}
	trd.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			eventChan <- "OnRspSettlementInfoConfirm"
		}
	}
	trd.OnRspQryInvestor = func(pInvestor *def.CThostFtdcInvestorField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInvestor != nil {
			fmt.Printf("OnRspQryInvestor %+v\n", pInvestor)
		}
		if bIsLast {
			eventChan <- "OnRspQryInvestor"
		}
	}
	trd.OnRspQryClassifiedInstrument = func(pInstrument *def.CThostFtdcInstrumentField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInstrument != nil {
			fmt.Printf("OnRspQryClassifiedInstrument %+v\n", pInstrument)
		}
		if bIsLast {
			eventChan <- "OnRspQryClassifiedInstrument"
		}
	}
	trd.OnRspQryOrder = func(pOrder *def.CThostFtdcOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pOrder != nil {
			fmt.Printf("OnRspQryOrder %+v\n", pOrder)
		}
		if bIsLast {
			eventChan <- "OnRspQryOrder"
		}
	}
	trd.OnRspQryTrade = func(pTrade *def.CThostFtdcTradeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTrade != nil {
			fmt.Printf("OnRspQryTrade %+v\n", pTrade)
		}
		if bIsLast {
			eventChan <- "OnRspQryTrade"
		}
	}
	trd.OnRspQryInvestorPosition = func(pInvestorPosition *def.CThostFtdcInvestorPositionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInvestorPosition != nil {
			fmt.Printf("OnRspQryInvestorPosition %+v\n", pInvestorPosition)
		}
		if bIsLast {
			eventChan <- "OnRspQryInvestorPosition"
		}
	}
	trd.OnRspQryInvestorPositionDetail = func(pInvestorPositionDetail *def.CThostFtdcInvestorPositionDetailField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInvestorPositionDetail != nil {
			fmt.Printf("OnRspQryInvestorPositionDetail %+v\n", pInvestorPositionDetail)
		}
		if bIsLast {
			eventChan <- "OnRspQryInvestorPositionDetail"
		}
	}
	trd.OnRspQryTradingAccount = func(pTradingAccount *def.CThostFtdcTradingAccountField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTradingAccount != nil {
			fmt.Printf("OnRspQryTradingAccount %+v\n", pTradingAccount)
		}
		if bIsLast {
			eventChan <- "OnRspQryTradingAccount"
		}
	}
	trd.OnRspOrderInsert = func(pInputOrder *def.CThostFtdcInputOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pInputOrder != nil {
			fmt.Printf("OnRspQryTradingAccount %+v\n", pInputOrder)
		}
		if bIsLast {
			eventChan <- "OnRspOrderInsert"
		}
	}
	trd.OnErrRtnOrderInsert = func(pInputOrder *def.CThostFtdcInputOrderField, pRspInfo *def.CThostFtdcRspInfoField) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else {
			fmt.Printf("OnErrRtnOrderInsert %+v\n", pRspInfo)
		}
	}
	trd.OnRtnOrder = func(pOrder *def.CThostFtdcOrderField) {
		fmt.Printf("OnRtnOrder %s %+v\n", pOrder.InstrumentID, pOrder.LimitPrice)
		sysID = pOrder.OrderSysID.String()
		fmt.Println(pOrder.RequestID, "::", sysID, "::", pOrder.StatusMsg)
	}
	trd.OnRspOrderAction = func(pInputOrderAction *def.CThostFtdcInputOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspOrderAction %+v\n", pRspInfo)
	}
	trd.OnErrRtnOrderAction = func(pOrderAction *def.CThostFtdcOrderActionField, pRspInfo *def.CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnOrderAction %+v\n", pRspInfo)
	}
	trd.OnRspQryDepthMarketData = func(pDepthMarketData *def.CThostFtdcDepthMarketDataField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if string(bytes.TrimRight(pDepthMarketData.InstrumentID[:], "\x00")) == "rb2305" {
			lastPrice = float64(pDepthMarketData.AskPrice1)
		}
		// fmt.Println(pDepthMarketData.InstrumentID, "::", lastPrice)
	}
	// 银转相关 -----------------
	var regInfo def.CThostFtdcAccountregisterField
	trd.OnRspQryAccountregister = func(pAccountregister *def.CThostFtdcAccountregisterField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pAccountregister != nil {
			regInfo = *pAccountregister
		}
		if bIsLast {
			eventChan <- "OnRspQryAccountregister"
		}
	}
	trd.OnRspQryTransferBank = func(pTransferBank *def.CThostFtdcTransferBankField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			errorChan <- *pRspInfo
		} else if pTransferBank != nil {
			fmt.Printf("%+v\n", pTransferBank)
		}
		if bIsLast {
			eventChan <- "OnRspQryTransferBank"
		}
	}
	trd.OnErrRtnBankToFutureByFuture = func(pReqTransfer *def.CThostFtdcReqTransferField, pRspInfo *def.CThostFtdcRspInfoField) {
		fmt.Printf("OnErrRtnBankToFutureByFuture %+v\n", pRspInfo)
	}
	trd.OnRspFromBankToFutureByFuture = func(pReqTransfer *def.CThostFtdcReqTransferField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		fmt.Printf("OnRspFromBankToFutureByFuture %+v\n", pRspInfo)
	}
	trd.OnRtnFromBankToFutureByFuture = func(pRspTransfer *def.CThostFtdcRspTransferField) {
		fmt.Printf("%+v\n", pRspTransfer)
	}

	trd.Subscribe(def.THOST_TERT_QUICK, def.THOST_TERT_RESTART)
	trd.ReqConnect("tcp://180.168.146.187:10130")

	var testAction = func() {
		trd.ReqQryDepthMarketData("SHFE", "rb2305")
		time.Sleep(1 * time.Second)
		if lastPrice == 0 {
			return
		}
		// trd.ReqOrderInsert("rb2305", "SHFE", def.THOST_FTDC_D_Buy, def.THOST_FTDC_OF_Open, lastPrice, 3, trd.InvestorID)
		trd.ReqOrderInsert("rb2305", "SHFE", def.THOST_FTDC_D_Sell, def.THOST_FTDC_OF_CloseToday, lastPrice+20, 3, trd.InvestorID)
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
		trd.ReqFromBankToFutureByFuture(regInfo, 100)
	}

	for {
		select {
		case cb := <-eventChan:
			switch cb {
			case "OnFrontConnected": // 连接
				trd.ReqAuthenticateField("9999", "008107", "simnow_client_test", "0000000000000000")
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
				// testAction() // 测试撤单
				testIn() // 测试入金
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
			case "OnRspQryTransferBank":
			case "OnRspQryAccountregister":
				
			default:
				fmt.Println("未处理标识")
			}
		case rspInfo := <-errorChan:
			fmt.Printf("%+v\n", rspInfo)
		}
	}
}
