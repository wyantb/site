# 'base' image
FROM golang:1.5
RUN apt-get -y update && apt-get install -y nodejs npm node-less
RUN ln -s ./nodejs /usr/bin/node

# 'server' image
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install
RUN make deps lessmin bindata

CMD make devserver
