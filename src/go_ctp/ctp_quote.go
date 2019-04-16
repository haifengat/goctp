package go_ctp

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type CThostFtdcMdSpi uintptr

type Quote struct {
	h                                                                                                                                                                                                                                                                    *syscall.DLL
	api, spi                                                                                                                                                                                                                                                             uintptr
	nRequestID                                                                                                                                                                                                                                                           int
	funcRelease, funcInit, funcJoin, funcRegisterFront, funcRegisterNameServer, funcRegisterFensUserInfo, funcRegisterSpi, funcSubscribeMarketData, funcUnSubscribeMarketData, funcSubscribeForQuoteRsp, funcUnSubscribeForQuoteRsp, funcReqUserLogin, funcReqUserLogout *syscall.Proc
}

func (p *Quote) Quote() {
	// 创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}

	p.h = syscall.MustLoadDLL("ctp_quote.dll")
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败

	p.api, _, _ = p.h.MustFindProc("CreateApi").Call()
	p.spi, _, _ = p.h.MustFindProc("CreateSpi").Call()
	p.h.MustFindProc("RegisterSpi").Call(p.api, p.spi)
	p.funcRelease = p.h.MustFindProc("Release")
	p.funcInit = p.h.MustFindProc("Init")
	p.funcJoin = p.h.MustFindProc("Join")
	p.funcRegisterFront = p.h.MustFindProc("RegisterFront")
	p.funcRegisterNameServer = p.h.MustFindProc("RegisterNameServer")
	p.funcRegisterFensUserInfo = p.h.MustFindProc("RegisterFensUserInfo")
	p.funcRegisterSpi = p.h.MustFindProc("RegisterSpi")
	p.funcSubscribeMarketData = p.h.MustFindProc("SubscribeMarketData")
	p.funcUnSubscribeMarketData = p.h.MustFindProc("UnSubscribeMarketData")
	p.funcSubscribeForQuoteRsp = p.h.MustFindProc("SubscribeForQuoteRsp")
	p.funcUnSubscribeForQuoteRsp = p.h.MustFindProc("UnSubscribeForQuoteRsp")
	p.funcReqUserLogin = p.h.MustFindProc("ReqUserLogin")
	p.funcReqUserLogout = p.h.MustFindProc("ReqUserLogout")

	p.h.MustFindProc("SetOnFrontConnected").Call(p.spi, syscall.NewCallback(p.OnFrontConnected))
	p.h.MustFindProc("SetOnFrontDisconnected").Call(p.spi, syscall.NewCallback(p.OnFrontDisconnected))
	p.h.MustFindProc("SetOnHeartBeatWarning").Call(p.spi, syscall.NewCallback(p.OnHeartBeatWarning))
	p.h.MustFindProc("SetOnRspUserLogin").Call(p.spi, syscall.NewCallback(p.OnRspUserLogin))
	p.h.MustFindProc("SetOnRspUserLogout").Call(p.spi, syscall.NewCallback(p.OnRspUserLogout))
	p.h.MustFindProc("SetOnRspError").Call(p.spi, syscall.NewCallback(p.OnRspError))
	p.h.MustFindProc("SetOnRspSubMarketData").Call(p.spi, syscall.NewCallback(p.OnRspSubMarketData))
	p.h.MustFindProc("SetOnRspUnSubMarketData").Call(p.spi, syscall.NewCallback(p.OnRspUnSubMarketData))
	p.h.MustFindProc("SetOnRspSubForQuoteRsp").Call(p.spi, syscall.NewCallback(p.OnRspSubForQuoteRsp))
	p.h.MustFindProc("SetOnRspUnSubForQuoteRsp").Call(p.spi, syscall.NewCallback(p.OnRspUnSubForQuoteRsp))
	p.h.MustFindProc("SetOnRtnDepthMarketData").Call(p.spi, syscall.NewCallback(p.OnRtnDepthMarketData))
	p.h.MustFindProc("SetOnRtnForQuoteRsp").Call(p.spi, syscall.NewCallback(p.OnRtnForQuoteRsp))

}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (p *Quote) OnFrontConnected() uintptr {
	return 0
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (p *Quote) OnFrontDisconnected(nReason int) uintptr {
	return 0
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (p *Quote) OnHeartBeatWarning(nTimeLapse int) uintptr {
	return 0
}

// 登录请求响应
func (p *Quote) OnRspUserLogin(pRspUserLogin uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcRspUserLoginField := (*CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin))
	fmt.Println(pCThostFtdcRspUserLoginField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 登出请求响应
func (p *Quote) OnRspUserLogout(pUserLogout uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcUserLogoutField := (*CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout))
	fmt.Println(pCThostFtdcUserLogoutField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 错误应答
func (p *Quote) OnRspError(pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 订阅行情应答
func (p *Quote) OnRspSubMarketData(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 取消订阅行情应答
func (p *Quote) OnRspUnSubMarketData(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 订阅询价应答
func (p *Quote) OnRspSubForQuoteRsp(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 取消订阅询价应答
func (p *Quote) OnRspUnSubForQuoteRsp(pSpecificInstrument uintptr, pRspInfo uintptr, nRequestID int, bIsLast bool) uintptr {
	pCThostFtdcSpecificInstrumentField := (*CThostFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument))
	fmt.Println(pCThostFtdcSpecificInstrumentField)
	pCThostFtdcRspInfoField := (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo))
	fmt.Println(pCThostFtdcRspInfoField)
	return 0
}

// 深度行情通知
func (p *Quote) OnRtnDepthMarketData(pDepthMarketData uintptr) uintptr {
	pCThostFtdcDepthMarketDataField := (*CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData))
	fmt.Println(pCThostFtdcDepthMarketDataField)
	return 0
}

// 询价通知
func (p *Quote) OnRtnForQuoteRsp(pForQuoteRsp uintptr) uintptr {
	pCThostFtdcForQuoteRspField := (*CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp))
	fmt.Println(pCThostFtdcForQuoteRspField)
	return 0
}

// 创建MdApi
func (p *Quote) Release() {
	p.funcRelease.Call(p.api)
}

// 初始化
func (p *Quote) Init() {
	p.funcInit.Call(p.api)
}

// 等待接口线程结束运行
func (p *Quote) Join() {
	p.funcJoin.Call(p.api)
}

// 注册前置机网络地址
func (p *Quote) RegisterFront(pszFrontAddress string) {
	bs, _ := syscall.BytePtrFromString(pszFrontAddress)
	p.funcRegisterFront.Call(p.api, uintptr(unsafe.Pointer(bs)))
}

// @remark RegisterNameServer优先于RegisterFront
func (p *Quote) RegisterNameServer(pszNsAddress string) {
	bs, _ := syscall.BytePtrFromString(pszNsAddress)
	p.funcRegisterNameServer.Call(p.api, uintptr(unsafe.Pointer(bs)))
}

// 注册名字服务器用户信息
func (p *Quote) RegisterFensUserInfo(pFensUserInfo CThostFtdcFensUserInfoField) {
	p.funcRegisterFensUserInfo.Call(p.api, uintptr(unsafe.Pointer(&pFensUserInfo)))
}

// 注册回调接口
func (p *Quote) RegisterSpi(pSpi CThostFtdcMdSpi) {
	p.funcRegisterSpi.Call(p.api, uintptr(unsafe.Pointer(&pSpi)))
}

// 订阅行情。
func (p *Quote) SubscribeMarketData(ppInstrumentID []string, nCount int) {
	p.funcSubscribeMarketData.Call(p.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 退订行情。
func (p *Quote) UnSubscribeMarketData(ppInstrumentID []string, nCount int) {
	p.funcUnSubscribeMarketData.Call(p.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 订阅询价。
func (p *Quote) SubscribeForQuoteRsp(ppInstrumentID []string, nCount int) {
	p.funcSubscribeForQuoteRsp.Call(p.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 退订询价。
func (p *Quote) UnSubscribeForQuoteRsp(ppInstrumentID []string, nCount int) {
	p.funcUnSubscribeForQuoteRsp.Call(p.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(nCount))
}

// 用户登录请求
func (p *Quote) ReqUserLogin(pReqUserLoginField CThostFtdcReqUserLoginField) {
	p.nRequestID++
	p.funcReqUserLogin.Call(p.api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(p.nRequestID))
}

// 登出请求
func (p *Quote) ReqUserLogout(pUserLogout CThostFtdcUserLogoutField) {
	p.nRequestID++
	p.funcReqUserLogout.Call(p.api, uintptr(unsafe.Pointer(&pUserLogout)), uintptr(p.nRequestID))
}
