FROM golang:1.9.2 as builder

WORKDIR /go/src/github.com/koffsson/shipper/consignment-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/koffsson/shipper/consignment-cli/consignment-cli .

CMD ["./consignment-cli"]
