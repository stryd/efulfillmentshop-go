# \AddressApi

All URIs are relative to *http://localhost/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAddressItem**](AddressApi.md#DeleteAddressItem) | **Delete** /addresses/{id} | Removes the Address resource.
[**GetAddressCollection**](AddressApi.md#GetAddressCollection) | **Get** /addresses | Retrieves the collection of Address resources.
[**GetAddressItem**](AddressApi.md#GetAddressItem) | **Get** /addresses/{id} | Retrieves a Address resource.
[**PatchAddressItem**](AddressApi.md#PatchAddressItem) | **Patch** /addresses/{id} | Updates the Address resource.
[**PostAddressCollection**](AddressApi.md#PostAddressCollection) | **Post** /addresses | Creates a Address resource.
[**PutAddressItem**](AddressApi.md#PutAddressItem) | **Put** /addresses/{id} | Replaces the Address resource.



## DeleteAddressItem

> DeleteAddressItem(ctx, id).Execute()

Removes the Address resource.

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
    resp, r, err := api_client.AddressApi.DeleteAddressItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.DeleteAddressItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAddressItemRequest struct via the builder pattern


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


## GetAddressCollection

> []AddressRead GetAddressCollection(ctx).Page(page).Items(items).Execute()

Retrieves the collection of Address resources.

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
    resp, r, err := api_client.AddressApi.GetAddressCollection(context.Background()).Page(page).Items(items).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.GetAddressCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressCollection`: []AddressRead
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.GetAddressCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The collection page number | [default to 1]
 **items** | **int32** | The number of items per page | [default to 30]

### Return type

[**[]AddressRead**](Address-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressItem

> AddressRead GetAddressItem(ctx, id).Execute()

Retrieves a Address resource.

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
    resp, r, err := api_client.AddressApi.GetAddressItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.GetAddressItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAddressItem`: AddressRead
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.GetAddressItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressRead**](Address-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAddressItem

> AddressRead PatchAddressItem(ctx, id).AddressWrite(addressWrite).Execute()

Updates the Address resource.

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
    addressWrite := *openapiclient.NewAddressWrite("City_example", "CountryCode_example", "Name_example", "Street_example", "Zip_example") // AddressWrite | The updated Address resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.PatchAddressItem(context.Background(), id).AddressWrite(addressWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.PatchAddressItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAddressItem`: AddressRead
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.PatchAddressItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAddressItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addressWrite** | [**AddressWrite**](AddressWrite.md) | The updated Address resource | 

### Return type

[**AddressRead**](Address-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddressCollection

> AddressRead PostAddressCollection(ctx).AddressWrite(addressWrite).Execute()

Creates a Address resource.

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
    addressWrite := *openapiclient.NewAddressWrite("City_example", "CountryCode_example", "Name_example", "Street_example", "Zip_example") // AddressWrite | The new Address resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.PostAddressCollection(context.Background()).AddressWrite(addressWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.PostAddressCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAddressCollection`: AddressRead
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.PostAddressCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAddressCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressWrite** | [**AddressWrite**](AddressWrite.md) | The new Address resource | 

### Return type

[**AddressRead**](Address-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAddressItem

> AddressRead PutAddressItem(ctx, id).AddressWrite(addressWrite).Execute()

Replaces the Address resource.

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
    addressWrite := *openapiclient.NewAddressWrite("City_example", "CountryCode_example", "Name_example", "Street_example", "Zip_example") // AddressWrite | The updated Address resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AddressApi.PutAddressItem(context.Background(), id).AddressWrite(addressWrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressApi.PutAddressItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAddressItem`: AddressRead
    fmt.Fprintf(os.Stdout, "Response from `AddressApi.PutAddressItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAddressItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addressWrite** | [**AddressWrite**](AddressWrite.md) | The updated Address resource | 

### Return type

[**AddressRead**](Address-read.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, application/ld+json, text/html
- **Accept**: application/json, application/ld+json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

