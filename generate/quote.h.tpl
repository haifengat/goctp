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

    [[ range . ]]
	// [[ .Comment ]]
	typedef int (WINAPI *[[ .FuncTypeName ]])([[ range $i, $v := .FuncFields ]][[if gt $i 0]], [[end]][[ .FieldType ]] [[ .FieldName ]][[ end ]]);
	void *_[[ .FuncName ]];
	virtual void On[[ .FuncName ]]([[ range $n, $var := .FuncFields ]][[if gt $n 0]], [[end]][[ .FieldType ]] [[ .FieldName ]][[ end ]]){if (_[[ .FuncName ]]) (([[ .FuncTypeName ]])_[[ .FuncName ]])([[ range $n, $var := .FuncFields ]][[if gt $n 0]], [[end]][[ trimStar .FieldName ]][[ end ]]);}
    [[ end ]]
};
