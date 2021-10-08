# SaleReadLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
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

### NewSaleReadLd

`func NewSaleReadLd(invoiceAddressId int32, shippingAddressId int32, ) *SaleReadLd`

NewSaleReadLd instantiates a new SaleReadLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleReadLdWithDefaults

`func NewSaleReadLdWithDefaults() *SaleReadLd`

NewSaleReadLdWithDefaults instantiates a new SaleReadLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *SaleReadLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *SaleReadLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *SaleReadLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *SaleReadLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *SaleReadLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *SaleReadLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *SaleReadLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *SaleReadLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *SaleReadLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *SaleReadLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *SaleReadLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *SaleReadLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetCarrierId

`func (o *SaleReadLd) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *SaleReadLd) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *SaleReadLd) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *SaleReadLd) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetId

`func (o *SaleReadLd) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaleReadLd) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaleReadLd) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SaleReadLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceAddressId

`func (o *SaleReadLd) GetInvoiceAddressId() int32`

GetInvoiceAddressId returns the InvoiceAddressId field if non-nil, zero value otherwise.

### GetInvoiceAddressIdOk

`func (o *SaleReadLd) GetInvoiceAddressIdOk() (*int32, bool)`

GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAddressId

`func (o *SaleReadLd) SetInvoiceAddressId(v int32)`

SetInvoiceAddressId sets InvoiceAddressId field to given value.


### GetLineIds

`func (o *SaleReadLd) GetLineIds() []int32`

GetLineIds returns the LineIds field if non-nil, zero value otherwise.

### GetLineIdsOk

`func (o *SaleReadLd) GetLineIdsOk() (*[]int32, bool)`

GetLineIdsOk returns a tuple with the LineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIds

`func (o *SaleReadLd) SetLineIds(v []int32)`

SetLineIds sets LineIds field to given value.

### HasLineIds

`func (o *SaleReadLd) HasLineIds() bool`

HasLineIds returns a boolean if a field has been set.

### GetName

`func (o *SaleReadLd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaleReadLd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaleReadLd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaleReadLd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShipmentIds

`func (o *SaleReadLd) GetShipmentIds() []int32`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SaleReadLd) GetShipmentIdsOk() (*[]int32, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SaleReadLd) SetShipmentIds(v []int32)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SaleReadLd) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetShippingAddressId

`func (o *SaleReadLd) GetShippingAddressId() int32`

GetShippingAddressId returns the ShippingAddressId field if non-nil, zero value otherwise.

### GetShippingAddressIdOk

`func (o *SaleReadLd) GetShippingAddressIdOk() (*int32, bool)`

GetShippingAddressIdOk returns a tuple with the ShippingAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddressId

`func (o *SaleReadLd) SetShippingAddressId(v int32)`

SetShippingAddressId sets ShippingAddressId field to given value.


### GetStatus

`func (o *SaleReadLd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SaleReadLd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SaleReadLd) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SaleReadLd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SaleReadLd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SaleReadLd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SaleReadLd) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SaleReadLd) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


