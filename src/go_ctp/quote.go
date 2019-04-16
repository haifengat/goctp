package go_ctp

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type CThostFtdcMdSpi uintptr

var (
	h                                                                                                                                                                                                                                                                    *syscall.DLL
	api, spi                                                                                                                                                                                                                                                             uintptr
	nRequestID                                                                                                                                                                                                                                                           int
	funcRelease, funcInit, funcJoin, funcRegisterFront, funcRegisterNameServer, funcRegisterFensUserInfo, funcRegisterSpi, funcSubscribeMarketData, funcUnSubscribeMarketData, funcSubscribeForQuoteRsp, funcUnSubscribeForQuoteRsp, funcReqUserLogin, funcReqUserLogout *syscall.Proc
)

func InitQuote() {
	// 创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}

	h = syscall.MustLoadDLL("ctp_quote.dll")
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败

	api, _, _ = h.MustFindProc("CreateApi").Call()
	spi, _, _ = h.MustFindProc("CreateSpi").Call()
	h.MustFindProc("RegisterSpi").Call(api, spi)
	funcRelease = h.MustFindProc("Release")
	funcInit = h.MustFindProc("Init")
	funcJoin = h.MustFindProc("Join")
	funcRegisterFront = h.MustFindProc("RegisterFront")
	funcRegisterNameServer = h.MustFindProc("RegisterNameServer")
	funcRegisterFensUserInfo = h.MustFindProc("RegisterFensUserInfo")
	funcRegisterSpi = h.MustFindProc("RegisterSpi")
	funcSubscribeMarketData = h.MustFindProc("SubscribeMarketData")
	funcUnSubscribeMarketData = h.MustFindProc("UnSubscribeMarketData")
	funcSubscribeForQuoteRsp = h.MustFindProc("SubscribeForQuoteRsp")
	funcUnSubscribeForQuoteRsp = h.MustFindProc("UnSubscribeForQuoteRsp")
	funcReqUserLogin = h.MustFindProc("ReqUserLogin")
	funcReqUserLogout = h.MustFindProc("ReqUserLogout")

	h.MustFindProc("SetOnFrontConnected").Call(spi, syscall.NewCallback(OnFrontConnected))
	h.MustFindProc("SetOnFrontDisconnected").Call(spi, syscall.NewCallback(OnFrontDisconnected))
	h.MustFindProc("SetOnHeartBeatWarning").Call(spi, syscall.NewCallback(OnHeartBeatWarning))
	h.MustFindProc("SetOnRspUserLogin").Call(spi, syscall.NewCallback(OnRspUserLogin))
	h.MustFindProc("SetOnRspUserLogout").Call(spi, syscall.NewCallback(OnRspUserLogout))
	h.MustFindProc("SetOnRspError").Call(spi, syscall.NewCallback(OnRspError))
	h.MustFindProc("SetOnRspSubMarketData").Call(spi, syscall.NewCallback(OnRspSubMarketData))
	h.MustFindProc("SetOnRspUnSubMarketData").Call(spi, syscall.NewCallback(OnRspUnSubMarketData))
	h.MustFindProc("SetOnRspSubForQuoteRsp").Call(spi, syscall.NewCallback(OnRspSubForQuoteRsp))
	h.MustFindProc("SetOnRspUnSubForQuoteRsp").Call(spi, syscall.NewCallback(OnRspUnSubForQuoteRsp))
	h.MustFindProc("SetOnRtnDepthMarketData").Call(spi, syscall.NewCallback(OnRtnDepthMarketData))
	h.MustFindProc("SetOnRtnForQuoteRsp").Call(spi, syscall.NewCallback(OnRtnForQuoteRsp))

}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func OnFrontConnected() uintptr {
	return 0
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func OnFrontDisconnected(nReason int) uintptr {
	return 0
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func OnHeartBeatWarning(nTimeLapse int) uintptr {
	return 0
}

// 登录请求响应
func OnRspUserLogin(pRspUserLogin uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcRspUserLoginField := (*CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin))
	fmt.Println(pCThostFtdcRspUserLoginField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 登出请求响应
func OnRspUserLogout(pUserLogout uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcUserLogoutField := (*CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout))
	fmt.Println(pCThostFtdcUserLogoutField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 错误应答
func OnRspError(pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 订阅行情应答
func OnRspSubMarketData(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 取消订阅行情应答
func OnRspUnSubMarketData(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 订阅询价应答
func OnRspSubForQuoteRsp(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 取消订阅询价应答
func OnRspUnSubForQuoteRsp(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 深度行情通知
func OnRtnDepthMarketData(pDepthMarketData uintptr) uintptr {
	pCThostFtdcDepthMarketDataField := (*CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData))
	fmt.Println(pCThostFtdcDepthMarketDataField)
	return 0
}

// 询价通知
func OnRtnForQuoteRsp(pForQuoteRsp uintptr) uintptr {
	pCThostFtdcForQuoteRspField := (*CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp))
	fmt.Println(pCThostFtdcForQuoteRspField)
	return 0
}

// 创建MdApi
func Release() {
	funcRelease.Call(api)
}

// 初始化
func Init() {
	funcInit.Call(api)
}

// 等待接口线程结束运行
func Join() {
	funcJoin.Call(api)
}

// 注册前置机网络地址
func RegisterFront(pszFrontAddress string) {
	bs, _ := syscall.BytePtrFromString(pszFrontAddress)
	funcRegisterFront.Call(api, uintptr(unsafe.Pointer(bs)))
}

// @remark RegisterNameServer优先于RegisterFront
func RegisterNameServer(pszNsAddress string) {
	bs, _ := syscall.BytePtrFromString(pszNsAddress)
	funcRegisterNameServer.Call(api, uintptr(unsafe.Pointer(bs)))
}

// 注册名字服务器用户信息
func RegisterFensUserInfo(pFensUserInfo CThostFtdcFensUserInfoField) {
	funcRegisterFensUserInfo.Call(api, uintptr(unsafe.Pointer(&pFensUserInfo)))
}

// 注册回调接口
func RegisterSpi(pSpi CThostFtdcMdSpi) {
	funcRegisterSpi.Call(api, uintptr(unsafe.Pointer(&pSpi)))
}

// 订阅行情。
func SubscribeMarketData(ppInstrumentID []string, nCount int) {
	funcSubscribeMarketData.Call(api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 退订行情。
func UnSubscribeMarketData(ppInstrumentID []string, nCount int) {
	funcUnSubscribeMarketData.Call(api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 订阅询价。
func SubscribeForQuoteRsp(ppInstrumentID []string, nCount int) {
	funcSubscribeForQuoteRsp.Call(api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 退订询价。
func UnSubscribeForQuoteRsp(ppInstrumentID []string, nCount int) {
	funcUnSubscribeForQuoteRsp.Call(api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 用户登录请求
func ReqUserLogin(pReqUserLoginField CThostFtdcReqUserLoginField, nRequestID int) {
	nRequestID++
	funcReqUserLogin.Call(api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(nRequestID))
}

// 登出请求
func ReqUserLogout(pUserLogout CThostFtdcUserLogoutField, nRequestID int) {
	nRequestID++
	funcReqUserLogout.Call(api, uintptr(unsafe.Pointer(&pUserLogout)), uintptr(nRequestID))
}
