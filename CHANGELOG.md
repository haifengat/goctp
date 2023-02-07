# 更新记录

## v0.0.2

新增: trade.go 交易接口
新增: trade.h trade.cpp 交易接口
新增: quote.go 封装
新增: demo/main.go 行情示例
新增: 完整封装 Quote 对应的 quote.h quote.cpp
新增: struct.go 生成器
优化: datatype 生成器, 模板中使用 $ 全局变量
新增: dataType.go 生成器
更新: 接口文档编码 gb2312->utf8

## v0.0.1

新增: 响应函数(OnRspUserLogin) 参数 struct CThostFtdcRspInfoField
新增: 响应函数(OnRspUserLogin)
新增: 响应函数(OnFrontConnected)
新增: 调用函数(CreateFtdcMdSpi,RegisterSpi,SetOnFrontConnected)
新增: main.go 中 -L 和 -rpath 参数指向 lib 目录,与 make.sh 一致
新增: make.sh 编译 so
新增: 调用函数(CreateFtdcMdApi,RegisterFront,Init)
