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

void SetOnFrontConnected(void*, void*);
int tFrontConnected();
void SetOnFrontDisconnected(void*, void*);
int tFrontDisConnected();
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
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"unsafe"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

// Trade 交易接口
type Trade struct {
	api               unsafe.Pointer
	InvestorID        string              // 帐号
	BrokerID          string              // 经纪商
	TradingDay        string              // 交易日
	Instruments       sync.Map            // 合约列表
	InstrumentStatuss sync.Map            // 合约状态
	Positions         sync.Map            // 持仓列表
	Orders            sync.Map            // 委托
	Trades            sync.Map            // 成交
	sysID4Order       sync.Map            // orderSysID 对应的 Order
	Account           *goctp.AccountField // 帐户权益
	IsLogin           bool                // 登录成功

	passWord  string         // 密码
	sessionID int            // 判断是否自己的委托用
	qryTicker *time.Ticker   // 循环查询
	waitGroup sync.WaitGroup // 登录信号
	reqID     int            // requestid
	cntOrder  int            // 计算order数量

	onFrontConnected      goctp.OnFrontConnectedType // 事件
	onFrontDisConnected   goctp.OnFrontDisConnectedType
	onRspUserLogin        goctp.OnRspUserLoginType
	onRtnOrder            goctp.OnRtnOrderType
	onRtnCancel           goctp.OnRtnOrderType
	onErrRtnOrder         goctp.OnRtnErrOrderType
	onErrAction           goctp.OnRtnErrActionType
	onRtnTrade            goctp.OnRtnTradeType
	onRtnInstrumentStatus goctp.OnRtnInstrumentStatusType
	onRtnBankToFuture     goctp.OnRtnFromBankToFutureByFuture
	onRtnFutureToBank     goctp.OnRtnFromFutureToBankByFuture
}

var t *Trade

func (t *Trade) getReqID() C.int {
	t.reqID++
	return C.int(t.reqID)
}

// NewTrade 实例化
func NewTrade() *Trade {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
	t = new(Trade)
	// 初始化变量
	t.waitGroup = sync.WaitGroup{}
	t.IsLogin = false
	t.Account = new(goctp.AccountField)

	t.api = C.CreateApi()
	spi := C.CreateSpi()
	C.RegisterSpi(t.api, spi)

	C.SetOnFrontConnected(spi, C.tFrontConnected)
	C.SetOnFrontDisconnected(spi, C.tFrontDisConnected)
	C.SetOnRspUserLogin(spi, C.tRspUserLogin)
	C.SetOnRspAuthenticate(spi, C.tRspAuthenticate)
	C.SetOnRspSettlementInfoConfirm(spi, C.tRspSettlementInfoConfirm)
	C.SetOnRspQryInstrument(spi, C.tRspQryInstrument)
	C.SetOnRspQryClassifiedInstrument(spi, C.tRspQryClassifiedInstrument)
	C.SetOnRspQryTradingAccount(spi, C.tRspQryTradingAccount)
	C.SetOnRspQryInvestorPosition(spi, C.tRspQryInvestorPosition)
	C.SetOnErrRtnOrderInsert(spi, C.tErrRtnOrderInsert)
	C.SetOnErrRtnOrderAction(spi, C.tErrRtnOrderAction)
	C.SetOnRtnOrder(spi, C.tRtnOrder)
	C.SetOnRtnTrade(spi, C.tRtnTrade)
	C.SetOnRtnInstrumentStatus(spi, C.tRtnInstrumentStatus)
	C.SetOnRtnFromBankToFutureByFuture(spi, C.tRtnFromBankToFutureByFuture)
	C.SetOnRtnFromFutureToBankByFuture(spi, C.tRtnFromFutureToBankByFuture)

	return t
}

// ********************** 主调函数 ************************

// Release 接口销毁处理
func (t *Trade) Release() {
	t.qryTicker.Stop()
	t.IsLogin = false
	C.Release(t.api)
	t.tFrontDisConnected(C.int(0))
}

// ReqConnect 连接;Join阻塞，用goroutine
func (t *Trade) ReqConnect(addr string) {
	front := C.CString(addr)
	C.RegisterFront(t.api, front)
	defer C.free(unsafe.Pointer(front))
	// C.SubscribePrivateTopic(t.api, C.int(ctp.THOST_TERT_RESTART))
	// C.SubscribePublicTopic(t.api, C.int(ctp.THOST_TERT_RESTART))
	C.Init(t.api)
	// C.Join(t.api)
}

