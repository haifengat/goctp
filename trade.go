package goctp

/*
#cgo CPPFLAGS: -fPIC -I./CTPv6.6.9_20220922
#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"

#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/lib -Wl,-rpath ${SRCDIR}/lib -l ctptrade -lstdc++

void* CreateFtdcTraderApi(char const*);
void* CreateFtdcTraderSpi();
void* GetVersion();

// 创建TraderApi
void tRelease(void *api);
// 初始化
void tInit(void *api);
// 等待接口线程结束运行
int tJoin(void *api);
// 注册前置机网络地址
void tRegisterFront(void *api, char *pszFrontAddress);
// @remark RegisterNameServer优先于RegisterFront
void tRegisterNameServer(void *api, char *pszNsAddress);
// 注册名字服务器用户信息
void tRegisterFensUserInfo(void *api, struct CThostFtdcFensUserInfoField *pFensUserInfo);
// 注册回调接口
void tRegisterSpi(void *api, void *pSpi);
// 订阅私有流。
void tSubscribePrivateTopic(void *api, int nResumeType);
// 订阅公共流。
void tSubscribePublicTopic(void *api, int nResumeType);
// 客户端认证请求
int tReqAuthenticate(void *api, struct CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID);
// 注册用户终端信息，用于中继服务器多连接模式
int tRegisterUserSystemInfo(void *api, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);
// 上报用户终端信息，用于中继服务器操作员登录模式
int tSubmitUserSystemInfo(void *api, struct CThostFtdcUserSystemInfoField *pUserSystemInfo);
// 用户登录请求
int tReqUserLogin(void *api, struct CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID);
// 登出请求
int tReqUserLogout(void *api, struct CThostFtdcUserLogoutField *pUserLogout, int nRequestID);
// 用户口令更新请求
int tReqUserPasswordUpdate(void *api, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID);
// 资金账户口令更新请求
int tReqTradingAccountPasswordUpdate(void *api, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID);
// 查询用户当前支持的认证模式
int tReqUserAuthMethod(void *api, struct CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID);
// 用户发出获取图形验证码请求
int tReqGenUserCaptcha(void *api, struct CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID);
// 用户发出获取短信验证码请求
int tReqGenUserText(void *api, struct CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID);
// 用户发出带有图片验证码的登陆请求
int tReqUserLoginWithCaptcha(void *api, struct CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID);
// 用户发出带有短信验证码的登陆请求
int tReqUserLoginWithText(void *api, struct CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID);
// 用户发出带有动态口令的登陆请求
int tReqUserLoginWithOTP(void *api, struct CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID);
// 报单录入请求
int tReqOrderInsert(void *api, struct CThostFtdcInputOrderField *pInputOrder, int nRequestID);
// 预埋单录入请求
int tReqParkedOrderInsert(void *api, struct CThostFtdcParkedOrderField *pParkedOrder, int nRequestID);
// 预埋撤单录入请求
int tReqParkedOrderAction(void *api, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID);
// 报单操作请求
int tReqOrderAction(void *api, struct CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID);
// 查询最大报单数量请求
int tReqQryMaxOrderVolume(void *api, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID);
// 投资者结算结果确认
int tReqSettlementInfoConfirm(void *api, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID);
// 请求删除预埋单
int tReqRemoveParkedOrder(void *api, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID);
// 请求删除预埋撤单
int tReqRemoveParkedOrderAction(void *api, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID);
// 执行宣告录入请求
int tReqExecOrderInsert(void *api, struct CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID);
// 执行宣告操作请求
int tReqExecOrderAction(void *api, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID);
// 询价录入请求
int tReqForQuoteInsert(void *api, struct CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID);
// 报价录入请求
int tReqQuoteInsert(void *api, struct CThostFtdcInputQuoteField *pInputQuote, int nRequestID);
// 报价操作请求
int tReqQuoteAction(void *api, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID);
// 批量报单操作请求
int tReqBatchOrderAction(void *api, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID);
// 期权自对冲录入请求
int tReqOptionSelfCloseInsert(void *api, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID);
// 期权自对冲操作请求
int tReqOptionSelfCloseAction(void *api, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID);
// 申请组合录入请求
int tReqCombActionInsert(void *api, struct CThostFtdcInputCombActionField *pInputCombAction, int nRequestID);
// 请求查询报单
int tReqQryOrder(void *api, struct CThostFtdcQryOrderField *pQryOrder, int nRequestID);
// 请求查询成交
int tReqQryTrade(void *api, struct CThostFtdcQryTradeField *pQryTrade, int nRequestID);
// 请求查询投资者持仓
int tReqQryInvestorPosition(void *api, struct CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID);
// 请求查询资金账户
int tReqQryTradingAccount(void *api, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);
// 请求查询投资者
int tReqQryInvestor(void *api, struct CThostFtdcQryInvestorField *pQryInvestor, int nRequestID);
// 请求查询交易编码
int tReqQryTradingCode(void *api, struct CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID);
// 请求查询合约保证金率
int tReqQryInstrumentMarginRate(void *api, struct CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID);
// 请求查询合约手续费率
int tReqQryInstrumentCommissionRate(void *api, struct CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID);
// 请求查询交易所
int tReqQryExchange(void *api, struct CThostFtdcQryExchangeField *pQryExchange, int nRequestID);
// 请求查询产品
int tReqQryProduct(void *api, struct CThostFtdcQryProductField *pQryProduct, int nRequestID);
// 请求查询合约
int tReqQryInstrument(void *api, struct CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID);
// 请求查询行情
int tReqQryDepthMarketData(void *api, struct CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID);
// 请求查询交易员报盘机
int tReqQryTraderOffer(void *api, struct CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID);
// 请求查询投资者结算结果
int tReqQrySettlementInfo(void *api, struct CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID);
// 请求查询转帐银行
int tReqQryTransferBank(void *api, struct CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID);
// 请求查询投资者持仓明细
int tReqQryInvestorPositionDetail(void *api, struct CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID);
// 请求查询客户通知
int tReqQryNotice(void *api, struct CThostFtdcQryNoticeField *pQryNotice, int nRequestID);
// 请求查询结算信息确认
int tReqQrySettlementInfoConfirm(void *api, struct CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID);
// 请求查询投资者持仓明细
int tReqQryInvestorPositionCombineDetail(void *api, struct CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID);
// 请求查询保证金监管系统经纪公司资金账户密钥
int tReqQryCFMMCTradingAccountKey(void *api, struct CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID);
// 请求查询仓单折抵信息
int tReqQryEWarrantOffset(void *api, struct CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID);
// 请求查询投资者品种/跨品种保证金
int tReqQryInvestorProductGroupMargin(void *api, struct CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID);
// 请求查询交易所保证金率
int tReqQryExchangeMarginRate(void *api, struct CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID);
// 请求查询交易所调整保证金率
int tReqQryExchangeMarginRateAdjust(void *api, struct CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID);
// 请求查询汇率
int tReqQryExchangeRate(void *api, struct CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID);
// 请求查询二级代理操作员银期权限
int tReqQrySecAgentACIDMap(void *api, struct CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID);
// 请求查询产品报价汇率
int tReqQryProductExchRate(void *api, struct CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID);
// 请求查询产品组
int tReqQryProductGroup(void *api, struct CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID);
// 请求查询做市商合约手续费率
int tReqQryMMInstrumentCommissionRate(void *api, struct CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID);
// 请求查询做市商期权合约手续费
int tReqQryMMOptionInstrCommRate(void *api, struct CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID);
// 请求查询报单手续费
int tReqQryInstrumentOrderCommRate(void *api, struct CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID);
// 请求查询资金账户
int tReqQrySecAgentTradingAccount(void *api, struct CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID);
// 请求查询二级代理商资金校验模式
int tReqQrySecAgentCheckMode(void *api, struct CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID);
// 请求查询二级代理商信息
int tReqQrySecAgentTradeInfo(void *api, struct CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID);
// 请求查询期权交易成本
int tReqQryOptionInstrTradeCost(void *api, struct CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID);
// 请求查询期权合约手续费
int tReqQryOptionInstrCommRate(void *api, struct CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID);
// 请求查询执行宣告
int tReqQryExecOrder(void *api, struct CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID);
// 请求查询询价
int tReqQryForQuote(void *api, struct CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID);
// 请求查询报价
int tReqQryQuote(void *api, struct CThostFtdcQryQuoteField *pQryQuote, int nRequestID);
// 请求查询期权自对冲
int tReqQryOptionSelfClose(void *api, struct CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID);
// 请求查询投资单元
int tReqQryInvestUnit(void *api, struct CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID);
// 请求查询组合合约安全系数
int tReqQryCombInstrumentGuard(void *api, struct CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID);
// 请求查询申请组合
int tReqQryCombAction(void *api, struct CThostFtdcQryCombActionField *pQryCombAction, int nRequestID);
// 请求查询转帐流水
int tReqQryTransferSerial(void *api, struct CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID);
// 请求查询银期签约关系
int tReqQryAccountregister(void *api, struct CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID);
// 请求查询签约银行
int tReqQryContractBank(void *api, struct CThostFtdcQryContractBankField *pQryContractBank, int nRequestID);
// 请求查询预埋单
int tReqQryParkedOrder(void *api, struct CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID);
// 请求查询预埋撤单
int tReqQryParkedOrderAction(void *api, struct CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID);
// 请求查询交易通知
int tReqQryTradingNotice(void *api, struct CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID);
// 请求查询经纪公司交易参数
int tReqQryBrokerTradingParams(void *api, struct CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID);
// 请求查询经纪公司交易算法
int tReqQryBrokerTradingAlgos(void *api, struct CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID);
// 请求查询监控中心用户令牌
int tReqQueryCFMMCTradingAccountToken(void *api, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID);
// 期货发起银行资金转期货请求
int tReqFromBankToFutureByFuture(void *api, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);
// 期货发起期货资金转银行请求
int tReqFromFutureToBankByFuture(void *api, struct CThostFtdcReqTransferField *pReqTransfer, int nRequestID);
// 期货发起查询银行余额请求
int tReqQueryBankAccountMoneyByFuture(void *api, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID);
// 请求查询分类合约
int tReqQryClassifiedInstrument(void *api, struct CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID);
// 请求组合优惠比例
int tReqQryCombPromotionParam(void *api, struct CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID);
// 投资者风险结算持仓查询
int tReqQryRiskSettleInvstPosition(void *api, struct CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID);
// 风险结算产品查询
int tReqQryRiskSettleProductStatus(void *api, struct CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID);
// SPBM期货合约参数查询
int tReqQrySPBMFutureParameter(void *api, struct CThostFtdcQrySPBMFutureParameterField *pQrySPBMFutureParameter, int nRequestID);
// SPBM期权合约参数查询
int tReqQrySPBMOptionParameter(void *api, struct CThostFtdcQrySPBMOptionParameterField *pQrySPBMOptionParameter, int nRequestID);
// SPBM品种内对锁仓折扣参数查询
int tReqQrySPBMIntraParameter(void *api, struct CThostFtdcQrySPBMIntraParameterField *pQrySPBMIntraParameter, int nRequestID);
// SPBM跨品种抵扣参数查询
int tReqQrySPBMInterParameter(void *api, struct CThostFtdcQrySPBMInterParameterField *pQrySPBMInterParameter, int nRequestID);
// SPBM组合保证金套餐查询
int tReqQrySPBMPortfDefinition(void *api, struct CThostFtdcQrySPBMPortfDefinitionField *pQrySPBMPortfDefinition, int nRequestID);
// 投资者SPBM套餐选择查询
int tReqQrySPBMInvestorPortfDef(void *api, struct CThostFtdcQrySPBMInvestorPortfDefField *pQrySPBMInvestorPortfDef, int nRequestID);
// 投资者新型组合保证金系数查询
int tReqQryInvestorPortfMarginRatio(void *api, struct CThostFtdcQryInvestorPortfMarginRatioField *pQryInvestorPortfMarginRatio, int nRequestID);
// 投资者产品SPBM明细查询
int tReqQryInvestorProdSPBMDetail(void *api, struct CThostFtdcQryInvestorProdSPBMDetailField *pQryInvestorProdSPBMDetail, int nRequestID);

// //////////////////////////////////////////////////////////////////////
void tSetOnFrontConnected(void *, void *);
void OnFrontConnected(void*);
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
void tSetOnFrontDisconnected(void *, void *);
void OnFrontDisconnected(void*, int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
void tSetOnHeartBeatWarning(void *, void *);
void OnHeartBeatWarning(void*, int nTimeLapse);
// 客户端认证响应
void tSetOnRspAuthenticate(void *, void *);
void OnRspAuthenticate(void*, struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登录请求响应
void tSetOnRspUserLogin(void *, void *);
void OnRspUserLogin(void*, struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登出请求响应
void tSetOnRspUserLogout(void *, void *);
void OnRspUserLogout(void*, struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 用户口令更新请求响应
void tSetOnRspUserPasswordUpdate(void *, void *);
void OnRspUserPasswordUpdate(void*, struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 资金账户口令更新请求响应
void tSetOnRspTradingAccountPasswordUpdate(void *, void *);
void OnRspTradingAccountPasswordUpdate(void*, struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询用户当前支持的认证模式的回复
void tSetOnRspUserAuthMethod(void *, void *);
void OnRspUserAuthMethod(void*, struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 获取图形验证码请求的回复
void tSetOnRspGenUserCaptcha(void *, void *);
void OnRspGenUserCaptcha(void*, struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 获取短信验证码请求的回复
void tSetOnRspGenUserText(void *, void *);
void OnRspGenUserText(void*, struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单录入请求响应
void tSetOnRspOrderInsert(void *, void *);
void OnRspOrderInsert(void*, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 预埋单录入请求响应
void tSetOnRspParkedOrderInsert(void *, void *);
void OnRspParkedOrderInsert(void*, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 预埋撤单录入请求响应
void tSetOnRspParkedOrderAction(void *, void *);
void OnRspParkedOrderAction(void*, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单操作请求响应
void tSetOnRspOrderAction(void *, void *);
void OnRspOrderAction(void*, struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询最大报单数量响应
void tSetOnRspQryMaxOrderVolume(void *, void *);
void OnRspQryMaxOrderVolume(void*, struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者结算结果确认响应
void tSetOnRspSettlementInfoConfirm(void *, void *);
void OnRspSettlementInfoConfirm(void*, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 删除预埋单响应
void tSetOnRspRemoveParkedOrder(void *, void *);
void OnRspRemoveParkedOrder(void*, struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 删除预埋撤单响应
void tSetOnRspRemoveParkedOrderAction(void *, void *);
void OnRspRemoveParkedOrderAction(void*, struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 执行宣告录入请求响应
void tSetOnRspExecOrderInsert(void *, void *);
void OnRspExecOrderInsert(void*, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 执行宣告操作请求响应
void tSetOnRspExecOrderAction(void *, void *);
void OnRspExecOrderAction(void*, struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 询价录入请求响应
void tSetOnRspForQuoteInsert(void *, void *);
void OnRspForQuoteInsert(void*, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报价录入请求响应
void tSetOnRspQuoteInsert(void *, void *);
void OnRspQuoteInsert(void*, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报价操作请求响应
void tSetOnRspQuoteAction(void *, void *);
void OnRspQuoteAction(void*, struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 批量报单操作请求响应
void tSetOnRspBatchOrderAction(void *, void *);
void OnRspBatchOrderAction(void*, struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期权自对冲录入请求响应
void tSetOnRspOptionSelfCloseInsert(void *, void *);
void OnRspOptionSelfCloseInsert(void*, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期权自对冲操作请求响应
void tSetOnRspOptionSelfCloseAction(void *, void *);
void OnRspOptionSelfCloseAction(void*, struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 申请组合录入请求响应
void tSetOnRspCombActionInsert(void *, void *);
void OnRspCombActionInsert(void*, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报单响应
void tSetOnRspQryOrder(void *, void *);
void OnRspQryOrder(void*, struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询成交响应
void tSetOnRspQryTrade(void *, void *);
void OnRspQryTrade(void*, struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓响应
void tSetOnRspQryInvestorPosition(void *, void *);
void OnRspQryInvestorPosition(void*, struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询资金账户响应
void tSetOnRspQryTradingAccount(void *, void *);
void OnRspQryTradingAccount(void*, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者响应
void tSetOnRspQryInvestor(void *, void *);
void OnRspQryInvestor(void*, struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易编码响应
void tSetOnRspQryTradingCode(void *, void *);
void OnRspQryTradingCode(void*, struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约保证金率响应
void tSetOnRspQryInstrumentMarginRate(void *, void *);
void OnRspQryInstrumentMarginRate(void*, struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约手续费率响应
void tSetOnRspQryInstrumentCommissionRate(void *, void *);
void OnRspQryInstrumentCommissionRate(void*, struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所响应
void tSetOnRspQryExchange(void *, void *);
void OnRspQryExchange(void*, struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品响应
void tSetOnRspQryProduct(void *, void *);
void OnRspQryProduct(void*, struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约响应
void tSetOnRspQryInstrument(void *, void *);
void OnRspQryInstrument(void*, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询行情响应
void tSetOnRspQryDepthMarketData(void *, void *);
void OnRspQryDepthMarketData(void*, struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易员报盘机响应
void tSetOnRspQryTraderOffer(void *, void *);
void OnRspQryTraderOffer(void*, struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者结算结果响应
void tSetOnRspQrySettlementInfo(void *, void *);
void OnRspQrySettlementInfo(void*, struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询转帐银行响应
void tSetOnRspQryTransferBank(void *, void *);
void OnRspQryTransferBank(void*, struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓明细响应
void tSetOnRspQryInvestorPositionDetail(void *, void *);
void OnRspQryInvestorPositionDetail(void*, struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询客户通知响应
void tSetOnRspQryNotice(void *, void *);
void OnRspQryNotice(void*, struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询结算信息确认响应
void tSetOnRspQrySettlementInfoConfirm(void *, void *);
void OnRspQrySettlementInfoConfirm(void*, struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓明细响应
void tSetOnRspQryInvestorPositionCombineDetail(void *, void *);
void OnRspQryInvestorPositionCombineDetail(void*, struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询保证金监管系统经纪公司资金账户密钥响应
void tSetOnRspQryCFMMCTradingAccountKey(void *, void *);
void OnRspQryCFMMCTradingAccountKey(void*, struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询仓单折抵信息响应
void tSetOnRspQryEWarrantOffset(void *, void *);
void OnRspQryEWarrantOffset(void*, struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者品种/跨品种保证金响应
void tSetOnRspQryInvestorProductGroupMargin(void *, void *);
void OnRspQryInvestorProductGroupMargin(void*, struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所保证金率响应
void tSetOnRspQryExchangeMarginRate(void *, void *);
void OnRspQryExchangeMarginRate(void*, struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所调整保证金率响应
void tSetOnRspQryExchangeMarginRateAdjust(void *, void *);
void OnRspQryExchangeMarginRateAdjust(void*, struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询汇率响应
void tSetOnRspQryExchangeRate(void *, void *);
void OnRspQryExchangeRate(void*, struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理操作员银期权限响应
void tSetOnRspQrySecAgentACIDMap(void *, void *);
void OnRspQrySecAgentACIDMap(void*, struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品报价汇率
void tSetOnRspQryProductExchRate(void *, void *);
void OnRspQryProductExchRate(void*, struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品组
void tSetOnRspQryProductGroup(void *, void *);
void OnRspQryProductGroup(void*, struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询做市商合约手续费率响应
void tSetOnRspQryMMInstrumentCommissionRate(void *, void *);
void OnRspQryMMInstrumentCommissionRate(void*, struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询做市商期权合约手续费响应
void tSetOnRspQryMMOptionInstrCommRate(void *, void *);
void OnRspQryMMOptionInstrCommRate(void*, struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报单手续费响应
void tSetOnRspQryInstrumentOrderCommRate(void *, void *);
void OnRspQryInstrumentOrderCommRate(void*, struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询资金账户响应
void tSetOnRspQrySecAgentTradingAccount(void *, void *);
void OnRspQrySecAgentTradingAccount(void*, struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理商资金校验模式响应
void tSetOnRspQrySecAgentCheckMode(void *, void *);
void OnRspQrySecAgentCheckMode(void*, struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理商信息响应
void tSetOnRspQrySecAgentTradeInfo(void *, void *);
void OnRspQrySecAgentTradeInfo(void*, struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权交易成本响应
void tSetOnRspQryOptionInstrTradeCost(void *, void *);
void OnRspQryOptionInstrTradeCost(void*, struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权合约手续费响应
void tSetOnRspQryOptionInstrCommRate(void *, void *);
void OnRspQryOptionInstrCommRate(void*, struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询执行宣告响应
void tSetOnRspQryExecOrder(void *, void *);
void OnRspQryExecOrder(void*, struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询询价响应
void tSetOnRspQryForQuote(void *, void *);
void OnRspQryForQuote(void*, struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报价响应
void tSetOnRspQryQuote(void *, void *);
void OnRspQryQuote(void*, struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权自对冲响应
void tSetOnRspQryOptionSelfClose(void *, void *);
void OnRspQryOptionSelfClose(void*, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资单元响应
void tSetOnRspQryInvestUnit(void *, void *);
void OnRspQryInvestUnit(void*, struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询组合合约安全系数响应
void tSetOnRspQryCombInstrumentGuard(void *, void *);
void OnRspQryCombInstrumentGuard(void*, struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询申请组合响应
void tSetOnRspQryCombAction(void *, void *);
void OnRspQryCombAction(void*, struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询转帐流水响应
void tSetOnRspQryTransferSerial(void *, void *);
void OnRspQryTransferSerial(void*, struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询银期签约关系响应
void tSetOnRspQryAccountregister(void *, void *);
void OnRspQryAccountregister(void*, struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 错误应答
void tSetOnRspError(void *, void *);
void OnRspError(void*, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单通知
void tSetOnRtnOrder(void *, void *);
void OnRtnOrder(void*, struct CThostFtdcOrderField *pOrder);
// 成交通知
void tSetOnRtnTrade(void *, void *);
void OnRtnTrade(void*, struct CThostFtdcTradeField *pTrade);
// 报单录入错误回报
void tSetOnErrRtnOrderInsert(void *, void *);
void OnErrRtnOrderInsert(void*, struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);
// 报单操作错误回报
void tSetOnErrRtnOrderAction(void *, void *);
void OnErrRtnOrderAction(void*, struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 合约交易状态通知
void tSetOnRtnInstrumentStatus(void *, void *);
void OnRtnInstrumentStatus(void*, struct CThostFtdcInstrumentStatusField *pInstrumentStatus);
// 交易所公告通知
void tSetOnRtnBulletin(void *, void *);
void OnRtnBulletin(void*, struct CThostFtdcBulletinField *pBulletin);
// 交易通知
void tSetOnRtnTradingNotice(void *, void *);
void OnRtnTradingNotice(void*, struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);
// 提示条件单校验错误
void tSetOnRtnErrorConditionalOrder(void *, void *);
void OnRtnErrorConditionalOrder(void*, struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);
// 执行宣告通知
void tSetOnRtnExecOrder(void *, void *);
void OnRtnExecOrder(void*, struct CThostFtdcExecOrderField *pExecOrder);
// 执行宣告录入错误回报
void tSetOnErrRtnExecOrderInsert(void *, void *);
void OnErrRtnExecOrderInsert(void*, struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo);
// 执行宣告操作错误回报
void tSetOnErrRtnExecOrderAction(void *, void *);
void OnErrRtnExecOrderAction(void*, struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 询价录入错误回报
void tSetOnErrRtnForQuoteInsert(void *, void *);
void OnErrRtnForQuoteInsert(void*, struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo);
// 报价通知
void tSetOnRtnQuote(void *, void *);
void OnRtnQuote(void*, struct CThostFtdcQuoteField *pQuote);
// 报价录入错误回报
void tSetOnErrRtnQuoteInsert(void *, void *);
void OnErrRtnQuoteInsert(void*, struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo);
// 报价操作错误回报
void tSetOnErrRtnQuoteAction(void *, void *);
void OnErrRtnQuoteAction(void*, struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo);
// 询价通知
void tSetOnRtnForQuoteRsp(void *, void *);
void OnRtnForQuoteRsp(void*, struct CThostFtdcForQuoteRspField *pForQuoteRsp);
// 保证金监控中心用户令牌
void tSetOnRtnCFMMCTradingAccountToken(void *, void *);
void OnRtnCFMMCTradingAccountToken(void*, struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);
// 批量报单操作错误回报
void tSetOnErrRtnBatchOrderAction(void *, void *);
void OnErrRtnBatchOrderAction(void*, struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 期权自对冲通知
void tSetOnRtnOptionSelfClose(void *, void *);
void OnRtnOptionSelfClose(void*, struct CThostFtdcOptionSelfCloseField *pOptionSelfClose);
// 期权自对冲录入错误回报
void tSetOnErrRtnOptionSelfCloseInsert(void *, void *);
void OnErrRtnOptionSelfCloseInsert(void*, struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo);
// 期权自对冲操作错误回报
void tSetOnErrRtnOptionSelfCloseAction(void *, void *);
void OnErrRtnOptionSelfCloseAction(void*, struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo);
// 申请组合通知
void tSetOnRtnCombAction(void *, void *);
void OnRtnCombAction(void*, struct CThostFtdcCombActionField *pCombAction);
// 申请组合录入错误回报
void tSetOnErrRtnCombActionInsert(void *, void *);
void OnErrRtnCombActionInsert(void*, struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo);
// 请求查询签约银行响应
void tSetOnRspQryContractBank(void *, void *);
void OnRspQryContractBank(void*, struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询预埋单响应
void tSetOnRspQryParkedOrder(void *, void *);
void OnRspQryParkedOrder(void*, struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询预埋撤单响应
void tSetOnRspQryParkedOrderAction(void *, void *);
void OnRspQryParkedOrderAction(void*, struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易通知响应
void tSetOnRspQryTradingNotice(void *, void *);
void OnRspQryTradingNotice(void*, struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询经纪公司交易参数响应
void tSetOnRspQryBrokerTradingParams(void *, void *);
void OnRspQryBrokerTradingParams(void*, struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询经纪公司交易算法响应
void tSetOnRspQryBrokerTradingAlgos(void *, void *);
void OnRspQryBrokerTradingAlgos(void*, struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询监控中心用户令牌
void tSetOnRspQueryCFMMCTradingAccountToken(void *, void *);
void OnRspQueryCFMMCTradingAccountToken(void*, struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 银行发起银行资金转期货通知
void tSetOnRtnFromBankToFutureByBank(void *, void *);
void OnRtnFromBankToFutureByBank(void*, struct CThostFtdcRspTransferField *pRspTransfer);
// 银行发起期货资金转银行通知
void tSetOnRtnFromFutureToBankByBank(void *, void *);
void OnRtnFromFutureToBankByBank(void*, struct CThostFtdcRspTransferField *pRspTransfer);
// 银行发起冲正银行转期货通知
void tSetOnRtnRepealFromBankToFutureByBank(void *, void *);
void OnRtnRepealFromBankToFutureByBank(void*, struct CThostFtdcRspRepealField *pRspRepeal);
// 银行发起冲正期货转银行通知
void tSetOnRtnRepealFromFutureToBankByBank(void *, void *);
void OnRtnRepealFromFutureToBankByBank(void*, struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起银行资金转期货通知
void tSetOnRtnFromBankToFutureByFuture(void *, void *);
void OnRtnFromBankToFutureByFuture(void*, struct CThostFtdcRspTransferField *pRspTransfer);
// 期货发起期货资金转银行通知
void tSetOnRtnFromFutureToBankByFuture(void *, void *);
void OnRtnFromFutureToBankByFuture(void*, struct CThostFtdcRspTransferField *pRspTransfer);
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromBankToFutureByFutureManual(void *, void *);
void OnRtnRepealFromBankToFutureByFutureManual(void*, struct CThostFtdcRspRepealField *pRspRepeal);
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromFutureToBankByFutureManual(void *, void *);
void OnRtnRepealFromFutureToBankByFutureManual(void*, struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起查询银行余额通知
void tSetOnRtnQueryBankBalanceByFuture(void *, void *);
void OnRtnQueryBankBalanceByFuture(void*, struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);
// 期货发起银行资金转期货错误回报
void tSetOnErrRtnBankToFutureByFuture(void *, void *);
void OnErrRtnBankToFutureByFuture(void*, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起期货资金转银行错误回报
void tSetOnErrRtnFutureToBankByFuture(void *, void *);
void OnErrRtnFutureToBankByFuture(void*, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);
// 系统运行时期货端手工发起冲正银行转期货错误回报
void tSetOnErrRtnRepealBankToFutureByFutureManual(void *, void *);
void OnErrRtnRepealBankToFutureByFutureManual(void*, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);
// 系统运行时期货端手工发起冲正期货转银行错误回报
void tSetOnErrRtnRepealFutureToBankByFutureManual(void *, void *);
void OnErrRtnRepealFutureToBankByFutureManual(void*, struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起查询银行余额错误回报
void tSetOnErrRtnQueryBankBalanceByFuture(void *, void *);
void OnErrRtnQueryBankBalanceByFuture(void*, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromBankToFutureByFuture(void *, void *);
void OnRtnRepealFromBankToFutureByFuture(void*, struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromFutureToBankByFuture(void *, void *);
void OnRtnRepealFromFutureToBankByFuture(void*, struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起银行资金转期货应答
void tSetOnRspFromBankToFutureByFuture(void *, void *);
void OnRspFromBankToFutureByFuture(void*, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期货发起期货资金转银行应答
void tSetOnRspFromFutureToBankByFuture(void *, void *);
void OnRspFromFutureToBankByFuture(void*, struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期货发起查询银行余额应答
void tSetOnRspQueryBankAccountMoneyByFuture(void *, void *);
void OnRspQueryBankAccountMoneyByFuture(void*, struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 银行发起银期开户通知
void tSetOnRtnOpenAccountByBank(void *, void *);
void OnRtnOpenAccountByBank(void*, struct CThostFtdcOpenAccountField *pOpenAccount);
// 银行发起银期销户通知
void tSetOnRtnCancelAccountByBank(void *, void *);
void OnRtnCancelAccountByBank(void*, struct CThostFtdcCancelAccountField *pCancelAccount);
// 银行发起变更银行账号通知
void tSetOnRtnChangeAccountByBank(void *, void *);
void OnRtnChangeAccountByBank(void*, struct CThostFtdcChangeAccountField *pChangeAccount);
// 请求查询分类合约响应
void tSetOnRspQryClassifiedInstrument(void *, void *);
void OnRspQryClassifiedInstrument(void*, struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求组合优惠比例响应
void tSetOnRspQryCombPromotionParam(void *, void *);
void OnRspQryCombPromotionParam(void*, struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者风险结算持仓查询响应
void tSetOnRspQryRiskSettleInvstPosition(void *, void *);
void OnRspQryRiskSettleInvstPosition(void*, struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 风险结算产品查询响应
void tSetOnRspQryRiskSettleProductStatus(void *, void *);
void OnRspQryRiskSettleProductStatus(void*, struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// SPBM期货合约参数查询响应
void tSetOnRspQrySPBMFutureParameter(void *, void *);
void OnRspQrySPBMFutureParameter(void*, struct CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// SPBM期权合约参数查询响应
void tSetOnRspQrySPBMOptionParameter(void *, void *);
void OnRspQrySPBMOptionParameter(void*, struct CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// SPBM品种内对锁仓折扣参数查询响应
void tSetOnRspQrySPBMIntraParameter(void *, void *);
void OnRspQrySPBMIntraParameter(void*, struct CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// SPBM跨品种抵扣参数查询响应
void tSetOnRspQrySPBMInterParameter(void *, void *);
void OnRspQrySPBMInterParameter(void*, struct CThostFtdcSPBMInterParameterField *pSPBMInterParameter, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// SPBM组合保证金套餐查询响应
void tSetOnRspQrySPBMPortfDefinition(void *, void *);
void OnRspQrySPBMPortfDefinition(void*, struct CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者SPBM套餐选择查询响应
void tSetOnRspQrySPBMInvestorPortfDef(void *, void *);
void OnRspQrySPBMInvestorPortfDef(void*, struct CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者新型组合保证金系数查询响应
void tSetOnRspQryInvestorPortfMarginRatio(void *, void *);
void OnRspQryInvestorPortfMarginRatio(void*, struct CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者产品SPBM明细查询响应
void tSetOnRspQryInvestorProdSPBMDetail(void *, void *);
void OnRspQryInvestorProdSPBMDetail(void*, struct CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);


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
	
	// //////////////////////////////////////////////////////////////////////
	OnFrontConnected func()
	// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	OnFrontDisconnected func(nReason int)
	// 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnHeartBeatWarning func(nTimeLapse int)
	// 客户端认证响应
	OnRspAuthenticate func(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 登录请求响应
	OnRspUserLogin func(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 登出请求响应
	OnRspUserLogout func(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 用户口令更新请求响应
	OnRspUserPasswordUpdate func(pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 资金账户口令更新请求响应
	OnRspTradingAccountPasswordUpdate func(pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 查询用户当前支持的认证模式的回复
	OnRspUserAuthMethod func(pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 获取图形验证码请求的回复
	OnRspGenUserCaptcha func(pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 获取短信验证码请求的回复
	OnRspGenUserText func(pRspGenUserText *CThostFtdcRspGenUserTextField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 报单录入请求响应
	OnRspOrderInsert func(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 预埋单录入请求响应
	OnRspParkedOrderInsert func(pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 预埋撤单录入请求响应
	OnRspParkedOrderAction func(pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 报单操作请求响应
	OnRspOrderAction func(pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 查询最大报单数量响应
	OnRspQryMaxOrderVolume func(pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 投资者结算结果确认响应
	OnRspSettlementInfoConfirm func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 删除预埋单响应
	OnRspRemoveParkedOrder func(pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 删除预埋撤单响应
	OnRspRemoveParkedOrderAction func(pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 执行宣告录入请求响应
	OnRspExecOrderInsert func(pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 执行宣告操作请求响应
	OnRspExecOrderAction func(pInputExecOrderAction *CThostFtdcInputExecOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 询价录入请求响应
	OnRspForQuoteInsert func(pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 报价录入请求响应
	OnRspQuoteInsert func(pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 报价操作请求响应
	OnRspQuoteAction func(pInputQuoteAction *CThostFtdcInputQuoteActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 批量报单操作请求响应
	OnRspBatchOrderAction func(pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 期权自对冲录入请求响应
	OnRspOptionSelfCloseInsert func(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 期权自对冲操作请求响应
	OnRspOptionSelfCloseAction func(pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 申请组合录入请求响应
	OnRspCombActionInsert func(pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询报单响应
	OnRspQryOrder func(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询成交响应
	OnRspQryTrade func(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资者持仓响应
	OnRspQryInvestorPosition func(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询资金账户响应
	OnRspQryTradingAccount func(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资者响应
	OnRspQryInvestor func(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询交易编码响应
	OnRspQryTradingCode func(pTradingCode *CThostFtdcTradingCodeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询合约保证金率响应
	OnRspQryInstrumentMarginRate func(pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询合约手续费率响应
	OnRspQryInstrumentCommissionRate func(pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询交易所响应
	OnRspQryExchange func(pExchange *CThostFtdcExchangeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询产品响应
	OnRspQryProduct func(pProduct *CThostFtdcProductField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询合约响应
	OnRspQryInstrument func(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询行情响应
	OnRspQryDepthMarketData func(pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询交易员报盘机响应
	OnRspQryTraderOffer func(pTraderOffer *CThostFtdcTraderOfferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资者结算结果响应
	OnRspQrySettlementInfo func(pSettlementInfo *CThostFtdcSettlementInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询转帐银行响应
	OnRspQryTransferBank func(pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资者持仓明细响应
	OnRspQryInvestorPositionDetail func(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询客户通知响应
	OnRspQryNotice func(pNotice *CThostFtdcNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询结算信息确认响应
	OnRspQrySettlementInfoConfirm func(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资者持仓明细响应
	OnRspQryInvestorPositionCombineDetail func(pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 查询保证金监管系统经纪公司资金账户密钥响应
	OnRspQryCFMMCTradingAccountKey func(pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询仓单折抵信息响应
	OnRspQryEWarrantOffset func(pEWarrantOffset *CThostFtdcEWarrantOffsetField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资者品种/跨品种保证金响应
	OnRspQryInvestorProductGroupMargin func(pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询交易所保证金率响应
	OnRspQryExchangeMarginRate func(pExchangeMarginRate *CThostFtdcExchangeMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询交易所调整保证金率响应
	OnRspQryExchangeMarginRateAdjust func(pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询汇率响应
	OnRspQryExchangeRate func(pExchangeRate *CThostFtdcExchangeRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询二级代理操作员银期权限响应
	OnRspQrySecAgentACIDMap func(pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询产品报价汇率
	OnRspQryProductExchRate func(pProductExchRate *CThostFtdcProductExchRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询产品组
	OnRspQryProductGroup func(pProductGroup *CThostFtdcProductGroupField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询做市商合约手续费率响应
	OnRspQryMMInstrumentCommissionRate func(pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询做市商期权合约手续费响应
	OnRspQryMMOptionInstrCommRate func(pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询报单手续费响应
	OnRspQryInstrumentOrderCommRate func(pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询资金账户响应
	OnRspQrySecAgentTradingAccount func(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询二级代理商资金校验模式响应
	OnRspQrySecAgentCheckMode func(pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询二级代理商信息响应
	OnRspQrySecAgentTradeInfo func(pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询期权交易成本响应
	OnRspQryOptionInstrTradeCost func(pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询期权合约手续费响应
	OnRspQryOptionInstrCommRate func(pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询执行宣告响应
	OnRspQryExecOrder func(pExecOrder *CThostFtdcExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询询价响应
	OnRspQryForQuote func(pForQuote *CThostFtdcForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询报价响应
	OnRspQryQuote func(pQuote *CThostFtdcQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询期权自对冲响应
	OnRspQryOptionSelfClose func(pOptionSelfClose *CThostFtdcOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询投资单元响应
	OnRspQryInvestUnit func(pInvestUnit *CThostFtdcInvestUnitField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询组合合约安全系数响应
	OnRspQryCombInstrumentGuard func(pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询申请组合响应
	OnRspQryCombAction func(pCombAction *CThostFtdcCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询转帐流水响应
	OnRspQryTransferSerial func(pTransferSerial *CThostFtdcTransferSerialField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询银期签约关系响应
	OnRspQryAccountregister func(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 错误应答
	OnRspError func(pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 报单通知
	OnRtnOrder func(pOrder *CThostFtdcOrderField)
	// 成交通知
	OnRtnTrade func(pTrade *CThostFtdcTradeField)
	// 报单录入错误回报
	OnErrRtnOrderInsert func(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField)
	// 报单操作错误回报
	OnErrRtnOrderAction func(pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField)
	// 合约交易状态通知
	OnRtnInstrumentStatus func(pInstrumentStatus *CThostFtdcInstrumentStatusField)
	// 交易所公告通知
	OnRtnBulletin func(pBulletin *CThostFtdcBulletinField)
	// 交易通知
	OnRtnTradingNotice func(pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField)
	// 提示条件单校验错误
	OnRtnErrorConditionalOrder func(pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField)
	// 执行宣告通知
	OnRtnExecOrder func(pExecOrder *CThostFtdcExecOrderField)
	// 执行宣告录入错误回报
	OnErrRtnExecOrderInsert func(pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField)
	// 执行宣告操作错误回报
	OnErrRtnExecOrderAction func(pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField)
	// 询价录入错误回报
	OnErrRtnForQuoteInsert func(pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField)
	// 报价通知
	OnRtnQuote func(pQuote *CThostFtdcQuoteField)
	// 报价录入错误回报
	OnErrRtnQuoteInsert func(pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField)
	// 报价操作错误回报
	OnErrRtnQuoteAction func(pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField)
	// 询价通知
	OnRtnForQuoteRsp func(pForQuoteRsp *CThostFtdcForQuoteRspField)
	// 保证金监控中心用户令牌
	OnRtnCFMMCTradingAccountToken func(pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField)
	// 批量报单操作错误回报
	OnErrRtnBatchOrderAction func(pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField)
	// 期权自对冲通知
	OnRtnOptionSelfClose func(pOptionSelfClose *CThostFtdcOptionSelfCloseField)
	// 期权自对冲录入错误回报
	OnErrRtnOptionSelfCloseInsert func(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField)
	// 期权自对冲操作错误回报
	OnErrRtnOptionSelfCloseAction func(pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField)
	// 申请组合通知
	OnRtnCombAction func(pCombAction *CThostFtdcCombActionField)
	// 申请组合录入错误回报
	OnErrRtnCombActionInsert func(pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField)
	// 请求查询签约银行响应
	OnRspQryContractBank func(pContractBank *CThostFtdcContractBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询预埋单响应
	OnRspQryParkedOrder func(pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询预埋撤单响应
	OnRspQryParkedOrderAction func(pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询交易通知响应
	OnRspQryTradingNotice func(pTradingNotice *CThostFtdcTradingNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询经纪公司交易参数响应
	OnRspQryBrokerTradingParams func(pBrokerTradingParams *CThostFtdcBrokerTradingParamsField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询经纪公司交易算法响应
	OnRspQryBrokerTradingAlgos func(pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求查询监控中心用户令牌
	OnRspQueryCFMMCTradingAccountToken func(pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 银行发起银行资金转期货通知
	OnRtnFromBankToFutureByBank func(pRspTransfer *CThostFtdcRspTransferField)
	// 银行发起期货资金转银行通知
	OnRtnFromFutureToBankByBank func(pRspTransfer *CThostFtdcRspTransferField)
	// 银行发起冲正银行转期货通知
	OnRtnRepealFromBankToFutureByBank func(pRspRepeal *CThostFtdcRspRepealField)
	// 银行发起冲正期货转银行通知
	OnRtnRepealFromFutureToBankByBank func(pRspRepeal *CThostFtdcRspRepealField)
	// 期货发起银行资金转期货通知
	OnRtnFromBankToFutureByFuture func(pRspTransfer *CThostFtdcRspTransferField)
	// 期货发起期货资金转银行通知
	OnRtnFromFutureToBankByFuture func(pRspTransfer *CThostFtdcRspTransferField)
	// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFutureManual func(pRspRepeal *CThostFtdcRspRepealField)
	// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFutureManual func(pRspRepeal *CThostFtdcRspRepealField)
	// 期货发起查询银行余额通知
	OnRtnQueryBankBalanceByFuture func(pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField)
	// 期货发起银行资金转期货错误回报
	OnErrRtnBankToFutureByFuture func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField)
	// 期货发起期货资金转银行错误回报
	OnErrRtnFutureToBankByFuture func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField)
	// 系统运行时期货端手工发起冲正银行转期货错误回报
	OnErrRtnRepealBankToFutureByFutureManual func(pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField)
	// 系统运行时期货端手工发起冲正期货转银行错误回报
	OnErrRtnRepealFutureToBankByFutureManual func(pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField)
	// 期货发起查询银行余额错误回报
	OnErrRtnQueryBankBalanceByFuture func(pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField)
	// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFuture func(pRspRepeal *CThostFtdcRspRepealField)
	// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFuture func(pRspRepeal *CThostFtdcRspRepealField)
	// 期货发起银行资金转期货应答
	OnRspFromBankToFutureByFuture func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 期货发起期货资金转银行应答
	OnRspFromFutureToBankByFuture func(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 期货发起查询银行余额应答
	OnRspQueryBankAccountMoneyByFuture func(pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 银行发起银期开户通知
	OnRtnOpenAccountByBank func(pOpenAccount *CThostFtdcOpenAccountField)
	// 银行发起银期销户通知
	OnRtnCancelAccountByBank func(pCancelAccount *CThostFtdcCancelAccountField)
	// 银行发起变更银行账号通知
	OnRtnChangeAccountByBank func(pChangeAccount *CThostFtdcChangeAccountField)
	// 请求查询分类合约响应
	OnRspQryClassifiedInstrument func(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 请求组合优惠比例响应
	OnRspQryCombPromotionParam func(pCombPromotionParam *CThostFtdcCombPromotionParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 投资者风险结算持仓查询响应
	OnRspQryRiskSettleInvstPosition func(pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 风险结算产品查询响应
	OnRspQryRiskSettleProductStatus func(pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// SPBM期货合约参数查询响应
	OnRspQrySPBMFutureParameter func(pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// SPBM期权合约参数查询响应
	OnRspQrySPBMOptionParameter func(pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// SPBM品种内对锁仓折扣参数查询响应
	OnRspQrySPBMIntraParameter func(pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// SPBM跨品种抵扣参数查询响应
	OnRspQrySPBMInterParameter func(pSPBMInterParameter *CThostFtdcSPBMInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// SPBM组合保证金套餐查询响应
	OnRspQrySPBMPortfDefinition func(pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 投资者SPBM套餐选择查询响应
	OnRspQrySPBMInvestorPortfDef func(pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 投资者新型组合保证金系数查询响应
	OnRspQryInvestorPortfMarginRatio func(pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
	// 投资者产品SPBM明细查询响应
	OnRspQryInvestorProdSPBMDetail func(pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
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

    C.tSetOnFrontConnected(t.spi, C.OnFrontConnected) // //////////////////////////////////////////////////////////////////////
    C.tSetOnFrontDisconnected(t.spi, C.OnFrontDisconnected) // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    C.tSetOnHeartBeatWarning(t.spi, C.OnHeartBeatWarning) // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    C.tSetOnRspAuthenticate(t.spi, C.OnRspAuthenticate) // 客户端认证响应
    C.tSetOnRspUserLogin(t.spi, C.OnRspUserLogin) // 登录请求响应
    C.tSetOnRspUserLogout(t.spi, C.OnRspUserLogout) // 登出请求响应
    C.tSetOnRspUserPasswordUpdate(t.spi, C.OnRspUserPasswordUpdate) // 用户口令更新请求响应
    C.tSetOnRspTradingAccountPasswordUpdate(t.spi, C.OnRspTradingAccountPasswordUpdate) // 资金账户口令更新请求响应
    C.tSetOnRspUserAuthMethod(t.spi, C.OnRspUserAuthMethod) // 查询用户当前支持的认证模式的回复
    C.tSetOnRspGenUserCaptcha(t.spi, C.OnRspGenUserCaptcha) // 获取图形验证码请求的回复
    C.tSetOnRspGenUserText(t.spi, C.OnRspGenUserText) // 获取短信验证码请求的回复
    C.tSetOnRspOrderInsert(t.spi, C.OnRspOrderInsert) // 报单录入请求响应
    C.tSetOnRspParkedOrderInsert(t.spi, C.OnRspParkedOrderInsert) // 预埋单录入请求响应
    C.tSetOnRspParkedOrderAction(t.spi, C.OnRspParkedOrderAction) // 预埋撤单录入请求响应
    C.tSetOnRspOrderAction(t.spi, C.OnRspOrderAction) // 报单操作请求响应
    C.tSetOnRspQryMaxOrderVolume(t.spi, C.OnRspQryMaxOrderVolume) // 查询最大报单数量响应
    C.tSetOnRspSettlementInfoConfirm(t.spi, C.OnRspSettlementInfoConfirm) // 投资者结算结果确认响应
    C.tSetOnRspRemoveParkedOrder(t.spi, C.OnRspRemoveParkedOrder) // 删除预埋单响应
    C.tSetOnRspRemoveParkedOrderAction(t.spi, C.OnRspRemoveParkedOrderAction) // 删除预埋撤单响应
    C.tSetOnRspExecOrderInsert(t.spi, C.OnRspExecOrderInsert) // 执行宣告录入请求响应
    C.tSetOnRspExecOrderAction(t.spi, C.OnRspExecOrderAction) // 执行宣告操作请求响应
    C.tSetOnRspForQuoteInsert(t.spi, C.OnRspForQuoteInsert) // 询价录入请求响应
    C.tSetOnRspQuoteInsert(t.spi, C.OnRspQuoteInsert) // 报价录入请求响应
    C.tSetOnRspQuoteAction(t.spi, C.OnRspQuoteAction) // 报价操作请求响应
    C.tSetOnRspBatchOrderAction(t.spi, C.OnRspBatchOrderAction) // 批量报单操作请求响应
    C.tSetOnRspOptionSelfCloseInsert(t.spi, C.OnRspOptionSelfCloseInsert) // 期权自对冲录入请求响应
    C.tSetOnRspOptionSelfCloseAction(t.spi, C.OnRspOptionSelfCloseAction) // 期权自对冲操作请求响应
    C.tSetOnRspCombActionInsert(t.spi, C.OnRspCombActionInsert) // 申请组合录入请求响应
    C.tSetOnRspQryOrder(t.spi, C.OnRspQryOrder) // 请求查询报单响应
    C.tSetOnRspQryTrade(t.spi, C.OnRspQryTrade) // 请求查询成交响应
    C.tSetOnRspQryInvestorPosition(t.spi, C.OnRspQryInvestorPosition) // 请求查询投资者持仓响应
    C.tSetOnRspQryTradingAccount(t.spi, C.OnRspQryTradingAccount) // 请求查询资金账户响应
    C.tSetOnRspQryInvestor(t.spi, C.OnRspQryInvestor) // 请求查询投资者响应
    C.tSetOnRspQryTradingCode(t.spi, C.OnRspQryTradingCode) // 请求查询交易编码响应
    C.tSetOnRspQryInstrumentMarginRate(t.spi, C.OnRspQryInstrumentMarginRate) // 请求查询合约保证金率响应
    C.tSetOnRspQryInstrumentCommissionRate(t.spi, C.OnRspQryInstrumentCommissionRate) // 请求查询合约手续费率响应
    C.tSetOnRspQryExchange(t.spi, C.OnRspQryExchange) // 请求查询交易所响应
    C.tSetOnRspQryProduct(t.spi, C.OnRspQryProduct) // 请求查询产品响应
    C.tSetOnRspQryInstrument(t.spi, C.OnRspQryInstrument) // 请求查询合约响应
    C.tSetOnRspQryDepthMarketData(t.spi, C.OnRspQryDepthMarketData) // 请求查询行情响应
    C.tSetOnRspQryTraderOffer(t.spi, C.OnRspQryTraderOffer) // 请求查询交易员报盘机响应
    C.tSetOnRspQrySettlementInfo(t.spi, C.OnRspQrySettlementInfo) // 请求查询投资者结算结果响应
    C.tSetOnRspQryTransferBank(t.spi, C.OnRspQryTransferBank) // 请求查询转帐银行响应
    C.tSetOnRspQryInvestorPositionDetail(t.spi, C.OnRspQryInvestorPositionDetail) // 请求查询投资者持仓明细响应
    C.tSetOnRspQryNotice(t.spi, C.OnRspQryNotice) // 请求查询客户通知响应
    C.tSetOnRspQrySettlementInfoConfirm(t.spi, C.OnRspQrySettlementInfoConfirm) // 请求查询结算信息确认响应
    C.tSetOnRspQryInvestorPositionCombineDetail(t.spi, C.OnRspQryInvestorPositionCombineDetail) // 请求查询投资者持仓明细响应
    C.tSetOnRspQryCFMMCTradingAccountKey(t.spi, C.OnRspQryCFMMCTradingAccountKey) // 查询保证金监管系统经纪公司资金账户密钥响应
    C.tSetOnRspQryEWarrantOffset(t.spi, C.OnRspQryEWarrantOffset) // 请求查询仓单折抵信息响应
    C.tSetOnRspQryInvestorProductGroupMargin(t.spi, C.OnRspQryInvestorProductGroupMargin) // 请求查询投资者品种/跨品种保证金响应
    C.tSetOnRspQryExchangeMarginRate(t.spi, C.OnRspQryExchangeMarginRate) // 请求查询交易所保证金率响应
    C.tSetOnRspQryExchangeMarginRateAdjust(t.spi, C.OnRspQryExchangeMarginRateAdjust) // 请求查询交易所调整保证金率响应
    C.tSetOnRspQryExchangeRate(t.spi, C.OnRspQryExchangeRate) // 请求查询汇率响应
    C.tSetOnRspQrySecAgentACIDMap(t.spi, C.OnRspQrySecAgentACIDMap) // 请求查询二级代理操作员银期权限响应
    C.tSetOnRspQryProductExchRate(t.spi, C.OnRspQryProductExchRate) // 请求查询产品报价汇率
    C.tSetOnRspQryProductGroup(t.spi, C.OnRspQryProductGroup) // 请求查询产品组
    C.tSetOnRspQryMMInstrumentCommissionRate(t.spi, C.OnRspQryMMInstrumentCommissionRate) // 请求查询做市商合约手续费率响应
    C.tSetOnRspQryMMOptionInstrCommRate(t.spi, C.OnRspQryMMOptionInstrCommRate) // 请求查询做市商期权合约手续费响应
    C.tSetOnRspQryInstrumentOrderCommRate(t.spi, C.OnRspQryInstrumentOrderCommRate) // 请求查询报单手续费响应
    C.tSetOnRspQrySecAgentTradingAccount(t.spi, C.OnRspQrySecAgentTradingAccount) // 请求查询资金账户响应
    C.tSetOnRspQrySecAgentCheckMode(t.spi, C.OnRspQrySecAgentCheckMode) // 请求查询二级代理商资金校验模式响应
    C.tSetOnRspQrySecAgentTradeInfo(t.spi, C.OnRspQrySecAgentTradeInfo) // 请求查询二级代理商信息响应
    C.tSetOnRspQryOptionInstrTradeCost(t.spi, C.OnRspQryOptionInstrTradeCost) // 请求查询期权交易成本响应
    C.tSetOnRspQryOptionInstrCommRate(t.spi, C.OnRspQryOptionInstrCommRate) // 请求查询期权合约手续费响应
    C.tSetOnRspQryExecOrder(t.spi, C.OnRspQryExecOrder) // 请求查询执行宣告响应
    C.tSetOnRspQryForQuote(t.spi, C.OnRspQryForQuote) // 请求查询询价响应
    C.tSetOnRspQryQuote(t.spi, C.OnRspQryQuote) // 请求查询报价响应
    C.tSetOnRspQryOptionSelfClose(t.spi, C.OnRspQryOptionSelfClose) // 请求查询期权自对冲响应
    C.tSetOnRspQryInvestUnit(t.spi, C.OnRspQryInvestUnit) // 请求查询投资单元响应
    C.tSetOnRspQryCombInstrumentGuard(t.spi, C.OnRspQryCombInstrumentGuard) // 请求查询组合合约安全系数响应
    C.tSetOnRspQryCombAction(t.spi, C.OnRspQryCombAction) // 请求查询申请组合响应
    C.tSetOnRspQryTransferSerial(t.spi, C.OnRspQryTransferSerial) // 请求查询转帐流水响应
    C.tSetOnRspQryAccountregister(t.spi, C.OnRspQryAccountregister) // 请求查询银期签约关系响应
    C.tSetOnRspError(t.spi, C.OnRspError) // 错误应答
    C.tSetOnRtnOrder(t.spi, C.OnRtnOrder) // 报单通知
    C.tSetOnRtnTrade(t.spi, C.OnRtnTrade) // 成交通知
    C.tSetOnErrRtnOrderInsert(t.spi, C.OnErrRtnOrderInsert) // 报单录入错误回报
    C.tSetOnErrRtnOrderAction(t.spi, C.OnErrRtnOrderAction) // 报单操作错误回报
    C.tSetOnRtnInstrumentStatus(t.spi, C.OnRtnInstrumentStatus) // 合约交易状态通知
    C.tSetOnRtnBulletin(t.spi, C.OnRtnBulletin) // 交易所公告通知
    C.tSetOnRtnTradingNotice(t.spi, C.OnRtnTradingNotice) // 交易通知
    C.tSetOnRtnErrorConditionalOrder(t.spi, C.OnRtnErrorConditionalOrder) // 提示条件单校验错误
    C.tSetOnRtnExecOrder(t.spi, C.OnRtnExecOrder) // 执行宣告通知
    C.tSetOnErrRtnExecOrderInsert(t.spi, C.OnErrRtnExecOrderInsert) // 执行宣告录入错误回报
    C.tSetOnErrRtnExecOrderAction(t.spi, C.OnErrRtnExecOrderAction) // 执行宣告操作错误回报
    C.tSetOnErrRtnForQuoteInsert(t.spi, C.OnErrRtnForQuoteInsert) // 询价录入错误回报
    C.tSetOnRtnQuote(t.spi, C.OnRtnQuote) // 报价通知
    C.tSetOnErrRtnQuoteInsert(t.spi, C.OnErrRtnQuoteInsert) // 报价录入错误回报
    C.tSetOnErrRtnQuoteAction(t.spi, C.OnErrRtnQuoteAction) // 报价操作错误回报
    C.tSetOnRtnForQuoteRsp(t.spi, C.OnRtnForQuoteRsp) // 询价通知
    C.tSetOnRtnCFMMCTradingAccountToken(t.spi, C.OnRtnCFMMCTradingAccountToken) // 保证金监控中心用户令牌
    C.tSetOnErrRtnBatchOrderAction(t.spi, C.OnErrRtnBatchOrderAction) // 批量报单操作错误回报
    C.tSetOnRtnOptionSelfClose(t.spi, C.OnRtnOptionSelfClose) // 期权自对冲通知
    C.tSetOnErrRtnOptionSelfCloseInsert(t.spi, C.OnErrRtnOptionSelfCloseInsert) // 期权自对冲录入错误回报
    C.tSetOnErrRtnOptionSelfCloseAction(t.spi, C.OnErrRtnOptionSelfCloseAction) // 期权自对冲操作错误回报
    C.tSetOnRtnCombAction(t.spi, C.OnRtnCombAction) // 申请组合通知
    C.tSetOnErrRtnCombActionInsert(t.spi, C.OnErrRtnCombActionInsert) // 申请组合录入错误回报
    C.tSetOnRspQryContractBank(t.spi, C.OnRspQryContractBank) // 请求查询签约银行响应
    C.tSetOnRspQryParkedOrder(t.spi, C.OnRspQryParkedOrder) // 请求查询预埋单响应
    C.tSetOnRspQryParkedOrderAction(t.spi, C.OnRspQryParkedOrderAction) // 请求查询预埋撤单响应
    C.tSetOnRspQryTradingNotice(t.spi, C.OnRspQryTradingNotice) // 请求查询交易通知响应
    C.tSetOnRspQryBrokerTradingParams(t.spi, C.OnRspQryBrokerTradingParams) // 请求查询经纪公司交易参数响应
    C.tSetOnRspQryBrokerTradingAlgos(t.spi, C.OnRspQryBrokerTradingAlgos) // 请求查询经纪公司交易算法响应
    C.tSetOnRspQueryCFMMCTradingAccountToken(t.spi, C.OnRspQueryCFMMCTradingAccountToken) // 请求查询监控中心用户令牌
    C.tSetOnRtnFromBankToFutureByBank(t.spi, C.OnRtnFromBankToFutureByBank) // 银行发起银行资金转期货通知
    C.tSetOnRtnFromFutureToBankByBank(t.spi, C.OnRtnFromFutureToBankByBank) // 银行发起期货资金转银行通知
    C.tSetOnRtnRepealFromBankToFutureByBank(t.spi, C.OnRtnRepealFromBankToFutureByBank) // 银行发起冲正银行转期货通知
    C.tSetOnRtnRepealFromFutureToBankByBank(t.spi, C.OnRtnRepealFromFutureToBankByBank) // 银行发起冲正期货转银行通知
    C.tSetOnRtnFromBankToFutureByFuture(t.spi, C.OnRtnFromBankToFutureByFuture) // 期货发起银行资金转期货通知
    C.tSetOnRtnFromFutureToBankByFuture(t.spi, C.OnRtnFromFutureToBankByFuture) // 期货发起期货资金转银行通知
    C.tSetOnRtnRepealFromBankToFutureByFutureManual(t.spi, C.OnRtnRepealFromBankToFutureByFutureManual) // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    C.tSetOnRtnRepealFromFutureToBankByFutureManual(t.spi, C.OnRtnRepealFromFutureToBankByFutureManual) // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    C.tSetOnRtnQueryBankBalanceByFuture(t.spi, C.OnRtnQueryBankBalanceByFuture) // 期货发起查询银行余额通知
    C.tSetOnErrRtnBankToFutureByFuture(t.spi, C.OnErrRtnBankToFutureByFuture) // 期货发起银行资金转期货错误回报
    C.tSetOnErrRtnFutureToBankByFuture(t.spi, C.OnErrRtnFutureToBankByFuture) // 期货发起期货资金转银行错误回报
    C.tSetOnErrRtnRepealBankToFutureByFutureManual(t.spi, C.OnErrRtnRepealBankToFutureByFutureManual) // 系统运行时期货端手工发起冲正银行转期货错误回报
    C.tSetOnErrRtnRepealFutureToBankByFutureManual(t.spi, C.OnErrRtnRepealFutureToBankByFutureManual) // 系统运行时期货端手工发起冲正期货转银行错误回报
    C.tSetOnErrRtnQueryBankBalanceByFuture(t.spi, C.OnErrRtnQueryBankBalanceByFuture) // 期货发起查询银行余额错误回报
    C.tSetOnRtnRepealFromBankToFutureByFuture(t.spi, C.OnRtnRepealFromBankToFutureByFuture) // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    C.tSetOnRtnRepealFromFutureToBankByFuture(t.spi, C.OnRtnRepealFromFutureToBankByFuture) // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    C.tSetOnRspFromBankToFutureByFuture(t.spi, C.OnRspFromBankToFutureByFuture) // 期货发起银行资金转期货应答
    C.tSetOnRspFromFutureToBankByFuture(t.spi, C.OnRspFromFutureToBankByFuture) // 期货发起期货资金转银行应答
    C.tSetOnRspQueryBankAccountMoneyByFuture(t.spi, C.OnRspQueryBankAccountMoneyByFuture) // 期货发起查询银行余额应答
    C.tSetOnRtnOpenAccountByBank(t.spi, C.OnRtnOpenAccountByBank) // 银行发起银期开户通知
    C.tSetOnRtnCancelAccountByBank(t.spi, C.OnRtnCancelAccountByBank) // 银行发起银期销户通知
    C.tSetOnRtnChangeAccountByBank(t.spi, C.OnRtnChangeAccountByBank) // 银行发起变更银行账号通知
    C.tSetOnRspQryClassifiedInstrument(t.spi, C.OnRspQryClassifiedInstrument) // 请求查询分类合约响应
    C.tSetOnRspQryCombPromotionParam(t.spi, C.OnRspQryCombPromotionParam) // 请求组合优惠比例响应
    C.tSetOnRspQryRiskSettleInvstPosition(t.spi, C.OnRspQryRiskSettleInvstPosition) // 投资者风险结算持仓查询响应
    C.tSetOnRspQryRiskSettleProductStatus(t.spi, C.OnRspQryRiskSettleProductStatus) // 风险结算产品查询响应
    C.tSetOnRspQrySPBMFutureParameter(t.spi, C.OnRspQrySPBMFutureParameter) // SPBM期货合约参数查询响应
    C.tSetOnRspQrySPBMOptionParameter(t.spi, C.OnRspQrySPBMOptionParameter) // SPBM期权合约参数查询响应
    C.tSetOnRspQrySPBMIntraParameter(t.spi, C.OnRspQrySPBMIntraParameter) // SPBM品种内对锁仓折扣参数查询响应
    C.tSetOnRspQrySPBMInterParameter(t.spi, C.OnRspQrySPBMInterParameter) // SPBM跨品种抵扣参数查询响应
    C.tSetOnRspQrySPBMPortfDefinition(t.spi, C.OnRspQrySPBMPortfDefinition) // SPBM组合保证金套餐查询响应
    C.tSetOnRspQrySPBMInvestorPortfDef(t.spi, C.OnRspQrySPBMInvestorPortfDef) // 投资者SPBM套餐选择查询响应
    C.tSetOnRspQryInvestorPortfMarginRatio(t.spi, C.OnRspQryInvestorPortfMarginRatio) // 投资者新型组合保证金系数查询响应
    C.tSetOnRspQryInvestorProdSPBMDetail(t.spi, C.OnRspQryInvestorProdSPBMDetail) // 投资者产品SPBM明细查询响应
    
	Trades[t.spi] = t
    return t
}

//export OnFrontConnected
func OnFrontConnected(spi unsafe.Pointer) {
	t := Trades[spi]
	if t.OnFrontConnected == nil {
		fmt.Println("OnFrontConnected")
	} else {
		t.OnFrontConnected()
	}
}
//export OnFrontDisconnected
func OnFrontDisconnected(spi unsafe.Pointer, nReason C.int) {
	t := Trades[spi]
	if t.OnFrontDisconnected == nil {
		fmt.Println("OnFrontDisconnected")
	} else {
		t.OnFrontDisconnected(int(nReason))
	}
}
//export OnHeartBeatWarning
func OnHeartBeatWarning(spi unsafe.Pointer, nTimeLapse C.int) {
	t := Trades[spi]
	if t.OnHeartBeatWarning == nil {
		fmt.Println("OnHeartBeatWarning")
	} else {
		t.OnHeartBeatWarning(int(nTimeLapse))
	}
}
//export OnRspAuthenticate
func OnRspAuthenticate(spi unsafe.Pointer, pRspAuthenticateField *C.struct_CThostFtdcRspAuthenticateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspAuthenticate == nil {
		fmt.Println("OnRspAuthenticate")
	} else {
		t.OnRspAuthenticate((*CThostFtdcRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspUserLogin
func OnRspUserLogin(spi unsafe.Pointer, pRspUserLogin *C.struct_CThostFtdcRspUserLoginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspUserLogin == nil {
		fmt.Println("OnRspUserLogin")
	} else {
		t.OnRspUserLogin((*CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspUserLogout
func OnRspUserLogout(spi unsafe.Pointer, pUserLogout *C.struct_CThostFtdcUserLogoutField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspUserLogout == nil {
		fmt.Println("OnRspUserLogout")
	} else {
		t.OnRspUserLogout((*CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspUserPasswordUpdate
func OnRspUserPasswordUpdate(spi unsafe.Pointer, pUserPasswordUpdate *C.struct_CThostFtdcUserPasswordUpdateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspUserPasswordUpdate == nil {
		fmt.Println("OnRspUserPasswordUpdate")
	} else {
		t.OnRspUserPasswordUpdate((*CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspTradingAccountPasswordUpdate
func OnRspTradingAccountPasswordUpdate(spi unsafe.Pointer, pTradingAccountPasswordUpdate *C.struct_CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspTradingAccountPasswordUpdate == nil {
		fmt.Println("OnRspTradingAccountPasswordUpdate")
	} else {
		t.OnRspTradingAccountPasswordUpdate((*CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspUserAuthMethod
func OnRspUserAuthMethod(spi unsafe.Pointer, pRspUserAuthMethod *C.struct_CThostFtdcRspUserAuthMethodField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspUserAuthMethod == nil {
		fmt.Println("OnRspUserAuthMethod")
	} else {
		t.OnRspUserAuthMethod((*CThostFtdcRspUserAuthMethodField)(unsafe.Pointer(pRspUserAuthMethod)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspGenUserCaptcha
func OnRspGenUserCaptcha(spi unsafe.Pointer, pRspGenUserCaptcha *C.struct_CThostFtdcRspGenUserCaptchaField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspGenUserCaptcha == nil {
		fmt.Println("OnRspGenUserCaptcha")
	} else {
		t.OnRspGenUserCaptcha((*CThostFtdcRspGenUserCaptchaField)(unsafe.Pointer(pRspGenUserCaptcha)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspGenUserText
func OnRspGenUserText(spi unsafe.Pointer, pRspGenUserText *C.struct_CThostFtdcRspGenUserTextField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspGenUserText == nil {
		fmt.Println("OnRspGenUserText")
	} else {
		t.OnRspGenUserText((*CThostFtdcRspGenUserTextField)(unsafe.Pointer(pRspGenUserText)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspOrderInsert
func OnRspOrderInsert(spi unsafe.Pointer, pInputOrder *C.struct_CThostFtdcInputOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspOrderInsert == nil {
		fmt.Println("OnRspOrderInsert")
	} else {
		t.OnRspOrderInsert((*CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspParkedOrderInsert
func OnRspParkedOrderInsert(spi unsafe.Pointer, pParkedOrder *C.struct_CThostFtdcParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspParkedOrderInsert == nil {
		fmt.Println("OnRspParkedOrderInsert")
	} else {
		t.OnRspParkedOrderInsert((*CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspParkedOrderAction
func OnRspParkedOrderAction(spi unsafe.Pointer, pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspParkedOrderAction == nil {
		fmt.Println("OnRspParkedOrderAction")
	} else {
		t.OnRspParkedOrderAction((*CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspOrderAction
func OnRspOrderAction(spi unsafe.Pointer, pInputOrderAction *C.struct_CThostFtdcInputOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspOrderAction == nil {
		fmt.Println("OnRspOrderAction")
	} else {
		t.OnRspOrderAction((*CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryMaxOrderVolume
func OnRspQryMaxOrderVolume(spi unsafe.Pointer, pQryMaxOrderVolume *C.struct_CThostFtdcQryMaxOrderVolumeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryMaxOrderVolume == nil {
		fmt.Println("OnRspQryMaxOrderVolume")
	} else {
		t.OnRspQryMaxOrderVolume((*CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspSettlementInfoConfirm
func OnRspSettlementInfoConfirm(spi unsafe.Pointer, pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspSettlementInfoConfirm == nil {
		fmt.Println("OnRspSettlementInfoConfirm")
	} else {
		t.OnRspSettlementInfoConfirm((*CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspRemoveParkedOrder
func OnRspRemoveParkedOrder(spi unsafe.Pointer, pRemoveParkedOrder *C.struct_CThostFtdcRemoveParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspRemoveParkedOrder == nil {
		fmt.Println("OnRspRemoveParkedOrder")
	} else {
		t.OnRspRemoveParkedOrder((*CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspRemoveParkedOrderAction
func OnRspRemoveParkedOrderAction(spi unsafe.Pointer, pRemoveParkedOrderAction *C.struct_CThostFtdcRemoveParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspRemoveParkedOrderAction == nil {
		fmt.Println("OnRspRemoveParkedOrderAction")
	} else {
		t.OnRspRemoveParkedOrderAction((*CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspExecOrderInsert
func OnRspExecOrderInsert(spi unsafe.Pointer, pInputExecOrder *C.struct_CThostFtdcInputExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspExecOrderInsert == nil {
		fmt.Println("OnRspExecOrderInsert")
	} else {
		t.OnRspExecOrderInsert((*CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspExecOrderAction
func OnRspExecOrderAction(spi unsafe.Pointer, pInputExecOrderAction *C.struct_CThostFtdcInputExecOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspExecOrderAction == nil {
		fmt.Println("OnRspExecOrderAction")
	} else {
		t.OnRspExecOrderAction((*CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspForQuoteInsert
func OnRspForQuoteInsert(spi unsafe.Pointer, pInputForQuote *C.struct_CThostFtdcInputForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspForQuoteInsert == nil {
		fmt.Println("OnRspForQuoteInsert")
	} else {
		t.OnRspForQuoteInsert((*CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQuoteInsert
func OnRspQuoteInsert(spi unsafe.Pointer, pInputQuote *C.struct_CThostFtdcInputQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQuoteInsert == nil {
		fmt.Println("OnRspQuoteInsert")
	} else {
		t.OnRspQuoteInsert((*CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQuoteAction
func OnRspQuoteAction(spi unsafe.Pointer, pInputQuoteAction *C.struct_CThostFtdcInputQuoteActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQuoteAction == nil {
		fmt.Println("OnRspQuoteAction")
	} else {
		t.OnRspQuoteAction((*CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspBatchOrderAction
func OnRspBatchOrderAction(spi unsafe.Pointer, pInputBatchOrderAction *C.struct_CThostFtdcInputBatchOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspBatchOrderAction == nil {
		fmt.Println("OnRspBatchOrderAction")
	} else {
		t.OnRspBatchOrderAction((*CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspOptionSelfCloseInsert
func OnRspOptionSelfCloseInsert(spi unsafe.Pointer, pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspOptionSelfCloseInsert == nil {
		fmt.Println("OnRspOptionSelfCloseInsert")
	} else {
		t.OnRspOptionSelfCloseInsert((*CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspOptionSelfCloseAction
func OnRspOptionSelfCloseAction(spi unsafe.Pointer, pInputOptionSelfCloseAction *C.struct_CThostFtdcInputOptionSelfCloseActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspOptionSelfCloseAction == nil {
		fmt.Println("OnRspOptionSelfCloseAction")
	} else {
		t.OnRspOptionSelfCloseAction((*CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspCombActionInsert
func OnRspCombActionInsert(spi unsafe.Pointer, pInputCombAction *C.struct_CThostFtdcInputCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspCombActionInsert == nil {
		fmt.Println("OnRspCombActionInsert")
	} else {
		t.OnRspCombActionInsert((*CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryOrder
func OnRspQryOrder(spi unsafe.Pointer, pOrder *C.struct_CThostFtdcOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryOrder == nil {
		fmt.Println("OnRspQryOrder")
	} else {
		t.OnRspQryOrder((*CThostFtdcOrderField)(unsafe.Pointer(pOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTrade
func OnRspQryTrade(spi unsafe.Pointer, pTrade *C.struct_CThostFtdcTradeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTrade == nil {
		fmt.Println("OnRspQryTrade")
	} else {
		t.OnRspQryTrade((*CThostFtdcTradeField)(unsafe.Pointer(pTrade)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestorPosition
func OnRspQryInvestorPosition(spi unsafe.Pointer, pInvestorPosition *C.struct_CThostFtdcInvestorPositionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestorPosition == nil {
		fmt.Println("OnRspQryInvestorPosition")
	} else {
		t.OnRspQryInvestorPosition((*CThostFtdcInvestorPositionField)(unsafe.Pointer(pInvestorPosition)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTradingAccount
func OnRspQryTradingAccount(spi unsafe.Pointer, pTradingAccount *C.struct_CThostFtdcTradingAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTradingAccount == nil {
		fmt.Println("OnRspQryTradingAccount")
	} else {
		t.OnRspQryTradingAccount((*CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestor
func OnRspQryInvestor(spi unsafe.Pointer, pInvestor *C.struct_CThostFtdcInvestorField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestor == nil {
		fmt.Println("OnRspQryInvestor")
	} else {
		t.OnRspQryInvestor((*CThostFtdcInvestorField)(unsafe.Pointer(pInvestor)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTradingCode
func OnRspQryTradingCode(spi unsafe.Pointer, pTradingCode *C.struct_CThostFtdcTradingCodeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTradingCode == nil {
		fmt.Println("OnRspQryTradingCode")
	} else {
		t.OnRspQryTradingCode((*CThostFtdcTradingCodeField)(unsafe.Pointer(pTradingCode)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInstrumentMarginRate
func OnRspQryInstrumentMarginRate(spi unsafe.Pointer, pInstrumentMarginRate *C.struct_CThostFtdcInstrumentMarginRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInstrumentMarginRate == nil {
		fmt.Println("OnRspQryInstrumentMarginRate")
	} else {
		t.OnRspQryInstrumentMarginRate((*CThostFtdcInstrumentMarginRateField)(unsafe.Pointer(pInstrumentMarginRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInstrumentCommissionRate
func OnRspQryInstrumentCommissionRate(spi unsafe.Pointer, pInstrumentCommissionRate *C.struct_CThostFtdcInstrumentCommissionRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInstrumentCommissionRate == nil {
		fmt.Println("OnRspQryInstrumentCommissionRate")
	} else {
		t.OnRspQryInstrumentCommissionRate((*CThostFtdcInstrumentCommissionRateField)(unsafe.Pointer(pInstrumentCommissionRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryExchange
func OnRspQryExchange(spi unsafe.Pointer, pExchange *C.struct_CThostFtdcExchangeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryExchange == nil {
		fmt.Println("OnRspQryExchange")
	} else {
		t.OnRspQryExchange((*CThostFtdcExchangeField)(unsafe.Pointer(pExchange)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryProduct
func OnRspQryProduct(spi unsafe.Pointer, pProduct *C.struct_CThostFtdcProductField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryProduct == nil {
		fmt.Println("OnRspQryProduct")
	} else {
		t.OnRspQryProduct((*CThostFtdcProductField)(unsafe.Pointer(pProduct)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInstrument
func OnRspQryInstrument(spi unsafe.Pointer, pInstrument *C.struct_CThostFtdcInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInstrument == nil {
		fmt.Println("OnRspQryInstrument")
	} else {
		t.OnRspQryInstrument((*CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryDepthMarketData
func OnRspQryDepthMarketData(spi unsafe.Pointer, pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryDepthMarketData == nil {
		fmt.Println("OnRspQryDepthMarketData")
	} else {
		t.OnRspQryDepthMarketData((*CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTraderOffer
func OnRspQryTraderOffer(spi unsafe.Pointer, pTraderOffer *C.struct_CThostFtdcTraderOfferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTraderOffer == nil {
		fmt.Println("OnRspQryTraderOffer")
	} else {
		t.OnRspQryTraderOffer((*CThostFtdcTraderOfferField)(unsafe.Pointer(pTraderOffer)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySettlementInfo
func OnRspQrySettlementInfo(spi unsafe.Pointer, pSettlementInfo *C.struct_CThostFtdcSettlementInfoField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySettlementInfo == nil {
		fmt.Println("OnRspQrySettlementInfo")
	} else {
		t.OnRspQrySettlementInfo((*CThostFtdcSettlementInfoField)(unsafe.Pointer(pSettlementInfo)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTransferBank
func OnRspQryTransferBank(spi unsafe.Pointer, pTransferBank *C.struct_CThostFtdcTransferBankField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTransferBank == nil {
		fmt.Println("OnRspQryTransferBank")
	} else {
		t.OnRspQryTransferBank((*CThostFtdcTransferBankField)(unsafe.Pointer(pTransferBank)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestorPositionDetail
func OnRspQryInvestorPositionDetail(spi unsafe.Pointer, pInvestorPositionDetail *C.struct_CThostFtdcInvestorPositionDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestorPositionDetail == nil {
		fmt.Println("OnRspQryInvestorPositionDetail")
	} else {
		t.OnRspQryInvestorPositionDetail((*CThostFtdcInvestorPositionDetailField)(unsafe.Pointer(pInvestorPositionDetail)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryNotice
func OnRspQryNotice(spi unsafe.Pointer, pNotice *C.struct_CThostFtdcNoticeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryNotice == nil {
		fmt.Println("OnRspQryNotice")
	} else {
		t.OnRspQryNotice((*CThostFtdcNoticeField)(unsafe.Pointer(pNotice)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySettlementInfoConfirm
func OnRspQrySettlementInfoConfirm(spi unsafe.Pointer, pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySettlementInfoConfirm == nil {
		fmt.Println("OnRspQrySettlementInfoConfirm")
	} else {
		t.OnRspQrySettlementInfoConfirm((*CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestorPositionCombineDetail
func OnRspQryInvestorPositionCombineDetail(spi unsafe.Pointer, pInvestorPositionCombineDetail *C.struct_CThostFtdcInvestorPositionCombineDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestorPositionCombineDetail == nil {
		fmt.Println("OnRspQryInvestorPositionCombineDetail")
	} else {
		t.OnRspQryInvestorPositionCombineDetail((*CThostFtdcInvestorPositionCombineDetailField)(unsafe.Pointer(pInvestorPositionCombineDetail)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryCFMMCTradingAccountKey
func OnRspQryCFMMCTradingAccountKey(spi unsafe.Pointer, pCFMMCTradingAccountKey *C.struct_CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryCFMMCTradingAccountKey == nil {
		fmt.Println("OnRspQryCFMMCTradingAccountKey")
	} else {
		t.OnRspQryCFMMCTradingAccountKey((*CThostFtdcCFMMCTradingAccountKeyField)(unsafe.Pointer(pCFMMCTradingAccountKey)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryEWarrantOffset
func OnRspQryEWarrantOffset(spi unsafe.Pointer, pEWarrantOffset *C.struct_CThostFtdcEWarrantOffsetField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryEWarrantOffset == nil {
		fmt.Println("OnRspQryEWarrantOffset")
	} else {
		t.OnRspQryEWarrantOffset((*CThostFtdcEWarrantOffsetField)(unsafe.Pointer(pEWarrantOffset)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestorProductGroupMargin
func OnRspQryInvestorProductGroupMargin(spi unsafe.Pointer, pInvestorProductGroupMargin *C.struct_CThostFtdcInvestorProductGroupMarginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestorProductGroupMargin == nil {
		fmt.Println("OnRspQryInvestorProductGroupMargin")
	} else {
		t.OnRspQryInvestorProductGroupMargin((*CThostFtdcInvestorProductGroupMarginField)(unsafe.Pointer(pInvestorProductGroupMargin)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryExchangeMarginRate
func OnRspQryExchangeMarginRate(spi unsafe.Pointer, pExchangeMarginRate *C.struct_CThostFtdcExchangeMarginRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryExchangeMarginRate == nil {
		fmt.Println("OnRspQryExchangeMarginRate")
	} else {
		t.OnRspQryExchangeMarginRate((*CThostFtdcExchangeMarginRateField)(unsafe.Pointer(pExchangeMarginRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryExchangeMarginRateAdjust
func OnRspQryExchangeMarginRateAdjust(spi unsafe.Pointer, pExchangeMarginRateAdjust *C.struct_CThostFtdcExchangeMarginRateAdjustField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryExchangeMarginRateAdjust == nil {
		fmt.Println("OnRspQryExchangeMarginRateAdjust")
	} else {
		t.OnRspQryExchangeMarginRateAdjust((*CThostFtdcExchangeMarginRateAdjustField)(unsafe.Pointer(pExchangeMarginRateAdjust)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryExchangeRate
func OnRspQryExchangeRate(spi unsafe.Pointer, pExchangeRate *C.struct_CThostFtdcExchangeRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryExchangeRate == nil {
		fmt.Println("OnRspQryExchangeRate")
	} else {
		t.OnRspQryExchangeRate((*CThostFtdcExchangeRateField)(unsafe.Pointer(pExchangeRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySecAgentACIDMap
func OnRspQrySecAgentACIDMap(spi unsafe.Pointer, pSecAgentACIDMap *C.struct_CThostFtdcSecAgentACIDMapField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySecAgentACIDMap == nil {
		fmt.Println("OnRspQrySecAgentACIDMap")
	} else {
		t.OnRspQrySecAgentACIDMap((*CThostFtdcSecAgentACIDMapField)(unsafe.Pointer(pSecAgentACIDMap)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryProductExchRate
func OnRspQryProductExchRate(spi unsafe.Pointer, pProductExchRate *C.struct_CThostFtdcProductExchRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryProductExchRate == nil {
		fmt.Println("OnRspQryProductExchRate")
	} else {
		t.OnRspQryProductExchRate((*CThostFtdcProductExchRateField)(unsafe.Pointer(pProductExchRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryProductGroup
func OnRspQryProductGroup(spi unsafe.Pointer, pProductGroup *C.struct_CThostFtdcProductGroupField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryProductGroup == nil {
		fmt.Println("OnRspQryProductGroup")
	} else {
		t.OnRspQryProductGroup((*CThostFtdcProductGroupField)(unsafe.Pointer(pProductGroup)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryMMInstrumentCommissionRate
func OnRspQryMMInstrumentCommissionRate(spi unsafe.Pointer, pMMInstrumentCommissionRate *C.struct_CThostFtdcMMInstrumentCommissionRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryMMInstrumentCommissionRate == nil {
		fmt.Println("OnRspQryMMInstrumentCommissionRate")
	} else {
		t.OnRspQryMMInstrumentCommissionRate((*CThostFtdcMMInstrumentCommissionRateField)(unsafe.Pointer(pMMInstrumentCommissionRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryMMOptionInstrCommRate
func OnRspQryMMOptionInstrCommRate(spi unsafe.Pointer, pMMOptionInstrCommRate *C.struct_CThostFtdcMMOptionInstrCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryMMOptionInstrCommRate == nil {
		fmt.Println("OnRspQryMMOptionInstrCommRate")
	} else {
		t.OnRspQryMMOptionInstrCommRate((*CThostFtdcMMOptionInstrCommRateField)(unsafe.Pointer(pMMOptionInstrCommRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInstrumentOrderCommRate
func OnRspQryInstrumentOrderCommRate(spi unsafe.Pointer, pInstrumentOrderCommRate *C.struct_CThostFtdcInstrumentOrderCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInstrumentOrderCommRate == nil {
		fmt.Println("OnRspQryInstrumentOrderCommRate")
	} else {
		t.OnRspQryInstrumentOrderCommRate((*CThostFtdcInstrumentOrderCommRateField)(unsafe.Pointer(pInstrumentOrderCommRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySecAgentTradingAccount
func OnRspQrySecAgentTradingAccount(spi unsafe.Pointer, pTradingAccount *C.struct_CThostFtdcTradingAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySecAgentTradingAccount == nil {
		fmt.Println("OnRspQrySecAgentTradingAccount")
	} else {
		t.OnRspQrySecAgentTradingAccount((*CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySecAgentCheckMode
func OnRspQrySecAgentCheckMode(spi unsafe.Pointer, pSecAgentCheckMode *C.struct_CThostFtdcSecAgentCheckModeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySecAgentCheckMode == nil {
		fmt.Println("OnRspQrySecAgentCheckMode")
	} else {
		t.OnRspQrySecAgentCheckMode((*CThostFtdcSecAgentCheckModeField)(unsafe.Pointer(pSecAgentCheckMode)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySecAgentTradeInfo
func OnRspQrySecAgentTradeInfo(spi unsafe.Pointer, pSecAgentTradeInfo *C.struct_CThostFtdcSecAgentTradeInfoField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySecAgentTradeInfo == nil {
		fmt.Println("OnRspQrySecAgentTradeInfo")
	} else {
		t.OnRspQrySecAgentTradeInfo((*CThostFtdcSecAgentTradeInfoField)(unsafe.Pointer(pSecAgentTradeInfo)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryOptionInstrTradeCost
func OnRspQryOptionInstrTradeCost(spi unsafe.Pointer, pOptionInstrTradeCost *C.struct_CThostFtdcOptionInstrTradeCostField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryOptionInstrTradeCost == nil {
		fmt.Println("OnRspQryOptionInstrTradeCost")
	} else {
		t.OnRspQryOptionInstrTradeCost((*CThostFtdcOptionInstrTradeCostField)(unsafe.Pointer(pOptionInstrTradeCost)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryOptionInstrCommRate
func OnRspQryOptionInstrCommRate(spi unsafe.Pointer, pOptionInstrCommRate *C.struct_CThostFtdcOptionInstrCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryOptionInstrCommRate == nil {
		fmt.Println("OnRspQryOptionInstrCommRate")
	} else {
		t.OnRspQryOptionInstrCommRate((*CThostFtdcOptionInstrCommRateField)(unsafe.Pointer(pOptionInstrCommRate)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryExecOrder
func OnRspQryExecOrder(spi unsafe.Pointer, pExecOrder *C.struct_CThostFtdcExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryExecOrder == nil {
		fmt.Println("OnRspQryExecOrder")
	} else {
		t.OnRspQryExecOrder((*CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryForQuote
func OnRspQryForQuote(spi unsafe.Pointer, pForQuote *C.struct_CThostFtdcForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryForQuote == nil {
		fmt.Println("OnRspQryForQuote")
	} else {
		t.OnRspQryForQuote((*CThostFtdcForQuoteField)(unsafe.Pointer(pForQuote)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryQuote
func OnRspQryQuote(spi unsafe.Pointer, pQuote *C.struct_CThostFtdcQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryQuote == nil {
		fmt.Println("OnRspQryQuote")
	} else {
		t.OnRspQryQuote((*CThostFtdcQuoteField)(unsafe.Pointer(pQuote)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryOptionSelfClose
func OnRspQryOptionSelfClose(spi unsafe.Pointer, pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryOptionSelfClose == nil {
		fmt.Println("OnRspQryOptionSelfClose")
	} else {
		t.OnRspQryOptionSelfClose((*CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestUnit
func OnRspQryInvestUnit(spi unsafe.Pointer, pInvestUnit *C.struct_CThostFtdcInvestUnitField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestUnit == nil {
		fmt.Println("OnRspQryInvestUnit")
	} else {
		t.OnRspQryInvestUnit((*CThostFtdcInvestUnitField)(unsafe.Pointer(pInvestUnit)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryCombInstrumentGuard
func OnRspQryCombInstrumentGuard(spi unsafe.Pointer, pCombInstrumentGuard *C.struct_CThostFtdcCombInstrumentGuardField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryCombInstrumentGuard == nil {
		fmt.Println("OnRspQryCombInstrumentGuard")
	} else {
		t.OnRspQryCombInstrumentGuard((*CThostFtdcCombInstrumentGuardField)(unsafe.Pointer(pCombInstrumentGuard)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryCombAction
func OnRspQryCombAction(spi unsafe.Pointer, pCombAction *C.struct_CThostFtdcCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryCombAction == nil {
		fmt.Println("OnRspQryCombAction")
	} else {
		t.OnRspQryCombAction((*CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTransferSerial
func OnRspQryTransferSerial(spi unsafe.Pointer, pTransferSerial *C.struct_CThostFtdcTransferSerialField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTransferSerial == nil {
		fmt.Println("OnRspQryTransferSerial")
	} else {
		t.OnRspQryTransferSerial((*CThostFtdcTransferSerialField)(unsafe.Pointer(pTransferSerial)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryAccountregister
func OnRspQryAccountregister(spi unsafe.Pointer, pAccountregister *C.struct_CThostFtdcAccountregisterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryAccountregister == nil {
		fmt.Println("OnRspQryAccountregister")
	} else {
		t.OnRspQryAccountregister((*CThostFtdcAccountregisterField)(unsafe.Pointer(pAccountregister)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspError
func OnRspError(spi unsafe.Pointer, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspError == nil {
		fmt.Println("OnRspError")
	} else {
		t.OnRspError((*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRtnOrder
func OnRtnOrder(spi unsafe.Pointer, pOrder *C.struct_CThostFtdcOrderField) {
	t := Trades[spi]
	if t.OnRtnOrder == nil {
		fmt.Println("OnRtnOrder")
	} else {
		t.OnRtnOrder((*CThostFtdcOrderField)(unsafe.Pointer(pOrder)))
	}
}
//export OnRtnTrade
func OnRtnTrade(spi unsafe.Pointer, pTrade *C.struct_CThostFtdcTradeField) {
	t := Trades[spi]
	if t.OnRtnTrade == nil {
		fmt.Println("OnRtnTrade")
	} else {
		t.OnRtnTrade((*CThostFtdcTradeField)(unsafe.Pointer(pTrade)))
	}
}
//export OnErrRtnOrderInsert
func OnErrRtnOrderInsert(spi unsafe.Pointer, pInputOrder *C.struct_CThostFtdcInputOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnOrderInsert == nil {
		fmt.Println("OnErrRtnOrderInsert")
	} else {
		t.OnErrRtnOrderInsert((*CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnOrderAction
func OnErrRtnOrderAction(spi unsafe.Pointer, pOrderAction *C.struct_CThostFtdcOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnOrderAction == nil {
		fmt.Println("OnErrRtnOrderAction")
	} else {
		t.OnErrRtnOrderAction((*CThostFtdcOrderActionField)(unsafe.Pointer(pOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRtnInstrumentStatus
func OnRtnInstrumentStatus(spi unsafe.Pointer, pInstrumentStatus *C.struct_CThostFtdcInstrumentStatusField) {
	t := Trades[spi]
	if t.OnRtnInstrumentStatus == nil {
		fmt.Println("OnRtnInstrumentStatus")
	} else {
		t.OnRtnInstrumentStatus((*CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)))
	}
}
//export OnRtnBulletin
func OnRtnBulletin(spi unsafe.Pointer, pBulletin *C.struct_CThostFtdcBulletinField) {
	t := Trades[spi]
	if t.OnRtnBulletin == nil {
		fmt.Println("OnRtnBulletin")
	} else {
		t.OnRtnBulletin((*CThostFtdcBulletinField)(unsafe.Pointer(pBulletin)))
	}
}
//export OnRtnTradingNotice
func OnRtnTradingNotice(spi unsafe.Pointer, pTradingNoticeInfo *C.struct_CThostFtdcTradingNoticeInfoField) {
	t := Trades[spi]
	if t.OnRtnTradingNotice == nil {
		fmt.Println("OnRtnTradingNotice")
	} else {
		t.OnRtnTradingNotice((*CThostFtdcTradingNoticeInfoField)(unsafe.Pointer(pTradingNoticeInfo)))
	}
}
//export OnRtnErrorConditionalOrder
func OnRtnErrorConditionalOrder(spi unsafe.Pointer, pErrorConditionalOrder *C.struct_CThostFtdcErrorConditionalOrderField) {
	t := Trades[spi]
	if t.OnRtnErrorConditionalOrder == nil {
		fmt.Println("OnRtnErrorConditionalOrder")
	} else {
		t.OnRtnErrorConditionalOrder((*CThostFtdcErrorConditionalOrderField)(unsafe.Pointer(pErrorConditionalOrder)))
	}
}
//export OnRtnExecOrder
func OnRtnExecOrder(spi unsafe.Pointer, pExecOrder *C.struct_CThostFtdcExecOrderField) {
	t := Trades[spi]
	if t.OnRtnExecOrder == nil {
		fmt.Println("OnRtnExecOrder")
	} else {
		t.OnRtnExecOrder((*CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)))
	}
}
//export OnErrRtnExecOrderInsert
func OnErrRtnExecOrderInsert(spi unsafe.Pointer, pInputExecOrder *C.struct_CThostFtdcInputExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnExecOrderInsert == nil {
		fmt.Println("OnErrRtnExecOrderInsert")
	} else {
		t.OnErrRtnExecOrderInsert((*CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnExecOrderAction
func OnErrRtnExecOrderAction(spi unsafe.Pointer, pExecOrderAction *C.struct_CThostFtdcExecOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnExecOrderAction == nil {
		fmt.Println("OnErrRtnExecOrderAction")
	} else {
		t.OnErrRtnExecOrderAction((*CThostFtdcExecOrderActionField)(unsafe.Pointer(pExecOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnForQuoteInsert
func OnErrRtnForQuoteInsert(spi unsafe.Pointer, pInputForQuote *C.struct_CThostFtdcInputForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnForQuoteInsert == nil {
		fmt.Println("OnErrRtnForQuoteInsert")
	} else {
		t.OnErrRtnForQuoteInsert((*CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRtnQuote
func OnRtnQuote(spi unsafe.Pointer, pQuote *C.struct_CThostFtdcQuoteField) {
	t := Trades[spi]
	if t.OnRtnQuote == nil {
		fmt.Println("OnRtnQuote")
	} else {
		t.OnRtnQuote((*CThostFtdcQuoteField)(unsafe.Pointer(pQuote)))
	}
}
//export OnErrRtnQuoteInsert
func OnErrRtnQuoteInsert(spi unsafe.Pointer, pInputQuote *C.struct_CThostFtdcInputQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnQuoteInsert == nil {
		fmt.Println("OnErrRtnQuoteInsert")
	} else {
		t.OnErrRtnQuoteInsert((*CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnQuoteAction
func OnErrRtnQuoteAction(spi unsafe.Pointer, pQuoteAction *C.struct_CThostFtdcQuoteActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnQuoteAction == nil {
		fmt.Println("OnErrRtnQuoteAction")
	} else {
		t.OnErrRtnQuoteAction((*CThostFtdcQuoteActionField)(unsafe.Pointer(pQuoteAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRtnForQuoteRsp
func OnRtnForQuoteRsp(spi unsafe.Pointer, pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField) {
	t := Trades[spi]
	if t.OnRtnForQuoteRsp == nil {
		fmt.Println("OnRtnForQuoteRsp")
	} else {
		t.OnRtnForQuoteRsp((*CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)))
	}
}
//export OnRtnCFMMCTradingAccountToken
func OnRtnCFMMCTradingAccountToken(spi unsafe.Pointer, pCFMMCTradingAccountToken *C.struct_CThostFtdcCFMMCTradingAccountTokenField) {
	t := Trades[spi]
	if t.OnRtnCFMMCTradingAccountToken == nil {
		fmt.Println("OnRtnCFMMCTradingAccountToken")
	} else {
		t.OnRtnCFMMCTradingAccountToken((*CThostFtdcCFMMCTradingAccountTokenField)(unsafe.Pointer(pCFMMCTradingAccountToken)))
	}
}
//export OnErrRtnBatchOrderAction
func OnErrRtnBatchOrderAction(spi unsafe.Pointer, pBatchOrderAction *C.struct_CThostFtdcBatchOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnBatchOrderAction == nil {
		fmt.Println("OnErrRtnBatchOrderAction")
	} else {
		t.OnErrRtnBatchOrderAction((*CThostFtdcBatchOrderActionField)(unsafe.Pointer(pBatchOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRtnOptionSelfClose
func OnRtnOptionSelfClose(spi unsafe.Pointer, pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField) {
	t := Trades[spi]
	if t.OnRtnOptionSelfClose == nil {
		fmt.Println("OnRtnOptionSelfClose")
	} else {
		t.OnRtnOptionSelfClose((*CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)))
	}
}
//export OnErrRtnOptionSelfCloseInsert
func OnErrRtnOptionSelfCloseInsert(spi unsafe.Pointer, pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnOptionSelfCloseInsert == nil {
		fmt.Println("OnErrRtnOptionSelfCloseInsert")
	} else {
		t.OnErrRtnOptionSelfCloseInsert((*CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnOptionSelfCloseAction
func OnErrRtnOptionSelfCloseAction(spi unsafe.Pointer, pOptionSelfCloseAction *C.struct_CThostFtdcOptionSelfCloseActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnOptionSelfCloseAction == nil {
		fmt.Println("OnErrRtnOptionSelfCloseAction")
	} else {
		t.OnErrRtnOptionSelfCloseAction((*CThostFtdcOptionSelfCloseActionField)(unsafe.Pointer(pOptionSelfCloseAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRtnCombAction
func OnRtnCombAction(spi unsafe.Pointer, pCombAction *C.struct_CThostFtdcCombActionField) {
	t := Trades[spi]
	if t.OnRtnCombAction == nil {
		fmt.Println("OnRtnCombAction")
	} else {
		t.OnRtnCombAction((*CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)))
	}
}
//export OnErrRtnCombActionInsert
func OnErrRtnCombActionInsert(spi unsafe.Pointer, pInputCombAction *C.struct_CThostFtdcInputCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnCombActionInsert == nil {
		fmt.Println("OnErrRtnCombActionInsert")
	} else {
		t.OnErrRtnCombActionInsert((*CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRspQryContractBank
func OnRspQryContractBank(spi unsafe.Pointer, pContractBank *C.struct_CThostFtdcContractBankField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryContractBank == nil {
		fmt.Println("OnRspQryContractBank")
	} else {
		t.OnRspQryContractBank((*CThostFtdcContractBankField)(unsafe.Pointer(pContractBank)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryParkedOrder
func OnRspQryParkedOrder(spi unsafe.Pointer, pParkedOrder *C.struct_CThostFtdcParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryParkedOrder == nil {
		fmt.Println("OnRspQryParkedOrder")
	} else {
		t.OnRspQryParkedOrder((*CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryParkedOrderAction
func OnRspQryParkedOrderAction(spi unsafe.Pointer, pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryParkedOrderAction == nil {
		fmt.Println("OnRspQryParkedOrderAction")
	} else {
		t.OnRspQryParkedOrderAction((*CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryTradingNotice
func OnRspQryTradingNotice(spi unsafe.Pointer, pTradingNotice *C.struct_CThostFtdcTradingNoticeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryTradingNotice == nil {
		fmt.Println("OnRspQryTradingNotice")
	} else {
		t.OnRspQryTradingNotice((*CThostFtdcTradingNoticeField)(unsafe.Pointer(pTradingNotice)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryBrokerTradingParams
func OnRspQryBrokerTradingParams(spi unsafe.Pointer, pBrokerTradingParams *C.struct_CThostFtdcBrokerTradingParamsField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryBrokerTradingParams == nil {
		fmt.Println("OnRspQryBrokerTradingParams")
	} else {
		t.OnRspQryBrokerTradingParams((*CThostFtdcBrokerTradingParamsField)(unsafe.Pointer(pBrokerTradingParams)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryBrokerTradingAlgos
func OnRspQryBrokerTradingAlgos(spi unsafe.Pointer, pBrokerTradingAlgos *C.struct_CThostFtdcBrokerTradingAlgosField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryBrokerTradingAlgos == nil {
		fmt.Println("OnRspQryBrokerTradingAlgos")
	} else {
		t.OnRspQryBrokerTradingAlgos((*CThostFtdcBrokerTradingAlgosField)(unsafe.Pointer(pBrokerTradingAlgos)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQueryCFMMCTradingAccountToken
func OnRspQueryCFMMCTradingAccountToken(spi unsafe.Pointer, pQueryCFMMCTradingAccountToken *C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQueryCFMMCTradingAccountToken == nil {
		fmt.Println("OnRspQueryCFMMCTradingAccountToken")
	} else {
		t.OnRspQueryCFMMCTradingAccountToken((*CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRtnFromBankToFutureByBank
func OnRtnFromBankToFutureByBank(spi unsafe.Pointer, pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	t := Trades[spi]
	if t.OnRtnFromBankToFutureByBank == nil {
		fmt.Println("OnRtnFromBankToFutureByBank")
	} else {
		t.OnRtnFromBankToFutureByBank((*CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}
//export OnRtnFromFutureToBankByBank
func OnRtnFromFutureToBankByBank(spi unsafe.Pointer, pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	t := Trades[spi]
	if t.OnRtnFromFutureToBankByBank == nil {
		fmt.Println("OnRtnFromFutureToBankByBank")
	} else {
		t.OnRtnFromFutureToBankByBank((*CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}
//export OnRtnRepealFromBankToFutureByBank
func OnRtnRepealFromBankToFutureByBank(spi unsafe.Pointer, pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	t := Trades[spi]
	if t.OnRtnRepealFromBankToFutureByBank == nil {
		fmt.Println("OnRtnRepealFromBankToFutureByBank")
	} else {
		t.OnRtnRepealFromBankToFutureByBank((*CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}
//export OnRtnRepealFromFutureToBankByBank
func OnRtnRepealFromFutureToBankByBank(spi unsafe.Pointer, pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	t := Trades[spi]
	if t.OnRtnRepealFromFutureToBankByBank == nil {
		fmt.Println("OnRtnRepealFromFutureToBankByBank")
	} else {
		t.OnRtnRepealFromFutureToBankByBank((*CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}
//export OnRtnFromBankToFutureByFuture
func OnRtnFromBankToFutureByFuture(spi unsafe.Pointer, pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	t := Trades[spi]
	if t.OnRtnFromBankToFutureByFuture == nil {
		fmt.Println("OnRtnFromBankToFutureByFuture")
	} else {
		t.OnRtnFromBankToFutureByFuture((*CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}
//export OnRtnFromFutureToBankByFuture
func OnRtnFromFutureToBankByFuture(spi unsafe.Pointer, pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	t := Trades[spi]
	if t.OnRtnFromFutureToBankByFuture == nil {
		fmt.Println("OnRtnFromFutureToBankByFuture")
	} else {
		t.OnRtnFromFutureToBankByFuture((*CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}
//export OnRtnRepealFromBankToFutureByFutureManual
func OnRtnRepealFromBankToFutureByFutureManual(spi unsafe.Pointer, pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	t := Trades[spi]
	if t.OnRtnRepealFromBankToFutureByFutureManual == nil {
		fmt.Println("OnRtnRepealFromBankToFutureByFutureManual")
	} else {
		t.OnRtnRepealFromBankToFutureByFutureManual((*CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}
//export OnRtnRepealFromFutureToBankByFutureManual
func OnRtnRepealFromFutureToBankByFutureManual(spi unsafe.Pointer, pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	t := Trades[spi]
	if t.OnRtnRepealFromFutureToBankByFutureManual == nil {
		fmt.Println("OnRtnRepealFromFutureToBankByFutureManual")
	} else {
		t.OnRtnRepealFromFutureToBankByFutureManual((*CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}
//export OnRtnQueryBankBalanceByFuture
func OnRtnQueryBankBalanceByFuture(spi unsafe.Pointer, pNotifyQueryAccount *C.struct_CThostFtdcNotifyQueryAccountField) {
	t := Trades[spi]
	if t.OnRtnQueryBankBalanceByFuture == nil {
		fmt.Println("OnRtnQueryBankBalanceByFuture")
	} else {
		t.OnRtnQueryBankBalanceByFuture((*CThostFtdcNotifyQueryAccountField)(unsafe.Pointer(pNotifyQueryAccount)))
	}
}
//export OnErrRtnBankToFutureByFuture
func OnErrRtnBankToFutureByFuture(spi unsafe.Pointer, pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnBankToFutureByFuture == nil {
		fmt.Println("OnErrRtnBankToFutureByFuture")
	} else {
		t.OnErrRtnBankToFutureByFuture((*CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnFutureToBankByFuture
func OnErrRtnFutureToBankByFuture(spi unsafe.Pointer, pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnFutureToBankByFuture == nil {
		fmt.Println("OnErrRtnFutureToBankByFuture")
	} else {
		t.OnErrRtnFutureToBankByFuture((*CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnRepealBankToFutureByFutureManual
func OnErrRtnRepealBankToFutureByFutureManual(spi unsafe.Pointer, pReqRepeal *C.struct_CThostFtdcReqRepealField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnRepealBankToFutureByFutureManual == nil {
		fmt.Println("OnErrRtnRepealBankToFutureByFutureManual")
	} else {
		t.OnErrRtnRepealBankToFutureByFutureManual((*CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnRepealFutureToBankByFutureManual
func OnErrRtnRepealFutureToBankByFutureManual(spi unsafe.Pointer, pReqRepeal *C.struct_CThostFtdcReqRepealField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnRepealFutureToBankByFutureManual == nil {
		fmt.Println("OnErrRtnRepealFutureToBankByFutureManual")
	} else {
		t.OnErrRtnRepealFutureToBankByFutureManual((*CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnErrRtnQueryBankBalanceByFuture
func OnErrRtnQueryBankBalanceByFuture(spi unsafe.Pointer, pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	t := Trades[spi]
	if t.OnErrRtnQueryBankBalanceByFuture == nil {
		fmt.Println("OnErrRtnQueryBankBalanceByFuture")
	} else {
		t.OnErrRtnQueryBankBalanceByFuture((*CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}
//export OnRtnRepealFromBankToFutureByFuture
func OnRtnRepealFromBankToFutureByFuture(spi unsafe.Pointer, pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	t := Trades[spi]
	if t.OnRtnRepealFromBankToFutureByFuture == nil {
		fmt.Println("OnRtnRepealFromBankToFutureByFuture")
	} else {
		t.OnRtnRepealFromBankToFutureByFuture((*CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}
//export OnRtnRepealFromFutureToBankByFuture
func OnRtnRepealFromFutureToBankByFuture(spi unsafe.Pointer, pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	t := Trades[spi]
	if t.OnRtnRepealFromFutureToBankByFuture == nil {
		fmt.Println("OnRtnRepealFromFutureToBankByFuture")
	} else {
		t.OnRtnRepealFromFutureToBankByFuture((*CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}
//export OnRspFromBankToFutureByFuture
func OnRspFromBankToFutureByFuture(spi unsafe.Pointer, pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspFromBankToFutureByFuture == nil {
		fmt.Println("OnRspFromBankToFutureByFuture")
	} else {
		t.OnRspFromBankToFutureByFuture((*CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspFromFutureToBankByFuture
func OnRspFromFutureToBankByFuture(spi unsafe.Pointer, pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspFromFutureToBankByFuture == nil {
		fmt.Println("OnRspFromFutureToBankByFuture")
	} else {
		t.OnRspFromFutureToBankByFuture((*CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQueryBankAccountMoneyByFuture
func OnRspQueryBankAccountMoneyByFuture(spi unsafe.Pointer, pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQueryBankAccountMoneyByFuture == nil {
		fmt.Println("OnRspQueryBankAccountMoneyByFuture")
	} else {
		t.OnRspQueryBankAccountMoneyByFuture((*CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRtnOpenAccountByBank
func OnRtnOpenAccountByBank(spi unsafe.Pointer, pOpenAccount *C.struct_CThostFtdcOpenAccountField) {
	t := Trades[spi]
	if t.OnRtnOpenAccountByBank == nil {
		fmt.Println("OnRtnOpenAccountByBank")
	} else {
		t.OnRtnOpenAccountByBank((*CThostFtdcOpenAccountField)(unsafe.Pointer(pOpenAccount)))
	}
}
//export OnRtnCancelAccountByBank
func OnRtnCancelAccountByBank(spi unsafe.Pointer, pCancelAccount *C.struct_CThostFtdcCancelAccountField) {
	t := Trades[spi]
	if t.OnRtnCancelAccountByBank == nil {
		fmt.Println("OnRtnCancelAccountByBank")
	} else {
		t.OnRtnCancelAccountByBank((*CThostFtdcCancelAccountField)(unsafe.Pointer(pCancelAccount)))
	}
}
//export OnRtnChangeAccountByBank
func OnRtnChangeAccountByBank(spi unsafe.Pointer, pChangeAccount *C.struct_CThostFtdcChangeAccountField) {
	t := Trades[spi]
	if t.OnRtnChangeAccountByBank == nil {
		fmt.Println("OnRtnChangeAccountByBank")
	} else {
		t.OnRtnChangeAccountByBank((*CThostFtdcChangeAccountField)(unsafe.Pointer(pChangeAccount)))
	}
}
//export OnRspQryClassifiedInstrument
func OnRspQryClassifiedInstrument(spi unsafe.Pointer, pInstrument *C.struct_CThostFtdcInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryClassifiedInstrument == nil {
		fmt.Println("OnRspQryClassifiedInstrument")
	} else {
		t.OnRspQryClassifiedInstrument((*CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryCombPromotionParam
func OnRspQryCombPromotionParam(spi unsafe.Pointer, pCombPromotionParam *C.struct_CThostFtdcCombPromotionParamField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryCombPromotionParam == nil {
		fmt.Println("OnRspQryCombPromotionParam")
	} else {
		t.OnRspQryCombPromotionParam((*CThostFtdcCombPromotionParamField)(unsafe.Pointer(pCombPromotionParam)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryRiskSettleInvstPosition
func OnRspQryRiskSettleInvstPosition(spi unsafe.Pointer, pRiskSettleInvstPosition *C.struct_CThostFtdcRiskSettleInvstPositionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryRiskSettleInvstPosition == nil {
		fmt.Println("OnRspQryRiskSettleInvstPosition")
	} else {
		t.OnRspQryRiskSettleInvstPosition((*CThostFtdcRiskSettleInvstPositionField)(unsafe.Pointer(pRiskSettleInvstPosition)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryRiskSettleProductStatus
func OnRspQryRiskSettleProductStatus(spi unsafe.Pointer, pRiskSettleProductStatus *C.struct_CThostFtdcRiskSettleProductStatusField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryRiskSettleProductStatus == nil {
		fmt.Println("OnRspQryRiskSettleProductStatus")
	} else {
		t.OnRspQryRiskSettleProductStatus((*CThostFtdcRiskSettleProductStatusField)(unsafe.Pointer(pRiskSettleProductStatus)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySPBMFutureParameter
func OnRspQrySPBMFutureParameter(spi unsafe.Pointer, pSPBMFutureParameter *C.struct_CThostFtdcSPBMFutureParameterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySPBMFutureParameter == nil {
		fmt.Println("OnRspQrySPBMFutureParameter")
	} else {
		t.OnRspQrySPBMFutureParameter((*CThostFtdcSPBMFutureParameterField)(unsafe.Pointer(pSPBMFutureParameter)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySPBMOptionParameter
func OnRspQrySPBMOptionParameter(spi unsafe.Pointer, pSPBMOptionParameter *C.struct_CThostFtdcSPBMOptionParameterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySPBMOptionParameter == nil {
		fmt.Println("OnRspQrySPBMOptionParameter")
	} else {
		t.OnRspQrySPBMOptionParameter((*CThostFtdcSPBMOptionParameterField)(unsafe.Pointer(pSPBMOptionParameter)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySPBMIntraParameter
func OnRspQrySPBMIntraParameter(spi unsafe.Pointer, pSPBMIntraParameter *C.struct_CThostFtdcSPBMIntraParameterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySPBMIntraParameter == nil {
		fmt.Println("OnRspQrySPBMIntraParameter")
	} else {
		t.OnRspQrySPBMIntraParameter((*CThostFtdcSPBMIntraParameterField)(unsafe.Pointer(pSPBMIntraParameter)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySPBMInterParameter
func OnRspQrySPBMInterParameter(spi unsafe.Pointer, pSPBMInterParameter *C.struct_CThostFtdcSPBMInterParameterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySPBMInterParameter == nil {
		fmt.Println("OnRspQrySPBMInterParameter")
	} else {
		t.OnRspQrySPBMInterParameter((*CThostFtdcSPBMInterParameterField)(unsafe.Pointer(pSPBMInterParameter)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySPBMPortfDefinition
func OnRspQrySPBMPortfDefinition(spi unsafe.Pointer, pSPBMPortfDefinition *C.struct_CThostFtdcSPBMPortfDefinitionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySPBMPortfDefinition == nil {
		fmt.Println("OnRspQrySPBMPortfDefinition")
	} else {
		t.OnRspQrySPBMPortfDefinition((*CThostFtdcSPBMPortfDefinitionField)(unsafe.Pointer(pSPBMPortfDefinition)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQrySPBMInvestorPortfDef
func OnRspQrySPBMInvestorPortfDef(spi unsafe.Pointer, pSPBMInvestorPortfDef *C.struct_CThostFtdcSPBMInvestorPortfDefField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQrySPBMInvestorPortfDef == nil {
		fmt.Println("OnRspQrySPBMInvestorPortfDef")
	} else {
		t.OnRspQrySPBMInvestorPortfDef((*CThostFtdcSPBMInvestorPortfDefField)(unsafe.Pointer(pSPBMInvestorPortfDef)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestorPortfMarginRatio
func OnRspQryInvestorPortfMarginRatio(spi unsafe.Pointer, pInvestorPortfMarginRatio *C.struct_CThostFtdcInvestorPortfMarginRatioField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestorPortfMarginRatio == nil {
		fmt.Println("OnRspQryInvestorPortfMarginRatio")
	} else {
		t.OnRspQryInvestorPortfMarginRatio((*CThostFtdcInvestorPortfMarginRatioField)(unsafe.Pointer(pInvestorPortfMarginRatio)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}
//export OnRspQryInvestorProdSPBMDetail
func OnRspQryInvestorProdSPBMDetail(spi unsafe.Pointer, pInvestorProdSPBMDetail *C.struct_CThostFtdcInvestorProdSPBMDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	t := Trades[spi]
	if t.OnRspQryInvestorProdSPBMDetail == nil {
		fmt.Println("OnRspQryInvestorProdSPBMDetail")
	} else {
		t.OnRspQryInvestorProdSPBMDetail((*CThostFtdcInvestorProdSPBMDetailField)(unsafe.Pointer(pInvestorProdSPBMDetail)), (*CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}


// 创建TraderApi
func (t *Trade)Release() {
	C.tRelease(t.api)
}
// 初始化
func (t *Trade)Init() {
	C.tInit(t.api)
}
// 等待接口线程结束运行
func (t *Trade)Join() int{
	return int(C.tJoin(t.api))
}
// 注册前置机网络地址
func (t *Trade)RegisterFront(pszFrontAddress string) {
	C.tRegisterFront(t.api, C.CString(pszFrontAddress))
}
// @remark RegisterNameServer优先于RegisterFront
func (t *Trade)RegisterNameServer(pszNsAddress string) {
	C.tRegisterNameServer(t.api, C.CString(pszNsAddress))
}
// 注册名字服务器用户信息
func (t *Trade)RegisterFensUserInfo(pFensUserInfo *CThostFtdcFensUserInfoField) {
	C.tRegisterFensUserInfo(t.api, (*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(pFensUserInfo)))
}
// 注册回调接口
func (t *Trade)RegisterSpi(pSpi unsafe.Pointer) {
	C.tRegisterSpi(t.api, pSpi)
}
// 订阅私有流。
func (t *Trade)SubscribePrivateTopic(nResumeType THOST_TE_RESUME_TYPE) {
	C.tSubscribePrivateTopic(t.api, C.int(nResumeType))
}
// 订阅公共流。
func (t *Trade)SubscribePublicTopic(nResumeType THOST_TE_RESUME_TYPE) {
	C.tSubscribePublicTopic(t.api, C.int(nResumeType))
}
// 客户端认证请求
func (t *Trade)ReqAuthenticate(pReqAuthenticateField *CThostFtdcReqAuthenticateField, nRequestID int) int{
	return int(C.tReqAuthenticate(t.api, (*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(pReqAuthenticateField)), C.int(nRequestID)))
}
// 注册用户终端信息，用于中继服务器多连接模式
func (t *Trade)RegisterUserSystemInfo(pUserSystemInfo *CThostFtdcUserSystemInfoField) int{
	return int(C.tRegisterUserSystemInfo(t.api, (*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo))))
}
// 上报用户终端信息，用于中继服务器操作员登录模式
func (t *Trade)SubmitUserSystemInfo(pUserSystemInfo *CThostFtdcUserSystemInfoField) int{
	return int(C.tSubmitUserSystemInfo(t.api, (*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo))))
}
// 用户登录请求
func (t *Trade)ReqUserLogin(pReqUserLoginField *CThostFtdcReqUserLoginField, nRequestID int) int{
	return int(C.tReqUserLogin(t.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(pReqUserLoginField)), C.int(nRequestID)))
}
// 登出请求
func (t *Trade)ReqUserLogout(pUserLogout *CThostFtdcUserLogoutField, nRequestID int) int{
	return int(C.tReqUserLogout(t.api, (*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), C.int(nRequestID)))
}
// 用户口令更新请求
func (t *Trade)ReqUserPasswordUpdate(pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, nRequestID int) int{
	return int(C.tReqUserPasswordUpdate(t.api, (*C.struct_CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)), C.int(nRequestID)))
}
// 资金账户口令更新请求
func (t *Trade)ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, nRequestID int) int{
	return int(C.tReqTradingAccountPasswordUpdate(t.api, (*C.struct_CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)), C.int(nRequestID)))
}
// 查询用户当前支持的认证模式
func (t *Trade)ReqUserAuthMethod(pReqUserAuthMethod *CThostFtdcReqUserAuthMethodField, nRequestID int) int{
	return int(C.tReqUserAuthMethod(t.api, (*C.struct_CThostFtdcReqUserAuthMethodField)(unsafe.Pointer(pReqUserAuthMethod)), C.int(nRequestID)))
}
// 用户发出获取图形验证码请求
func (t *Trade)ReqGenUserCaptcha(pReqGenUserCaptcha *CThostFtdcReqGenUserCaptchaField, nRequestID int) int{
	return int(C.tReqGenUserCaptcha(t.api, (*C.struct_CThostFtdcReqGenUserCaptchaField)(unsafe.Pointer(pReqGenUserCaptcha)), C.int(nRequestID)))
}
// 用户发出获取短信验证码请求
func (t *Trade)ReqGenUserText(pReqGenUserText *CThostFtdcReqGenUserTextField, nRequestID int) int{
	return int(C.tReqGenUserText(t.api, (*C.struct_CThostFtdcReqGenUserTextField)(unsafe.Pointer(pReqGenUserText)), C.int(nRequestID)))
}
// 用户发出带有图片验证码的登陆请求
func (t *Trade)ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha *CThostFtdcReqUserLoginWithCaptchaField, nRequestID int) int{
	return int(C.tReqUserLoginWithCaptcha(t.api, (*C.struct_CThostFtdcReqUserLoginWithCaptchaField)(unsafe.Pointer(pReqUserLoginWithCaptcha)), C.int(nRequestID)))
}
// 用户发出带有短信验证码的登陆请求
func (t *Trade)ReqUserLoginWithText(pReqUserLoginWithText *CThostFtdcReqUserLoginWithTextField, nRequestID int) int{
	return int(C.tReqUserLoginWithText(t.api, (*C.struct_CThostFtdcReqUserLoginWithTextField)(unsafe.Pointer(pReqUserLoginWithText)), C.int(nRequestID)))
}
// 用户发出带有动态口令的登陆请求
func (t *Trade)ReqUserLoginWithOTP(pReqUserLoginWithOTP *CThostFtdcReqUserLoginWithOTPField, nRequestID int) int{
	return int(C.tReqUserLoginWithOTP(t.api, (*C.struct_CThostFtdcReqUserLoginWithOTPField)(unsafe.Pointer(pReqUserLoginWithOTP)), C.int(nRequestID)))
}
// 报单录入请求
func (t *Trade)ReqOrderInsert(pInputOrder *CThostFtdcInputOrderField, nRequestID int) int{
	return int(C.tReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), C.int(nRequestID)))
}
// 预埋单录入请求
func (t *Trade)ReqParkedOrderInsert(pParkedOrder *CThostFtdcParkedOrderField, nRequestID int) int{
	return int(C.tReqParkedOrderInsert(t.api, (*C.struct_CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), C.int(nRequestID)))
}
// 预埋撤单录入请求
func (t *Trade)ReqParkedOrderAction(pParkedOrderAction *CThostFtdcParkedOrderActionField, nRequestID int) int{
	return int(C.tReqParkedOrderAction(t.api, (*C.struct_CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), C.int(nRequestID)))
}
// 报单操作请求
func (t *Trade)ReqOrderAction(pInputOrderAction *CThostFtdcInputOrderActionField, nRequestID int) int{
	return int(C.tReqOrderAction(t.api, (*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)), C.int(nRequestID)))
}
// 查询最大报单数量请求
func (t *Trade)ReqQryMaxOrderVolume(pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, nRequestID int) int{
	return int(C.tReqQryMaxOrderVolume(t.api, (*C.struct_CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)), C.int(nRequestID)))
}
// 投资者结算结果确认
func (t *Trade)ReqSettlementInfoConfirm(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, nRequestID int) int{
	return int(C.tReqSettlementInfoConfirm(t.api, (*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), C.int(nRequestID)))
}
// 请求删除预埋单
func (t *Trade)ReqRemoveParkedOrder(pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, nRequestID int) int{
	return int(C.tReqRemoveParkedOrder(t.api, (*C.struct_CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)), C.int(nRequestID)))
}
// 请求删除预埋撤单
func (t *Trade)ReqRemoveParkedOrderAction(pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, nRequestID int) int{
	return int(C.tReqRemoveParkedOrderAction(t.api, (*C.struct_CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)), C.int(nRequestID)))
}
// 执行宣告录入请求
func (t *Trade)ReqExecOrderInsert(pInputExecOrder *CThostFtdcInputExecOrderField, nRequestID int) int{
	return int(C.tReqExecOrderInsert(t.api, (*C.struct_CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), C.int(nRequestID)))
}
// 执行宣告操作请求
func (t *Trade)ReqExecOrderAction(pInputExecOrderAction *CThostFtdcInputExecOrderActionField, nRequestID int) int{
	return int(C.tReqExecOrderAction(t.api, (*C.struct_CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)), C.int(nRequestID)))
}
// 询价录入请求
func (t *Trade)ReqForQuoteInsert(pInputForQuote *CThostFtdcInputForQuoteField, nRequestID int) int{
	return int(C.tReqForQuoteInsert(t.api, (*C.struct_CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), C.int(nRequestID)))
}
// 报价录入请求
func (t *Trade)ReqQuoteInsert(pInputQuote *CThostFtdcInputQuoteField, nRequestID int) int{
	return int(C.tReqQuoteInsert(t.api, (*C.struct_CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), C.int(nRequestID)))
}
// 报价操作请求
func (t *Trade)ReqQuoteAction(pInputQuoteAction *CThostFtdcInputQuoteActionField, nRequestID int) int{
	return int(C.tReqQuoteAction(t.api, (*C.struct_CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)), C.int(nRequestID)))
}
// 批量报单操作请求
func (t *Trade)ReqBatchOrderAction(pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, nRequestID int) int{
	return int(C.tReqBatchOrderAction(t.api, (*C.struct_CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)), C.int(nRequestID)))
}
// 期权自对冲录入请求
func (t *Trade)ReqOptionSelfCloseInsert(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, nRequestID int) int{
	return int(C.tReqOptionSelfCloseInsert(t.api, (*C.struct_CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), C.int(nRequestID)))
}
// 期权自对冲操作请求
func (t *Trade)ReqOptionSelfCloseAction(pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, nRequestID int) int{
	return int(C.tReqOptionSelfCloseAction(t.api, (*C.struct_CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)), C.int(nRequestID)))
}
// 申请组合录入请求
func (t *Trade)ReqCombActionInsert(pInputCombAction *CThostFtdcInputCombActionField, nRequestID int) int{
	return int(C.tReqCombActionInsert(t.api, (*C.struct_CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), C.int(nRequestID)))
}
// 请求查询报单
func (t *Trade)ReqQryOrder(pQryOrder *CThostFtdcQryOrderField, nRequestID int) int{
	return int(C.tReqQryOrder(t.api, (*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(pQryOrder)), C.int(nRequestID)))
}
// 请求查询成交
func (t *Trade)ReqQryTrade(pQryTrade *CThostFtdcQryTradeField, nRequestID int) int{
	return int(C.tReqQryTrade(t.api, (*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(pQryTrade)), C.int(nRequestID)))
}
// 请求查询投资者持仓
func (t *Trade)ReqQryInvestorPosition(pQryInvestorPosition *CThostFtdcQryInvestorPositionField, nRequestID int) int{
	return int(C.tReqQryInvestorPosition(t.api, (*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(pQryInvestorPosition)), C.int(nRequestID)))
}
// 请求查询资金账户
func (t *Trade)ReqQryTradingAccount(pQryTradingAccount *CThostFtdcQryTradingAccountField, nRequestID int) int{
	return int(C.tReqQryTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(pQryTradingAccount)), C.int(nRequestID)))
}
// 请求查询投资者
func (t *Trade)ReqQryInvestor(pQryInvestor *CThostFtdcQryInvestorField, nRequestID int) int{
	return int(C.tReqQryInvestor(t.api, (*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(pQryInvestor)), C.int(nRequestID)))
}
// 请求查询交易编码
func (t *Trade)ReqQryTradingCode(pQryTradingCode *CThostFtdcQryTradingCodeField, nRequestID int) int{
	return int(C.tReqQryTradingCode(t.api, (*C.struct_CThostFtdcQryTradingCodeField)(unsafe.Pointer(pQryTradingCode)), C.int(nRequestID)))
}
// 请求查询合约保证金率
func (t *Trade)ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *CThostFtdcQryInstrumentMarginRateField, nRequestID int) int{
	return int(C.tReqQryInstrumentMarginRate(t.api, (*C.struct_CThostFtdcQryInstrumentMarginRateField)(unsafe.Pointer(pQryInstrumentMarginRate)), C.int(nRequestID)))
}
// 请求查询合约手续费率
func (t *Trade)ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *CThostFtdcQryInstrumentCommissionRateField, nRequestID int) int{
	return int(C.tReqQryInstrumentCommissionRate(t.api, (*C.struct_CThostFtdcQryInstrumentCommissionRateField)(unsafe.Pointer(pQryInstrumentCommissionRate)), C.int(nRequestID)))
}
// 请求查询交易所
func (t *Trade)ReqQryExchange(pQryExchange *CThostFtdcQryExchangeField, nRequestID int) int{
	return int(C.tReqQryExchange(t.api, (*C.struct_CThostFtdcQryExchangeField)(unsafe.Pointer(pQryExchange)), C.int(nRequestID)))
}
// 请求查询产品
func (t *Trade)ReqQryProduct(pQryProduct *CThostFtdcQryProductField, nRequestID int) int{
	return int(C.tReqQryProduct(t.api, (*C.struct_CThostFtdcQryProductField)(unsafe.Pointer(pQryProduct)), C.int(nRequestID)))
}
// 请求查询合约
func (t *Trade)ReqQryInstrument(pQryInstrument *CThostFtdcQryInstrumentField, nRequestID int) int{
	return int(C.tReqQryInstrument(t.api, (*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(pQryInstrument)), C.int(nRequestID)))
}
// 请求查询行情
func (t *Trade)ReqQryDepthMarketData(pQryDepthMarketData *CThostFtdcQryDepthMarketDataField, nRequestID int) int{
	return int(C.tReqQryDepthMarketData(t.api, (*C.struct_CThostFtdcQryDepthMarketDataField)(unsafe.Pointer(pQryDepthMarketData)), C.int(nRequestID)))
}
// 请求查询交易员报盘机
func (t *Trade)ReqQryTraderOffer(pQryTraderOffer *CThostFtdcQryTraderOfferField, nRequestID int) int{
	return int(C.tReqQryTraderOffer(t.api, (*C.struct_CThostFtdcQryTraderOfferField)(unsafe.Pointer(pQryTraderOffer)), C.int(nRequestID)))
}
// 请求查询投资者结算结果
func (t *Trade)ReqQrySettlementInfo(pQrySettlementInfo *CThostFtdcQrySettlementInfoField, nRequestID int) int{
	return int(C.tReqQrySettlementInfo(t.api, (*C.struct_CThostFtdcQrySettlementInfoField)(unsafe.Pointer(pQrySettlementInfo)), C.int(nRequestID)))
}
// 请求查询转帐银行
func (t *Trade)ReqQryTransferBank(pQryTransferBank *CThostFtdcQryTransferBankField, nRequestID int) int{
	return int(C.tReqQryTransferBank(t.api, (*C.struct_CThostFtdcQryTransferBankField)(unsafe.Pointer(pQryTransferBank)), C.int(nRequestID)))
}
// 请求查询投资者持仓明细
func (t *Trade)ReqQryInvestorPositionDetail(pQryInvestorPositionDetail *CThostFtdcQryInvestorPositionDetailField, nRequestID int) int{
	return int(C.tReqQryInvestorPositionDetail(t.api, (*C.struct_CThostFtdcQryInvestorPositionDetailField)(unsafe.Pointer(pQryInvestorPositionDetail)), C.int(nRequestID)))
}
// 请求查询客户通知
func (t *Trade)ReqQryNotice(pQryNotice *CThostFtdcQryNoticeField, nRequestID int) int{
	return int(C.tReqQryNotice(t.api, (*C.struct_CThostFtdcQryNoticeField)(unsafe.Pointer(pQryNotice)), C.int(nRequestID)))
}
// 请求查询结算信息确认
func (t *Trade)ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *CThostFtdcQrySettlementInfoConfirmField, nRequestID int) int{
	return int(C.tReqQrySettlementInfoConfirm(t.api, (*C.struct_CThostFtdcQrySettlementInfoConfirmField)(unsafe.Pointer(pQrySettlementInfoConfirm)), C.int(nRequestID)))
}
// 请求查询投资者持仓明细
func (t *Trade)ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail *CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int) int{
	return int(C.tReqQryInvestorPositionCombineDetail(t.api, (*C.struct_CThostFtdcQryInvestorPositionCombineDetailField)(unsafe.Pointer(pQryInvestorPositionCombineDetail)), C.int(nRequestID)))
}
// 请求查询保证金监管系统经纪公司资金账户密钥
func (t *Trade)ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey *CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int) int{
	return int(C.tReqQryCFMMCTradingAccountKey(t.api, (*C.struct_CThostFtdcQryCFMMCTradingAccountKeyField)(unsafe.Pointer(pQryCFMMCTradingAccountKey)), C.int(nRequestID)))
}
// 请求查询仓单折抵信息
func (t *Trade)ReqQryEWarrantOffset(pQryEWarrantOffset *CThostFtdcQryEWarrantOffsetField, nRequestID int) int{
	return int(C.tReqQryEWarrantOffset(t.api, (*C.struct_CThostFtdcQryEWarrantOffsetField)(unsafe.Pointer(pQryEWarrantOffset)), C.int(nRequestID)))
}
// 请求查询投资者品种/跨品种保证金
func (t *Trade)ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin *CThostFtdcQryInvestorProductGroupMarginField, nRequestID int) int{
	return int(C.tReqQryInvestorProductGroupMargin(t.api, (*C.struct_CThostFtdcQryInvestorProductGroupMarginField)(unsafe.Pointer(pQryInvestorProductGroupMargin)), C.int(nRequestID)))
}
// 请求查询交易所保证金率
func (t *Trade)ReqQryExchangeMarginRate(pQryExchangeMarginRate *CThostFtdcQryExchangeMarginRateField, nRequestID int) int{
	return int(C.tReqQryExchangeMarginRate(t.api, (*C.struct_CThostFtdcQryExchangeMarginRateField)(unsafe.Pointer(pQryExchangeMarginRate)), C.int(nRequestID)))
}
// 请求查询交易所调整保证金率
func (t *Trade)ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust *CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int) int{
	return int(C.tReqQryExchangeMarginRateAdjust(t.api, (*C.struct_CThostFtdcQryExchangeMarginRateAdjustField)(unsafe.Pointer(pQryExchangeMarginRateAdjust)), C.int(nRequestID)))
}
// 请求查询汇率
func (t *Trade)ReqQryExchangeRate(pQryExchangeRate *CThostFtdcQryExchangeRateField, nRequestID int) int{
	return int(C.tReqQryExchangeRate(t.api, (*C.struct_CThostFtdcQryExchangeRateField)(unsafe.Pointer(pQryExchangeRate)), C.int(nRequestID)))
}
// 请求查询二级代理操作员银期权限
func (t *Trade)ReqQrySecAgentACIDMap(pQrySecAgentACIDMap *CThostFtdcQrySecAgentACIDMapField, nRequestID int) int{
	return int(C.tReqQrySecAgentACIDMap(t.api, (*C.struct_CThostFtdcQrySecAgentACIDMapField)(unsafe.Pointer(pQrySecAgentACIDMap)), C.int(nRequestID)))
}
// 请求查询产品报价汇率
func (t *Trade)ReqQryProductExchRate(pQryProductExchRate *CThostFtdcQryProductExchRateField, nRequestID int) int{
	return int(C.tReqQryProductExchRate(t.api, (*C.struct_CThostFtdcQryProductExchRateField)(unsafe.Pointer(pQryProductExchRate)), C.int(nRequestID)))
}
// 请求查询产品组
func (t *Trade)ReqQryProductGroup(pQryProductGroup *CThostFtdcQryProductGroupField, nRequestID int) int{
	return int(C.tReqQryProductGroup(t.api, (*C.struct_CThostFtdcQryProductGroupField)(unsafe.Pointer(pQryProductGroup)), C.int(nRequestID)))
}
// 请求查询做市商合约手续费率
func (t *Trade)ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate *CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int) int{
	return int(C.tReqQryMMInstrumentCommissionRate(t.api, (*C.struct_CThostFtdcQryMMInstrumentCommissionRateField)(unsafe.Pointer(pQryMMInstrumentCommissionRate)), C.int(nRequestID)))
}
// 请求查询做市商期权合约手续费
func (t *Trade)ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate *CThostFtdcQryMMOptionInstrCommRateField, nRequestID int) int{
	return int(C.tReqQryMMOptionInstrCommRate(t.api, (*C.struct_CThostFtdcQryMMOptionInstrCommRateField)(unsafe.Pointer(pQryMMOptionInstrCommRate)), C.int(nRequestID)))
}
// 请求查询报单手续费
func (t *Trade)ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate *CThostFtdcQryInstrumentOrderCommRateField, nRequestID int) int{
	return int(C.tReqQryInstrumentOrderCommRate(t.api, (*C.struct_CThostFtdcQryInstrumentOrderCommRateField)(unsafe.Pointer(pQryInstrumentOrderCommRate)), C.int(nRequestID)))
}
// 请求查询资金账户
func (t *Trade)ReqQrySecAgentTradingAccount(pQryTradingAccount *CThostFtdcQryTradingAccountField, nRequestID int) int{
	return int(C.tReqQrySecAgentTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(pQryTradingAccount)), C.int(nRequestID)))
}
// 请求查询二级代理商资金校验模式
func (t *Trade)ReqQrySecAgentCheckMode(pQrySecAgentCheckMode *CThostFtdcQrySecAgentCheckModeField, nRequestID int) int{
	return int(C.tReqQrySecAgentCheckMode(t.api, (*C.struct_CThostFtdcQrySecAgentCheckModeField)(unsafe.Pointer(pQrySecAgentCheckMode)), C.int(nRequestID)))
}
// 请求查询二级代理商信息
func (t *Trade)ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo *CThostFtdcQrySecAgentTradeInfoField, nRequestID int) int{
	return int(C.tReqQrySecAgentTradeInfo(t.api, (*C.struct_CThostFtdcQrySecAgentTradeInfoField)(unsafe.Pointer(pQrySecAgentTradeInfo)), C.int(nRequestID)))
}
// 请求查询期权交易成本
func (t *Trade)ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost *CThostFtdcQryOptionInstrTradeCostField, nRequestID int) int{
	return int(C.tReqQryOptionInstrTradeCost(t.api, (*C.struct_CThostFtdcQryOptionInstrTradeCostField)(unsafe.Pointer(pQryOptionInstrTradeCost)), C.int(nRequestID)))
}
// 请求查询期权合约手续费
func (t *Trade)ReqQryOptionInstrCommRate(pQryOptionInstrCommRate *CThostFtdcQryOptionInstrCommRateField, nRequestID int) int{
	return int(C.tReqQryOptionInstrCommRate(t.api, (*C.struct_CThostFtdcQryOptionInstrCommRateField)(unsafe.Pointer(pQryOptionInstrCommRate)), C.int(nRequestID)))
}
// 请求查询执行宣告
func (t *Trade)ReqQryExecOrder(pQryExecOrder *CThostFtdcQryExecOrderField, nRequestID int) int{
	return int(C.tReqQryExecOrder(t.api, (*C.struct_CThostFtdcQryExecOrderField)(unsafe.Pointer(pQryExecOrder)), C.int(nRequestID)))
}
// 请求查询询价
func (t *Trade)ReqQryForQuote(pQryForQuote *CThostFtdcQryForQuoteField, nRequestID int) int{
	return int(C.tReqQryForQuote(t.api, (*C.struct_CThostFtdcQryForQuoteField)(unsafe.Pointer(pQryForQuote)), C.int(nRequestID)))
}
// 请求查询报价
func (t *Trade)ReqQryQuote(pQryQuote *CThostFtdcQryQuoteField, nRequestID int) int{
	return int(C.tReqQryQuote(t.api, (*C.struct_CThostFtdcQryQuoteField)(unsafe.Pointer(pQryQuote)), C.int(nRequestID)))
}
// 请求查询期权自对冲
func (t *Trade)ReqQryOptionSelfClose(pQryOptionSelfClose *CThostFtdcQryOptionSelfCloseField, nRequestID int) int{
	return int(C.tReqQryOptionSelfClose(t.api, (*C.struct_CThostFtdcQryOptionSelfCloseField)(unsafe.Pointer(pQryOptionSelfClose)), C.int(nRequestID)))
}
// 请求查询投资单元
func (t *Trade)ReqQryInvestUnit(pQryInvestUnit *CThostFtdcQryInvestUnitField, nRequestID int) int{
	return int(C.tReqQryInvestUnit(t.api, (*C.struct_CThostFtdcQryInvestUnitField)(unsafe.Pointer(pQryInvestUnit)), C.int(nRequestID)))
}
// 请求查询组合合约安全系数
func (t *Trade)ReqQryCombInstrumentGuard(pQryCombInstrumentGuard *CThostFtdcQryCombInstrumentGuardField, nRequestID int) int{
	return int(C.tReqQryCombInstrumentGuard(t.api, (*C.struct_CThostFtdcQryCombInstrumentGuardField)(unsafe.Pointer(pQryCombInstrumentGuard)), C.int(nRequestID)))
}
// 请求查询申请组合
func (t *Trade)ReqQryCombAction(pQryCombAction *CThostFtdcQryCombActionField, nRequestID int) int{
	return int(C.tReqQryCombAction(t.api, (*C.struct_CThostFtdcQryCombActionField)(unsafe.Pointer(pQryCombAction)), C.int(nRequestID)))
}
// 请求查询转帐流水
func (t *Trade)ReqQryTransferSerial(pQryTransferSerial *CThostFtdcQryTransferSerialField, nRequestID int) int{
	return int(C.tReqQryTransferSerial(t.api, (*C.struct_CThostFtdcQryTransferSerialField)(unsafe.Pointer(pQryTransferSerial)), C.int(nRequestID)))
}
// 请求查询银期签约关系
func (t *Trade)ReqQryAccountregister(pQryAccountregister *CThostFtdcQryAccountregisterField, nRequestID int) int{
	return int(C.tReqQryAccountregister(t.api, (*C.struct_CThostFtdcQryAccountregisterField)(unsafe.Pointer(pQryAccountregister)), C.int(nRequestID)))
}
// 请求查询签约银行
func (t *Trade)ReqQryContractBank(pQryContractBank *CThostFtdcQryContractBankField, nRequestID int) int{
	return int(C.tReqQryContractBank(t.api, (*C.struct_CThostFtdcQryContractBankField)(unsafe.Pointer(pQryContractBank)), C.int(nRequestID)))
}
// 请求查询预埋单
func (t *Trade)ReqQryParkedOrder(pQryParkedOrder *CThostFtdcQryParkedOrderField, nRequestID int) int{
	return int(C.tReqQryParkedOrder(t.api, (*C.struct_CThostFtdcQryParkedOrderField)(unsafe.Pointer(pQryParkedOrder)), C.int(nRequestID)))
}
// 请求查询预埋撤单
func (t *Trade)ReqQryParkedOrderAction(pQryParkedOrderAction *CThostFtdcQryParkedOrderActionField, nRequestID int) int{
	return int(C.tReqQryParkedOrderAction(t.api, (*C.struct_CThostFtdcQryParkedOrderActionField)(unsafe.Pointer(pQryParkedOrderAction)), C.int(nRequestID)))
}
// 请求查询交易通知
func (t *Trade)ReqQryTradingNotice(pQryTradingNotice *CThostFtdcQryTradingNoticeField, nRequestID int) int{
	return int(C.tReqQryTradingNotice(t.api, (*C.struct_CThostFtdcQryTradingNoticeField)(unsafe.Pointer(pQryTradingNotice)), C.int(nRequestID)))
}
// 请求查询经纪公司交易参数
func (t *Trade)ReqQryBrokerTradingParams(pQryBrokerTradingParams *CThostFtdcQryBrokerTradingParamsField, nRequestID int) int{
	return int(C.tReqQryBrokerTradingParams(t.api, (*C.struct_CThostFtdcQryBrokerTradingParamsField)(unsafe.Pointer(pQryBrokerTradingParams)), C.int(nRequestID)))
}
// 请求查询经纪公司交易算法
func (t *Trade)ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos *CThostFtdcQryBrokerTradingAlgosField, nRequestID int) int{
	return int(C.tReqQryBrokerTradingAlgos(t.api, (*C.struct_CThostFtdcQryBrokerTradingAlgosField)(unsafe.Pointer(pQryBrokerTradingAlgos)), C.int(nRequestID)))
}
// 请求查询监控中心用户令牌
func (t *Trade)ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int) int{
	return int(C.tReqQueryCFMMCTradingAccountToken(t.api, (*C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)), C.int(nRequestID)))
}
// 期货发起银行资金转期货请求
func (t *Trade)ReqFromBankToFutureByFuture(pReqTransfer *CThostFtdcReqTransferField, nRequestID int) int{
	return int(C.tReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), C.int(nRequestID)))
}
// 期货发起期货资金转银行请求
func (t *Trade)ReqFromFutureToBankByFuture(pReqTransfer *CThostFtdcReqTransferField, nRequestID int) int{
	return int(C.tReqFromFutureToBankByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), C.int(nRequestID)))
}
// 期货发起查询银行余额请求
func (t *Trade)ReqQueryBankAccountMoneyByFuture(pReqQueryAccount *CThostFtdcReqQueryAccountField, nRequestID int) int{
	return int(C.tReqQueryBankAccountMoneyByFuture(t.api, (*C.struct_CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), C.int(nRequestID)))
}
// 请求查询分类合约
func (t *Trade)ReqQryClassifiedInstrument(pQryClassifiedInstrument *CThostFtdcQryClassifiedInstrumentField, nRequestID int) int{
	return int(C.tReqQryClassifiedInstrument(t.api, (*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(pQryClassifiedInstrument)), C.int(nRequestID)))
}
// 请求组合优惠比例
func (t *Trade)ReqQryCombPromotionParam(pQryCombPromotionParam *CThostFtdcQryCombPromotionParamField, nRequestID int) int{
	return int(C.tReqQryCombPromotionParam(t.api, (*C.struct_CThostFtdcQryCombPromotionParamField)(unsafe.Pointer(pQryCombPromotionParam)), C.int(nRequestID)))
}
// 投资者风险结算持仓查询
func (t *Trade)ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition *CThostFtdcQryRiskSettleInvstPositionField, nRequestID int) int{
	return int(C.tReqQryRiskSettleInvstPosition(t.api, (*C.struct_CThostFtdcQryRiskSettleInvstPositionField)(unsafe.Pointer(pQryRiskSettleInvstPosition)), C.int(nRequestID)))
}
// 风险结算产品查询
func (t *Trade)ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus *CThostFtdcQryRiskSettleProductStatusField, nRequestID int) int{
	return int(C.tReqQryRiskSettleProductStatus(t.api, (*C.struct_CThostFtdcQryRiskSettleProductStatusField)(unsafe.Pointer(pQryRiskSettleProductStatus)), C.int(nRequestID)))
}
// SPBM期货合约参数查询
func (t *Trade)ReqQrySPBMFutureParameter(pQrySPBMFutureParameter *CThostFtdcQrySPBMFutureParameterField, nRequestID int) int{
	return int(C.tReqQrySPBMFutureParameter(t.api, (*C.struct_CThostFtdcQrySPBMFutureParameterField)(unsafe.Pointer(pQrySPBMFutureParameter)), C.int(nRequestID)))
}
// SPBM期权合约参数查询
func (t *Trade)ReqQrySPBMOptionParameter(pQrySPBMOptionParameter *CThostFtdcQrySPBMOptionParameterField, nRequestID int) int{
	return int(C.tReqQrySPBMOptionParameter(t.api, (*C.struct_CThostFtdcQrySPBMOptionParameterField)(unsafe.Pointer(pQrySPBMOptionParameter)), C.int(nRequestID)))
}
// SPBM品种内对锁仓折扣参数查询
func (t *Trade)ReqQrySPBMIntraParameter(pQrySPBMIntraParameter *CThostFtdcQrySPBMIntraParameterField, nRequestID int) int{
	return int(C.tReqQrySPBMIntraParameter(t.api, (*C.struct_CThostFtdcQrySPBMIntraParameterField)(unsafe.Pointer(pQrySPBMIntraParameter)), C.int(nRequestID)))
}
// SPBM跨品种抵扣参数查询
func (t *Trade)ReqQrySPBMInterParameter(pQrySPBMInterParameter *CThostFtdcQrySPBMInterParameterField, nRequestID int) int{
	return int(C.tReqQrySPBMInterParameter(t.api, (*C.struct_CThostFtdcQrySPBMInterParameterField)(unsafe.Pointer(pQrySPBMInterParameter)), C.int(nRequestID)))
}
// SPBM组合保证金套餐查询
func (t *Trade)ReqQrySPBMPortfDefinition(pQrySPBMPortfDefinition *CThostFtdcQrySPBMPortfDefinitionField, nRequestID int) int{
	return int(C.tReqQrySPBMPortfDefinition(t.api, (*C.struct_CThostFtdcQrySPBMPortfDefinitionField)(unsafe.Pointer(pQrySPBMPortfDefinition)), C.int(nRequestID)))
}
// 投资者SPBM套餐选择查询
func (t *Trade)ReqQrySPBMInvestorPortfDef(pQrySPBMInvestorPortfDef *CThostFtdcQrySPBMInvestorPortfDefField, nRequestID int) int{
	return int(C.tReqQrySPBMInvestorPortfDef(t.api, (*C.struct_CThostFtdcQrySPBMInvestorPortfDefField)(unsafe.Pointer(pQrySPBMInvestorPortfDef)), C.int(nRequestID)))
}
// 投资者新型组合保证金系数查询
func (t *Trade)ReqQryInvestorPortfMarginRatio(pQryInvestorPortfMarginRatio *CThostFtdcQryInvestorPortfMarginRatioField, nRequestID int) int{
	return int(C.tReqQryInvestorPortfMarginRatio(t.api, (*C.struct_CThostFtdcQryInvestorPortfMarginRatioField)(unsafe.Pointer(pQryInvestorPortfMarginRatio)), C.int(nRequestID)))
}
// 投资者产品SPBM明细查询
func (t *Trade)ReqQryInvestorProdSPBMDetail(pQryInvestorProdSPBMDetail *CThostFtdcQryInvestorProdSPBMDetailField, nRequestID int) int{
	return int(C.tReqQryInvestorProdSPBMDetail(t.api, (*C.struct_CThostFtdcQryInvestorProdSPBMDetailField)(unsafe.Pointer(pQryInvestorProdSPBMDetail)), C.int(nRequestID)))
}
