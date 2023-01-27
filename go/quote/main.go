package main

import (
	"fmt"
	"os"
)

/*
// #cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../lib -Wl,-rpath ${SRCDIR}/../lib -lthostmduserapi_se -lctpquote -lstdc++
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../lib -Wl,-rpath ${SRCDIR}/../lib -l ctpquote -lstdc++

//void* _ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbb(char const*, bool, bool)
// void* _ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbb(char const*, _Bool, _Bool);
void* CreateFtdcMdApi(char const*, _Bool, _Bool);

//DLL_EXPORT void RegisterFront(CThostFtdcMdApi *api, char *pszFrontAddress)
void RegisterFront(void *, char*);

//DLL_EXPORT void Init(CThostFtdcMdApi *api, char *pszFrontAddress)
void Init(void *);

#include <stdlib.h>
#include <stdint.h>
*/
import "C"

func main() {
	path := C.CString("./log/")
	os.MkdirAll("./log/", os.ModePerm)
	// api := C._ZN15CThostFtdcMdApi15CreateFtdcMdApiEPKcbb(path, false, false)
	api := C.CreateFtdcMdApi(path, false, false)
	fmt.Println(api)

	front := C.CString("tcp://111.111.111.111:8080")
	C.RegisterFront(api, front)

	C.Init(api)
}
