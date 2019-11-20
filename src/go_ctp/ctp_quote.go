package go_ctp

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

type CThostFtdcMdSpi uintptr
type onFrontConnectedType func() uintptr
type onFrontDisconnectedType func(int) uintptr
type onHeartBeatWarningType func(int) uintptr
type onRspUserLoginType func(*CThostFtdcRspUserLoginField, *CThostFtdcRspInfoField, int, bool) uintptr
type onRspUserLogoutType func(*CThostFtdcUserLogoutField, *CThostFtdcRspInfoField, int, bool) uintptr
type onRspErrorType func(*CThostFtdcRspInfoField, int, bool) uintptr
type onRspSubMarketDataType func(*CThostFtdcSpecificInstrumentField, *CThostFtdcRspInfoField, int, bool) uintptr
type onRspUnSubMarketDataType func(*CThostFtdcSpecificInstrumentField, *CThostFtdcRspInfoField, int, bool) uintptr
type onRspSubForQuoteRspType func(*CThostFtdcSpecificInstrumentField, *CThostFtdcRspInfoField, int, bool) uintptr
type onRspUnSubForQuoteRspType func(*CThostFtdcSpecificInstrumentField, *CThostFtdcRspInfoField, int, bool) uintptr
type onRtnDepthMarketDataType func(*CThostFtdcDepthMarketDataField) uintptr
type onRtnForQuoteRspType func(*CThostFtdcForQuoteRspField) uintptr

type quote struct {
	h                            *syscall.DLL
	api, spi                     uintptr
	nRequestID                   int
	funcCreateApi, funcCreateSpi *syscall.Proc
	funcRelease                  *syscall.Proc
	funcInit                     *syscall.Proc
	funcJoin                     *syscall.Proc
	funcRegisterFront            *syscall.Proc
	funcRegisterNameServer       *syscall.Proc
	funcRegisterFensUserInfo     *syscall.Proc
	funcRegisterSpi              *syscall.Proc
	funcSubscribeMarketData      *syscall.Proc
	funcUnSubscribeMarketData    *syscall.Proc
	funcSubscribeForQuoteRsp     *syscall.Proc
	funcUnSubscribeForQuoteRsp   *syscall.Proc
	funcReqUserLogin             *syscall.Proc
	funcReqUserLogout            *syscall.Proc
}

