# SaleLineReadLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The sale line description | 
**Id** | Pointer to **int32** | The sale line ID | [optional] 
**ProductId** | **int32** | The product ID | 
**Quantity** | **int32** | The sale line quantity | 
**QuantityDelivered** | Pointer to **int32** | The sale line quantity delivered | [optional] 
**SaleId** | **int32** | The sale ID | 

## Methods

### NewSaleLineReadLd

`func NewSaleLineReadLd(description string, productId int32, quantity int32, saleId int32, ) *SaleLineReadLd`

NewSaleLineReadLd instantiates a new SaleLineReadLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleLineReadLdWithDefaults

`func NewSaleLineReadLdWithDefaults() *SaleLineReadLd`

NewSaleLineReadLdWithDefaults instantiates a new SaleLineReadLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *SaleLineReadLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *SaleLineReadLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *SaleLineReadLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *SaleLineReadLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *SaleLineReadLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *SaleLineReadLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *SaleLineReadLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *SaleLineReadLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *SaleLineReadLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *SaleLineReadLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *SaleLineReadLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *SaleLineReadLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetDescription

`func (o *SaleLineReadLd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaleLineReadLd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaleLineReadLd) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SaleLineReadLd) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleLineReadLd) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleLineReadLd) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SaleLineReadLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *SaleLineReadLd) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SaleLineReadLd) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SaleLineReadLd) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *SaleLineReadLd) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SaleLineReadLd) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SaleLineReadLd) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetQuantityDelivered

`func (o *SaleLineReadLd) GetQuantityDelivered() int32`

GetQuantityDelivered returns the QuantityDelivered field if non-nil, zero value otherwise.

### GetQuantityDeliveredOk

`func (o *SaleLineReadLd) GetQuantityDeliveredOk() (*int32, bool)`

GetQuantityDeliveredOk returns a tuple with the QuantityDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDelivered

`func (o *SaleLineReadLd) SetQuantityDelivered(v int32)`

SetQuantityDelivered sets QuantityDelivered field to given value.

### HasQuantityDelivered

`func (o *SaleLineReadLd) HasQuantityDelivered() bool`

HasQuantityDelivered returns a boolean if a field has been set.

### GetSaleId

`func (o *SaleLineReadLd) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *SaleLineReadLd) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *SaleLineReadLd) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


