#!/bin/bash

# Alpha-Omega Dashboard Start Script
# Starts the Alpha-Omega dashboard server

PROJECT_DIR="/Users/corelogic/satori-dev/clients/alpha-omega"
PID_FILE="$PROJECT_DIR/.server.pid"

echo "🚀 Starting Alpha-Omega Dashboard..."

cd "$PROJECT_DIR"

# Check if server is already running
if [ -f "$PID_FILE" ]; then
    PID=$(cat "$PID_FILE")
    if ps -p $PID > /dev/null 2>&1; then
        echo "⚠️  Alpha-Omega is already running (PID: $PID)"
        echo "🌐 Dashboard available at: http://localhost:9999"
        exit 1
    else
        echo "🧹 Cleaning up stale PID file..."
        rm -f "$PID_FILE"
    fi
fi

# Install dependencies if needed
if [ ! -f "go.sum" ]; then
    echo "📦 Installing Go dependencies..."
    go mod tidy
fi

# Build and start the server
echo "🔨 Building Alpha-Omega..."
go build -o alpha-omega main.go

if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi

echo "🌟 Starting Alpha-Omega server..."
nohup ./alpha-omega > server.log 2>&1 &
SERVER_PID=$!

echo $SERVER_PID > "$PID_FILE"

# Wait a moment and check if server started successfully
sleep 2
if ps -p $SERVER_PID > /dev/null 2>&1; then
    echo "✅ Alpha-Omega started successfully!"
    echo "🌐 Dashboard: http://localhost:9999"
    echo "📊 Monitoring:"
    echo "   - /Users/corelogic/satori-dev/clients"
    echo "   - /Users/corelogic/satori-dev/engagements"
    echo "   - /Users/corelogic/satori-dev/platform"
    echo "📋 PID: $SERVER_PID"
    echo "📝 Logs: $PROJECT_DIR/server.log"
else
    echo "❌ Failed to start Alpha-Omega"
    rm -f "$PID_FILE"
    exit 1
fi