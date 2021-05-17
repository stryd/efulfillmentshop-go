# \PurchaseApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchaseItem**](PurchaseApi.md#DeletePurchaseItem) | **Delete** /purchases/{id} | Removes the Purchase resource.
[**GetPurchaseCollection**](PurchaseApi.md#GetPurchaseCollection) | **Get** /purchases | Retrieves the collection of Purchase resources.
[**GetPurchaseItem**](PurchaseApi.md#GetPurchaseItem) | **Get** /purchases/{id} | Retrieves a Purchase resource.
[**PatchPurchaseItem**](PurchaseApi.md#PatchPurchaseItem) | **Patch** /purchases/{id} | Updates the Purchase resource.
[**PostPurchaseCollection**](PurchaseApi.md#PostPurchaseCollection) | **Post** /purchases | Creates a Purchase resource.
[**PutPurchaseItem**](PurchaseApi.md#PutPurchaseItem) | **Put** /purchases/{id} | Replaces the Purchase resource.



## DeletePurchaseItem

> DeletePurchaseItem(ctx, id).Execute()

Removes the Purchase resource.

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
    resp, r, err := api_client.PurchaseApi.DeletePurchaseItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.DeletePurchaseItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePurchaseItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseCollection

> []PurchaseRead GetPurchaseCollection(ctx).ChannelReference(channelReference).Reference(reference).Page(page).Items(items).Execute()

Retrieves the collection of Purchase resources.

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
    channelReference := "channelReference_example" // string | The purchase channel reference (DEPRECATED) (optional)
    reference := "reference_example" // string | Your purchase reference (This could be your purchase ID) (optional)
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseApi.GetPurchaseCollection(context.Background()).ChannelReference(channelReference).Reference(reference).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.GetPurchaseCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchaseCollection`: []PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.GetPurchaseCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelReference** | **string** | The purchase channel reference (DEPRECATED) | 
 **reference** | **string** | Your purchase reference (This could be your purchase ID) | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]PurchaseRead**](Purchase-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchaseItem

> PurchaseRead GetPurchaseItem(ctx, id).Execute()

Retrieves a Purchase resource.

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
    resp, r, err := api_client.PurchaseApi.GetPurchaseItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.GetPurchaseItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchaseItem`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.GetPurchaseItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseRead**](Purchase-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPurchaseItem

> PurchaseRead PatchPurchaseItem(ctx, id).PurchaseWrite(purchaseWrite).Execute()

Updates the Purchase resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    purchaseWrite := *openapiclient.NewPurchaseWrite(time.Now(), int32(123)) // PurchaseWrite | The updated Purchase resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseApi.PatchPurchaseItem(context.Background(), id).PurchaseWrite(purchaseWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.PatchPurchaseItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPurchaseItem`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.PatchPurchaseItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPurchaseItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseWrite** | [**PurchaseWrite**](PurchaseWrite.md) | The updated Purchase resource | 

### Return type

[**PurchaseRead**](Purchase-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPurchaseCollection

> PurchaseRead PostPurchaseCollection(ctx).PurchaseWrite(purchaseWrite).Execute()

Creates a Purchase resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    purchaseWrite := *openapiclient.NewPurchaseWrite(time.Now(), int32(123)) // PurchaseWrite | The new Purchase resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseApi.PostPurchaseCollection(context.Background()).PurchaseWrite(purchaseWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.PostPurchaseCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPurchaseCollection`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.PostPurchaseCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPurchaseCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseWrite** | [**PurchaseWrite**](PurchaseWrite.md) | The new Purchase resource | 

### Return type

[**PurchaseRead**](Purchase-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPurchaseItem

> PurchaseRead PutPurchaseItem(ctx, id).PurchaseWrite(purchaseWrite).Execute()

Replaces the Purchase resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    purchaseWrite := *openapiclient.NewPurchaseWrite(time.Now(), int32(123)) // PurchaseWrite | The updated Purchase resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PurchaseApi.PutPurchaseItem(context.Background(), id).PurchaseWrite(purchaseWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.PutPurchaseItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPurchaseItem`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.PutPurchaseItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPurchaseItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseWrite** | [**PurchaseWrite**](PurchaseWrite.md) | The updated Purchase resource | 

### Return type

[**PurchaseRead**](Purchase-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

