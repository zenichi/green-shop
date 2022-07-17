module github.com/zenichi/green-shop/green-api

go 1.18

require (
	github.com/go-playground/validator/v10 v10.11.0
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.48.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/rs/xid v1.4.0
	github.com/stretchr/testify v1.8.0
	github.com/zenichi/green-api/pricing-service v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20220712014510-0a85c31ab51e // indirect
)

replace github.com/zenichi/green-api/pricing-service => ../pricing-service
