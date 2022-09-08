#include "trade.h"
#include <string.h>
int nReq;

Trade::Trade(void)
{	
	[[ range .On ]]_[[ .FuncName ]] = NULL;
	[[ end ]]
}

[[ range .On ]]
// [[ .Comment ]]
DLL_EXPORT_C_DECL void WINAPI tSetOn[[ .FuncName ]](Trade* spi, void* func){spi->_[[ .FuncName ]] = func;}
[[ end ]]
DLL_EXPORT_C_DECL void* WINAPI CreateApi(){return CThostFtdcTraderApi::CreateFtdcTraderApi("./log/");}
DLL_EXPORT_C_DECL void* WINAPI CreateSpi(){return new Trade();}

[[ range .Fn ]]
// [[ .Comment ]]
DLL_EXPORT_C_DECL void* WINAPI t[[ .FuncName ]](CThostFtdcTraderApi *api[[ range .FuncFields ]], [[.FieldType]] [[.FieldName]][[end]]){api->[[.FuncName]]([[ range $i, $v := .FuncFields ]][[if gt $i 0]], [[end]][[trimStar .FieldName]][[end]]); return 0;}
[[ end ]]
