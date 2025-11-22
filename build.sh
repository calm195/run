#!/bin/bash

set -e

# === 配置区 ===
APP_NAME="run"
MAIN_FILE="."
OUTPUT_DIR="./bin"
ENV_SOURCE_DIR="./env"

REMOTE_USER="root"
REMOTE_HOST="kurous.cn"
REMOTE_DIR="/root/run"

VERSION="v1.0.0"
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
LD_FLAGS="-s -w -X main.version=${VERSION} -X main.buildTime=${BUILD_TIME}"

# 创建输出目录
mkdir -p "${OUTPUT_DIR}"

# === 1. 编译 Linux/amd64 ===
echo "🚀 编译 Linux/amd64..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build \
  -ldflags="${LD_FLAGS}" \
  -o "${OUTPUT_DIR}/${APP_NAME}" \
  "${MAIN_FILE}"

echo "✅ 二进制构建完成"

# === 2. 复制配置文件 ===
if [ -d "${ENV_SOURCE_DIR}" ]; then
  echo "📁 复制配置目录..."
  rm -rf "${OUTPUT_DIR}/env"
  cp -r "${ENV_SOURCE_DIR}" "${OUTPUT_DIR}/"
else
  echo "⚠️ 配置目录 ${ENV_SOURCE_DIR} 不存在，跳过"
fi

# === 3. 上传到服务器 ===
echo "📡 正在上传到 ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR} ..."

# 确保远程目录存在
ssh "${REMOTE_USER}@${REMOTE_HOST}" "mkdir -p ${REMOTE_DIR}"

# 上传整个 bin/ 内容（覆盖）
scp -r "${OUTPUT_DIR}/"* "${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_DIR}/"

echo "✅ 上传完成！"
echo "💡 登录服务器后，可执行："
echo "   cd /root/run && ./run"