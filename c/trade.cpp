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
DLL_EXPORT void tRelease(CThostFtdcTraderApi *api){
    cout << "Release" << endl;
    return api->Release();
}
// 初始化
DLL_EXPORT void tInit(CThostFtdcTraderApi *api){
    cout << "Init" << endl;
    return api->Init();
}
// 等待接口线程结束运行
DLL_EXPORT int tJoin(CThostFtdcTraderApi *api){
    cout << "Join" << endl;
    return api->Join();
}
// 注册前置机网络地址
DLL_EXPORT void tRegisterFront(CThostFtdcTraderApi *api, char *pszFrontAddress){
    cout << "RegisterFront" << endl;
    return api->RegisterFront(pszFrontAddress);
}
// @remark RegisterNameServer优先于RegisterFront
DLL_EXPORT void tRegisterNameServer(CThostFtdcTraderApi *api, char *pszNsAddress){
    cout << "RegisterNameServer" << endl;
    return api->RegisterNameServer(pszNsAddress);
}
// 注册名字服务器用户信息
DLL_EXPORT void tRegisterFensUserInfo(CThostFtdcTraderApi *api, CThostFtdcFensUserInfoField *pFensUserInfo){
    cout << "RegisterFensUserInfo" << endl;
    return api->RegisterFensUserInfo(pFensUserInfo);
}
// 注册回调接口
DLL_EXPORT void tRegisterSpi(CThostFtdcTraderApi *api, CThostFtdcTraderSpi *pSpi){
    cout << "RegisterSpi" << endl;
    return api->RegisterSpi(pSpi);
}
// 订阅私有流。
DLL_EXPORT void tSubscribePrivateTopic(CThostFtdcTraderApi *api, THOST_TE_RESUME_TYPE nResumeType){
    cout << "SubscribePrivateTopic" << endl;
    return api->SubscribePrivateTopic(nResumeType);
}
// 订阅公共流。
DLL_EXPORT void tSubscribePublicTopic(CThostFtdcTraderApi *api, THOST_TE_RESUME_TYPE nResumeType){
    cout << "SubscribePublicTopic" << endl;
    return api->SubscribePublicTopic(nResumeType);
}
// 客户端认证请求
DLL_EXPORT int tReqAuthenticate(CThostFtdcTraderApi *api, CThostFtdcReqAuthenticateField *pReqAuthenticateField, int nRequestID){
    cout << "ReqAuthenticate" << endl;
    return api->ReqAuthenticate(pReqAuthenticateField, nRequestID);
}
// 注册用户终端信息，用于中继服务器多连接模式
DLL_EXPORT int tRegisterUserSystemInfo(CThostFtdcTraderApi *api, CThostFtdcUserSystemInfoField *pUserSystemInfo){
    cout << "RegisterUserSystemInfo" << endl;
    return api->RegisterUserSystemInfo(pUserSystemInfo);
}
// 上报用户终端信息，用于中继服务器操作员登录模式
DLL_EXPORT int tSubmitUserSystemInfo(CThostFtdcTraderApi *api, CThostFtdcUserSystemInfoField *pUserSystemInfo){
    cout << "SubmitUserSystemInfo" << endl;
    return api->SubmitUserSystemInfo(pUserSystemInfo);
}
// 用户登录请求
DLL_EXPORT int tReqUserLogin(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID){
    cout << "ReqUserLogin" << endl;
    return api->ReqUserLogin(pReqUserLoginField, nRequestID);
}
// 登出请求
DLL_EXPORT int tReqUserLogout(CThostFtdcTraderApi *api, CThostFtdcUserLogoutField *pUserLogout, int nRequestID){
    cout << "ReqUserLogout" << endl;
    return api->ReqUserLogout(pUserLogout, nRequestID);
}
// 用户口令更新请求
DLL_EXPORT int tReqUserPasswordUpdate(CThostFtdcTraderApi *api, CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, int nRequestID){
    cout << "ReqUserPasswordUpdate" << endl;
    return api->ReqUserPasswordUpdate(pUserPasswordUpdate, nRequestID);
}
// 资金账户口令更新请求
DLL_EXPORT int tReqTradingAccountPasswordUpdate(CThostFtdcTraderApi *api, CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, int nRequestID){
    cout << "ReqTradingAccountPasswordUpdate" << endl;
    return api->ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate, nRequestID);
}
// 查询用户当前支持的认证模式
DLL_EXPORT int tReqUserAuthMethod(CThostFtdcTraderApi *api, CThostFtdcReqUserAuthMethodField *pReqUserAuthMethod, int nRequestID){
    cout << "ReqUserAuthMethod" << endl;
    return api->ReqUserAuthMethod(pReqUserAuthMethod, nRequestID);
}
// 用户发出获取图形验证码请求
DLL_EXPORT int tReqGenUserCaptcha(CThostFtdcTraderApi *api, CThostFtdcReqGenUserCaptchaField *pReqGenUserCaptcha, int nRequestID){
    cout << "ReqGenUserCaptcha" << endl;
    return api->ReqGenUserCaptcha(pReqGenUserCaptcha, nRequestID);
}
// 用户发出获取短信验证码请求
DLL_EXPORT int tReqGenUserText(CThostFtdcTraderApi *api, CThostFtdcReqGenUserTextField *pReqGenUserText, int nRequestID){
    cout << "ReqGenUserText" << endl;
    return api->ReqGenUserText(pReqGenUserText, nRequestID);
}
// 用户发出带有图片验证码的登陆请求
DLL_EXPORT int tReqUserLoginWithCaptcha(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithCaptchaField *pReqUserLoginWithCaptcha, int nRequestID){
    cout << "ReqUserLoginWithCaptcha" << endl;
    return api->ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha, nRequestID);
}
// 用户发出带有短信验证码的登陆请求
DLL_EXPORT int tReqUserLoginWithText(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithTextField *pReqUserLoginWithText, int nRequestID){
    cout << "ReqUserLoginWithText" << endl;
    return api->ReqUserLoginWithText(pReqUserLoginWithText, nRequestID);
}
// 用户发出带有动态口令的登陆请求
DLL_EXPORT int tReqUserLoginWithOTP(CThostFtdcTraderApi *api, CThostFtdcReqUserLoginWithOTPField *pReqUserLoginWithOTP, int nRequestID){
    cout << "ReqUserLoginWithOTP" << endl;
    return api->ReqUserLoginWithOTP(pReqUserLoginWithOTP, nRequestID);
}
// 报单录入请求
DLL_EXPORT int tReqOrderInsert(CThostFtdcTraderApi *api, CThostFtdcInputOrderField *pInputOrder, int nRequestID){
    cout << "ReqOrderInsert" << endl;
    return api->ReqOrderInsert(pInputOrder, nRequestID);
}
// 预埋单录入请求
DLL_EXPORT int tReqParkedOrderInsert(CThostFtdcTraderApi *api, CThostFtdcParkedOrderField *pParkedOrder, int nRequestID){
    cout << "ReqParkedOrderInsert" << endl;
    return api->ReqParkedOrderInsert(pParkedOrder, nRequestID);
}
// 预埋撤单录入请求
DLL_EXPORT int tReqParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcParkedOrderActionField *pParkedOrderAction, int nRequestID){
    cout << "ReqParkedOrderAction" << endl;
    return api->ReqParkedOrderAction(pParkedOrderAction, nRequestID);
}
// 报单操作请求
DLL_EXPORT int tReqOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputOrderActionField *pInputOrderAction, int nRequestID){
    cout << "ReqOrderAction" << endl;
    return api->ReqOrderAction(pInputOrderAction, nRequestID);
}
// 查询最大报单数量请求
DLL_EXPORT int tReqQryMaxOrderVolume(CThostFtdcTraderApi *api, CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, int nRequestID){
    cout << "ReqQryMaxOrderVolume" << endl;
    return api->ReqQryMaxOrderVolume(pQryMaxOrderVolume, nRequestID);
}
// 投资者结算结果确认
DLL_EXPORT int tReqSettlementInfoConfirm(CThostFtdcTraderApi *api, CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, int nRequestID){
    cout << "ReqSettlementInfoConfirm" << endl;
    return api->ReqSettlementInfoConfirm(pSettlementInfoConfirm, nRequestID);
}
// 请求删除预埋单
DLL_EXPORT int tReqRemoveParkedOrder(CThostFtdcTraderApi *api, CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, int nRequestID){
    cout << "ReqRemoveParkedOrder" << endl;
    return api->ReqRemoveParkedOrder(pRemoveParkedOrder, nRequestID);
}
// 请求删除预埋撤单
DLL_EXPORT int tReqRemoveParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, int nRequestID){
    cout << "ReqRemoveParkedOrderAction" << endl;
    return api->ReqRemoveParkedOrderAction(pRemoveParkedOrderAction, nRequestID);
}
// 执行宣告录入请求
DLL_EXPORT int tReqExecOrderInsert(CThostFtdcTraderApi *api, CThostFtdcInputExecOrderField *pInputExecOrder, int nRequestID){
    cout << "ReqExecOrderInsert" << endl;
    return api->ReqExecOrderInsert(pInputExecOrder, nRequestID);
}
// 执行宣告操作请求
DLL_EXPORT int tReqExecOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputExecOrderActionField *pInputExecOrderAction, int nRequestID){
    cout << "ReqExecOrderAction" << endl;
    return api->ReqExecOrderAction(pInputExecOrderAction, nRequestID);
}
// 询价录入请求
DLL_EXPORT int tReqForQuoteInsert(CThostFtdcTraderApi *api, CThostFtdcInputForQuoteField *pInputForQuote, int nRequestID){
    cout << "ReqForQuoteInsert" << endl;
    return api->ReqForQuoteInsert(pInputForQuote, nRequestID);
}
// 报价录入请求
DLL_EXPORT int tReqQuoteInsert(CThostFtdcTraderApi *api, CThostFtdcInputQuoteField *pInputQuote, int nRequestID){
    cout << "ReqQuoteInsert" << endl;
    return api->ReqQuoteInsert(pInputQuote, nRequestID);
}
// 报价操作请求
DLL_EXPORT int tReqQuoteAction(CThostFtdcTraderApi *api, CThostFtdcInputQuoteActionField *pInputQuoteAction, int nRequestID){
    cout << "ReqQuoteAction" << endl;
    return api->ReqQuoteAction(pInputQuoteAction, nRequestID);
}
// 批量报单操作请求
DLL_EXPORT int tReqBatchOrderAction(CThostFtdcTraderApi *api, CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, int nRequestID){
    cout << "ReqBatchOrderAction" << endl;
    return api->ReqBatchOrderAction(pInputBatchOrderAction, nRequestID);
}
// 期权自对冲录入请求
DLL_EXPORT int tReqOptionSelfCloseInsert(CThostFtdcTraderApi *api, CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, int nRequestID){
    cout << "ReqOptionSelfCloseInsert" << endl;
    return api->ReqOptionSelfCloseInsert(pInputOptionSelfClose, nRequestID);
}
// 期权自对冲操作请求
DLL_EXPORT int tReqOptionSelfCloseAction(CThostFtdcTraderApi *api, CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, int nRequestID){
    cout << "ReqOptionSelfCloseAction" << endl;
    return api->ReqOptionSelfCloseAction(pInputOptionSelfCloseAction, nRequestID);
}
// 申请组合录入请求
DLL_EXPORT int tReqCombActionInsert(CThostFtdcTraderApi *api, CThostFtdcInputCombActionField *pInputCombAction, int nRequestID){
    cout << "ReqCombActionInsert" << endl;
    return api->ReqCombActionInsert(pInputCombAction, nRequestID);
}
// 请求查询报单
DLL_EXPORT int tReqQryOrder(CThostFtdcTraderApi *api, CThostFtdcQryOrderField *pQryOrder, int nRequestID){
    cout << "ReqQryOrder" << endl;
    return api->ReqQryOrder(pQryOrder, nRequestID);
}
// 请求查询成交
DLL_EXPORT int tReqQryTrade(CThostFtdcTraderApi *api, CThostFtdcQryTradeField *pQryTrade, int nRequestID){
    cout << "ReqQryTrade" << endl;
    return api->ReqQryTrade(pQryTrade, nRequestID);
}
// 请求查询投资者持仓
DLL_EXPORT int tReqQryInvestorPosition(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionField *pQryInvestorPosition, int nRequestID){
    cout << "ReqQryInvestorPosition" << endl;
    return api->ReqQryInvestorPosition(pQryInvestorPosition, nRequestID);
}
// 请求查询资金账户
DLL_EXPORT int tReqQryTradingAccount(CThostFtdcTraderApi *api, CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID){
    cout << "ReqQryTradingAccount" << endl;
    return api->ReqQryTradingAccount(pQryTradingAccount, nRequestID);
}
// 请求查询投资者
DLL_EXPORT int tReqQryInvestor(CThostFtdcTraderApi *api, CThostFtdcQryInvestorField *pQryInvestor, int nRequestID){
    cout << "ReqQryInvestor" << endl;
    return api->ReqQryInvestor(pQryInvestor, nRequestID);
}
// 请求查询交易编码
DLL_EXPORT int tReqQryTradingCode(CThostFtdcTraderApi *api, CThostFtdcQryTradingCodeField *pQryTradingCode, int nRequestID){
    cout << "ReqQryTradingCode" << endl;
    return api->ReqQryTradingCode(pQryTradingCode, nRequestID);
}
// 请求查询合约保证金率
DLL_EXPORT int tReqQryInstrumentMarginRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentMarginRateField *pQryInstrumentMarginRate, int nRequestID){
    cout << "ReqQryInstrumentMarginRate" << endl;
    return api->ReqQryInstrumentMarginRate(pQryInstrumentMarginRate, nRequestID);
}
// 请求查询合约手续费率
DLL_EXPORT int tReqQryInstrumentCommissionRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentCommissionRateField *pQryInstrumentCommissionRate, int nRequestID){
    cout << "ReqQryInstrumentCommissionRate" << endl;
    return api->ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate, nRequestID);
}
// 请求查询交易所
DLL_EXPORT int tReqQryExchange(CThostFtdcTraderApi *api, CThostFtdcQryExchangeField *pQryExchange, int nRequestID){
    cout << "ReqQryExchange" << endl;
    return api->ReqQryExchange(pQryExchange, nRequestID);
}
// 请求查询产品
DLL_EXPORT int tReqQryProduct(CThostFtdcTraderApi *api, CThostFtdcQryProductField *pQryProduct, int nRequestID){
    cout << "ReqQryProduct" << endl;
    return api->ReqQryProduct(pQryProduct, nRequestID);
}
// 请求查询合约
DLL_EXPORT int tReqQryInstrument(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentField *pQryInstrument, int nRequestID){
    cout << "ReqQryInstrument" << endl;
    return api->ReqQryInstrument(pQryInstrument, nRequestID);
}
// 请求查询行情
DLL_EXPORT int tReqQryDepthMarketData(CThostFtdcTraderApi *api, CThostFtdcQryDepthMarketDataField *pQryDepthMarketData, int nRequestID){
    cout << "ReqQryDepthMarketData" << endl;
    return api->ReqQryDepthMarketData(pQryDepthMarketData, nRequestID);
}
// 请求查询交易员报盘机
DLL_EXPORT int tReqQryTraderOffer(CThostFtdcTraderApi *api, CThostFtdcQryTraderOfferField *pQryTraderOffer, int nRequestID){
    cout << "ReqQryTraderOffer" << endl;
    return api->ReqQryTraderOffer(pQryTraderOffer, nRequestID);
}
// 请求查询投资者结算结果
DLL_EXPORT int tReqQrySettlementInfo(CThostFtdcTraderApi *api, CThostFtdcQrySettlementInfoField *pQrySettlementInfo, int nRequestID){
    cout << "ReqQrySettlementInfo" << endl;
    return api->ReqQrySettlementInfo(pQrySettlementInfo, nRequestID);
}
// 请求查询转帐银行
DLL_EXPORT int tReqQryTransferBank(CThostFtdcTraderApi *api, CThostFtdcQryTransferBankField *pQryTransferBank, int nRequestID){
    cout << "ReqQryTransferBank" << endl;
    return api->ReqQryTransferBank(pQryTransferBank, nRequestID);
}
// 请求查询投资者持仓明细
DLL_EXPORT int tReqQryInvestorPositionDetail(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionDetailField *pQryInvestorPositionDetail, int nRequestID){
    cout << "ReqQryInvestorPositionDetail" << endl;
    return api->ReqQryInvestorPositionDetail(pQryInvestorPositionDetail, nRequestID);
}
// 请求查询客户通知
DLL_EXPORT int tReqQryNotice(CThostFtdcTraderApi *api, CThostFtdcQryNoticeField *pQryNotice, int nRequestID){
    cout << "ReqQryNotice" << endl;
    return api->ReqQryNotice(pQryNotice, nRequestID);
}
// 请求查询结算信息确认
DLL_EXPORT int tReqQrySettlementInfoConfirm(CThostFtdcTraderApi *api, CThostFtdcQrySettlementInfoConfirmField *pQrySettlementInfoConfirm, int nRequestID){
    cout << "ReqQrySettlementInfoConfirm" << endl;
    return api->ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm, nRequestID);
}
// 请求查询投资者持仓明细
DLL_EXPORT int tReqQryInvestorPositionCombineDetail(CThostFtdcTraderApi *api, CThostFtdcQryInvestorPositionCombineDetailField *pQryInvestorPositionCombineDetail, int nRequestID){
    cout << "ReqQryInvestorPositionCombineDetail" << endl;
    return api->ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail, nRequestID);
}
// 请求查询保证金监管系统经纪公司资金账户密钥
DLL_EXPORT int tReqQryCFMMCTradingAccountKey(CThostFtdcTraderApi *api, CThostFtdcQryCFMMCTradingAccountKeyField *pQryCFMMCTradingAccountKey, int nRequestID){
    cout << "ReqQryCFMMCTradingAccountKey" << endl;
    return api->ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey, nRequestID);
}
// 请求查询仓单折抵信息
DLL_EXPORT int tReqQryEWarrantOffset(CThostFtdcTraderApi *api, CThostFtdcQryEWarrantOffsetField *pQryEWarrantOffset, int nRequestID){
    cout << "ReqQryEWarrantOffset" << endl;
    return api->ReqQryEWarrantOffset(pQryEWarrantOffset, nRequestID);
}
// 请求查询投资者品种/跨品种保证金
DLL_EXPORT int tReqQryInvestorProductGroupMargin(CThostFtdcTraderApi *api, CThostFtdcQryInvestorProductGroupMarginField *pQryInvestorProductGroupMargin, int nRequestID){
    cout << "ReqQryInvestorProductGroupMargin" << endl;
    return api->ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin, nRequestID);
}
// 请求查询交易所保证金率
DLL_EXPORT int tReqQryExchangeMarginRate(CThostFtdcTraderApi *api, CThostFtdcQryExchangeMarginRateField *pQryExchangeMarginRate, int nRequestID){
    cout << "ReqQryExchangeMarginRate" << endl;
    return api->ReqQryExchangeMarginRate(pQryExchangeMarginRate, nRequestID);
}
// 请求查询交易所调整保证金率
DLL_EXPORT int tReqQryExchangeMarginRateAdjust(CThostFtdcTraderApi *api, CThostFtdcQryExchangeMarginRateAdjustField *pQryExchangeMarginRateAdjust, int nRequestID){
    cout << "ReqQryExchangeMarginRateAdjust" << endl;
    return api->ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust, nRequestID);
}
// 请求查询汇率
DLL_EXPORT int tReqQryExchangeRate(CThostFtdcTraderApi *api, CThostFtdcQryExchangeRateField *pQryExchangeRate, int nRequestID){
    cout << "ReqQryExchangeRate" << endl;
    return api->ReqQryExchangeRate(pQryExchangeRate, nRequestID);
}
// 请求查询二级代理操作员银期权限
DLL_EXPORT int tReqQrySecAgentACIDMap(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentACIDMapField *pQrySecAgentACIDMap, int nRequestID){
    cout << "ReqQrySecAgentACIDMap" << endl;
    return api->ReqQrySecAgentACIDMap(pQrySecAgentACIDMap, nRequestID);
}
// 请求查询产品报价汇率
DLL_EXPORT int tReqQryProductExchRate(CThostFtdcTraderApi *api, CThostFtdcQryProductExchRateField *pQryProductExchRate, int nRequestID){
    cout << "ReqQryProductExchRate" << endl;
    return api->ReqQryProductExchRate(pQryProductExchRate, nRequestID);
}
// 请求查询产品组
DLL_EXPORT int tReqQryProductGroup(CThostFtdcTraderApi *api, CThostFtdcQryProductGroupField *pQryProductGroup, int nRequestID){
    cout << "ReqQryProductGroup" << endl;
    return api->ReqQryProductGroup(pQryProductGroup, nRequestID);
}
// 请求查询做市商合约手续费率
DLL_EXPORT int tReqQryMMInstrumentCommissionRate(CThostFtdcTraderApi *api, CThostFtdcQryMMInstrumentCommissionRateField *pQryMMInstrumentCommissionRate, int nRequestID){
    cout << "ReqQryMMInstrumentCommissionRate" << endl;
    return api->ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate, nRequestID);
}
// 请求查询做市商期权合约手续费
DLL_EXPORT int tReqQryMMOptionInstrCommRate(CThostFtdcTraderApi *api, CThostFtdcQryMMOptionInstrCommRateField *pQryMMOptionInstrCommRate, int nRequestID){
    cout << "ReqQryMMOptionInstrCommRate" << endl;
    return api->ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate, nRequestID);
}
// 请求查询报单手续费
DLL_EXPORT int tReqQryInstrumentOrderCommRate(CThostFtdcTraderApi *api, CThostFtdcQryInstrumentOrderCommRateField *pQryInstrumentOrderCommRate, int nRequestID){
    cout << "ReqQryInstrumentOrderCommRate" << endl;
    return api->ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate, nRequestID);
}
// 请求查询资金账户
DLL_EXPORT int tReqQrySecAgentTradingAccount(CThostFtdcTraderApi *api, CThostFtdcQryTradingAccountField *pQryTradingAccount, int nRequestID){
    cout << "ReqQrySecAgentTradingAccount" << endl;
    return api->ReqQrySecAgentTradingAccount(pQryTradingAccount, nRequestID);
}
// 请求查询二级代理商资金校验模式
DLL_EXPORT int tReqQrySecAgentCheckMode(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentCheckModeField *pQrySecAgentCheckMode, int nRequestID){
    cout << "ReqQrySecAgentCheckMode" << endl;
    return api->ReqQrySecAgentCheckMode(pQrySecAgentCheckMode, nRequestID);
}
// 请求查询二级代理商信息
DLL_EXPORT int tReqQrySecAgentTradeInfo(CThostFtdcTraderApi *api, CThostFtdcQrySecAgentTradeInfoField *pQrySecAgentTradeInfo, int nRequestID){
    cout << "ReqQrySecAgentTradeInfo" << endl;
    return api->ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo, nRequestID);
}
// 请求查询期权交易成本
DLL_EXPORT int tReqQryOptionInstrTradeCost(CThostFtdcTraderApi *api, CThostFtdcQryOptionInstrTradeCostField *pQryOptionInstrTradeCost, int nRequestID){
    cout << "ReqQryOptionInstrTradeCost" << endl;
    return api->ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost, nRequestID);
}
// 请求查询期权合约手续费
DLL_EXPORT int tReqQryOptionInstrCommRate(CThostFtdcTraderApi *api, CThostFtdcQryOptionInstrCommRateField *pQryOptionInstrCommRate, int nRequestID){
    cout << "ReqQryOptionInstrCommRate" << endl;
    return api->ReqQryOptionInstrCommRate(pQryOptionInstrCommRate, nRequestID);
}
// 请求查询执行宣告
DLL_EXPORT int tReqQryExecOrder(CThostFtdcTraderApi *api, CThostFtdcQryExecOrderField *pQryExecOrder, int nRequestID){
    cout << "ReqQryExecOrder" << endl;
    return api->ReqQryExecOrder(pQryExecOrder, nRequestID);
}
// 请求查询询价
DLL_EXPORT int tReqQryForQuote(CThostFtdcTraderApi *api, CThostFtdcQryForQuoteField *pQryForQuote, int nRequestID){
    cout << "ReqQryForQuote" << endl;
    return api->ReqQryForQuote(pQryForQuote, nRequestID);
}
// 请求查询报价
DLL_EXPORT int tReqQryQuote(CThostFtdcTraderApi *api, CThostFtdcQryQuoteField *pQryQuote, int nRequestID){
    cout << "ReqQryQuote" << endl;
    return api->ReqQryQuote(pQryQuote, nRequestID);
}
// 请求查询期权自对冲
DLL_EXPORT int tReqQryOptionSelfClose(CThostFtdcTraderApi *api, CThostFtdcQryOptionSelfCloseField *pQryOptionSelfClose, int nRequestID){
    cout << "ReqQryOptionSelfClose" << endl;
    return api->ReqQryOptionSelfClose(pQryOptionSelfClose, nRequestID);
}
// 请求查询投资单元
DLL_EXPORT int tReqQryInvestUnit(CThostFtdcTraderApi *api, CThostFtdcQryInvestUnitField *pQryInvestUnit, int nRequestID){
    cout << "ReqQryInvestUnit" << endl;
    return api->ReqQryInvestUnit(pQryInvestUnit, nRequestID);
}
// 请求查询组合合约安全系数
DLL_EXPORT int tReqQryCombInstrumentGuard(CThostFtdcTraderApi *api, CThostFtdcQryCombInstrumentGuardField *pQryCombInstrumentGuard, int nRequestID){
    cout << "ReqQryCombInstrumentGuard" << endl;
    return api->ReqQryCombInstrumentGuard(pQryCombInstrumentGuard, nRequestID);
}
// 请求查询申请组合
DLL_EXPORT int tReqQryCombAction(CThostFtdcTraderApi *api, CThostFtdcQryCombActionField *pQryCombAction, int nRequestID){
    cout << "ReqQryCombAction" << endl;
    return api->ReqQryCombAction(pQryCombAction, nRequestID);
}
// 请求查询转帐流水
DLL_EXPORT int tReqQryTransferSerial(CThostFtdcTraderApi *api, CThostFtdcQryTransferSerialField *pQryTransferSerial, int nRequestID){
    cout << "ReqQryTransferSerial" << endl;
    return api->ReqQryTransferSerial(pQryTransferSerial, nRequestID);
}
// 请求查询银期签约关系
DLL_EXPORT int tReqQryAccountregister(CThostFtdcTraderApi *api, CThostFtdcQryAccountregisterField *pQryAccountregister, int nRequestID){
    cout << "ReqQryAccountregister" << endl;
    return api->ReqQryAccountregister(pQryAccountregister, nRequestID);
}
// 请求查询签约银行
DLL_EXPORT int tReqQryContractBank(CThostFtdcTraderApi *api, CThostFtdcQryContractBankField *pQryContractBank, int nRequestID){
    cout << "ReqQryContractBank" << endl;
    return api->ReqQryContractBank(pQryContractBank, nRequestID);
}
// 请求查询预埋单
DLL_EXPORT int tReqQryParkedOrder(CThostFtdcTraderApi *api, CThostFtdcQryParkedOrderField *pQryParkedOrder, int nRequestID){
    cout << "ReqQryParkedOrder" << endl;
    return api->ReqQryParkedOrder(pQryParkedOrder, nRequestID);
}
// 请求查询预埋撤单
DLL_EXPORT int tReqQryParkedOrderAction(CThostFtdcTraderApi *api, CThostFtdcQryParkedOrderActionField *pQryParkedOrderAction, int nRequestID){
    cout << "ReqQryParkedOrderAction" << endl;
    return api->ReqQryParkedOrderAction(pQryParkedOrderAction, nRequestID);
}
// 请求查询交易通知
DLL_EXPORT int tReqQryTradingNotice(CThostFtdcTraderApi *api, CThostFtdcQryTradingNoticeField *pQryTradingNotice, int nRequestID){
    cout << "ReqQryTradingNotice" << endl;
    return api->ReqQryTradingNotice(pQryTradingNotice, nRequestID);
}
// 请求查询经纪公司交易参数
DLL_EXPORT int tReqQryBrokerTradingParams(CThostFtdcTraderApi *api, CThostFtdcQryBrokerTradingParamsField *pQryBrokerTradingParams, int nRequestID){
    cout << "ReqQryBrokerTradingParams" << endl;
    return api->ReqQryBrokerTradingParams(pQryBrokerTradingParams, nRequestID);
}
// 请求查询经纪公司交易算法
DLL_EXPORT int tReqQryBrokerTradingAlgos(CThostFtdcTraderApi *api, CThostFtdcQryBrokerTradingAlgosField *pQryBrokerTradingAlgos, int nRequestID){
    cout << "ReqQryBrokerTradingAlgos" << endl;
    return api->ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos, nRequestID);
}
// 请求查询监控中心用户令牌
DLL_EXPORT int tReqQueryCFMMCTradingAccountToken(CThostFtdcTraderApi *api, CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, int nRequestID){
    cout << "ReqQueryCFMMCTradingAccountToken" << endl;
    return api->ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken, nRequestID);
}
// 期货发起银行资金转期货请求
DLL_EXPORT int tReqFromBankToFutureByFuture(CThostFtdcTraderApi *api, CThostFtdcReqTransferField *pReqTransfer, int nRequestID){
    cout << "ReqFromBankToFutureByFuture" << endl;
    return api->ReqFromBankToFutureByFuture(pReqTransfer, nRequestID);
}
// 期货发起期货资金转银行请求
DLL_EXPORT int tReqFromFutureToBankByFuture(CThostFtdcTraderApi *api, CThostFtdcReqTransferField *pReqTransfer, int nRequestID){
    cout << "ReqFromFutureToBankByFuture" << endl;
    return api->ReqFromFutureToBankByFuture(pReqTransfer, nRequestID);
}
// 期货发起查询银行余额请求
DLL_EXPORT int tReqQueryBankAccountMoneyByFuture(CThostFtdcTraderApi *api, CThostFtdcReqQueryAccountField *pReqQueryAccount, int nRequestID){
    cout << "ReqQueryBankAccountMoneyByFuture" << endl;
    return api->ReqQueryBankAccountMoneyByFuture(pReqQueryAccount, nRequestID);
}
// 请求查询分类合约
DLL_EXPORT int tReqQryClassifiedInstrument(CThostFtdcTraderApi *api, CThostFtdcQryClassifiedInstrumentField *pQryClassifiedInstrument, int nRequestID){
    cout << "ReqQryClassifiedInstrument" << endl;
    return api->ReqQryClassifiedInstrument(pQryClassifiedInstrument, nRequestID);
}
// 请求组合优惠比例
DLL_EXPORT int tReqQryCombPromotionParam(CThostFtdcTraderApi *api, CThostFtdcQryCombPromotionParamField *pQryCombPromotionParam, int nRequestID){
    cout << "ReqQryCombPromotionParam" << endl;
    return api->ReqQryCombPromotionParam(pQryCombPromotionParam, nRequestID);
}
// 投资者风险结算持仓查询
DLL_EXPORT int tReqQryRiskSettleInvstPosition(CThostFtdcTraderApi *api, CThostFtdcQryRiskSettleInvstPositionField *pQryRiskSettleInvstPosition, int nRequestID){
    cout << "ReqQryRiskSettleInvstPosition" << endl;
    return api->ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition, nRequestID);
}
// 风险结算产品查询
DLL_EXPORT int tReqQryRiskSettleProductStatus(CThostFtdcTraderApi *api, CThostFtdcQryRiskSettleProductStatusField *pQryRiskSettleProductStatus, int nRequestID){
    cout << "ReqQryRiskSettleProductStatus" << endl;
    return api->ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus, nRequestID);
}


