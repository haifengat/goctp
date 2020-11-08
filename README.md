# goctp

### 介绍
CTP封装之golang版,支持Windows Linux x64.
采用二次封装，将C++封装成C，并export所用函数.

### 软件架构
代码与C#/PYTHON版本逻辑相同

### 使用说明
#### 安装
```
go get github.com/haifengat/goctp
```

#### 示例
`https://github.com/haifengat/goctp/raw/master/demo/main.go`

#### 编译
##### linux
复制所有 so文件到系统 lib下,或把当前路径加入 LD_LIBRARY_PATH 中.
##### windows
编译后复制 dll到应用程序目录下即可

### QA
#### linux中quote与trade不能同时载入的问题
经测试，在trade创建子目录test_quote并载入quote测试代码可行。（test_quote下放quote代码，test_trade下放trade代码，亦报错）
##### 解决
> 从python项目中复制quote.h quote.cpp过来,修改所有函数名,离开前缀q. 重新编译ctp_quote.so
`cd go_ctp_lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_quote.so ../../ctp_c/quote.cpp  thostmduserapi_se.so && cd ..`

#### VSCode launch.json 配置
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
            "env": {"LD_LIBRARY_PATH":"${workspaceFolder}/go_ctp/lib64/"},
            "args": []
        }
    ]
}
```