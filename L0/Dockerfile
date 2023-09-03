FROM golang:1.21

COPY . /go/src/app

WORKDIR /go/src/app/cmd/L0

RUN go mod download
RUN go build main.go

EXPOSE 9090

CMD ["./main"]