#include "trade.h"
#include <string.h>
int nReq;

Trade::Trade(void)
{	
	_FrontConnected = NULL;
	_FrontDisconnected = NULL;
	_HeartBeatWarning = NULL;
	_RspAuthenticate = NULL;
	_RspUserLogin = NULL;
	_RspUserLogout = NULL;
	_RspUserPasswordUpdate = NULL;
	_RspTradingAccountPasswordUpdate = NULL;
	_RspUserAuthMethod = NULL;
	_RspGenUserCaptcha = NULL;
	_RspGenUserText = NULL;
	_RspOrderInsert = NULL;
	_RspParkedOrderInsert = NULL;
	_RspParkedOrderAction = NULL;
	_RspOrderAction = NULL;
	_RspQryMaxOrderVolume = NULL;
	_RspSettlementInfoConfirm = NULL;
	_RspRemoveParkedOrder = NULL;
	_RspRemoveParkedOrderAction = NULL;
	_RspExecOrderInsert = NULL;
	_RspExecOrderAction = NULL;
	_RspForQuoteInsert = NULL;
	_RspQuoteInsert = NULL;
	_RspQuoteAction = NULL;
	_RspBatchOrderAction = NULL;
	_RspOptionSelfCloseInsert = NULL;
	_RspOptionSelfCloseAction = NULL;
	_RspCombActionInsert = NULL;
	_RspQryOrder = NULL;
	_RspQryTrade = NULL;
	_RspQryInvestorPosition = NULL;
	_RspQryTradingAccount = NULL;
	_RspQryInvestor = NULL;
	_RspQryTradingCode = NULL;
	_RspQryInstrumentMarginRate = NULL;
	_RspQryInstrumentCommissionRate = NULL;
	_RspQryExchange = NULL;
	_RspQryProduct = NULL;
	_RspQryInstrument = NULL;
	_RspQryDepthMarketData = NULL;
	_RspQryTraderOffer = NULL;
	_RspQrySettlementInfo = NULL;
	_RspQryTransferBank = NULL;
	_RspQryInvestorPositionDetail = NULL;
	_RspQryNotice = NULL;
	_RspQrySettlementInfoConfirm = NULL;
	_RspQryInvestorPositionCombineDetail = NULL;
	_RspQryCFMMCTradingAccountKey = NULL;
	_RspQryEWarrantOffset = NULL;
	_RspQryInvestorProductGroupMargin = NULL;
	_RspQryExchangeMarginRate = NULL;
	_RspQryExchangeMarginRateAdjust = NULL;
	_RspQryExchangeRate = NULL;
	_RspQrySecAgentACIDMap = NULL;
	_RspQryProductExchRate = NULL;
	_RspQryProductGroup = NULL;
	_RspQryMMInstrumentCommissionRate = NULL;
	_RspQryMMOptionInstrCommRate = NULL;
	_RspQryInstrumentOrderCommRate = NULL;
	_RspQrySecAgentTradingAccount = NULL;
	_RspQrySecAgentCheckMode = NULL;
	_RspQrySecAgentTradeInfo = NULL;
	_RspQryOptionInstrTradeCost = NULL;
	_RspQryOptionInstrCommRate = NULL;
	_RspQryExecOrder = NULL;
	_RspQryForQuote = NULL;
	_RspQryQuote = NULL;
	_RspQryOptionSelfClose = NULL;
	_RspQryInvestUnit = NULL;
	_RspQryCombInstrumentGuard = NULL;
	_RspQryCombAction = NULL;
	_RspQryTransferSerial = NULL;
	_RspQryAccountregister = NULL;
	_RspError = NULL;
	_RtnOrder = NULL;
	_RtnTrade = NULL;
	_ErrRtnOrderInsert = NULL;
	_ErrRtnOrderAction = NULL;
	_RtnInstrumentStatus = NULL;
	_RtnBulletin = NULL;
	_RtnTradingNotice = NULL;
	_RtnErrorConditionalOrder = NULL;
	_RtnExecOrder = NULL;
	_ErrRtnExecOrderInsert = NULL;
	_ErrRtnExecOrderAction = NULL;
	_ErrRtnForQuoteInsert = NULL;
	_RtnQuote = NULL;
	_ErrRtnQuoteInsert = NULL;
	_ErrRtnQuoteAction = NULL;
	_RtnForQuoteRsp = NULL;
	_RtnCFMMCTradingAccountToken = NULL;
	_ErrRtnBatchOrderAction = NULL;
	_RtnOptionSelfClose = NULL;
	_ErrRtnOptionSelfCloseInsert = NULL;
	_ErrRtnOptionSelfCloseAction = NULL;
	_RtnCombAction = NULL;
	_ErrRtnCombActionInsert = NULL;
	_RspQryContractBank = NULL;
	_RspQryParkedOrder = NULL;
	_RspQryParkedOrderAction = NULL;
	_RspQryTradingNotice = NULL;
	_RspQryBrokerTradingParams = NULL;
	_RspQryBrokerTradingAlgos = NULL;
	_RspQueryCFMMCTradingAccountToken = NULL;
	_RtnFromBankToFutureByBank = NULL;
	_RtnFromFutureToBankByBank = NULL;
	_RtnRepealFromBankToFutureByBank = NULL;
	_RtnRepealFromFutureToBankByBank = NULL;
	_RtnFromBankToFutureByFuture = NULL;
	_RtnFromFutureToBankByFuture = NULL;
	_RtnRepealFromBankToFutureByFutureManual = NULL;
	_RtnRepealFromFutureToBankByFutureManual = NULL;
	_RtnQueryBankBalanceByFuture = NULL;
	_ErrRtnBankToFutureByFuture = NULL;
	_ErrRtnFutureToBankByFuture = NULL;
	_ErrRtnRepealBankToFutureByFutureManual = NULL;
	_ErrRtnRepealFutureToBankByFutureManual = NULL;
	_ErrRtnQueryBankBalanceByFuture = NULL;
	_RtnRepealFromBankToFutureByFuture = NULL;
	_RtnRepealFromFutureToBankByFuture = NULL;
	_RspFromBankToFutureByFuture = NULL;
	_RspFromFutureToBankByFuture = NULL;
	_RspQueryBankAccountMoneyByFuture = NULL;
	_RtnOpenAccountByBank = NULL;
	_RtnCancelAccountByBank = NULL;
	_RtnChangeAccountByBank = NULL;
	_RspQryClassifiedInstrument = NULL;
	_RspQryCombPromotionParam = NULL;
	_RspQryRiskSettleInvstPosition = NULL;
	_RspQryRiskSettleProductStatus = NULL;
	
}


// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
DLL_EXPORT_C_DECL void WINAPI tSetOnFrontConnected(Trade* spi, void* func){spi->_FrontConnected = func;}

// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
DLL_EXPORT_C_DECL void WINAPI tSetOnFrontDisconnected(Trade* spi, void* func){spi->_FrontDisconnected = func;}

// 心跳超时警告。当长时间未收到报文时，该方法被调用。
DLL_EXPORT_C_DECL void WINAPI tSetOnHeartBeatWarning(Trade* spi, void* func){spi->_HeartBeatWarning = func;}

// 客户端认证响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspAuthenticate(Trade* spi, void* func){spi->_RspAuthenticate = func;}

// 登录请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspUserLogin(Trade* spi, void* func){spi->_RspUserLogin = func;}

// 登出请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspUserLogout(Trade* spi, void* func){spi->_RspUserLogout = func;}

// 用户口令更新请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspUserPasswordUpdate(Trade* spi, void* func){spi->_RspUserPasswordUpdate = func;}

// 资金账户口令更新请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspTradingAccountPasswordUpdate(Trade* spi, void* func){spi->_RspTradingAccountPasswordUpdate = func;}

// 查询用户当前支持的认证模式的回复
DLL_EXPORT_C_DECL void WINAPI tSetOnRspUserAuthMethod(Trade* spi, void* func){spi->_RspUserAuthMethod = func;}

// 获取图形验证码请求的回复
DLL_EXPORT_C_DECL void WINAPI tSetOnRspGenUserCaptcha(Trade* spi, void* func){spi->_RspGenUserCaptcha = func;}

// 获取短信验证码请求的回复
DLL_EXPORT_C_DECL void WINAPI tSetOnRspGenUserText(Trade* spi, void* func){spi->_RspGenUserText = func;}

// 报单录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspOrderInsert(Trade* spi, void* func){spi->_RspOrderInsert = func;}

// 预埋单录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspParkedOrderInsert(Trade* spi, void* func){spi->_RspParkedOrderInsert = func;}

// 预埋撤单录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspParkedOrderAction(Trade* spi, void* func){spi->_RspParkedOrderAction = func;}

// 报单操作请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspOrderAction(Trade* spi, void* func){spi->_RspOrderAction = func;}

// 查询最大报单数量响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryMaxOrderVolume(Trade* spi, void* func){spi->_RspQryMaxOrderVolume = func;}

// 投资者结算结果确认响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspSettlementInfoConfirm(Trade* spi, void* func){spi->_RspSettlementInfoConfirm = func;}

// 删除预埋单响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspRemoveParkedOrder(Trade* spi, void* func){spi->_RspRemoveParkedOrder = func;}

// 删除预埋撤单响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspRemoveParkedOrderAction(Trade* spi, void* func){spi->_RspRemoveParkedOrderAction = func;}

// 执行宣告录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspExecOrderInsert(Trade* spi, void* func){spi->_RspExecOrderInsert = func;}

// 执行宣告操作请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspExecOrderAction(Trade* spi, void* func){spi->_RspExecOrderAction = func;}

// 询价录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspForQuoteInsert(Trade* spi, void* func){spi->_RspForQuoteInsert = func;}

// 报价录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQuoteInsert(Trade* spi, void* func){spi->_RspQuoteInsert = func;}

// 报价操作请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQuoteAction(Trade* spi, void* func){spi->_RspQuoteAction = func;}

// 批量报单操作请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspBatchOrderAction(Trade* spi, void* func){spi->_RspBatchOrderAction = func;}

// 期权自对冲录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspOptionSelfCloseInsert(Trade* spi, void* func){spi->_RspOptionSelfCloseInsert = func;}

// 期权自对冲操作请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspOptionSelfCloseAction(Trade* spi, void* func){spi->_RspOptionSelfCloseAction = func;}

// 申请组合录入请求响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspCombActionInsert(Trade* spi, void* func){spi->_RspCombActionInsert = func;}

// 请求查询报单响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryOrder(Trade* spi, void* func){spi->_RspQryOrder = func;}

// 请求查询成交响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTrade(Trade* spi, void* func){spi->_RspQryTrade = func;}

// 请求查询投资者持仓响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInvestorPosition(Trade* spi, void* func){spi->_RspQryInvestorPosition = func;}

// 请求查询资金账户响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTradingAccount(Trade* spi, void* func){spi->_RspQryTradingAccount = func;}

// 请求查询投资者响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInvestor(Trade* spi, void* func){spi->_RspQryInvestor = func;}

// 请求查询交易编码响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTradingCode(Trade* spi, void* func){spi->_RspQryTradingCode = func;}

// 请求查询合约保证金率响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInstrumentMarginRate(Trade* spi, void* func){spi->_RspQryInstrumentMarginRate = func;}

// 请求查询合约手续费率响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInstrumentCommissionRate(Trade* spi, void* func){spi->_RspQryInstrumentCommissionRate = func;}

// 请求查询交易所响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryExchange(Trade* spi, void* func){spi->_RspQryExchange = func;}

// 请求查询产品响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryProduct(Trade* spi, void* func){spi->_RspQryProduct = func;}

// 请求查询合约响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInstrument(Trade* spi, void* func){spi->_RspQryInstrument = func;}

// 请求查询行情响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryDepthMarketData(Trade* spi, void* func){spi->_RspQryDepthMarketData = func;}

// 请求查询交易员报盘机响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTraderOffer(Trade* spi, void* func){spi->_RspQryTraderOffer = func;}

// 请求查询投资者结算结果响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQrySettlementInfo(Trade* spi, void* func){spi->_RspQrySettlementInfo = func;}

// 请求查询转帐银行响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTransferBank(Trade* spi, void* func){spi->_RspQryTransferBank = func;}

// 请求查询投资者持仓明细响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInvestorPositionDetail(Trade* spi, void* func){spi->_RspQryInvestorPositionDetail = func;}

