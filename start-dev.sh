#!/bin/bash

# MCP Manager 项目启动脚本

echo "🎯 MCP Manager 开发环境启动"
echo "=================================="

# 检查Go环境
if ! command -v go &> /dev/null; then
    echo "❌ 未找到Go环境，请先安装Go 1.23+"
    exit 1
fi

# 检查Node.js环境
if ! command -v node &> /dev/null; then
    echo "❌ 未找到Node.js环境，请先安装Node.js 18+"
    exit 1
fi

echo "✅ 环境检查通过"
echo ""

# 启动后端服务（后台运行）
echo "🔧 启动后端服务..."
cd /Users/wanghao/Desktop/github/go/mcp-manager
go mod tidy
go run main.go &
BACKEND_PID=$!
echo "后端服务PID: $BACKEND_PID"
echo "后端地址: http://localhost:8080"

# 等待后端服务启动
sleep 3

# 启动前端服务
echo ""
echo "🎨 启动前端服务..."
cd /Users/wanghao/Desktop/github/go/mcp-manager/web

# 安装前端依赖（如果需要）
if [ ! -d "node_modules" ]; then
    echo "📦 安装前端依赖..."
    npm install
fi

echo "前端地址: http://localhost:3000"
echo ""
echo "🎉 所有服务已启动！"
echo "=================================="
echo "前端: http://localhost:3000"
echo "后端: http://localhost:8080"  
echo "API文档: http://localhost:8080/swagger/index.html"
echo ""
echo "按 Ctrl+C 停止所有服务"

# 启动前端（前台运行）
npm run dev

# 清理后台进程
trap "echo ''; echo '🛑 正在停止服务...'; kill $BACKEND_PID 2>/dev/null; exit 0" INT TERM
