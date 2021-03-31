# ProductjsonldReadCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | **string** | The product barcode | 
**ChannelId** | **int32** | The product channel | 
**ChannelReference** | **string** | The product channel reference | 
**ChannelSku** | Pointer to **string** | The product channel SKU | [optional] 
**CostPrice** | Pointer to **float32** | The product cost price (Excluding taxes) | [optional] 
**DimensionX** | Pointer to **float32** | The product dimension X in cm | [optional] 
**DimensionY** | Pointer to **float32** | The product dimension Y in cm | [optional] 
**DimensionZ** | Pointer to **float32** | The product dimension Z in cm | [optional] 
**HsCode** | Pointer to **string** | The product HS code | [optional] 
**Id** | Pointer to **int32** | The product ID | [optional] 
**Name** | **string** | The product name | 
**QuantityAvailable** | Pointer to **int32** | The physical available product quantity | [optional] 
**QuantityIncoming** | Pointer to **int32** | The incoming product quantity | [optional] 
**QuantityOnHand** | Pointer to **int32** | The product quantity on hand | [optional] 
**QuantityOutgoing** | Pointer to **int32** | The outgoing product quantity | [optional] 
**QuantitySold** | Pointer to **int32** | The sold product quantity | [optional] 
**SellingPrice** | Pointer to **float32** | The product selling price (Excluding taxes) | [optional] 
**Sku** | Pointer to **string** | The product warehouse SKU | [optional] 
**Volume** | Pointer to **float32** | The product volume in L (Calculated using dimensions) | [optional] 
**Weight** | Pointer to **float32** | The product weight in kg | [optional] 

## Methods

### NewProductjsonldReadCollection

`func NewProductjsonldReadCollection(barcode string, channelId int32, channelReference string, name string, ) *ProductjsonldReadCollection`

NewProductjsonldReadCollection instantiates a new ProductjsonldReadCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductjsonldReadCollectionWithDefaults

`func NewProductjsonldReadCollectionWithDefaults() *ProductjsonldReadCollection`

NewProductjsonldReadCollectionWithDefaults instantiates a new ProductjsonldReadCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ProductjsonldReadCollection) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductjsonldReadCollection) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductjsonldReadCollection) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetChannelId

`func (o *ProductjsonldReadCollection) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ProductjsonldReadCollection) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ProductjsonldReadCollection) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *ProductjsonldReadCollection) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ProductjsonldReadCollection) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ProductjsonldReadCollection) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetChannelSku

`func (o *ProductjsonldReadCollection) GetChannelSku() string`

GetChannelSku returns the ChannelSku field if non-nil, zero value otherwise.

### GetChannelSkuOk

`func (o *ProductjsonldReadCollection) GetChannelSkuOk() (*string, bool)`

GetChannelSkuOk returns a tuple with the ChannelSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSku

`func (o *ProductjsonldReadCollection) SetChannelSku(v string)`

SetChannelSku sets ChannelSku field to given value.

### HasChannelSku

`func (o *ProductjsonldReadCollection) HasChannelSku() bool`

HasChannelSku returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductjsonldReadCollection) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductjsonldReadCollection) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductjsonldReadCollection) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductjsonldReadCollection) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetDimensionX

`func (o *ProductjsonldReadCollection) GetDimensionX() float32`

GetDimensionX returns the DimensionX field if non-nil, zero value otherwise.

### GetDimensionXOk

`func (o *ProductjsonldReadCollection) GetDimensionXOk() (*float32, bool)`

GetDimensionXOk returns a tuple with the DimensionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionX

`func (o *ProductjsonldReadCollection) SetDimensionX(v float32)`

SetDimensionX sets DimensionX field to given value.

### HasDimensionX

`func (o *ProductjsonldReadCollection) HasDimensionX() bool`

HasDimensionX returns a boolean if a field has been set.

### GetDimensionY

`func (o *ProductjsonldReadCollection) GetDimensionY() float32`

GetDimensionY returns the DimensionY field if non-nil, zero value otherwise.

### GetDimensionYOk

`func (o *ProductjsonldReadCollection) GetDimensionYOk() (*float32, bool)`

GetDimensionYOk returns a tuple with the DimensionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionY

`func (o *ProductjsonldReadCollection) SetDimensionY(v float32)`

SetDimensionY sets DimensionY field to given value.

### HasDimensionY

`func (o *ProductjsonldReadCollection) HasDimensionY() bool`

HasDimensionY returns a boolean if a field has been set.

### GetDimensionZ

`func (o *ProductjsonldReadCollection) GetDimensionZ() float32`

GetDimensionZ returns the DimensionZ field if non-nil, zero value otherwise.

### GetDimensionZOk

`func (o *ProductjsonldReadCollection) GetDimensionZOk() (*float32, bool)`

GetDimensionZOk returns a tuple with the DimensionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionZ

`func (o *ProductjsonldReadCollection) SetDimensionZ(v float32)`

SetDimensionZ sets DimensionZ field to given value.

### HasDimensionZ

`func (o *ProductjsonldReadCollection) HasDimensionZ() bool`

HasDimensionZ returns a boolean if a field has been set.

### GetHsCode

`func (o *ProductjsonldReadCollection) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProductjsonldReadCollection) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProductjsonldReadCollection) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProductjsonldReadCollection) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetId

