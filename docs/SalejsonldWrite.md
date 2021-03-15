# SalejsonldWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**CarrierId** | Pointer to **int32** | The sale carrier ID | [optional] 
**ChannelId** | **int32** | The sale channel | 
**ChannelReference** | **string** | The sale channel reference | 
**InvoiceAddressId** | **int32** | The sale invoice address ID | 
**ShippingAddressId** | **int32** | The sale shipping address ID | 

## Methods

### NewSalejsonldWrite

`func NewSalejsonldWrite(channelId int32, channelReference string, invoiceAddressId int32, shippingAddressId int32, ) *SalejsonldWrite`

NewSalejsonldWrite instantiates a new SalejsonldWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalejsonldWriteWithDefaults

`func NewSalejsonldWriteWithDefaults() *SalejsonldWrite`

NewSalejsonldWriteWithDefaults instantiates a new SalejsonldWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SalejsonldWrite) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SalejsonldWrite) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SalejsonldWrite) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *SalejsonldWrite) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *SalejsonldWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalejsonldWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalejsonldWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SalejsonldWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SalejsonldWrite) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SalejsonldWrite) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SalejsonldWrite) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SalejsonldWrite) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCarrierId

`func (o *SalejsonldWrite) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SalejsonldWrite) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SalejsonldWrite) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SalejsonldWrite) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetChannelId

`func (o *SalejsonldWrite) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SalejsonldWrite) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SalejsonldWrite) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *SalejsonldWrite) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *SalejsonldWrite) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *SalejsonldWrite) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetInvoiceAddressId

`func (o *SalejsonldWrite) GetInvoiceAddressId() int32`

GetInvoiceAddressId returns the InvoiceAddressId field if non-nil, zero value otherwise.

### GetInvoiceAddressIdOk

`func (o *SalejsonldWrite) GetInvoiceAddressIdOk() (*int32, bool)`

GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressId

`func (o *SalejsonldWrite) SetInvoiceAddressId(v int32)`

SetInvoiceAddressId sets InvoiceAddressId field to given value.


### GetShippingAddressId

`func (o *SalejsonldWrite) GetShippingAddressId() int32`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *SalejsonldWrite) GetShippingAddressIdOk() (*int32, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *SalejsonldWrite) SetShippingAddressId(v int32)`

SetShippingAddressId sets ShippingAddressId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


