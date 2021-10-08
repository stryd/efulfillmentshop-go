# \AuthenticationTokenApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAuthentication**](AuthenticationTokenApi.md#PostAuthentication) | **Post** /authentication_token | Performs login, returning a token on success.



## PostAuthentication

> InlineResponse200 PostAuthentication(ctx).UserAuthenticate(userAuthenticate).Execute()

Performs login, returning a token on success.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userAuthenticate := *openapiclient.NewUserAuthenticate("Email_example", "Password_example") // UserAuthenticate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthenticationTokenApi.PostAuthentication(context.Background()).UserAuthenticate(userAuthenticate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationTokenApi.PostAuthentication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAuthentication`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationTokenApi.PostAuthentication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuthenticate** | [**UserAuthenticate**](UserAuthenticate.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