// 请求查询客户通知响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryNotice(Trade* spi, void* func){spi->_RspQryNotice = func;}

// 请求查询结算信息确认响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQrySettlementInfoConfirm(Trade* spi, void* func){spi->_RspQrySettlementInfoConfirm = func;}

// 请求查询投资者持仓明细响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInvestorPositionCombineDetail(Trade* spi, void* func){spi->_RspQryInvestorPositionCombineDetail = func;}

// 查询保证金监管系统经纪公司资金账户密钥响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryCFMMCTradingAccountKey(Trade* spi, void* func){spi->_RspQryCFMMCTradingAccountKey = func;}

// 请求查询仓单折抵信息响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryEWarrantOffset(Trade* spi, void* func){spi->_RspQryEWarrantOffset = func;}

// 请求查询投资者品种/跨品种保证金响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInvestorProductGroupMargin(Trade* spi, void* func){spi->_RspQryInvestorProductGroupMargin = func;}

// 请求查询交易所保证金率响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryExchangeMarginRate(Trade* spi, void* func){spi->_RspQryExchangeMarginRate = func;}

// 请求查询交易所调整保证金率响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryExchangeMarginRateAdjust(Trade* spi, void* func){spi->_RspQryExchangeMarginRateAdjust = func;}

// 请求查询汇率响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryExchangeRate(Trade* spi, void* func){spi->_RspQryExchangeRate = func;}

// 请求查询二级代理操作员银期权限响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQrySecAgentACIDMap(Trade* spi, void* func){spi->_RspQrySecAgentACIDMap = func;}

// 请求查询产品报价汇率
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryProductExchRate(Trade* spi, void* func){spi->_RspQryProductExchRate = func;}

// 请求查询产品组
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryProductGroup(Trade* spi, void* func){spi->_RspQryProductGroup = func;}

// 请求查询做市商合约手续费率响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryMMInstrumentCommissionRate(Trade* spi, void* func){spi->_RspQryMMInstrumentCommissionRate = func;}

// 请求查询做市商期权合约手续费响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryMMOptionInstrCommRate(Trade* spi, void* func){spi->_RspQryMMOptionInstrCommRate = func;}

// 请求查询报单手续费响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInstrumentOrderCommRate(Trade* spi, void* func){spi->_RspQryInstrumentOrderCommRate = func;}

// 请求查询资金账户响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQrySecAgentTradingAccount(Trade* spi, void* func){spi->_RspQrySecAgentTradingAccount = func;}

// 请求查询二级代理商资金校验模式响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQrySecAgentCheckMode(Trade* spi, void* func){spi->_RspQrySecAgentCheckMode = func;}

// 请求查询二级代理商信息响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQrySecAgentTradeInfo(Trade* spi, void* func){spi->_RspQrySecAgentTradeInfo = func;}

// 请求查询期权交易成本响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryOptionInstrTradeCost(Trade* spi, void* func){spi->_RspQryOptionInstrTradeCost = func;}

// 请求查询期权合约手续费响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryOptionInstrCommRate(Trade* spi, void* func){spi->_RspQryOptionInstrCommRate = func;}

// 请求查询执行宣告响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryExecOrder(Trade* spi, void* func){spi->_RspQryExecOrder = func;}

// 请求查询询价响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryForQuote(Trade* spi, void* func){spi->_RspQryForQuote = func;}

// 请求查询报价响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryQuote(Trade* spi, void* func){spi->_RspQryQuote = func;}

// 请求查询期权自对冲响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryOptionSelfClose(Trade* spi, void* func){spi->_RspQryOptionSelfClose = func;}

// 请求查询投资单元响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryInvestUnit(Trade* spi, void* func){spi->_RspQryInvestUnit = func;}

// 请求查询组合合约安全系数响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryCombInstrumentGuard(Trade* spi, void* func){spi->_RspQryCombInstrumentGuard = func;}

// 请求查询申请组合响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryCombAction(Trade* spi, void* func){spi->_RspQryCombAction = func;}

// 请求查询转帐流水响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTransferSerial(Trade* spi, void* func){spi->_RspQryTransferSerial = func;}

// 请求查询银期签约关系响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryAccountregister(Trade* spi, void* func){spi->_RspQryAccountregister = func;}

// 错误应答
DLL_EXPORT_C_DECL void WINAPI tSetOnRspError(Trade* spi, void* func){spi->_RspError = func;}

// 报单通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnOrder(Trade* spi, void* func){spi->_RtnOrder = func;}

// 成交通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnTrade(Trade* spi, void* func){spi->_RtnTrade = func;}

// 报单录入错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnOrderInsert(Trade* spi, void* func){spi->_ErrRtnOrderInsert = func;}

// 报单操作错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnOrderAction(Trade* spi, void* func){spi->_ErrRtnOrderAction = func;}

// 合约交易状态通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnInstrumentStatus(Trade* spi, void* func){spi->_RtnInstrumentStatus = func;}

// 交易所公告通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnBulletin(Trade* spi, void* func){spi->_RtnBulletin = func;}

// 交易通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnTradingNotice(Trade* spi, void* func){spi->_RtnTradingNotice = func;}

// 提示条件单校验错误
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnErrorConditionalOrder(Trade* spi, void* func){spi->_RtnErrorConditionalOrder = func;}

// 执行宣告通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnExecOrder(Trade* spi, void* func){spi->_RtnExecOrder = func;}

// 执行宣告录入错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnExecOrderInsert(Trade* spi, void* func){spi->_ErrRtnExecOrderInsert = func;}

// 执行宣告操作错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnExecOrderAction(Trade* spi, void* func){spi->_ErrRtnExecOrderAction = func;}

// 询价录入错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnForQuoteInsert(Trade* spi, void* func){spi->_ErrRtnForQuoteInsert = func;}

// 报价通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnQuote(Trade* spi, void* func){spi->_RtnQuote = func;}

// 报价录入错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnQuoteInsert(Trade* spi, void* func){spi->_ErrRtnQuoteInsert = func;}

// 报价操作错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnQuoteAction(Trade* spi, void* func){spi->_ErrRtnQuoteAction = func;}

// 询价通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnForQuoteRsp(Trade* spi, void* func){spi->_RtnForQuoteRsp = func;}

