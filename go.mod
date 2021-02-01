module github.com/bhbosman/goMessages

go 1.15

require (
	github.com/bhbosman/gocommon v0.0.0-20201004145117-eae02ab42c9a
	github.com/bhbosman/goerrors v0.0.0-20200918064252-e47717b09c4f
	github.com/bhbosman/goprotoextra v0.0.1
	github.com/golang/protobuf v1.4.2
	google.golang.org/protobuf v1.25.0
)

replace github.com/reactivex/rxgo/v2 v2.1.0 => github.com/bhbosman/rxgo/v2 v2.1.1-0.20200922152528-6aef42e76e00

replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/bhbosman/goprotoextra => ../goprotoextra
