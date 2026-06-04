#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REDIS_SERVER="/Volumes/SSD4T/dev/redis/bin/redis-server"
REDIS_CONF="/Volumes/SSD4T/dev/redis/redis.conf"
REDIS_PIDFILE="/Volumes/SSD4T/dev/redis/redis.pid"
PIDS=()
STARTED_REDIS=false

cleanup() {
  echo ""
  echo "正在停止服务..."
  for pid in "${PIDS[@]}"; do
    kill "$pid" 2>/dev/null || true
  done
  if $STARTED_REDIS && [[ -f "$REDIS_PIDFILE" ]]; then
    kill "$(cat "$REDIS_PIDFILE")" 2>/dev/null || true
  fi
  wait 2>/dev/null || true
  echo "已停止"
  exit 0
}

trap cleanup SIGINT SIGTERM

if [[ -f "$REDIS_PIDFILE" ]] && kill -0 "$(cat "$REDIS_PIDFILE")" 2>/dev/null; then
  echo "Redis 已在运行 (PID: $(cat "$REDIS_PIDFILE"))"
else
  echo "启动 Redis..."
  "$REDIS_SERVER" "$REDIS_CONF"
  STARTED_REDIS=true
  echo "Redis 已启动 (PID: $(cat "$REDIS_PIDFILE"))"
fi

echo "启动后端 (server)..."
(cd "$ROOT_DIR/server" && gf run main.go) &
PIDS+=($!)

echo "启动前端 (web)..."
(cd "$ROOT_DIR/web" && pnpm dev) &
PIDS+=($!)

echo ""
echo "后端 PID: ${PIDS[0]}"
echo "前端 PID: ${PIDS[1]}"
echo "按 Ctrl+C 停止所有服务"
echo ""

wait
