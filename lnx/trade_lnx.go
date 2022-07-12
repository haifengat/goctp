package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.5.1_20200908
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR}  -lctp_trade -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* CreateApi();
void* CreateSpi();
void* RegisterSpi(void*, void*);
void* RegisterFront(void*, char*);
void* SubscribePublicTopic(void*, int);
void* SubscribePrivateTopic(void*, int);
void* Init(void*);
void* Join(void*);
void* GetVersion();
void* Release(void*);
void* ReqAuthenticate(void*, struct CThostFtdcReqAuthenticateField*, int);
void* ReqUserLogin(void*, struct CThostFtdcReqUserLoginField*, int);
void* ReqSettlementInfoConfirm(void*, struct CThostFtdcSettlementInfoConfirmField*, int);
void* ReqQryTradingAccount(void*, struct CThostFtdcQryTradingAccountField*, int);
void* ReqQryInvestorPosition(void*, struct CThostFtdcQryInvestorPositionField*, int);
void* ReqQryInstrument(void*, struct CThostFtdcQryInstrumentField*, int);
void* ReqQryClassifiedInstrument(void*, struct CThostFtdcQryClassifiedInstrumentField*, int);
void* ReqOrderInsert(void*, struct CThostFtdcInputOrderField*, int);
void* ReqOrderAction(void*, struct CThostFtdcInputOrderActionField*, int);
void* ReqFromBankToFutureByFuture(void*, struct CThostFtdcReqTransferField *, int);
void* ReqFromFutureToBankByFuture(void*, struct CThostFtdcReqTransferField *, int);

// 查询用户信息(交易员查询)
void* ReqQryInvestor(void*, struct CThostFtdcQryInvestorField*, int);
void SetOnRspQryInvestor(void*, void*);
int tRspQryInvestor(struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);

