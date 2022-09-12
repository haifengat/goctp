package win

import (
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"

	"gitee.com/haifengat/goctp"
	ctp "gitee.com/haifengat/goctp/ctpdefine"
)

type Trade struct {
	goctp.HFTrade
	h        *syscall.DLL
	api, spi uintptr

    // 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
    _FrontConnected func()
    // 当客户端与交易后台通信连接断开时，该方法被调用。当发生这个情况后，API会自动重新连接，客户端可不做处理。
    _FrontDisconnected func(nReason int)
    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
    _HeartBeatWarning func(nTimeLapse int)
    // 客户端认证响应
    _RspAuthenticate func(pRspAuthenticateField *ctp.CThostFtdcRspAuthenticateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 登录请求响应
    _RspUserLogin func(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 登出请求响应
    _RspUserLogout func(pUserLogout *ctp.CThostFtdcUserLogoutField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 用户口令更新请求响应
    _RspUserPasswordUpdate func(pUserPasswordUpdate *ctp.CThostFtdcUserPasswordUpdateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 资金账户口令更新请求响应
    _RspTradingAccountPasswordUpdate func(pTradingAccountPasswordUpdate *ctp.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 查询用户当前支持的认证模式的回复
    _RspUserAuthMethod func(pRspUserAuthMethod *ctp.CThostFtdcRspUserAuthMethodField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 获取图形验证码请求的回复
    _RspGenUserCaptcha func(pRspGenUserCaptcha *ctp.CThostFtdcRspGenUserCaptchaField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 获取短信验证码请求的回复
    _RspGenUserText func(pRspGenUserText *ctp.CThostFtdcRspGenUserTextField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报单录入请求响应
    _RspOrderInsert func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 预埋单录入请求响应
    _RspParkedOrderInsert func(pParkedOrder *ctp.CThostFtdcParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 预埋撤单录入请求响应
    _RspParkedOrderAction func(pParkedOrderAction *ctp.CThostFtdcParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报单操作请求响应
    _RspOrderAction func(pInputOrderAction *ctp.CThostFtdcInputOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 查询最大报单数量响应
    _RspQryMaxOrderVolume func(pQryMaxOrderVolume *ctp.CThostFtdcQryMaxOrderVolumeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 投资者结算结果确认响应
    _RspSettlementInfoConfirm func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 删除预埋单响应
    _RspRemoveParkedOrder func(pRemoveParkedOrder *ctp.CThostFtdcRemoveParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 删除预埋撤单响应
    _RspRemoveParkedOrderAction func(pRemoveParkedOrderAction *ctp.CThostFtdcRemoveParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 执行宣告录入请求响应
    _RspExecOrderInsert func(pInputExecOrder *ctp.CThostFtdcInputExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 执行宣告操作请求响应
    _RspExecOrderAction func(pInputExecOrderAction *ctp.CThostFtdcInputExecOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 询价录入请求响应
    _RspForQuoteInsert func(pInputForQuote *ctp.CThostFtdcInputForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报价录入请求响应
    _RspQuoteInsert func(pInputQuote *ctp.CThostFtdcInputQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报价操作请求响应
    _RspQuoteAction func(pInputQuoteAction *ctp.CThostFtdcInputQuoteActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 批量报单操作请求响应
    _RspBatchOrderAction func(pInputBatchOrderAction *ctp.CThostFtdcInputBatchOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期权自对冲录入请求响应
    _RspOptionSelfCloseInsert func(pInputOptionSelfClose *ctp.CThostFtdcInputOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期权自对冲操作请求响应
    _RspOptionSelfCloseAction func(pInputOptionSelfCloseAction *ctp.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 申请组合录入请求响应
    _RspCombActionInsert func(pInputCombAction *ctp.CThostFtdcInputCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询报单响应
    _RspQryOrder func(pOrder *ctp.CThostFtdcOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询成交响应
    _RspQryTrade func(pTrade *ctp.CThostFtdcTradeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者持仓响应
    _RspQryInvestorPosition func(pInvestorPosition *ctp.CThostFtdcInvestorPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询资金账户响应
    _RspQryTradingAccount func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者响应
    _RspQryInvestor func(pInvestor *ctp.CThostFtdcInvestorField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易编码响应
    _RspQryTradingCode func(pTradingCode *ctp.CThostFtdcTradingCodeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询合约保证金率响应
    _RspQryInstrumentMarginRate func(pInstrumentMarginRate *ctp.CThostFtdcInstrumentMarginRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询合约手续费率响应
    _RspQryInstrumentCommissionRate func(pInstrumentCommissionRate *ctp.CThostFtdcInstrumentCommissionRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易所响应
    _RspQryExchange func(pExchange *ctp.CThostFtdcExchangeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询产品响应
    _RspQryProduct func(pProduct *ctp.CThostFtdcProductField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询合约响应
    _RspQryInstrument func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询行情响应
    _RspQryDepthMarketData func(pDepthMarketData *ctp.CThostFtdcDepthMarketDataField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易员报盘机响应
    _RspQryTraderOffer func(pTraderOffer *ctp.CThostFtdcTraderOfferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者结算结果响应
    _RspQrySettlementInfo func(pSettlementInfo *ctp.CThostFtdcSettlementInfoField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询转帐银行响应
    _RspQryTransferBank func(pTransferBank *ctp.CThostFtdcTransferBankField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者持仓明细响应
    _RspQryInvestorPositionDetail func(pInvestorPositionDetail *ctp.CThostFtdcInvestorPositionDetailField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询客户通知响应
    _RspQryNotice func(pNotice *ctp.CThostFtdcNoticeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询结算信息确认响应
    _RspQrySettlementInfoConfirm func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者持仓明细响应
    _RspQryInvestorPositionCombineDetail func(pInvestorPositionCombineDetail *ctp.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 查询保证金监管系统经纪公司资金账户密钥响应
    _RspQryCFMMCTradingAccountKey func(pCFMMCTradingAccountKey *ctp.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询仓单折抵信息响应
    _RspQryEWarrantOffset func(pEWarrantOffset *ctp.CThostFtdcEWarrantOffsetField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资者品种/跨品种保证金响应
    _RspQryInvestorProductGroupMargin func(pInvestorProductGroupMargin *ctp.CThostFtdcInvestorProductGroupMarginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易所保证金率响应
    _RspQryExchangeMarginRate func(pExchangeMarginRate *ctp.CThostFtdcExchangeMarginRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易所调整保证金率响应
    _RspQryExchangeMarginRateAdjust func(pExchangeMarginRateAdjust *ctp.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询汇率响应
    _RspQryExchangeRate func(pExchangeRate *ctp.CThostFtdcExchangeRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询二级代理操作员银期权限响应
    _RspQrySecAgentACIDMap func(pSecAgentACIDMap *ctp.CThostFtdcSecAgentACIDMapField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询产品报价汇率
    _RspQryProductExchRate func(pProductExchRate *ctp.CThostFtdcProductExchRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询产品组
    _RspQryProductGroup func(pProductGroup *ctp.CThostFtdcProductGroupField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询做市商合约手续费率响应
    _RspQryMMInstrumentCommissionRate func(pMMInstrumentCommissionRate *ctp.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询做市商期权合约手续费响应
    _RspQryMMOptionInstrCommRate func(pMMOptionInstrCommRate *ctp.CThostFtdcMMOptionInstrCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询报单手续费响应
    _RspQryInstrumentOrderCommRate func(pInstrumentOrderCommRate *ctp.CThostFtdcInstrumentOrderCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询资金账户响应
    _RspQrySecAgentTradingAccount func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询二级代理商资金校验模式响应
    _RspQrySecAgentCheckMode func(pSecAgentCheckMode *ctp.CThostFtdcSecAgentCheckModeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询二级代理商信息响应
    _RspQrySecAgentTradeInfo func(pSecAgentTradeInfo *ctp.CThostFtdcSecAgentTradeInfoField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询期权交易成本响应
    _RspQryOptionInstrTradeCost func(pOptionInstrTradeCost *ctp.CThostFtdcOptionInstrTradeCostField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询期权合约手续费响应
    _RspQryOptionInstrCommRate func(pOptionInstrCommRate *ctp.CThostFtdcOptionInstrCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询执行宣告响应
    _RspQryExecOrder func(pExecOrder *ctp.CThostFtdcExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询询价响应
    _RspQryForQuote func(pForQuote *ctp.CThostFtdcForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询报价响应
    _RspQryQuote func(pQuote *ctp.CThostFtdcQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询期权自对冲响应
    _RspQryOptionSelfClose func(pOptionSelfClose *ctp.CThostFtdcOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询投资单元响应
    _RspQryInvestUnit func(pInvestUnit *ctp.CThostFtdcInvestUnitField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询组合合约安全系数响应
    _RspQryCombInstrumentGuard func(pCombInstrumentGuard *ctp.CThostFtdcCombInstrumentGuardField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询申请组合响应
    _RspQryCombAction func(pCombAction *ctp.CThostFtdcCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询转帐流水响应
    _RspQryTransferSerial func(pTransferSerial *ctp.CThostFtdcTransferSerialField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询银期签约关系响应
    _RspQryAccountregister func(pAccountregister *ctp.CThostFtdcAccountregisterField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 错误应答
    _RspError func(pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 报单通知
    _RtnOrder func(pOrder *ctp.CThostFtdcOrderField)
    // 成交通知
    _RtnTrade func(pTrade *ctp.CThostFtdcTradeField)
    // 报单录入错误回报
    _ErrRtnOrderInsert func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 报单操作错误回报
    _ErrRtnOrderAction func(pOrderAction *ctp.CThostFtdcOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 合约交易状态通知
    _RtnInstrumentStatus func(pInstrumentStatus *ctp.CThostFtdcInstrumentStatusField)
    // 交易所公告通知
    _RtnBulletin func(pBulletin *ctp.CThostFtdcBulletinField)
    // 交易通知
    _RtnTradingNotice func(pTradingNoticeInfo *ctp.CThostFtdcTradingNoticeInfoField)
    // 提示条件单校验错误
    _RtnErrorConditionalOrder func(pErrorConditionalOrder *ctp.CThostFtdcErrorConditionalOrderField)
    // 执行宣告通知
    _RtnExecOrder func(pExecOrder *ctp.CThostFtdcExecOrderField)
    // 执行宣告录入错误回报
    _ErrRtnExecOrderInsert func(pInputExecOrder *ctp.CThostFtdcInputExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 执行宣告操作错误回报
    _ErrRtnExecOrderAction func(pExecOrderAction *ctp.CThostFtdcExecOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 询价录入错误回报
    _ErrRtnForQuoteInsert func(pInputForQuote *ctp.CThostFtdcInputForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 报价通知
    _RtnQuote func(pQuote *ctp.CThostFtdcQuoteField)
    // 报价录入错误回报
    _ErrRtnQuoteInsert func(pInputQuote *ctp.CThostFtdcInputQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 报价操作错误回报
    _ErrRtnQuoteAction func(pQuoteAction *ctp.CThostFtdcQuoteActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 询价通知
    _RtnForQuoteRsp func(pForQuoteRsp *ctp.CThostFtdcForQuoteRspField)
    // 保证金监控中心用户令牌
    _RtnCFMMCTradingAccountToken func(pCFMMCTradingAccountToken *ctp.CThostFtdcCFMMCTradingAccountTokenField)
    // 批量报单操作错误回报
    _ErrRtnBatchOrderAction func(pBatchOrderAction *ctp.CThostFtdcBatchOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期权自对冲通知
    _RtnOptionSelfClose func(pOptionSelfClose *ctp.CThostFtdcOptionSelfCloseField)
    // 期权自对冲录入错误回报
    _ErrRtnOptionSelfCloseInsert func(pInputOptionSelfClose *ctp.CThostFtdcInputOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期权自对冲操作错误回报
    _ErrRtnOptionSelfCloseAction func(pOptionSelfCloseAction *ctp.CThostFtdcOptionSelfCloseActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 申请组合通知
    _RtnCombAction func(pCombAction *ctp.CThostFtdcCombActionField)
    // 申请组合录入错误回报
    _ErrRtnCombActionInsert func(pInputCombAction *ctp.CThostFtdcInputCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 请求查询签约银行响应
    _RspQryContractBank func(pContractBank *ctp.CThostFtdcContractBankField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询预埋单响应
    _RspQryParkedOrder func(pParkedOrder *ctp.CThostFtdcParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询预埋撤单响应
    _RspQryParkedOrderAction func(pParkedOrderAction *ctp.CThostFtdcParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询交易通知响应
    _RspQryTradingNotice func(pTradingNotice *ctp.CThostFtdcTradingNoticeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询经纪公司交易参数响应
    _RspQryBrokerTradingParams func(pBrokerTradingParams *ctp.CThostFtdcBrokerTradingParamsField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询经纪公司交易算法响应
    _RspQryBrokerTradingAlgos func(pBrokerTradingAlgos *ctp.CThostFtdcBrokerTradingAlgosField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求查询监控中心用户令牌
    _RspQueryCFMMCTradingAccountToken func(pQueryCFMMCTradingAccountToken *ctp.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 银行发起银行资金转期货通知
    _RtnFromBankToFutureByBank func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 银行发起期货资金转银行通知
    _RtnFromFutureToBankByBank func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 银行发起冲正银行转期货通知
    _RtnRepealFromBankToFutureByBank func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 银行发起冲正期货转银行通知
    _RtnRepealFromFutureToBankByBank func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起银行资金转期货通知
    _RtnFromBankToFutureByFuture func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 期货发起期货资金转银行通知
    _RtnFromFutureToBankByFuture func(pRspTransfer *ctp.CThostFtdcRspTransferField)
    // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromBankToFutureByFutureManual func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromFutureToBankByFutureManual func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起查询银行余额通知
    _RtnQueryBankBalanceByFuture func(pNotifyQueryAccount *ctp.CThostFtdcNotifyQueryAccountField)
    // 期货发起银行资金转期货错误回报
    _ErrRtnBankToFutureByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期货发起期货资金转银行错误回报
    _ErrRtnFutureToBankByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 系统运行时期货端手工发起冲正银行转期货错误回报
    _ErrRtnRepealBankToFutureByFutureManual func(pReqRepeal *ctp.CThostFtdcReqRepealField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 系统运行时期货端手工发起冲正期货转银行错误回报
    _ErrRtnRepealFutureToBankByFutureManual func(pReqRepeal *ctp.CThostFtdcReqRepealField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期货发起查询银行余额错误回报
    _ErrRtnQueryBankBalanceByFuture func(pReqQueryAccount *ctp.CThostFtdcReqQueryAccountField, pRspInfo *ctp.CThostFtdcRspInfoField)
    // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromBankToFutureByFuture func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
    _RtnRepealFromFutureToBankByFuture func(pRspRepeal *ctp.CThostFtdcRspRepealField)
    // 期货发起银行资金转期货应答
    _RspFromBankToFutureByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期货发起期货资金转银行应答
    _RspFromFutureToBankByFuture func(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 期货发起查询银行余额应答
    _RspQueryBankAccountMoneyByFuture func(pReqQueryAccount *ctp.CThostFtdcReqQueryAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 银行发起银期开户通知
    _RtnOpenAccountByBank func(pOpenAccount *ctp.CThostFtdcOpenAccountField)
    // 银行发起银期销户通知
    _RtnCancelAccountByBank func(pCancelAccount *ctp.CThostFtdcCancelAccountField)
    // 银行发起变更银行账号通知
    _RtnChangeAccountByBank func(pChangeAccount *ctp.CThostFtdcChangeAccountField)
    // 请求查询分类合约响应
    _RspQryClassifiedInstrument func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 请求组合优惠比例响应
    _RspQryCombPromotionParam func(pCombPromotionParam *ctp.CThostFtdcCombPromotionParamField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 投资者风险结算持仓查询响应
    _RspQryRiskSettleInvstPosition func(pRiskSettleInvstPosition *ctp.CThostFtdcRiskSettleInvstPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    // 风险结算产品查询响应
    _RspQryRiskSettleProductStatus func(pRiskSettleProductStatus *ctp.CThostFtdcRiskSettleProductStatusField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool)
    
}

func (t *Trade) loadDll() {
	workPath, _ := os.Getwd()
	_, curFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("取当前文件路径失败")
	}
	dllPath := filepath.Dir(curFile)
	_ = os.Chdir(dllPath)
	t.h = syscall.MustLoadDLL("ctp_trade.dll")

	// 还原到之前的工作目录
	os.Chdir(workPath)
	//defer h.Release() // 函数结束后会释放导致后续函数执行失败
}

// NewTrade 实例化
func NewTrade() *Trade {
	t := new(Trade)

	// 主调函数封装 手动添加
	t.HFTrade.GetVersion = func() string {
		ver, _, _ := t.h.MustFindProc("tGetVersion").Call()
		p := (*byte)(unsafe.Pointer(ver))
		data := make([]byte, 0)
		for *p != 0 {
			data = append(data, *p)
			ver += unsafe.Sizeof(byte(0))
			p = (*byte)(unsafe.Pointer(ver))
		}
		return string(data)
	}
	t.HFTrade.ReqConnect = func(addr string) {
		bs, _ := syscall.BytePtrFromString(addr)
		t.h.MustFindProc("tRegisterFront").Call(t.api, uintptr(unsafe.Pointer(bs)))
		t.h.MustFindProc("tSubscribePrivateTopic").Call(t.api, uintptr(t.PrivateMode))
		t.h.MustFindProc("tSubscribePublicTopic").Call(t.api, uintptr(ctp.THOST_TERT_RESTART))
		t.h.MustFindProc("tInit").Call(t.api)
		// C.Join(t.api)
	}
	t.HFTrade.ReleaseAPI = func() {
		t.h.MustFindProc("tRelease").Call(t.api)
	}
	t.HFTrade.ReqAuthenticate = func(f *ctp.CThostFtdcReqAuthenticateField, i int) {
		t.h.MustFindProc("tReqAuthenticate").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqUserLogin = func(f *ctp.CThostFtdcReqUserLoginField, i int) {
		t.h.MustFindProc("tReqUserLogin").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqSettlementInfoConfirm = func(f *ctp.CThostFtdcSettlementInfoConfirmField, i int) {
		t.h.MustFindProc("tReqSettlementInfoConfirm").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryInstrument = func(f *ctp.CThostFtdcQryInstrumentField, i int) {
		t.h.MustFindProc("tReqQryInstrument").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryInvestor = func(f *ctp.CThostFtdcQryInvestorField, i int) {
		t.h.MustFindProc("tReqQryInvestor").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryClassifiedInstrument = func(f *ctp.CThostFtdcQryClassifiedInstrumentField, i int) {
		t.h.MustFindProc("tReqQryClassifiedInstrument").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryTradingAccount = func(f *ctp.CThostFtdcQryTradingAccountField, i int) {
		t.h.MustFindProc("tReqQryTradingAccount").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryInvestorPosition = func(f *ctp.CThostFtdcQryInvestorPositionField, i int) {
		t.h.MustFindProc("tReqQryInvestorPosition").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqOrder = func(f *ctp.CThostFtdcInputOrderField, i int) {
		t.h.MustFindProc("tReqOrderInsert").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqAction = func(f *ctp.CThostFtdcInputOrderActionField, i int) {
		t.h.MustFindProc("tReqOrderAction").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqFromBankToFutureByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		t.h.MustFindProc("tReqFromBankToFutureByFuture").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqFromFutureToBankByFuture = func(f *ctp.CThostFtdcReqTransferField, i int) {
		t.h.MustFindProc("tReqFromFutureToBankByFuture").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryOrder = func(f *ctp.CThostFtdcQryOrderField, i int) {
		t.h.MustFindProc("tReqQryOrder").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	t.HFTrade.ReqQryTrade = func(f *ctp.CThostFtdcQryTradeField, i int) {
		t.h.MustFindProc("tReqQryTrade").Call(t.api, uintptr(unsafe.Pointer(f)), uintptr(i))
	}
	
	// HFTrade 响应 手动添加即可增加新功能
	t._RtnFromFutureToBankByFuture = func(pRspTransfer *ctp.CThostFtdcRspTransferField) {
		t.HFTrade.RtnFromFutureToBankByFuture(pRspTransfer)
	}
	t._RtnFromBankToFutureByFuture = func(pRspTransfer *ctp.CThostFtdcRspTransferField) {
		t.HFTrade.RtnFromBankToFutureByFuture(pRspTransfer)
	}
	t._ErrRtnOrderAction = func(pOrderAction *ctp.CThostFtdcOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField) {
		t.HFTrade.ErrRtnOrderAction(pOrderAction, pRspInfo)
	}
	t._ErrRtnOrderInsert = func(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField) {
		t.HFTrade.ErrRtnOrderInsert(pInputOrder, pRspInfo)
	}
	t._RtnTrade = func(pTrade *ctp.CThostFtdcTradeField) {
		t.HFTrade.RtnTrade(pTrade)
	}
	t._RtnOrder = func(pOrder *ctp.CThostFtdcOrderField) {
		t.HFTrade.RtnOrder(pOrder)
	}
	t._RtnInstrumentStatus = func(pInstrumentStatus *ctp.CThostFtdcInstrumentStatusField) {
		t.HFTrade.RtnInstrumentStatus(pInstrumentStatus)
	}
	t._RspQryInvestorPosition = func(pInvestorPosition *ctp.CThostFtdcInvestorPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInvestorPosition(pInvestorPosition, bIsLast)
	}
	t._RspQryTradingAccount = func(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryTradingAccount(pTradingAccount, bIsLast)
	}
	t._RspQryTrade = func(pTrade *ctp.CThostFtdcTradeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryTrade(pTrade, bIsLast)
	}
	t._RspQryOrder = func(pOrder *ctp.CThostFtdcOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryOrder(pOrder, bIsLast)
	}
	t._RspQryInvestor = func(pInvestor *ctp.CThostFtdcInvestorField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInvestor(pInvestor, bIsLast)
	}
	t._RspQryInstrument = func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInstrument(pInstrument, bIsLast)
	}
	t._RspQryClassifiedInstrument = func(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspQryInstrument(pInstrument, bIsLast)
	}
	t._RspSettlementInfoConfirm = func(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspSettlementInfoConfirm()
	}
	t._RspUserLogin = func(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspUserLogin(pRspUserLogin, pRspInfo)
	}
	t._RspAuthenticate = func(pRspAuthenticateField *ctp.CThostFtdcRspAuthenticateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) {
		t.HFTrade.RspAuthenticate(pRspInfo)
	}
	t._FrontConnected = func() {
		t.HFTrade.FrontConnected()
	}
	t._FrontDisconnected = func(nReason int) {
		t.HFTrade.FrontDisConnected(nReason)
	}

	t.loadDll()
	t.HFTrade.Init() // 能够在 hftrade 中调用函数

	t.api, _, _ = t.h.MustFindProc("tCreateApi").Call()
	t.spi, _, _ = t.h.MustFindProc("tCreateSpi").Call()
	t.h.MustFindProc("tRegisterSpi").Call(t.api, uintptr(t.spi))

    t.h.MustFindProc("tSetOnFrontConnected").Call(t.spi, syscall.NewCallback(t.OnFrontConnected))
    t.h.MustFindProc("tSetOnFrontDisconnected").Call(t.spi, syscall.NewCallback(t.OnFrontDisconnected))
    t.h.MustFindProc("tSetOnHeartBeatWarning").Call(t.spi, syscall.NewCallback(t.OnHeartBeatWarning))
    t.h.MustFindProc("tSetOnRspAuthenticate").Call(t.spi, syscall.NewCallback(t.OnRspAuthenticate))
    t.h.MustFindProc("tSetOnRspUserLogin").Call(t.spi, syscall.NewCallback(t.OnRspUserLogin))
    t.h.MustFindProc("tSetOnRspUserLogout").Call(t.spi, syscall.NewCallback(t.OnRspUserLogout))
    t.h.MustFindProc("tSetOnRspUserPasswordUpdate").Call(t.spi, syscall.NewCallback(t.OnRspUserPasswordUpdate))
    t.h.MustFindProc("tSetOnRspTradingAccountPasswordUpdate").Call(t.spi, syscall.NewCallback(t.OnRspTradingAccountPasswordUpdate))
    t.h.MustFindProc("tSetOnRspUserAuthMethod").Call(t.spi, syscall.NewCallback(t.OnRspUserAuthMethod))
    t.h.MustFindProc("tSetOnRspGenUserCaptcha").Call(t.spi, syscall.NewCallback(t.OnRspGenUserCaptcha))
    t.h.MustFindProc("tSetOnRspGenUserText").Call(t.spi, syscall.NewCallback(t.OnRspGenUserText))
    t.h.MustFindProc("tSetOnRspOrderInsert").Call(t.spi, syscall.NewCallback(t.OnRspOrderInsert))
    t.h.MustFindProc("tSetOnRspParkedOrderInsert").Call(t.spi, syscall.NewCallback(t.OnRspParkedOrderInsert))
    t.h.MustFindProc("tSetOnRspParkedOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspParkedOrderAction))
    t.h.MustFindProc("tSetOnRspOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspOrderAction))
    t.h.MustFindProc("tSetOnRspQryMaxOrderVolume").Call(t.spi, syscall.NewCallback(t.OnRspQryMaxOrderVolume))
    t.h.MustFindProc("tSetOnRspSettlementInfoConfirm").Call(t.spi, syscall.NewCallback(t.OnRspSettlementInfoConfirm))
    t.h.MustFindProc("tSetOnRspRemoveParkedOrder").Call(t.spi, syscall.NewCallback(t.OnRspRemoveParkedOrder))
    t.h.MustFindProc("tSetOnRspRemoveParkedOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspRemoveParkedOrderAction))
    t.h.MustFindProc("tSetOnRspExecOrderInsert").Call(t.spi, syscall.NewCallback(t.OnRspExecOrderInsert))
    t.h.MustFindProc("tSetOnRspExecOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspExecOrderAction))
    t.h.MustFindProc("tSetOnRspForQuoteInsert").Call(t.spi, syscall.NewCallback(t.OnRspForQuoteInsert))
    t.h.MustFindProc("tSetOnRspQuoteInsert").Call(t.spi, syscall.NewCallback(t.OnRspQuoteInsert))
    t.h.MustFindProc("tSetOnRspQuoteAction").Call(t.spi, syscall.NewCallback(t.OnRspQuoteAction))
    t.h.MustFindProc("tSetOnRspBatchOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspBatchOrderAction))
    t.h.MustFindProc("tSetOnRspOptionSelfCloseInsert").Call(t.spi, syscall.NewCallback(t.OnRspOptionSelfCloseInsert))
    t.h.MustFindProc("tSetOnRspOptionSelfCloseAction").Call(t.spi, syscall.NewCallback(t.OnRspOptionSelfCloseAction))
    t.h.MustFindProc("tSetOnRspCombActionInsert").Call(t.spi, syscall.NewCallback(t.OnRspCombActionInsert))
    t.h.MustFindProc("tSetOnRspQryOrder").Call(t.spi, syscall.NewCallback(t.OnRspQryOrder))
    t.h.MustFindProc("tSetOnRspQryTrade").Call(t.spi, syscall.NewCallback(t.OnRspQryTrade))
    t.h.MustFindProc("tSetOnRspQryInvestorPosition").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestorPosition))
    t.h.MustFindProc("tSetOnRspQryTradingAccount").Call(t.spi, syscall.NewCallback(t.OnRspQryTradingAccount))
    t.h.MustFindProc("tSetOnRspQryInvestor").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestor))
    t.h.MustFindProc("tSetOnRspQryTradingCode").Call(t.spi, syscall.NewCallback(t.OnRspQryTradingCode))
    t.h.MustFindProc("tSetOnRspQryInstrumentMarginRate").Call(t.spi, syscall.NewCallback(t.OnRspQryInstrumentMarginRate))
    t.h.MustFindProc("tSetOnRspQryInstrumentCommissionRate").Call(t.spi, syscall.NewCallback(t.OnRspQryInstrumentCommissionRate))
    t.h.MustFindProc("tSetOnRspQryExchange").Call(t.spi, syscall.NewCallback(t.OnRspQryExchange))
    t.h.MustFindProc("tSetOnRspQryProduct").Call(t.spi, syscall.NewCallback(t.OnRspQryProduct))
    t.h.MustFindProc("tSetOnRspQryInstrument").Call(t.spi, syscall.NewCallback(t.OnRspQryInstrument))
    t.h.MustFindProc("tSetOnRspQryDepthMarketData").Call(t.spi, syscall.NewCallback(t.OnRspQryDepthMarketData))
    t.h.MustFindProc("tSetOnRspQryTraderOffer").Call(t.spi, syscall.NewCallback(t.OnRspQryTraderOffer))
    t.h.MustFindProc("tSetOnRspQrySettlementInfo").Call(t.spi, syscall.NewCallback(t.OnRspQrySettlementInfo))
    t.h.MustFindProc("tSetOnRspQryTransferBank").Call(t.spi, syscall.NewCallback(t.OnRspQryTransferBank))
    t.h.MustFindProc("tSetOnRspQryInvestorPositionDetail").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestorPositionDetail))
    t.h.MustFindProc("tSetOnRspQryNotice").Call(t.spi, syscall.NewCallback(t.OnRspQryNotice))
    t.h.MustFindProc("tSetOnRspQrySettlementInfoConfirm").Call(t.spi, syscall.NewCallback(t.OnRspQrySettlementInfoConfirm))
    t.h.MustFindProc("tSetOnRspQryInvestorPositionCombineDetail").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestorPositionCombineDetail))
    t.h.MustFindProc("tSetOnRspQryCFMMCTradingAccountKey").Call(t.spi, syscall.NewCallback(t.OnRspQryCFMMCTradingAccountKey))
    t.h.MustFindProc("tSetOnRspQryEWarrantOffset").Call(t.spi, syscall.NewCallback(t.OnRspQryEWarrantOffset))
    t.h.MustFindProc("tSetOnRspQryInvestorProductGroupMargin").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestorProductGroupMargin))
    t.h.MustFindProc("tSetOnRspQryExchangeMarginRate").Call(t.spi, syscall.NewCallback(t.OnRspQryExchangeMarginRate))
    t.h.MustFindProc("tSetOnRspQryExchangeMarginRateAdjust").Call(t.spi, syscall.NewCallback(t.OnRspQryExchangeMarginRateAdjust))
    t.h.MustFindProc("tSetOnRspQryExchangeRate").Call(t.spi, syscall.NewCallback(t.OnRspQryExchangeRate))
    t.h.MustFindProc("tSetOnRspQrySecAgentACIDMap").Call(t.spi, syscall.NewCallback(t.OnRspQrySecAgentACIDMap))
    t.h.MustFindProc("tSetOnRspQryProductExchRate").Call(t.spi, syscall.NewCallback(t.OnRspQryProductExchRate))
    t.h.MustFindProc("tSetOnRspQryProductGroup").Call(t.spi, syscall.NewCallback(t.OnRspQryProductGroup))
    t.h.MustFindProc("tSetOnRspQryMMInstrumentCommissionRate").Call(t.spi, syscall.NewCallback(t.OnRspQryMMInstrumentCommissionRate))
    t.h.MustFindProc("tSetOnRspQryMMOptionInstrCommRate").Call(t.spi, syscall.NewCallback(t.OnRspQryMMOptionInstrCommRate))
    t.h.MustFindProc("tSetOnRspQryInstrumentOrderCommRate").Call(t.spi, syscall.NewCallback(t.OnRspQryInstrumentOrderCommRate))
    t.h.MustFindProc("tSetOnRspQrySecAgentTradingAccount").Call(t.spi, syscall.NewCallback(t.OnRspQrySecAgentTradingAccount))
    t.h.MustFindProc("tSetOnRspQrySecAgentCheckMode").Call(t.spi, syscall.NewCallback(t.OnRspQrySecAgentCheckMode))
    t.h.MustFindProc("tSetOnRspQrySecAgentTradeInfo").Call(t.spi, syscall.NewCallback(t.OnRspQrySecAgentTradeInfo))
    t.h.MustFindProc("tSetOnRspQryOptionInstrTradeCost").Call(t.spi, syscall.NewCallback(t.OnRspQryOptionInstrTradeCost))
    t.h.MustFindProc("tSetOnRspQryOptionInstrCommRate").Call(t.spi, syscall.NewCallback(t.OnRspQryOptionInstrCommRate))
    t.h.MustFindProc("tSetOnRspQryExecOrder").Call(t.spi, syscall.NewCallback(t.OnRspQryExecOrder))
    t.h.MustFindProc("tSetOnRspQryForQuote").Call(t.spi, syscall.NewCallback(t.OnRspQryForQuote))
    t.h.MustFindProc("tSetOnRspQryQuote").Call(t.spi, syscall.NewCallback(t.OnRspQryQuote))
    t.h.MustFindProc("tSetOnRspQryOptionSelfClose").Call(t.spi, syscall.NewCallback(t.OnRspQryOptionSelfClose))
    t.h.MustFindProc("tSetOnRspQryInvestUnit").Call(t.spi, syscall.NewCallback(t.OnRspQryInvestUnit))
    t.h.MustFindProc("tSetOnRspQryCombInstrumentGuard").Call(t.spi, syscall.NewCallback(t.OnRspQryCombInstrumentGuard))
    t.h.MustFindProc("tSetOnRspQryCombAction").Call(t.spi, syscall.NewCallback(t.OnRspQryCombAction))
    t.h.MustFindProc("tSetOnRspQryTransferSerial").Call(t.spi, syscall.NewCallback(t.OnRspQryTransferSerial))
    t.h.MustFindProc("tSetOnRspQryAccountregister").Call(t.spi, syscall.NewCallback(t.OnRspQryAccountregister))
    t.h.MustFindProc("tSetOnRspError").Call(t.spi, syscall.NewCallback(t.OnRspError))
    t.h.MustFindProc("tSetOnRtnOrder").Call(t.spi, syscall.NewCallback(t.OnRtnOrder))
    t.h.MustFindProc("tSetOnRtnTrade").Call(t.spi, syscall.NewCallback(t.OnRtnTrade))
    t.h.MustFindProc("tSetOnErrRtnOrderInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnOrderInsert))
    t.h.MustFindProc("tSetOnErrRtnOrderAction").Call(t.spi, syscall.NewCallback(t.OnErrRtnOrderAction))
    t.h.MustFindProc("tSetOnRtnInstrumentStatus").Call(t.spi, syscall.NewCallback(t.OnRtnInstrumentStatus))
    t.h.MustFindProc("tSetOnRtnBulletin").Call(t.spi, syscall.NewCallback(t.OnRtnBulletin))
    t.h.MustFindProc("tSetOnRtnTradingNotice").Call(t.spi, syscall.NewCallback(t.OnRtnTradingNotice))
    t.h.MustFindProc("tSetOnRtnErrorConditionalOrder").Call(t.spi, syscall.NewCallback(t.OnRtnErrorConditionalOrder))
    t.h.MustFindProc("tSetOnRtnExecOrder").Call(t.spi, syscall.NewCallback(t.OnRtnExecOrder))
    t.h.MustFindProc("tSetOnErrRtnExecOrderInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnExecOrderInsert))
    t.h.MustFindProc("tSetOnErrRtnExecOrderAction").Call(t.spi, syscall.NewCallback(t.OnErrRtnExecOrderAction))
    t.h.MustFindProc("tSetOnErrRtnForQuoteInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnForQuoteInsert))
    t.h.MustFindProc("tSetOnRtnQuote").Call(t.spi, syscall.NewCallback(t.OnRtnQuote))
    t.h.MustFindProc("tSetOnErrRtnQuoteInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnQuoteInsert))
    t.h.MustFindProc("tSetOnErrRtnQuoteAction").Call(t.spi, syscall.NewCallback(t.OnErrRtnQuoteAction))
    t.h.MustFindProc("tSetOnRtnForQuoteRsp").Call(t.spi, syscall.NewCallback(t.OnRtnForQuoteRsp))
    t.h.MustFindProc("tSetOnRtnCFMMCTradingAccountToken").Call(t.spi, syscall.NewCallback(t.OnRtnCFMMCTradingAccountToken))
    t.h.MustFindProc("tSetOnErrRtnBatchOrderAction").Call(t.spi, syscall.NewCallback(t.OnErrRtnBatchOrderAction))
    t.h.MustFindProc("tSetOnRtnOptionSelfClose").Call(t.spi, syscall.NewCallback(t.OnRtnOptionSelfClose))
    t.h.MustFindProc("tSetOnErrRtnOptionSelfCloseInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnOptionSelfCloseInsert))
    t.h.MustFindProc("tSetOnErrRtnOptionSelfCloseAction").Call(t.spi, syscall.NewCallback(t.OnErrRtnOptionSelfCloseAction))
    t.h.MustFindProc("tSetOnRtnCombAction").Call(t.spi, syscall.NewCallback(t.OnRtnCombAction))
    t.h.MustFindProc("tSetOnErrRtnCombActionInsert").Call(t.spi, syscall.NewCallback(t.OnErrRtnCombActionInsert))
    t.h.MustFindProc("tSetOnRspQryContractBank").Call(t.spi, syscall.NewCallback(t.OnRspQryContractBank))
    t.h.MustFindProc("tSetOnRspQryParkedOrder").Call(t.spi, syscall.NewCallback(t.OnRspQryParkedOrder))
    t.h.MustFindProc("tSetOnRspQryParkedOrderAction").Call(t.spi, syscall.NewCallback(t.OnRspQryParkedOrderAction))
    t.h.MustFindProc("tSetOnRspQryTradingNotice").Call(t.spi, syscall.NewCallback(t.OnRspQryTradingNotice))
    t.h.MustFindProc("tSetOnRspQryBrokerTradingParams").Call(t.spi, syscall.NewCallback(t.OnRspQryBrokerTradingParams))
    t.h.MustFindProc("tSetOnRspQryBrokerTradingAlgos").Call(t.spi, syscall.NewCallback(t.OnRspQryBrokerTradingAlgos))
    t.h.MustFindProc("tSetOnRspQueryCFMMCTradingAccountToken").Call(t.spi, syscall.NewCallback(t.OnRspQueryCFMMCTradingAccountToken))
    t.h.MustFindProc("tSetOnRtnFromBankToFutureByBank").Call(t.spi, syscall.NewCallback(t.OnRtnFromBankToFutureByBank))
    t.h.MustFindProc("tSetOnRtnFromFutureToBankByBank").Call(t.spi, syscall.NewCallback(t.OnRtnFromFutureToBankByBank))
    t.h.MustFindProc("tSetOnRtnRepealFromBankToFutureByBank").Call(t.spi, syscall.NewCallback(t.OnRtnRepealFromBankToFutureByBank))
    t.h.MustFindProc("tSetOnRtnRepealFromFutureToBankByBank").Call(t.spi, syscall.NewCallback(t.OnRtnRepealFromFutureToBankByBank))
    t.h.MustFindProc("tSetOnRtnFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnFromBankToFutureByFuture))
    t.h.MustFindProc("tSetOnRtnFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnFromFutureToBankByFuture))
    t.h.MustFindProc("tSetOnRtnRepealFromBankToFutureByFutureManual").Call(t.spi, syscall.NewCallback(t.OnRtnRepealFromBankToFutureByFutureManual))
    t.h.MustFindProc("tSetOnRtnRepealFromFutureToBankByFutureManual").Call(t.spi, syscall.NewCallback(t.OnRtnRepealFromFutureToBankByFutureManual))
    t.h.MustFindProc("tSetOnRtnQueryBankBalanceByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnQueryBankBalanceByFuture))
    t.h.MustFindProc("tSetOnErrRtnBankToFutureByFuture").Call(t.spi, syscall.NewCallback(t.OnErrRtnBankToFutureByFuture))
    t.h.MustFindProc("tSetOnErrRtnFutureToBankByFuture").Call(t.spi, syscall.NewCallback(t.OnErrRtnFutureToBankByFuture))
    t.h.MustFindProc("tSetOnErrRtnRepealBankToFutureByFutureManual").Call(t.spi, syscall.NewCallback(t.OnErrRtnRepealBankToFutureByFutureManual))
    t.h.MustFindProc("tSetOnErrRtnRepealFutureToBankByFutureManual").Call(t.spi, syscall.NewCallback(t.OnErrRtnRepealFutureToBankByFutureManual))
    t.h.MustFindProc("tSetOnErrRtnQueryBankBalanceByFuture").Call(t.spi, syscall.NewCallback(t.OnErrRtnQueryBankBalanceByFuture))
    t.h.MustFindProc("tSetOnRtnRepealFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnRepealFromBankToFutureByFuture))
    t.h.MustFindProc("tSetOnRtnRepealFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(t.OnRtnRepealFromFutureToBankByFuture))
    t.h.MustFindProc("tSetOnRspFromBankToFutureByFuture").Call(t.spi, syscall.NewCallback(t.OnRspFromBankToFutureByFuture))
    t.h.MustFindProc("tSetOnRspFromFutureToBankByFuture").Call(t.spi, syscall.NewCallback(t.OnRspFromFutureToBankByFuture))
    t.h.MustFindProc("tSetOnRspQueryBankAccountMoneyByFuture").Call(t.spi, syscall.NewCallback(t.OnRspQueryBankAccountMoneyByFuture))
    t.h.MustFindProc("tSetOnRtnOpenAccountByBank").Call(t.spi, syscall.NewCallback(t.OnRtnOpenAccountByBank))
    t.h.MustFindProc("tSetOnRtnCancelAccountByBank").Call(t.spi, syscall.NewCallback(t.OnRtnCancelAccountByBank))
    t.h.MustFindProc("tSetOnRtnChangeAccountByBank").Call(t.spi, syscall.NewCallback(t.OnRtnChangeAccountByBank))
    t.h.MustFindProc("tSetOnRspQryClassifiedInstrument").Call(t.spi, syscall.NewCallback(t.OnRspQryClassifiedInstrument))
    t.h.MustFindProc("tSetOnRspQryCombPromotionParam").Call(t.spi, syscall.NewCallback(t.OnRspQryCombPromotionParam))
    t.h.MustFindProc("tSetOnRspQryRiskSettleInvstPosition").Call(t.spi, syscall.NewCallback(t.OnRspQryRiskSettleInvstPosition))
    t.h.MustFindProc("tSetOnRspQryRiskSettleProductStatus").Call(t.spi, syscall.NewCallback(t.OnRspQryRiskSettleProductStatus))
    
	return t
}

//export tFrontConnected
func (t *Trade) OnFrontConnected() uintptr {
    if t._FrontConnected != nil{
        t._FrontConnected()
    }
	return 0
}

//export tFrontDisconnected
func (t *Trade) OnFrontDisconnected(nReason int) uintptr {
    if t._FrontDisconnected != nil{
        t._FrontDisconnected(nReason)
    }
	return 0
}

//export tHeartBeatWarning
func (t *Trade) OnHeartBeatWarning(nTimeLapse int) uintptr {
    if t._HeartBeatWarning != nil{
        t._HeartBeatWarning(nTimeLapse)
    }
	return 0
}

//export tRspAuthenticate
func (t *Trade) OnRspAuthenticate(pRspAuthenticateField *ctp.CThostFtdcRspAuthenticateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspAuthenticate != nil{
        t._RspAuthenticate(pRspAuthenticateField, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspUserLogin
func (t *Trade) OnRspUserLogin(pRspUserLogin *ctp.CThostFtdcRspUserLoginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspUserLogin != nil{
        t._RspUserLogin(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspUserLogout
func (t *Trade) OnRspUserLogout(pUserLogout *ctp.CThostFtdcUserLogoutField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspUserLogout != nil{
        t._RspUserLogout(pUserLogout, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspUserPasswordUpdate
func (t *Trade) OnRspUserPasswordUpdate(pUserPasswordUpdate *ctp.CThostFtdcUserPasswordUpdateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspUserPasswordUpdate != nil{
        t._RspUserPasswordUpdate(pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspTradingAccountPasswordUpdate
func (t *Trade) OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *ctp.CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspTradingAccountPasswordUpdate != nil{
        t._RspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspUserAuthMethod
func (t *Trade) OnRspUserAuthMethod(pRspUserAuthMethod *ctp.CThostFtdcRspUserAuthMethodField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspUserAuthMethod != nil{
        t._RspUserAuthMethod(pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspGenUserCaptcha
func (t *Trade) OnRspGenUserCaptcha(pRspGenUserCaptcha *ctp.CThostFtdcRspGenUserCaptchaField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspGenUserCaptcha != nil{
        t._RspGenUserCaptcha(pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspGenUserText
func (t *Trade) OnRspGenUserText(pRspGenUserText *ctp.CThostFtdcRspGenUserTextField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspGenUserText != nil{
        t._RspGenUserText(pRspGenUserText, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspOrderInsert
func (t *Trade) OnRspOrderInsert(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspOrderInsert != nil{
        t._RspOrderInsert(pInputOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspParkedOrderInsert
func (t *Trade) OnRspParkedOrderInsert(pParkedOrder *ctp.CThostFtdcParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspParkedOrderInsert != nil{
        t._RspParkedOrderInsert(pParkedOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspParkedOrderAction
func (t *Trade) OnRspParkedOrderAction(pParkedOrderAction *ctp.CThostFtdcParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspParkedOrderAction != nil{
        t._RspParkedOrderAction(pParkedOrderAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspOrderAction
func (t *Trade) OnRspOrderAction(pInputOrderAction *ctp.CThostFtdcInputOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspOrderAction != nil{
        t._RspOrderAction(pInputOrderAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryMaxOrderVolume
func (t *Trade) OnRspQryMaxOrderVolume(pQryMaxOrderVolume *ctp.CThostFtdcQryMaxOrderVolumeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryMaxOrderVolume != nil{
        t._RspQryMaxOrderVolume(pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspSettlementInfoConfirm
func (t *Trade) OnRspSettlementInfoConfirm(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspSettlementInfoConfirm != nil{
        t._RspSettlementInfoConfirm(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspRemoveParkedOrder
func (t *Trade) OnRspRemoveParkedOrder(pRemoveParkedOrder *ctp.CThostFtdcRemoveParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspRemoveParkedOrder != nil{
        t._RspRemoveParkedOrder(pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspRemoveParkedOrderAction
func (t *Trade) OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *ctp.CThostFtdcRemoveParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspRemoveParkedOrderAction != nil{
        t._RspRemoveParkedOrderAction(pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspExecOrderInsert
func (t *Trade) OnRspExecOrderInsert(pInputExecOrder *ctp.CThostFtdcInputExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspExecOrderInsert != nil{
        t._RspExecOrderInsert(pInputExecOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspExecOrderAction
func (t *Trade) OnRspExecOrderAction(pInputExecOrderAction *ctp.CThostFtdcInputExecOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspExecOrderAction != nil{
        t._RspExecOrderAction(pInputExecOrderAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspForQuoteInsert
func (t *Trade) OnRspForQuoteInsert(pInputForQuote *ctp.CThostFtdcInputForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspForQuoteInsert != nil{
        t._RspForQuoteInsert(pInputForQuote, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQuoteInsert
func (t *Trade) OnRspQuoteInsert(pInputQuote *ctp.CThostFtdcInputQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQuoteInsert != nil{
        t._RspQuoteInsert(pInputQuote, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQuoteAction
func (t *Trade) OnRspQuoteAction(pInputQuoteAction *ctp.CThostFtdcInputQuoteActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQuoteAction != nil{
        t._RspQuoteAction(pInputQuoteAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspBatchOrderAction
func (t *Trade) OnRspBatchOrderAction(pInputBatchOrderAction *ctp.CThostFtdcInputBatchOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspBatchOrderAction != nil{
        t._RspBatchOrderAction(pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspOptionSelfCloseInsert
func (t *Trade) OnRspOptionSelfCloseInsert(pInputOptionSelfClose *ctp.CThostFtdcInputOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspOptionSelfCloseInsert != nil{
        t._RspOptionSelfCloseInsert(pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspOptionSelfCloseAction
func (t *Trade) OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *ctp.CThostFtdcInputOptionSelfCloseActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspOptionSelfCloseAction != nil{
        t._RspOptionSelfCloseAction(pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspCombActionInsert
func (t *Trade) OnRspCombActionInsert(pInputCombAction *ctp.CThostFtdcInputCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspCombActionInsert != nil{
        t._RspCombActionInsert(pInputCombAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryOrder
func (t *Trade) OnRspQryOrder(pOrder *ctp.CThostFtdcOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryOrder != nil{
        t._RspQryOrder(pOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTrade
func (t *Trade) OnRspQryTrade(pTrade *ctp.CThostFtdcTradeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTrade != nil{
        t._RspQryTrade(pTrade, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInvestorPosition
func (t *Trade) OnRspQryInvestorPosition(pInvestorPosition *ctp.CThostFtdcInvestorPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInvestorPosition != nil{
        t._RspQryInvestorPosition(pInvestorPosition, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTradingAccount
func (t *Trade) OnRspQryTradingAccount(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTradingAccount != nil{
        t._RspQryTradingAccount(pTradingAccount, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInvestor
func (t *Trade) OnRspQryInvestor(pInvestor *ctp.CThostFtdcInvestorField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInvestor != nil{
        t._RspQryInvestor(pInvestor, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTradingCode
func (t *Trade) OnRspQryTradingCode(pTradingCode *ctp.CThostFtdcTradingCodeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTradingCode != nil{
        t._RspQryTradingCode(pTradingCode, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInstrumentMarginRate
func (t *Trade) OnRspQryInstrumentMarginRate(pInstrumentMarginRate *ctp.CThostFtdcInstrumentMarginRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInstrumentMarginRate != nil{
        t._RspQryInstrumentMarginRate(pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInstrumentCommissionRate
func (t *Trade) OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *ctp.CThostFtdcInstrumentCommissionRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInstrumentCommissionRate != nil{
        t._RspQryInstrumentCommissionRate(pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryExchange
func (t *Trade) OnRspQryExchange(pExchange *ctp.CThostFtdcExchangeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryExchange != nil{
        t._RspQryExchange(pExchange, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryProduct
func (t *Trade) OnRspQryProduct(pProduct *ctp.CThostFtdcProductField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryProduct != nil{
        t._RspQryProduct(pProduct, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInstrument
func (t *Trade) OnRspQryInstrument(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInstrument != nil{
        t._RspQryInstrument(pInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryDepthMarketData
func (t *Trade) OnRspQryDepthMarketData(pDepthMarketData *ctp.CThostFtdcDepthMarketDataField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryDepthMarketData != nil{
        t._RspQryDepthMarketData(pDepthMarketData, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTraderOffer
func (t *Trade) OnRspQryTraderOffer(pTraderOffer *ctp.CThostFtdcTraderOfferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTraderOffer != nil{
        t._RspQryTraderOffer(pTraderOffer, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQrySettlementInfo
func (t *Trade) OnRspQrySettlementInfo(pSettlementInfo *ctp.CThostFtdcSettlementInfoField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQrySettlementInfo != nil{
        t._RspQrySettlementInfo(pSettlementInfo, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTransferBank
func (t *Trade) OnRspQryTransferBank(pTransferBank *ctp.CThostFtdcTransferBankField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTransferBank != nil{
        t._RspQryTransferBank(pTransferBank, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInvestorPositionDetail
func (t *Trade) OnRspQryInvestorPositionDetail(pInvestorPositionDetail *ctp.CThostFtdcInvestorPositionDetailField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInvestorPositionDetail != nil{
        t._RspQryInvestorPositionDetail(pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryNotice
func (t *Trade) OnRspQryNotice(pNotice *ctp.CThostFtdcNoticeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryNotice != nil{
        t._RspQryNotice(pNotice, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQrySettlementInfoConfirm
func (t *Trade) OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *ctp.CThostFtdcSettlementInfoConfirmField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQrySettlementInfoConfirm != nil{
        t._RspQrySettlementInfoConfirm(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInvestorPositionCombineDetail
func (t *Trade) OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *ctp.CThostFtdcInvestorPositionCombineDetailField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInvestorPositionCombineDetail != nil{
        t._RspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryCFMMCTradingAccountKey
func (t *Trade) OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *ctp.CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryCFMMCTradingAccountKey != nil{
        t._RspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryEWarrantOffset
func (t *Trade) OnRspQryEWarrantOffset(pEWarrantOffset *ctp.CThostFtdcEWarrantOffsetField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryEWarrantOffset != nil{
        t._RspQryEWarrantOffset(pEWarrantOffset, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInvestorProductGroupMargin
func (t *Trade) OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *ctp.CThostFtdcInvestorProductGroupMarginField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInvestorProductGroupMargin != nil{
        t._RspQryInvestorProductGroupMargin(pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryExchangeMarginRate
func (t *Trade) OnRspQryExchangeMarginRate(pExchangeMarginRate *ctp.CThostFtdcExchangeMarginRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryExchangeMarginRate != nil{
        t._RspQryExchangeMarginRate(pExchangeMarginRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryExchangeMarginRateAdjust
func (t *Trade) OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *ctp.CThostFtdcExchangeMarginRateAdjustField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryExchangeMarginRateAdjust != nil{
        t._RspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryExchangeRate
func (t *Trade) OnRspQryExchangeRate(pExchangeRate *ctp.CThostFtdcExchangeRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryExchangeRate != nil{
        t._RspQryExchangeRate(pExchangeRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQrySecAgentACIDMap
func (t *Trade) OnRspQrySecAgentACIDMap(pSecAgentACIDMap *ctp.CThostFtdcSecAgentACIDMapField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQrySecAgentACIDMap != nil{
        t._RspQrySecAgentACIDMap(pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryProductExchRate
func (t *Trade) OnRspQryProductExchRate(pProductExchRate *ctp.CThostFtdcProductExchRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryProductExchRate != nil{
        t._RspQryProductExchRate(pProductExchRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryProductGroup
func (t *Trade) OnRspQryProductGroup(pProductGroup *ctp.CThostFtdcProductGroupField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryProductGroup != nil{
        t._RspQryProductGroup(pProductGroup, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryMMInstrumentCommissionRate
func (t *Trade) OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *ctp.CThostFtdcMMInstrumentCommissionRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryMMInstrumentCommissionRate != nil{
        t._RspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryMMOptionInstrCommRate
func (t *Trade) OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *ctp.CThostFtdcMMOptionInstrCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryMMOptionInstrCommRate != nil{
        t._RspQryMMOptionInstrCommRate(pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInstrumentOrderCommRate
func (t *Trade) OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *ctp.CThostFtdcInstrumentOrderCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInstrumentOrderCommRate != nil{
        t._RspQryInstrumentOrderCommRate(pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQrySecAgentTradingAccount
func (t *Trade) OnRspQrySecAgentTradingAccount(pTradingAccount *ctp.CThostFtdcTradingAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQrySecAgentTradingAccount != nil{
        t._RspQrySecAgentTradingAccount(pTradingAccount, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQrySecAgentCheckMode
func (t *Trade) OnRspQrySecAgentCheckMode(pSecAgentCheckMode *ctp.CThostFtdcSecAgentCheckModeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQrySecAgentCheckMode != nil{
        t._RspQrySecAgentCheckMode(pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQrySecAgentTradeInfo
func (t *Trade) OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *ctp.CThostFtdcSecAgentTradeInfoField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQrySecAgentTradeInfo != nil{
        t._RspQrySecAgentTradeInfo(pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryOptionInstrTradeCost
func (t *Trade) OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *ctp.CThostFtdcOptionInstrTradeCostField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryOptionInstrTradeCost != nil{
        t._RspQryOptionInstrTradeCost(pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryOptionInstrCommRate
func (t *Trade) OnRspQryOptionInstrCommRate(pOptionInstrCommRate *ctp.CThostFtdcOptionInstrCommRateField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryOptionInstrCommRate != nil{
        t._RspQryOptionInstrCommRate(pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryExecOrder
func (t *Trade) OnRspQryExecOrder(pExecOrder *ctp.CThostFtdcExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryExecOrder != nil{
        t._RspQryExecOrder(pExecOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryForQuote
func (t *Trade) OnRspQryForQuote(pForQuote *ctp.CThostFtdcForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryForQuote != nil{
        t._RspQryForQuote(pForQuote, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryQuote
func (t *Trade) OnRspQryQuote(pQuote *ctp.CThostFtdcQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryQuote != nil{
        t._RspQryQuote(pQuote, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryOptionSelfClose
func (t *Trade) OnRspQryOptionSelfClose(pOptionSelfClose *ctp.CThostFtdcOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryOptionSelfClose != nil{
        t._RspQryOptionSelfClose(pOptionSelfClose, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryInvestUnit
func (t *Trade) OnRspQryInvestUnit(pInvestUnit *ctp.CThostFtdcInvestUnitField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryInvestUnit != nil{
        t._RspQryInvestUnit(pInvestUnit, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryCombInstrumentGuard
func (t *Trade) OnRspQryCombInstrumentGuard(pCombInstrumentGuard *ctp.CThostFtdcCombInstrumentGuardField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryCombInstrumentGuard != nil{
        t._RspQryCombInstrumentGuard(pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryCombAction
func (t *Trade) OnRspQryCombAction(pCombAction *ctp.CThostFtdcCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryCombAction != nil{
        t._RspQryCombAction(pCombAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTransferSerial
func (t *Trade) OnRspQryTransferSerial(pTransferSerial *ctp.CThostFtdcTransferSerialField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTransferSerial != nil{
        t._RspQryTransferSerial(pTransferSerial, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryAccountregister
func (t *Trade) OnRspQryAccountregister(pAccountregister *ctp.CThostFtdcAccountregisterField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryAccountregister != nil{
        t._RspQryAccountregister(pAccountregister, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspError
func (t *Trade) OnRspError(pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspError != nil{
        t._RspError(pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRtnOrder
func (t *Trade) OnRtnOrder(pOrder *ctp.CThostFtdcOrderField) uintptr {
    if t._RtnOrder != nil{
        t._RtnOrder(pOrder)
    }
	return 0
}

//export tRtnTrade
func (t *Trade) OnRtnTrade(pTrade *ctp.CThostFtdcTradeField) uintptr {
    if t._RtnTrade != nil{
        t._RtnTrade(pTrade)
    }
	return 0
}

//export tErrRtnOrderInsert
func (t *Trade) OnErrRtnOrderInsert(pInputOrder *ctp.CThostFtdcInputOrderField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnOrderInsert != nil{
        t._ErrRtnOrderInsert(pInputOrder, pRspInfo)
    }
	return 0
}

//export tErrRtnOrderAction
func (t *Trade) OnErrRtnOrderAction(pOrderAction *ctp.CThostFtdcOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnOrderAction != nil{
        t._ErrRtnOrderAction(pOrderAction, pRspInfo)
    }
	return 0
}

//export tRtnInstrumentStatus
func (t *Trade) OnRtnInstrumentStatus(pInstrumentStatus *ctp.CThostFtdcInstrumentStatusField) uintptr {
    if t._RtnInstrumentStatus != nil{
        t._RtnInstrumentStatus(pInstrumentStatus)
    }
	return 0
}

//export tRtnBulletin
func (t *Trade) OnRtnBulletin(pBulletin *ctp.CThostFtdcBulletinField) uintptr {
    if t._RtnBulletin != nil{
        t._RtnBulletin(pBulletin)
    }
	return 0
}

//export tRtnTradingNotice
func (t *Trade) OnRtnTradingNotice(pTradingNoticeInfo *ctp.CThostFtdcTradingNoticeInfoField) uintptr {
    if t._RtnTradingNotice != nil{
        t._RtnTradingNotice(pTradingNoticeInfo)
    }
	return 0
}

//export tRtnErrorConditionalOrder
func (t *Trade) OnRtnErrorConditionalOrder(pErrorConditionalOrder *ctp.CThostFtdcErrorConditionalOrderField) uintptr {
    if t._RtnErrorConditionalOrder != nil{
        t._RtnErrorConditionalOrder(pErrorConditionalOrder)
    }
	return 0
}

//export tRtnExecOrder
func (t *Trade) OnRtnExecOrder(pExecOrder *ctp.CThostFtdcExecOrderField) uintptr {
    if t._RtnExecOrder != nil{
        t._RtnExecOrder(pExecOrder)
    }
	return 0
}

//export tErrRtnExecOrderInsert
func (t *Trade) OnErrRtnExecOrderInsert(pInputExecOrder *ctp.CThostFtdcInputExecOrderField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnExecOrderInsert != nil{
        t._ErrRtnExecOrderInsert(pInputExecOrder, pRspInfo)
    }
	return 0
}

//export tErrRtnExecOrderAction
func (t *Trade) OnErrRtnExecOrderAction(pExecOrderAction *ctp.CThostFtdcExecOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnExecOrderAction != nil{
        t._ErrRtnExecOrderAction(pExecOrderAction, pRspInfo)
    }
	return 0
}

//export tErrRtnForQuoteInsert
func (t *Trade) OnErrRtnForQuoteInsert(pInputForQuote *ctp.CThostFtdcInputForQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnForQuoteInsert != nil{
        t._ErrRtnForQuoteInsert(pInputForQuote, pRspInfo)
    }
	return 0
}

//export tRtnQuote
func (t *Trade) OnRtnQuote(pQuote *ctp.CThostFtdcQuoteField) uintptr {
    if t._RtnQuote != nil{
        t._RtnQuote(pQuote)
    }
	return 0
}

//export tErrRtnQuoteInsert
func (t *Trade) OnErrRtnQuoteInsert(pInputQuote *ctp.CThostFtdcInputQuoteField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnQuoteInsert != nil{
        t._ErrRtnQuoteInsert(pInputQuote, pRspInfo)
    }
	return 0
}

//export tErrRtnQuoteAction
func (t *Trade) OnErrRtnQuoteAction(pQuoteAction *ctp.CThostFtdcQuoteActionField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnQuoteAction != nil{
        t._ErrRtnQuoteAction(pQuoteAction, pRspInfo)
    }
	return 0
}

//export tRtnForQuoteRsp
func (t *Trade) OnRtnForQuoteRsp(pForQuoteRsp *ctp.CThostFtdcForQuoteRspField) uintptr {
    if t._RtnForQuoteRsp != nil{
        t._RtnForQuoteRsp(pForQuoteRsp)
    }
	return 0
}

//export tRtnCFMMCTradingAccountToken
func (t *Trade) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *ctp.CThostFtdcCFMMCTradingAccountTokenField) uintptr {
    if t._RtnCFMMCTradingAccountToken != nil{
        t._RtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken)
    }
	return 0
}

//export tErrRtnBatchOrderAction
func (t *Trade) OnErrRtnBatchOrderAction(pBatchOrderAction *ctp.CThostFtdcBatchOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnBatchOrderAction != nil{
        t._ErrRtnBatchOrderAction(pBatchOrderAction, pRspInfo)
    }
	return 0
}

//export tRtnOptionSelfClose
func (t *Trade) OnRtnOptionSelfClose(pOptionSelfClose *ctp.CThostFtdcOptionSelfCloseField) uintptr {
    if t._RtnOptionSelfClose != nil{
        t._RtnOptionSelfClose(pOptionSelfClose)
    }
	return 0
}

//export tErrRtnOptionSelfCloseInsert
func (t *Trade) OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *ctp.CThostFtdcInputOptionSelfCloseField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnOptionSelfCloseInsert != nil{
        t._ErrRtnOptionSelfCloseInsert(pInputOptionSelfClose, pRspInfo)
    }
	return 0
}

//export tErrRtnOptionSelfCloseAction
func (t *Trade) OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *ctp.CThostFtdcOptionSelfCloseActionField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnOptionSelfCloseAction != nil{
        t._ErrRtnOptionSelfCloseAction(pOptionSelfCloseAction, pRspInfo)
    }
	return 0
}

//export tRtnCombAction
func (t *Trade) OnRtnCombAction(pCombAction *ctp.CThostFtdcCombActionField) uintptr {
    if t._RtnCombAction != nil{
        t._RtnCombAction(pCombAction)
    }
	return 0
}

//export tErrRtnCombActionInsert
func (t *Trade) OnErrRtnCombActionInsert(pInputCombAction *ctp.CThostFtdcInputCombActionField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnCombActionInsert != nil{
        t._ErrRtnCombActionInsert(pInputCombAction, pRspInfo)
    }
	return 0
}

//export tRspQryContractBank
func (t *Trade) OnRspQryContractBank(pContractBank *ctp.CThostFtdcContractBankField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryContractBank != nil{
        t._RspQryContractBank(pContractBank, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryParkedOrder
func (t *Trade) OnRspQryParkedOrder(pParkedOrder *ctp.CThostFtdcParkedOrderField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryParkedOrder != nil{
        t._RspQryParkedOrder(pParkedOrder, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryParkedOrderAction
func (t *Trade) OnRspQryParkedOrderAction(pParkedOrderAction *ctp.CThostFtdcParkedOrderActionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryParkedOrderAction != nil{
        t._RspQryParkedOrderAction(pParkedOrderAction, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryTradingNotice
func (t *Trade) OnRspQryTradingNotice(pTradingNotice *ctp.CThostFtdcTradingNoticeField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryTradingNotice != nil{
        t._RspQryTradingNotice(pTradingNotice, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryBrokerTradingParams
func (t *Trade) OnRspQryBrokerTradingParams(pBrokerTradingParams *ctp.CThostFtdcBrokerTradingParamsField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryBrokerTradingParams != nil{
        t._RspQryBrokerTradingParams(pBrokerTradingParams, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryBrokerTradingAlgos
func (t *Trade) OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *ctp.CThostFtdcBrokerTradingAlgosField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryBrokerTradingAlgos != nil{
        t._RspQryBrokerTradingAlgos(pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQueryCFMMCTradingAccountToken
func (t *Trade) OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *ctp.CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQueryCFMMCTradingAccountToken != nil{
        t._RspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRtnFromBankToFutureByBank
func (t *Trade) OnRtnFromBankToFutureByBank(pRspTransfer *ctp.CThostFtdcRspTransferField) uintptr {
    if t._RtnFromBankToFutureByBank != nil{
        t._RtnFromBankToFutureByBank(pRspTransfer)
    }
	return 0
}

//export tRtnFromFutureToBankByBank
func (t *Trade) OnRtnFromFutureToBankByBank(pRspTransfer *ctp.CThostFtdcRspTransferField) uintptr {
    if t._RtnFromFutureToBankByBank != nil{
        t._RtnFromFutureToBankByBank(pRspTransfer)
    }
	return 0
}

//export tRtnRepealFromBankToFutureByBank
func (t *Trade) OnRtnRepealFromBankToFutureByBank(pRspRepeal *ctp.CThostFtdcRspRepealField) uintptr {
    if t._RtnRepealFromBankToFutureByBank != nil{
        t._RtnRepealFromBankToFutureByBank(pRspRepeal)
    }
	return 0
}

//export tRtnRepealFromFutureToBankByBank
func (t *Trade) OnRtnRepealFromFutureToBankByBank(pRspRepeal *ctp.CThostFtdcRspRepealField) uintptr {
    if t._RtnRepealFromFutureToBankByBank != nil{
        t._RtnRepealFromFutureToBankByBank(pRspRepeal)
    }
	return 0
}

//export tRtnFromBankToFutureByFuture
func (t *Trade) OnRtnFromBankToFutureByFuture(pRspTransfer *ctp.CThostFtdcRspTransferField) uintptr {
    if t._RtnFromBankToFutureByFuture != nil{
        t._RtnFromBankToFutureByFuture(pRspTransfer)
    }
	return 0
}

//export tRtnFromFutureToBankByFuture
func (t *Trade) OnRtnFromFutureToBankByFuture(pRspTransfer *ctp.CThostFtdcRspTransferField) uintptr {
    if t._RtnFromFutureToBankByFuture != nil{
        t._RtnFromFutureToBankByFuture(pRspTransfer)
    }
	return 0
}

//export tRtnRepealFromBankToFutureByFutureManual
func (t *Trade) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *ctp.CThostFtdcRspRepealField) uintptr {
    if t._RtnRepealFromBankToFutureByFutureManual != nil{
        t._RtnRepealFromBankToFutureByFutureManual(pRspRepeal)
    }
	return 0
}

//export tRtnRepealFromFutureToBankByFutureManual
func (t *Trade) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *ctp.CThostFtdcRspRepealField) uintptr {
    if t._RtnRepealFromFutureToBankByFutureManual != nil{
        t._RtnRepealFromFutureToBankByFutureManual(pRspRepeal)
    }
	return 0
}

//export tRtnQueryBankBalanceByFuture
func (t *Trade) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *ctp.CThostFtdcNotifyQueryAccountField) uintptr {
    if t._RtnQueryBankBalanceByFuture != nil{
        t._RtnQueryBankBalanceByFuture(pNotifyQueryAccount)
    }
	return 0
}

//export tErrRtnBankToFutureByFuture
func (t *Trade) OnErrRtnBankToFutureByFuture(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnBankToFutureByFuture != nil{
        t._ErrRtnBankToFutureByFuture(pReqTransfer, pRspInfo)
    }
	return 0
}

//export tErrRtnFutureToBankByFuture
func (t *Trade) OnErrRtnFutureToBankByFuture(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnFutureToBankByFuture != nil{
        t._ErrRtnFutureToBankByFuture(pReqTransfer, pRspInfo)
    }
	return 0
}

//export tErrRtnRepealBankToFutureByFutureManual
func (t *Trade) OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *ctp.CThostFtdcReqRepealField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnRepealBankToFutureByFutureManual != nil{
        t._ErrRtnRepealBankToFutureByFutureManual(pReqRepeal, pRspInfo)
    }
	return 0
}

//export tErrRtnRepealFutureToBankByFutureManual
func (t *Trade) OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *ctp.CThostFtdcReqRepealField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnRepealFutureToBankByFutureManual != nil{
        t._ErrRtnRepealFutureToBankByFutureManual(pReqRepeal, pRspInfo)
    }
	return 0
}

//export tErrRtnQueryBankBalanceByFuture
func (t *Trade) OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *ctp.CThostFtdcReqQueryAccountField, pRspInfo *ctp.CThostFtdcRspInfoField) uintptr {
    if t._ErrRtnQueryBankBalanceByFuture != nil{
        t._ErrRtnQueryBankBalanceByFuture(pReqQueryAccount, pRspInfo)
    }
	return 0
}

//export tRtnRepealFromBankToFutureByFuture
func (t *Trade) OnRtnRepealFromBankToFutureByFuture(pRspRepeal *ctp.CThostFtdcRspRepealField) uintptr {
    if t._RtnRepealFromBankToFutureByFuture != nil{
        t._RtnRepealFromBankToFutureByFuture(pRspRepeal)
    }
	return 0
}

//export tRtnRepealFromFutureToBankByFuture
func (t *Trade) OnRtnRepealFromFutureToBankByFuture(pRspRepeal *ctp.CThostFtdcRspRepealField) uintptr {
    if t._RtnRepealFromFutureToBankByFuture != nil{
        t._RtnRepealFromFutureToBankByFuture(pRspRepeal)
    }
	return 0
}

//export tRspFromBankToFutureByFuture
func (t *Trade) OnRspFromBankToFutureByFuture(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspFromBankToFutureByFuture != nil{
        t._RspFromBankToFutureByFuture(pReqTransfer, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspFromFutureToBankByFuture
func (t *Trade) OnRspFromFutureToBankByFuture(pReqTransfer *ctp.CThostFtdcReqTransferField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspFromFutureToBankByFuture != nil{
        t._RspFromFutureToBankByFuture(pReqTransfer, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQueryBankAccountMoneyByFuture
func (t *Trade) OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *ctp.CThostFtdcReqQueryAccountField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQueryBankAccountMoneyByFuture != nil{
        t._RspQueryBankAccountMoneyByFuture(pReqQueryAccount, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRtnOpenAccountByBank
func (t *Trade) OnRtnOpenAccountByBank(pOpenAccount *ctp.CThostFtdcOpenAccountField) uintptr {
    if t._RtnOpenAccountByBank != nil{
        t._RtnOpenAccountByBank(pOpenAccount)
    }
	return 0
}

//export tRtnCancelAccountByBank
func (t *Trade) OnRtnCancelAccountByBank(pCancelAccount *ctp.CThostFtdcCancelAccountField) uintptr {
    if t._RtnCancelAccountByBank != nil{
        t._RtnCancelAccountByBank(pCancelAccount)
    }
	return 0
}

//export tRtnChangeAccountByBank
func (t *Trade) OnRtnChangeAccountByBank(pChangeAccount *ctp.CThostFtdcChangeAccountField) uintptr {
    if t._RtnChangeAccountByBank != nil{
        t._RtnChangeAccountByBank(pChangeAccount)
    }
	return 0
}

//export tRspQryClassifiedInstrument
func (t *Trade) OnRspQryClassifiedInstrument(pInstrument *ctp.CThostFtdcInstrumentField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryClassifiedInstrument != nil{
        t._RspQryClassifiedInstrument(pInstrument, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryCombPromotionParam
func (t *Trade) OnRspQryCombPromotionParam(pCombPromotionParam *ctp.CThostFtdcCombPromotionParamField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryCombPromotionParam != nil{
        t._RspQryCombPromotionParam(pCombPromotionParam, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryRiskSettleInvstPosition
func (t *Trade) OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *ctp.CThostFtdcRiskSettleInvstPositionField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryRiskSettleInvstPosition != nil{
        t._RspQryRiskSettleInvstPosition(pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

//export tRspQryRiskSettleProductStatus
func (t *Trade) OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *ctp.CThostFtdcRiskSettleProductStatusField, pRspInfo *ctp.CThostFtdcRspInfoField, nRequestID int, bIsLast bool) uintptr {
    if t._RspQryRiskSettleProductStatus != nil{
        t._RspQryRiskSettleProductStatus(pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast)
    }
	return 0
}

