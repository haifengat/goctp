#define DLL_EXPORT extern "C"

#include "trade.h"
#include <iostream>

using namespace std;

DLL_EXPORT void* CreateFtdcTraderApi(const char *pszFlowPath = ""){
    cout << pszFlowPath << endl;
    return CThostFtdcTraderApi::CreateFtdcTraderApi(pszFlowPath);
}

DLL_EXPORT void* GetVersion(){
    return (void*)CThostFtdcTraderApi::GetApiVersion();
}

// 创建 Trade 实例
DLL_EXPORT void* CreateFtdcTraderSpi() {
    return new Trade(); 
}

// ******** 调用函数 *********
// 创建TraderApi
DLL_EXPORT void Release(CThostFtdcTraderApi *api){
    cout << "Release" << endl;
    return api->Release();
}
// 初始化
DLL_EXPORT void Init(CThostFtdcTraderApi *api){
    cout << "Init" << endl;
    return api->Init();
}
// 等待接口线程结束运行
DLL_EXPORT int Join(CThostFtdcTraderApi *api){
    cout << "Join" << endl;
    return api->Join();
}
// 注册前置机网络地址
DLL_EXPORT void RegisterFront(CThostFtdcTraderApi *api, char *pszFrontAddress){
    cout << "RegisterFront" << endl;
    return api->RegisterFront(pszFrontAddress);
}
// @remark RegisterNameServer优先于RegisterFront
DLL_EXPORT void RegisterNameServer(CThostFtdcTraderApi *api, char *pszNsAddress){
    cout << "RegisterNameServer" << endl;
    return api->RegisterNameServer(pszNsAddress);
}
// 注册名字服务器用户信息
DLL_EXPORT void RegisterFensUserInfo(CThostFtdcTraderApi *api, CThostFtdcFensUserInfoField *pFensUserInfo){
    cout << "RegisterFensUserInfo" << endl;
    return api->RegisterFensUserInfo(pFensUserInfo);
}
// 注册回调接口
DLL_EXPORT void RegisterSpi(CThostFtdcTraderApi *api, CThostFtdcTraderSpi *pSpi){
    cout << "RegisterSpi" << endl;
    return api->RegisterSpi(pSpi);
}
// 订阅私有流。
DLL_EXPORT void SubscribePrivateTopic(CThostFtdcTraderApi *api, THOST_TE_RESUME_TYPE nResumeType){
    cout << "SubscribePrivateTopic" << endl;
    return api->SubscribePrivateTopic(nResumeType);
}
// 订阅公共流。
DLL_EXPORT void SubscribePublicTopic(CThostFtdcTraderApi *api, THOST_TE_RESUME_TYPE nResumeType){
    cout << "SubscribePublicTopic" << endl;
    return api->SubscribePublicTopic(nResumeType);
}
// 客户端认证请求
DLL_EXPORT int ReqAuthenticate(CThostFtdcTraderApi *api, CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID){
    cout << "ReqAuthenticate" << endl;
    return api->ReqAuthenticate(pReqAuthenticateField, nRequestID);
}
// 注册用户终端信息，用于中继服务器多连接模式
DLL_EXPORT int RegisterUserSystemInfo(CThostFtdcTraderApi *api, CThostFtdcUserSystemInfoField *pUserSystemInfo){
    cout << "RegisterUserSystemInfo" << endl;
    return api->RegisterUserSystemInfo(pUserSystemInfo);
}
// 上报用户终端信息，用于中继服务器操作员登录模式
DLL_EXPORT int SubmitUserSystemInfo(CThostFtdcTraderApi *api, CThostFtdcUserSystemInfoField *pUserSystemInfo){
    cout << "SubmitUserSystemInfo" << endl;
    return api->SubmitUserSystemInfo(pUserSystemInfo);
}
// 用户登录请求
DLL_EXPORT int ReqUserLogin(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID){
    cout << "ReqUserLogin" << endl;
    return api->ReqUserLogin(pReqUserLoginField, nRequestID);
}
// 登出请求
DLL_EXPORT int ReqUserLogout(CThostFtdcTraderApi *api, CThostFtdcUserLogoutField *pUserLogout, int nRequestID){
    cout << "ReqUserLogout" << endl;
    return api->ReqUserLogout(pUserLogout, nRequestID);
}
// 用户口令更新请求
DLL_EXPORT int ReqUserPasswordUpdate(CThostFtdcTraderApi *api, CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID){
    cout << "ReqUserPasswordUpdate" << endl;
    return api->ReqUserPasswordUpdate(pUserPasswordUpdate, nRequestID);
}
// 资金账户口令更新请求
DLL_EXPORT int ReqTradingAccountPasswordUpdate(CThostFtdcTraderApi *api, CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID){
    cout << "ReqTradingAccountPasswordUpdate" << endl;
    return api->ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate, nRequestID);
}
// 查询用户当前支持的认证模式
DLL_EXPORT int ReqUserAuthMethod(CThostFtdcTraderApi *api, CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID){
    cout << "ReqUserAuthMethod" << endl;
    return api->ReqUserAuthMethod(pReqUserAuthMethod, nRequestID);
}
// 用户发出获取图形验证码请求
DLL_EXPORT int ReqGenUserCaptcha(CThostFtdcTraderApi *api, CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID){
    cout << "ReqGenUserCaptcha" << endl;
    return api->ReqGenUserCaptcha(pReqGenUserCaptcha, nRequestID);
}
// 用户发出获取短信验证码请求
DLL_EXPORT int ReqGenUserText(CThostFtdcTraderApi *api, CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID){
    cout << "ReqGenUserText" << endl;
    return api->ReqGenUserText(pReqGenUserText, nRequestID);
}
// 用户发出带有图片验证码的登陆请求
DLL_EXPORT int ReqUserLoginWithCaptcha(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID){
    cout << "ReqUserLoginWithCaptcha" << endl;
    return api->ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha, nRequestID);
}
// 用户发出带有短信验证码的登陆请求
DLL_EXPORT int ReqUserLoginWithText(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID){
    cout << "ReqUserLoginWithText" << endl;
    return api->ReqUserLoginWithText(pReqUserLoginWithText, nRequestID);
}
// 用户发出带有动态口令的登陆请求
DLL_EXPORT int ReqUserLoginWithOTP(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID){
    cout << "ReqUserLoginWithOTP" << endl;
    return api->ReqUserLoginWithOTP(pReqUserLoginWithOTP, nRequestID);
}
// 报单录入请求
DLL_EXPORT int ReqOrderInsert(CThostFtdcTraderApi *api, CThostFtdcInputOrderField *pInputOrder, int nRequestID){
    cout << "ReqOrderInsert" << endl;
    return api->ReqOrderInsert(pInputOrder, nRequestID);
}
// 预埋单录入请求
DLL_EXPORT int ReqParkedOrderInsert(CThostFtdcTraderApi *api, CThostFtdcParkedOrderField *pParkedOrder, int nRequestID){
    cout << "ReqParkedOrderInsert" << endl;
    return api->ReqParkedOrderInsert(pParkedOrder, nRequestID);
}
// 预埋撤单录入请求
DLL_EXPORT int ReqParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID){
    cout << "ReqParkedOrderAction" << endl;
    return api->ReqParkedOrderAction(pParkedOrderAction, nRequestID);
}
// 报单操作请求
DLL_EXPORT int ReqOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID){
    cout << "ReqOrderAction" << endl;
    return api->ReqOrderAction(pInputOrderAction, nRequestID);
}
// 查询最大报单数量请求
DLL_EXPORT int ReqQryMaxOrderVolume(CThostFtdcTraderApi *api, CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID){
    cout << "ReqQryMaxOrderVolume" << endl;
    return api->ReqQryMaxOrderVolume(pQryMaxOrderVolume, nRequestID);
}
// 投资者结算结果确认
DLL_EXPORT int ReqSettlementInfoConfirm(CThostFtdcTraderApi *api, CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID){
    cout << "ReqSettlementInfoConfirm" << endl;
    return api->ReqSettlementInfoConfirm(pSettlementInfoConfirm, nRequestID);
}
// 请求删除预埋单
DLL_EXPORT int ReqRemoveParkedOrder(CThostFtdcTraderApi *api, CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID){
    cout << "ReqRemoveParkedOrder" << endl;
    return api->ReqRemoveParkedOrder(pRemoveParkedOrder, nRequestID);
}
// 请求删除预埋撤单
DLL_EXPORT int ReqRemoveParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID){
    cout << "ReqRemoveParkedOrderAction" << endl;
    return api->ReqRemoveParkedOrderAction(pRemoveParkedOrderAction, nRequestID);
}
// 执行宣告录入请求
DLL_EXPORT int ReqExecOrderInsert(CThostFtdcTraderApi *api, CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID){
    cout << "ReqExecOrderInsert" << endl;
    return api->ReqExecOrderInsert(pInputExecOrder, nRequestID);
}
// 执行宣告操作请求
DLL_EXPORT int ReqExecOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID){
    cout << "ReqExecOrderAction" << endl;
    return api->ReqExecOrderAction(pInputExecOrderAction, nRequestID);
}
// 询价录入请求
DLL_EXPORT int ReqForQuoteInsert(CThostFtdcTraderApi *api, CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID){
    cout << "ReqForQuoteInsert" << endl;
    return api->ReqForQuoteInsert(pInputForQuote, nRequestID);
}
// 报价录入请求
DLL_EXPORT int ReqQuoteInsert(CThostFtdcTraderApi *api, CThostFtdcInputQuoteField *pInputQuote, int nRequestID){
    cout << "ReqQuoteInsert" << endl;
    return api->ReqQuoteInsert(pInputQuote, nRequestID);
}
// 报价操作请求
DLL_EXPORT int ReqQuoteAction(CThostFtdcTraderApi *api, CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID){
    cout << "ReqQuoteAction" << endl;
    return api->ReqQuoteAction(pInputQuoteAction, nRequestID);
}
// 批量报单操作请求
DLL_EXPORT int ReqBatchOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID){
    cout << "ReqBatchOrderAction" << endl;
    return api->ReqBatchOrderAction(pInputBatchOrderAction, nRequestID);
}
// 期权自对冲录入请求
DLL_EXPORT int ReqOptionSelfCloseInsert(CThostFtdcTraderApi *api, CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID){
    cout << "ReqOptionSelfCloseInsert" << endl;
    return api->ReqOptionSelfCloseInsert(pInputOptionSelfClose, nRequestID);
}
// 期权自对冲操作请求
DLL_EXPORT int ReqOptionSelfCloseAction(CThostFtdcTraderApi *api, CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID){
    cout << "ReqOptionSelfCloseAction" << endl;
    return api->ReqOptionSelfCloseAction(pInputOptionSelfCloseAction, nRequestID);
}
// 申请组合录入请求
DLL_EXPORT int ReqCombActionInsert(CThostFtdcTraderApi *api, CThostFtdcInputCombActionField *pInputCombAction, int nRequestID){
    cout << "ReqCombActionInsert" << endl;
    return api->ReqCombActionInsert(pInputCombAction, nRequestID);
}
// 请求查询报单
DLL_EXPORT int ReqQryOrder(CThostFtdcTraderApi *api, CThostFtdcQryOrderField *pQryOrder, int nRequestID){
    cout << "ReqQryOrder" << endl;
    return api->ReqQryOrder(pQryOrder, nRequestID);
}
// 请求查询成交
DLL_EXPORT int ReqQryTrade(CThostFtdcTraderApi *api, CThostFtdcQryTradeField *pQryTrade, int nRequestID){
    cout << "ReqQryTrade" << endl;
    return api->ReqQryTrade(pQryTrade, nRequestID);
}
// 请求查询投资者持仓
DLL_EXPORT int ReqQryInvestorPosition(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID){
    cout << "ReqQryInvestorPosition" << endl;
    return api->ReqQryInvestorPosition(pQryInvestorPosition, nRequestID);
}
// 请求查询资金账户
DLL_EXPORT int ReqQryTradingAccount(CThostFtdcTraderApi *api, CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID){
    cout << "ReqQryTradingAccount" << endl;
    return api->ReqQryTradingAccount(pQryTradingAccount, nRequestID);
}
// 请求查询投资者
DLL_EXPORT int ReqQryInvestor(CThostFtdcTraderApi *api, CThostFtdcQryInvestorField *pQryInvestor, int nRequestID){
    cout << "ReqQryInvestor" << endl;
    return api->ReqQryInvestor(pQryInvestor, nRequestID);
}
// 请求查询交易编码
DLL_EXPORT int ReqQryTradingCode(CThostFtdcTraderApi *api, CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID){
    cout << "ReqQryTradingCode" << endl;
    return api->ReqQryTradingCode(pQryTradingCode, nRequestID);
}
// 请求查询合约保证金率
DLL_EXPORT int ReqQryInstrumentMarginRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID){
    cout << "ReqQryInstrumentMarginRate" << endl;
    return api->ReqQryInstrumentMarginRate(pQryInstrumentMarginRate, nRequestID);
}
// 请求查询合约手续费率
DLL_EXPORT int ReqQryInstrumentCommissionRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID){
    cout << "ReqQryInstrumentCommissionRate" << endl;
    return api->ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate, nRequestID);
}
// 请求查询交易所
DLL_EXPORT int ReqQryExchange(CThostFtdcTraderApi *api, CThostFtdcQryExchangeField *pQryExchange, int nRequestID){
    cout << "ReqQryExchange" << endl;
    return api->ReqQryExchange(pQryExchange, nRequestID);
}
// 请求查询产品
DLL_EXPORT int ReqQryProduct(CThostFtdcTraderApi *api, CThostFtdcQryProductField *pQryProduct, int nRequestID){
    cout << "ReqQryProduct" << endl;
    return api->ReqQryProduct(pQryProduct, nRequestID);
}
// 请求查询合约
DLL_EXPORT int ReqQryInstrument(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID){
    cout << "ReqQryInstrument" << endl;
    return api->ReqQryInstrument(pQryInstrument, nRequestID);
}
// 请求查询行情
DLL_EXPORT int ReqQryDepthMarketData(CThostFtdcTraderApi *api, CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID){
    cout << "ReqQryDepthMarketData" << endl;
    return api->ReqQryDepthMarketData(pQryDepthMarketData, nRequestID);
}
// 请求查询交易员报盘机
DLL_EXPORT int ReqQryTraderOffer(CThostFtdcTraderApi *api, CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID){
    cout << "ReqQryTraderOffer" << endl;
    return api->ReqQryTraderOffer(pQryTraderOffer, nRequestID);
}
// 请求查询投资者结算结果
DLL_EXPORT int ReqQrySettlementInfo(CThostFtdcTraderApi *api, CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID){
    cout << "ReqQrySettlementInfo" << endl;
    return api->ReqQrySettlementInfo(pQrySettlementInfo, nRequestID);
}
// 请求查询转帐银行
DLL_EXPORT int ReqQryTransferBank(CThostFtdcTraderApi *api, CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID){
    cout << "ReqQryTransferBank" << endl;
    return api->ReqQryTransferBank(pQryTransferBank, nRequestID);
}
// 请求查询投资者持仓明细
DLL_EXPORT int ReqQryInvestorPositionDetail(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID){
    cout << "ReqQryInvestorPositionDetail" << endl;
    return api->ReqQryInvestorPositionDetail(pQryInvestorPositionDetail, nRequestID);
}
// 请求查询客户通知
DLL_EXPORT int ReqQryNotice(CThostFtdcTraderApi *api, CThostFtdcQryNoticeField *pQryNotice, int nRequestID){
    cout << "ReqQryNotice" << endl;
    return api->ReqQryNotice(pQryNotice, nRequestID);
}
// 请求查询结算信息确认
DLL_EXPORT int ReqQrySettlementInfoConfirm(CThostFtdcTraderApi *api, CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID){
    cout << "ReqQrySettlementInfoConfirm" << endl;
    return api->ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm, nRequestID);
}
// 请求查询投资者持仓明细
DLL_EXPORT int ReqQryInvestorPositionCombineDetail(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID){
    cout << "ReqQryInvestorPositionCombineDetail" << endl;
    return api->ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail, nRequestID);
}
// 请求查询保证金监管系统经纪公司资金账户密钥
DLL_EXPORT int ReqQryCFMMCTradingAccountKey(CThostFtdcTraderApi *api, CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID){
    cout << "ReqQryCFMMCTradingAccountKey" << endl;
    return api->ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey, nRequestID);
}
// 请求查询仓单折抵信息
DLL_EXPORT int ReqQryEWarrantOffset(CThostFtdcTraderApi *api, CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID){
    cout << "ReqQryEWarrantOffset" << endl;
    return api->ReqQryEWarrantOffset(pQryEWarrantOffset, nRequestID);
}
// 请求查询投资者品种/跨品种保证金
DLL_EXPORT int ReqQryInvestorProductGroupMargin(CThostFtdcTraderApi *api, CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID){
    cout << "ReqQryInvestorProductGroupMargin" << endl;
    return api->ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin, nRequestID);
}
// 请求查询交易所保证金率
DLL_EXPORT int ReqQryExchangeMarginRate(CThostFtdcTraderApi *api, CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID){
    cout << "ReqQryExchangeMarginRate" << endl;
    return api->ReqQryExchangeMarginRate(pQryExchangeMarginRate, nRequestID);
}
// 请求查询交易所调整保证金率
DLL_EXPORT int ReqQryExchangeMarginRateAdjust(CThostFtdcTraderApi *api, CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID){
    cout << "ReqQryExchangeMarginRateAdjust" << endl;
    return api->ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust, nRequestID);
}
// 请求查询汇率
DLL_EXPORT int ReqQryExchangeRate(CThostFtdcTraderApi *api, CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID){
    cout << "ReqQryExchangeRate" << endl;
    return api->ReqQryExchangeRate(pQryExchangeRate, nRequestID);
}
// 请求查询二级代理操作员银期权限
DLL_EXPORT int ReqQrySecAgentACIDMap(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID){
    cout << "ReqQrySecAgentACIDMap" << endl;
    return api->ReqQrySecAgentACIDMap(pQrySecAgentACIDMap, nRequestID);
}
// 请求查询产品报价汇率
DLL_EXPORT int ReqQryProductExchRate(CThostFtdcTraderApi *api, CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID){
    cout << "ReqQryProductExchRate" << endl;
    return api->ReqQryProductExchRate(pQryProductExchRate, nRequestID);
}
// 请求查询产品组
DLL_EXPORT int ReqQryProductGroup(CThostFtdcTraderApi *api, CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID){
    cout << "ReqQryProductGroup" << endl;
    return api->ReqQryProductGroup(pQryProductGroup, nRequestID);
}
// 请求查询做市商合约手续费率
DLL_EXPORT int ReqQryMMInstrumentCommissionRate(CThostFtdcTraderApi *api, CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID){
    cout << "ReqQryMMInstrumentCommissionRate" << endl;
    return api->ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate, nRequestID);
}
// 请求查询做市商期权合约手续费
DLL_EXPORT int ReqQryMMOptionInstrCommRate(CThostFtdcTraderApi *api, CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID){
    cout << "ReqQryMMOptionInstrCommRate" << endl;
    return api->ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate, nRequestID);
}
// 请求查询报单手续费
DLL_EXPORT int ReqQryInstrumentOrderCommRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID){
    cout << "ReqQryInstrumentOrderCommRate" << endl;
    return api->ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate, nRequestID);
}
// 请求查询资金账户
DLL_EXPORT int ReqQrySecAgentTradingAccount(CThostFtdcTraderApi *api, CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID){
    cout << "ReqQrySecAgentTradingAccount" << endl;
    return api->ReqQrySecAgentTradingAccount(pQryTradingAccount, nRequestID);
}
// 请求查询二级代理商资金校验模式
DLL_EXPORT int ReqQrySecAgentCheckMode(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID){
    cout << "ReqQrySecAgentCheckMode" << endl;
    return api->ReqQrySecAgentCheckMode(pQrySecAgentCheckMode, nRequestID);
}
// 请求查询二级代理商信息
DLL_EXPORT int ReqQrySecAgentTradeInfo(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID){
    cout << "ReqQrySecAgentTradeInfo" << endl;
    return api->ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo, nRequestID);
}
// 请求查询期权交易成本
DLL_EXPORT int ReqQryOptionInstrTradeCost(CThostFtdcTraderApi *api, CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID){
    cout << "ReqQryOptionInstrTradeCost" << endl;
    return api->ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost, nRequestID);
}
// 请求查询期权合约手续费
DLL_EXPORT int ReqQryOptionInstrCommRate(CThostFtdcTraderApi *api, CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID){
    cout << "ReqQryOptionInstrCommRate" << endl;
    return api->ReqQryOptionInstrCommRate(pQryOptionInstrCommRate, nRequestID);
}
// 请求查询执行宣告
DLL_EXPORT int ReqQryExecOrder(CThostFtdcTraderApi *api, CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID){
    cout << "ReqQryExecOrder" << endl;
    return api->ReqQryExecOrder(pQryExecOrder, nRequestID);
}
// 请求查询询价
DLL_EXPORT int ReqQryForQuote(CThostFtdcTraderApi *api, CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID){
    cout << "ReqQryForQuote" << endl;
    return api->ReqQryForQuote(pQryForQuote, nRequestID);
}
// 请求查询报价
DLL_EXPORT int ReqQryQuote(CThostFtdcTraderApi *api, CThostFtdcQryQuoteField *pQryQuote, int nRequestID){
    cout << "ReqQryQuote" << endl;
    return api->ReqQryQuote(pQryQuote, nRequestID);
}
// 请求查询期权自对冲
DLL_EXPORT int ReqQryOptionSelfClose(CThostFtdcTraderApi *api, CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID){
    cout << "ReqQryOptionSelfClose" << endl;
    return api->ReqQryOptionSelfClose(pQryOptionSelfClose, nRequestID);
}
// 请求查询投资单元
DLL_EXPORT int ReqQryInvestUnit(CThostFtdcTraderApi *api, CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID){
    cout << "ReqQryInvestUnit" << endl;
    return api->ReqQryInvestUnit(pQryInvestUnit, nRequestID);
}
// 请求查询组合合约安全系数
DLL_EXPORT int ReqQryCombInstrumentGuard(CThostFtdcTraderApi *api, CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID){
    cout << "ReqQryCombInstrumentGuard" << endl;
    return api->ReqQryCombInstrumentGuard(pQryCombInstrumentGuard, nRequestID);
}
// 请求查询申请组合
DLL_EXPORT int ReqQryCombAction(CThostFtdcTraderApi *api, CThostFtdcQryCombActionField *pQryCombAction, int nRequestID){
    cout << "ReqQryCombAction" << endl;
    return api->ReqQryCombAction(pQryCombAction, nRequestID);
}
// 请求查询转帐流水
DLL_EXPORT int ReqQryTransferSerial(CThostFtdcTraderApi *api, CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID){
    cout << "ReqQryTransferSerial" << endl;
    return api->ReqQryTransferSerial(pQryTransferSerial, nRequestID);
}
// 请求查询银期签约关系
DLL_EXPORT int ReqQryAccountregister(CThostFtdcTraderApi *api, CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID){
    cout << "ReqQryAccountregister" << endl;
    return api->ReqQryAccountregister(pQryAccountregister, nRequestID);
}
// 请求查询签约银行
DLL_EXPORT int ReqQryContractBank(CThostFtdcTraderApi *api, CThostFtdcQryContractBankField *pQryContractBank, int nRequestID){
    cout << "ReqQryContractBank" << endl;
    return api->ReqQryContractBank(pQryContractBank, nRequestID);
}
// 请求查询预埋单
DLL_EXPORT int ReqQryParkedOrder(CThostFtdcTraderApi *api, CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID){
    cout << "ReqQryParkedOrder" << endl;
    return api->ReqQryParkedOrder(pQryParkedOrder, nRequestID);
}
// 请求查询预埋撤单
DLL_EXPORT int ReqQryParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID){
    cout << "ReqQryParkedOrderAction" << endl;
    return api->ReqQryParkedOrderAction(pQryParkedOrderAction, nRequestID);
}
// 请求查询交易通知
DLL_EXPORT int ReqQryTradingNotice(CThostFtdcTraderApi *api, CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID){
    cout << "ReqQryTradingNotice" << endl;
    return api->ReqQryTradingNotice(pQryTradingNotice, nRequestID);
}
// 请求查询经纪公司交易参数
DLL_EXPORT int ReqQryBrokerTradingParams(CThostFtdcTraderApi *api, CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID){
    cout << "ReqQryBrokerTradingParams" << endl;
    return api->ReqQryBrokerTradingParams(pQryBrokerTradingParams, nRequestID);
}
// 请求查询经纪公司交易算法
DLL_EXPORT int ReqQryBrokerTradingAlgos(CThostFtdcTraderApi *api, CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID){
    cout << "ReqQryBrokerTradingAlgos" << endl;
    return api->ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos, nRequestID);
}
// 请求查询监控中心用户令牌
DLL_EXPORT int ReqQueryCFMMCTradingAccountToken(CThostFtdcTraderApi *api, CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID){
    cout << "ReqQueryCFMMCTradingAccountToken" << endl;
    return api->ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken, nRequestID);
}
// 期货发起银行资金转期货请求
DLL_EXPORT int ReqFromBankToFutureByFuture(CThostFtdcTraderApi *api, CThostFtdcReqTransferField *pReqTransfer, int nRequestID){
    cout << "ReqFromBankToFutureByFuture" << endl;
    return api->ReqFromBankToFutureByFuture(pReqTransfer, nRequestID);
}
// 期货发起期货资金转银行请求
DLL_EXPORT int ReqFromFutureToBankByFuture(CThostFtdcTraderApi *api, CThostFtdcReqTransferField *pReqTransfer, int nRequestID){
    cout << "ReqFromFutureToBankByFuture" << endl;
    return api->ReqFromFutureToBankByFuture(pReqTransfer, nRequestID);
}
// 期货发起查询银行余额请求
DLL_EXPORT int ReqQueryBankAccountMoneyByFuture(CThostFtdcTraderApi *api, CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID){
    cout << "ReqQueryBankAccountMoneyByFuture" << endl;
    return api->ReqQueryBankAccountMoneyByFuture(pReqQueryAccount, nRequestID);
}
// 请求查询分类合约
DLL_EXPORT int ReqQryClassifiedInstrument(CThostFtdcTraderApi *api, CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID){
    cout << "ReqQryClassifiedInstrument" << endl;
    return api->ReqQryClassifiedInstrument(pQryClassifiedInstrument, nRequestID);
}
// 请求组合优惠比例
DLL_EXPORT int ReqQryCombPromotionParam(CThostFtdcTraderApi *api, CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID){
    cout << "ReqQryCombPromotionParam" << endl;
    return api->ReqQryCombPromotionParam(pQryCombPromotionParam, nRequestID);
}
// 投资者风险结算持仓查询
DLL_EXPORT int ReqQryRiskSettleInvstPosition(CThostFtdcTraderApi *api, CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID){
    cout << "ReqQryRiskSettleInvstPosition" << endl;
    return api->ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition, nRequestID);
}
// 风险结算产品查询
DLL_EXPORT int ReqQryRiskSettleProductStatus(CThostFtdcTraderApi *api, CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID){
    cout << "ReqQryRiskSettleProductStatus" << endl;
    return api->ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus, nRequestID);
}


