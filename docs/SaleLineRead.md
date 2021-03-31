# SaleLineRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The sale line description | 
**Id** | Pointer to **int32** | The sale line ID | [optional] 
**ProductId** | **int32** | The product ID | 
**Quantity** | **int32** | The sale line quantity | 
**QuantityDelivered** | Pointer to **int32** | The sale line quantity delivered | [optional] 
**SaleId** | **int32** | The sale ID | 

## Methods

### NewSaleLineRead

`func NewSaleLineRead(description string, productId int32, quantity int32, saleId int32, ) *SaleLineRead`

NewSaleLineRead instantiates a new SaleLineRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleLineReadWithDefaults

`func NewSaleLineReadWithDefaults() *SaleLineRead`

NewSaleLineReadWithDefaults instantiates a new SaleLineRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SaleLineRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaleLineRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaleLineRead) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SaleLineRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleLineRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleLineRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SaleLineRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *SaleLineRead) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SaleLineRead) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SaleLineRead) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *SaleLineRead) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SaleLineRead) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SaleLineRead) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetQuantityDelivered

`func (o *SaleLineRead) GetQuantityDelivered() int32`

GetQuantityDelivered returns the QuantityDelivered field if non-nil, zero value otherwise.

### GetQuantityDeliveredOk

`func (o *SaleLineRead) GetQuantityDeliveredOk() (*int32, bool)`

GetQuantityDeliveredOk returns a tuple with the QuantityDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDelivered

`func (o *SaleLineRead) SetQuantityDelivered(v int32)`

SetQuantityDelivered sets QuantityDelivered field to given value.

### HasQuantityDelivered

`func (o *SaleLineRead) HasQuantityDelivered() bool`

HasQuantityDelivered returns a boolean if a field has been set.

### GetSaleId

`func (o *SaleLineRead) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *SaleLineRead) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *SaleLineRead) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


