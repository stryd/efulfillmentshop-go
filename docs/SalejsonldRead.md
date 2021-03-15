# SalejsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**CarrierId** | Pointer to **int32** | The sale carrier ID | [optional] 
**ChannelId** | **int32** | The sale channel | 
**ChannelReference** | **string** | The sale channel reference | 
**Id** | Pointer to **int32** | The sale ID | [optional] [readonly] 
**InvoiceAddressId** | **int32** | The sale invoice address ID | 
**LineIds** | Pointer to **[]int32** | The sale line IDs | [optional] [readonly] 
**Name** | Pointer to **string** | The sale name | [optional] [readonly] 
**ShipmentIds** | Pointer to **[]int32** | The sale shipment IDs | [optional] [readonly] 
**ShippingAddressId** | **int32** | The sale shipping address ID | 
**Status** | Pointer to **string** | The sale warehouse status | [optional] [readonly] 

## Methods

### NewSalejsonldRead

`func NewSalejsonldRead(channelId int32, channelReference string, invoiceAddressId int32, shippingAddressId int32, ) *SalejsonldRead`

NewSalejsonldRead instantiates a new SalejsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalejsonldReadWithDefaults

`func NewSalejsonldReadWithDefaults() *SalejsonldRead`

NewSalejsonldReadWithDefaults instantiates a new SalejsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SalejsonldRead) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SalejsonldRead) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SalejsonldRead) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *SalejsonldRead) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *SalejsonldRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalejsonldRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalejsonldRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SalejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SalejsonldRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SalejsonldRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SalejsonldRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SalejsonldRead) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCarrierId

`func (o *SalejsonldRead) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SalejsonldRead) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SalejsonldRead) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SalejsonldRead) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetChannelId

`func (o *SalejsonldRead) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SalejsonldRead) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SalejsonldRead) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *SalejsonldRead) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *SalejsonldRead) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *SalejsonldRead) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetId

`func (o *SalejsonldRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalejsonldRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalejsonldRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalejsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceAddressId

`func (o *SalejsonldRead) GetInvoiceAddressId() int32`

GetInvoiceAddressId returns the InvoiceAddressId field if non-nil, zero value otherwise.

### GetInvoiceAddressIdOk

`func (o *SalejsonldRead) GetInvoiceAddressIdOk() (*int32, bool)`

GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressId

`func (o *SalejsonldRead) SetInvoiceAddressId(v int32)`

SetInvoiceAddressId sets InvoiceAddressId field to given value.


### GetLineIds

`func (o *SalejsonldRead) GetLineIds() []int32`

GetLineIds returns the LineIds field if non-nil, zero value otherwise.

### GetLineIdsOk

`func (o *SalejsonldRead) GetLineIdsOk() (*[]int32, bool)`

GetLineIdsOk returns a tuple with the LineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIds

`func (o *SalejsonldRead) SetLineIds(v []int32)`

SetLineIds sets LineIds field to given value.

### HasLineIds

`func (o *SalejsonldRead) HasLineIds() bool`

HasLineIds returns a boolean if a field has been set.

### GetName

`func (o *SalejsonldRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SalejsonldRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SalejsonldRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SalejsonldRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SalejsonldRead) GetShipmentIds() []int32`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SalejsonldRead) GetShipmentIdsOk() (*[]int32, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SalejsonldRead) SetShipmentIds(v []int32)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SalejsonldRead) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *SalejsonldRead) GetShippingAddressId() int32`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *SalejsonldRead) GetShippingAddressIdOk() (*int32, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *SalejsonldRead) SetShippingAddressId(v int32)`

SetShippingAddressId sets ShippingAddressId field to given value.


### GetStatus

`func (o *SalejsonldRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SalejsonldRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SalejsonldRead) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SalejsonldRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


