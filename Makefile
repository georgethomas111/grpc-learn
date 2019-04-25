.PHONY: generate
generate:
	protoc --proto_path=learn  learn/simple.proto --go_out=plugins=grpc:learn
