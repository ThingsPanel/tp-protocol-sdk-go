# TP Protocol SDK Go

ThingsPanel Protocol SDK for Go，用于快速开发ThingsPanel插件。

## 功能特性

- 设备配置管理
- 服务接入管理
- MQTT消息通信
- HTTP回调处理

## 安装

```bash
go get github.com/ThingsPanel/tp-protocol-sdk-go
```

## 快速开始

### HTTP回调服务

```go
package main

import (
    "log"
    "os"

    "github.com/ThingsPanel/tp-protocol-sdk-go/handler"
)

func main() {
    // 创建处理器
    h := handler.NewHandler(handler.HandlerConfig{
        Logger: log.New(os.Stdout, "[TP] ", log.LstdFlags),
    })

    // 设置表单配置处理函数
    h.SetFormConfigHandler(func(req *handler.GetFormConfigRequest) (interface{}, error) {
        return map[string]interface{}{
            "fields": []map[string]interface{}{
                {
                    "name":  "host",
                    "type":  "string",
                    "label": "服务器地址",
                },
            },
        }, nil
    })

    // 启动HTTP服务
    h.Start(":8080")
}
```

### MQTT客户端

```go
package main

import (
    "log"

    tp "github.com/ThingsPanel/tp-protocol-sdk-go/client"
)

func main() {
    // 创建客户端
    client := tp.NewClient("tcp://localhost:1883")

    // 连接MQTT服务器
    err := client.Connect()
    if err != nil {
        log.Fatal(err)
    }

    // 发送设备状态
    client.SendStatus("device-001", "1")
}
```

## API说明

### HTTP回调接口

- `/api/v1/form/config` - 获取表单配置
- `/api/v1/device/disconnect` - 设备断开通知
- `/api/v1/plugin/notification` - 事件通知
- `/api/v1/plugin/device/list` - 获取设备列表

### MQTT主题

- `devices/status/{device_id}` - 设备状态上报
- `plugin/{服务标识符}/` - 设备数据上报主题前缀，后面跟平台规范的直连设备主题
- `plugin/{服务标识符}/#` - 订阅平台数据主题，# 位置会是平台规范的直连设备订阅主题

## 目录结构

```text
tp-protocol-sdk-go/
├── client/       - 客户端实现
├── handler/      - HTTP回调处理
├── types/        - 数据类型定义
└── examples/     - 使用示例
```

## 开发文档

更多详细信息请参考[开发文档](https://docs.thingspanel.cn/docs/protocol-sdk-go/)

## License

This project is licensed under the GNU Affero General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
