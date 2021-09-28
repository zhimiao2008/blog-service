NAME := $(shell basename $(shell git config --get remote.origin.url) | sed 's/\.git//')
BRANCH := $(shell git symbolic-ref --short HEAD 2>/dev/null)

RELEASESDIR=releases

default: all

all: pack ## all

test: ## 检查&测试

build: ## 构建
	go build -o blog-service -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`" *.go

build-for-docker: ## 初始化
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o blog-service -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`" *.go

pack: clean test build ## 打包项目
	mkdir -p $(RELEASESDIR)/configs
	cp -r configs  $(RELEASESDIR)/
	mv blog-service  $(RELEASESDIR)/
	tar czf blog-service.tar.gz $(RELEASESDIR)/
	mv blog-service.tar.gz $(RELEASESDIR)/
docker: clean test build-for-docker ## 打包镜像
	docker build . -t blog-service:latest

clean: ## 清除目录
	rm -rf $(RELEASESDIR) blog-service blog-service.tar.gz
	go clean

.DEFAULT: all


help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'