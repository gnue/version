module github.com/gnue/version/examples/grpc-version

go 1.19

//replace github.com/gnue/version => ../..

require (
	github.com/gnue/version v0.0.0-00010101000000-000000000000
	github.com/jessevdk/go-flags v1.5.0
	golang.org/x/net v0.4.0
	google.golang.org/grpc v1.51.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
