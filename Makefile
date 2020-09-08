.PHONY: polygon
all: polygon



polygon:
	protoc ./polygon/greetings.proto --go_out=./polygon
	protoc ./polygon/forex.proto --go_out=./polygon
	protoc ./polygon/rest.proto --go_out=./polygon



