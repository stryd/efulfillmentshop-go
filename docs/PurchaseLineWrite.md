# PurchaseLineWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The purchase line description | 
**PlannedDate** | Pointer to **time.Time** | The purchase line planned delivery date | [optional] 
**ProductId** | **int32** | The product ID | 
**PurchaseId** | **int32** | The purchase ID | 
**Quantity** | **int32** | The purchase line quantity | 

## Methods

### NewPurchaseLineWrite

`func NewPurchaseLineWrite(description string, productId int32, purchaseId int32, quantity int32, ) *PurchaseLineWrite`

NewPurchaseLineWrite instantiates a new PurchaseLineWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseLineWriteWithDefaults

`func NewPurchaseLineWriteWithDefaults() *PurchaseLineWrite`

NewPurchaseLineWriteWithDefaults instantiates a new PurchaseLineWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PurchaseLineWrite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseLineWrite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseLineWrite) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPlannedDate

`func (o *PurchaseLineWrite) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseLineWrite) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseLineWrite) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.

### HasPlannedDate

`func (o *PurchaseLineWrite) HasPlannedDate() bool`

HasPlannedDate returns a boolean if a field has been set.

### GetProductId

`func (o *PurchaseLineWrite) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PurchaseLineWrite) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PurchaseLineWrite) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetPurchaseId

`func (o *PurchaseLineWrite) GetPurchaseId() int32`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *PurchaseLineWrite) GetPurchaseIdOk() (*int32, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *PurchaseLineWrite) SetPurchaseId(v int32)`

SetPurchaseId sets PurchaseId field to given value.


### GetQuantity

`func (o *PurchaseLineWrite) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseLineWrite) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseLineWrite) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


