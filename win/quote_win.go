package win

import (
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

type Quote struct {
	goctp.HFQuote
	h        *syscall.DLL
	api, spi uintptr
}

func (q *Quote) loadDll() {
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
func NewQuote() *Quote {
	q := new(Quote)
	q.HFQuote.ReqConnect = func(addr string) {
		bs, _ := syscall.BytePtrFromString(addr)
		q.h.MustFindProc("qRegisterFront").Call(q.api, uintptr(unsafe.Pointer(bs)))
		q.h.MustFindProc("qInit").Call(q.api)
		// q.h.MustFindProc("Join").Call(q.api)
	}
	q.HFQuote.ReleaseAPI = func() {
		q.h.MustFindProc("qRelease").Call(q.api)
	}
	q.HFQuote.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		q.h.MustFindProc("qReqUserLogin").Call(q.api, uintptr(unsafe.Pointer(&f)), uintptr(1))
	}
	q.HFQuote.ReqSubMarketData = func(instrument ...string) {
		ppInstrumentID := make([][]byte, len(instrument)) // [][]byte{[]byte(instrument)}
		for i := 0; i < len(instrument); i++ {
			copy(ppInstrumentID[i], []byte(instrument[i]))
		}
		q.h.MustFindProc("qSubscribeMarketData").Call(q.api, uintptr(unsafe.Pointer(&ppInstrumentID)), uintptr(len(instrument)))
	}
	q.HFQuote.Init() // 初始化

	q.loadDll()
	q.api, _, _ = q.h.MustFindProc("qCreateApi").Call()
	q.spi, _, _ = q.h.MustFindProc("qCreateSpi").Call()

	q.h.MustFindProc("qRegisterSpi").Call(q.api, uintptr(q.spi))

	q.h.MustFindProc("qSetOnFrontConnected").Call(q.spi, syscall.NewCallback(q.OnFrontConnected))
	q.h.MustFindProc("qSetOnFrontDisconnected").Call(q.spi, syscall.NewCallback(q.OnFrontDisconnected))
	q.h.MustFindProc("qSetOnRspUserLogin").Call(q.spi, syscall.NewCallback(q.OnRspUserLogin))
	q.h.MustFindProc("qSetOnRtnDepthMarketData").Call(q.spi, syscall.NewCallback(q.OnRtnDepthMarketData))
	return q
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (q *Quote) OnFrontConnected() uintptr {
	q.HFQuote.FrontConnected()
	return 0
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (q *Quote) OnFrontDisconnected(reason int) uintptr {
	q.HFQuote.FrontDisConnected(reason)
	return 0
}

// 登录请求响应
func (q *Quote) OnRspUserLogin(field *ctp.CThostFtdcRspUserLoginField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	q.HFQuote.RspUserLogin(field, info)
	return 0
}

// 深度行情通知
func (q *Quote) OnRtnDepthMarketData(field *ctp.CThostFtdcDepthMarketDataField) uintptr {
	q.HFQuote.RtnDepthMarketData(field)
	return 0
}
