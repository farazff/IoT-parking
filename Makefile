#@IgnoreInspection BashAddShebang

export APP=cluboffice

export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

export BUILD_INFO_PKG="github.com/okian/servo/v2/config"

export LDFLAGS="-w -s -X $(BUILD_INFO_PKG).date=$$(TZ=Asia/Tehran date '+%FT%T') -X $(BUILD_INFO_PKG).commit=$$(git rev-parse HEAD | cut -c 1-8) -X $(BUILD_INFO_PKG).branch=$$([ -z \"$$CI_COMMIT_BRANCH\" ] && echo $$CI_COMMIT_BRANCH || git rev-parse --abbrev-ref HEAD) -X $(BUILD_INFO_PKG).tag=$$([ -z \"$$CI_COMMIT_TAG\" ] && echo $$CI_COMMIT_TAG || git describe --exact-match --tags $$(git log -n1 --pretty='%h') 2> /dev/null|| echo NONE) -X $(BUILD_INFO_PKG).app=$(APP)" 

export GOFLAGS=-mod=vendor

swagger: 
	which swagger  || go get -u github.com/go-swagger/go-swagger/cmd/swagger
	GO111MODULE=on swagger generate spec -o ./io/rest/swagger.yaml --scan-models

