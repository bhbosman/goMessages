module github.com/bhbosman/goMessages

go 1.18

require (
	github.com/bhbosman/gocommon v0.0.0-20230328150634-566a0f916878
	github.com/bhbosman/goerrors v0.0.0-20220623084908-4d7bbcd178cf
	github.com/bhbosman/goprotoextra v0.0.2
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/bhbosman/goCommsDefinitions v0.0.0-20230308211643-24fe88b2378e // indirect
	github.com/bhbosman/gomessageblock v0.0.0-20220617132215-32f430d7de62 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/reactivex/rxgo/v2 v2.5.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/fx v1.19.2 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/sys v0.0.0-20220318055525-2edf467146b5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//replace github.com/bhbosman/gocomms => ../gocomms


//replace github.com/bhbosman/gocommon => ../gocommon

//replace github.com/bhbosman/goprotoextra => ../goprotoextra

replace github.com/bhbosman/goCommsDefinitions => ../goCommsDefinitions

replace github.com/bhbosman/goCommsNetListener => ../goCommsNetListener

replace github.com/bhbosman/goConn => ../goConn
