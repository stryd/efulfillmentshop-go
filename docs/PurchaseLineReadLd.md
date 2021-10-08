# PurchaseLineReadLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The purchase line description | 
**Id** | Pointer to **int32** | The purchase line ID | [optional] 
**PlannedDate** | Pointer to **time.Time** | The purchase line planned delivery date | [optional] 
**ProductId** | **int32** | The product ID | 
**PurchaseId** | **int32** | The purchase ID | 
**QtyReceived** | Pointer to **int32** | The purchase line quantity received | [optional] 
**Quantity** | **int32** | The purchase line quantity | 

## Methods

### NewPurchaseLineReadLd

`func NewPurchaseLineReadLd(description string, productId int32, purchaseId int32, quantity int32, ) *PurchaseLineReadLd`

NewPurchaseLineReadLd instantiates a new PurchaseLineReadLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseLineReadLdWithDefaults

`func NewPurchaseLineReadLdWithDefaults() *PurchaseLineReadLd`

NewPurchaseLineReadLdWithDefaults instantiates a new PurchaseLineReadLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *PurchaseLineReadLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *PurchaseLineReadLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *PurchaseLineReadLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *PurchaseLineReadLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *PurchaseLineReadLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *PurchaseLineReadLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *PurchaseLineReadLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *PurchaseLineReadLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *PurchaseLineReadLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *PurchaseLineReadLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *PurchaseLineReadLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *PurchaseLineReadLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetDescription

`func (o *PurchaseLineReadLd) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseLineReadLd) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseLineReadLd) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *PurchaseLineReadLd) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseLineReadLd) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseLineReadLd) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseLineReadLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchaseLineReadLd) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseLineReadLd) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseLineReadLd) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.

### HasPlannedDate

`func (o *PurchaseLineReadLd) HasPlannedDate() bool`

HasPlannedDate returns a boolean if a field has been set.

### GetProductId

`func (o *PurchaseLineReadLd) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PurchaseLineReadLd) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PurchaseLineReadLd) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetPurchaseId

`func (o *PurchaseLineReadLd) GetPurchaseId() int32`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *PurchaseLineReadLd) GetPurchaseIdOk() (*int32, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *PurchaseLineReadLd) SetPurchaseId(v int32)`

SetPurchaseId sets PurchaseId field to given value.


### GetQtyReceived

`func (o *PurchaseLineReadLd) GetQtyReceived() int32`

GetQtyReceived returns the QtyReceived field if non-nil, zero value otherwise.

### GetQtyReceivedOk

`func (o *PurchaseLineReadLd) GetQtyReceivedOk() (*int32, bool)`

GetQtyReceivedOk returns a tuple with the QtyReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyReceived

`func (o *PurchaseLineReadLd) SetQtyReceived(v int32)`

SetQtyReceived sets QtyReceived field to given value.

### HasQtyReceived

`func (o *PurchaseLineReadLd) HasQtyReceived() bool`

HasQtyReceived returns a boolean if a field has been set.

### GetQuantity

`func (o *PurchaseLineReadLd) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseLineReadLd) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseLineReadLd) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


