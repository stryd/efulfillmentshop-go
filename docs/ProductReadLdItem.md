# ProductReadLdItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**Barcode** | **string** | The product barcode | 
**ChannelReference** | Pointer to **string** | The product channel reference (DEPRECATED) | [optional] 
**CostPrice** | Pointer to **float32** | The product cost price (Excluding taxes) | [optional] 
**DimensionX** | Pointer to **float32** | The product dimension X in cm | [optional] 
**DimensionY** | Pointer to **float32** | The product dimension Y in cm | [optional] 
**DimensionZ** | Pointer to **float32** | The product dimension Z in cm | [optional] 
**HsCode** | Pointer to **string** | The product HS code | [optional] 
**Id** | Pointer to **int32** | The product ID | [optional] 
**Image** | Pointer to **string** | The product image (As a base64 encoded string) | [optional] 
**Name** | **string** | The product name | 
**QuantityAvailable** | Pointer to **int32** | The physical available product quantity | [optional] 
**QuantityIncoming** | Pointer to **int32** | The incoming product quantity | [optional] 
**QuantityOnHand** | Pointer to **int32** | The product quantity on hand | [optional] 
**QuantityOutgoing** | Pointer to **int32** | The outgoing product quantity | [optional] 
**QuantitySold** | Pointer to **int32** | The sold product quantity | [optional] 
**Reference** | Pointer to **string** | Your product reference (This could be your product ID) | [optional] 
**SellingPrice** | Pointer to **float32** | The product selling price (Excluding taxes) | [optional] 
**Sku** | Pointer to **string** | The product warehouse SKU | [optional] 
**Volume** | Pointer to **float32** | The product volume in L (Calculated using dimensions) | [optional] 
**Weight** | Pointer to **float32** | The product weight in kg | [optional] 

## Methods

### NewProductReadLdItem

`func NewProductReadLdItem(barcode string, name string, ) *ProductReadLdItem`

NewProductReadLdItem instantiates a new ProductReadLdItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductReadLdItemWithDefaults

`func NewProductReadLdItemWithDefaults() *ProductReadLdItem`

NewProductReadLdItemWithDefaults instantiates a new ProductReadLdItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *ProductReadLdItem) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *ProductReadLdItem) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *ProductReadLdItem) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *ProductReadLdItem) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *ProductReadLdItem) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *ProductReadLdItem) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *ProductReadLdItem) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *ProductReadLdItem) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *ProductReadLdItem) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *ProductReadLdItem) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *ProductReadLdItem) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *ProductReadLdItem) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductReadLdItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductReadLdItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductReadLdItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetChannelReference

`func (o *ProductReadLdItem) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ProductReadLdItem) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ProductReadLdItem) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.

### HasChannelReference

`func (o *ProductReadLdItem) HasChannelReference() bool`

HasChannelReference returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductReadLdItem) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductReadLdItem) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductReadLdItem) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductReadLdItem) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetDimensionX

`func (o *ProductReadLdItem) GetDimensionX() float32`

GetDimensionX returns the DimensionX field if non-nil, zero value otherwise.

### GetDimensionXOk

`func (o *ProductReadLdItem) GetDimensionXOk() (*float32, bool)`

GetDimensionXOk returns a tuple with the DimensionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionX

`func (o *ProductReadLdItem) SetDimensionX(v float32)`

SetDimensionX sets DimensionX field to given value.

### HasDimensionX

`func (o *ProductReadLdItem) HasDimensionX() bool`

HasDimensionX returns a boolean if a field has been set.

### GetDimensionY

`func (o *ProductReadLdItem) GetDimensionY() float32`

GetDimensionY returns the DimensionY field if non-nil, zero value otherwise.

### GetDimensionYOk

`func (o *ProductReadLdItem) GetDimensionYOk() (*float32, bool)`

GetDimensionYOk returns a tuple with the DimensionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionY

`func (o *ProductReadLdItem) SetDimensionY(v float32)`

SetDimensionY sets DimensionY field to given value.

### HasDimensionY

`func (o *ProductReadLdItem) HasDimensionY() bool`

HasDimensionY returns a boolean if a field has been set.

### GetDimensionZ

`func (o *ProductReadLdItem) GetDimensionZ() float32`

GetDimensionZ returns the DimensionZ field if non-nil, zero value otherwise.

### GetDimensionZOk

`func (o *ProductReadLdItem) GetDimensionZOk() (*float32, bool)`

GetDimensionZOk returns a tuple with the DimensionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionZ

`func (o *ProductReadLdItem) SetDimensionZ(v float32)`

SetDimensionZ sets DimensionZ field to given value.

### HasDimensionZ

`func (o *ProductReadLdItem) HasDimensionZ() bool`

HasDimensionZ returns a boolean if a field has been set.

### GetHsCode

`func (o *ProductReadLdItem) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProductReadLdItem) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProductReadLdItem) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProductReadLdItem) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetId

