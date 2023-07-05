#!/bin/sh

# curl访问https
# -k 忽略证书验证
# --tlsv1.2 指定TLS版本(这里是指定最低版本，实际使用的版本是服务端和客户端协商后决定的)
# -s 静默模式
curl --tlsv1.2  -s -k https://localhost:8443/hello | xxd

# 注 `xxd`是一个十六进制查看工具, 可以用来查看二进制文件