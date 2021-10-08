# SaleLineWriteLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The sale line description | 
**ProductId** | **int32** | The product ID | 
**Quantity** | **int32** | The sale line quantity | 
**SaleId** | **int32** | The sale ID | 

## Methods

### NewSaleLineWriteLd

`func NewSaleLineWriteLd(description string, productId int32, quantity int32, saleId int32, ) *SaleLineWriteLd`

NewSaleLineWriteLd instantiates a new SaleLineWriteLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleLineWriteLdWithDefaults

`func NewSaleLineWriteLdWithDefaults() *SaleLineWriteLd`

NewSaleLineWriteLdWithDefaults instantiates a new SaleLineWriteLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *SaleLineWriteLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *SaleLineWriteLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *SaleLineWriteLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *SaleLineWriteLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *SaleLineWriteLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *SaleLineWriteLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *SaleLineWriteLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *SaleLineWriteLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *SaleLineWriteLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *SaleLineWriteLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *SaleLineWriteLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *SaleLineWriteLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetDescription

`func (o *SaleLineWriteLd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaleLineWriteLd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaleLineWriteLd) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProductId

`func (o *SaleLineWriteLd) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SaleLineWriteLd) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SaleLineWriteLd) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *SaleLineWriteLd) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SaleLineWriteLd) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SaleLineWriteLd) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSaleId

`func (o *SaleLineWriteLd) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *SaleLineWriteLd) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *SaleLineWriteLd) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


