run:
	go run ./cmd/pricingapp
protos:
	protoc --proto_path=pkg/protos --go_out=pkg/protos/rates --go_opt=paths=source_relative --go-grpc_out=pkg/protos/rates --go-grpc_opt=paths=source_relative pkg/protos/rates.proto
lint: # requires https://golangci-lint.run/ to be installed
	golangci-lint run