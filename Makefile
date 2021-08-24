NAME := $(shell basename $(shell git config --get remote.origin.url) | sed 's/\.git//')
BRANCH := $(shell git symbolic-ref --short HEAD 2>/dev/null)

TMPDIR=blog.com
RELEASESDIR=releases

default: all

all: pack ## all

test: ## 检查&测试

build:  ## 初始化
	mkdir -p $(TMPDIR)/bin
	go build -o $(TMPDIR)/bin/blog -ldflags "-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD` "
	cp -r configs $(TMPDIR)/
pack: clean test build ## 打包项目
	mkdir -p $(RELEASESDIR)
	tar czf $(RELEASESDIR)/$(TMPDIR).tar.gz $(TMPDIR)/

clean: ## 清除目录
	rm -rf $(TMPDIR)
	rm -rf $(RELEASESDIR)
	go clean

.DEFAULT: all
