#!/bin/bash

# MCP Manager 前端启动脚本

echo "🚀 启动 MCP Manager 前端项目..."

# 检查是否存在 node_modules
if [ ! -d "node_modules" ]; then
    echo "📦 安装项目依赖..."
    npm install
fi

# 启动开发服务器
echo "🌟 启动开发服务器..."
echo "前端地址: http://localhost:3000"
echo "API代理: http://localhost:8080"
echo ""
echo "请确保后端服务已在 8080 端口启动"
echo ""

npm run dev
