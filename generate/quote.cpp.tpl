#include "quote.h"
#include <string.h>
int nReq;

Quote::Quote(void)
{
	[[ range .On ]]_[[ .FuncName ]] = NULL;
	[[ end ]]
}

[[ range .On ]]DLL_EXPORT_C_DECL void WINAPI qSetOn[[ .FuncName ]](Quote* spi, void* func){spi->_[[ .FuncName ]] = func;}
[[ end ]]
DLL_EXPORT_C_DECL void* WINAPI qCreateApi(){return CThostFtdcMdApi::CreateFtdcMdApi("./log/");}
DLL_EXPORT_C_DECL void* WINAPI qCreateSpi(){return new Quote();}

[[ range .Fn ]]DLL_EXPORT_C_DECL void* WINAPI q[[ .FuncName ]](CThostFtdcMdApi *api[[ range .FuncFields ]], [[.FieldType]] [[.FieldName]][[if eq .FieldName "*ppInstrumentID"]][][[end]][[end]]){api->[[.FuncName]]([[ range $i, $v := .FuncFields ]][[if gt $i 0]], [[end]][[trimStar .FieldName]][[end]]); return 0;}
[[ end ]]
