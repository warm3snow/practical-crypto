#!/bin/sh

#set -x

# 在itrustee_sdk目录下执行
WKDIR_PATH="`pwd`"
ITRUSTEE_SDK_PATH="$WKDIR_PATH/itrustee_sdk"
ROOT_PATH="$WKDIR_PATH/root"

echo "手动修改CA，指定TA应用sec路径"
exit 1
# compile cert manager CA tool
# 注：这里需要手动修改CA下的itrustee_sdk/test/CA/cert_manager/cert_manager.c (TODO 注：目前需要手动修改，后续会在renew_root_pub_tool中支持自动化修改）
# 修改点1, line26:  #define TA_PATH "/usr/bin/4acaf7c8-c652-4643-9b7a-cc07e7a3187a.sec"
# 修改点2, line256: static TEEC_UUID g_taId = {
      # 0x4acaf7c8, 0xc652, 0x4643,
      # { 0x9b, 0x7a, 0xcc, 0x07, 0xe7, 0xa3, 0x18, 0x7a }

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
mkdir -p $ITRUSTEE_SDK_PATH/build/signtools/signed_config
mkdir -p $ITRUSTEE_SDK_PATH/build/signtools/TA_cert
cp $ROOT_PATH/config $ITRUSTEE_SDK_PATH/build/signtools/signed_config
cp $ROOT_PATH/private_key.pem $ITRUSTEE_SDK_PATH/build/signtools/TA_cert/private_key.pem

cp -f $ROOT_PATH/manifest.txt $ITRUSTEE_SDK_PATH/test/TA/cert_manager/manifest.txt
cd $ITRUSTEE_SDK_PATH/test/TA/cert_manager/
make
cp *.sec /usr/bin
