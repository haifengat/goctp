package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.5.1_20200908
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
	"os"
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
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
	q = new(Quote)
	q.api = C.qCreateApi()
	spi := C.qCreateSpi()
	C.qRegisterSpi(q.api, spi)

	C.qSetOnFrontConnected(spi, C.qFrontConnected)
	C.qSetOnFrontDisconnected(spi, C.qFrontDisConnected)
	C.qSetOnRspUserLogin(spi, C.qRspUserLogin)
	C.qSetOnRtnDepthMarketData(spi, C.qRtnDepthMarketData)
	return q
}

// Release 接口消毁
func (q *Quote) Release() {
	C.qRelease(q.api)
}

// ReqConnect 连接前置;Join阻塞，请用goroutine
func (q *Quote) ReqConnect(addr string) {
	front := C.CString(addr)
	C.qRegisterFront(q.api, front)
	defer C.free(unsafe.Pointer(front))
	C.qInit(q.api)
	// C.qJoin(q.api)
}

// ReqLogin 登录
func (q *Quote) ReqLogin(investor, pwd, broker string) {
	q.InvestorID = investor
	q.BrokerID = broker
	f := ctp.CThostFtdcReqUserLoginField{}
	copy(f.UserID[:], q.InvestorID)
	copy(f.BrokerID[:], q.BrokerID)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@HF")
	C.qReqUserLogin(q.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(&f)), C.int(1))
}

// ReqSubscript 订阅行情
func (q *Quote) ReqSubscript(instrument string) {
	inst := make([]*C.char, 1)
	inst[0] = (*C.char)(unsafe.Pointer(C.CString(instrument)))
	C.qSubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&inst[0])), C.int(1))
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
