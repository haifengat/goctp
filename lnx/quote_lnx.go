package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.6.8_20220712
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR} -lctp_quote -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* qCreateApi();
void* qCreateSpi();
// 创建MdApi
void* qRelease(void* api);
// 初始化
void* qInit(void* api);
// 等待接口线程结束运行
void* qJoin(void* api);
// 注册前置机网络地址
void* qRegisterFront(void* api, char *pszFrontAddress);
// @remark RegisterNameServer优先于RegisterFront
void* qRegisterNameServer(void* api, char *pszNsAddress);
// 注册名字服务器用户信息
void* qRegisterFensUserInfo(void* api, struct CThostFtdcFensUserInfoField * pFensUserInfo);
// 注册回调接口
void* qRegisterSpi(void* api, void *pSpi);
// 订阅行情。
void* qSubscribeMarketData(void* api, char *ppInstrumentID[], int nCount);
// 退订行情。
void* qUnSubscribeMarketData(void* api, char *ppInstrumentID[], int nCount);
// 订阅询价。
void* qSubscribeForQuoteRsp(void* api, char *ppInstrumentID[], int nCount);
// 退订询价。
void* qUnSubscribeForQuoteRsp(void* api, char *ppInstrumentID[], int nCount);
// 用户登录请求
void* qReqUserLogin(void* api, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);
// 登出请求
void* qReqUserLogout(void* api, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);
// 请求查询组播合约
void* qReqQryMulticastInstrument(void* api, struct CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID);

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
void qSetOnFrontConnected(void*, void*);
int qFrontConnected();
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
void qSetOnFrontDisconnected(void*, void*);
int qFrontDisconnected(int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
void qSetOnHeartBeatWarning(void*, void*);
int qHeartBeatWarning(int nTimeLapse);
// 登录请求响应
void qSetOnRspUserLogin(void*, void*);
int qRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登出请求响应
void qSetOnRspUserLogout(void*, void*);
int qRspUserLogout(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询组播合约响应
void qSetOnRspQryMulticastInstrument(void*, void*);
int qRspQryMulticastInstrument(struct CThostFtdcMulticastInstrumentField *pMulticastInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 错误应答
void qSetOnRspError(void*, void*);
int qRspError(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 订阅行情应答
void qSetOnRspSubMarketData(void*, void*);
int qRspSubMarketData(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 取消订阅行情应答
void qSetOnRspUnSubMarketData(void*, void*);
int qRspUnSubMarketData(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 订阅询价应答
void qSetOnRspSubForQuoteRsp(void*, void*);
int qRspSubForQuoteRsp(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 取消订阅询价应答
void qSetOnRspUnSubForQuoteRsp(void*, void*);
int qRspUnSubForQuoteRsp(struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 深度行情通知
void qSetOnRtnDepthMarketData(void*, void*);
int qRtnDepthMarketData(struct CThostFtdcDepthMarketDataField *pDepthMarketData);
// 询价通知
void qSetOnRtnForQuoteRsp(void*, void*);
int qRtnForQuoteRsp(struct CThostFtdcForQuoteRspField *pForQuoteRsp);

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

    // 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
    _FrontConnected func()
    // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    _FrontDisconnected func(nReason int)
    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    _HeartBeatWarning func(nTimeLapse int)
    // 登录请求响应
    _RspUserLogin func(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 登出请求响应
    _RspUserLogout func(pUserLogout *ctp.CThostFtdcUserLogoutField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询组播合约响应
    _RspQryMulticastInstrument func(pMulticastInstrument *ctp.CThostFtdcMulticastInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 错误应答
    _RspError func(pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 订阅行情应答
    _RspSubMarketData func(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 取消订阅行情应答
    _RspUnSubMarketData func(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 订阅询价应答
    _RspSubForQuoteRsp func(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 取消订阅询价应答
    _RspUnSubForQuoteRsp func(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 深度行情通知
    _RtnDepthMarketData func(pDepthMarketData *ctp.CThostFtdcDepthMarketDataField)
    // 询价通知
    _RtnForQuoteRsp func(pForQuoteRsp *ctp.CThostFtdcForQuoteRspField)
    
}

var q *Quote

// NewQuote 实例化
func NewQuote() *Quote {
	q = new(Quote)
	
	// 主调函数封装 手动添加
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
	 
	// HFQuote 响应  手动添加即可增加新功能
	q._RtnDepthMarketData = func(f *ctp.CThostFtdcDepthMarketDataField) {
		q.HFQuote.RtnDepthMarketData(f)
	}
	q._RspUserLogin = func(f *ctp.CThostFtdcRspUserLoginField, i *ctp.CThostFtdcRspInfoField, n int, b bool) {
		q.HFQuote.RspUserLogin(f, i)
	}
	q._FrontConnected = func() {
		q.HFQuote.FrontConnected()
	}
	q._FrontDisconnected = func(n int) {
		q.HFQuote.FrontDisConnected(n)
	}
	
	q.HFQuote.Init() // 初始化

	q.api = C.qCreateApi()
	spi := C.qCreateSpi()
	C.qRegisterSpi(q.api, spi)

    C.qSetOnFrontConnected(spi, C.qFrontConnected)
    C.qSetOnFrontDisconnected(spi, C.qFrontDisconnected)
    C.qSetOnHeartBeatWarning(spi, C.qHeartBeatWarning)
    C.qSetOnRspUserLogin(spi, C.qRspUserLogin)
    C.qSetOnRspUserLogout(spi, C.qRspUserLogout)
    C.qSetOnRspQryMulticastInstrument(spi, C.qRspQryMulticastInstrument)
    C.qSetOnRspError(spi, C.qRspError)
    C.qSetOnRspSubMarketData(spi, C.qRspSubMarketData)
    C.qSetOnRspUnSubMarketData(spi, C.qRspUnSubMarketData)
    C.qSetOnRspSubForQuoteRsp(spi, C.qRspSubForQuoteRsp)
    C.qSetOnRspUnSubForQuoteRsp(spi, C.qRspUnSubForQuoteRsp)
    C.qSetOnRtnDepthMarketData(spi, C.qRtnDepthMarketData)
    C.qSetOnRtnForQuoteRsp(spi, C.qRtnForQuoteRsp)
    
	return q
}

//export qFrontConnected
func qFrontConnected() C.int {
    if q._FrontConnected != nil{
        q._FrontConnected()
    }
	return 0
}

//export qFrontDisconnected
func qFrontDisconnected(nReason C.int) C.int {
    if q._FrontDisconnected != nil{
        q._FrontDisconnected(int(nReason))
    }
	return 0
}

//export qHeartBeatWarning
func qHeartBeatWarning(nTimeLapse C.int) C.int {
    if q._HeartBeatWarning != nil{
        q._HeartBeatWarning(int(nTimeLapse))
    }
	return 0
}

//export qRspUserLogin
func qRspUserLogin(pRspUserLogin *C.struct_CThostFtdcRspUserLoginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspUserLogin != nil{
        q._RspUserLogin((*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspUserLogout
func qRspUserLogout(pUserLogout *C.struct_CThostFtdcUserLogoutField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspUserLogout != nil{
        q._RspUserLogout((*ctp.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspQryMulticastInstrument
func qRspQryMulticastInstrument(pMulticastInstrument *C.struct_CThostFtdcMulticastInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspQryMulticastInstrument != nil{
        q._RspQryMulticastInstrument((*ctp.CThostFtdcMulticastInstrumentField)(unsafe.Pointer(pMulticastInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspError
func qRspError(pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspError != nil{
        q._RspError((*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspSubMarketData
func qRspSubMarketData(pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspSubMarketData != nil{
        q._RspSubMarketData((*ctp.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspUnSubMarketData
func qRspUnSubMarketData(pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspUnSubMarketData != nil{
        q._RspUnSubMarketData((*ctp.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspSubForQuoteRsp
func qRspSubForQuoteRsp(pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspSubForQuoteRsp != nil{
        q._RspSubForQuoteRsp((*ctp.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRspUnSubForQuoteRsp
func qRspUnSubForQuoteRsp(pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if q._RspUnSubForQuoteRsp != nil{
        q._RspUnSubForQuoteRsp((*ctp.CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export qRtnDepthMarketData
func qRtnDepthMarketData(pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField) C.int {
    if q._RtnDepthMarketData != nil{
        q._RtnDepthMarketData((*ctp.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)))
    }
	return 0
}

//export qRtnForQuoteRsp
func qRtnForQuoteRsp(pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField) C.int {
    if q._RtnForQuoteRsp != nil{
        q._RtnForQuoteRsp((*ctp.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)))
    }
	return 0
}

