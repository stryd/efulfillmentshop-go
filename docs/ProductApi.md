# \ProductApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProduct**](ProductApi.md#DeleteProduct) | **Delete** /products/{id} | Removes the Product resource.
[**GetProduct**](ProductApi.md#GetProduct) | **Get** /products/{id} | Retrieves a Product resource.
[**GetProducts**](ProductApi.md#GetProducts) | **Get** /products | Retrieves the collection of Product resources.
[**PatchProduct**](ProductApi.md#PatchProduct) | **Patch** /products/{id} | Updates the Product resource.
[**PostProduct**](ProductApi.md#PostProduct) | **Post** /products | Creates a Product resource.
[**PutProduct**](ProductApi.md#PutProduct) | **Put** /products/{id} | Replaces the Product resource.



## DeleteProduct

> DeleteProduct(ctx, id).Execute()

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
    resp, r, err := api_client.ProductApi.DeleteProduct(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.DeleteProduct``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


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


## GetProduct

> ProductReadItem GetProduct(ctx, id).Execute()

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
    resp, r, err := api_client.ProductApi.GetProduct(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProduct`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProductReadItem**](ProductReadItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> []ProductReadCollection GetProducts(ctx).Name(name).Sku(sku).Barcode(barcode).ChannelReference(channelReference).Reference(reference).Page(page).Items(items).Execute()

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
    channelReference := "channelReference_example" // string | The product channel reference (DEPRECATED) (optional)
    reference := "reference_example" // string | Your product reference (This could be your product ID) (optional)
    page := int32(56) // int32 | The collection page number (optional) (default to 1)
    items := int32(56) // int32 | The number of items per page (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.GetProducts(context.Background()).Name(name).Sku(sku).Barcode(barcode).ChannelReference(channelReference).Reference(reference).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.GetProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProducts`: []ProductReadCollection
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The product name | 
 **sku** | **string** | The product warehouse SKU | 
 **barcode** | **string** | The product barcode | 
 **channelReference** | **string** | The product channel reference (DEPRECATED) | 
 **reference** | **string** | Your product reference (This could be your product ID) | 
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]ProductReadCollection**](ProductReadCollection.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchProduct

> ProductReadItem PatchProduct(ctx, id).ProductWrite(productWrite).Execute()

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
    productWrite := *openapiclient.NewProductWrite("Barcode_example", "Name_example") // ProductWrite | The updated Product resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.PatchProduct(context.Background(), id).ProductWrite(productWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.PatchProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProduct`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.PatchProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productWrite** | [**ProductWrite**](ProductWrite.md) | The updated Product resource | 

### Return type

[**ProductReadItem**](ProductReadItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostProduct

> ProductReadItem PostProduct(ctx).ProductWrite(productWrite).Execute()

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
    productWrite := *openapiclient.NewProductWrite("Barcode_example", "Name_example") // ProductWrite | The new Product resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.PostProduct(context.Background()).ProductWrite(productWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.PostProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostProduct`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.PostProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productWrite** | [**ProductWrite**](ProductWrite.md) | The new Product resource | 

### Return type

[**ProductReadItem**](ProductReadItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutProduct

> ProductReadItem PutProduct(ctx, id).ProductWrite(productWrite).Execute()

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
    productWrite := *openapiclient.NewProductWrite("Barcode_example", "Name_example") // ProductWrite | The updated Product resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ProductApi.PutProduct(context.Background(), id).ProductWrite(productWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductApi.PutProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutProduct`: ProductReadItem
    fmt.Fprintf(os.Stdout, "Response from `ProductApi.PutProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productWrite** | [**ProductWrite**](ProductWrite.md) | The updated Product resource | 

### Return type

[**ProductReadItem**](ProductReadItem.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

