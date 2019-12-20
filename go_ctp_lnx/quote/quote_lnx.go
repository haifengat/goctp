package quote

/*
#cgo CPPFLAGS: -fPIC -I${SRCDIR}
#cgo LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath-link,${SRCDIR}  -lctp_quote -lthostmduserapi_se -lstdc++

#include "../../go_ctp/ctp_20190220_se_x64/ThostFtdcUserApiDataType.h"
#include "../../go_ctp/ctp_20190220_se_x64/ThostFtdcUserApiStruct.h"
void* CreateApi();
void* CreateSpi();
void* RegisterSpi(void*, void*);
void* RegisterFront(void*, char*);
void* Init(void*);
void* Release(void*);
void* ReqUserLogin(void*, struct CThostFtdcReqUserLoginField*, int);
void* SubscribeMarketData(void*, char *ppInstrumentID[], int nCount);

void SetOnFrontConnected(void*, void*);
int OnFrontConnected();
void SetOnRspUserLogin(void*, void*);
int OnRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
void SetOnRtnDepthMarketData(void*, void*);
int OnRtnDepthMarketData(struct CThostFtdcDepthMarketDataField *pDepthMarketData);
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"hf_go_ctp"
	"hf_go_ctp/go_ctp"
	"unsafe"
)

type Quote struct {
	api unsafe.Pointer
	// 帐号
	InvestorID string
	// 经纪商
	BrokerID string
	// 交易日
	TradingDay       string
	onFrontConnected hf_go_ctp.OnFrontConnectedType
	onRspUserLogin   hf_go_ctp.OnRspUserLoginType
	onTick           hf_go_ctp.OnTickType
	reqID            int
}

func (q *Quote) getReqID() C.int {
	q.reqID++
	return C.int(q.reqID)
}

var q *Quote

func NewQuote() *Quote {
	q = new(Quote)
	q.api = C.CreateApi()
	spi := C.CreateSpi()
	C.RegisterSpi(q.api, spi)

	C.SetOnFrontConnected(spi, C.OnFrontConnected)
	C.SetOnRspUserLogin(spi, C.OnRspUserLogin)
	C.SetOnRtnDepthMarketData(spi, C.OnRtnDepthMarketData)
	return q
}

func (q *Quote) Release() {
	C.Release(q.api)
}

func (q *Quote) ReqConnect(addr string) {
	front := C.CString(addr)
	C.RegisterFront(q.api, front)
	defer C.free(unsafe.Pointer(front))
	C.Init(q.api)
}

func (q *Quote) ReqLogin(investor, pwd, broker string) {
	q.InvestorID = investor
	q.BrokerID = broker
	f := go_ctp.CThostFtdcReqUserLoginField{}
	copy(f.UserID[:], q.InvestorID)
	copy(f.BrokerID[:], q.BrokerID)
	copy(f.Password[:], pwd)
	copy(f.UserProductInfo[:], "@HF")
	C.ReqUserLogin(q.api, (* C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(&f)), q.getReqID())
}

func (q *Quote) ReqSubscript(instrument string) {
	inst := make([]*C.char, 1)
	inst[0] = (*C.char)(unsafe.Pointer(C.CString(instrument)))
	C.SubscribeMarketData(q.api, (**C.char)(unsafe.Pointer(&inst[0])), 1)
}

func (q *Quote) RegOnFrontConnected(on hf_go_ctp.OnFrontConnectedType) {
	q.onFrontConnected = on
}
func (q *Quote) RegOnRspUserLogin(on hf_go_ctp.OnRspUserLoginType) {
	q.onRspUserLogin = on
}
func (q *Quote) RegOnTick(on hf_go_ctp.OnTickType) {
	q.onTick = on
}

//export OnRtnDepthMarketData
func OnRtnDepthMarketData(field *C.struct_CThostFtdcDepthMarketDataField) C.int {
	dataField := (*go_ctp.CThostFtdcDepthMarketDataField)(unsafe.Pointer(field))
	if q.onTick == nil {
		return 0
	}
	tick := hf_go_ctp.TickField{
		TradingDay:      hf_go_ctp.Bytes2String(dataField.TradingDay[:]),
		InstrumentID:    hf_go_ctp.Bytes2String(dataField.InstrumentID[:]),
		ExchangeID:      hf_go_ctp.Bytes2String(dataField.ExchangeID[:]),
		LastPrice:       float64(dataField.LastPrice),
		OpenPrice:       float64(dataField.OpenPrice),
		HighestPrice:    float64(dataField.HighestPrice),
		LowestPrice:     float64(dataField.LowestPrice),
		Volume:          int(dataField.Volume),
		Turnover:        float64(dataField.Turnover),
		OpenInterest:    float64(dataField.OpenInterest),
		SettlementPrice: float64(dataField.SettlementPrice),
		UpperLimitPrice: float64(dataField.UpperLimitPrice),
		LowerLimitPrice: float64(dataField.LowerLimitPrice),
		CurrDelta:       float64(dataField.CurrDelta),
		UpdateTime:      hf_go_ctp.Bytes2String(dataField.UpdateTime[:]),
		UpdateMillisec:  int(dataField.UpdateMillisec),
		BidPrice1:       float64(dataField.BidPrice1),
		BidVolume1:      int(dataField.BidVolume1),
		AskPrice1:       float64(dataField.AskPrice1),
		AskVolume1:      int(dataField.AskVolume1),
		BidPrice2:       float64(dataField.BidPrice2),
		BidVolume2:      int(dataField.BidVolume2),
		AskPrice2:       float64(dataField.AskPrice2),
		AskVolume2:      int(dataField.AskVolume2),
		BidPrice3:       float64(dataField.BidPrice3),
		BidVolume3:      int(dataField.BidVolume3),
		AskPrice3:       float64(dataField.AskPrice3),
		AskVolume3:      int(dataField.AskVolume3),
		BidPrice4:       float64(dataField.BidPrice4),
		BidVolume4:      int(dataField.BidVolume4),
		AskPrice4:       float64(dataField.AskPrice4),
		AskVolume4:      int(dataField.AskVolume4),
		BidPrice5:       float64(dataField.BidPrice5),
		BidVolume5:      int(dataField.BidVolume5),
		AskPrice5:       float64(dataField.AskPrice5),
		AskVolume5:      int(dataField.AskVolume5),
		AveragePrice:    float64(dataField.AskPrice5),
		ActionDay:       hf_go_ctp.Bytes2String(dataField.ActionDay[:]),
	}
	q.onTick(&tick)
	return 0
}

// 登陆响应
//export OnRspUserLogin
func OnRspUserLogin(field *C.struct_CThostFtdcRspUserLoginField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {

	loginField := (* go_ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
	infoField := (* go_ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if q.onRspUserLogin == nil {
		return 0
	}
	q.onRspUserLogin(&hf_go_ctp.RspUserLoginField{
		TradingDay:  string(loginField.TradingDay[:]),
		LoginTime:   string(loginField.LoginTime[:]),
		BrokerID:    string(loginField.BrokerID[:]),
		UserID:      string(loginField.UserID[:]),
		FrontID:     int(loginField.FrontID),
		SessionID:   int(loginField.SessionID),
		MaxOrderRef: string(loginField.MaxOrderRef[:]),
	}, &hf_go_ctp.RspInfoField{
		ErrorID:  int(infoField.ErrorID),
		ErrorMsg: hf_go_ctp.Bytes2String(infoField.ErrorMsg[:]),
	})
	return 0
}

// 连接前置响应
//export OnFrontConnected
func OnFrontConnected() C.int {
	if q.onFrontConnected != nil {
		q.onFrontConnected()
	}
	return 0
}
