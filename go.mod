module github.com/bhbosman/goMessages

go 1.18

require (
	github.com/bhbosman/gocommon v0.0.0-20220620062839-d40b45675900
	github.com/bhbosman/goerrors v0.0.0-20220623084908-4d7bbcd178cf
	github.com/bhbosman/goprotoextra v0.0.2-0.20210817141206-117becbef7c7
	github.com/reactivex/rxgo/v2 v2.5.0
	google.golang.org/protobuf v1.25.0
)

require (
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	github.com/bhbosman/gomessageblock v0.0.0-20220617132215-32f430d7de62 // indirect
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

replace github.com/bhbosman/goprotoextra => ../goprotoextra
