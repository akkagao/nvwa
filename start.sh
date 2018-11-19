#!/bin/bash
echo "清理编译环境"
rm -rf etc nvwa
echo "静态文件打包"
go-bindata  -pkg etc -o etc/bindata.go template/*
echo "启动"
go run main.go service --name goods
