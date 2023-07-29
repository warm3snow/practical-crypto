#!/bin/sh

# 本脚本用于生成sm2自签名证书，仅用于测试目的

# 生成自签名证书
# 方式1: 生成sm2 key => 生成证书签名请求 => 生成证书
# 方式2: 生成sm2 key, 同时生成自签名证书

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=Test/OU=Test/CN=localhost"

# 方式1
# generate sm2 key => key
openssl gensm2 -out ./gmcerts/server.key 2048

# generate certificate signing request => csr
openssl req -new -key ./gmcerts/server.key -out ./gmcerts/server.csr -subj $SUBJ

# issue certificate => crt
openssl x509 -req -sha256 -days 365 -in ./gmcerts/server.csr -signkey ./gmcerts/server.key -out ./gmcerts/server.crt


# 方式2
#openssl req -x509 -newkey sm2:2048 -keyout ./gmcerts/server.key -out ./gmcerts/server.crt -days 3650  -nodes \
#        -subj $SUBJ