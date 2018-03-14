protoc -I. --go_out=plugins=micro:. ./proto/consignment/consignment.proto
docker build -t consignment-service .