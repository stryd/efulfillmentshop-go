# ProductjsonldWrite

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
**Image** | Pointer to **string** | The product image (As a base64 encoded string) | [optional] 
**Name** | **string** | The product name | 
**SellingPrice** | Pointer to **float32** | The product selling price (Excluding taxes) | [optional] 
**Weight** | Pointer to **float32** | The product weight in kg | [optional] 

## Methods

### NewProductjsonldWrite

`func NewProductjsonldWrite(barcode string, channelId int32, channelReference string, name string, ) *ProductjsonldWrite`

NewProductjsonldWrite instantiates a new ProductjsonldWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductjsonldWriteWithDefaults

`func NewProductjsonldWriteWithDefaults() *ProductjsonldWrite`

NewProductjsonldWriteWithDefaults instantiates a new ProductjsonldWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ProductjsonldWrite) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ProductjsonldWrite) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ProductjsonldWrite) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *ProductjsonldWrite) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *ProductjsonldWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductjsonldWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductjsonldWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProductjsonldWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ProductjsonldWrite) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductjsonldWrite) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductjsonldWrite) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProductjsonldWrite) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductjsonldWrite) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductjsonldWrite) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductjsonldWrite) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetChannelId

`func (o *ProductjsonldWrite) GetChannelId() int32`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ProductjsonldWrite) GetChannelIdOk() (*int32, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ProductjsonldWrite) SetChannelId(v int32)`

SetChannelId sets ChannelId field to given value.


### GetChannelReference

`func (o *ProductjsonldWrite) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ProductjsonldWrite) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ProductjsonldWrite) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.


### GetChannelSku

`func (o *ProductjsonldWrite) GetChannelSku() string`

GetChannelSku returns the ChannelSku field if non-nil, zero value otherwise.

### GetChannelSkuOk

`func (o *ProductjsonldWrite) GetChannelSkuOk() (*string, bool)`

GetChannelSkuOk returns a tuple with the ChannelSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelSku

`func (o *ProductjsonldWrite) SetChannelSku(v string)`

SetChannelSku sets ChannelSku field to given value.

### HasChannelSku

`func (o *ProductjsonldWrite) HasChannelSku() bool`

HasChannelSku returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductjsonldWrite) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductjsonldWrite) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductjsonldWrite) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductjsonldWrite) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetDimensionX

`func (o *ProductjsonldWrite) GetDimensionX() float32`

GetDimensionX returns the DimensionX field if non-nil, zero value otherwise.

### GetDimensionXOk

`func (o *ProductjsonldWrite) GetDimensionXOk() (*float32, bool)`

GetDimensionXOk returns a tuple with the DimensionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionX

`func (o *ProductjsonldWrite) SetDimensionX(v float32)`

SetDimensionX sets DimensionX field to given value.

### HasDimensionX

`func (o *ProductjsonldWrite) HasDimensionX() bool`

HasDimensionX returns a boolean if a field has been set.

### GetDimensionY

`func (o *ProductjsonldWrite) GetDimensionY() float32`

GetDimensionY returns the DimensionY field if non-nil, zero value otherwise.

### GetDimensionYOk

`func (o *ProductjsonldWrite) GetDimensionYOk() (*float32, bool)`

GetDimensionYOk returns a tuple with the DimensionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionY

`func (o *ProductjsonldWrite) SetDimensionY(v float32)`

SetDimensionY sets DimensionY field to given value.

### HasDimensionY

`func (o *ProductjsonldWrite) HasDimensionY() bool`

HasDimensionY returns a boolean if a field has been set.

### GetDimensionZ

`func (o *ProductjsonldWrite) GetDimensionZ() float32`

GetDimensionZ returns the DimensionZ field if non-nil, zero value otherwise.

### GetDimensionZOk

`func (o *ProductjsonldWrite) GetDimensionZOk() (*float32, bool)`

GetDimensionZOk returns a tuple with the DimensionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionZ

`func (o *ProductjsonldWrite) SetDimensionZ(v float32)`

SetDimensionZ sets DimensionZ field to given value.

### HasDimensionZ

`func (o *ProductjsonldWrite) HasDimensionZ() bool`

HasDimensionZ returns a boolean if a field has been set.

### GetHsCode

`func (o *ProductjsonldWrite) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProductjsonldWrite) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProductjsonldWrite) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProductjsonldWrite) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetImage

`func (o *ProductjsonldWrite) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductjsonldWrite) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductjsonldWrite) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductjsonldWrite) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *ProductjsonldWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductjsonldWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductjsonldWrite) SetName(v string)`

SetName sets Name field to given value.


### GetSellingPrice

`func (o *ProductjsonldWrite) GetSellingPrice() float32`

GetSellingPrice returns the SellingPrice field if non-nil, zero value otherwise.

### GetSellingPriceOk

`func (o *ProductjsonldWrite) GetSellingPriceOk() (*float32, bool)`

GetSellingPriceOk returns a tuple with the SellingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPrice

`func (o *ProductjsonldWrite) SetSellingPrice(v float32)`

SetSellingPrice sets SellingPrice field to given value.

### HasSellingPrice

`func (o *ProductjsonldWrite) HasSellingPrice() bool`

HasSellingPrice returns a boolean if a field has been set.

### GetWeight

`func (o *ProductjsonldWrite) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductjsonldWrite) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductjsonldWrite) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductjsonldWrite) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