// ReqLogin 登录
func (t *Trade) ReqLogin(investor, pwd, broker, appID, authCode string) {
	t.InvestorID = investor
	t.passWord = pwd
	t.BrokerID = broker
	f := ctp.CThostFtdcReqAuthenticateField{}
	copy(f.BrokerID[:], broker)
	copy(f.UserID[:], investor)
	copy(f.AppID[:], appID)
	copy(f.AuthCode[:], authCode)
	C.ReqAuthenticate(t.api, (*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(&f)), t.getReqID())
}

// ReqOrderInsert 限价委托
func (t *Trade) ReqOrderInsert(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_GFD
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	C.ReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(&f)), id)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertMarket 市价委托
func (t *Trade) ReqOrderInsertMarket(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_AnyPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_IOC
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(0)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	C.ReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(&f)), id)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertFOK FOK委托[部成撤单]
func (t *Trade) ReqOrderInsertFOK(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_GFD
	f.VolumeCondition = ctp.THOST_FTDC_VC_CV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	C.ReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(&f)), id)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderInsertFAK FAK委托[全成or撤单]
func (t *Trade) ReqOrderInsertFAK(instrument string, buySell goctp.DirectionType, openClose goctp.OffsetFlagType, price float64, volume int) string {
	f := ctp.CThostFtdcInputOrderField{}
	copy(f.BrokerID[:], t.BrokerID)
	if info, ok := t.Instruments.Load(instrument); ok {
		copy(f.ExchangeID[:], info.(*goctp.InstrumentField).ExchangeID)
	}
	copy(f.UserID[:], t.InvestorID)
	copy(f.InvestorID[:], t.InvestorID)
	copy(f.AccountID[:], t.InvestorID)
	f.IsAutoSuspend = ctp.TThostFtdcBoolType(0)
	f.IsSwapOrder = ctp.TThostFtdcBoolType(0)
	f.ForceCloseReason = ctp.THOST_FTDC_FCC_NotForceClose
	// 参数赋值
	id := t.getReqID()
	copy(f.OrderRef[:], fmt.Sprintf("%012d", id))
	copy(f.InstrumentID[:], instrument)
	f.Direction = ctp.TThostFtdcDirectionType(buySell)
	f.CombOffsetFlag[0] = byte(openClose)
	f.CombHedgeFlag[0] = byte(goctp.HedgeFlagSpeculation)
	// 不同类型的Order
	f.OrderPriceType = ctp.THOST_FTDC_OPT_LimitPrice
	f.TimeCondition = ctp.THOST_FTDC_TC_IOC
	f.VolumeCondition = ctp.THOST_FTDC_VC_AV
	f.ContingentCondition = ctp.THOST_FTDC_CC_Immediately
	f.LimitPrice = ctp.TThostFtdcPriceType(price)
	f.VolumeTotalOriginal = ctp.TThostFtdcVolumeType(volume)
	C.ReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(&f)), id)
	return fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(f.OrderRef[:]))
}

// ReqOrderAction 撤单
func (t *Trade) ReqOrderAction(orderID string) int {
	if o, ok := t.Orders.Load(orderID); ok {
		var order = o.(*goctp.OrderField)
		f := ctp.CThostFtdcInputOrderActionField{}
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.UserID[:], t.InvestorID)
		copy(f.InstrumentID[:], order.InstrumentID)
		copy(f.ExchangeID[:], order.ExchangeID)
		copy(f.OrderRef[:], order.OrderRef)
		f.ActionFlag = ctp.THOST_FTDC_AF_Delete
		f.FrontID = ctp.TThostFtdcFrontIDType(order.FrontID)
		f.SessionID = ctp.TThostFtdcSessionIDType(order.SessionID)
		C.ReqOrderAction(t.api, (*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(&f)), t.getReqID())
		return 0
	}
	return -1
}

// ReqBankToFuture 银行转期货
func (t *Trade) ReqBankToFuture(bankID, bankAccount, bankPwd string, amount float64) {
	f := ctp.CThostFtdcReqTransferField{}
	copy(f.TradeCode[:], "202001")
	copy(f.BankBranchID[:], "0000")
	copy(f.BrokerID[:], []byte(t.BrokerID))
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.Password[:], []byte(t.passWord))
	copy(f.CurrencyID[:], "CNY")
	f.LastFragment = ctp.THOST_FTDC_LF_Yes
	f.IdCardType = ctp.THOST_FTDC_ICT_IDCard
	f.CustType = ctp.THOST_FTDC_CUSTT_Person
	f.InstallID = 1
	f.FutureSerial = 0
	f.VerifyCertNoFlag = ctp.THOST_FTDC_YNI_No
	f.FutureFetchAmount = 0
	f.CustFee = 0
	f.BrokerFee = 0
	f.SecuPwdFlag = ctp.THOST_FTDC_BPWDF_BlankCheck
	f.RequestID = ctp.TThostFtdcRequestIDType(t.getReqID())
	f.TID = 0
	copy(f.BankID[:], []byte(bankID))
	copy(f.BankAccount[:], []byte(bankAccount))
	copy(f.BankPassWord[:], []byte(bankPwd))
	f.TradeAmount = ctp.TThostFtdcTradeAmountType(amount)

	C.ReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(&f)), t.getReqID())
}

