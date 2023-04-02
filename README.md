# goctp

## 介绍

CTP golang 封装, 封装分三个层次:

- 基础封装
  - 函数用法与官方一致
- 简易封装
  - 常用请求函数参数由 struct 变为基础数据类型
- 易用封装
  - 登录过程,包括基础数据查询
  - 委托与成交业务处理
  - 多种委托类型: 市价,FOK,FAK
  - 实现权益与持仓实时更新
  - 交易员模式(UserID 交易员, InvestorID 投资者)

## 交易员模式

> 普通模式中 InvestorID 与 UserID 相同
> 交易员模式中 UserID 作为交易员帐号登录后,可处理多个 InvestorID 投资者帐号的业务