// 保证金监控中心用户令牌
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnCFMMCTradingAccountToken(Trade* spi, void* func){spi->_RtnCFMMCTradingAccountToken = func;}

// 批量报单操作错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnBatchOrderAction(Trade* spi, void* func){spi->_ErrRtnBatchOrderAction = func;}

// 期权自对冲通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnOptionSelfClose(Trade* spi, void* func){spi->_RtnOptionSelfClose = func;}

// 期权自对冲录入错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnOptionSelfCloseInsert(Trade* spi, void* func){spi->_ErrRtnOptionSelfCloseInsert = func;}

// 期权自对冲操作错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnOptionSelfCloseAction(Trade* spi, void* func){spi->_ErrRtnOptionSelfCloseAction = func;}

// 申请组合通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnCombAction(Trade* spi, void* func){spi->_RtnCombAction = func;}

// 申请组合录入错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnCombActionInsert(Trade* spi, void* func){spi->_ErrRtnCombActionInsert = func;}

// 请求查询签约银行响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryContractBank(Trade* spi, void* func){spi->_RspQryContractBank = func;}

// 请求查询预埋单响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryParkedOrder(Trade* spi, void* func){spi->_RspQryParkedOrder = func;}

// 请求查询预埋撤单响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryParkedOrderAction(Trade* spi, void* func){spi->_RspQryParkedOrderAction = func;}

// 请求查询交易通知响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryTradingNotice(Trade* spi, void* func){spi->_RspQryTradingNotice = func;}

// 请求查询经纪公司交易参数响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryBrokerTradingParams(Trade* spi, void* func){spi->_RspQryBrokerTradingParams = func;}

// 请求查询经纪公司交易算法响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryBrokerTradingAlgos(Trade* spi, void* func){spi->_RspQryBrokerTradingAlgos = func;}

// 请求查询监控中心用户令牌
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQueryCFMMCTradingAccountToken(Trade* spi, void* func){spi->_RspQueryCFMMCTradingAccountToken = func;}

// 银行发起银行资金转期货通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnFromBankToFutureByBank(Trade* spi, void* func){spi->_RtnFromBankToFutureByBank = func;}

// 银行发起期货资金转银行通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnFromFutureToBankByBank(Trade* spi, void* func){spi->_RtnFromFutureToBankByBank = func;}

// 银行发起冲正银行转期货通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnRepealFromBankToFutureByBank(Trade* spi, void* func){spi->_RtnRepealFromBankToFutureByBank = func;}

// 银行发起冲正期货转银行通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnRepealFromFutureToBankByBank(Trade* spi, void* func){spi->_RtnRepealFromFutureToBankByBank = func;}

// 期货发起银行资金转期货通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnFromBankToFutureByFuture(Trade* spi, void* func){spi->_RtnFromBankToFutureByFuture = func;}

// 期货发起期货资金转银行通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnFromFutureToBankByFuture(Trade* spi, void* func){spi->_RtnFromFutureToBankByFuture = func;}

// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnRepealFromBankToFutureByFutureManual(Trade* spi, void* func){spi->_RtnRepealFromBankToFutureByFutureManual = func;}

// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnRepealFromFutureToBankByFutureManual(Trade* spi, void* func){spi->_RtnRepealFromFutureToBankByFutureManual = func;}

// 期货发起查询银行余额通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnQueryBankBalanceByFuture(Trade* spi, void* func){spi->_RtnQueryBankBalanceByFuture = func;}

// 期货发起银行资金转期货错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnBankToFutureByFuture(Trade* spi, void* func){spi->_ErrRtnBankToFutureByFuture = func;}

// 期货发起期货资金转银行错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnFutureToBankByFuture(Trade* spi, void* func){spi->_ErrRtnFutureToBankByFuture = func;}

// 系统运行时期货端手工发起冲正银行转期货错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnRepealBankToFutureByFutureManual(Trade* spi, void* func){spi->_ErrRtnRepealBankToFutureByFutureManual = func;}

// 系统运行时期货端手工发起冲正期货转银行错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnRepealFutureToBankByFutureManual(Trade* spi, void* func){spi->_ErrRtnRepealFutureToBankByFutureManual = func;}

// 期货发起查询银行余额错误回报
DLL_EXPORT_C_DECL void WINAPI tSetOnErrRtnQueryBankBalanceByFuture(Trade* spi, void* func){spi->_ErrRtnQueryBankBalanceByFuture = func;}

// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnRepealFromBankToFutureByFuture(Trade* spi, void* func){spi->_RtnRepealFromBankToFutureByFuture = func;}

// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnRepealFromFutureToBankByFuture(Trade* spi, void* func){spi->_RtnRepealFromFutureToBankByFuture = func;}

// 期货发起银行资金转期货应答
DLL_EXPORT_C_DECL void WINAPI tSetOnRspFromBankToFutureByFuture(Trade* spi, void* func){spi->_RspFromBankToFutureByFuture = func;}

// 期货发起期货资金转银行应答
DLL_EXPORT_C_DECL void WINAPI tSetOnRspFromFutureToBankByFuture(Trade* spi, void* func){spi->_RspFromFutureToBankByFuture = func;}

// 期货发起查询银行余额应答
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQueryBankAccountMoneyByFuture(Trade* spi, void* func){spi->_RspQueryBankAccountMoneyByFuture = func;}

// 银行发起银期开户通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnOpenAccountByBank(Trade* spi, void* func){spi->_RtnOpenAccountByBank = func;}

// 银行发起银期销户通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnCancelAccountByBank(Trade* spi, void* func){spi->_RtnCancelAccountByBank = func;}

// 银行发起变更银行账号通知
DLL_EXPORT_C_DECL void WINAPI tSetOnRtnChangeAccountByBank(Trade* spi, void* func){spi->_RtnChangeAccountByBank = func;}

// 请求查询分类合约响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryClassifiedInstrument(Trade* spi, void* func){spi->_RspQryClassifiedInstrument = func;}

// 请求组合优惠比例响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryCombPromotionParam(Trade* spi, void* func){spi->_RspQryCombPromotionParam = func;}

// 投资者风险结算持仓查询响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryRiskSettleInvstPosition(Trade* spi, void* func){spi->_RspQryRiskSettleInvstPosition = func;}

