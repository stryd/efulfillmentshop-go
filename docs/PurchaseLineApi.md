# \PurchaseLineApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPurchaseLineCollection**](PurchaseLineApi.md#GetPurchaseLineCollection) | **Get** /purchase_lines | Retrieves the collection of PurchaseLine resources.
[**GetPurchaseLineItem**](PurchaseLineApi.md#GetPurchaseLineItem) | **Get** /purchase_lines/{id} | Retrieves a PurchaseLine resource.
[**PatchPurchaseLineItem**](PurchaseLineApi.md#PatchPurchaseLineItem) | **Patch** /purchase_lines/{id} | Updates the PurchaseLine resource.
[**PostPurchaseLineCollection**](PurchaseLineApi.md#PostPurchaseLineCollection) | **Post** /purchase_lines | Creates a PurchaseLine resource.
[**PutPurchaseLineItem**](PurchaseLineApi.md#PutPurchaseLineItem) | **Put** /purchase_lines/{id} | Replaces the PurchaseLine resource.



## GetPurchaseLineCollection

> []PurchaseLineRead GetPurchaseLineCollection(ctx).PurchaseId(purchaseId).Page(page).Items(items).Execute()

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
    resp, r, err := api_client.PurchaseLineApi.GetPurchaseLineCollection(context.Background()).PurchaseId(purchaseId).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.GetPurchaseLineCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchaseLineCollection`: []PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.GetPurchaseLineCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseLineCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseId** | **int32** | The purchase ID | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]PurchaseLineRead**](PurchaseLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseLineItem

> PurchaseLineRead GetPurchaseLineItem(ctx, id).Execute()

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
    resp, r, err := api_client.PurchaseLineApi.GetPurchaseLineItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.GetPurchaseLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchaseLineItem`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.GetPurchaseLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseLineRead**](PurchaseLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPurchaseLineItem

> PurchaseLineRead PatchPurchaseLineItem(ctx, id).PurchaseLineWrite(purchaseLineWrite).Execute()

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
    resp, r, err := api_client.PurchaseLineApi.PatchPurchaseLineItem(context.Background(), id).PurchaseLineWrite(purchaseLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.PatchPurchaseLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPurchaseLineItem`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.PatchPurchaseLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPurchaseLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseLineWrite** | [**PurchaseLineWrite**](PurchaseLineWrite.md) | The updated PurchaseLine resource | 

### Return type

[**PurchaseLineRead**](PurchaseLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPurchaseLineCollection

> PurchaseLineRead PostPurchaseLineCollection(ctx).PurchaseLineWrite(purchaseLineWrite).Execute()

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
    resp, r, err := api_client.PurchaseLineApi.PostPurchaseLineCollection(context.Background()).PurchaseLineWrite(purchaseLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.PostPurchaseLineCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPurchaseLineCollection`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.PostPurchaseLineCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPurchaseLineCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseLineWrite** | [**PurchaseLineWrite**](PurchaseLineWrite.md) | The new PurchaseLine resource | 

### Return type

[**PurchaseLineRead**](PurchaseLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPurchaseLineItem

> PurchaseLineRead PutPurchaseLineItem(ctx, id).PurchaseLineWrite(purchaseLineWrite).Execute()

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
    resp, r, err := api_client.PurchaseLineApi.PutPurchaseLineItem(context.Background(), id).PurchaseLineWrite(purchaseLineWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseLineApi.PutPurchaseLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPurchaseLineItem`: PurchaseLineRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseLineApi.PutPurchaseLineItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPurchaseLineItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseLineWrite** | [**PurchaseLineWrite**](PurchaseLineWrite.md) | The updated PurchaseLine resource | 

### Return type

[**PurchaseLineRead**](PurchaseLine-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

