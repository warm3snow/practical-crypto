#!/bin/sh

#set -x

# 在itrustee_sdk目录下执行
ITRUSTEE_SDK_PATH="`pwd`"

cd $ITRUSTEE_SDK_PATH
openssl genrsa -out second.key 4096
openssl req -new -key second.key -out second.csr -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=Secondary CA"

openssl x509 -req -CA root.crt -CAkey root.key -CAcreateserial -in second.csr -out second.crt -sha256 -extfile openssl.cnf -extensions v3_ca -days 3650
openssl x509 -in second.crt -outform der -out second.der

# 导入TA二级证书
# 注：成功导入后，用户可基于二级证书自行颁发TA证书以及签发config，然后参与TA编译形成业务TA，
# 此时安全OS仅允许加载三方TA（由用户二级证书签发的TA证书）以及系统TA（configs.xml注有<sys_verify_ta>标签）加载运行，
# 原由华为签发的TA证书的TA将会加载失败。
/usr/bin/certmanager import second.der

rm -rf second.csr root.srl