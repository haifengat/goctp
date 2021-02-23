package win

import (
	ctp "gitee.com/haifengat/goctp/ctpdefine"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

type qOnFrontConnectedType func() uintptr
type qOnFrontDisconnectedType func(int) uintptr
type qOnHeartBeatWarningType func(int) uintptr
type qOnRspUserLoginType func(*ctp.CThostFtdcRspUserLoginField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRspUserLogoutType func(*ctp.CThostFtdcUserLogoutField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRspErrorType func(*ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRspSubMarketDataType func(*ctp.CThostFtdcSpecificInstrumentField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRspUnSubMarketDataType func(*ctp.CThostFtdcSpecificInstrumentField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRspSubForQuoteRspType func(*ctp.CThostFtdcSpecificInstrumentField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRspUnSubForQuoteRspType func(*ctp.CThostFtdcSpecificInstrumentField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type qOnRtnDepthMarketDataType func(*ctp.CThostFtdcDepthMarketDataField) uintptr
type qOnRtnForQuoteRspType func(*ctp.CThostFtdcForQuoteRspField) uintptr

type quote struct {
	h        *syscall.DLL
	api, spi  uintptr
	nRequestID      int
}

func (q *quote) loadDll() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	_ = os.Chdir(dllPath)
	q.h = syscall.MustLoadDLL("ctp_quote.dll")
	
	// 还原到之前的工作目录
	os.Chdir(workPath)
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 接口
func newQuote() *quote {
	q := new(quote)

	q.loadDll()
	q.api, _, _ = q.h.MustFindProc("CreateApi").Call()
	q.spi, _, _ = q.h.MustFindProc("CreateSpi").Call()
	_, _, _ = q.h.MustFindProc("RegisterSpi").Call(q.api, uintptr(unsafe.Pointer(q.spi)))
	return q
}


// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (q *quote) regOnFrontConnected(on qOnFrontConnectedType){
	_, _, _ = q.h.MustFindProc("SetOnFrontConnected").Call(q.spi, syscall.NewCallback(on))
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (q *quote) regOnFrontDisconnected(on qOnFrontDisconnectedType){
	_, _, _ = q.h.MustFindProc("SetOnFrontDisconnected").Call(q.spi, syscall.NewCallback(on))
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (q *quote) regOnHeartBeatWarning(on qOnHeartBeatWarningType){
	_, _, _ = q.h.MustFindProc("SetOnHeartBeatWarning").Call(q.spi, syscall.NewCallback(on))
}

// 登录请求响应
func (q *quote) regOnRspUserLogin(on qOnRspUserLoginType){
	_, _, _ = q.h.MustFindProc("SetOnRspUserLogin").Call(q.spi, syscall.NewCallback(on))
}

// 登出请求响应
func (q *quote) regOnRspUserLogout(on qOnRspUserLogoutType){
	_, _, _ = q.h.MustFindProc("SetOnRspUserLogout").Call(q.spi, syscall.NewCallback(on))
}

// 错误应答
func (q *quote) regOnRspError(on qOnRspErrorType){
	_, _, _ = q.h.MustFindProc("SetOnRspError").Call(q.spi, syscall.NewCallback(on))
}

// 订阅行情应答
func (q *quote) regOnRspSubMarketData(on qOnRspSubMarketDataType){
	_, _, _ = q.h.MustFindProc("SetOnRspSubMarketData").Call(q.spi, syscall.NewCallback(on))
}

// 取消订阅行情应答
func (q *quote) regOnRspUnSubMarketData(on qOnRspUnSubMarketDataType){
	_, _, _ = q.h.MustFindProc("SetOnRspUnSubMarketData").Call(q.spi, syscall.NewCallback(on))
}

// 订阅询价应答
func (q *quote) regOnRspSubForQuoteRsp(on qOnRspSubForQuoteRspType){
	_, _, _ = q.h.MustFindProc("SetOnRspSubForQuoteRsp").Call(q.spi, syscall.NewCallback(on))
}

// 取消订阅询价应答
func (q *quote) regOnRspUnSubForQuoteRsp(on qOnRspUnSubForQuoteRspType){
	_, _, _ = q.h.MustFindProc("SetOnRspUnSubForQuoteRsp").Call(q.spi, syscall.NewCallback(on))
}

// 深度行情通知
func (q *quote) regOnRtnDepthMarketData(on qOnRtnDepthMarketDataType){
	_, _, _ = q.h.MustFindProc("SetOnRtnDepthMarketData").Call(q.spi, syscall.NewCallback(on))
}

// 询价通知
func (q *quote) regOnRtnForQuoteRsp(on qOnRtnForQuoteRspType){
	_, _, _ = q.h.MustFindProc("SetOnRtnForQuoteRsp").Call(q.spi, syscall.NewCallback(on))
}



// 创建MdApi
func (q *quote) Release(){ 
	_, _, _ = q.h.MustFindProc("Release").Call(q.api)
}

// 初始化
func (q *quote) Init(){ 
	_, _, _ = q.h.MustFindProc("Init").Call(q.api)
}

// 等待接口线程结束运行
func (q *quote) Join(){ 
	_, _, _ = q.h.MustFindProc("Join").Call(q.api)
}

// 注册前置机网络地址
func (q *quote) RegisterFront(pszFrontAddress string){ 
	bs, _ := syscall.BytePtrFromString(pszFrontAddress)
	_, _, _ = q.h.MustFindProc("RegisterFront").Call(q.api, uintptr(unsafe.Pointer(bs)))
}

// @remark RegisterNameServer优先于RegisterFront
func (q *quote) RegisterNameServer(pszNsAddress string){ 
	bs, _ := syscall.BytePtrFromString(pszNsAddress)
	_, _, _ = q.h.MustFindProc("RegisterNameServer").Call(q.api, uintptr(unsafe.Pointer(bs)))
}

// 注册名字服务器用户信息
func (q *quote) RegisterFensUserInfo(){ 
	_, _, _ = q.h.MustFindProc("RegisterFensUserInfo").Call(q.api)
}

// 订阅行情。
func (q *quote) SubscribeMarketData(ppInstrumentID [1][]byte, nCount int){ 
	_, _, _ = q.h.MustFindProc("SubscribeMarketData").Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 退订行情。
func (q *quote) UnSubscribeMarketData(ppInstrumentID [1][]byte, nCount int){ 
	_, _, _ = q.h.MustFindProc("UnSubscribeMarketData").Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 订阅询价。
func (q *quote) SubscribeForQuoteRsp(ppInstrumentID [1][]byte, nCount int){ 
	_, _, _ = q.h.MustFindProc("SubscribeForQuoteRsp").Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 退订询价。
func (q *quote) UnSubscribeForQuoteRsp(ppInstrumentID [1][]byte, nCount int){ 
	_, _, _ = q.h.MustFindProc("UnSubscribeForQuoteRsp").Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(1))
}

// 用户登录请求
func (q *quote) ReqUserLogin(pReqUserLoginField ctp.CThostFtdcReqUserLoginField){ 
	q.nRequestID++
	_, _, _ = q.h.MustFindProc("ReqUserLogin").Call(q.api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(q.nRequestID))
}

// 登出请求
func (q *quote) ReqUserLogout(pUserLogout ctp.CThostFtdcUserLogoutField){ 
	q.nRequestID++
	_, _, _ = q.h.MustFindProc("ReqUserLogout").Call(q.api, uintptr(unsafe.Pointer(&pUserLogout)), uintptr(q.nRequestID))
}