// ReqFutureToBank 期货转银行
func (t *Trade) ReqFutureToBank(bankID, bankAccount string, amount float64) {
	f := ctp.CThostFtdcReqTransferField{}
	copy(f.TradeCode[:], "202002")
	copy(f.BankBranchID[:], "0000")
	copy(f.BrokerID[:], []byte(t.BrokerID))
	copy(f.AccountID[:], []byte(t.InvestorID))
	copy(f.Password[:], []byte(t.passWord))
	copy(f.CurrencyID[:], "CNY")
	f.LastFragment = ctp.THOST_FTDC_LF_Yes
	f.IdCardType = ctp.THOST_FTDC_ICT_IDCard
	f.CustType = ctp.THOST_FTDC_CUSTT_Person
	f.InstallID = 1
	f.FutureSerial = 0
	f.VerifyCertNoFlag = ctp.THOST_FTDC_YNI_No
	f.FutureFetchAmount = 0
	f.CustFee = 0
	f.BrokerFee = 0
	f.SecuPwdFlag = ctp.THOST_FTDC_BPWDF_BlankCheck
	f.RequestID = ctp.TThostFtdcRequestIDType(t.getReqID())
	f.TID = 0
	copy(f.BankID[:], []byte(bankID))
	copy(f.BankAccount[:], []byte(bankAccount))
	// copy(f.BankPassWord[:], []byte(bankPwd))
	f.TradeAmount = ctp.TThostFtdcTradeAmountType(amount)

	C.ReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(&f)), t.getReqID())
}

// ********************** 注册客户响应 ************************

// RegOnFrontConnected 注册连接响应
func (t *Trade) RegOnFrontConnected(on goctp.OnFrontConnectedType) {
	t.onFrontConnected = on
}

// RegOnFrontDisConnected 注册连接响应
func (t *Trade) RegOnFrontDisConnected(on goctp.OnFrontDisConnectedType) {
	t.onFrontDisConnected = on
}

// RegOnRspUserLogin 注册登陆响应
func (t *Trade) RegOnRspUserLogin(on goctp.OnRspUserLoginType) {
	t.onRspUserLogin = on
}

// RegOnRtnOrder 注册委托响应
func (t *Trade) RegOnRtnOrder(on goctp.OnRtnOrderType) {
	t.onRtnOrder = on
}

// RegOnErrRtnOrder 注册委托响应
func (t *Trade) RegOnErrRtnOrder(on goctp.OnRtnErrOrderType) {
	t.onErrRtnOrder = on
}

// RegOnErrAction 注册撤单响应
func (t *Trade) RegOnErrAction(on goctp.OnRtnErrActionType) {
	t.onErrAction = on
}

// RegOnRtnCancel 注册撤单响应
func (t *Trade) RegOnRtnCancel(on goctp.OnRtnOrderType) {
	t.onRtnCancel = on
}

// RegOnRtnTrade 注册成交响应
func (t *Trade) RegOnRtnTrade(on goctp.OnRtnTradeType) {
	t.onRtnTrade = on
}

// RegOnRtnInstrumentStatus 注册合约状态变化
func (t *Trade) RegOnRtnInstrumentStatus(on goctp.OnRtnInstrumentStatusType) {
	t.onRtnInstrumentStatus = on
}

func (t *Trade) RegOnRtnFromBankToFuture(on goctp.OnRtnFromBankToFutureByFuture) {
	t.onRtnBankToFuture = on
}

func (t *Trade) RegOnRtnFromFutureToBank(on goctp.OnRtnFromFutureToBankByFuture) {
	t.onRtnFutureToBank = on
}

// ********************** 底层接口响应处理 **********************************

//export tRtnFromBankToFutureByFuture
func tRtnFromBankToFutureByFuture(field *C.struct_CThostFtdcRspTransferField) C.int {
	transField := (*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(field))
	if t.onRtnBankToFuture != nil {
		f := goctp.TransferField{
			Amout:      float64(transField.TradeAmount),
			CurrencyID: goctp.Bytes2String(transField.CurrencyID[:]),
			ErrorID:    int(transField.ErrorID),
			ErrorMsg:   goctp.Bytes2String(transField.ErrorMsg[:]),
		}
		t.onRtnBankToFuture(&f)
	}
	return 0
}

