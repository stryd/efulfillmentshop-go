# PurchasejsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**ChannelId** | **int32** | The purchase channel | 
**ChannelReference** | **string** | The purchase channel reference | 
**Id** | Pointer to **int32** | The purchase ID | [optional] [readonly] 
**Name** | Pointer to **string** | The purchase name | [optional] [readonly] 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchasejsonldRead

`func NewPurchasejsonldRead(channelId int32, channelReference string, plannedDate time.Time, supplierId int32, ) *PurchasejsonldRead`

NewPurchasejsonldRead instantiates a new PurchasejsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchasejsonldReadWithDefaults

`func NewPurchasejsonldReadWithDefaults() *PurchasejsonldRead`

NewPurchasejsonldReadWithDefaults instantiates a new PurchasejsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PurchasejsonldRead) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PurchasejsonldRead) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PurchasejsonldRead) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PurchasejsonldRead) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *PurchasejsonldRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchasejsonldRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchasejsonldRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PurchasejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PurchasejsonldRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchasejsonldRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchasejsonldRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PurchasejsonldRead) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChannelId

`func (o *PurchasejsonldRead) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PurchasejsonldRead) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PurchasejsonldRead) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *PurchasejsonldRead) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *PurchasejsonldRead) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *PurchasejsonldRead) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetId

`func (o *PurchasejsonldRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchasejsonldRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchasejsonldRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchasejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PurchasejsonldRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PurchasejsonldRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PurchasejsonldRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PurchasejsonldRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchasejsonldRead) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchasejsonldRead) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchasejsonldRead) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.


### GetSupplierId

`func (o *PurchasejsonldRead) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *PurchasejsonldRead) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *PurchasejsonldRead) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


