#define DLL_EXPORT extern "C"

#include "trade.h"
#include <iostream>

using namespace std;

DLL_EXPORT void* CreateFtdcTraderApi(const char *pszFlowPath = ""){
    cout << pszFlowPath << endl;
    return CThostFtdcTraderApi::CreateFtdcTraderApi(pszFlowPath);
}

DLL_EXPORT void* GetVersion(){
    return (void*)CThostFtdcTraderApi::GetApiVersion();
}

// 创建 Trade 实例
DLL_EXPORT void* CreateFtdcTraderSpi() {
    return new Trade(); 
}

// ******** 调用函数 *********
[[ range .Fn]]// [[ .Comment ]]
DLL_EXPORT [[ .RtnType ]] [[ .Name ]](CThostFtdcTraderApi *api[[ range .Params ]], [[ .Type ]] [[ if .HasStar ]]*[[ end ]][[ .Var ]][[ if eq .Var "ppInstrumentID" ]][][[ end ]][[ end ]]){
    cout << "[[ .Name ]]" << endl;
    return api->[[ .Name ]]([[ range $idx, $param := .Params ]][[ if gt $idx 0 ]], [[ end ]][[ .Var ]][[ end ]]);
}
[[ end ]]

// **** 用 Set 函数将 go 函数指针赋值给 C 函数指针 ****
[[ range .On ]]// [[ .Comment ]]
DLL_EXPORT void Set[[ .Name ]](Trade *spi, void *onFunc){
    spi->_[[ .Name ]] = onFunc;
}
[[ end ]]