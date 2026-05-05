#!/usr/bin/env bash
go env -w GOPROXY=https://goproxy.cn,direct && go mod tidy

go build -o /go/src/app/main main.go

cd /go/src/app/ && ./main
