
.PHONY: all clean deps lint bindata less lessmin devserver dev assets

all: less assets bindata lint

dev: all devserver

devserver:
	GOBIN=./ go install server.go
	./server 1323
	# If I want to force all assets to be bundled in binary?
	#(cd ../ && site/server)

less: lessmin

lessmin:
	#lessc css/global.less assets/css/global.css
	lessc css/global.less assets/css/global.css --clean-css="--s1 --advanced --compatibility=ie8"

bindata:
	@go-bindata -pkg=data -prefix=raw-data/ -o=data/raw.go raw-data/
	#@go-bindata -pkg=assets -prefix=assets/ -o=assets/raw.go assets/

assets:
	mkdir -p assets/css/fonts
	cp css/fonts/* assets/css/fonts

lint:
	@-go vet ./...

deps:
	npm install -g less
	npm install -g less-plugin-clean-css
	@go get -u -v github.com/jteeuwen/go-bindata/...
#	@go get -u -v github.com/golang/lint/golint
#	@go get -u -v github.com/labstack/echo
