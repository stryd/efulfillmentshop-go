# ProductWriteLd

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
**Image** | Pointer to **string** | The product image (As a base64 encoded string) | [optional] 
**Name** | **string** | The product name | 
**Reference** | Pointer to **string** | Your product reference (This could be your product ID) | [optional] 
**SellingPrice** | Pointer to **float32** | The product selling price (Excluding taxes) | [optional] 
**Weight** | Pointer to **float32** | The product weight in kg | [optional] 

## Methods

### NewProductWriteLd

`func NewProductWriteLd(barcode string, name string, ) *ProductWriteLd`

NewProductWriteLd instantiates a new ProductWriteLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWriteLdWithDefaults

`func NewProductWriteLdWithDefaults() *ProductWriteLd`

NewProductWriteLdWithDefaults instantiates a new ProductWriteLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *ProductWriteLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *ProductWriteLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *ProductWriteLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *ProductWriteLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *ProductWriteLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *ProductWriteLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *ProductWriteLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *ProductWriteLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *ProductWriteLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *ProductWriteLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *ProductWriteLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *ProductWriteLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetBarcode

`func (o *ProductWriteLd) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *ProductWriteLd) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *ProductWriteLd) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.


### GetChannelReference

`func (o *ProductWriteLd) GetChannelReference() string`

GetChannelReference returns the ChannelReference field if non-nil, zero value otherwise.

### GetChannelReferenceOk

`func (o *ProductWriteLd) GetChannelReferenceOk() (*string, bool)`

GetChannelReferenceOk returns a tuple with the ChannelReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelReference

`func (o *ProductWriteLd) SetChannelReference(v string)`

SetChannelReference sets ChannelReference field to given value.

### HasChannelReference

`func (o *ProductWriteLd) HasChannelReference() bool`

HasChannelReference returns a boolean if a field has been set.

### GetCostPrice

`func (o *ProductWriteLd) GetCostPrice() float32`

GetCostPrice returns the CostPrice field if non-nil, zero value otherwise.

### GetCostPriceOk

`func (o *ProductWriteLd) GetCostPriceOk() (*float32, bool)`

GetCostPriceOk returns a tuple with the CostPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPrice

`func (o *ProductWriteLd) SetCostPrice(v float32)`

SetCostPrice sets CostPrice field to given value.

### HasCostPrice

`func (o *ProductWriteLd) HasCostPrice() bool`

HasCostPrice returns a boolean if a field has been set.

### GetDimensionX

`func (o *ProductWriteLd) GetDimensionX() float32`

GetDimensionX returns the DimensionX field if non-nil, zero value otherwise.

### GetDimensionXOk

`func (o *ProductWriteLd) GetDimensionXOk() (*float32, bool)`

GetDimensionXOk returns a tuple with the DimensionX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionX

`func (o *ProductWriteLd) SetDimensionX(v float32)`

SetDimensionX sets DimensionX field to given value.

### HasDimensionX

`func (o *ProductWriteLd) HasDimensionX() bool`

HasDimensionX returns a boolean if a field has been set.

### GetDimensionY

`func (o *ProductWriteLd) GetDimensionY() float32`

GetDimensionY returns the DimensionY field if non-nil, zero value otherwise.

### GetDimensionYOk

`func (o *ProductWriteLd) GetDimensionYOk() (*float32, bool)`

GetDimensionYOk returns a tuple with the DimensionY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionY

`func (o *ProductWriteLd) SetDimensionY(v float32)`

SetDimensionY sets DimensionY field to given value.

### HasDimensionY

`func (o *ProductWriteLd) HasDimensionY() bool`

HasDimensionY returns a boolean if a field has been set.

### GetDimensionZ

`func (o *ProductWriteLd) GetDimensionZ() float32`

GetDimensionZ returns the DimensionZ field if non-nil, zero value otherwise.

### GetDimensionZOk

`func (o *ProductWriteLd) GetDimensionZOk() (*float32, bool)`

GetDimensionZOk returns a tuple with the DimensionZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionZ

`func (o *ProductWriteLd) SetDimensionZ(v float32)`

SetDimensionZ sets DimensionZ field to given value.

### HasDimensionZ

`func (o *ProductWriteLd) HasDimensionZ() bool`

HasDimensionZ returns a boolean if a field has been set.

### GetHsCode

`func (o *ProductWriteLd) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *ProductWriteLd) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *ProductWriteLd) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *ProductWriteLd) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetImage

`func (o *ProductWriteLd) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ProductWriteLd) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ProductWriteLd) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ProductWriteLd) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *ProductWriteLd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductWriteLd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductWriteLd) SetName(v string)`

SetName sets Name field to given value.


### GetReference

`func (o *ProductWriteLd) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ProductWriteLd) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ProductWriteLd) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ProductWriteLd) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSellingPrice

`func (o *ProductWriteLd) GetSellingPrice() float32`

GetSellingPrice returns the SellingPrice field if non-nil, zero value otherwise.

### GetSellingPriceOk

`func (o *ProductWriteLd) GetSellingPriceOk() (*float32, bool)`

GetSellingPriceOk returns a tuple with the SellingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingPrice

`func (o *ProductWriteLd) SetSellingPrice(v float32)`

SetSellingPrice sets SellingPrice field to given value.

### HasSellingPrice

`func (o *ProductWriteLd) HasSellingPrice() bool`

HasSellingPrice returns a boolean if a field has been set.

### GetWeight

`func (o *ProductWriteLd) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ProductWriteLd) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ProductWriteLd) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ProductWriteLd) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


