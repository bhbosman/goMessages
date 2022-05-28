module github.com/bhbosman/goMessages

go 1.15

require (
	github.com/bhbosman/gocommon v0.0.0-20210817162338-4486f68fa3f4
	github.com/bhbosman/goerrors v0.0.0-20210201065523-bb3e832fa9ab
	github.com/bhbosman/goprotoextra v0.0.2-0.20210817141206-117becbef7c7
	google.golang.org/protobuf v1.25.0
)

//replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/golang/mock => ../gomock

replace github.com/bhbosman/gocommon => ../gocommon
