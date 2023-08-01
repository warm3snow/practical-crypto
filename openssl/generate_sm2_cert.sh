#!/bin/sh

# 本脚本用于生成sm2自签名证书，仅用于测试目的
# 需要使用gmssl、openssl 1.1.x以上版本、或者tassl

# 1. 设置证书主体subject
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=warm3snow/OU=Server/CN=localhost"

# 2. 生成签名证书
gmssl ecparam -genkey -name SM2 -text -out ./gmcerts/server_sign.key
gmssl req -new -key ./gmcerts/server_sign.key -out ./gmcerts/server.csr -subj $SUBJ
gmssl x509 -req -days 3650 -sm3 -in ./gmcerts/server.csr  -signkey ./gmcerts/server_sign.key -out ./gmcerts/server_sign.crt \
-extfile ./openssl.cnf -extensions v3_req_sign


# 3。 生成加密证书
gmssl ecparam -genkey -name SM2 -text -out ./gmcerts/server_enc.key
gmssl req -new -key ./gmcerts/server_enc.key -out ./gmcerts/server.csr -subj $SUBJ
gmssl x509 -req -days 3650 -sm3 -in ./gmcerts/server.csr  -signkey ./gmcerts/server_enc.key -out ./gmcerts/server_enc.crt \
-extfile ./openssl.cnf -extensions v3_req_enc

# clean
rm -rf ./gmcerts/server.csr