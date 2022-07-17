## Instructions

Example requests/responses:
```
Request:
HttpGET http:localhost:9081/products/USD

Response:
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 17 Jul 2022 16:51:45 GMT
Content-Length: 193

[
  {
    "id": 1,
    "name": "Garlic",
    "description": "Spicy Chinese garlic",
    "price": 1.5,
    "externalId": "G-45646"
  },
  {
    "id": 2,
    "name": "Onion",
    "description": "Small yellow onion",
    "price": 1.1,
    "externalId": "O-45234"
  }
]

```
```
Request:
HttpGET http:localhost:9081/products/2/USD

Response:
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 17 Jul 2022 16:53:23 GMT
Content-Length: 94

{
  "id": 2,
  "name": "Onion",
  "description": "Small yellow onion",
  "price": 1.1,
  "externalId": "O-45234"
}

```

```
Request:
HttpGET http:localhost:9081/products/2/USD

Response:
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 17 Jul 2022 16:53:23 GMT
Content-Length: 94

{
  "id": 2,
  "name": "Onion",
  "description": "Small yellow onion",
  "price": 1.1,
  "externalId": "O-45234"
}

```
```
Request:
HttpPOST http:localhost:9081/products
{ 
  "name": "Tomato",
  "description": "Small red",
  "price": 0.8,
  "externalId": "O-33234"
}

Response:
HTTP/1.1 200 OK
Date: Sun, 17 Jul 2022 16:55:55 GMT
Content-Length: 0

```
```
Request:
HttpPUT http:localhost:9081/products/2
{ 
  "id":2,
  "name": "Green Tomato",
  "description": "Small green",
  "price": 0.9,
  "externalId": "O-133234"
}

Response:
HTTP/1.1 200 OK
Date: Sun, 17 Jul 2022 16:58:55 GMT
Content-Length: 0

```
```
Request:
HttpDELETE http:localhost:9081/products/2

Response:
HTTP/1.1 200 OK
Date: Sun, 17 Jul 2022 17:00:06 GMT
Content-Length: 0

```
```
Request:
HttpGET http:localhost:9081/info

Response:
HTTP/1.1 200 OK
Date: Sun, 17 Jul 2022 17:07:52 GMT
Content-Length: 80

{
  "message": "Server time: 2022-07-17 19:07:51.293562 +0200 CEST m=+3.734373626"
}

```
