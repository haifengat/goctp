#include "quote.h"
#include <string.h>
int nReq;

Quote::Quote(void)
{	
	_FrontConnected = NULL;
	_FrontDisconnected = NULL;
	_HeartBeatWarning = NULL;
	_RspUserLogin = NULL;
	_RspUserLogout = NULL;
	_RspError = NULL;
	_RspSubMarketData = NULL;
	_RspUnSubMarketData = NULL;
	_RspSubForQuoteRsp = NULL;
	_RspUnSubForQuoteRsp = NULL;
	_RtnDepthMarketData = NULL;
	_RtnForQuoteRsp = NULL;
}

DLL_EXPORT_C_DECL void WINAPI qSetOnFrontConnected(Quote* spi, void* func){spi->_FrontConnected = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnFrontDisconnected(Quote* spi, void* func){spi->_FrontDisconnected = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnHeartBeatWarning(Quote* spi, void* func){spi->_HeartBeatWarning = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspUserLogin(Quote* spi, void* func){spi->_RspUserLogin = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspUserLogout(Quote* spi, void* func){spi->_RspUserLogout = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspError(Quote* spi, void* func){spi->_RspError = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspSubMarketData(Quote* spi, void* func){spi->_RspSubMarketData = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspUnSubMarketData(Quote* spi, void* func){spi->_RspUnSubMarketData = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspSubForQuoteRsp(Quote* spi, void* func){spi->_RspSubForQuoteRsp = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRspUnSubForQuoteRsp(Quote* spi, void* func){spi->_RspUnSubForQuoteRsp = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRtnDepthMarketData(Quote* spi, void* func){spi->_RtnDepthMarketData = func;}
DLL_EXPORT_C_DECL void WINAPI qSetOnRtnForQuoteRsp(Quote* spi, void* func){spi->_RtnForQuoteRsp = func;}

DLL_EXPORT_C_DECL void* WINAPI qCreateApi(){return CThostFtdcMdApi::CreateFtdcMdApi("./log/");}
DLL_EXPORT_C_DECL void* WINAPI qCreateSpi(){return new Quote();}

DLL_EXPORT_C_DECL void* WINAPI qRelease(CThostFtdcMdApi *api){api->Release(); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qInit(CThostFtdcMdApi *api){api->Init(); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qJoin(CThostFtdcMdApi *api){api->Join(); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qRegisterFront(CThostFtdcMdApi *api, char *pszFrontAddress){api->RegisterFront(pszFrontAddress); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qRegisterNameServer(CThostFtdcMdApi *api, char *pszNsAddress){api->RegisterNameServer(pszNsAddress); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qRegisterFensUserInfo(CThostFtdcMdApi *api, CThostFtdcFensUserInfoField * pFensUserInfo){api->RegisterFensUserInfo(pFensUserInfo); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qRegisterSpi(CThostFtdcMdApi *api, CThostFtdcMdSpi *pSpi){api->RegisterSpi(pSpi); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qSubscribeMarketData(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount){api->SubscribeMarketData(ppInstrumentID, nCount); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qUnSubscribeMarketData(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount){api->UnSubscribeMarketData(ppInstrumentID, nCount); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qSubscribeForQuoteRsp(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount){api->SubscribeForQuoteRsp(ppInstrumentID, nCount); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qUnSubscribeForQuoteRsp(CThostFtdcMdApi *api, char *ppInstrumentID[], int nCount){api->UnSubscribeForQuoteRsp(ppInstrumentID, nCount); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qReqUserLogin(CThostFtdcMdApi *api, CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID){api->ReqUserLogin(pReqUserLoginField, nRequestID); return 0;}
DLL_EXPORT_C_DECL void* WINAPI qReqUserLogout(CThostFtdcMdApi *api, CThostFtdcUserLogoutField *pUserLogout, int nRequestID){api->ReqUserLogout(pUserLogout, nRequestID); return 0;}