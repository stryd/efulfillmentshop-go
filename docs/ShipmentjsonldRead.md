# ShipmentjsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonContext** | Pointer to **string** |  | [optional] [readonly] 
**JsonId** | Pointer to **string** |  | [optional] [readonly] 
**JsonType** | Pointer to **string** |  | [optional] [readonly] 
**CarrierId** | Pointer to **int32** | The shipment carrier ID | [optional] 
**Id** | Pointer to **int32** | The shipment ID | [optional] 
**SaleId** | **int32** | The sale ID | 
**ShippedAt** | **time.Time** | The shipment date | 
**TrackingCodes** | Pointer to **[]string** | The shipment tracking codes | [optional] 

## Methods

### NewShipmentjsonldRead

`func NewShipmentjsonldRead(saleId int32, shippedAt time.Time, ) *ShipmentjsonldRead`

NewShipmentjsonldRead instantiates a new ShipmentjsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentjsonldReadWithDefaults

`func NewShipmentjsonldReadWithDefaults() *ShipmentjsonldRead`

NewShipmentjsonldReadWithDefaults instantiates a new ShipmentjsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonContext

`func (o *ShipmentjsonldRead) GetJsonContext() string`

GetJsonContext returns the JsonContext field if non-nil, zero value otherwise.

### GetJsonContextOk

`func (o *ShipmentjsonldRead) GetJsonContextOk() (*string, bool)`

GetJsonContextOk returns a tuple with the JsonContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonContext

`func (o *ShipmentjsonldRead) SetJsonContext(v string)`

SetJsonContext sets JsonContext field to given value.

### HasJsonContext

`func (o *ShipmentjsonldRead) HasJsonContext() bool`

HasJsonContext returns a boolean if a field has been set.

### GetJsonId

`func (o *ShipmentjsonldRead) GetJsonId() string`

GetJsonId returns the JsonId field if non-nil, zero value otherwise.

### GetJsonIdOk

`func (o *ShipmentjsonldRead) GetJsonIdOk() (*string, bool)`

GetJsonIdOk returns a tuple with the JsonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonId

`func (o *ShipmentjsonldRead) SetJsonId(v string)`

SetJsonId sets JsonId field to given value.

### HasJsonId

`func (o *ShipmentjsonldRead) HasJsonId() bool`

HasJsonId returns a boolean if a field has been set.

### GetJsonType

`func (o *ShipmentjsonldRead) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *ShipmentjsonldRead) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *ShipmentjsonldRead) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.

### HasJsonType

`func (o *ShipmentjsonldRead) HasJsonType() bool`

HasJsonType returns a boolean if a field has been set.

### GetCarrierId

`func (o *ShipmentjsonldRead) GetCarrierId() int32`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *ShipmentjsonldRead) GetCarrierIdOk() (*int32, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *ShipmentjsonldRead) SetCarrierId(v int32)`

SetCarrierId sets CarrierId field to given value.

### HasCarrierId

`func (o *ShipmentjsonldRead) HasCarrierId() bool`

HasCarrierId returns a boolean if a field has been set.

### GetId

`func (o *ShipmentjsonldRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentjsonldRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentjsonldRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentjsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSaleId

`func (o *ShipmentjsonldRead) GetSaleId() int32`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *ShipmentjsonldRead) GetSaleIdOk() (*int32, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *ShipmentjsonldRead) SetSaleId(v int32)`

SetSaleId sets SaleId field to given value.


### GetShippedAt

`func (o *ShipmentjsonldRead) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *ShipmentjsonldRead) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *ShipmentjsonldRead) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.


### GetTrackingCodes

`func (o *ShipmentjsonldRead) GetTrackingCodes() []string`

GetTrackingCodes returns the TrackingCodes field if non-nil, zero value otherwise.

### GetTrackingCodesOk

`func (o *ShipmentjsonldRead) GetTrackingCodesOk() (*[]string, bool)`

GetTrackingCodesOk returns a tuple with the TrackingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingCodes

`func (o *ShipmentjsonldRead) SetTrackingCodes(v []string)`

SetTrackingCodes sets TrackingCodes field to given value.

### HasTrackingCodes

`func (o *ShipmentjsonldRead) HasTrackingCodes() bool`

HasTrackingCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


