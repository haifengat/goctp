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
#include "../v6.6.8_20220712/ThostFtdcMdApi.h"
#pragma comment(lib, "../v6.6.8_20220712/thostmduserapi_se.lib")
#else
#define WINAPI      __stdcall
#include "../v6.6.8_20220712/ThostFtdcMdApi.h"
#pragma comment(lib, "../v6.6.8_20220712/thostmduserapi_se.lib")
#endif
#else
#define WINAPI
#include "../v6.6.8_20220712/ThostFtdcMdApi.h"
#endif

#include <string.h>

class Quote: CThostFtdcMdSpi
{
public:
    Quote(void);
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
    
	// 登录请求响应
	typedef int (WINAPI *RspUserLoginType)(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUserLogin;
	virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUserLogin) ((RspUserLoginType)_RspUserLogin)(pRspUserLogin, pRspInfo, nRequestID, bIsLast);}
    
	// 登出请求响应
	typedef int (WINAPI *RspUserLogoutType)(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUserLogout;
	virtual void OnRspUserLogout(CThostFtdcUserLogoutField *pUserLogout, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUserLogout) ((RspUserLogoutType)_RspUserLogout)(pUserLogout, pRspInfo, nRequestID, bIsLast);}
    
	// 请求查询组播合约响应
	typedef int (WINAPI *RspQryMulticastInstrumentType)(CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspQryMulticastInstrument;
	virtual void OnRspQryMulticastInstrument(CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspQryMulticastInstrument) ((RspQryMulticastInstrumentType)_RspQryMulticastInstrument)(pMulticastInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 错误应答
	typedef int (WINAPI *RspErrorType)(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspError;
	virtual void OnRspError(CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspError) ((RspErrorType)_RspError)(pRspInfo, nRequestID, bIsLast);}
    
	// 订阅行情应答
	typedef int (WINAPI *RspSubMarketDataType)(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspSubMarketData;
	virtual void OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspSubMarketData) ((RspSubMarketDataType)_RspSubMarketData)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 取消订阅行情应答
	typedef int (WINAPI *RspUnSubMarketDataType)(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUnSubMarketData;
	virtual void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUnSubMarketData) ((RspUnSubMarketDataType)_RspUnSubMarketData)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 订阅询价应答
	typedef int (WINAPI *RspSubForQuoteRspType)(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspSubForQuoteRsp;
	virtual void OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspSubForQuoteRsp) ((RspSubForQuoteRspType)_RspSubForQuoteRsp)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 取消订阅询价应答
	typedef int (WINAPI *RspUnSubForQuoteRspType)(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
	void *_RspUnSubForQuoteRsp;
	virtual void OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){if (_RspUnSubForQuoteRsp) ((RspUnSubForQuoteRspType)_RspUnSubForQuoteRsp)(pSpecificInstrument, pRspInfo, nRequestID, bIsLast);}
    
	// 深度行情通知
	typedef int (WINAPI *RtnDepthMarketDataType)(CThostFtdcDepthMarketDataField *pDepthMarketData);
	void *_RtnDepthMarketData;
	virtual void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData){if (_RtnDepthMarketData) ((RtnDepthMarketDataType)_RtnDepthMarketData)(pDepthMarketData);}
    
	// 询价通知
	typedef int (WINAPI *RtnForQuoteRspType)(CThostFtdcForQuoteRspField *pForQuoteRsp);
	void *_RtnForQuoteRsp;
	virtual void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField *pForQuoteRsp){if (_RtnForQuoteRsp) ((RtnForQuoteRspType)_RtnForQuoteRsp)(pForQuoteRsp);}
    
};
