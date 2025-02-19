package goctp

import (
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type QuotePro struct {
	*QuoteExt
	IsLogin bool

	OnRtnTick func(pDepthMarketData *CThostFtdcDepthMarketDataField)

	// 行情 key: InstrumentID
	Ticks map[string]CThostFtdcDepthMarketDataField
}

func NewQuotePro() *QuotePro {
	q := &QuotePro{}
	q.QuoteExt = NewQuoteExt()

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
	done := make(chan struct{})
	first := true
	q.Quote.OnFrontConnected = func() {
		q.ReqUserLogin(cfg.Broker, cfg.UserID, cfg.Password)
	}
	q.QuoteExt.OnFrontDisconnected = func(nReason int) {
		q.IsLogin = false
	}
	q.Quote.OnRspUserLogin = func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if bIsLast {
			if pRspInfo.ErrorID == 0 {
				q.IsLogin = true
				logInfo = *pRspUserLogin
			} else {
				rsp = *pRspInfo
			}
			if first {
				first = false
				done <- struct{}{}
			}
		}
	}
	q.Quote.RegisterFront(cfg.Front)
	q.Quote.Init()

	// 连接
	select {
	case <-done:
	case <-time.NewTimer(5 * time.Second).C:
		bs, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("连接超时 5s"))
		rsp.ErrorID = -1
		copy(rsp.ErrorMsg[:], bs)
	}
	return
}
