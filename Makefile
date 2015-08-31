
.PHONY: all clean deps lint bindata less lessmin devserver dev

all: less bindata lint

dev: all devserver

devserver:
	go run server.go

less:
	lessc css/global.less css/global.css

lessmin:
	lessc css/global.less css/global.css --clean-css="--s1 --advanced --compatibility=ie8"

bindata:
	@go-bindata -pkg=data -prefix=raw-data/ -o=data/raw.go raw-data/

lint:
	@-go vet ./...

deps:
	npm install -g less
	npm install -g less-plugin-clean-css
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/jteeuwen/go-bindata/...
	@go get -u -v github.com/labstack/echo
