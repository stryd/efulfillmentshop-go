# PurchaseLinejsonldWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Description** | **string** | The purchase line description | 
**PlannedDate** | Pointer to **time.Time** | The purchase line planned delivery date | [optional] 
**ProductId** | **int32** | The product ID | 
**PurchaseId** | **int32** | The purchase ID | 
**Quantity** | **int32** | The purchase line quantity | 

## Methods

### NewPurchaseLinejsonldWrite

`func NewPurchaseLinejsonldWrite(description string, productId int32, purchaseId int32, quantity int32, ) *PurchaseLinejsonldWrite`

NewPurchaseLinejsonldWrite instantiates a new PurchaseLinejsonldWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseLinejsonldWriteWithDefaults

`func NewPurchaseLinejsonldWriteWithDefaults() *PurchaseLinejsonldWrite`

NewPurchaseLinejsonldWriteWithDefaults instantiates a new PurchaseLinejsonldWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PurchaseLinejsonldWrite) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PurchaseLinejsonldWrite) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PurchaseLinejsonldWrite) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *PurchaseLinejsonldWrite) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *PurchaseLinejsonldWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseLinejsonldWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseLinejsonldWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseLinejsonldWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PurchaseLinejsonldWrite) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseLinejsonldWrite) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseLinejsonldWrite) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PurchaseLinejsonldWrite) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *PurchaseLinejsonldWrite) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseLinejsonldWrite) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseLinejsonldWrite) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPlannedDate

`func (o *PurchaseLinejsonldWrite) GetPlannedDate() time.Time`

GetPlannedDate returns the PlannedDate field if non-nil, zero value otherwise.

### GetPlannedDateOk

`func (o *PurchaseLinejsonldWrite) GetPlannedDateOk() (*time.Time, bool)`

GetPlannedDateOk returns a tuple with the PlannedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDate

`func (o *PurchaseLinejsonldWrite) SetPlannedDate(v time.Time)`

SetPlannedDate sets PlannedDate field to given value.

### HasPlannedDate

`func (o *PurchaseLinejsonldWrite) HasPlannedDate() bool`

HasPlannedDate returns a boolean if a field has been set.

### GetProductId

`func (o *PurchaseLinejsonldWrite) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *PurchaseLinejsonldWrite) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *PurchaseLinejsonldWrite) SetProductId(v int32)`

SetProductId sets ProductId field to given value.


### GetPurchaseId

`func (o *PurchaseLinejsonldWrite) GetPurchaseId() int32`

GetPurchaseId returns the PurchaseId field if non-nil, zero value otherwise.

### GetPurchaseIdOk

`func (o *PurchaseLinejsonldWrite) GetPurchaseIdOk() (*int32, bool)`

GetPurchaseIdOk returns a tuple with the PurchaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseId

`func (o *PurchaseLinejsonldWrite) SetPurchaseId(v int32)`

SetPurchaseId sets PurchaseId field to given value.


### GetQuantity

`func (o *PurchaseLinejsonldWrite) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseLinejsonldWrite) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseLinejsonldWrite) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


