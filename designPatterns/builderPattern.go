package main

import (
	"fmt"
)

type APIRequest struct {
	method   string
	endpoint string
	headers  map[string]string
	data     map[string]interface{}
}

func (req *APIRequest) ToString() string {
	return fmt.Sprintf("APIRequest: { Method: %s, Endpoint: %s, Headers: %v, Data: %v }", req.method, req.endpoint, req.headers, req.data)
}

type APIRequestBuilder struct {
	apiRequest *APIRequest
}

/* Constructor function */
func newAPIRequestBuilder() *APIRequestBuilder {
	return &APIRequestBuilder{
		apiRequest: &APIRequest{},
	}
}

/* Methods */
func (reqBuilder *APIRequestBuilder) WithMethod(method string) *APIRequestBuilder {
	reqBuilder.apiRequest.method = method
	return reqBuilder
}

func (reqBuilder *APIRequestBuilder) WithEndpoint(endpoint string) *APIRequestBuilder {
	reqBuilder.apiRequest.endpoint = endpoint
	return reqBuilder
}

func (reqBuilder *APIRequestBuilder) WithHeader(key, value string) *APIRequestBuilder {
	reqBuilder.apiRequest.headers = map[string]string{key: value}
	return reqBuilder
}

func (reqBuilder *APIRequestBuilder) WithData(data map[string]interface{}) *APIRequestBuilder {
	reqBuilder.apiRequest.data = data
	return reqBuilder
}

func (reqBuilder *APIRequestBuilder) Build() *APIRequest {
	return reqBuilder.apiRequest
}

func main() {
	newApiRequestBuilder := newAPIRequestBuilder()
	newApiRequestBuilder2 := newAPIRequestBuilder()

	newApiRequest := newApiRequestBuilder.WithMethod("GET").WithEndpoint("/resource1").WithHeader("Authorization", "Bearer token1").Build()

	newApiRequest2 := newApiRequestBuilder2.WithMethod("POST").WithEndpoint("/resource2").WithHeader("Authorization", "Bearer token2").WithData(map[string]interface{}{"name": "kenny"}).Build()

	fmt.Println(newApiRequest.ToString())
	fmt.Println(newApiRequest2.ToString())
}
