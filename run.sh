#!/bin/bash
kill $(lsof -t -i:8080)
kill $(lsof -t -i:3000)
cd backend
go mod tidy
go run main.go &

GO_PID=$!

cd ..
npm start
wait $GO_PID