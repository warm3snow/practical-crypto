#!/bin/sh

# 生成自签名证书:
# -x509 生成自签名证书
# -newkey 生成新的RSA密钥
# -keyout 指定私钥文件
# -out 指定证书文件
# -days 证书有效期
# -nodes 不使用密码
# -subj 证书主题
CA_SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=Test/OU=Test/CN=localhost/emailAddress=localhost@test.com"
openssl req -x509 -newkey rsa:2048 -keyout ../testdata/server.key -out ../testdata/server.crt -days 3650  -nodes \
        -subj $CA_SUBJ