//export tRtnFromFutureToBankByFuture
func tRtnFromFutureToBankByFuture(field *C.struct_CThostFtdcRspTransferField) C.int {
	transField := (*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(field))
	if t.onRtnFutureToBank != nil {
		f := goctp.TransferField{
			Amout:      float64(transField.TradeAmount),
			CurrencyID: goctp.Bytes2String(transField.CurrencyID[:]),
			ErrorID:    int(transField.ErrorID),
			ErrorMsg:   goctp.Bytes2String(transField.ErrorMsg[:]),
		}
		t.onRtnFutureToBank(&f)
	}
	return 0
}

//export tRtnInstrumentStatus
func tRtnInstrumentStatus(field *C.struct_CThostFtdcInstrumentStatusField) C.int {
	statusField := (*ctp.CThostFtdcInstrumentStatusField)(unsafe.Pointer(field))
	status, loaded := t.InstrumentStatuss.LoadOrStore(goctp.Bytes2String(statusField.InstrumentID[:]), &goctp.InstrumentStatus{
		ExchangeID:       goctp.Bytes2String(statusField.ExchangeID[:]),
		InstrumentID:     goctp.Bytes2String(statusField.InstrumentID[:]),
		InstrumentStatus: goctp.InstrumentStatusType(statusField.InstrumentStatus),
		EnterTime:        goctp.Bytes2String(statusField.EnterTime[:]),
	})
	if loaded {
		status.(*goctp.InstrumentStatus).InstrumentStatus = goctp.InstrumentStatusType(statusField.InstrumentStatus)
		status.(*goctp.InstrumentStatus).EnterTime = goctp.Bytes2String(statusField.EnterTime[:])
	}

	if t.onRtnInstrumentStatus != nil {
		t.onRtnInstrumentStatus(status.(*goctp.InstrumentStatus))
	}
	return 0
}

//export tRtnTrade
func tRtnTrade(field *C.struct_CThostFtdcTradeField) C.int {
	tradeField := (*ctp.CThostFtdcTradeField)(unsafe.Pointer(field))
	key := fmt.Sprintf("%s_%c", tradeField.TradeID, tradeField.Direction)
	tf, _ := t.Trades.LoadOrStore(key, &goctp.TradeField{
		Direction:    goctp.DirectionType(tradeField.Direction),
		HedgeFlag:    goctp.HedgeFlagType(tradeField.HedgeFlag),
		InstrumentID: goctp.Bytes2String(tradeField.InstrumentID[:]),
		ExchangeID:   goctp.Bytes2String(tradeField.ExchangeID[:]),
		TradingDay:   goctp.Bytes2String(tradeField.TradingDay[:]),
		Volume:       int(tradeField.Volume),
		OffsetFlag:   goctp.OffsetFlagType(tradeField.OffsetFlag),
		OrderSysID:   goctp.Bytes2String(tradeField.OrderSysID[:]),
		Price:        float64(tradeField.Price),
		TradeDate:    goctp.Bytes2String(tradeField.TradeDate[:]),
		TradeTime:    goctp.Bytes2String(tradeField.TradeTime[:]),
		TradeID:      key,
	})
	var f = tf.(*goctp.TradeField)
	// 更新持仓
	if f.OffsetFlag == goctp.OffsetFlagOpen {
		var key string
		if f.Direction == goctp.DirectionBuy {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionLong, f.HedgeFlag)
		} else {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionShort, f.HedgeFlag)
		}
		pf, _ := t.Positions.LoadOrStore(key, &goctp.PositionField{
			InstrumentID:      f.InstrumentID,
			PositionDirection: goctp.PosiDirectionLong,
			HedgeFlag:         f.HedgeFlag,
			ExchangeID:        f.ExchangeID,
		})
		var p = pf.(*goctp.PositionField)
		p.OpenVolume += f.Volume
		p.OpenAmount += f.Price * float64(f.Volume)
		if info, ok := t.Instruments.Load(f.InstrumentID); ok {
			p.OpenCost += f.Price * float64(f.Volume) * float64(info.(*goctp.InstrumentField).VolumeMultiple)
		}
		p.Position += f.Volume
		p.TodayPosition += f.Volume
	} else {
		var key string
		if f.Direction == goctp.DirectionBuy {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionShort, f.HedgeFlag)
		} else {
			key = fmt.Sprintf("%s_%c_%c", f.InstrumentID, goctp.PosiDirectionLong, f.HedgeFlag)
		}
		if posi, ok := t.Positions.Load(key); ok {
			var p = posi.(*goctp.PositionField)
			p.OpenVolume -= f.Volume
			p.OpenAmount -= f.Price * float64(f.Volume)
			if info, ok := t.Instruments.Load(f.InstrumentID); ok {
				p.OpenCost -= f.Price * float64(f.Volume) * float64(info.(*goctp.InstrumentField).VolumeMultiple)
			}
			p.Position -= f.Volume
			if f.OffsetFlag == goctp.OffsetFlagCloseToday {
				p.TodayPosition -= f.Volume
			} else {
				p.YdPosition -= f.Volume
			}
		}
	}
	// 处理对应的Order
	if ord, ok := t.sysID4Order.Load(f.OrderSysID); ok {
		var o = ord.(*goctp.OrderField)
		o.LastTradeTime = f.TradeTime
		o.VolumeTraded = f.Volume
		o.VolumeLeft -= f.Volume
		if o.VolumeLeft == 0 {
			o.OrderStatus = goctp.OrderStatusAllTraded
			o.StatusMsg = "全部成交"
		} else {
			o.OrderStatus = goctp.OrderStatusPartTradedQueueing
			o.StatusMsg = "部分成交"
		}
		if t.onRtnOrder != nil {
			t.onRtnOrder(o)
		}
	}
	// 客户端响应
	if t.onRtnTrade != nil {
		t.onRtnTrade(f)
	}
	return 0
}

