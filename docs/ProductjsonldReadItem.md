# ProductjsonldReadItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**Barcode** | **string** | The product barcode | 
**ChannelId** | **int32** | The product channel | 
**ChannelReference** | **string** | The product channel reference | 
**ChannelSku** | Pointer to **string** | The product channel SKU | [optional] 
**CostPrice** | Pointer to **float32** | The product cost price (Excluding taxes) | [optional] 
**DimensionX** | Pointer to **float32** | The product dimension X in cm | [optional] 
**DimensionY** | Pointer to **float32** | The product dimension Y in cm | [optional] 
**DimensionZ** | Pointer to **float32** | The product dimension Z in cm | [optional] 
**HsCode** | Pointer to **string** | The product HS code | [optional] 
**Id** | Pointer to **int32** | The product ID | [optional] [readonly] 
**Image** | Pointer to **string** | The product image (As a base64 encoded string) | [optional] 
**Name** | **string** | The product name | 
**QuantityAvailable** | Pointer to **int32** | The physical available product quantity | [optional] [readonly] 
**QuantityIncoming** | Pointer to **int32** | The incoming product quantity | [optional] [readonly] 
**QuantityOnHand** | Pointer to **int32** | The product quantity on hand | [optional] [readonly] 
**QuantityOutgoing** | Pointer to **int32** | The outgoing product quantity | [optional] [readonly] 
**QuantitySold** | Pointer to **int32** | The sold product quantity | [optional] [readonly] 
**SellingPrice** | Pointer to **float32** | The product selling price (Excluding taxes) | [optional] 
**Sku** | Pointer to **string** | The product warehouse SKU | [optional] [readonly] 
**Volume** | Pointer to **float32** | The product volume in L (Calculated using dimensions) | [optional] [readonly] 
**Weight** | Pointer to **float32** | The product weight in kg | [optional] 

## Methods

### NewProductjsonldReadItem

`func NewProductjsonldReadItem(barcode string, channelId int32, channelReference string, name string, ) *ProductjsonldReadItem`

NewProductjsonldReadItem instantiates a new ProductjsonldReadItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductjsonldReadItemWithDefaults

`func NewProductjsonldReadItemWithDefaults() *ProductjsonldReadItem`

NewProductjsonldReadItemWithDefaults instantiates a new ProductjsonldReadItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ProductjsonldReadItem) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ProductjsonldReadItem) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ProductjsonldReadItem) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ProductjsonldReadItem) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *ProductjsonldReadItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductjsonldReadItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductjsonldReadItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductjsonldReadItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ProductjsonldReadItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductjsonldReadItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductjsonldReadItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductjsonldReadItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductjsonldReadItem) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductjsonldReadItem) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductjsonldReadItem) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetChannelId

`func (o *ProductjsonldReadItem) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ProductjsonldReadItem) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ProductjsonldReadItem) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *ProductjsonldReadItem) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ProductjsonldReadItem) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ProductjsonldReadItem) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetChannelSku

`func (o *ProductjsonldReadItem) GetChannelSku() string`

GetChannelSku returns the ChannelSku field if non-nil, zero value otherwise.

### GetChannelSkuOk

`func (o *ProductjsonldReadItem) GetChannelSkuOk() (*string, bool)`

GetChannelSkuOk returns a tuple with the ChannelSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSku

`func (o *ProductjsonldReadItem) SetChannelSku(v string)`

SetChannelSku sets ChannelSku field to given value.

### HasChannelSku

`func (o *ProductjsonldReadItem) HasChannelSku() bool`

HasChannelSku returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductjsonldReadItem) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductjsonldReadItem) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductjsonldReadItem) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductjsonldReadItem) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetDimensionX

`func (o *ProductjsonldReadItem) GetDimensionX() float32`

GetDimensionX returns the DimensionX field if non-nil, zero value otherwise.

### GetDimensionXOk

`func (o *ProductjsonldReadItem) GetDimensionXOk() (*float32, bool)`

GetDimensionXOk returns a tuple with the DimensionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionX

`func (o *ProductjsonldReadItem) SetDimensionX(v float32)`

SetDimensionX sets DimensionX field to given value.

### HasDimensionX

`func (o *ProductjsonldReadItem) HasDimensionX() bool`

HasDimensionX returns a boolean if a field has been set.

### GetDimensionY

`func (o *ProductjsonldReadItem) GetDimensionY() float32`

GetDimensionY returns the DimensionY field if non-nil, zero value otherwise.

### GetDimensionYOk

`func (o *ProductjsonldReadItem) GetDimensionYOk() (*float32, bool)`

GetDimensionYOk returns a tuple with the DimensionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionY

`func (o *ProductjsonldReadItem) SetDimensionY(v float32)`

SetDimensionY sets DimensionY field to given value.

### HasDimensionY

`func (o *ProductjsonldReadItem) HasDimensionY() bool`

HasDimensionY returns a boolean if a field has been set.

### GetDimensionZ

`func (o *ProductjsonldReadItem) GetDimensionZ() float32`

GetDimensionZ returns the DimensionZ field if non-nil, zero value otherwise.

### GetDimensionZOk

`func (o *ProductjsonldReadItem) GetDimensionZOk() (*float32, bool)`

GetDimensionZOk returns a tuple with the DimensionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionZ

`func (o *ProductjsonldReadItem) SetDimensionZ(v float32)`

SetDimensionZ sets DimensionZ field to given value.

### HasDimensionZ

`func (o *ProductjsonldReadItem) HasDimensionZ() bool`

HasDimensionZ returns a boolean if a field has been set.

### GetHsCode

`func (o *ProductjsonldReadItem) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProductjsonldReadItem) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProductjsonldReadItem) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProductjsonldReadItem) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetId

