#!/bin/bash
cd backend
go mod tidy
go run main.go
cd ..