void SetOnFrontConnected(void*, void*);
int tFrontConnected();
void SetOnFrontDisconnected(void*, void*);
int tFrontDisConnected(int);
void SetOnRspUserLogin(void*, void*);
int tRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspAuthenticate(void*, void*);
int tRspAuthenticate(struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspSettlementInfoConfirm(void*, void*);
int tRspSettlementInfoConfirm(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspQryInstrument(void*, void*);
int tRspQryInstrument(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspQryClassifiedInstrument(void*, void*);
int tRspQryClassifiedInstrument(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspQryTradingAccount(void*, void*);
int tRspQryTradingAccount(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspQryInvestorPosition(void*, void*);
int tRspQryInvestorPosition(struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRspOrderInsert(void*, void*);
int tRspOrderInsert(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnErrRtnOrderInsert(void*, void*);
int tErrRtnOrderInsert(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);
void SetOnRtnOrder(void*, void*);
int tRtnOrder(struct CThostFtdcOrderField *pOrder);
void SetOnRtnTrade(void*, void*);
int tRtnTrade(struct CThostFtdcTradeField *pTrade);
void SetOnRtnInstrumentStatus(void*, void*);
int tRtnInstrumentStatus(struct CThostFtdcInstrumentStatusField *pInstrumentStatus);
void SetOnErrRtnOrderAction(void*, void*);
int tErrRtnOrderAction(struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
void SetOnRtnFromFutureToBankByFuture(void*, void*);
int tRtnFromFutureToBankByFuture(struct CThostFtdcRspTransferField *pRspTransfer);
void SetOnRtnFromBankToFutureByFuture(void*, void*);
int tRtnFromBankToFutureByFuture(struct CThostFtdcRspTransferField *pRspTransfer);
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"unsafe"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

// Trade 交易接口
type Trade struct {
	goctp.HFTrade
	spi, api unsafe.Pointer

	passWord string // 密码
}

var t *Trade
var mode = ctp.THOST_TERT_RESTART

// SetQuick 以quick模式启动(须在NewTrade前调用)
func SetQuick() {
	mode = ctp.THOST_TERT_QUICK
}

// NewTrade 实例化
func NewTrade() *Trade {
	t = new(Trade)
	t.HFTrade.GetVersion = func() string {
		return C.GoString((*C.char)(C.GetVersion()))
	}
	t.HFTrade.ReqConnect = func(addr string) {
		front := C.CString(addr)
		C.RegisterFront(t.api, front)
		defer C.free(unsafe.Pointer(front))
		C.SubscribePrivateTopic(t.api, C.int(mode))
		C.SubscribePublicTopic(t.api, C.int(mode))
		C.Init(t.api)
		// C.Join(t.api)
	}
	t.HFTrade.ReleaseAPI = func() {
		C.RegisterSpi(t.api, nil)
		t.spi = nil
		C.Release(t.api)
		t.api = nil
	}
	t.HFTrade.ReqAuthenticate = func(f *ctp.CThostFtdcReqAuthenticateField, i int) {
		C.ReqAuthenticate(t.api, (*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		C.ReqUserLogin(t.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqSettlementInfoConfirm = func(f *ctp.CThostFtdcSettlementInfoConfirmField, i int) {
		C.ReqSettlementInfoConfirm(t.api, (*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInstrument = func(f *ctp.CThostFtdcQryInstrumentField, i int) {
		C.ReqQryInstrument(t.api, (*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryClassifiedInstrument = func(f *ctp.CThostFtdcQryClassifiedInstrumentField, i int) {
		C.ReqQryClassifiedInstrument(t.api, (*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryTradingAccount = func(f *ctp.CThostFtdcQryTradingAccountField, i int) {
		C.ReqQryTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInvestorPosition = func(f *ctp.CThostFtdcQryInvestorPositionField, i int) {
		C.ReqQryInvestorPosition(t.api, (*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqOrder = func(f *ctp.CThostFtdcInputOrderField, i int) {
		C.ReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqAction = func(f *ctp.CThostFtdcInputOrderActionField, i int) {
		C.ReqOrderAction(t.api, (*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(&f)), C.int(i))
	}
	t.HFTrade.ReqFromBankToFutureByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		C.ReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqFromFutureToBankByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		C.ReqFromFutureToBankByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(f)), C.int(i))
	}
	// 查询用户信息
	t.HFTrade.ReqQryInvestor = func(f *ctp.CThostFtdcQryInvestorField, i int) {
		C.ReqQryInvestor(t.api, (*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(f)), C.int(i))
	}

	t.HFTrade.Init() // 初始化

	t.api = C.CreateApi()
	t.spi = C.CreateSpi()
	C.RegisterSpi(t.api, t.spi)

	C.SetOnFrontConnected(t.spi, C.tFrontConnected)
	C.SetOnFrontDisconnected(t.spi, C.tFrontDisConnected)
	C.SetOnRspUserLogin(t.spi, C.tRspUserLogin)
	C.SetOnRspAuthenticate(t.spi, C.tRspAuthenticate)
	C.SetOnRspSettlementInfoConfirm(t.spi, C.tRspSettlementInfoConfirm)
	C.SetOnRspQryInstrument(t.spi, C.tRspQryInstrument)
	C.SetOnRspQryClassifiedInstrument(t.spi, C.tRspQryClassifiedInstrument)
	C.SetOnRspQryTradingAccount(t.spi, C.tRspQryTradingAccount)
	C.SetOnRspQryInvestorPosition(t.spi, C.tRspQryInvestorPosition)
	C.SetOnRspOrderInsert(t.spi, C.tRspOrderInsert)
	C.SetOnRtnOrder(t.spi, C.tRtnOrder)
	C.SetOnRtnTrade(t.spi, C.tRtnTrade)
	C.SetOnErrRtnOrderInsert(t.spi, C.tErrRtnOrderInsert)
	C.SetOnErrRtnOrderAction(t.spi, C.tErrRtnOrderAction)
	C.SetOnRtnInstrumentStatus(t.spi, C.tRtnInstrumentStatus)
	C.SetOnRtnFromBankToFutureByFuture(t.spi, C.tRtnFromBankToFutureByFuture)
	C.SetOnRtnFromFutureToBankByFuture(t.spi, C.tRtnFromFutureToBankByFuture)

	C.SetOnRspQryInvestor(t.spi, C.tRspQryInvestor)
	return t
}

//export tRtnFromBankToFutureByFuture
func tRtnFromBankToFutureByFuture(field *C.struct_CThostFtdcRspTransferField) C.int {
	transField := (*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(field))
	t.HFTrade.RtnFromBankToFutureByFuture(transField)
	return 0
}

//export tRtnFromFutureToBankByFuture
func tRtnFromFutureToBankByFuture(field *C.struct_CThostFtdcRspTransferField) C.int {
	transField := (*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(field))
	t.HFTrade.RtnFromFutureToBankByFuture(transField)
	return 0
}

//export tRtnInstrumentStatus
func tRtnInstrumentStatus(field *C.struct_CThostFtdcInstrumentStatusField) C.int {
	statusField := (*ctp.CThostFtdcInstrumentStatusField)(unsafe.Pointer(field))
	t.HFTrade.RtnInstrumentStatus(statusField)
	return 0
}

//export tRtnTrade
func tRtnTrade(field *C.struct_CThostFtdcTradeField) C.int {
	tradeField := (*ctp.CThostFtdcTradeField)(unsafe.Pointer(field))
	t.HFTrade.RtnTrade(tradeField)
	return 0
}

//export tRtnOrder
func tRtnOrder(field *C.struct_CThostFtdcOrderField) C.int {
	orderField := (*ctp.CThostFtdcOrderField)(unsafe.Pointer(field))
	t.HFTrade.RtnOrder(orderField)
	return 0
}

//export tErrRtnOrderAction
func tErrRtnOrderAction(field *C.struct_CThostFtdcOrderActionField, info *C.struct_CThostFtdcRspInfoField) C.int {
	actionField := (*ctp.CThostFtdcOrderActionField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.HFTrade.ErrRtnOrderAction(actionField, infoField)
	return 0
}

//export tRspOrderInsert
func tRspOrderInsert(field *C.struct_CThostFtdcInputOrderField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	orderField := (*ctp.CThostFtdcInputOrderField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.HFTrade.ErrRtnOrderInsert(orderField, infoField)
	return 0
}

//export tErrRtnOrderInsert
func tErrRtnOrderInsert(field *C.struct_CThostFtdcInputOrderField, info *C.struct_CThostFtdcRspInfoField) C.int {
	orderField := (*ctp.CThostFtdcInputOrderField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.HFTrade.ErrRtnOrderInsert(orderField, infoField)
	return 0
}

//export tRspQryInvestorPosition
func tRspQryInvestorPosition(field *C.struct_CThostFtdcInvestorPositionField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	posiField := (*ctp.CThostFtdcInvestorPositionField)(unsafe.Pointer(field))
	t.HFTrade.RspQryInvestorPosition(posiField, bool(b))
	return 0
}

//export tRspQryTradingAccount
func tRspQryTradingAccount(field *C.struct_CThostFtdcTradingAccountField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	accountField := (*ctp.CThostFtdcTradingAccountField)(unsafe.Pointer(field))
	t.HFTrade.RspQryTradingAccount(accountField)
	return 0
}

//export tRspQryClassifiedInstrument
func tRspQryClassifiedInstrument(field *C.struct_CThostFtdcInstrumentField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	instrumentField := (*ctp.CThostFtdcInstrumentField)(unsafe.Pointer(field))
	t.HFTrade.RspQryInstrument(instrumentField, bool(b))
	return 0
}

//export tRspQryInstrument
func tRspQryInstrument(field *C.struct_CThostFtdcInstrumentField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	instrumentField := (*ctp.CThostFtdcInstrumentField)(unsafe.Pointer(field))
	t.HFTrade.RspQryInstrument(instrumentField, bool(b))
	return 0
}

//export tRspSettlementInfoConfirm
func tRspSettlementInfoConfirm(field *C.struct_CThostFtdcSettlementInfoConfirmField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	t.HFTrade.RspSettlementInfoConfirm()
	return 0
}

//export tRspUserLogin
func tRspUserLogin(field *C.struct_CThostFtdcRspUserLoginField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	loginField := (*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.HFTrade.RspUserLogin(loginField, infoField)
	return 0
}

//export tRspAuthenticate
func tRspAuthenticate(field *C.struct_CThostFtdcRspAuthenticateField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	//authField := (* ctp.CThostFtdcRspAuthenticateField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.HFTrade.RspAuthenticate(infoField)
	return 0
}

//export tFrontDisConnected
func tFrontDisConnected(reason C.int) C.int {
	t.HFTrade.FrontDisConnected(int(reason))
	return 0
}

//export tFrontConnected
func tFrontConnected() C.int {
	t.HFTrade.FrontConnected()
	return 0
}

//export tRspQryInvestor
func tRspQryInvestor(field *C.struct_CThostFtdcInvestorField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	investorField := (*ctp.CThostFtdcInvestorField)(unsafe.Pointer(field))
	t.HFTrade.RspQryInvestor(investorField, bool(b))
	return 0
}
