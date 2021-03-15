# PurchaseLineRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The purchase line description | 
**Id** | Pointer to **int32** | The purchase line ID | [optional] [readonly] 
**PlannedDate** | Pointer to **time.Time** | The purchase line planned delivery date | [optional] 
**ProductId** | **int32** | The product ID | 
**PurchaseId** | **int32** | The purchase ID | 
**QtyReceived** | Pointer to **int32** | The purchase line quantity received | [optional] [readonly] 
**Quantity** | **int32** | The purchase line quantity | 

## Methods

### NewPurchaseLineRead

`func NewPurchaseLineRead(description string, productId int32, purchaseId int32, quantity int32, ) *PurchaseLineRead`

NewPurchaseLineRead instantiates a new PurchaseLineRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseLineReadWithDefaults

`func NewPurchaseLineReadWithDefaults() *PurchaseLineRead`

NewPurchaseLineReadWithDefaults instantiates a new PurchaseLineRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PurchaseLineRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseLineRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseLineRead) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *PurchaseLineRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseLineRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseLineRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseLineRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchaseLineRead) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseLineRead) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseLineRead) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.

### HasPlannedDate

`func (o *PurchaseLineRead) HasPlannedDate() bool`

HasPlannedDate returns a boolean if a field has been set.

### GetProductId

`func (o *PurchaseLineRead) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PurchaseLineRead) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PurchaseLineRead) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetPurchaseId

`func (o *PurchaseLineRead) GetPurchaseId() int32`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *PurchaseLineRead) GetPurchaseIdOk() (*int32, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *PurchaseLineRead) SetPurchaseId(v int32)`

SetPurchaseId sets PurchaseId field to given value.


### GetQtyReceived

`func (o *PurchaseLineRead) GetQtyReceived() int32`

GetQtyReceived returns the QtyReceived field if non-nil, zero value otherwise.

### GetQtyReceivedOk

`func (o *PurchaseLineRead) GetQtyReceivedOk() (*int32, bool)`

GetQtyReceivedOk returns a tuple with the QtyReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyReceived

`func (o *PurchaseLineRead) SetQtyReceived(v int32)`

SetQtyReceived sets QtyReceived field to given value.

### HasQtyReceived

`func (o *PurchaseLineRead) HasQtyReceived() bool`

HasQtyReceived returns a boolean if a field has been set.

### GetQuantity

`func (o *PurchaseLineRead) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseLineRead) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseLineRead) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


