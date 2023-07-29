#!/bin/sh

# 生成自签名证书
# 方式1: 生成rsa key => 生成证书签名请求 => 生成证书
# 方式2: 生成rsa key, 同时生成自签名证书

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=Test/OU=Test/CN=localhost"

# 方式1
# generate rsa key => key
openssl genrsa -out ./rsacerts/server.key 2048

# generate certificate signing request => csr
openssl req -new -key ./rsacerts/server.key -out ./rsacerts/server.csr -subj $SUBJ

# issue certificate => crt
openssl x509 -req -sha256 -days 365 -in ./rsacerts/server.csr -signkey ./rsacerts/server.key -out ./rsacerts/server.crt


# 方式2
#openssl req -x509 -newkey rsa:2048 -keyout ./rsacerts/server.key -out ./rsacerts/server.crt -days 3650  -nodes \
#        -subj $SUBJ