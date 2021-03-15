# \SaleLineApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSaleLineCollection**](SaleLineApi.md#GetSaleLineCollection) | **Get** /sale_lines | Retrieves the collection of SaleLine resources.
[**GetSaleLineItem**](SaleLineApi.md#GetSaleLineItem) | **Get** /sale_lines/{id} | Retrieves a SaleLine resource.
[**PatchSaleLineItem**](SaleLineApi.md#PatchSaleLineItem) | **Patch** /sale_lines/{id} | Updates the SaleLine resource.
[**PostSaleLineCollection**](SaleLineApi.md#PostSaleLineCollection) | **Post** /sale_lines | Creates a SaleLine resource.
[**PutSaleLineItem**](SaleLineApi.md#PutSaleLineItem) | **Put** /sale_lines/{id} | Replaces the SaleLine resource.



## GetSaleLineCollection

> []SaleLineRead GetSaleLineCollection(ctx).SaleId(saleId).Page(page).Items(items).Execute()

Retrieves the collection of SaleLine resources.

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
    saleId := int32(56) // int32 | The sale ID (optional)
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleLineApi.GetSaleLineCollection(context.Background()).SaleId(saleId).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.GetSaleLineCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaleLineCollection`: []SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.GetSaleLineCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSaleLineCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saleId** | **int32** | The sale ID | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]SaleLineRead**](SaleLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSaleLineItem

> SaleLineRead GetSaleLineItem(ctx, id).Execute()

Retrieves a SaleLine resource.

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
    resp, r, err := api_client.SaleLineApi.GetSaleLineItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.GetSaleLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaleLineItem`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.GetSaleLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSaleLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SaleLineRead**](SaleLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSaleLineItem

> SaleLineRead PatchSaleLineItem(ctx, id).SaleLineWrite(saleLineWrite).Execute()

Updates the SaleLine resource.

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
    saleLineWrite := *openapiclient.NewSaleLineWrite("Description_example", int32(123), int32(123), int32(123)) // SaleLineWrite | The updated SaleLine resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleLineApi.PatchSaleLineItem(context.Background(), id).SaleLineWrite(saleLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.PatchSaleLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSaleLineItem`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.PatchSaleLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSaleLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saleLineWrite** | [**SaleLineWrite**](SaleLineWrite.md) | The updated SaleLine resource | 

### Return type

[**SaleLineRead**](SaleLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSaleLineCollection

> SaleLineRead PostSaleLineCollection(ctx).SaleLineWrite(saleLineWrite).Execute()

Creates a SaleLine resource.

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
    saleLineWrite := *openapiclient.NewSaleLineWrite("Description_example", int32(123), int32(123), int32(123)) // SaleLineWrite | The new SaleLine resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleLineApi.PostSaleLineCollection(context.Background()).SaleLineWrite(saleLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.PostSaleLineCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSaleLineCollection`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.PostSaleLineCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSaleLineCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saleLineWrite** | [**SaleLineWrite**](SaleLineWrite.md) | The new SaleLine resource | 

### Return type

[**SaleLineRead**](SaleLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSaleLineItem

> SaleLineRead PutSaleLineItem(ctx, id).SaleLineWrite(saleLineWrite).Execute()

Replaces the SaleLine resource.

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
    saleLineWrite := *openapiclient.NewSaleLineWrite("Description_example", int32(123), int32(123), int32(123)) // SaleLineWrite | The updated SaleLine resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleLineApi.PutSaleLineItem(context.Background(), id).SaleLineWrite(saleLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.PutSaleLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSaleLineItem`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.PutSaleLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSaleLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saleLineWrite** | [**SaleLineWrite**](SaleLineWrite.md) | The updated SaleLine resource | 

### Return type

[**SaleLineRead**](SaleLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

