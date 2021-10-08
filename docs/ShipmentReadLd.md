# ShipmentReadLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**CarrierId** | Pointer to **int32** | The shipment carrier ID | [optional] 
**Id** | Pointer to **int32** | The shipment ID | [optional] 
**SaleId** | **int32** | The sale ID | 
**ShippedAt** | **time.Time** | The shipment date | 
**TrackingCodes** | Pointer to **[]string** | The shipment tracking codes | [optional] 

## Methods

### NewShipmentReadLd

`func NewShipmentReadLd(saleId int32, shippedAt time.Time, ) *ShipmentReadLd`

NewShipmentReadLd instantiates a new ShipmentReadLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentReadLdWithDefaults

`func NewShipmentReadLdWithDefaults() *ShipmentReadLd`

NewShipmentReadLdWithDefaults instantiates a new ShipmentReadLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *ShipmentReadLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *ShipmentReadLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *ShipmentReadLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *ShipmentReadLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *ShipmentReadLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *ShipmentReadLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *ShipmentReadLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *ShipmentReadLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *ShipmentReadLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *ShipmentReadLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *ShipmentReadLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *ShipmentReadLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetCarrierId

`func (o *ShipmentReadLd) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *ShipmentReadLd) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *ShipmentReadLd) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *ShipmentReadLd) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetId

`func (o *ShipmentReadLd) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentReadLd) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentReadLd) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentReadLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSaleId

`func (o *ShipmentReadLd) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *ShipmentReadLd) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *ShipmentReadLd) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.


### GetShippedAt

`func (o *ShipmentReadLd) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *ShipmentReadLd) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *ShipmentReadLd) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.


### GetTrackingCodes

`func (o *ShipmentReadLd) GetTrackingCodes() []string`

GetTrackingCodes returns the TrackingCodes field if non-nil, zero value otherwise.

### GetTrackingCodesOk

`func (o *ShipmentReadLd) GetTrackingCodesOk() (*[]string, bool)`

GetTrackingCodesOk returns a tuple with the TrackingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCodes

`func (o *ShipmentReadLd) SetTrackingCodes(v []string)`

SetTrackingCodes sets TrackingCodes field to given value.

### HasTrackingCodes

`func (o *ShipmentReadLd) HasTrackingCodes() bool`

HasTrackingCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


