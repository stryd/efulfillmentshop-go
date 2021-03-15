# ShipmentRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierId** | Pointer to **int32** | The shipment carrier ID | [optional] [readonly] 
**Id** | Pointer to **int32** | The shipment ID | [optional] [readonly] 
**SaleId** | Pointer to **int32** | The sale ID | [optional] [readonly] 
**ShippedAt** | Pointer to **time.Time** | The shipment date | [optional] [readonly] 
**TrackingCodes** | Pointer to **[]string** | The shipment tracking codes | [optional] [readonly] 

## Methods

### NewShipmentRead

`func NewShipmentRead() *ShipmentRead`

NewShipmentRead instantiates a new ShipmentRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentReadWithDefaults

`func NewShipmentReadWithDefaults() *ShipmentRead`

NewShipmentReadWithDefaults instantiates a new ShipmentRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierId

`func (o *ShipmentRead) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *ShipmentRead) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *ShipmentRead) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *ShipmentRead) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetId

`func (o *ShipmentRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSaleId

`func (o *ShipmentRead) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *ShipmentRead) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *ShipmentRead) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.

### HasSaleId

`func (o *ShipmentRead) HasSaleId() bool`

HasSaleId returns a boolean if a field has been set.

### GetShippedAt

`func (o *ShipmentRead) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *ShipmentRead) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *ShipmentRead) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *ShipmentRead) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### GetTrackingCodes

`func (o *ShipmentRead) GetTrackingCodes() []string`

GetTrackingCodes returns the TrackingCodes field if non-nil, zero value otherwise.

### GetTrackingCodesOk

`func (o *ShipmentRead) GetTrackingCodesOk() (*[]string, bool)`

GetTrackingCodesOk returns a tuple with the TrackingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCodes

`func (o *ShipmentRead) SetTrackingCodes(v []string)`

SetTrackingCodes sets TrackingCodes field to given value.

### HasTrackingCodes

`func (o *ShipmentRead) HasTrackingCodes() bool`

HasTrackingCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


