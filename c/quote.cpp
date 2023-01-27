#define DLL_EXPORT extern "C"
// #include "../CTPv6.6.8_20220712/ThostFtdcMdApi.h"
#include "quote.h"
#include <iostream>


using namespace std;

DLL_EXPORT void Init(CThostFtdcMdApi *api, char *pszFrontAddress){
    cout << "init" << endl;
    return api->Init();
}

DLL_EXPORT void RegisterFront(CThostFtdcMdApi *api, char *pszFrontAddress){
    cout << pszFrontAddress << endl;
    return api->RegisterFront(pszFrontAddress);
}

DLL_EXPORT void* CreateFtdcMdApi(const char *pszFlowPath = "", const bool bIsUsingUdp=false, const bool bIsMulticast=false){
    cout << pszFlowPath << endl;
    return CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath, bIsUsingUdp, bIsMulticast);
}

// 创建 Quote 实例
DLL_EXPORT void* CreateFtdcMdSpi() {
    return new Quote(); 
}

// 注册 Quote 实例给 Api
DLL_EXPORT void RegisterSpi(CThostFtdcMdApi *api, CThostFtdcMdSpi *spi) {
    api->RegisterSpi(spi);
}

// 用 Set 函数将 go 函数指针赋值给 C 函数指针
DLL_EXPORT void SetOnFrontConnected(Quote *spi, void *onFunc){
    spi->_OnFrontConnected = onFunc;
}