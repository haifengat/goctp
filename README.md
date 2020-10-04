# hf_go_ctp

#### 介绍
CTP封装之golang版,支持Windows x86/x64 Linux x64.
采用二次封装，将C++封装成C，并export所用函数.

#### 软件架构
代码与C#/PYTHON版本逻辑相同

#### 使用说明

##### linux 使用
1. 基础查询完成
2. OnRtnOrder旧数据完成
3. 响应中相比CTP,增加了OnCancel,以处理委托被撤单.

#### quote与trade不能同时载入的问题
经测试，在trade创建子目录test_quote并载入quote测试代码可行。（test_quote下放quote代码，test_trade下放trade代码，亦报错）

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