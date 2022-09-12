# goctp

## 介绍

CTP 封装之 golang 版,支持 Windows Linux x64.
封装代码由 [ctp_generate](https://gitee.com/haifengat/ctp_generate.git) 生成

## 运行

### 环境变量

```json
{
  "userID": "008105",
  "investorID": "",
  "password": "1",
  "brokerID": "9999",
  "appID": "simnow_client_test",
  "authCode": "看穿码",
  "tradeFront": "tcp://180.168.146.187:10202",
  "quoteFront": "tcp://180.168.146.187:10212"
}
```

### 示例

[main.go](https://gitee.com/haifengat/goctp/blob/master/demo/main.go)

## 版本切换

复制官方库文件(\_se.so \_se.dll)覆盖到 lnx win 下同名文件即可。

## QA

### operator delete(void\*, unsigned long)@CXXABI_1.3.9’未定义的引用

> 不同系统，不同版本的底层依赖不同
> 重新编译即可


## 视频资料

[golang 环境-linux](https://flowus.cn/haifengat/share/a24501e7-5510-40bf-9e70-0bbb3baca118)
[goctp 环境 -win](https://flowus.cn/haifengat/share/1d17f90b-c050-49bc-b6d8-26d08566a1e1)
[行情 demo](https://flowus.cn/haifengat/share/85dc16bb-c19e-4a1e-8e92-26da077b225d)
[交易 demo](https://flowus.cn/haifengat/share/b64721ad-ddd8-4015-8325-a54ef37ed601)
[交易 板价追单](https://flowus.cn/haifengat/share/045c5c02-c01a-4a69-a75b-e6631cd1b0a8)
[交易 追单-01](https://flowus.cn/haifengat/share/433e2f06-c0d2-4b9b-a5ca-45edda8afdd4)
[交易 追单-02](https://flowus.cn/haifengat/share/96d87afa-4123-426d-80c7-c26e006ac303)
[交易 板价止损](https://flowus.cn/haifengat/share/5ab4b345-f407-44dd-8c8a-20f70c844066)
[交易 追单止盈](https://flowus.cn/haifengat/share/f0ca84a1-8485-4c12-a4b3-a20837a914aa)