`func (o *ProductjsonldReadCollection) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductjsonldReadCollection) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductjsonldReadCollection) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductjsonldReadCollection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProductjsonldReadCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductjsonldReadCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductjsonldReadCollection) SetName(v string)`

SetName sets Name field to given value.


### GetQuantityAvailable

`func (o *ProductjsonldReadCollection) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ProductjsonldReadCollection) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ProductjsonldReadCollection) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *ProductjsonldReadCollection) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### GetQuantityIncoming

`func (o *ProductjsonldReadCollection) GetQuantityIncoming() int32`

GetQuantityIncoming returns the QuantityIncoming field if non-nil, zero value otherwise.

### GetQuantityIncomingOk

`func (o *ProductjsonldReadCollection) GetQuantityIncomingOk() (*int32, bool)`

GetQuantityIncomingOk returns a tuple with the QuantityIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityIncoming

`func (o *ProductjsonldReadCollection) SetQuantityIncoming(v int32)`

SetQuantityIncoming sets QuantityIncoming field to given value.

### HasQuantityIncoming

`func (o *ProductjsonldReadCollection) HasQuantityIncoming() bool`

HasQuantityIncoming returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *ProductjsonldReadCollection) GetQuantityOnHand() int32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *ProductjsonldReadCollection) GetQuantityOnHandOk() (*int32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *ProductjsonldReadCollection) SetQuantityOnHand(v int32)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *ProductjsonldReadCollection) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### GetQuantityOutgoing

`func (o *ProductjsonldReadCollection) GetQuantityOutgoing() int32`

GetQuantityOutgoing returns the QuantityOutgoing field if non-nil, zero value otherwise.

### GetQuantityOutgoingOk

`func (o *ProductjsonldReadCollection) GetQuantityOutgoingOk() (*int32, bool)`

GetQuantityOutgoingOk returns a tuple with the QuantityOutgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOutgoing

`func (o *ProductjsonldReadCollection) SetQuantityOutgoing(v int32)`

SetQuantityOutgoing sets QuantityOutgoing field to given value.

### HasQuantityOutgoing

`func (o *ProductjsonldReadCollection) HasQuantityOutgoing() bool`

HasQuantityOutgoing returns a boolean if a field has been set.

### GetQuantitySold

`func (o *ProductjsonldReadCollection) GetQuantitySold() int32`

GetQuantitySold returns the QuantitySold field if non-nil, zero value otherwise.

### GetQuantitySoldOk

`func (o *ProductjsonldReadCollection) GetQuantitySoldOk() (*int32, bool)`

GetQuantitySoldOk returns a tuple with the QuantitySold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantitySold

`func (o *ProductjsonldReadCollection) SetQuantitySold(v int32)`

SetQuantitySold sets QuantitySold field to given value.

### HasQuantitySold

`func (o *ProductjsonldReadCollection) HasQuantitySold() bool`

HasQuantitySold returns a boolean if a field has been set.

### GetSellingPrice

`func (o *ProductjsonldReadCollection) GetSellingPrice() float32`

GetSellingPrice returns the SellingPrice field if non-nil, zero value otherwise.

### GetSellingPriceOk

`func (o *ProductjsonldReadCollection) GetSellingPriceOk() (*float32, bool)`

GetSellingPriceOk returns a tuple with the SellingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPrice

`func (o *ProductjsonldReadCollection) SetSellingPrice(v float32)`

SetSellingPrice sets SellingPrice field to given value.

### HasSellingPrice

`func (o *ProductjsonldReadCollection) HasSellingPrice() bool`

HasSellingPrice returns a boolean if a field has been set.

### GetSku

`func (o *ProductjsonldReadCollection) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductjsonldReadCollection) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductjsonldReadCollection) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductjsonldReadCollection) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVolume

`func (o *ProductjsonldReadCollection) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ProductjsonldReadCollection) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ProductjsonldReadCollection) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ProductjsonldReadCollection) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWeight

`func (o *ProductjsonldReadCollection) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductjsonldReadCollection) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductjsonldReadCollection) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductjsonldReadCollection) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


