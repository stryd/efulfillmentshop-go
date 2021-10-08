# PurchaseWriteLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**ChannelReference** | Pointer to **string** | The purchase channel reference (DEPRECATED) | [optional] 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**Reference** | Pointer to **string** | Your purchase reference (This could be your purchase ID) | [optional] 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchaseWriteLd

`func NewPurchaseWriteLd(plannedDate time.Time, supplierId int32, ) *PurchaseWriteLd`

NewPurchaseWriteLd instantiates a new PurchaseWriteLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWriteLdWithDefaults

`func NewPurchaseWriteLdWithDefaults() *PurchaseWriteLd`

NewPurchaseWriteLdWithDefaults instantiates a new PurchaseWriteLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *PurchaseWriteLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *PurchaseWriteLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *PurchaseWriteLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *PurchaseWriteLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *PurchaseWriteLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *PurchaseWriteLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *PurchaseWriteLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *PurchaseWriteLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *PurchaseWriteLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *PurchaseWriteLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *PurchaseWriteLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *PurchaseWriteLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetChannelReference

`func (o *PurchaseWriteLd) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *PurchaseWriteLd) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *PurchaseWriteLd) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.

### HasChannelReference

`func (o *PurchaseWriteLd) HasChannelReference() bool`

HasChannelReference returns a boolean if a field has been set.

### GetPlannedDate

`func (o *PurchaseWriteLd) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseWriteLd) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseWriteLd) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.


### GetReference

`func (o *PurchaseWriteLd) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PurchaseWriteLd) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PurchaseWriteLd) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PurchaseWriteLd) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSupplierId

`func (o *PurchaseWriteLd) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *PurchaseWriteLd) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *PurchaseWriteLd) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


