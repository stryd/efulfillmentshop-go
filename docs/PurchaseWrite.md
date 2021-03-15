# PurchaseWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **int32** | The purchase channel | 
**ChannelReference** | **string** | The purchase channel reference | 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchaseWrite

`func NewPurchaseWrite(channelId int32, channelReference string, plannedDate time.Time, supplierId int32, ) *PurchaseWrite`

NewPurchaseWrite instantiates a new PurchaseWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWriteWithDefaults

`func NewPurchaseWriteWithDefaults() *PurchaseWrite`

NewPurchaseWriteWithDefaults instantiates a new PurchaseWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *PurchaseWrite) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PurchaseWrite) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PurchaseWrite) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


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


