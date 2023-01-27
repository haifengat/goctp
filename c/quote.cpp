#define DLL_EXPORT extern "C"
#include "../CTPv6.6.8_20220712/ThostFtdcMdApi.h"
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