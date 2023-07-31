#!/bin/sh

# 本脚本用于生成sm2自签名证书，仅用于测试目的

# 生成自签名证书
# 方式1: 生成sm2 key => 生成证书签名请求 => 生成证书
# 方式2: 生成sm2 key, 同时生成自签名证书

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=warm3snow/OU=Server/CN=localhost"

# 签名证书 (openssl 1.1.x以上版本支持sm2, 或者tassl、gmssl)
gmssl ecparam -genkey -name SM2 -text -out ./gmcerts/server_sign.key
gmssl req -new -key ./gmcerts/server_sign.key -out ./gmcerts/server.csr -subj $SUBJ
gmssl x509 -req -days 3650 -sm3 -in ./gmcerts/server.csr  -signkey ./gmcerts/server_sign.key -out ./gmcerts/server_sign.crt \
-extfile ./openssl.cnf -extensions v3_req


# 加密证书
gmssl ecparam -genkey -name SM2 -text -out ./gmcerts/server_enc.key
gmssl req -new -key ./gmcerts/server_enc.key -out ./gmcerts/server.csr -subj $SUBJ
gmssl x509 -req -days 3650 -sm3 -in ./gmcerts/server.csr  -signkey ./gmcerts/server_enc.key -out ./gmcerts/server_enc.crt \
-extfile ./openssl.cnf -extensions v3_req

# clean
rm -rf ./gmcerts/server.csr