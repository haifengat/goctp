package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.6.8_20220712
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR} -lctp_trade -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* tCreateApi();
void* tCreateSpi();
void* tGetVersion();
[[ range .Fn ]]// [[ .Comment ]]
void* t[[ .FuncName ]]([[ range $i, $v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ .FieldType|struct_Type ]] [[ .FieldName ]][[ if eq .FieldName "*ppInstrumentID" ]][][[ end ]][[end]]);
[[ end ]]
[[ range .On ]]// [[ .Comment ]]
void tSetOn[[ .FuncName ]](void*, void*);
int t[[ .FuncName ]]([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldType|struct_Type ]] [[ .FieldName ]][[ end ]]);
[[ end ]]
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

    [[ range .On]]// [[ .Comment ]]
    _[[ .FuncName ]] func([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[ end ]])
    [[ end]]
}

var t *Trade

// NewTrade 实例化
func NewTrade() *Trade {
	t = new(Trade)

	// 主调函数封装 手动添加
	t.HFTrade.GetVersion = func() string {
		return C.GoString((*C.char)(C.tGetVersion()))
	}
	t.HFTrade.ReqConnect = func(addr string) {
		front := C.CString(addr)
		C.tRegisterFront(t.api, front)
		defer C.free(unsafe.Pointer(front))
		C.tSubscribePrivateTopic(t.api, C.int(t.PrivateMode))
		C.tSubscribePublicTopic(t.api, C.int(ctp.THOST_TERT_RESTART)) // 公有流不能用 quick, 会导致交易所状态不更新
		C.tInit(t.api)
		// C.Join(t.api)
	}
	t.HFTrade.ReleaseAPI = func() {
		// C.RegisterSpi(t.api, nil) // 6.6.1说明中提到,会导致程序崩溃
		C.tRelease(t.api) // 前置开着,后台关闭报错. DesignError:pthread_mutex_unlock in line 116 of file ../../source/event/Mutex.h
		// 若不release 则会返回n个 4097后,程序崩溃
		t.spi = nil
		t.api = nil
	}
	t.HFTrade.ReqAuthenticate = func(f *ctp.CThostFtdcReqAuthenticateField, i int) {
		C.tReqAuthenticate(t.api, (*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		C.tReqUserLogin(t.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqOrder = func(f *ctp.CThostFtdcInputOrderField, i int) {
		C.tReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqAction = func(f *ctp.CThostFtdcInputOrderActionField, i int) {
		C.tReqOrderAction(t.api, (*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqFromBankToFutureByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		C.tReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqFromFutureToBankByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		C.tReqFromFutureToBankByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqSettlementInfoConfirm = func(f *ctp.CThostFtdcSettlementInfoConfirmField, i int) {
		C.tReqSettlementInfoConfirm(t.api, (*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInstrument = func(f *ctp.CThostFtdcQryInstrumentField, i int) {
		C.tReqQryInstrument(t.api, (*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryClassifiedInstrument = func(f *ctp.CThostFtdcQryClassifiedInstrumentField, i int) {
		C.tReqQryClassifiedInstrument(t.api, (*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryTradingAccount = func(f *ctp.CThostFtdcQryTradingAccountField, i int) {
		C.tReqQryTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInvestorPosition = func(f *ctp.CThostFtdcQryInvestorPositionField, i int) {
		C.tReqQryInvestorPosition(t.api, (*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInvestor = func(f *ctp.CThostFtdcQryInvestorField, i int) {
		C.tReqQryInvestor(t.api, (*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryOrder = func(f *ctp.CThostFtdcQryOrderField, i int) {
		C.tReqQryOrder(t.api, (*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryTrade = func(f *ctp.CThostFtdcQryTradeField, i int) {
		C.tReqQryTrade(t.api, (*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(f)), C.int(i))
	}
	
	// HFTrade 响应 手动添加即可增加新功能
	t._RtnFromFutureToBankByFuture = func(pRspTransfer *ctp.CThostFtdcRspTransferField) {
		t.HFTrade.RtnFromFutureToBankByFuture(pRspTransfer)
	}
	t._RtnFromBankToFutureByFuture = func(pRspTransfer *ctp.CThostFtdcRspTransferField) {
		t.HFTrade.RtnFromBankToFutureByFuture(pRspTransfer)
	}
	t._ErrRtnOrderAction = func(pOrderAction *ctp.CThostFtdcOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField) {
		t.HFTrade.ErrRtnOrderAction(pOrderAction, pRspInfo)
	}
	t._ErrRtnOrderInsert = func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField) {
		t.HFTrade.ErrRtnOrderInsert(pInputOrder, pRspInfo)
	}
	t._RtnTrade = func(pTrade *ctp.CThostFtdcTradeField) {
		t.HFTrade.RtnTrade(pTrade)
	}
	t._RtnOrder = func(pOrder *ctp.CThostFtdcOrderField) {
		t.HFTrade.RtnOrder(pOrder)
	}
	t._RtnInstrumentStatus = func(pInstrumentStatus *ctp.CThostFtdcInstrumentStatusField) {
		t.HFTrade.RtnInstrumentStatus(pInstrumentStatus)
	}
	t._RspQryInvestorPosition = func(pInvestorPosition *ctp.CThostFtdcInvestorPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInvestorPosition(pInvestorPosition, bIsLast)
	}
	t._RspQryTradingAccount = func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryTradingAccount(pTradingAccount, bIsLast)
	}
	t._RspQryTrade = func(pTrade *ctp.CThostFtdcTradeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryTrade(pTrade, bIsLast)
	}
	t._RspQryOrder = func(pOrder *ctp.CThostFtdcOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryOrder(pOrder, bIsLast)
	}
	t._RspQryInvestor = func(pInvestor *ctp.CThostFtdcInvestorField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInvestor(pInvestor, bIsLast)
	}
	t._RspQryInstrument = func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInstrument(pInstrument, bIsLast)
	}
	t._RspQryClassifiedInstrument = func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInstrument(pInstrument, bIsLast)
	}
	t._RspSettlementInfoConfirm = func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspSettlementInfoConfirm()
	}
	t._RspUserLogin = func(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspUserLogin(pRspUserLogin, pRspInfo)
	}
	t._RspAuthenticate = func(pRspAuthenticateField *ctp.CThostFtdcRspAuthenticateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspAuthenticate(pRspInfo)
	}
	t._FrontConnected = func() {
		t.HFTrade.FrontConnected()
	}
	t._FrontDisconnected = func(nReason int) {
		t.HFTrade.FrontDisConnected(nReason)
	}

	t.HFTrade.Init() // 初始化

	t.api = C.tCreateApi()
	t.spi = C.tCreateSpi()
	C.tRegisterSpi(t.api, t.spi)

    [[ range .On ]]C.tSetOn[[ .FuncName ]](t.spi, C.t[[ .FuncName ]])
    [[ end ]]
	return t
}
[[ range .On ]]
//export t[[ .FuncName ]]
func t[[ .FuncName ]]([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|C_struct]][[ end ]]) C.int {
    if t._[[ .FuncName ]] != nil{
        t._[[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ ctp_param .FieldType .FieldName ]][[ end ]])
    }
	return 0
}
[[ end ]]
