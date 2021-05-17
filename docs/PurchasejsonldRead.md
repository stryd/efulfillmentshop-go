# PurchasejsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonContext** | Pointer to **string** |  | [optional] [readonly] 
**JsonId** | Pointer to **string** |  | [optional] [readonly] 
**JsonType** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **int32** | The purchase ID | [optional] 
**Name** | Pointer to **string** | The purchase name | [optional] 
**PlannedDate** | **time.Time** | The purchase planned delivery date | 
**SupplierId** | **int32** | The purchase supplier ID | 

## Methods

### NewPurchasejsonldRead

`func NewPurchasejsonldRead(plannedDate time.Time, supplierId int32, ) *PurchasejsonldRead`

NewPurchasejsonldRead instantiates a new PurchasejsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchasejsonldReadWithDefaults

`func NewPurchasejsonldReadWithDefaults() *PurchasejsonldRead`

NewPurchasejsonldReadWithDefaults instantiates a new PurchasejsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonContext

`func (o *PurchasejsonldRead) GetJsonContext() string`

GetJsonContext returns the JsonContext field if non-nil, zero value otherwise.

### GetJsonContextOk

`func (o *PurchasejsonldRead) GetJsonContextOk() (*string, bool)`

GetJsonContextOk returns a tuple with the JsonContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonContext

`func (o *PurchasejsonldRead) SetJsonContext(v string)`

SetJsonContext sets JsonContext field to given value.

### HasJsonContext

`func (o *PurchasejsonldRead) HasJsonContext() bool`

HasJsonContext returns a boolean if a field has been set.

### GetJsonId

`func (o *PurchasejsonldRead) GetJsonId() string`

GetJsonId returns the JsonId field if non-nil, zero value otherwise.

### GetJsonIdOk

`func (o *PurchasejsonldRead) GetJsonIdOk() (*string, bool)`

GetJsonIdOk returns a tuple with the JsonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonId

`func (o *PurchasejsonldRead) SetJsonId(v string)`

SetJsonId sets JsonId field to given value.

### HasJsonId

`func (o *PurchasejsonldRead) HasJsonId() bool`

HasJsonId returns a boolean if a field has been set.

### GetJsonType

`func (o *PurchasejsonldRead) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *PurchasejsonldRead) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *PurchasejsonldRead) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.

### HasJsonType

`func (o *PurchasejsonldRead) HasJsonType() bool`

HasJsonType returns a boolean if a field has been set.

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


