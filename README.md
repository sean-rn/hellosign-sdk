hellosign-sdk
=============

The `hellosign-sdk` module provides a convenient to use client to invoking a subset of 
Hellosign (now called Dropbox Sign) API endpoints.  It provides structs modeling the JSON
request and resposnes for the supported endpoints as well as supported webhook event requests.

The https://openapi-generator.tech/ was used to generate some of the model files

Installation
============

To install Testify, use `go get`:

    go get github.com/sean-rn/hellosign-sdk

Example Usage
=============

Create a client instance
```go
client := hellosign.NewClient(hellosign.WithApiKey("my-api-key"))
```

Invoke endpoints (error handling omitted for brevity)
```go
ctx := context.Background()
resp, err := client.GetEmbeddedSignUrl(ctx, "a-signature-id")
return resp.Embedded.SignURL
```