// **** 用 Set 函数将 go 函数指针赋值给 C 函数指针 ****
// //////////////////////////////////////////////////////////////////////
DLL_EXPORT void SetOnFrontConnected(Trade *spi, void *onFunc){
    spi->_OnFrontConnected = onFunc;
}
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
DLL_EXPORT void SetOnFrontDisconnected(Trade *spi, void *onFunc){
    spi->_OnFrontDisconnected = onFunc;
}
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
DLL_EXPORT void SetOnHeartBeatWarning(Trade *spi, void *onFunc){
    spi->_OnHeartBeatWarning = onFunc;
}
// 客户端认证响应
DLL_EXPORT void SetOnRspAuthenticate(Trade *spi, void *onFunc){
    spi->_OnRspAuthenticate = onFunc;
}
// 登录请求响应
DLL_EXPORT void SetOnRspUserLogin(Trade *spi, void *onFunc){
    spi->_OnRspUserLogin = onFunc;
}
// 登出请求响应
DLL_EXPORT void SetOnRspUserLogout(Trade *spi, void *onFunc){
    spi->_OnRspUserLogout = onFunc;
}
// 用户口令更新请求响应
DLL_EXPORT void SetOnRspUserPasswordUpdate(Trade *spi, void *onFunc){
    spi->_OnRspUserPasswordUpdate = onFunc;
}
// 资金账户口令更新请求响应
DLL_EXPORT void SetOnRspTradingAccountPasswordUpdate(Trade *spi, void *onFunc){
    spi->_OnRspTradingAccountPasswordUpdate = onFunc;
}
// 查询用户当前支持的认证模式的回复
DLL_EXPORT void SetOnRspUserAuthMethod(Trade *spi, void *onFunc){
    spi->_OnRspUserAuthMethod = onFunc;
}
// 获取图形验证码请求的回复
DLL_EXPORT void SetOnRspGenUserCaptcha(Trade *spi, void *onFunc){
    spi->_OnRspGenUserCaptcha = onFunc;
}
// 获取短信验证码请求的回复
DLL_EXPORT void SetOnRspGenUserText(Trade *spi, void *onFunc){
    spi->_OnRspGenUserText = onFunc;
}
// 报单录入请求响应
DLL_EXPORT void SetOnRspOrderInsert(Trade *spi, void *onFunc){
    spi->_OnRspOrderInsert = onFunc;
}
// 预埋单录入请求响应
DLL_EXPORT void SetOnRspParkedOrderInsert(Trade *spi, void *onFunc){
    spi->_OnRspParkedOrderInsert = onFunc;
}
// 预埋撤单录入请求响应
DLL_EXPORT void SetOnRspParkedOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspParkedOrderAction = onFunc;
}
// 报单操作请求响应
DLL_EXPORT void SetOnRspOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspOrderAction = onFunc;
}
// 查询最大报单数量响应
DLL_EXPORT void SetOnRspQryMaxOrderVolume(Trade *spi, void *onFunc){
    spi->_OnRspQryMaxOrderVolume = onFunc;
}
// 投资者结算结果确认响应
DLL_EXPORT void SetOnRspSettlementInfoConfirm(Trade *spi, void *onFunc){
    spi->_OnRspSettlementInfoConfirm = onFunc;
}
// 删除预埋单响应
DLL_EXPORT void SetOnRspRemoveParkedOrder(Trade *spi, void *onFunc){
    spi->_OnRspRemoveParkedOrder = onFunc;
}
// 删除预埋撤单响应
DLL_EXPORT void SetOnRspRemoveParkedOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspRemoveParkedOrderAction = onFunc;
}
// 执行宣告录入请求响应
DLL_EXPORT void SetOnRspExecOrderInsert(Trade *spi, void *onFunc){
    spi->_OnRspExecOrderInsert = onFunc;
}
// 执行宣告操作请求响应
DLL_EXPORT void SetOnRspExecOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspExecOrderAction = onFunc;
}
// 询价录入请求响应
DLL_EXPORT void SetOnRspForQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnRspForQuoteInsert = onFunc;
}
// 报价录入请求响应
DLL_EXPORT void SetOnRspQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnRspQuoteInsert = onFunc;
}
// 报价操作请求响应
DLL_EXPORT void SetOnRspQuoteAction(Trade *spi, void *onFunc){
    spi->_OnRspQuoteAction = onFunc;
}
// 批量报单操作请求响应
DLL_EXPORT void SetOnRspBatchOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspBatchOrderAction = onFunc;
}
// 期权自对冲录入请求响应
DLL_EXPORT void SetOnRspOptionSelfCloseInsert(Trade *spi, void *onFunc){
    spi->_OnRspOptionSelfCloseInsert = onFunc;
}
// 期权自对冲操作请求响应
DLL_EXPORT void SetOnRspOptionSelfCloseAction(Trade *spi, void *onFunc){
    spi->_OnRspOptionSelfCloseAction = onFunc;
}
// 申请组合录入请求响应
DLL_EXPORT void SetOnRspCombActionInsert(Trade *spi, void *onFunc){
    spi->_OnRspCombActionInsert = onFunc;
}
// 请求查询报单响应
DLL_EXPORT void SetOnRspQryOrder(Trade *spi, void *onFunc){
    spi->_OnRspQryOrder = onFunc;
}
// 请求查询成交响应
DLL_EXPORT void SetOnRspQryTrade(Trade *spi, void *onFunc){
    spi->_OnRspQryTrade = onFunc;
}
// 请求查询投资者持仓响应
DLL_EXPORT void SetOnRspQryInvestorPosition(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorPosition = onFunc;
}
// 请求查询资金账户响应
DLL_EXPORT void SetOnRspQryTradingAccount(Trade *spi, void *onFunc){
    spi->_OnRspQryTradingAccount = onFunc;
}
// 请求查询投资者响应
DLL_EXPORT void SetOnRspQryInvestor(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestor = onFunc;
}
// 请求查询交易编码响应
DLL_EXPORT void SetOnRspQryTradingCode(Trade *spi, void *onFunc){
    spi->_OnRspQryTradingCode = onFunc;
}
// 请求查询合约保证金率响应
DLL_EXPORT void SetOnRspQryInstrumentMarginRate(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrumentMarginRate = onFunc;
}
// 请求查询合约手续费率响应
DLL_EXPORT void SetOnRspQryInstrumentCommissionRate(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrumentCommissionRate = onFunc;
}
// 请求查询交易所响应
DLL_EXPORT void SetOnRspQryExchange(Trade *spi, void *onFunc){
    spi->_OnRspQryExchange = onFunc;
}
// 请求查询产品响应
DLL_EXPORT void SetOnRspQryProduct(Trade *spi, void *onFunc){
    spi->_OnRspQryProduct = onFunc;
}
// 请求查询合约响应
DLL_EXPORT void SetOnRspQryInstrument(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrument = onFunc;
}
// 请求查询行情响应
DLL_EXPORT void SetOnRspQryDepthMarketData(Trade *spi, void *onFunc){
    spi->_OnRspQryDepthMarketData = onFunc;
}
// 请求查询交易员报盘机响应
DLL_EXPORT void SetOnRspQryTraderOffer(Trade *spi, void *onFunc){
    spi->_OnRspQryTraderOffer = onFunc;
}
// 请求查询投资者结算结果响应
DLL_EXPORT void SetOnRspQrySettlementInfo(Trade *spi, void *onFunc){
    spi->_OnRspQrySettlementInfo = onFunc;
}
// 请求查询转帐银行响应
DLL_EXPORT void SetOnRspQryTransferBank(Trade *spi, void *onFunc){
    spi->_OnRspQryTransferBank = onFunc;
}
// 请求查询投资者持仓明细响应
DLL_EXPORT void SetOnRspQryInvestorPositionDetail(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorPositionDetail = onFunc;
}
// 请求查询客户通知响应
DLL_EXPORT void SetOnRspQryNotice(Trade *spi, void *onFunc){
    spi->_OnRspQryNotice = onFunc;
}
// 请求查询结算信息确认响应
DLL_EXPORT void SetOnRspQrySettlementInfoConfirm(Trade *spi, void *onFunc){
    spi->_OnRspQrySettlementInfoConfirm = onFunc;
}
// 请求查询投资者持仓明细响应
DLL_EXPORT void SetOnRspQryInvestorPositionCombineDetail(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorPositionCombineDetail = onFunc;
}
// 查询保证金监管系统经纪公司资金账户密钥响应
DLL_EXPORT void SetOnRspQryCFMMCTradingAccountKey(Trade *spi, void *onFunc){
    spi->_OnRspQryCFMMCTradingAccountKey = onFunc;
}
// 请求查询仓单折抵信息响应
DLL_EXPORT void SetOnRspQryEWarrantOffset(Trade *spi, void *onFunc){
    spi->_OnRspQryEWarrantOffset = onFunc;
}
// 请求查询投资者品种/跨品种保证金响应
DLL_EXPORT void SetOnRspQryInvestorProductGroupMargin(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorProductGroupMargin = onFunc;
}
// 请求查询交易所保证金率响应
DLL_EXPORT void SetOnRspQryExchangeMarginRate(Trade *spi, void *onFunc){
    spi->_OnRspQryExchangeMarginRate = onFunc;
}
// 请求查询交易所调整保证金率响应
DLL_EXPORT void SetOnRspQryExchangeMarginRateAdjust(Trade *spi, void *onFunc){
    spi->_OnRspQryExchangeMarginRateAdjust = onFunc;
}
// 请求查询汇率响应
DLL_EXPORT void SetOnRspQryExchangeRate(Trade *spi, void *onFunc){
    spi->_OnRspQryExchangeRate = onFunc;
}
// 请求查询二级代理操作员银期权限响应
DLL_EXPORT void SetOnRspQrySecAgentACIDMap(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentACIDMap = onFunc;
}
// 请求查询产品报价汇率
DLL_EXPORT void SetOnRspQryProductExchRate(Trade *spi, void *onFunc){
    spi->_OnRspQryProductExchRate = onFunc;
}
// 请求查询产品组
DLL_EXPORT void SetOnRspQryProductGroup(Trade *spi, void *onFunc){
    spi->_OnRspQryProductGroup = onFunc;
}
// 请求查询做市商合约手续费率响应
DLL_EXPORT void SetOnRspQryMMInstrumentCommissionRate(Trade *spi, void *onFunc){
    spi->_OnRspQryMMInstrumentCommissionRate = onFunc;
}
// 请求查询做市商期权合约手续费响应
DLL_EXPORT void SetOnRspQryMMOptionInstrCommRate(Trade *spi, void *onFunc){
    spi->_OnRspQryMMOptionInstrCommRate = onFunc;
}
// 请求查询报单手续费响应
DLL_EXPORT void SetOnRspQryInstrumentOrderCommRate(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrumentOrderCommRate = onFunc;
}
// 请求查询资金账户响应
DLL_EXPORT void SetOnRspQrySecAgentTradingAccount(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentTradingAccount = onFunc;
}
// 请求查询二级代理商资金校验模式响应
DLL_EXPORT void SetOnRspQrySecAgentCheckMode(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentCheckMode = onFunc;
}
// 请求查询二级代理商信息响应
DLL_EXPORT void SetOnRspQrySecAgentTradeInfo(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentTradeInfo = onFunc;
}
// 请求查询期权交易成本响应
DLL_EXPORT void SetOnRspQryOptionInstrTradeCost(Trade *spi, void *onFunc){
    spi->_OnRspQryOptionInstrTradeCost = onFunc;
}
// 请求查询期权合约手续费响应
DLL_EXPORT void SetOnRspQryOptionInstrCommRate(Trade *spi, void *onFunc){
    spi->_OnRspQryOptionInstrCommRate = onFunc;
}
// 请求查询执行宣告响应
DLL_EXPORT void SetOnRspQryExecOrder(Trade *spi, void *onFunc){
    spi->_OnRspQryExecOrder = onFunc;
}
// 请求查询询价响应
DLL_EXPORT void SetOnRspQryForQuote(Trade *spi, void *onFunc){
    spi->_OnRspQryForQuote = onFunc;
}
// 请求查询报价响应
DLL_EXPORT void SetOnRspQryQuote(Trade *spi, void *onFunc){
    spi->_OnRspQryQuote = onFunc;
}
// 请求查询期权自对冲响应
DLL_EXPORT void SetOnRspQryOptionSelfClose(Trade *spi, void *onFunc){
    spi->_OnRspQryOptionSelfClose = onFunc;
}
// 请求查询投资单元响应
DLL_EXPORT void SetOnRspQryInvestUnit(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestUnit = onFunc;
}
// 请求查询组合合约安全系数响应
DLL_EXPORT void SetOnRspQryCombInstrumentGuard(Trade *spi, void *onFunc){
    spi->_OnRspQryCombInstrumentGuard = onFunc;
}
// 请求查询申请组合响应
DLL_EXPORT void SetOnRspQryCombAction(Trade *spi, void *onFunc){
    spi->_OnRspQryCombAction = onFunc;
}
// 请求查询转帐流水响应
DLL_EXPORT void SetOnRspQryTransferSerial(Trade *spi, void *onFunc){
    spi->_OnRspQryTransferSerial = onFunc;
}
// 请求查询银期签约关系响应
DLL_EXPORT void SetOnRspQryAccountregister(Trade *spi, void *onFunc){
    spi->_OnRspQryAccountregister = onFunc;
}
// 错误应答
DLL_EXPORT void SetOnRspError(Trade *spi, void *onFunc){
    spi->_OnRspError = onFunc;
}
// 报单通知
DLL_EXPORT void SetOnRtnOrder(Trade *spi, void *onFunc){
    spi->_OnRtnOrder = onFunc;
}
// 成交通知
DLL_EXPORT void SetOnRtnTrade(Trade *spi, void *onFunc){
    spi->_OnRtnTrade = onFunc;
}
// 报单录入错误回报
DLL_EXPORT void SetOnErrRtnOrderInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnOrderInsert = onFunc;
}
// 报单操作错误回报
DLL_EXPORT void SetOnErrRtnOrderAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnOrderAction = onFunc;
}
// 合约交易状态通知
DLL_EXPORT void SetOnRtnInstrumentStatus(Trade *spi, void *onFunc){
    spi->_OnRtnInstrumentStatus = onFunc;
}
// 交易所公告通知
DLL_EXPORT void SetOnRtnBulletin(Trade *spi, void *onFunc){
    spi->_OnRtnBulletin = onFunc;
}
// 交易通知
DLL_EXPORT void SetOnRtnTradingNotice(Trade *spi, void *onFunc){
    spi->_OnRtnTradingNotice = onFunc;
}
// 提示条件单校验错误
DLL_EXPORT void SetOnRtnErrorConditionalOrder(Trade *spi, void *onFunc){
    spi->_OnRtnErrorConditionalOrder = onFunc;
}
// 执行宣告通知
DLL_EXPORT void SetOnRtnExecOrder(Trade *spi, void *onFunc){
    spi->_OnRtnExecOrder = onFunc;
}
// 执行宣告录入错误回报
DLL_EXPORT void SetOnErrRtnExecOrderInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnExecOrderInsert = onFunc;
}
// 执行宣告操作错误回报
DLL_EXPORT void SetOnErrRtnExecOrderAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnExecOrderAction = onFunc;
}
// 询价录入错误回报
DLL_EXPORT void SetOnErrRtnForQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnForQuoteInsert = onFunc;
}
// 报价通知
DLL_EXPORT void SetOnRtnQuote(Trade *spi, void *onFunc){
    spi->_OnRtnQuote = onFunc;
}
// 报价录入错误回报
DLL_EXPORT void SetOnErrRtnQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnQuoteInsert = onFunc;
}
// 报价操作错误回报
DLL_EXPORT void SetOnErrRtnQuoteAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnQuoteAction = onFunc;
}
// 询价通知
DLL_EXPORT void SetOnRtnForQuoteRsp(Trade *spi, void *onFunc){
    spi->_OnRtnForQuoteRsp = onFunc;
}
// 保证金监控中心用户令牌
DLL_EXPORT void SetOnRtnCFMMCTradingAccountToken(Trade *spi, void *onFunc){
    spi->_OnRtnCFMMCTradingAccountToken = onFunc;
}
// 批量报单操作错误回报
DLL_EXPORT void SetOnErrRtnBatchOrderAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnBatchOrderAction = onFunc;
}
// 期权自对冲通知
DLL_EXPORT void SetOnRtnOptionSelfClose(Trade *spi, void *onFunc){
    spi->_OnRtnOptionSelfClose = onFunc;
}
// 期权自对冲录入错误回报
DLL_EXPORT void SetOnErrRtnOptionSelfCloseInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnOptionSelfCloseInsert = onFunc;
}
// 期权自对冲操作错误回报
DLL_EXPORT void SetOnErrRtnOptionSelfCloseAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnOptionSelfCloseAction = onFunc;
}
// 申请组合通知
DLL_EXPORT void SetOnRtnCombAction(Trade *spi, void *onFunc){
    spi->_OnRtnCombAction = onFunc;
}
// 申请组合录入错误回报
DLL_EXPORT void SetOnErrRtnCombActionInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnCombActionInsert = onFunc;
}
// 请求查询签约银行响应
DLL_EXPORT void SetOnRspQryContractBank(Trade *spi, void *onFunc){
    spi->_OnRspQryContractBank = onFunc;
}
// 请求查询预埋单响应
DLL_EXPORT void SetOnRspQryParkedOrder(Trade *spi, void *onFunc){
    spi->_OnRspQryParkedOrder = onFunc;
}
// 请求查询预埋撤单响应
DLL_EXPORT void SetOnRspQryParkedOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspQryParkedOrderAction = onFunc;
}
// 请求查询交易通知响应
DLL_EXPORT void SetOnRspQryTradingNotice(Trade *spi, void *onFunc){
    spi->_OnRspQryTradingNotice = onFunc;
}
// 请求查询经纪公司交易参数响应
DLL_EXPORT void SetOnRspQryBrokerTradingParams(Trade *spi, void *onFunc){
    spi->_OnRspQryBrokerTradingParams = onFunc;
}
// 请求查询经纪公司交易算法响应
DLL_EXPORT void SetOnRspQryBrokerTradingAlgos(Trade *spi, void *onFunc){
    spi->_OnRspQryBrokerTradingAlgos = onFunc;
}
// 请求查询监控中心用户令牌
DLL_EXPORT void SetOnRspQueryCFMMCTradingAccountToken(Trade *spi, void *onFunc){
    spi->_OnRspQueryCFMMCTradingAccountToken = onFunc;
}
// 银行发起银行资金转期货通知
DLL_EXPORT void SetOnRtnFromBankToFutureByBank(Trade *spi, void *onFunc){
    spi->_OnRtnFromBankToFutureByBank = onFunc;
}
// 银行发起期货资金转银行通知
DLL_EXPORT void SetOnRtnFromFutureToBankByBank(Trade *spi, void *onFunc){
    spi->_OnRtnFromFutureToBankByBank = onFunc;
}
// 银行发起冲正银行转期货通知
DLL_EXPORT void SetOnRtnRepealFromBankToFutureByBank(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromBankToFutureByBank = onFunc;
}
// 银行发起冲正期货转银行通知
DLL_EXPORT void SetOnRtnRepealFromFutureToBankByBank(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromFutureToBankByBank = onFunc;
}
// 期货发起银行资金转期货通知
DLL_EXPORT void SetOnRtnFromBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnFromBankToFutureByFuture = onFunc;
}
// 期货发起期货资金转银行通知
DLL_EXPORT void SetOnRtnFromFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnFromFutureToBankByFuture = onFunc;
}
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void SetOnRtnRepealFromBankToFutureByFutureManual(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromBankToFutureByFutureManual = onFunc;
}
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void SetOnRtnRepealFromFutureToBankByFutureManual(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromFutureToBankByFutureManual = onFunc;
}
// 期货发起查询银行余额通知
DLL_EXPORT void SetOnRtnQueryBankBalanceByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnQueryBankBalanceByFuture = onFunc;
}
// 期货发起银行资金转期货错误回报
DLL_EXPORT void SetOnErrRtnBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnErrRtnBankToFutureByFuture = onFunc;
}
// 期货发起期货资金转银行错误回报
DLL_EXPORT void SetOnErrRtnFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnErrRtnFutureToBankByFuture = onFunc;
}
// 系统运行时期货端手工发起冲正银行转期货错误回报
DLL_EXPORT void SetOnErrRtnRepealBankToFutureByFutureManual(Trade *spi, void *onFunc){
    spi->_OnErrRtnRepealBankToFutureByFutureManual = onFunc;
}
// 系统运行时期货端手工发起冲正期货转银行错误回报
DLL_EXPORT void SetOnErrRtnRepealFutureToBankByFutureManual(Trade *spi, void *onFunc){
    spi->_OnErrRtnRepealFutureToBankByFutureManual = onFunc;
}
// 期货发起查询银行余额错误回报
DLL_EXPORT void SetOnErrRtnQueryBankBalanceByFuture(Trade *spi, void *onFunc){
    spi->_OnErrRtnQueryBankBalanceByFuture = onFunc;
}
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void SetOnRtnRepealFromBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromBankToFutureByFuture = onFunc;
}
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void SetOnRtnRepealFromFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromFutureToBankByFuture = onFunc;
}
// 期货发起银行资金转期货应答
DLL_EXPORT void SetOnRspFromBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnRspFromBankToFutureByFuture = onFunc;
}
// 期货发起期货资金转银行应答
DLL_EXPORT void SetOnRspFromFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnRspFromFutureToBankByFuture = onFunc;
}
// 期货发起查询银行余额应答
DLL_EXPORT void SetOnRspQueryBankAccountMoneyByFuture(Trade *spi, void *onFunc){
    spi->_OnRspQueryBankAccountMoneyByFuture = onFunc;
}
// 银行发起银期开户通知
DLL_EXPORT void SetOnRtnOpenAccountByBank(Trade *spi, void *onFunc){
    spi->_OnRtnOpenAccountByBank = onFunc;
}
// 银行发起银期销户通知
DLL_EXPORT void SetOnRtnCancelAccountByBank(Trade *spi, void *onFunc){
    spi->_OnRtnCancelAccountByBank = onFunc;
}
// 银行发起变更银行账号通知
DLL_EXPORT void SetOnRtnChangeAccountByBank(Trade *spi, void *onFunc){
    spi->_OnRtnChangeAccountByBank = onFunc;
}
// 请求查询分类合约响应
DLL_EXPORT void SetOnRspQryClassifiedInstrument(Trade *spi, void *onFunc){
    spi->_OnRspQryClassifiedInstrument = onFunc;
}
// 请求组合优惠比例响应
DLL_EXPORT void SetOnRspQryCombPromotionParam(Trade *spi, void *onFunc){
    spi->_OnRspQryCombPromotionParam = onFunc;
}
// 投资者风险结算持仓查询响应
DLL_EXPORT void SetOnRspQryRiskSettleInvstPosition(Trade *spi, void *onFunc){
    spi->_OnRspQryRiskSettleInvstPosition = onFunc;
}
// 风险结算产品查询响应
DLL_EXPORT void SetOnRspQryRiskSettleProductStatus(Trade *spi, void *onFunc){
    spi->_OnRspQryRiskSettleProductStatus = onFunc;
}
