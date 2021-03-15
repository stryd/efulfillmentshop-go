# \ProductApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProductItem**](ProductApi.md#DeleteProductItem) | **Delete** /products/{id} | Removes the Product resource.
[**GetProductCollection**](ProductApi.md#GetProductCollection) | **Get** /products | Retrieves the collection of Product resources.
[**GetProductItem**](ProductApi.md#GetProductItem) | **Get** /products/{id} | Retrieves a Product resource.
[**PatchProductItem**](ProductApi.md#PatchProductItem) | **Patch** /products/{id} | Updates the Product resource.
[**PostProductCollection**](ProductApi.md#PostProductCollection) | **Post** /products | Creates a Product resource.
[**PutProductItem**](ProductApi.md#PutProductItem) | **Put** /products/{id} | Replaces the Product resource.



## DeleteProductItem

> DeleteProductItem(ctx, id).Execute()

Removes the Product resource.

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
    resp, r, err := api_client.ProductApi.DeleteProductItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.DeleteProductItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProductItemRequest struct via the builder pattern


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


## GetProductCollection

> []ProductReadCollection GetProductCollection(ctx).Name(name).Sku(sku).Barcode(barcode).ChannelId(channelId).ChannelReference(channelReference).ChannelSku(channelSku).Page(page).Items(items).Execute()

Retrieves the collection of Product resources.

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
    name := "name_example" // string | The product name (optional)
    sku := "sku_example" // string | The product warehouse SKU (optional)
    barcode := "barcode_example" // string | The product barcode (optional)
    channelId := int32(56) // int32 | The product channel (optional)
    channelReference := "channelReference_example" // string | The product channel reference (optional)
    channelSku := "channelSku_example" // string | The product channel SKU (optional)
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.GetProductCollection(context.Background()).Name(name).Sku(sku).Barcode(barcode).ChannelId(channelId).ChannelReference(channelReference).ChannelSku(channelSku).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductCollection`: []ProductReadCollection
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The product name | 
 **sku** | **string** | The product warehouse SKU | 
 **barcode** | **string** | The product barcode | 
 **channelId** | **int32** | The product channel | 
 **channelReference** | **string** | The product channel reference | 
 **channelSku** | **string** | The product channel SKU | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]ProductReadCollection**](Product-read_collection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductItem

> ProductReadItem GetProductItem(ctx, id).Execute()

Retrieves a Product resource.

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
    resp, r, err := api_client.ProductApi.GetProductItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProductItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductItem`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProductItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductReadItem**](Product-read_item.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProductItem

> ProductReadItem PatchProductItem(ctx, id).ProductWrite(productWrite).Execute()

Updates the Product resource.

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
    productWrite := *openapiclient.NewProductWrite("Barcode_example", int32(123), "ChannelReference_example", "Name_example") // ProductWrite | The updated Product resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.PatchProductItem(context.Background(), id).ProductWrite(productWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.PatchProductItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProductItem`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.PatchProductItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProductItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productWrite** | [**ProductWrite**](ProductWrite.md) | The updated Product resource | 

### Return type

[**ProductReadItem**](Product-read_item.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProductCollection

> ProductReadItem PostProductCollection(ctx).ProductWrite(productWrite).Execute()

Creates a Product resource.

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
    productWrite := *openapiclient.NewProductWrite("Barcode_example", int32(123), "ChannelReference_example", "Name_example") // ProductWrite | The new Product resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.PostProductCollection(context.Background()).ProductWrite(productWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.PostProductCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostProductCollection`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.PostProductCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productWrite** | [**ProductWrite**](ProductWrite.md) | The new Product resource | 

### Return type

[**ProductReadItem**](Product-read_item.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProductItem

> ProductReadItem PutProductItem(ctx, id).ProductWrite(productWrite).Execute()

Replaces the Product resource.

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
    productWrite := *openapiclient.NewProductWrite("Barcode_example", int32(123), "ChannelReference_example", "Name_example") // ProductWrite | The updated Product resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.PutProductItem(context.Background(), id).ProductWrite(productWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.PutProductItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutProductItem`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.PutProductItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProductItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productWrite** | [**ProductWrite**](ProductWrite.md) | The updated Product resource | 

### Return type

[**ProductReadItem**](Product-read_item.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

