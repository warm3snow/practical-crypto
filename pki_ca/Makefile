VERSION=v1.0.0
BRANCH=master
OS := $(shell uname -s)
PLATFORM=$(shell uname -m)
ARCH := $(shell uname -m)
DATETIME=$(shell date "+%Y/%m/%d %H:%M:%S")
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
# 注意：这里的路径是你自己的交叉编译工具链的路径, 下载地址https://github.com/messense/homebrew-macos-cross-toolchains/releases
CROSS_BUILD_LINUX_GUN_GCC = /Users/hxy/wkdir/x86_64-unknown-linux-gnu/build/x86_64-unknown-linux-gnu-gcc

LOCALCONF_HOME=github.com/warm3snow/crypto-service-backend/config
GOLDFLAGS += -X "${LOCALCONF_HOME}.CurrentVersion=${VERSION}"
GOLDFLAGS += -X "${LOCALCONF_HOME}.BuildDateTime=${DATETIME}"
GOLDFLAGS += -X "${LOCALCONF_HOME}.GitBranch=${GIT_BRANCH}"
GOLDFLAGS += -X "${LOCALCONF_HOME}.GitCommit=${GIT_COMMIT}"

build_local:
	#go mod tidy && CGO_LDFLAGS="-g -O2 -L$(LD_LIBRARY_PATH)" go build -o ./build/crypto-service-backend cmd/main.go
	go mod tidy && go build -ldflags '${GOLDFLAGS}' -o ./build/crypto-service-backend ./cmd/main.go

# 本地需要安装交叉编译器
build_linux:
	# mkdir build
	@rm -rf build/linux_amd64 && mkdir -p build/linux_amd64
	# linux
	go mod tidy && CGO_ENABLED=1 CC=${CROSS_BUILD_LINUX_GUN_GCC}  GOARCH=amd64 GOOS=linux go build  -ldflags '${GOLDFLAGS}' -o ./build/crypto-service-backend ./cmd/main.go	

	md5 ./build/crypto-service-backend | awk '{print $NF}' > build/md5.txt
generate:
	go generate ./...

ut:
	go test ./...
