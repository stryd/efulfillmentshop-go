# PurchasejsonldWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
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

### GetContext

`func (o *PurchasejsonldWrite) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PurchasejsonldWrite) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PurchasejsonldWrite) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PurchasejsonldWrite) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *PurchasejsonldWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchasejsonldWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchasejsonldWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PurchasejsonldWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PurchasejsonldWrite) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchasejsonldWrite) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchasejsonldWrite) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PurchasejsonldWrite) HasType() bool`

HasType returns a boolean if a field has been set.

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


