# PurchaseLinejsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The purchase line description | 
**Id** | Pointer to **int32** | The purchase line ID | [optional] [readonly] 
**PlannedDate** | Pointer to **time.Time** | The purchase line planned delivery date | [optional] 
**ProductId** | **int32** | The product ID | 
**PurchaseId** | **int32** | The purchase ID | 
**QtyReceived** | Pointer to **int32** | The purchase line quantity received | [optional] [readonly] 
**Quantity** | **int32** | The purchase line quantity | 

## Methods

### NewPurchaseLinejsonldRead

`func NewPurchaseLinejsonldRead(description string, productId int32, purchaseId int32, quantity int32, ) *PurchaseLinejsonldRead`

NewPurchaseLinejsonldRead instantiates a new PurchaseLinejsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseLinejsonldReadWithDefaults

`func NewPurchaseLinejsonldReadWithDefaults() *PurchaseLinejsonldRead`

NewPurchaseLinejsonldReadWithDefaults instantiates a new PurchaseLinejsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PurchaseLinejsonldRead) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PurchaseLinejsonldRead) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PurchaseLinejsonldRead) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PurchaseLinejsonldRead) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *PurchaseLinejsonldRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseLinejsonldRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseLinejsonldRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseLinejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PurchaseLinejsonldRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseLinejsonldRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseLinejsonldRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PurchaseLinejsonldRead) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *PurchaseLinejsonldRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseLinejsonldRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseLinejsonldRead) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *PurchaseLinejsonldRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseLinejsonldRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseLinejsonldRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseLinejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchaseLinejsonldRead) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseLinejsonldRead) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseLinejsonldRead) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.

### HasPlannedDate

`func (o *PurchaseLinejsonldRead) HasPlannedDate() bool`

HasPlannedDate returns a boolean if a field has been set.

### GetProductId

`func (o *PurchaseLinejsonldRead) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PurchaseLinejsonldRead) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PurchaseLinejsonldRead) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetPurchaseId

`func (o *PurchaseLinejsonldRead) GetPurchaseId() int32`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *PurchaseLinejsonldRead) GetPurchaseIdOk() (*int32, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *PurchaseLinejsonldRead) SetPurchaseId(v int32)`

SetPurchaseId sets PurchaseId field to given value.


### GetQtyReceived

`func (o *PurchaseLinejsonldRead) GetQtyReceived() int32`

GetQtyReceived returns the QtyReceived field if non-nil, zero value otherwise.

### GetQtyReceivedOk

`func (o *PurchaseLinejsonldRead) GetQtyReceivedOk() (*int32, bool)`

GetQtyReceivedOk returns a tuple with the QtyReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyReceived

`func (o *PurchaseLinejsonldRead) SetQtyReceived(v int32)`

SetQtyReceived sets QtyReceived field to given value.

### HasQtyReceived

`func (o *PurchaseLinejsonldRead) HasQtyReceived() bool`

HasQtyReceived returns a boolean if a field has been set.

### GetQuantity

`func (o *PurchaseLinejsonldRead) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseLinejsonldRead) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseLinejsonldRead) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


