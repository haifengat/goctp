package main

import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

/*
#cgo CPPFLAGS: -fPIC -I../../CTPv6.6.8_20220712
#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"

// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../lib -Wl,-rpath ${SRCDIR}/../lib -lthostmduserapi_se -lctpquote -lstdc++
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../lib -Wl,-rpath ${SRCDIR}/../lib -l ctpquote -lstdc++

//void* _ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbb(char const*, bool, bool)
// void* _ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbb(char const*, _Bool, _Bool);
void* CreateFtdcMdApi(char const*, _Bool, _Bool);

//DLL_EXPORT void RegisterFront(CThostFtdcMdApi *api, char *pszFrontAddress)
void RegisterFront(void *, char*);

//DLL_EXPORT void Init(CThostFtdcMdApi *api, char *pszFrontAddress)
void Init(void *);

//DLL_EXPORT int ReqUserLogin(CThostFtdcMdApi *api, CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID)
int ReqUserLogin(void *api, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);

void* CreateFtdcMdSpi();
void RegisterSpi(void *, void*);

void SetOnFrontConnected(void *, void *);
void exOnFrontConnected();

void SetOnRspUserLogin(void *, void *);
void exOnRspUserLogin(struct CThostFtdcRspUserLoginField*, struct CThostFtdcRspInfoField*, int, _Bool);

#include <stdlib.h>
#include <stdint.h>
*/
import "C"

func main() {
	path := C.CString("./log/")
	os.MkdirAll("./log/", os.ModePerm)
	// api := C._ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbb(path, false, false)
	var api unsafe.Pointer = C.CreateFtdcMdApi(path, false, false)
	fmt.Println(api)

	var spi unsafe.Pointer = C.CreateFtdcMdSpi()
	C.RegisterSpi(api, spi)

	C.SetOnFrontConnected(spi, C.exOnFrontConnected)
	C.SetOnRspUserLogin(spi, C.exOnRspUserLogin)

	front := C.CString("tcp://180.168.146.187:10131")
	C.RegisterFront(api, front)

	C.Init(api)

	time.Sleep(time.Second)

	var f = struct {
		// 交易日
		TradingDay [9]byte
		// 经纪公司代码
		BrokerID [11]byte
		// 用户代码
		UserID [13]byte
		// 密码
		Password [41]byte
		// 用户端产品信息
		UserProductInfo [11]byte
		// 接口端产品信息
		InterfaceProductInfo [11]byte
		// 协议信息
		ProtocolInfo [11]byte
		// Mac地址
		MacAddress [21]byte
		// 动态密码
		OneTimePassword [41]byte
		// 保留的无效字段
		reserve1 [16]byte
		// 登录备注
		LoginRemark [36]byte
		// 终端IP端口
		ClientIPPort int32
		// 终端IP地址
		ClientIPAddress [33]byte
	}{}
	copy(f.BrokerID[:], "9999")
	copy(f.UserID[:], "008105")
	copy(f.Password[:], "1")

	nRequestID := C.int(1)
	C.ReqUserLogin(api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(&f)), nRequestID)

	select {}
}

//export exOnFrontConnected
func exOnFrontConnected() {
	fmt.Println("行情接口连接")
}

//export exOnRspUserLogin
func exOnRspUserLogin(loginField *C.struct_CThostFtdcRspUserLoginField, rspInfo *C.struct_CThostFtdcRspInfoField, requestID C.int, isLast C._Bool) {
	fmt.Println("登录")
}
