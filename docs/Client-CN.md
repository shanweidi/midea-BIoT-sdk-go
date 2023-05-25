# 客户端
您可以同时创建多个不同的客户端，每个客户端都可以有独立的配置。

## 客户端类型
作为iBuilding的主推接入方式，目前版本的sdk只支持以MQTT的方式接入云平台并与之通信。实际使用，应采用带配置的API创建客户端。
> 示例代码：创建一个客户端。


```go
package main

import (
        "github.com/shanweidi/midea-BIoT-sdk-go/sdk"
	"github.com/shanweidi/midea-BIoT-sdk-go/services"
)

config := sdk.NewConfig().WithClientId("sdk_test").WithServerUri("mqtt://127.0.0.1:1883")
client, err := services.NewClientWithOptions(config)
```

## 客户端配置

| 属性 | 描述 |默认 |
| -------- | -------------- | -------- |
| GwType  | 网关类型   | 无默认，必需指定 |
| GwSn   | 网关SN | 无默认，必需指定 |
| Key    | 注册密钥   | 无默认，必需指定 |
| ProductKey    | 产品密钥   | 无默认|
| ServerUri    | emqx地址   | 无默认，必需指定 |
| ClientId    | 客户端ID   | 无默认，必需指定 |
| Protocol    | 协议   | MQTT |
| GoRoutinePoolSize    | 最大并发数   | 5 |
| MaxTaskQueueSize    | 可缓存的最大任务数  | 1000 |
| EnableAsync    | 是否开启异步  | 不开启 |
| Qos    | Qos等级  | 0 |
| Username    | 用户名  | 默认为空 |
| Password    | 密码  | 默认为空 |
| LogLevel    | 日志级别  | INFO |
| KeepAlive    | 心跳间隔  | 60秒 |
| Timeout    | 超时时间  | 无默认 |
