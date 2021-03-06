# goctp

### 介绍
CTP封装之golang版,支持Windows Linux x64.
采用二次封装，将C++封装成C，并export所用函数.

### 软件架构
代码与C#/PYTHON版本逻辑相同

### 使用说明
#### 源码测试
> 需卸载已安装
`go clean -i github.com/haifengat/goctp'
#### 安装
```
go get github.com/haifengat/goctp
```
#### 提交(默认gitee)
```bash
# $1 comment $2 vn.n.n tag
sed -i "s#github.com#gitee.com#g" go.mod
sed -i "s#github.com#gitee.com#g" lnx/*.go
sed -i "s#github.com#gitee.com#g" win/*.go
sed -i "s#github.com#gitee.com#g" demo/*.go
git commit -am "$1"
git push gitee
git tag -a $2 -m "$1"
git push gitee $2
```
#### github 提交
```bash
sed -i "s#gitee.com#github.com#g" go.mod
sed -i "s#gitee.com#github.com#g" lnx/*.go
sed -i "s#gitee.com#github.com#g" win/*.go
sed -i "s#gitee.com#github.com#g" demo/*.go
git commit -am $1
git push github
git tag -a $2
git push github $2
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
> 从python项目中复制quote.h quote.cpp过来,修改所有函数和回调函数,增加前缀q. 重新编译ctp_quote.so
```bash
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_quote.so ../generate/quote.cpp  thostmduserapi_se.so && cd ..
cd lnx && g++ -shared -fPIC -Wl,-rpath . -o ./libctp_trade.so ../generate/trade.cpp  thosttraderapi_se.so && cd ..
```

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