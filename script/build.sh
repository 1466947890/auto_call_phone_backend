#!/usr/bin/env bash
go build -o /go/src/app/main main.go

cd /go/src/app/ && ./main
