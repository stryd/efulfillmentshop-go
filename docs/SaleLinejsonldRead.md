# SaleLinejsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The sale line description | 
**Id** | Pointer to **int32** | The sale line ID | [optional] [readonly] 
**ProductId** | **int32** | The product ID | 
**Quantity** | **int32** | The sale line quantity | 
**QuantityDelivered** | Pointer to **int32** | The sale line quantity delivered | [optional] [readonly] 
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

### GetContext

`func (o *SaleLinejsonldRead) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SaleLinejsonldRead) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SaleLinejsonldRead) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *SaleLinejsonldRead) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *SaleLinejsonldRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleLinejsonldRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleLinejsonldRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SaleLinejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SaleLinejsonldRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleLinejsonldRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleLinejsonldRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaleLinejsonldRead) HasType() bool`

HasType returns a boolean if a field has been set.

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


