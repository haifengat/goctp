
#include "../CTPv6.6.8_20220712/ThostFtdcMdApi.h"

class Quote: CThostFtdcMdSpi{
public:
    // 定义响应函数类型
    typedef void OnFrontConnectedType();
    // 声明响应函数指针的变量
    void *_OnFrontConnected;
    // 调用函数指针的变量
    virtual void OnFrontConnected(){
        if (_OnFrontConnected) { ((OnFrontConnectedType*)_OnFrontConnected)(); }
    }

    typedef void OnRspUserLoginType(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast);
    void *_OnRspUserLogin;
    virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {
        if (_OnRspUserLogin){
            ((OnRspUserLoginType *)_OnRspUserLogin)(pRspUserLogin, pRspInfo, nRequestID, bIsLast);
        }
    };
};