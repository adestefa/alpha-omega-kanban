#!/bin/bash

# Alpha-Omega Dashboard Stop Script
# Stops the Alpha-Omega dashboard server

PROJECT_DIR="/Users/corelogic/satori-dev/clients/alpha-omega"
PID_FILE="$PROJECT_DIR/.server.pid"

echo "🛑 Stopping Alpha-Omega Dashboard..."

cd "$PROJECT_DIR"

if [ ! -f "$PID_FILE" ]; then
    echo "⚠️  Alpha-Omega is not running (no PID file found)"
    exit 1
fi

PID=$(cat "$PID_FILE")

if ps -p $PID > /dev/null 2>&1; then
    echo "🔄 Stopping Alpha-Omega (PID: $PID)..."
    kill $PID
    
    # Wait for graceful shutdown
    sleep 3
    
    if ps -p $PID > /dev/null 2>&1; then
        echo "⚠️  Graceful shutdown failed, forcing..."
        kill -9 $PID
    fi
    
    rm -f "$PID_FILE"
    echo "✅ Alpha-Omega stopped successfully"
else
    echo "⚠️  Process not found, cleaning up PID file..."
    rm -f "$PID_FILE"
fi