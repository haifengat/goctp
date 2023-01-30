#define DLL_EXPORT extern "C"

#include "quote.h"
#include <iostream>

using namespace std;

DLL_EXPORT void* CreateFtdcMdApi(const char *pszFlowPath = "", const bool bIsUsingUdp=false, const bool bIsMulticast=false){
    cout << pszFlowPath << endl;
    return CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath, bIsUsingUdp, bIsMulticast);
}

// 创建 Quote 实例
DLL_EXPORT void* CreateFtdcMdSpi() {
    return new Quote(); 
}

// ******** 调用函数 *********
[[ range .Fn]]// [[ .Comment ]]
DLL_EXPORT [[ .RtnType ]] [[ .Name ]](CThostFtdcMdApi *api[[ range .Params ]], [[ .Type ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ if eq .Var "ppInstrumentID" ]][][[ end ]][[ end ]]){
    cout << "[[ .Name ]]" << endl;
    return api->[[ .Name ]]([[ range $idx, $param := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]][[ end ]]);
}
[[ end ]]

// **** 用 Set 函数将 go 函数指针赋值给 C 函数指针 ****
[[ range .On ]]// [[ .Comment ]]
DLL_EXPORT void Set[[ .Name ]](Quote *spi, void *onFunc){
    spi->_[[ .Name ]] = onFunc;
}
[[ end ]]