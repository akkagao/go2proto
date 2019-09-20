.PHONY: proto
proto:

	go2proto -t proto -f server/common/common.go
	go2proto -t proto -f server/demo.go

	protoc --go_out=plugins=grpc:. proto/common.proto
	protoc --go_out=plugins=grpc:. proto/demo.proto
# 	protoc --micro_out=. --go_out=. proto/common.proto
# 	protoc --micro_out=. --go_out=. proto/demo.proto