`func (o *ProductjsonldReadItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductjsonldReadItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductjsonldReadItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductjsonldReadItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *ProductjsonldReadItem) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductjsonldReadItem) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductjsonldReadItem) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductjsonldReadItem) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *ProductjsonldReadItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductjsonldReadItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductjsonldReadItem) SetName(v string)`

SetName sets Name field to given value.


### GetQuantityAvailable

`func (o *ProductjsonldReadItem) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ProductjsonldReadItem) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ProductjsonldReadItem) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *ProductjsonldReadItem) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### GetQuantityIncoming

`func (o *ProductjsonldReadItem) GetQuantityIncoming() int32`

GetQuantityIncoming returns the QuantityIncoming field if non-nil, zero value otherwise.

### GetQuantityIncomingOk

`func (o *ProductjsonldReadItem) GetQuantityIncomingOk() (*int32, bool)`

GetQuantityIncomingOk returns a tuple with the QuantityIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityIncoming

`func (o *ProductjsonldReadItem) SetQuantityIncoming(v int32)`

SetQuantityIncoming sets QuantityIncoming field to given value.

### HasQuantityIncoming

`func (o *ProductjsonldReadItem) HasQuantityIncoming() bool`

HasQuantityIncoming returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *ProductjsonldReadItem) GetQuantityOnHand() int32`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *ProductjsonldReadItem) GetQuantityOnHandOk() (*int32, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *ProductjsonldReadItem) SetQuantityOnHand(v int32)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *ProductjsonldReadItem) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### GetQuantityOutgoing

`func (o *ProductjsonldReadItem) GetQuantityOutgoing() int32`

GetQuantityOutgoing returns the QuantityOutgoing field if non-nil, zero value otherwise.

### GetQuantityOutgoingOk

`func (o *ProductjsonldReadItem) GetQuantityOutgoingOk() (*int32, bool)`

GetQuantityOutgoingOk returns a tuple with the QuantityOutgoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOutgoing

`func (o *ProductjsonldReadItem) SetQuantityOutgoing(v int32)`

SetQuantityOutgoing sets QuantityOutgoing field to given value.

### HasQuantityOutgoing

`func (o *ProductjsonldReadItem) HasQuantityOutgoing() bool`

HasQuantityOutgoing returns a boolean if a field has been set.

### GetQuantitySold

`func (o *ProductjsonldReadItem) GetQuantitySold() int32`

GetQuantitySold returns the QuantitySold field if non-nil, zero value otherwise.

### GetQuantitySoldOk

`func (o *ProductjsonldReadItem) GetQuantitySoldOk() (*int32, bool)`

GetQuantitySoldOk returns a tuple with the QuantitySold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantitySold

`func (o *ProductjsonldReadItem) SetQuantitySold(v int32)`

SetQuantitySold sets QuantitySold field to given value.

### HasQuantitySold

`func (o *ProductjsonldReadItem) HasQuantitySold() bool`

HasQuantitySold returns a boolean if a field has been set.

### GetSellingPrice

`func (o *ProductjsonldReadItem) GetSellingPrice() float32`

GetSellingPrice returns the SellingPrice field if non-nil, zero value otherwise.

### GetSellingPriceOk

`func (o *ProductjsonldReadItem) GetSellingPriceOk() (*float32, bool)`

GetSellingPriceOk returns a tuple with the SellingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPrice

`func (o *ProductjsonldReadItem) SetSellingPrice(v float32)`

SetSellingPrice sets SellingPrice field to given value.

### HasSellingPrice

`func (o *ProductjsonldReadItem) HasSellingPrice() bool`

HasSellingPrice returns a boolean if a field has been set.

### GetSku

`func (o *ProductjsonldReadItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ProductjsonldReadItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ProductjsonldReadItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ProductjsonldReadItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetVolume

`func (o *ProductjsonldReadItem) GetVolume() float32`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ProductjsonldReadItem) GetVolumeOk() (*float32, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ProductjsonldReadItem) SetVolume(v float32)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *ProductjsonldReadItem) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetWeight

`func (o *ProductjsonldReadItem) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductjsonldReadItem) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductjsonldReadItem) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductjsonldReadItem) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