func (q *quote) loadDll() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		checkErr(os.Mkdir("log", os.ModePerm))
	}
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	checkErr(os.Chdir(path.Join(dllPath, "lib64")))
	q.h = syscall.MustLoadDLL("ctp_quote.dll")
	// 还原到之前的工作目录
	checkErr(os.Chdir(workPath))
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 行情接口
func newQuote() *quote {
	q := new(quote)

	q.loadDll()
	q.funcRelease = q.h.MustFindProc("Release")
	q.funcInit = q.h.MustFindProc("Init")
	q.funcJoin = q.h.MustFindProc("Join")
	q.funcRegisterFront = q.h.MustFindProc("RegisterFront")
	q.funcRegisterNameServer = q.h.MustFindProc("RegisterNameServer")
	q.funcRegisterFensUserInfo = q.h.MustFindProc("RegisterFensUserInfo")
	q.funcRegisterSpi = q.h.MustFindProc("RegisterSpi")
	q.funcSubscribeMarketData = q.h.MustFindProc("SubscribeMarketData")
	q.funcUnSubscribeMarketData = q.h.MustFindProc("UnSubscribeMarketData")
	q.funcSubscribeForQuoteRsp = q.h.MustFindProc("SubscribeForQuoteRsp")
	q.funcUnSubscribeForQuoteRsp = q.h.MustFindProc("UnSubscribeForQuoteRsp")
	q.funcReqUserLogin = q.h.MustFindProc("ReqUserLogin")
	q.funcReqUserLogout = q.h.MustFindProc("ReqUserLogout")

	q.api, _, _ = q.h.MustFindProc("CreateApi").Call()
	q.spi, _, _ = q.h.MustFindProc("CreateSpi").Call()
	_, _, _ = q.funcRegisterSpi.Call(q.api, uintptr(unsafe.Pointer(q.spi)))
	return q
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (q *quote) regOnFrontConnected(on onFrontConnectedType) {
	_, _, _ = q.h.MustFindProc("SetOnFrontConnected").Call(q.spi, syscall.NewCallback(on))
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (q *quote) regOnFrontDisconnected(on onFrontDisconnectedType) {
	_, _, _ = q.h.MustFindProc("SetOnFrontDisconnected").Call(q.spi, syscall.NewCallback(on))
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (q *quote) regOnHeartBeatWarning(on onHeartBeatWarningType) {
	_, _, _ = q.h.MustFindProc("SetOnHeartBeatWarning").Call(q.spi, syscall.NewCallback(on))
}

// 登录请求响应
func (q *quote) regOnRspUserLogin(on onRspUserLoginType) {
	_, _, _ = q.h.MustFindProc("SetOnRspUserLogin").Call(q.spi, syscall.NewCallback(on))
}

// 登出请求响应
func (q *quote) regOnRspUserLogout(on onRspUserLogoutType) {
	_, _, _ = q.h.MustFindProc("SetOnRspUserLogout").Call(q.spi, syscall.NewCallback(on))
}

// 错误应答
func (q *quote) regOnRspError(on onRspErrorType) {
	_, _, _ = q.h.MustFindProc("SetOnRspError").Call(q.spi, syscall.NewCallback(on))
}

// 订阅行情应答
func (q *quote) regOnRspSubMarketData(on onRspSubMarketDataType) {
	_, _, _ = q.h.MustFindProc("SetOnRspSubMarketData").Call(q.spi, syscall.NewCallback(on))
}

// 取消订阅行情应答
func (q *quote) regOnRspUnSubMarketData(on onRspUnSubMarketDataType) {
	_, _, _ = q.h.MustFindProc("SetOnRspUnSubMarketData").Call(q.spi, syscall.NewCallback(on))
}

// 订阅询价应答
func (q *quote) regOnRspSubForQuoteRsp(on onRspSubForQuoteRspType) {
	_, _, _ = q.h.MustFindProc("SetOnRspSubForQuoteRsp").Call(q.spi, syscall.NewCallback(on))
}

// 取消订阅询价应答
func (q *quote) regOnRspUnSubForQuoteRsp(on onRspUnSubForQuoteRspType) {
	_, _, _ = q.h.MustFindProc("SetOnRspUnSubForQuoteRsp").Call(q.spi, syscall.NewCallback(on))
}

// 深度行情通知
func (q *quote) regOnRtnDepthMarketData(on onRtnDepthMarketDataType) {
	_, _, _ = q.h.MustFindProc("SetOnRtnDepthMarketData").Call(q.spi, syscall.NewCallback(on))
}

// 询价通知
func (q *quote) regOnRtnForQuoteRsp(on onRtnForQuoteRspType) {
	_, _, _ = q.h.MustFindProc("SetOnRtnForQuoteRsp").Call(q.spi, syscall.NewCallback(on))
}

// 创建MdApi
func (q *quote) Release() {
	_, _, _ = q.funcRelease.Call(q.api)
}

// 初始化
func (q *quote) Init() {
	_, _, _ = q.funcInit.Call(q.api)
}

// 等待接口线程结束运行
func (q *quote) Join() {
	_, _, _ = q.funcJoin.Call(q.api)
}

// 注册前置机网络地址
func (q *quote) RegisterFront(pszFrontAddress string) {
	bs, _ := syscall.BytePtrFromString(pszFrontAddress)
	_, _, _ = q.funcRegisterFront.Call(q.api, uintptr(unsafe.Pointer(bs)))
}

// @remark RegisterNameServer优先于RegisterFront
func (q *quote) RegisterNameServer(pszNsAddress string) {
	bs, _ := syscall.BytePtrFromString(pszNsAddress)
	_, _, _ = q.funcRegisterNameServer.Call(q.api, uintptr(unsafe.Pointer(bs)))
}

// 注册名字服务器用户信息
func (q *quote) RegisterFensUserInfo(pFensUserInfo CThostFtdcFensUserInfoField) {
	_, _, _ = q.funcRegisterFensUserInfo.Call(q.api, uintptr(unsafe.Pointer(&pFensUserInfo)))
}

// 注册回调接口
func (q *quote) RegisterSpi(pSpi CThostFtdcMdSpi) {
	_, _, _ = q.funcRegisterSpi.Call(q.api, uintptr(unsafe.Pointer(&pSpi)))
}

// 订阅行情。
func (q *quote) SubscribeMarketData(ppInstrumentID [1][]byte, nCount int) {
	_, _, _ = q.funcSubscribeMarketData.Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 退订行情。
func (q *quote) UnSubscribeMarketData(ppInstrumentID [1][]byte, nCount int) {
	_, _, _ = q.funcUnSubscribeMarketData.Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 订阅询价。
func (q *quote) SubscribeForQuoteRsp(ppInstrumentID [1][]byte, nCount int) {
	_, _, _ = q.funcSubscribeForQuoteRsp.Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 退订询价。
func (q *quote) UnSubscribeForQuoteRsp(ppInstrumentID [1][]byte, nCount int) {
	_, _, _ = q.funcUnSubscribeForQuoteRsp.Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 用户登录请求
func (q *quote) ReqUserLogin(pReqUserLoginField CThostFtdcReqUserLoginField) {
	q.nRequestID++
	_, _, _ = q.funcReqUserLogin.Call(q.api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(q.nRequestID))
}

// 登出请求
func (q *quote) ReqUserLogout(pUserLogout CThostFtdcUserLogoutField) {
	q.nRequestID++
	_, _, _ = q.funcReqUserLogout.Call(q.api, uintptr(unsafe.Pointer(&pUserLogout)), uintptr(q.nRequestID))
}
