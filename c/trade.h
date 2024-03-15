#include "../CTPv6.7.2_20230913/ThostFtdcTraderApi.h"

class Trade: CThostFtdcTraderSpi{
public:
	// //////////////////////////////////////////////////////////////////////    
    typedef void OnFrontConnectedType(void*);
    void *_OnFrontConnected;
    virtual void OnFrontConnected(){
        if (_OnFrontConnected) {
			((OnFrontConnectedType*)_OnFrontConnected)(this);
		}
    }
	// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。    
    typedef void OnFrontDisconnectedType(void*, int nReason);
    void *_OnFrontDisconnected;
    virtual void OnFrontDisconnected(int nReason){
        if (_OnFrontDisconnected) {
			((OnFrontDisconnectedType*)_OnFrontDisconnected)(this, nReason);
		}
    }
	// 心跳超时警告。当长时间未收到报文时，该方法被调用。    
    typedef void OnHeartBeatWarningType(void*, int nTimeLapse);
    void *_OnHeartBeatWarning;
    virtual void OnHeartBeatWarning(int nTimeLapse){
        if (_OnHeartBeatWarning) {
			((OnHeartBeatWarningType*)_OnHeartBeatWarning)(this, nTimeLapse);
		}
    }
	// 客户端认证响应    
    typedef void OnRspAuthenticateType(void*, CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspAuthenticate;
    virtual void OnRspAuthenticate(CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspAuthenticate) {
			((OnRspAuthenticateType*)_OnRspAuthenticate)(this, pRspAuthenticateField, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 登录请求响应    
    typedef void OnRspUserLoginType(void*, CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserLogin;
    virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspUserLogin) {
			((OnRspUserLoginType*)_OnRspUserLogin)(this, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 登出请求响应    
    typedef void OnRspUserLogoutType(void*, CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserLogout;
    virtual void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspUserLogout) {
			((OnRspUserLogoutType*)_OnRspUserLogout)(this, pUserLogout, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 用户口令更新请求响应    
    typedef void OnRspUserPasswordUpdateType(void*, CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserPasswordUpdate;
    virtual void OnRspUserPasswordUpdate(CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspUserPasswordUpdate) {
			((OnRspUserPasswordUpdateType*)_OnRspUserPasswordUpdate)(this, pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 资金账户口令更新请求响应    
    typedef void OnRspTradingAccountPasswordUpdateType(void*, CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspTradingAccountPasswordUpdate;
    virtual void OnRspTradingAccountPasswordUpdate(CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspTradingAccountPasswordUpdate) {
			((OnRspTradingAccountPasswordUpdateType*)_OnRspTradingAccountPasswordUpdate)(this, pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 查询用户当前支持的认证模式的回复    
    typedef void OnRspUserAuthMethodType(void*, CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserAuthMethod;
    virtual void OnRspUserAuthMethod(CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspUserAuthMethod) {
			((OnRspUserAuthMethodType*)_OnRspUserAuthMethod)(this, pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 获取图形验证码请求的回复    
    typedef void OnRspGenUserCaptchaType(void*, CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspGenUserCaptcha;
    virtual void OnRspGenUserCaptcha(CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspGenUserCaptcha) {
			((OnRspGenUserCaptchaType*)_OnRspGenUserCaptcha)(this, pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 获取短信验证码请求的回复    
    typedef void OnRspGenUserTextType(void*, CThostFtdcRspGenUserTextField *pRspGenUserText, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspGenUserText;
    virtual void OnRspGenUserText(CThostFtdcRspGenUserTextField *pRspGenUserText, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspGenUserText) {
			((OnRspGenUserTextType*)_OnRspGenUserText)(this, pRspGenUserText, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 报单录入请求响应    
    typedef void OnRspOrderInsertType(void*, CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOrderInsert;
    virtual void OnRspOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspOrderInsert) {
			((OnRspOrderInsertType*)_OnRspOrderInsert)(this, pInputOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 预埋单录入请求响应    
    typedef void OnRspParkedOrderInsertType(void*, CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspParkedOrderInsert;
    virtual void OnRspParkedOrderInsert(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspParkedOrderInsert) {
			((OnRspParkedOrderInsertType*)_OnRspParkedOrderInsert)(this, pParkedOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 预埋撤单录入请求响应    
    typedef void OnRspParkedOrderActionType(void*, CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspParkedOrderAction;
    virtual void OnRspParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspParkedOrderAction) {
			((OnRspParkedOrderActionType*)_OnRspParkedOrderAction)(this, pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 报单操作请求响应    
    typedef void OnRspOrderActionType(void*, CThostFtdcInputOrderActionField *pInputOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOrderAction;
    virtual void OnRspOrderAction(CThostFtdcInputOrderActionField *pInputOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspOrderAction) {
			((OnRspOrderActionType*)_OnRspOrderAction)(this, pInputOrderAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 查询最大报单数量响应    
    typedef void OnRspQryMaxOrderVolumeType(void*, CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMaxOrderVolume;
    virtual void OnRspQryMaxOrderVolume(CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryMaxOrderVolume) {
			((OnRspQryMaxOrderVolumeType*)_OnRspQryMaxOrderVolume)(this, pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者结算结果确认响应    
    typedef void OnRspSettlementInfoConfirmType(void*, CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspSettlementInfoConfirm;
    virtual void OnRspSettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspSettlementInfoConfirm) {
			((OnRspSettlementInfoConfirmType*)_OnRspSettlementInfoConfirm)(this, pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 删除预埋单响应    
    typedef void OnRspRemoveParkedOrderType(void*, CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspRemoveParkedOrder;
    virtual void OnRspRemoveParkedOrder(CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspRemoveParkedOrder) {
			((OnRspRemoveParkedOrderType*)_OnRspRemoveParkedOrder)(this, pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 删除预埋撤单响应    
    typedef void OnRspRemoveParkedOrderActionType(void*, CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspRemoveParkedOrderAction;
    virtual void OnRspRemoveParkedOrderAction(CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspRemoveParkedOrderAction) {
			((OnRspRemoveParkedOrderActionType*)_OnRspRemoveParkedOrderAction)(this, pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 执行宣告录入请求响应    
    typedef void OnRspExecOrderInsertType(void*, CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspExecOrderInsert;
    virtual void OnRspExecOrderInsert(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspExecOrderInsert) {
			((OnRspExecOrderInsertType*)_OnRspExecOrderInsert)(this, pInputExecOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 执行宣告操作请求响应    
    typedef void OnRspExecOrderActionType(void*, CThostFtdcInputExecOrderActionField *pInputExecOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspExecOrderAction;
    virtual void OnRspExecOrderAction(CThostFtdcInputExecOrderActionField *pInputExecOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspExecOrderAction) {
			((OnRspExecOrderActionType*)_OnRspExecOrderAction)(this, pInputExecOrderAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 询价录入请求响应    
    typedef void OnRspForQuoteInsertType(void*, CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspForQuoteInsert;
    virtual void OnRspForQuoteInsert(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspForQuoteInsert) {
			((OnRspForQuoteInsertType*)_OnRspForQuoteInsert)(this, pInputForQuote, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 报价录入请求响应    
    typedef void OnRspQuoteInsertType(void*, CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQuoteInsert;
    virtual void OnRspQuoteInsert(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQuoteInsert) {
			((OnRspQuoteInsertType*)_OnRspQuoteInsert)(this, pInputQuote, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 报价操作请求响应    
    typedef void OnRspQuoteActionType(void*, CThostFtdcInputQuoteActionField *pInputQuoteAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQuoteAction;
    virtual void OnRspQuoteAction(CThostFtdcInputQuoteActionField *pInputQuoteAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQuoteAction) {
			((OnRspQuoteActionType*)_OnRspQuoteAction)(this, pInputQuoteAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 批量报单操作请求响应    
    typedef void OnRspBatchOrderActionType(void*, CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspBatchOrderAction;
    virtual void OnRspBatchOrderAction(CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspBatchOrderAction) {
			((OnRspBatchOrderActionType*)_OnRspBatchOrderAction)(this, pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 期权自对冲录入请求响应    
    typedef void OnRspOptionSelfCloseInsertType(void*, CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOptionSelfCloseInsert;
    virtual void OnRspOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspOptionSelfCloseInsert) {
			((OnRspOptionSelfCloseInsertType*)_OnRspOptionSelfCloseInsert)(this, pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 期权自对冲操作请求响应    
    typedef void OnRspOptionSelfCloseActionType(void*, CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOptionSelfCloseAction;
    virtual void OnRspOptionSelfCloseAction(CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspOptionSelfCloseAction) {
			((OnRspOptionSelfCloseActionType*)_OnRspOptionSelfCloseAction)(this, pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 申请组合录入请求响应    
    typedef void OnRspCombActionInsertType(void*, CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspCombActionInsert;
    virtual void OnRspCombActionInsert(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspCombActionInsert) {
			((OnRspCombActionInsertType*)_OnRspCombActionInsert)(this, pInputCombAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询报单响应    
    typedef void OnRspQryOrderType(void*, CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOrder;
    virtual void OnRspQryOrder(CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryOrder) {
			((OnRspQryOrderType*)_OnRspQryOrder)(this, pOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询成交响应    
    typedef void OnRspQryTradeType(void*, CThostFtdcTradeField *pTrade, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTrade;
    virtual void OnRspQryTrade(CThostFtdcTradeField *pTrade, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTrade) {
			((OnRspQryTradeType*)_OnRspQryTrade)(this, pTrade, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资者持仓响应    
    typedef void OnRspQryInvestorPositionType(void*, CThostFtdcInvestorPositionField *pInvestorPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPosition;
    virtual void OnRspQryInvestorPosition(CThostFtdcInvestorPositionField *pInvestorPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorPosition) {
			((OnRspQryInvestorPositionType*)_OnRspQryInvestorPosition)(this, pInvestorPosition, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询资金账户响应    
    typedef void OnRspQryTradingAccountType(void*, CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTradingAccount;
    virtual void OnRspQryTradingAccount(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTradingAccount) {
			((OnRspQryTradingAccountType*)_OnRspQryTradingAccount)(this, pTradingAccount, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资者响应    
    typedef void OnRspQryInvestorType(void*, CThostFtdcInvestorField *pInvestor, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestor;
    virtual void OnRspQryInvestor(CThostFtdcInvestorField *pInvestor, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestor) {
			((OnRspQryInvestorType*)_OnRspQryInvestor)(this, pInvestor, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询交易编码响应    
    typedef void OnRspQryTradingCodeType(void*, CThostFtdcTradingCodeField *pTradingCode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTradingCode;
    virtual void OnRspQryTradingCode(CThostFtdcTradingCodeField *pTradingCode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTradingCode) {
			((OnRspQryTradingCodeType*)_OnRspQryTradingCode)(this, pTradingCode, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询合约保证金率响应    
    typedef void OnRspQryInstrumentMarginRateType(void*, CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrumentMarginRate;
    virtual void OnRspQryInstrumentMarginRate(CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInstrumentMarginRate) {
			((OnRspQryInstrumentMarginRateType*)_OnRspQryInstrumentMarginRate)(this, pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询合约手续费率响应    
    typedef void OnRspQryInstrumentCommissionRateType(void*, CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrumentCommissionRate;
    virtual void OnRspQryInstrumentCommissionRate(CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInstrumentCommissionRate) {
			((OnRspQryInstrumentCommissionRateType*)_OnRspQryInstrumentCommissionRate)(this, pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询交易所响应    
    typedef void OnRspQryExchangeType(void*, CThostFtdcExchangeField *pExchange, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchange;
    virtual void OnRspQryExchange(CThostFtdcExchangeField *pExchange, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryExchange) {
			((OnRspQryExchangeType*)_OnRspQryExchange)(this, pExchange, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询产品响应    
    typedef void OnRspQryProductType(void*, CThostFtdcProductField *pProduct, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryProduct;
    virtual void OnRspQryProduct(CThostFtdcProductField *pProduct, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryProduct) {
			((OnRspQryProductType*)_OnRspQryProduct)(this, pProduct, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询合约响应    
    typedef void OnRspQryInstrumentType(void*, CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrument;
    virtual void OnRspQryInstrument(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInstrument) {
			((OnRspQryInstrumentType*)_OnRspQryInstrument)(this, pInstrument, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询行情响应    
    typedef void OnRspQryDepthMarketDataType(void*, CThostFtdcDepthMarketDataField *pDepthMarketData, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryDepthMarketData;
    virtual void OnRspQryDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryDepthMarketData) {
			((OnRspQryDepthMarketDataType*)_OnRspQryDepthMarketData)(this, pDepthMarketData, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询交易员报盘机响应    
    typedef void OnRspQryTraderOfferType(void*, CThostFtdcTraderOfferField *pTraderOffer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTraderOffer;
    virtual void OnRspQryTraderOffer(CThostFtdcTraderOfferField *pTraderOffer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTraderOffer) {
			((OnRspQryTraderOfferType*)_OnRspQryTraderOffer)(this, pTraderOffer, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资者结算结果响应    
    typedef void OnRspQrySettlementInfoType(void*, CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySettlementInfo;
    virtual void OnRspQrySettlementInfo(CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySettlementInfo) {
			((OnRspQrySettlementInfoType*)_OnRspQrySettlementInfo)(this, pSettlementInfo, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询转帐银行响应    
    typedef void OnRspQryTransferBankType(void*, CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTransferBank;
    virtual void OnRspQryTransferBank(CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTransferBank) {
			((OnRspQryTransferBankType*)_OnRspQryTransferBank)(this, pTransferBank, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资者持仓明细响应    
    typedef void OnRspQryInvestorPositionDetailType(void*, CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPositionDetail;
    virtual void OnRspQryInvestorPositionDetail(CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorPositionDetail) {
			((OnRspQryInvestorPositionDetailType*)_OnRspQryInvestorPositionDetail)(this, pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询客户通知响应    
    typedef void OnRspQryNoticeType(void*, CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryNotice;
    virtual void OnRspQryNotice(CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryNotice) {
			((OnRspQryNoticeType*)_OnRspQryNotice)(this, pNotice, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询结算信息确认响应    
    typedef void OnRspQrySettlementInfoConfirmType(void*, CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySettlementInfoConfirm;
    virtual void OnRspQrySettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySettlementInfoConfirm) {
			((OnRspQrySettlementInfoConfirmType*)_OnRspQrySettlementInfoConfirm)(this, pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资者持仓明细响应    
    typedef void OnRspQryInvestorPositionCombineDetailType(void*, CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPositionCombineDetail;
    virtual void OnRspQryInvestorPositionCombineDetail(CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorPositionCombineDetail) {
			((OnRspQryInvestorPositionCombineDetailType*)_OnRspQryInvestorPositionCombineDetail)(this, pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 查询保证金监管系统经纪公司资金账户密钥响应    
    typedef void OnRspQryCFMMCTradingAccountKeyType(void*, CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCFMMCTradingAccountKey;
    virtual void OnRspQryCFMMCTradingAccountKey(CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryCFMMCTradingAccountKey) {
			((OnRspQryCFMMCTradingAccountKeyType*)_OnRspQryCFMMCTradingAccountKey)(this, pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询仓单折抵信息响应    
    typedef void OnRspQryEWarrantOffsetType(void*, CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryEWarrantOffset;
    virtual void OnRspQryEWarrantOffset(CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryEWarrantOffset) {
			((OnRspQryEWarrantOffsetType*)_OnRspQryEWarrantOffset)(this, pEWarrantOffset, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资者品种/跨品种保证金响应    
    typedef void OnRspQryInvestorProductGroupMarginType(void*, CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorProductGroupMargin;
    virtual void OnRspQryInvestorProductGroupMargin(CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorProductGroupMargin) {
			((OnRspQryInvestorProductGroupMarginType*)_OnRspQryInvestorProductGroupMargin)(this, pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询交易所保证金率响应    
    typedef void OnRspQryExchangeMarginRateType(void*, CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchangeMarginRate;
    virtual void OnRspQryExchangeMarginRate(CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryExchangeMarginRate) {
			((OnRspQryExchangeMarginRateType*)_OnRspQryExchangeMarginRate)(this, pExchangeMarginRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询交易所调整保证金率响应    
    typedef void OnRspQryExchangeMarginRateAdjustType(void*, CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchangeMarginRateAdjust;
    virtual void OnRspQryExchangeMarginRateAdjust(CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryExchangeMarginRateAdjust) {
			((OnRspQryExchangeMarginRateAdjustType*)_OnRspQryExchangeMarginRateAdjust)(this, pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询汇率响应    
    typedef void OnRspQryExchangeRateType(void*, CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchangeRate;
    virtual void OnRspQryExchangeRate(CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryExchangeRate) {
			((OnRspQryExchangeRateType*)_OnRspQryExchangeRate)(this, pExchangeRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询二级代理操作员银期权限响应    
    typedef void OnRspQrySecAgentACIDMapType(void*, CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentACIDMap;
    virtual void OnRspQrySecAgentACIDMap(CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySecAgentACIDMap) {
			((OnRspQrySecAgentACIDMapType*)_OnRspQrySecAgentACIDMap)(this, pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询产品报价汇率    
    typedef void OnRspQryProductExchRateType(void*, CThostFtdcProductExchRateField *pProductExchRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryProductExchRate;
    virtual void OnRspQryProductExchRate(CThostFtdcProductExchRateField *pProductExchRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryProductExchRate) {
			((OnRspQryProductExchRateType*)_OnRspQryProductExchRate)(this, pProductExchRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询产品组    
    typedef void OnRspQryProductGroupType(void*, CThostFtdcProductGroupField *pProductGroup, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryProductGroup;
    virtual void OnRspQryProductGroup(CThostFtdcProductGroupField *pProductGroup, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryProductGroup) {
			((OnRspQryProductGroupType*)_OnRspQryProductGroup)(this, pProductGroup, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询做市商合约手续费率响应    
    typedef void OnRspQryMMInstrumentCommissionRateType(void*, CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMMInstrumentCommissionRate;
    virtual void OnRspQryMMInstrumentCommissionRate(CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryMMInstrumentCommissionRate) {
			((OnRspQryMMInstrumentCommissionRateType*)_OnRspQryMMInstrumentCommissionRate)(this, pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询做市商期权合约手续费响应    
    typedef void OnRspQryMMOptionInstrCommRateType(void*, CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMMOptionInstrCommRate;
    virtual void OnRspQryMMOptionInstrCommRate(CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryMMOptionInstrCommRate) {
			((OnRspQryMMOptionInstrCommRateType*)_OnRspQryMMOptionInstrCommRate)(this, pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询报单手续费响应    
    typedef void OnRspQryInstrumentOrderCommRateType(void*, CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrumentOrderCommRate;
    virtual void OnRspQryInstrumentOrderCommRate(CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInstrumentOrderCommRate) {
			((OnRspQryInstrumentOrderCommRateType*)_OnRspQryInstrumentOrderCommRate)(this, pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询资金账户响应    
    typedef void OnRspQrySecAgentTradingAccountType(void*, CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentTradingAccount;
    virtual void OnRspQrySecAgentTradingAccount(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySecAgentTradingAccount) {
			((OnRspQrySecAgentTradingAccountType*)_OnRspQrySecAgentTradingAccount)(this, pTradingAccount, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询二级代理商资金校验模式响应    
    typedef void OnRspQrySecAgentCheckModeType(void*, CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentCheckMode;
    virtual void OnRspQrySecAgentCheckMode(CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySecAgentCheckMode) {
			((OnRspQrySecAgentCheckModeType*)_OnRspQrySecAgentCheckMode)(this, pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询二级代理商信息响应    
    typedef void OnRspQrySecAgentTradeInfoType(void*, CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentTradeInfo;
    virtual void OnRspQrySecAgentTradeInfo(CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySecAgentTradeInfo) {
			((OnRspQrySecAgentTradeInfoType*)_OnRspQrySecAgentTradeInfo)(this, pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询期权交易成本响应    
    typedef void OnRspQryOptionInstrTradeCostType(void*, CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOptionInstrTradeCost;
    virtual void OnRspQryOptionInstrTradeCost(CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryOptionInstrTradeCost) {
			((OnRspQryOptionInstrTradeCostType*)_OnRspQryOptionInstrTradeCost)(this, pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询期权合约手续费响应    
    typedef void OnRspQryOptionInstrCommRateType(void*, CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOptionInstrCommRate;
    virtual void OnRspQryOptionInstrCommRate(CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryOptionInstrCommRate) {
			((OnRspQryOptionInstrCommRateType*)_OnRspQryOptionInstrCommRate)(this, pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询执行宣告响应    
    typedef void OnRspQryExecOrderType(void*, CThostFtdcExecOrderField *pExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExecOrder;
    virtual void OnRspQryExecOrder(CThostFtdcExecOrderField *pExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryExecOrder) {
			((OnRspQryExecOrderType*)_OnRspQryExecOrder)(this, pExecOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询询价响应    
    typedef void OnRspQryForQuoteType(void*, CThostFtdcForQuoteField *pForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryForQuote;
    virtual void OnRspQryForQuote(CThostFtdcForQuoteField *pForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryForQuote) {
			((OnRspQryForQuoteType*)_OnRspQryForQuote)(this, pForQuote, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询报价响应    
    typedef void OnRspQryQuoteType(void*, CThostFtdcQuoteField *pQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryQuote;
    virtual void OnRspQryQuote(CThostFtdcQuoteField *pQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryQuote) {
			((OnRspQryQuoteType*)_OnRspQryQuote)(this, pQuote, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询期权自对冲响应    
    typedef void OnRspQryOptionSelfCloseType(void*, CThostFtdcOptionSelfCloseField *pOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOptionSelfClose;
    virtual void OnRspQryOptionSelfClose(CThostFtdcOptionSelfCloseField *pOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryOptionSelfClose) {
			((OnRspQryOptionSelfCloseType*)_OnRspQryOptionSelfClose)(this, pOptionSelfClose, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询投资单元响应    
    typedef void OnRspQryInvestUnitType(void*, CThostFtdcInvestUnitField *pInvestUnit, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestUnit;
    virtual void OnRspQryInvestUnit(CThostFtdcInvestUnitField *pInvestUnit, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestUnit) {
			((OnRspQryInvestUnitType*)_OnRspQryInvestUnit)(this, pInvestUnit, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询组合合约安全系数响应    
    typedef void OnRspQryCombInstrumentGuardType(void*, CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCombInstrumentGuard;
    virtual void OnRspQryCombInstrumentGuard(CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryCombInstrumentGuard) {
			((OnRspQryCombInstrumentGuardType*)_OnRspQryCombInstrumentGuard)(this, pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询申请组合响应    
    typedef void OnRspQryCombActionType(void*, CThostFtdcCombActionField *pCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCombAction;
    virtual void OnRspQryCombAction(CThostFtdcCombActionField *pCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryCombAction) {
			((OnRspQryCombActionType*)_OnRspQryCombAction)(this, pCombAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询转帐流水响应    
    typedef void OnRspQryTransferSerialType(void*, CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTransferSerial;
    virtual void OnRspQryTransferSerial(CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTransferSerial) {
			((OnRspQryTransferSerialType*)_OnRspQryTransferSerial)(this, pTransferSerial, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询银期签约关系响应    
    typedef void OnRspQryAccountregisterType(void*, CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryAccountregister;
    virtual void OnRspQryAccountregister(CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryAccountregister) {
			((OnRspQryAccountregisterType*)_OnRspQryAccountregister)(this, pAccountregister, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 错误应答    
    typedef void OnRspErrorType(void*, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspError;
    virtual void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspError) {
			((OnRspErrorType*)_OnRspError)(this, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 报单通知    
    typedef void OnRtnOrderType(void*, CThostFtdcOrderField *pOrder);
    void *_OnRtnOrder;
    virtual void OnRtnOrder(CThostFtdcOrderField *pOrder){
        if (_OnRtnOrder) {
			((OnRtnOrderType*)_OnRtnOrder)(this, pOrder);
		}
    }
	// 成交通知    
    typedef void OnRtnTradeType(void*, CThostFtdcTradeField *pTrade);
    void *_OnRtnTrade;
    virtual void OnRtnTrade(CThostFtdcTradeField *pTrade){
        if (_OnRtnTrade) {
			((OnRtnTradeType*)_OnRtnTrade)(this, pTrade);
		}
    }
	// 报单录入错误回报    
    typedef void OnErrRtnOrderInsertType(void*, CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOrderInsert;
    virtual void OnErrRtnOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnOrderInsert) {
			((OnErrRtnOrderInsertType*)_OnErrRtnOrderInsert)(this, pInputOrder, pRspInfo);
		}
    }
	// 报单操作错误回报    
    typedef void OnErrRtnOrderActionType(void*, CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOrderAction;
    virtual void OnErrRtnOrderAction(CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnOrderAction) {
			((OnErrRtnOrderActionType*)_OnErrRtnOrderAction)(this, pOrderAction, pRspInfo);
		}
    }
	// 合约交易状态通知    
    typedef void OnRtnInstrumentStatusType(void*, CThostFtdcInstrumentStatusField *pInstrumentStatus);
    void *_OnRtnInstrumentStatus;
    virtual void OnRtnInstrumentStatus(CThostFtdcInstrumentStatusField *pInstrumentStatus){
        if (_OnRtnInstrumentStatus) {
			((OnRtnInstrumentStatusType*)_OnRtnInstrumentStatus)(this, pInstrumentStatus);
		}
    }
	// 交易所公告通知    
    typedef void OnRtnBulletinType(void*, CThostFtdcBulletinField *pBulletin);
    void *_OnRtnBulletin;
    virtual void OnRtnBulletin(CThostFtdcBulletinField *pBulletin){
        if (_OnRtnBulletin) {
			((OnRtnBulletinType*)_OnRtnBulletin)(this, pBulletin);
		}
    }
	// 交易通知    
    typedef void OnRtnTradingNoticeType(void*, CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);
    void *_OnRtnTradingNotice;
    virtual void OnRtnTradingNotice(CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo){
        if (_OnRtnTradingNotice) {
			((OnRtnTradingNoticeType*)_OnRtnTradingNotice)(this, pTradingNoticeInfo);
		}
    }
	// 提示条件单校验错误    
    typedef void OnRtnErrorConditionalOrderType(void*, CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);
    void *_OnRtnErrorConditionalOrder;
    virtual void OnRtnErrorConditionalOrder(CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder){
        if (_OnRtnErrorConditionalOrder) {
			((OnRtnErrorConditionalOrderType*)_OnRtnErrorConditionalOrder)(this, pErrorConditionalOrder);
		}
    }
	// 执行宣告通知    
    typedef void OnRtnExecOrderType(void*, CThostFtdcExecOrderField *pExecOrder);
    void *_OnRtnExecOrder;
    virtual void OnRtnExecOrder(CThostFtdcExecOrderField *pExecOrder){
        if (_OnRtnExecOrder) {
			((OnRtnExecOrderType*)_OnRtnExecOrder)(this, pExecOrder);
		}
    }
	// 执行宣告录入错误回报    
    typedef void OnErrRtnExecOrderInsertType(void*, CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnExecOrderInsert;
    virtual void OnErrRtnExecOrderInsert(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnExecOrderInsert) {
			((OnErrRtnExecOrderInsertType*)_OnErrRtnExecOrderInsert)(this, pInputExecOrder, pRspInfo);
		}
    }
	// 执行宣告操作错误回报    
    typedef void OnErrRtnExecOrderActionType(void*, CThostFtdcExecOrderActionField *pExecOrderAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnExecOrderAction;
    virtual void OnErrRtnExecOrderAction(CThostFtdcExecOrderActionField *pExecOrderAction, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnExecOrderAction) {
			((OnErrRtnExecOrderActionType*)_OnErrRtnExecOrderAction)(this, pExecOrderAction, pRspInfo);
		}
    }
	// 询价录入错误回报    
    typedef void OnErrRtnForQuoteInsertType(void*, CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnForQuoteInsert;
    virtual void OnErrRtnForQuoteInsert(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnForQuoteInsert) {
			((OnErrRtnForQuoteInsertType*)_OnErrRtnForQuoteInsert)(this, pInputForQuote, pRspInfo);
		}
    }
	// 报价通知    
    typedef void OnRtnQuoteType(void*, CThostFtdcQuoteField *pQuote);
    void *_OnRtnQuote;
    virtual void OnRtnQuote(CThostFtdcQuoteField *pQuote){
        if (_OnRtnQuote) {
			((OnRtnQuoteType*)_OnRtnQuote)(this, pQuote);
		}
    }
	// 报价录入错误回报    
    typedef void OnErrRtnQuoteInsertType(void*, CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnQuoteInsert;
    virtual void OnErrRtnQuoteInsert(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnQuoteInsert) {
			((OnErrRtnQuoteInsertType*)_OnErrRtnQuoteInsert)(this, pInputQuote, pRspInfo);
		}
    }
	// 报价操作错误回报    
    typedef void OnErrRtnQuoteActionType(void*, CThostFtdcQuoteActionField *pQuoteAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnQuoteAction;
    virtual void OnErrRtnQuoteAction(CThostFtdcQuoteActionField *pQuoteAction, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnQuoteAction) {
			((OnErrRtnQuoteActionType*)_OnErrRtnQuoteAction)(this, pQuoteAction, pRspInfo);
		}
    }
	// 询价通知    
    typedef void OnRtnForQuoteRspType(void*, CThostFtdcForQuoteRspField *pForQuoteRsp);
    void *_OnRtnForQuoteRsp;
    virtual void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp){
        if (_OnRtnForQuoteRsp) {
			((OnRtnForQuoteRspType*)_OnRtnForQuoteRsp)(this, pForQuoteRsp);
		}
    }
	// 保证金监控中心用户令牌    
    typedef void OnRtnCFMMCTradingAccountTokenType(void*, CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);
    void *_OnRtnCFMMCTradingAccountToken;
    virtual void OnRtnCFMMCTradingAccountToken(CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken){
        if (_OnRtnCFMMCTradingAccountToken) {
			((OnRtnCFMMCTradingAccountTokenType*)_OnRtnCFMMCTradingAccountToken)(this, pCFMMCTradingAccountToken);
		}
    }
	// 批量报单操作错误回报    
    typedef void OnErrRtnBatchOrderActionType(void*, CThostFtdcBatchOrderActionField *pBatchOrderAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnBatchOrderAction;
    virtual void OnErrRtnBatchOrderAction(CThostFtdcBatchOrderActionField *pBatchOrderAction, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnBatchOrderAction) {
			((OnErrRtnBatchOrderActionType*)_OnErrRtnBatchOrderAction)(this, pBatchOrderAction, pRspInfo);
		}
    }
	// 期权自对冲通知    
    typedef void OnRtnOptionSelfCloseType(void*, CThostFtdcOptionSelfCloseField *pOptionSelfClose);
    void *_OnRtnOptionSelfClose;
    virtual void OnRtnOptionSelfClose(CThostFtdcOptionSelfCloseField *pOptionSelfClose){
        if (_OnRtnOptionSelfClose) {
			((OnRtnOptionSelfCloseType*)_OnRtnOptionSelfClose)(this, pOptionSelfClose);
		}
    }
	// 期权自对冲录入错误回报    
    typedef void OnErrRtnOptionSelfCloseInsertType(void*, CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOptionSelfCloseInsert;
    virtual void OnErrRtnOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnOptionSelfCloseInsert) {
			((OnErrRtnOptionSelfCloseInsertType*)_OnErrRtnOptionSelfCloseInsert)(this, pInputOptionSelfClose, pRspInfo);
		}
    }
	// 期权自对冲操作错误回报    
    typedef void OnErrRtnOptionSelfCloseActionType(void*, CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOptionSelfCloseAction;
    virtual void OnErrRtnOptionSelfCloseAction(CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnOptionSelfCloseAction) {
			((OnErrRtnOptionSelfCloseActionType*)_OnErrRtnOptionSelfCloseAction)(this, pOptionSelfCloseAction, pRspInfo);
		}
    }
	// 申请组合通知    
    typedef void OnRtnCombActionType(void*, CThostFtdcCombActionField *pCombAction);
    void *_OnRtnCombAction;
    virtual void OnRtnCombAction(CThostFtdcCombActionField *pCombAction){
        if (_OnRtnCombAction) {
			((OnRtnCombActionType*)_OnRtnCombAction)(this, pCombAction);
		}
    }
	// 申请组合录入错误回报    
    typedef void OnErrRtnCombActionInsertType(void*, CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnCombActionInsert;
    virtual void OnErrRtnCombActionInsert(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnCombActionInsert) {
			((OnErrRtnCombActionInsertType*)_OnErrRtnCombActionInsert)(this, pInputCombAction, pRspInfo);
		}
    }
	// 请求查询签约银行响应    
    typedef void OnRspQryContractBankType(void*, CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryContractBank;
    virtual void OnRspQryContractBank(CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryContractBank) {
			((OnRspQryContractBankType*)_OnRspQryContractBank)(this, pContractBank, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询预埋单响应    
    typedef void OnRspQryParkedOrderType(void*, CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryParkedOrder;
    virtual void OnRspQryParkedOrder(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryParkedOrder) {
			((OnRspQryParkedOrderType*)_OnRspQryParkedOrder)(this, pParkedOrder, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询预埋撤单响应    
    typedef void OnRspQryParkedOrderActionType(void*, CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryParkedOrderAction;
    virtual void OnRspQryParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryParkedOrderAction) {
			((OnRspQryParkedOrderActionType*)_OnRspQryParkedOrderAction)(this, pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询交易通知响应    
    typedef void OnRspQryTradingNoticeType(void*, CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTradingNotice;
    virtual void OnRspQryTradingNotice(CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryTradingNotice) {
			((OnRspQryTradingNoticeType*)_OnRspQryTradingNotice)(this, pTradingNotice, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询经纪公司交易参数响应    
    typedef void OnRspQryBrokerTradingParamsType(void*, CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryBrokerTradingParams;
    virtual void OnRspQryBrokerTradingParams(CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryBrokerTradingParams) {
			((OnRspQryBrokerTradingParamsType*)_OnRspQryBrokerTradingParams)(this, pBrokerTradingParams, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询经纪公司交易算法响应    
    typedef void OnRspQryBrokerTradingAlgosType(void*, CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryBrokerTradingAlgos;
    virtual void OnRspQryBrokerTradingAlgos(CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryBrokerTradingAlgos) {
			((OnRspQryBrokerTradingAlgosType*)_OnRspQryBrokerTradingAlgos)(this, pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求查询监控中心用户令牌    
    typedef void OnRspQueryCFMMCTradingAccountTokenType(void*, CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQueryCFMMCTradingAccountToken;
    virtual void OnRspQueryCFMMCTradingAccountToken(CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQueryCFMMCTradingAccountToken) {
			((OnRspQueryCFMMCTradingAccountTokenType*)_OnRspQueryCFMMCTradingAccountToken)(this, pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 银行发起银行资金转期货通知    
    typedef void OnRtnFromBankToFutureByBankType(void*, CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromBankToFutureByBank;
    virtual void OnRtnFromBankToFutureByBank(CThostFtdcRspTransferField *pRspTransfer){
        if (_OnRtnFromBankToFutureByBank) {
			((OnRtnFromBankToFutureByBankType*)_OnRtnFromBankToFutureByBank)(this, pRspTransfer);
		}
    }
	// 银行发起期货资金转银行通知    
    typedef void OnRtnFromFutureToBankByBankType(void*, CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromFutureToBankByBank;
    virtual void OnRtnFromFutureToBankByBank(CThostFtdcRspTransferField *pRspTransfer){
        if (_OnRtnFromFutureToBankByBank) {
			((OnRtnFromFutureToBankByBankType*)_OnRtnFromFutureToBankByBank)(this, pRspTransfer);
		}
    }
	// 银行发起冲正银行转期货通知    
    typedef void OnRtnRepealFromBankToFutureByBankType(void*, CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromBankToFutureByBank;
    virtual void OnRtnRepealFromBankToFutureByBank(CThostFtdcRspRepealField *pRspRepeal){
        if (_OnRtnRepealFromBankToFutureByBank) {
			((OnRtnRepealFromBankToFutureByBankType*)_OnRtnRepealFromBankToFutureByBank)(this, pRspRepeal);
		}
    }
	// 银行发起冲正期货转银行通知    
    typedef void OnRtnRepealFromFutureToBankByBankType(void*, CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromFutureToBankByBank;
    virtual void OnRtnRepealFromFutureToBankByBank(CThostFtdcRspRepealField *pRspRepeal){
        if (_OnRtnRepealFromFutureToBankByBank) {
			((OnRtnRepealFromFutureToBankByBankType*)_OnRtnRepealFromFutureToBankByBank)(this, pRspRepeal);
		}
    }
	// 期货发起银行资金转期货通知    
    typedef void OnRtnFromBankToFutureByFutureType(void*, CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromBankToFutureByFuture;
    virtual void OnRtnFromBankToFutureByFuture(CThostFtdcRspTransferField *pRspTransfer){
        if (_OnRtnFromBankToFutureByFuture) {
			((OnRtnFromBankToFutureByFutureType*)_OnRtnFromBankToFutureByFuture)(this, pRspTransfer);
		}
    }
	// 期货发起期货资金转银行通知    
    typedef void OnRtnFromFutureToBankByFutureType(void*, CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromFutureToBankByFuture;
    virtual void OnRtnFromFutureToBankByFuture(CThostFtdcRspTransferField *pRspTransfer){
        if (_OnRtnFromFutureToBankByFuture) {
			((OnRtnFromFutureToBankByFutureType*)_OnRtnFromFutureToBankByFuture)(this, pRspTransfer);
		}
    }
	// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知    
    typedef void OnRtnRepealFromBankToFutureByFutureManualType(void*, CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromBankToFutureByFutureManual;
    virtual void OnRtnRepealFromBankToFutureByFutureManual(CThostFtdcRspRepealField *pRspRepeal){
        if (_OnRtnRepealFromBankToFutureByFutureManual) {
			((OnRtnRepealFromBankToFutureByFutureManualType*)_OnRtnRepealFromBankToFutureByFutureManual)(this, pRspRepeal);
		}
    }
	// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知    
    typedef void OnRtnRepealFromFutureToBankByFutureManualType(void*, CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromFutureToBankByFutureManual;
    virtual void OnRtnRepealFromFutureToBankByFutureManual(CThostFtdcRspRepealField *pRspRepeal){
        if (_OnRtnRepealFromFutureToBankByFutureManual) {
			((OnRtnRepealFromFutureToBankByFutureManualType*)_OnRtnRepealFromFutureToBankByFutureManual)(this, pRspRepeal);
		}
    }
	// 期货发起查询银行余额通知    
    typedef void OnRtnQueryBankBalanceByFutureType(void*, CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);
    void *_OnRtnQueryBankBalanceByFuture;
    virtual void OnRtnQueryBankBalanceByFuture(CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount){
        if (_OnRtnQueryBankBalanceByFuture) {
			((OnRtnQueryBankBalanceByFutureType*)_OnRtnQueryBankBalanceByFuture)(this, pNotifyQueryAccount);
		}
    }
	// 期货发起银行资金转期货错误回报    
    typedef void OnErrRtnBankToFutureByFutureType(void*, CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnBankToFutureByFuture;
    virtual void OnErrRtnBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnBankToFutureByFuture) {
			((OnErrRtnBankToFutureByFutureType*)_OnErrRtnBankToFutureByFuture)(this, pReqTransfer, pRspInfo);
		}
    }
	// 期货发起期货资金转银行错误回报    
    typedef void OnErrRtnFutureToBankByFutureType(void*, CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnFutureToBankByFuture;
    virtual void OnErrRtnFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnFutureToBankByFuture) {
			((OnErrRtnFutureToBankByFutureType*)_OnErrRtnFutureToBankByFuture)(this, pReqTransfer, pRspInfo);
		}
    }
	// 系统运行时期货端手工发起冲正银行转期货错误回报    
    typedef void OnErrRtnRepealBankToFutureByFutureManualType(void*, CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnRepealBankToFutureByFutureManual;
    virtual void OnErrRtnRepealBankToFutureByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnRepealBankToFutureByFutureManual) {
			((OnErrRtnRepealBankToFutureByFutureManualType*)_OnErrRtnRepealBankToFutureByFutureManual)(this, pReqRepeal, pRspInfo);
		}
    }
	// 系统运行时期货端手工发起冲正期货转银行错误回报    
    typedef void OnErrRtnRepealFutureToBankByFutureManualType(void*, CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnRepealFutureToBankByFutureManual;
    virtual void OnErrRtnRepealFutureToBankByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnRepealFutureToBankByFutureManual) {
			((OnErrRtnRepealFutureToBankByFutureManualType*)_OnErrRtnRepealFutureToBankByFutureManual)(this, pReqRepeal, pRspInfo);
		}
    }
	// 期货发起查询银行余额错误回报    
    typedef void OnErrRtnQueryBankBalanceByFutureType(void*, CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnQueryBankBalanceByFuture;
    virtual void OnErrRtnQueryBankBalanceByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo){
        if (_OnErrRtnQueryBankBalanceByFuture) {
			((OnErrRtnQueryBankBalanceByFutureType*)_OnErrRtnQueryBankBalanceByFuture)(this, pReqQueryAccount, pRspInfo);
		}
    }
	// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知    
    typedef void OnRtnRepealFromBankToFutureByFutureType(void*, CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromBankToFutureByFuture;
    virtual void OnRtnRepealFromBankToFutureByFuture(CThostFtdcRspRepealField *pRspRepeal){
        if (_OnRtnRepealFromBankToFutureByFuture) {
			((OnRtnRepealFromBankToFutureByFutureType*)_OnRtnRepealFromBankToFutureByFuture)(this, pRspRepeal);
		}
    }
	// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知    
    typedef void OnRtnRepealFromFutureToBankByFutureType(void*, CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromFutureToBankByFuture;
    virtual void OnRtnRepealFromFutureToBankByFuture(CThostFtdcRspRepealField *pRspRepeal){
        if (_OnRtnRepealFromFutureToBankByFuture) {
			((OnRtnRepealFromFutureToBankByFutureType*)_OnRtnRepealFromFutureToBankByFuture)(this, pRspRepeal);
		}
    }
	// 期货发起银行资金转期货应答    
    typedef void OnRspFromBankToFutureByFutureType(void*, CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspFromBankToFutureByFuture;
    virtual void OnRspFromBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspFromBankToFutureByFuture) {
			((OnRspFromBankToFutureByFutureType*)_OnRspFromBankToFutureByFuture)(this, pReqTransfer, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 期货发起期货资金转银行应答    
    typedef void OnRspFromFutureToBankByFutureType(void*, CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspFromFutureToBankByFuture;
    virtual void OnRspFromFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspFromFutureToBankByFuture) {
			((OnRspFromFutureToBankByFutureType*)_OnRspFromFutureToBankByFuture)(this, pReqTransfer, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 期货发起查询银行余额应答    
    typedef void OnRspQueryBankAccountMoneyByFutureType(void*, CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQueryBankAccountMoneyByFuture;
    virtual void OnRspQueryBankAccountMoneyByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQueryBankAccountMoneyByFuture) {
			((OnRspQueryBankAccountMoneyByFutureType*)_OnRspQueryBankAccountMoneyByFuture)(this, pReqQueryAccount, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 银行发起银期开户通知    
    typedef void OnRtnOpenAccountByBankType(void*, CThostFtdcOpenAccountField *pOpenAccount);
    void *_OnRtnOpenAccountByBank;
    virtual void OnRtnOpenAccountByBank(CThostFtdcOpenAccountField *pOpenAccount){
        if (_OnRtnOpenAccountByBank) {
			((OnRtnOpenAccountByBankType*)_OnRtnOpenAccountByBank)(this, pOpenAccount);
		}
    }
	// 银行发起银期销户通知    
    typedef void OnRtnCancelAccountByBankType(void*, CThostFtdcCancelAccountField *pCancelAccount);
    void *_OnRtnCancelAccountByBank;
    virtual void OnRtnCancelAccountByBank(CThostFtdcCancelAccountField *pCancelAccount){
        if (_OnRtnCancelAccountByBank) {
			((OnRtnCancelAccountByBankType*)_OnRtnCancelAccountByBank)(this, pCancelAccount);
		}
    }
	// 银行发起变更银行账号通知    
    typedef void OnRtnChangeAccountByBankType(void*, CThostFtdcChangeAccountField *pChangeAccount);
    void *_OnRtnChangeAccountByBank;
    virtual void OnRtnChangeAccountByBank(CThostFtdcChangeAccountField *pChangeAccount){
        if (_OnRtnChangeAccountByBank) {
			((OnRtnChangeAccountByBankType*)_OnRtnChangeAccountByBank)(this, pChangeAccount);
		}
    }
	// 请求查询分类合约响应    
    typedef void OnRspQryClassifiedInstrumentType(void*, CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryClassifiedInstrument;
    virtual void OnRspQryClassifiedInstrument(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryClassifiedInstrument) {
			((OnRspQryClassifiedInstrumentType*)_OnRspQryClassifiedInstrument)(this, pInstrument, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 请求组合优惠比例响应    
    typedef void OnRspQryCombPromotionParamType(void*, CThostFtdcCombPromotionParamField *pCombPromotionParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCombPromotionParam;
    virtual void OnRspQryCombPromotionParam(CThostFtdcCombPromotionParamField *pCombPromotionParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryCombPromotionParam) {
			((OnRspQryCombPromotionParamType*)_OnRspQryCombPromotionParam)(this, pCombPromotionParam, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者风险结算持仓查询响应    
    typedef void OnRspQryRiskSettleInvstPositionType(void*, CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRiskSettleInvstPosition;
    virtual void OnRspQryRiskSettleInvstPosition(CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRiskSettleInvstPosition) {
			((OnRspQryRiskSettleInvstPositionType*)_OnRspQryRiskSettleInvstPosition)(this, pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 风险结算产品查询响应    
    typedef void OnRspQryRiskSettleProductStatusType(void*, CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRiskSettleProductStatus;
    virtual void OnRspQryRiskSettleProductStatus(CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRiskSettleProductStatus) {
			((OnRspQryRiskSettleProductStatusType*)_OnRspQryRiskSettleProductStatus)(this, pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPBM期货合约参数查询响应    
    typedef void OnRspQrySPBMFutureParameterType(void*, CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMFutureParameter;
    virtual void OnRspQrySPBMFutureParameter(CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMFutureParameter) {
			((OnRspQrySPBMFutureParameterType*)_OnRspQrySPBMFutureParameter)(this, pSPBMFutureParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPBM期权合约参数查询响应    
    typedef void OnRspQrySPBMOptionParameterType(void*, CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMOptionParameter;
    virtual void OnRspQrySPBMOptionParameter(CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMOptionParameter) {
			((OnRspQrySPBMOptionParameterType*)_OnRspQrySPBMOptionParameter)(this, pSPBMOptionParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPBM品种内对锁仓折扣参数查询响应    
    typedef void OnRspQrySPBMIntraParameterType(void*, CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMIntraParameter;
    virtual void OnRspQrySPBMIntraParameter(CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMIntraParameter) {
			((OnRspQrySPBMIntraParameterType*)_OnRspQrySPBMIntraParameter)(this, pSPBMIntraParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPBM跨品种抵扣参数查询响应    
    typedef void OnRspQrySPBMInterParameterType(void*, CThostFtdcSPBMInterParameterField *pSPBMInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMInterParameter;
    virtual void OnRspQrySPBMInterParameter(CThostFtdcSPBMInterParameterField *pSPBMInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMInterParameter) {
			((OnRspQrySPBMInterParameterType*)_OnRspQrySPBMInterParameter)(this, pSPBMInterParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPBM组合保证金套餐查询响应    
    typedef void OnRspQrySPBMPortfDefinitionType(void*, CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMPortfDefinition;
    virtual void OnRspQrySPBMPortfDefinition(CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMPortfDefinition) {
			((OnRspQrySPBMPortfDefinitionType*)_OnRspQrySPBMPortfDefinition)(this, pSPBMPortfDefinition, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者SPBM套餐选择查询响应    
    typedef void OnRspQrySPBMInvestorPortfDefType(void*, CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMInvestorPortfDef;
    virtual void OnRspQrySPBMInvestorPortfDef(CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMInvestorPortfDef) {
			((OnRspQrySPBMInvestorPortfDefType*)_OnRspQrySPBMInvestorPortfDef)(this, pSPBMInvestorPortfDef, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者新型组合保证金系数查询响应    
    typedef void OnRspQryInvestorPortfMarginRatioType(void*, CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPortfMarginRatio;
    virtual void OnRspQryInvestorPortfMarginRatio(CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorPortfMarginRatio) {
			((OnRspQryInvestorPortfMarginRatioType*)_OnRspQryInvestorPortfMarginRatio)(this, pInvestorPortfMarginRatio, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者产品SPBM明细查询响应    
    typedef void OnRspQryInvestorProdSPBMDetailType(void*, CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorProdSPBMDetail;
    virtual void OnRspQryInvestorProdSPBMDetail(CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorProdSPBMDetail) {
			((OnRspQryInvestorProdSPBMDetailType*)_OnRspQryInvestorProdSPBMDetail)(this, pInvestorProdSPBMDetail, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者商品组SPMM记录查询响应    
    typedef void OnRspQryInvestorCommoditySPMMMarginType(void*, CThostFtdcInvestorCommoditySPMMMarginField *pInvestorCommoditySPMMMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorCommoditySPMMMargin;
    virtual void OnRspQryInvestorCommoditySPMMMargin(CThostFtdcInvestorCommoditySPMMMarginField *pInvestorCommoditySPMMMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorCommoditySPMMMargin) {
			((OnRspQryInvestorCommoditySPMMMarginType*)_OnRspQryInvestorCommoditySPMMMargin)(this, pInvestorCommoditySPMMMargin, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者商品群SPMM记录查询响应    
    typedef void OnRspQryInvestorCommodityGroupSPMMMarginType(void*, CThostFtdcInvestorCommodityGroupSPMMMarginField *pInvestorCommodityGroupSPMMMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorCommodityGroupSPMMMargin;
    virtual void OnRspQryInvestorCommodityGroupSPMMMargin(CThostFtdcInvestorCommodityGroupSPMMMarginField *pInvestorCommodityGroupSPMMMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorCommodityGroupSPMMMargin) {
			((OnRspQryInvestorCommodityGroupSPMMMarginType*)_OnRspQryInvestorCommodityGroupSPMMMargin)(this, pInvestorCommodityGroupSPMMMargin, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPMM合约参数查询响应    
    typedef void OnRspQrySPMMInstParamType(void*, CThostFtdcSPMMInstParamField *pSPMMInstParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPMMInstParam;
    virtual void OnRspQrySPMMInstParam(CThostFtdcSPMMInstParamField *pSPMMInstParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPMMInstParam) {
			((OnRspQrySPMMInstParamType*)_OnRspQrySPMMInstParam)(this, pSPMMInstParam, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPMM产品参数查询响应    
    typedef void OnRspQrySPMMProductParamType(void*, CThostFtdcSPMMProductParamField *pSPMMProductParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPMMProductParam;
    virtual void OnRspQrySPMMProductParam(CThostFtdcSPMMProductParamField *pSPMMProductParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPMMProductParam) {
			((OnRspQrySPMMProductParamType*)_OnRspQrySPMMProductParam)(this, pSPMMProductParam, pRspInfo, nRequestID, bIsLast);
		}
    }
	// SPBM附加跨品种抵扣参数查询响应    
    typedef void OnRspQrySPBMAddOnInterParameterType(void*, CThostFtdcSPBMAddOnInterParameterField *pSPBMAddOnInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMAddOnInterParameter;
    virtual void OnRspQrySPBMAddOnInterParameter(CThostFtdcSPBMAddOnInterParameterField *pSPBMAddOnInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQrySPBMAddOnInterParameter) {
			((OnRspQrySPBMAddOnInterParameterType*)_OnRspQrySPBMAddOnInterParameter)(this, pSPBMAddOnInterParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RCAMS产品组合信息查询响应    
    typedef void OnRspQryRCAMSCombProductInfoType(void*, CThostFtdcRCAMSCombProductInfoField *pRCAMSCombProductInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRCAMSCombProductInfo;
    virtual void OnRspQryRCAMSCombProductInfo(CThostFtdcRCAMSCombProductInfoField *pRCAMSCombProductInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRCAMSCombProductInfo) {
			((OnRspQryRCAMSCombProductInfoType*)_OnRspQryRCAMSCombProductInfo)(this, pRCAMSCombProductInfo, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RCAMS同合约风险对冲参数查询响应    
    typedef void OnRspQryRCAMSInstrParameterType(void*, CThostFtdcRCAMSInstrParameterField *pRCAMSInstrParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRCAMSInstrParameter;
    virtual void OnRspQryRCAMSInstrParameter(CThostFtdcRCAMSInstrParameterField *pRCAMSInstrParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRCAMSInstrParameter) {
			((OnRspQryRCAMSInstrParameterType*)_OnRspQryRCAMSInstrParameter)(this, pRCAMSInstrParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RCAMS品种内风险对冲参数查询响应    
    typedef void OnRspQryRCAMSIntraParameterType(void*, CThostFtdcRCAMSIntraParameterField *pRCAMSIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRCAMSIntraParameter;
    virtual void OnRspQryRCAMSIntraParameter(CThostFtdcRCAMSIntraParameterField *pRCAMSIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRCAMSIntraParameter) {
			((OnRspQryRCAMSIntraParameterType*)_OnRspQryRCAMSIntraParameter)(this, pRCAMSIntraParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RCAMS跨品种风险折抵参数查询响应    
    typedef void OnRspQryRCAMSInterParameterType(void*, CThostFtdcRCAMSInterParameterField *pRCAMSInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRCAMSInterParameter;
    virtual void OnRspQryRCAMSInterParameter(CThostFtdcRCAMSInterParameterField *pRCAMSInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRCAMSInterParameter) {
			((OnRspQryRCAMSInterParameterType*)_OnRspQryRCAMSInterParameter)(this, pRCAMSInterParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RCAMS空头期权风险调整参数查询响应    
    typedef void OnRspQryRCAMSShortOptAdjustParamType(void*, CThostFtdcRCAMSShortOptAdjustParamField *pRCAMSShortOptAdjustParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRCAMSShortOptAdjustParam;
    virtual void OnRspQryRCAMSShortOptAdjustParam(CThostFtdcRCAMSShortOptAdjustParamField *pRCAMSShortOptAdjustParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRCAMSShortOptAdjustParam) {
			((OnRspQryRCAMSShortOptAdjustParamType*)_OnRspQryRCAMSShortOptAdjustParam)(this, pRCAMSShortOptAdjustParam, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RCAMS策略组合持仓查询响应    
    typedef void OnRspQryRCAMSInvestorCombPositionType(void*, CThostFtdcRCAMSInvestorCombPositionField *pRCAMSInvestorCombPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRCAMSInvestorCombPosition;
    virtual void OnRspQryRCAMSInvestorCombPosition(CThostFtdcRCAMSInvestorCombPositionField *pRCAMSInvestorCombPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRCAMSInvestorCombPosition) {
			((OnRspQryRCAMSInvestorCombPositionType*)_OnRspQryRCAMSInvestorCombPosition)(this, pRCAMSInvestorCombPosition, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者品种RCAMS保证金查询响应    
    typedef void OnRspQryInvestorProdRCAMSMarginType(void*, CThostFtdcInvestorProdRCAMSMarginField *pInvestorProdRCAMSMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorProdRCAMSMargin;
    virtual void OnRspQryInvestorProdRCAMSMargin(CThostFtdcInvestorProdRCAMSMarginField *pInvestorProdRCAMSMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorProdRCAMSMargin) {
			((OnRspQryInvestorProdRCAMSMarginType*)_OnRspQryInvestorProdRCAMSMargin)(this, pInvestorProdRCAMSMargin, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RULE合约保证金参数查询响应    
    typedef void OnRspQryRULEInstrParameterType(void*, CThostFtdcRULEInstrParameterField *pRULEInstrParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRULEInstrParameter;
    virtual void OnRspQryRULEInstrParameter(CThostFtdcRULEInstrParameterField *pRULEInstrParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRULEInstrParameter) {
			((OnRspQryRULEInstrParameterType*)_OnRspQryRULEInstrParameter)(this, pRULEInstrParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RULE品种内对锁仓折扣参数查询响应    
    typedef void OnRspQryRULEIntraParameterType(void*, CThostFtdcRULEIntraParameterField *pRULEIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRULEIntraParameter;
    virtual void OnRspQryRULEIntraParameter(CThostFtdcRULEIntraParameterField *pRULEIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRULEIntraParameter) {
			((OnRspQryRULEIntraParameterType*)_OnRspQryRULEIntraParameter)(this, pRULEIntraParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// RULE跨品种抵扣参数查询响应    
    typedef void OnRspQryRULEInterParameterType(void*, CThostFtdcRULEInterParameterField *pRULEInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRULEInterParameter;
    virtual void OnRspQryRULEInterParameter(CThostFtdcRULEInterParameterField *pRULEInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryRULEInterParameter) {
			((OnRspQryRULEInterParameterType*)_OnRspQryRULEInterParameter)(this, pRULEInterParameter, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 投资者产品RULE保证金查询响应    
    typedef void OnRspQryInvestorProdRULEMarginType(void*, CThostFtdcInvestorProdRULEMarginField *pInvestorProdRULEMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorProdRULEMargin;
    virtual void OnRspQryInvestorProdRULEMargin(CThostFtdcInvestorProdRULEMarginField *pInvestorProdRULEMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryInvestorProdRULEMargin) {
			((OnRspQryInvestorProdRULEMarginType*)_OnRspQryInvestorProdRULEMargin)(this, pInvestorProdRULEMargin, pRspInfo, nRequestID, bIsLast);
		}
    }
	
};