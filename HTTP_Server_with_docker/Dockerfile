FROM golang
MAINTAINER Jacksontmm

RUN mkdir /server
WORKDIR /server

COPY . .

RUN go mod download
RUN go build main.go

CMD ["./main","run"]