#!/bin/sh

#set -x

# 在itrustee_sdk目录下执行
ITRUSTEE_SDK_PATH="`pwd`/itrustee_sdk"

# 1. 生成root证书
#openssl genrsa -out root.key 4096
openssl req -new -x509 -key root.key -out root.crt -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=Root CA" -days 3650

#openssl x509 -in root.crt -pubkey -noout -C

# backup cert_config.h
CertManagerPath="$ITRUSTEE_SDK_PATH/test/TA/cert_manager"
cp $CertManagerPath/include/cert_config.h $CertManagerPath/include/cert_config.h_.`date +%Y%m%d%H%M%S`


# compile renew_root_pub_tool
go build -o renew_root_pub_tool ../renew_root_pub_tool/renew_root_pub_tool.go

# renew root public key
./renew_root_pub_tool \
-certConfigFilePath $CertManagerPath/include/cert_config.h \
-rootCrt root.crt