# \SaleLineApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSaleLine**](SaleLineApi.md#GetSaleLine) | **Get** /sale_lines/{id} | Retrieves a SaleLine resource.
[**GetSaleLines**](SaleLineApi.md#GetSaleLines) | **Get** /sale_lines | Retrieves the collection of SaleLine resources.
[**PatchSaleLine**](SaleLineApi.md#PatchSaleLine) | **Patch** /sale_lines/{id} | Updates the SaleLine resource.
[**PostSaleLine**](SaleLineApi.md#PostSaleLine) | **Post** /sale_lines | Creates a SaleLine resource.
[**PutSaleLine**](SaleLineApi.md#PutSaleLine) | **Put** /sale_lines/{id} | Replaces the SaleLine resource.



## GetSaleLine

> SaleLineRead GetSaleLine(ctx, id).Execute()

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
    resp, r, err := api_client.SaleLineApi.GetSaleLine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.GetSaleLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaleLine`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.GetSaleLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSaleLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SaleLineRead**](SaleLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSaleLines

> []SaleLineRead GetSaleLines(ctx).SaleId(saleId).Page(page).Items(items).Execute()

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
    resp, r, err := api_client.SaleLineApi.GetSaleLines(context.Background()).SaleId(saleId).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.GetSaleLines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaleLines`: []SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.GetSaleLines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSaleLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saleId** | **int32** | The sale ID | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]SaleLineRead**](SaleLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSaleLine

> SaleLineRead PatchSaleLine(ctx, id).SaleLineWrite(saleLineWrite).Execute()

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
    resp, r, err := api_client.SaleLineApi.PatchSaleLine(context.Background(), id).SaleLineWrite(saleLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.PatchSaleLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSaleLine`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.PatchSaleLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSaleLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saleLineWrite** | [**SaleLineWrite**](SaleLineWrite.md) | The updated SaleLine resource | 

### Return type

[**SaleLineRead**](SaleLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSaleLine

> SaleLineRead PostSaleLine(ctx).SaleLineWrite(saleLineWrite).Execute()

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
    resp, r, err := api_client.SaleLineApi.PostSaleLine(context.Background()).SaleLineWrite(saleLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.PostSaleLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSaleLine`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.PostSaleLine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSaleLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saleLineWrite** | [**SaleLineWrite**](SaleLineWrite.md) | The new SaleLine resource | 

### Return type

[**SaleLineRead**](SaleLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSaleLine

> SaleLineRead PutSaleLine(ctx, id).SaleLineWrite(saleLineWrite).Execute()

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
    resp, r, err := api_client.SaleLineApi.PutSaleLine(context.Background(), id).SaleLineWrite(saleLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleLineApi.PutSaleLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSaleLine`: SaleLineRead
    fmt.Fprintf(os.Stdout, "Response from `SaleLineApi.PutSaleLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSaleLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **saleLineWrite** | [**SaleLineWrite**](SaleLineWrite.md) | The updated SaleLine resource | 

### Return type

[**SaleLineRead**](SaleLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