//export tRtnOrder
func tRtnOrder(field *C.struct_CThostFtdcOrderField) C.int {
	t.cntOrder++
	orderField := (*ctp.CThostFtdcOrderField)(unsafe.Pointer(field))
	key := fmt.Sprintf("%d_%s", orderField.SessionID, goctp.Bytes2String(orderField.OrderRef[:]))
	if of, ok := t.Orders.LoadOrStore(key, &goctp.OrderField{
		InstrumentID:        goctp.Bytes2String(orderField.InstrumentID[:]),
		SessionID:           int(orderField.SessionID),
		FrontID:             int(orderField.FrontID),
		OrderRef:            goctp.Bytes2String(orderField.OrderRef[:]),
		Direction:           goctp.DirectionType(orderField.Direction),
		OffsetFlag:          goctp.OffsetFlagType(orderField.CombOffsetFlag[0]),
		HedgeFlag:           goctp.HedgeFlagType(orderField.CombHedgeFlag[0]),
		LimitPrice:          float64(orderField.LimitPrice),
		VolumeTotalOriginal: int(orderField.VolumeTotalOriginal),
		VolumeLeft:          int(orderField.VolumeTotalOriginal),
		ExchangeID:          goctp.Bytes2String(orderField.ExchangeID[:]),
		InsertDate:          goctp.Bytes2String(orderField.InsertDate[:]),
		InsertTime:          goctp.Bytes2String(orderField.InsertTime[:]),
		OrderStatus:         goctp.OrderStatusNoTradeQueueing, // OrderStatusType(orderField.OrderStatus)
		StatusMsg:           "委托已提交",                          // bytes2GBKbytes2GBKString(orderField.StatusMsg[:])
		IsLocal:             int(orderField.SessionID) == t.sessionID,
	}); !ok { // 新添加
		if t.onRtnOrder != nil {
			t.onRtnOrder(of.(*goctp.OrderField))
		}
	} else {
		var o = of.(*goctp.OrderField)
		if o.OrderStatus == goctp.OrderStatusCanceled {
			o.CancelTime = goctp.Bytes2String(orderField.CancelTime[:])
			// 错单
			if strings.Contains(o.StatusMsg, "被拒绝") {
				if t.onErrRtnOrder != nil {
					t.onErrRtnOrder(o, &goctp.RspInfoField{
						ErrorID:  -1,
						ErrorMsg: o.StatusMsg,
					})
				}
			} else if t.onRtnCancel != nil {
				t.onRtnCancel(o)
			}
		} else {
			o.OrderSysID = goctp.Bytes2String(orderField.OrderSysID[:])
			if len(o.OrderSysID) > 0 {
				t.sysID4Order.Store(o.OrderSysID, o)
			}
		}
	}
	return 0
}

//export tErrRtnOrderAction
func tErrRtnOrderAction(field *C.struct_CThostFtdcOrderActionField, info *C.struct_CThostFtdcRspInfoField) C.int {
	actionField := (*ctp.CThostFtdcOrderActionField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if t.onErrAction != nil {
		t.onErrAction(fmt.Sprintf("%d_%s", actionField.SessionID, actionField.OrderRef), &goctp.RspInfoField{
			ErrorID:  int(infoField.ErrorID),
			ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:]),
		})
	}
	return 0
}

