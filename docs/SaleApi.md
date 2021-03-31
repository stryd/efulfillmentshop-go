# \SaleApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmSale**](SaleApi.md#ConfirmSale) | **Post** /sales/{id}/confirm | Confirms a sale.
[**GetSaleCollection**](SaleApi.md#GetSaleCollection) | **Get** /sales | Retrieves the collection of Sale resources.
[**GetSaleItem**](SaleApi.md#GetSaleItem) | **Get** /sales/{id} | Retrieves a Sale resource.
[**PostSaleCollection**](SaleApi.md#PostSaleCollection) | **Post** /sales | Creates a Sale resource.



## ConfirmSale

> SaleRead ConfirmSale(ctx, id).Body(body).Execute()

Confirms a sale.

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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleApi.ConfirmSale(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleApi.ConfirmSale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmSale`: SaleRead
    fmt.Fprintf(os.Stdout, "Response from `SaleApi.ConfirmSale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmSaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**SaleRead**](Sale-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json
- **Accept**: application/json, application/ld+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSaleCollection

> []SaleRead GetSaleCollection(ctx).ChannelId(channelId).ChannelReference(channelReference).Page(page).Items(items).Execute()

Retrieves the collection of Sale resources.

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
    channelId := int32(56) // int32 | The sale channel (optional)
    channelReference := "channelReference_example" // string | The sale channel reference (optional)
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleApi.GetSaleCollection(context.Background()).ChannelId(channelId).ChannelReference(channelReference).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleApi.GetSaleCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaleCollection`: []SaleRead
    fmt.Fprintf(os.Stdout, "Response from `SaleApi.GetSaleCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSaleCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **int32** | The sale channel | 
 **channelReference** | **string** | The sale channel reference | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]SaleRead**](Sale-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSaleItem

> SaleRead GetSaleItem(ctx, id).Execute()

Retrieves a Sale resource.

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
    resp, r, err := api_client.SaleApi.GetSaleItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleApi.GetSaleItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSaleItem`: SaleRead
    fmt.Fprintf(os.Stdout, "Response from `SaleApi.GetSaleItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSaleItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SaleRead**](Sale-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSaleCollection

> SaleRead PostSaleCollection(ctx).SaleWrite(saleWrite).Execute()

Creates a Sale resource.

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
    saleWrite := *openapiclient.NewSaleWrite(int32(123), "ChannelReference_example", int32(123), int32(123)) // SaleWrite | The new Sale resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SaleApi.PostSaleCollection(context.Background()).SaleWrite(saleWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SaleApi.PostSaleCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSaleCollection`: SaleRead
    fmt.Fprintf(os.Stdout, "Response from `SaleApi.PostSaleCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSaleCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saleWrite** | [**SaleWrite**](SaleWrite.md) | The new Sale resource | 

### Return type

[**SaleRead**](Sale-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

