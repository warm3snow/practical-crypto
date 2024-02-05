FROM golang:1.18

#RUN GO_VERSION=1.19.8                                                 &&\
#  cd /root                                                            &&\
#  wget "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"  &&\
#  tar -xzf go${GO_VERSION}.linux-amd64.tar.gz                         &&\
#  rm -f go${GO_VERSION}.linux-amd64.tar.gz                            &&\
#  mv go /usr/local

#ENV GOPROXY="https://goproxy.cn"  \
#  GOPATH="/root/go"               \
#  GOROOT="/usr/local/go"          \
#  CGO_ENABLED="1"                 \
#  PATH="/usr/local/go/bin:$PATH"

COPY ./crypto/hsmimpl/lib/libswsds.so /usr/local/lib64/
COPY ./crypto/hsmimpl/lib/* /etc/


RUN mkdir -p /root/go/src/github.com/warm3snow/practical-crypto
WORKDIR /root/go/src/github.com/warm3snow/practical-crypto
COPY . .
RUN go mod tidy

WORKDIR /root/go/src/github.com/warm3snow/practical-crypto/crypto/hsmimpl

#go test -v -test.run TestVerifySign
#go test -v -test.run TestSM3HMac