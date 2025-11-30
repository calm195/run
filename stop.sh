#!/bin/bash

APP_NAME="run"
PID_FILE="/var/run/${APP_NAME}.pid"

if [ ! -f "$PID_FILE" ]; then
  echo "❌ $APP_NAME 未运行（PID 文件不存在）"
  exit 1
fi

PID=$(cat "$PID_FILE")
if kill -0 "$PID" 2>/dev/null; then
  kill "$PID"
  rm -f "$PID_FILE"
  echo "✅ $APP_NAME 已停止 (PID: $PID)"
else
  echo "⚠️ 进程已结束，清理 PID 文件"
  rm -f "$PID_FILE"
fi
