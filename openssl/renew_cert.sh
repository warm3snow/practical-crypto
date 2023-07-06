#!/bin/sh


caCrt=./testcerts/ca_cert.pem
caKey=./testcerts/ca_key.pem
crt=./testcerts/cert.pem
key=./testcerts/key.pem
csr=./testcerts/csr.pem
newCrt=./testcerts/new_cert.pem

# renew csr
openssl x509 -x509toreq -sha256 -in ${crt} -signkey ${key} \
-out ${csr}

# renew cert, days 3650 represents 10 years. old cert only 1 year.
openssl x509 -req -sha256 -CAkey ${caKey} -CA ${caCrt} -CAcreateserial \
-in ${csr} -days 3650 \
-out ${newCrt}

openssl verify -CAfile ${caCrt} ${newCrt}
