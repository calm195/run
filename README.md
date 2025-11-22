# 跑步性能追踪系统 - 未完成

一个基于 Go 语言开发的跑步性能追踪系统，使用 Gin 框架、PostgreSQL 数据库和 GORM ORM。

## 项目概述

这是一个用于追踪和管理跑步性能数据的 Web 应用程序。系统提供了用户注册/登录、跑步记录管理、数据分析和统计等功能。

## 功能特性

- 用户注册和登录
- 跑步记录的创建、读取、更新和删除 (CRUD)
- 跑步数据统计和分析
- RESTful API 接口
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

## 项目架构

```
├── api/                    # API 控制器
│   ├── enter.go
│   ├── game.go
│   └── record.go
├── config/                 # 配置结构体
│   ├── enter.go
│   ├── pgsql.go
│   ├── system.go
│   └── zap.go
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
│   ├── base.go
│   ├── constant/
│   ├── module.go
│   ├── request/
│   └── response/
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

3. 配置环境变量：
    - 复制 `.env.example` 为 `.env`
    - 根据需要修改数据库连接信息

4. 运行应用：
   ```bash
   go run main.go
   ```

## 配置

系统使用 Viper 进行配置管理，支持多种配置源：

- 环境变量
- 配置文件
- 命令行参数

主要配置项包括：
- 数据库连接信息
- 服务器端口
- 日志级别
- CORS 设置

## API 接口

### 用户相关接口

- `POST /api/v1/user/register` - 用户注册
- `POST /api/v1/user/login` - 用户登录
- `GET /api/v1/user/profile` - 获取用户信息

### 跑步记录接口

- `POST /api/v1/record` - 创建跑步记录
- `GET /api/v1/record/{id}` - 获取单条记录
- `PUT /api/v1/record/{id}` - 更新记录
- `DELETE /api/v1/record/{id}` - 删除记录
- `GET /api/v1/records` - 获取记录列表

### 游戏相关接口

- `POST /api/v1/game` - 创建游戏
- `GET /api/v1/game/{id}` - 获取游戏信息
- `PUT /api/v1/game/{id}` - 更新游戏
- `DELETE /api/v1/game/{id}` - 删除游戏

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 待办事项 (TODO)

### 核心功能改进
- [ ] 添加完整的单元测试和集成测试
- [ ] 实现用户认证和授权系统
- [ ] 添加 API 文档 (Swagger)
- [ ] 实现数据验证和错误处理
- [ ] 添加请求限流功能
- [ ] 实现缓存机制

### 数据库相关
- [ ] 添加对多种数据库类型的支持 (MySQL, SQLite)
- [ ] 实现数据库迁移功能
- [ ] 添加数据备份和恢复功能

### 部署和运维
- [ ] 创建 Docker 配置文件
- [ ] 实现 CI/CD 流水线
- [ ] 添加监控和告警系统

### 用户体验
- [ ] 实现数据导出功能 (CSV, Excel)
- [ ] 添加实时性能追踪功能
- [ ] 创建 Web 界面 (前端)
- [ ] 添加数据可视化功能

### 安全性
- [ ] 实现 API 版本控制
- [ ] 添加输入验证和过滤
- [ ] 实现安全头设置 (CSP, XSS protection)

### 扩展功能
- [ ] 添加从健身设备导入数据的功能
- [ ] 实现数据分析和报告生成
- [ ] 添加社交功能 (好友系统、排行榜)
- rfc3339 时间格式

## 许可证

该项目采用 [MIT 许可证](LICENSE)。

## 联系方式

如有问题或建议，请提交 Issue