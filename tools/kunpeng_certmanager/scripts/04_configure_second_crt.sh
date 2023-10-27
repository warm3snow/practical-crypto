#!/bin/sh

#set -x

# 在itrustee_sdk目录下执行
ITRUSTEE_SDK_PATH="`pwd`"
ROOT_PATH="`pwd`/root"
SECOND_PATH="`pwd`/second"

# 仅执行一次，执行后，需要安全保存second.key、second.crt、second.der三个文件
openssl genrsa -out $SECOND_PATH/second.key 4096
openssl req -new -key $SECOND_PATH/second.key -out $SECOND_PATH/second.csr -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=Secondary CA"

openssl x509 -req -CA $ROOT_PATH/root.crt -CAkey $ROOT_PATH/private_key.pem -CAcreateserial \
-in $SECOND_PATH/second.csr -out  $SECOND_PATH/second.crt \
-sha256 -extfile $SECOND_PATH/openssl.cnf -extensions v3_ca -days 3650
openssl x509 -in $SECOND_PATH/second.crt -outform $SECOND_PATH/der -out $SECOND_PATH/second.der

# 导入TA二级证书
# 注：成功导入后，用户可基于二级证书自行颁发TA证书以及签发config，然后参与TA编译形成业务TA，
# 此时安全OS仅允许加载三方TA（由用户二级证书签发的TA证书）以及系统TA（configs.xml注有<sys_verify_ta>标签）加载运行，
# 原由华为签发的TA证书的TA将会加载失败。
/usr/bin/certmanager import $SECOND_PATH/second.der