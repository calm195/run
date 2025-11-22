# Run

一个基于 Go 语言开发的跑步性能追踪系统，使用 Gin 框架、PostgreSQL 数据库和 GORM ORM。

## 项目概述

这是一个用于追踪和管理跑步性能数据的 Web 应用程序。系统提供了跑步记录管理功能。

## 功能特性

- 跑步记录的创建、读取、更新和删除 (CRUD)
- 数据库集成 (PostgreSQL)
- 配置管理 (Viper)
- 日志记录 (Zap)

## 技术栈

- **语言**: Go (Golang)
- **Web 框架**: Gin
- **数据库**: PostgreSQL
- **ORM**: GORM
- **配置管理**: Viper
- **日志**: Zap
- **路由**: Gin
- **热重载**: air

## 项目架构

```
├── api/                    # API 控制器
├── config/                 # 配置结构体
├── core/                   # 核心功能
│   ├── constant.go
│   ├── cors.go
│   ├── cutter.go
│   ├── gin.go
│   ├── router.go
│   ├── viper.go
│   ├── zap.go
│   └── zap_core.go
├── env/                    # 环境配置
├── global/                 # 全局变量
├── models/                 # 数据模型
├── orm/                    # ORM 配置
├── router/                 # 路由定义
├── service/                # 业务逻辑
├── util/                   # 工具函数
└── main.go                 # 应用入口
```

## 环境要求

- Go 1.19+
- PostgreSQL 数据库

## 安装步骤

1. 克隆项目：

   ```bash
   git clone <repository-url>
   ```

2. 安装依赖：

   ```bash
   go mod tidy
   ```

3. 配置信息：
    - `/env/dev/config.yaml`样例
    - 根据需要修改数据库连接信息

4. 运行应用：

   ```bash
   go run main.go
   ```

## 配置

系统使用 Viper 进行配置管理，支持命令行参数切换配置文件

- `-f filename` 用于指定配置文件，默认为 `core.DefaultConfigFileName`
- `-p` 用于开启发布模式，不需要传值，读取`/env/prod/config.yaml`

主要配置项包括：
- 数据库连接信息
- 服务器端口
- 日志级别
- CORS 设置

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 待办事项 (TODO)

### 核心功能改进
- [ ] 实现用户认证和授权系统
- [ ] 实现缓存机制
- [ ] 常量入库 - 运动类型、分级成绩
- [ ] 成绩排名、成绩状态、用户绑定
- [ ] 文件解析、导出文件 csv

### 数据库相关
- [ ] 添加对多种数据库类型的支持 (MySQL, SQLite)
- [ ] 实现数据库迁移功能
- [ ] 添加数据备份和恢复功能

### 部署和运维
- [ ] 实现 CI/CD 流水线

### 用户体验
- [ ] 实现数据导出功能 (CSV, Excel)
- [ ] 添加实时性能追踪功能
- [ ] 添加数据可视化功能

### 安全性
- [ ] 实现 API 版本控制
- [ ] 实现安全头设置 (CSP, XSS protection)

## 许可证

该项目采用 [MIT 许可证](LICENSE)。

## 联系方式

如有问题或建议，请提交 Issue