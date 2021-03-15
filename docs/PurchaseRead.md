# PurchaseRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **int32** | The purchase channel | 
**ChannelReference** | **string** | The purchase channel reference | 
**Id** | Pointer to **int32** | The purchase ID | [optional] [readonly] 
**Name** | Pointer to **string** | The purchase name | [optional] [readonly] 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchaseRead

`func NewPurchaseRead(channelId int32, channelReference string, plannedDate time.Time, supplierId int32, ) *PurchaseRead`

NewPurchaseRead instantiates a new PurchaseRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseReadWithDefaults

`func NewPurchaseReadWithDefaults() *PurchaseRead`

NewPurchaseReadWithDefaults instantiates a new PurchaseRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *PurchaseRead) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PurchaseRead) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PurchaseRead) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *PurchaseRead) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *PurchaseRead) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *PurchaseRead) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetId

`func (o *PurchaseRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PurchaseRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PurchaseRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PurchaseRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PurchaseRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchaseRead) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseRead) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseRead) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.


### GetSupplierId

`func (o *PurchaseRead) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *PurchaseRead) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *PurchaseRead) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


