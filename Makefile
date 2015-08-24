
.PHONY: all clean deps lint

lint:
	@-go vet ./...
	@-golint ./... |grep -v bindata/template/* |grep -v bindata/static/*

deps:
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/jteeuwen/go-bindata/...
