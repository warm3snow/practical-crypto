#!/bin/sh


#openssl genrsa -out ca.key 4096
#openssl req -new -x509 -days 3650 -config openssl.cnf  -key ca.key -out ca.crt
#openssl x509 -text -noout -in ca.crt | grep -A10 "X509v3 extensions"
#
#openssl genrsa -out server.key.pem 4096
#openssl req -config server.openssl.cnf -new -key server.key.pem -out server.csr
##openssl req -text -in server.csr | grep -A 6 "Requested Extensions:"
##openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
#openssl x509 -req -in server.csr  -CA ca.crt -CAkey ca.key -out server.crt  -CAcreateserial -CAserial serial -days 365 -sha512 -extfile custom_server.openssl.cnf
##openssl x509 -text -noout -in server.crt | grep -A10 "X509v3 extensions:"


SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=root-cert/CN=ca.wx-org5.chainmaker.org"
openssl ecparam -genkey -name prime256v1 -out ca.key -noout
openssl req -new -x509 -days 3650 -config openssl.cnf  -key ca.key -out ca.crt  -subj $SUBJ

# 生成节点证书
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=common/CN=common1.sign.wx-org5.chainmaker.org"
openssl ecparam -out common1.sign.key -name prime256v1 -genkey -noout
openssl req -key common1.sign.key -new -out common1.sign.csr -subj $SUBJ -config server.openssl.cnf
openssl x509 -req -in common1.sign.csr -CA ca.crt -CAkey ca.key -out common1.sign.crt -days 3650 -extfile custom_server.openssl.cnf

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=common/CN=common1.tls.wx-org5.chainmaker.org"
openssl ecparam -out common1.tls.key -name prime256v1 -genkey -noout
openssl req -key common1.tls.key -new -out common1.tls.csr -subj $SUBJ -config server.openssl.cnf
openssl x509 -req -in common1.tls.csr -CA ca.crt -CAkey ca.key -out common1.tls.crt -days 3650 -extfile custom_server.openssl.cnf

rm -rf common1.sign.csr common1.tls.csr


# 生成管理员证书
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=admin/CN=admin1.sign.wx-org5.chainmaker.org"
openssl ecparam -out admin1.sign.key -name prime256v1 -genkey -noout
openssl req -key admin1.sign.key -new -out admin1.sign.csr -subj $SUBJ -config server.openssl.cnf
openssl x509 -req -in admin1.sign.csr -CA ca.crt -CAkey ca.key -out admin1.sign.crt -days 3650 -extfile custom_server.openssl.cnf

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=admin/CN=admin1.tls.wx-org5.chainmaker.org"
openssl ecparam -out admin1.tls.key -name prime256v1 -genkey -noout
openssl req -key admin1.tls.key -new -out admin1.tls.csr -subj $SUBJ -config server.openssl.cnf
openssl x509 -req -in admin1.tls.csr -CA ca.crt -CAkey ca.key -out admin1.tls.crt -days 3650 -extfile custom_server.openssl.cnf

rm -rf admin1.sign.csr admin1.tls.csr

# 生成普通用户证书
SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=client/CN=client1.sign.wx-org5.chainmaker.org"
openssl ecparam -out client1.sign.key -name prime256v1 -genkey -noout
openssl req -key client1.sign.key -new -out client1.sign.csr -subj $SUBJ  -config server.openssl.cnf
openssl x509 -req -in client1.sign.csr -CA ca.crt -CAkey ca.key -out client1.sign.crt -days 3650 -extfile custom_server.openssl.cnf

SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=wx-org5.chainmaker.org/OU=client/CN=client1.tls.wx-org5.chainmaker.org"
openssl ecparam -out client1.tls.key -name prime256v1 -genkey -noout
openssl req -key client1.tls.key -new -out client1.tls.csr -subj $SUBJ -config server.openssl.cnf
openssl x509 -req -in client1.tls.csr -CA ca.crt -CAkey ca.key -out client1.tls.crt -days 3650 -extfile custom_server.openssl.cnf

rm -rf client1.sign.csr client1.tls.csr