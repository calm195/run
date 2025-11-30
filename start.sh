#!/bin/bash

# start.sh - 静默启动 Go 应用（无 systemd）

APP_NAME="run"
APP_PATH="/root/run/run"
PID_FILE="/var/run/${APP_NAME}.pid"
LOG_FILE="/var/log/${APP_NAME}.log"

# 检查是否已在运行
if [ -f "$PID_FILE" ]; then
  PID=$(cat "$PID_FILE")
  if kill -0 "$PID" 2>/dev/null; then
    echo "⚠️  $APP_NAME 已在运行 (PID: $PID)"
    exit 1
  else
    echo "🧹 清理残留 PID 文件"
    rm -f "$PID_FILE"
  fi
fi

# 启动应用（完全静默）
echo "🚀 静默启动 $APP_NAME ..."
nohup "$APP_PATH" -p > "$LOG_FILE" 2>&1 &
echo $! > "$PID_FILE"

echo "✅ $APP_NAME 已在后台启动，PID: $(cat $PID_FILE)"
echo "📝 日志路径: $LOG_FILE"
