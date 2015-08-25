
.PHONY: all clean deps lint bindata

all: deps bindata lint

bindata:
	@go-bindata -prefix=raw-data/ raw-data/

lint:
	@-go vet ./...

deps:
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/jteeuwen/go-bindata/...
	@go get -u -v github.com/labstack/echo
