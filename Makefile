.PHONY: polygon marketData luno pingpong kraken valr
all: polygon luno marketData pingpong

polygon:
	protoc ./polygon/rest.proto --go_out=./polygon
	protoc ./polygon/stream.proto --go_out=./polygon

luno:
	protoc ./luno/stream.proto --go_out=./luno

marketData:
	protoc ./marketData/marketData.proto --go_out=./marketData
	protoc ./marketData/marketData.proto --csharp_out=./marketData


pingpong:
	protoc ./pingpong/pingpong.proto --go_out=./pingpong

kraken:
	protoc ./kraken/kraken.proto --go_out=./kraken

valr:
	protoc ./valr/valr.proto --go_out=./valr

