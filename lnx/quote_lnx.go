package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.6.8_20220712
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR} -lctp_quote -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* qCreateApi();
void* qCreateSpi();
void* qRegisterSpi(void*, void*);
void* qRegisterFront(void*, char*);
void* qInit(void*);
void* qJoin(void*);
void* qRelease(void*);
void* qReqUserLogin(void*, struct CThostFtdcReqUserLoginField*, int);
void* qSubscribeMarketData(void*, char *ppInstrumentID[], int nCount);

void qSetOnFrontConnected(void*, void*);
int qFrontConnected();
void qSetOnFrontDisconnected(void*, void*);
int qFrontDisConnected();
void qSetOnRspUserLogin(void*, void*);
int qRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void qSetOnRtnDepthMarketData(void*, void*);
int qRtnDepthMarketData(struct CThostFtdcDepthMarketDataField *pDepthMarketData);
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"unsafe"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

// Quote 行情接口
type Quote struct {
	goctp.HFQuote // 组合

	api unsafe.Pointer
}

var q *Quote

// NewQuote new md api instanse
func NewQuote() *Quote {
	q = new(Quote)
	q.HFQuote.ReqConnect = func(addr string) {
		front := C.CString(addr)
		C.qRegisterFront(q.api, front)
		defer C.free(unsafe.Pointer(front))
		C.qInit(q.api)
		// C.qJoin(q.api)
	}
	q.HFQuote.ReleaseAPI = func() {
		C.qRegisterSpi(q.api, nil)
		C.qRelease(q.api)
		q.api = nil
	}
	q.HFQuote.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		C.qReqUserLogin(q.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(f)), C.int(1))
	}
	q.HFQuote.ReqSubMarketData = func(instrument ...string) {
		ppInstrumentID := make([]*C.char, len(instrument))
		for i := 0; i < len(instrument); i++ {
			ppInstrumentID[i] = (*C.char)(unsafe.Pointer(C.CString(instrument[i])))
		}
		C.qSubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(len(instrument)))
	}
	q.HFQuote.Init() // 初始化

	q.api = C.qCreateApi()
	spi := C.qCreateSpi()
	C.qRegisterSpi(q.api, spi)

	C.qSetOnFrontConnected(spi, C.qFrontConnected)
	C.qSetOnFrontDisconnected(spi, C.qFrontDisConnected)
	C.qSetOnRspUserLogin(spi, C.qRspUserLogin)
	C.qSetOnRtnDepthMarketData(spi, C.qRtnDepthMarketData)
	return q
}

//export qRtnDepthMarketData
func qRtnDepthMarketData(field *C.struct_CThostFtdcDepthMarketDataField) C.int {
	dataField := (*ctp.CThostFtdcDepthMarketDataField)(unsafe.Pointer(field))
	q.HFQuote.RtnDepthMarketData(dataField)
	return 0
}

//export qRspUserLogin
func qRspUserLogin(field *C.struct_CThostFtdcRspUserLoginField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	loginField := (*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	q.HFQuote.RspUserLogin(loginField, infoField)
	return 0
}

//export qFrontConnected
func qFrontConnected() C.int {
	q.HFQuote.FrontConnected()
	return 0
}

//export qFrontDisConnected
func qFrontDisConnected(reason C.int) C.int {
	q.HFQuote.FrontDisConnected(int(reason))
	return 0
}
