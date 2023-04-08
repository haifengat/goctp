#define DLL_EXPORT extern "C"

#include "quote.h"
#include <iostream>

using namespace std;

DLL_EXPORT void *CreateFtdcMdApi(const char *pszFlowPath = "", const bool bIsUsingUdp = false, const bool bIsMulticast = false)
{
    cout << pszFlowPath << endl;
    return CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath, bIsUsingUdp, bIsMulticast);
}

// 创建 Quote 实例
DLL_EXPORT void *CreateFtdcMdSpi()
{
    return new Quote();
}

// ******** 调用函数 *********
// 创建MdApi
DLL_EXPORT void Release(CThostFtdcMdApi *api)
{
    cout << "Release" << endl;
    return api->Release();
}
// 初始化
DLL_EXPORT void Init(CThostFtdcMdApi *api)
{
    cout << "Init" << endl;
    return api->Init();
}
// 等待接口线程结束运行
DLL_EXPORT int Join(CThostFtdcMdApi *api)
{
    cout << "Join" << endl;
    return api->Join();
}
// 注册前置机网络地址
DLL_EXPORT void RegisterFront(CThostFtdcMdApi *api, char *pszFrontAddress)
{
    cout << "RegisterFront" << endl;
    return api->RegisterFront(pszFrontAddress);
}
// @remark RegisterNameServer优先于RegisterFront
DLL_EXPORT void RegisterNameServer(CThostFtdcMdApi *api, char *pszNsAddress)
{
    cout << "RegisterNameServer" << endl;
    return api->RegisterNameServer(pszNsAddress);
}
// 注册名字服务器用户信息
DLL_EXPORT void RegisterFensUserInfo(CThostFtdcMdApi *api, CThostFtdcFensUserInfoField *pFensUserInfo)
{
    cout << "RegisterFensUserInfo" << endl;
    return api->RegisterFensUserInfo(pFensUserInfo);
}
// 注册回调接口
DLL_EXPORT void RegisterSpi(CThostFtdcMdApi *api, CThostFtdcMdSpi *pSpi)
{
    cout << "RegisterSpi" << endl;
    return api->RegisterSpi(pSpi);
}
// 订阅行情。
DLL_EXPORT int SubscribeMarketData(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount)
{
    cout << "SubscribeMarketData" << endl;
    return api->SubscribeMarketData(ppInstrumentID, nCount);
}
// 退订行情。
DLL_EXPORT int UnSubscribeMarketData(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount)
{
    cout << "UnSubscribeMarketData" << endl;
    return api->UnSubscribeMarketData(ppInstrumentID, nCount);
}
// 订阅询价。
DLL_EXPORT int SubscribeForQuoteRsp(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount)
{
    cout << "SubscribeForQuoteRsp" << endl;
    return api->SubscribeForQuoteRsp(ppInstrumentID, nCount);
}
// 退订询价。
DLL_EXPORT int UnSubscribeForQuoteRsp(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount)
{
    cout << "UnSubscribeForQuoteRsp" << endl;
    return api->UnSubscribeForQuoteRsp(ppInstrumentID, nCount);
}
// 用户登录请求
DLL_EXPORT int ReqUserLogin(CThostFtdcMdApi *api, CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID)
{
    cout << "ReqUserLogin" << endl;
    return api->ReqUserLogin(pReqUserLoginField, nRequestID);
}
// 登出请求
DLL_EXPORT int ReqUserLogout(CThostFtdcMdApi *api, CThostFtdcUserLogoutField *pUserLogout, int nRequestID)
{
    cout << "ReqUserLogout" << endl;
    return api->ReqUserLogout(pUserLogout, nRequestID);
}
// 请求查询组播合约
DLL_EXPORT int ReqQryMulticastInstrument(CThostFtdcMdApi *api, CThostFtdcQryMulticastInstrumentField *pQryMulticastInstrument, int nRequestID)
{
    cout << "ReqQryMulticastInstrument" << endl;
    return api->ReqQryMulticastInstrument(pQryMulticastInstrument, nRequestID);
}

// **** 用 Set 函数将 go 函数指针赋值给 C 函数指针 ****
// //////////////////////////////////////////////////////////////////////
DLL_EXPORT void SetOnFrontConnected(Quote *spi, void *onFunc)
{
    spi->_OnFrontConnected = onFunc;
}
// 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
DLL_EXPORT void SetOnFrontDisconnected(Quote *spi, void *onFunc)
{
    spi->_OnFrontDisconnected = onFunc;
}
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
DLL_EXPORT void SetOnHeartBeatWarning(Quote *spi, void *onFunc)
{
    spi->_OnHeartBeatWarning = onFunc;
}
// 登录请求响应
DLL_EXPORT void SetOnRspUserLogin(Quote *spi, void *onFunc)
{
    spi->_OnRspUserLogin = onFunc;
}
// 登出请求响应
DLL_EXPORT void SetOnRspUserLogout(Quote *spi, void *onFunc)
{
    spi->_OnRspUserLogout = onFunc;
}
// 请求查询组播合约响应
DLL_EXPORT void SetOnRspQryMulticastInstrument(Quote *spi, void *onFunc)
{
    spi->_OnRspQryMulticastInstrument = onFunc;
}
// 错误应答
DLL_EXPORT void SetOnRspError(Quote *spi, void *onFunc)
{
    spi->_OnRspError = onFunc;
}
// 订阅行情应答
DLL_EXPORT void SetOnRspSubMarketData(Quote *spi, void *onFunc)
{
    spi->_OnRspSubMarketData = onFunc;
}
// 取消订阅行情应答
DLL_EXPORT void SetOnRspUnSubMarketData(Quote *spi, void *onFunc)
{
    spi->_OnRspUnSubMarketData = onFunc;
}
// 订阅询价应答
DLL_EXPORT void SetOnRspSubForQuoteRsp(Quote *spi, void *onFunc)
{
    spi->_OnRspSubForQuoteRsp = onFunc;
}
// 取消订阅询价应答
DLL_EXPORT void SetOnRspUnSubForQuoteRsp(Quote *spi, void *onFunc)
{
    spi->_OnRspUnSubForQuoteRsp = onFunc;
}
// 深度行情通知
DLL_EXPORT void SetOnRtnDepthMarketData(Quote *spi, void *onFunc)
{
    spi->_OnRtnDepthMarketData = onFunc;
}
// 询价通知
DLL_EXPORT void SetOnRtnForQuoteRsp(Quote *spi, void *onFunc)
{
    spi->_OnRtnForQuoteRsp = onFunc;
}
