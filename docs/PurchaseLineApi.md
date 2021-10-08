# \PurchaseLineApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPurchaseLine**](PurchaseLineApi.md#GetPurchaseLine) | **Get** /purchase_lines/{id} | Retrieves a PurchaseLine resource.
[**GetPurchaseLines**](PurchaseLineApi.md#GetPurchaseLines) | **Get** /purchase_lines | Retrieves the collection of PurchaseLine resources.
[**PatchPurchaseLine**](PurchaseLineApi.md#PatchPurchaseLine) | **Patch** /purchase_lines/{id} | Updates the PurchaseLine resource.
[**PostPurchaseLine**](PurchaseLineApi.md#PostPurchaseLine) | **Post** /purchase_lines | Creates a PurchaseLine resource.
[**PutPurchaseLine**](PurchaseLineApi.md#PutPurchaseLine) | **Put** /purchase_lines/{id} | Replaces the PurchaseLine resource.



## GetPurchaseLine

> PurchaseLineRead GetPurchaseLine(ctx, id).Execute()

Retrieves a PurchaseLine resource.

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
    resp, r, err := api_client.PurchaseLineApi.GetPurchaseLine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.GetPurchaseLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchaseLine`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.GetPurchaseLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseLineRead**](PurchaseLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseLines

> []PurchaseLineRead GetPurchaseLines(ctx).PurchaseId(purchaseId).Page(page).Items(items).Execute()

Retrieves the collection of PurchaseLine resources.

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
    purchaseId := int32(56) // int32 | The purchase ID (optional)
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseLineApi.GetPurchaseLines(context.Background()).PurchaseId(purchaseId).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.GetPurchaseLines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchaseLines`: []PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.GetPurchaseLines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseLinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseId** | **int32** | The purchase ID | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]PurchaseLineRead**](PurchaseLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPurchaseLine

> PurchaseLineRead PatchPurchaseLine(ctx, id).PurchaseLineWrite(purchaseLineWrite).Execute()

Updates the PurchaseLine resource.

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
    purchaseLineWrite := *openapiclient.NewPurchaseLineWrite("Description_example", int32(123), int32(123), int32(123)) // PurchaseLineWrite | The updated PurchaseLine resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseLineApi.PatchPurchaseLine(context.Background(), id).PurchaseLineWrite(purchaseLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.PatchPurchaseLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPurchaseLine`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.PatchPurchaseLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPurchaseLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseLineWrite** | [**PurchaseLineWrite**](PurchaseLineWrite.md) | The updated PurchaseLine resource | 

### Return type

[**PurchaseLineRead**](PurchaseLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPurchaseLine

> PurchaseLineRead PostPurchaseLine(ctx).PurchaseLineWrite(purchaseLineWrite).Execute()

Creates a PurchaseLine resource.

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
    purchaseLineWrite := *openapiclient.NewPurchaseLineWrite("Description_example", int32(123), int32(123), int32(123)) // PurchaseLineWrite | The new PurchaseLine resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseLineApi.PostPurchaseLine(context.Background()).PurchaseLineWrite(purchaseLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.PostPurchaseLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPurchaseLine`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.PostPurchaseLine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPurchaseLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseLineWrite** | [**PurchaseLineWrite**](PurchaseLineWrite.md) | The new PurchaseLine resource | 

### Return type

[**PurchaseLineRead**](PurchaseLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPurchaseLine

> PurchaseLineRead PutPurchaseLine(ctx, id).PurchaseLineWrite(purchaseLineWrite).Execute()

Replaces the PurchaseLine resource.

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
    purchaseLineWrite := *openapiclient.NewPurchaseLineWrite("Description_example", int32(123), int32(123), int32(123)) // PurchaseLineWrite | The updated PurchaseLine resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseLineApi.PutPurchaseLine(context.Background(), id).PurchaseLineWrite(purchaseLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.PutPurchaseLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPurchaseLine`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.PutPurchaseLine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPurchaseLineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseLineWrite** | [**PurchaseLineWrite**](PurchaseLineWrite.md) | The updated PurchaseLine resource | 

### Return type

[**PurchaseLineRead**](PurchaseLineRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

