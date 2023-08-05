package goctp

/*
#cgo CPPFLAGS: -fPIC -I./CTPv6.6.9_20220922
#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"

#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/lib -Wl,-rpath ${SRCDIR}/lib -l ctptrade -lstdc++

void* CreateFtdcTraderApi(char const*);
void* CreateFtdcTraderSpi();
void* GetVersion();

[[ range .Fn ]]// [[ .Comment ]]
[[ .RtnType ]] t[[ .Name ]](void *api[[ range .Params ]], [[ .Type|toCGo ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ if eq .Var "ppInstrumentID" ]][][[ end ]][[ end ]]);
[[ end ]]
[[ range .On ]]// [[ .Comment ]]
void tSet[[ .Name ]](void *, void *);
void [[ .Name ]](void*[[ range .Params ]], [[ .Type|toCGo ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ if eq .Var "ppInstrumentID" ]][][[ end ]][[ end ]]);
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

type Trade struct {
	api, spi unsafe.Pointer
	Version string

	// ************ 响应函数变量 ******************
	[[ range .On]]
	// [[ .Comment ]]
	[[ .Name ]] func([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]] [[ toGoType .Type .Var ]][[ end ]])
	[[- end]]
}

var Trades = make(map[unsafe.Pointer]*Trade)

func NewTrade() *Trade {
    t := &Trade{}
	path := C.CString("./log/")
	os.MkdirAll("./log/", os.ModePerm)

	t.api = C.CreateFtdcTraderApi(path)
	t.Version = C.GoString((*C.char)(C.GetVersion()))
	fmt.Println(t.Version)

	t.spi  = C.CreateFtdcTraderSpi()
	C.tRegisterSpi(t.api, t.spi)

    [[ range .On -]]	
	C.tSet[[ .Name ]](t.spi, C.[[ .Name ]]) // [[ .Comment ]]
    [[ end ]]
	Trades[t.spi] = t
    return t
}

[[ range .On -]]
//export [[ .Name ]]
func [[ .Name ]](spi unsafe.Pointer[[ range .Params ]], [[ .Var ]] [[ if .HasStar ]]*[[ end ]][[ .Type|exToCGo ]][[ end ]]) {
	t := Trades[spi]
	if t.[[ .Name ]] == nil {
		fmt.Println("[[ .Name ]]")
	} else {
		t.[[ .Name ]]([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ onVar .Type .Var ]][[ end ]])
	}
}
[[ end ]]

[[ range .Fn ]]// [[ .Comment ]]
func (t *Trade)[[ .Name ]]([[ range $idx, $p := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]] [[ toGoType .Type .Var ]][[ end ]]) [[ if eq .RtnType "int" ]]int[[ end ]]{
	[[ if eq .RtnType "int" ]]return int([[ end ]]C.t[[ .Name ]](t.api[[ range .Params ]], [[ fnVar .Type .Var ]][[ end ]])[[ if eq .RtnType "int" ]])[[ end ]]
}
[[ end ]]