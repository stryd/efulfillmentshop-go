# SaleLinejsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonContext** | Pointer to **string** |  | [optional] [readonly] 
**JsonId** | Pointer to **string** |  | [optional] [readonly] 
**JsonType** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The sale line description | 
**Id** | Pointer to **int32** | The sale line ID | [optional] 
**ProductId** | **int32** | The product ID | 
**Quantity** | **int32** | The sale line quantity | 
**QuantityDelivered** | Pointer to **int32** | The sale line quantity delivered | [optional] 
**SaleId** | **int32** | The sale ID | 

## Methods

### NewSaleLinejsonldRead

`func NewSaleLinejsonldRead(description string, productId int32, quantity int32, saleId int32, ) *SaleLinejsonldRead`

NewSaleLinejsonldRead instantiates a new SaleLinejsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleLinejsonldReadWithDefaults

`func NewSaleLinejsonldReadWithDefaults() *SaleLinejsonldRead`

NewSaleLinejsonldReadWithDefaults instantiates a new SaleLinejsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonContext

`func (o *SaleLinejsonldRead) GetJsonContext() string`

GetJsonContext returns the JsonContext field if non-nil, zero value otherwise.

### GetJsonContextOk

`func (o *SaleLinejsonldRead) GetJsonContextOk() (*string, bool)`

GetJsonContextOk returns a tuple with the JsonContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonContext

`func (o *SaleLinejsonldRead) SetJsonContext(v string)`

SetJsonContext sets JsonContext field to given value.

### HasJsonContext

`func (o *SaleLinejsonldRead) HasJsonContext() bool`

HasJsonContext returns a boolean if a field has been set.

### GetJsonId

`func (o *SaleLinejsonldRead) GetJsonId() string`

GetJsonId returns the JsonId field if non-nil, zero value otherwise.

### GetJsonIdOk

`func (o *SaleLinejsonldRead) GetJsonIdOk() (*string, bool)`

GetJsonIdOk returns a tuple with the JsonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonId

`func (o *SaleLinejsonldRead) SetJsonId(v string)`

SetJsonId sets JsonId field to given value.

### HasJsonId

`func (o *SaleLinejsonldRead) HasJsonId() bool`

HasJsonId returns a boolean if a field has been set.

### GetJsonType

`func (o *SaleLinejsonldRead) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *SaleLinejsonldRead) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *SaleLinejsonldRead) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.

### HasJsonType

`func (o *SaleLinejsonldRead) HasJsonType() bool`

HasJsonType returns a boolean if a field has been set.

### GetDescription

`func (o *SaleLinejsonldRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaleLinejsonldRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaleLinejsonldRead) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SaleLinejsonldRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleLinejsonldRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleLinejsonldRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SaleLinejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductId

`func (o *SaleLinejsonldRead) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SaleLinejsonldRead) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SaleLinejsonldRead) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetQuantity

`func (o *SaleLinejsonldRead) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SaleLinejsonldRead) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SaleLinejsonldRead) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetQuantityDelivered

`func (o *SaleLinejsonldRead) GetQuantityDelivered() int32`

GetQuantityDelivered returns the QuantityDelivered field if non-nil, zero value otherwise.

### GetQuantityDeliveredOk

`func (o *SaleLinejsonldRead) GetQuantityDeliveredOk() (*int32, bool)`

GetQuantityDeliveredOk returns a tuple with the QuantityDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityDelivered

`func (o *SaleLinejsonldRead) SetQuantityDelivered(v int32)`

SetQuantityDelivered sets QuantityDelivered field to given value.

### HasQuantityDelivered

`func (o *SaleLinejsonldRead) HasQuantityDelivered() bool`

HasQuantityDelivered returns a boolean if a field has been set.

### GetSaleId

`func (o *SaleLinejsonldRead) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *SaleLinejsonldRead) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *SaleLinejsonldRead) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


