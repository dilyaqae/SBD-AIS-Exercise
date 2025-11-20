#!/bin/sh
cd /app
go mod download
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem
chmod +x /app/ordersystem
