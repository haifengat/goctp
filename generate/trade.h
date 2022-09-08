#pragma once
#ifdef _WINDOWS  //64位系统没有预定义 WIN32
#ifdef __cplusplus
#define DLL_EXPORT_C_DECL extern "C" __declspec(dllexport)
#else
#define DLL_EXPORT_DECL __declspec(dllexport)
#endif
#else
#ifdef __cplusplus
#define DLL_EXPORT_C_DECL extern "C"
#else
#define DLL_EXPORT_DECL extern
#endif
#endif

#ifdef _WINDOWS
#define WIN32_LEAN_AND_MEAN             //  从 Windows 头文件中排除极少使用的信息
#include "stddef.h"
#ifdef WIN32
#define WINAPI      __cdecl
#include "../v6.6.8_20220712/ThostFtdcTraderApi.h"
#pragma comment(lib, "../v6.6.8_20220712/thosttraderapi_se.lib")
#else
#define WINAPI      __stdcall
#include "../v6.6.8_20220712/ThostFtdcTraderApi.h"
#pragma comment(lib, "../v6.6.8_20220712/thosttraderapi_se.lib")
#endif
#else
#define WINAPI
#include "../v6.6.8_20220712/ThostFtdcTraderApi.h"
#endif

#include <string.h>

class Trade: CThostFtdcTraderSpi
{
public:
    Trade(void);
    //针对收到空反馈的处理
    CThostFtdcRspInfoField rif;
    CThostFtdcRspInfoField* repare(CThostFtdcRspInfoField *pRspInfo)
    {
        if (pRspInfo == NULL)
        {
            memset(&rif, 0, sizeof(rif));
            return &rif;
        }
        else
            return pRspInfo;
    }

    
	// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	typedef int (WINAPI *FrontConnectedType)();
	void *_FrontConnected;
	virtual void OnFrontConnected(){if (_FrontConnected) ((FrontConnectedType)_FrontConnected)();}
    
	// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
	typedef int (WINAPI *FrontDisconnectedType)(int nReason);
	void *_FrontDisconnected;
	virtual void OnFrontDisconnected(int nReason){if (_FrontDisconnected) ((FrontDisconnectedType)_FrontDisconnected)(nReason);}
    
	// 心跳超时警告。当长时间未收到报文时，该方法被调用。
	typedef int (WINAPI *HeartBeatWarningType)(int nTimeLapse);
	void *_HeartBeatWarning;
	virtual void OnHeartBeatWarning(int nTimeLapse){if (_HeartBeatWarning) ((HeartBeatWarningType)_HeartBeatWarning)(nTimeLapse);}
    
	// 客户端认证响应
	typedef int (WINAPI *RspAuthenticateType)(CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspAuthenticate;
	virtual void OnRspAuthenticate(CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspAuthenticate) ((RspAuthenticateType)_RspAuthenticate)(pRspAuthenticateField, pRspInfo, nRequestID, bIsLast);}
    
