# goctp

## 介绍

CTP 封装之 golang 版,支持 Windows Linux x64.
采用二次封装，将 C++封装成 C，并 export 所用函数.

## 软件架构

代码与 C#/PYTHON 版本逻辑相同

## 使用说明

### 源码测试

> 需卸载已安装
> `go clean -i github.com/haifengat/goctp'

### 安装

```bash
go get github.com/haifengat/goctp
```

### github 提交

```bash
sed -i "s#gitee.com#github.com#g" go.mod
sed -i "s#gitee.com#github.com#g" lnx/*.go
sed -i "s#gitee.com#github.com#g" win/*.go
sed -i "s#gitee.com#github.com#g" demo/*.go
# 提取最后一次提交的说明
export comment=$(git log --pretty=%s -1) ver=v0.2.2
git commit -am $comment
git tag -a $ver -m $comment
git push github && git push github $ver
# 恢复为gitee状态
sed -i "s#github.com#gitee.com#g" go.mod
sed -i "s#github.com#gitee.com#g" lnx/*.go
sed -i "s#github.com#gitee.com#g" win/*.go
sed -i "s#github.com#gitee.com#g" demo/*.go
```

### 示例

`https://github.com/haifengat/goctp/raw/master/demo/main.go`

### 编译

#### linux

复制所有 so 文件到系统 lib 下,或把当前路径加入 LD_LIBRARY_PATH 中.

#### windows

编译后复制 dll 到应用程序目录下即可

## QA

### linux 中 quote 与 trade 不能同时载入的问题

经测试，在 trade 创建子目录 test_quote 并载入 quote 测试代码可行。（test_quote 下放 quote 代码，test_trade 下放 trade 代码，亦报错）

#### 解决

> 从 python 项目中复制 quote.h quote.cpp 过来,修改所有函数和回调函数,增加前缀 q. 重新编译 ctp_quote.so

```bash
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_quote.so ../generate/quote.cpp  thostmduserapi_se.so && cd ..
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_trade.so ../generate/trade.cpp  thosttraderapi_se.so && cd ..
```

### VSCode launch.json 配置

```json
{
  // 使用 IntelliSense 了解相关属性。
  // 悬停以查看现有属性的描述。
  // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/demo/",
      "cwd": "${workspaceFolder}",
      "env": { "LD_LIBRARY_PATH": "${workspaceFolder}/go_ctp/lib64/" },
      "args": []
    }
  ]
}
```
