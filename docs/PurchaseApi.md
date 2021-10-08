# \PurchaseApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePurchase**](PurchaseApi.md#DeletePurchase) | **Delete** /purchases/{id} | Removes the Purchase resource.
[**GetPurchase**](PurchaseApi.md#GetPurchase) | **Get** /purchases/{id} | Retrieves a Purchase resource.
[**GetPurchases**](PurchaseApi.md#GetPurchases) | **Get** /purchases | Retrieves the collection of Purchase resources.
[**PatchPurchase**](PurchaseApi.md#PatchPurchase) | **Patch** /purchases/{id} | Updates the Purchase resource.
[**PostPurchase**](PurchaseApi.md#PostPurchase) | **Post** /purchases | Creates a Purchase resource.
[**PutPurchase**](PurchaseApi.md#PutPurchase) | **Put** /purchases/{id} | Replaces the Purchase resource.



## DeletePurchase

> DeletePurchase(ctx, id).Execute()

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
    resp, r, err := api_client.PurchaseApi.DeletePurchase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.DeletePurchase``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePurchaseRequest struct via the builder pattern


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


## GetPurchase

> PurchaseRead GetPurchase(ctx, id).Execute()

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
    resp, r, err := api_client.PurchaseApi.GetPurchase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.GetPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchase`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.GetPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PurchaseRead**](PurchaseRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchases

> []PurchaseRead GetPurchases(ctx).ChannelReference(channelReference).Reference(reference).Page(page).Items(items).Execute()

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
    resp, r, err := api_client.PurchaseApi.GetPurchases(context.Background()).ChannelReference(channelReference).Reference(reference).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.GetPurchases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchases`: []PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.GetPurchases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelReference** | **string** | The purchase channel reference (DEPRECATED) | 
 **reference** | **string** | Your purchase reference (This could be your purchase ID) | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]PurchaseRead**](PurchaseRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPurchase

> PurchaseRead PatchPurchase(ctx, id).PurchaseWrite(purchaseWrite).Execute()

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
    resp, r, err := api_client.PurchaseApi.PatchPurchase(context.Background(), id).PurchaseWrite(purchaseWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.PatchPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPurchase`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.PatchPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseWrite** | [**PurchaseWrite**](PurchaseWrite.md) | The updated Purchase resource | 

### Return type

[**PurchaseRead**](PurchaseRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPurchase

> PurchaseRead PostPurchase(ctx).PurchaseWrite(purchaseWrite).Execute()

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
    resp, r, err := api_client.PurchaseApi.PostPurchase(context.Background()).PurchaseWrite(purchaseWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.PostPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostPurchase`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.PostPurchase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purchaseWrite** | [**PurchaseWrite**](PurchaseWrite.md) | The new Purchase resource | 

### Return type

[**PurchaseRead**](PurchaseRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPurchase

> PurchaseRead PutPurchase(ctx, id).PurchaseWrite(purchaseWrite).Execute()

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
    resp, r, err := api_client.PurchaseApi.PutPurchase(context.Background(), id).PurchaseWrite(purchaseWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PurchaseApi.PutPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPurchase`: PurchaseRead
    fmt.Fprintf(os.Stdout, "Response from `PurchaseApi.PutPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **purchaseWrite** | [**PurchaseWrite**](PurchaseWrite.md) | The updated Purchase resource | 

### Return type

[**PurchaseRead**](PurchaseRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