// 风险结算产品查询响应
DLL_EXPORT_C_DECL void WINAPI tSetOnRspQryRiskSettleProductStatus(Trade* spi, void* func){spi->_RspQryRiskSettleProductStatus = func;}

DLL_EXPORT_C_DECL void* WINAPI tCreateApi(){return CThostFtdcTraderApi::CreateFtdcTraderApi("./log/");}
DLL_EXPORT_C_DECL void* WINAPI tCreateSpi(){return new Trade();}
DLL_EXPORT_C_DECL void* WINAPI tGetVersion() { return (void*)CThostFtdcTraderApi::GetApiVersion(); }


// 创建TraderApi
DLL_EXPORT_C_DECL void* WINAPI tRelease(CThostFtdcTraderApi *api){api->Release(); return 0;}

// 初始化
DLL_EXPORT_C_DECL void* WINAPI tInit(CThostFtdcTraderApi *api){api->Init(); return 0;}

// 等待接口线程结束运行
DLL_EXPORT_C_DECL void* WINAPI tJoin(CThostFtdcTraderApi *api){api->Join(); return 0;}

// 注册前置机网络地址
DLL_EXPORT_C_DECL void* WINAPI tRegisterFront(CThostFtdcTraderApi *api, char *pszFrontAddress){api->RegisterFront(pszFrontAddress); return 0;}

// @remark RegisterNameServer优先于RegisterFront
DLL_EXPORT_C_DECL void* WINAPI tRegisterNameServer(CThostFtdcTraderApi *api, char *pszNsAddress){api->RegisterNameServer(pszNsAddress); return 0;}

// 注册名字服务器用户信息
DLL_EXPORT_C_DECL void* WINAPI tRegisterFensUserInfo(CThostFtdcTraderApi *api, CThostFtdcFensUserInfoField * pFensUserInfo){api->RegisterFensUserInfo( pFensUserInfo); return 0;}

// 注册回调接口
DLL_EXPORT_C_DECL void* WINAPI tRegisterSpi(CThostFtdcTraderApi *api, CThostFtdcTraderSpi *pSpi){api->RegisterSpi(pSpi); return 0;}

// 订阅私有流。
DLL_EXPORT_C_DECL void* WINAPI tSubscribePrivateTopic(CThostFtdcTraderApi *api, THOST_TE_RESUME_TYPE nResumeType){api->SubscribePrivateTopic(nResumeType); return 0;}

// 订阅公共流。
DLL_EXPORT_C_DECL void* WINAPI tSubscribePublicTopic(CThostFtdcTraderApi *api, THOST_TE_RESUME_TYPE nResumeType){api->SubscribePublicTopic(nResumeType); return 0;}

// 客户端认证请求
DLL_EXPORT_C_DECL void* WINAPI tReqAuthenticate(CThostFtdcTraderApi *api, CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID){api->ReqAuthenticate(pReqAuthenticateField, nRequestID); return 0;}

// 注册用户终端信息，用于中继服务器多连接模式
DLL_EXPORT_C_DECL void* WINAPI tRegisterUserSystemInfo(CThostFtdcTraderApi *api, CThostFtdcUserSystemInfoField *pUserSystemInfo){api->RegisterUserSystemInfo(pUserSystemInfo); return 0;}

// 上报用户终端信息，用于中继服务器操作员登录模式
DLL_EXPORT_C_DECL void* WINAPI tSubmitUserSystemInfo(CThostFtdcTraderApi *api, CThostFtdcUserSystemInfoField *pUserSystemInfo){api->SubmitUserSystemInfo(pUserSystemInfo); return 0;}

// 用户登录请求
DLL_EXPORT_C_DECL void* WINAPI tReqUserLogin(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID){api->ReqUserLogin(pReqUserLoginField, nRequestID); return 0;}

// 登出请求
DLL_EXPORT_C_DECL void* WINAPI tReqUserLogout(CThostFtdcTraderApi *api, CThostFtdcUserLogoutField *pUserLogout, int nRequestID){api->ReqUserLogout(pUserLogout, nRequestID); return 0;}

// 用户口令更新请求
DLL_EXPORT_C_DECL void* WINAPI tReqUserPasswordUpdate(CThostFtdcTraderApi *api, CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID){api->ReqUserPasswordUpdate(pUserPasswordUpdate, nRequestID); return 0;}

// 资金账户口令更新请求
DLL_EXPORT_C_DECL void* WINAPI tReqTradingAccountPasswordUpdate(CThostFtdcTraderApi *api, CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID){api->ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate, nRequestID); return 0;}

// 查询用户当前支持的认证模式
DLL_EXPORT_C_DECL void* WINAPI tReqUserAuthMethod(CThostFtdcTraderApi *api, CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID){api->ReqUserAuthMethod(pReqUserAuthMethod, nRequestID); return 0;}

// 用户发出获取图形验证码请求
DLL_EXPORT_C_DECL void* WINAPI tReqGenUserCaptcha(CThostFtdcTraderApi *api, CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID){api->ReqGenUserCaptcha(pReqGenUserCaptcha, nRequestID); return 0;}

// 用户发出获取短信验证码请求
DLL_EXPORT_C_DECL void* WINAPI tReqGenUserText(CThostFtdcTraderApi *api, CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID){api->ReqGenUserText(pReqGenUserText, nRequestID); return 0;}

// 用户发出带有图片验证码的登陆请求
DLL_EXPORT_C_DECL void* WINAPI tReqUserLoginWithCaptcha(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID){api->ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha, nRequestID); return 0;}

// 用户发出带有短信验证码的登陆请求
DLL_EXPORT_C_DECL void* WINAPI tReqUserLoginWithText(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID){api->ReqUserLoginWithText(pReqUserLoginWithText, nRequestID); return 0;}

// 用户发出带有动态口令的登陆请求
DLL_EXPORT_C_DECL void* WINAPI tReqUserLoginWithOTP(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID){api->ReqUserLoginWithOTP(pReqUserLoginWithOTP, nRequestID); return 0;}

// 报单录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqOrderInsert(CThostFtdcTraderApi *api, CThostFtdcInputOrderField *pInputOrder, int nRequestID){api->ReqOrderInsert(pInputOrder, nRequestID); return 0;}

