package trade

/*
#cgo CPPFLAGS: -fPIC -I../CTPv6.6.8_20220712
#include "ThostFtdcUserApiDataType.h"
#include "ThostFtdcUserApiStruct.h"

#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/../lib -Wl,-rpath ${SRCDIR}/../lib -l ctptrade -lstdc++

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

// //////////////////////////////////////////////////////////////////////
void tSetOnFrontConnected(void *, void *);
void OnFrontConnected();
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
void tSetOnFrontDisconnected(void *, void *);
void OnFrontDisconnected(int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
void tSetOnHeartBeatWarning(void *, void *);
void OnHeartBeatWarning(int nTimeLapse);
// 客户端认证响应
void tSetOnRspAuthenticate(void *, void *);
void OnRspAuthenticate(struct CThostFtdcRspAuthenticateField *pRspAuthenticateField, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登录请求响应
void tSetOnRspUserLogin(void *, void *);
void OnRspUserLogin(struct CThostFtdcRspUserLoginField *pRspUserLogin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 登出请求响应
void tSetOnRspUserLogout(void *, void *);
void OnRspUserLogout(struct CThostFtdcUserLogoutField *pUserLogout, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 用户口令更新请求响应
void tSetOnRspUserPasswordUpdate(void *, void *);
void OnRspUserPasswordUpdate(struct CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 资金账户口令更新请求响应
void tSetOnRspTradingAccountPasswordUpdate(void *, void *);
void OnRspTradingAccountPasswordUpdate(struct CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询用户当前支持的认证模式的回复
void tSetOnRspUserAuthMethod(void *, void *);
void OnRspUserAuthMethod(struct CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 获取图形验证码请求的回复
void tSetOnRspGenUserCaptcha(void *, void *);
void OnRspGenUserCaptcha(struct CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 获取短信验证码请求的回复
void tSetOnRspGenUserText(void *, void *);
void OnRspGenUserText(struct CThostFtdcRspGenUserTextField *pRspGenUserText, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单录入请求响应
void tSetOnRspOrderInsert(void *, void *);
void OnRspOrderInsert(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 预埋单录入请求响应
void tSetOnRspParkedOrderInsert(void *, void *);
void OnRspParkedOrderInsert(struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 预埋撤单录入请求响应
void tSetOnRspParkedOrderAction(void *, void *);
void OnRspParkedOrderAction(struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单操作请求响应
void tSetOnRspOrderAction(void *, void *);
void OnRspOrderAction(struct CThostFtdcInputOrderActionField *pInputOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询最大报单数量响应
void tSetOnRspQryMaxOrderVolume(void *, void *);
void OnRspQryMaxOrderVolume(struct CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者结算结果确认响应
void tSetOnRspSettlementInfoConfirm(void *, void *);
void OnRspSettlementInfoConfirm(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 删除预埋单响应
void tSetOnRspRemoveParkedOrder(void *, void *);
void OnRspRemoveParkedOrder(struct CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 删除预埋撤单响应
void tSetOnRspRemoveParkedOrderAction(void *, void *);
void OnRspRemoveParkedOrderAction(struct CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 执行宣告录入请求响应
void tSetOnRspExecOrderInsert(void *, void *);
void OnRspExecOrderInsert(struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 执行宣告操作请求响应
void tSetOnRspExecOrderAction(void *, void *);
void OnRspExecOrderAction(struct CThostFtdcInputExecOrderActionField *pInputExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 询价录入请求响应
void tSetOnRspForQuoteInsert(void *, void *);
void OnRspForQuoteInsert(struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报价录入请求响应
void tSetOnRspQuoteInsert(void *, void *);
void OnRspQuoteInsert(struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报价操作请求响应
void tSetOnRspQuoteAction(void *, void *);
void OnRspQuoteAction(struct CThostFtdcInputQuoteActionField *pInputQuoteAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 批量报单操作请求响应
void tSetOnRspBatchOrderAction(void *, void *);
void OnRspBatchOrderAction(struct CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期权自对冲录入请求响应
void tSetOnRspOptionSelfCloseInsert(void *, void *);
void OnRspOptionSelfCloseInsert(struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期权自对冲操作请求响应
void tSetOnRspOptionSelfCloseAction(void *, void *);
void OnRspOptionSelfCloseAction(struct CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 申请组合录入请求响应
void tSetOnRspCombActionInsert(void *, void *);
void OnRspCombActionInsert(struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报单响应
void tSetOnRspQryOrder(void *, void *);
void OnRspQryOrder(struct CThostFtdcOrderField *pOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询成交响应
void tSetOnRspQryTrade(void *, void *);
void OnRspQryTrade(struct CThostFtdcTradeField *pTrade, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓响应
void tSetOnRspQryInvestorPosition(void *, void *);
void OnRspQryInvestorPosition(struct CThostFtdcInvestorPositionField *pInvestorPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询资金账户响应
void tSetOnRspQryTradingAccount(void *, void *);
void OnRspQryTradingAccount(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者响应
void tSetOnRspQryInvestor(void *, void *);
void OnRspQryInvestor(struct CThostFtdcInvestorField *pInvestor, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易编码响应
void tSetOnRspQryTradingCode(void *, void *);
void OnRspQryTradingCode(struct CThostFtdcTradingCodeField *pTradingCode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约保证金率响应
void tSetOnRspQryInstrumentMarginRate(void *, void *);
void OnRspQryInstrumentMarginRate(struct CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约手续费率响应
void tSetOnRspQryInstrumentCommissionRate(void *, void *);
void OnRspQryInstrumentCommissionRate(struct CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所响应
void tSetOnRspQryExchange(void *, void *);
void OnRspQryExchange(struct CThostFtdcExchangeField *pExchange, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品响应
void tSetOnRspQryProduct(void *, void *);
void OnRspQryProduct(struct CThostFtdcProductField *pProduct, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询合约响应
void tSetOnRspQryInstrument(void *, void *);
void OnRspQryInstrument(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询行情响应
void tSetOnRspQryDepthMarketData(void *, void *);
void OnRspQryDepthMarketData(struct CThostFtdcDepthMarketDataField *pDepthMarketData, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易员报盘机响应
void tSetOnRspQryTraderOffer(void *, void *);
void OnRspQryTraderOffer(struct CThostFtdcTraderOfferField *pTraderOffer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者结算结果响应
void tSetOnRspQrySettlementInfo(void *, void *);
void OnRspQrySettlementInfo(struct CThostFtdcSettlementInfoField *pSettlementInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询转帐银行响应
void tSetOnRspQryTransferBank(void *, void *);
void OnRspQryTransferBank(struct CThostFtdcTransferBankField *pTransferBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓明细响应
void tSetOnRspQryInvestorPositionDetail(void *, void *);
void OnRspQryInvestorPositionDetail(struct CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询客户通知响应
void tSetOnRspQryNotice(void *, void *);
void OnRspQryNotice(struct CThostFtdcNoticeField *pNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询结算信息确认响应
void tSetOnRspQrySettlementInfoConfirm(void *, void *);
void OnRspQrySettlementInfoConfirm(struct CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者持仓明细响应
void tSetOnRspQryInvestorPositionCombineDetail(void *, void *);
void OnRspQryInvestorPositionCombineDetail(struct CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 查询保证金监管系统经纪公司资金账户密钥响应
void tSetOnRspQryCFMMCTradingAccountKey(void *, void *);
void OnRspQryCFMMCTradingAccountKey(struct CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询仓单折抵信息响应
void tSetOnRspQryEWarrantOffset(void *, void *);
void OnRspQryEWarrantOffset(struct CThostFtdcEWarrantOffsetField *pEWarrantOffset, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资者品种/跨品种保证金响应
void tSetOnRspQryInvestorProductGroupMargin(void *, void *);
void OnRspQryInvestorProductGroupMargin(struct CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所保证金率响应
void tSetOnRspQryExchangeMarginRate(void *, void *);
void OnRspQryExchangeMarginRate(struct CThostFtdcExchangeMarginRateField *pExchangeMarginRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易所调整保证金率响应
void tSetOnRspQryExchangeMarginRateAdjust(void *, void *);
void OnRspQryExchangeMarginRateAdjust(struct CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询汇率响应
void tSetOnRspQryExchangeRate(void *, void *);
void OnRspQryExchangeRate(struct CThostFtdcExchangeRateField *pExchangeRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理操作员银期权限响应
void tSetOnRspQrySecAgentACIDMap(void *, void *);
void OnRspQrySecAgentACIDMap(struct CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品报价汇率
void tSetOnRspQryProductExchRate(void *, void *);
void OnRspQryProductExchRate(struct CThostFtdcProductExchRateField *pProductExchRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询产品组
void tSetOnRspQryProductGroup(void *, void *);
void OnRspQryProductGroup(struct CThostFtdcProductGroupField *pProductGroup, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询做市商合约手续费率响应
void tSetOnRspQryMMInstrumentCommissionRate(void *, void *);
void OnRspQryMMInstrumentCommissionRate(struct CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询做市商期权合约手续费响应
void tSetOnRspQryMMOptionInstrCommRate(void *, void *);
void OnRspQryMMOptionInstrCommRate(struct CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报单手续费响应
void tSetOnRspQryInstrumentOrderCommRate(void *, void *);
void OnRspQryInstrumentOrderCommRate(struct CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询资金账户响应
void tSetOnRspQrySecAgentTradingAccount(void *, void *);
void OnRspQrySecAgentTradingAccount(struct CThostFtdcTradingAccountField *pTradingAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理商资金校验模式响应
void tSetOnRspQrySecAgentCheckMode(void *, void *);
void OnRspQrySecAgentCheckMode(struct CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询二级代理商信息响应
void tSetOnRspQrySecAgentTradeInfo(void *, void *);
void OnRspQrySecAgentTradeInfo(struct CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权交易成本响应
void tSetOnRspQryOptionInstrTradeCost(void *, void *);
void OnRspQryOptionInstrTradeCost(struct CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权合约手续费响应
void tSetOnRspQryOptionInstrCommRate(void *, void *);
void OnRspQryOptionInstrCommRate(struct CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询执行宣告响应
void tSetOnRspQryExecOrder(void *, void *);
void OnRspQryExecOrder(struct CThostFtdcExecOrderField *pExecOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询询价响应
void tSetOnRspQryForQuote(void *, void *);
void OnRspQryForQuote(struct CThostFtdcForQuoteField *pForQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询报价响应
void tSetOnRspQryQuote(void *, void *);
void OnRspQryQuote(struct CThostFtdcQuoteField *pQuote, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询期权自对冲响应
void tSetOnRspQryOptionSelfClose(void *, void *);
void OnRspQryOptionSelfClose(struct CThostFtdcOptionSelfCloseField *pOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询投资单元响应
void tSetOnRspQryInvestUnit(void *, void *);
void OnRspQryInvestUnit(struct CThostFtdcInvestUnitField *pInvestUnit, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询组合合约安全系数响应
void tSetOnRspQryCombInstrumentGuard(void *, void *);
void OnRspQryCombInstrumentGuard(struct CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询申请组合响应
void tSetOnRspQryCombAction(void *, void *);
void OnRspQryCombAction(struct CThostFtdcCombActionField *pCombAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询转帐流水响应
void tSetOnRspQryTransferSerial(void *, void *);
void OnRspQryTransferSerial(struct CThostFtdcTransferSerialField *pTransferSerial, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询银期签约关系响应
void tSetOnRspQryAccountregister(void *, void *);
void OnRspQryAccountregister(struct CThostFtdcAccountregisterField *pAccountregister, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 错误应答
void tSetOnRspError(void *, void *);
void OnRspError(struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 报单通知
void tSetOnRtnOrder(void *, void *);
void OnRtnOrder(struct CThostFtdcOrderField *pOrder);
// 成交通知
void tSetOnRtnTrade(void *, void *);
void OnRtnTrade(struct CThostFtdcTradeField *pTrade);
// 报单录入错误回报
void tSetOnErrRtnOrderInsert(void *, void *);
void OnErrRtnOrderInsert(struct CThostFtdcInputOrderField *pInputOrder, struct CThostFtdcRspInfoField *pRspInfo);
// 报单操作错误回报
void tSetOnErrRtnOrderAction(void *, void *);
void OnErrRtnOrderAction(struct CThostFtdcOrderActionField *pOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 合约交易状态通知
void tSetOnRtnInstrumentStatus(void *, void *);
void OnRtnInstrumentStatus(struct CThostFtdcInstrumentStatusField *pInstrumentStatus);
// 交易所公告通知
void tSetOnRtnBulletin(void *, void *);
void OnRtnBulletin(struct CThostFtdcBulletinField *pBulletin);
// 交易通知
void tSetOnRtnTradingNotice(void *, void *);
void OnRtnTradingNotice(struct CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);
// 提示条件单校验错误
void tSetOnRtnErrorConditionalOrder(void *, void *);
void OnRtnErrorConditionalOrder(struct CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);
// 执行宣告通知
void tSetOnRtnExecOrder(void *, void *);
void OnRtnExecOrder(struct CThostFtdcExecOrderField *pExecOrder);
// 执行宣告录入错误回报
void tSetOnErrRtnExecOrderInsert(void *, void *);
void OnErrRtnExecOrderInsert(struct CThostFtdcInputExecOrderField *pInputExecOrder, struct CThostFtdcRspInfoField *pRspInfo);
// 执行宣告操作错误回报
void tSetOnErrRtnExecOrderAction(void *, void *);
void OnErrRtnExecOrderAction(struct CThostFtdcExecOrderActionField *pExecOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 询价录入错误回报
void tSetOnErrRtnForQuoteInsert(void *, void *);
void OnErrRtnForQuoteInsert(struct CThostFtdcInputForQuoteField *pInputForQuote, struct CThostFtdcRspInfoField *pRspInfo);
// 报价通知
void tSetOnRtnQuote(void *, void *);
void OnRtnQuote(struct CThostFtdcQuoteField *pQuote);
// 报价录入错误回报
void tSetOnErrRtnQuoteInsert(void *, void *);
void OnErrRtnQuoteInsert(struct CThostFtdcInputQuoteField *pInputQuote, struct CThostFtdcRspInfoField *pRspInfo);
// 报价操作错误回报
void tSetOnErrRtnQuoteAction(void *, void *);
void OnErrRtnQuoteAction(struct CThostFtdcQuoteActionField *pQuoteAction, struct CThostFtdcRspInfoField *pRspInfo);
// 询价通知
void tSetOnRtnForQuoteRsp(void *, void *);
void OnRtnForQuoteRsp(struct CThostFtdcForQuoteRspField *pForQuoteRsp);
// 保证金监控中心用户令牌
void tSetOnRtnCFMMCTradingAccountToken(void *, void *);
void OnRtnCFMMCTradingAccountToken(struct CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);
// 批量报单操作错误回报
void tSetOnErrRtnBatchOrderAction(void *, void *);
void OnErrRtnBatchOrderAction(struct CThostFtdcBatchOrderActionField *pBatchOrderAction, struct CThostFtdcRspInfoField *pRspInfo);
// 期权自对冲通知
void tSetOnRtnOptionSelfClose(void *, void *);
void OnRtnOptionSelfClose(struct CThostFtdcOptionSelfCloseField *pOptionSelfClose);
// 期权自对冲录入错误回报
void tSetOnErrRtnOptionSelfCloseInsert(void *, void *);
void OnErrRtnOptionSelfCloseInsert(struct CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, struct CThostFtdcRspInfoField *pRspInfo);
// 期权自对冲操作错误回报
void tSetOnErrRtnOptionSelfCloseAction(void *, void *);
void OnErrRtnOptionSelfCloseAction(struct CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, struct CThostFtdcRspInfoField *pRspInfo);
// 申请组合通知
void tSetOnRtnCombAction(void *, void *);
void OnRtnCombAction(struct CThostFtdcCombActionField *pCombAction);
// 申请组合录入错误回报
void tSetOnErrRtnCombActionInsert(void *, void *);
void OnErrRtnCombActionInsert(struct CThostFtdcInputCombActionField *pInputCombAction, struct CThostFtdcRspInfoField *pRspInfo);
// 请求查询签约银行响应
void tSetOnRspQryContractBank(void *, void *);
void OnRspQryContractBank(struct CThostFtdcContractBankField *pContractBank, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询预埋单响应
void tSetOnRspQryParkedOrder(void *, void *);
void OnRspQryParkedOrder(struct CThostFtdcParkedOrderField *pParkedOrder, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询预埋撤单响应
void tSetOnRspQryParkedOrderAction(void *, void *);
void OnRspQryParkedOrderAction(struct CThostFtdcParkedOrderActionField *pParkedOrderAction, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询交易通知响应
void tSetOnRspQryTradingNotice(void *, void *);
void OnRspQryTradingNotice(struct CThostFtdcTradingNoticeField *pTradingNotice, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询经纪公司交易参数响应
void tSetOnRspQryBrokerTradingParams(void *, void *);
void OnRspQryBrokerTradingParams(struct CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询经纪公司交易算法响应
void tSetOnRspQryBrokerTradingAlgos(void *, void *);
void OnRspQryBrokerTradingAlgos(struct CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求查询监控中心用户令牌
void tSetOnRspQueryCFMMCTradingAccountToken(void *, void *);
void OnRspQueryCFMMCTradingAccountToken(struct CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 银行发起银行资金转期货通知
void tSetOnRtnFromBankToFutureByBank(void *, void *);
void OnRtnFromBankToFutureByBank(struct CThostFtdcRspTransferField *pRspTransfer);
// 银行发起期货资金转银行通知
void tSetOnRtnFromFutureToBankByBank(void *, void *);
void OnRtnFromFutureToBankByBank(struct CThostFtdcRspTransferField *pRspTransfer);
// 银行发起冲正银行转期货通知
void tSetOnRtnRepealFromBankToFutureByBank(void *, void *);
void OnRtnRepealFromBankToFutureByBank(struct CThostFtdcRspRepealField *pRspRepeal);
// 银行发起冲正期货转银行通知
void tSetOnRtnRepealFromFutureToBankByBank(void *, void *);
void OnRtnRepealFromFutureToBankByBank(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起银行资金转期货通知
void tSetOnRtnFromBankToFutureByFuture(void *, void *);
void OnRtnFromBankToFutureByFuture(struct CThostFtdcRspTransferField *pRspTransfer);
// 期货发起期货资金转银行通知
void tSetOnRtnFromFutureToBankByFuture(void *, void *);
void OnRtnFromFutureToBankByFuture(struct CThostFtdcRspTransferField *pRspTransfer);
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromBankToFutureByFutureManual(void *, void *);
void OnRtnRepealFromBankToFutureByFutureManual(struct CThostFtdcRspRepealField *pRspRepeal);
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromFutureToBankByFutureManual(void *, void *);
void OnRtnRepealFromFutureToBankByFutureManual(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起查询银行余额通知
void tSetOnRtnQueryBankBalanceByFuture(void *, void *);
void OnRtnQueryBankBalanceByFuture(struct CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);
// 期货发起银行资金转期货错误回报
void tSetOnErrRtnBankToFutureByFuture(void *, void *);
void OnErrRtnBankToFutureByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起期货资金转银行错误回报
void tSetOnErrRtnFutureToBankByFuture(void *, void *);
void OnErrRtnFutureToBankByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo);
// 系统运行时期货端手工发起冲正银行转期货错误回报
void tSetOnErrRtnRepealBankToFutureByFutureManual(void *, void *);
void OnErrRtnRepealBankToFutureByFutureManual(struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);
// 系统运行时期货端手工发起冲正期货转银行错误回报
void tSetOnErrRtnRepealFutureToBankByFutureManual(void *, void *);
void OnErrRtnRepealFutureToBankByFutureManual(struct CThostFtdcReqRepealField *pReqRepeal, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起查询银行余额错误回报
void tSetOnErrRtnQueryBankBalanceByFuture(void *, void *);
void OnErrRtnQueryBankBalanceByFuture(struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo);
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromBankToFutureByFuture(void *, void *);
void OnRtnRepealFromBankToFutureByFuture(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
void tSetOnRtnRepealFromFutureToBankByFuture(void *, void *);
void OnRtnRepealFromFutureToBankByFuture(struct CThostFtdcRspRepealField *pRspRepeal);
// 期货发起银行资金转期货应答
void tSetOnRspFromBankToFutureByFuture(void *, void *);
void OnRspFromBankToFutureByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期货发起期货资金转银行应答
void tSetOnRspFromFutureToBankByFuture(void *, void *);
void OnRspFromFutureToBankByFuture(struct CThostFtdcReqTransferField *pReqTransfer, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 期货发起查询银行余额应答
void tSetOnRspQueryBankAccountMoneyByFuture(void *, void *);
void OnRspQueryBankAccountMoneyByFuture(struct CThostFtdcReqQueryAccountField *pReqQueryAccount, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 银行发起银期开户通知
void tSetOnRtnOpenAccountByBank(void *, void *);
void OnRtnOpenAccountByBank(struct CThostFtdcOpenAccountField *pOpenAccount);
// 银行发起银期销户通知
void tSetOnRtnCancelAccountByBank(void *, void *);
void OnRtnCancelAccountByBank(struct CThostFtdcCancelAccountField *pCancelAccount);
// 银行发起变更银行账号通知
void tSetOnRtnChangeAccountByBank(void *, void *);
void OnRtnChangeAccountByBank(struct CThostFtdcChangeAccountField *pChangeAccount);
// 请求查询分类合约响应
void tSetOnRspQryClassifiedInstrument(void *, void *);
void OnRspQryClassifiedInstrument(struct CThostFtdcInstrumentField *pInstrument, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 请求组合优惠比例响应
void tSetOnRspQryCombPromotionParam(void *, void *);
void OnRspQryCombPromotionParam(struct CThostFtdcCombPromotionParamField *pCombPromotionParam, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 投资者风险结算持仓查询响应
void tSetOnRspQryRiskSettleInvstPosition(void *, void *);
void OnRspQryRiskSettleInvstPosition(struct CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);
// 风险结算产品查询响应
void tSetOnRspQryRiskSettleProductStatus(void *, void *);
void OnRspQryRiskSettleProductStatus(struct CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, struct CThostFtdcRspInfoField *pRspInfo, int nRequestID, _Bool bIsLast);


#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"

	"gitee.com/haifengat/goctp/def"
)

type Trade struct {
	api, spi unsafe.Pointer
	Version  string

	// ************ 响应函数变量 ******************
	// //////////////////////////////////////////////////////////////////////
	OnFrontConnected                          func()                                                                                                                                                     // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	OnFrontDisconnected                       func(nReason int)                                                                                                                                          // 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnHeartBeatWarning                        func(nTimeLapse int)                                                                                                                                       // 客户端认证响应
	OnRspAuthenticate                         func(pRspAuthenticateField *def.CThostFtdcRspAuthenticateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                        // 登录请求响应
	OnRspUserLogin                            func(pRspUserLogin *def.CThostFtdcRspUserLoginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                   // 登出请求响应
	OnRspUserLogout                           func(pUserLogout *def.CThostFtdcUserLogoutField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 用户口令更新请求响应
	OnRspUserPasswordUpdate                   func(pUserPasswordUpdate *def.CThostFtdcUserPasswordUpdateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                       // 资金账户口令更新请求响应
	OnRspTradingAccountPasswordUpdate         func(pTradingAccountPasswordUpdate *def.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)   // 查询用户当前支持的认证模式的回复
	OnRspUserAuthMethod                       func(pRspUserAuthMethod *def.CThostFtdcRspUserAuthMethodField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 获取图形验证码请求的回复
	OnRspGenUserCaptcha                       func(pRspGenUserCaptcha *def.CThostFtdcRspGenUserCaptchaField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 获取短信验证码请求的回复
	OnRspGenUserText                          func(pRspGenUserText *def.CThostFtdcRspGenUserTextField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 报单录入请求响应
	OnRspOrderInsert                          func(pInputOrder *def.CThostFtdcInputOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 预埋单录入请求响应
	OnRspParkedOrderInsert                    func(pParkedOrder *def.CThostFtdcParkedOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                     // 预埋撤单录入请求响应
	OnRspParkedOrderAction                    func(pParkedOrderAction *def.CThostFtdcParkedOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 报单操作请求响应
	OnRspOrderAction                          func(pInputOrderAction *def.CThostFtdcInputOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                           // 查询最大报单数量响应
	OnRspQryMaxOrderVolume                    func(pQryMaxOrderVolume *def.CThostFtdcQryMaxOrderVolumeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 投资者结算结果确认响应
	OnRspSettlementInfoConfirm                func(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                 // 删除预埋单响应
	OnRspRemoveParkedOrder                    func(pRemoveParkedOrder *def.CThostFtdcRemoveParkedOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 删除预埋撤单响应
	OnRspRemoveParkedOrderAction              func(pRemoveParkedOrderAction *def.CThostFtdcRemoveParkedOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)             // 执行宣告录入请求响应
	OnRspExecOrderInsert                      func(pInputExecOrder *def.CThostFtdcInputExecOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 执行宣告操作请求响应
	OnRspExecOrderAction                      func(pInputExecOrderAction *def.CThostFtdcInputExecOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                   // 询价录入请求响应
	OnRspForQuoteInsert                       func(pInputForQuote *def.CThostFtdcInputForQuoteField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                 // 报价录入请求响应
	OnRspQuoteInsert                          func(pInputQuote *def.CThostFtdcInputQuoteField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 报价操作请求响应
	OnRspQuoteAction                          func(pInputQuoteAction *def.CThostFtdcInputQuoteActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                           // 批量报单操作请求响应
	OnRspBatchOrderAction                     func(pInputBatchOrderAction *def.CThostFtdcInputBatchOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                 // 期权自对冲录入请求响应
	OnRspOptionSelfCloseInsert                func(pInputOptionSelfClose *def.CThostFtdcInputOptionSelfCloseField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                   // 期权自对冲操作请求响应
	OnRspOptionSelfCloseAction                func(pInputOptionSelfCloseAction *def.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)       // 申请组合录入请求响应
	OnRspCombActionInsert                     func(pInputCombAction *def.CThostFtdcInputCombActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 请求查询报单响应
	OnRspQryOrder                             func(pOrder *def.CThostFtdcOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                                 // 请求查询成交响应
	OnRspQryTrade                             func(pTrade *def.CThostFtdcTradeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                                 // 请求查询投资者持仓响应
	OnRspQryInvestorPosition                  func(pInvestorPosition *def.CThostFtdcInvestorPositionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                           // 请求查询资金账户响应
	OnRspQryTradingAccount                    func(pTradingAccount *def.CThostFtdcTradingAccountField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 请求查询投资者响应
	OnRspQryInvestor                          func(pInvestor *def.CThostFtdcInvestorField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                           // 请求查询交易编码响应
	OnRspQryTradingCode                       func(pTradingCode *def.CThostFtdcTradingCodeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                     // 请求查询合约保证金率响应
	OnRspQryInstrumentMarginRate              func(pInstrumentMarginRate *def.CThostFtdcInstrumentMarginRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                   // 请求查询合约手续费率响应
	OnRspQryInstrumentCommissionRate          func(pInstrumentCommissionRate *def.CThostFtdcInstrumentCommissionRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)           // 请求查询交易所响应
	OnRspQryExchange                          func(pExchange *def.CThostFtdcExchangeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                           // 请求查询产品响应
	OnRspQryProduct                           func(pProduct *def.CThostFtdcProductField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                             // 请求查询合约响应
	OnRspQryInstrument                        func(pInstrument *def.CThostFtdcInstrumentField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 请求查询行情响应
	OnRspQryDepthMarketData                   func(pDepthMarketData *def.CThostFtdcDepthMarketDataField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 请求查询交易员报盘机响应
	OnRspQryTraderOffer                       func(pTraderOffer *def.CThostFtdcTraderOfferField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                     // 请求查询投资者结算结果响应
	OnRspQrySettlementInfo                    func(pSettlementInfo *def.CThostFtdcSettlementInfoField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 请求查询转帐银行响应
	OnRspQryTransferBank                      func(pTransferBank *def.CThostFtdcTransferBankField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                   // 请求查询投资者持仓明细响应
	OnRspQryInvestorPositionDetail            func(pInvestorPositionDetail *def.CThostFtdcInvestorPositionDetailField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)               // 请求查询客户通知响应
	OnRspQryNotice                            func(pNotice *def.CThostFtdcNoticeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                               // 请求查询结算信息确认响应
	OnRspQrySettlementInfoConfirm             func(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                 // 请求查询投资者持仓明细响应
	OnRspQryInvestorPositionCombineDetail     func(pInvestorPositionCombineDetail *def.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) // 查询保证金监管系统经纪公司资金账户密钥响应
	OnRspQryCFMMCTradingAccountKey            func(pCFMMCTradingAccountKey *def.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)               // 请求查询仓单折抵信息响应
	OnRspQryEWarrantOffset                    func(pEWarrantOffset *def.CThostFtdcEWarrantOffsetField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 请求查询投资者品种/跨品种保证金响应
	OnRspQryInvestorProductGroupMargin        func(pInvestorProductGroupMargin *def.CThostFtdcInvestorProductGroupMarginField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)       // 请求查询交易所保证金率响应
	OnRspQryExchangeMarginRate                func(pExchangeMarginRate *def.CThostFtdcExchangeMarginRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                       // 请求查询交易所调整保证金率响应
	OnRspQryExchangeMarginRateAdjust          func(pExchangeMarginRateAdjust *def.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)           // 请求查询汇率响应
	OnRspQryExchangeRate                      func(pExchangeRate *def.CThostFtdcExchangeRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                   // 请求查询二级代理操作员银期权限响应
	OnRspQrySecAgentACIDMap                   func(pSecAgentACIDMap *def.CThostFtdcSecAgentACIDMapField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 请求查询产品报价汇率
	OnRspQryProductExchRate                   func(pProductExchRate *def.CThostFtdcProductExchRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 请求查询产品组
	OnRspQryProductGroup                      func(pProductGroup *def.CThostFtdcProductGroupField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                   // 请求查询做市商合约手续费率响应
	OnRspQryMMInstrumentCommissionRate        func(pMMInstrumentCommissionRate *def.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)       // 请求查询做市商期权合约手续费响应
	OnRspQryMMOptionInstrCommRate             func(pMMOptionInstrCommRate *def.CThostFtdcMMOptionInstrCommRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                 // 请求查询报单手续费响应
	OnRspQryInstrumentOrderCommRate           func(pInstrumentOrderCommRate *def.CThostFtdcInstrumentOrderCommRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)             // 请求查询资金账户响应
	OnRspQrySecAgentTradingAccount            func(pTradingAccount *def.CThostFtdcTradingAccountField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 请求查询二级代理商资金校验模式响应
	OnRspQrySecAgentCheckMode                 func(pSecAgentCheckMode *def.CThostFtdcSecAgentCheckModeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 请求查询二级代理商信息响应
	OnRspQrySecAgentTradeInfo                 func(pSecAgentTradeInfo *def.CThostFtdcSecAgentTradeInfoField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 请求查询期权交易成本响应
	OnRspQryOptionInstrTradeCost              func(pOptionInstrTradeCost *def.CThostFtdcOptionInstrTradeCostField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                   // 请求查询期权合约手续费响应
	OnRspQryOptionInstrCommRate               func(pOptionInstrCommRate *def.CThostFtdcOptionInstrCommRateField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                     // 请求查询执行宣告响应
	OnRspQryExecOrder                         func(pExecOrder *def.CThostFtdcExecOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                         // 请求查询询价响应
	OnRspQryForQuote                          func(pForQuote *def.CThostFtdcForQuoteField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                           // 请求查询报价响应
	OnRspQryQuote                             func(pQuote *def.CThostFtdcQuoteField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                                 // 请求查询期权自对冲响应
	OnRspQryOptionSelfClose                   func(pOptionSelfClose *def.CThostFtdcOptionSelfCloseField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 请求查询投资单元响应
	OnRspQryInvestUnit                        func(pInvestUnit *def.CThostFtdcInvestUnitField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 请求查询组合合约安全系数响应
	OnRspQryCombInstrumentGuard               func(pCombInstrumentGuard *def.CThostFtdcCombInstrumentGuardField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                     // 请求查询申请组合响应
	OnRspQryCombAction                        func(pCombAction *def.CThostFtdcCombActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 请求查询转帐流水响应
	OnRspQryTransferSerial                    func(pTransferSerial *def.CThostFtdcTransferSerialField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                               // 请求查询银期签约关系响应
	OnRspQryAccountregister                   func(pAccountregister *def.CThostFtdcAccountregisterField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 错误应答
	OnRspError                                func(pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                                                                   // 报单通知
	OnRtnOrder                                func(pOrder *def.CThostFtdcOrderField)                                                                                                                     // 成交通知
	OnRtnTrade                                func(pTrade *def.CThostFtdcTradeField)                                                                                                                     // 报单录入错误回报
	OnErrRtnOrderInsert                       func(pInputOrder *def.CThostFtdcInputOrderField, pRspInfo *def.CThostFtdcRspInfoField)                                                                     // 报单操作错误回报
	OnErrRtnOrderAction                       func(pOrderAction *def.CThostFtdcOrderActionField, pRspInfo *def.CThostFtdcRspInfoField)                                                                   // 合约交易状态通知
	OnRtnInstrumentStatus                     func(pInstrumentStatus *def.CThostFtdcInstrumentStatusField)                                                                                               // 交易所公告通知
	OnRtnBulletin                             func(pBulletin *def.CThostFtdcBulletinField)                                                                                                               // 交易通知
	OnRtnTradingNotice                        func(pTradingNoticeInfo *def.CThostFtdcTradingNoticeInfoField)                                                                                             // 提示条件单校验错误
	OnRtnErrorConditionalOrder                func(pErrorConditionalOrder *def.CThostFtdcErrorConditionalOrderField)                                                                                     // 执行宣告通知
	OnRtnExecOrder                            func(pExecOrder *def.CThostFtdcExecOrderField)                                                                                                             // 执行宣告录入错误回报
	OnErrRtnExecOrderInsert                   func(pInputExecOrder *def.CThostFtdcInputExecOrderField, pRspInfo *def.CThostFtdcRspInfoField)                                                             // 执行宣告操作错误回报
	OnErrRtnExecOrderAction                   func(pExecOrderAction *def.CThostFtdcExecOrderActionField, pRspInfo *def.CThostFtdcRspInfoField)                                                           // 询价录入错误回报
	OnErrRtnForQuoteInsert                    func(pInputForQuote *def.CThostFtdcInputForQuoteField, pRspInfo *def.CThostFtdcRspInfoField)                                                               // 报价通知
	OnRtnQuote                                func(pQuote *def.CThostFtdcQuoteField)                                                                                                                     // 报价录入错误回报
	OnErrRtnQuoteInsert                       func(pInputQuote *def.CThostFtdcInputQuoteField, pRspInfo *def.CThostFtdcRspInfoField)                                                                     // 报价操作错误回报
	OnErrRtnQuoteAction                       func(pQuoteAction *def.CThostFtdcQuoteActionField, pRspInfo *def.CThostFtdcRspInfoField)                                                                   // 询价通知
	OnRtnForQuoteRsp                          func(pForQuoteRsp *def.CThostFtdcForQuoteRspField)                                                                                                         // 保证金监控中心用户令牌
	OnRtnCFMMCTradingAccountToken             func(pCFMMCTradingAccountToken *def.CThostFtdcCFMMCTradingAccountTokenField)                                                                               // 批量报单操作错误回报
	OnErrRtnBatchOrderAction                  func(pBatchOrderAction *def.CThostFtdcBatchOrderActionField, pRspInfo *def.CThostFtdcRspInfoField)                                                         // 期权自对冲通知
	OnRtnOptionSelfClose                      func(pOptionSelfClose *def.CThostFtdcOptionSelfCloseField)                                                                                                 // 期权自对冲录入错误回报
	OnErrRtnOptionSelfCloseInsert             func(pInputOptionSelfClose *def.CThostFtdcInputOptionSelfCloseField, pRspInfo *def.CThostFtdcRspInfoField)                                                 // 期权自对冲操作错误回报
	OnErrRtnOptionSelfCloseAction             func(pOptionSelfCloseAction *def.CThostFtdcOptionSelfCloseActionField, pRspInfo *def.CThostFtdcRspInfoField)                                               // 申请组合通知
	OnRtnCombAction                           func(pCombAction *def.CThostFtdcCombActionField)                                                                                                           // 申请组合录入错误回报
	OnErrRtnCombActionInsert                  func(pInputCombAction *def.CThostFtdcInputCombActionField, pRspInfo *def.CThostFtdcRspInfoField)                                                           // 请求查询签约银行响应
	OnRspQryContractBank                      func(pContractBank *def.CThostFtdcContractBankField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                   // 请求查询预埋单响应
	OnRspQryParkedOrder                       func(pParkedOrder *def.CThostFtdcParkedOrderField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                     // 请求查询预埋撤单响应
	OnRspQryParkedOrderAction                 func(pParkedOrderAction *def.CThostFtdcParkedOrderActionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                         // 请求查询交易通知响应
	OnRspQryTradingNotice                     func(pTradingNotice *def.CThostFtdcTradingNoticeField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                 // 请求查询经纪公司交易参数响应
	OnRspQryBrokerTradingParams               func(pBrokerTradingParams *def.CThostFtdcBrokerTradingParamsField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                     // 请求查询经纪公司交易算法响应
	OnRspQryBrokerTradingAlgos                func(pBrokerTradingAlgos *def.CThostFtdcBrokerTradingAlgosField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                       // 请求查询监控中心用户令牌
	OnRspQueryCFMMCTradingAccountToken        func(pQueryCFMMCTradingAccountToken *def.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) // 银行发起银行资金转期货通知
	OnRtnFromBankToFutureByBank               func(pRspTransfer *def.CThostFtdcRspTransferField)                                                                                                         // 银行发起期货资金转银行通知
	OnRtnFromFutureToBankByBank               func(pRspTransfer *def.CThostFtdcRspTransferField)                                                                                                         // 银行发起冲正银行转期货通知
	OnRtnRepealFromBankToFutureByBank         func(pRspRepeal *def.CThostFtdcRspRepealField)                                                                                                             // 银行发起冲正期货转银行通知
	OnRtnRepealFromFutureToBankByBank         func(pRspRepeal *def.CThostFtdcRspRepealField)                                                                                                             // 期货发起银行资金转期货通知
	OnRtnFromBankToFutureByFuture             func(pRspTransfer *def.CThostFtdcRspTransferField)                                                                                                         // 期货发起期货资金转银行通知
	OnRtnFromFutureToBankByFuture             func(pRspTransfer *def.CThostFtdcRspTransferField)                                                                                                         // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFutureManual func(pRspRepeal *def.CThostFtdcRspRepealField)                                                                                                             // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFutureManual func(pRspRepeal *def.CThostFtdcRspRepealField)                                                                                                             // 期货发起查询银行余额通知
	OnRtnQueryBankBalanceByFuture             func(pNotifyQueryAccount *def.CThostFtdcNotifyQueryAccountField)                                                                                           // 期货发起银行资金转期货错误回报
	OnErrRtnBankToFutureByFuture              func(pReqTransfer *def.CThostFtdcReqTransferField, pRspInfo *def.CThostFtdcRspInfoField)                                                                   // 期货发起期货资金转银行错误回报
	OnErrRtnFutureToBankByFuture              func(pReqTransfer *def.CThostFtdcReqTransferField, pRspInfo *def.CThostFtdcRspInfoField)                                                                   // 系统运行时期货端手工发起冲正银行转期货错误回报
	OnErrRtnRepealBankToFutureByFutureManual  func(pReqRepeal *def.CThostFtdcReqRepealField, pRspInfo *def.CThostFtdcRspInfoField)                                                                       // 系统运行时期货端手工发起冲正期货转银行错误回报
	OnErrRtnRepealFutureToBankByFutureManual  func(pReqRepeal *def.CThostFtdcReqRepealField, pRspInfo *def.CThostFtdcRspInfoField)                                                                       // 期货发起查询银行余额错误回报
	OnErrRtnQueryBankBalanceByFuture          func(pReqQueryAccount *def.CThostFtdcReqQueryAccountField, pRspInfo *def.CThostFtdcRspInfoField)                                                           // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromBankToFutureByFuture       func(pRspRepeal *def.CThostFtdcRspRepealField)                                                                                                             // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFuture       func(pRspRepeal *def.CThostFtdcRspRepealField)                                                                                                             // 期货发起银行资金转期货应答
	OnRspFromBankToFutureByFuture             func(pReqTransfer *def.CThostFtdcReqTransferField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                     // 期货发起期货资金转银行应答
	OnRspFromFutureToBankByFuture             func(pReqTransfer *def.CThostFtdcReqTransferField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                     // 期货发起查询银行余额应答
	OnRspQueryBankAccountMoneyByFuture        func(pReqQueryAccount *def.CThostFtdcReqQueryAccountField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                             // 银行发起银期开户通知
	OnRtnOpenAccountByBank                    func(pOpenAccount *def.CThostFtdcOpenAccountField)                                                                                                         // 银行发起银期销户通知
	OnRtnCancelAccountByBank                  func(pCancelAccount *def.CThostFtdcCancelAccountField)                                                                                                     // 银行发起变更银行账号通知
	OnRtnChangeAccountByBank                  func(pChangeAccount *def.CThostFtdcChangeAccountField)                                                                                                     // 请求查询分类合约响应
	OnRspQryClassifiedInstrument              func(pInstrument *def.CThostFtdcInstrumentField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                                       // 请求组合优惠比例响应
	OnRspQryCombPromotionParam                func(pCombPromotionParam *def.CThostFtdcCombPromotionParamField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)                       // 投资者风险结算持仓查询响应
	OnRspQryRiskSettleInvstPosition           func(pRiskSettleInvstPosition *def.CThostFtdcRiskSettleInvstPositionField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)             // 风险结算产品查询响应
	OnRspQryRiskSettleProductStatus           func(pRiskSettleProductStatus *def.CThostFtdcRiskSettleProductStatusField, pRspInfo *def.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
}

var t *Trade

func NewTrade() *Trade {
	t = &Trade{}
	path := C.CString("./log/")
	os.MkdirAll("./log/", os.ModePerm)

	t.api = C.CreateFtdcTraderApi(path)
	t.Version = C.GoString((*C.char)(C.GetVersion()))
	fmt.Println(t.Version)

	t.spi = C.CreateFtdcTraderSpi()
	C.tRegisterSpi(t.api, t.spi)

	C.tSetOnFrontConnected(t.spi, C.OnFrontConnected)                                                   // //////////////////////////////////////////////////////////////////////
	C.tSetOnFrontDisconnected(t.spi, C.OnFrontDisconnected)                                             // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	C.tSetOnHeartBeatWarning(t.spi, C.OnHeartBeatWarning)                                               // 心跳超时警告。当长时间未收到报文时，该方法被调用。
	C.tSetOnRspAuthenticate(t.spi, C.OnRspAuthenticate)                                                 // 客户端认证响应
	C.tSetOnRspUserLogin(t.spi, C.OnRspUserLogin)                                                       // 登录请求响应
	C.tSetOnRspUserLogout(t.spi, C.OnRspUserLogout)                                                     // 登出请求响应
	C.tSetOnRspUserPasswordUpdate(t.spi, C.OnRspUserPasswordUpdate)                                     // 用户口令更新请求响应
	C.tSetOnRspTradingAccountPasswordUpdate(t.spi, C.OnRspTradingAccountPasswordUpdate)                 // 资金账户口令更新请求响应
	C.tSetOnRspUserAuthMethod(t.spi, C.OnRspUserAuthMethod)                                             // 查询用户当前支持的认证模式的回复
	C.tSetOnRspGenUserCaptcha(t.spi, C.OnRspGenUserCaptcha)                                             // 获取图形验证码请求的回复
	C.tSetOnRspGenUserText(t.spi, C.OnRspGenUserText)                                                   // 获取短信验证码请求的回复
	C.tSetOnRspOrderInsert(t.spi, C.OnRspOrderInsert)                                                   // 报单录入请求响应
	C.tSetOnRspParkedOrderInsert(t.spi, C.OnRspParkedOrderInsert)                                       // 预埋单录入请求响应
	C.tSetOnRspParkedOrderAction(t.spi, C.OnRspParkedOrderAction)                                       // 预埋撤单录入请求响应
	C.tSetOnRspOrderAction(t.spi, C.OnRspOrderAction)                                                   // 报单操作请求响应
	C.tSetOnRspQryMaxOrderVolume(t.spi, C.OnRspQryMaxOrderVolume)                                       // 查询最大报单数量响应
	C.tSetOnRspSettlementInfoConfirm(t.spi, C.OnRspSettlementInfoConfirm)                               // 投资者结算结果确认响应
	C.tSetOnRspRemoveParkedOrder(t.spi, C.OnRspRemoveParkedOrder)                                       // 删除预埋单响应
	C.tSetOnRspRemoveParkedOrderAction(t.spi, C.OnRspRemoveParkedOrderAction)                           // 删除预埋撤单响应
	C.tSetOnRspExecOrderInsert(t.spi, C.OnRspExecOrderInsert)                                           // 执行宣告录入请求响应
	C.tSetOnRspExecOrderAction(t.spi, C.OnRspExecOrderAction)                                           // 执行宣告操作请求响应
	C.tSetOnRspForQuoteInsert(t.spi, C.OnRspForQuoteInsert)                                             // 询价录入请求响应
	C.tSetOnRspQuoteInsert(t.spi, C.OnRspQuoteInsert)                                                   // 报价录入请求响应
	C.tSetOnRspQuoteAction(t.spi, C.OnRspQuoteAction)                                                   // 报价操作请求响应
	C.tSetOnRspBatchOrderAction(t.spi, C.OnRspBatchOrderAction)                                         // 批量报单操作请求响应
	C.tSetOnRspOptionSelfCloseInsert(t.spi, C.OnRspOptionSelfCloseInsert)                               // 期权自对冲录入请求响应
	C.tSetOnRspOptionSelfCloseAction(t.spi, C.OnRspOptionSelfCloseAction)                               // 期权自对冲操作请求响应
	C.tSetOnRspCombActionInsert(t.spi, C.OnRspCombActionInsert)                                         // 申请组合录入请求响应
	C.tSetOnRspQryOrder(t.spi, C.OnRspQryOrder)                                                         // 请求查询报单响应
	C.tSetOnRspQryTrade(t.spi, C.OnRspQryTrade)                                                         // 请求查询成交响应
	C.tSetOnRspQryInvestorPosition(t.spi, C.OnRspQryInvestorPosition)                                   // 请求查询投资者持仓响应
	C.tSetOnRspQryTradingAccount(t.spi, C.OnRspQryTradingAccount)                                       // 请求查询资金账户响应
	C.tSetOnRspQryInvestor(t.spi, C.OnRspQryInvestor)                                                   // 请求查询投资者响应
	C.tSetOnRspQryTradingCode(t.spi, C.OnRspQryTradingCode)                                             // 请求查询交易编码响应
	C.tSetOnRspQryInstrumentMarginRate(t.spi, C.OnRspQryInstrumentMarginRate)                           // 请求查询合约保证金率响应
	C.tSetOnRspQryInstrumentCommissionRate(t.spi, C.OnRspQryInstrumentCommissionRate)                   // 请求查询合约手续费率响应
	C.tSetOnRspQryExchange(t.spi, C.OnRspQryExchange)                                                   // 请求查询交易所响应
	C.tSetOnRspQryProduct(t.spi, C.OnRspQryProduct)                                                     // 请求查询产品响应
	C.tSetOnRspQryInstrument(t.spi, C.OnRspQryInstrument)                                               // 请求查询合约响应
	C.tSetOnRspQryDepthMarketData(t.spi, C.OnRspQryDepthMarketData)                                     // 请求查询行情响应
	C.tSetOnRspQryTraderOffer(t.spi, C.OnRspQryTraderOffer)                                             // 请求查询交易员报盘机响应
	C.tSetOnRspQrySettlementInfo(t.spi, C.OnRspQrySettlementInfo)                                       // 请求查询投资者结算结果响应
	C.tSetOnRspQryTransferBank(t.spi, C.OnRspQryTransferBank)                                           // 请求查询转帐银行响应
	C.tSetOnRspQryInvestorPositionDetail(t.spi, C.OnRspQryInvestorPositionDetail)                       // 请求查询投资者持仓明细响应
	C.tSetOnRspQryNotice(t.spi, C.OnRspQryNotice)                                                       // 请求查询客户通知响应
	C.tSetOnRspQrySettlementInfoConfirm(t.spi, C.OnRspQrySettlementInfoConfirm)                         // 请求查询结算信息确认响应
	C.tSetOnRspQryInvestorPositionCombineDetail(t.spi, C.OnRspQryInvestorPositionCombineDetail)         // 请求查询投资者持仓明细响应
	C.tSetOnRspQryCFMMCTradingAccountKey(t.spi, C.OnRspQryCFMMCTradingAccountKey)                       // 查询保证金监管系统经纪公司资金账户密钥响应
	C.tSetOnRspQryEWarrantOffset(t.spi, C.OnRspQryEWarrantOffset)                                       // 请求查询仓单折抵信息响应
	C.tSetOnRspQryInvestorProductGroupMargin(t.spi, C.OnRspQryInvestorProductGroupMargin)               // 请求查询投资者品种/跨品种保证金响应
	C.tSetOnRspQryExchangeMarginRate(t.spi, C.OnRspQryExchangeMarginRate)                               // 请求查询交易所保证金率响应
	C.tSetOnRspQryExchangeMarginRateAdjust(t.spi, C.OnRspQryExchangeMarginRateAdjust)                   // 请求查询交易所调整保证金率响应
	C.tSetOnRspQryExchangeRate(t.spi, C.OnRspQryExchangeRate)                                           // 请求查询汇率响应
	C.tSetOnRspQrySecAgentACIDMap(t.spi, C.OnRspQrySecAgentACIDMap)                                     // 请求查询二级代理操作员银期权限响应
	C.tSetOnRspQryProductExchRate(t.spi, C.OnRspQryProductExchRate)                                     // 请求查询产品报价汇率
	C.tSetOnRspQryProductGroup(t.spi, C.OnRspQryProductGroup)                                           // 请求查询产品组
	C.tSetOnRspQryMMInstrumentCommissionRate(t.spi, C.OnRspQryMMInstrumentCommissionRate)               // 请求查询做市商合约手续费率响应
	C.tSetOnRspQryMMOptionInstrCommRate(t.spi, C.OnRspQryMMOptionInstrCommRate)                         // 请求查询做市商期权合约手续费响应
	C.tSetOnRspQryInstrumentOrderCommRate(t.spi, C.OnRspQryInstrumentOrderCommRate)                     // 请求查询报单手续费响应
	C.tSetOnRspQrySecAgentTradingAccount(t.spi, C.OnRspQrySecAgentTradingAccount)                       // 请求查询资金账户响应
	C.tSetOnRspQrySecAgentCheckMode(t.spi, C.OnRspQrySecAgentCheckMode)                                 // 请求查询二级代理商资金校验模式响应
	C.tSetOnRspQrySecAgentTradeInfo(t.spi, C.OnRspQrySecAgentTradeInfo)                                 // 请求查询二级代理商信息响应
	C.tSetOnRspQryOptionInstrTradeCost(t.spi, C.OnRspQryOptionInstrTradeCost)                           // 请求查询期权交易成本响应
	C.tSetOnRspQryOptionInstrCommRate(t.spi, C.OnRspQryOptionInstrCommRate)                             // 请求查询期权合约手续费响应
	C.tSetOnRspQryExecOrder(t.spi, C.OnRspQryExecOrder)                                                 // 请求查询执行宣告响应
	C.tSetOnRspQryForQuote(t.spi, C.OnRspQryForQuote)                                                   // 请求查询询价响应
	C.tSetOnRspQryQuote(t.spi, C.OnRspQryQuote)                                                         // 请求查询报价响应
	C.tSetOnRspQryOptionSelfClose(t.spi, C.OnRspQryOptionSelfClose)                                     // 请求查询期权自对冲响应
	C.tSetOnRspQryInvestUnit(t.spi, C.OnRspQryInvestUnit)                                               // 请求查询投资单元响应
	C.tSetOnRspQryCombInstrumentGuard(t.spi, C.OnRspQryCombInstrumentGuard)                             // 请求查询组合合约安全系数响应
	C.tSetOnRspQryCombAction(t.spi, C.OnRspQryCombAction)                                               // 请求查询申请组合响应
	C.tSetOnRspQryTransferSerial(t.spi, C.OnRspQryTransferSerial)                                       // 请求查询转帐流水响应
	C.tSetOnRspQryAccountregister(t.spi, C.OnRspQryAccountregister)                                     // 请求查询银期签约关系响应
	C.tSetOnRspError(t.spi, C.OnRspError)                                                               // 错误应答
	C.tSetOnRtnOrder(t.spi, C.OnRtnOrder)                                                               // 报单通知
	C.tSetOnRtnTrade(t.spi, C.OnRtnTrade)                                                               // 成交通知
	C.tSetOnErrRtnOrderInsert(t.spi, C.OnErrRtnOrderInsert)                                             // 报单录入错误回报
	C.tSetOnErrRtnOrderAction(t.spi, C.OnErrRtnOrderAction)                                             // 报单操作错误回报
	C.tSetOnRtnInstrumentStatus(t.spi, C.OnRtnInstrumentStatus)                                         // 合约交易状态通知
	C.tSetOnRtnBulletin(t.spi, C.OnRtnBulletin)                                                         // 交易所公告通知
	C.tSetOnRtnTradingNotice(t.spi, C.OnRtnTradingNotice)                                               // 交易通知
	C.tSetOnRtnErrorConditionalOrder(t.spi, C.OnRtnErrorConditionalOrder)                               // 提示条件单校验错误
	C.tSetOnRtnExecOrder(t.spi, C.OnRtnExecOrder)                                                       // 执行宣告通知
	C.tSetOnErrRtnExecOrderInsert(t.spi, C.OnErrRtnExecOrderInsert)                                     // 执行宣告录入错误回报
	C.tSetOnErrRtnExecOrderAction(t.spi, C.OnErrRtnExecOrderAction)                                     // 执行宣告操作错误回报
	C.tSetOnErrRtnForQuoteInsert(t.spi, C.OnErrRtnForQuoteInsert)                                       // 询价录入错误回报
	C.tSetOnRtnQuote(t.spi, C.OnRtnQuote)                                                               // 报价通知
	C.tSetOnErrRtnQuoteInsert(t.spi, C.OnErrRtnQuoteInsert)                                             // 报价录入错误回报
	C.tSetOnErrRtnQuoteAction(t.spi, C.OnErrRtnQuoteAction)                                             // 报价操作错误回报
	C.tSetOnRtnForQuoteRsp(t.spi, C.OnRtnForQuoteRsp)                                                   // 询价通知
	C.tSetOnRtnCFMMCTradingAccountToken(t.spi, C.OnRtnCFMMCTradingAccountToken)                         // 保证金监控中心用户令牌
	C.tSetOnErrRtnBatchOrderAction(t.spi, C.OnErrRtnBatchOrderAction)                                   // 批量报单操作错误回报
	C.tSetOnRtnOptionSelfClose(t.spi, C.OnRtnOptionSelfClose)                                           // 期权自对冲通知
	C.tSetOnErrRtnOptionSelfCloseInsert(t.spi, C.OnErrRtnOptionSelfCloseInsert)                         // 期权自对冲录入错误回报
	C.tSetOnErrRtnOptionSelfCloseAction(t.spi, C.OnErrRtnOptionSelfCloseAction)                         // 期权自对冲操作错误回报
	C.tSetOnRtnCombAction(t.spi, C.OnRtnCombAction)                                                     // 申请组合通知
	C.tSetOnErrRtnCombActionInsert(t.spi, C.OnErrRtnCombActionInsert)                                   // 申请组合录入错误回报
	C.tSetOnRspQryContractBank(t.spi, C.OnRspQryContractBank)                                           // 请求查询签约银行响应
	C.tSetOnRspQryParkedOrder(t.spi, C.OnRspQryParkedOrder)                                             // 请求查询预埋单响应
	C.tSetOnRspQryParkedOrderAction(t.spi, C.OnRspQryParkedOrderAction)                                 // 请求查询预埋撤单响应
	C.tSetOnRspQryTradingNotice(t.spi, C.OnRspQryTradingNotice)                                         // 请求查询交易通知响应
	C.tSetOnRspQryBrokerTradingParams(t.spi, C.OnRspQryBrokerTradingParams)                             // 请求查询经纪公司交易参数响应
	C.tSetOnRspQryBrokerTradingAlgos(t.spi, C.OnRspQryBrokerTradingAlgos)                               // 请求查询经纪公司交易算法响应
	C.tSetOnRspQueryCFMMCTradingAccountToken(t.spi, C.OnRspQueryCFMMCTradingAccountToken)               // 请求查询监控中心用户令牌
	C.tSetOnRtnFromBankToFutureByBank(t.spi, C.OnRtnFromBankToFutureByBank)                             // 银行发起银行资金转期货通知
	C.tSetOnRtnFromFutureToBankByBank(t.spi, C.OnRtnFromFutureToBankByBank)                             // 银行发起期货资金转银行通知
	C.tSetOnRtnRepealFromBankToFutureByBank(t.spi, C.OnRtnRepealFromBankToFutureByBank)                 // 银行发起冲正银行转期货通知
	C.tSetOnRtnRepealFromFutureToBankByBank(t.spi, C.OnRtnRepealFromFutureToBankByBank)                 // 银行发起冲正期货转银行通知
	C.tSetOnRtnFromBankToFutureByFuture(t.spi, C.OnRtnFromBankToFutureByFuture)                         // 期货发起银行资金转期货通知
	C.tSetOnRtnFromFutureToBankByFuture(t.spi, C.OnRtnFromFutureToBankByFuture)                         // 期货发起期货资金转银行通知
	C.tSetOnRtnRepealFromBankToFutureByFutureManual(t.spi, C.OnRtnRepealFromBankToFutureByFutureManual) // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	C.tSetOnRtnRepealFromFutureToBankByFutureManual(t.spi, C.OnRtnRepealFromFutureToBankByFutureManual) // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	C.tSetOnRtnQueryBankBalanceByFuture(t.spi, C.OnRtnQueryBankBalanceByFuture)                         // 期货发起查询银行余额通知
	C.tSetOnErrRtnBankToFutureByFuture(t.spi, C.OnErrRtnBankToFutureByFuture)                           // 期货发起银行资金转期货错误回报
	C.tSetOnErrRtnFutureToBankByFuture(t.spi, C.OnErrRtnFutureToBankByFuture)                           // 期货发起期货资金转银行错误回报
	C.tSetOnErrRtnRepealBankToFutureByFutureManual(t.spi, C.OnErrRtnRepealBankToFutureByFutureManual)   // 系统运行时期货端手工发起冲正银行转期货错误回报
	C.tSetOnErrRtnRepealFutureToBankByFutureManual(t.spi, C.OnErrRtnRepealFutureToBankByFutureManual)   // 系统运行时期货端手工发起冲正期货转银行错误回报
	C.tSetOnErrRtnQueryBankBalanceByFuture(t.spi, C.OnErrRtnQueryBankBalanceByFuture)                   // 期货发起查询银行余额错误回报
	C.tSetOnRtnRepealFromBankToFutureByFuture(t.spi, C.OnRtnRepealFromBankToFutureByFuture)             // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	C.tSetOnRtnRepealFromFutureToBankByFuture(t.spi, C.OnRtnRepealFromFutureToBankByFuture)             // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	C.tSetOnRspFromBankToFutureByFuture(t.spi, C.OnRspFromBankToFutureByFuture)                         // 期货发起银行资金转期货应答
	C.tSetOnRspFromFutureToBankByFuture(t.spi, C.OnRspFromFutureToBankByFuture)                         // 期货发起期货资金转银行应答
	C.tSetOnRspQueryBankAccountMoneyByFuture(t.spi, C.OnRspQueryBankAccountMoneyByFuture)               // 期货发起查询银行余额应答
	C.tSetOnRtnOpenAccountByBank(t.spi, C.OnRtnOpenAccountByBank)                                       // 银行发起银期开户通知
	C.tSetOnRtnCancelAccountByBank(t.spi, C.OnRtnCancelAccountByBank)                                   // 银行发起银期销户通知
	C.tSetOnRtnChangeAccountByBank(t.spi, C.OnRtnChangeAccountByBank)                                   // 银行发起变更银行账号通知
	C.tSetOnRspQryClassifiedInstrument(t.spi, C.OnRspQryClassifiedInstrument)                           // 请求查询分类合约响应
	C.tSetOnRspQryCombPromotionParam(t.spi, C.OnRspQryCombPromotionParam)                               // 请求组合优惠比例响应
	C.tSetOnRspQryRiskSettleInvstPosition(t.spi, C.OnRspQryRiskSettleInvstPosition)                     // 投资者风险结算持仓查询响应
	C.tSetOnRspQryRiskSettleProductStatus(t.spi, C.OnRspQryRiskSettleProductStatus)                     // 风险结算产品查询响应

	return t
}

//export OnFrontConnected
func OnFrontConnected() {
	if t.OnFrontConnected == nil {
		fmt.Println("OnFrontConnected")
	} else {
		t.OnFrontConnected()
	}
}

//export OnFrontDisconnected
func OnFrontDisconnected(nReason C.int) {
	if t.OnFrontDisconnected == nil {
		fmt.Println("OnFrontDisconnected")
	} else {
		t.OnFrontDisconnected(int(nReason))
	}
}

//export OnHeartBeatWarning
func OnHeartBeatWarning(nTimeLapse C.int) {
	if t.OnHeartBeatWarning == nil {
		fmt.Println("OnHeartBeatWarning")
	} else {
		t.OnHeartBeatWarning(int(nTimeLapse))
	}
}

//export OnRspAuthenticate
func OnRspAuthenticate(pRspAuthenticateField *C.struct_CThostFtdcRspAuthenticateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspAuthenticate == nil {
		fmt.Println("OnRspAuthenticate")
	} else {
		t.OnRspAuthenticate((*def.CThostFtdcRspAuthenticateField)(unsafe.Pointer(pRspAuthenticateField)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspUserLogin
func OnRspUserLogin(pRspUserLogin *C.struct_CThostFtdcRspUserLoginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspUserLogin == nil {
		fmt.Println("OnRspUserLogin")
	} else {
		t.OnRspUserLogin((*def.CThostFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspUserLogout
func OnRspUserLogout(pUserLogout *C.struct_CThostFtdcUserLogoutField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspUserLogout == nil {
		fmt.Println("OnRspUserLogout")
	} else {
		t.OnRspUserLogout((*def.CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspUserPasswordUpdate
func OnRspUserPasswordUpdate(pUserPasswordUpdate *C.struct_CThostFtdcUserPasswordUpdateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspUserPasswordUpdate == nil {
		fmt.Println("OnRspUserPasswordUpdate")
	} else {
		t.OnRspUserPasswordUpdate((*def.CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspTradingAccountPasswordUpdate
func OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *C.struct_CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspTradingAccountPasswordUpdate == nil {
		fmt.Println("OnRspTradingAccountPasswordUpdate")
	} else {
		t.OnRspTradingAccountPasswordUpdate((*def.CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspUserAuthMethod
func OnRspUserAuthMethod(pRspUserAuthMethod *C.struct_CThostFtdcRspUserAuthMethodField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspUserAuthMethod == nil {
		fmt.Println("OnRspUserAuthMethod")
	} else {
		t.OnRspUserAuthMethod((*def.CThostFtdcRspUserAuthMethodField)(unsafe.Pointer(pRspUserAuthMethod)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspGenUserCaptcha
func OnRspGenUserCaptcha(pRspGenUserCaptcha *C.struct_CThostFtdcRspGenUserCaptchaField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspGenUserCaptcha == nil {
		fmt.Println("OnRspGenUserCaptcha")
	} else {
		t.OnRspGenUserCaptcha((*def.CThostFtdcRspGenUserCaptchaField)(unsafe.Pointer(pRspGenUserCaptcha)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspGenUserText
func OnRspGenUserText(pRspGenUserText *C.struct_CThostFtdcRspGenUserTextField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspGenUserText == nil {
		fmt.Println("OnRspGenUserText")
	} else {
		t.OnRspGenUserText((*def.CThostFtdcRspGenUserTextField)(unsafe.Pointer(pRspGenUserText)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspOrderInsert
func OnRspOrderInsert(pInputOrder *C.struct_CThostFtdcInputOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspOrderInsert == nil {
		fmt.Println("OnRspOrderInsert")
	} else {
		t.OnRspOrderInsert((*def.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspParkedOrderInsert
func OnRspParkedOrderInsert(pParkedOrder *C.struct_CThostFtdcParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspParkedOrderInsert == nil {
		fmt.Println("OnRspParkedOrderInsert")
	} else {
		t.OnRspParkedOrderInsert((*def.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspParkedOrderAction
func OnRspParkedOrderAction(pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspParkedOrderAction == nil {
		fmt.Println("OnRspParkedOrderAction")
	} else {
		t.OnRspParkedOrderAction((*def.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspOrderAction
func OnRspOrderAction(pInputOrderAction *C.struct_CThostFtdcInputOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspOrderAction == nil {
		fmt.Println("OnRspOrderAction")
	} else {
		t.OnRspOrderAction((*def.CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryMaxOrderVolume
func OnRspQryMaxOrderVolume(pQryMaxOrderVolume *C.struct_CThostFtdcQryMaxOrderVolumeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryMaxOrderVolume == nil {
		fmt.Println("OnRspQryMaxOrderVolume")
	} else {
		t.OnRspQryMaxOrderVolume((*def.CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspSettlementInfoConfirm
func OnRspSettlementInfoConfirm(pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspSettlementInfoConfirm == nil {
		fmt.Println("OnRspSettlementInfoConfirm")
	} else {
		t.OnRspSettlementInfoConfirm((*def.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspRemoveParkedOrder
func OnRspRemoveParkedOrder(pRemoveParkedOrder *C.struct_CThostFtdcRemoveParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspRemoveParkedOrder == nil {
		fmt.Println("OnRspRemoveParkedOrder")
	} else {
		t.OnRspRemoveParkedOrder((*def.CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspRemoveParkedOrderAction
func OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *C.struct_CThostFtdcRemoveParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspRemoveParkedOrderAction == nil {
		fmt.Println("OnRspRemoveParkedOrderAction")
	} else {
		t.OnRspRemoveParkedOrderAction((*def.CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspExecOrderInsert
func OnRspExecOrderInsert(pInputExecOrder *C.struct_CThostFtdcInputExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspExecOrderInsert == nil {
		fmt.Println("OnRspExecOrderInsert")
	} else {
		t.OnRspExecOrderInsert((*def.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspExecOrderAction
func OnRspExecOrderAction(pInputExecOrderAction *C.struct_CThostFtdcInputExecOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspExecOrderAction == nil {
		fmt.Println("OnRspExecOrderAction")
	} else {
		t.OnRspExecOrderAction((*def.CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspForQuoteInsert
func OnRspForQuoteInsert(pInputForQuote *C.struct_CThostFtdcInputForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspForQuoteInsert == nil {
		fmt.Println("OnRspForQuoteInsert")
	} else {
		t.OnRspForQuoteInsert((*def.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQuoteInsert
func OnRspQuoteInsert(pInputQuote *C.struct_CThostFtdcInputQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQuoteInsert == nil {
		fmt.Println("OnRspQuoteInsert")
	} else {
		t.OnRspQuoteInsert((*def.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQuoteAction
func OnRspQuoteAction(pInputQuoteAction *C.struct_CThostFtdcInputQuoteActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQuoteAction == nil {
		fmt.Println("OnRspQuoteAction")
	} else {
		t.OnRspQuoteAction((*def.CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspBatchOrderAction
func OnRspBatchOrderAction(pInputBatchOrderAction *C.struct_CThostFtdcInputBatchOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspBatchOrderAction == nil {
		fmt.Println("OnRspBatchOrderAction")
	} else {
		t.OnRspBatchOrderAction((*def.CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspOptionSelfCloseInsert
func OnRspOptionSelfCloseInsert(pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspOptionSelfCloseInsert == nil {
		fmt.Println("OnRspOptionSelfCloseInsert")
	} else {
		t.OnRspOptionSelfCloseInsert((*def.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspOptionSelfCloseAction
func OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *C.struct_CThostFtdcInputOptionSelfCloseActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspOptionSelfCloseAction == nil {
		fmt.Println("OnRspOptionSelfCloseAction")
	} else {
		t.OnRspOptionSelfCloseAction((*def.CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspCombActionInsert
func OnRspCombActionInsert(pInputCombAction *C.struct_CThostFtdcInputCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspCombActionInsert == nil {
		fmt.Println("OnRspCombActionInsert")
	} else {
		t.OnRspCombActionInsert((*def.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryOrder
func OnRspQryOrder(pOrder *C.struct_CThostFtdcOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryOrder == nil {
		fmt.Println("OnRspQryOrder")
	} else {
		t.OnRspQryOrder((*def.CThostFtdcOrderField)(unsafe.Pointer(pOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTrade
func OnRspQryTrade(pTrade *C.struct_CThostFtdcTradeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTrade == nil {
		fmt.Println("OnRspQryTrade")
	} else {
		t.OnRspQryTrade((*def.CThostFtdcTradeField)(unsafe.Pointer(pTrade)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInvestorPosition
func OnRspQryInvestorPosition(pInvestorPosition *C.struct_CThostFtdcInvestorPositionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInvestorPosition == nil {
		fmt.Println("OnRspQryInvestorPosition")
	} else {
		t.OnRspQryInvestorPosition((*def.CThostFtdcInvestorPositionField)(unsafe.Pointer(pInvestorPosition)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTradingAccount
func OnRspQryTradingAccount(pTradingAccount *C.struct_CThostFtdcTradingAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTradingAccount == nil {
		fmt.Println("OnRspQryTradingAccount")
	} else {
		t.OnRspQryTradingAccount((*def.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInvestor
func OnRspQryInvestor(pInvestor *C.struct_CThostFtdcInvestorField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInvestor == nil {
		fmt.Println("OnRspQryInvestor")
	} else {
		t.OnRspQryInvestor((*def.CThostFtdcInvestorField)(unsafe.Pointer(pInvestor)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTradingCode
func OnRspQryTradingCode(pTradingCode *C.struct_CThostFtdcTradingCodeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTradingCode == nil {
		fmt.Println("OnRspQryTradingCode")
	} else {
		t.OnRspQryTradingCode((*def.CThostFtdcTradingCodeField)(unsafe.Pointer(pTradingCode)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInstrumentMarginRate
func OnRspQryInstrumentMarginRate(pInstrumentMarginRate *C.struct_CThostFtdcInstrumentMarginRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInstrumentMarginRate == nil {
		fmt.Println("OnRspQryInstrumentMarginRate")
	} else {
		t.OnRspQryInstrumentMarginRate((*def.CThostFtdcInstrumentMarginRateField)(unsafe.Pointer(pInstrumentMarginRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInstrumentCommissionRate
func OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *C.struct_CThostFtdcInstrumentCommissionRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInstrumentCommissionRate == nil {
		fmt.Println("OnRspQryInstrumentCommissionRate")
	} else {
		t.OnRspQryInstrumentCommissionRate((*def.CThostFtdcInstrumentCommissionRateField)(unsafe.Pointer(pInstrumentCommissionRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryExchange
func OnRspQryExchange(pExchange *C.struct_CThostFtdcExchangeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryExchange == nil {
		fmt.Println("OnRspQryExchange")
	} else {
		t.OnRspQryExchange((*def.CThostFtdcExchangeField)(unsafe.Pointer(pExchange)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryProduct
func OnRspQryProduct(pProduct *C.struct_CThostFtdcProductField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryProduct == nil {
		fmt.Println("OnRspQryProduct")
	} else {
		t.OnRspQryProduct((*def.CThostFtdcProductField)(unsafe.Pointer(pProduct)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInstrument
func OnRspQryInstrument(pInstrument *C.struct_CThostFtdcInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInstrument == nil {
		fmt.Println("OnRspQryInstrument")
	} else {
		t.OnRspQryInstrument((*def.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryDepthMarketData
func OnRspQryDepthMarketData(pDepthMarketData *C.struct_CThostFtdcDepthMarketDataField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryDepthMarketData == nil {
		fmt.Println("OnRspQryDepthMarketData")
	} else {
		t.OnRspQryDepthMarketData((*def.CThostFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTraderOffer
func OnRspQryTraderOffer(pTraderOffer *C.struct_CThostFtdcTraderOfferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTraderOffer == nil {
		fmt.Println("OnRspQryTraderOffer")
	} else {
		t.OnRspQryTraderOffer((*def.CThostFtdcTraderOfferField)(unsafe.Pointer(pTraderOffer)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQrySettlementInfo
func OnRspQrySettlementInfo(pSettlementInfo *C.struct_CThostFtdcSettlementInfoField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQrySettlementInfo == nil {
		fmt.Println("OnRspQrySettlementInfo")
	} else {
		t.OnRspQrySettlementInfo((*def.CThostFtdcSettlementInfoField)(unsafe.Pointer(pSettlementInfo)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTransferBank
func OnRspQryTransferBank(pTransferBank *C.struct_CThostFtdcTransferBankField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTransferBank == nil {
		fmt.Println("OnRspQryTransferBank")
	} else {
		t.OnRspQryTransferBank((*def.CThostFtdcTransferBankField)(unsafe.Pointer(pTransferBank)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInvestorPositionDetail
func OnRspQryInvestorPositionDetail(pInvestorPositionDetail *C.struct_CThostFtdcInvestorPositionDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInvestorPositionDetail == nil {
		fmt.Println("OnRspQryInvestorPositionDetail")
	} else {
		t.OnRspQryInvestorPositionDetail((*def.CThostFtdcInvestorPositionDetailField)(unsafe.Pointer(pInvestorPositionDetail)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryNotice
func OnRspQryNotice(pNotice *C.struct_CThostFtdcNoticeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryNotice == nil {
		fmt.Println("OnRspQryNotice")
	} else {
		t.OnRspQryNotice((*def.CThostFtdcNoticeField)(unsafe.Pointer(pNotice)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQrySettlementInfoConfirm
func OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *C.struct_CThostFtdcSettlementInfoConfirmField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQrySettlementInfoConfirm == nil {
		fmt.Println("OnRspQrySettlementInfoConfirm")
	} else {
		t.OnRspQrySettlementInfoConfirm((*def.CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInvestorPositionCombineDetail
func OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *C.struct_CThostFtdcInvestorPositionCombineDetailField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInvestorPositionCombineDetail == nil {
		fmt.Println("OnRspQryInvestorPositionCombineDetail")
	} else {
		t.OnRspQryInvestorPositionCombineDetail((*def.CThostFtdcInvestorPositionCombineDetailField)(unsafe.Pointer(pInvestorPositionCombineDetail)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryCFMMCTradingAccountKey
func OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *C.struct_CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryCFMMCTradingAccountKey == nil {
		fmt.Println("OnRspQryCFMMCTradingAccountKey")
	} else {
		t.OnRspQryCFMMCTradingAccountKey((*def.CThostFtdcCFMMCTradingAccountKeyField)(unsafe.Pointer(pCFMMCTradingAccountKey)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryEWarrantOffset
func OnRspQryEWarrantOffset(pEWarrantOffset *C.struct_CThostFtdcEWarrantOffsetField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryEWarrantOffset == nil {
		fmt.Println("OnRspQryEWarrantOffset")
	} else {
		t.OnRspQryEWarrantOffset((*def.CThostFtdcEWarrantOffsetField)(unsafe.Pointer(pEWarrantOffset)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInvestorProductGroupMargin
func OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *C.struct_CThostFtdcInvestorProductGroupMarginField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInvestorProductGroupMargin == nil {
		fmt.Println("OnRspQryInvestorProductGroupMargin")
	} else {
		t.OnRspQryInvestorProductGroupMargin((*def.CThostFtdcInvestorProductGroupMarginField)(unsafe.Pointer(pInvestorProductGroupMargin)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryExchangeMarginRate
func OnRspQryExchangeMarginRate(pExchangeMarginRate *C.struct_CThostFtdcExchangeMarginRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryExchangeMarginRate == nil {
		fmt.Println("OnRspQryExchangeMarginRate")
	} else {
		t.OnRspQryExchangeMarginRate((*def.CThostFtdcExchangeMarginRateField)(unsafe.Pointer(pExchangeMarginRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryExchangeMarginRateAdjust
func OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *C.struct_CThostFtdcExchangeMarginRateAdjustField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryExchangeMarginRateAdjust == nil {
		fmt.Println("OnRspQryExchangeMarginRateAdjust")
	} else {
		t.OnRspQryExchangeMarginRateAdjust((*def.CThostFtdcExchangeMarginRateAdjustField)(unsafe.Pointer(pExchangeMarginRateAdjust)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryExchangeRate
func OnRspQryExchangeRate(pExchangeRate *C.struct_CThostFtdcExchangeRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryExchangeRate == nil {
		fmt.Println("OnRspQryExchangeRate")
	} else {
		t.OnRspQryExchangeRate((*def.CThostFtdcExchangeRateField)(unsafe.Pointer(pExchangeRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQrySecAgentACIDMap
func OnRspQrySecAgentACIDMap(pSecAgentACIDMap *C.struct_CThostFtdcSecAgentACIDMapField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQrySecAgentACIDMap == nil {
		fmt.Println("OnRspQrySecAgentACIDMap")
	} else {
		t.OnRspQrySecAgentACIDMap((*def.CThostFtdcSecAgentACIDMapField)(unsafe.Pointer(pSecAgentACIDMap)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryProductExchRate
func OnRspQryProductExchRate(pProductExchRate *C.struct_CThostFtdcProductExchRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryProductExchRate == nil {
		fmt.Println("OnRspQryProductExchRate")
	} else {
		t.OnRspQryProductExchRate((*def.CThostFtdcProductExchRateField)(unsafe.Pointer(pProductExchRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryProductGroup
func OnRspQryProductGroup(pProductGroup *C.struct_CThostFtdcProductGroupField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryProductGroup == nil {
		fmt.Println("OnRspQryProductGroup")
	} else {
		t.OnRspQryProductGroup((*def.CThostFtdcProductGroupField)(unsafe.Pointer(pProductGroup)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryMMInstrumentCommissionRate
func OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *C.struct_CThostFtdcMMInstrumentCommissionRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryMMInstrumentCommissionRate == nil {
		fmt.Println("OnRspQryMMInstrumentCommissionRate")
	} else {
		t.OnRspQryMMInstrumentCommissionRate((*def.CThostFtdcMMInstrumentCommissionRateField)(unsafe.Pointer(pMMInstrumentCommissionRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryMMOptionInstrCommRate
func OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *C.struct_CThostFtdcMMOptionInstrCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryMMOptionInstrCommRate == nil {
		fmt.Println("OnRspQryMMOptionInstrCommRate")
	} else {
		t.OnRspQryMMOptionInstrCommRate((*def.CThostFtdcMMOptionInstrCommRateField)(unsafe.Pointer(pMMOptionInstrCommRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInstrumentOrderCommRate
func OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *C.struct_CThostFtdcInstrumentOrderCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInstrumentOrderCommRate == nil {
		fmt.Println("OnRspQryInstrumentOrderCommRate")
	} else {
		t.OnRspQryInstrumentOrderCommRate((*def.CThostFtdcInstrumentOrderCommRateField)(unsafe.Pointer(pInstrumentOrderCommRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQrySecAgentTradingAccount
func OnRspQrySecAgentTradingAccount(pTradingAccount *C.struct_CThostFtdcTradingAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQrySecAgentTradingAccount == nil {
		fmt.Println("OnRspQrySecAgentTradingAccount")
	} else {
		t.OnRspQrySecAgentTradingAccount((*def.CThostFtdcTradingAccountField)(unsafe.Pointer(pTradingAccount)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQrySecAgentCheckMode
func OnRspQrySecAgentCheckMode(pSecAgentCheckMode *C.struct_CThostFtdcSecAgentCheckModeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQrySecAgentCheckMode == nil {
		fmt.Println("OnRspQrySecAgentCheckMode")
	} else {
		t.OnRspQrySecAgentCheckMode((*def.CThostFtdcSecAgentCheckModeField)(unsafe.Pointer(pSecAgentCheckMode)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQrySecAgentTradeInfo
func OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *C.struct_CThostFtdcSecAgentTradeInfoField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQrySecAgentTradeInfo == nil {
		fmt.Println("OnRspQrySecAgentTradeInfo")
	} else {
		t.OnRspQrySecAgentTradeInfo((*def.CThostFtdcSecAgentTradeInfoField)(unsafe.Pointer(pSecAgentTradeInfo)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryOptionInstrTradeCost
func OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *C.struct_CThostFtdcOptionInstrTradeCostField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryOptionInstrTradeCost == nil {
		fmt.Println("OnRspQryOptionInstrTradeCost")
	} else {
		t.OnRspQryOptionInstrTradeCost((*def.CThostFtdcOptionInstrTradeCostField)(unsafe.Pointer(pOptionInstrTradeCost)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryOptionInstrCommRate
func OnRspQryOptionInstrCommRate(pOptionInstrCommRate *C.struct_CThostFtdcOptionInstrCommRateField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryOptionInstrCommRate == nil {
		fmt.Println("OnRspQryOptionInstrCommRate")
	} else {
		t.OnRspQryOptionInstrCommRate((*def.CThostFtdcOptionInstrCommRateField)(unsafe.Pointer(pOptionInstrCommRate)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryExecOrder
func OnRspQryExecOrder(pExecOrder *C.struct_CThostFtdcExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryExecOrder == nil {
		fmt.Println("OnRspQryExecOrder")
	} else {
		t.OnRspQryExecOrder((*def.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryForQuote
func OnRspQryForQuote(pForQuote *C.struct_CThostFtdcForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryForQuote == nil {
		fmt.Println("OnRspQryForQuote")
	} else {
		t.OnRspQryForQuote((*def.CThostFtdcForQuoteField)(unsafe.Pointer(pForQuote)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryQuote
func OnRspQryQuote(pQuote *C.struct_CThostFtdcQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryQuote == nil {
		fmt.Println("OnRspQryQuote")
	} else {
		t.OnRspQryQuote((*def.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryOptionSelfClose
func OnRspQryOptionSelfClose(pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryOptionSelfClose == nil {
		fmt.Println("OnRspQryOptionSelfClose")
	} else {
		t.OnRspQryOptionSelfClose((*def.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryInvestUnit
func OnRspQryInvestUnit(pInvestUnit *C.struct_CThostFtdcInvestUnitField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryInvestUnit == nil {
		fmt.Println("OnRspQryInvestUnit")
	} else {
		t.OnRspQryInvestUnit((*def.CThostFtdcInvestUnitField)(unsafe.Pointer(pInvestUnit)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryCombInstrumentGuard
func OnRspQryCombInstrumentGuard(pCombInstrumentGuard *C.struct_CThostFtdcCombInstrumentGuardField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryCombInstrumentGuard == nil {
		fmt.Println("OnRspQryCombInstrumentGuard")
	} else {
		t.OnRspQryCombInstrumentGuard((*def.CThostFtdcCombInstrumentGuardField)(unsafe.Pointer(pCombInstrumentGuard)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryCombAction
func OnRspQryCombAction(pCombAction *C.struct_CThostFtdcCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryCombAction == nil {
		fmt.Println("OnRspQryCombAction")
	} else {
		t.OnRspQryCombAction((*def.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTransferSerial
func OnRspQryTransferSerial(pTransferSerial *C.struct_CThostFtdcTransferSerialField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTransferSerial == nil {
		fmt.Println("OnRspQryTransferSerial")
	} else {
		t.OnRspQryTransferSerial((*def.CThostFtdcTransferSerialField)(unsafe.Pointer(pTransferSerial)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryAccountregister
func OnRspQryAccountregister(pAccountregister *C.struct_CThostFtdcAccountregisterField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryAccountregister == nil {
		fmt.Println("OnRspQryAccountregister")
	} else {
		t.OnRspQryAccountregister((*def.CThostFtdcAccountregisterField)(unsafe.Pointer(pAccountregister)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspError
func OnRspError(pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspError == nil {
		fmt.Println("OnRspError")
	} else {
		t.OnRspError((*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRtnOrder
func OnRtnOrder(pOrder *C.struct_CThostFtdcOrderField) {
	if t.OnRtnOrder == nil {
		fmt.Println("OnRtnOrder")
	} else {
		t.OnRtnOrder((*def.CThostFtdcOrderField)(unsafe.Pointer(pOrder)))
	}
}

//export OnRtnTrade
func OnRtnTrade(pTrade *C.struct_CThostFtdcTradeField) {
	if t.OnRtnTrade == nil {
		fmt.Println("OnRtnTrade")
	} else {
		t.OnRtnTrade((*def.CThostFtdcTradeField)(unsafe.Pointer(pTrade)))
	}
}

//export OnErrRtnOrderInsert
func OnErrRtnOrderInsert(pInputOrder *C.struct_CThostFtdcInputOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnOrderInsert == nil {
		fmt.Println("OnErrRtnOrderInsert")
	} else {
		t.OnErrRtnOrderInsert((*def.CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnOrderAction
func OnErrRtnOrderAction(pOrderAction *C.struct_CThostFtdcOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnOrderAction == nil {
		fmt.Println("OnErrRtnOrderAction")
	} else {
		t.OnErrRtnOrderAction((*def.CThostFtdcOrderActionField)(unsafe.Pointer(pOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRtnInstrumentStatus
func OnRtnInstrumentStatus(pInstrumentStatus *C.struct_CThostFtdcInstrumentStatusField) {
	if t.OnRtnInstrumentStatus == nil {
		fmt.Println("OnRtnInstrumentStatus")
	} else {
		t.OnRtnInstrumentStatus((*def.CThostFtdcInstrumentStatusField)(unsafe.Pointer(pInstrumentStatus)))
	}
}

//export OnRtnBulletin
func OnRtnBulletin(pBulletin *C.struct_CThostFtdcBulletinField) {
	if t.OnRtnBulletin == nil {
		fmt.Println("OnRtnBulletin")
	} else {
		t.OnRtnBulletin((*def.CThostFtdcBulletinField)(unsafe.Pointer(pBulletin)))
	}
}

//export OnRtnTradingNotice
func OnRtnTradingNotice(pTradingNoticeInfo *C.struct_CThostFtdcTradingNoticeInfoField) {
	if t.OnRtnTradingNotice == nil {
		fmt.Println("OnRtnTradingNotice")
	} else {
		t.OnRtnTradingNotice((*def.CThostFtdcTradingNoticeInfoField)(unsafe.Pointer(pTradingNoticeInfo)))
	}
}

//export OnRtnErrorConditionalOrder
func OnRtnErrorConditionalOrder(pErrorConditionalOrder *C.struct_CThostFtdcErrorConditionalOrderField) {
	if t.OnRtnErrorConditionalOrder == nil {
		fmt.Println("OnRtnErrorConditionalOrder")
	} else {
		t.OnRtnErrorConditionalOrder((*def.CThostFtdcErrorConditionalOrderField)(unsafe.Pointer(pErrorConditionalOrder)))
	}
}

//export OnRtnExecOrder
func OnRtnExecOrder(pExecOrder *C.struct_CThostFtdcExecOrderField) {
	if t.OnRtnExecOrder == nil {
		fmt.Println("OnRtnExecOrder")
	} else {
		t.OnRtnExecOrder((*def.CThostFtdcExecOrderField)(unsafe.Pointer(pExecOrder)))
	}
}

//export OnErrRtnExecOrderInsert
func OnErrRtnExecOrderInsert(pInputExecOrder *C.struct_CThostFtdcInputExecOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnExecOrderInsert == nil {
		fmt.Println("OnErrRtnExecOrderInsert")
	} else {
		t.OnErrRtnExecOrderInsert((*def.CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnExecOrderAction
func OnErrRtnExecOrderAction(pExecOrderAction *C.struct_CThostFtdcExecOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnExecOrderAction == nil {
		fmt.Println("OnErrRtnExecOrderAction")
	} else {
		t.OnErrRtnExecOrderAction((*def.CThostFtdcExecOrderActionField)(unsafe.Pointer(pExecOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnForQuoteInsert
func OnErrRtnForQuoteInsert(pInputForQuote *C.struct_CThostFtdcInputForQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnForQuoteInsert == nil {
		fmt.Println("OnErrRtnForQuoteInsert")
	} else {
		t.OnErrRtnForQuoteInsert((*def.CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRtnQuote
func OnRtnQuote(pQuote *C.struct_CThostFtdcQuoteField) {
	if t.OnRtnQuote == nil {
		fmt.Println("OnRtnQuote")
	} else {
		t.OnRtnQuote((*def.CThostFtdcQuoteField)(unsafe.Pointer(pQuote)))
	}
}

//export OnErrRtnQuoteInsert
func OnErrRtnQuoteInsert(pInputQuote *C.struct_CThostFtdcInputQuoteField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnQuoteInsert == nil {
		fmt.Println("OnErrRtnQuoteInsert")
	} else {
		t.OnErrRtnQuoteInsert((*def.CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnQuoteAction
func OnErrRtnQuoteAction(pQuoteAction *C.struct_CThostFtdcQuoteActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnQuoteAction == nil {
		fmt.Println("OnErrRtnQuoteAction")
	} else {
		t.OnErrRtnQuoteAction((*def.CThostFtdcQuoteActionField)(unsafe.Pointer(pQuoteAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRtnForQuoteRsp
func OnRtnForQuoteRsp(pForQuoteRsp *C.struct_CThostFtdcForQuoteRspField) {
	if t.OnRtnForQuoteRsp == nil {
		fmt.Println("OnRtnForQuoteRsp")
	} else {
		t.OnRtnForQuoteRsp((*def.CThostFtdcForQuoteRspField)(unsafe.Pointer(pForQuoteRsp)))
	}
}

//export OnRtnCFMMCTradingAccountToken
func OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *C.struct_CThostFtdcCFMMCTradingAccountTokenField) {
	if t.OnRtnCFMMCTradingAccountToken == nil {
		fmt.Println("OnRtnCFMMCTradingAccountToken")
	} else {
		t.OnRtnCFMMCTradingAccountToken((*def.CThostFtdcCFMMCTradingAccountTokenField)(unsafe.Pointer(pCFMMCTradingAccountToken)))
	}
}

//export OnErrRtnBatchOrderAction
func OnErrRtnBatchOrderAction(pBatchOrderAction *C.struct_CThostFtdcBatchOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnBatchOrderAction == nil {
		fmt.Println("OnErrRtnBatchOrderAction")
	} else {
		t.OnErrRtnBatchOrderAction((*def.CThostFtdcBatchOrderActionField)(unsafe.Pointer(pBatchOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRtnOptionSelfClose
func OnRtnOptionSelfClose(pOptionSelfClose *C.struct_CThostFtdcOptionSelfCloseField) {
	if t.OnRtnOptionSelfClose == nil {
		fmt.Println("OnRtnOptionSelfClose")
	} else {
		t.OnRtnOptionSelfClose((*def.CThostFtdcOptionSelfCloseField)(unsafe.Pointer(pOptionSelfClose)))
	}
}

//export OnErrRtnOptionSelfCloseInsert
func OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *C.struct_CThostFtdcInputOptionSelfCloseField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnOptionSelfCloseInsert == nil {
		fmt.Println("OnErrRtnOptionSelfCloseInsert")
	} else {
		t.OnErrRtnOptionSelfCloseInsert((*def.CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnOptionSelfCloseAction
func OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *C.struct_CThostFtdcOptionSelfCloseActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnOptionSelfCloseAction == nil {
		fmt.Println("OnErrRtnOptionSelfCloseAction")
	} else {
		t.OnErrRtnOptionSelfCloseAction((*def.CThostFtdcOptionSelfCloseActionField)(unsafe.Pointer(pOptionSelfCloseAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRtnCombAction
func OnRtnCombAction(pCombAction *C.struct_CThostFtdcCombActionField) {
	if t.OnRtnCombAction == nil {
		fmt.Println("OnRtnCombAction")
	} else {
		t.OnRtnCombAction((*def.CThostFtdcCombActionField)(unsafe.Pointer(pCombAction)))
	}
}

//export OnErrRtnCombActionInsert
func OnErrRtnCombActionInsert(pInputCombAction *C.struct_CThostFtdcInputCombActionField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnCombActionInsert == nil {
		fmt.Println("OnErrRtnCombActionInsert")
	} else {
		t.OnErrRtnCombActionInsert((*def.CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRspQryContractBank
func OnRspQryContractBank(pContractBank *C.struct_CThostFtdcContractBankField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryContractBank == nil {
		fmt.Println("OnRspQryContractBank")
	} else {
		t.OnRspQryContractBank((*def.CThostFtdcContractBankField)(unsafe.Pointer(pContractBank)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryParkedOrder
func OnRspQryParkedOrder(pParkedOrder *C.struct_CThostFtdcParkedOrderField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryParkedOrder == nil {
		fmt.Println("OnRspQryParkedOrder")
	} else {
		t.OnRspQryParkedOrder((*def.CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryParkedOrderAction
func OnRspQryParkedOrderAction(pParkedOrderAction *C.struct_CThostFtdcParkedOrderActionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryParkedOrderAction == nil {
		fmt.Println("OnRspQryParkedOrderAction")
	} else {
		t.OnRspQryParkedOrderAction((*def.CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryTradingNotice
func OnRspQryTradingNotice(pTradingNotice *C.struct_CThostFtdcTradingNoticeField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryTradingNotice == nil {
		fmt.Println("OnRspQryTradingNotice")
	} else {
		t.OnRspQryTradingNotice((*def.CThostFtdcTradingNoticeField)(unsafe.Pointer(pTradingNotice)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryBrokerTradingParams
func OnRspQryBrokerTradingParams(pBrokerTradingParams *C.struct_CThostFtdcBrokerTradingParamsField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryBrokerTradingParams == nil {
		fmt.Println("OnRspQryBrokerTradingParams")
	} else {
		t.OnRspQryBrokerTradingParams((*def.CThostFtdcBrokerTradingParamsField)(unsafe.Pointer(pBrokerTradingParams)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryBrokerTradingAlgos
func OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *C.struct_CThostFtdcBrokerTradingAlgosField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryBrokerTradingAlgos == nil {
		fmt.Println("OnRspQryBrokerTradingAlgos")
	} else {
		t.OnRspQryBrokerTradingAlgos((*def.CThostFtdcBrokerTradingAlgosField)(unsafe.Pointer(pBrokerTradingAlgos)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQueryCFMMCTradingAccountToken
func OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQueryCFMMCTradingAccountToken == nil {
		fmt.Println("OnRspQueryCFMMCTradingAccountToken")
	} else {
		t.OnRspQueryCFMMCTradingAccountToken((*def.CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRtnFromBankToFutureByBank
func OnRtnFromBankToFutureByBank(pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	if t.OnRtnFromBankToFutureByBank == nil {
		fmt.Println("OnRtnFromBankToFutureByBank")
	} else {
		t.OnRtnFromBankToFutureByBank((*def.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}

//export OnRtnFromFutureToBankByBank
func OnRtnFromFutureToBankByBank(pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	if t.OnRtnFromFutureToBankByBank == nil {
		fmt.Println("OnRtnFromFutureToBankByBank")
	} else {
		t.OnRtnFromFutureToBankByBank((*def.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}

//export OnRtnRepealFromBankToFutureByBank
func OnRtnRepealFromBankToFutureByBank(pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	if t.OnRtnRepealFromBankToFutureByBank == nil {
		fmt.Println("OnRtnRepealFromBankToFutureByBank")
	} else {
		t.OnRtnRepealFromBankToFutureByBank((*def.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}

//export OnRtnRepealFromFutureToBankByBank
func OnRtnRepealFromFutureToBankByBank(pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	if t.OnRtnRepealFromFutureToBankByBank == nil {
		fmt.Println("OnRtnRepealFromFutureToBankByBank")
	} else {
		t.OnRtnRepealFromFutureToBankByBank((*def.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}

//export OnRtnFromBankToFutureByFuture
func OnRtnFromBankToFutureByFuture(pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	if t.OnRtnFromBankToFutureByFuture == nil {
		fmt.Println("OnRtnFromBankToFutureByFuture")
	} else {
		t.OnRtnFromBankToFutureByFuture((*def.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}

//export OnRtnFromFutureToBankByFuture
func OnRtnFromFutureToBankByFuture(pRspTransfer *C.struct_CThostFtdcRspTransferField) {
	if t.OnRtnFromFutureToBankByFuture == nil {
		fmt.Println("OnRtnFromFutureToBankByFuture")
	} else {
		t.OnRtnFromFutureToBankByFuture((*def.CThostFtdcRspTransferField)(unsafe.Pointer(pRspTransfer)))
	}
}

//export OnRtnRepealFromBankToFutureByFutureManual
func OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	if t.OnRtnRepealFromBankToFutureByFutureManual == nil {
		fmt.Println("OnRtnRepealFromBankToFutureByFutureManual")
	} else {
		t.OnRtnRepealFromBankToFutureByFutureManual((*def.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}

//export OnRtnRepealFromFutureToBankByFutureManual
func OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	if t.OnRtnRepealFromFutureToBankByFutureManual == nil {
		fmt.Println("OnRtnRepealFromFutureToBankByFutureManual")
	} else {
		t.OnRtnRepealFromFutureToBankByFutureManual((*def.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}

//export OnRtnQueryBankBalanceByFuture
func OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *C.struct_CThostFtdcNotifyQueryAccountField) {
	if t.OnRtnQueryBankBalanceByFuture == nil {
		fmt.Println("OnRtnQueryBankBalanceByFuture")
	} else {
		t.OnRtnQueryBankBalanceByFuture((*def.CThostFtdcNotifyQueryAccountField)(unsafe.Pointer(pNotifyQueryAccount)))
	}
}

//export OnErrRtnBankToFutureByFuture
func OnErrRtnBankToFutureByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnBankToFutureByFuture == nil {
		fmt.Println("OnErrRtnBankToFutureByFuture")
	} else {
		t.OnErrRtnBankToFutureByFuture((*def.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnFutureToBankByFuture
func OnErrRtnFutureToBankByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnFutureToBankByFuture == nil {
		fmt.Println("OnErrRtnFutureToBankByFuture")
	} else {
		t.OnErrRtnFutureToBankByFuture((*def.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnRepealBankToFutureByFutureManual
func OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *C.struct_CThostFtdcReqRepealField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnRepealBankToFutureByFutureManual == nil {
		fmt.Println("OnErrRtnRepealBankToFutureByFutureManual")
	} else {
		t.OnErrRtnRepealBankToFutureByFutureManual((*def.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnRepealFutureToBankByFutureManual
func OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *C.struct_CThostFtdcReqRepealField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnRepealFutureToBankByFutureManual == nil {
		fmt.Println("OnErrRtnRepealFutureToBankByFutureManual")
	} else {
		t.OnErrRtnRepealFutureToBankByFutureManual((*def.CThostFtdcReqRepealField)(unsafe.Pointer(pReqRepeal)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnErrRtnQueryBankBalanceByFuture
func OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField) {
	if t.OnErrRtnQueryBankBalanceByFuture == nil {
		fmt.Println("OnErrRtnQueryBankBalanceByFuture")
	} else {
		t.OnErrRtnQueryBankBalanceByFuture((*def.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)))
	}
}

//export OnRtnRepealFromBankToFutureByFuture
func OnRtnRepealFromBankToFutureByFuture(pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	if t.OnRtnRepealFromBankToFutureByFuture == nil {
		fmt.Println("OnRtnRepealFromBankToFutureByFuture")
	} else {
		t.OnRtnRepealFromBankToFutureByFuture((*def.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}

//export OnRtnRepealFromFutureToBankByFuture
func OnRtnRepealFromFutureToBankByFuture(pRspRepeal *C.struct_CThostFtdcRspRepealField) {
	if t.OnRtnRepealFromFutureToBankByFuture == nil {
		fmt.Println("OnRtnRepealFromFutureToBankByFuture")
	} else {
		t.OnRtnRepealFromFutureToBankByFuture((*def.CThostFtdcRspRepealField)(unsafe.Pointer(pRspRepeal)))
	}
}

//export OnRspFromBankToFutureByFuture
func OnRspFromBankToFutureByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspFromBankToFutureByFuture == nil {
		fmt.Println("OnRspFromBankToFutureByFuture")
	} else {
		t.OnRspFromBankToFutureByFuture((*def.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspFromFutureToBankByFuture
func OnRspFromFutureToBankByFuture(pReqTransfer *C.struct_CThostFtdcReqTransferField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspFromFutureToBankByFuture == nil {
		fmt.Println("OnRspFromFutureToBankByFuture")
	} else {
		t.OnRspFromFutureToBankByFuture((*def.CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQueryBankAccountMoneyByFuture
func OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *C.struct_CThostFtdcReqQueryAccountField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQueryBankAccountMoneyByFuture == nil {
		fmt.Println("OnRspQueryBankAccountMoneyByFuture")
	} else {
		t.OnRspQueryBankAccountMoneyByFuture((*def.CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRtnOpenAccountByBank
func OnRtnOpenAccountByBank(pOpenAccount *C.struct_CThostFtdcOpenAccountField) {
	if t.OnRtnOpenAccountByBank == nil {
		fmt.Println("OnRtnOpenAccountByBank")
	} else {
		t.OnRtnOpenAccountByBank((*def.CThostFtdcOpenAccountField)(unsafe.Pointer(pOpenAccount)))
	}
}

//export OnRtnCancelAccountByBank
func OnRtnCancelAccountByBank(pCancelAccount *C.struct_CThostFtdcCancelAccountField) {
	if t.OnRtnCancelAccountByBank == nil {
		fmt.Println("OnRtnCancelAccountByBank")
	} else {
		t.OnRtnCancelAccountByBank((*def.CThostFtdcCancelAccountField)(unsafe.Pointer(pCancelAccount)))
	}
}

//export OnRtnChangeAccountByBank
func OnRtnChangeAccountByBank(pChangeAccount *C.struct_CThostFtdcChangeAccountField) {
	if t.OnRtnChangeAccountByBank == nil {
		fmt.Println("OnRtnChangeAccountByBank")
	} else {
		t.OnRtnChangeAccountByBank((*def.CThostFtdcChangeAccountField)(unsafe.Pointer(pChangeAccount)))
	}
}

//export OnRspQryClassifiedInstrument
func OnRspQryClassifiedInstrument(pInstrument *C.struct_CThostFtdcInstrumentField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryClassifiedInstrument == nil {
		fmt.Println("OnRspQryClassifiedInstrument")
	} else {
		t.OnRspQryClassifiedInstrument((*def.CThostFtdcInstrumentField)(unsafe.Pointer(pInstrument)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryCombPromotionParam
func OnRspQryCombPromotionParam(pCombPromotionParam *C.struct_CThostFtdcCombPromotionParamField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryCombPromotionParam == nil {
		fmt.Println("OnRspQryCombPromotionParam")
	} else {
		t.OnRspQryCombPromotionParam((*def.CThostFtdcCombPromotionParamField)(unsafe.Pointer(pCombPromotionParam)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryRiskSettleInvstPosition
func OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *C.struct_CThostFtdcRiskSettleInvstPositionField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryRiskSettleInvstPosition == nil {
		fmt.Println("OnRspQryRiskSettleInvstPosition")
	} else {
		t.OnRspQryRiskSettleInvstPosition((*def.CThostFtdcRiskSettleInvstPositionField)(unsafe.Pointer(pRiskSettleInvstPosition)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

//export OnRspQryRiskSettleProductStatus
func OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *C.struct_CThostFtdcRiskSettleProductStatusField, pRspInfo *C.struct_CThostFtdcRspInfoField, nRequestID C.int, bIsLast C._Bool) {
	if t.OnRspQryRiskSettleProductStatus == nil {
		fmt.Println("OnRspQryRiskSettleProductStatus")
	} else {
		t.OnRspQryRiskSettleProductStatus((*def.CThostFtdcRiskSettleProductStatusField)(unsafe.Pointer(pRiskSettleProductStatus)), (*def.CThostFtdcRspInfoField)(unsafe.Pointer(pRspInfo)), int(nRequestID), bool(bIsLast))
	}
}

// 创建TraderApi
func (t *Trade) Release() {
	C.tRelease(t.api)
}

// 初始化
func (t *Trade) Init() {
	C.tInit(t.api)
}

// 等待接口线程结束运行
func (t *Trade) Join() {
	C.tJoin(t.api)
}

// 注册前置机网络地址
func (t *Trade) RegisterFront(pszFrontAddress string) {
	C.tRegisterFront(t.api, C.CString(pszFrontAddress))
}

// @remark RegisterNameServer优先于RegisterFront
func (t *Trade) RegisterNameServer(pszNsAddress string) {
	C.tRegisterNameServer(t.api, C.CString(pszNsAddress))
}

// 注册名字服务器用户信息
func (t *Trade) RegisterFensUserInfo(pFensUserInfo *def.CThostFtdcFensUserInfoField) {
	C.tRegisterFensUserInfo(t.api, (*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(pFensUserInfo)))
}

// 注册回调接口
func (t *Trade) RegisterSpi(pSpi unsafe.Pointer) {
	C.tRegisterSpi(t.api, pSpi)
}

// 订阅私有流。
func (t *Trade) SubscribePrivateTopic(nResumeType def.THOST_TE_RESUME_TYPE) {
	C.tSubscribePrivateTopic(t.api, C.int(nResumeType))
}

// 订阅公共流。
func (t *Trade) SubscribePublicTopic(nResumeType def.THOST_TE_RESUME_TYPE) {
	C.tSubscribePublicTopic(t.api, C.int(nResumeType))
}

// 客户端认证请求
func (t *Trade) ReqAuthenticate(pReqAuthenticateField *def.CThostFtdcReqAuthenticateField, nRequestID int) {
	C.tReqAuthenticate(t.api, (*C.struct_CThostFtdcReqAuthenticateField)(unsafe.Pointer(pReqAuthenticateField)), C.int(nRequestID))
}

// 注册用户终端信息，用于中继服务器多连接模式
func (t *Trade) RegisterUserSystemInfo(pUserSystemInfo *def.CThostFtdcUserSystemInfoField) {
	C.tRegisterUserSystemInfo(t.api, (*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo)))
}

// 上报用户终端信息，用于中继服务器操作员登录模式
func (t *Trade) SubmitUserSystemInfo(pUserSystemInfo *def.CThostFtdcUserSystemInfoField) {
	C.tSubmitUserSystemInfo(t.api, (*C.struct_CThostFtdcUserSystemInfoField)(unsafe.Pointer(pUserSystemInfo)))
}

// 用户登录请求
func (t *Trade) ReqUserLogin(pReqUserLoginField *def.CThostFtdcReqUserLoginField, nRequestID int) {
	C.tReqUserLogin(t.api, (*C.struct_CThostFtdcReqUserLoginField)(unsafe.Pointer(pReqUserLoginField)), C.int(nRequestID))
}

// 登出请求
func (t *Trade) ReqUserLogout(pUserLogout *def.CThostFtdcUserLogoutField, nRequestID int) {
	C.tReqUserLogout(t.api, (*C.struct_CThostFtdcUserLogoutField)(unsafe.Pointer(pUserLogout)), C.int(nRequestID))
}

// 用户口令更新请求
func (t *Trade) ReqUserPasswordUpdate(pUserPasswordUpdate *def.CThostFtdcUserPasswordUpdateField, nRequestID int) {
	C.tReqUserPasswordUpdate(t.api, (*C.struct_CThostFtdcUserPasswordUpdateField)(unsafe.Pointer(pUserPasswordUpdate)), C.int(nRequestID))
}

// 资金账户口令更新请求
func (t *Trade) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *def.CThostFtdcTradingAccountPasswordUpdateField, nRequestID int) {
	C.tReqTradingAccountPasswordUpdate(t.api, (*C.struct_CThostFtdcTradingAccountPasswordUpdateField)(unsafe.Pointer(pTradingAccountPasswordUpdate)), C.int(nRequestID))
}

// 查询用户当前支持的认证模式
func (t *Trade) ReqUserAuthMethod(pReqUserAuthMethod *def.CThostFtdcReqUserAuthMethodField, nRequestID int) {
	C.tReqUserAuthMethod(t.api, (*C.struct_CThostFtdcReqUserAuthMethodField)(unsafe.Pointer(pReqUserAuthMethod)), C.int(nRequestID))
}

// 用户发出获取图形验证码请求
func (t *Trade) ReqGenUserCaptcha(pReqGenUserCaptcha *def.CThostFtdcReqGenUserCaptchaField, nRequestID int) {
	C.tReqGenUserCaptcha(t.api, (*C.struct_CThostFtdcReqGenUserCaptchaField)(unsafe.Pointer(pReqGenUserCaptcha)), C.int(nRequestID))
}

// 用户发出获取短信验证码请求
func (t *Trade) ReqGenUserText(pReqGenUserText *def.CThostFtdcReqGenUserTextField, nRequestID int) {
	C.tReqGenUserText(t.api, (*C.struct_CThostFtdcReqGenUserTextField)(unsafe.Pointer(pReqGenUserText)), C.int(nRequestID))
}

// 用户发出带有图片验证码的登陆请求
func (t *Trade) ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha *def.CThostFtdcReqUserLoginWithCaptchaField, nRequestID int) {
	C.tReqUserLoginWithCaptcha(t.api, (*C.struct_CThostFtdcReqUserLoginWithCaptchaField)(unsafe.Pointer(pReqUserLoginWithCaptcha)), C.int(nRequestID))
}

// 用户发出带有短信验证码的登陆请求
func (t *Trade) ReqUserLoginWithText(pReqUserLoginWithText *def.CThostFtdcReqUserLoginWithTextField, nRequestID int) {
	C.tReqUserLoginWithText(t.api, (*C.struct_CThostFtdcReqUserLoginWithTextField)(unsafe.Pointer(pReqUserLoginWithText)), C.int(nRequestID))
}

// 用户发出带有动态口令的登陆请求
func (t *Trade) ReqUserLoginWithOTP(pReqUserLoginWithOTP *def.CThostFtdcReqUserLoginWithOTPField, nRequestID int) {
	C.tReqUserLoginWithOTP(t.api, (*C.struct_CThostFtdcReqUserLoginWithOTPField)(unsafe.Pointer(pReqUserLoginWithOTP)), C.int(nRequestID))
}

// 报单录入请求
func (t *Trade) ReqOrderInsert(pInputOrder *def.CThostFtdcInputOrderField, nRequestID int) {
	C.tReqOrderInsert(t.api, (*C.struct_CThostFtdcInputOrderField)(unsafe.Pointer(pInputOrder)), C.int(nRequestID))
}

// 预埋单录入请求
func (t *Trade) ReqParkedOrderInsert(pParkedOrder *def.CThostFtdcParkedOrderField, nRequestID int) {
	C.tReqParkedOrderInsert(t.api, (*C.struct_CThostFtdcParkedOrderField)(unsafe.Pointer(pParkedOrder)), C.int(nRequestID))
}

// 预埋撤单录入请求
func (t *Trade) ReqParkedOrderAction(pParkedOrderAction *def.CThostFtdcParkedOrderActionField, nRequestID int) {
	C.tReqParkedOrderAction(t.api, (*C.struct_CThostFtdcParkedOrderActionField)(unsafe.Pointer(pParkedOrderAction)), C.int(nRequestID))
}

// 报单操作请求
func (t *Trade) ReqOrderAction(pInputOrderAction *def.CThostFtdcInputOrderActionField, nRequestID int) {
	C.tReqOrderAction(t.api, (*C.struct_CThostFtdcInputOrderActionField)(unsafe.Pointer(pInputOrderAction)), C.int(nRequestID))
}

// 查询最大报单数量请求
func (t *Trade) ReqQryMaxOrderVolume(pQryMaxOrderVolume *def.CThostFtdcQryMaxOrderVolumeField, nRequestID int) {
	C.tReqQryMaxOrderVolume(t.api, (*C.struct_CThostFtdcQryMaxOrderVolumeField)(unsafe.Pointer(pQryMaxOrderVolume)), C.int(nRequestID))
}

// 投资者结算结果确认
func (t *Trade) ReqSettlementInfoConfirm(pSettlementInfoConfirm *def.CThostFtdcSettlementInfoConfirmField, nRequestID int) {
	C.tReqSettlementInfoConfirm(t.api, (*C.struct_CThostFtdcSettlementInfoConfirmField)(unsafe.Pointer(pSettlementInfoConfirm)), C.int(nRequestID))
}

// 请求删除预埋单
func (t *Trade) ReqRemoveParkedOrder(pRemoveParkedOrder *def.CThostFtdcRemoveParkedOrderField, nRequestID int) {
	C.tReqRemoveParkedOrder(t.api, (*C.struct_CThostFtdcRemoveParkedOrderField)(unsafe.Pointer(pRemoveParkedOrder)), C.int(nRequestID))
}

// 请求删除预埋撤单
func (t *Trade) ReqRemoveParkedOrderAction(pRemoveParkedOrderAction *def.CThostFtdcRemoveParkedOrderActionField, nRequestID int) {
	C.tReqRemoveParkedOrderAction(t.api, (*C.struct_CThostFtdcRemoveParkedOrderActionField)(unsafe.Pointer(pRemoveParkedOrderAction)), C.int(nRequestID))
}

// 执行宣告录入请求
func (t *Trade) ReqExecOrderInsert(pInputExecOrder *def.CThostFtdcInputExecOrderField, nRequestID int) {
	C.tReqExecOrderInsert(t.api, (*C.struct_CThostFtdcInputExecOrderField)(unsafe.Pointer(pInputExecOrder)), C.int(nRequestID))
}

// 执行宣告操作请求
func (t *Trade) ReqExecOrderAction(pInputExecOrderAction *def.CThostFtdcInputExecOrderActionField, nRequestID int) {
	C.tReqExecOrderAction(t.api, (*C.struct_CThostFtdcInputExecOrderActionField)(unsafe.Pointer(pInputExecOrderAction)), C.int(nRequestID))
}

// 询价录入请求
func (t *Trade) ReqForQuoteInsert(pInputForQuote *def.CThostFtdcInputForQuoteField, nRequestID int) {
	C.tReqForQuoteInsert(t.api, (*C.struct_CThostFtdcInputForQuoteField)(unsafe.Pointer(pInputForQuote)), C.int(nRequestID))
}

// 报价录入请求
func (t *Trade) ReqQuoteInsert(pInputQuote *def.CThostFtdcInputQuoteField, nRequestID int) {
	C.tReqQuoteInsert(t.api, (*C.struct_CThostFtdcInputQuoteField)(unsafe.Pointer(pInputQuote)), C.int(nRequestID))
}

// 报价操作请求
func (t *Trade) ReqQuoteAction(pInputQuoteAction *def.CThostFtdcInputQuoteActionField, nRequestID int) {
	C.tReqQuoteAction(t.api, (*C.struct_CThostFtdcInputQuoteActionField)(unsafe.Pointer(pInputQuoteAction)), C.int(nRequestID))
}

// 批量报单操作请求
func (t *Trade) ReqBatchOrderAction(pInputBatchOrderAction *def.CThostFtdcInputBatchOrderActionField, nRequestID int) {
	C.tReqBatchOrderAction(t.api, (*C.struct_CThostFtdcInputBatchOrderActionField)(unsafe.Pointer(pInputBatchOrderAction)), C.int(nRequestID))
}

// 期权自对冲录入请求
func (t *Trade) ReqOptionSelfCloseInsert(pInputOptionSelfClose *def.CThostFtdcInputOptionSelfCloseField, nRequestID int) {
	C.tReqOptionSelfCloseInsert(t.api, (*C.struct_CThostFtdcInputOptionSelfCloseField)(unsafe.Pointer(pInputOptionSelfClose)), C.int(nRequestID))
}

// 期权自对冲操作请求
func (t *Trade) ReqOptionSelfCloseAction(pInputOptionSelfCloseAction *def.CThostFtdcInputOptionSelfCloseActionField, nRequestID int) {
	C.tReqOptionSelfCloseAction(t.api, (*C.struct_CThostFtdcInputOptionSelfCloseActionField)(unsafe.Pointer(pInputOptionSelfCloseAction)), C.int(nRequestID))
}

// 申请组合录入请求
func (t *Trade) ReqCombActionInsert(pInputCombAction *def.CThostFtdcInputCombActionField, nRequestID int) {
	C.tReqCombActionInsert(t.api, (*C.struct_CThostFtdcInputCombActionField)(unsafe.Pointer(pInputCombAction)), C.int(nRequestID))
}

// 请求查询报单
func (t *Trade) ReqQryOrder(pQryOrder *def.CThostFtdcQryOrderField, nRequestID int) {
	C.tReqQryOrder(t.api, (*C.struct_CThostFtdcQryOrderField)(unsafe.Pointer(pQryOrder)), C.int(nRequestID))
}

// 请求查询成交
func (t *Trade) ReqQryTrade(pQryTrade *def.CThostFtdcQryTradeField, nRequestID int) {
	C.tReqQryTrade(t.api, (*C.struct_CThostFtdcQryTradeField)(unsafe.Pointer(pQryTrade)), C.int(nRequestID))
}

// 请求查询投资者持仓
func (t *Trade) ReqQryInvestorPosition(pQryInvestorPosition *def.CThostFtdcQryInvestorPositionField, nRequestID int) {
	C.tReqQryInvestorPosition(t.api, (*C.struct_CThostFtdcQryInvestorPositionField)(unsafe.Pointer(pQryInvestorPosition)), C.int(nRequestID))
}

// 请求查询资金账户
func (t *Trade) ReqQryTradingAccount(pQryTradingAccount *def.CThostFtdcQryTradingAccountField, nRequestID int) {
	C.tReqQryTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(pQryTradingAccount)), C.int(nRequestID))
}

// 请求查询投资者
func (t *Trade) ReqQryInvestor(pQryInvestor *def.CThostFtdcQryInvestorField, nRequestID int) {
	C.tReqQryInvestor(t.api, (*C.struct_CThostFtdcQryInvestorField)(unsafe.Pointer(pQryInvestor)), C.int(nRequestID))
}

// 请求查询交易编码
func (t *Trade) ReqQryTradingCode(pQryTradingCode *def.CThostFtdcQryTradingCodeField, nRequestID int) {
	C.tReqQryTradingCode(t.api, (*C.struct_CThostFtdcQryTradingCodeField)(unsafe.Pointer(pQryTradingCode)), C.int(nRequestID))
}

// 请求查询合约保证金率
func (t *Trade) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *def.CThostFtdcQryInstrumentMarginRateField, nRequestID int) {
	C.tReqQryInstrumentMarginRate(t.api, (*C.struct_CThostFtdcQryInstrumentMarginRateField)(unsafe.Pointer(pQryInstrumentMarginRate)), C.int(nRequestID))
}

// 请求查询合约手续费率
func (t *Trade) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *def.CThostFtdcQryInstrumentCommissionRateField, nRequestID int) {
	C.tReqQryInstrumentCommissionRate(t.api, (*C.struct_CThostFtdcQryInstrumentCommissionRateField)(unsafe.Pointer(pQryInstrumentCommissionRate)), C.int(nRequestID))
}

// 请求查询交易所
func (t *Trade) ReqQryExchange(pQryExchange *def.CThostFtdcQryExchangeField, nRequestID int) {
	C.tReqQryExchange(t.api, (*C.struct_CThostFtdcQryExchangeField)(unsafe.Pointer(pQryExchange)), C.int(nRequestID))
}

// 请求查询产品
func (t *Trade) ReqQryProduct(pQryProduct *def.CThostFtdcQryProductField, nRequestID int) {
	C.tReqQryProduct(t.api, (*C.struct_CThostFtdcQryProductField)(unsafe.Pointer(pQryProduct)), C.int(nRequestID))
}

// 请求查询合约
func (t *Trade) ReqQryInstrument(pQryInstrument *def.CThostFtdcQryInstrumentField, nRequestID int) {
	C.tReqQryInstrument(t.api, (*C.struct_CThostFtdcQryInstrumentField)(unsafe.Pointer(pQryInstrument)), C.int(nRequestID))
}

// 请求查询行情
func (t *Trade) ReqQryDepthMarketData(pQryDepthMarketData *def.CThostFtdcQryDepthMarketDataField, nRequestID int) {
	C.tReqQryDepthMarketData(t.api, (*C.struct_CThostFtdcQryDepthMarketDataField)(unsafe.Pointer(pQryDepthMarketData)), C.int(nRequestID))
}

// 请求查询交易员报盘机
func (t *Trade) ReqQryTraderOffer(pQryTraderOffer *def.CThostFtdcQryTraderOfferField, nRequestID int) {
	C.tReqQryTraderOffer(t.api, (*C.struct_CThostFtdcQryTraderOfferField)(unsafe.Pointer(pQryTraderOffer)), C.int(nRequestID))
}

// 请求查询投资者结算结果
func (t *Trade) ReqQrySettlementInfo(pQrySettlementInfo *def.CThostFtdcQrySettlementInfoField, nRequestID int) {
	C.tReqQrySettlementInfo(t.api, (*C.struct_CThostFtdcQrySettlementInfoField)(unsafe.Pointer(pQrySettlementInfo)), C.int(nRequestID))
}

// 请求查询转帐银行
func (t *Trade) ReqQryTransferBank(pQryTransferBank *def.CThostFtdcQryTransferBankField, nRequestID int) {
	C.tReqQryTransferBank(t.api, (*C.struct_CThostFtdcQryTransferBankField)(unsafe.Pointer(pQryTransferBank)), C.int(nRequestID))
}

// 请求查询投资者持仓明细
func (t *Trade) ReqQryInvestorPositionDetail(pQryInvestorPositionDetail *def.CThostFtdcQryInvestorPositionDetailField, nRequestID int) {
	C.tReqQryInvestorPositionDetail(t.api, (*C.struct_CThostFtdcQryInvestorPositionDetailField)(unsafe.Pointer(pQryInvestorPositionDetail)), C.int(nRequestID))
}

// 请求查询客户通知
func (t *Trade) ReqQryNotice(pQryNotice *def.CThostFtdcQryNoticeField, nRequestID int) {
	C.tReqQryNotice(t.api, (*C.struct_CThostFtdcQryNoticeField)(unsafe.Pointer(pQryNotice)), C.int(nRequestID))
}

// 请求查询结算信息确认
func (t *Trade) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *def.CThostFtdcQrySettlementInfoConfirmField, nRequestID int) {
	C.tReqQrySettlementInfoConfirm(t.api, (*C.struct_CThostFtdcQrySettlementInfoConfirmField)(unsafe.Pointer(pQrySettlementInfoConfirm)), C.int(nRequestID))
}

// 请求查询投资者持仓明细
func (t *Trade) ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail *def.CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int) {
	C.tReqQryInvestorPositionCombineDetail(t.api, (*C.struct_CThostFtdcQryInvestorPositionCombineDetailField)(unsafe.Pointer(pQryInvestorPositionCombineDetail)), C.int(nRequestID))
}

// 请求查询保证金监管系统经纪公司资金账户密钥
func (t *Trade) ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey *def.CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int) {
	C.tReqQryCFMMCTradingAccountKey(t.api, (*C.struct_CThostFtdcQryCFMMCTradingAccountKeyField)(unsafe.Pointer(pQryCFMMCTradingAccountKey)), C.int(nRequestID))
}

// 请求查询仓单折抵信息
func (t *Trade) ReqQryEWarrantOffset(pQryEWarrantOffset *def.CThostFtdcQryEWarrantOffsetField, nRequestID int) {
	C.tReqQryEWarrantOffset(t.api, (*C.struct_CThostFtdcQryEWarrantOffsetField)(unsafe.Pointer(pQryEWarrantOffset)), C.int(nRequestID))
}

// 请求查询投资者品种/跨品种保证金
func (t *Trade) ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin *def.CThostFtdcQryInvestorProductGroupMarginField, nRequestID int) {
	C.tReqQryInvestorProductGroupMargin(t.api, (*C.struct_CThostFtdcQryInvestorProductGroupMarginField)(unsafe.Pointer(pQryInvestorProductGroupMargin)), C.int(nRequestID))
}

// 请求查询交易所保证金率
func (t *Trade) ReqQryExchangeMarginRate(pQryExchangeMarginRate *def.CThostFtdcQryExchangeMarginRateField, nRequestID int) {
	C.tReqQryExchangeMarginRate(t.api, (*C.struct_CThostFtdcQryExchangeMarginRateField)(unsafe.Pointer(pQryExchangeMarginRate)), C.int(nRequestID))
}

// 请求查询交易所调整保证金率
func (t *Trade) ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust *def.CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int) {
	C.tReqQryExchangeMarginRateAdjust(t.api, (*C.struct_CThostFtdcQryExchangeMarginRateAdjustField)(unsafe.Pointer(pQryExchangeMarginRateAdjust)), C.int(nRequestID))
}

// 请求查询汇率
func (t *Trade) ReqQryExchangeRate(pQryExchangeRate *def.CThostFtdcQryExchangeRateField, nRequestID int) {
	C.tReqQryExchangeRate(t.api, (*C.struct_CThostFtdcQryExchangeRateField)(unsafe.Pointer(pQryExchangeRate)), C.int(nRequestID))
}

// 请求查询二级代理操作员银期权限
func (t *Trade) ReqQrySecAgentACIDMap(pQrySecAgentACIDMap *def.CThostFtdcQrySecAgentACIDMapField, nRequestID int) {
	C.tReqQrySecAgentACIDMap(t.api, (*C.struct_CThostFtdcQrySecAgentACIDMapField)(unsafe.Pointer(pQrySecAgentACIDMap)), C.int(nRequestID))
}

// 请求查询产品报价汇率
func (t *Trade) ReqQryProductExchRate(pQryProductExchRate *def.CThostFtdcQryProductExchRateField, nRequestID int) {
	C.tReqQryProductExchRate(t.api, (*C.struct_CThostFtdcQryProductExchRateField)(unsafe.Pointer(pQryProductExchRate)), C.int(nRequestID))
}

// 请求查询产品组
func (t *Trade) ReqQryProductGroup(pQryProductGroup *def.CThostFtdcQryProductGroupField, nRequestID int) {
	C.tReqQryProductGroup(t.api, (*C.struct_CThostFtdcQryProductGroupField)(unsafe.Pointer(pQryProductGroup)), C.int(nRequestID))
}

// 请求查询做市商合约手续费率
func (t *Trade) ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate *def.CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int) {
	C.tReqQryMMInstrumentCommissionRate(t.api, (*C.struct_CThostFtdcQryMMInstrumentCommissionRateField)(unsafe.Pointer(pQryMMInstrumentCommissionRate)), C.int(nRequestID))
}

// 请求查询做市商期权合约手续费
func (t *Trade) ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate *def.CThostFtdcQryMMOptionInstrCommRateField, nRequestID int) {
	C.tReqQryMMOptionInstrCommRate(t.api, (*C.struct_CThostFtdcQryMMOptionInstrCommRateField)(unsafe.Pointer(pQryMMOptionInstrCommRate)), C.int(nRequestID))
}

// 请求查询报单手续费
func (t *Trade) ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate *def.CThostFtdcQryInstrumentOrderCommRateField, nRequestID int) {
	C.tReqQryInstrumentOrderCommRate(t.api, (*C.struct_CThostFtdcQryInstrumentOrderCommRateField)(unsafe.Pointer(pQryInstrumentOrderCommRate)), C.int(nRequestID))
}

// 请求查询资金账户
func (t *Trade) ReqQrySecAgentTradingAccount(pQryTradingAccount *def.CThostFtdcQryTradingAccountField, nRequestID int) {
	C.tReqQrySecAgentTradingAccount(t.api, (*C.struct_CThostFtdcQryTradingAccountField)(unsafe.Pointer(pQryTradingAccount)), C.int(nRequestID))
}

// 请求查询二级代理商资金校验模式
func (t *Trade) ReqQrySecAgentCheckMode(pQrySecAgentCheckMode *def.CThostFtdcQrySecAgentCheckModeField, nRequestID int) {
	C.tReqQrySecAgentCheckMode(t.api, (*C.struct_CThostFtdcQrySecAgentCheckModeField)(unsafe.Pointer(pQrySecAgentCheckMode)), C.int(nRequestID))
}

// 请求查询二级代理商信息
func (t *Trade) ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo *def.CThostFtdcQrySecAgentTradeInfoField, nRequestID int) {
	C.tReqQrySecAgentTradeInfo(t.api, (*C.struct_CThostFtdcQrySecAgentTradeInfoField)(unsafe.Pointer(pQrySecAgentTradeInfo)), C.int(nRequestID))
}

// 请求查询期权交易成本
func (t *Trade) ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost *def.CThostFtdcQryOptionInstrTradeCostField, nRequestID int) {
	C.tReqQryOptionInstrTradeCost(t.api, (*C.struct_CThostFtdcQryOptionInstrTradeCostField)(unsafe.Pointer(pQryOptionInstrTradeCost)), C.int(nRequestID))
}

// 请求查询期权合约手续费
func (t *Trade) ReqQryOptionInstrCommRate(pQryOptionInstrCommRate *def.CThostFtdcQryOptionInstrCommRateField, nRequestID int) {
	C.tReqQryOptionInstrCommRate(t.api, (*C.struct_CThostFtdcQryOptionInstrCommRateField)(unsafe.Pointer(pQryOptionInstrCommRate)), C.int(nRequestID))
}

// 请求查询执行宣告
func (t *Trade) ReqQryExecOrder(pQryExecOrder *def.CThostFtdcQryExecOrderField, nRequestID int) {
	C.tReqQryExecOrder(t.api, (*C.struct_CThostFtdcQryExecOrderField)(unsafe.Pointer(pQryExecOrder)), C.int(nRequestID))
}

// 请求查询询价
func (t *Trade) ReqQryForQuote(pQryForQuote *def.CThostFtdcQryForQuoteField, nRequestID int) {
	C.tReqQryForQuote(t.api, (*C.struct_CThostFtdcQryForQuoteField)(unsafe.Pointer(pQryForQuote)), C.int(nRequestID))
}

// 请求查询报价
func (t *Trade) ReqQryQuote(pQryQuote *def.CThostFtdcQryQuoteField, nRequestID int) {
	C.tReqQryQuote(t.api, (*C.struct_CThostFtdcQryQuoteField)(unsafe.Pointer(pQryQuote)), C.int(nRequestID))
}

// 请求查询期权自对冲
func (t *Trade) ReqQryOptionSelfClose(pQryOptionSelfClose *def.CThostFtdcQryOptionSelfCloseField, nRequestID int) {
	C.tReqQryOptionSelfClose(t.api, (*C.struct_CThostFtdcQryOptionSelfCloseField)(unsafe.Pointer(pQryOptionSelfClose)), C.int(nRequestID))
}

// 请求查询投资单元
func (t *Trade) ReqQryInvestUnit(pQryInvestUnit *def.CThostFtdcQryInvestUnitField, nRequestID int) {
	C.tReqQryInvestUnit(t.api, (*C.struct_CThostFtdcQryInvestUnitField)(unsafe.Pointer(pQryInvestUnit)), C.int(nRequestID))
}

// 请求查询组合合约安全系数
func (t *Trade) ReqQryCombInstrumentGuard(pQryCombInstrumentGuard *def.CThostFtdcQryCombInstrumentGuardField, nRequestID int) {
	C.tReqQryCombInstrumentGuard(t.api, (*C.struct_CThostFtdcQryCombInstrumentGuardField)(unsafe.Pointer(pQryCombInstrumentGuard)), C.int(nRequestID))
}

// 请求查询申请组合
func (t *Trade) ReqQryCombAction(pQryCombAction *def.CThostFtdcQryCombActionField, nRequestID int) {
	C.tReqQryCombAction(t.api, (*C.struct_CThostFtdcQryCombActionField)(unsafe.Pointer(pQryCombAction)), C.int(nRequestID))
}

// 请求查询转帐流水
func (t *Trade) ReqQryTransferSerial(pQryTransferSerial *def.CThostFtdcQryTransferSerialField, nRequestID int) {
	C.tReqQryTransferSerial(t.api, (*C.struct_CThostFtdcQryTransferSerialField)(unsafe.Pointer(pQryTransferSerial)), C.int(nRequestID))
}

// 请求查询银期签约关系
func (t *Trade) ReqQryAccountregister(pQryAccountregister *def.CThostFtdcQryAccountregisterField, nRequestID int) {
	C.tReqQryAccountregister(t.api, (*C.struct_CThostFtdcQryAccountregisterField)(unsafe.Pointer(pQryAccountregister)), C.int(nRequestID))
}

// 请求查询签约银行
func (t *Trade) ReqQryContractBank(pQryContractBank *def.CThostFtdcQryContractBankField, nRequestID int) {
	C.tReqQryContractBank(t.api, (*C.struct_CThostFtdcQryContractBankField)(unsafe.Pointer(pQryContractBank)), C.int(nRequestID))
}

// 请求查询预埋单
func (t *Trade) ReqQryParkedOrder(pQryParkedOrder *def.CThostFtdcQryParkedOrderField, nRequestID int) {
	C.tReqQryParkedOrder(t.api, (*C.struct_CThostFtdcQryParkedOrderField)(unsafe.Pointer(pQryParkedOrder)), C.int(nRequestID))
}

// 请求查询预埋撤单
func (t *Trade) ReqQryParkedOrderAction(pQryParkedOrderAction *def.CThostFtdcQryParkedOrderActionField, nRequestID int) {
	C.tReqQryParkedOrderAction(t.api, (*C.struct_CThostFtdcQryParkedOrderActionField)(unsafe.Pointer(pQryParkedOrderAction)), C.int(nRequestID))
}

// 请求查询交易通知
func (t *Trade) ReqQryTradingNotice(pQryTradingNotice *def.CThostFtdcQryTradingNoticeField, nRequestID int) {
	C.tReqQryTradingNotice(t.api, (*C.struct_CThostFtdcQryTradingNoticeField)(unsafe.Pointer(pQryTradingNotice)), C.int(nRequestID))
}

// 请求查询经纪公司交易参数
func (t *Trade) ReqQryBrokerTradingParams(pQryBrokerTradingParams *def.CThostFtdcQryBrokerTradingParamsField, nRequestID int) {
	C.tReqQryBrokerTradingParams(t.api, (*C.struct_CThostFtdcQryBrokerTradingParamsField)(unsafe.Pointer(pQryBrokerTradingParams)), C.int(nRequestID))
}

// 请求查询经纪公司交易算法
func (t *Trade) ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos *def.CThostFtdcQryBrokerTradingAlgosField, nRequestID int) {
	C.tReqQryBrokerTradingAlgos(t.api, (*C.struct_CThostFtdcQryBrokerTradingAlgosField)(unsafe.Pointer(pQryBrokerTradingAlgos)), C.int(nRequestID))
}

// 请求查询监控中心用户令牌
func (t *Trade) ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *def.CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int) {
	C.tReqQueryCFMMCTradingAccountToken(t.api, (*C.struct_CThostFtdcQueryCFMMCTradingAccountTokenField)(unsafe.Pointer(pQueryCFMMCTradingAccountToken)), C.int(nRequestID))
}

// 期货发起银行资金转期货请求
func (t *Trade) ReqFromBankToFutureByFuture(pReqTransfer *def.CThostFtdcReqTransferField, nRequestID int) {
	C.tReqFromBankToFutureByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), C.int(nRequestID))
}

// 期货发起期货资金转银行请求
func (t *Trade) ReqFromFutureToBankByFuture(pReqTransfer *def.CThostFtdcReqTransferField, nRequestID int) {
	C.tReqFromFutureToBankByFuture(t.api, (*C.struct_CThostFtdcReqTransferField)(unsafe.Pointer(pReqTransfer)), C.int(nRequestID))
}

// 期货发起查询银行余额请求
func (t *Trade) ReqQueryBankAccountMoneyByFuture(pReqQueryAccount *def.CThostFtdcReqQueryAccountField, nRequestID int) {
	C.tReqQueryBankAccountMoneyByFuture(t.api, (*C.struct_CThostFtdcReqQueryAccountField)(unsafe.Pointer(pReqQueryAccount)), C.int(nRequestID))
}

// 请求查询分类合约
func (t *Trade) ReqQryClassifiedInstrument(pQryClassifiedInstrument *def.CThostFtdcQryClassifiedInstrumentField, nRequestID int) {
	C.tReqQryClassifiedInstrument(t.api, (*C.struct_CThostFtdcQryClassifiedInstrumentField)(unsafe.Pointer(pQryClassifiedInstrument)), C.int(nRequestID))
}

// 请求组合优惠比例
func (t *Trade) ReqQryCombPromotionParam(pQryCombPromotionParam *def.CThostFtdcQryCombPromotionParamField, nRequestID int) {
	C.tReqQryCombPromotionParam(t.api, (*C.struct_CThostFtdcQryCombPromotionParamField)(unsafe.Pointer(pQryCombPromotionParam)), C.int(nRequestID))
}

// 投资者风险结算持仓查询
func (t *Trade) ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition *def.CThostFtdcQryRiskSettleInvstPositionField, nRequestID int) {
	C.tReqQryRiskSettleInvstPosition(t.api, (*C.struct_CThostFtdcQryRiskSettleInvstPositionField)(unsafe.Pointer(pQryRiskSettleInvstPosition)), C.int(nRequestID))
}

// 风险结算产品查询
func (t *Trade) ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus *def.CThostFtdcQryRiskSettleProductStatusField, nRequestID int) {
	C.tReqQryRiskSettleProductStatus(t.api, (*C.struct_CThostFtdcQryRiskSettleProductStatusField)(unsafe.Pointer(pQryRiskSettleProductStatus)), C.int(nRequestID))
}
