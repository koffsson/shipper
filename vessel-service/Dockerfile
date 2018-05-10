FROM golang:1.9.2 as builder

WORKDIR /go/src/github.com/koffsson/shipper/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/koffsson/shipper/vessel-service/vessel-service .

CMD ["./vessel-service"]
