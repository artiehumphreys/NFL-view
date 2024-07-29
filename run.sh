#!/bin/bash
kill -9 15008
cd backend
go mod tidy
go run main.go &

GO_PID=$!

cd ..
npm start
wait $GO_PID