package src

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"syscall"
	"unsafe"
)

var (
	h        *syscall.DLL
	api, spi uintptr
)

type CThostFtdcReqUserLoginField struct {
	TradingDay           [9]byte
	BrokerID             [11]byte
	UserID               [16]byte
	Password             [41]byte
	UserProductInfo      [11]byte
	InterfaceProductInfo [11]byte
	ProtocolInfo         [11]byte
	MacAddress           [21]byte
	OneTimePassword      [41]byte
	ClientIPAddress      [16]byte
	LoginRemark          [36]byte
	ClientIPPort         int32
}

func OnConnect() uintptr {
	fmt.Println("connected")
	pReqUserLogin := h.MustFindProc("ReqUserLogin")
	var a [16]byte
	copy(a[:], "008105")
	var b [41]byte
	copy(b[:], "1")
	var c [11]byte
	copy(c[:], "9999")
	var pLogin = CThostFtdcReqUserLoginField{
		UserID:   a,
		Password: b,
		BrokerID: c,
	}

	_, _, err := pReqUserLogin.Call(api, uintptr(unsafe.Pointer(&pLogin)), 1)
	fmt.Println(err)
	return 0
}

type CThostFtdcRspInfoField struct {
	ErrorID  int32
	ErrorMsg [81]byte
}

func OnUserLogin(login uintptr, info uintptr, id int, last bool) uintptr {
	fmt.Println(login)
	//buf:=make([]byte, unsafe.Sizeof(info))
	i := (*CThostFtdcRspInfoField)(unsafe.Pointer(info))

	bs, err := simplifiedchinese.GB18030.NewDecoder().Bytes(i.ErrorMsg[:bytes.IndexByte(i.ErrorMsg[:], 0)])
	fmt.Println(err)
	fmt.Println(string(bs))
	fmt.Println("login succeed")
	return 0
}

func Text() {
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}

	h = syscall.MustLoadDLL("ctp_trade.dll")
	defer h.Release()
	pCreateApi := h.MustFindProc("CreateApi")
	pCreateSpi := h.MustFindProc("CreateSpi")
	pRegisterSpi := h.MustFindProc("RegisterSpi")
	pRegisterFront := h.MustFindProc("RegisterFront")
	pSubscribePrivateTopic := h.MustFindProc("SubscribePrivateTopic")
	pSubscribePublicTopic := h.MustFindProc("SubscribePublicTopic")
	pInit := h.MustFindProc("Init")

	api, _, err = pCreateApi.Call()
	spi, _, err = pCreateSpi.Call()

	_, _, err = pRegisterSpi.Call(api, spi)

	pSetOnFrontConnected := h.MustFindProc("SetOnFrontConnected")
	_, _, err = pSetOnFrontConnected.Call(spi, syscall.NewCallback(OnConnect))
	pSetOnRspUserLogin := h.MustFindProc("SetOnRspUserLogin")
	_, _, err = pSetOnRspUserLogin.Call(spi, syscall.NewCallback(OnUserLogin))

	addr := "tcp://180.168.146.187:10000"
	// 用byteptr 而不是utf16ptr
	bs, err := syscall.BytePtrFromString(addr)
	_, _, err = pRegisterFront.Call(api, uintptr(unsafe.Pointer(bs)))

	_, _, err = pSubscribePrivateTopic.Call(api, 0)
	_, _, err = pSubscribePublicTopic.Call(api, 0)
	_, _, err = pInit.Call(api)
	fmt.Println(err)

	fmt.Scanln()
	//select {}
}
