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

type trade struct {
	goctp.HFTrade
	h        *syscall.DLL
	api, spi uintptr

	passWord string
}

func (t *trade) loadDll() {
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	_ = os.Chdir(dllPath)
	t.h = syscall.MustLoadDLL("ctp_trade.dll")

	// 还原到之前的工作目录
	os.Chdir(workPath)
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 接口
func newTrade() *trade {
	t := new(trade)

	t.HFTrade.ReqConnect = func(addr string) {
		bs, _ := syscall.BytePtrFromString(addr)
		t.h.MustFindProc("RegisterFront").Call(t.api, uintptr(unsafe.Pointer(bs)))
		t.h.MustFindProc("SubscribePrivateTopic").Call(t.api, uintptr(ctp.THOST_TERT_RESTART))
		t.h.MustFindProc("SubscribePublicTopic").Call(t.api, uintptr(ctp.THOST_TERT_RESTART))
		t.h.MustFindProc("Init").Call(t.api)
		// C.Join(t.api)
	}
	t.HFTrade.ReleaseAPI = func() {
		t.h.MustFindProc("Release").Call(t.api)
	}
	t.HFTrade.ReqAuthenticate = func(f *ctp.CThostFtdcReqAuthenticateField, i int) {
		t.h.MustFindProc("ReqAuthenticate").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		t.h.MustFindProc("ReqUserLogin").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqSettlementInfoConfirm = func(f *ctp.CThostFtdcSettlementInfoConfirmField, i int) {
		t.h.MustFindProc("ReqSettlementInfoConfirm").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqQryInstrument = func(f *ctp.CThostFtdcQryInstrumentField, i int) {
		t.h.MustFindProc("ReqQryInstrument").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqQryClassifiedInstrument = func(f *ctp.CThostFtdcQryClassifiedInstrumentField, i int) {
		t.h.MustFindProc("ReqQryClassifiedInstrument").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqQryTradingAccount = func(f *ctp.CThostFtdcQryTradingAccountField, i int) {
		t.h.MustFindProc("ReqQryTradingAccount").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqQryInvestorPosition = func(f *ctp.CThostFtdcQryInvestorPositionField, i int) {
		t.h.MustFindProc("ReqQryInvestorPosition").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqOrder = func(f *ctp.CThostFtdcInputOrderField, i int) {
		t.h.MustFindProc("ReqOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqAction = func(f *ctp.CThostFtdcInputOrderActionField, i int) {
		t.h.MustFindProc("ReqOrderAction").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqFromBankToFutureByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		t.h.MustFindProc("ReqFromBankToFutureByFuture").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.ReqFromFutureToBankByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		t.h.MustFindProc("ReqFromFutureToBankByFuture").Call(t.api, uintptr(unsafe.Pointer(&f)), uintptr(i))
	}
	t.HFTrade.Init()

	t.loadDll()
	t.api, _, _ = t.h.MustFindProc("CreateApi").Call()
	t.spi, _, _ = t.h.MustFindProc("CreateSpi").Call()
	t.h.MustFindProc("RegisterSpi").Call(t.api, uintptr(unsafe.Pointer(t.spi)))

	t.h.MustFindProc("SetOnFrontConnected").Call(t.spi, syscall.NewCallback(t.OnFrontConnected))
	t.h.MustFindProc("SetOnFrontDisconnected").Call(t.spi, syscall.NewCallback(t.OnFrontDisconnected))
	t.h.MustFindProc("SetOnRspAuthenticate").Call(t.spi, syscall.NewCallback(t.OnRspAuthenticate))
	t.h.MustFindProc("SetOnRspUserLogin").Call(t.spi, syscall.NewCallback(t.OnRspUserLogin))
	t.h.MustFindProc("SetOnRspOrderInsert").Call(t.spi, syscall.NewCallback(t.OnRspOrderInsert))
	t.h.MustFindProc("SetOnRspOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspOrderAction))
	t.h.MustFindProc("SetOnRspSettlementInfoConfirm").Call(t.spi, syscall.NewCallback(t.OnRspSettlementInfoConfirm))
	t.h.MustFindProc("SetOnRspQryOrder").Call(t.spi, syscall.NewCallback(t.OnRspQryOrder))
	t.h.MustFindProc("SetOnRspQryTrade").Call(t.spi, syscall.NewCallback(t.OnRspQryTrade))
	t.h.MustFindProc("SetOnRspQryInvestorPosition").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestorPosition))
	t.h.MustFindProc("SetOnRspQryTradingAccount").Call(t.spi, syscall.NewCallback(t.OnRspQryTradingAccount))
	t.h.MustFindProc("SetOnRtnOrder").Call(t.spi, syscall.NewCallback(t.OnRtnOrder))
	t.h.MustFindProc("SetOnRtnTrade").Call(t.spi, syscall.NewCallback(t.OnRtnTrade))
	t.h.MustFindProc("SetOnErrRtnOrderInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnOrderInsert))
	t.h.MustFindProc("SetOnErrRtnOrderAction").Call(t.spi, syscall.NewCallback(t.OnErrRtnOrderAction))
	t.h.MustFindProc("SetOnRtnInstrumentStatus").Call(t.spi, syscall.NewCallback(t.OnRtnInstrumentStatus))
	t.h.MustFindProc("SetOnRtnFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnFromBankToFutureByFuture))
	t.h.MustFindProc("SetOnRtnFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnFromFutureToBankByFuture))
	return t
}

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (t *trade) OnFrontConnected() uintptr {
	t.HFTrade.FrontConnected()
	return 0
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (t *trade) OnFrontDisconnected(reason int) uintptr {
	t.HFTrade.FrontDisConnected(reason)
	return 0
}

// 客户端认证响应
func (t *trade) OnRspAuthenticate(field *ctp.CThostFtdcRspAuthenticateField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.HFTrade.RspAuthenticate(info)
	return 0
}

// 登录请求响应
func (t *trade) OnRspUserLogin(field *ctp.CThostFtdcRspUserLoginField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.HFTrade.RspUserLogin(field, info)
	return 0
}

// 报单录入请求响应
func (t *trade) OnRspOrderInsert(field *ctp.CThostFtdcInputOrderField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.HFTrade.ErrRtnOrderInsert(field, info)
	return 0
}

// 报单操作请求响应
func (t *trade) OnRspOrderAction(field *ctp.CThostFtdcInputOrderActionField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	// t.HFTrade.ErrRtnOrderAction(field, info)
	return 0
}

// 投资者结算结果确认响应
func (t *trade) OnRspSettlementInfoConfirm(field *ctp.CThostFtdcSettlementInfoConfirmField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.HFTrade.RspSettlementInfoConfirm()
	return 0
}

// 请求查询投资者持仓响应
func (t *trade) OnRspQryInvestorPosition(field *ctp.CThostFtdcInvestorPositionField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.HFTrade.RspQryInvestorPosition(field, b)
	return 0
}

// 请求查询资金账户响应
func (t *trade) OnRspQryTradingAccount(field *ctp.CThostFtdcTradingAccountField, info *ctp.CThostFtdcRspInfoField, i int, b bool) uintptr {
	t.HFTrade.RspQryTradingAccount(field)
	return 0
}

// 报单通知
func (t *trade) OnRtnOrder(field *ctp.CThostFtdcOrderField) uintptr {
	t.HFTrade.RtnOrder(field)
	return 0
}

// 成交通知
func (t *trade) OnRtnTrade(field *ctp.CThostFtdcTradeField) uintptr {
	t.HFTrade.RtnTrade(field)
	return 0
}

// 报单录入错误回报
func (t *trade) OnErrRtnOrderInsert(field *ctp.CThostFtdcInputOrderField, info *ctp.CThostFtdcRspInfoField) uintptr {
	t.HFTrade.ErrRtnOrderInsert(field, info)
	return 0
}

// 报单操作错误回报
func (t *trade) OnErrRtnOrderAction(field *ctp.CThostFtdcOrderActionField, info *ctp.CThostFtdcRspInfoField) uintptr {
	t.HFTrade.ErrRtnOrderAction(field, info)
	return 0
}

// 合约交易状态通知
func (t *trade) OnRtnInstrumentStatus(field *ctp.CThostFtdcInstrumentStatusField) uintptr {
	t.HFTrade.RtnInstrumentStatus(field)
	return 0
}

// 期货发起银行资金转期货通知
func (t *trade) OnRtnFromBankToFutureByFuture(field *ctp.CThostFtdcRspTransferField) uintptr {
	t.HFTrade.RtnFromBankToFutureByFuture(field)
	return 0
}

// 期货发起期货资金转银行通知
func (t *trade) OnRtnFromFutureToBankByFuture(field *ctp.CThostFtdcRspTransferField) uintptr {
	t.HFTrade.RtnFromFutureToBankByFuture(field)
	return 0
}