//export tErrRtnOrderInsert
func tErrRtnOrderInsert(field *C.struct_CThostFtdcInputOrderField, info *C.struct_CThostFtdcRspInfoField) C.int {
	orderField := (*ctp.CThostFtdcInputOrderField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	key := fmt.Sprintf("%d_%s", t.sessionID, goctp.Bytes2String(orderField.OrderRef[:]))
	of, _ := t.Orders.LoadOrStore(key, &goctp.OrderField{
		InstrumentID:        goctp.Bytes2String(orderField.InstrumentID[:]),
		SessionID:           t.sessionID,
		FrontID:             0,
		OrderRef:            goctp.Bytes2String(orderField.OrderRef[:]),
		Direction:           goctp.DirectionType(orderField.Direction),
		OffsetFlag:          goctp.OffsetFlagType(orderField.CombOffsetFlag[0]),
		HedgeFlag:           goctp.HedgeFlagType(orderField.CombHedgeFlag[0]),
		LimitPrice:          float64(orderField.LimitPrice),
		VolumeTotalOriginal: int(orderField.VolumeTotalOriginal),
		VolumeLeft:          int(orderField.VolumeTotalOriginal),
		ExchangeID:          goctp.Bytes2String(orderField.ExchangeID[:]),
		IsLocal:             true,
	})
	var o = of.(*goctp.OrderField)
	o.OrderStatus = goctp.OrderStatusCanceled
	if t.onErrRtnOrder != nil {
		t.onErrRtnOrder(o, &goctp.RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

//export tRspQryInvestorPosition
func tRspQryInvestorPosition(field *C.struct_CThostFtdcInvestorPositionField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	positionField := (*ctp.CThostFtdcInvestorPositionField)(unsafe.Pointer(field))
	//infoField := (* ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if len(goctp.Bytes2String(positionField.InstrumentID[:])) > 0 {
		key := fmt.Sprintf("%s_%c_%c", goctp.Bytes2String(positionField.InstrumentID[:]), goctp.PosiDirectionType(positionField.PosiDirection), goctp.HedgeFlagType(positionField.HedgeFlag))
		pf, _ := t.Positions.LoadOrStore(key, &goctp.PositionField{
			InstrumentID:      goctp.Bytes2String(positionField.InstrumentID[:]),
			PositionDirection: goctp.PosiDirectionType(positionField.PosiDirection),
			HedgeFlag:         goctp.HedgeFlagType(positionField.HedgeFlag),
			ExchangeID:        goctp.Bytes2String(positionField.ExchangeID[:]),
		})
		var p = pf.(*goctp.PositionField)
		p.YdPosition = int(positionField.YdPosition)
		p.Position = int(positionField.Position)
		p.LongFrozen = int(positionField.LongFrozen)
		p.ShortFrozen = int(positionField.ShortFrozen)
		p.LongFrozenAmount = float64(positionField.LongFrozenAmount)
		p.ShortFrozenAmount = float64(positionField.ShortFrozenAmount)
		p.OpenVolume = int(positionField.OpenVolume)
		p.CloseVolume = int(positionField.CloseVolume)
		p.OpenAmount = float64(positionField.OpenAmount)
		p.CloseAmount = float64(positionField.CloseAmount)
		p.PositionCost = float64(positionField.PositionCost)
		p.PreMargin = float64(positionField.PreMargin)
		p.UseMargin = float64(positionField.UseMargin)
		p.FrozenMargin = float64(positionField.FrozenMargin)
		p.FrozenCash = float64(positionField.FrozenCash)
		p.FrozenCommission = float64(positionField.FrozenCommission)
		p.CashIn = float64(positionField.CashIn)
		p.Commission = float64(positionField.Commission)
		p.CloseProfit = float64(positionField.CloseProfit)
		p.PositionProfit = float64(positionField.PositionProfit)
		p.PreSettlementPrice = float64(positionField.PreSettlementPrice)
		p.SettlementPrice = float64(positionField.SettlementPrice)
		p.OpenCost = float64(positionField.OpenCost)
		p.ExchangeMargin = float64(positionField.ExchangeMargin)
		p.CombPosition = int(positionField.CombPosition)
		p.CombLongFrozen = int(positionField.CombLongFrozen)
		p.CombShortFrozen = int(positionField.CombShortFrozen)
		p.CloseProfitByDate = float64(positionField.CloseProfitByDate)
		p.CloseProfitByTrade = float64(positionField.CloseProfitByTrade)
		p.TodayPosition = int(positionField.TodayPosition)
		p.StrikeFrozen = int(positionField.StrikeFrozen)
		p.StrikeFrozenAmount = float64(positionField.StrikeFrozenAmount)
		p.AbandonFrozen = int(positionField.AbandonFrozen)
		p.YdStrikeFrozen = int(positionField.YdStrikeFrozen)
		p.PositionCostOffset = float64(positionField.PositionCostOffset)
	}
	return 0
}

//export tRspQryTradingAccount
func tRspQryTradingAccount(field *C.struct_CThostFtdcTradingAccountField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	accountField := (*ctp.CThostFtdcTradingAccountField)(unsafe.Pointer(field))
	//infoField := (* ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	t.Account.PreMortgage = float64(accountField.PreMortgage)
	t.Account.PreDeposit = float64(accountField.PreDeposit)
	t.Account.PreBalance = float64(accountField.PreBalance)
	t.Account.PreMargin = float64(accountField.PreMargin)
	t.Account.InterestBase = float64(accountField.InterestBase)
	t.Account.Interest = float64(accountField.Interest)
	t.Account.Deposit = float64(accountField.Deposit)
	t.Account.Withdraw = float64(accountField.Withdraw)
	t.Account.FrozenMargin = float64(accountField.FrozenMargin)
	t.Account.FrozenCash = float64(accountField.FrozenCash)
	t.Account.FrozenCommission = float64(accountField.FrozenCommission)
	t.Account.CurrMargin = float64(accountField.CurrMargin)
	t.Account.CashIn = float64(accountField.CashIn)
	t.Account.Commission = float64(accountField.Commission)
	t.Account.CloseProfit = float64(accountField.CloseProfit)
	t.Account.PositionProfit = float64(accountField.PositionProfit)
	t.Account.Balance = float64(accountField.Balance)
	t.Account.Available = float64(accountField.Available)
	t.Account.WithdrawQuota = float64(accountField.WithdrawQuota)
	t.Account.Reserve = float64(accountField.Reserve)
	t.Account.Credit = float64(accountField.Credit)
	t.Account.Mortgage = float64(accountField.Mortgage)
	t.Account.ExchangeMargin = float64(accountField.ExchangeMargin)
	t.Account.DeliveryMargin = float64(accountField.DeliveryMargin)
	t.Account.ExchangeDeliveryMargin = float64(accountField.ExchangeDeliveryMargin)
	t.Account.ReserveBalance = float64(accountField.ReserveBalance)
	t.Account.CurrencyID = goctp.Bytes2String(accountField.CurrencyID[:])
	t.Account.PreFundMortgageIn = float64(accountField.PreFundMortgageIn)
	t.Account.PreFundMortgageOut = float64(accountField.PreFundMortgageOut)
	t.Account.FundMortgageIn = float64(accountField.FundMortgageIn)
	t.Account.FundMortgageOut = float64(accountField.FundMortgageOut)
	t.Account.FundMortgageAvailable = float64(accountField.FundMortgageAvailable)
	t.Account.MortgageableFund = float64(accountField.MortgageableFund)
	return 0
}

//export tRspQryClassifiedInstrument
func tRspQryClassifiedInstrument(field *C.struct_CThostFtdcInstrumentField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	return tRspQryInstrument(field, info, i, b)
}

//export tRspQryInstrument
func tRspQryInstrument(field *C.struct_CThostFtdcInstrumentField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	instrumentField := (*ctp.CThostFtdcInstrumentField)(unsafe.Pointer(field))
	//infoField := (* ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if instrumentField != nil {
		t.Instruments.Store(goctp.Bytes2String(instrumentField.InstrumentID[:]), &goctp.InstrumentField{
			InstrumentID:              goctp.Bytes2String(instrumentField.InstrumentID[:]),
			ExchangeID:                goctp.Bytes2String(instrumentField.ExchangeID[:]),
			ProductID:                 goctp.Bytes2String(instrumentField.ProductID[:]),
			ProductClass:              goctp.ProductClassType(instrumentField.ProductClass),
			MaxMarketOrderVolume:      int(instrumentField.MaxMarketOrderVolume),
			MinMarketOrderVolume:      int(instrumentField.MinMarketOrderVolume),
			MaxLimitOrderVolume:       int(instrumentField.MaxLimitOrderVolume),
			MinLimitOrderVolume:       int(instrumentField.MinLimitOrderVolume),
			VolumeMultiple:            int(instrumentField.VolumeMultiple),
			PriceTick:                 float64(instrumentField.PriceTick),
			PositionType:              goctp.PositionTypeType(instrumentField.PositionType),
			UseMaxMarginSideAlgorithm: instrumentField.MaxMarginSideAlgorithm == '1',
			UnderlyingInstrID:         goctp.Bytes2String(instrumentField.UnderlyingInstrID[:]),
			StrikePrice:               float64(instrumentField.StrikePrice),
			OptionsType:               goctp.OptionsTypeType(instrumentField.OptionsType),
			UnderlyingMultiple:        float64(instrumentField.UnderlyingMultiple),
			CombinationType:           goctp.CombinationTypeType(instrumentField.CombinationType),
		})
	}
	if b {
		// fmt.Printf("收到合约: %d\n", len(t.Instruments))
		// 循环查询持仓与权益
		go t.qry()
	}
	return 0
}

// 循环查询持仓&资金
func (t *Trade) qry() {
	t.qryTicker = time.NewTicker(1100 * time.Millisecond)
	// 等待之前的Order响应完再发送登录通知
	ordCnt := t.cntOrder
	for range t.qryTicker.C {
		if ordCnt == t.cntOrder {
			break
		}
		ordCnt = t.cntOrder
	}

	qryAccount := ctp.CThostFtdcQryTradingAccountField{}
	copy(qryAccount.InvestorID[:], t.InvestorID)
	copy(qryAccount.BrokerID[:], t.BrokerID)
	qryPosition := ctp.CThostFtdcQryInvestorPositionField{}
	copy(qryPosition.InvestorID[:], t.InvestorID)
	copy(qryPosition.BrokerID[:], t.BrokerID)
	// 启动查询
	bQryAccount := false
	for range t.qryTicker.C {
		if bQryAccount {
			C.ReqQryTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(&qryAccount)), t.getReqID())
		} else {
			C.ReqQryInvestorPosition(t.api, (*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(&qryPosition)), t.getReqID())
		}
		bQryAccount = !bQryAccount
		if !t.IsLogin && !bQryAccount { // account position 都查一遍后再通知
			// 登录成功响应
			t.IsLogin = true
			t.waitGroup.Done() // 通知:登录响应可以发了
		}
	}
}

