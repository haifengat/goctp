package go_ctp_win

import (
	"hf_go_ctp"
	"hf_go_ctp/go_ctp"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

type CThostFtdcTraderSpi uintptr
type OnFrontConnectedType func() uintptr
type OnFrontDisconnectedType func(int) uintptr
type OnHeartBeatWarningType func(int) uintptr
type OnRspAuthenticateType func(*go_ctp.CThostFtdcRspAuthenticateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspUserLoginType func(*go_ctp.CThostFtdcRspUserLoginField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspUserLogoutType func(*go_ctp.CThostFtdcUserLogoutField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspUserPasswordUpdateType func(*go_ctp.CThostFtdcUserPasswordUpdateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspTradingAccountPasswordUpdateType func(*go_ctp.CThostFtdcTradingAccountPasswordUpdateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspUserAuthMethodType func(*go_ctp.CThostFtdcRspUserAuthMethodField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspGenUserCaptchaType func(*go_ctp.CThostFtdcRspGenUserCaptchaField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspGenUserTextType func(*go_ctp.CThostFtdcRspGenUserTextField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspOrderInsertType func(*go_ctp.CThostFtdcInputOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspParkedOrderInsertType func(*go_ctp.CThostFtdcParkedOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspParkedOrderActionType func(*go_ctp.CThostFtdcParkedOrderActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspOrderActionType func(*go_ctp.CThostFtdcInputOrderActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQueryMaxOrderVolumeType func(*go_ctp.CThostFtdcQueryMaxOrderVolumeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspSettlementInfoConfirmType func(*go_ctp.CThostFtdcSettlementInfoConfirmField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspRemoveParkedOrderType func(*go_ctp.CThostFtdcRemoveParkedOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspRemoveParkedOrderActionType func(*go_ctp.CThostFtdcRemoveParkedOrderActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspExecOrderInsertType func(*go_ctp.CThostFtdcInputExecOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspExecOrderActionType func(*go_ctp.CThostFtdcInputExecOrderActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspForQuoteInsertType func(*go_ctp.CThostFtdcInputForQuoteField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQuoteInsertType func(*go_ctp.CThostFtdcInputQuoteField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQuoteActionType func(*go_ctp.CThostFtdcInputQuoteActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspBatchOrderActionType func(*go_ctp.CThostFtdcInputBatchOrderActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspOptionSelfCloseInsertType func(*go_ctp.CThostFtdcInputOptionSelfCloseField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspOptionSelfCloseActionType func(*go_ctp.CThostFtdcInputOptionSelfCloseActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspCombActionInsertType func(*go_ctp.CThostFtdcInputCombActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryOrderType func(*go_ctp.CThostFtdcOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryTradeType func(*go_ctp.CThostFtdcTradeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInvestorPositionType func(*go_ctp.CThostFtdcInvestorPositionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryTradingAccountType func(*go_ctp.CThostFtdcTradingAccountField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInvestorType func(*go_ctp.CThostFtdcInvestorField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryTradingCodeType func(*go_ctp.CThostFtdcTradingCodeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInstrumentMarginRateType func(*go_ctp.CThostFtdcInstrumentMarginRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInstrumentCommissionRateType func(*go_ctp.CThostFtdcInstrumentCommissionRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryExchangeType func(*go_ctp.CThostFtdcExchangeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryProductType func(*go_ctp.CThostFtdcProductField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInstrumentType func(*go_ctp.CThostFtdcInstrumentField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryDepthMarketDataType func(*go_ctp.CThostFtdcDepthMarketDataField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQrySettlementInfoType func(*go_ctp.CThostFtdcSettlementInfoField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryTransferBankType func(*go_ctp.CThostFtdcTransferBankField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInvestorPositionDetailType func(*go_ctp.CThostFtdcInvestorPositionDetailField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryNoticeType func(*go_ctp.CThostFtdcNoticeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQrySettlementInfoConfirmType func(*go_ctp.CThostFtdcSettlementInfoConfirmField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInvestorPositionCombineDetailType func(*go_ctp.CThostFtdcInvestorPositionCombineDetailField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryCFMMCTradingAccountKeyType func(*go_ctp.CThostFtdcCFMMCTradingAccountKeyField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryEWarrantOffsetType func(*go_ctp.CThostFtdcEWarrantOffsetField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInvestorProductGroupMarginType func(*go_ctp.CThostFtdcInvestorProductGroupMarginField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryExchangeMarginRateType func(*go_ctp.CThostFtdcExchangeMarginRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryExchangeMarginRateAdjustType func(*go_ctp.CThostFtdcExchangeMarginRateAdjustField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryExchangeRateType func(*go_ctp.CThostFtdcExchangeRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQrySecAgentACIDMapType func(*go_ctp.CThostFtdcSecAgentACIDMapField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryProductExchRateType func(*go_ctp.CThostFtdcProductExchRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryProductGroupType func(*go_ctp.CThostFtdcProductGroupField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryMMInstrumentCommissionRateType func(*go_ctp.CThostFtdcMMInstrumentCommissionRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryMMOptionInstrCommRateType func(*go_ctp.CThostFtdcMMOptionInstrCommRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInstrumentOrderCommRateType func(*go_ctp.CThostFtdcInstrumentOrderCommRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQrySecAgentTradingAccountType func(*go_ctp.CThostFtdcTradingAccountField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQrySecAgentCheckModeType func(*go_ctp.CThostFtdcSecAgentCheckModeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQrySecAgentTradeInfoType func(*go_ctp.CThostFtdcSecAgentTradeInfoField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryOptionInstrTradeCostType func(*go_ctp.CThostFtdcOptionInstrTradeCostField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryOptionInstrCommRateType func(*go_ctp.CThostFtdcOptionInstrCommRateField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryExecOrderType func(*go_ctp.CThostFtdcExecOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryForQuoteType func(*go_ctp.CThostFtdcForQuoteField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryQuoteType func(*go_ctp.CThostFtdcQuoteField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryOptionSelfCloseType func(*go_ctp.CThostFtdcOptionSelfCloseField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryInvestUnitType func(*go_ctp.CThostFtdcInvestUnitField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryCombInstrumentGuardType func(*go_ctp.CThostFtdcCombInstrumentGuardField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryCombActionType func(*go_ctp.CThostFtdcCombActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryTransferSerialType func(*go_ctp.CThostFtdcTransferSerialField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryAccountregisterType func(*go_ctp.CThostFtdcAccountregisterField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspErrorType func(*go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRtnOrderType func(*go_ctp.CThostFtdcOrderField) uintptr
type OnRtnTradeType func(*go_ctp.CThostFtdcTradeField) uintptr
type OnErrRtnOrderInsertType func(*go_ctp.CThostFtdcInputOrderField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnOrderActionType func(*go_ctp.CThostFtdcOrderActionField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRtnInstrumentStatusType func(*go_ctp.CThostFtdcInstrumentStatusField) uintptr
type OnRtnBulletinType func(*go_ctp.CThostFtdcBulletinField) uintptr
type OnRtnTradingNoticeType func(*go_ctp.CThostFtdcTradingNoticeInfoField) uintptr
type OnRtnErrorConditionalOrderType func(*go_ctp.CThostFtdcErrorConditionalOrderField) uintptr
type OnRtnExecOrderType func(*go_ctp.CThostFtdcExecOrderField) uintptr
type OnErrRtnExecOrderInsertType func(*go_ctp.CThostFtdcInputExecOrderField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnExecOrderActionType func(*go_ctp.CThostFtdcExecOrderActionField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnForQuoteInsertType func(*go_ctp.CThostFtdcInputForQuoteField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRtnQuoteType func(*go_ctp.CThostFtdcQuoteField) uintptr
type OnErrRtnQuoteInsertType func(*go_ctp.CThostFtdcInputQuoteField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnQuoteActionType func(*go_ctp.CThostFtdcQuoteActionField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRtnForQuoteRspType func(*go_ctp.CThostFtdcForQuoteRspField) uintptr
type OnRtnCFMMCTradingAccountTokenType func(*go_ctp.CThostFtdcCFMMCTradingAccountTokenField) uintptr
type OnErrRtnBatchOrderActionType func(*go_ctp.CThostFtdcBatchOrderActionField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRtnOptionSelfCloseType func(*go_ctp.CThostFtdcOptionSelfCloseField) uintptr
type OnErrRtnOptionSelfCloseInsertType func(*go_ctp.CThostFtdcInputOptionSelfCloseField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnOptionSelfCloseActionType func(*go_ctp.CThostFtdcOptionSelfCloseActionField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRtnCombActionType func(*go_ctp.CThostFtdcCombActionField) uintptr
type OnErrRtnCombActionInsertType func(*go_ctp.CThostFtdcInputCombActionField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRspQryContractBankType func(*go_ctp.CThostFtdcContractBankField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryParkedOrderType func(*go_ctp.CThostFtdcParkedOrderField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryParkedOrderActionType func(*go_ctp.CThostFtdcParkedOrderActionField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryTradingNoticeType func(*go_ctp.CThostFtdcTradingNoticeField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryBrokerTradingParamsType func(*go_ctp.CThostFtdcBrokerTradingParamsField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQryBrokerTradingAlgosType func(*go_ctp.CThostFtdcBrokerTradingAlgosField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQueryCFMMCTradingAccountTokenType func(*go_ctp.CThostFtdcQueryCFMMCTradingAccountTokenField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRtnFromBankToFutureByBankType func(*go_ctp.CThostFtdcRspTransferField) uintptr
type OnRtnFromFutureToBankByBankType func(*go_ctp.CThostFtdcRspTransferField) uintptr
type OnRtnRepealFromBankToFutureByBankType func(*go_ctp.CThostFtdcRspRepealField) uintptr
type OnRtnRepealFromFutureToBankByBankType func(*go_ctp.CThostFtdcRspRepealField) uintptr
type OnRtnFromBankToFutureByFutureType func(*go_ctp.CThostFtdcRspTransferField) uintptr
type OnRtnFromFutureToBankByFutureType func(*go_ctp.CThostFtdcRspTransferField) uintptr
type OnRtnRepealFromBankToFutureByFutureManualType func(*go_ctp.CThostFtdcRspRepealField) uintptr
type OnRtnRepealFromFutureToBankByFutureManualType func(*go_ctp.CThostFtdcRspRepealField) uintptr
type OnRtnQueryBankBalanceByFutureType func(*go_ctp.CThostFtdcNotifyQueryAccountField) uintptr
type OnErrRtnBankToFutureByFutureType func(*go_ctp.CThostFtdcReqTransferField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnFutureToBankByFutureType func(*go_ctp.CThostFtdcReqTransferField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnRepealBankToFutureByFutureManualType func(*go_ctp.CThostFtdcReqRepealField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnRepealFutureToBankByFutureManualType func(*go_ctp.CThostFtdcReqRepealField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnErrRtnQueryBankBalanceByFutureType func(*go_ctp.CThostFtdcReqQueryAccountField, *go_ctp.CThostFtdcRspInfoField) uintptr
type OnRtnRepealFromBankToFutureByFutureType func(*go_ctp.CThostFtdcRspRepealField) uintptr
type OnRtnRepealFromFutureToBankByFutureType func(*go_ctp.CThostFtdcRspRepealField) uintptr
type OnRspFromBankToFutureByFutureType func(*go_ctp.CThostFtdcReqTransferField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspFromFutureToBankByFutureType func(*go_ctp.CThostFtdcReqTransferField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRspQueryBankAccountMoneyByFutureType func(*go_ctp.CThostFtdcReqQueryAccountField, *go_ctp.CThostFtdcRspInfoField, int, bool) uintptr
type OnRtnOpenAccountByBankType func(*go_ctp.CThostFtdcOpenAccountField) uintptr
type OnRtnCancelAccountByBankType func(*go_ctp.CThostFtdcCancelAccountField) uintptr
type OnRtnChangeAccountByBankType func(*go_ctp.CThostFtdcChangeAccountField) uintptr

type trade struct {
	h          *syscall.DLL
	api, spi   uintptr
	nRequestID int
}

func (t *trade) loadDll() {
	// 执行目录下创建 log目录
	_, err := os.Stat("log")
	if err != nil {
		hf_go_ctp.CheckErr(os.Mkdir("log", os.ModePerm))
	}
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)                      // 当前文件所在文件夹
	dllPath = dllPath[0:strings.LastIndex(dllPath, "\\")] // 上级文件夹
	dllPath = path.Join(dllPath, "go_ctp")                // 上级下的go_ctp
	if 32<<(^uint(0)>>63) == 64 {
		err = os.Chdir(path.Join(dllPath, "lib64"))
	} else {
		err = os.Chdir(path.Join(dllPath, "lib32"))
	}
	hf_go_ctp.CheckErr(err)
	t.h = syscall.MustLoadDLL("ctp_trade.dll")

	// 还原到之前的工作目录
	hf_go_ctp.CheckErr(os.Chdir(workPath))
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
func (t *trade) regOnFrontConnected(on OnFrontConnectedType) {
	_, _, _ = t.h.MustFindProc("SetOnFrontConnected").Call(t.spi, syscall.NewCallback(on))
}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
func (t *trade) regOnFrontDisconnected(on OnFrontDisconnectedType) {
	_, _, _ = t.h.MustFindProc("SetOnFrontDisconnected").Call(t.spi, syscall.NewCallback(on))
}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (t *trade) regOnHeartBeatWarning(on OnHeartBeatWarningType) {
	_, _, _ = t.h.MustFindProc("SetOnHeartBeatWarning").Call(t.spi, syscall.NewCallback(on))
}

// 客户端认证响应
func (t *trade) regOnRspAuthenticate(on OnRspAuthenticateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspAuthenticate").Call(t.spi, syscall.NewCallback(on))
}

// 登录请求响应
func (t *trade) regOnRspUserLogin(on OnRspUserLoginType) {
	_, _, _ = t.h.MustFindProc("SetOnRspUserLogin").Call(t.spi, syscall.NewCallback(on))
}

// 登出请求响应
func (t *trade) regOnRspUserLogout(on OnRspUserLogoutType) {
	_, _, _ = t.h.MustFindProc("SetOnRspUserLogout").Call(t.spi, syscall.NewCallback(on))
}

// 用户口令更新请求响应
func (t *trade) regOnRspUserPasswordUpdate(on OnRspUserPasswordUpdateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspUserPasswordUpdate").Call(t.spi, syscall.NewCallback(on))
}

// 资金账户口令更新请求响应
func (t *trade) regOnRspTradingAccountPasswordUpdate(on OnRspTradingAccountPasswordUpdateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspTradingAccountPasswordUpdate").Call(t.spi, syscall.NewCallback(on))
}

// 查询用户当前支持的认证模式的回复
func (t *trade) regOnRspUserAuthMethod(on OnRspUserAuthMethodType) {
	_, _, _ = t.h.MustFindProc("SetOnRspUserAuthMethod").Call(t.spi, syscall.NewCallback(on))
}

// 获取图形验证码请求的回复
func (t *trade) regOnRspGenUserCaptcha(on OnRspGenUserCaptchaType) {
	_, _, _ = t.h.MustFindProc("SetOnRspGenUserCaptcha").Call(t.spi, syscall.NewCallback(on))
}

// 获取短信验证码请求的回复
func (t *trade) regOnRspGenUserText(on OnRspGenUserTextType) {
	_, _, _ = t.h.MustFindProc("SetOnRspGenUserText").Call(t.spi, syscall.NewCallback(on))
}

// 报单录入请求响应
func (t *trade) regOnRspOrderInsert(on OnRspOrderInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 预埋单录入请求响应
func (t *trade) regOnRspParkedOrderInsert(on OnRspParkedOrderInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspParkedOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 预埋撤单录入请求响应
func (t *trade) regOnRspParkedOrderAction(on OnRspParkedOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspParkedOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 报单操作请求响应
func (t *trade) regOnRspOrderAction(on OnRspOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 查询最大报单数量响应
func (t *trade) regOnRspQueryMaxOrderVolume(on OnRspQueryMaxOrderVolumeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQueryMaxOrderVolume").Call(t.spi, syscall.NewCallback(on))
}

// 投资者结算结果确认响应
func (t *trade) regOnRspSettlementInfoConfirm(on OnRspSettlementInfoConfirmType) {
	_, _, _ = t.h.MustFindProc("SetOnRspSettlementInfoConfirm").Call(t.spi, syscall.NewCallback(on))
}

// 删除预埋单响应
func (t *trade) regOnRspRemoveParkedOrder(on OnRspRemoveParkedOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRspRemoveParkedOrder").Call(t.spi, syscall.NewCallback(on))
}

// 删除预埋撤单响应
func (t *trade) regOnRspRemoveParkedOrderAction(on OnRspRemoveParkedOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspRemoveParkedOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告录入请求响应
func (t *trade) regOnRspExecOrderInsert(on OnRspExecOrderInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspExecOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告操作请求响应
func (t *trade) regOnRspExecOrderAction(on OnRspExecOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspExecOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 询价录入请求响应
func (t *trade) regOnRspForQuoteInsert(on OnRspForQuoteInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspForQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价录入请求响应
func (t *trade) regOnRspQuoteInsert(on OnRspQuoteInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价操作请求响应
func (t *trade) regOnRspQuoteAction(on OnRspQuoteActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQuoteAction").Call(t.spi, syscall.NewCallback(on))
}

// 批量报单操作请求响应
func (t *trade) regOnRspBatchOrderAction(on OnRspBatchOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspBatchOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲录入请求响应
func (t *trade) regOnRspOptionSelfCloseInsert(on OnRspOptionSelfCloseInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspOptionSelfCloseInsert").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲操作请求响应
func (t *trade) regOnRspOptionSelfCloseAction(on OnRspOptionSelfCloseActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspOptionSelfCloseAction").Call(t.spi, syscall.NewCallback(on))
}

// 申请组合录入请求响应
func (t *trade) regOnRspCombActionInsert(on OnRspCombActionInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnRspCombActionInsert").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询报单响应
func (t *trade) regOnRspQryOrder(on OnRspQryOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryOrder").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询成交响应
func (t *trade) regOnRspQryTrade(on OnRspQryTradeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryTrade").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者持仓响应
func (t *trade) regOnRspQryInvestorPosition(on OnRspQryInvestorPositionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorPosition").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询资金账户响应
func (t *trade) regOnRspQryTradingAccount(on OnRspQryTradingAccountType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryTradingAccount").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者响应
func (t *trade) regOnRspQryInvestor(on OnRspQryInvestorType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestor").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易编码响应
func (t *trade) regOnRspQryTradingCode(on OnRspQryTradingCodeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryTradingCode").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询合约保证金率响应
func (t *trade) regOnRspQryInstrumentMarginRate(on OnRspQryInstrumentMarginRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrumentMarginRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询合约手续费率响应
func (t *trade) regOnRspQryInstrumentCommissionRate(on OnRspQryInstrumentCommissionRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrumentCommissionRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易所响应
func (t *trade) regOnRspQryExchange(on OnRspQryExchangeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchange").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询产品响应
func (t *trade) regOnRspQryProduct(on OnRspQryProductType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryProduct").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询合约响应
func (t *trade) regOnRspQryInstrument(on OnRspQryInstrumentType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrument").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询行情响应
func (t *trade) regOnRspQryDepthMarketData(on OnRspQryDepthMarketDataType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryDepthMarketData").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者结算结果响应
func (t *trade) regOnRspQrySettlementInfo(on OnRspQrySettlementInfoType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQrySettlementInfo").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询转帐银行响应
func (t *trade) regOnRspQryTransferBank(on OnRspQryTransferBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryTransferBank").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者持仓明细响应
func (t *trade) regOnRspQryInvestorPositionDetail(on OnRspQryInvestorPositionDetailType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorPositionDetail").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询客户通知响应
func (t *trade) regOnRspQryNotice(on OnRspQryNoticeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryNotice").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询结算信息确认响应
func (t *trade) regOnRspQrySettlementInfoConfirm(on OnRspQrySettlementInfoConfirmType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQrySettlementInfoConfirm").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者持仓明细响应
func (t *trade) regOnRspQryInvestorPositionCombineDetail(on OnRspQryInvestorPositionCombineDetailType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorPositionCombineDetail").Call(t.spi, syscall.NewCallback(on))
}

// 查询保证金监管系统经纪公司资金账户密钥响应
func (t *trade) regOnRspQryCFMMCTradingAccountKey(on OnRspQryCFMMCTradingAccountKeyType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryCFMMCTradingAccountKey").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询仓单折抵信息响应
func (t *trade) regOnRspQryEWarrantOffset(on OnRspQryEWarrantOffsetType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryEWarrantOffset").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资者品种/跨品种保证金响应
func (t *trade) regOnRspQryInvestorProductGroupMargin(on OnRspQryInvestorProductGroupMarginType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestorProductGroupMargin").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易所保证金率响应
func (t *trade) regOnRspQryExchangeMarginRate(on OnRspQryExchangeMarginRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchangeMarginRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易所调整保证金率响应
func (t *trade) regOnRspQryExchangeMarginRateAdjust(on OnRspQryExchangeMarginRateAdjustType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchangeMarginRateAdjust").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询汇率响应
func (t *trade) regOnRspQryExchangeRate(on OnRspQryExchangeRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryExchangeRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询二级代理操作员银期权限响应
func (t *trade) regOnRspQrySecAgentACIDMap(on OnRspQrySecAgentACIDMapType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentACIDMap").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询产品报价汇率
func (t *trade) regOnRspQryProductExchRate(on OnRspQryProductExchRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryProductExchRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询产品组
func (t *trade) regOnRspQryProductGroup(on OnRspQryProductGroupType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryProductGroup").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询做市商合约手续费率响应
func (t *trade) regOnRspQryMMInstrumentCommissionRate(on OnRspQryMMInstrumentCommissionRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryMMInstrumentCommissionRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询做市商期权合约手续费响应
func (t *trade) regOnRspQryMMOptionInstrCommRate(on OnRspQryMMOptionInstrCommRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryMMOptionInstrCommRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询报单手续费响应
func (t *trade) regOnRspQryInstrumentOrderCommRate(on OnRspQryInstrumentOrderCommRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInstrumentOrderCommRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询资金账户响应
func (t *trade) regOnRspQrySecAgentTradingAccount(on OnRspQrySecAgentTradingAccountType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentTradingAccount").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询二级代理商资金校验模式响应
func (t *trade) regOnRspQrySecAgentCheckMode(on OnRspQrySecAgentCheckModeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentCheckMode").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询二级代理商信息响应
func (t *trade) regOnRspQrySecAgentTradeInfo(on OnRspQrySecAgentTradeInfoType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQrySecAgentTradeInfo").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询期权交易成本响应
func (t *trade) regOnRspQryOptionInstrTradeCost(on OnRspQryOptionInstrTradeCostType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryOptionInstrTradeCost").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询期权合约手续费响应
func (t *trade) regOnRspQryOptionInstrCommRate(on OnRspQryOptionInstrCommRateType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryOptionInstrCommRate").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询执行宣告响应
func (t *trade) regOnRspQryExecOrder(on OnRspQryExecOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryExecOrder").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询询价响应
func (t *trade) regOnRspQryForQuote(on OnRspQryForQuoteType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryForQuote").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询报价响应
func (t *trade) regOnRspQryQuote(on OnRspQryQuoteType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryQuote").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询期权自对冲响应
func (t *trade) regOnRspQryOptionSelfClose(on OnRspQryOptionSelfCloseType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryOptionSelfClose").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询投资单元响应
func (t *trade) regOnRspQryInvestUnit(on OnRspQryInvestUnitType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryInvestUnit").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询组合合约安全系数响应
func (t *trade) regOnRspQryCombInstrumentGuard(on OnRspQryCombInstrumentGuardType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryCombInstrumentGuard").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询申请组合响应
func (t *trade) regOnRspQryCombAction(on OnRspQryCombActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryCombAction").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询转帐流水响应
func (t *trade) regOnRspQryTransferSerial(on OnRspQryTransferSerialType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryTransferSerial").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询银期签约关系响应
func (t *trade) regOnRspQryAccountregister(on OnRspQryAccountregisterType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryAccountregister").Call(t.spi, syscall.NewCallback(on))
}

// 错误应答
func (t *trade) regOnRspError(on OnRspErrorType) {
	_, _, _ = t.h.MustFindProc("SetOnRspError").Call(t.spi, syscall.NewCallback(on))
}

// 报单通知
func (t *trade) regOnRtnOrder(on OnRtnOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnOrder").Call(t.spi, syscall.NewCallback(on))
}

// 成交通知
func (t *trade) regOnRtnTrade(on OnRtnTradeType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnTrade").Call(t.spi, syscall.NewCallback(on))
}

// 报单录入错误回报
func (t *trade) regOnErrRtnOrderInsert(on OnErrRtnOrderInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报单操作错误回报
func (t *trade) regOnErrRtnOrderAction(on OnErrRtnOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 合约交易状态通知
func (t *trade) regOnRtnInstrumentStatus(on OnRtnInstrumentStatusType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnInstrumentStatus").Call(t.spi, syscall.NewCallback(on))
}

// 交易所公告通知
func (t *trade) regOnRtnBulletin(on OnRtnBulletinType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnBulletin").Call(t.spi, syscall.NewCallback(on))
}

// 交易通知
func (t *trade) regOnRtnTradingNotice(on OnRtnTradingNoticeType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnTradingNotice").Call(t.spi, syscall.NewCallback(on))
}

// 提示条件单校验错误
func (t *trade) regOnRtnErrorConditionalOrder(on OnRtnErrorConditionalOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnErrorConditionalOrder").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告通知
func (t *trade) regOnRtnExecOrder(on OnRtnExecOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnExecOrder").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告录入错误回报
func (t *trade) regOnErrRtnExecOrderInsert(on OnErrRtnExecOrderInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnExecOrderInsert").Call(t.spi, syscall.NewCallback(on))
}

// 执行宣告操作错误回报
func (t *trade) regOnErrRtnExecOrderAction(on OnErrRtnExecOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnExecOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 询价录入错误回报
func (t *trade) regOnErrRtnForQuoteInsert(on OnErrRtnForQuoteInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnForQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价通知
func (t *trade) regOnRtnQuote(on OnRtnQuoteType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnQuote").Call(t.spi, syscall.NewCallback(on))
}

// 报价录入错误回报
func (t *trade) regOnErrRtnQuoteInsert(on OnErrRtnQuoteInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnQuoteInsert").Call(t.spi, syscall.NewCallback(on))
}

// 报价操作错误回报
func (t *trade) regOnErrRtnQuoteAction(on OnErrRtnQuoteActionType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnQuoteAction").Call(t.spi, syscall.NewCallback(on))
}

// 询价通知
func (t *trade) regOnRtnForQuoteRsp(on OnRtnForQuoteRspType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnForQuoteRsp").Call(t.spi, syscall.NewCallback(on))
}

// 保证金监控中心用户令牌
func (t *trade) regOnRtnCFMMCTradingAccountToken(on OnRtnCFMMCTradingAccountTokenType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnCFMMCTradingAccountToken").Call(t.spi, syscall.NewCallback(on))
}

// 批量报单操作错误回报
func (t *trade) regOnErrRtnBatchOrderAction(on OnErrRtnBatchOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnBatchOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲通知
func (t *trade) regOnRtnOptionSelfClose(on OnRtnOptionSelfCloseType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnOptionSelfClose").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲录入错误回报
func (t *trade) regOnErrRtnOptionSelfCloseInsert(on OnErrRtnOptionSelfCloseInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOptionSelfCloseInsert").Call(t.spi, syscall.NewCallback(on))
}

// 期权自对冲操作错误回报
func (t *trade) regOnErrRtnOptionSelfCloseAction(on OnErrRtnOptionSelfCloseActionType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnOptionSelfCloseAction").Call(t.spi, syscall.NewCallback(on))
}

// 申请组合通知
func (t *trade) regOnRtnCombAction(on OnRtnCombActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnCombAction").Call(t.spi, syscall.NewCallback(on))
}

// 申请组合录入错误回报
func (t *trade) regOnErrRtnCombActionInsert(on OnErrRtnCombActionInsertType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnCombActionInsert").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询签约银行响应
func (t *trade) regOnRspQryContractBank(on OnRspQryContractBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryContractBank").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询预埋单响应
func (t *trade) regOnRspQryParkedOrder(on OnRspQryParkedOrderType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryParkedOrder").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询预埋撤单响应
func (t *trade) regOnRspQryParkedOrderAction(on OnRspQryParkedOrderActionType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryParkedOrderAction").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询交易通知响应
func (t *trade) regOnRspQryTradingNotice(on OnRspQryTradingNoticeType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryTradingNotice").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询经纪公司交易参数响应
func (t *trade) regOnRspQryBrokerTradingParams(on OnRspQryBrokerTradingParamsType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryBrokerTradingParams").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询经纪公司交易算法响应
func (t *trade) regOnRspQryBrokerTradingAlgos(on OnRspQryBrokerTradingAlgosType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQryBrokerTradingAlgos").Call(t.spi, syscall.NewCallback(on))
}

// 请求查询监控中心用户令牌
func (t *trade) regOnRspQueryCFMMCTradingAccountToken(on OnRspQueryCFMMCTradingAccountTokenType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQueryCFMMCTradingAccountToken").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起银行资金转期货通知
func (t *trade) regOnRtnFromBankToFutureByBank(on OnRtnFromBankToFutureByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnFromBankToFutureByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起期货资金转银行通知
func (t *trade) regOnRtnFromFutureToBankByBank(on OnRtnFromFutureToBankByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnFromFutureToBankByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起冲正银行转期货通知
func (t *trade) regOnRtnRepealFromBankToFutureByBank(on OnRtnRepealFromBankToFutureByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromBankToFutureByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起冲正期货转银行通知
func (t *trade) regOnRtnRepealFromFutureToBankByBank(on OnRtnRepealFromFutureToBankByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromFutureToBankByBank").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起银行资金转期货通知
func (t *trade) regOnRtnFromBankToFutureByFuture(on OnRtnFromBankToFutureByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起期货资金转银行通知
func (t *trade) regOnRtnFromFutureToBankByFuture(on OnRtnFromFutureToBankByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromBankToFutureByFutureManual(on OnRtnRepealFromBankToFutureByFutureManualType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromBankToFutureByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromFutureToBankByFutureManual(on OnRtnRepealFromFutureToBankByFutureManualType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromFutureToBankByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起查询银行余额通知
func (t *trade) regOnRtnQueryBankBalanceByFuture(on OnRtnQueryBankBalanceByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnQueryBankBalanceByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起银行资金转期货错误回报
func (t *trade) regOnErrRtnBankToFutureByFuture(on OnErrRtnBankToFutureByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起期货资金转银行错误回报
func (t *trade) regOnErrRtnFutureToBankByFuture(on OnErrRtnFutureToBankByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正银行转期货错误回报
func (t *trade) regOnErrRtnRepealBankToFutureByFutureManual(on OnErrRtnRepealBankToFutureByFutureManualType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnRepealBankToFutureByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 系统运行时期货端手工发起冲正期货转银行错误回报
func (t *trade) regOnErrRtnRepealFutureToBankByFutureManual(on OnErrRtnRepealFutureToBankByFutureManualType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnRepealFutureToBankByFutureManual").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起查询银行余额错误回报
func (t *trade) regOnErrRtnQueryBankBalanceByFuture(on OnErrRtnQueryBankBalanceByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnErrRtnQueryBankBalanceByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromBankToFutureByFuture(on OnRtnRepealFromBankToFutureByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (t *trade) regOnRtnRepealFromFutureToBankByFuture(on OnRtnRepealFromFutureToBankByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnRepealFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起银行资金转期货应答
func (t *trade) regOnRspFromBankToFutureByFuture(on OnRspFromBankToFutureByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRspFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起期货资金转银行应答
func (t *trade) regOnRspFromFutureToBankByFuture(on OnRspFromFutureToBankByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRspFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 期货发起查询银行余额应答
func (t *trade) regOnRspQueryBankAccountMoneyByFuture(on OnRspQueryBankAccountMoneyByFutureType) {
	_, _, _ = t.h.MustFindProc("SetOnRspQueryBankAccountMoneyByFuture").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起银期开户通知
func (t *trade) regOnRtnOpenAccountByBank(on OnRtnOpenAccountByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnOpenAccountByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起银期销户通知
func (t *trade) regOnRtnCancelAccountByBank(on OnRtnCancelAccountByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnCancelAccountByBank").Call(t.spi, syscall.NewCallback(on))
}

// 银行发起变更银行账号通知
func (t *trade) regOnRtnChangeAccountByBank(on OnRtnChangeAccountByBankType) {
	_, _, _ = t.h.MustFindProc("SetOnRtnChangeAccountByBank").Call(t.spi, syscall.NewCallback(on))
}

// 创建TraderApi
func (t *trade) Release() {
	_, _, _ = t.h.MustFindProc("Release").Call(t.api)
}

// 初始化
func (t *trade) Init() {
	_, _, _ = t.h.MustFindProc("Init").Call(t.api)
}

// 等待接口线程结束运行
func (t *trade) Join() {
	_, _, _ = t.h.MustFindProc("Join").Call(t.api)
}

// 注册前置机网络地址
func (t *trade) RegisterFront(pszFrontAddress string) {
	bs, _ := syscall.BytePtrFromString(pszFrontAddress)
	_, _, _ = t.h.MustFindProc("RegisterFront").Call(t.api, uintptr(unsafe.Pointer(bs)))
}

// @remark RegisterNameServer优先于RegisterFront
func (t *trade) RegisterNameServer(pszNsAddress string) {
	bs, _ := syscall.BytePtrFromString(pszNsAddress)
	_, _, _ = t.h.MustFindProc("RegisterNameServer").Call(t.api, uintptr(unsafe.Pointer(bs)))
}

// 注册名字服务器用户信息
func (t *trade) RegisterFensUserInfo() {
	_, _, _ = t.h.MustFindProc("RegisterFensUserInfo").Call(t.api)
}

// 订阅私有流。
func (t *trade) SubscribePrivateTopic(nResumeType go_ctp.THOST_TE_RESUME_TYPE) {
	_, _, _ = t.h.MustFindProc("SubscribePrivateTopic").Call(t.api, uintptr(nResumeType))
}

// 订阅公共流。
func (t *trade) SubscribePublicTopic(nResumeType go_ctp.THOST_TE_RESUME_TYPE) {
	_, _, _ = t.h.MustFindProc("SubscribePublicTopic").Call(t.api, uintptr(nResumeType))
}

// 客户端认证请求
func (t *trade) ReqAuthenticate(pReqAuthenticateField go_ctp.CThostFtdcReqAuthenticateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqAuthenticate").Call(t.api, uintptr(unsafe.Pointer(&pReqAuthenticateField)), uintptr(t.nRequestID))
}

// 注册用户终端信息，用于中继服务器多连接模式
func (t *trade) RegisterUserSystemInfo(pUserSystemInfo go_ctp.CThostFtdcUserSystemInfoField) {
	_, _, _ = t.h.MustFindProc("RegisterUserSystemInfo").Call(t.api, uintptr(unsafe.Pointer(&pUserSystemInfo)))
}

// 上报用户终端信息，用于中继服务器操作员登录模式
func (t *trade) SubmitUserSystemInfo(pUserSystemInfo go_ctp.CThostFtdcUserSystemInfoField) {
	_, _, _ = t.h.MustFindProc("SubmitUserSystemInfo").Call(t.api, uintptr(unsafe.Pointer(&pUserSystemInfo)))
}

// 用户登录请求
func (t *trade) ReqUserLogin(pReqUserLoginField go_ctp.CThostFtdcReqUserLoginField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLogin").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginField)), uintptr(t.nRequestID))
}

// 登出请求
func (t *trade) ReqUserLogout(pUserLogout go_ctp.CThostFtdcUserLogoutField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLogout").Call(t.api, uintptr(unsafe.Pointer(&pUserLogout)), uintptr(t.nRequestID))
}

// 用户口令更新请求
func (t *trade) ReqUserPasswordUpdate(pUserPasswordUpdate go_ctp.CThostFtdcUserPasswordUpdateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserPasswordUpdate").Call(t.api, uintptr(unsafe.Pointer(&pUserPasswordUpdate)), uintptr(t.nRequestID))
}

// 资金账户口令更新请求
func (t *trade) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate go_ctp.CThostFtdcTradingAccountPasswordUpdateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqTradingAccountPasswordUpdate").Call(t.api, uintptr(unsafe.Pointer(&pTradingAccountPasswordUpdate)), uintptr(t.nRequestID))
}

// 查询用户当前支持的认证模式
func (t *trade) ReqUserAuthMethod(pReqUserAuthMethod go_ctp.CThostFtdcReqUserAuthMethodField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserAuthMethod").Call(t.api, uintptr(unsafe.Pointer(&pReqUserAuthMethod)), uintptr(t.nRequestID))
}

// 用户发出获取图形验证码请求
func (t *trade) ReqGenUserCaptcha(pReqGenUserCaptcha go_ctp.CThostFtdcReqGenUserCaptchaField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqGenUserCaptcha").Call(t.api, uintptr(unsafe.Pointer(&pReqGenUserCaptcha)), uintptr(t.nRequestID))
}

// 用户发出获取短信验证码请求
func (t *trade) ReqGenUserText(pReqGenUserText go_ctp.CThostFtdcReqGenUserTextField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqGenUserText").Call(t.api, uintptr(unsafe.Pointer(&pReqGenUserText)), uintptr(t.nRequestID))
}

// 用户发出带有图片验证码的登陆请求
func (t *trade) ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha go_ctp.CThostFtdcReqUserLoginWithCaptchaField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLoginWithCaptcha").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginWithCaptcha)), uintptr(t.nRequestID))
}

// 用户发出带有短信验证码的登陆请求
func (t *trade) ReqUserLoginWithText(pReqUserLoginWithText go_ctp.CThostFtdcReqUserLoginWithTextField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLoginWithText").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginWithText)), uintptr(t.nRequestID))
}

// 用户发出带有动态口令的登陆请求
func (t *trade) ReqUserLoginWithOTP(pReqUserLoginWithOTP go_ctp.CThostFtdcReqUserLoginWithOTPField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqUserLoginWithOTP").Call(t.api, uintptr(unsafe.Pointer(&pReqUserLoginWithOTP)), uintptr(t.nRequestID))
}

// 报单录入请求
func (t *trade) ReqOrderInsert(pInputOrder go_ctp.CThostFtdcInputOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputOrder)), uintptr(t.nRequestID))
}

// 预埋单录入请求
func (t *trade) ReqParkedOrderInsert(pParkedOrder go_ctp.CThostFtdcParkedOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqParkedOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&pParkedOrder)), uintptr(t.nRequestID))
}

// 预埋撤单录入请求
func (t *trade) ReqParkedOrderAction(pParkedOrderAction go_ctp.CThostFtdcParkedOrderActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqParkedOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pParkedOrderAction)), uintptr(t.nRequestID))
}

// 报单操作请求
func (t *trade) ReqOrderAction(pInputOrderAction go_ctp.CThostFtdcInputOrderActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pInputOrderAction)), uintptr(t.nRequestID))
}

// 查询最大报单数量请求
func (t *trade) ReqQueryMaxOrderVolume(pQueryMaxOrderVolume go_ctp.CThostFtdcQueryMaxOrderVolumeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQueryMaxOrderVolume").Call(t.api, uintptr(unsafe.Pointer(&pQueryMaxOrderVolume)), uintptr(t.nRequestID))
}

// 投资者结算结果确认
func (t *trade) ReqSettlementInfoConfirm(pSettlementInfoConfirm go_ctp.CThostFtdcSettlementInfoConfirmField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqSettlementInfoConfirm").Call(t.api, uintptr(unsafe.Pointer(&pSettlementInfoConfirm)), uintptr(t.nRequestID))
}

// 请求删除预埋单
func (t *trade) ReqRemoveParkedOrder(pRemoveParkedOrder go_ctp.CThostFtdcRemoveParkedOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqRemoveParkedOrder").Call(t.api, uintptr(unsafe.Pointer(&pRemoveParkedOrder)), uintptr(t.nRequestID))
}

// 请求删除预埋撤单
func (t *trade) ReqRemoveParkedOrderAction(pRemoveParkedOrderAction go_ctp.CThostFtdcRemoveParkedOrderActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqRemoveParkedOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pRemoveParkedOrderAction)), uintptr(t.nRequestID))
}

// 执行宣告录入请求
func (t *trade) ReqExecOrderInsert(pInputExecOrder go_ctp.CThostFtdcInputExecOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqExecOrderInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputExecOrder)), uintptr(t.nRequestID))
}

// 执行宣告操作请求
func (t *trade) ReqExecOrderAction(pInputExecOrderAction go_ctp.CThostFtdcInputExecOrderActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqExecOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pInputExecOrderAction)), uintptr(t.nRequestID))
}

// 询价录入请求
func (t *trade) ReqForQuoteInsert(pInputForQuote go_ctp.CThostFtdcInputForQuoteField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqForQuoteInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputForQuote)), uintptr(t.nRequestID))
}

// 报价录入请求
func (t *trade) ReqQuoteInsert(pInputQuote go_ctp.CThostFtdcInputQuoteField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQuoteInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputQuote)), uintptr(t.nRequestID))
}

// 报价操作请求
func (t *trade) ReqQuoteAction(pInputQuoteAction go_ctp.CThostFtdcInputQuoteActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQuoteAction").Call(t.api, uintptr(unsafe.Pointer(&pInputQuoteAction)), uintptr(t.nRequestID))
}

// 批量报单操作请求
func (t *trade) ReqBatchOrderAction(pInputBatchOrderAction go_ctp.CThostFtdcInputBatchOrderActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqBatchOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pInputBatchOrderAction)), uintptr(t.nRequestID))
}

// 期权自对冲录入请求
func (t *trade) ReqOptionSelfCloseInsert(pInputOptionSelfClose go_ctp.CThostFtdcInputOptionSelfCloseField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOptionSelfCloseInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputOptionSelfClose)), uintptr(t.nRequestID))
}

// 期权自对冲操作请求
func (t *trade) ReqOptionSelfCloseAction(pInputOptionSelfCloseAction go_ctp.CThostFtdcInputOptionSelfCloseActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqOptionSelfCloseAction").Call(t.api, uintptr(unsafe.Pointer(&pInputOptionSelfCloseAction)), uintptr(t.nRequestID))
}

// 申请组合录入请求
func (t *trade) ReqCombActionInsert(pInputCombAction go_ctp.CThostFtdcInputCombActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqCombActionInsert").Call(t.api, uintptr(unsafe.Pointer(&pInputCombAction)), uintptr(t.nRequestID))
}

// 请求查询报单
func (t *trade) ReqQryOrder(pQryOrder go_ctp.CThostFtdcQryOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOrder").Call(t.api, uintptr(unsafe.Pointer(&pQryOrder)), uintptr(t.nRequestID))
}

// 请求查询成交
func (t *trade) ReqQryTrade(pQryTrade go_ctp.CThostFtdcQryTradeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTrade").Call(t.api, uintptr(unsafe.Pointer(&pQryTrade)), uintptr(t.nRequestID))
}

// 请求查询投资者持仓
func (t *trade) ReqQryInvestorPosition(pQryInvestorPosition go_ctp.CThostFtdcQryInvestorPositionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorPosition").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorPosition)), uintptr(t.nRequestID))
}

// 请求查询资金账户
func (t *trade) ReqQryTradingAccount(pQryTradingAccount go_ctp.CThostFtdcQryTradingAccountField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTradingAccount").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingAccount)), uintptr(t.nRequestID))
}

// 请求查询投资者
func (t *trade) ReqQryInvestor(pQryInvestor go_ctp.CThostFtdcQryInvestorField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestor").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestor)), uintptr(t.nRequestID))
}

// 请求查询交易编码
func (t *trade) ReqQryTradingCode(pQryTradingCode go_ctp.CThostFtdcQryTradingCodeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTradingCode").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingCode)), uintptr(t.nRequestID))
}

// 请求查询合约保证金率
func (t *trade) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate go_ctp.CThostFtdcQryInstrumentMarginRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrumentMarginRate").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrumentMarginRate)), uintptr(t.nRequestID))
}

// 请求查询合约手续费率
func (t *trade) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate go_ctp.CThostFtdcQryInstrumentCommissionRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrumentCommissionRate").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrumentCommissionRate)), uintptr(t.nRequestID))
}

// 请求查询交易所
func (t *trade) ReqQryExchange(pQryExchange go_ctp.CThostFtdcQryExchangeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchange").Call(t.api, uintptr(unsafe.Pointer(&pQryExchange)), uintptr(t.nRequestID))
}

// 请求查询产品
func (t *trade) ReqQryProduct(pQryProduct go_ctp.CThostFtdcQryProductField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryProduct").Call(t.api, uintptr(unsafe.Pointer(&pQryProduct)), uintptr(t.nRequestID))
}

// 请求查询合约
func (t *trade) ReqQryInstrument(pQryInstrument go_ctp.CThostFtdcQryInstrumentField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrument").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrument)), uintptr(t.nRequestID))
}

// 请求查询行情
func (t *trade) ReqQryDepthMarketData(pQryDepthMarketData go_ctp.CThostFtdcQryDepthMarketDataField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryDepthMarketData").Call(t.api, uintptr(unsafe.Pointer(&pQryDepthMarketData)), uintptr(t.nRequestID))
}

// 请求查询投资者结算结果
func (t *trade) ReqQrySettlementInfo(pQrySettlementInfo go_ctp.CThostFtdcQrySettlementInfoField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySettlementInfo").Call(t.api, uintptr(unsafe.Pointer(&pQrySettlementInfo)), uintptr(t.nRequestID))
}

// 请求查询转帐银行
func (t *trade) ReqQryTransferBank(pQryTransferBank go_ctp.CThostFtdcQryTransferBankField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTransferBank").Call(t.api, uintptr(unsafe.Pointer(&pQryTransferBank)), uintptr(t.nRequestID))
}

// 请求查询投资者持仓明细
func (t *trade) ReqQryInvestorPositionDetail(pQryInvestorPositionDetail go_ctp.CThostFtdcQryInvestorPositionDetailField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorPositionDetail").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorPositionDetail)), uintptr(t.nRequestID))
}

// 请求查询客户通知
func (t *trade) ReqQryNotice(pQryNotice go_ctp.CThostFtdcQryNoticeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryNotice").Call(t.api, uintptr(unsafe.Pointer(&pQryNotice)), uintptr(t.nRequestID))
}

// 请求查询结算信息确认
func (t *trade) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm go_ctp.CThostFtdcQrySettlementInfoConfirmField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySettlementInfoConfirm").Call(t.api, uintptr(unsafe.Pointer(&pQrySettlementInfoConfirm)), uintptr(t.nRequestID))
}

// 请求查询投资者持仓明细
func (t *trade) ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail go_ctp.CThostFtdcQryInvestorPositionCombineDetailField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorPositionCombineDetail").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorPositionCombineDetail)), uintptr(t.nRequestID))
}

// 请求查询保证金监管系统经纪公司资金账户密钥
func (t *trade) ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey go_ctp.CThostFtdcQryCFMMCTradingAccountKeyField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryCFMMCTradingAccountKey").Call(t.api, uintptr(unsafe.Pointer(&pQryCFMMCTradingAccountKey)), uintptr(t.nRequestID))
}

// 请求查询仓单折抵信息
func (t *trade) ReqQryEWarrantOffset(pQryEWarrantOffset go_ctp.CThostFtdcQryEWarrantOffsetField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryEWarrantOffset").Call(t.api, uintptr(unsafe.Pointer(&pQryEWarrantOffset)), uintptr(t.nRequestID))
}

// 请求查询投资者品种/跨品种保证金
func (t *trade) ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin go_ctp.CThostFtdcQryInvestorProductGroupMarginField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestorProductGroupMargin").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestorProductGroupMargin)), uintptr(t.nRequestID))
}

// 请求查询交易所保证金率
func (t *trade) ReqQryExchangeMarginRate(pQryExchangeMarginRate go_ctp.CThostFtdcQryExchangeMarginRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchangeMarginRate").Call(t.api, uintptr(unsafe.Pointer(&pQryExchangeMarginRate)), uintptr(t.nRequestID))
}

// 请求查询交易所调整保证金率
func (t *trade) ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust go_ctp.CThostFtdcQryExchangeMarginRateAdjustField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchangeMarginRateAdjust").Call(t.api, uintptr(unsafe.Pointer(&pQryExchangeMarginRateAdjust)), uintptr(t.nRequestID))
}

// 请求查询汇率
func (t *trade) ReqQryExchangeRate(pQryExchangeRate go_ctp.CThostFtdcQryExchangeRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExchangeRate").Call(t.api, uintptr(unsafe.Pointer(&pQryExchangeRate)), uintptr(t.nRequestID))
}

// 请求查询二级代理操作员银期权限
func (t *trade) ReqQrySecAgentACIDMap(pQrySecAgentACIDMap go_ctp.CThostFtdcQrySecAgentACIDMapField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentACIDMap").Call(t.api, uintptr(unsafe.Pointer(&pQrySecAgentACIDMap)), uintptr(t.nRequestID))
}

// 请求查询产品报价汇率
func (t *trade) ReqQryProductExchRate(pQryProductExchRate go_ctp.CThostFtdcQryProductExchRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryProductExchRate").Call(t.api, uintptr(unsafe.Pointer(&pQryProductExchRate)), uintptr(t.nRequestID))
}

// 请求查询产品组
func (t *trade) ReqQryProductGroup(pQryProductGroup go_ctp.CThostFtdcQryProductGroupField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryProductGroup").Call(t.api, uintptr(unsafe.Pointer(&pQryProductGroup)), uintptr(t.nRequestID))
}

// 请求查询做市商合约手续费率
func (t *trade) ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate go_ctp.CThostFtdcQryMMInstrumentCommissionRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryMMInstrumentCommissionRate").Call(t.api, uintptr(unsafe.Pointer(&pQryMMInstrumentCommissionRate)), uintptr(t.nRequestID))
}

// 请求查询做市商期权合约手续费
func (t *trade) ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate go_ctp.CThostFtdcQryMMOptionInstrCommRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryMMOptionInstrCommRate").Call(t.api, uintptr(unsafe.Pointer(&pQryMMOptionInstrCommRate)), uintptr(t.nRequestID))
}

// 请求查询报单手续费
func (t *trade) ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate go_ctp.CThostFtdcQryInstrumentOrderCommRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInstrumentOrderCommRate").Call(t.api, uintptr(unsafe.Pointer(&pQryInstrumentOrderCommRate)), uintptr(t.nRequestID))
}

// 请求查询资金账户
func (t *trade) ReqQrySecAgentTradingAccount(pQryTradingAccount go_ctp.CThostFtdcQryTradingAccountField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentTradingAccount").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingAccount)), uintptr(t.nRequestID))
}

// 请求查询二级代理商资金校验模式
func (t *trade) ReqQrySecAgentCheckMode(pQrySecAgentCheckMode go_ctp.CThostFtdcQrySecAgentCheckModeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentCheckMode").Call(t.api, uintptr(unsafe.Pointer(&pQrySecAgentCheckMode)), uintptr(t.nRequestID))
}

// 请求查询二级代理商信息
func (t *trade) ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo go_ctp.CThostFtdcQrySecAgentTradeInfoField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQrySecAgentTradeInfo").Call(t.api, uintptr(unsafe.Pointer(&pQrySecAgentTradeInfo)), uintptr(t.nRequestID))
}

// 请求查询期权交易成本
func (t *trade) ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost go_ctp.CThostFtdcQryOptionInstrTradeCostField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOptionInstrTradeCost").Call(t.api, uintptr(unsafe.Pointer(&pQryOptionInstrTradeCost)), uintptr(t.nRequestID))
}

// 请求查询期权合约手续费
func (t *trade) ReqQryOptionInstrCommRate(pQryOptionInstrCommRate go_ctp.CThostFtdcQryOptionInstrCommRateField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOptionInstrCommRate").Call(t.api, uintptr(unsafe.Pointer(&pQryOptionInstrCommRate)), uintptr(t.nRequestID))
}

// 请求查询执行宣告
func (t *trade) ReqQryExecOrder(pQryExecOrder go_ctp.CThostFtdcQryExecOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryExecOrder").Call(t.api, uintptr(unsafe.Pointer(&pQryExecOrder)), uintptr(t.nRequestID))
}

// 请求查询询价
func (t *trade) ReqQryForQuote(pQryForQuote go_ctp.CThostFtdcQryForQuoteField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryForQuote").Call(t.api, uintptr(unsafe.Pointer(&pQryForQuote)), uintptr(t.nRequestID))
}

// 请求查询报价
func (t *trade) ReqQryQuote(pQryQuote go_ctp.CThostFtdcQryQuoteField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryQuote").Call(t.api, uintptr(unsafe.Pointer(&pQryQuote)), uintptr(t.nRequestID))
}

// 请求查询期权自对冲
func (t *trade) ReqQryOptionSelfClose(pQryOptionSelfClose go_ctp.CThostFtdcQryOptionSelfCloseField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryOptionSelfClose").Call(t.api, uintptr(unsafe.Pointer(&pQryOptionSelfClose)), uintptr(t.nRequestID))
}

// 请求查询投资单元
func (t *trade) ReqQryInvestUnit(pQryInvestUnit go_ctp.CThostFtdcQryInvestUnitField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryInvestUnit").Call(t.api, uintptr(unsafe.Pointer(&pQryInvestUnit)), uintptr(t.nRequestID))
}

// 请求查询组合合约安全系数
func (t *trade) ReqQryCombInstrumentGuard(pQryCombInstrumentGuard go_ctp.CThostFtdcQryCombInstrumentGuardField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryCombInstrumentGuard").Call(t.api, uintptr(unsafe.Pointer(&pQryCombInstrumentGuard)), uintptr(t.nRequestID))
}

// 请求查询申请组合
func (t *trade) ReqQryCombAction(pQryCombAction go_ctp.CThostFtdcQryCombActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryCombAction").Call(t.api, uintptr(unsafe.Pointer(&pQryCombAction)), uintptr(t.nRequestID))
}

// 请求查询转帐流水
func (t *trade) ReqQryTransferSerial(pQryTransferSerial go_ctp.CThostFtdcQryTransferSerialField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTransferSerial").Call(t.api, uintptr(unsafe.Pointer(&pQryTransferSerial)), uintptr(t.nRequestID))
}

// 请求查询银期签约关系
func (t *trade) ReqQryAccountregister(pQryAccountregister go_ctp.CThostFtdcQryAccountregisterField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryAccountregister").Call(t.api, uintptr(unsafe.Pointer(&pQryAccountregister)), uintptr(t.nRequestID))
}

// 请求查询签约银行
func (t *trade) ReqQryContractBank(pQryContractBank go_ctp.CThostFtdcQryContractBankField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryContractBank").Call(t.api, uintptr(unsafe.Pointer(&pQryContractBank)), uintptr(t.nRequestID))
}

// 请求查询预埋单
func (t *trade) ReqQryParkedOrder(pQryParkedOrder go_ctp.CThostFtdcQryParkedOrderField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryParkedOrder").Call(t.api, uintptr(unsafe.Pointer(&pQryParkedOrder)), uintptr(t.nRequestID))
}

// 请求查询预埋撤单
func (t *trade) ReqQryParkedOrderAction(pQryParkedOrderAction go_ctp.CThostFtdcQryParkedOrderActionField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryParkedOrderAction").Call(t.api, uintptr(unsafe.Pointer(&pQryParkedOrderAction)), uintptr(t.nRequestID))
}

// 请求查询交易通知
func (t *trade) ReqQryTradingNotice(pQryTradingNotice go_ctp.CThostFtdcQryTradingNoticeField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryTradingNotice").Call(t.api, uintptr(unsafe.Pointer(&pQryTradingNotice)), uintptr(t.nRequestID))
}

// 请求查询经纪公司交易参数
func (t *trade) ReqQryBrokerTradingParams(pQryBrokerTradingParams go_ctp.CThostFtdcQryBrokerTradingParamsField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryBrokerTradingParams").Call(t.api, uintptr(unsafe.Pointer(&pQryBrokerTradingParams)), uintptr(t.nRequestID))
}

// 请求查询经纪公司交易算法
func (t *trade) ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos go_ctp.CThostFtdcQryBrokerTradingAlgosField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQryBrokerTradingAlgos").Call(t.api, uintptr(unsafe.Pointer(&pQryBrokerTradingAlgos)), uintptr(t.nRequestID))
}

// 请求查询监控中心用户令牌
func (t *trade) ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken go_ctp.CThostFtdcQueryCFMMCTradingAccountTokenField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQueryCFMMCTradingAccountToken").Call(t.api, uintptr(unsafe.Pointer(&pQueryCFMMCTradingAccountToken)), uintptr(t.nRequestID))
}

// 期货发起银行资金转期货请求
func (t *trade) ReqFromBankToFutureByFuture(pReqTransfer go_ctp.CThostFtdcReqTransferField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqFromBankToFutureByFuture").Call(t.api, uintptr(unsafe.Pointer(&pReqTransfer)), uintptr(t.nRequestID))
}

// 期货发起期货资金转银行请求
func (t *trade) ReqFromFutureToBankByFuture(pReqTransfer go_ctp.CThostFtdcReqTransferField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqFromFutureToBankByFuture").Call(t.api, uintptr(unsafe.Pointer(&pReqTransfer)), uintptr(t.nRequestID))
}

// 期货发起查询银行余额请求
func (t *trade) ReqQueryBankAccountMoneyByFuture(pReqQueryAccount go_ctp.CThostFtdcReqQueryAccountField) {
	t.nRequestID++
	_, _, _ = t.h.MustFindProc("ReqQueryBankAccountMoneyByFuture").Call(t.api, uintptr(unsafe.Pointer(&pReqQueryAccount)), uintptr(t.nRequestID))
}
