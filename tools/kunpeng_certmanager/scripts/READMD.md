# 鲲鹏TEE certmanager
## 根证书和CertManger TA开发者证书
先向华为申请certmanager TA开发者证书，申请前需准备好以下信息
- configs.xml
```xml
<?xml version="1.0" encoding="utf-8"?>
<ConfigInfo>
    <TA_Basic_Info>
        <service_name>certmanager</service_name>
        <uuid>d354a48a-e6b1-4651-b7b5-c79c28f13870</uuid>
    </TA_Basic_Info>
    <TA_Manifest_Info>
        <instance_keep_alive>false</instance_keep_alive>
        <stack_size>32768</stack_size>
        <heap_size>2097152</heap_size>
        <multi_command>false</multi_command>
        <multi_session>true</multi_session>
        <single_instance>true</single_instance>
        <sys_verify_ta>true</sys_verify_ta>
    </TA_Manifest_Info>
    <TA_Control_Info>
        <CERT_Info>
            <CERT_Permission>true</CERT_Permission>
        </CERT_Info>
    </TA_Control_Info>
</ConfigInfo>
```
- manifest.txt 
本地保存，需要与configs.xml参数对应
```sh
gpd.ta.appID:   		d354a48a-e6b1-4651-b7b5-c79c28f13870
gpd.ta.service_name:		certmanager
gpd.ta.singleInstance:		true
gpd.ta.multiSession: 		true
gpd.ta.instanceKeepAlive:	false
gpd.ta.dataSize:		2097152
gpd.ta.stackSize:		32768
gpd.ta.sys_verify_ta:   true
```
- csr文件
```sh
# 生成certmanager TA私钥
$ openssl genrsa -out private_key.pem 4096

# 生成certmanager TA证书请求
$ openssl req -new -key private_key.pem -out cert_req_01.csr -subj "/C=CN/O=Huawei/OU=Huawei iTrustee Production/CN=d354a48a-e6b1-4651-b7b5-c79c28f13870_certmanager"

# 自签证书
openssl req -new -x509 -key private_key.pem -out root.crt -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=Root CA" -days 3650
```
注：private_key.pem，root.crt, manifest.txt 需要妥善保管

## 二级证书生成

## Scripts功能介绍
- 首先需要将private_key.pem/config/manifest.txt配置到scripts/certmanager目录下（该目录如果没有需要手工创建）
- 运行脚本[01_download_itrustee_depends.sh](01_download_itrustee_depends.sh)下载itrustee依赖库
- 运行脚本[02_configure_root_crt.sh](02_configure_root_crt.sh) 生成root证书，并提取root公钥，配置到certmanager TA程序中（硬编码）
- 运行脚本[03_compile_certmanager.sh](03_compile_certmanager.sh)生成certmanager CA/TA程序
- 运行脚本[04_configure_second_crt.sh](04_configure_second_crt.sh)生成certmanager 二级证书并导入（需要uncomment脚本中的部分代码）
- 运行脚本[05_issue_ta_crt.sh](05_issue_ta_crt.sh)生成TA开发者证书