`func (o *ProductReadLdItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductReadLdItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductReadLdItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductReadLdItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ProductReadLdItem) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductReadLdItem) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductReadLdItem) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductReadLdItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *ProductReadLdItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductReadLdItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductReadLdItem) SetName(v string)`

SetName sets Name field to given value.


### GetQuantityAvailable

`func (o *ProductReadLdItem) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ProductReadLdItem) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ProductReadLdItem) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *ProductReadLdItem) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### GetQuantityIncoming

`func (o *ProductReadLdItem) GetQuantityIncoming() int32`

GetQuantityIncoming returns the QuantityIncoming field if non-nil, zero value otherwise.

### GetQuantityIncomingOk

`func (o *ProductReadLdItem) GetQuantityIncomingOk() (*int32, bool)`

GetQuantityIncomingOk returns a tuple with the QuantityIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityIncoming

`func (o *ProductReadLdItem) SetQuantityIncoming(v int32)`

SetQuantityIncoming sets QuantityIncoming field to given value.

### HasQuantityIncoming

`func (o *ProductReadLdItem) HasQuantityIncoming() bool`

HasQuantityIncoming returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *ProductReadLdItem) GetQuantityOnHand() int32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *ProductReadLdItem) GetQuantityOnHandOk() (*int32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *ProductReadLdItem) SetQuantityOnHand(v int32)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *ProductReadLdItem) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### GetQuantityOutgoing

`func (o *ProductReadLdItem) GetQuantityOutgoing() int32`

GetQuantityOutgoing returns the QuantityOutgoing field if non-nil, zero value otherwise.

### GetQuantityOutgoingOk

`func (o *ProductReadLdItem) GetQuantityOutgoingOk() (*int32, bool)`

GetQuantityOutgoingOk returns a tuple with the QuantityOutgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOutgoing

`func (o *ProductReadLdItem) SetQuantityOutgoing(v int32)`

SetQuantityOutgoing sets QuantityOutgoing field to given value.

### HasQuantityOutgoing

`func (o *ProductReadLdItem) HasQuantityOutgoing() bool`

HasQuantityOutgoing returns a boolean if a field has been set.

### GetQuantitySold

`func (o *ProductReadLdItem) GetQuantitySold() int32`

GetQuantitySold returns the QuantitySold field if non-nil, zero value otherwise.

### GetQuantitySoldOk

`func (o *ProductReadLdItem) GetQuantitySoldOk() (*int32, bool)`

GetQuantitySoldOk returns a tuple with the QuantitySold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantitySold

`func (o *ProductReadLdItem) SetQuantitySold(v int32)`

SetQuantitySold sets QuantitySold field to given value.

### HasQuantitySold

`func (o *ProductReadLdItem) HasQuantitySold() bool`

HasQuantitySold returns a boolean if a field has been set.

### GetReference

`func (o *ProductReadLdItem) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ProductReadLdItem) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ProductReadLdItem) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ProductReadLdItem) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSellingPrice

`func (o *ProductReadLdItem) GetSellingPrice() float32`

GetSellingPrice returns the SellingPrice field if non-nil, zero value otherwise.

### GetSellingPriceOk

`func (o *ProductReadLdItem) GetSellingPriceOk() (*float32, bool)`

GetSellingPriceOk returns a tuple with the SellingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPrice

`func (o *ProductReadLdItem) SetSellingPrice(v float32)`

SetSellingPrice sets SellingPrice field to given value.

### HasSellingPrice

`func (o *ProductReadLdItem) HasSellingPrice() bool`

HasSellingPrice returns a boolean if a field has been set.

### GetSku

`func (o *ProductReadLdItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductReadLdItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductReadLdItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductReadLdItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVolume

`func (o *ProductReadLdItem) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ProductReadLdItem) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ProductReadLdItem) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ProductReadLdItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWeight

`func (o *ProductReadLdItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductReadLdItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductReadLdItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductReadLdItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


