#!/bin/bash
kill_process_on_port() {
  PORT=$1
  PID=$(lsof -t -i:$PORT)
  if [ -n "$PID" ]; then
    echo "Killing process $PID on port $PORT"
    kill $PID
  fi
}
kill_process_on_port 8080
kill_process_on_port 3000

cd backend
go run main.go &

GO_PID=$!
echo "Started Go server with PID $GO_PID"

cd ..
npm start

wait $GO_PID