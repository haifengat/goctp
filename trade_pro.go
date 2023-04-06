package goctp

import (
	"fmt"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type Event int

const (
	onFrontConnected Event = iota
	onRspAuthenticate
	onRspUserLogin
	onRspSettlementInfoConfirm
	onRspQryInvestor
	onRspQryClassifiedInstrument
	onRspQryOrder
	onRspQryTrade
	onRspQryInvestorPosition
	onRspQryInvestorPositionDetail
	onRspQryTradingAccount
	onRspQryAccountregister
	onRspQryTransferBank
)

// TradePro 业务逻辑封装
type TradePro struct {
	*TradeExt
	// 响应事件
	eventChan chan Event
	// 错误
	errorChan chan CThostFtdcRspInfoField
}

func NewTradePro() *TradePro {
	trd := TradePro{}
	trd.TradeExt = NewTradeExt()

	trd.eventChan = make(chan Event)
	trd.errorChan = make(chan CThostFtdcRspInfoField)

	trd.OnRspQryInvestorPosition = func(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pInvestorPosition != nil {
			fmt.Printf("OnRspQryInvestorPosition %+v\n", pInvestorPosition)
		}
		if bIsLast {
			trd.eventChan <- onRspQryInvestorPosition
		}
	}
	trd.OnRspQryInvestorPositionDetail = func(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pInvestorPositionDetail != nil {
			fmt.Printf("OnRspQryInvestorPositionDetail %+v\n", pInvestorPositionDetail)
		}
		if bIsLast {
			trd.eventChan <- onRspQryInvestorPositionDetail
		}
	}
	trd.OnRspQryTradingAccount = func(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pTradingAccount != nil {
			fmt.Printf("OnRspQryTradingAccount %+v\n", pTradingAccount)
		}
		if bIsLast {
			trd.eventChan <- onRspQryTradingAccount
		}
	}
	trd.OnRspQryAccountregister = func(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
			// } else if pAccountregister != nil {
			// 	//
		}
		if bIsLast {
			trd.eventChan <- onRspQryAccountregister
		}
	}
	trd.OnRspQryTransferBank = func(pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pTransferBank != nil {
			fmt.Printf("%+v\n", pTransferBank)
		}
		if bIsLast {
			trd.eventChan <- onRspQryTransferBank
		}
	}
	return &trd
}

type LoginConfig struct {
	Front, Broker, UserID, Password, AppID, AuthCode string
}

func (trd *TradePro) Start(cfg LoginConfig) (logInfo CThostFtdcRspUserLoginField, rsp CThostFtdcRspInfoField) {
	trd.Trade.OnFrontConnected = func() {
		trd.eventChan <- onFrontConnected
	}
	trd.Trade.OnRspAuthenticate = func(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else {
			trd.eventChan <- onRspAuthenticate
		}
	}
	trd.Trade.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else {
			fmt.Printf("%+v\n", pRspUserLogin)
			fmt.Printf("%+v\n", pRspInfo)
			logInfo = *pRspUserLogin
			trd.eventChan <- onRspUserLogin
		}
	}
	trd.Trade.OnRspSettlementInfoConfirm = func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else {
			trd.eventChan <- onRspSettlementInfoConfirm
		}
	}
	trd.Trade.OnRspQryInvestor = func(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pInvestor != nil {
			fmt.Printf("OnRspQryInvestor %+v\n", pInvestor)
		}
		if bIsLast {
			trd.eventChan <- onRspQryInvestor
		}
	}
	trd.Trade.OnRspQryClassifiedInstrument = func(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pInstrument != nil {
			fmt.Printf("OnRspQryClassifiedInstrument %+v\n", pInstrument)
		}
		if bIsLast {
			trd.eventChan <- onRspQryClassifiedInstrument
		}
	}
	trd.Trade.OnRspQryOrder = func(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pOrder != nil {
			fmt.Printf("OnRspQryOrder %+v\n", pOrder)
		}
		if bIsLast {
			trd.eventChan <- onRspQryOrder
		}
	}
	trd.Trade.OnRspQryTrade = func(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pRspInfo != nil && pRspInfo.ErrorID != 0 {
			trd.errorChan <- *pRspInfo
		} else if pTrade != nil {
			fmt.Printf("OnRspQryTrade %+v\n", pTrade)
		}
		if bIsLast {
			trd.eventChan <- onRspQryTrade
		}
	}

	trd.TradeExt.RegisterFront(cfg.Front)
	trd.TradeExt.SubscribePrivateTopic(THOST_TERT_QUICK)
	trd.TradeExt.SubscribePublicTopic(THOST_TERT_RESTART)
	trd.Init()

	// 登录过程
	for {
		select {
		case cb := <-trd.eventChan:
			switch cb {
			case onFrontConnected: // 连接
				trd.ReqAuthenticate(cfg.Broker, cfg.UserID, cfg.AppID, cfg.AuthCode)
			case onRspAuthenticate: // 认证
				trd.ReqUserLogin(cfg.Password)
			case onRspUserLogin: // 登录
				trd.ReqSettlementInfoConfirm()
			case onRspSettlementInfoConfirm: // 确认结算
				trd.ReqQryInvestor()
			case onRspQryInvestor: // 查用户
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryClassifiedInstrument()
			case onRspQryClassifiedInstrument: // 查合约
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryOrder()
			case onRspQryOrder: // 查委托
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryTrade()
			case onRspQryTrade: // 查成交
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryPosition()
			case onRspQryInvestorPosition: //查持仓
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryPositionDetail()
			case onRspQryInvestorPositionDetail: // 查持仓明细
				time.Sleep(time.Millisecond * 1100)
				trd.ReqQryTradingAccount()
			case onRspQryTradingAccount: // 查持仓后,发送委托
				fmt.Println("登录过程完成")
				bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("正确"))
				copy(rsp.ErrorMsg[:], bs)
				return
			case onRspQryTransferBank:
			case onRspQryAccountregister:

			default:
				fmt.Println("未处理标识")
			}
		case rspInfo := <-trd.errorChan:
			fmt.Printf("%+v\n", rspInfo)
			return
		case <-time.NewTimer(5 * time.Second).C:
			fmt.Println("登录超时")
			copy(rsp.ErrorMsg[:], "TimeOut")
			return
		}
	}
}