// 预埋单录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqParkedOrderInsert(CThostFtdcTraderApi *api, CThostFtdcParkedOrderField *pParkedOrder, int nRequestID){api->ReqParkedOrderInsert(pParkedOrder, nRequestID); return 0;}

// 预埋撤单录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID){api->ReqParkedOrderAction(pParkedOrderAction, nRequestID); return 0;}

// 报单操作请求
DLL_EXPORT_C_DECL void* WINAPI tReqOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID){api->ReqOrderAction(pInputOrderAction, nRequestID); return 0;}

// 查询最大报单数量请求
DLL_EXPORT_C_DECL void* WINAPI tReqQryMaxOrderVolume(CThostFtdcTraderApi *api, CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID){api->ReqQryMaxOrderVolume(pQryMaxOrderVolume, nRequestID); return 0;}

// 投资者结算结果确认
DLL_EXPORT_C_DECL void* WINAPI tReqSettlementInfoConfirm(CThostFtdcTraderApi *api, CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID){api->ReqSettlementInfoConfirm(pSettlementInfoConfirm, nRequestID); return 0;}

// 请求删除预埋单
DLL_EXPORT_C_DECL void* WINAPI tReqRemoveParkedOrder(CThostFtdcTraderApi *api, CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID){api->ReqRemoveParkedOrder(pRemoveParkedOrder, nRequestID); return 0;}

// 请求删除预埋撤单
DLL_EXPORT_C_DECL void* WINAPI tReqRemoveParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID){api->ReqRemoveParkedOrderAction(pRemoveParkedOrderAction, nRequestID); return 0;}

// 执行宣告录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqExecOrderInsert(CThostFtdcTraderApi *api, CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID){api->ReqExecOrderInsert(pInputExecOrder, nRequestID); return 0;}

// 执行宣告操作请求
DLL_EXPORT_C_DECL void* WINAPI tReqExecOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID){api->ReqExecOrderAction(pInputExecOrderAction, nRequestID); return 0;}

// 询价录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqForQuoteInsert(CThostFtdcTraderApi *api, CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID){api->ReqForQuoteInsert(pInputForQuote, nRequestID); return 0;}

// 报价录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqQuoteInsert(CThostFtdcTraderApi *api, CThostFtdcInputQuoteField *pInputQuote, int nRequestID){api->ReqQuoteInsert(pInputQuote, nRequestID); return 0;}

// 报价操作请求
DLL_EXPORT_C_DECL void* WINAPI tReqQuoteAction(CThostFtdcTraderApi *api, CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID){api->ReqQuoteAction(pInputQuoteAction, nRequestID); return 0;}

// 批量报单操作请求
DLL_EXPORT_C_DECL void* WINAPI tReqBatchOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID){api->ReqBatchOrderAction(pInputBatchOrderAction, nRequestID); return 0;}

// 期权自对冲录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqOptionSelfCloseInsert(CThostFtdcTraderApi *api, CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID){api->ReqOptionSelfCloseInsert(pInputOptionSelfClose, nRequestID); return 0;}

// 期权自对冲操作请求
DLL_EXPORT_C_DECL void* WINAPI tReqOptionSelfCloseAction(CThostFtdcTraderApi *api, CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID){api->ReqOptionSelfCloseAction(pInputOptionSelfCloseAction, nRequestID); return 0;}

// 申请组合录入请求
DLL_EXPORT_C_DECL void* WINAPI tReqCombActionInsert(CThostFtdcTraderApi *api, CThostFtdcInputCombActionField *pInputCombAction, int nRequestID){api->ReqCombActionInsert(pInputCombAction, nRequestID); return 0;}

// 请求查询报单
DLL_EXPORT_C_DECL void* WINAPI tReqQryOrder(CThostFtdcTraderApi *api, CThostFtdcQryOrderField *pQryOrder, int nRequestID){api->ReqQryOrder(pQryOrder, nRequestID); return 0;}

// 请求查询成交
DLL_EXPORT_C_DECL void* WINAPI tReqQryTrade(CThostFtdcTraderApi *api, CThostFtdcQryTradeField *pQryTrade, int nRequestID){api->ReqQryTrade(pQryTrade, nRequestID); return 0;}

// 请求查询投资者持仓
DLL_EXPORT_C_DECL void* WINAPI tReqQryInvestorPosition(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID){api->ReqQryInvestorPosition(pQryInvestorPosition, nRequestID); return 0;}

// 请求查询资金账户
DLL_EXPORT_C_DECL void* WINAPI tReqQryTradingAccount(CThostFtdcTraderApi *api, CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID){api->ReqQryTradingAccount(pQryTradingAccount, nRequestID); return 0;}

// 请求查询投资者
DLL_EXPORT_C_DECL void* WINAPI tReqQryInvestor(CThostFtdcTraderApi *api, CThostFtdcQryInvestorField *pQryInvestor, int nRequestID){api->ReqQryInvestor(pQryInvestor, nRequestID); return 0;}

// 请求查询交易编码
DLL_EXPORT_C_DECL void* WINAPI tReqQryTradingCode(CThostFtdcTraderApi *api, CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID){api->ReqQryTradingCode(pQryTradingCode, nRequestID); return 0;}

// 请求查询合约保证金率
DLL_EXPORT_C_DECL void* WINAPI tReqQryInstrumentMarginRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID){api->ReqQryInstrumentMarginRate(pQryInstrumentMarginRate, nRequestID); return 0;}

// 请求查询合约手续费率
DLL_EXPORT_C_DECL void* WINAPI tReqQryInstrumentCommissionRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID){api->ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate, nRequestID); return 0;}

// 请求查询交易所
DLL_EXPORT_C_DECL void* WINAPI tReqQryExchange(CThostFtdcTraderApi *api, CThostFtdcQryExchangeField *pQryExchange, int nRequestID){api->ReqQryExchange(pQryExchange, nRequestID); return 0;}

// 请求查询产品
DLL_EXPORT_C_DECL void* WINAPI tReqQryProduct(CThostFtdcTraderApi *api, CThostFtdcQryProductField *pQryProduct, int nRequestID){api->ReqQryProduct(pQryProduct, nRequestID); return 0;}

