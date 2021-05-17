# PurchaseWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelReference** | Pointer to **string** | The purchase channel reference (DEPRECATED) | [optional] 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**Reference** | Pointer to **string** | Your purchase reference (This could be your purchase ID) | [optional] 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchaseWrite

`func NewPurchaseWrite(plannedDate time.Time, supplierId int32, ) *PurchaseWrite`

NewPurchaseWrite instantiates a new PurchaseWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWriteWithDefaults

`func NewPurchaseWriteWithDefaults() *PurchaseWrite`

NewPurchaseWriteWithDefaults instantiates a new PurchaseWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelReference

`func (o *PurchaseWrite) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *PurchaseWrite) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *PurchaseWrite) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.

### HasChannelReference

`func (o *PurchaseWrite) HasChannelReference() bool`

HasChannelReference returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchaseWrite) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseWrite) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseWrite) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.


### GetReference

`func (o *PurchaseWrite) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PurchaseWrite) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PurchaseWrite) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PurchaseWrite) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSupplierId

`func (o *PurchaseWrite) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *PurchaseWrite) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *PurchaseWrite) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


