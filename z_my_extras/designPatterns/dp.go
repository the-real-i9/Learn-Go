package main

import "fmt"

func main() {
	newApiRequestBuilder := newAPIRequestBuilder()
	newApiRequestBuilder2 := newAPIRequestBuilder()

	newApiRequest := newApiRequestBuilder.WithMethod("GET").WithEndpoint("/resource1").WithHeader("Authorization", "Bearer token1").Build()

	newApiRequest2 := newApiRequestBuilder2.WithMethod("POST").WithEndpoint("/resource2").WithHeader("Authorization", "Bearer token2").WithData(map[string]interface{}{"name": "kenny"}).Build()

	fmt.Println(newApiRequest.ToString())
	fmt.Println(newApiRequest2.ToString())
}
