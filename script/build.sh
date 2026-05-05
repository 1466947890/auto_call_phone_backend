#!/usr/bin/env bash
# 创建app目录

go build -o /go/src/app/main main.go

cd /go/src/app/ && ./main
