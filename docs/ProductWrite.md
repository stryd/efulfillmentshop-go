# ProductWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | **string** | The product barcode | 
**ChannelReference** | Pointer to **string** | The product channel reference (DEPRECATED) | [optional] 
**CostPrice** | Pointer to **float32** | The product cost price (Excluding taxes) | [optional] 
**DimensionX** | Pointer to **float32** | The product dimension X in cm | [optional] 
**DimensionY** | Pointer to **float32** | The product dimension Y in cm | [optional] 
**DimensionZ** | Pointer to **float32** | The product dimension Z in cm | [optional] 
**HsCode** | Pointer to **string** | The product HS code | [optional] 
**Image** | Pointer to **string** | The product image (As a base64 encoded string) | [optional] 
**Name** | **string** | The product name | 
**Reference** | Pointer to **string** | Your product reference (This could be your product ID) | [optional] 
**SellingPrice** | Pointer to **float32** | The product selling price (Excluding taxes) | [optional] 
**Weight** | Pointer to **float32** | The product weight in kg | [optional] 

## Methods

### NewProductWrite

`func NewProductWrite(barcode string, name string, ) *ProductWrite`

NewProductWrite instantiates a new ProductWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWriteWithDefaults

`func NewProductWriteWithDefaults() *ProductWrite`

NewProductWriteWithDefaults instantiates a new ProductWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcode

`func (o *ProductWrite) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductWrite) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductWrite) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetChannelReference

`func (o *ProductWrite) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ProductWrite) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ProductWrite) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.

### HasChannelReference

`func (o *ProductWrite) HasChannelReference() bool`

HasChannelReference returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductWrite) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductWrite) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductWrite) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductWrite) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetDimensionX

`func (o *ProductWrite) GetDimensionX() float32`

GetDimensionX returns the DimensionX field if non-nil, zero value otherwise.

### GetDimensionXOk

`func (o *ProductWrite) GetDimensionXOk() (*float32, bool)`

GetDimensionXOk returns a tuple with the DimensionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionX

`func (o *ProductWrite) SetDimensionX(v float32)`

SetDimensionX sets DimensionX field to given value.

### HasDimensionX

`func (o *ProductWrite) HasDimensionX() bool`

HasDimensionX returns a boolean if a field has been set.

### GetDimensionY

`func (o *ProductWrite) GetDimensionY() float32`

GetDimensionY returns the DimensionY field if non-nil, zero value otherwise.

### GetDimensionYOk

`func (o *ProductWrite) GetDimensionYOk() (*float32, bool)`

GetDimensionYOk returns a tuple with the DimensionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionY

`func (o *ProductWrite) SetDimensionY(v float32)`

SetDimensionY sets DimensionY field to given value.

### HasDimensionY

`func (o *ProductWrite) HasDimensionY() bool`

HasDimensionY returns a boolean if a field has been set.

### GetDimensionZ

`func (o *ProductWrite) GetDimensionZ() float32`

GetDimensionZ returns the DimensionZ field if non-nil, zero value otherwise.

### GetDimensionZOk

`func (o *ProductWrite) GetDimensionZOk() (*float32, bool)`

GetDimensionZOk returns a tuple with the DimensionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionZ

`func (o *ProductWrite) SetDimensionZ(v float32)`

SetDimensionZ sets DimensionZ field to given value.

### HasDimensionZ

`func (o *ProductWrite) HasDimensionZ() bool`

HasDimensionZ returns a boolean if a field has been set.

### GetHsCode

`func (o *ProductWrite) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProductWrite) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProductWrite) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProductWrite) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetImage

`func (o *ProductWrite) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductWrite) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductWrite) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductWrite) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *ProductWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductWrite) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *ProductWrite) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ProductWrite) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ProductWrite) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ProductWrite) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSellingPrice

`func (o *ProductWrite) GetSellingPrice() float32`

GetSellingPrice returns the SellingPrice field if non-nil, zero value otherwise.

### GetSellingPriceOk

`func (o *ProductWrite) GetSellingPriceOk() (*float32, bool)`

GetSellingPriceOk returns a tuple with the SellingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPrice

`func (o *ProductWrite) SetSellingPrice(v float32)`

SetSellingPrice sets SellingPrice field to given value.

### HasSellingPrice

`func (o *ProductWrite) HasSellingPrice() bool`

HasSellingPrice returns a boolean if a field has been set.

### GetWeight

`func (o *ProductWrite) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductWrite) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductWrite) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductWrite) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


