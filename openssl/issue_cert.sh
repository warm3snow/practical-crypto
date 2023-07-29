#!/bin/sh

caCrt=./rsacerts/ca.crt
caKey=./rsacerts/ca.key
crt=./rsacerts/server.crt
key=./rsacerts/server.key
csr=./rsacerts/server.csr

# self-signed ca certificate
openssl req -x509 -newkey rsa:4096 -keyout ${caKey} -out ${caCrt} -days 365 -nodes \
-subj "/C=CN/ST=Beijing/L=Beijing/O=CA/OU=BJ-CA/CN=www.ca.com"

# certificate signing request
openssl req -newkey rsa:4096 -keyout ${key} -out ${csr} -nodes \
-subj "/C=CN/ST=Beijing/L=Beijing/O=TX/OU=YL/CN=localhost"

# sign certificate
openssl x509 -req  -CA ${caCrt} -CAkey ${caKey} -CAcreateserial -days 365 \
-in ${csr} -out ${crt}