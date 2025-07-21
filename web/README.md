# MCP Manager Web Frontend

基于 React + TypeScript + TDesign 开发的 MCP 工具市场管理平台前端。

## 功能特性

### 已实现功能
1. **Swagger文档导入**
   - 支持文件上传和文本粘贴两种方式
   - 实时校验文档格式和内容
   - 解析并保存API接口到数据库
   - 预览解析后的接口列表

2. **API接口管理**
   - 接口列表展示（支持分页）
   - 接口详情查看
   - 接口信息编辑
   - 接口删除操作
   - 接口测试功能

### 技术栈
- **前端框架**: React 18
- **类型系统**: TypeScript
- **UI组件库**: TDesign React
- **图标库**: TDesign Icons
- **HTTP客户端**: Axios
- **构建工具**: Vite
- **代码规范**: ESLint

## 项目结构

```
web/
├── src/
│   ├── pages/              # 页面组件
│   │   ├── SwaggerImport.tsx    # Swagger文档导入页面
│   │   └── EndpointManagement.tsx # API接口管理页面
│   ├── services/           # 服务层
│   │   ├── api.ts          # Axios配置
│   │   └── swagger.ts      # Swagger相关API服务
│   ├── types/              # 类型定义
│   │   └── swagger.ts      # Swagger相关类型定义
│   ├── App.tsx             # 主应用组件
│   ├── main.tsx            # 应用入口
│   └── index.css           # 全局样式
├── package.json
├── vite.config.ts
├── tsconfig.json
└── README.md
```

## 快速开始

### 安装依赖
```bash
cd web
npm install
```

### 开发模式
```bash
npm run dev
```
项目将运行在 http://localhost:3000

### 构建生产版本
```bash
npm run build
```

### 预览构建结果
```bash
npm run preview
```

## 开发配置

### 代理配置
开发环境下，所有 `/api` 开头的请求会被代理到后端服务 `http://localhost:8080`。

### 接口对接
前端与后端 API 的对应关系：

1. **Swagger校验 (文件)**: `POST /api/swagger/validate/file`
2. **Swagger校验 (文本)**: `POST /api/swagger/validate/text`
3. **解析并保存**: `POST /api/swagger/parse`
4. **获取接口列表**: `GET /api/swagger/endpoints?swagger_id={id}`
5. **获取单个接口**: `GET /api/swagger/endpoint/{id}`
6. **更新接口**: `PUT /api/swagger/endpoint`
7. **删除接口**: `DELETE /api/swagger/endpoint/{id}`
8. **测试接口**: `POST /api/swagger/endpoint/test?base_url={url}`

## 页面功能说明

### Swagger文档导入页面
- **导入方式选择**: 支持文件上传或文本粘贴
- **实时校验**: 上传或粘贴后立即校验文档格式
- **解析保存**: 校验通过后可解析并保存所有API接口
- **接口预览**: 显示解析出的接口列表，包括方法、路径、描述等信息

### API接口管理页面
- **接口列表**: 表格形式展示所有API接口，支持分页
- **操作功能**:
  - 👁️ 查看：打开侧边抽屉显示接口详细信息
  - ✏️ 编辑：弹窗编辑接口基本信息
  - ▶️ 测试：弹窗测试接口，需要提供服务器基础URL
  - 🗑️ 删除：删除接口（需要确认）

## 待实现功能

1. **MCP服务集成**
   - MCP服务器生成
   - 工具注册管理
   - MCP配置管理

2. **用户权限管理**
   - 用户登录注册
   - 权限控制
   - 角色管理

3. **工具市场**
   - 工具浏览
   - 工具搜索
   - 工具申请

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。
