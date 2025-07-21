# MCP Manager - 使用指南

## 项目概述

MCP Manager 是一个基于 Golang + React + TDesign 开发的 MCP（Model Context Protocol）工具市场管理平台。本demo实现了 "OPENAPI文档导入与转换" 的核心功能。

## 快速启动

### 方式一：使用一键启动脚本

```bash
# 在项目根目录执行
./start-dev.sh
```

这个脚本会自动：
- 检查环境依赖（Go 1.23+、Node.js 18+）
- 启动后端服务 (http://localhost:8080)
- 安装前端依赖（如果需要）
- 启动前端服务 (http://localhost:3000)

### 方式二：分别启动

#### 启动后端
```bash
# 在项目根目录
go mod tidy
go run main.go
```

#### 启动前端
```bash
# 在 web 目录
cd web
npm install  # 首次运行需要
npm run dev
```

## 功能演示

### 1. Swagger 文档导入

访问 http://localhost:3000，默认进入 "Swagger文档导入" 页面。

#### 文件上传方式：
1. 选择 "文件上传" 方式
2. 点击上传区域或拖拽文件
3. 支持 .json、.yaml、.yml 格式
4. 系统会自动校验文档格式
5. 校验通过后可点击 "解析并保存到数据库"
6. 点击 "预览API接口" 查看解析结果

#### 文本粘贴方式：
1. 选择 "文本粘贴" 方式  
2. 在文本框中粘贴完整的 Swagger/OpenAPI 文档内容
3. 点击 "校验文档格式"
4. 校验通过后可解析并保存

#### 测试用例：
项目中提供了示例文档：`web/public/example-swagger.json`，包含用户管理和产品管理的API接口。

### 2. API 接口管理

点击左侧菜单 "API接口管理" 进入接口管理页面。

#### 功能包括：
- **接口列表展示**：表格形式显示所有解析的API接口
- **查看详情** 👁️：侧边抽屉显示接口的详细信息，包括参数、响应等
- **编辑接口** ✏️：弹窗编辑接口的基本信息（名称、描述、方法、路径、标签）
- **接口测试** ▶️：弹窗测试接口，需要提供服务器基础URL
- **删除接口** 🗑️：删除选中的接口（需要确认）

## API 接口说明

后端提供的API接口与swagger.yaml文档对应：

```
POST /api/swagger/validate/file     - 文件校验
POST /api/swagger/validate/text     - 文本校验  
POST /api/swagger/parse             - 解析并保存
GET  /api/swagger/endpoints         - 获取接口列表
GET  /api/swagger/endpoint/{id}     - 获取单个接口
PUT  /api/swagger/endpoint          - 更新接口
DELETE /api/swagger/endpoint/{id}   - 删除接口
POST /api/swagger/endpoint/test     - 测试接口
```

## 目录结构

```
mcp-manager/
├── web/                    # 前端项目
│   ├── src/
│   │   ├── pages/          # 页面组件
│   │   │   ├── SwaggerImport.tsx      # Swagger导入页面
│   │   │   └── EndpointManagement.tsx # 接口管理页面
│   │   ├── services/       # API服务层
│   │   │   ├── api.ts      # Axios配置
│   │   │   └── swagger.ts  # Swagger相关API
│   │   ├── types/          # TypeScript类型定义
│   │   │   └── swagger.ts  # Swagger相关类型
│   │   ├── App.tsx         # 主应用
│   │   ├── main.tsx        # 入口文件
│   │   └── index.css       # 全局样式
│   ├── public/
│   │   └── example-swagger.json  # 示例Swagger文档
│   ├── package.json
│   ├── vite.config.ts      # Vite配置
│   ├── tsconfig.json       # TypeScript配置
│   └── README.md           # 前端文档
├── internal/               # 后端业务逻辑
├── docs/swagger.yaml       # API文档
├── start-dev.sh           # 一键启动脚本
└── README.md              # 项目文档
```

## 技术特性

### 前端技术栈
- **React 18**: 现代React框架，支持Hooks和函数组件
- **TypeScript**: 静态类型检查，提高代码质量
- **TDesign React**: 腾讯TDesign企业级UI组件库
- **Vite**: 快速的前端构建工具
- **Axios**: HTTP客户端，支持请求/响应拦截

### 后端技术栈  
- **Golang 1.23**: 高性能后端语言
- **Gin**: Web框架
- **Swagger**: API文档自动生成
- **数据库**: 支持MySQL/PostgreSQL（配置中）

## 开发特性

### 前端特性
- **组件化开发**: 页面拆分为独立的React组件
- **类型安全**: 全程TypeScript开发，类型安全
- **响应式设计**: 适配不同屏幕尺寸
- **错误处理**: 完善的错误提示和异常处理
- **用户体验**: Loading状态、确认弹窗、实时校验等

### 开发体验
- **热重载**: 代码修改后自动刷新
- **代理配置**: 开发环境自动代理后端API
- **代码规范**: ESLint + TypeScript 代码检查
- **一键启动**: 脚本自动启动前后端服务

## 下一步计划

基于当前的基础，接下来可以扩展以下功能：

1. **MCP服务集成**
   - MCP Server 自动生成
   - 工具注册和管理
   - MCP配置管理界面

2. **用户系统**
   - 用户注册登录
   - 权限管理
   - 多租户支持

3. **工具市场**
   - 工具浏览和搜索
   - 工具分类和标签
   - 工具申请和审核

4. **监控和日志**
   - 接口调用监控
   - 错误日志记录
   - 性能指标统计

## 常见问题

### Q: 前端启动失败？
A: 确保Node.js版本 >= 18，执行 `npm install` 安装依赖。

### Q: 后端服务无法访问？  
A: 确保Go版本 >= 1.23，检查8080端口是否被占用。

### Q: Swagger解析失败？
A: 确保上传的文档符合OpenAPI 2.0或3.0规范，格式正确。

### Q: 接口测试失败？
A: 确保提供的baseUrl正确，目标服务可访问。

## 贡献说明

欢迎提交Issue和Pull Request来完善这个项目！

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支  
5. 创建Pull Request
