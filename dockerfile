FROM golang:1.10.2
ENV GOPATH /go
COPY . /go/src/github.com/dijckstra/cartola-data-scrapper

RUN cd /go/src/github.com/dijckstra/cartola-data-scrapper && go get ./... && go build -o main.go

CMD ["/go/bin/cartola-data-scrapper"]