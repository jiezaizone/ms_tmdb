#!/bin/sh
set -eu

CONFIG_FILE="${TMDB_CONFIG_FILE:-/app/etc/tmdb.yaml}"

/app/tmdb -f "${CONFIG_FILE}" &
backend_pid="$!"

cleanup() {
  kill "${backend_pid}" >/dev/null 2>&1 || true
  if [ -n "${nginx_pid:-}" ]; then
    kill "${nginx_pid}" >/dev/null 2>&1 || true
  fi
}

trap cleanup INT TERM EXIT

nginx -g 'daemon off;' &
nginx_pid="$!"

while true; do
  if ! kill -0 "${backend_pid}" >/dev/null 2>&1; then
    echo "backend exited unexpectedly" >&2
    exit 1
  fi

  if ! kill -0 "${nginx_pid}" >/dev/null 2>&1; then
    wait "${nginx_pid}"
    exit $?
  fi

  sleep 1
done
