#!/bin/bash
cd backend
go mod tidy
go run main.go &

# Change back to the root directory and start the React frontend
cd ..
npm start &