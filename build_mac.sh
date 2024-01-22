#!/bin/bash

mkdir -p build
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/yuque_export_mac_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/yuque_export_mac_arm64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/yuque_export_linux_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/yuque_export_linux_arm64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/yuque_export_win_amd64.exe
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/yuque_export_win_arm64.exe
chmod +x build/yuque_export*
