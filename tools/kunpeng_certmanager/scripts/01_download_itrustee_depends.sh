#!/bin/sh

#set -x

ITRUSTEE_SDK_PATH="`pwd`/itrustee_sdk"

# clone itrustee_sdk
git clone https://gitee.com/openeuler/itrustee_sdk.git -b master
cd itrustee_sdk
git checkout 22c68de6cac1810c927a91bf86a40cc07b5005a2

# clone libboundscheck & build
cd $ITRUSTEE_SDK_PATH/thirdparty/open_source
git clone https://gitee.com/openeuler/libboundscheck.git
cd libboundscheck
make
cp lib/libboundscheck.so /usr/lib64