// 请求查询合约
DLL_EXPORT_C_DECL void* WINAPI tReqQryInstrument(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID){api->ReqQryInstrument(pQryInstrument, nRequestID); return 0;}

// 请求查询行情
DLL_EXPORT_C_DECL void* WINAPI tReqQryDepthMarketData(CThostFtdcTraderApi *api, CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID){api->ReqQryDepthMarketData(pQryDepthMarketData, nRequestID); return 0;}

// 请求查询交易员报盘机
DLL_EXPORT_C_DECL void* WINAPI tReqQryTraderOffer(CThostFtdcTraderApi *api, CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID){api->ReqQryTraderOffer(pQryTraderOffer, nRequestID); return 0;}

// 请求查询投资者结算结果
DLL_EXPORT_C_DECL void* WINAPI tReqQrySettlementInfo(CThostFtdcTraderApi *api, CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID){api->ReqQrySettlementInfo(pQrySettlementInfo, nRequestID); return 0;}

// 请求查询转帐银行
DLL_EXPORT_C_DECL void* WINAPI tReqQryTransferBank(CThostFtdcTraderApi *api, CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID){api->ReqQryTransferBank(pQryTransferBank, nRequestID); return 0;}

// 请求查询投资者持仓明细
DLL_EXPORT_C_DECL void* WINAPI tReqQryInvestorPositionDetail(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID){api->ReqQryInvestorPositionDetail(pQryInvestorPositionDetail, nRequestID); return 0;}

// 请求查询客户通知
DLL_EXPORT_C_DECL void* WINAPI tReqQryNotice(CThostFtdcTraderApi *api, CThostFtdcQryNoticeField *pQryNotice, int nRequestID){api->ReqQryNotice(pQryNotice, nRequestID); return 0;}

// 请求查询结算信息确认
DLL_EXPORT_C_DECL void* WINAPI tReqQrySettlementInfoConfirm(CThostFtdcTraderApi *api, CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID){api->ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm, nRequestID); return 0;}

// 请求查询投资者持仓明细
DLL_EXPORT_C_DECL void* WINAPI tReqQryInvestorPositionCombineDetail(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID){api->ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail, nRequestID); return 0;}

// 请求查询保证金监管系统经纪公司资金账户密钥
DLL_EXPORT_C_DECL void* WINAPI tReqQryCFMMCTradingAccountKey(CThostFtdcTraderApi *api, CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID){api->ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey, nRequestID); return 0;}

// 请求查询仓单折抵信息
DLL_EXPORT_C_DECL void* WINAPI tReqQryEWarrantOffset(CThostFtdcTraderApi *api, CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID){api->ReqQryEWarrantOffset(pQryEWarrantOffset, nRequestID); return 0;}

// 请求查询投资者品种/跨品种保证金
DLL_EXPORT_C_DECL void* WINAPI tReqQryInvestorProductGroupMargin(CThostFtdcTraderApi *api, CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID){api->ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin, nRequestID); return 0;}

// 请求查询交易所保证金率
DLL_EXPORT_C_DECL void* WINAPI tReqQryExchangeMarginRate(CThostFtdcTraderApi *api, CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID){api->ReqQryExchangeMarginRate(pQryExchangeMarginRate, nRequestID); return 0;}

// 请求查询交易所调整保证金率
DLL_EXPORT_C_DECL void* WINAPI tReqQryExchangeMarginRateAdjust(CThostFtdcTraderApi *api, CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID){api->ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust, nRequestID); return 0;}

// 请求查询汇率
DLL_EXPORT_C_DECL void* WINAPI tReqQryExchangeRate(CThostFtdcTraderApi *api, CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID){api->ReqQryExchangeRate(pQryExchangeRate, nRequestID); return 0;}

// 请求查询二级代理操作员银期权限
DLL_EXPORT_C_DECL void* WINAPI tReqQrySecAgentACIDMap(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID){api->ReqQrySecAgentACIDMap(pQrySecAgentACIDMap, nRequestID); return 0;}

// 请求查询产品报价汇率
DLL_EXPORT_C_DECL void* WINAPI tReqQryProductExchRate(CThostFtdcTraderApi *api, CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID){api->ReqQryProductExchRate(pQryProductExchRate, nRequestID); return 0;}

// 请求查询产品组
DLL_EXPORT_C_DECL void* WINAPI tReqQryProductGroup(CThostFtdcTraderApi *api, CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID){api->ReqQryProductGroup(pQryProductGroup, nRequestID); return 0;}

// 请求查询做市商合约手续费率
DLL_EXPORT_C_DECL void* WINAPI tReqQryMMInstrumentCommissionRate(CThostFtdcTraderApi *api, CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID){api->ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate, nRequestID); return 0;}

// 请求查询做市商期权合约手续费
DLL_EXPORT_C_DECL void* WINAPI tReqQryMMOptionInstrCommRate(CThostFtdcTraderApi *api, CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID){api->ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate, nRequestID); return 0;}

// 请求查询报单手续费
DLL_EXPORT_C_DECL void* WINAPI tReqQryInstrumentOrderCommRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID){api->ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate, nRequestID); return 0;}

// 请求查询资金账户
DLL_EXPORT_C_DECL void* WINAPI tReqQrySecAgentTradingAccount(CThostFtdcTraderApi *api, CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID){api->ReqQrySecAgentTradingAccount(pQryTradingAccount, nRequestID); return 0;}

// 请求查询二级代理商资金校验模式
DLL_EXPORT_C_DECL void* WINAPI tReqQrySecAgentCheckMode(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID){api->ReqQrySecAgentCheckMode(pQrySecAgentCheckMode, nRequestID); return 0;}

// 请求查询二级代理商信息
DLL_EXPORT_C_DECL void* WINAPI tReqQrySecAgentTradeInfo(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID){api->ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo, nRequestID); return 0;}

// 请求查询期权交易成本
DLL_EXPORT_C_DECL void* WINAPI tReqQryOptionInstrTradeCost(CThostFtdcTraderApi *api, CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID){api->ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost, nRequestID); return 0;}

// 请求查询期权合约手续费
DLL_EXPORT_C_DECL void* WINAPI tReqQryOptionInstrCommRate(CThostFtdcTraderApi *api, CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID){api->ReqQryOptionInstrCommRate(pQryOptionInstrCommRate, nRequestID); return 0;}

