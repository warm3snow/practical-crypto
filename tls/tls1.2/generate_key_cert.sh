#!/bin/sh

CA_SUBJ="/C=CN/ST=BeiJing/L=BeiJing/O=Test/OU=Test/CN=localhost/emailAddress=localhost@test.com"
openssl req -x509 -newkey rsa:2048 -keyout server.key -out server.crt -days 3650  -nodes \
        -subj $CA_SUBJ