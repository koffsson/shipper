FROM golang:1.9.2 as builder

WORKDIR /go/src/github.com/koffsson/shipper/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o user-service -a main.go handler.go repository.go database.go token_service.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/koffsson/shipper/user-service/user-service .

CMD ["./user-service"]
