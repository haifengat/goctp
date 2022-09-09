package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.6.8_20220712
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR} -lctp_quote -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* qCreateApi();
void* qCreateSpi();
[[ range .Fn ]]// [[ .Comment ]]
void* q[[ .FuncName ]]([[ range $i, $v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ .FieldType|struct_Type ]] [[ .FieldName ]][[ if eq .FieldName "*ppInstrumentID" ]][][[ end ]][[end]]);
[[ end ]]
[[ range .On ]]// [[ .Comment ]]
void qSetOn[[ .FuncName ]](void*, void*);
int q[[ .FuncName ]]([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldType|struct_Type ]] [[ .FieldName ]][[ end ]]);
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

// Quote 行情接口
type Quote struct {
	goctp.HFQuote // 组合

	api unsafe.Pointer

    [[ range .On]]// [[ .Comment ]]
    _[[ .FuncName ]] func([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[ end ]])
    [[ end]]
}

var q *Quote

// NewQuote 实例化
func NewQuote() *Quote {
	q = new(Quote)
	
	// 主调函数封装 手动添加
	q.HFQuote.ReqConnect = func(addr string) {
		front := C.CString(addr)
		C.qRegisterFront(q.api, front)
		defer C.free(unsafe.Pointer(front))
		C.qInit(q.api)
		// C.qJoin(q.api)
	}
	q.HFQuote.ReleaseAPI = func() {
		C.qRegisterSpi(q.api, nil)
		C.qRelease(q.api)
		q.api = nil
	}
	q.HFQuote.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		C.qReqUserLogin(q.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(f)), C.int(1))
	}
	q.HFQuote.ReqSubMarketData = func(instrument ...string) {
		ppInstrumentID := make([]*C.char, len(instrument))
		for i := 0; i < len(instrument); i++ {
			ppInstrumentID[i] = (*C.char)(unsafe.Pointer(C.CString(instrument[i])))
		}
		C.qSubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&ppInstrumentID[0])), C.int(len(instrument)))
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

	q.api = C.qCreateApi()
	spi := C.qCreateSpi()
	C.qRegisterSpi(q.api, spi)

    [[ range .On ]]C.qSetOn[[ .FuncName ]](spi, C.q[[ .FuncName ]])
    [[ end ]]
	return q
}
[[ range .On ]]
//export q[[ .FuncName ]]
func q[[ .FuncName ]]([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|C_struct]][[ end ]]) C.int {
    if q._[[ .FuncName ]] != nil{
        q._[[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ ctp_param .FieldType .FieldName ]][[ end ]])
    }
	return 0
}
[[ end ]]
