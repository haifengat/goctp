#include "../CTPv6.6.9_20220922/ThostFtdcMdApi.h"

class Quote : CThostFtdcMdSpi
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
    // 请求查询组播合约响应
    typedef void OnRspQryMulticastInstrumentType(CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMulticastInstrument;
    virtual void OnRspQryMulticastInstrument(CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspQryMulticastInstrument)
        {
            ((OnRspQryMulticastInstrumentType *)_OnRspQryMulticastInstrument)(pMulticastInstrument, pRspInfo, nRequestID, bIsLast);
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
    // 订阅行情应答
    typedef void OnRspSubMarketDataType(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspSubMarketData;
    virtual void OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspSubMarketData)
        {
            ((OnRspSubMarketDataType *)_OnRspSubMarketData)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 取消订阅行情应答
    typedef void OnRspUnSubMarketDataType(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUnSubMarketData;
    virtual void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspUnSubMarketData)
        {
            ((OnRspUnSubMarketDataType *)_OnRspUnSubMarketData)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 订阅询价应答
    typedef void OnRspSubForQuoteRspType(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspSubForQuoteRsp;
    virtual void OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspSubForQuoteRsp)
        {
            ((OnRspSubForQuoteRspType *)_OnRspSubForQuoteRsp)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 取消订阅询价应答
    typedef void OnRspUnSubForQuoteRspType(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUnSubForQuoteRsp;
    virtual void OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast)
    {
        if (_OnRspUnSubForQuoteRsp)
        {
            ((OnRspUnSubForQuoteRspType *)_OnRspUnSubForQuoteRsp)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }
    // 深度行情通知
    typedef void OnRtnDepthMarketDataType(CThostFtdcDepthMarketDataField *pDepthMarketData);
    void *_OnRtnDepthMarketData;
    virtual void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData)
    {
        if (_OnRtnDepthMarketData)
        {
            ((OnRtnDepthMarketDataType *)_OnRtnDepthMarketData)(pDepthMarketData);
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
};