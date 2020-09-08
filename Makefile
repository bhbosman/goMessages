.PHONY: polygon marketData luno
all: polygon

polygon:
	protoc ./polygon/rest.proto --go_out=./polygon
	protoc ./polygon/stream.proto --go_out=./polygon

luno:
	protoc ./luno/stream.proto --go_out=./luno

marketData:
	protoc ./marketData/marketData.proto --go_out=./marketData



