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

func (q *Quote) loadDll() {
	workPath, _ := os.Getwd()
	_, curFile, _, _ := runtime.Caller(1)
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

	// 主调函数封装 手动添加
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

	q.loadDll()
	q.api, _, _ = q.h.MustFindProc("qCreateApi").Call()
	q.spi, _, _ = q.h.MustFindProc("qCreateSpi").Call()
	q.h.MustFindProc("qRegisterSpi").Call(q.api, uintptr(q.spi))

    q.h.MustFindProc("qSetOnFrontConnected").Call(q.spi, syscall.NewCallback(q.OnFrontConnected))
    q.h.MustFindProc("qSetOnFrontDisconnected").Call(q.spi, syscall.NewCallback(q.OnFrontDisconnected))
    q.h.MustFindProc("qSetOnHeartBeatWarning").Call(q.spi, syscall.NewCallback(q.OnHeartBeatWarning))
    q.h.MustFindProc("qSetOnRspUserLogin").Call(q.spi, syscall.NewCallback(q.OnRspUserLogin))
    q.h.MustFindProc("qSetOnRspUserLogout").Call(q.spi, syscall.NewCallback(q.OnRspUserLogout))
    q.h.MustFindProc("qSetOnRspQryMulticastInstrument").Call(q.spi, syscall.NewCallback(q.OnRspQryMulticastInstrument))
    q.h.MustFindProc("qSetOnRspError").Call(q.spi, syscall.NewCallback(q.OnRspError))
    q.h.MustFindProc("qSetOnRspSubMarketData").Call(q.spi, syscall.NewCallback(q.OnRspSubMarketData))
    q.h.MustFindProc("qSetOnRspUnSubMarketData").Call(q.spi, syscall.NewCallback(q.OnRspUnSubMarketData))
    q.h.MustFindProc("qSetOnRspSubForQuoteRsp").Call(q.spi, syscall.NewCallback(q.OnRspSubForQuoteRsp))
    q.h.MustFindProc("qSetOnRspUnSubForQuoteRsp").Call(q.spi, syscall.NewCallback(q.OnRspUnSubForQuoteRsp))
    q.h.MustFindProc("qSetOnRtnDepthMarketData").Call(q.spi, syscall.NewCallback(q.OnRtnDepthMarketData))
    q.h.MustFindProc("qSetOnRtnForQuoteRsp").Call(q.spi, syscall.NewCallback(q.OnRtnForQuoteRsp))
    
	return q
}


// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (q *Quote) OnFrontConnected() uintptr {
    if q._FrontConnected != nil {    
	    q._FrontConnected()
    }
	return 0
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (q *Quote) OnFrontDisconnected(nReason int) uintptr {
    if q._FrontDisconnected != nil {    
	    q._FrontDisconnected(nReason)
    }
	return 0
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (q *Quote) OnHeartBeatWarning(nTimeLapse int) uintptr {
    if q._HeartBeatWarning != nil {    
	    q._HeartBeatWarning(nTimeLapse)
    }
	return 0
}

// 登录请求响应
func (q *Quote) OnRspUserLogin(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspUserLogin != nil {    
	    q._RspUserLogin(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 登出请求响应
func (q *Quote) OnRspUserLogout(pUserLogout *ctp.CThostFtdcUserLogoutField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspUserLogout != nil {    
	    q._RspUserLogout(pUserLogout, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 请求查询组播合约响应
func (q *Quote) OnRspQryMulticastInstrument(pMulticastInstrument *ctp.CThostFtdcMulticastInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspQryMulticastInstrument != nil {    
	    q._RspQryMulticastInstrument(pMulticastInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 错误应答
func (q *Quote) OnRspError(pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspError != nil {    
	    q._RspError(pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 订阅行情应答
func (q *Quote) OnRspSubMarketData(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspSubMarketData != nil {    
	    q._RspSubMarketData(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 取消订阅行情应答
func (q *Quote) OnRspUnSubMarketData(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspUnSubMarketData != nil {    
	    q._RspUnSubMarketData(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 订阅询价应答
func (q *Quote) OnRspSubForQuoteRsp(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspSubForQuoteRsp != nil {    
	    q._RspSubForQuoteRsp(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 取消订阅询价应答
func (q *Quote) OnRspUnSubForQuoteRsp(pSpecificInstrument *ctp.CThostFtdcSpecificInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if q._RspUnSubForQuoteRsp != nil {    
	    q._RspUnSubForQuoteRsp(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

// 深度行情通知
func (q *Quote) OnRtnDepthMarketData(pDepthMarketData *ctp.CThostFtdcDepthMarketDataField) uintptr {
    if q._RtnDepthMarketData != nil {    
	    q._RtnDepthMarketData(pDepthMarketData)
    }
	return 0
}

// 询价通知
func (q *Quote) OnRtnForQuoteRsp(pForQuoteRsp *ctp.CThostFtdcForQuoteRspField) uintptr {
    if q._RtnForQuoteRsp != nil {    
	    q._RtnForQuoteRsp(pForQuoteRsp)
    }
	return 0
}

