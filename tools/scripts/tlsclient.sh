#!/bin/sh

# 生成tlsclient
if [ ! -d "../bin" ]; then
    mkdir ../bin
fi

if [ ! -f "../bin/tlsClient" ]; then
  go build -o ../bin/tlsClient ../tls/tlsclient.go
fi


#➜  tls git:(master) ✗ ./tlsClient -h
#Usage of ./dial:
#  -addr string
#        addr (default "https://localhost:8443")
#  -skipVerify
#        skipVerify (default true)
#  -tlsVersion string 协商时指定使用的tls版本（注与curl指定的--tlsv1.2的区别）
#        tls version (default "1.2")
../bin/tlsClient -addr https://localhost:8443 -tlsVersion 1.2 -skipVerify true