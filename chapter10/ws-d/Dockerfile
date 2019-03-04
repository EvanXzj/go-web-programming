FROM golang

ADD . /go/src/github.com/EvanXzj/ws-d
WORKDIR /go/src/github.com/EvanXzj/ws-d

RUN go get github.com/lib/pq
RUN go install github.com/EvanXzj/go-web-programming/chapter10/ws-d

ENTRYPOINT [ "/go/bin/ws-d" ]

EXPOSE 8080