//export tRspSettlementInfoConfirm
func tRspSettlementInfoConfirm(field *C.struct_CThostFtdcSettlementInfoConfirmField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	if t.BrokerID == "9999" {
		C.ReqQryInstrument(t.api, (*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(&ctp.CThostFtdcQryInstrumentField{})), t.getReqID())
	} else {
		// TradingType = THOST_FTDC_TD_TRADE,
		req := ctp.CThostFtdcQryClassifiedInstrumentField{
			TradingType: ctp.THOST_FTDC_TD_TRADE,
			ClassType:   ctp.THOST_FTDC_INS_ALL,
		}
		C.ReqQryClassifiedInstrument(t.api, (*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(&req)), t.getReqID())
	}
	return 0
}

//export tRspUserLogin
func tRspUserLogin(field *C.struct_CThostFtdcRspUserLoginField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	loginField := (*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
	infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if infoField.ErrorID == 0 {
		t.sessionID = int(loginField.SessionID)
		t.TradingDay = goctp.Bytes2String(loginField.TradingDay[:])
		f := ctp.CThostFtdcSettlementInfoConfirmField{}
		copy(f.InvestorID[:], t.InvestorID)
		copy(f.AccountID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		C.ReqSettlementInfoConfirm(t.api, (*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(&f)), t.getReqID())

		// 用waitgroup控制登录消息发送信号
		if t.onRspUserLogin != nil {
			t.waitGroup.Add(1)
			go func(field *goctp.RspUserLoginField) {
				t.waitGroup.Wait()
				t.onRspUserLogin(field, &goctp.RspInfoField{ErrorID: 0, ErrorMsg: "成功"})
			}(&goctp.RspUserLoginField{
				TradingDay:  t.TradingDay,
				LoginTime:   goctp.Bytes2String(loginField.LoginTime[:]),
				BrokerID:    t.BrokerID,
				UserID:      t.InvestorID,
				FrontID:     int(loginField.FrontID),
				SessionID:   t.sessionID,
				MaxOrderRef: string(loginField.MaxOrderRef[:]),
			})
		}
	} else {
		t.onRspUserLogin(&goctp.RspUserLoginField{}, &goctp.RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

//export tRspAuthenticate
func tRspAuthenticate(field *C.struct_CThostFtdcRspAuthenticateField, info *C.struct_CThostFtdcRspInfoField, i C.int, b C._Bool) C.int {
	//authField := (* ctp.CThostFtdcRspAuthenticateField)(unsafe.Pointer(field))
	//infoField := (* ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
	if info.ErrorID == 0 {
		f := ctp.CThostFtdcReqUserLoginField{}
		copy(f.UserID[:], t.InvestorID)
		copy(f.BrokerID[:], t.BrokerID)
		copy(f.Password[:], t.passWord)
		copy(f.UserProductInfo[:], "@HF")
		C.ReqUserLogin(t.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(&f)), t.getReqID())
	} else if t.onRspUserLogin != nil {
		infoField := (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(info))
		t.onRspUserLogin(&goctp.RspUserLoginField{}, &goctp.RspInfoField{ErrorID: int(infoField.ErrorID), ErrorMsg: goctp.Bytes2String(infoField.ErrorMsg[:])})
	}
	return 0
}

//export tFrontDisConnected
func tFrontDisConnected(reason C.int) C.int {
	if t.onFrontDisConnected != nil {
		t.onFrontDisConnected(int(reason))
	}
	return 0
}

//export tFrontConnected
func tFrontConnected() C.int {
	if t.onFrontConnected != nil {
		t.onFrontConnected()
	}
	return 0
}
