package win

import (
	ctp "goctp/ctpdefine"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

type tOnFrontConnectedType func() uintptr
type tOnFrontDisconnectedType func(int) uintptr
type tOnHeartBeatWarningType func(int) uintptr
type tOnRspAuthenticateType func(*ctp.CThostFtdcRspAuthenticateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspUserLoginType func(*ctp.CThostFtdcRspUserLoginField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspUserLogoutType func(*ctp.CThostFtdcUserLogoutField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspUserPasswordUpdateType func(*ctp.CThostFtdcUserPasswordUpdateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspTradingAccountPasswordUpdateType func(*ctp.CThostFtdcTradingAccountPasswordUpdateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspUserAuthMethodType func(*ctp.CThostFtdcRspUserAuthMethodField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspGenUserCaptchaType func(*ctp.CThostFtdcRspGenUserCaptchaField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspGenUserTextType func(*ctp.CThostFtdcRspGenUserTextField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspOrderInsertType func(*ctp.CThostFtdcInputOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspParkedOrderInsertType func(*ctp.CThostFtdcParkedOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspParkedOrderActionType func(*ctp.CThostFtdcParkedOrderActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspOrderActionType func(*ctp.CThostFtdcInputOrderActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQueryMaxOrderVolumeType func(*ctp.CThostFtdcQueryMaxOrderVolumeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspSettlementInfoConfirmType func(*ctp.CThostFtdcSettlementInfoConfirmField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspRemoveParkedOrderType func(*ctp.CThostFtdcRemoveParkedOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspRemoveParkedOrderActionType func(*ctp.CThostFtdcRemoveParkedOrderActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspExecOrderInsertType func(*ctp.CThostFtdcInputExecOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspExecOrderActionType func(*ctp.CThostFtdcInputExecOrderActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspForQuoteInsertType func(*ctp.CThostFtdcInputForQuoteField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQuoteInsertType func(*ctp.CThostFtdcInputQuoteField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQuoteActionType func(*ctp.CThostFtdcInputQuoteActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspBatchOrderActionType func(*ctp.CThostFtdcInputBatchOrderActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspOptionSelfCloseInsertType func(*ctp.CThostFtdcInputOptionSelfCloseField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspOptionSelfCloseActionType func(*ctp.CThostFtdcInputOptionSelfCloseActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspCombActionInsertType func(*ctp.CThostFtdcInputCombActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryOrderType func(*ctp.CThostFtdcOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryTradeType func(*ctp.CThostFtdcTradeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInvestorPositionType func(*ctp.CThostFtdcInvestorPositionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryTradingAccountType func(*ctp.CThostFtdcTradingAccountField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInvestorType func(*ctp.CThostFtdcInvestorField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryTradingCodeType func(*ctp.CThostFtdcTradingCodeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInstrumentMarginRateType func(*ctp.CThostFtdcInstrumentMarginRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInstrumentCommissionRateType func(*ctp.CThostFtdcInstrumentCommissionRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryExchangeType func(*ctp.CThostFtdcExchangeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryProductType func(*ctp.CThostFtdcProductField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInstrumentType func(*ctp.CThostFtdcInstrumentField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryDepthMarketDataType func(*ctp.CThostFtdcDepthMarketDataField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQrySettlementInfoType func(*ctp.CThostFtdcSettlementInfoField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryTransferBankType func(*ctp.CThostFtdcTransferBankField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInvestorPositionDetailType func(*ctp.CThostFtdcInvestorPositionDetailField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryNoticeType func(*ctp.CThostFtdcNoticeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQrySettlementInfoConfirmType func(*ctp.CThostFtdcSettlementInfoConfirmField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInvestorPositionCombineDetailType func(*ctp.CThostFtdcInvestorPositionCombineDetailField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryCFMMCTradingAccountKeyType func(*ctp.CThostFtdcCFMMCTradingAccountKeyField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryEWarrantOffsetType func(*ctp.CThostFtdcEWarrantOffsetField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInvestorProductGroupMarginType func(*ctp.CThostFtdcInvestorProductGroupMarginField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryExchangeMarginRateType func(*ctp.CThostFtdcExchangeMarginRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryExchangeMarginRateAdjustType func(*ctp.CThostFtdcExchangeMarginRateAdjustField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryExchangeRateType func(*ctp.CThostFtdcExchangeRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQrySecAgentACIDMapType func(*ctp.CThostFtdcSecAgentACIDMapField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryProductExchRateType func(*ctp.CThostFtdcProductExchRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryProductGroupType func(*ctp.CThostFtdcProductGroupField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryMMInstrumentCommissionRateType func(*ctp.CThostFtdcMMInstrumentCommissionRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryMMOptionInstrCommRateType func(*ctp.CThostFtdcMMOptionInstrCommRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInstrumentOrderCommRateType func(*ctp.CThostFtdcInstrumentOrderCommRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQrySecAgentTradingAccountType func(*ctp.CThostFtdcTradingAccountField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQrySecAgentCheckModeType func(*ctp.CThostFtdcSecAgentCheckModeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQrySecAgentTradeInfoType func(*ctp.CThostFtdcSecAgentTradeInfoField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryOptionInstrTradeCostType func(*ctp.CThostFtdcOptionInstrTradeCostField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryOptionInstrCommRateType func(*ctp.CThostFtdcOptionInstrCommRateField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryExecOrderType func(*ctp.CThostFtdcExecOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryForQuoteType func(*ctp.CThostFtdcForQuoteField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryQuoteType func(*ctp.CThostFtdcQuoteField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryOptionSelfCloseType func(*ctp.CThostFtdcOptionSelfCloseField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryInvestUnitType func(*ctp.CThostFtdcInvestUnitField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryCombInstrumentGuardType func(*ctp.CThostFtdcCombInstrumentGuardField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryCombActionType func(*ctp.CThostFtdcCombActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryTransferSerialType func(*ctp.CThostFtdcTransferSerialField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryAccountregisterType func(*ctp.CThostFtdcAccountregisterField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspErrorType func(*ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRtnOrderType func(*ctp.CThostFtdcOrderField) uintptr
type tOnRtnTradeType func(*ctp.CThostFtdcTradeField) uintptr
type tOnErrRtnOrderInsertType func(*ctp.CThostFtdcInputOrderField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnOrderActionType func(*ctp.CThostFtdcOrderActionField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRtnInstrumentStatusType func(*ctp.CThostFtdcInstrumentStatusField) uintptr
type tOnRtnBulletinType func(*ctp.CThostFtdcBulletinField) uintptr
type tOnRtnTradingNoticeType func(*ctp.CThostFtdcTradingNoticeInfoField) uintptr
type tOnRtnErrorConditionalOrderType func(*ctp.CThostFtdcErrorConditionalOrderField) uintptr
type tOnRtnExecOrderType func(*ctp.CThostFtdcExecOrderField) uintptr
type tOnErrRtnExecOrderInsertType func(*ctp.CThostFtdcInputExecOrderField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnExecOrderActionType func(*ctp.CThostFtdcExecOrderActionField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnForQuoteInsertType func(*ctp.CThostFtdcInputForQuoteField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRtnQuoteType func(*ctp.CThostFtdcQuoteField) uintptr
type tOnErrRtnQuoteInsertType func(*ctp.CThostFtdcInputQuoteField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnQuoteActionType func(*ctp.CThostFtdcQuoteActionField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRtnForQuoteRspType func(*ctp.CThostFtdcForQuoteRspField) uintptr
type tOnRtnCFMMCTradingAccountTokenType func(*ctp.CThostFtdcCFMMCTradingAccountTokenField) uintptr
type tOnErrRtnBatchOrderActionType func(*ctp.CThostFtdcBatchOrderActionField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRtnOptionSelfCloseType func(*ctp.CThostFtdcOptionSelfCloseField) uintptr
type tOnErrRtnOptionSelfCloseInsertType func(*ctp.CThostFtdcInputOptionSelfCloseField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnOptionSelfCloseActionType func(*ctp.CThostFtdcOptionSelfCloseActionField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRtnCombActionType func(*ctp.CThostFtdcCombActionField) uintptr
type tOnErrRtnCombActionInsertType func(*ctp.CThostFtdcInputCombActionField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRspQryContractBankType func(*ctp.CThostFtdcContractBankField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryParkedOrderType func(*ctp.CThostFtdcParkedOrderField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryParkedOrderActionType func(*ctp.CThostFtdcParkedOrderActionField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryTradingNoticeType func(*ctp.CThostFtdcTradingNoticeField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryBrokerTradingParamsType func(*ctp.CThostFtdcBrokerTradingParamsField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQryBrokerTradingAlgosType func(*ctp.CThostFtdcBrokerTradingAlgosField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQueryCFMMCTradingAccountTokenType func(*ctp.CThostFtdcQueryCFMMCTradingAccountTokenField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRtnFromBankToFutureByBankType func(*ctp.CThostFtdcRspTransferField) uintptr
type tOnRtnFromFutureToBankByBankType func(*ctp.CThostFtdcRspTransferField) uintptr
type tOnRtnRepealFromBankToFutureByBankType func(*ctp.CThostFtdcRspRepealField) uintptr
type tOnRtnRepealFromFutureToBankByBankType func(*ctp.CThostFtdcRspRepealField) uintptr
type tOnRtnFromBankToFutureByFutureType func(*ctp.CThostFtdcRspTransferField) uintptr
type tOnRtnFromFutureToBankByFutureType func(*ctp.CThostFtdcRspTransferField) uintptr
type tOnRtnRepealFromBankToFutureByFutureManualType func(*ctp.CThostFtdcRspRepealField) uintptr
type tOnRtnRepealFromFutureToBankByFutureManualType func(*ctp.CThostFtdcRspRepealField) uintptr
type tOnRtnQueryBankBalanceByFutureType func(*ctp.CThostFtdcNotifyQueryAccountField) uintptr
type tOnErrRtnBankToFutureByFutureType func(*ctp.CThostFtdcReqTransferField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnFutureToBankByFutureType func(*ctp.CThostFtdcReqTransferField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnRepealBankToFutureByFutureManualType func(*ctp.CThostFtdcReqRepealField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnRepealFutureToBankByFutureManualType func(*ctp.CThostFtdcReqRepealField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnErrRtnQueryBankBalanceByFutureType func(*ctp.CThostFtdcReqQueryAccountField, *ctp.CThostFtdcRspInfoField) uintptr
type tOnRtnRepealFromBankToFutureByFutureType func(*ctp.CThostFtdcRspRepealField) uintptr
type tOnRtnRepealFromFutureToBankByFutureType func(*ctp.CThostFtdcRspRepealField) uintptr
type tOnRspFromBankToFutureByFutureType func(*ctp.CThostFtdcReqTransferField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspFromFutureToBankByFutureType func(*ctp.CThostFtdcReqTransferField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRspQueryBankAccountMoneyByFutureType func(*ctp.CThostFtdcReqQueryAccountField, *ctp.CThostFtdcRspInfoField, int, bool) uintptr
type tOnRtnOpenAccountByBankType func(*ctp.CThostFtdcOpenAccountField) uintptr
type tOnRtnCancelAccountByBankType func(*ctp.CThostFtdcCancelAccountField) uintptr
type tOnRtnChangeAccountByBankType func(*ctp.CThostFtdcChangeAccountField) uintptr

type trade struct {
	h        *syscall.DLL
	api, spi  uintptr
	nRequestID      int
}

func (t *trade) loadDll() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		os.Mkdir("log", os.ModePerm)
	}
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	_ = os.Chdir(dllPath)
	t.h = syscall.MustLoadDLL("ctp_trade.dll")
	
	// 还原到之前的工作目录
	os.Chdir(workPath)
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// 接口
func newTrade() *trade {
	t := new(trade)

	t.loadDll()
	t.api, _, _ = t.h.MustFindProc("CreateApi").Call()
	t.spi, _, _ = t.h.MustFindProc("CreateSpi").Call()
	_, _, _ = t.h.MustFindProc("RegisterSpi").Call(t.api, uintptr(unsafe.Pointer(t.spi)))
	return t
}


// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (t *trade) regOnFrontConnected(on tOnFrontConnectedType){
	_, _, _ = t.h.MustFindProc("SetOnFrontConnected").Call(t.spi, syscall.NewCallback(on))
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (t *trade) regOnFrontDisconnected(on tOnFrontDisconnectedType){
	_, _, _ = t.h.MustFindProc("SetOnFrontDisconnected").Call(t.spi, syscall.NewCallback(on))
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (t *trade) regOnHeartBeatWarning(on tOnHeartBeatWarningType){
	_, _, _ = t.h.MustFindProc("SetOnHeartBeatWarning").Call(t.spi, syscall.NewCallback(on))
}

// 客户端认证响应
func (t *trade) regOnRspAuthenticate(on tOnRspAuthenticateType){
	_, _, _ = t.h.MustFindProc("SetOnRspAuthenticate").Call(t.spi, syscall.NewCallback(on))
}

// 登录请求响应
func (t *trade) regOnRspUserLogin(on tOnRspUserLoginType){
	_, _, _ = t.h.MustFindProc("SetOnRspUserLogin").Call(t.spi, syscall.NewCallback(on))
}

// 登出请求响应
func (t *trade) regOnRspUserLogout(on tOnRspUserLogoutType){
	_, _, _ = t.h.MustFindProc("SetOnRspUserLogout").Call(t.spi, syscall.NewCallback(on))
}

// 用户口令更新请求响应
func (t *trade) regOnRspUserPasswordUpdate(on tOnRspUserPasswordUpdateType){
	_, _, _ = t.h.MustFindProc("SetOnRspUserPasswordUpdate").Call(t.spi, syscall.NewCallback(on))
}

// 资金账户口令更新请求响应
func (t *trade) regOnRspTradingAccountPasswordUpdate(on tOnRspTradingAccountPasswordUpdateType){
	_, _, _ = t.h.MustFindProc("SetOnRspTradingAccountPasswordUpdate").Call(t.spi, syscall.NewCallback(on))
}

// 查询用户当前支持的认证模式的回复
func (t *trade) regOnRspUserAuthMethod(on tOnRspUserAuthMethodType){
	_, _, _ = t.h.MustFindProc("SetOnRspUserAuthMethod").Call(t.spi, syscall.NewCallback(on))
}

// 获取图形验证码请求的回复
func (t *trade) regOnRspGenUserCaptcha(on tOnRspGenUserCaptchaType){
	_, _, _ = t.h.MustFindProc("SetOnRspGenUserCaptcha").Call(t.spi, syscall.NewCallback(on))
}

// 获取短信验证码请求的回复
func (t *trade) regOnRspGenUserText(on tOnRspGenUserTextType){
	_, _, _ = t.h.MustFindProc("SetOnRspGenUserText").Call(t.spi, syscall.NewCallback(on))
}

// 报单录入请求响应
func (t *trade) regOnRspOrderInsert(on tOnRspOrderInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 预埋单录入请求响应
func (t *trade) regOnRspParkedOrderInsert(on tOnRspParkedOrderInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspParkedOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 预埋撤单录入请求响应
func (t *trade) regOnRspParkedOrderAction(on tOnRspParkedOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspParkedOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 报单操作请求响应
func (t *trade) regOnRspOrderAction(on tOnRspOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 查询最大报单数量响应
func (t *trade) regOnRspQueryMaxOrderVolume(on tOnRspQueryMaxOrderVolumeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQueryMaxOrderVolume").Call(t.spi, syscall.NewCallback(on))
}

// 投资者结算结果确认响应
func (t *trade) regOnRspSettlementInfoConfirm(on tOnRspSettlementInfoConfirmType){
	_, _, _ = t.h.MustFindProc("SetOnRspSettlementInfoConfirm").Call(t.spi, syscall.NewCallback(on))
}

// 删除预埋单响应
func (t *trade) regOnRspRemoveParkedOrder(on tOnRspRemoveParkedOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRspRemoveParkedOrder").Call(t.spi, syscall.NewCallback(on))
}

// 删除预埋撤单响应
func (t *trade) regOnRspRemoveParkedOrderAction(on tOnRspRemoveParkedOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspRemoveParkedOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告录入请求响应
func (t *trade) regOnRspExecOrderInsert(on tOnRspExecOrderInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspExecOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告操作请求响应
func (t *trade) regOnRspExecOrderAction(on tOnRspExecOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspExecOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 询价录入请求响应
func (t *trade) regOnRspForQuoteInsert(on tOnRspForQuoteInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspForQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价录入请求响应
func (t *trade) regOnRspQuoteInsert(on tOnRspQuoteInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价操作请求响应
func (t *trade) regOnRspQuoteAction(on tOnRspQuoteActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspQuoteAction").Call(t.spi, syscall.NewCallback(on))
}

// 批量报单操作请求响应
func (t *trade) regOnRspBatchOrderAction(on tOnRspBatchOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspBatchOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲录入请求响应
func (t *trade) regOnRspOptionSelfCloseInsert(on tOnRspOptionSelfCloseInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspOptionSelfCloseInsert").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲操作请求响应
func (t *trade) regOnRspOptionSelfCloseAction(on tOnRspOptionSelfCloseActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspOptionSelfCloseAction").Call(t.spi, syscall.NewCallback(on))
}

// 申请组合录入请求响应
func (t *trade) regOnRspCombActionInsert(on tOnRspCombActionInsertType){
	_, _, _ = t.h.MustFindProc("SetOnRspCombActionInsert").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询报单响应
func (t *trade) regOnRspQryOrder(on tOnRspQryOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryOrder").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询成交响应
func (t *trade) regOnRspQryTrade(on tOnRspQryTradeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryTrade").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者持仓响应
func (t *trade) regOnRspQryInvestorPosition(on tOnRspQryInvestorPositionType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorPosition").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询资金账户响应
func (t *trade) regOnRspQryTradingAccount(on tOnRspQryTradingAccountType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryTradingAccount").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者响应
func (t *trade) regOnRspQryInvestor(on tOnRspQryInvestorType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestor").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易编码响应
func (t *trade) regOnRspQryTradingCode(on tOnRspQryTradingCodeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryTradingCode").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询合约保证金率响应
func (t *trade) regOnRspQryInstrumentMarginRate(on tOnRspQryInstrumentMarginRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrumentMarginRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询合约手续费率响应
func (t *trade) regOnRspQryInstrumentCommissionRate(on tOnRspQryInstrumentCommissionRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrumentCommissionRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易所响应
func (t *trade) regOnRspQryExchange(on tOnRspQryExchangeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchange").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询产品响应
func (t *trade) regOnRspQryProduct(on tOnRspQryProductType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryProduct").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询合约响应
func (t *trade) regOnRspQryInstrument(on tOnRspQryInstrumentType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrument").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询行情响应
func (t *trade) regOnRspQryDepthMarketData(on tOnRspQryDepthMarketDataType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryDepthMarketData").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者结算结果响应
func (t *trade) regOnRspQrySettlementInfo(on tOnRspQrySettlementInfoType){
	_, _, _ = t.h.MustFindProc("SetOnRspQrySettlementInfo").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询转帐银行响应
func (t *trade) regOnRspQryTransferBank(on tOnRspQryTransferBankType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryTransferBank").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者持仓明细响应
func (t *trade) regOnRspQryInvestorPositionDetail(on tOnRspQryInvestorPositionDetailType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorPositionDetail").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询客户通知响应
func (t *trade) regOnRspQryNotice(on tOnRspQryNoticeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryNotice").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询结算信息确认响应
func (t *trade) regOnRspQrySettlementInfoConfirm(on tOnRspQrySettlementInfoConfirmType){
	_, _, _ = t.h.MustFindProc("SetOnRspQrySettlementInfoConfirm").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者持仓明细响应
func (t *trade) regOnRspQryInvestorPositionCombineDetail(on tOnRspQryInvestorPositionCombineDetailType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorPositionCombineDetail").Call(t.spi, syscall.NewCallback(on))
}

// 查询保证金监管系统经纪公司资金账户密钥响应
func (t *trade) regOnRspQryCFMMCTradingAccountKey(on tOnRspQryCFMMCTradingAccountKeyType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryCFMMCTradingAccountKey").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询仓单折抵信息响应
func (t *trade) regOnRspQryEWarrantOffset(on tOnRspQryEWarrantOffsetType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryEWarrantOffset").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者品种/跨品种保证金响应
func (t *trade) regOnRspQryInvestorProductGroupMargin(on tOnRspQryInvestorProductGroupMarginType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorProductGroupMargin").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易所保证金率响应
func (t *trade) regOnRspQryExchangeMarginRate(on tOnRspQryExchangeMarginRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchangeMarginRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易所调整保证金率响应
func (t *trade) regOnRspQryExchangeMarginRateAdjust(on tOnRspQryExchangeMarginRateAdjustType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchangeMarginRateAdjust").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询汇率响应
func (t *trade) regOnRspQryExchangeRate(on tOnRspQryExchangeRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchangeRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询二级代理操作员银期权限响应
func (t *trade) regOnRspQrySecAgentACIDMap(on tOnRspQrySecAgentACIDMapType){
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentACIDMap").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询产品报价汇率
func (t *trade) regOnRspQryProductExchRate(on tOnRspQryProductExchRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryProductExchRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询产品组
func (t *trade) regOnRspQryProductGroup(on tOnRspQryProductGroupType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryProductGroup").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询做市商合约手续费率响应
func (t *trade) regOnRspQryMMInstrumentCommissionRate(on tOnRspQryMMInstrumentCommissionRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryMMInstrumentCommissionRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询做市商期权合约手续费响应
func (t *trade) regOnRspQryMMOptionInstrCommRate(on tOnRspQryMMOptionInstrCommRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryMMOptionInstrCommRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询报单手续费响应
func (t *trade) regOnRspQryInstrumentOrderCommRate(on tOnRspQryInstrumentOrderCommRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrumentOrderCommRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询资金账户响应
func (t *trade) regOnRspQrySecAgentTradingAccount(on tOnRspQrySecAgentTradingAccountType){
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentTradingAccount").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询二级代理商资金校验模式响应
func (t *trade) regOnRspQrySecAgentCheckMode(on tOnRspQrySecAgentCheckModeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentCheckMode").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询二级代理商信息响应
func (t *trade) regOnRspQrySecAgentTradeInfo(on tOnRspQrySecAgentTradeInfoType){
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentTradeInfo").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询期权交易成本响应
func (t *trade) regOnRspQryOptionInstrTradeCost(on tOnRspQryOptionInstrTradeCostType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryOptionInstrTradeCost").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询期权合约手续费响应
func (t *trade) regOnRspQryOptionInstrCommRate(on tOnRspQryOptionInstrCommRateType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryOptionInstrCommRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询执行宣告响应
func (t *trade) regOnRspQryExecOrder(on tOnRspQryExecOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryExecOrder").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询询价响应
func (t *trade) regOnRspQryForQuote(on tOnRspQryForQuoteType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryForQuote").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询报价响应
func (t *trade) regOnRspQryQuote(on tOnRspQryQuoteType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryQuote").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询期权自对冲响应
func (t *trade) regOnRspQryOptionSelfClose(on tOnRspQryOptionSelfCloseType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryOptionSelfClose").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资单元响应
func (t *trade) regOnRspQryInvestUnit(on tOnRspQryInvestUnitType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestUnit").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询组合合约安全系数响应
func (t *trade) regOnRspQryCombInstrumentGuard(on tOnRspQryCombInstrumentGuardType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryCombInstrumentGuard").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询申请组合响应
func (t *trade) regOnRspQryCombAction(on tOnRspQryCombActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryCombAction").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询转帐流水响应
func (t *trade) regOnRspQryTransferSerial(on tOnRspQryTransferSerialType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryTransferSerial").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询银期签约关系响应
func (t *trade) regOnRspQryAccountregister(on tOnRspQryAccountregisterType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryAccountregister").Call(t.spi, syscall.NewCallback(on))
}

// 错误应答
func (t *trade) regOnRspError(on tOnRspErrorType){
	_, _, _ = t.h.MustFindProc("SetOnRspError").Call(t.spi, syscall.NewCallback(on))
}

// 报单通知
func (t *trade) regOnRtnOrder(on tOnRtnOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRtnOrder").Call(t.spi, syscall.NewCallback(on))
}

// 成交通知
func (t *trade) regOnRtnTrade(on tOnRtnTradeType){
	_, _, _ = t.h.MustFindProc("SetOnRtnTrade").Call(t.spi, syscall.NewCallback(on))
}

// 报单录入错误回报
func (t *trade) regOnErrRtnOrderInsert(on tOnErrRtnOrderInsertType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报单操作错误回报
func (t *trade) regOnErrRtnOrderAction(on tOnErrRtnOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 合约交易状态通知
func (t *trade) regOnRtnInstrumentStatus(on tOnRtnInstrumentStatusType){
	_, _, _ = t.h.MustFindProc("SetOnRtnInstrumentStatus").Call(t.spi, syscall.NewCallback(on))
}

// 交易所公告通知
func (t *trade) regOnRtnBulletin(on tOnRtnBulletinType){
	_, _, _ = t.h.MustFindProc("SetOnRtnBulletin").Call(t.spi, syscall.NewCallback(on))
}

// 交易通知
func (t *trade) regOnRtnTradingNotice(on tOnRtnTradingNoticeType){
	_, _, _ = t.h.MustFindProc("SetOnRtnTradingNotice").Call(t.spi, syscall.NewCallback(on))
}

// 提示条件单校验错误
func (t *trade) regOnRtnErrorConditionalOrder(on tOnRtnErrorConditionalOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRtnErrorConditionalOrder").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告通知
func (t *trade) regOnRtnExecOrder(on tOnRtnExecOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRtnExecOrder").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告录入错误回报
func (t *trade) regOnErrRtnExecOrderInsert(on tOnErrRtnExecOrderInsertType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnExecOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告操作错误回报
func (t *trade) regOnErrRtnExecOrderAction(on tOnErrRtnExecOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnExecOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 询价录入错误回报
func (t *trade) regOnErrRtnForQuoteInsert(on tOnErrRtnForQuoteInsertType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnForQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价通知
func (t *trade) regOnRtnQuote(on tOnRtnQuoteType){
	_, _, _ = t.h.MustFindProc("SetOnRtnQuote").Call(t.spi, syscall.NewCallback(on))
}

// 报价录入错误回报
func (t *trade) regOnErrRtnQuoteInsert(on tOnErrRtnQuoteInsertType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价操作错误回报
func (t *trade) regOnErrRtnQuoteAction(on tOnErrRtnQuoteActionType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnQuoteAction").Call(t.spi, syscall.NewCallback(on))
}

// 询价通知
func (t *trade) regOnRtnForQuoteRsp(on tOnRtnForQuoteRspType){
	_, _, _ = t.h.MustFindProc("SetOnRtnForQuoteRsp").Call(t.spi, syscall.NewCallback(on))
}

// 保证金监控中心用户令牌
func (t *trade) regOnRtnCFMMCTradingAccountToken(on tOnRtnCFMMCTradingAccountTokenType){
	_, _, _ = t.h.MustFindProc("SetOnRtnCFMMCTradingAccountToken").Call(t.spi, syscall.NewCallback(on))
}

// 批量报单操作错误回报
func (t *trade) regOnErrRtnBatchOrderAction(on tOnErrRtnBatchOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnBatchOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲通知
func (t *trade) regOnRtnOptionSelfClose(on tOnRtnOptionSelfCloseType){
	_, _, _ = t.h.MustFindProc("SetOnRtnOptionSelfClose").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲录入错误回报
func (t *trade) regOnErrRtnOptionSelfCloseInsert(on tOnErrRtnOptionSelfCloseInsertType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOptionSelfCloseInsert").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲操作错误回报
func (t *trade) regOnErrRtnOptionSelfCloseAction(on tOnErrRtnOptionSelfCloseActionType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOptionSelfCloseAction").Call(t.spi, syscall.NewCallback(on))
}

// 申请组合通知
func (t *trade) regOnRtnCombAction(on tOnRtnCombActionType){
	_, _, _ = t.h.MustFindProc("SetOnRtnCombAction").Call(t.spi, syscall.NewCallback(on))
}

// 申请组合录入错误回报
func (t *trade) regOnErrRtnCombActionInsert(on tOnErrRtnCombActionInsertType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnCombActionInsert").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询签约银行响应
func (t *trade) regOnRspQryContractBank(on tOnRspQryContractBankType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryContractBank").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询预埋单响应
func (t *trade) regOnRspQryParkedOrder(on tOnRspQryParkedOrderType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryParkedOrder").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询预埋撤单响应
func (t *trade) regOnRspQryParkedOrderAction(on tOnRspQryParkedOrderActionType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryParkedOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易通知响应
func (t *trade) regOnRspQryTradingNotice(on tOnRspQryTradingNoticeType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryTradingNotice").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询经纪公司交易参数响应
func (t *trade) regOnRspQryBrokerTradingParams(on tOnRspQryBrokerTradingParamsType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryBrokerTradingParams").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询经纪公司交易算法响应
func (t *trade) regOnRspQryBrokerTradingAlgos(on tOnRspQryBrokerTradingAlgosType){
	_, _, _ = t.h.MustFindProc("SetOnRspQryBrokerTradingAlgos").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询监控中心用户令牌
func (t *trade) regOnRspQueryCFMMCTradingAccountToken(on tOnRspQueryCFMMCTradingAccountTokenType){
	_, _, _ = t.h.MustFindProc("SetOnRspQueryCFMMCTradingAccountToken").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起银行资金转期货通知
func (t *trade) regOnRtnFromBankToFutureByBank(on tOnRtnFromBankToFutureByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnFromBankToFutureByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起期货资金转银行通知
func (t *trade) regOnRtnFromFutureToBankByBank(on tOnRtnFromFutureToBankByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnFromFutureToBankByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起冲正银行转期货通知
func (t *trade) regOnRtnRepealFromBankToFutureByBank(on tOnRtnRepealFromBankToFutureByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromBankToFutureByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起冲正期货转银行通知
func (t *trade) regOnRtnRepealFromFutureToBankByBank(on tOnRtnRepealFromFutureToBankByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromFutureToBankByBank").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起银行资金转期货通知
func (t *trade) regOnRtnFromBankToFutureByFuture(on tOnRtnFromBankToFutureByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRtnFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起期货资金转银行通知
func (t *trade) regOnRtnFromFutureToBankByFuture(on tOnRtnFromFutureToBankByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRtnFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromBankToFutureByFutureManual(on tOnRtnRepealFromBankToFutureByFutureManualType){
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromBankToFutureByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromFutureToBankByFutureManual(on tOnRtnRepealFromFutureToBankByFutureManualType){
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromFutureToBankByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起查询银行余额通知
func (t *trade) regOnRtnQueryBankBalanceByFuture(on tOnRtnQueryBankBalanceByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRtnQueryBankBalanceByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起银行资金转期货错误回报
func (t *trade) regOnErrRtnBankToFutureByFuture(on tOnErrRtnBankToFutureByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起期货资金转银行错误回报
func (t *trade) regOnErrRtnFutureToBankByFuture(on tOnErrRtnFutureToBankByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正银行转期货错误回报
func (t *trade) regOnErrRtnRepealBankToFutureByFutureManual(on tOnErrRtnRepealBankToFutureByFutureManualType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnRepealBankToFutureByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正期货转银行错误回报
func (t *trade) regOnErrRtnRepealFutureToBankByFutureManual(on tOnErrRtnRepealFutureToBankByFutureManualType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnRepealFutureToBankByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起查询银行余额错误回报
func (t *trade) regOnErrRtnQueryBankBalanceByFuture(on tOnErrRtnQueryBankBalanceByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnErrRtnQueryBankBalanceByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromBankToFutureByFuture(on tOnRtnRepealFromBankToFutureByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromFutureToBankByFuture(on tOnRtnRepealFromFutureToBankByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起银行资金转期货应答
func (t *trade) regOnRspFromBankToFutureByFuture(on tOnRspFromBankToFutureByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRspFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起期货资金转银行应答
func (t *trade) regOnRspFromFutureToBankByFuture(on tOnRspFromFutureToBankByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRspFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起查询银行余额应答
func (t *trade) regOnRspQueryBankAccountMoneyByFuture(on tOnRspQueryBankAccountMoneyByFutureType){
	_, _, _ = t.h.MustFindProc("SetOnRspQueryBankAccountMoneyByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起银期开户通知
func (t *trade) regOnRtnOpenAccountByBank(on tOnRtnOpenAccountByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnOpenAccountByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起银期销户通知
func (t *trade) regOnRtnCancelAccountByBank(on tOnRtnCancelAccountByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnCancelAccountByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起变更银行账号通知
func (t *trade) regOnRtnChangeAccountByBank(on tOnRtnChangeAccountByBankType){
	_, _, _ = t.h.MustFindProc("SetOnRtnChangeAccountByBank").Call(t.spi, syscall.NewCallback(on))
}



// 创建TraderApi
func (t *trade) Release(){ 
	_, _, _ = t.h.MustFindProc("Release").Call(t.api)
}

// 初始化
func (t *trade) Init(){ 
	_, _, _ = t.h.MustFindProc("Init").Call(t.api)
}

// 等待接口线程结束运行
func (t *trade) Join(){ 
	_, _, _ = t.h.MustFindProc("Join").Call(t.api)
}

// 注册前置机网络地址
func (t *trade) RegisterFront(pszFrontAddress string){ 
	bs, _ := syscall.BytePtrFromString(pszFrontAddress)
	_, _, _ = t.h.MustFindProc("RegisterFront").Call(t.api, uintptr(unsafe.Pointer(bs)))
}

// @remark RegisterNameServer优先于RegisterFront
func (t *trade) RegisterNameServer(pszNsAddress string){ 
	bs, _ := syscall.BytePtrFromString(pszNsAddress)
	_, _, _ = t.h.MustFindProc("RegisterNameServer").Call(t.api, uintptr(unsafe.Pointer(bs)))
}

// 注册名字服务器用户信息
func (t *trade) RegisterFensUserInfo(){ 
	_, _, _ = t.h.MustFindProc("RegisterFensUserInfo").Call(t.api)
}

// 订阅私有流。
func (t *trade) SubscribePrivateTopic(nResumeType ctp.THOST_TE_RESUME_TYPE){ 
	_, _, _ = t.h.MustFindProc("SubscribePrivateTopic").Call(t.api, uintptr(nResumeType))
}

// 订阅公共流。
func (t *trade) SubscribePublicTopic(nResumeType ctp.THOST_TE_RESUME_TYPE){ 
	_, _, _ = t.h.MustFindProc("SubscribePublicTopic").Call(t.api, uintptr(nResumeType))
}

// 客户端认证请求
func (t *trade) ReqAuthenticate(pReqAuthenticateField ctp.CThostFtdcReqAuthenticateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqAuthenticate").Call(t.api, uintptr(unsafe.Pointer(&pReqAuthenticateField)), uintptr(t.nRequestID))
}

// 注册用户终端信息，用于中继服务器多连接模式
func (t *trade) RegisterUserSystemInfo(pUserSystemInfo ctp.CThostFtdcUserSystemInfoField){ 
	_, _, _ = t.h.MustFindProc("RegisterUserSystemInfo").Call(t.api, uintptr(unsafe.Pointer(&pUserSystemInfo)))
}

// 上报用户终端信息，用于中继服务器操作员登录模式
func (t *trade) SubmitUserSystemInfo(pUserSystemInfo ctp.CThostFtdcUserSystemInfoField){ 
	_, _, _ = t.h.MustFindProc("SubmitUserSystemInfo").Call(t.api, uintptr(unsafe.Pointer(&pUserSystemInfo)))
}

// 用户登录请求
func (t *trade) ReqUserLogin(pReqUserLoginField ctp.CThostFtdcReqUserLoginField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLogin").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(t.nRequestID))
}

// 登出请求
func (t *trade) ReqUserLogout(pUserLogout ctp.CThostFtdcUserLogoutField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLogout").Call(t.api, uintptr(unsafe.Pointer(&pUserLogout)), uintptr(t.nRequestID))
}

// 用户口令更新请求
func (t *trade) ReqUserPasswordUpdate(pUserPasswordUpdate ctp.CThostFtdcUserPasswordUpdateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserPasswordUpdate").Call(t.api, uintptr(unsafe.Pointer(&pUserPasswordUpdate)), uintptr(t.nRequestID))
}

// 资金账户口令更新请求
func (t *trade) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate ctp.CThostFtdcTradingAccountPasswordUpdateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqTradingAccountPasswordUpdate").Call(t.api, uintptr(unsafe.Pointer(&pTradingAccountPasswordUpdate)), uintptr(t.nRequestID))
}

// 查询用户当前支持的认证模式
func (t *trade) ReqUserAuthMethod(pReqUserAuthMethod ctp.CThostFtdcReqUserAuthMethodField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserAuthMethod").Call(t.api, uintptr(unsafe.Pointer(&pReqUserAuthMethod)), uintptr(t.nRequestID))
}

// 用户发出获取图形验证码请求
func (t *trade) ReqGenUserCaptcha(pReqGenUserCaptcha ctp.CThostFtdcReqGenUserCaptchaField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqGenUserCaptcha").Call(t.api, uintptr(unsafe.Pointer(&pReqGenUserCaptcha)), uintptr(t.nRequestID))
}

// 用户发出获取短信验证码请求
func (t *trade) ReqGenUserText(pReqGenUserText ctp.CThostFtdcReqGenUserTextField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqGenUserText").Call(t.api, uintptr(unsafe.Pointer(&pReqGenUserText)), uintptr(t.nRequestID))
}

// 用户发出带有图片验证码的登陆请求
func (t *trade) ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha ctp.CThostFtdcReqUserLoginWithCaptchaField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLoginWithCaptcha").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginWithCaptcha)), uintptr(t.nRequestID))
}

// 用户发出带有短信验证码的登陆请求
func (t *trade) ReqUserLoginWithText(pReqUserLoginWithText ctp.CThostFtdcReqUserLoginWithTextField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLoginWithText").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginWithText)), uintptr(t.nRequestID))
}

// 用户发出带有动态口令的登陆请求
func (t *trade) ReqUserLoginWithOTP(pReqUserLoginWithOTP ctp.CThostFtdcReqUserLoginWithOTPField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLoginWithOTP").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginWithOTP)), uintptr(t.nRequestID))
}

// 报单录入请求
func (t *trade) ReqOrderInsert(pInputOrder ctp.CThostFtdcInputOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputOrder)), uintptr(t.nRequestID))
}

// 预埋单录入请求
func (t *trade) ReqParkedOrderInsert(pParkedOrder ctp.CThostFtdcParkedOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqParkedOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&pParkedOrder)), uintptr(t.nRequestID))
}

// 预埋撤单录入请求
func (t *trade) ReqParkedOrderAction(pParkedOrderAction ctp.CThostFtdcParkedOrderActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqParkedOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pParkedOrderAction)), uintptr(t.nRequestID))
}

// 报单操作请求
func (t *trade) ReqOrderAction(pInputOrderAction ctp.CThostFtdcInputOrderActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pInputOrderAction)), uintptr(t.nRequestID))
}

// 查询最大报单数量请求
func (t *trade) ReqQueryMaxOrderVolume(pQueryMaxOrderVolume ctp.CThostFtdcQueryMaxOrderVolumeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQueryMaxOrderVolume").Call(t.api, uintptr(unsafe.Pointer(&pQueryMaxOrderVolume)), uintptr(t.nRequestID))
}

// 投资者结算结果确认
func (t *trade) ReqSettlementInfoConfirm(pSettlementInfoConfirm ctp.CThostFtdcSettlementInfoConfirmField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqSettlementInfoConfirm").Call(t.api, uintptr(unsafe.Pointer(&pSettlementInfoConfirm)), uintptr(t.nRequestID))
}

// 请求删除预埋单
func (t *trade) ReqRemoveParkedOrder(pRemoveParkedOrder ctp.CThostFtdcRemoveParkedOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqRemoveParkedOrder").Call(t.api, uintptr(unsafe.Pointer(&pRemoveParkedOrder)), uintptr(t.nRequestID))
}

// 请求删除预埋撤单
func (t *trade) ReqRemoveParkedOrderAction(pRemoveParkedOrderAction ctp.CThostFtdcRemoveParkedOrderActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqRemoveParkedOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pRemoveParkedOrderAction)), uintptr(t.nRequestID))
}

// 执行宣告录入请求
func (t *trade) ReqExecOrderInsert(pInputExecOrder ctp.CThostFtdcInputExecOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqExecOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputExecOrder)), uintptr(t.nRequestID))
}

// 执行宣告操作请求
func (t *trade) ReqExecOrderAction(pInputExecOrderAction ctp.CThostFtdcInputExecOrderActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqExecOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pInputExecOrderAction)), uintptr(t.nRequestID))
}

// 询价录入请求
func (t *trade) ReqForQuoteInsert(pInputForQuote ctp.CThostFtdcInputForQuoteField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqForQuoteInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputForQuote)), uintptr(t.nRequestID))
}

// 报价录入请求
func (t *trade) ReqQuoteInsert(pInputQuote ctp.CThostFtdcInputQuoteField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQuoteInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputQuote)), uintptr(t.nRequestID))
}

// 报价操作请求
func (t *trade) ReqQuoteAction(pInputQuoteAction ctp.CThostFtdcInputQuoteActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQuoteAction").Call(t.api, uintptr(unsafe.Pointer(&pInputQuoteAction)), uintptr(t.nRequestID))
}

// 批量报单操作请求
func (t *trade) ReqBatchOrderAction(pInputBatchOrderAction ctp.CThostFtdcInputBatchOrderActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqBatchOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pInputBatchOrderAction)), uintptr(t.nRequestID))
}

// 期权自对冲录入请求
func (t *trade) ReqOptionSelfCloseInsert(pInputOptionSelfClose ctp.CThostFtdcInputOptionSelfCloseField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOptionSelfCloseInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputOptionSelfClose)), uintptr(t.nRequestID))
}

// 期权自对冲操作请求
func (t *trade) ReqOptionSelfCloseAction(pInputOptionSelfCloseAction ctp.CThostFtdcInputOptionSelfCloseActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOptionSelfCloseAction").Call(t.api, uintptr(unsafe.Pointer(&pInputOptionSelfCloseAction)), uintptr(t.nRequestID))
}

// 申请组合录入请求
func (t *trade) ReqCombActionInsert(pInputCombAction ctp.CThostFtdcInputCombActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqCombActionInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputCombAction)), uintptr(t.nRequestID))
}

// 请求查询报单
func (t *trade) ReqQryOrder(pQryOrder ctp.CThostFtdcQryOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOrder").Call(t.api, uintptr(unsafe.Pointer(&pQryOrder)), uintptr(t.nRequestID))
}

// 请求查询成交
func (t *trade) ReqQryTrade(pQryTrade ctp.CThostFtdcQryTradeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTrade").Call(t.api, uintptr(unsafe.Pointer(&pQryTrade)), uintptr(t.nRequestID))
}

// 请求查询投资者持仓
func (t *trade) ReqQryInvestorPosition(pQryInvestorPosition ctp.CThostFtdcQryInvestorPositionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorPosition").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorPosition)), uintptr(t.nRequestID))
}

// 请求查询资金账户
func (t *trade) ReqQryTradingAccount(pQryTradingAccount ctp.CThostFtdcQryTradingAccountField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTradingAccount").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingAccount)), uintptr(t.nRequestID))
}

// 请求查询投资者
func (t *trade) ReqQryInvestor(pQryInvestor ctp.CThostFtdcQryInvestorField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestor").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestor)), uintptr(t.nRequestID))
}

// 请求查询交易编码
func (t *trade) ReqQryTradingCode(pQryTradingCode ctp.CThostFtdcQryTradingCodeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTradingCode").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingCode)), uintptr(t.nRequestID))
}

// 请求查询合约保证金率
func (t *trade) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate ctp.CThostFtdcQryInstrumentMarginRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrumentMarginRate").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrumentMarginRate)), uintptr(t.nRequestID))
}

// 请求查询合约手续费率
func (t *trade) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate ctp.CThostFtdcQryInstrumentCommissionRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrumentCommissionRate").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrumentCommissionRate)), uintptr(t.nRequestID))
}

// 请求查询交易所
func (t *trade) ReqQryExchange(pQryExchange ctp.CThostFtdcQryExchangeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchange").Call(t.api, uintptr(unsafe.Pointer(&pQryExchange)), uintptr(t.nRequestID))
}

// 请求查询产品
func (t *trade) ReqQryProduct(pQryProduct ctp.CThostFtdcQryProductField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryProduct").Call(t.api, uintptr(unsafe.Pointer(&pQryProduct)), uintptr(t.nRequestID))
}

// 请求查询合约
func (t *trade) ReqQryInstrument(pQryInstrument ctp.CThostFtdcQryInstrumentField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrument").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrument)), uintptr(t.nRequestID))
}

// 请求查询行情
func (t *trade) ReqQryDepthMarketData(pQryDepthMarketData ctp.CThostFtdcQryDepthMarketDataField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryDepthMarketData").Call(t.api, uintptr(unsafe.Pointer(&pQryDepthMarketData)), uintptr(t.nRequestID))
}

// 请求查询投资者结算结果
func (t *trade) ReqQrySettlementInfo(pQrySettlementInfo ctp.CThostFtdcQrySettlementInfoField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySettlementInfo").Call(t.api, uintptr(unsafe.Pointer(&pQrySettlementInfo)), uintptr(t.nRequestID))
}

// 请求查询转帐银行
func (t *trade) ReqQryTransferBank(pQryTransferBank ctp.CThostFtdcQryTransferBankField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTransferBank").Call(t.api, uintptr(unsafe.Pointer(&pQryTransferBank)), uintptr(t.nRequestID))
}

// 请求查询投资者持仓明细
func (t *trade) ReqQryInvestorPositionDetail(pQryInvestorPositionDetail ctp.CThostFtdcQryInvestorPositionDetailField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorPositionDetail").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorPositionDetail)), uintptr(t.nRequestID))
}

// 请求查询客户通知
func (t *trade) ReqQryNotice(pQryNotice ctp.CThostFtdcQryNoticeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryNotice").Call(t.api, uintptr(unsafe.Pointer(&pQryNotice)), uintptr(t.nRequestID))
}

// 请求查询结算信息确认
func (t *trade) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm ctp.CThostFtdcQrySettlementInfoConfirmField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySettlementInfoConfirm").Call(t.api, uintptr(unsafe.Pointer(&pQrySettlementInfoConfirm)), uintptr(t.nRequestID))
}

// 请求查询投资者持仓明细
func (t *trade) ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail ctp.CThostFtdcQryInvestorPositionCombineDetailField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorPositionCombineDetail").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorPositionCombineDetail)), uintptr(t.nRequestID))
}

// 请求查询保证金监管系统经纪公司资金账户密钥
func (t *trade) ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey ctp.CThostFtdcQryCFMMCTradingAccountKeyField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryCFMMCTradingAccountKey").Call(t.api, uintptr(unsafe.Pointer(&pQryCFMMCTradingAccountKey)), uintptr(t.nRequestID))
}

// 请求查询仓单折抵信息
func (t *trade) ReqQryEWarrantOffset(pQryEWarrantOffset ctp.CThostFtdcQryEWarrantOffsetField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryEWarrantOffset").Call(t.api, uintptr(unsafe.Pointer(&pQryEWarrantOffset)), uintptr(t.nRequestID))
}

// 请求查询投资者品种/跨品种保证金
func (t *trade) ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin ctp.CThostFtdcQryInvestorProductGroupMarginField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorProductGroupMargin").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorProductGroupMargin)), uintptr(t.nRequestID))
}

// 请求查询交易所保证金率
func (t *trade) ReqQryExchangeMarginRate(pQryExchangeMarginRate ctp.CThostFtdcQryExchangeMarginRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchangeMarginRate").Call(t.api, uintptr(unsafe.Pointer(&pQryExchangeMarginRate)), uintptr(t.nRequestID))
}

// 请求查询交易所调整保证金率
func (t *trade) ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust ctp.CThostFtdcQryExchangeMarginRateAdjustField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchangeMarginRateAdjust").Call(t.api, uintptr(unsafe.Pointer(&pQryExchangeMarginRateAdjust)), uintptr(t.nRequestID))
}

// 请求查询汇率
func (t *trade) ReqQryExchangeRate(pQryExchangeRate ctp.CThostFtdcQryExchangeRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchangeRate").Call(t.api, uintptr(unsafe.Pointer(&pQryExchangeRate)), uintptr(t.nRequestID))
}

// 请求查询二级代理操作员银期权限
func (t *trade) ReqQrySecAgentACIDMap(pQrySecAgentACIDMap ctp.CThostFtdcQrySecAgentACIDMapField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentACIDMap").Call(t.api, uintptr(unsafe.Pointer(&pQrySecAgentACIDMap)), uintptr(t.nRequestID))
}

// 请求查询产品报价汇率
func (t *trade) ReqQryProductExchRate(pQryProductExchRate ctp.CThostFtdcQryProductExchRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryProductExchRate").Call(t.api, uintptr(unsafe.Pointer(&pQryProductExchRate)), uintptr(t.nRequestID))
}

// 请求查询产品组
func (t *trade) ReqQryProductGroup(pQryProductGroup ctp.CThostFtdcQryProductGroupField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryProductGroup").Call(t.api, uintptr(unsafe.Pointer(&pQryProductGroup)), uintptr(t.nRequestID))
}

// 请求查询做市商合约手续费率
func (t *trade) ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate ctp.CThostFtdcQryMMInstrumentCommissionRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryMMInstrumentCommissionRate").Call(t.api, uintptr(unsafe.Pointer(&pQryMMInstrumentCommissionRate)), uintptr(t.nRequestID))
}

// 请求查询做市商期权合约手续费
func (t *trade) ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate ctp.CThostFtdcQryMMOptionInstrCommRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryMMOptionInstrCommRate").Call(t.api, uintptr(unsafe.Pointer(&pQryMMOptionInstrCommRate)), uintptr(t.nRequestID))
}

// 请求查询报单手续费
func (t *trade) ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate ctp.CThostFtdcQryInstrumentOrderCommRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrumentOrderCommRate").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrumentOrderCommRate)), uintptr(t.nRequestID))
}

// 请求查询资金账户
func (t *trade) ReqQrySecAgentTradingAccount(pQryTradingAccount ctp.CThostFtdcQryTradingAccountField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentTradingAccount").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingAccount)), uintptr(t.nRequestID))
}

// 请求查询二级代理商资金校验模式
func (t *trade) ReqQrySecAgentCheckMode(pQrySecAgentCheckMode ctp.CThostFtdcQrySecAgentCheckModeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentCheckMode").Call(t.api, uintptr(unsafe.Pointer(&pQrySecAgentCheckMode)), uintptr(t.nRequestID))
}

// 请求查询二级代理商信息
func (t *trade) ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo ctp.CThostFtdcQrySecAgentTradeInfoField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentTradeInfo").Call(t.api, uintptr(unsafe.Pointer(&pQrySecAgentTradeInfo)), uintptr(t.nRequestID))
}

// 请求查询期权交易成本
func (t *trade) ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost ctp.CThostFtdcQryOptionInstrTradeCostField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOptionInstrTradeCost").Call(t.api, uintptr(unsafe.Pointer(&pQryOptionInstrTradeCost)), uintptr(t.nRequestID))
}

// 请求查询期权合约手续费
func (t *trade) ReqQryOptionInstrCommRate(pQryOptionInstrCommRate ctp.CThostFtdcQryOptionInstrCommRateField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOptionInstrCommRate").Call(t.api, uintptr(unsafe.Pointer(&pQryOptionInstrCommRate)), uintptr(t.nRequestID))
}

// 请求查询执行宣告
func (t *trade) ReqQryExecOrder(pQryExecOrder ctp.CThostFtdcQryExecOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExecOrder").Call(t.api, uintptr(unsafe.Pointer(&pQryExecOrder)), uintptr(t.nRequestID))
}

// 请求查询询价
func (t *trade) ReqQryForQuote(pQryForQuote ctp.CThostFtdcQryForQuoteField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryForQuote").Call(t.api, uintptr(unsafe.Pointer(&pQryForQuote)), uintptr(t.nRequestID))
}

// 请求查询报价
func (t *trade) ReqQryQuote(pQryQuote ctp.CThostFtdcQryQuoteField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryQuote").Call(t.api, uintptr(unsafe.Pointer(&pQryQuote)), uintptr(t.nRequestID))
}

// 请求查询期权自对冲
func (t *trade) ReqQryOptionSelfClose(pQryOptionSelfClose ctp.CThostFtdcQryOptionSelfCloseField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOptionSelfClose").Call(t.api, uintptr(unsafe.Pointer(&pQryOptionSelfClose)), uintptr(t.nRequestID))
}

// 请求查询投资单元
func (t *trade) ReqQryInvestUnit(pQryInvestUnit ctp.CThostFtdcQryInvestUnitField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestUnit").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestUnit)), uintptr(t.nRequestID))
}

// 请求查询组合合约安全系数
func (t *trade) ReqQryCombInstrumentGuard(pQryCombInstrumentGuard ctp.CThostFtdcQryCombInstrumentGuardField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryCombInstrumentGuard").Call(t.api, uintptr(unsafe.Pointer(&pQryCombInstrumentGuard)), uintptr(t.nRequestID))
}

// 请求查询申请组合
func (t *trade) ReqQryCombAction(pQryCombAction ctp.CThostFtdcQryCombActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryCombAction").Call(t.api, uintptr(unsafe.Pointer(&pQryCombAction)), uintptr(t.nRequestID))
}

// 请求查询转帐流水
func (t *trade) ReqQryTransferSerial(pQryTransferSerial ctp.CThostFtdcQryTransferSerialField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTransferSerial").Call(t.api, uintptr(unsafe.Pointer(&pQryTransferSerial)), uintptr(t.nRequestID))
}

// 请求查询银期签约关系
func (t *trade) ReqQryAccountregister(pQryAccountregister ctp.CThostFtdcQryAccountregisterField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryAccountregister").Call(t.api, uintptr(unsafe.Pointer(&pQryAccountregister)), uintptr(t.nRequestID))
}

// 请求查询签约银行
func (t *trade) ReqQryContractBank(pQryContractBank ctp.CThostFtdcQryContractBankField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryContractBank").Call(t.api, uintptr(unsafe.Pointer(&pQryContractBank)), uintptr(t.nRequestID))
}

// 请求查询预埋单
func (t *trade) ReqQryParkedOrder(pQryParkedOrder ctp.CThostFtdcQryParkedOrderField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryParkedOrder").Call(t.api, uintptr(unsafe.Pointer(&pQryParkedOrder)), uintptr(t.nRequestID))
}

// 请求查询预埋撤单
func (t *trade) ReqQryParkedOrderAction(pQryParkedOrderAction ctp.CThostFtdcQryParkedOrderActionField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryParkedOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pQryParkedOrderAction)), uintptr(t.nRequestID))
}

// 请求查询交易通知
func (t *trade) ReqQryTradingNotice(pQryTradingNotice ctp.CThostFtdcQryTradingNoticeField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTradingNotice").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingNotice)), uintptr(t.nRequestID))
}

// 请求查询经纪公司交易参数
func (t *trade) ReqQryBrokerTradingParams(pQryBrokerTradingParams ctp.CThostFtdcQryBrokerTradingParamsField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryBrokerTradingParams").Call(t.api, uintptr(unsafe.Pointer(&pQryBrokerTradingParams)), uintptr(t.nRequestID))
}

// 请求查询经纪公司交易算法
func (t *trade) ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos ctp.CThostFtdcQryBrokerTradingAlgosField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryBrokerTradingAlgos").Call(t.api, uintptr(unsafe.Pointer(&pQryBrokerTradingAlgos)), uintptr(t.nRequestID))
}

// 请求查询监控中心用户令牌
func (t *trade) ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken ctp.CThostFtdcQueryCFMMCTradingAccountTokenField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQueryCFMMCTradingAccountToken").Call(t.api, uintptr(unsafe.Pointer(&pQueryCFMMCTradingAccountToken)), uintptr(t.nRequestID))
}

// 期货发起银行资金转期货请求
func (t *trade) ReqFromBankToFutureByFuture(pReqTransfer ctp.CThostFtdcReqTransferField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqFromBankToFutureByFuture").Call(t.api, uintptr(unsafe.Pointer(&pReqTransfer)), uintptr(t.nRequestID))
}

// 期货发起期货资金转银行请求
func (t *trade) ReqFromFutureToBankByFuture(pReqTransfer ctp.CThostFtdcReqTransferField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqFromFutureToBankByFuture").Call(t.api, uintptr(unsafe.Pointer(&pReqTransfer)), uintptr(t.nRequestID))
}

// 期货发起查询银行余额请求
func (t *trade) ReqQueryBankAccountMoneyByFuture(pReqQueryAccount ctp.CThostFtdcReqQueryAccountField){ 
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQueryBankAccountMoneyByFuture").Call(t.api, uintptr(unsafe.Pointer(&pReqQueryAccount)), uintptr(t.nRequestID))
}
