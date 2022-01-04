# goctp

## 介绍

CTP 封装之 golang 版,支持 Windows Linux x64.
采用二次封装，将 C++封装成 C，并 export.
win lnx 封装逻辑相同: trade.go quote.go

## 运行

### 环境变量

```json
"tradeFront": "tcp://180.168.146.187:10130",
"quoteFront": "tcp://180.168.146.187:10131",
"loginInfo": "9999/008107/1/simnow_client_test/0000000000000000"
```

### 看穿式评测

```shell
# lnx
cp CTPv6.3.19_20200423_cp/*.so lnx/
# win
cp CTPv6.3.19_20200423_cp/*.dll win/
GOOS=windows GOARCH=amd64 go build -o ctp_auth.exe demo/main.go
tar -czf ctp_auth.tgz ctp_auth.exe ./win/*.dll
# 解压后,复制win下dll到 ctp_auth.exe 目录下
set tradeFront=tcp://180.168.146.187:10130
set quoteFront=tcp://180.168.146.187:10131
set loginInfo=9999/008106/jie123465/simnow_client_test/0000000000000000
ctp_auth.exe
```

### 示例

[main.go](https://gitee.com/haifengat/goctp/blob/v6.3.15/demo/main.go)

## QA

### operator delete(void\*, unsigned long)@CXXABI_1.3.9’未定义的引用

> 不同系统，不同版本的底层依赖不同
> 重新编译即可

```bash
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_quote.so ../generate/quote.cpp  thostmduserapi_se.so && cd ..
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_trade.so ../generate/trade.cpp  thosttraderapi_se.so && cd ..
```
