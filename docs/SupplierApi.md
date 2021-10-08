# \SupplierApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSupplier**](SupplierApi.md#DeleteSupplier) | **Delete** /suppliers/{id} | Removes the Supplier resource.
[**GetSupplier**](SupplierApi.md#GetSupplier) | **Get** /suppliers/{id} | Retrieves a Supplier resource.
[**GetSuppliers**](SupplierApi.md#GetSuppliers) | **Get** /suppliers | Retrieves the collection of Supplier resources.
[**PatchSupplier**](SupplierApi.md#PatchSupplier) | **Patch** /suppliers/{id} | Updates the Supplier resource.
[**PostSupplier**](SupplierApi.md#PostSupplier) | **Post** /suppliers | Creates a Supplier resource.
[**PutSupplier**](SupplierApi.md#PutSupplier) | **Put** /suppliers/{id} | Replaces the Supplier resource.



## DeleteSupplier

> DeleteSupplier(ctx, id).Execute()

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
    resp, r, err := api_client.SupplierApi.DeleteSupplier(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.DeleteSupplier``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSupplierRequest struct via the builder pattern


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


## GetSupplier

> SupplierRead GetSupplier(ctx, id).Execute()

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
    resp, r, err := api_client.SupplierApi.GetSupplier(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.GetSupplier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupplier`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.GetSupplier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupplierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupplierRead**](SupplierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuppliers

> []SupplierRead GetSuppliers(ctx).Page(page).Items(items).Execute()

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
    resp, r, err := api_client.SupplierApi.GetSuppliers(context.Background()).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.GetSuppliers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSuppliers`: []SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.GetSuppliers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSuppliersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]SupplierRead**](SupplierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSupplier

> SupplierRead PatchSupplier(ctx, id).SupplierWrite(supplierWrite).Execute()

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
    resp, r, err := api_client.SupplierApi.PatchSupplier(context.Background(), id).SupplierWrite(supplierWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.PatchSupplier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSupplier`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.PatchSupplier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSupplierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplierWrite** | [**SupplierWrite**](SupplierWrite.md) | The updated Supplier resource | 

### Return type

[**SupplierRead**](SupplierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSupplier

> SupplierRead PostSupplier(ctx).SupplierWrite(supplierWrite).Execute()

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
    resp, r, err := api_client.SupplierApi.PostSupplier(context.Background()).SupplierWrite(supplierWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.PostSupplier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSupplier`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.PostSupplier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSupplierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supplierWrite** | [**SupplierWrite**](SupplierWrite.md) | The new Supplier resource | 

### Return type

[**SupplierRead**](SupplierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSupplier

> SupplierRead PutSupplier(ctx, id).SupplierWrite(supplierWrite).Execute()

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
    resp, r, err := api_client.SupplierApi.PutSupplier(context.Background(), id).SupplierWrite(supplierWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SupplierApi.PutSupplier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSupplier`: SupplierRead
    fmt.Fprintf(os.Stdout, "Response from `SupplierApi.PutSupplier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSupplierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **supplierWrite** | [**SupplierWrite**](SupplierWrite.md) | The updated Supplier resource | 

### Return type

[**SupplierRead**](SupplierRead.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

