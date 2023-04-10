package goctp

import (
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type QuotePro struct {
	*QuoteExt

	OnRtnTick func(pDepthMarketData *CThostFtdcDepthMarketDataField)

	// 行情 key: InstrumentID
	Ticks map[string]CThostFtdcDepthMarketDataField

	// 响应事件
	eventChan chan Event
	// 错误
	errorChan chan CThostFtdcRspInfoField
}

func NewQuotePro() *QuotePro {
	q := &QuotePro{}
	q.QuoteExt = NewQuoteExt()
	q.eventChan = make(chan Event)
	q.errorChan = make(chan CThostFtdcRspInfoField)

	q.Ticks = make(map[string]CThostFtdcDepthMarketDataField)

	q.Quote.OnRtnDepthMarketData = func(pDepthMarketData *CThostFtdcDepthMarketDataField) {
		q.Ticks[pDepthMarketData.InstrumentID.String()] = *pDepthMarketData
		if q.OnRtnTick != nil {
			q.OnRtnTick(pDepthMarketData)
		}
	}
	return q
}

func (q *QuotePro) Start(cfg LoginConfig) (logInfo CThostFtdcRspUserLoginField, rsp CThostFtdcRspInfoField) {
	q.Quote.OnFrontConnected = func() {
		q.eventChan <- onFrontConnected
	}
	q.Quote.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			if pRspInfo.ErrorID == 0 {
				logInfo = *pRspUserLogin
				q.eventChan <- onRspUserLogin
			} else {
				q.errorChan <- *pRspInfo
			}
		}
	}
	q.Quote.RegisterFront(cfg.Front)
	q.Quote.Init()

	// 连接
	select {
	case <-q.eventChan: // 连接
		q.ReqUserLogin(cfg.Broker, cfg.UserID, cfg.Password)
	case <-time.NewTimer(5 * time.Second).C:
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("连接超时 5s"))
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], bs)
		return
	}

	// 登录
	select {
	case <-q.eventChan:
	case rsp = <-q.errorChan:
	}
	return
}
