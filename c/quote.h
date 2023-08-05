#include "../CTPv6.6.9_20220922/ThostFtdcMdApi.h"

class Quote: CThostFtdcMdSpi{
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
	// 请求查询组播合约响应    
    typedef void OnRspQryMulticastInstrumentType(void*, CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspQryMulticastInstrument;
    virtual void OnRspQryMulticastInstrument(CThostFtdcMulticastInstrumentField *pMulticastInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspQryMulticastInstrument) {
			((OnRspQryMulticastInstrumentType*)_OnRspQryMulticastInstrument)(this, pMulticastInstrument, pRspInfo, nRequestID, bIsLast);
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
	// 订阅行情应答    
    typedef void OnRspSubMarketDataType(void*, CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspSubMarketData;
    virtual void OnRspSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspSubMarketData) {
			((OnRspSubMarketDataType*)_OnRspSubMarketData)(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 取消订阅行情应答    
    typedef void OnRspUnSubMarketDataType(void*, CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUnSubMarketData;
    virtual void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspUnSubMarketData) {
			((OnRspUnSubMarketDataType*)_OnRspUnSubMarketData)(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 订阅询价应答    
    typedef void OnRspSubForQuoteRspType(void*, CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspSubForQuoteRsp;
    virtual void OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspSubForQuoteRsp) {
			((OnRspSubForQuoteRspType*)_OnRspSubForQuoteRsp)(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 取消订阅询价应答    
    typedef void OnRspUnSubForQuoteRspType(void*, CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUnSubForQuoteRsp;
    virtual void OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField *pSpecificInstrument, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast){
        if (_OnRspUnSubForQuoteRsp) {
			((OnRspUnSubForQuoteRspType*)_OnRspUnSubForQuoteRsp)(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
		}
    }
	// 深度行情通知    
    typedef void OnRtnDepthMarketDataType(void*, CThostFtdcDepthMarketDataField *pDepthMarketData);
    void *_OnRtnDepthMarketData;
    virtual void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField *pDepthMarketData){
        if (_OnRtnDepthMarketData) {
			((OnRtnDepthMarketDataType*)_OnRtnDepthMarketData)(this, pDepthMarketData);
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
	
};