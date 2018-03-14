protoc -I. --go_out=plugins=micro:. ./proto/vessel/vessel.proto
docker build -t vessel-service .