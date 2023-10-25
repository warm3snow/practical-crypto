#!/bin/sh

#set -x
# 在itrustee_sdk目录下执行
ITRUSTEE_SDK_PATH="`pwd`"

# 1. 生成config证书
#注：生成config证书CSR请求文件时，证书主题“CN”字段内容必须为“iTrustee_Config”。
openssl genrsa -out taconfig.key 4096
openssl req -new -key taconfig.key -out cert_req_01.csr -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=iTrustee_Config"
openssl x509 -req -in cert_req_01.csr -CA ./second.crt -CAkey./second.key -CAcreateserial -out taconfig.crt -sha256 -days 3650
openssl x509 -in taconfig.crt -outform der -out taconfig.der

# 2. 生成TA证书
openssl genrsa -out private_key.pem 4096
#生成业务TA证书CSR请求文件,其中“CN”字段格式为“{业务TA uuid}_{业务TA server name}”，请根据实际情况替换!!!
openssl req -new -key private_key.pem -out cert_req_02.csr -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=87709191-1633-4d4e-b39a-951a59d82fb2_third-demo"
openssl x509 -req -in cert_req_02.csr -CA ./second.crt -CAkey./second/second.key -CAcreateserial -out tacert.crt -sha256 -days 3650
openssl x509 -in tacert.crt -outform der -out ta_cert.der

# 3. 准备TA配置文件configs.xml
# configs.xml示例：
#<?xml version="1.0" encoding="utf-8"?>
#<ConfigInfo>
#  <TA_Basic_Info>
#      <service_name>third-demo</service_name>
#      <uuid>87709191-1633-4d4e-b39a-951a59d82fb2</uuid>
#  </TA_Basic_Info>
#  <TA_Manifest_Info>
#    <instance_keep_alive>false</instance_keep_alive>
#    <stack_size>40960</stack_size>
#    <heap_size>67928064</heap_size>
#    <multi_command>false</multi_command>
#    <multi_session>false</multi_session>
#    <single_instance>true</single_instance>
#  </TA_Manifest_Info>
#</ConfigInfo>

# 4. 基于itrustee sdk，编译TA config
mkdir -p $ITRUSTEE_SDK_PATH/build/pack-Config/xml2tlv_tools/csv
cp itrustee_sdk/build/signtools/tag_parse_dict.csv itrustee_sdk/build/pack-Config/xml2tlv_tools/csv
cp -f configs.xml itrustee_sdk/build/pack-Config/input
mv taconfig.key itrustee_sdk/build/pack-Config/config_cert/config_cert_private.key
mv taconfig.der itrustee_sdk/build/pack-Config/config_cert/
mv ta_cert.der itrustee_sdk/build/pack-Config/ta_cert/
cd $ITRUSTEE_SDK_PATH/build/pack-Config
bash local_sign.sh
ll $ITRUSTEE_SDK_PATH/build/pack-Config/output/

# 5. 基于生成的config二进制、业务TA公私钥对，编译业务TA。
# 参考脚本 '03_compile_CATA.sh'


