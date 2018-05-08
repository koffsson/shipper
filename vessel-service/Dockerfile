FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/ewanvalentine.io/shipper/vessel-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-ache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from-builder /go/src/github.com/ewanvalentine.io/shipper/vessel-service/vessel-service .

CMD ["./vessel-service"]
