# PurchasejsonldWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonContext** | Pointer to **string** |  | [optional] [readonly] 
**JsonId** | Pointer to **string** |  | [optional] [readonly] 
**JsonType** | Pointer to **string** |  | [optional] [readonly] 
**ChannelId** | **int32** | The purchase channel | 
**ChannelReference** | **string** | The purchase channel reference | 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchasejsonldWrite

`func NewPurchasejsonldWrite(channelId int32, channelReference string, plannedDate time.Time, supplierId int32, ) *PurchasejsonldWrite`

NewPurchasejsonldWrite instantiates a new PurchasejsonldWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchasejsonldWriteWithDefaults

`func NewPurchasejsonldWriteWithDefaults() *PurchasejsonldWrite`

NewPurchasejsonldWriteWithDefaults instantiates a new PurchasejsonldWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonContext

`func (o *PurchasejsonldWrite) GetJsonContext() string`

GetJsonContext returns the JsonContext field if non-nil, zero value otherwise.

### GetJsonContextOk

`func (o *PurchasejsonldWrite) GetJsonContextOk() (*string, bool)`

GetJsonContextOk returns a tuple with the JsonContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonContext

`func (o *PurchasejsonldWrite) SetJsonContext(v string)`

SetJsonContext sets JsonContext field to given value.

### HasJsonContext

`func (o *PurchasejsonldWrite) HasJsonContext() bool`

HasJsonContext returns a boolean if a field has been set.

### GetJsonId

`func (o *PurchasejsonldWrite) GetJsonId() string`

GetJsonId returns the JsonId field if non-nil, zero value otherwise.

### GetJsonIdOk

`func (o *PurchasejsonldWrite) GetJsonIdOk() (*string, bool)`

GetJsonIdOk returns a tuple with the JsonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonId

`func (o *PurchasejsonldWrite) SetJsonId(v string)`

SetJsonId sets JsonId field to given value.

### HasJsonId

`func (o *PurchasejsonldWrite) HasJsonId() bool`

HasJsonId returns a boolean if a field has been set.

### GetJsonType

`func (o *PurchasejsonldWrite) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *PurchasejsonldWrite) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *PurchasejsonldWrite) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.

### HasJsonType

`func (o *PurchasejsonldWrite) HasJsonType() bool`

HasJsonType returns a boolean if a field has been set.

### GetChannelId

`func (o *PurchasejsonldWrite) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PurchasejsonldWrite) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PurchasejsonldWrite) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *PurchasejsonldWrite) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *PurchasejsonldWrite) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *PurchasejsonldWrite) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetPlannedDate

`func (o *PurchasejsonldWrite) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchasejsonldWrite) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchasejsonldWrite) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.


### GetSupplierId

`func (o *PurchasejsonldWrite) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *PurchasejsonldWrite) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *PurchasejsonldWrite) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


