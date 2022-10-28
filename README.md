# fwds

## 简介

Fitz's web development scaffolding.

fitz的Web开发脚手架,`fwds`是进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

Go 语言中常用的 API 风格是 RPC 和 REST，常用的媒体类型是 JSON、XML 和 Protobuf。

将支持在 Go API 开发中常用组合 `gRPC + Protobuf` (更适合调用频繁的微服务场景) 和 `REST + JSON`。

集成组件：
1. 支持 [jwt](https://github.com/dgrijalva/jwt-go) 接口鉴权
1. 支持 sms和email 的封装,开箱即用.
1. 支持 [cors](https://github.com/rs/cors) 接口跨域 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成  
1. 支持 [zap](https://go.uber.org/zap) 日志收集
1. 支持 [viper](https://github.com/spf13/viper) 配置文件解析
1. 支持 [gorm](https://gorm.io/gorm) 数据库组件
1. 支持 [rate](https://golang.org/x/time/rate) 接口限流
1. 支持 panic 异常时邮件通知


1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录

1. 支持 [GraphQL](https://github.com/99designs/gqlgen) 查询语言
1. 支持 trace 项目内部链路追踪
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析

1. 支持 errno 统一定义错误码

1. 支持 [go-redis](https://github.com/go-redis/redis/v8) 组件
1. 支持 RESTful API 返回值规范
1. 支持 生成数据表 CURD、控制器方法 等代码生成器

 
## 目录

```
.
├── app   
├── boot
├── cmd
├── config
├── deploy
├── docs
├── fwds
├── internal
├── pkg
├── scripts
└── storage
```
