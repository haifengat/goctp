package goctp

/*
#cgo CPPFLAGS: -fPIC -I./CTPv6.6.8_20220712
#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"

#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/lib -Wl,-rpath ${SRCDIR}/lib -l ctpquote -lstdc++

void* CreateFtdcMdApi(char const*, _Bool, _Bool);
void* CreateFtdcMdSpi();

[[ range .Fn ]]// [[ .Comment ]]
[[ .RtnType ]] [[ .Name ]](void *api[[ range .Params ]], [[ .Type|toCGo ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ if eq .Var "ppInstrumentID" ]][][[ end ]][[ end ]]);
[[ end ]]
[[ range .On ]]// [[ .Comment ]]
void Set[[ .Name ]](void *, void *);
void ex[[ .Name ]]([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end]][[ .Type|toCGo ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ if eq .Var "ppInstrumentID" ]][][[ end ]][[ end ]]);
[[ end ]]

#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

type Quote struct {
	api, spi unsafe.Pointer
	// ************ 响应函数变量 ******************
	[[ range .On -]]
	// [[ .Comment ]]
	[[ .Name ]] func([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]] [[ toGoType .Type .Var ]][[ end ]])
	[[- end]]
}

var q *Quote

func NewQuote() *Quote {
    q = &Quote{}
	path := C.CString("./log/")
	os.MkdirAll("./log/", os.ModePerm)

	q.api = C.CreateFtdcMdApi(path, false, false)

	q.spi  = C.CreateFtdcMdSpi()
	C.RegisterSpi(q.api, q.spi)

    [[ range .On -]]	
	C.Set[[ .Name ]](q.spi, C.ex[[ .Name ]]) // [[ .Comment ]]
    [[ end ]]
    return q
}

[[ range .On -]]
// [[ .Comment ]]
//
//export ex[[ .Name ]]
func ex[[ .Name ]]([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]] [[ if .HasStar ]]*[[ end ]][[ .Type|exToCGo ]][[ end ]]) {
	if q.[[ .Name ]] == nil {
		fmt.Println("[[ .Name ]]")
	} else {
		q.[[ .Name ]]([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ onVar .Type .Var ]][[ end ]])
	}
}
[[ end ]]

[[ range .Fn ]]// [[ .Comment ]]
func (q *Quote)[[ .Name ]]([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]] [[ toGoType .Type .Var ]][[ end ]]){
    [[ if and (gt (len .Params) 0) (eq (index .Params 0).Var "ppInstrumentID") -]]
	instruments := make([]*C.char, len(ppInstrumentID))
	for i := 0; i < len(instruments); i++ {
		instruments[i] = (*C.char)(unsafe.Pointer(C.CString(ppInstrumentID[i])))
	}
	C.[[ .Name ]](q.api, (**C.char)(unsafe.Pointer(&instruments[0])), C.int(len(instruments)))
	[[- else -]]
	C.[[ .Name ]](q.api[[ range .Params ]], [[ fnVar .Type .Var ]][[ end ]])
	[[- end ]]
}
[[ end ]]