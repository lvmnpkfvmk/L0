FROM golang:1.21

COPY . /go/src/app

WORKDIR /go/src/app/cmd/

RUN go build main.go

EXPOSE 9090

CMD ["./main"]