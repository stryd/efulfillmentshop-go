# SaleRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierId** | Pointer to **int32** | The sale carrier ID | [optional] 
**Id** | Pointer to **int32** | The sale ID | [optional] 
**InvoiceAddressId** | **int32** | The sale invoice address ID | 
**LineIds** | Pointer to **[]int32** | The sale line IDs | [optional] 
**Name** | Pointer to **string** | The sale name | [optional] 
**ShipmentIds** | Pointer to **[]int32** | The sale shipment IDs | [optional] 
**ShippingAddressId** | **int32** | The sale shipping address ID | 
**Status** | Pointer to **string** | The sale warehouse status | [optional] 
**Type** | Pointer to **string** | The sale type | [optional] 

## Methods

### NewSaleRead

`func NewSaleRead(invoiceAddressId int32, shippingAddressId int32, ) *SaleRead`

NewSaleRead instantiates a new SaleRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleReadWithDefaults

`func NewSaleReadWithDefaults() *SaleRead`

NewSaleReadWithDefaults instantiates a new SaleRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierId

`func (o *SaleRead) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SaleRead) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SaleRead) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SaleRead) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetId

`func (o *SaleRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SaleRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceAddressId

`func (o *SaleRead) GetInvoiceAddressId() int32`

GetInvoiceAddressId returns the InvoiceAddressId field if non-nil, zero value otherwise.

### GetInvoiceAddressIdOk

`func (o *SaleRead) GetInvoiceAddressIdOk() (*int32, bool)`

GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressId

`func (o *SaleRead) SetInvoiceAddressId(v int32)`

SetInvoiceAddressId sets InvoiceAddressId field to given value.


### GetLineIds

`func (o *SaleRead) GetLineIds() []int32`

GetLineIds returns the LineIds field if non-nil, zero value otherwise.

### GetLineIdsOk

`func (o *SaleRead) GetLineIdsOk() (*[]int32, bool)`

GetLineIdsOk returns a tuple with the LineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIds

`func (o *SaleRead) SetLineIds(v []int32)`

SetLineIds sets LineIds field to given value.

### HasLineIds

`func (o *SaleRead) HasLineIds() bool`

HasLineIds returns a boolean if a field has been set.

### GetName

`func (o *SaleRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaleRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaleRead) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaleRead) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SaleRead) GetShipmentIds() []int32`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SaleRead) GetShipmentIdsOk() (*[]int32, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SaleRead) SetShipmentIds(v []int32)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SaleRead) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *SaleRead) GetShippingAddressId() int32`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *SaleRead) GetShippingAddressIdOk() (*int32, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *SaleRead) SetShippingAddressId(v int32)`

SetShippingAddressId sets ShippingAddressId field to given value.


### GetStatus

`func (o *SaleRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaleRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaleRead) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SaleRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SaleRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaleRead) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


