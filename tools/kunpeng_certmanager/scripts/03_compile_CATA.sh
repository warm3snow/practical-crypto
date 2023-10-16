#!/bin/sh

#set -x

# 在itrustee_sdk目录下执行
ITRUSTEE_SDK_PATH="`pwd`"

# compile cert manager CA tool
cd $ITRUSTEE_SDK_PATH/test/CA/cert_manager
make
cp certmanager /usr/bin

# compile cert manager TA tool (注：需先获取华为颁发的TA开发者证书，参考https://support.huawei.com/enterprise/zh/doc/EDOC1100315175/97c8190e)
# configs.xml示例：（特别注意：sys_verify_ta必须为true，CERT_Permission必须为true）
#<?xml version="1.0" encoding="utf-8"?>
#<ConfigInfo>
#  <TA_Basic_Info>
#      <service_name>rsa-demo</service_name>
#    <uuid>f68fd704-6eb1-4d14-b218-722850eb3ef0</uuid>
#  </TA_Basic_Info>
#  <TA_Manifest_Info>
#    <instance_keep_alive>false</instance_keep_alive>
#    <stack_size>40960</stack_size>
#    <heap_size>67928064</heap_size>
#    <multi_command>false</multi_command>
#    <multi_session>false</multi_session>
#    <single_instance>true</single_instance>
#    <sys_verify_ta>true</sys_verify_ta>
#  </TA_Manifest_Info>
#  <TA_Control_Info>
#	<CERT_Info>
#		<CERT_Permission>true</CERT_Permission>
#	</CERT_Info>
#  </TA_Control_Info>
#</ConfigInfo>
#获取cert manager的TA开发者证书、签发的config二进制后，将cert manager 的TA私钥、config放到ITRUSTEE_SDK_PATH项目根目录
mkdir -p $ITRUSTEE_SDK_PATH/test/TA/certmanager/cloud/signed_config
mkdir -p $ITRUSTEE_SDK_PATH/test/TA/certmanager/cloud/TA_cert
cp $ITRUSTEE_SDK_PATH/config $ITRUSTEE_SDK_PATH/test/TA/certmanager/cloud/signed_config
cp $ITRUSTEE_SDK_PATH/root.key $ITRUSTEE_SDK_PATH/test/TA/certmanager/cloud/TA_cert/private_key.pem
cd $ITRUSTEE_SDK_PATH/test/TA/certmanager/cloud
make
cp *.sec /usr/bin
