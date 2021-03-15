# \SupplierApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSupplierItem**](SupplierApi.md#DeleteSupplierItem) | **Delete** /suppliers/{id} | Removes the Supplier resource.
[**GetSupplierCollection**](SupplierApi.md#GetSupplierCollection) | **Get** /suppliers | Retrieves the collection of Supplier resources.
[**GetSupplierItem**](SupplierApi.md#GetSupplierItem) | **Get** /suppliers/{id} | Retrieves a Supplier resource.
[**PatchSupplierItem**](SupplierApi.md#PatchSupplierItem) | **Patch** /suppliers/{id} | Updates the Supplier resource.
[**PostSupplierCollection**](SupplierApi.md#PostSupplierCollection) | **Post** /suppliers | Creates a Supplier resource.
[**PutSupplierItem**](SupplierApi.md#PutSupplierItem) | **Put** /suppliers/{id} | Replaces the Supplier resource.



## DeleteSupplierItem

> DeleteSupplierItem(ctx, id).Execute()

Removes the Supplier resource.

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
    resp, r, err := api_client.SupplierApi.DeleteSupplierItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.DeleteSupplierItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSupplierItemRequest struct via the builder pattern


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


## GetSupplierCollection

> []SupplierRead GetSupplierCollection(ctx).Page(page).Items(items).Execute()

Retrieves the collection of Supplier resources.

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
    resp, r, err := api_client.SupplierApi.GetSupplierCollection(context.Background()).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.GetSupplierCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupplierCollection`: []SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.GetSupplierCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplierCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]SupplierRead**](Supplier-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupplierItem

> SupplierRead GetSupplierItem(ctx, id).Execute()

Retrieves a Supplier resource.

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
    resp, r, err := api_client.SupplierApi.GetSupplierItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.GetSupplierItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupplierItem`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.GetSupplierItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplierItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupplierRead**](Supplier-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSupplierItem

> SupplierRead PatchSupplierItem(ctx, id).SupplierWrite(supplierWrite).Execute()

Updates the Supplier resource.

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
    supplierWrite := *openapiclient.NewSupplierWrite("Name_example") // SupplierWrite | The updated Supplier resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupplierApi.PatchSupplierItem(context.Background(), id).SupplierWrite(supplierWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.PatchSupplierItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSupplierItem`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.PatchSupplierItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSupplierItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplierWrite** | [**SupplierWrite**](SupplierWrite.md) | The updated Supplier resource | 

### Return type

[**SupplierRead**](Supplier-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSupplierCollection

> SupplierRead PostSupplierCollection(ctx).SupplierWrite(supplierWrite).Execute()

Creates a Supplier resource.

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
    supplierWrite := *openapiclient.NewSupplierWrite("Name_example") // SupplierWrite | The new Supplier resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupplierApi.PostSupplierCollection(context.Background()).SupplierWrite(supplierWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.PostSupplierCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSupplierCollection`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.PostSupplierCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSupplierCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supplierWrite** | [**SupplierWrite**](SupplierWrite.md) | The new Supplier resource | 

### Return type

[**SupplierRead**](Supplier-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSupplierItem

> SupplierRead PutSupplierItem(ctx, id).SupplierWrite(supplierWrite).Execute()

Replaces the Supplier resource.

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
    supplierWrite := *openapiclient.NewSupplierWrite("Name_example") // SupplierWrite | The updated Supplier resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SupplierApi.PutSupplierItem(context.Background(), id).SupplierWrite(supplierWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.PutSupplierItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSupplierItem`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.PutSupplierItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSupplierItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplierWrite** | [**SupplierWrite**](SupplierWrite.md) | The updated Supplier resource | 

### Return type

[**SupplierRead**](Supplier-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

