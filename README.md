# green-shop
E-commerce exercise

This is an early stage E-commerce website. 

## Services

### Product API [./green-api](./green-api)
Simple Go based JSON API built using the Gorilla framework and GRPC Client. The API allows CRUD based operations on a product list.

#### The Interface
Example http request/responses ara placed in the [green-api-doc](./green-api/README.md)

#### Instructions

- You can run the server simply with `$ make run` in the [green-api](./green-api/) directory
- Tests are run with `make tests` in the [green-api](./green-api/) directory


### Pricing GRPC Server [./pricing-service](./pricing-service)
Simple Go based grpc Server to serve currency rates exchange. It's consumed by Product API.

#### The Interface
```
$ grpcurl --plaintext localhost:9085 describe RateService

RateService is a service:
service RateService {
  rpc GetRate ( .RateRequest ) returns ( .RateResponse );
}
```
```
$ grpcurl --plaintext localhost:9085 describe .RateRequest  

RateRequest is a message:
message RateRequest {
  string FromCurrency = 1;
  string ToCurrency = 2;
}
```
```
$ grpcurl --plaintext localhost:9085 describe.RateResponse

RateResponse is a message:
message RateResponse {
  double Rate = 1;
}
```



#### Instructions

- You can run the server simply with `$ make run` in the [pricing-service](./pricing-service/) directory
- Proto file can be re-generated with `$ make protos` in the [pricing-service](./pricing-service/) directory

### Frontend website [./frontend-vue](./frontend-vue)
VueJS website for presenting the Product API information. It consumes Product API endpoints.

#### Instructions

- You can run the SPA server with `$ yarn serve` in the [frontend-vue](./frontend-vue/) directory



## Next steps (todo)
* API - add postgress DB instead of in memory storage
* API - add swagger doc
* All - add docker files 
* All - add configuration files
* Front-End - integrate with all endpoints of REST API
* Pricing-service - read currency rates from the external API