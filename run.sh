#!/bin/bash

kill_process_on_port() {
  PORT=$1
  PIDS=$(lsof -t -i:$PORT)
  if [ -n "$PIDS" ]; then
    for PID in $PIDS; do
      echo "Killing process $PID on port $PORT"
      kill $PID
    done
  fi
}

kill_process_on_port 8080
kill_process_on_port 3000

cd backend
go install github.com/air-verse/air@latest
air &

GO_PID=$!
echo "Started Go server with PID $GO_PID"

cd ..
npm start

wait $GO_PID