// **** 用 Set 函数将 go 函数指针赋值给 C 函数指针 ****
// //////////////////////////////////////////////////////////////////////
DLL_EXPORT void tSetOnFrontConnected(Trade *spi, void *onFunc){
    spi->_OnFrontConnected = onFunc;
}
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
DLL_EXPORT void tSetOnFrontDisconnected(Trade *spi, void *onFunc){
    spi->_OnFrontDisconnected = onFunc;
}
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
DLL_EXPORT void tSetOnHeartBeatWarning(Trade *spi, void *onFunc){
    spi->_OnHeartBeatWarning = onFunc;
}
// 客户端认证响应
DLL_EXPORT void tSetOnRspAuthenticate(Trade *spi, void *onFunc){
    spi->_OnRspAuthenticate = onFunc;
}
// 登录请求响应
DLL_EXPORT void tSetOnRspUserLogin(Trade *spi, void *onFunc){
    spi->_OnRspUserLogin = onFunc;
}
// 登出请求响应
DLL_EXPORT void tSetOnRspUserLogout(Trade *spi, void *onFunc){
    spi->_OnRspUserLogout = onFunc;
}
// 用户口令更新请求响应
DLL_EXPORT void tSetOnRspUserPasswordUpdate(Trade *spi, void *onFunc){
    spi->_OnRspUserPasswordUpdate = onFunc;
}
// 资金账户口令更新请求响应
DLL_EXPORT void tSetOnRspTradingAccountPasswordUpdate(Trade *spi, void *onFunc){
    spi->_OnRspTradingAccountPasswordUpdate = onFunc;
}
// 查询用户当前支持的认证模式的回复
DLL_EXPORT void tSetOnRspUserAuthMethod(Trade *spi, void *onFunc){
    spi->_OnRspUserAuthMethod = onFunc;
}
// 获取图形验证码请求的回复
DLL_EXPORT void tSetOnRspGenUserCaptcha(Trade *spi, void *onFunc){
    spi->_OnRspGenUserCaptcha = onFunc;
}
// 获取短信验证码请求的回复
DLL_EXPORT void tSetOnRspGenUserText(Trade *spi, void *onFunc){
    spi->_OnRspGenUserText = onFunc;
}
// 报单录入请求响应
DLL_EXPORT void tSetOnRspOrderInsert(Trade *spi, void *onFunc){
    spi->_OnRspOrderInsert = onFunc;
}
// 预埋单录入请求响应
DLL_EXPORT void tSetOnRspParkedOrderInsert(Trade *spi, void *onFunc){
    spi->_OnRspParkedOrderInsert = onFunc;
}
// 预埋撤单录入请求响应
DLL_EXPORT void tSetOnRspParkedOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspParkedOrderAction = onFunc;
}
// 报单操作请求响应
DLL_EXPORT void tSetOnRspOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspOrderAction = onFunc;
}
// 查询最大报单数量响应
DLL_EXPORT void tSetOnRspQryMaxOrderVolume(Trade *spi, void *onFunc){
    spi->_OnRspQryMaxOrderVolume = onFunc;
}
// 投资者结算结果确认响应
DLL_EXPORT void tSetOnRspSettlementInfoConfirm(Trade *spi, void *onFunc){
    spi->_OnRspSettlementInfoConfirm = onFunc;
}
// 删除预埋单响应
DLL_EXPORT void tSetOnRspRemoveParkedOrder(Trade *spi, void *onFunc){
    spi->_OnRspRemoveParkedOrder = onFunc;
}
// 删除预埋撤单响应
DLL_EXPORT void tSetOnRspRemoveParkedOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspRemoveParkedOrderAction = onFunc;
}
// 执行宣告录入请求响应
DLL_EXPORT void tSetOnRspExecOrderInsert(Trade *spi, void *onFunc){
    spi->_OnRspExecOrderInsert = onFunc;
}
// 执行宣告操作请求响应
DLL_EXPORT void tSetOnRspExecOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspExecOrderAction = onFunc;
}
// 询价录入请求响应
DLL_EXPORT void tSetOnRspForQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnRspForQuoteInsert = onFunc;
}
// 报价录入请求响应
DLL_EXPORT void tSetOnRspQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnRspQuoteInsert = onFunc;
}
// 报价操作请求响应
DLL_EXPORT void tSetOnRspQuoteAction(Trade *spi, void *onFunc){
    spi->_OnRspQuoteAction = onFunc;
}
// 批量报单操作请求响应
DLL_EXPORT void tSetOnRspBatchOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspBatchOrderAction = onFunc;
}
// 期权自对冲录入请求响应
DLL_EXPORT void tSetOnRspOptionSelfCloseInsert(Trade *spi, void *onFunc){
    spi->_OnRspOptionSelfCloseInsert = onFunc;
}
// 期权自对冲操作请求响应
DLL_EXPORT void tSetOnRspOptionSelfCloseAction(Trade *spi, void *onFunc){
    spi->_OnRspOptionSelfCloseAction = onFunc;
}
// 申请组合录入请求响应
DLL_EXPORT void tSetOnRspCombActionInsert(Trade *spi, void *onFunc){
    spi->_OnRspCombActionInsert = onFunc;
}
// 请求查询报单响应
DLL_EXPORT void tSetOnRspQryOrder(Trade *spi, void *onFunc){
    spi->_OnRspQryOrder = onFunc;
}
// 请求查询成交响应
DLL_EXPORT void tSetOnRspQryTrade(Trade *spi, void *onFunc){
    spi->_OnRspQryTrade = onFunc;
}
// 请求查询投资者持仓响应
DLL_EXPORT void tSetOnRspQryInvestorPosition(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorPosition = onFunc;
}
// 请求查询资金账户响应
DLL_EXPORT void tSetOnRspQryTradingAccount(Trade *spi, void *onFunc){
    spi->_OnRspQryTradingAccount = onFunc;
}
// 请求查询投资者响应
DLL_EXPORT void tSetOnRspQryInvestor(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestor = onFunc;
}
// 请求查询交易编码响应
DLL_EXPORT void tSetOnRspQryTradingCode(Trade *spi, void *onFunc){
    spi->_OnRspQryTradingCode = onFunc;
}
// 请求查询合约保证金率响应
DLL_EXPORT void tSetOnRspQryInstrumentMarginRate(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrumentMarginRate = onFunc;
}
// 请求查询合约手续费率响应
DLL_EXPORT void tSetOnRspQryInstrumentCommissionRate(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrumentCommissionRate = onFunc;
}
// 请求查询交易所响应
DLL_EXPORT void tSetOnRspQryExchange(Trade *spi, void *onFunc){
    spi->_OnRspQryExchange = onFunc;
}
// 请求查询产品响应
DLL_EXPORT void tSetOnRspQryProduct(Trade *spi, void *onFunc){
    spi->_OnRspQryProduct = onFunc;
}
// 请求查询合约响应
DLL_EXPORT void tSetOnRspQryInstrument(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrument = onFunc;
}
// 请求查询行情响应
DLL_EXPORT void tSetOnRspQryDepthMarketData(Trade *spi, void *onFunc){
    spi->_OnRspQryDepthMarketData = onFunc;
}
// 请求查询交易员报盘机响应
DLL_EXPORT void tSetOnRspQryTraderOffer(Trade *spi, void *onFunc){
    spi->_OnRspQryTraderOffer = onFunc;
}
// 请求查询投资者结算结果响应
DLL_EXPORT void tSetOnRspQrySettlementInfo(Trade *spi, void *onFunc){
    spi->_OnRspQrySettlementInfo = onFunc;
}
// 请求查询转帐银行响应
DLL_EXPORT void tSetOnRspQryTransferBank(Trade *spi, void *onFunc){
    spi->_OnRspQryTransferBank = onFunc;
}
// 请求查询投资者持仓明细响应
DLL_EXPORT void tSetOnRspQryInvestorPositionDetail(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorPositionDetail = onFunc;
}
// 请求查询客户通知响应
DLL_EXPORT void tSetOnRspQryNotice(Trade *spi, void *onFunc){
    spi->_OnRspQryNotice = onFunc;
}
// 请求查询结算信息确认响应
DLL_EXPORT void tSetOnRspQrySettlementInfoConfirm(Trade *spi, void *onFunc){
    spi->_OnRspQrySettlementInfoConfirm = onFunc;
}
// 请求查询投资者持仓明细响应
DLL_EXPORT void tSetOnRspQryInvestorPositionCombineDetail(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorPositionCombineDetail = onFunc;
}
// 查询保证金监管系统经纪公司资金账户密钥响应
DLL_EXPORT void tSetOnRspQryCFMMCTradingAccountKey(Trade *spi, void *onFunc){
    spi->_OnRspQryCFMMCTradingAccountKey = onFunc;
}
// 请求查询仓单折抵信息响应
DLL_EXPORT void tSetOnRspQryEWarrantOffset(Trade *spi, void *onFunc){
    spi->_OnRspQryEWarrantOffset = onFunc;
}
// 请求查询投资者品种/跨品种保证金响应
DLL_EXPORT void tSetOnRspQryInvestorProductGroupMargin(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestorProductGroupMargin = onFunc;
}
// 请求查询交易所保证金率响应
DLL_EXPORT void tSetOnRspQryExchangeMarginRate(Trade *spi, void *onFunc){
    spi->_OnRspQryExchangeMarginRate = onFunc;
}
// 请求查询交易所调整保证金率响应
DLL_EXPORT void tSetOnRspQryExchangeMarginRateAdjust(Trade *spi, void *onFunc){
    spi->_OnRspQryExchangeMarginRateAdjust = onFunc;
}
// 请求查询汇率响应
DLL_EXPORT void tSetOnRspQryExchangeRate(Trade *spi, void *onFunc){
    spi->_OnRspQryExchangeRate = onFunc;
}
// 请求查询二级代理操作员银期权限响应
DLL_EXPORT void tSetOnRspQrySecAgentACIDMap(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentACIDMap = onFunc;
}
// 请求查询产品报价汇率
DLL_EXPORT void tSetOnRspQryProductExchRate(Trade *spi, void *onFunc){
    spi->_OnRspQryProductExchRate = onFunc;
}
// 请求查询产品组
DLL_EXPORT void tSetOnRspQryProductGroup(Trade *spi, void *onFunc){
    spi->_OnRspQryProductGroup = onFunc;
}
// 请求查询做市商合约手续费率响应
DLL_EXPORT void tSetOnRspQryMMInstrumentCommissionRate(Trade *spi, void *onFunc){
    spi->_OnRspQryMMInstrumentCommissionRate = onFunc;
}
// 请求查询做市商期权合约手续费响应
DLL_EXPORT void tSetOnRspQryMMOptionInstrCommRate(Trade *spi, void *onFunc){
    spi->_OnRspQryMMOptionInstrCommRate = onFunc;
}
// 请求查询报单手续费响应
DLL_EXPORT void tSetOnRspQryInstrumentOrderCommRate(Trade *spi, void *onFunc){
    spi->_OnRspQryInstrumentOrderCommRate = onFunc;
}
// 请求查询资金账户响应
DLL_EXPORT void tSetOnRspQrySecAgentTradingAccount(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentTradingAccount = onFunc;
}
// 请求查询二级代理商资金校验模式响应
DLL_EXPORT void tSetOnRspQrySecAgentCheckMode(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentCheckMode = onFunc;
}
// 请求查询二级代理商信息响应
DLL_EXPORT void tSetOnRspQrySecAgentTradeInfo(Trade *spi, void *onFunc){
    spi->_OnRspQrySecAgentTradeInfo = onFunc;
}
// 请求查询期权交易成本响应
DLL_EXPORT void tSetOnRspQryOptionInstrTradeCost(Trade *spi, void *onFunc){
    spi->_OnRspQryOptionInstrTradeCost = onFunc;
}
// 请求查询期权合约手续费响应
DLL_EXPORT void tSetOnRspQryOptionInstrCommRate(Trade *spi, void *onFunc){
    spi->_OnRspQryOptionInstrCommRate = onFunc;
}
// 请求查询执行宣告响应
DLL_EXPORT void tSetOnRspQryExecOrder(Trade *spi, void *onFunc){
    spi->_OnRspQryExecOrder = onFunc;
}
// 请求查询询价响应
DLL_EXPORT void tSetOnRspQryForQuote(Trade *spi, void *onFunc){
    spi->_OnRspQryForQuote = onFunc;
}
// 请求查询报价响应
DLL_EXPORT void tSetOnRspQryQuote(Trade *spi, void *onFunc){
    spi->_OnRspQryQuote = onFunc;
}
// 请求查询期权自对冲响应
DLL_EXPORT void tSetOnRspQryOptionSelfClose(Trade *spi, void *onFunc){
    spi->_OnRspQryOptionSelfClose = onFunc;
}
// 请求查询投资单元响应
DLL_EXPORT void tSetOnRspQryInvestUnit(Trade *spi, void *onFunc){
    spi->_OnRspQryInvestUnit = onFunc;
}
// 请求查询组合合约安全系数响应
DLL_EXPORT void tSetOnRspQryCombInstrumentGuard(Trade *spi, void *onFunc){
    spi->_OnRspQryCombInstrumentGuard = onFunc;
}
// 请求查询申请组合响应
DLL_EXPORT void tSetOnRspQryCombAction(Trade *spi, void *onFunc){
    spi->_OnRspQryCombAction = onFunc;
}
// 请求查询转帐流水响应
DLL_EXPORT void tSetOnRspQryTransferSerial(Trade *spi, void *onFunc){
    spi->_OnRspQryTransferSerial = onFunc;
}
// 请求查询银期签约关系响应
DLL_EXPORT void tSetOnRspQryAccountregister(Trade *spi, void *onFunc){
    spi->_OnRspQryAccountregister = onFunc;
}
// 错误应答
DLL_EXPORT void tSetOnRspError(Trade *spi, void *onFunc){
    spi->_OnRspError = onFunc;
}
// 报单通知
DLL_EXPORT void tSetOnRtnOrder(Trade *spi, void *onFunc){
    spi->_OnRtnOrder = onFunc;
}
// 成交通知
DLL_EXPORT void tSetOnRtnTrade(Trade *spi, void *onFunc){
    spi->_OnRtnTrade = onFunc;
}
// 报单录入错误回报
DLL_EXPORT void tSetOnErrRtnOrderInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnOrderInsert = onFunc;
}
// 报单操作错误回报
DLL_EXPORT void tSetOnErrRtnOrderAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnOrderAction = onFunc;
}
// 合约交易状态通知
DLL_EXPORT void tSetOnRtnInstrumentStatus(Trade *spi, void *onFunc){
    spi->_OnRtnInstrumentStatus = onFunc;
}
// 交易所公告通知
DLL_EXPORT void tSetOnRtnBulletin(Trade *spi, void *onFunc){
    spi->_OnRtnBulletin = onFunc;
}
// 交易通知
DLL_EXPORT void tSetOnRtnTradingNotice(Trade *spi, void *onFunc){
    spi->_OnRtnTradingNotice = onFunc;
}
// 提示条件单校验错误
DLL_EXPORT void tSetOnRtnErrorConditionalOrder(Trade *spi, void *onFunc){
    spi->_OnRtnErrorConditionalOrder = onFunc;
}
// 执行宣告通知
DLL_EXPORT void tSetOnRtnExecOrder(Trade *spi, void *onFunc){
    spi->_OnRtnExecOrder = onFunc;
}
// 执行宣告录入错误回报
DLL_EXPORT void tSetOnErrRtnExecOrderInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnExecOrderInsert = onFunc;
}
// 执行宣告操作错误回报
DLL_EXPORT void tSetOnErrRtnExecOrderAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnExecOrderAction = onFunc;
}
// 询价录入错误回报
DLL_EXPORT void tSetOnErrRtnForQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnForQuoteInsert = onFunc;
}
// 报价通知
DLL_EXPORT void tSetOnRtnQuote(Trade *spi, void *onFunc){
    spi->_OnRtnQuote = onFunc;
}
// 报价录入错误回报
DLL_EXPORT void tSetOnErrRtnQuoteInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnQuoteInsert = onFunc;
}
// 报价操作错误回报
DLL_EXPORT void tSetOnErrRtnQuoteAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnQuoteAction = onFunc;
}
// 询价通知
DLL_EXPORT void tSetOnRtnForQuoteRsp(Trade *spi, void *onFunc){
    spi->_OnRtnForQuoteRsp = onFunc;
}
// 保证金监控中心用户令牌
DLL_EXPORT void tSetOnRtnCFMMCTradingAccountToken(Trade *spi, void *onFunc){
    spi->_OnRtnCFMMCTradingAccountToken = onFunc;
}
// 批量报单操作错误回报
DLL_EXPORT void tSetOnErrRtnBatchOrderAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnBatchOrderAction = onFunc;
}
// 期权自对冲通知
DLL_EXPORT void tSetOnRtnOptionSelfClose(Trade *spi, void *onFunc){
    spi->_OnRtnOptionSelfClose = onFunc;
}
// 期权自对冲录入错误回报
DLL_EXPORT void tSetOnErrRtnOptionSelfCloseInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnOptionSelfCloseInsert = onFunc;
}
// 期权自对冲操作错误回报
DLL_EXPORT void tSetOnErrRtnOptionSelfCloseAction(Trade *spi, void *onFunc){
    spi->_OnErrRtnOptionSelfCloseAction = onFunc;
}
// 申请组合通知
DLL_EXPORT void tSetOnRtnCombAction(Trade *spi, void *onFunc){
    spi->_OnRtnCombAction = onFunc;
}
// 申请组合录入错误回报
DLL_EXPORT void tSetOnErrRtnCombActionInsert(Trade *spi, void *onFunc){
    spi->_OnErrRtnCombActionInsert = onFunc;
}
// 请求查询签约银行响应
DLL_EXPORT void tSetOnRspQryContractBank(Trade *spi, void *onFunc){
    spi->_OnRspQryContractBank = onFunc;
}
// 请求查询预埋单响应
DLL_EXPORT void tSetOnRspQryParkedOrder(Trade *spi, void *onFunc){
    spi->_OnRspQryParkedOrder = onFunc;
}
// 请求查询预埋撤单响应
DLL_EXPORT void tSetOnRspQryParkedOrderAction(Trade *spi, void *onFunc){
    spi->_OnRspQryParkedOrderAction = onFunc;
}
// 请求查询交易通知响应
DLL_EXPORT void tSetOnRspQryTradingNotice(Trade *spi, void *onFunc){
    spi->_OnRspQryTradingNotice = onFunc;
}
// 请求查询经纪公司交易参数响应
DLL_EXPORT void tSetOnRspQryBrokerTradingParams(Trade *spi, void *onFunc){
    spi->_OnRspQryBrokerTradingParams = onFunc;
}
// 请求查询经纪公司交易算法响应
DLL_EXPORT void tSetOnRspQryBrokerTradingAlgos(Trade *spi, void *onFunc){
    spi->_OnRspQryBrokerTradingAlgos = onFunc;
}
// 请求查询监控中心用户令牌
DLL_EXPORT void tSetOnRspQueryCFMMCTradingAccountToken(Trade *spi, void *onFunc){
    spi->_OnRspQueryCFMMCTradingAccountToken = onFunc;
}
// 银行发起银行资金转期货通知
DLL_EXPORT void tSetOnRtnFromBankToFutureByBank(Trade *spi, void *onFunc){
    spi->_OnRtnFromBankToFutureByBank = onFunc;
}
// 银行发起期货资金转银行通知
DLL_EXPORT void tSetOnRtnFromFutureToBankByBank(Trade *spi, void *onFunc){
    spi->_OnRtnFromFutureToBankByBank = onFunc;
}
// 银行发起冲正银行转期货通知
DLL_EXPORT void tSetOnRtnRepealFromBankToFutureByBank(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromBankToFutureByBank = onFunc;
}
// 银行发起冲正期货转银行通知
DLL_EXPORT void tSetOnRtnRepealFromFutureToBankByBank(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromFutureToBankByBank = onFunc;
}
// 期货发起银行资金转期货通知
DLL_EXPORT void tSetOnRtnFromBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnFromBankToFutureByFuture = onFunc;
}
// 期货发起期货资金转银行通知
DLL_EXPORT void tSetOnRtnFromFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnFromFutureToBankByFuture = onFunc;
}
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void tSetOnRtnRepealFromBankToFutureByFutureManual(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromBankToFutureByFutureManual = onFunc;
}
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void tSetOnRtnRepealFromFutureToBankByFutureManual(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromFutureToBankByFutureManual = onFunc;
}
// 期货发起查询银行余额通知
DLL_EXPORT void tSetOnRtnQueryBankBalanceByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnQueryBankBalanceByFuture = onFunc;
}
// 期货发起银行资金转期货错误回报
DLL_EXPORT void tSetOnErrRtnBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnErrRtnBankToFutureByFuture = onFunc;
}
// 期货发起期货资金转银行错误回报
DLL_EXPORT void tSetOnErrRtnFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnErrRtnFutureToBankByFuture = onFunc;
}
// 系统运行时期货端手工发起冲正银行转期货错误回报
DLL_EXPORT void tSetOnErrRtnRepealBankToFutureByFutureManual(Trade *spi, void *onFunc){
    spi->_OnErrRtnRepealBankToFutureByFutureManual = onFunc;
}
// 系统运行时期货端手工发起冲正期货转银行错误回报
DLL_EXPORT void tSetOnErrRtnRepealFutureToBankByFutureManual(Trade *spi, void *onFunc){
    spi->_OnErrRtnRepealFutureToBankByFutureManual = onFunc;
}
// 期货发起查询银行余额错误回报
DLL_EXPORT void tSetOnErrRtnQueryBankBalanceByFuture(Trade *spi, void *onFunc){
    spi->_OnErrRtnQueryBankBalanceByFuture = onFunc;
}
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void tSetOnRtnRepealFromBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromBankToFutureByFuture = onFunc;
}
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
DLL_EXPORT void tSetOnRtnRepealFromFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnRtnRepealFromFutureToBankByFuture = onFunc;
}
// 期货发起银行资金转期货应答
DLL_EXPORT void tSetOnRspFromBankToFutureByFuture(Trade *spi, void *onFunc){
    spi->_OnRspFromBankToFutureByFuture = onFunc;
}
// 期货发起期货资金转银行应答
DLL_EXPORT void tSetOnRspFromFutureToBankByFuture(Trade *spi, void *onFunc){
    spi->_OnRspFromFutureToBankByFuture = onFunc;
}
// 期货发起查询银行余额应答
DLL_EXPORT void tSetOnRspQueryBankAccountMoneyByFuture(Trade *spi, void *onFunc){
    spi->_OnRspQueryBankAccountMoneyByFuture = onFunc;
}
// 银行发起银期开户通知
DLL_EXPORT void tSetOnRtnOpenAccountByBank(Trade *spi, void *onFunc){
    spi->_OnRtnOpenAccountByBank = onFunc;
}
// 银行发起银期销户通知
DLL_EXPORT void tSetOnRtnCancelAccountByBank(Trade *spi, void *onFunc){
    spi->_OnRtnCancelAccountByBank = onFunc;
}
// 银行发起变更银行账号通知
DLL_EXPORT void tSetOnRtnChangeAccountByBank(Trade *spi, void *onFunc){
    spi->_OnRtnChangeAccountByBank = onFunc;
}
// 请求查询分类合约响应
DLL_EXPORT void tSetOnRspQryClassifiedInstrument(Trade *spi, void *onFunc){
    spi->_OnRspQryClassifiedInstrument = onFunc;
}
// 请求组合优惠比例响应
DLL_EXPORT void tSetOnRspQryCombPromotionParam(Trade *spi, void *onFunc){
    spi->_OnRspQryCombPromotionParam = onFunc;
}
// 投资者风险结算持仓查询响应
DLL_EXPORT void tSetOnRspQryRiskSettleInvstPosition(Trade *spi, void *onFunc){
    spi->_OnRspQryRiskSettleInvstPosition = onFunc;
}
// 风险结算产品查询响应
DLL_EXPORT void tSetOnRspQryRiskSettleProductStatus(Trade *spi, void *onFunc){
    spi->_OnRspQryRiskSettleProductStatus = onFunc;
}
