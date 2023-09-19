# 开发帮助
## 目录结构说明
tp-protocol-sdk-go/
├── client.go          # SDK客户端和主要逻辑
├── models/           # 存放数据模型的目录
│   ├── device.go     
│   └── ...
├── api/              # API层，处理HTTP请求和响应
│   ├── handler.go   
│   └── ...
├── utils/           # 用于存放工具函数的目录
│   └── ...
├── examples/        # 示例代码目录
│   └── main.go
├── README.md        # 项目README文件，提供基本的项目信息和使用示例
└── go.mod and go.sum # Go模块文件
