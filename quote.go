package goctp

/*
#cgo CPPFLAGS: -fPIC -I./CTPv6.7.2_20230913
#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"

#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/lib -Wl,-rpath ${SRCDIR}/lib -l ctpquote -lstdc++

void* CreateFtdcMdApi(char const*, _Bool, _Bool);
void* CreateFtdcMdSpi();

// 创建MdApi
void Release(void *api);
// 初始化
void Init(void *api);
// 等待接口线程结束运行
int Join(void *api);
// 注册前置机网络地址
void RegisterFront(void *api, char *pszFrontAddress);
// @remark RegisterNameServer优先于RegisterFront
void RegisterNameServer(void *api, char *pszNsAddress);
// 注册名字服务器用户信息
void RegisterFensUserInfo(void *api, struct CThostFtdcFensUserInfoField *pFensUserInfo);
// 注册回调接口
void RegisterSpi(void *api, void *pSpi);
// 订阅行情。
int SubscribeMarketData(void *api, char *ppInstrumentID[], int nCount);
// 退订行情。
int UnSubscribeMarketData(void *api, char *ppInstrumentID[], int nCount);
// 订阅询价。
int SubscribeForQuoteRsp(void *api, char *ppInstrumentID[], int nCount);
// 退订询价。
int UnSubscribeForQuoteRsp(void *api, char *ppInstrumentID[], int nCount);
// 用户登录请求
int ReqUserLogin(void *api, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);
// 登出请求
int ReqUserLogout(void *api, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);
// 请求查询组播合约
int ReqQryMulticastInstrument(void *api, struct CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID);

// //////////////////////////////////////////////////////////////////////
void SetOnFrontConnected(void *, void *);
void exOnFrontConnected(void*);
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
void SetOnFrontDisconnected(void *, void *);
void exOnFrontDisconnected(void*, int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
void SetOnHeartBeatWarning(void *, void *);
void exOnHeartBeatWarning(void*, int nTimeLapse);
// 登录请求响应
void SetOnRspUserLogin(void *, void *);
void exOnRspUserLogin(void*, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登出请求响应
void SetOnRspUserLogout(void *, void *);
void exOnRspUserLogout(void*, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询组播合约响应
void SetOnRspQryMulticastInstrument(void *, void *);
void exOnRspQryMulticastInstrument(void*, struct CThostFtdcMulticastInstrumentField *pMulticastInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 错误应答
void SetOnRspError(void *, void *);
void exOnRspError(void*, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 订阅行情应答
void SetOnRspSubMarketData(void *, void *);
void exOnRspSubMarketData(void*, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 取消订阅行情应答
void SetOnRspUnSubMarketData(void *, void *);
void exOnRspUnSubMarketData(void*, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 订阅询价应答
void SetOnRspSubForQuoteRsp(void *, void *);
void exOnRspSubForQuoteRsp(void*, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 取消订阅询价应答
void SetOnRspUnSubForQuoteRsp(void *, void *);
void exOnRspUnSubForQuoteRsp(void*, struct CThostFtdcSpecificInstrumentField *pSpecificInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 深度行情通知
void SetOnRtnDepthMarketData(void *, void *);
void exOnRtnDepthMarketData(void*, struct CThostFtdcDepthMarketDataField *pDepthMarketData);
// 询价通知
void SetOnRtnForQuoteRsp(void *, void *);
void exOnRtnForQuoteRsp(void*, struct CThostFtdcForQuoteRspField *pForQuoteRsp);


#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

type Quote struct {
	api, spi unsafe.Pointer
	// ************ 响应函数变量 ******************
	// //////////////////////////////////////////////////////////////////////
	OnFrontConnected func()// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	OnFrontDisconnected func(nReason int)// 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnHeartBeatWarning func(nTimeLapse int)// 登录请求响应
	OnRspUserLogin func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 登出请求响应
	OnRspUserLogout func(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 请求查询组播合约响应
	OnRspQryMulticastInstrument func(pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 错误应答
	OnRspError func(pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 订阅行情应答
	OnRspSubMarketData func(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 取消订阅行情应答
	OnRspUnSubMarketData func(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 订阅询价应答
	OnRspSubForQuoteRsp func(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 取消订阅询价应答
	OnRspUnSubForQuoteRsp func(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)// 深度行情通知
	OnRtnDepthMarketData func(pDepthMarketData *CThostFtdcDepthMarketDataField)// 询价通知
	OnRtnForQuoteRsp func(pForQuoteRsp *CThostFtdcForQuoteRspField)
}

var Quotes = make(map[unsafe.Pointer]*Quote)

func NewQuote() *Quote {
    q := &Quote{}
	path := C.CString("./log/")
	os.MkdirAll("./log/", os.ModePerm)

	q.api = C.CreateFtdcMdApi(path, false, false)

	q.spi  = C.CreateFtdcMdSpi()
	C.RegisterSpi(q.api, q.spi)

    C.SetOnFrontConnected(q.spi, C.exOnFrontConnected) // //////////////////////////////////////////////////////////////////////
    C.SetOnFrontDisconnected(q.spi, C.exOnFrontDisconnected) // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    C.SetOnHeartBeatWarning(q.spi, C.exOnHeartBeatWarning) // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    C.SetOnRspUserLogin(q.spi, C.exOnRspUserLogin) // 登录请求响应
    C.SetOnRspUserLogout(q.spi, C.exOnRspUserLogout) // 登出请求响应
    C.SetOnRspQryMulticastInstrument(q.spi, C.exOnRspQryMulticastInstrument) // 请求查询组播合约响应
    C.SetOnRspError(q.spi, C.exOnRspError) // 错误应答
    C.SetOnRspSubMarketData(q.spi, C.exOnRspSubMarketData) // 订阅行情应答
    C.SetOnRspUnSubMarketData(q.spi, C.exOnRspUnSubMarketData) // 取消订阅行情应答
    C.SetOnRspSubForQuoteRsp(q.spi, C.exOnRspSubForQuoteRsp) // 订阅询价应答
    C.SetOnRspUnSubForQuoteRsp(q.spi, C.exOnRspUnSubForQuoteRsp) // 取消订阅询价应答
    C.SetOnRtnDepthMarketData(q.spi, C.exOnRtnDepthMarketData) // 深度行情通知
    C.SetOnRtnForQuoteRsp(q.spi, C.exOnRtnForQuoteRsp) // 询价通知
    
	Quotes[q.spi] = q
    return q
}

// //////////////////////////////////////////////////////////////////////
//
//export exOnFrontConnected
func exOnFrontConnected(spi unsafe.Pointer) {
	q := Quotes[spi]
	if q.OnFrontConnected == nil {
		fmt.Println("OnFrontConnected")
	} else {
		q.OnFrontConnected()
	}
}
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
//
//export exOnFrontDisconnected
func exOnFrontDisconnected(spi unsafe.Pointer, nReason C.int) {
	q := Quotes[spi]
	if q.OnFrontDisconnected == nil {
		fmt.Println("OnFrontDisconnected")
	} else {
		q.OnFrontDisconnected(int(nReason))
	}
}
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
//
//export exOnHeartBeatWarning
func exOnHeartBeatWarning(spi unsafe.Pointer, nTimeLapse C.int) {
	q := Quotes[spi]
	if q.OnHeartBeatWarning == nil {
		fmt.Println("OnHeartBeatWarning")
	} else {
		q.OnHeartBeatWarning(int(nTimeLapse))
	}
}
// 登录请求响应
//
//export exOnRspUserLogin
func exOnRspUserLogin(spi unsafe.Pointer, pRspUserLogin *C.struct_CThostFtdcRspUserLoginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspUserLogin == nil {
		fmt.Println("OnRspUserLogin")
	} else {
		q.OnRspUserLogin((*CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 登出请求响应
//
//export exOnRspUserLogout
func exOnRspUserLogout(spi unsafe.Pointer, pUserLogout *C.struct_CThostFtdcUserLogoutField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspUserLogout == nil {
		fmt.Println("OnRspUserLogout")
	} else {
		q.OnRspUserLogout((*CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 请求查询组播合约响应
//
//export exOnRspQryMulticastInstrument
func exOnRspQryMulticastInstrument(spi unsafe.Pointer, pMulticastInstrument *C.struct_CThostFtdcMulticastInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspQryMulticastInstrument == nil {
		fmt.Println("OnRspQryMulticastInstrument")
	} else {
		q.OnRspQryMulticastInstrument((*CThostFtdcMulticastInstrumentField)(unsafe.Pointer(pMulticastInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 错误应答
//
//export exOnRspError
func exOnRspError(spi unsafe.Pointer, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspError == nil {
		fmt.Println("OnRspError")
	} else {
		q.OnRspError((*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 订阅行情应答
//
//export exOnRspSubMarketData
func exOnRspSubMarketData(spi unsafe.Pointer, pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspSubMarketData == nil {
		fmt.Println("OnRspSubMarketData")
	} else {
		q.OnRspSubMarketData((*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 取消订阅行情应答
//
//export exOnRspUnSubMarketData
func exOnRspUnSubMarketData(spi unsafe.Pointer, pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspUnSubMarketData == nil {
		fmt.Println("OnRspUnSubMarketData")
	} else {
		q.OnRspUnSubMarketData((*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 订阅询价应答
//
//export exOnRspSubForQuoteRsp
func exOnRspSubForQuoteRsp(spi unsafe.Pointer, pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspSubForQuoteRsp == nil {
		fmt.Println("OnRspSubForQuoteRsp")
	} else {
		q.OnRspSubForQuoteRsp((*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 取消订阅询价应答
//
//export exOnRspUnSubForQuoteRsp
func exOnRspUnSubForQuoteRsp(spi unsafe.Pointer, pSpecificInstrument *C.struct_CThostFtdcSpecificInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	q := Quotes[spi]
	if q.OnRspUnSubForQuoteRsp == nil {
		fmt.Println("OnRspUnSubForQuoteRsp")
	} else {
		q.OnRspUnSubForQuoteRsp((*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
// 深度行情通知
//
//export exOnRtnDepthMarketData
func exOnRtnDepthMarketData(spi unsafe.Pointer, pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField) {
	q := Quotes[spi]
	if q.OnRtnDepthMarketData == nil {
		fmt.Println("OnRtnDepthMarketData")
	} else {
		q.OnRtnDepthMarketData((*CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)))
	}
}
// 询价通知
//
//export exOnRtnForQuoteRsp
func exOnRtnForQuoteRsp(spi unsafe.Pointer, pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField) {
	q := Quotes[spi]
	if q.OnRtnForQuoteRsp == nil {
		fmt.Println("OnRtnForQuoteRsp")
	} else {
		q.OnRtnForQuoteRsp((*CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)))
	}
}


// 创建MdApi
func (q *Quote)Release(){
    C.Release(q.api)
}
// 初始化
func (q *Quote)Init(){
    C.Init(q.api)
}
// 等待接口线程结束运行
func (q *Quote)Join(){
    C.Join(q.api)
}
// 注册前置机网络地址
func (q *Quote)RegisterFront(pszFrontAddress string){
    C.RegisterFront(q.api, C.CString(pszFrontAddress))
}
// @remark RegisterNameServer优先于RegisterFront
func (q *Quote)RegisterNameServer(pszNsAddress string){
    C.RegisterNameServer(q.api, C.CString(pszNsAddress))
}
// 注册名字服务器用户信息
func (q *Quote)RegisterFensUserInfo(pFensUserInfo *CThostFtdcFensUserInfoField){
    C.RegisterFensUserInfo(q.api, (*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(pFensUserInfo)))
}
// 注册回调接口
func (q *Quote)RegisterSpi(pSpi unsafe.Pointer){
    C.RegisterSpi(q.api, pSpi)
}
// 订阅行情。
func (q *Quote)SubscribeMarketData(ppInstrumentID []string, nCount int){
    instruments := make([]*C.char, len(ppInstrumentID))
	for i := 0; i < len(instruments); i++ {
		instruments[i] = (*C.char)(unsafe.Pointer(C.CString(ppInstrumentID[i])))
	}
	C.SubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&instruments[0])), C.int(len(instruments)))
}
// 退订行情。
func (q *Quote)UnSubscribeMarketData(ppInstrumentID []string, nCount int){
    instruments := make([]*C.char, len(ppInstrumentID))
	for i := 0; i < len(instruments); i++ {
		instruments[i] = (*C.char)(unsafe.Pointer(C.CString(ppInstrumentID[i])))
	}
	C.UnSubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&instruments[0])), C.int(len(instruments)))
}
// 订阅询价。
func (q *Quote)SubscribeForQuoteRsp(ppInstrumentID []string, nCount int){
    instruments := make([]*C.char, len(ppInstrumentID))
	for i := 0; i < len(instruments); i++ {
		instruments[i] = (*C.char)(unsafe.Pointer(C.CString(ppInstrumentID[i])))
	}
	C.SubscribeForQuoteRsp(q.api, (**C.char)(unsafe.Pointer(&instruments[0])), C.int(len(instruments)))
}
// 退订询价。
func (q *Quote)UnSubscribeForQuoteRsp(ppInstrumentID []string, nCount int){
    instruments := make([]*C.char, len(ppInstrumentID))
	for i := 0; i < len(instruments); i++ {
		instruments[i] = (*C.char)(unsafe.Pointer(C.CString(ppInstrumentID[i])))
	}
	C.UnSubscribeForQuoteRsp(q.api, (**C.char)(unsafe.Pointer(&instruments[0])), C.int(len(instruments)))
}
// 用户登录请求
func (q *Quote)ReqUserLogin(pReqUserLoginField *CThostFtdcReqUserLoginField, nRequestID int){
    C.ReqUserLogin(q.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(pReqUserLoginField)), C.int(nRequestID))
}
// 登出请求
func (q *Quote)ReqUserLogout(pUserLogout *CThostFtdcUserLogoutField, nRequestID int){
    C.ReqUserLogout(q.api, (*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), C.int(nRequestID))
}
// 请求查询组播合约
func (q *Quote)ReqQryMulticastInstrument(pQryMulticastInstrument *CThostFtdcQryMulticastInstrumentField, nRequestID int){
    C.ReqQryMulticastInstrument(q.api, (*C.struct_CThostFtdcQryMulticastInstrumentField)(unsafe.Pointer(pQryMulticastInstrument)), C.int(nRequestID))
}
