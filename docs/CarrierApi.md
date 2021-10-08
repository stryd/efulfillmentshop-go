# \CarrierApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCarrier**](CarrierApi.md#GetCarrier) | **Get** /carriers/{id} | Retrieves a Carrier resource.
[**GetCarriers**](CarrierApi.md#GetCarriers) | **Get** /carriers | Retrieves the collection of Carrier resources.



## GetCarrier

> CarrierRead GetCarrier(ctx, id).Execute()

Retrieves a Carrier resource.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarrierApi.GetCarrier(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierApi.GetCarrier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCarrier`: CarrierRead
    fmt.Fprintf(os.Stdout, "Response from `CarrierApi.GetCarrier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCarrierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CarrierRead**](CarrierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCarriers

> []CarrierRead GetCarriers(ctx).Page(page).Items(items).Execute()

Retrieves the collection of Carrier resources.

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
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CarrierApi.GetCarriers(context.Background()).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CarrierApi.GetCarriers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCarriers`: []CarrierRead
    fmt.Fprintf(os.Stdout, "Response from `CarrierApi.GetCarriers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCarriersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]CarrierRead**](CarrierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

