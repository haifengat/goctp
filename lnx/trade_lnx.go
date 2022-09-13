package lnx

/*
#cgo CPPFLAGS: -fPIC -I../v6.6.8_20220712
#cgo linux LDFLAGS: -fPIC -L${SRCDIR} -Wl,-rpath ${SRCDIR} -lctp_trade -lstdc++

#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"
void* tCreateApi();
void* tCreateSpi();
void* tGetVersion();
// 创建TraderApi
void* tRelease(void* api);
// 初始化
void* tInit(void* api);
// 等待接口线程结束运行
void* tJoin(void* api);
// 注册前置机网络地址
void* tRegisterFront(void* api, char *pszFrontAddress);
// @remark RegisterNameServer优先于RegisterFront
void* tRegisterNameServer(void* api, char *pszNsAddress);
// 注册名字服务器用户信息
void* tRegisterFensUserInfo(void* api, struct CThostFtdcFensUserInfoField * pFensUserInfo);
// 注册回调接口
void* tRegisterSpi(void* api, void *pSpi);
// 订阅私有流。
void* tSubscribePrivateTopic(void* api, int nResumeType);
// 订阅公共流。
void* tSubscribePublicTopic(void* api, int nResumeType);
// 客户端认证请求
void* tReqAuthenticate(void* api, struct CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID);
// 注册用户终端信息，用于中继服务器多连接模式
void* tRegisterUserSystemInfo(void* api, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);
// 上报用户终端信息，用于中继服务器操作员登录模式
void* tSubmitUserSystemInfo(void* api, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);
// 用户登录请求
void* tReqUserLogin(void* api, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);
// 登出请求
void* tReqUserLogout(void* api, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);
// 用户口令更新请求
void* tReqUserPasswordUpdate(void* api, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID);
// 资金账户口令更新请求
void* tReqTradingAccountPasswordUpdate(void* api, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID);
// 查询用户当前支持的认证模式
void* tReqUserAuthMethod(void* api, struct CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID);
// 用户发出获取图形验证码请求
void* tReqGenUserCaptcha(void* api, struct CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID);
// 用户发出获取短信验证码请求
void* tReqGenUserText(void* api, struct CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID);
// 用户发出带有图片验证码的登陆请求
void* tReqUserLoginWithCaptcha(void* api, struct CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID);
// 用户发出带有短信验证码的登陆请求
void* tReqUserLoginWithText(void* api, struct CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID);
// 用户发出带有动态口令的登陆请求
void* tReqUserLoginWithOTP(void* api, struct CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID);
// 报单录入请求
void* tReqOrderInsert(void* api, struct CThostFtdcInputOrderField *pInputOrder, int nRequestID);
// 预埋单录入请求
void* tReqParkedOrderInsert(void* api, struct CThostFtdcParkedOrderField *pParkedOrder, int nRequestID);
// 预埋撤单录入请求
void* tReqParkedOrderAction(void* api, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID);
// 报单操作请求
void* tReqOrderAction(void* api, struct CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID);
// 查询最大报单数量请求
void* tReqQryMaxOrderVolume(void* api, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID);
// 投资者结算结果确认
void* tReqSettlementInfoConfirm(void* api, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID);
// 请求删除预埋单
void* tReqRemoveParkedOrder(void* api, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID);
// 请求删除预埋撤单
void* tReqRemoveParkedOrderAction(void* api, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID);
// 执行宣告录入请求
void* tReqExecOrderInsert(void* api, struct CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID);
// 执行宣告操作请求
void* tReqExecOrderAction(void* api, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID);
// 询价录入请求
void* tReqForQuoteInsert(void* api, struct CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID);
// 报价录入请求
void* tReqQuoteInsert(void* api, struct CThostFtdcInputQuoteField *pInputQuote, int nRequestID);
// 报价操作请求
void* tReqQuoteAction(void* api, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID);
// 批量报单操作请求
void* tReqBatchOrderAction(void* api, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID);
// 期权自对冲录入请求
void* tReqOptionSelfCloseInsert(void* api, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID);
// 期权自对冲操作请求
void* tReqOptionSelfCloseAction(void* api, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID);
// 申请组合录入请求
void* tReqCombActionInsert(void* api, struct CThostFtdcInputCombActionField *pInputCombAction, int nRequestID);
// 请求查询报单
void* tReqQryOrder(void* api, struct CThostFtdcQryOrderField *pQryOrder, int nRequestID);
// 请求查询成交
void* tReqQryTrade(void* api, struct CThostFtdcQryTradeField *pQryTrade, int nRequestID);
// 请求查询投资者持仓
void* tReqQryInvestorPosition(void* api, struct CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID);
// 请求查询资金账户
void* tReqQryTradingAccount(void* api, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);
// 请求查询投资者
void* tReqQryInvestor(void* api, struct CThostFtdcQryInvestorField *pQryInvestor, int nRequestID);
// 请求查询交易编码
void* tReqQryTradingCode(void* api, struct CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID);
// 请求查询合约保证金率
void* tReqQryInstrumentMarginRate(void* api, struct CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID);
// 请求查询合约手续费率
void* tReqQryInstrumentCommissionRate(void* api, struct CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID);
// 请求查询交易所
void* tReqQryExchange(void* api, struct CThostFtdcQryExchangeField *pQryExchange, int nRequestID);
// 请求查询产品
void* tReqQryProduct(void* api, struct CThostFtdcQryProductField *pQryProduct, int nRequestID);
// 请求查询合约
void* tReqQryInstrument(void* api, struct CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID);
// 请求查询行情
void* tReqQryDepthMarketData(void* api, struct CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID);
// 请求查询交易员报盘机
void* tReqQryTraderOffer(void* api, struct CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID);
// 请求查询投资者结算结果
void* tReqQrySettlementInfo(void* api, struct CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID);
// 请求查询转帐银行
void* tReqQryTransferBank(void* api, struct CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID);
// 请求查询投资者持仓明细
void* tReqQryInvestorPositionDetail(void* api, struct CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID);
// 请求查询客户通知
void* tReqQryNotice(void* api, struct CThostFtdcQryNoticeField *pQryNotice, int nRequestID);
// 请求查询结算信息确认
void* tReqQrySettlementInfoConfirm(void* api, struct CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID);
// 请求查询投资者持仓明细
void* tReqQryInvestorPositionCombineDetail(void* api, struct CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID);
// 请求查询保证金监管系统经纪公司资金账户密钥
void* tReqQryCFMMCTradingAccountKey(void* api, struct CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID);
// 请求查询仓单折抵信息
void* tReqQryEWarrantOffset(void* api, struct CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID);
// 请求查询投资者品种/跨品种保证金
void* tReqQryInvestorProductGroupMargin(void* api, struct CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID);
// 请求查询交易所保证金率
void* tReqQryExchangeMarginRate(void* api, struct CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID);
// 请求查询交易所调整保证金率
void* tReqQryExchangeMarginRateAdjust(void* api, struct CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID);
// 请求查询汇率
void* tReqQryExchangeRate(void* api, struct CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID);
// 请求查询二级代理操作员银期权限
void* tReqQrySecAgentACIDMap(void* api, struct CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID);
// 请求查询产品报价汇率
void* tReqQryProductExchRate(void* api, struct CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID);
// 请求查询产品组
void* tReqQryProductGroup(void* api, struct CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID);
// 请求查询做市商合约手续费率
void* tReqQryMMInstrumentCommissionRate(void* api, struct CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID);
// 请求查询做市商期权合约手续费
void* tReqQryMMOptionInstrCommRate(void* api, struct CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID);
// 请求查询报单手续费
void* tReqQryInstrumentOrderCommRate(void* api, struct CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID);
// 请求查询资金账户
void* tReqQrySecAgentTradingAccount(void* api, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);
// 请求查询二级代理商资金校验模式
void* tReqQrySecAgentCheckMode(void* api, struct CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID);
// 请求查询二级代理商信息
void* tReqQrySecAgentTradeInfo(void* api, struct CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID);
// 请求查询期权交易成本
void* tReqQryOptionInstrTradeCost(void* api, struct CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID);
// 请求查询期权合约手续费
void* tReqQryOptionInstrCommRate(void* api, struct CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID);
// 请求查询执行宣告
void* tReqQryExecOrder(void* api, struct CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID);
// 请求查询询价
void* tReqQryForQuote(void* api, struct CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID);
// 请求查询报价
void* tReqQryQuote(void* api, struct CThostFtdcQryQuoteField *pQryQuote, int nRequestID);
// 请求查询期权自对冲
void* tReqQryOptionSelfClose(void* api, struct CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID);
// 请求查询投资单元
void* tReqQryInvestUnit(void* api, struct CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID);
// 请求查询组合合约安全系数
void* tReqQryCombInstrumentGuard(void* api, struct CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID);
// 请求查询申请组合
void* tReqQryCombAction(void* api, struct CThostFtdcQryCombActionField *pQryCombAction, int nRequestID);
// 请求查询转帐流水
void* tReqQryTransferSerial(void* api, struct CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID);
// 请求查询银期签约关系
void* tReqQryAccountregister(void* api, struct CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID);
// 请求查询签约银行
void* tReqQryContractBank(void* api, struct CThostFtdcQryContractBankField *pQryContractBank, int nRequestID);
// 请求查询预埋单
void* tReqQryParkedOrder(void* api, struct CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID);
// 请求查询预埋撤单
void* tReqQryParkedOrderAction(void* api, struct CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID);
// 请求查询交易通知
void* tReqQryTradingNotice(void* api, struct CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID);
// 请求查询经纪公司交易参数
void* tReqQryBrokerTradingParams(void* api, struct CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID);
// 请求查询经纪公司交易算法
void* tReqQryBrokerTradingAlgos(void* api, struct CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID);
// 请求查询监控中心用户令牌
void* tReqQueryCFMMCTradingAccountToken(void* api, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID);
// 期货发起银行资金转期货请求
void* tReqFromBankToFutureByFuture(void* api, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);
// 期货发起期货资金转银行请求
void* tReqFromFutureToBankByFuture(void* api, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);
// 期货发起查询银行余额请求
void* tReqQueryBankAccountMoneyByFuture(void* api, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID);
// 请求查询分类合约
void* tReqQryClassifiedInstrument(void* api, struct CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID);
// 请求组合优惠比例
void* tReqQryCombPromotionParam(void* api, struct CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID);
// 投资者风险结算持仓查询
void* tReqQryRiskSettleInvstPosition(void* api, struct CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID);
// 风险结算产品查询
void* tReqQryRiskSettleProductStatus(void* api, struct CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID);

// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
void tSetOnFrontConnected(void*, void*);
int tFrontConnected();
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
void tSetOnFrontDisconnected(void*, void*);
int tFrontDisconnected(int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
void tSetOnHeartBeatWarning(void*, void*);
int tHeartBeatWarning(int nTimeLapse);
// 客户端认证响应
void tSetOnRspAuthenticate(void*, void*);
int tRspAuthenticate(struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登录请求响应
void tSetOnRspUserLogin(void*, void*);
int tRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登出请求响应
void tSetOnRspUserLogout(void*, void*);
int tRspUserLogout(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 用户口令更新请求响应
void tSetOnRspUserPasswordUpdate(void*, void*);
int tRspUserPasswordUpdate(struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 资金账户口令更新请求响应
void tSetOnRspTradingAccountPasswordUpdate(void*, void*);
int tRspTradingAccountPasswordUpdate(struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询用户当前支持的认证模式的回复
void tSetOnRspUserAuthMethod(void*, void*);
int tRspUserAuthMethod(struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 获取图形验证码请求的回复
void tSetOnRspGenUserCaptcha(void*, void*);
int tRspGenUserCaptcha(struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 获取短信验证码请求的回复
void tSetOnRspGenUserText(void*, void*);
int tRspGenUserText(struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单录入请求响应
void tSetOnRspOrderInsert(void*, void*);
int tRspOrderInsert(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 预埋单录入请求响应
void tSetOnRspParkedOrderInsert(void*, void*);
int tRspParkedOrderInsert(struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 预埋撤单录入请求响应
void tSetOnRspParkedOrderAction(void*, void*);
int tRspParkedOrderAction(struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单操作请求响应
void tSetOnRspOrderAction(void*, void*);
int tRspOrderAction(struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询最大报单数量响应
void tSetOnRspQryMaxOrderVolume(void*, void*);
int tRspQryMaxOrderVolume(struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者结算结果确认响应
void tSetOnRspSettlementInfoConfirm(void*, void*);
int tRspSettlementInfoConfirm(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 删除预埋单响应
void tSetOnRspRemoveParkedOrder(void*, void*);
int tRspRemoveParkedOrder(struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 删除预埋撤单响应
void tSetOnRspRemoveParkedOrderAction(void*, void*);
int tRspRemoveParkedOrderAction(struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 执行宣告录入请求响应
void tSetOnRspExecOrderInsert(void*, void*);
int tRspExecOrderInsert(struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 执行宣告操作请求响应
void tSetOnRspExecOrderAction(void*, void*);
int tRspExecOrderAction(struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 询价录入请求响应
void tSetOnRspForQuoteInsert(void*, void*);
int tRspForQuoteInsert(struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报价录入请求响应
void tSetOnRspQuoteInsert(void*, void*);
int tRspQuoteInsert(struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报价操作请求响应
void tSetOnRspQuoteAction(void*, void*);
int tRspQuoteAction(struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 批量报单操作请求响应
void tSetOnRspBatchOrderAction(void*, void*);
int tRspBatchOrderAction(struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期权自对冲录入请求响应
void tSetOnRspOptionSelfCloseInsert(void*, void*);
int tRspOptionSelfCloseInsert(struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期权自对冲操作请求响应
void tSetOnRspOptionSelfCloseAction(void*, void*);
int tRspOptionSelfCloseAction(struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 申请组合录入请求响应
void tSetOnRspCombActionInsert(void*, void*);
int tRspCombActionInsert(struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报单响应
void tSetOnRspQryOrder(void*, void*);
int tRspQryOrder(struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询成交响应
void tSetOnRspQryTrade(void*, void*);
int tRspQryTrade(struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓响应
void tSetOnRspQryInvestorPosition(void*, void*);
int tRspQryInvestorPosition(struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询资金账户响应
void tSetOnRspQryTradingAccount(void*, void*);
int tRspQryTradingAccount(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者响应
void tSetOnRspQryInvestor(void*, void*);
int tRspQryInvestor(struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易编码响应
void tSetOnRspQryTradingCode(void*, void*);
int tRspQryTradingCode(struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约保证金率响应
void tSetOnRspQryInstrumentMarginRate(void*, void*);
int tRspQryInstrumentMarginRate(struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约手续费率响应
void tSetOnRspQryInstrumentCommissionRate(void*, void*);
int tRspQryInstrumentCommissionRate(struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所响应
void tSetOnRspQryExchange(void*, void*);
int tRspQryExchange(struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品响应
void tSetOnRspQryProduct(void*, void*);
int tRspQryProduct(struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约响应
void tSetOnRspQryInstrument(void*, void*);
int tRspQryInstrument(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询行情响应
void tSetOnRspQryDepthMarketData(void*, void*);
int tRspQryDepthMarketData(struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易员报盘机响应
void tSetOnRspQryTraderOffer(void*, void*);
int tRspQryTraderOffer(struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者结算结果响应
void tSetOnRspQrySettlementInfo(void*, void*);
int tRspQrySettlementInfo(struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询转帐银行响应
void tSetOnRspQryTransferBank(void*, void*);
int tRspQryTransferBank(struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓明细响应
void tSetOnRspQryInvestorPositionDetail(void*, void*);
int tRspQryInvestorPositionDetail(struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询客户通知响应
void tSetOnRspQryNotice(void*, void*);
int tRspQryNotice(struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询结算信息确认响应
void tSetOnRspQrySettlementInfoConfirm(void*, void*);
int tRspQrySettlementInfoConfirm(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓明细响应
void tSetOnRspQryInvestorPositionCombineDetail(void*, void*);
int tRspQryInvestorPositionCombineDetail(struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询保证金监管系统经纪公司资金账户密钥响应
void tSetOnRspQryCFMMCTradingAccountKey(void*, void*);
int tRspQryCFMMCTradingAccountKey(struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询仓单折抵信息响应
void tSetOnRspQryEWarrantOffset(void*, void*);
int tRspQryEWarrantOffset(struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者品种/跨品种保证金响应
void tSetOnRspQryInvestorProductGroupMargin(void*, void*);
int tRspQryInvestorProductGroupMargin(struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所保证金率响应
void tSetOnRspQryExchangeMarginRate(void*, void*);
int tRspQryExchangeMarginRate(struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所调整保证金率响应
void tSetOnRspQryExchangeMarginRateAdjust(void*, void*);
int tRspQryExchangeMarginRateAdjust(struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询汇率响应
void tSetOnRspQryExchangeRate(void*, void*);
int tRspQryExchangeRate(struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理操作员银期权限响应
void tSetOnRspQrySecAgentACIDMap(void*, void*);
int tRspQrySecAgentACIDMap(struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品报价汇率
void tSetOnRspQryProductExchRate(void*, void*);
int tRspQryProductExchRate(struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品组
void tSetOnRspQryProductGroup(void*, void*);
int tRspQryProductGroup(struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询做市商合约手续费率响应
void tSetOnRspQryMMInstrumentCommissionRate(void*, void*);
int tRspQryMMInstrumentCommissionRate(struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询做市商期权合约手续费响应
void tSetOnRspQryMMOptionInstrCommRate(void*, void*);
int tRspQryMMOptionInstrCommRate(struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报单手续费响应
void tSetOnRspQryInstrumentOrderCommRate(void*, void*);
int tRspQryInstrumentOrderCommRate(struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询资金账户响应
void tSetOnRspQrySecAgentTradingAccount(void*, void*);
int tRspQrySecAgentTradingAccount(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理商资金校验模式响应
void tSetOnRspQrySecAgentCheckMode(void*, void*);
int tRspQrySecAgentCheckMode(struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理商信息响应
void tSetOnRspQrySecAgentTradeInfo(void*, void*);
int tRspQrySecAgentTradeInfo(struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权交易成本响应
void tSetOnRspQryOptionInstrTradeCost(void*, void*);
int tRspQryOptionInstrTradeCost(struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权合约手续费响应
void tSetOnRspQryOptionInstrCommRate(void*, void*);
int tRspQryOptionInstrCommRate(struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询执行宣告响应
void tSetOnRspQryExecOrder(void*, void*);
int tRspQryExecOrder(struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询询价响应
void tSetOnRspQryForQuote(void*, void*);
int tRspQryForQuote(struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报价响应
void tSetOnRspQryQuote(void*, void*);
int tRspQryQuote(struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权自对冲响应
void tSetOnRspQryOptionSelfClose(void*, void*);
int tRspQryOptionSelfClose(struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资单元响应
void tSetOnRspQryInvestUnit(void*, void*);
int tRspQryInvestUnit(struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询组合合约安全系数响应
void tSetOnRspQryCombInstrumentGuard(void*, void*);
int tRspQryCombInstrumentGuard(struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询申请组合响应
void tSetOnRspQryCombAction(void*, void*);
int tRspQryCombAction(struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询转帐流水响应
void tSetOnRspQryTransferSerial(void*, void*);
int tRspQryTransferSerial(struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询银期签约关系响应
void tSetOnRspQryAccountregister(void*, void*);
int tRspQryAccountregister(struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 错误应答
void tSetOnRspError(void*, void*);
int tRspError(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单通知
void tSetOnRtnOrder(void*, void*);
int tRtnOrder(struct CThostFtdcOrderField *pOrder);
// 成交通知
void tSetOnRtnTrade(void*, void*);
int tRtnTrade(struct CThostFtdcTradeField *pTrade);
// 报单录入错误回报
void tSetOnErrRtnOrderInsert(void*, void*);
int tErrRtnOrderInsert(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);
// 报单操作错误回报
void tSetOnErrRtnOrderAction(void*, void*);
int tErrRtnOrderAction(struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 合约交易状态通知
void tSetOnRtnInstrumentStatus(void*, void*);
int tRtnInstrumentStatus(struct CThostFtdcInstrumentStatusField *pInstrumentStatus);
// 交易所公告通知
void tSetOnRtnBulletin(void*, void*);
int tRtnBulletin(struct CThostFtdcBulletinField *pBulletin);
// 交易通知
void tSetOnRtnTradingNotice(void*, void*);
int tRtnTradingNotice(struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);
// 提示条件单校验错误
void tSetOnRtnErrorConditionalOrder(void*, void*);
int tRtnErrorConditionalOrder(struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);
// 执行宣告通知
void tSetOnRtnExecOrder(void*, void*);
int tRtnExecOrder(struct CThostFtdcExecOrderField *pExecOrder);
// 执行宣告录入错误回报
void tSetOnErrRtnExecOrderInsert(void*, void*);
int tErrRtnExecOrderInsert(struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo);
// 执行宣告操作错误回报
void tSetOnErrRtnExecOrderAction(void*, void*);
int tErrRtnExecOrderAction(struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 询价录入错误回报
void tSetOnErrRtnForQuoteInsert(void*, void*);
int tErrRtnForQuoteInsert(struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo);
// 报价通知
void tSetOnRtnQuote(void*, void*);
int tRtnQuote(struct CThostFtdcQuoteField *pQuote);
// 报价录入错误回报
void tSetOnErrRtnQuoteInsert(void*, void*);
int tErrRtnQuoteInsert(struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo);
// 报价操作错误回报
void tSetOnErrRtnQuoteAction(void*, void*);
int tErrRtnQuoteAction(struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo);
// 询价通知
void tSetOnRtnForQuoteRsp(void*, void*);
int tRtnForQuoteRsp(struct CThostFtdcForQuoteRspField *pForQuoteRsp);
// 保证金监控中心用户令牌
void tSetOnRtnCFMMCTradingAccountToken(void*, void*);
int tRtnCFMMCTradingAccountToken(struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);
// 批量报单操作错误回报
void tSetOnErrRtnBatchOrderAction(void*, void*);
int tErrRtnBatchOrderAction(struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 期权自对冲通知
void tSetOnRtnOptionSelfClose(void*, void*);
int tRtnOptionSelfClose(struct CThostFtdcOptionSelfCloseField *pOptionSelfClose);
// 期权自对冲录入错误回报
void tSetOnErrRtnOptionSelfCloseInsert(void*, void*);
int tErrRtnOptionSelfCloseInsert(struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo);
// 期权自对冲操作错误回报
void tSetOnErrRtnOptionSelfCloseAction(void*, void*);
int tErrRtnOptionSelfCloseAction(struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo);
// 申请组合通知
void tSetOnRtnCombAction(void*, void*);
int tRtnCombAction(struct CThostFtdcCombActionField *pCombAction);
// 申请组合录入错误回报
void tSetOnErrRtnCombActionInsert(void*, void*);
int tErrRtnCombActionInsert(struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo);
// 请求查询签约银行响应
void tSetOnRspQryContractBank(void*, void*);
int tRspQryContractBank(struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询预埋单响应
void tSetOnRspQryParkedOrder(void*, void*);
int tRspQryParkedOrder(struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询预埋撤单响应
void tSetOnRspQryParkedOrderAction(void*, void*);
int tRspQryParkedOrderAction(struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易通知响应
void tSetOnRspQryTradingNotice(void*, void*);
int tRspQryTradingNotice(struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询经纪公司交易参数响应
void tSetOnRspQryBrokerTradingParams(void*, void*);
int tRspQryBrokerTradingParams(struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询经纪公司交易算法响应
void tSetOnRspQryBrokerTradingAlgos(void*, void*);
int tRspQryBrokerTradingAlgos(struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询监控中心用户令牌
void tSetOnRspQueryCFMMCTradingAccountToken(void*, void*);
int tRspQueryCFMMCTradingAccountToken(struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 银行发起银行资金转期货通知
void tSetOnRtnFromBankToFutureByBank(void*, void*);
int tRtnFromBankToFutureByBank(struct CThostFtdcRspTransferField *pRspTransfer);
// 银行发起期货资金转银行通知
void tSetOnRtnFromFutureToBankByBank(void*, void*);
int tRtnFromFutureToBankByBank(struct CThostFtdcRspTransferField *pRspTransfer);
// 银行发起冲正银行转期货通知
void tSetOnRtnRepealFromBankToFutureByBank(void*, void*);
int tRtnRepealFromBankToFutureByBank(struct CThostFtdcRspRepealField *pRspRepeal);
// 银行发起冲正期货转银行通知
void tSetOnRtnRepealFromFutureToBankByBank(void*, void*);
int tRtnRepealFromFutureToBankByBank(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起银行资金转期货通知
void tSetOnRtnFromBankToFutureByFuture(void*, void*);
int tRtnFromBankToFutureByFuture(struct CThostFtdcRspTransferField *pRspTransfer);
// 期货发起期货资金转银行通知
void tSetOnRtnFromFutureToBankByFuture(void*, void*);
int tRtnFromFutureToBankByFuture(struct CThostFtdcRspTransferField *pRspTransfer);
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromBankToFutureByFutureManual(void*, void*);
int tRtnRepealFromBankToFutureByFutureManual(struct CThostFtdcRspRepealField *pRspRepeal);
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromFutureToBankByFutureManual(void*, void*);
int tRtnRepealFromFutureToBankByFutureManual(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起查询银行余额通知
void tSetOnRtnQueryBankBalanceByFuture(void*, void*);
int tRtnQueryBankBalanceByFuture(struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);
// 期货发起银行资金转期货错误回报
void tSetOnErrRtnBankToFutureByFuture(void*, void*);
int tErrRtnBankToFutureByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起期货资金转银行错误回报
void tSetOnErrRtnFutureToBankByFuture(void*, void*);
int tErrRtnFutureToBankByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);
// 系统运行时期货端手工发起冲正银行转期货错误回报
void tSetOnErrRtnRepealBankToFutureByFutureManual(void*, void*);
int tErrRtnRepealBankToFutureByFutureManual(struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);
// 系统运行时期货端手工发起冲正期货转银行错误回报
void tSetOnErrRtnRepealFutureToBankByFutureManual(void*, void*);
int tErrRtnRepealFutureToBankByFutureManual(struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起查询银行余额错误回报
void tSetOnErrRtnQueryBankBalanceByFuture(void*, void*);
int tErrRtnQueryBankBalanceByFuture(struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromBankToFutureByFuture(void*, void*);
int tRtnRepealFromBankToFutureByFuture(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromFutureToBankByFuture(void*, void*);
int tRtnRepealFromFutureToBankByFuture(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起银行资金转期货应答
void tSetOnRspFromBankToFutureByFuture(void*, void*);
int tRspFromBankToFutureByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期货发起期货资金转银行应答
void tSetOnRspFromFutureToBankByFuture(void*, void*);
int tRspFromFutureToBankByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期货发起查询银行余额应答
void tSetOnRspQueryBankAccountMoneyByFuture(void*, void*);
int tRspQueryBankAccountMoneyByFuture(struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 银行发起银期开户通知
void tSetOnRtnOpenAccountByBank(void*, void*);
int tRtnOpenAccountByBank(struct CThostFtdcOpenAccountField *pOpenAccount);
// 银行发起银期销户通知
void tSetOnRtnCancelAccountByBank(void*, void*);
int tRtnCancelAccountByBank(struct CThostFtdcCancelAccountField *pCancelAccount);
// 银行发起变更银行账号通知
void tSetOnRtnChangeAccountByBank(void*, void*);
int tRtnChangeAccountByBank(struct CThostFtdcChangeAccountField *pChangeAccount);
// 请求查询分类合约响应
void tSetOnRspQryClassifiedInstrument(void*, void*);
int tRspQryClassifiedInstrument(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求组合优惠比例响应
void tSetOnRspQryCombPromotionParam(void*, void*);
int tRspQryCombPromotionParam(struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者风险结算持仓查询响应
void tSetOnRspQryRiskSettleInvstPosition(void*, void*);
int tRspQryRiskSettleInvstPosition(struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 风险结算产品查询响应
void tSetOnRspQryRiskSettleProductStatus(void*, void*);
int tRspQryRiskSettleProductStatus(struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);

#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"unsafe"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

// Trade 交易接口
type Trade struct {
	goctp.HFTrade
	spi, api unsafe.Pointer
	passWord string // 密码

    // 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
    _FrontConnected func()
    // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    _FrontDisconnected func(nReason int)
    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    _HeartBeatWarning func(nTimeLapse int)
    // 客户端认证响应
    _RspAuthenticate func(pRspAuthenticateField *ctp.CThostFtdcRspAuthenticateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 登录请求响应
    _RspUserLogin func(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 登出请求响应
    _RspUserLogout func(pUserLogout *ctp.CThostFtdcUserLogoutField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 用户口令更新请求响应
    _RspUserPasswordUpdate func(pUserPasswordUpdate *ctp.CThostFtdcUserPasswordUpdateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 资金账户口令更新请求响应
    _RspTradingAccountPasswordUpdate func(pTradingAccountPasswordUpdate *ctp.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 查询用户当前支持的认证模式的回复
    _RspUserAuthMethod func(pRspUserAuthMethod *ctp.CThostFtdcRspUserAuthMethodField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 获取图形验证码请求的回复
    _RspGenUserCaptcha func(pRspGenUserCaptcha *ctp.CThostFtdcRspGenUserCaptchaField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 获取短信验证码请求的回复
    _RspGenUserText func(pRspGenUserText *ctp.CThostFtdcRspGenUserTextField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报单录入请求响应
    _RspOrderInsert func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 预埋单录入请求响应
    _RspParkedOrderInsert func(pParkedOrder *ctp.CThostFtdcParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 预埋撤单录入请求响应
    _RspParkedOrderAction func(pParkedOrderAction *ctp.CThostFtdcParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报单操作请求响应
    _RspOrderAction func(pInputOrderAction *ctp.CThostFtdcInputOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 查询最大报单数量响应
    _RspQryMaxOrderVolume func(pQryMaxOrderVolume *ctp.CThostFtdcQryMaxOrderVolumeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 投资者结算结果确认响应
    _RspSettlementInfoConfirm func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 删除预埋单响应
    _RspRemoveParkedOrder func(pRemoveParkedOrder *ctp.CThostFtdcRemoveParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 删除预埋撤单响应
    _RspRemoveParkedOrderAction func(pRemoveParkedOrderAction *ctp.CThostFtdcRemoveParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 执行宣告录入请求响应
    _RspExecOrderInsert func(pInputExecOrder *ctp.CThostFtdcInputExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 执行宣告操作请求响应
    _RspExecOrderAction func(pInputExecOrderAction *ctp.CThostFtdcInputExecOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 询价录入请求响应
    _RspForQuoteInsert func(pInputForQuote *ctp.CThostFtdcInputForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报价录入请求响应
    _RspQuoteInsert func(pInputQuote *ctp.CThostFtdcInputQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报价操作请求响应
    _RspQuoteAction func(pInputQuoteAction *ctp.CThostFtdcInputQuoteActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 批量报单操作请求响应
    _RspBatchOrderAction func(pInputBatchOrderAction *ctp.CThostFtdcInputBatchOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期权自对冲录入请求响应
    _RspOptionSelfCloseInsert func(pInputOptionSelfClose *ctp.CThostFtdcInputOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期权自对冲操作请求响应
    _RspOptionSelfCloseAction func(pInputOptionSelfCloseAction *ctp.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 申请组合录入请求响应
    _RspCombActionInsert func(pInputCombAction *ctp.CThostFtdcInputCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询报单响应
    _RspQryOrder func(pOrder *ctp.CThostFtdcOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询成交响应
    _RspQryTrade func(pTrade *ctp.CThostFtdcTradeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者持仓响应
    _RspQryInvestorPosition func(pInvestorPosition *ctp.CThostFtdcInvestorPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询资金账户响应
    _RspQryTradingAccount func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者响应
    _RspQryInvestor func(pInvestor *ctp.CThostFtdcInvestorField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易编码响应
    _RspQryTradingCode func(pTradingCode *ctp.CThostFtdcTradingCodeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询合约保证金率响应
    _RspQryInstrumentMarginRate func(pInstrumentMarginRate *ctp.CThostFtdcInstrumentMarginRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询合约手续费率响应
    _RspQryInstrumentCommissionRate func(pInstrumentCommissionRate *ctp.CThostFtdcInstrumentCommissionRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易所响应
    _RspQryExchange func(pExchange *ctp.CThostFtdcExchangeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询产品响应
    _RspQryProduct func(pProduct *ctp.CThostFtdcProductField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询合约响应
    _RspQryInstrument func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询行情响应
    _RspQryDepthMarketData func(pDepthMarketData *ctp.CThostFtdcDepthMarketDataField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易员报盘机响应
    _RspQryTraderOffer func(pTraderOffer *ctp.CThostFtdcTraderOfferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者结算结果响应
    _RspQrySettlementInfo func(pSettlementInfo *ctp.CThostFtdcSettlementInfoField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询转帐银行响应
    _RspQryTransferBank func(pTransferBank *ctp.CThostFtdcTransferBankField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者持仓明细响应
    _RspQryInvestorPositionDetail func(pInvestorPositionDetail *ctp.CThostFtdcInvestorPositionDetailField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询客户通知响应
    _RspQryNotice func(pNotice *ctp.CThostFtdcNoticeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询结算信息确认响应
    _RspQrySettlementInfoConfirm func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者持仓明细响应
    _RspQryInvestorPositionCombineDetail func(pInvestorPositionCombineDetail *ctp.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 查询保证金监管系统经纪公司资金账户密钥响应
    _RspQryCFMMCTradingAccountKey func(pCFMMCTradingAccountKey *ctp.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询仓单折抵信息响应
    _RspQryEWarrantOffset func(pEWarrantOffset *ctp.CThostFtdcEWarrantOffsetField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者品种/跨品种保证金响应
    _RspQryInvestorProductGroupMargin func(pInvestorProductGroupMargin *ctp.CThostFtdcInvestorProductGroupMarginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易所保证金率响应
    _RspQryExchangeMarginRate func(pExchangeMarginRate *ctp.CThostFtdcExchangeMarginRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易所调整保证金率响应
    _RspQryExchangeMarginRateAdjust func(pExchangeMarginRateAdjust *ctp.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询汇率响应
    _RspQryExchangeRate func(pExchangeRate *ctp.CThostFtdcExchangeRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询二级代理操作员银期权限响应
    _RspQrySecAgentACIDMap func(pSecAgentACIDMap *ctp.CThostFtdcSecAgentACIDMapField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询产品报价汇率
    _RspQryProductExchRate func(pProductExchRate *ctp.CThostFtdcProductExchRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询产品组
    _RspQryProductGroup func(pProductGroup *ctp.CThostFtdcProductGroupField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询做市商合约手续费率响应
    _RspQryMMInstrumentCommissionRate func(pMMInstrumentCommissionRate *ctp.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询做市商期权合约手续费响应
    _RspQryMMOptionInstrCommRate func(pMMOptionInstrCommRate *ctp.CThostFtdcMMOptionInstrCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询报单手续费响应
    _RspQryInstrumentOrderCommRate func(pInstrumentOrderCommRate *ctp.CThostFtdcInstrumentOrderCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询资金账户响应
    _RspQrySecAgentTradingAccount func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询二级代理商资金校验模式响应
    _RspQrySecAgentCheckMode func(pSecAgentCheckMode *ctp.CThostFtdcSecAgentCheckModeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询二级代理商信息响应
    _RspQrySecAgentTradeInfo func(pSecAgentTradeInfo *ctp.CThostFtdcSecAgentTradeInfoField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询期权交易成本响应
    _RspQryOptionInstrTradeCost func(pOptionInstrTradeCost *ctp.CThostFtdcOptionInstrTradeCostField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询期权合约手续费响应
    _RspQryOptionInstrCommRate func(pOptionInstrCommRate *ctp.CThostFtdcOptionInstrCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询执行宣告响应
    _RspQryExecOrder func(pExecOrder *ctp.CThostFtdcExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询询价响应
    _RspQryForQuote func(pForQuote *ctp.CThostFtdcForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询报价响应
    _RspQryQuote func(pQuote *ctp.CThostFtdcQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询期权自对冲响应
    _RspQryOptionSelfClose func(pOptionSelfClose *ctp.CThostFtdcOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资单元响应
    _RspQryInvestUnit func(pInvestUnit *ctp.CThostFtdcInvestUnitField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询组合合约安全系数响应
    _RspQryCombInstrumentGuard func(pCombInstrumentGuard *ctp.CThostFtdcCombInstrumentGuardField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询申请组合响应
    _RspQryCombAction func(pCombAction *ctp.CThostFtdcCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询转帐流水响应
    _RspQryTransferSerial func(pTransferSerial *ctp.CThostFtdcTransferSerialField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询银期签约关系响应
    _RspQryAccountregister func(pAccountregister *ctp.CThostFtdcAccountregisterField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 错误应答
    _RspError func(pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报单通知
    _RtnOrder func(pOrder *ctp.CThostFtdcOrderField)
    // 成交通知
    _RtnTrade func(pTrade *ctp.CThostFtdcTradeField)
    // 报单录入错误回报
    _ErrRtnOrderInsert func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 报单操作错误回报
    _ErrRtnOrderAction func(pOrderAction *ctp.CThostFtdcOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 合约交易状态通知
    _RtnInstrumentStatus func(pInstrumentStatus *ctp.CThostFtdcInstrumentStatusField)
    // 交易所公告通知
    _RtnBulletin func(pBulletin *ctp.CThostFtdcBulletinField)
    // 交易通知
    _RtnTradingNotice func(pTradingNoticeInfo *ctp.CThostFtdcTradingNoticeInfoField)
    // 提示条件单校验错误
    _RtnErrorConditionalOrder func(pErrorConditionalOrder *ctp.CThostFtdcErrorConditionalOrderField)
    // 执行宣告通知
    _RtnExecOrder func(pExecOrder *ctp.CThostFtdcExecOrderField)
    // 执行宣告录入错误回报
    _ErrRtnExecOrderInsert func(pInputExecOrder *ctp.CThostFtdcInputExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 执行宣告操作错误回报
    _ErrRtnExecOrderAction func(pExecOrderAction *ctp.CThostFtdcExecOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 询价录入错误回报
    _ErrRtnForQuoteInsert func(pInputForQuote *ctp.CThostFtdcInputForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 报价通知
    _RtnQuote func(pQuote *ctp.CThostFtdcQuoteField)
    // 报价录入错误回报
    _ErrRtnQuoteInsert func(pInputQuote *ctp.CThostFtdcInputQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 报价操作错误回报
    _ErrRtnQuoteAction func(pQuoteAction *ctp.CThostFtdcQuoteActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 询价通知
    _RtnForQuoteRsp func(pForQuoteRsp *ctp.CThostFtdcForQuoteRspField)
    // 保证金监控中心用户令牌
    _RtnCFMMCTradingAccountToken func(pCFMMCTradingAccountToken *ctp.CThostFtdcCFMMCTradingAccountTokenField)
    // 批量报单操作错误回报
    _ErrRtnBatchOrderAction func(pBatchOrderAction *ctp.CThostFtdcBatchOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期权自对冲通知
    _RtnOptionSelfClose func(pOptionSelfClose *ctp.CThostFtdcOptionSelfCloseField)
    // 期权自对冲录入错误回报
    _ErrRtnOptionSelfCloseInsert func(pInputOptionSelfClose *ctp.CThostFtdcInputOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期权自对冲操作错误回报
    _ErrRtnOptionSelfCloseAction func(pOptionSelfCloseAction *ctp.CThostFtdcOptionSelfCloseActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 申请组合通知
    _RtnCombAction func(pCombAction *ctp.CThostFtdcCombActionField)
    // 申请组合录入错误回报
    _ErrRtnCombActionInsert func(pInputCombAction *ctp.CThostFtdcInputCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 请求查询签约银行响应
    _RspQryContractBank func(pContractBank *ctp.CThostFtdcContractBankField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询预埋单响应
    _RspQryParkedOrder func(pParkedOrder *ctp.CThostFtdcParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询预埋撤单响应
    _RspQryParkedOrderAction func(pParkedOrderAction *ctp.CThostFtdcParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易通知响应
    _RspQryTradingNotice func(pTradingNotice *ctp.CThostFtdcTradingNoticeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询经纪公司交易参数响应
    _RspQryBrokerTradingParams func(pBrokerTradingParams *ctp.CThostFtdcBrokerTradingParamsField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询经纪公司交易算法响应
    _RspQryBrokerTradingAlgos func(pBrokerTradingAlgos *ctp.CThostFtdcBrokerTradingAlgosField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询监控中心用户令牌
    _RspQueryCFMMCTradingAccountToken func(pQueryCFMMCTradingAccountToken *ctp.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 银行发起银行资金转期货通知
    _RtnFromBankToFutureByBank func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 银行发起期货资金转银行通知
    _RtnFromFutureToBankByBank func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 银行发起冲正银行转期货通知
    _RtnRepealFromBankToFutureByBank func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 银行发起冲正期货转银行通知
    _RtnRepealFromFutureToBankByBank func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起银行资金转期货通知
    _RtnFromBankToFutureByFuture func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 期货发起期货资金转银行通知
    _RtnFromFutureToBankByFuture func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromBankToFutureByFutureManual func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromFutureToBankByFutureManual func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起查询银行余额通知
    _RtnQueryBankBalanceByFuture func(pNotifyQueryAccount *ctp.CThostFtdcNotifyQueryAccountField)
    // 期货发起银行资金转期货错误回报
    _ErrRtnBankToFutureByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期货发起期货资金转银行错误回报
    _ErrRtnFutureToBankByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 系统运行时期货端手工发起冲正银行转期货错误回报
    _ErrRtnRepealBankToFutureByFutureManual func(pReqRepeal *ctp.CThostFtdcReqRepealField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 系统运行时期货端手工发起冲正期货转银行错误回报
    _ErrRtnRepealFutureToBankByFutureManual func(pReqRepeal *ctp.CThostFtdcReqRepealField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期货发起查询银行余额错误回报
    _ErrRtnQueryBankBalanceByFuture func(pReqQueryAccount *ctp.CThostFtdcReqQueryAccountField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromBankToFutureByFuture func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromFutureToBankByFuture func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起银行资金转期货应答
    _RspFromBankToFutureByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期货发起期货资金转银行应答
    _RspFromFutureToBankByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期货发起查询银行余额应答
    _RspQueryBankAccountMoneyByFuture func(pReqQueryAccount *ctp.CThostFtdcReqQueryAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 银行发起银期开户通知
    _RtnOpenAccountByBank func(pOpenAccount *ctp.CThostFtdcOpenAccountField)
    // 银行发起银期销户通知
    _RtnCancelAccountByBank func(pCancelAccount *ctp.CThostFtdcCancelAccountField)
    // 银行发起变更银行账号通知
    _RtnChangeAccountByBank func(pChangeAccount *ctp.CThostFtdcChangeAccountField)
    // 请求查询分类合约响应
    _RspQryClassifiedInstrument func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求组合优惠比例响应
    _RspQryCombPromotionParam func(pCombPromotionParam *ctp.CThostFtdcCombPromotionParamField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 投资者风险结算持仓查询响应
    _RspQryRiskSettleInvstPosition func(pRiskSettleInvstPosition *ctp.CThostFtdcRiskSettleInvstPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 风险结算产品查询响应
    _RspQryRiskSettleProductStatus func(pRiskSettleProductStatus *ctp.CThostFtdcRiskSettleProductStatusField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    
}

var t *Trade

// NewTrade 实例化
func NewTrade() *Trade {
	t = new(Trade)

	// 主调函数封装 手动添加
	t.HFTrade.GetVersion = func() string {
		return C.GoString((*C.char)(C.tGetVersion()))
	}
	t.HFTrade.ReqConnect = func(addr string) {
		front := C.CString(addr)
		C.tRegisterFront(t.api, front)
		defer C.free(unsafe.Pointer(front))
		C.tSubscribePrivateTopic(t.api, C.int(t.PrivateMode))
		C.tSubscribePublicTopic(t.api, C.int(ctp.THOST_TERT_RESTART)) // 公有流不能用 quick, 会导致交易所状态不更新
		C.tInit(t.api)
		// C.Join(t.api)
	}
	t.HFTrade.ReleaseAPI = func() {
		// C.RegisterSpi(t.api, nil) // 6.6.1说明中提到,会导致程序崩溃
		C.tRelease(t.api) // 前置开着,后台关闭报错. DesignError:pthread_mutex_unlock in line 116 of file ../../source/event/Mutex.h
		// 若不release 则会返回n个 4097后,程序崩溃
		t.spi = nil
		t.api = nil
	}
	t.HFTrade.ReqAuthenticate = func(f *ctp.CThostFtdcReqAuthenticateField, i int) {
		C.tReqAuthenticate(t.api, (*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		C.tReqUserLogin(t.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqOrder = func(f *ctp.CThostFtdcInputOrderField, i int) {
		C.tReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqAction = func(f *ctp.CThostFtdcInputOrderActionField, i int) {
		C.tReqOrderAction(t.api, (*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqFromBankToFutureByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		C.tReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqFromFutureToBankByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		C.tReqFromFutureToBankByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqSettlementInfoConfirm = func(f *ctp.CThostFtdcSettlementInfoConfirmField, i int) {
		C.tReqSettlementInfoConfirm(t.api, (*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInstrument = func(f *ctp.CThostFtdcQryInstrumentField, i int) {
		C.tReqQryInstrument(t.api, (*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryClassifiedInstrument = func(f *ctp.CThostFtdcQryClassifiedInstrumentField, i int) {
		C.tReqQryClassifiedInstrument(t.api, (*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryTradingAccount = func(f *ctp.CThostFtdcQryTradingAccountField, i int) {
		C.tReqQryTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInvestorPosition = func(f *ctp.CThostFtdcQryInvestorPositionField, i int) {
		C.tReqQryInvestorPosition(t.api, (*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryInvestor = func(f *ctp.CThostFtdcQryInvestorField, i int) {
		C.tReqQryInvestor(t.api, (*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryOrder = func(f *ctp.CThostFtdcQryOrderField, i int) {
		C.tReqQryOrder(t.api, (*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(f)), C.int(i))
	}
	t.HFTrade.ReqQryTrade = func(f *ctp.CThostFtdcQryTradeField, i int) {
		C.tReqQryTrade(t.api, (*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(f)), C.int(i))
	}
	
	// HFTrade 响应 手动添加即可增加新功能
	t._RtnFromFutureToBankByFuture = func(pRspTransfer *ctp.CThostFtdcRspTransferField) {
		t.HFTrade.RtnFromFutureToBankByFuture(pRspTransfer)
	}
	t._RtnFromBankToFutureByFuture = func(pRspTransfer *ctp.CThostFtdcRspTransferField) {
		t.HFTrade.RtnFromBankToFutureByFuture(pRspTransfer)
	}
	t._ErrRtnOrderAction = func(pOrderAction *ctp.CThostFtdcOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField) {
		t.HFTrade.ErrRtnOrderAction(pOrderAction, pRspInfo)
	}
	t._ErrRtnOrderInsert = func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField) {
		t.HFTrade.ErrRtnOrderInsert(pInputOrder, pRspInfo)
	}
	t._RtnTrade = func(pTrade *ctp.CThostFtdcTradeField) {
		t.HFTrade.RtnTrade(pTrade)
	}
	t._RtnOrder = func(pOrder *ctp.CThostFtdcOrderField) {
		t.HFTrade.RtnOrder(pOrder)
	}
	t._RtnInstrumentStatus = func(pInstrumentStatus *ctp.CThostFtdcInstrumentStatusField) {
		t.HFTrade.RtnInstrumentStatus(pInstrumentStatus)
	}
	t._RspQryInvestorPosition = func(pInvestorPosition *ctp.CThostFtdcInvestorPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pInvestorPosition == nil{ // 处理空指针
			pInvestorPosition = &ctp.CThostFtdcInvestorPositionField{}
		}
		t.HFTrade.RspQryInvestorPosition(pInvestorPosition, bIsLast)
	}
	t._RspQryTradingAccount = func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryTradingAccount(pTradingAccount, bIsLast)
	}
	t._RspQryTrade = func(pTrade *ctp.CThostFtdcTradeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pTrade == nil{ // 处理空指针
			pTrade = &ctp.CThostFtdcTradeField{}
		}
		t.HFTrade.RspQryTrade(pTrade, bIsLast)
	}
	t._RspQryOrder = func(pOrder *ctp.CThostFtdcOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		if pOrder == nil{ // 处理空指针
			pOrder = &ctp.CThostFtdcOrderField{}
		}
		t.HFTrade.RspQryOrder(pOrder, bIsLast)
	}
	t._RspQryInvestor = func(pInvestor *ctp.CThostFtdcInvestorField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInvestor(pInvestor, bIsLast)
	}
	t._RspQryInstrument = func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInstrument(pInstrument, bIsLast)
	}
	t._RspQryClassifiedInstrument = func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInstrument(pInstrument, bIsLast)
	}
	t._RspSettlementInfoConfirm = func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspSettlementInfoConfirm()
	}
	t._RspUserLogin = func(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspUserLogin(pRspUserLogin, pRspInfo)
	}
	t._RspAuthenticate = func(pRspAuthenticateField *ctp.CThostFtdcRspAuthenticateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspAuthenticate(pRspInfo)
	}
	t._FrontConnected = func() {
		t.HFTrade.FrontConnected()
	}
	t._FrontDisconnected = func(nReason int) {
		t.HFTrade.FrontDisConnected(nReason)
	}

	t.HFTrade.Init() // 初始化

	t.api = C.tCreateApi()
	t.spi = C.tCreateSpi()
	C.tRegisterSpi(t.api, t.spi)

    C.tSetOnFrontConnected(t.spi, C.tFrontConnected)
    C.tSetOnFrontDisconnected(t.spi, C.tFrontDisconnected)
    C.tSetOnHeartBeatWarning(t.spi, C.tHeartBeatWarning)
    C.tSetOnRspAuthenticate(t.spi, C.tRspAuthenticate)
    C.tSetOnRspUserLogin(t.spi, C.tRspUserLogin)
    C.tSetOnRspUserLogout(t.spi, C.tRspUserLogout)
    C.tSetOnRspUserPasswordUpdate(t.spi, C.tRspUserPasswordUpdate)
    C.tSetOnRspTradingAccountPasswordUpdate(t.spi, C.tRspTradingAccountPasswordUpdate)
    C.tSetOnRspUserAuthMethod(t.spi, C.tRspUserAuthMethod)
    C.tSetOnRspGenUserCaptcha(t.spi, C.tRspGenUserCaptcha)
    C.tSetOnRspGenUserText(t.spi, C.tRspGenUserText)
    C.tSetOnRspOrderInsert(t.spi, C.tRspOrderInsert)
    C.tSetOnRspParkedOrderInsert(t.spi, C.tRspParkedOrderInsert)
    C.tSetOnRspParkedOrderAction(t.spi, C.tRspParkedOrderAction)
    C.tSetOnRspOrderAction(t.spi, C.tRspOrderAction)
    C.tSetOnRspQryMaxOrderVolume(t.spi, C.tRspQryMaxOrderVolume)
    C.tSetOnRspSettlementInfoConfirm(t.spi, C.tRspSettlementInfoConfirm)
    C.tSetOnRspRemoveParkedOrder(t.spi, C.tRspRemoveParkedOrder)
    C.tSetOnRspRemoveParkedOrderAction(t.spi, C.tRspRemoveParkedOrderAction)
    C.tSetOnRspExecOrderInsert(t.spi, C.tRspExecOrderInsert)
    C.tSetOnRspExecOrderAction(t.spi, C.tRspExecOrderAction)
    C.tSetOnRspForQuoteInsert(t.spi, C.tRspForQuoteInsert)
    C.tSetOnRspQuoteInsert(t.spi, C.tRspQuoteInsert)
    C.tSetOnRspQuoteAction(t.spi, C.tRspQuoteAction)
    C.tSetOnRspBatchOrderAction(t.spi, C.tRspBatchOrderAction)
    C.tSetOnRspOptionSelfCloseInsert(t.spi, C.tRspOptionSelfCloseInsert)
    C.tSetOnRspOptionSelfCloseAction(t.spi, C.tRspOptionSelfCloseAction)
    C.tSetOnRspCombActionInsert(t.spi, C.tRspCombActionInsert)
    C.tSetOnRspQryOrder(t.spi, C.tRspQryOrder)
    C.tSetOnRspQryTrade(t.spi, C.tRspQryTrade)
    C.tSetOnRspQryInvestorPosition(t.spi, C.tRspQryInvestorPosition)
    C.tSetOnRspQryTradingAccount(t.spi, C.tRspQryTradingAccount)
    C.tSetOnRspQryInvestor(t.spi, C.tRspQryInvestor)
    C.tSetOnRspQryTradingCode(t.spi, C.tRspQryTradingCode)
    C.tSetOnRspQryInstrumentMarginRate(t.spi, C.tRspQryInstrumentMarginRate)
    C.tSetOnRspQryInstrumentCommissionRate(t.spi, C.tRspQryInstrumentCommissionRate)
    C.tSetOnRspQryExchange(t.spi, C.tRspQryExchange)
    C.tSetOnRspQryProduct(t.spi, C.tRspQryProduct)
    C.tSetOnRspQryInstrument(t.spi, C.tRspQryInstrument)
    C.tSetOnRspQryDepthMarketData(t.spi, C.tRspQryDepthMarketData)
    C.tSetOnRspQryTraderOffer(t.spi, C.tRspQryTraderOffer)
    C.tSetOnRspQrySettlementInfo(t.spi, C.tRspQrySettlementInfo)
    C.tSetOnRspQryTransferBank(t.spi, C.tRspQryTransferBank)
    C.tSetOnRspQryInvestorPositionDetail(t.spi, C.tRspQryInvestorPositionDetail)
    C.tSetOnRspQryNotice(t.spi, C.tRspQryNotice)
    C.tSetOnRspQrySettlementInfoConfirm(t.spi, C.tRspQrySettlementInfoConfirm)
    C.tSetOnRspQryInvestorPositionCombineDetail(t.spi, C.tRspQryInvestorPositionCombineDetail)
    C.tSetOnRspQryCFMMCTradingAccountKey(t.spi, C.tRspQryCFMMCTradingAccountKey)
    C.tSetOnRspQryEWarrantOffset(t.spi, C.tRspQryEWarrantOffset)
    C.tSetOnRspQryInvestorProductGroupMargin(t.spi, C.tRspQryInvestorProductGroupMargin)
    C.tSetOnRspQryExchangeMarginRate(t.spi, C.tRspQryExchangeMarginRate)
    C.tSetOnRspQryExchangeMarginRateAdjust(t.spi, C.tRspQryExchangeMarginRateAdjust)
    C.tSetOnRspQryExchangeRate(t.spi, C.tRspQryExchangeRate)
    C.tSetOnRspQrySecAgentACIDMap(t.spi, C.tRspQrySecAgentACIDMap)
    C.tSetOnRspQryProductExchRate(t.spi, C.tRspQryProductExchRate)
    C.tSetOnRspQryProductGroup(t.spi, C.tRspQryProductGroup)
    C.tSetOnRspQryMMInstrumentCommissionRate(t.spi, C.tRspQryMMInstrumentCommissionRate)
    C.tSetOnRspQryMMOptionInstrCommRate(t.spi, C.tRspQryMMOptionInstrCommRate)
    C.tSetOnRspQryInstrumentOrderCommRate(t.spi, C.tRspQryInstrumentOrderCommRate)
    C.tSetOnRspQrySecAgentTradingAccount(t.spi, C.tRspQrySecAgentTradingAccount)
    C.tSetOnRspQrySecAgentCheckMode(t.spi, C.tRspQrySecAgentCheckMode)
    C.tSetOnRspQrySecAgentTradeInfo(t.spi, C.tRspQrySecAgentTradeInfo)
    C.tSetOnRspQryOptionInstrTradeCost(t.spi, C.tRspQryOptionInstrTradeCost)
    C.tSetOnRspQryOptionInstrCommRate(t.spi, C.tRspQryOptionInstrCommRate)
    C.tSetOnRspQryExecOrder(t.spi, C.tRspQryExecOrder)
    C.tSetOnRspQryForQuote(t.spi, C.tRspQryForQuote)
    C.tSetOnRspQryQuote(t.spi, C.tRspQryQuote)
    C.tSetOnRspQryOptionSelfClose(t.spi, C.tRspQryOptionSelfClose)
    C.tSetOnRspQryInvestUnit(t.spi, C.tRspQryInvestUnit)
    C.tSetOnRspQryCombInstrumentGuard(t.spi, C.tRspQryCombInstrumentGuard)
    C.tSetOnRspQryCombAction(t.spi, C.tRspQryCombAction)
    C.tSetOnRspQryTransferSerial(t.spi, C.tRspQryTransferSerial)
    C.tSetOnRspQryAccountregister(t.spi, C.tRspQryAccountregister)
    C.tSetOnRspError(t.spi, C.tRspError)
    C.tSetOnRtnOrder(t.spi, C.tRtnOrder)
    C.tSetOnRtnTrade(t.spi, C.tRtnTrade)
    C.tSetOnErrRtnOrderInsert(t.spi, C.tErrRtnOrderInsert)
    C.tSetOnErrRtnOrderAction(t.spi, C.tErrRtnOrderAction)
    C.tSetOnRtnInstrumentStatus(t.spi, C.tRtnInstrumentStatus)
    C.tSetOnRtnBulletin(t.spi, C.tRtnBulletin)
    C.tSetOnRtnTradingNotice(t.spi, C.tRtnTradingNotice)
    C.tSetOnRtnErrorConditionalOrder(t.spi, C.tRtnErrorConditionalOrder)
    C.tSetOnRtnExecOrder(t.spi, C.tRtnExecOrder)
    C.tSetOnErrRtnExecOrderInsert(t.spi, C.tErrRtnExecOrderInsert)
    C.tSetOnErrRtnExecOrderAction(t.spi, C.tErrRtnExecOrderAction)
    C.tSetOnErrRtnForQuoteInsert(t.spi, C.tErrRtnForQuoteInsert)
    C.tSetOnRtnQuote(t.spi, C.tRtnQuote)
    C.tSetOnErrRtnQuoteInsert(t.spi, C.tErrRtnQuoteInsert)
    C.tSetOnErrRtnQuoteAction(t.spi, C.tErrRtnQuoteAction)
    C.tSetOnRtnForQuoteRsp(t.spi, C.tRtnForQuoteRsp)
    C.tSetOnRtnCFMMCTradingAccountToken(t.spi, C.tRtnCFMMCTradingAccountToken)
    C.tSetOnErrRtnBatchOrderAction(t.spi, C.tErrRtnBatchOrderAction)
    C.tSetOnRtnOptionSelfClose(t.spi, C.tRtnOptionSelfClose)
    C.tSetOnErrRtnOptionSelfCloseInsert(t.spi, C.tErrRtnOptionSelfCloseInsert)
    C.tSetOnErrRtnOptionSelfCloseAction(t.spi, C.tErrRtnOptionSelfCloseAction)
    C.tSetOnRtnCombAction(t.spi, C.tRtnCombAction)
    C.tSetOnErrRtnCombActionInsert(t.spi, C.tErrRtnCombActionInsert)
    C.tSetOnRspQryContractBank(t.spi, C.tRspQryContractBank)
    C.tSetOnRspQryParkedOrder(t.spi, C.tRspQryParkedOrder)
    C.tSetOnRspQryParkedOrderAction(t.spi, C.tRspQryParkedOrderAction)
    C.tSetOnRspQryTradingNotice(t.spi, C.tRspQryTradingNotice)
    C.tSetOnRspQryBrokerTradingParams(t.spi, C.tRspQryBrokerTradingParams)
    C.tSetOnRspQryBrokerTradingAlgos(t.spi, C.tRspQryBrokerTradingAlgos)
    C.tSetOnRspQueryCFMMCTradingAccountToken(t.spi, C.tRspQueryCFMMCTradingAccountToken)
    C.tSetOnRtnFromBankToFutureByBank(t.spi, C.tRtnFromBankToFutureByBank)
    C.tSetOnRtnFromFutureToBankByBank(t.spi, C.tRtnFromFutureToBankByBank)
    C.tSetOnRtnRepealFromBankToFutureByBank(t.spi, C.tRtnRepealFromBankToFutureByBank)
    C.tSetOnRtnRepealFromFutureToBankByBank(t.spi, C.tRtnRepealFromFutureToBankByBank)
    C.tSetOnRtnFromBankToFutureByFuture(t.spi, C.tRtnFromBankToFutureByFuture)
    C.tSetOnRtnFromFutureToBankByFuture(t.spi, C.tRtnFromFutureToBankByFuture)
    C.tSetOnRtnRepealFromBankToFutureByFutureManual(t.spi, C.tRtnRepealFromBankToFutureByFutureManual)
    C.tSetOnRtnRepealFromFutureToBankByFutureManual(t.spi, C.tRtnRepealFromFutureToBankByFutureManual)
    C.tSetOnRtnQueryBankBalanceByFuture(t.spi, C.tRtnQueryBankBalanceByFuture)
    C.tSetOnErrRtnBankToFutureByFuture(t.spi, C.tErrRtnBankToFutureByFuture)
    C.tSetOnErrRtnFutureToBankByFuture(t.spi, C.tErrRtnFutureToBankByFuture)
    C.tSetOnErrRtnRepealBankToFutureByFutureManual(t.spi, C.tErrRtnRepealBankToFutureByFutureManual)
    C.tSetOnErrRtnRepealFutureToBankByFutureManual(t.spi, C.tErrRtnRepealFutureToBankByFutureManual)
    C.tSetOnErrRtnQueryBankBalanceByFuture(t.spi, C.tErrRtnQueryBankBalanceByFuture)
    C.tSetOnRtnRepealFromBankToFutureByFuture(t.spi, C.tRtnRepealFromBankToFutureByFuture)
    C.tSetOnRtnRepealFromFutureToBankByFuture(t.spi, C.tRtnRepealFromFutureToBankByFuture)
    C.tSetOnRspFromBankToFutureByFuture(t.spi, C.tRspFromBankToFutureByFuture)
    C.tSetOnRspFromFutureToBankByFuture(t.spi, C.tRspFromFutureToBankByFuture)
    C.tSetOnRspQueryBankAccountMoneyByFuture(t.spi, C.tRspQueryBankAccountMoneyByFuture)
    C.tSetOnRtnOpenAccountByBank(t.spi, C.tRtnOpenAccountByBank)
    C.tSetOnRtnCancelAccountByBank(t.spi, C.tRtnCancelAccountByBank)
    C.tSetOnRtnChangeAccountByBank(t.spi, C.tRtnChangeAccountByBank)
    C.tSetOnRspQryClassifiedInstrument(t.spi, C.tRspQryClassifiedInstrument)
    C.tSetOnRspQryCombPromotionParam(t.spi, C.tRspQryCombPromotionParam)
    C.tSetOnRspQryRiskSettleInvstPosition(t.spi, C.tRspQryRiskSettleInvstPosition)
    C.tSetOnRspQryRiskSettleProductStatus(t.spi, C.tRspQryRiskSettleProductStatus)
    
	return t
}

//export tFrontConnected
func tFrontConnected() C.int {
    if t._FrontConnected != nil{
        t._FrontConnected()
    }
	return 0
}

//export tFrontDisconnected
func tFrontDisconnected(nReason C.int) C.int {
    if t._FrontDisconnected != nil{
        t._FrontDisconnected(int(nReason))
    }
	return 0
}

//export tHeartBeatWarning
func tHeartBeatWarning(nTimeLapse C.int) C.int {
    if t._HeartBeatWarning != nil{
        t._HeartBeatWarning(int(nTimeLapse))
    }
	return 0
}

//export tRspAuthenticate
func tRspAuthenticate(pRspAuthenticateField *C.struct_CThostFtdcRspAuthenticateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspAuthenticate != nil{
        t._RspAuthenticate((*ctp.CThostFtdcRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspUserLogin
func tRspUserLogin(pRspUserLogin *C.struct_CThostFtdcRspUserLoginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspUserLogin != nil{
        t._RspUserLogin((*ctp.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspUserLogout
func tRspUserLogout(pUserLogout *C.struct_CThostFtdcUserLogoutField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspUserLogout != nil{
        t._RspUserLogout((*ctp.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspUserPasswordUpdate
func tRspUserPasswordUpdate(pUserPasswordUpdate *C.struct_CThostFtdcUserPasswordUpdateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspUserPasswordUpdate != nil{
        t._RspUserPasswordUpdate((*ctp.CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspTradingAccountPasswordUpdate
func tRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *C.struct_CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspTradingAccountPasswordUpdate != nil{
        t._RspTradingAccountPasswordUpdate((*ctp.CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspUserAuthMethod
func tRspUserAuthMethod(pRspUserAuthMethod *C.struct_CThostFtdcRspUserAuthMethodField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspUserAuthMethod != nil{
        t._RspUserAuthMethod((*ctp.CThostFtdcRspUserAuthMethodField)(unsafe.Pointer(pRspUserAuthMethod)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspGenUserCaptcha
func tRspGenUserCaptcha(pRspGenUserCaptcha *C.struct_CThostFtdcRspGenUserCaptchaField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspGenUserCaptcha != nil{
        t._RspGenUserCaptcha((*ctp.CThostFtdcRspGenUserCaptchaField)(unsafe.Pointer(pRspGenUserCaptcha)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspGenUserText
func tRspGenUserText(pRspGenUserText *C.struct_CThostFtdcRspGenUserTextField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspGenUserText != nil{
        t._RspGenUserText((*ctp.CThostFtdcRspGenUserTextField)(unsafe.Pointer(pRspGenUserText)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspOrderInsert
func tRspOrderInsert(pInputOrder *C.struct_CThostFtdcInputOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspOrderInsert != nil{
        t._RspOrderInsert((*ctp.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspParkedOrderInsert
func tRspParkedOrderInsert(pParkedOrder *C.struct_CThostFtdcParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspParkedOrderInsert != nil{
        t._RspParkedOrderInsert((*ctp.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspParkedOrderAction
func tRspParkedOrderAction(pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspParkedOrderAction != nil{
        t._RspParkedOrderAction((*ctp.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspOrderAction
func tRspOrderAction(pInputOrderAction *C.struct_CThostFtdcInputOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspOrderAction != nil{
        t._RspOrderAction((*ctp.CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryMaxOrderVolume
func tRspQryMaxOrderVolume(pQryMaxOrderVolume *C.struct_CThostFtdcQryMaxOrderVolumeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryMaxOrderVolume != nil{
        t._RspQryMaxOrderVolume((*ctp.CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspSettlementInfoConfirm
func tRspSettlementInfoConfirm(pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspSettlementInfoConfirm != nil{
        t._RspSettlementInfoConfirm((*ctp.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspRemoveParkedOrder
func tRspRemoveParkedOrder(pRemoveParkedOrder *C.struct_CThostFtdcRemoveParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspRemoveParkedOrder != nil{
        t._RspRemoveParkedOrder((*ctp.CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspRemoveParkedOrderAction
func tRspRemoveParkedOrderAction(pRemoveParkedOrderAction *C.struct_CThostFtdcRemoveParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspRemoveParkedOrderAction != nil{
        t._RspRemoveParkedOrderAction((*ctp.CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspExecOrderInsert
func tRspExecOrderInsert(pInputExecOrder *C.struct_CThostFtdcInputExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspExecOrderInsert != nil{
        t._RspExecOrderInsert((*ctp.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspExecOrderAction
func tRspExecOrderAction(pInputExecOrderAction *C.struct_CThostFtdcInputExecOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspExecOrderAction != nil{
        t._RspExecOrderAction((*ctp.CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspForQuoteInsert
func tRspForQuoteInsert(pInputForQuote *C.struct_CThostFtdcInputForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspForQuoteInsert != nil{
        t._RspForQuoteInsert((*ctp.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQuoteInsert
func tRspQuoteInsert(pInputQuote *C.struct_CThostFtdcInputQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQuoteInsert != nil{
        t._RspQuoteInsert((*ctp.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQuoteAction
func tRspQuoteAction(pInputQuoteAction *C.struct_CThostFtdcInputQuoteActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQuoteAction != nil{
        t._RspQuoteAction((*ctp.CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspBatchOrderAction
func tRspBatchOrderAction(pInputBatchOrderAction *C.struct_CThostFtdcInputBatchOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspBatchOrderAction != nil{
        t._RspBatchOrderAction((*ctp.CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspOptionSelfCloseInsert
func tRspOptionSelfCloseInsert(pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspOptionSelfCloseInsert != nil{
        t._RspOptionSelfCloseInsert((*ctp.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspOptionSelfCloseAction
func tRspOptionSelfCloseAction(pInputOptionSelfCloseAction *C.struct_CThostFtdcInputOptionSelfCloseActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspOptionSelfCloseAction != nil{
        t._RspOptionSelfCloseAction((*ctp.CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspCombActionInsert
func tRspCombActionInsert(pInputCombAction *C.struct_CThostFtdcInputCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspCombActionInsert != nil{
        t._RspCombActionInsert((*ctp.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryOrder
func tRspQryOrder(pOrder *C.struct_CThostFtdcOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryOrder != nil{
        t._RspQryOrder((*ctp.CThostFtdcOrderField)(unsafe.Pointer(pOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTrade
func tRspQryTrade(pTrade *C.struct_CThostFtdcTradeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTrade != nil{
        t._RspQryTrade((*ctp.CThostFtdcTradeField)(unsafe.Pointer(pTrade)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInvestorPosition
func tRspQryInvestorPosition(pInvestorPosition *C.struct_CThostFtdcInvestorPositionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInvestorPosition != nil{
        t._RspQryInvestorPosition((*ctp.CThostFtdcInvestorPositionField)(unsafe.Pointer(pInvestorPosition)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTradingAccount
func tRspQryTradingAccount(pTradingAccount *C.struct_CThostFtdcTradingAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTradingAccount != nil{
        t._RspQryTradingAccount((*ctp.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInvestor
func tRspQryInvestor(pInvestor *C.struct_CThostFtdcInvestorField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInvestor != nil{
        t._RspQryInvestor((*ctp.CThostFtdcInvestorField)(unsafe.Pointer(pInvestor)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTradingCode
func tRspQryTradingCode(pTradingCode *C.struct_CThostFtdcTradingCodeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTradingCode != nil{
        t._RspQryTradingCode((*ctp.CThostFtdcTradingCodeField)(unsafe.Pointer(pTradingCode)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInstrumentMarginRate
func tRspQryInstrumentMarginRate(pInstrumentMarginRate *C.struct_CThostFtdcInstrumentMarginRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInstrumentMarginRate != nil{
        t._RspQryInstrumentMarginRate((*ctp.CThostFtdcInstrumentMarginRateField)(unsafe.Pointer(pInstrumentMarginRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInstrumentCommissionRate
func tRspQryInstrumentCommissionRate(pInstrumentCommissionRate *C.struct_CThostFtdcInstrumentCommissionRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInstrumentCommissionRate != nil{
        t._RspQryInstrumentCommissionRate((*ctp.CThostFtdcInstrumentCommissionRateField)(unsafe.Pointer(pInstrumentCommissionRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryExchange
func tRspQryExchange(pExchange *C.struct_CThostFtdcExchangeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryExchange != nil{
        t._RspQryExchange((*ctp.CThostFtdcExchangeField)(unsafe.Pointer(pExchange)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryProduct
func tRspQryProduct(pProduct *C.struct_CThostFtdcProductField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryProduct != nil{
        t._RspQryProduct((*ctp.CThostFtdcProductField)(unsafe.Pointer(pProduct)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInstrument
func tRspQryInstrument(pInstrument *C.struct_CThostFtdcInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInstrument != nil{
        t._RspQryInstrument((*ctp.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryDepthMarketData
func tRspQryDepthMarketData(pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryDepthMarketData != nil{
        t._RspQryDepthMarketData((*ctp.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTraderOffer
func tRspQryTraderOffer(pTraderOffer *C.struct_CThostFtdcTraderOfferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTraderOffer != nil{
        t._RspQryTraderOffer((*ctp.CThostFtdcTraderOfferField)(unsafe.Pointer(pTraderOffer)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQrySettlementInfo
func tRspQrySettlementInfo(pSettlementInfo *C.struct_CThostFtdcSettlementInfoField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQrySettlementInfo != nil{
        t._RspQrySettlementInfo((*ctp.CThostFtdcSettlementInfoField)(unsafe.Pointer(pSettlementInfo)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTransferBank
func tRspQryTransferBank(pTransferBank *C.struct_CThostFtdcTransferBankField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTransferBank != nil{
        t._RspQryTransferBank((*ctp.CThostFtdcTransferBankField)(unsafe.Pointer(pTransferBank)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInvestorPositionDetail
func tRspQryInvestorPositionDetail(pInvestorPositionDetail *C.struct_CThostFtdcInvestorPositionDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInvestorPositionDetail != nil{
        t._RspQryInvestorPositionDetail((*ctp.CThostFtdcInvestorPositionDetailField)(unsafe.Pointer(pInvestorPositionDetail)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryNotice
func tRspQryNotice(pNotice *C.struct_CThostFtdcNoticeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryNotice != nil{
        t._RspQryNotice((*ctp.CThostFtdcNoticeField)(unsafe.Pointer(pNotice)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQrySettlementInfoConfirm
func tRspQrySettlementInfoConfirm(pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQrySettlementInfoConfirm != nil{
        t._RspQrySettlementInfoConfirm((*ctp.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInvestorPositionCombineDetail
func tRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *C.struct_CThostFtdcInvestorPositionCombineDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInvestorPositionCombineDetail != nil{
        t._RspQryInvestorPositionCombineDetail((*ctp.CThostFtdcInvestorPositionCombineDetailField)(unsafe.Pointer(pInvestorPositionCombineDetail)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryCFMMCTradingAccountKey
func tRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *C.struct_CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryCFMMCTradingAccountKey != nil{
        t._RspQryCFMMCTradingAccountKey((*ctp.CThostFtdcCFMMCTradingAccountKeyField)(unsafe.Pointer(pCFMMCTradingAccountKey)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryEWarrantOffset
func tRspQryEWarrantOffset(pEWarrantOffset *C.struct_CThostFtdcEWarrantOffsetField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryEWarrantOffset != nil{
        t._RspQryEWarrantOffset((*ctp.CThostFtdcEWarrantOffsetField)(unsafe.Pointer(pEWarrantOffset)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInvestorProductGroupMargin
func tRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *C.struct_CThostFtdcInvestorProductGroupMarginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInvestorProductGroupMargin != nil{
        t._RspQryInvestorProductGroupMargin((*ctp.CThostFtdcInvestorProductGroupMarginField)(unsafe.Pointer(pInvestorProductGroupMargin)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryExchangeMarginRate
func tRspQryExchangeMarginRate(pExchangeMarginRate *C.struct_CThostFtdcExchangeMarginRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryExchangeMarginRate != nil{
        t._RspQryExchangeMarginRate((*ctp.CThostFtdcExchangeMarginRateField)(unsafe.Pointer(pExchangeMarginRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryExchangeMarginRateAdjust
func tRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *C.struct_CThostFtdcExchangeMarginRateAdjustField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryExchangeMarginRateAdjust != nil{
        t._RspQryExchangeMarginRateAdjust((*ctp.CThostFtdcExchangeMarginRateAdjustField)(unsafe.Pointer(pExchangeMarginRateAdjust)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryExchangeRate
func tRspQryExchangeRate(pExchangeRate *C.struct_CThostFtdcExchangeRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryExchangeRate != nil{
        t._RspQryExchangeRate((*ctp.CThostFtdcExchangeRateField)(unsafe.Pointer(pExchangeRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQrySecAgentACIDMap
func tRspQrySecAgentACIDMap(pSecAgentACIDMap *C.struct_CThostFtdcSecAgentACIDMapField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQrySecAgentACIDMap != nil{
        t._RspQrySecAgentACIDMap((*ctp.CThostFtdcSecAgentACIDMapField)(unsafe.Pointer(pSecAgentACIDMap)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryProductExchRate
func tRspQryProductExchRate(pProductExchRate *C.struct_CThostFtdcProductExchRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryProductExchRate != nil{
        t._RspQryProductExchRate((*ctp.CThostFtdcProductExchRateField)(unsafe.Pointer(pProductExchRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryProductGroup
func tRspQryProductGroup(pProductGroup *C.struct_CThostFtdcProductGroupField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryProductGroup != nil{
        t._RspQryProductGroup((*ctp.CThostFtdcProductGroupField)(unsafe.Pointer(pProductGroup)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryMMInstrumentCommissionRate
func tRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *C.struct_CThostFtdcMMInstrumentCommissionRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryMMInstrumentCommissionRate != nil{
        t._RspQryMMInstrumentCommissionRate((*ctp.CThostFtdcMMInstrumentCommissionRateField)(unsafe.Pointer(pMMInstrumentCommissionRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryMMOptionInstrCommRate
func tRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *C.struct_CThostFtdcMMOptionInstrCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryMMOptionInstrCommRate != nil{
        t._RspQryMMOptionInstrCommRate((*ctp.CThostFtdcMMOptionInstrCommRateField)(unsafe.Pointer(pMMOptionInstrCommRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInstrumentOrderCommRate
func tRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *C.struct_CThostFtdcInstrumentOrderCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInstrumentOrderCommRate != nil{
        t._RspQryInstrumentOrderCommRate((*ctp.CThostFtdcInstrumentOrderCommRateField)(unsafe.Pointer(pInstrumentOrderCommRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQrySecAgentTradingAccount
func tRspQrySecAgentTradingAccount(pTradingAccount *C.struct_CThostFtdcTradingAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQrySecAgentTradingAccount != nil{
        t._RspQrySecAgentTradingAccount((*ctp.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQrySecAgentCheckMode
func tRspQrySecAgentCheckMode(pSecAgentCheckMode *C.struct_CThostFtdcSecAgentCheckModeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQrySecAgentCheckMode != nil{
        t._RspQrySecAgentCheckMode((*ctp.CThostFtdcSecAgentCheckModeField)(unsafe.Pointer(pSecAgentCheckMode)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQrySecAgentTradeInfo
func tRspQrySecAgentTradeInfo(pSecAgentTradeInfo *C.struct_CThostFtdcSecAgentTradeInfoField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQrySecAgentTradeInfo != nil{
        t._RspQrySecAgentTradeInfo((*ctp.CThostFtdcSecAgentTradeInfoField)(unsafe.Pointer(pSecAgentTradeInfo)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryOptionInstrTradeCost
func tRspQryOptionInstrTradeCost(pOptionInstrTradeCost *C.struct_CThostFtdcOptionInstrTradeCostField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryOptionInstrTradeCost != nil{
        t._RspQryOptionInstrTradeCost((*ctp.CThostFtdcOptionInstrTradeCostField)(unsafe.Pointer(pOptionInstrTradeCost)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryOptionInstrCommRate
func tRspQryOptionInstrCommRate(pOptionInstrCommRate *C.struct_CThostFtdcOptionInstrCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryOptionInstrCommRate != nil{
        t._RspQryOptionInstrCommRate((*ctp.CThostFtdcOptionInstrCommRateField)(unsafe.Pointer(pOptionInstrCommRate)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryExecOrder
func tRspQryExecOrder(pExecOrder *C.struct_CThostFtdcExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryExecOrder != nil{
        t._RspQryExecOrder((*ctp.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryForQuote
func tRspQryForQuote(pForQuote *C.struct_CThostFtdcForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryForQuote != nil{
        t._RspQryForQuote((*ctp.CThostFtdcForQuoteField)(unsafe.Pointer(pForQuote)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryQuote
func tRspQryQuote(pQuote *C.struct_CThostFtdcQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryQuote != nil{
        t._RspQryQuote((*ctp.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryOptionSelfClose
func tRspQryOptionSelfClose(pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryOptionSelfClose != nil{
        t._RspQryOptionSelfClose((*ctp.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryInvestUnit
func tRspQryInvestUnit(pInvestUnit *C.struct_CThostFtdcInvestUnitField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryInvestUnit != nil{
        t._RspQryInvestUnit((*ctp.CThostFtdcInvestUnitField)(unsafe.Pointer(pInvestUnit)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryCombInstrumentGuard
func tRspQryCombInstrumentGuard(pCombInstrumentGuard *C.struct_CThostFtdcCombInstrumentGuardField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryCombInstrumentGuard != nil{
        t._RspQryCombInstrumentGuard((*ctp.CThostFtdcCombInstrumentGuardField)(unsafe.Pointer(pCombInstrumentGuard)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryCombAction
func tRspQryCombAction(pCombAction *C.struct_CThostFtdcCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryCombAction != nil{
        t._RspQryCombAction((*ctp.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTransferSerial
func tRspQryTransferSerial(pTransferSerial *C.struct_CThostFtdcTransferSerialField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTransferSerial != nil{
        t._RspQryTransferSerial((*ctp.CThostFtdcTransferSerialField)(unsafe.Pointer(pTransferSerial)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryAccountregister
func tRspQryAccountregister(pAccountregister *C.struct_CThostFtdcAccountregisterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryAccountregister != nil{
        t._RspQryAccountregister((*ctp.CThostFtdcAccountregisterField)(unsafe.Pointer(pAccountregister)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspError
func tRspError(pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspError != nil{
        t._RspError((*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRtnOrder
func tRtnOrder(pOrder *C.struct_CThostFtdcOrderField) C.int {
    if t._RtnOrder != nil{
        t._RtnOrder((*ctp.CThostFtdcOrderField)(unsafe.Pointer(pOrder)))
    }
	return 0
}

//export tRtnTrade
func tRtnTrade(pTrade *C.struct_CThostFtdcTradeField) C.int {
    if t._RtnTrade != nil{
        t._RtnTrade((*ctp.CThostFtdcTradeField)(unsafe.Pointer(pTrade)))
    }
	return 0
}

//export tErrRtnOrderInsert
func tErrRtnOrderInsert(pInputOrder *C.struct_CThostFtdcInputOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnOrderInsert != nil{
        t._ErrRtnOrderInsert((*ctp.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnOrderAction
func tErrRtnOrderAction(pOrderAction *C.struct_CThostFtdcOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnOrderAction != nil{
        t._ErrRtnOrderAction((*ctp.CThostFtdcOrderActionField)(unsafe.Pointer(pOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRtnInstrumentStatus
func tRtnInstrumentStatus(pInstrumentStatus *C.struct_CThostFtdcInstrumentStatusField) C.int {
    if t._RtnInstrumentStatus != nil{
        t._RtnInstrumentStatus((*ctp.CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)))
    }
	return 0
}

//export tRtnBulletin
func tRtnBulletin(pBulletin *C.struct_CThostFtdcBulletinField) C.int {
    if t._RtnBulletin != nil{
        t._RtnBulletin((*ctp.CThostFtdcBulletinField)(unsafe.Pointer(pBulletin)))
    }
	return 0
}

//export tRtnTradingNotice
func tRtnTradingNotice(pTradingNoticeInfo *C.struct_CThostFtdcTradingNoticeInfoField) C.int {
    if t._RtnTradingNotice != nil{
        t._RtnTradingNotice((*ctp.CThostFtdcTradingNoticeInfoField)(unsafe.Pointer(pTradingNoticeInfo)))
    }
	return 0
}

//export tRtnErrorConditionalOrder
func tRtnErrorConditionalOrder(pErrorConditionalOrder *C.struct_CThostFtdcErrorConditionalOrderField) C.int {
    if t._RtnErrorConditionalOrder != nil{
        t._RtnErrorConditionalOrder((*ctp.CThostFtdcErrorConditionalOrderField)(unsafe.Pointer(pErrorConditionalOrder)))
    }
	return 0
}

//export tRtnExecOrder
func tRtnExecOrder(pExecOrder *C.struct_CThostFtdcExecOrderField) C.int {
    if t._RtnExecOrder != nil{
        t._RtnExecOrder((*ctp.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)))
    }
	return 0
}

//export tErrRtnExecOrderInsert
func tErrRtnExecOrderInsert(pInputExecOrder *C.struct_CThostFtdcInputExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnExecOrderInsert != nil{
        t._ErrRtnExecOrderInsert((*ctp.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnExecOrderAction
func tErrRtnExecOrderAction(pExecOrderAction *C.struct_CThostFtdcExecOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnExecOrderAction != nil{
        t._ErrRtnExecOrderAction((*ctp.CThostFtdcExecOrderActionField)(unsafe.Pointer(pExecOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnForQuoteInsert
func tErrRtnForQuoteInsert(pInputForQuote *C.struct_CThostFtdcInputForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnForQuoteInsert != nil{
        t._ErrRtnForQuoteInsert((*ctp.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRtnQuote
func tRtnQuote(pQuote *C.struct_CThostFtdcQuoteField) C.int {
    if t._RtnQuote != nil{
        t._RtnQuote((*ctp.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)))
    }
	return 0
}

//export tErrRtnQuoteInsert
func tErrRtnQuoteInsert(pInputQuote *C.struct_CThostFtdcInputQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnQuoteInsert != nil{
        t._ErrRtnQuoteInsert((*ctp.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnQuoteAction
func tErrRtnQuoteAction(pQuoteAction *C.struct_CThostFtdcQuoteActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnQuoteAction != nil{
        t._ErrRtnQuoteAction((*ctp.CThostFtdcQuoteActionField)(unsafe.Pointer(pQuoteAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRtnForQuoteRsp
func tRtnForQuoteRsp(pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField) C.int {
    if t._RtnForQuoteRsp != nil{
        t._RtnForQuoteRsp((*ctp.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)))
    }
	return 0
}

//export tRtnCFMMCTradingAccountToken
func tRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *C.struct_CThostFtdcCFMMCTradingAccountTokenField) C.int {
    if t._RtnCFMMCTradingAccountToken != nil{
        t._RtnCFMMCTradingAccountToken((*ctp.CThostFtdcCFMMCTradingAccountTokenField)(unsafe.Pointer(pCFMMCTradingAccountToken)))
    }
	return 0
}

//export tErrRtnBatchOrderAction
func tErrRtnBatchOrderAction(pBatchOrderAction *C.struct_CThostFtdcBatchOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnBatchOrderAction != nil{
        t._ErrRtnBatchOrderAction((*ctp.CThostFtdcBatchOrderActionField)(unsafe.Pointer(pBatchOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRtnOptionSelfClose
func tRtnOptionSelfClose(pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField) C.int {
    if t._RtnOptionSelfClose != nil{
        t._RtnOptionSelfClose((*ctp.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)))
    }
	return 0
}

//export tErrRtnOptionSelfCloseInsert
func tErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnOptionSelfCloseInsert != nil{
        t._ErrRtnOptionSelfCloseInsert((*ctp.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnOptionSelfCloseAction
func tErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *C.struct_CThostFtdcOptionSelfCloseActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnOptionSelfCloseAction != nil{
        t._ErrRtnOptionSelfCloseAction((*ctp.CThostFtdcOptionSelfCloseActionField)(unsafe.Pointer(pOptionSelfCloseAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRtnCombAction
func tRtnCombAction(pCombAction *C.struct_CThostFtdcCombActionField) C.int {
    if t._RtnCombAction != nil{
        t._RtnCombAction((*ctp.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)))
    }
	return 0
}

//export tErrRtnCombActionInsert
func tErrRtnCombActionInsert(pInputCombAction *C.struct_CThostFtdcInputCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnCombActionInsert != nil{
        t._ErrRtnCombActionInsert((*ctp.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRspQryContractBank
func tRspQryContractBank(pContractBank *C.struct_CThostFtdcContractBankField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryContractBank != nil{
        t._RspQryContractBank((*ctp.CThostFtdcContractBankField)(unsafe.Pointer(pContractBank)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryParkedOrder
func tRspQryParkedOrder(pParkedOrder *C.struct_CThostFtdcParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryParkedOrder != nil{
        t._RspQryParkedOrder((*ctp.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryParkedOrderAction
func tRspQryParkedOrderAction(pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryParkedOrderAction != nil{
        t._RspQryParkedOrderAction((*ctp.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryTradingNotice
func tRspQryTradingNotice(pTradingNotice *C.struct_CThostFtdcTradingNoticeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryTradingNotice != nil{
        t._RspQryTradingNotice((*ctp.CThostFtdcTradingNoticeField)(unsafe.Pointer(pTradingNotice)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryBrokerTradingParams
func tRspQryBrokerTradingParams(pBrokerTradingParams *C.struct_CThostFtdcBrokerTradingParamsField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryBrokerTradingParams != nil{
        t._RspQryBrokerTradingParams((*ctp.CThostFtdcBrokerTradingParamsField)(unsafe.Pointer(pBrokerTradingParams)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryBrokerTradingAlgos
func tRspQryBrokerTradingAlgos(pBrokerTradingAlgos *C.struct_CThostFtdcBrokerTradingAlgosField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryBrokerTradingAlgos != nil{
        t._RspQryBrokerTradingAlgos((*ctp.CThostFtdcBrokerTradingAlgosField)(unsafe.Pointer(pBrokerTradingAlgos)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQueryCFMMCTradingAccountToken
func tRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQueryCFMMCTradingAccountToken != nil{
        t._RspQueryCFMMCTradingAccountToken((*ctp.CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRtnFromBankToFutureByBank
func tRtnFromBankToFutureByBank(pRspTransfer *C.struct_CThostFtdcRspTransferField) C.int {
    if t._RtnFromBankToFutureByBank != nil{
        t._RtnFromBankToFutureByBank((*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
    }
	return 0
}

//export tRtnFromFutureToBankByBank
func tRtnFromFutureToBankByBank(pRspTransfer *C.struct_CThostFtdcRspTransferField) C.int {
    if t._RtnFromFutureToBankByBank != nil{
        t._RtnFromFutureToBankByBank((*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
    }
	return 0
}

//export tRtnRepealFromBankToFutureByBank
func tRtnRepealFromBankToFutureByBank(pRspRepeal *C.struct_CThostFtdcRspRepealField) C.int {
    if t._RtnRepealFromBankToFutureByBank != nil{
        t._RtnRepealFromBankToFutureByBank((*ctp.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
    }
	return 0
}

//export tRtnRepealFromFutureToBankByBank
func tRtnRepealFromFutureToBankByBank(pRspRepeal *C.struct_CThostFtdcRspRepealField) C.int {
    if t._RtnRepealFromFutureToBankByBank != nil{
        t._RtnRepealFromFutureToBankByBank((*ctp.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
    }
	return 0
}

//export tRtnFromBankToFutureByFuture
func tRtnFromBankToFutureByFuture(pRspTransfer *C.struct_CThostFtdcRspTransferField) C.int {
    if t._RtnFromBankToFutureByFuture != nil{
        t._RtnFromBankToFutureByFuture((*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
    }
	return 0
}

//export tRtnFromFutureToBankByFuture
func tRtnFromFutureToBankByFuture(pRspTransfer *C.struct_CThostFtdcRspTransferField) C.int {
    if t._RtnFromFutureToBankByFuture != nil{
        t._RtnFromFutureToBankByFuture((*ctp.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
    }
	return 0
}

//export tRtnRepealFromBankToFutureByFutureManual
func tRtnRepealFromBankToFutureByFutureManual(pRspRepeal *C.struct_CThostFtdcRspRepealField) C.int {
    if t._RtnRepealFromBankToFutureByFutureManual != nil{
        t._RtnRepealFromBankToFutureByFutureManual((*ctp.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
    }
	return 0
}

//export tRtnRepealFromFutureToBankByFutureManual
func tRtnRepealFromFutureToBankByFutureManual(pRspRepeal *C.struct_CThostFtdcRspRepealField) C.int {
    if t._RtnRepealFromFutureToBankByFutureManual != nil{
        t._RtnRepealFromFutureToBankByFutureManual((*ctp.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
    }
	return 0
}

//export tRtnQueryBankBalanceByFuture
func tRtnQueryBankBalanceByFuture(pNotifyQueryAccount *C.struct_CThostFtdcNotifyQueryAccountField) C.int {
    if t._RtnQueryBankBalanceByFuture != nil{
        t._RtnQueryBankBalanceByFuture((*ctp.CThostFtdcNotifyQueryAccountField)(unsafe.Pointer(pNotifyQueryAccount)))
    }
	return 0
}

//export tErrRtnBankToFutureByFuture
func tErrRtnBankToFutureByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnBankToFutureByFuture != nil{
        t._ErrRtnBankToFutureByFuture((*ctp.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnFutureToBankByFuture
func tErrRtnFutureToBankByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnFutureToBankByFuture != nil{
        t._ErrRtnFutureToBankByFuture((*ctp.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnRepealBankToFutureByFutureManual
func tErrRtnRepealBankToFutureByFutureManual(pReqRepeal *C.struct_CThostFtdcReqRepealField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnRepealBankToFutureByFutureManual != nil{
        t._ErrRtnRepealBankToFutureByFutureManual((*ctp.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnRepealFutureToBankByFutureManual
func tErrRtnRepealFutureToBankByFutureManual(pReqRepeal *C.struct_CThostFtdcReqRepealField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnRepealFutureToBankByFutureManual != nil{
        t._ErrRtnRepealFutureToBankByFutureManual((*ctp.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tErrRtnQueryBankBalanceByFuture
func tErrRtnQueryBankBalanceByFuture(pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField) C.int {
    if t._ErrRtnQueryBankBalanceByFuture != nil{
        t._ErrRtnQueryBankBalanceByFuture((*ctp.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
    }
	return 0
}

//export tRtnRepealFromBankToFutureByFuture
func tRtnRepealFromBankToFutureByFuture(pRspRepeal *C.struct_CThostFtdcRspRepealField) C.int {
    if t._RtnRepealFromBankToFutureByFuture != nil{
        t._RtnRepealFromBankToFutureByFuture((*ctp.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
    }
	return 0
}

//export tRtnRepealFromFutureToBankByFuture
func tRtnRepealFromFutureToBankByFuture(pRspRepeal *C.struct_CThostFtdcRspRepealField) C.int {
    if t._RtnRepealFromFutureToBankByFuture != nil{
        t._RtnRepealFromFutureToBankByFuture((*ctp.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
    }
	return 0
}

//export tRspFromBankToFutureByFuture
func tRspFromBankToFutureByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspFromBankToFutureByFuture != nil{
        t._RspFromBankToFutureByFuture((*ctp.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspFromFutureToBankByFuture
func tRspFromFutureToBankByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspFromFutureToBankByFuture != nil{
        t._RspFromFutureToBankByFuture((*ctp.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQueryBankAccountMoneyByFuture
func tRspQueryBankAccountMoneyByFuture(pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQueryBankAccountMoneyByFuture != nil{
        t._RspQueryBankAccountMoneyByFuture((*ctp.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRtnOpenAccountByBank
func tRtnOpenAccountByBank(pOpenAccount *C.struct_CThostFtdcOpenAccountField) C.int {
    if t._RtnOpenAccountByBank != nil{
        t._RtnOpenAccountByBank((*ctp.CThostFtdcOpenAccountField)(unsafe.Pointer(pOpenAccount)))
    }
	return 0
}

//export tRtnCancelAccountByBank
func tRtnCancelAccountByBank(pCancelAccount *C.struct_CThostFtdcCancelAccountField) C.int {
    if t._RtnCancelAccountByBank != nil{
        t._RtnCancelAccountByBank((*ctp.CThostFtdcCancelAccountField)(unsafe.Pointer(pCancelAccount)))
    }
	return 0
}

//export tRtnChangeAccountByBank
func tRtnChangeAccountByBank(pChangeAccount *C.struct_CThostFtdcChangeAccountField) C.int {
    if t._RtnChangeAccountByBank != nil{
        t._RtnChangeAccountByBank((*ctp.CThostFtdcChangeAccountField)(unsafe.Pointer(pChangeAccount)))
    }
	return 0
}

//export tRspQryClassifiedInstrument
func tRspQryClassifiedInstrument(pInstrument *C.struct_CThostFtdcInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryClassifiedInstrument != nil{
        t._RspQryClassifiedInstrument((*ctp.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryCombPromotionParam
func tRspQryCombPromotionParam(pCombPromotionParam *C.struct_CThostFtdcCombPromotionParamField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryCombPromotionParam != nil{
        t._RspQryCombPromotionParam((*ctp.CThostFtdcCombPromotionParamField)(unsafe.Pointer(pCombPromotionParam)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryRiskSettleInvstPosition
func tRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *C.struct_CThostFtdcRiskSettleInvstPositionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryRiskSettleInvstPosition != nil{
        t._RspQryRiskSettleInvstPosition((*ctp.CThostFtdcRiskSettleInvstPositionField)(unsafe.Pointer(pRiskSettleInvstPosition)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

//export tRspQryRiskSettleProductStatus
func tRspQryRiskSettleProductStatus(pRiskSettleProductStatus *C.struct_CThostFtdcRiskSettleProductStatusField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) C.int {
    if t._RspQryRiskSettleProductStatus != nil{
        t._RspQryRiskSettleProductStatus((*ctp.CThostFtdcRiskSettleProductStatusField)(unsafe.Pointer(pRiskSettleProductStatus)), (*ctp.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
    }
	return 0
}

