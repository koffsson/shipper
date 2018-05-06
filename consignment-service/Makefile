build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/ewanvalentine.io/shipper/consignment-service \
proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service
