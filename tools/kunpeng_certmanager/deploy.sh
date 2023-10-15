#!/bin/sh

# kunpeng certmanager deploy script

set -x

itrustee_sdk_path="`pwd`/itrustee_sdk"

# clone itrustee_sdk
git clone https://gitee.com/openeuler/itrustee_sdk.git -b master
cd itrustee_sdk
git checkout 22c68de6cac1810c927a91bf86a40cc07b5005a2

# clone libboundscheck & build
cd itrustee_sdk/thirdparty/open_source
git clone https://gitee.com/openeuler/libboundscheck.git
cd libboundscheck
make
cp lib/libboundscheck.so /usr/lib64

# set certmanager root public key
cd $itrustee_sdk_path
openssl genrsa -out root.key 4096
openssl req -new -x509 -key root.key -out root.crt -subj "/C=CN/L=F/O=testRootCA/OU=ACS/CN=Root CA" -days 3650

#openssl x509 -in root.crt -pubkey -noout -C

# backup cert_config.h
CertManagerPath="$itrustee_sdk_path/test/TA/cert_manager"
cp $CertManagerPath/include/cert_config.h $PWDPAHT $itrustee_sdk_path/cert_config.h_.`date +%Y%m%d%H%M%S`

# renew root public key
./renew_root_pub_tool \
-certConfigFilePath $CertManagerPath/include/cert_config.h \
-rootCrt $itrustee_sdk_path/root.crt