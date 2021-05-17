# SaleWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierId** | Pointer to **int32** | The sale carrier ID | [optional] 
**ChannelReference** | Pointer to **string** | The sale channel reference (DEPRECATED) | [optional] 
**InvoiceAddressId** | **int32** | The sale invoice address ID | 
**Reference** | Pointer to **string** | Your sale reference (This could be your sale ID) | [optional] 
**ShippingAddressId** | **int32** | The sale shipping address ID | 

## Methods

### NewSaleWrite

`func NewSaleWrite(invoiceAddressId int32, shippingAddressId int32, ) *SaleWrite`

NewSaleWrite instantiates a new SaleWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleWriteWithDefaults

`func NewSaleWriteWithDefaults() *SaleWrite`

NewSaleWriteWithDefaults instantiates a new SaleWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierId

`func (o *SaleWrite) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SaleWrite) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SaleWrite) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SaleWrite) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetChannelReference

`func (o *SaleWrite) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *SaleWrite) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *SaleWrite) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.

### HasChannelReference

`func (o *SaleWrite) HasChannelReference() bool`

HasChannelReference returns a boolean if a field has been set.

### GetInvoiceAddressId

`func (o *SaleWrite) GetInvoiceAddressId() int32`

GetInvoiceAddressId returns the InvoiceAddressId field if non-nil, zero value otherwise.

### GetInvoiceAddressIdOk

`func (o *SaleWrite) GetInvoiceAddressIdOk() (*int32, bool)`

GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressId

`func (o *SaleWrite) SetInvoiceAddressId(v int32)`

SetInvoiceAddressId sets InvoiceAddressId field to given value.


### GetReference

`func (o *SaleWrite) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SaleWrite) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SaleWrite) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SaleWrite) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *SaleWrite) GetShippingAddressId() int32`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *SaleWrite) GetShippingAddressIdOk() (*int32, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *SaleWrite) SetShippingAddressId(v int32)`

SetShippingAddressId sets ShippingAddressId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