// 请求查询执行宣告
DLL_EXPORT_C_DECL void* WINAPI tReqQryExecOrder(CThostFtdcTraderApi *api, CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID){api->ReqQryExecOrder(pQryExecOrder, nRequestID); return 0;}

// 请求查询询价
DLL_EXPORT_C_DECL void* WINAPI tReqQryForQuote(CThostFtdcTraderApi *api, CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID){api->ReqQryForQuote(pQryForQuote, nRequestID); return 0;}

// 请求查询报价
DLL_EXPORT_C_DECL void* WINAPI tReqQryQuote(CThostFtdcTraderApi *api, CThostFtdcQryQuoteField *pQryQuote, int nRequestID){api->ReqQryQuote(pQryQuote, nRequestID); return 0;}

// 请求查询期权自对冲
DLL_EXPORT_C_DECL void* WINAPI tReqQryOptionSelfClose(CThostFtdcTraderApi *api, CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID){api->ReqQryOptionSelfClose(pQryOptionSelfClose, nRequestID); return 0;}

// 请求查询投资单元
DLL_EXPORT_C_DECL void* WINAPI tReqQryInvestUnit(CThostFtdcTraderApi *api, CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID){api->ReqQryInvestUnit(pQryInvestUnit, nRequestID); return 0;}

// 请求查询组合合约安全系数
DLL_EXPORT_C_DECL void* WINAPI tReqQryCombInstrumentGuard(CThostFtdcTraderApi *api, CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID){api->ReqQryCombInstrumentGuard(pQryCombInstrumentGuard, nRequestID); return 0;}

// 请求查询申请组合
DLL_EXPORT_C_DECL void* WINAPI tReqQryCombAction(CThostFtdcTraderApi *api, CThostFtdcQryCombActionField *pQryCombAction, int nRequestID){api->ReqQryCombAction(pQryCombAction, nRequestID); return 0;}

// 请求查询转帐流水
DLL_EXPORT_C_DECL void* WINAPI tReqQryTransferSerial(CThostFtdcTraderApi *api, CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID){api->ReqQryTransferSerial(pQryTransferSerial, nRequestID); return 0;}

// 请求查询银期签约关系
DLL_EXPORT_C_DECL void* WINAPI tReqQryAccountregister(CThostFtdcTraderApi *api, CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID){api->ReqQryAccountregister(pQryAccountregister, nRequestID); return 0;}

// 请求查询签约银行
DLL_EXPORT_C_DECL void* WINAPI tReqQryContractBank(CThostFtdcTraderApi *api, CThostFtdcQryContractBankField *pQryContractBank, int nRequestID){api->ReqQryContractBank(pQryContractBank, nRequestID); return 0;}

// 请求查询预埋单
DLL_EXPORT_C_DECL void* WINAPI tReqQryParkedOrder(CThostFtdcTraderApi *api, CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID){api->ReqQryParkedOrder(pQryParkedOrder, nRequestID); return 0;}

// 请求查询预埋撤单
DLL_EXPORT_C_DECL void* WINAPI tReqQryParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID){api->ReqQryParkedOrderAction(pQryParkedOrderAction, nRequestID); return 0;}

// 请求查询交易通知
DLL_EXPORT_C_DECL void* WINAPI tReqQryTradingNotice(CThostFtdcTraderApi *api, CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID){api->ReqQryTradingNotice(pQryTradingNotice, nRequestID); return 0;}

// 请求查询经纪公司交易参数
DLL_EXPORT_C_DECL void* WINAPI tReqQryBrokerTradingParams(CThostFtdcTraderApi *api, CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID){api->ReqQryBrokerTradingParams(pQryBrokerTradingParams, nRequestID); return 0;}

// 请求查询经纪公司交易算法
DLL_EXPORT_C_DECL void* WINAPI tReqQryBrokerTradingAlgos(CThostFtdcTraderApi *api, CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID){api->ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos, nRequestID); return 0;}

// 请求查询监控中心用户令牌
DLL_EXPORT_C_DECL void* WINAPI tReqQueryCFMMCTradingAccountToken(CThostFtdcTraderApi *api, CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID){api->ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken, nRequestID); return 0;}

// 期货发起银行资金转期货请求
DLL_EXPORT_C_DECL void* WINAPI tReqFromBankToFutureByFuture(CThostFtdcTraderApi *api, CThostFtdcReqTransferField *pReqTransfer, int nRequestID){api->ReqFromBankToFutureByFuture(pReqTransfer, nRequestID); return 0;}

// 期货发起期货资金转银行请求
DLL_EXPORT_C_DECL void* WINAPI tReqFromFutureToBankByFuture(CThostFtdcTraderApi *api, CThostFtdcReqTransferField *pReqTransfer, int nRequestID){api->ReqFromFutureToBankByFuture(pReqTransfer, nRequestID); return 0;}

// 期货发起查询银行余额请求
DLL_EXPORT_C_DECL void* WINAPI tReqQueryBankAccountMoneyByFuture(CThostFtdcTraderApi *api, CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID){api->ReqQueryBankAccountMoneyByFuture(pReqQueryAccount, nRequestID); return 0;}

// 请求查询分类合约
DLL_EXPORT_C_DECL void* WINAPI tReqQryClassifiedInstrument(CThostFtdcTraderApi *api, CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID){api->ReqQryClassifiedInstrument(pQryClassifiedInstrument, nRequestID); return 0;}

// 请求组合优惠比例
DLL_EXPORT_C_DECL void* WINAPI tReqQryCombPromotionParam(CThostFtdcTraderApi *api, CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID){api->ReqQryCombPromotionParam(pQryCombPromotionParam, nRequestID); return 0;}

// 投资者风险结算持仓查询
DLL_EXPORT_C_DECL void* WINAPI tReqQryRiskSettleInvstPosition(CThostFtdcTraderApi *api, CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID){api->ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition, nRequestID); return 0;}

// 风险结算产品查询
DLL_EXPORT_C_DECL void* WINAPI tReqQryRiskSettleProductStatus(CThostFtdcTraderApi *api, CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID){api->ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus, nRequestID); return 0;}

