.PHONY: polygon marketData
all: polygon

polygon:
	protoc ./polygon/rest.proto --go_out=./polygon
	protoc ./polygon/stream.proto --go_out=./polygon


marketData:
	protoc ./marketData/marketData.proto --go_out=./marketData



