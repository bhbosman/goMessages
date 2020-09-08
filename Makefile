.PHONY: polygon
all: polygon

polygon:
	protoc ./polygon/rest.proto --go_out=./polygon
	protoc ./polygon/stream.proto --go_out=./polygon



