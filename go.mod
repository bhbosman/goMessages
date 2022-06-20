module github.com/bhbosman/goMessages

go 1.18

require (
	github.com/bhbosman/gocommon v0.0.0-20220620062839-d40b45675900
	github.com/bhbosman/goerrors v0.0.0-20210201065523-bb3e832fa9ab
	github.com/bhbosman/goprotoextra v0.0.2-0.20210817141206-117becbef7c7
	google.golang.org/protobuf v1.25.0
)

require github.com/bhbosman/gologging v0.0.0-20200921180328-d29fc55c00bc // indirect

require (
	//github.com/bhbosman/gologging v0.0.0-20200921180328-d29fc55c00bc // indirect
	github.com/bhbosman/gomessageblock v0.0.0-20210901070622-be36a3f8d303 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.12.0 // indirect
	go.uber.org/fx v1.14.2 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/sys v0.0.0-20220318055525-2edf467146b5 // indirect
)

//replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/golang/mock => ../gomock

replace github.com/bhbosman/gocommon => ../gocommon
