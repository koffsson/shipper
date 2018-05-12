FROM golang:1.9.2 as builder

WORKDIR /go/src/github.com/koffsson/shipper/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 build -a .


FROM alpine:latest

RUN apt --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/koffsson/shipper/user-service/user-service .

CMD ["./user-service"]
