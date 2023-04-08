#include "../CTPv6.6.9_20220922/ThostFtdcTraderApi.h"

class Trade : CThostFtdcTraderSpi
{
public:
    // //////////////////////////////////////////////////////////////////////
    typedef void OnFrontConnectedType();
    void *_OnFrontConnected;
    virtual void OnFrontConnected()
    {
        if (_OnFrontConnected)
        {
            ((OnFrontConnectedType *)_OnFrontConnected)();
        }
    }
    // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    typedef void OnFrontDisconnectedType(int nReason);
    void *_OnFrontDisconnected;
    virtual void OnFrontDisconnected(int nReason)
    {
        if (_OnFrontDisconnected)
        {
            ((OnFrontDisconnectedType *)_OnFrontDisconnected)(nReason);
        }
    }
    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    typedef void OnHeartBeatWarningType(int nTimeLapse);
    void *_OnHeartBeatWarning;
    virtual void OnHeartBeatWarning(int nTimeLapse)
    {
        if (_OnHeartBeatWarning)
        {
            ((OnHeartBeatWarningType *)_OnHeartBeatWarning)(nTimeLapse);
        }
    }
    // 客户端认证响应
    typedef void OnRspAuthenticateType(CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspAuthenticate;
    virtual void OnRspAuthenticate(CThostFtdcRspAuthenticateField *pRspAuthenticateField, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspAuthenticate)
        {
            ((OnRspAuthenticateType *)_OnRspAuthenticate)(pRspAuthenticateField, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 登录请求响应
    typedef void OnRspUserLoginType(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserLogin;
    virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspUserLogin)
        {
            ((OnRspUserLoginType *)_OnRspUserLogin)(pRspUserLogin, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 登出请求响应
    typedef void OnRspUserLogoutType(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserLogout;
    virtual void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspUserLogout)
        {
            ((OnRspUserLogoutType *)_OnRspUserLogout)(pUserLogout, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 用户口令更新请求响应
    typedef void OnRspUserPasswordUpdateType(CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserPasswordUpdate;
    virtual void OnRspUserPasswordUpdate(CThostFtdcUserPasswordUpdateField *pUserPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspUserPasswordUpdate)
        {
            ((OnRspUserPasswordUpdateType *)_OnRspUserPasswordUpdate)(pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 资金账户口令更新请求响应
    typedef void OnRspTradingAccountPasswordUpdateType(CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspTradingAccountPasswordUpdate;
    virtual void OnRspTradingAccountPasswordUpdate(CThostFtdcTradingAccountPasswordUpdateField *pTradingAccountPasswordUpdate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspTradingAccountPasswordUpdate)
        {
            ((OnRspTradingAccountPasswordUpdateType *)_OnRspTradingAccountPasswordUpdate)(pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 查询用户当前支持的认证模式的回复
    typedef void OnRspUserAuthMethodType(CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserAuthMethod;
    virtual void OnRspUserAuthMethod(CThostFtdcRspUserAuthMethodField *pRspUserAuthMethod, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspUserAuthMethod)
        {
            ((OnRspUserAuthMethodType *)_OnRspUserAuthMethod)(pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 获取图形验证码请求的回复
    typedef void OnRspGenUserCaptchaType(CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspGenUserCaptcha;
    virtual void OnRspGenUserCaptcha(CThostFtdcRspGenUserCaptchaField *pRspGenUserCaptcha, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspGenUserCaptcha)
        {
            ((OnRspGenUserCaptchaType *)_OnRspGenUserCaptcha)(pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 获取短信验证码请求的回复
    typedef void OnRspGenUserTextType(CThostFtdcRspGenUserTextField *pRspGenUserText, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspGenUserText;
    virtual void OnRspGenUserText(CThostFtdcRspGenUserTextField *pRspGenUserText, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspGenUserText)
        {
            ((OnRspGenUserTextType *)_OnRspGenUserText)(pRspGenUserText, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 报单录入请求响应
    typedef void OnRspOrderInsertType(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOrderInsert;
    virtual void OnRspOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspOrderInsert)
        {
            ((OnRspOrderInsertType *)_OnRspOrderInsert)(pInputOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 预埋单录入请求响应
    typedef void OnRspParkedOrderInsertType(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspParkedOrderInsert;
    virtual void OnRspParkedOrderInsert(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspParkedOrderInsert)
        {
            ((OnRspParkedOrderInsertType *)_OnRspParkedOrderInsert)(pParkedOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 预埋撤单录入请求响应
    typedef void OnRspParkedOrderActionType(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspParkedOrderAction;
    virtual void OnRspParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspParkedOrderAction)
        {
            ((OnRspParkedOrderActionType *)_OnRspParkedOrderAction)(pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 报单操作请求响应
    typedef void OnRspOrderActionType(CThostFtdcInputOrderActionField *pInputOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOrderAction;
    virtual void OnRspOrderAction(CThostFtdcInputOrderActionField *pInputOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspOrderAction)
        {
            ((OnRspOrderActionType *)_OnRspOrderAction)(pInputOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 查询最大报单数量响应
    typedef void OnRspQryMaxOrderVolumeType(CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMaxOrderVolume;
    virtual void OnRspQryMaxOrderVolume(CThostFtdcQryMaxOrderVolumeField *pQryMaxOrderVolume, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryMaxOrderVolume)
        {
            ((OnRspQryMaxOrderVolumeType *)_OnRspQryMaxOrderVolume)(pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 投资者结算结果确认响应
    typedef void OnRspSettlementInfoConfirmType(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspSettlementInfoConfirm;
    virtual void OnRspSettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspSettlementInfoConfirm)
        {
            ((OnRspSettlementInfoConfirmType *)_OnRspSettlementInfoConfirm)(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 删除预埋单响应
    typedef void OnRspRemoveParkedOrderType(CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspRemoveParkedOrder;
    virtual void OnRspRemoveParkedOrder(CThostFtdcRemoveParkedOrderField *pRemoveParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspRemoveParkedOrder)
        {
            ((OnRspRemoveParkedOrderType *)_OnRspRemoveParkedOrder)(pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 删除预埋撤单响应
    typedef void OnRspRemoveParkedOrderActionType(CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspRemoveParkedOrderAction;
    virtual void OnRspRemoveParkedOrderAction(CThostFtdcRemoveParkedOrderActionField *pRemoveParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspRemoveParkedOrderAction)
        {
            ((OnRspRemoveParkedOrderActionType *)_OnRspRemoveParkedOrderAction)(pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 执行宣告录入请求响应
    typedef void OnRspExecOrderInsertType(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspExecOrderInsert;
    virtual void OnRspExecOrderInsert(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspExecOrderInsert)
        {
            ((OnRspExecOrderInsertType *)_OnRspExecOrderInsert)(pInputExecOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 执行宣告操作请求响应
    typedef void OnRspExecOrderActionType(CThostFtdcInputExecOrderActionField *pInputExecOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspExecOrderAction;
    virtual void OnRspExecOrderAction(CThostFtdcInputExecOrderActionField *pInputExecOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspExecOrderAction)
        {
            ((OnRspExecOrderActionType *)_OnRspExecOrderAction)(pInputExecOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 询价录入请求响应
    typedef void OnRspForQuoteInsertType(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspForQuoteInsert;
    virtual void OnRspForQuoteInsert(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspForQuoteInsert)
        {
            ((OnRspForQuoteInsertType *)_OnRspForQuoteInsert)(pInputForQuote, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 报价录入请求响应
    typedef void OnRspQuoteInsertType(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQuoteInsert;
    virtual void OnRspQuoteInsert(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQuoteInsert)
        {
            ((OnRspQuoteInsertType *)_OnRspQuoteInsert)(pInputQuote, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 报价操作请求响应
    typedef void OnRspQuoteActionType(CThostFtdcInputQuoteActionField *pInputQuoteAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQuoteAction;
    virtual void OnRspQuoteAction(CThostFtdcInputQuoteActionField *pInputQuoteAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQuoteAction)
        {
            ((OnRspQuoteActionType *)_OnRspQuoteAction)(pInputQuoteAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 批量报单操作请求响应
    typedef void OnRspBatchOrderActionType(CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspBatchOrderAction;
    virtual void OnRspBatchOrderAction(CThostFtdcInputBatchOrderActionField *pInputBatchOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspBatchOrderAction)
        {
            ((OnRspBatchOrderActionType *)_OnRspBatchOrderAction)(pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 期权自对冲录入请求响应
    typedef void OnRspOptionSelfCloseInsertType(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOptionSelfCloseInsert;
    virtual void OnRspOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspOptionSelfCloseInsert)
        {
            ((OnRspOptionSelfCloseInsertType *)_OnRspOptionSelfCloseInsert)(pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 期权自对冲操作请求响应
    typedef void OnRspOptionSelfCloseActionType(CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspOptionSelfCloseAction;
    virtual void OnRspOptionSelfCloseAction(CThostFtdcInputOptionSelfCloseActionField *pInputOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspOptionSelfCloseAction)
        {
            ((OnRspOptionSelfCloseActionType *)_OnRspOptionSelfCloseAction)(pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 申请组合录入请求响应
    typedef void OnRspCombActionInsertType(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspCombActionInsert;
    virtual void OnRspCombActionInsert(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspCombActionInsert)
        {
            ((OnRspCombActionInsertType *)_OnRspCombActionInsert)(pInputCombAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询报单响应
    typedef void OnRspQryOrderType(CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOrder;
    virtual void OnRspQryOrder(CThostFtdcOrderField *pOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryOrder)
        {
            ((OnRspQryOrderType *)_OnRspQryOrder)(pOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询成交响应
    typedef void OnRspQryTradeType(CThostFtdcTradeField *pTrade, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTrade;
    virtual void OnRspQryTrade(CThostFtdcTradeField *pTrade, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTrade)
        {
            ((OnRspQryTradeType *)_OnRspQryTrade)(pTrade, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资者持仓响应
    typedef void OnRspQryInvestorPositionType(CThostFtdcInvestorPositionField *pInvestorPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPosition;
    virtual void OnRspQryInvestorPosition(CThostFtdcInvestorPositionField *pInvestorPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestorPosition)
        {
            ((OnRspQryInvestorPositionType *)_OnRspQryInvestorPosition)(pInvestorPosition, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询资金账户响应
    typedef void OnRspQryTradingAccountType(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTradingAccount;
    virtual void OnRspQryTradingAccount(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTradingAccount)
        {
            ((OnRspQryTradingAccountType *)_OnRspQryTradingAccount)(pTradingAccount, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资者响应
    typedef void OnRspQryInvestorType(CThostFtdcInvestorField *pInvestor, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestor;
    virtual void OnRspQryInvestor(CThostFtdcInvestorField *pInvestor, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestor)
        {
            ((OnRspQryInvestorType *)_OnRspQryInvestor)(pInvestor, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询交易编码响应
    typedef void OnRspQryTradingCodeType(CThostFtdcTradingCodeField *pTradingCode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTradingCode;
    virtual void OnRspQryTradingCode(CThostFtdcTradingCodeField *pTradingCode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTradingCode)
        {
            ((OnRspQryTradingCodeType *)_OnRspQryTradingCode)(pTradingCode, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询合约保证金率响应
    typedef void OnRspQryInstrumentMarginRateType(CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrumentMarginRate;
    virtual void OnRspQryInstrumentMarginRate(CThostFtdcInstrumentMarginRateField *pInstrumentMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInstrumentMarginRate)
        {
            ((OnRspQryInstrumentMarginRateType *)_OnRspQryInstrumentMarginRate)(pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询合约手续费率响应
    typedef void OnRspQryInstrumentCommissionRateType(CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrumentCommissionRate;
    virtual void OnRspQryInstrumentCommissionRate(CThostFtdcInstrumentCommissionRateField *pInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInstrumentCommissionRate)
        {
            ((OnRspQryInstrumentCommissionRateType *)_OnRspQryInstrumentCommissionRate)(pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询交易所响应
    typedef void OnRspQryExchangeType(CThostFtdcExchangeField *pExchange, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchange;
    virtual void OnRspQryExchange(CThostFtdcExchangeField *pExchange, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryExchange)
        {
            ((OnRspQryExchangeType *)_OnRspQryExchange)(pExchange, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询产品响应
    typedef void OnRspQryProductType(CThostFtdcProductField *pProduct, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryProduct;
    virtual void OnRspQryProduct(CThostFtdcProductField *pProduct, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryProduct)
        {
            ((OnRspQryProductType *)_OnRspQryProduct)(pProduct, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询合约响应
    typedef void OnRspQryInstrumentType(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrument;
    virtual void OnRspQryInstrument(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInstrument)
        {
            ((OnRspQryInstrumentType *)_OnRspQryInstrument)(pInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询行情响应
    typedef void OnRspQryDepthMarketDataType(CThostFtdcDepthMarketDataField *pDepthMarketData, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryDepthMarketData;
    virtual void OnRspQryDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryDepthMarketData)
        {
            ((OnRspQryDepthMarketDataType *)_OnRspQryDepthMarketData)(pDepthMarketData, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询交易员报盘机响应
    typedef void OnRspQryTraderOfferType(CThostFtdcTraderOfferField *pTraderOffer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTraderOffer;
    virtual void OnRspQryTraderOffer(CThostFtdcTraderOfferField *pTraderOffer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTraderOffer)
        {
            ((OnRspQryTraderOfferType *)_OnRspQryTraderOffer)(pTraderOffer, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资者结算结果响应
    typedef void OnRspQrySettlementInfoType(CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySettlementInfo;
    virtual void OnRspQrySettlementInfo(CThostFtdcSettlementInfoField *pSettlementInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySettlementInfo)
        {
            ((OnRspQrySettlementInfoType *)_OnRspQrySettlementInfo)(pSettlementInfo, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询转帐银行响应
    typedef void OnRspQryTransferBankType(CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTransferBank;
    virtual void OnRspQryTransferBank(CThostFtdcTransferBankField *pTransferBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTransferBank)
        {
            ((OnRspQryTransferBankType *)_OnRspQryTransferBank)(pTransferBank, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资者持仓明细响应
    typedef void OnRspQryInvestorPositionDetailType(CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPositionDetail;
    virtual void OnRspQryInvestorPositionDetail(CThostFtdcInvestorPositionDetailField *pInvestorPositionDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestorPositionDetail)
        {
            ((OnRspQryInvestorPositionDetailType *)_OnRspQryInvestorPositionDetail)(pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询客户通知响应
    typedef void OnRspQryNoticeType(CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryNotice;
    virtual void OnRspQryNotice(CThostFtdcNoticeField *pNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryNotice)
        {
            ((OnRspQryNoticeType *)_OnRspQryNotice)(pNotice, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询结算信息确认响应
    typedef void OnRspQrySettlementInfoConfirmType(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySettlementInfoConfirm;
    virtual void OnRspQrySettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField *pSettlementInfoConfirm, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySettlementInfoConfirm)
        {
            ((OnRspQrySettlementInfoConfirmType *)_OnRspQrySettlementInfoConfirm)(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资者持仓明细响应
    typedef void OnRspQryInvestorPositionCombineDetailType(CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPositionCombineDetail;
    virtual void OnRspQryInvestorPositionCombineDetail(CThostFtdcInvestorPositionCombineDetailField *pInvestorPositionCombineDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestorPositionCombineDetail)
        {
            ((OnRspQryInvestorPositionCombineDetailType *)_OnRspQryInvestorPositionCombineDetail)(pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 查询保证金监管系统经纪公司资金账户密钥响应
    typedef void OnRspQryCFMMCTradingAccountKeyType(CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCFMMCTradingAccountKey;
    virtual void OnRspQryCFMMCTradingAccountKey(CThostFtdcCFMMCTradingAccountKeyField *pCFMMCTradingAccountKey, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryCFMMCTradingAccountKey)
        {
            ((OnRspQryCFMMCTradingAccountKeyType *)_OnRspQryCFMMCTradingAccountKey)(pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询仓单折抵信息响应
    typedef void OnRspQryEWarrantOffsetType(CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryEWarrantOffset;
    virtual void OnRspQryEWarrantOffset(CThostFtdcEWarrantOffsetField *pEWarrantOffset, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryEWarrantOffset)
        {
            ((OnRspQryEWarrantOffsetType *)_OnRspQryEWarrantOffset)(pEWarrantOffset, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资者品种/跨品种保证金响应
    typedef void OnRspQryInvestorProductGroupMarginType(CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorProductGroupMargin;
    virtual void OnRspQryInvestorProductGroupMargin(CThostFtdcInvestorProductGroupMarginField *pInvestorProductGroupMargin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestorProductGroupMargin)
        {
            ((OnRspQryInvestorProductGroupMarginType *)_OnRspQryInvestorProductGroupMargin)(pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询交易所保证金率响应
    typedef void OnRspQryExchangeMarginRateType(CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchangeMarginRate;
    virtual void OnRspQryExchangeMarginRate(CThostFtdcExchangeMarginRateField *pExchangeMarginRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryExchangeMarginRate)
        {
            ((OnRspQryExchangeMarginRateType *)_OnRspQryExchangeMarginRate)(pExchangeMarginRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询交易所调整保证金率响应
    typedef void OnRspQryExchangeMarginRateAdjustType(CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchangeMarginRateAdjust;
    virtual void OnRspQryExchangeMarginRateAdjust(CThostFtdcExchangeMarginRateAdjustField *pExchangeMarginRateAdjust, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryExchangeMarginRateAdjust)
        {
            ((OnRspQryExchangeMarginRateAdjustType *)_OnRspQryExchangeMarginRateAdjust)(pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询汇率响应
    typedef void OnRspQryExchangeRateType(CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExchangeRate;
    virtual void OnRspQryExchangeRate(CThostFtdcExchangeRateField *pExchangeRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryExchangeRate)
        {
            ((OnRspQryExchangeRateType *)_OnRspQryExchangeRate)(pExchangeRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询二级代理操作员银期权限响应
    typedef void OnRspQrySecAgentACIDMapType(CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentACIDMap;
    virtual void OnRspQrySecAgentACIDMap(CThostFtdcSecAgentACIDMapField *pSecAgentACIDMap, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySecAgentACIDMap)
        {
            ((OnRspQrySecAgentACIDMapType *)_OnRspQrySecAgentACIDMap)(pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询产品报价汇率
    typedef void OnRspQryProductExchRateType(CThostFtdcProductExchRateField *pProductExchRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryProductExchRate;
    virtual void OnRspQryProductExchRate(CThostFtdcProductExchRateField *pProductExchRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryProductExchRate)
        {
            ((OnRspQryProductExchRateType *)_OnRspQryProductExchRate)(pProductExchRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询产品组
    typedef void OnRspQryProductGroupType(CThostFtdcProductGroupField *pProductGroup, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryProductGroup;
    virtual void OnRspQryProductGroup(CThostFtdcProductGroupField *pProductGroup, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryProductGroup)
        {
            ((OnRspQryProductGroupType *)_OnRspQryProductGroup)(pProductGroup, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询做市商合约手续费率响应
    typedef void OnRspQryMMInstrumentCommissionRateType(CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMMInstrumentCommissionRate;
    virtual void OnRspQryMMInstrumentCommissionRate(CThostFtdcMMInstrumentCommissionRateField *pMMInstrumentCommissionRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryMMInstrumentCommissionRate)
        {
            ((OnRspQryMMInstrumentCommissionRateType *)_OnRspQryMMInstrumentCommissionRate)(pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询做市商期权合约手续费响应
    typedef void OnRspQryMMOptionInstrCommRateType(CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMMOptionInstrCommRate;
    virtual void OnRspQryMMOptionInstrCommRate(CThostFtdcMMOptionInstrCommRateField *pMMOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryMMOptionInstrCommRate)
        {
            ((OnRspQryMMOptionInstrCommRateType *)_OnRspQryMMOptionInstrCommRate)(pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询报单手续费响应
    typedef void OnRspQryInstrumentOrderCommRateType(CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInstrumentOrderCommRate;
    virtual void OnRspQryInstrumentOrderCommRate(CThostFtdcInstrumentOrderCommRateField *pInstrumentOrderCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInstrumentOrderCommRate)
        {
            ((OnRspQryInstrumentOrderCommRateType *)_OnRspQryInstrumentOrderCommRate)(pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询资金账户响应
    typedef void OnRspQrySecAgentTradingAccountType(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentTradingAccount;
    virtual void OnRspQrySecAgentTradingAccount(CThostFtdcTradingAccountField *pTradingAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySecAgentTradingAccount)
        {
            ((OnRspQrySecAgentTradingAccountType *)_OnRspQrySecAgentTradingAccount)(pTradingAccount, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询二级代理商资金校验模式响应
    typedef void OnRspQrySecAgentCheckModeType(CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentCheckMode;
    virtual void OnRspQrySecAgentCheckMode(CThostFtdcSecAgentCheckModeField *pSecAgentCheckMode, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySecAgentCheckMode)
        {
            ((OnRspQrySecAgentCheckModeType *)_OnRspQrySecAgentCheckMode)(pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询二级代理商信息响应
    typedef void OnRspQrySecAgentTradeInfoType(CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySecAgentTradeInfo;
    virtual void OnRspQrySecAgentTradeInfo(CThostFtdcSecAgentTradeInfoField *pSecAgentTradeInfo, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySecAgentTradeInfo)
        {
            ((OnRspQrySecAgentTradeInfoType *)_OnRspQrySecAgentTradeInfo)(pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询期权交易成本响应
    typedef void OnRspQryOptionInstrTradeCostType(CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOptionInstrTradeCost;
    virtual void OnRspQryOptionInstrTradeCost(CThostFtdcOptionInstrTradeCostField *pOptionInstrTradeCost, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryOptionInstrTradeCost)
        {
            ((OnRspQryOptionInstrTradeCostType *)_OnRspQryOptionInstrTradeCost)(pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询期权合约手续费响应
    typedef void OnRspQryOptionInstrCommRateType(CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOptionInstrCommRate;
    virtual void OnRspQryOptionInstrCommRate(CThostFtdcOptionInstrCommRateField *pOptionInstrCommRate, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryOptionInstrCommRate)
        {
            ((OnRspQryOptionInstrCommRateType *)_OnRspQryOptionInstrCommRate)(pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询执行宣告响应
    typedef void OnRspQryExecOrderType(CThostFtdcExecOrderField *pExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryExecOrder;
    virtual void OnRspQryExecOrder(CThostFtdcExecOrderField *pExecOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryExecOrder)
        {
            ((OnRspQryExecOrderType *)_OnRspQryExecOrder)(pExecOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询询价响应
    typedef void OnRspQryForQuoteType(CThostFtdcForQuoteField *pForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryForQuote;
    virtual void OnRspQryForQuote(CThostFtdcForQuoteField *pForQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryForQuote)
        {
            ((OnRspQryForQuoteType *)_OnRspQryForQuote)(pForQuote, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询报价响应
    typedef void OnRspQryQuoteType(CThostFtdcQuoteField *pQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryQuote;
    virtual void OnRspQryQuote(CThostFtdcQuoteField *pQuote, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryQuote)
        {
            ((OnRspQryQuoteType *)_OnRspQryQuote)(pQuote, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询期权自对冲响应
    typedef void OnRspQryOptionSelfCloseType(CThostFtdcOptionSelfCloseField *pOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryOptionSelfClose;
    virtual void OnRspQryOptionSelfClose(CThostFtdcOptionSelfCloseField *pOptionSelfClose, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryOptionSelfClose)
        {
            ((OnRspQryOptionSelfCloseType *)_OnRspQryOptionSelfClose)(pOptionSelfClose, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询投资单元响应
    typedef void OnRspQryInvestUnitType(CThostFtdcInvestUnitField *pInvestUnit, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestUnit;
    virtual void OnRspQryInvestUnit(CThostFtdcInvestUnitField *pInvestUnit, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestUnit)
        {
            ((OnRspQryInvestUnitType *)_OnRspQryInvestUnit)(pInvestUnit, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询组合合约安全系数响应
    typedef void OnRspQryCombInstrumentGuardType(CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCombInstrumentGuard;
    virtual void OnRspQryCombInstrumentGuard(CThostFtdcCombInstrumentGuardField *pCombInstrumentGuard, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryCombInstrumentGuard)
        {
            ((OnRspQryCombInstrumentGuardType *)_OnRspQryCombInstrumentGuard)(pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询申请组合响应
    typedef void OnRspQryCombActionType(CThostFtdcCombActionField *pCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCombAction;
    virtual void OnRspQryCombAction(CThostFtdcCombActionField *pCombAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryCombAction)
        {
            ((OnRspQryCombActionType *)_OnRspQryCombAction)(pCombAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询转帐流水响应
    typedef void OnRspQryTransferSerialType(CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTransferSerial;
    virtual void OnRspQryTransferSerial(CThostFtdcTransferSerialField *pTransferSerial, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTransferSerial)
        {
            ((OnRspQryTransferSerialType *)_OnRspQryTransferSerial)(pTransferSerial, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询银期签约关系响应
    typedef void OnRspQryAccountregisterType(CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryAccountregister;
    virtual void OnRspQryAccountregister(CThostFtdcAccountregisterField *pAccountregister, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryAccountregister)
        {
            ((OnRspQryAccountregisterType *)_OnRspQryAccountregister)(pAccountregister, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 错误应答
    typedef void OnRspErrorType(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspError;
    virtual void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspError)
        {
            ((OnRspErrorType *)_OnRspError)(pRspInfo, nRequestID, bIsLast);
        }
    }
    // 报单通知
    typedef void OnRtnOrderType(CThostFtdcOrderField *pOrder);
    void *_OnRtnOrder;
    virtual void OnRtnOrder(CThostFtdcOrderField *pOrder)
    {
        if (_OnRtnOrder)
        {
            ((OnRtnOrderType *)_OnRtnOrder)(pOrder);
        }
    }
    // 成交通知
    typedef void OnRtnTradeType(CThostFtdcTradeField *pTrade);
    void *_OnRtnTrade;
    virtual void OnRtnTrade(CThostFtdcTradeField *pTrade)
    {
        if (_OnRtnTrade)
        {
            ((OnRtnTradeType *)_OnRtnTrade)(pTrade);
        }
    }
    // 报单录入错误回报
    typedef void OnErrRtnOrderInsertType(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOrderInsert;
    virtual void OnErrRtnOrderInsert(CThostFtdcInputOrderField *pInputOrder, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnOrderInsert)
        {
            ((OnErrRtnOrderInsertType *)_OnErrRtnOrderInsert)(pInputOrder, pRspInfo);
        }
    }
    // 报单操作错误回报
    typedef void OnErrRtnOrderActionType(CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOrderAction;
    virtual void OnErrRtnOrderAction(CThostFtdcOrderActionField *pOrderAction, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnOrderAction)
        {
            ((OnErrRtnOrderActionType *)_OnErrRtnOrderAction)(pOrderAction, pRspInfo);
        }
    }
    // 合约交易状态通知
    typedef void OnRtnInstrumentStatusType(CThostFtdcInstrumentStatusField *pInstrumentStatus);
    void *_OnRtnInstrumentStatus;
    virtual void OnRtnInstrumentStatus(CThostFtdcInstrumentStatusField *pInstrumentStatus)
    {
        if (_OnRtnInstrumentStatus)
        {
            ((OnRtnInstrumentStatusType *)_OnRtnInstrumentStatus)(pInstrumentStatus);
        }
    }
    // 交易所公告通知
    typedef void OnRtnBulletinType(CThostFtdcBulletinField *pBulletin);
    void *_OnRtnBulletin;
    virtual void OnRtnBulletin(CThostFtdcBulletinField *pBulletin)
    {
        if (_OnRtnBulletin)
        {
            ((OnRtnBulletinType *)_OnRtnBulletin)(pBulletin);
        }
    }
    // 交易通知
    typedef void OnRtnTradingNoticeType(CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo);
    void *_OnRtnTradingNotice;
    virtual void OnRtnTradingNotice(CThostFtdcTradingNoticeInfoField *pTradingNoticeInfo)
    {
        if (_OnRtnTradingNotice)
        {
            ((OnRtnTradingNoticeType *)_OnRtnTradingNotice)(pTradingNoticeInfo);
        }
    }
    // 提示条件单校验错误
    typedef void OnRtnErrorConditionalOrderType(CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder);
    void *_OnRtnErrorConditionalOrder;
    virtual void OnRtnErrorConditionalOrder(CThostFtdcErrorConditionalOrderField *pErrorConditionalOrder)
    {
        if (_OnRtnErrorConditionalOrder)
        {
            ((OnRtnErrorConditionalOrderType *)_OnRtnErrorConditionalOrder)(pErrorConditionalOrder);
        }
    }
    // 执行宣告通知
    typedef void OnRtnExecOrderType(CThostFtdcExecOrderField *pExecOrder);
    void *_OnRtnExecOrder;
    virtual void OnRtnExecOrder(CThostFtdcExecOrderField *pExecOrder)
    {
        if (_OnRtnExecOrder)
        {
            ((OnRtnExecOrderType *)_OnRtnExecOrder)(pExecOrder);
        }
    }
    // 执行宣告录入错误回报
    typedef void OnErrRtnExecOrderInsertType(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnExecOrderInsert;
    virtual void OnErrRtnExecOrderInsert(CThostFtdcInputExecOrderField *pInputExecOrder, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnExecOrderInsert)
        {
            ((OnErrRtnExecOrderInsertType *)_OnErrRtnExecOrderInsert)(pInputExecOrder, pRspInfo);
        }
    }
    // 执行宣告操作错误回报
    typedef void OnErrRtnExecOrderActionType(CThostFtdcExecOrderActionField *pExecOrderAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnExecOrderAction;
    virtual void OnErrRtnExecOrderAction(CThostFtdcExecOrderActionField *pExecOrderAction, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnExecOrderAction)
        {
            ((OnErrRtnExecOrderActionType *)_OnErrRtnExecOrderAction)(pExecOrderAction, pRspInfo);
        }
    }
    // 询价录入错误回报
    typedef void OnErrRtnForQuoteInsertType(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnForQuoteInsert;
    virtual void OnErrRtnForQuoteInsert(CThostFtdcInputForQuoteField *pInputForQuote, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnForQuoteInsert)
        {
            ((OnErrRtnForQuoteInsertType *)_OnErrRtnForQuoteInsert)(pInputForQuote, pRspInfo);
        }
    }
    // 报价通知
    typedef void OnRtnQuoteType(CThostFtdcQuoteField *pQuote);
    void *_OnRtnQuote;
    virtual void OnRtnQuote(CThostFtdcQuoteField *pQuote)
    {
        if (_OnRtnQuote)
        {
            ((OnRtnQuoteType *)_OnRtnQuote)(pQuote);
        }
    }
    // 报价录入错误回报
    typedef void OnErrRtnQuoteInsertType(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnQuoteInsert;
    virtual void OnErrRtnQuoteInsert(CThostFtdcInputQuoteField *pInputQuote, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnQuoteInsert)
        {
            ((OnErrRtnQuoteInsertType *)_OnErrRtnQuoteInsert)(pInputQuote, pRspInfo);
        }
    }
    // 报价操作错误回报
    typedef void OnErrRtnQuoteActionType(CThostFtdcQuoteActionField *pQuoteAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnQuoteAction;
    virtual void OnErrRtnQuoteAction(CThostFtdcQuoteActionField *pQuoteAction, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnQuoteAction)
        {
            ((OnErrRtnQuoteActionType *)_OnErrRtnQuoteAction)(pQuoteAction, pRspInfo);
        }
    }
    // 询价通知
    typedef void OnRtnForQuoteRspType(CThostFtdcForQuoteRspField *pForQuoteRsp);
    void *_OnRtnForQuoteRsp;
    virtual void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp)
    {
        if (_OnRtnForQuoteRsp)
        {
            ((OnRtnForQuoteRspType *)_OnRtnForQuoteRsp)(pForQuoteRsp);
        }
    }
    // 保证金监控中心用户令牌
    typedef void OnRtnCFMMCTradingAccountTokenType(CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken);
    void *_OnRtnCFMMCTradingAccountToken;
    virtual void OnRtnCFMMCTradingAccountToken(CThostFtdcCFMMCTradingAccountTokenField *pCFMMCTradingAccountToken)
    {
        if (_OnRtnCFMMCTradingAccountToken)
        {
            ((OnRtnCFMMCTradingAccountTokenType *)_OnRtnCFMMCTradingAccountToken)(pCFMMCTradingAccountToken);
        }
    }
    // 批量报单操作错误回报
    typedef void OnErrRtnBatchOrderActionType(CThostFtdcBatchOrderActionField *pBatchOrderAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnBatchOrderAction;
    virtual void OnErrRtnBatchOrderAction(CThostFtdcBatchOrderActionField *pBatchOrderAction, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnBatchOrderAction)
        {
            ((OnErrRtnBatchOrderActionType *)_OnErrRtnBatchOrderAction)(pBatchOrderAction, pRspInfo);
        }
    }
    // 期权自对冲通知
    typedef void OnRtnOptionSelfCloseType(CThostFtdcOptionSelfCloseField *pOptionSelfClose);
    void *_OnRtnOptionSelfClose;
    virtual void OnRtnOptionSelfClose(CThostFtdcOptionSelfCloseField *pOptionSelfClose)
    {
        if (_OnRtnOptionSelfClose)
        {
            ((OnRtnOptionSelfCloseType *)_OnRtnOptionSelfClose)(pOptionSelfClose);
        }
    }
    // 期权自对冲录入错误回报
    typedef void OnErrRtnOptionSelfCloseInsertType(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOptionSelfCloseInsert;
    virtual void OnErrRtnOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField *pInputOptionSelfClose, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnOptionSelfCloseInsert)
        {
            ((OnErrRtnOptionSelfCloseInsertType *)_OnErrRtnOptionSelfCloseInsert)(pInputOptionSelfClose, pRspInfo);
        }
    }
    // 期权自对冲操作错误回报
    typedef void OnErrRtnOptionSelfCloseActionType(CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnOptionSelfCloseAction;
    virtual void OnErrRtnOptionSelfCloseAction(CThostFtdcOptionSelfCloseActionField *pOptionSelfCloseAction, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnOptionSelfCloseAction)
        {
            ((OnErrRtnOptionSelfCloseActionType *)_OnErrRtnOptionSelfCloseAction)(pOptionSelfCloseAction, pRspInfo);
        }
    }
    // 申请组合通知
    typedef void OnRtnCombActionType(CThostFtdcCombActionField *pCombAction);
    void *_OnRtnCombAction;
    virtual void OnRtnCombAction(CThostFtdcCombActionField *pCombAction)
    {
        if (_OnRtnCombAction)
        {
            ((OnRtnCombActionType *)_OnRtnCombAction)(pCombAction);
        }
    }
    // 申请组合录入错误回报
    typedef void OnErrRtnCombActionInsertType(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnCombActionInsert;
    virtual void OnErrRtnCombActionInsert(CThostFtdcInputCombActionField *pInputCombAction, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnCombActionInsert)
        {
            ((OnErrRtnCombActionInsertType *)_OnErrRtnCombActionInsert)(pInputCombAction, pRspInfo);
        }
    }
    // 请求查询签约银行响应
    typedef void OnRspQryContractBankType(CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryContractBank;
    virtual void OnRspQryContractBank(CThostFtdcContractBankField *pContractBank, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryContractBank)
        {
            ((OnRspQryContractBankType *)_OnRspQryContractBank)(pContractBank, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询预埋单响应
    typedef void OnRspQryParkedOrderType(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryParkedOrder;
    virtual void OnRspQryParkedOrder(CThostFtdcParkedOrderField *pParkedOrder, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryParkedOrder)
        {
            ((OnRspQryParkedOrderType *)_OnRspQryParkedOrder)(pParkedOrder, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询预埋撤单响应
    typedef void OnRspQryParkedOrderActionType(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryParkedOrderAction;
    virtual void OnRspQryParkedOrderAction(CThostFtdcParkedOrderActionField *pParkedOrderAction, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryParkedOrderAction)
        {
            ((OnRspQryParkedOrderActionType *)_OnRspQryParkedOrderAction)(pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询交易通知响应
    typedef void OnRspQryTradingNoticeType(CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryTradingNotice;
    virtual void OnRspQryTradingNotice(CThostFtdcTradingNoticeField *pTradingNotice, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryTradingNotice)
        {
            ((OnRspQryTradingNoticeType *)_OnRspQryTradingNotice)(pTradingNotice, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询经纪公司交易参数响应
    typedef void OnRspQryBrokerTradingParamsType(CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryBrokerTradingParams;
    virtual void OnRspQryBrokerTradingParams(CThostFtdcBrokerTradingParamsField *pBrokerTradingParams, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryBrokerTradingParams)
        {
            ((OnRspQryBrokerTradingParamsType *)_OnRspQryBrokerTradingParams)(pBrokerTradingParams, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询经纪公司交易算法响应
    typedef void OnRspQryBrokerTradingAlgosType(CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryBrokerTradingAlgos;
    virtual void OnRspQryBrokerTradingAlgos(CThostFtdcBrokerTradingAlgosField *pBrokerTradingAlgos, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryBrokerTradingAlgos)
        {
            ((OnRspQryBrokerTradingAlgosType *)_OnRspQryBrokerTradingAlgos)(pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求查询监控中心用户令牌
    typedef void OnRspQueryCFMMCTradingAccountTokenType(CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQueryCFMMCTradingAccountToken;
    virtual void OnRspQueryCFMMCTradingAccountToken(CThostFtdcQueryCFMMCTradingAccountTokenField *pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQueryCFMMCTradingAccountToken)
        {
            ((OnRspQueryCFMMCTradingAccountTokenType *)_OnRspQueryCFMMCTradingAccountToken)(pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 银行发起银行资金转期货通知
    typedef void OnRtnFromBankToFutureByBankType(CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromBankToFutureByBank;
    virtual void OnRtnFromBankToFutureByBank(CThostFtdcRspTransferField *pRspTransfer)
    {
        if (_OnRtnFromBankToFutureByBank)
        {
            ((OnRtnFromBankToFutureByBankType *)_OnRtnFromBankToFutureByBank)(pRspTransfer);
        }
    }
    // 银行发起期货资金转银行通知
    typedef void OnRtnFromFutureToBankByBankType(CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromFutureToBankByBank;
    virtual void OnRtnFromFutureToBankByBank(CThostFtdcRspTransferField *pRspTransfer)
    {
        if (_OnRtnFromFutureToBankByBank)
        {
            ((OnRtnFromFutureToBankByBankType *)_OnRtnFromFutureToBankByBank)(pRspTransfer);
        }
    }
    // 银行发起冲正银行转期货通知
    typedef void OnRtnRepealFromBankToFutureByBankType(CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromBankToFutureByBank;
    virtual void OnRtnRepealFromBankToFutureByBank(CThostFtdcRspRepealField *pRspRepeal)
    {
        if (_OnRtnRepealFromBankToFutureByBank)
        {
            ((OnRtnRepealFromBankToFutureByBankType *)_OnRtnRepealFromBankToFutureByBank)(pRspRepeal);
        }
    }
    // 银行发起冲正期货转银行通知
    typedef void OnRtnRepealFromFutureToBankByBankType(CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromFutureToBankByBank;
    virtual void OnRtnRepealFromFutureToBankByBank(CThostFtdcRspRepealField *pRspRepeal)
    {
        if (_OnRtnRepealFromFutureToBankByBank)
        {
            ((OnRtnRepealFromFutureToBankByBankType *)_OnRtnRepealFromFutureToBankByBank)(pRspRepeal);
        }
    }
    // 期货发起银行资金转期货通知
    typedef void OnRtnFromBankToFutureByFutureType(CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromBankToFutureByFuture;
    virtual void OnRtnFromBankToFutureByFuture(CThostFtdcRspTransferField *pRspTransfer)
    {
        if (_OnRtnFromBankToFutureByFuture)
        {
            ((OnRtnFromBankToFutureByFutureType *)_OnRtnFromBankToFutureByFuture)(pRspTransfer);
        }
    }
    // 期货发起期货资金转银行通知
    typedef void OnRtnFromFutureToBankByFutureType(CThostFtdcRspTransferField *pRspTransfer);
    void *_OnRtnFromFutureToBankByFuture;
    virtual void OnRtnFromFutureToBankByFuture(CThostFtdcRspTransferField *pRspTransfer)
    {
        if (_OnRtnFromFutureToBankByFuture)
        {
            ((OnRtnFromFutureToBankByFutureType *)_OnRtnFromFutureToBankByFuture)(pRspTransfer);
        }
    }
    // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    typedef void OnRtnRepealFromBankToFutureByFutureManualType(CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromBankToFutureByFutureManual;
    virtual void OnRtnRepealFromBankToFutureByFutureManual(CThostFtdcRspRepealField *pRspRepeal)
    {
        if (_OnRtnRepealFromBankToFutureByFutureManual)
        {
            ((OnRtnRepealFromBankToFutureByFutureManualType *)_OnRtnRepealFromBankToFutureByFutureManual)(pRspRepeal);
        }
    }
    // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    typedef void OnRtnRepealFromFutureToBankByFutureManualType(CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromFutureToBankByFutureManual;
    virtual void OnRtnRepealFromFutureToBankByFutureManual(CThostFtdcRspRepealField *pRspRepeal)
    {
        if (_OnRtnRepealFromFutureToBankByFutureManual)
        {
            ((OnRtnRepealFromFutureToBankByFutureManualType *)_OnRtnRepealFromFutureToBankByFutureManual)(pRspRepeal);
        }
    }
    // 期货发起查询银行余额通知
    typedef void OnRtnQueryBankBalanceByFutureType(CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount);
    void *_OnRtnQueryBankBalanceByFuture;
    virtual void OnRtnQueryBankBalanceByFuture(CThostFtdcNotifyQueryAccountField *pNotifyQueryAccount)
    {
        if (_OnRtnQueryBankBalanceByFuture)
        {
            ((OnRtnQueryBankBalanceByFutureType *)_OnRtnQueryBankBalanceByFuture)(pNotifyQueryAccount);
        }
    }
    // 期货发起银行资金转期货错误回报
    typedef void OnErrRtnBankToFutureByFutureType(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnBankToFutureByFuture;
    virtual void OnErrRtnBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnBankToFutureByFuture)
        {
            ((OnErrRtnBankToFutureByFutureType *)_OnErrRtnBankToFutureByFuture)(pReqTransfer, pRspInfo);
        }
    }
    // 期货发起期货资金转银行错误回报
    typedef void OnErrRtnFutureToBankByFutureType(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnFutureToBankByFuture;
    virtual void OnErrRtnFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnFutureToBankByFuture)
        {
            ((OnErrRtnFutureToBankByFutureType *)_OnErrRtnFutureToBankByFuture)(pReqTransfer, pRspInfo);
        }
    }
    // 系统运行时期货端手工发起冲正银行转期货错误回报
    typedef void OnErrRtnRepealBankToFutureByFutureManualType(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnRepealBankToFutureByFutureManual;
    virtual void OnErrRtnRepealBankToFutureByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnRepealBankToFutureByFutureManual)
        {
            ((OnErrRtnRepealBankToFutureByFutureManualType *)_OnErrRtnRepealBankToFutureByFutureManual)(pReqRepeal, pRspInfo);
        }
    }
    // 系统运行时期货端手工发起冲正期货转银行错误回报
    typedef void OnErrRtnRepealFutureToBankByFutureManualType(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnRepealFutureToBankByFutureManual;
    virtual void OnErrRtnRepealFutureToBankByFutureManual(CThostFtdcReqRepealField *pReqRepeal, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnRepealFutureToBankByFutureManual)
        {
            ((OnErrRtnRepealFutureToBankByFutureManualType *)_OnErrRtnRepealFutureToBankByFutureManual)(pReqRepeal, pRspInfo);
        }
    }
    // 期货发起查询银行余额错误回报
    typedef void OnErrRtnQueryBankBalanceByFutureType(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo);
    void *_OnErrRtnQueryBankBalanceByFuture;
    virtual void OnErrRtnQueryBankBalanceByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo)
    {
        if (_OnErrRtnQueryBankBalanceByFuture)
        {
            ((OnErrRtnQueryBankBalanceByFutureType *)_OnErrRtnQueryBankBalanceByFuture)(pReqQueryAccount, pRspInfo);
        }
    }
    // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    typedef void OnRtnRepealFromBankToFutureByFutureType(CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromBankToFutureByFuture;
    virtual void OnRtnRepealFromBankToFutureByFuture(CThostFtdcRspRepealField *pRspRepeal)
    {
        if (_OnRtnRepealFromBankToFutureByFuture)
        {
            ((OnRtnRepealFromBankToFutureByFutureType *)_OnRtnRepealFromBankToFutureByFuture)(pRspRepeal);
        }
    }
    // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    typedef void OnRtnRepealFromFutureToBankByFutureType(CThostFtdcRspRepealField *pRspRepeal);
    void *_OnRtnRepealFromFutureToBankByFuture;
    virtual void OnRtnRepealFromFutureToBankByFuture(CThostFtdcRspRepealField *pRspRepeal)
    {
        if (_OnRtnRepealFromFutureToBankByFuture)
        {
            ((OnRtnRepealFromFutureToBankByFutureType *)_OnRtnRepealFromFutureToBankByFuture)(pRspRepeal);
        }
    }
    // 期货发起银行资金转期货应答
    typedef void OnRspFromBankToFutureByFutureType(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspFromBankToFutureByFuture;
    virtual void OnRspFromBankToFutureByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspFromBankToFutureByFuture)
        {
            ((OnRspFromBankToFutureByFutureType *)_OnRspFromBankToFutureByFuture)(pReqTransfer, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 期货发起期货资金转银行应答
    typedef void OnRspFromFutureToBankByFutureType(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspFromFutureToBankByFuture;
    virtual void OnRspFromFutureToBankByFuture(CThostFtdcReqTransferField *pReqTransfer, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspFromFutureToBankByFuture)
        {
            ((OnRspFromFutureToBankByFutureType *)_OnRspFromFutureToBankByFuture)(pReqTransfer, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 期货发起查询银行余额应答
    typedef void OnRspQueryBankAccountMoneyByFutureType(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQueryBankAccountMoneyByFuture;
    virtual void OnRspQueryBankAccountMoneyByFuture(CThostFtdcReqQueryAccountField *pReqQueryAccount, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQueryBankAccountMoneyByFuture)
        {
            ((OnRspQueryBankAccountMoneyByFutureType *)_OnRspQueryBankAccountMoneyByFuture)(pReqQueryAccount, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 银行发起银期开户通知
    typedef void OnRtnOpenAccountByBankType(CThostFtdcOpenAccountField *pOpenAccount);
    void *_OnRtnOpenAccountByBank;
    virtual void OnRtnOpenAccountByBank(CThostFtdcOpenAccountField *pOpenAccount)
    {
        if (_OnRtnOpenAccountByBank)
        {
            ((OnRtnOpenAccountByBankType *)_OnRtnOpenAccountByBank)(pOpenAccount);
        }
    }
    // 银行发起银期销户通知
    typedef void OnRtnCancelAccountByBankType(CThostFtdcCancelAccountField *pCancelAccount);
    void *_OnRtnCancelAccountByBank;
    virtual void OnRtnCancelAccountByBank(CThostFtdcCancelAccountField *pCancelAccount)
    {
        if (_OnRtnCancelAccountByBank)
        {
            ((OnRtnCancelAccountByBankType *)_OnRtnCancelAccountByBank)(pCancelAccount);
        }
    }
    // 银行发起变更银行账号通知
    typedef void OnRtnChangeAccountByBankType(CThostFtdcChangeAccountField *pChangeAccount);
    void *_OnRtnChangeAccountByBank;
    virtual void OnRtnChangeAccountByBank(CThostFtdcChangeAccountField *pChangeAccount)
    {
        if (_OnRtnChangeAccountByBank)
        {
            ((OnRtnChangeAccountByBankType *)_OnRtnChangeAccountByBank)(pChangeAccount);
        }
    }
    // 请求查询分类合约响应
    typedef void OnRspQryClassifiedInstrumentType(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryClassifiedInstrument;
    virtual void OnRspQryClassifiedInstrument(CThostFtdcInstrumentField *pInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryClassifiedInstrument)
        {
            ((OnRspQryClassifiedInstrumentType *)_OnRspQryClassifiedInstrument)(pInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 请求组合优惠比例响应
    typedef void OnRspQryCombPromotionParamType(CThostFtdcCombPromotionParamField *pCombPromotionParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryCombPromotionParam;
    virtual void OnRspQryCombPromotionParam(CThostFtdcCombPromotionParamField *pCombPromotionParam, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryCombPromotionParam)
        {
            ((OnRspQryCombPromotionParamType *)_OnRspQryCombPromotionParam)(pCombPromotionParam, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 投资者风险结算持仓查询响应
    typedef void OnRspQryRiskSettleInvstPositionType(CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRiskSettleInvstPosition;
    virtual void OnRspQryRiskSettleInvstPosition(CThostFtdcRiskSettleInvstPositionField *pRiskSettleInvstPosition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryRiskSettleInvstPosition)
        {
            ((OnRspQryRiskSettleInvstPositionType *)_OnRspQryRiskSettleInvstPosition)(pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 风险结算产品查询响应
    typedef void OnRspQryRiskSettleProductStatusType(CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryRiskSettleProductStatus;
    virtual void OnRspQryRiskSettleProductStatus(CThostFtdcRiskSettleProductStatusField *pRiskSettleProductStatus, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryRiskSettleProductStatus)
        {
            ((OnRspQryRiskSettleProductStatusType *)_OnRspQryRiskSettleProductStatus)(pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast);
        }
    }
    // SPBM期货合约参数查询响应
    typedef void OnRspQrySPBMFutureParameterType(CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMFutureParameter;
    virtual void OnRspQrySPBMFutureParameter(CThostFtdcSPBMFutureParameterField *pSPBMFutureParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySPBMFutureParameter)
        {
            ((OnRspQrySPBMFutureParameterType *)_OnRspQrySPBMFutureParameter)(pSPBMFutureParameter, pRspInfo, nRequestID, bIsLast);
        }
    }
    // SPBM期权合约参数查询响应
    typedef void OnRspQrySPBMOptionParameterType(CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMOptionParameter;
    virtual void OnRspQrySPBMOptionParameter(CThostFtdcSPBMOptionParameterField *pSPBMOptionParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySPBMOptionParameter)
        {
            ((OnRspQrySPBMOptionParameterType *)_OnRspQrySPBMOptionParameter)(pSPBMOptionParameter, pRspInfo, nRequestID, bIsLast);
        }
    }
    // SPBM品种内对锁仓折扣参数查询响应
    typedef void OnRspQrySPBMIntraParameterType(CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMIntraParameter;
    virtual void OnRspQrySPBMIntraParameter(CThostFtdcSPBMIntraParameterField *pSPBMIntraParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySPBMIntraParameter)
        {
            ((OnRspQrySPBMIntraParameterType *)_OnRspQrySPBMIntraParameter)(pSPBMIntraParameter, pRspInfo, nRequestID, bIsLast);
        }
    }
    // SPBM跨品种抵扣参数查询响应
    typedef void OnRspQrySPBMInterParameterType(CThostFtdcSPBMInterParameterField *pSPBMInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMInterParameter;
    virtual void OnRspQrySPBMInterParameter(CThostFtdcSPBMInterParameterField *pSPBMInterParameter, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySPBMInterParameter)
        {
            ((OnRspQrySPBMInterParameterType *)_OnRspQrySPBMInterParameter)(pSPBMInterParameter, pRspInfo, nRequestID, bIsLast);
        }
    }
    // SPBM组合保证金套餐查询响应
    typedef void OnRspQrySPBMPortfDefinitionType(CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMPortfDefinition;
    virtual void OnRspQrySPBMPortfDefinition(CThostFtdcSPBMPortfDefinitionField *pSPBMPortfDefinition, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySPBMPortfDefinition)
        {
            ((OnRspQrySPBMPortfDefinitionType *)_OnRspQrySPBMPortfDefinition)(pSPBMPortfDefinition, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 投资者SPBM套餐选择查询响应
    typedef void OnRspQrySPBMInvestorPortfDefType(CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQrySPBMInvestorPortfDef;
    virtual void OnRspQrySPBMInvestorPortfDef(CThostFtdcSPBMInvestorPortfDefField *pSPBMInvestorPortfDef, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQrySPBMInvestorPortfDef)
        {
            ((OnRspQrySPBMInvestorPortfDefType *)_OnRspQrySPBMInvestorPortfDef)(pSPBMInvestorPortfDef, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 投资者新型组合保证金系数查询响应
    typedef void OnRspQryInvestorPortfMarginRatioType(CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorPortfMarginRatio;
    virtual void OnRspQryInvestorPortfMarginRatio(CThostFtdcInvestorPortfMarginRatioField *pInvestorPortfMarginRatio, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestorPortfMarginRatio)
        {
            ((OnRspQryInvestorPortfMarginRatioType *)_OnRspQryInvestorPortfMarginRatio)(pInvestorPortfMarginRatio, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 投资者产品SPBM明细查询响应
    typedef void OnRspQryInvestorProdSPBMDetailType(CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryInvestorProdSPBMDetail;
    virtual void OnRspQryInvestorProdSPBMDetail(CThostFtdcInvestorProdSPBMDetailField *pInvestorProdSPBMDetail, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryInvestorProdSPBMDetail)
        {
            ((OnRspQryInvestorProdSPBMDetailType *)_OnRspQryInvestorProdSPBMDetail)(pInvestorProdSPBMDetail, pRspInfo, nRequestID, bIsLast);
        }
    }
};