	// 登录请求响应
	typedef int (WINAPI *RspUserLoginType)(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUserLogin;
	virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUserLogin) ((RspUserLoginType)_RspUserLogin)(pRspUserLogin, pRspInfo, nRequestID, bIsLast);}
    
	// 登出请求响应
	typedef int (WINAPI *RspUserLogoutType)(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUserLogout;
	virtual void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUserLogout) ((RspUserLogoutType)_RspUserLogout)(pUserLogout, pRspInfo, nRequestID, bIsLast);}
    
	// 用户口令更新请求响应
	typedef int (WINAPI *RspUserPasswordUpdateType)(CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUserPasswordUpdate;
	virtual void OnRspUserPasswordUpdate(CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUserPasswordUpdate) ((RspUserPasswordUpdateType)_RspUserPasswordUpdate)(pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);}
    
	// 资金账户口令更新请求响应
	typedef int (WINAPI *RspTradingAccountPasswordUpdateType)(CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspTradingAccountPasswordUpdate;
	virtual void OnRspTradingAccountPasswordUpdate(CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspTradingAccountPasswordUpdate) ((RspTradingAccountPasswordUpdateType)_RspTradingAccountPasswordUpdate)(pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast);}
    
	// 查询用户当前支持的认证模式的回复
	typedef int (WINAPI *RspUserAuthMethodType)(CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUserAuthMethod;
	virtual void OnRspUserAuthMethod(CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUserAuthMethod) ((RspUserAuthMethodType)_RspUserAuthMethod)(pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast);}
    
	// 获取图形验证码请求的回复
	typedef int (WINAPI *RspGenUserCaptchaType)(CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspGenUserCaptcha;
	virtual void OnRspGenUserCaptcha(CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspGenUserCaptcha) ((RspGenUserCaptchaType)_RspGenUserCaptcha)(pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast);}
    
	// 获取短信验证码请求的回复
	typedef int (WINAPI *RspGenUserTextType)(CThostFtdcRspGenUserTextField *pRspGenUserText, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspGenUserText;
	virtual void OnRspGenUserText(CThostFtdcRspGenUserTextField *pRspGenUserText, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspGenUserText) ((RspGenUserTextType)_RspGenUserText)(pRspGenUserText, pRspInfo, nRequestID, bIsLast);}
    
	// 报单录入请求响应
	typedef int (WINAPI *RspOrderInsertType)(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspOrderInsert;
	virtual void OnRspOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspOrderInsert) ((RspOrderInsertType)_RspOrderInsert)(pInputOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 预埋单录入请求响应
	typedef int (WINAPI *RspParkedOrderInsertType)(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspParkedOrderInsert;
	virtual void OnRspParkedOrderInsert(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspParkedOrderInsert) ((RspParkedOrderInsertType)_RspParkedOrderInsert)(pParkedOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 预埋撤单录入请求响应
	typedef int (WINAPI *RspParkedOrderActionType)(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspParkedOrderAction;
	virtual void OnRspParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspParkedOrderAction) ((RspParkedOrderActionType)_RspParkedOrderAction)(pParkedOrderAction, pRspInfo, nRequestID, bIsLast);}
    
	// 报单操作请求响应
	typedef int (WINAPI *RspOrderActionType)(CThostFtdcInputOrderActionField *pInputOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspOrderAction;
	virtual void OnRspOrderAction(CThostFtdcInputOrderActionField *pInputOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspOrderAction) ((RspOrderActionType)_RspOrderAction)(pInputOrderAction, pRspInfo, nRequestID, bIsLast);}
    
	// 查询最大报单数量响应
	typedef int (WINAPI *RspQryMaxOrderVolumeType)(CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryMaxOrderVolume;
	virtual void OnRspQryMaxOrderVolume(CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryMaxOrderVolume) ((RspQryMaxOrderVolumeType)_RspQryMaxOrderVolume)(pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);}
    
	// 投资者结算结果确认响应
	typedef int (WINAPI *RspSettlementInfoConfirmType)(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspSettlementInfoConfirm;
	virtual void OnRspSettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspSettlementInfoConfirm) ((RspSettlementInfoConfirmType)_RspSettlementInfoConfirm)(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);}
    
	// 删除预埋单响应
	typedef int (WINAPI *RspRemoveParkedOrderType)(CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspRemoveParkedOrder;
	virtual void OnRspRemoveParkedOrder(CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspRemoveParkedOrder) ((RspRemoveParkedOrderType)_RspRemoveParkedOrder)(pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 删除预埋撤单响应
	typedef int (WINAPI *RspRemoveParkedOrderActionType)(CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspRemoveParkedOrderAction;
	virtual void OnRspRemoveParkedOrderAction(CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspRemoveParkedOrderAction) ((RspRemoveParkedOrderActionType)_RspRemoveParkedOrderAction)(pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast);}
    
	// 执行宣告录入请求响应
	typedef int (WINAPI *RspExecOrderInsertType)(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspExecOrderInsert;
	virtual void OnRspExecOrderInsert(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspExecOrderInsert) ((RspExecOrderInsertType)_RspExecOrderInsert)(pInputExecOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 执行宣告操作请求响应
	typedef int (WINAPI *RspExecOrderActionType)(CThostFtdcInputExecOrderActionField *pInputExecOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspExecOrderAction;
	virtual void OnRspExecOrderAction(CThostFtdcInputExecOrderActionField *pInputExecOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspExecOrderAction) ((RspExecOrderActionType)_RspExecOrderAction)(pInputExecOrderAction, pRspInfo, nRequestID, bIsLast);}
    
	// 询价录入请求响应
	typedef int (WINAPI *RspForQuoteInsertType)(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspForQuoteInsert;
	virtual void OnRspForQuoteInsert(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspForQuoteInsert) ((RspForQuoteInsertType)_RspForQuoteInsert)(pInputForQuote, pRspInfo, nRequestID, bIsLast);}
    
	// 报价录入请求响应
	typedef int (WINAPI *RspQuoteInsertType)(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQuoteInsert;
	virtual void OnRspQuoteInsert(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQuoteInsert) ((RspQuoteInsertType)_RspQuoteInsert)(pInputQuote, pRspInfo, nRequestID, bIsLast);}
    
	// 报价操作请求响应
	typedef int (WINAPI *RspQuoteActionType)(CThostFtdcInputQuoteActionField *pInputQuoteAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQuoteAction;
	virtual void OnRspQuoteAction(CThostFtdcInputQuoteActionField *pInputQuoteAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQuoteAction) ((RspQuoteActionType)_RspQuoteAction)(pInputQuoteAction, pRspInfo, nRequestID, bIsLast);}
    
	// 批量报单操作请求响应
	typedef int (WINAPI *RspBatchOrderActionType)(CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspBatchOrderAction;
	virtual void OnRspBatchOrderAction(CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspBatchOrderAction) ((RspBatchOrderActionType)_RspBatchOrderAction)(pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast);}
    
	// 期权自对冲录入请求响应
	typedef int (WINAPI *RspOptionSelfCloseInsertType)(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspOptionSelfCloseInsert;
	virtual void OnRspOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspOptionSelfCloseInsert) ((RspOptionSelfCloseInsertType)_RspOptionSelfCloseInsert)(pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast);}
    
	// 期权自对冲操作请求响应
	typedef int (WINAPI *RspOptionSelfCloseActionType)(CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspOptionSelfCloseAction;
	virtual void OnRspOptionSelfCloseAction(CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspOptionSelfCloseAction) ((RspOptionSelfCloseActionType)_RspOptionSelfCloseAction)(pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);}
    
	// 申请组合录入请求响应
	typedef int (WINAPI *RspCombActionInsertType)(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspCombActionInsert;
	virtual void OnRspCombActionInsert(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspCombActionInsert) ((RspCombActionInsertType)_RspCombActionInsert)(pInputCombAction, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询报单响应
	typedef int (WINAPI *RspQryOrderType)(CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryOrder;
	virtual void OnRspQryOrder(CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryOrder) ((RspQryOrderType)_RspQryOrder)(pOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询成交响应
	typedef int (WINAPI *RspQryTradeType)(CThostFtdcTradeField *pTrade, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTrade;
	virtual void OnRspQryTrade(CThostFtdcTradeField *pTrade, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTrade) ((RspQryTradeType)_RspQryTrade)(pTrade, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资者持仓响应
	typedef int (WINAPI *RspQryInvestorPositionType)(CThostFtdcInvestorPositionField *pInvestorPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInvestorPosition;
	virtual void OnRspQryInvestorPosition(CThostFtdcInvestorPositionField *pInvestorPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInvestorPosition) ((RspQryInvestorPositionType)_RspQryInvestorPosition)(pInvestorPosition, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询资金账户响应
	typedef int (WINAPI *RspQryTradingAccountType)(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTradingAccount;
	virtual void OnRspQryTradingAccount(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTradingAccount) ((RspQryTradingAccountType)_RspQryTradingAccount)(pTradingAccount, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资者响应
	typedef int (WINAPI *RspQryInvestorType)(CThostFtdcInvestorField *pInvestor, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInvestor;
	virtual void OnRspQryInvestor(CThostFtdcInvestorField *pInvestor, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInvestor) ((RspQryInvestorType)_RspQryInvestor)(pInvestor, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询交易编码响应
	typedef int (WINAPI *RspQryTradingCodeType)(CThostFtdcTradingCodeField *pTradingCode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTradingCode;
	virtual void OnRspQryTradingCode(CThostFtdcTradingCodeField *pTradingCode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTradingCode) ((RspQryTradingCodeType)_RspQryTradingCode)(pTradingCode, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询合约保证金率响应
	typedef int (WINAPI *RspQryInstrumentMarginRateType)(CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInstrumentMarginRate;
	virtual void OnRspQryInstrumentMarginRate(CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInstrumentMarginRate) ((RspQryInstrumentMarginRateType)_RspQryInstrumentMarginRate)(pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询合约手续费率响应
	typedef int (WINAPI *RspQryInstrumentCommissionRateType)(CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInstrumentCommissionRate;
	virtual void OnRspQryInstrumentCommissionRate(CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInstrumentCommissionRate) ((RspQryInstrumentCommissionRateType)_RspQryInstrumentCommissionRate)(pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询交易所响应
	typedef int (WINAPI *RspQryExchangeType)(CThostFtdcExchangeField *pExchange, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryExchange;
	virtual void OnRspQryExchange(CThostFtdcExchangeField *pExchange, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryExchange) ((RspQryExchangeType)_RspQryExchange)(pExchange, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询产品响应
	typedef int (WINAPI *RspQryProductType)(CThostFtdcProductField *pProduct, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryProduct;
	virtual void OnRspQryProduct(CThostFtdcProductField *pProduct, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryProduct) ((RspQryProductType)_RspQryProduct)(pProduct, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询合约响应
	typedef int (WINAPI *RspQryInstrumentType)(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInstrument;
	virtual void OnRspQryInstrument(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInstrument) ((RspQryInstrumentType)_RspQryInstrument)(pInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询行情响应
	typedef int (WINAPI *RspQryDepthMarketDataType)(CThostFtdcDepthMarketDataField *pDepthMarketData, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryDepthMarketData;
	virtual void OnRspQryDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryDepthMarketData) ((RspQryDepthMarketDataType)_RspQryDepthMarketData)(pDepthMarketData, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询交易员报盘机响应
	typedef int (WINAPI *RspQryTraderOfferType)(CThostFtdcTraderOfferField *pTraderOffer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTraderOffer;
	virtual void OnRspQryTraderOffer(CThostFtdcTraderOfferField *pTraderOffer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTraderOffer) ((RspQryTraderOfferType)_RspQryTraderOffer)(pTraderOffer, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资者结算结果响应
	typedef int (WINAPI *RspQrySettlementInfoType)(CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQrySettlementInfo;
	virtual void OnRspQrySettlementInfo(CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQrySettlementInfo) ((RspQrySettlementInfoType)_RspQrySettlementInfo)(pSettlementInfo, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询转帐银行响应
	typedef int (WINAPI *RspQryTransferBankType)(CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTransferBank;
	virtual void OnRspQryTransferBank(CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTransferBank) ((RspQryTransferBankType)_RspQryTransferBank)(pTransferBank, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资者持仓明细响应
	typedef int (WINAPI *RspQryInvestorPositionDetailType)(CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInvestorPositionDetail;
	virtual void OnRspQryInvestorPositionDetail(CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInvestorPositionDetail) ((RspQryInvestorPositionDetailType)_RspQryInvestorPositionDetail)(pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询客户通知响应
	typedef int (WINAPI *RspQryNoticeType)(CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryNotice;
	virtual void OnRspQryNotice(CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryNotice) ((RspQryNoticeType)_RspQryNotice)(pNotice, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询结算信息确认响应
	typedef int (WINAPI *RspQrySettlementInfoConfirmType)(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQrySettlementInfoConfirm;
	virtual void OnRspQrySettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQrySettlementInfoConfirm) ((RspQrySettlementInfoConfirmType)_RspQrySettlementInfoConfirm)(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资者持仓明细响应
	typedef int (WINAPI *RspQryInvestorPositionCombineDetailType)(CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInvestorPositionCombineDetail;
	virtual void OnRspQryInvestorPositionCombineDetail(CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInvestorPositionCombineDetail) ((RspQryInvestorPositionCombineDetailType)_RspQryInvestorPositionCombineDetail)(pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast);}
    
	// 查询保证金监管系统经纪公司资金账户密钥响应
	typedef int (WINAPI *RspQryCFMMCTradingAccountKeyType)(CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryCFMMCTradingAccountKey;
	virtual void OnRspQryCFMMCTradingAccountKey(CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryCFMMCTradingAccountKey) ((RspQryCFMMCTradingAccountKeyType)_RspQryCFMMCTradingAccountKey)(pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询仓单折抵信息响应
	typedef int (WINAPI *RspQryEWarrantOffsetType)(CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryEWarrantOffset;
	virtual void OnRspQryEWarrantOffset(CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryEWarrantOffset) ((RspQryEWarrantOffsetType)_RspQryEWarrantOffset)(pEWarrantOffset, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资者品种/跨品种保证金响应
	typedef int (WINAPI *RspQryInvestorProductGroupMarginType)(CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInvestorProductGroupMargin;
	virtual void OnRspQryInvestorProductGroupMargin(CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInvestorProductGroupMargin) ((RspQryInvestorProductGroupMarginType)_RspQryInvestorProductGroupMargin)(pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询交易所保证金率响应
	typedef int (WINAPI *RspQryExchangeMarginRateType)(CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryExchangeMarginRate;
	virtual void OnRspQryExchangeMarginRate(CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryExchangeMarginRate) ((RspQryExchangeMarginRateType)_RspQryExchangeMarginRate)(pExchangeMarginRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询交易所调整保证金率响应
	typedef int (WINAPI *RspQryExchangeMarginRateAdjustType)(CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryExchangeMarginRateAdjust;
	virtual void OnRspQryExchangeMarginRateAdjust(CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryExchangeMarginRateAdjust) ((RspQryExchangeMarginRateAdjustType)_RspQryExchangeMarginRateAdjust)(pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询汇率响应
	typedef int (WINAPI *RspQryExchangeRateType)(CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryExchangeRate;
	virtual void OnRspQryExchangeRate(CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryExchangeRate) ((RspQryExchangeRateType)_RspQryExchangeRate)(pExchangeRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询二级代理操作员银期权限响应
	typedef int (WINAPI *RspQrySecAgentACIDMapType)(CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQrySecAgentACIDMap;
	virtual void OnRspQrySecAgentACIDMap(CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQrySecAgentACIDMap) ((RspQrySecAgentACIDMapType)_RspQrySecAgentACIDMap)(pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询产品报价汇率
	typedef int (WINAPI *RspQryProductExchRateType)(CThostFtdcProductExchRateField *pProductExchRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryProductExchRate;
	virtual void OnRspQryProductExchRate(CThostFtdcProductExchRateField *pProductExchRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryProductExchRate) ((RspQryProductExchRateType)_RspQryProductExchRate)(pProductExchRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询产品组
	typedef int (WINAPI *RspQryProductGroupType)(CThostFtdcProductGroupField *pProductGroup, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryProductGroup;
	virtual void OnRspQryProductGroup(CThostFtdcProductGroupField *pProductGroup, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryProductGroup) ((RspQryProductGroupType)_RspQryProductGroup)(pProductGroup, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询做市商合约手续费率响应
	typedef int (WINAPI *RspQryMMInstrumentCommissionRateType)(CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryMMInstrumentCommissionRate;
	virtual void OnRspQryMMInstrumentCommissionRate(CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryMMInstrumentCommissionRate) ((RspQryMMInstrumentCommissionRateType)_RspQryMMInstrumentCommissionRate)(pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询做市商期权合约手续费响应
	typedef int (WINAPI *RspQryMMOptionInstrCommRateType)(CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryMMOptionInstrCommRate;
	virtual void OnRspQryMMOptionInstrCommRate(CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryMMOptionInstrCommRate) ((RspQryMMOptionInstrCommRateType)_RspQryMMOptionInstrCommRate)(pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询报单手续费响应
	typedef int (WINAPI *RspQryInstrumentOrderCommRateType)(CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInstrumentOrderCommRate;
	virtual void OnRspQryInstrumentOrderCommRate(CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInstrumentOrderCommRate) ((RspQryInstrumentOrderCommRateType)_RspQryInstrumentOrderCommRate)(pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询资金账户响应
	typedef int (WINAPI *RspQrySecAgentTradingAccountType)(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQrySecAgentTradingAccount;
	virtual void OnRspQrySecAgentTradingAccount(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQrySecAgentTradingAccount) ((RspQrySecAgentTradingAccountType)_RspQrySecAgentTradingAccount)(pTradingAccount, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询二级代理商资金校验模式响应
	typedef int (WINAPI *RspQrySecAgentCheckModeType)(CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQrySecAgentCheckMode;
	virtual void OnRspQrySecAgentCheckMode(CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQrySecAgentCheckMode) ((RspQrySecAgentCheckModeType)_RspQrySecAgentCheckMode)(pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询二级代理商信息响应
	typedef int (WINAPI *RspQrySecAgentTradeInfoType)(CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQrySecAgentTradeInfo;
	virtual void OnRspQrySecAgentTradeInfo(CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQrySecAgentTradeInfo) ((RspQrySecAgentTradeInfoType)_RspQrySecAgentTradeInfo)(pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询期权交易成本响应
	typedef int (WINAPI *RspQryOptionInstrTradeCostType)(CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryOptionInstrTradeCost;
	virtual void OnRspQryOptionInstrTradeCost(CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryOptionInstrTradeCost) ((RspQryOptionInstrTradeCostType)_RspQryOptionInstrTradeCost)(pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询期权合约手续费响应
	typedef int (WINAPI *RspQryOptionInstrCommRateType)(CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryOptionInstrCommRate;
	virtual void OnRspQryOptionInstrCommRate(CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryOptionInstrCommRate) ((RspQryOptionInstrCommRateType)_RspQryOptionInstrCommRate)(pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询执行宣告响应
	typedef int (WINAPI *RspQryExecOrderType)(CThostFtdcExecOrderField *pExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryExecOrder;
	virtual void OnRspQryExecOrder(CThostFtdcExecOrderField *pExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryExecOrder) ((RspQryExecOrderType)_RspQryExecOrder)(pExecOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询询价响应
	typedef int (WINAPI *RspQryForQuoteType)(CThostFtdcForQuoteField *pForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryForQuote;
	virtual void OnRspQryForQuote(CThostFtdcForQuoteField *pForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryForQuote) ((RspQryForQuoteType)_RspQryForQuote)(pForQuote, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询报价响应
	typedef int (WINAPI *RspQryQuoteType)(CThostFtdcQuoteField *pQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryQuote;
	virtual void OnRspQryQuote(CThostFtdcQuoteField *pQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryQuote) ((RspQryQuoteType)_RspQryQuote)(pQuote, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询期权自对冲响应
	typedef int (WINAPI *RspQryOptionSelfCloseType)(CThostFtdcOptionSelfCloseField *pOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryOptionSelfClose;
	virtual void OnRspQryOptionSelfClose(CThostFtdcOptionSelfCloseField *pOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryOptionSelfClose) ((RspQryOptionSelfCloseType)_RspQryOptionSelfClose)(pOptionSelfClose, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询投资单元响应
	typedef int (WINAPI *RspQryInvestUnitType)(CThostFtdcInvestUnitField *pInvestUnit, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryInvestUnit;
	virtual void OnRspQryInvestUnit(CThostFtdcInvestUnitField *pInvestUnit, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryInvestUnit) ((RspQryInvestUnitType)_RspQryInvestUnit)(pInvestUnit, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询组合合约安全系数响应
	typedef int (WINAPI *RspQryCombInstrumentGuardType)(CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryCombInstrumentGuard;
	virtual void OnRspQryCombInstrumentGuard(CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryCombInstrumentGuard) ((RspQryCombInstrumentGuardType)_RspQryCombInstrumentGuard)(pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询申请组合响应
	typedef int (WINAPI *RspQryCombActionType)(CThostFtdcCombActionField *pCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryCombAction;
	virtual void OnRspQryCombAction(CThostFtdcCombActionField *pCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryCombAction) ((RspQryCombActionType)_RspQryCombAction)(pCombAction, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询转帐流水响应
	typedef int (WINAPI *RspQryTransferSerialType)(CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTransferSerial;
	virtual void OnRspQryTransferSerial(CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTransferSerial) ((RspQryTransferSerialType)_RspQryTransferSerial)(pTransferSerial, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询银期签约关系响应
	typedef int (WINAPI *RspQryAccountregisterType)(CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryAccountregister;
	virtual void OnRspQryAccountregister(CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryAccountregister) ((RspQryAccountregisterType)_RspQryAccountregister)(pAccountregister, pRspInfo, nRequestID, bIsLast);}
    
	// 错误应答
	typedef int (WINAPI *RspErrorType)(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspError;
	virtual void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspError) ((RspErrorType)_RspError)(pRspInfo, nRequestID, bIsLast);}
    
	// 报单通知
	typedef int (WINAPI *RtnOrderType)(CThostFtdcOrderField *pOrder);
	void *_RtnOrder;
	virtual void OnRtnOrder(CThostFtdcOrderField *pOrder){if (_RtnOrder) ((RtnOrderType)_RtnOrder)(pOrder);}
    
	// 成交通知
	typedef int (WINAPI *RtnTradeType)(CThostFtdcTradeField *pTrade);
	void *_RtnTrade;
	virtual void OnRtnTrade(CThostFtdcTradeField *pTrade){if (_RtnTrade) ((RtnTradeType)_RtnTrade)(pTrade);}
    
	// 报单录入错误回报
	typedef int (WINAPI *ErrRtnOrderInsertType)(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnOrderInsert;
	virtual void OnErrRtnOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnOrderInsert) ((ErrRtnOrderInsertType)_ErrRtnOrderInsert)(pInputOrder, pRspInfo);}
    
	// 报单操作错误回报
	typedef int (WINAPI *ErrRtnOrderActionType)(CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnOrderAction;
	virtual void OnErrRtnOrderAction(CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnOrderAction) ((ErrRtnOrderActionType)_ErrRtnOrderAction)(pOrderAction, pRspInfo);}
    
	// 合约交易状态通知
	typedef int (WINAPI *RtnInstrumentStatusType)(CThostFtdcInstrumentStatusField *pInstrumentStatus);
	void *_RtnInstrumentStatus;
	virtual void OnRtnInstrumentStatus(CThostFtdcInstrumentStatusField *pInstrumentStatus){if (_RtnInstrumentStatus) ((RtnInstrumentStatusType)_RtnInstrumentStatus)(pInstrumentStatus);}
    
	// 交易所公告通知
	typedef int (WINAPI *RtnBulletinType)(CThostFtdcBulletinField *pBulletin);
	void *_RtnBulletin;
	virtual void OnRtnBulletin(CThostFtdcBulletinField *pBulletin){if (_RtnBulletin) ((RtnBulletinType)_RtnBulletin)(pBulletin);}
    
	// 交易通知
	typedef int (WINAPI *RtnTradingNoticeType)(CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);
	void *_RtnTradingNotice;
	virtual void OnRtnTradingNotice(CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo){if (_RtnTradingNotice) ((RtnTradingNoticeType)_RtnTradingNotice)(pTradingNoticeInfo);}
    
	// 提示条件单校验错误
	typedef int (WINAPI *RtnErrorConditionalOrderType)(CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);
	void *_RtnErrorConditionalOrder;
	virtual void OnRtnErrorConditionalOrder(CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder){if (_RtnErrorConditionalOrder) ((RtnErrorConditionalOrderType)_RtnErrorConditionalOrder)(pErrorConditionalOrder);}
    
	// 执行宣告通知
	typedef int (WINAPI *RtnExecOrderType)(CThostFtdcExecOrderField *pExecOrder);
	void *_RtnExecOrder;
	virtual void OnRtnExecOrder(CThostFtdcExecOrderField *pExecOrder){if (_RtnExecOrder) ((RtnExecOrderType)_RtnExecOrder)(pExecOrder);}
    
	// 执行宣告录入错误回报
	typedef int (WINAPI *ErrRtnExecOrderInsertType)(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnExecOrderInsert;
	virtual void OnErrRtnExecOrderInsert(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnExecOrderInsert) ((ErrRtnExecOrderInsertType)_ErrRtnExecOrderInsert)(pInputExecOrder, pRspInfo);}
    
	// 执行宣告操作错误回报
	typedef int (WINAPI *ErrRtnExecOrderActionType)(CThostFtdcExecOrderActionField *pExecOrderAction, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnExecOrderAction;
	virtual void OnErrRtnExecOrderAction(CThostFtdcExecOrderActionField *pExecOrderAction, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnExecOrderAction) ((ErrRtnExecOrderActionType)_ErrRtnExecOrderAction)(pExecOrderAction, pRspInfo);}
    
	// 询价录入错误回报
	typedef int (WINAPI *ErrRtnForQuoteInsertType)(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnForQuoteInsert;
	virtual void OnErrRtnForQuoteInsert(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnForQuoteInsert) ((ErrRtnForQuoteInsertType)_ErrRtnForQuoteInsert)(pInputForQuote, pRspInfo);}
    
	// 报价通知
	typedef int (WINAPI *RtnQuoteType)(CThostFtdcQuoteField *pQuote);
	void *_RtnQuote;
	virtual void OnRtnQuote(CThostFtdcQuoteField *pQuote){if (_RtnQuote) ((RtnQuoteType)_RtnQuote)(pQuote);}
    
	// 报价录入错误回报
	typedef int (WINAPI *ErrRtnQuoteInsertType)(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnQuoteInsert;
	virtual void OnErrRtnQuoteInsert(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnQuoteInsert) ((ErrRtnQuoteInsertType)_ErrRtnQuoteInsert)(pInputQuote, pRspInfo);}
    
	// 报价操作错误回报
	typedef int (WINAPI *ErrRtnQuoteActionType)(CThostFtdcQuoteActionField *pQuoteAction, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnQuoteAction;
	virtual void OnErrRtnQuoteAction(CThostFtdcQuoteActionField *pQuoteAction, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnQuoteAction) ((ErrRtnQuoteActionType)_ErrRtnQuoteAction)(pQuoteAction, pRspInfo);}
    
	// 询价通知
	typedef int (WINAPI *RtnForQuoteRspType)(CThostFtdcForQuoteRspField *pForQuoteRsp);
	void *_RtnForQuoteRsp;
	virtual void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp){if (_RtnForQuoteRsp) ((RtnForQuoteRspType)_RtnForQuoteRsp)(pForQuoteRsp);}
    
	// 保证金监控中心用户令牌
	typedef int (WINAPI *RtnCFMMCTradingAccountTokenType)(CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);
	void *_RtnCFMMCTradingAccountToken;
	virtual void OnRtnCFMMCTradingAccountToken(CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken){if (_RtnCFMMCTradingAccountToken) ((RtnCFMMCTradingAccountTokenType)_RtnCFMMCTradingAccountToken)(pCFMMCTradingAccountToken);}
    
	// 批量报单操作错误回报
	typedef int (WINAPI *ErrRtnBatchOrderActionType)(CThostFtdcBatchOrderActionField *pBatchOrderAction, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnBatchOrderAction;
	virtual void OnErrRtnBatchOrderAction(CThostFtdcBatchOrderActionField *pBatchOrderAction, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnBatchOrderAction) ((ErrRtnBatchOrderActionType)_ErrRtnBatchOrderAction)(pBatchOrderAction, pRspInfo);}
    
	// 期权自对冲通知
	typedef int (WINAPI *RtnOptionSelfCloseType)(CThostFtdcOptionSelfCloseField *pOptionSelfClose);
	void *_RtnOptionSelfClose;
	virtual void OnRtnOptionSelfClose(CThostFtdcOptionSelfCloseField *pOptionSelfClose){if (_RtnOptionSelfClose) ((RtnOptionSelfCloseType)_RtnOptionSelfClose)(pOptionSelfClose);}
    
	// 期权自对冲录入错误回报
	typedef int (WINAPI *ErrRtnOptionSelfCloseInsertType)(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnOptionSelfCloseInsert;
	virtual void OnErrRtnOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnOptionSelfCloseInsert) ((ErrRtnOptionSelfCloseInsertType)_ErrRtnOptionSelfCloseInsert)(pInputOptionSelfClose, pRspInfo);}
    
	// 期权自对冲操作错误回报
	typedef int (WINAPI *ErrRtnOptionSelfCloseActionType)(CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnOptionSelfCloseAction;
	virtual void OnErrRtnOptionSelfCloseAction(CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnOptionSelfCloseAction) ((ErrRtnOptionSelfCloseActionType)_ErrRtnOptionSelfCloseAction)(pOptionSelfCloseAction, pRspInfo);}
    
	// 申请组合通知
	typedef int (WINAPI *RtnCombActionType)(CThostFtdcCombActionField *pCombAction);
	void *_RtnCombAction;
	virtual void OnRtnCombAction(CThostFtdcCombActionField *pCombAction){if (_RtnCombAction) ((RtnCombActionType)_RtnCombAction)(pCombAction);}
    
	// 申请组合录入错误回报
	typedef int (WINAPI *ErrRtnCombActionInsertType)(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnCombActionInsert;
	virtual void OnErrRtnCombActionInsert(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnCombActionInsert) ((ErrRtnCombActionInsertType)_ErrRtnCombActionInsert)(pInputCombAction, pRspInfo);}
    
	// 请求查询签约银行响应
	typedef int (WINAPI *RspQryContractBankType)(CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryContractBank;
	virtual void OnRspQryContractBank(CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryContractBank) ((RspQryContractBankType)_RspQryContractBank)(pContractBank, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询预埋单响应
	typedef int (WINAPI *RspQryParkedOrderType)(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryParkedOrder;
	virtual void OnRspQryParkedOrder(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryParkedOrder) ((RspQryParkedOrderType)_RspQryParkedOrder)(pParkedOrder, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询预埋撤单响应
	typedef int (WINAPI *RspQryParkedOrderActionType)(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryParkedOrderAction;
	virtual void OnRspQryParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryParkedOrderAction) ((RspQryParkedOrderActionType)_RspQryParkedOrderAction)(pParkedOrderAction, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询交易通知响应
	typedef int (WINAPI *RspQryTradingNoticeType)(CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryTradingNotice;
	virtual void OnRspQryTradingNotice(CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryTradingNotice) ((RspQryTradingNoticeType)_RspQryTradingNotice)(pTradingNotice, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询经纪公司交易参数响应
	typedef int (WINAPI *RspQryBrokerTradingParamsType)(CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryBrokerTradingParams;
	virtual void OnRspQryBrokerTradingParams(CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryBrokerTradingParams) ((RspQryBrokerTradingParamsType)_RspQryBrokerTradingParams)(pBrokerTradingParams, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询经纪公司交易算法响应
	typedef int (WINAPI *RspQryBrokerTradingAlgosType)(CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryBrokerTradingAlgos;
	virtual void OnRspQryBrokerTradingAlgos(CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryBrokerTradingAlgos) ((RspQryBrokerTradingAlgosType)_RspQryBrokerTradingAlgos)(pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询监控中心用户令牌
	typedef int (WINAPI *RspQueryCFMMCTradingAccountTokenType)(CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQueryCFMMCTradingAccountToken;
	virtual void OnRspQueryCFMMCTradingAccountToken(CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQueryCFMMCTradingAccountToken) ((RspQueryCFMMCTradingAccountTokenType)_RspQueryCFMMCTradingAccountToken)(pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast);}
    
	// 银行发起银行资金转期货通知
	typedef int (WINAPI *RtnFromBankToFutureByBankType)(CThostFtdcRspTransferField *pRspTransfer);
	void *_RtnFromBankToFutureByBank;
	virtual void OnRtnFromBankToFutureByBank(CThostFtdcRspTransferField *pRspTransfer){if (_RtnFromBankToFutureByBank) ((RtnFromBankToFutureByBankType)_RtnFromBankToFutureByBank)(pRspTransfer);}
    
	// 银行发起期货资金转银行通知
	typedef int (WINAPI *RtnFromFutureToBankByBankType)(CThostFtdcRspTransferField *pRspTransfer);
	void *_RtnFromFutureToBankByBank;
	virtual void OnRtnFromFutureToBankByBank(CThostFtdcRspTransferField *pRspTransfer){if (_RtnFromFutureToBankByBank) ((RtnFromFutureToBankByBankType)_RtnFromFutureToBankByBank)(pRspTransfer);}
    
	// 银行发起冲正银行转期货通知
	typedef int (WINAPI *RtnRepealFromBankToFutureByBankType)(CThostFtdcRspRepealField *pRspRepeal);
	void *_RtnRepealFromBankToFutureByBank;
	virtual void OnRtnRepealFromBankToFutureByBank(CThostFtdcRspRepealField *pRspRepeal){if (_RtnRepealFromBankToFutureByBank) ((RtnRepealFromBankToFutureByBankType)_RtnRepealFromBankToFutureByBank)(pRspRepeal);}
    
	// 银行发起冲正期货转银行通知
	typedef int (WINAPI *RtnRepealFromFutureToBankByBankType)(CThostFtdcRspRepealField *pRspRepeal);
	void *_RtnRepealFromFutureToBankByBank;
	virtual void OnRtnRepealFromFutureToBankByBank(CThostFtdcRspRepealField *pRspRepeal){if (_RtnRepealFromFutureToBankByBank) ((RtnRepealFromFutureToBankByBankType)_RtnRepealFromFutureToBankByBank)(pRspRepeal);}
    
	// 期货发起银行资金转期货通知
	typedef int (WINAPI *RtnFromBankToFutureByFutureType)(CThostFtdcRspTransferField *pRspTransfer);
	void *_RtnFromBankToFutureByFuture;
	virtual void OnRtnFromBankToFutureByFuture(CThostFtdcRspTransferField *pRspTransfer){if (_RtnFromBankToFutureByFuture) ((RtnFromBankToFutureByFutureType)_RtnFromBankToFutureByFuture)(pRspTransfer);}
    
	// 期货发起期货资金转银行通知
	typedef int (WINAPI *RtnFromFutureToBankByFutureType)(CThostFtdcRspTransferField *pRspTransfer);
	void *_RtnFromFutureToBankByFuture;
	virtual void OnRtnFromFutureToBankByFuture(CThostFtdcRspTransferField *pRspTransfer){if (_RtnFromFutureToBankByFuture) ((RtnFromFutureToBankByFutureType)_RtnFromFutureToBankByFuture)(pRspTransfer);}
    
	// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	typedef int (WINAPI *RtnRepealFromBankToFutureByFutureManualType)(CThostFtdcRspRepealField *pRspRepeal);
	void *_RtnRepealFromBankToFutureByFutureManual;
	virtual void OnRtnRepealFromBankToFutureByFutureManual(CThostFtdcRspRepealField *pRspRepeal){if (_RtnRepealFromBankToFutureByFutureManual) ((RtnRepealFromBankToFutureByFutureManualType)_RtnRepealFromBankToFutureByFutureManual)(pRspRepeal);}
    
	// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	typedef int (WINAPI *RtnRepealFromFutureToBankByFutureManualType)(CThostFtdcRspRepealField *pRspRepeal);
	void *_RtnRepealFromFutureToBankByFutureManual;
	virtual void OnRtnRepealFromFutureToBankByFutureManual(CThostFtdcRspRepealField *pRspRepeal){if (_RtnRepealFromFutureToBankByFutureManual) ((RtnRepealFromFutureToBankByFutureManualType)_RtnRepealFromFutureToBankByFutureManual)(pRspRepeal);}
    
	// 期货发起查询银行余额通知
	typedef int (WINAPI *RtnQueryBankBalanceByFutureType)(CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);
	void *_RtnQueryBankBalanceByFuture;
	virtual void OnRtnQueryBankBalanceByFuture(CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount){if (_RtnQueryBankBalanceByFuture) ((RtnQueryBankBalanceByFutureType)_RtnQueryBankBalanceByFuture)(pNotifyQueryAccount);}
    
	// 期货发起银行资金转期货错误回报
	typedef int (WINAPI *ErrRtnBankToFutureByFutureType)(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnBankToFutureByFuture;
	virtual void OnErrRtnBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnBankToFutureByFuture) ((ErrRtnBankToFutureByFutureType)_ErrRtnBankToFutureByFuture)(pReqTransfer, pRspInfo);}
    
	// 期货发起期货资金转银行错误回报
	typedef int (WINAPI *ErrRtnFutureToBankByFutureType)(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnFutureToBankByFuture;
	virtual void OnErrRtnFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnFutureToBankByFuture) ((ErrRtnFutureToBankByFutureType)_ErrRtnFutureToBankByFuture)(pReqTransfer, pRspInfo);}
    
	// 系统运行时期货端手工发起冲正银行转期货错误回报
	typedef int (WINAPI *ErrRtnRepealBankToFutureByFutureManualType)(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnRepealBankToFutureByFutureManual;
	virtual void OnErrRtnRepealBankToFutureByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnRepealBankToFutureByFutureManual) ((ErrRtnRepealBankToFutureByFutureManualType)_ErrRtnRepealBankToFutureByFutureManual)(pReqRepeal, pRspInfo);}
    
	// 系统运行时期货端手工发起冲正期货转银行错误回报
	typedef int (WINAPI *ErrRtnRepealFutureToBankByFutureManualType)(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnRepealFutureToBankByFutureManual;
	virtual void OnErrRtnRepealFutureToBankByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnRepealFutureToBankByFutureManual) ((ErrRtnRepealFutureToBankByFutureManualType)_ErrRtnRepealFutureToBankByFutureManual)(pReqRepeal, pRspInfo);}
    
	// 期货发起查询银行余额错误回报
	typedef int (WINAPI *ErrRtnQueryBankBalanceByFutureType)(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo);
	void *_ErrRtnQueryBankBalanceByFuture;
	virtual void OnErrRtnQueryBankBalanceByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo){if (_ErrRtnQueryBankBalanceByFuture) ((ErrRtnQueryBankBalanceByFutureType)_ErrRtnQueryBankBalanceByFuture)(pReqQueryAccount, pRspInfo);}
    
	// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	typedef int (WINAPI *RtnRepealFromBankToFutureByFutureType)(CThostFtdcRspRepealField *pRspRepeal);
	void *_RtnRepealFromBankToFutureByFuture;
	virtual void OnRtnRepealFromBankToFutureByFuture(CThostFtdcRspRepealField *pRspRepeal){if (_RtnRepealFromBankToFutureByFuture) ((RtnRepealFromBankToFutureByFutureType)_RtnRepealFromBankToFutureByFuture)(pRspRepeal);}
    
	// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	typedef int (WINAPI *RtnRepealFromFutureToBankByFutureType)(CThostFtdcRspRepealField *pRspRepeal);
	void *_RtnRepealFromFutureToBankByFuture;
	virtual void OnRtnRepealFromFutureToBankByFuture(CThostFtdcRspRepealField *pRspRepeal){if (_RtnRepealFromFutureToBankByFuture) ((RtnRepealFromFutureToBankByFutureType)_RtnRepealFromFutureToBankByFuture)(pRspRepeal);}
    
	// 期货发起银行资金转期货应答
	typedef int (WINAPI *RspFromBankToFutureByFutureType)(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspFromBankToFutureByFuture;
	virtual void OnRspFromBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspFromBankToFutureByFuture) ((RspFromBankToFutureByFutureType)_RspFromBankToFutureByFuture)(pReqTransfer, pRspInfo, nRequestID, bIsLast);}
    
	// 期货发起期货资金转银行应答
	typedef int (WINAPI *RspFromFutureToBankByFutureType)(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspFromFutureToBankByFuture;
	virtual void OnRspFromFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspFromFutureToBankByFuture) ((RspFromFutureToBankByFutureType)_RspFromFutureToBankByFuture)(pReqTransfer, pRspInfo, nRequestID, bIsLast);}
    
	// 期货发起查询银行余额应答
	typedef int (WINAPI *RspQueryBankAccountMoneyByFutureType)(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQueryBankAccountMoneyByFuture;
	virtual void OnRspQueryBankAccountMoneyByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQueryBankAccountMoneyByFuture) ((RspQueryBankAccountMoneyByFutureType)_RspQueryBankAccountMoneyByFuture)(pReqQueryAccount, pRspInfo, nRequestID, bIsLast);}
    
	// 银行发起银期开户通知
	typedef int (WINAPI *RtnOpenAccountByBankType)(CThostFtdcOpenAccountField *pOpenAccount);
	void *_RtnOpenAccountByBank;
	virtual void OnRtnOpenAccountByBank(CThostFtdcOpenAccountField *pOpenAccount){if (_RtnOpenAccountByBank) ((RtnOpenAccountByBankType)_RtnOpenAccountByBank)(pOpenAccount);}
    
	// 银行发起银期销户通知
	typedef int (WINAPI *RtnCancelAccountByBankType)(CThostFtdcCancelAccountField *pCancelAccount);
	void *_RtnCancelAccountByBank;
	virtual void OnRtnCancelAccountByBank(CThostFtdcCancelAccountField *pCancelAccount){if (_RtnCancelAccountByBank) ((RtnCancelAccountByBankType)_RtnCancelAccountByBank)(pCancelAccount);}
    
	// 银行发起变更银行账号通知
	typedef int (WINAPI *RtnChangeAccountByBankType)(CThostFtdcChangeAccountField *pChangeAccount);
	void *_RtnChangeAccountByBank;
	virtual void OnRtnChangeAccountByBank(CThostFtdcChangeAccountField *pChangeAccount){if (_RtnChangeAccountByBank) ((RtnChangeAccountByBankType)_RtnChangeAccountByBank)(pChangeAccount);}
    
	// 请求查询分类合约响应
	typedef int (WINAPI *RspQryClassifiedInstrumentType)(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryClassifiedInstrument;
	virtual void OnRspQryClassifiedInstrument(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryClassifiedInstrument) ((RspQryClassifiedInstrumentType)_RspQryClassifiedInstrument)(pInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 请求组合优惠比例响应
	typedef int (WINAPI *RspQryCombPromotionParamType)(CThostFtdcCombPromotionParamField *pCombPromotionParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryCombPromotionParam;
	virtual void OnRspQryCombPromotionParam(CThostFtdcCombPromotionParamField *pCombPromotionParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryCombPromotionParam) ((RspQryCombPromotionParamType)_RspQryCombPromotionParam)(pCombPromotionParam, pRspInfo, nRequestID, bIsLast);}
    
	// 投资者风险结算持仓查询响应
	typedef int (WINAPI *RspQryRiskSettleInvstPositionType)(CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryRiskSettleInvstPosition;
	virtual void OnRspQryRiskSettleInvstPosition(CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryRiskSettleInvstPosition) ((RspQryRiskSettleInvstPositionType)_RspQryRiskSettleInvstPosition)(pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast);}
    
	// 风险结算产品查询响应
	typedef int (WINAPI *RspQryRiskSettleProductStatusType)(CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryRiskSettleProductStatus;
	virtual void OnRspQryRiskSettleProductStatus(CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryRiskSettleProductStatus) ((RspQryRiskSettleProductStatusType)_RspQryRiskSettleProductStatus)(pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast);}
    
};
