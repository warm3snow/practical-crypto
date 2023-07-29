#!/bin/sh


caCrt=./rsacerts/ca.crt
caKey=./rsacerts/ca.key
crt=./rsacerts/server.crt
key=./rsacerts/server.key
csr=./rsacerts/server.csr
newCrt=./rsacerts/new_server.crt

# renew csr
openssl x509 -x509toreq -sha256 -in ${crt} -signkey ${key} \
-out ${csr}

# renew cert, days 3650 represents 10 years. old cert only 1 year.
openssl x509 -req -sha256 -CAkey ${caKey} -CA ${caCrt} -CAcreateserial \
-in ${csr} -days 3650 \
-out ${newCrt}

openssl verify -CAfile ${caCrt} ${newCrt}
