#!/bin/sh

caCrt=./testcerts/ca_cert.pem
caKey=./testcerts/ca_key.pem
crt=./testcerts/cert.pem
key=./testcerts/key.pem
csr=./testcerts/csr.pem

# self-signed ca certificate
openssl req -x509 -newkey rsa:4096 -keyout ${caKey} -out ${caCrt} -days 365 -nodes \
-subj "/C=CN/ST=Beijing/L=Beijing/O=CA/OU=BJ-CA/CN=www.ca.com"

# certificate signing request
openssl req -newkey rsa:4096 -keyout ${key} -out ${csr} -nodes \
-subj "/C=CN/ST=Beijing/L=Beijing/O=TX/OU=YL/CN=localhost"

# sign certificate
openssl x509 -req  -CA ${caCrt} -CAkey ${caKey} -CAcreateserial -days 365 \
-in ${csr} -out ${crt}