package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.3.15_20190220
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR} -lctp_quote -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* qCreateApi();
void* qCreateSpi();
void* qRegisterSpi(void*, void*);
void* qRegisterFront(void*, char*);
void* qInit(void*);
void* qRelease(void*);
void* qReqUserLogin(void*, struct CThostFtdcReqUserLoginField*, int);
void* qSubscribeMarketData(void*, char *ppInstrumentID[], int nCount);

void qSetOnFrontConnected(void*, void*);
int QOnFrontConnected();
void qSetOnRspUserLogin(void*, void*);
int QOnRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void qSetOnRtnDepthMarketData(void*, void*);
int OnRtnDepthMarketData(struct CThostFtdcDepthMarketDataField *pDepthMarketData);
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"unsafe"

	"github.com/haifengat/goctp"
	ctp "github.com/haifengat/goctp/ctpdefine"
)

// Quote 行情接口
type Quote struct {
	api unsafe.Pointer
	// 帐号
	InvestorID string
	// 经纪商
	BrokerID string
	// 交易日
	TradingDay       string
	onFrontConnected goctp.OnFrontConnectedType
	onRspUserLogin   goctp.OnRspUserLoginType
	onTick           goctp.OnTickType
	reqID            int
}

func (q *Quote) getReqID() C.int {
	q.reqID++
	return C.int(q.reqID)
}

var q *Quote

// NewQuote new md api instanse
func NewQuote() *Quote {
	q = new(Quote)
	q.api = C.qCreateApi()
	spi := C.qCreateSpi()
	C.qRegisterSpi(q.api, spi)

	C.qSetOnFrontConnected(spi, C.QOnFrontConnected)
	C.qSetOnRspUserLogin(spi, C.QOnRspUserLogin)
	C.qSetOnRtnDepthMarketData(spi, C.OnRtnDepthMarketData)
	return q
}

func (q *Quote) Release() {
	C.qRelease(q.api)
}

func (q *Quote) ReqConnect(addr string) {
	front := C.CString(addr)
	C.qRegisterFront(q.api, front)
	defer C.free(unsafe.Pointer(front))
	C.qInit(q.api)
}

func (q *Quote) ReqLogin(investor, pwd, broker string) {
	q.InvestorID = investor
	q.BrokerID = broker
	f := ctp.CThostFtdcReqUserLoginField{}
	copy(f.UserID[:], q.InvestorID)
	copy(f.BrokerID[:], q.BrokerID)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@HF")
	C.qReqUserLogin(q.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(&f)), q.getReqID())
}

func (q *Quote) ReqSubscript(instrument string) {
	inst := make([]*C.char, 1)
	inst[0] = (*C.char)(unsafe.Pointer(C.CString(instrument)))
	C.qSubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&inst[0])), 1)
}

func (q *Quote) RegOnFrontConnected(on goctp.OnFrontConnectedType) {
	q.onFrontConnected = on
}
func (q *Quote) RegOnRspUserLogin(on goctp.OnRspUserLoginType) {
	q.onRspUserLogin = on
}
func (q *Quote) RegOnTick(on goctp.OnTickType) {
	q.onTick = on
}

//export OnRtnDepthMarketData
func OnRtnDepthMarketData(field *C.struct_CThostFtdcDepthMarketDataField) C.int {
	dataField := (*ctp.CThostFtdcDepthMarketDataField)(unsafe.Pointer(field))
	if q.onTick == nil {
		return 0
	}
	tick := goctp.TickField{
		TradingDay:      goctp.Bytes2String(dataField.TradingDay[:]),
		InstrumentID:    goctp.Bytes2String(dataField.InstrumentID[:]),
		ExchangeID:      goctp.Bytes2String(dataField.ExchangeID[:]),
		LastPrice:       float64(dataField.LastPrice),
		OpenPrice:       float64(dataField.OpenPrice),
		HighestPrice:    float64(dataField.HighestPrice),
		LowestPrice:     float64(dataField.LowestPrice),
		Volume:          int(dataField.Volume),
		Turnover:        float64(dataField.Turnover),
		OpenInterest:    float64(dataField.OpenInterest),
		SettlementPrice: float64(dataField.SettlementPrice),
		UpperLimitPrice: float64(dataField.UpperLimitPrice),
		LowerLimitPrice: float64(dataField.LowerLimitPrice),
		CurrDelta:       float64(dataField.CurrDelta),
		UpdateTime:      goctp.Bytes2String(dataField.UpdateTime[:]),
		UpdateMillisec:  int(dataField.UpdateMillisec),
		BidPrice1:       float64(dataField.BidPrice1),
		BidVolume1:      int(dataField.BidVolume1),
		AskPrice1:       float64(dataField.AskPrice1),
		AskVolume1:      int(dataField.AskVolume1),
		BidPrice2:       float64(dataField.BidPrice2),
		BidVolume2:      int(dataField.BidVolume2),
		AskPrice2:       float64(dataField.AskPrice2),
		AskVolume2:      int(dataField.AskVolume2),
		BidPrice3:       float64(dataField.BidPrice3),
		BidVolume3:      int(dataField.BidVolume3),
		AskPrice3:       float64(dataField.AskPrice3),
		AskVolume3:      int(dataField.AskVolume3),
		BidPrice4:       float64(dataField.BidPrice4),
		BidVolume4:      int(dataField.BidVolume4),
		AskPrice4:       float64(dataField.AskPrice4),
		AskVolume4:      int(dataField.AskVolume4),
		BidPrice5:       float64(dataField.BidPrice5),
		BidVolume5:      int(dataField.BidVolume5),
		AskPrice5:       float64(dataField.AskPrice5),
		AskVolume5:      int(dataField.AskVolume5),
		AveragePrice:    float64(dataField.AskPrice5),
		ActionDay:       goctp.Bytes2String(dataField.ActionDay[:]),
	}
	q.onTick(&tick)
	return 0
}

// 登陆响应
//export QOnRspUserLogin
func QOnRspUserLogin(field *C.struct_CThostFtdcRspUserLoginField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {

	loginField := (*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if q.onRspUserLogin == nil {
		return 0
	}
	q.onRspUserLogin(&goctp.RspUserLoginField{
		TradingDay:  string(loginField.TradingDay[:]),
		LoginTime:   string(loginField.LoginTime[:]),
		BrokerID:    string(loginField.BrokerID[:]),
		UserID:      string(loginField.UserID[:]),
		FrontID:     int(loginField.FrontID),
		SessionID:   int(loginField.SessionID),
		MaxOrderRef: string(loginField.MaxOrderRef[:]),
	}, &goctp.RspInfoField{
		ErrorID:  int(infoField.ErrorID),
		ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:]),
	})
	return 0
}

// QOnFrontConnected 连接前置响应
//export QOnFrontConnected
func QOnFrontConnected() C.int {
	if q.onFrontConnected != nil {
		q.onFrontConnected()
	}
	return 0
}
