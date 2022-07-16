module github.com/zenichi/green-shop/green-api

go 1.18

require github.com/sirupsen/logrus v1.8.1

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/rs/xid v1.4.0
	github.com/stretchr/testify v1.8.0
	github.com/zenichi/green-api/pricing-service v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20220712014510-0a85c31ab51e // indirect
)

replace github.com/zenichi/green-api/pricing-service => ../pricing-service
