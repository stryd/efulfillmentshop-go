# ShipmentjsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**CarrierId** | Pointer to **int32** | The shipment carrier ID | [optional] [readonly] 
**Id** | Pointer to **int32** | The shipment ID | [optional] [readonly] 
**SaleId** | Pointer to **int32** | The sale ID | [optional] [readonly] 
**ShippedAt** | Pointer to **time.Time** | The shipment date | [optional] [readonly] 
**TrackingCodes** | Pointer to **[]string** | The shipment tracking codes | [optional] [readonly] 

## Methods

### NewShipmentjsonldRead

`func NewShipmentjsonldRead() *ShipmentjsonldRead`

NewShipmentjsonldRead instantiates a new ShipmentjsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentjsonldReadWithDefaults

`func NewShipmentjsonldReadWithDefaults() *ShipmentjsonldRead`

NewShipmentjsonldReadWithDefaults instantiates a new ShipmentjsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ShipmentjsonldRead) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ShipmentjsonldRead) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ShipmentjsonldRead) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ShipmentjsonldRead) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *ShipmentjsonldRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentjsonldRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentjsonldRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentjsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ShipmentjsonldRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipmentjsonldRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipmentjsonldRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShipmentjsonldRead) HasType() bool`

HasType returns a boolean if a field has been set.

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

### HasSaleId

`func (o *ShipmentjsonldRead) HasSaleId() bool`

HasSaleId returns a boolean if a field has been set.

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

### HasShippedAt

`func (o *ShipmentjsonldRead) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

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


