FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/koffsson/shipper/consignment-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux CGO_ENABLED=0 go build -a .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/koffsson/shipper/consignment-service/consignment-service .

CMD ["./consignment-service"]
