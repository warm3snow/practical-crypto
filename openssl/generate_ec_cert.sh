#!/bin/sh

# 本脚本用于生成sm2自签名证书，仅用于测试目的

# 生成自签名证书
# 方式1: 生成sm2 key => 生成证书签名请求 => 生成证书
# 方式2: 生成sm2 key, 同时生成自签名证书

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=warm3snow/OU=Server/CN=localhost"

# generate sm2 key => key
gmssl ecparam -genkey -name prime256v1 -text -out ./eccerts/server.key

# generate certificate signing request => csr
gmssl req -new -key ./eccerts/server.key -out ./eccerts/server.csr -subj $SUBJ

# issue certificate => crt
gmssl x509 -req -days 3650 -sha256 -in ./eccerts/server.csr -signkey ./eccerts/server.key -out ./eccerts/server.crt


