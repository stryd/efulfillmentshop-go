/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
)

// ProductjsonldWrite Product entity
type ProductjsonldWrite struct {
	JsonContext *string `json:"json_context,omitempty"`
	JsonId *string `json:"json_id,omitempty"`
	JsonType *string `json:"json_type,omitempty"`
	// The product barcode
	Barcode string `json:"barcode"`
	// The product channel
	ChannelId int32 `json:"channelId"`
	// The product channel reference
	ChannelReference string `json:"channelReference"`
	// The product channel SKU
	ChannelSku *string `json:"channelSku,omitempty"`
	// The product cost price (Excluding taxes)
	CostPrice *float32 `json:"costPrice,omitempty"`
	// The product dimension X in cm
	DimensionX *float32 `json:"dimensionX,omitempty"`
	// The product dimension Y in cm
	DimensionY *float32 `json:"dimensionY,omitempty"`
	// The product dimension Z in cm
	DimensionZ *float32 `json:"dimensionZ,omitempty"`
	// The product HS code
	HsCode *string `json:"hsCode,omitempty"`
	// The product image (As a base64 encoded string)
	Image *string `json:"image,omitempty"`
	// The product name
	Name string `json:"name"`
	// The product selling price (Excluding taxes)
	SellingPrice *float32 `json:"sellingPrice,omitempty"`
	// The product weight in kg
	Weight *float32 `json:"weight,omitempty"`
}

// NewProductjsonldWrite instantiates a new ProductjsonldWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductjsonldWrite(barcode string, channelId int32, channelReference string, name string, ) *ProductjsonldWrite {
	this := ProductjsonldWrite{}
	this.Barcode = barcode
	this.ChannelId = channelId
	this.ChannelReference = channelReference
	this.Name = name
	return &this
}

// NewProductjsonldWriteWithDefaults instantiates a new ProductjsonldWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductjsonldWriteWithDefaults() *ProductjsonldWrite {
	this := ProductjsonldWrite{}
	return &this
}

// GetJsonContext returns the JsonContext field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetJsonContext() string {
	if o == nil || o.JsonContext == nil {
		var ret string
		return ret
	}
	return *o.JsonContext
}

// GetJsonContextOk returns a tuple with the JsonContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetJsonContextOk() (*string, bool) {
	if o == nil || o.JsonContext == nil {
		return nil, false
	}
	return o.JsonContext, true
}

// HasJsonContext returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasJsonContext() bool {
	if o != nil && o.JsonContext != nil {
		return true
	}

	return false
}

// SetJsonContext gets a reference to the given string and assigns it to the JsonContext field.
func (o *ProductjsonldWrite) SetJsonContext(v string) {
	o.JsonContext = &v
}

// GetJsonId returns the JsonId field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetJsonId() string {
	if o == nil || o.JsonId == nil {
		var ret string
		return ret
	}
	return *o.JsonId
}

// GetJsonIdOk returns a tuple with the JsonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetJsonIdOk() (*string, bool) {
	if o == nil || o.JsonId == nil {
		return nil, false
	}
	return o.JsonId, true
}

// HasJsonId returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasJsonId() bool {
	if o != nil && o.JsonId != nil {
		return true
	}

	return false
}

// SetJsonId gets a reference to the given string and assigns it to the JsonId field.
func (o *ProductjsonldWrite) SetJsonId(v string) {
	o.JsonId = &v
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *ProductjsonldWrite) SetJsonType(v string) {
	o.JsonType = &v
}

// GetBarcode returns the Barcode field value
func (o *ProductjsonldWrite) GetBarcode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetBarcodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *ProductjsonldWrite) SetBarcode(v string) {
	o.Barcode = v
}

// GetChannelId returns the ChannelId field value
func (o *ProductjsonldWrite) GetChannelId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetChannelIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *ProductjsonldWrite) SetChannelId(v int32) {
	o.ChannelId = v
}

// GetChannelReference returns the ChannelReference field value
func (o *ProductjsonldWrite) GetChannelReference() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetChannelReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelReference, true
}

// SetChannelReference sets field value
func (o *ProductjsonldWrite) SetChannelReference(v string) {
	o.ChannelReference = v
}

// GetChannelSku returns the ChannelSku field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetChannelSku() string {
	if o == nil || o.ChannelSku == nil {
		var ret string
		return ret
	}
	return *o.ChannelSku
}

// GetChannelSkuOk returns a tuple with the ChannelSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetChannelSkuOk() (*string, bool) {
	if o == nil || o.ChannelSku == nil {
		return nil, false
	}
	return o.ChannelSku, true
}

// HasChannelSku returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasChannelSku() bool {
	if o != nil && o.ChannelSku != nil {
		return true
	}

	return false
}

// SetChannelSku gets a reference to the given string and assigns it to the ChannelSku field.
func (o *ProductjsonldWrite) SetChannelSku(v string) {
	o.ChannelSku = &v
}

// GetCostPrice returns the CostPrice field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetCostPrice() float32 {
	if o == nil || o.CostPrice == nil {
		var ret float32
		return ret
	}
	return *o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetCostPriceOk() (*float32, bool) {
	if o == nil || o.CostPrice == nil {
		return nil, false
	}
	return o.CostPrice, true
}

// HasCostPrice returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasCostPrice() bool {
	if o != nil && o.CostPrice != nil {
		return true
	}

	return false
}

// SetCostPrice gets a reference to the given float32 and assigns it to the CostPrice field.
func (o *ProductjsonldWrite) SetCostPrice(v float32) {
	o.CostPrice = &v
}

// GetDimensionX returns the DimensionX field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetDimensionX() float32 {
	if o == nil || o.DimensionX == nil {
		var ret float32
		return ret
	}
	return *o.DimensionX
}

// GetDimensionXOk returns a tuple with the DimensionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetDimensionXOk() (*float32, bool) {
	if o == nil || o.DimensionX == nil {
		return nil, false
	}
	return o.DimensionX, true
}

// HasDimensionX returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasDimensionX() bool {
	if o != nil && o.DimensionX != nil {
		return true
	}

	return false
}

// SetDimensionX gets a reference to the given float32 and assigns it to the DimensionX field.
func (o *ProductjsonldWrite) SetDimensionX(v float32) {
	o.DimensionX = &v
}

// GetDimensionY returns the DimensionY field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetDimensionY() float32 {
	if o == nil || o.DimensionY == nil {
		var ret float32
		return ret
	}
	return *o.DimensionY
}

// GetDimensionYOk returns a tuple with the DimensionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetDimensionYOk() (*float32, bool) {
	if o == nil || o.DimensionY == nil {
		return nil, false
	}
	return o.DimensionY, true
}

// HasDimensionY returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasDimensionY() bool {
	if o != nil && o.DimensionY != nil {
		return true
	}

	return false
}

// SetDimensionY gets a reference to the given float32 and assigns it to the DimensionY field.
func (o *ProductjsonldWrite) SetDimensionY(v float32) {
	o.DimensionY = &v
}

// GetDimensionZ returns the DimensionZ field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetDimensionZ() float32 {
	if o == nil || o.DimensionZ == nil {
		var ret float32
		return ret
	}
	return *o.DimensionZ
}

// GetDimensionZOk returns a tuple with the DimensionZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetDimensionZOk() (*float32, bool) {
	if o == nil || o.DimensionZ == nil {
		return nil, false
	}
	return o.DimensionZ, true
}

// HasDimensionZ returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasDimensionZ() bool {
	if o != nil && o.DimensionZ != nil {
		return true
	}

	return false
}

// SetDimensionZ gets a reference to the given float32 and assigns it to the DimensionZ field.
func (o *ProductjsonldWrite) SetDimensionZ(v float32) {
	o.DimensionZ = &v
}

// GetHsCode returns the HsCode field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetHsCode() string {
	if o == nil || o.HsCode == nil {
		var ret string
		return ret
	}
	return *o.HsCode
}

// GetHsCodeOk returns a tuple with the HsCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetHsCodeOk() (*string, bool) {
	if o == nil || o.HsCode == nil {
		return nil, false
	}
	return o.HsCode, true
}

// HasHsCode returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasHsCode() bool {
	if o != nil && o.HsCode != nil {
		return true
	}

	return false
}

// SetHsCode gets a reference to the given string and assigns it to the HsCode field.
func (o *ProductjsonldWrite) SetHsCode(v string) {
	o.HsCode = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ProductjsonldWrite) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value
func (o *ProductjsonldWrite) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductjsonldWrite) SetName(v string) {
	o.Name = v
}

// GetSellingPrice returns the SellingPrice field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetSellingPrice() float32 {
	if o == nil || o.SellingPrice == nil {
		var ret float32
		return ret
	}
	return *o.SellingPrice
}

// GetSellingPriceOk returns a tuple with the SellingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetSellingPriceOk() (*float32, bool) {
	if o == nil || o.SellingPrice == nil {
		return nil, false
	}
	return o.SellingPrice, true
}

// HasSellingPrice returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasSellingPrice() bool {
	if o != nil && o.SellingPrice != nil {
		return true
	}

	return false
}

// SetSellingPrice gets a reference to the given float32 and assigns it to the SellingPrice field.
func (o *ProductjsonldWrite) SetSellingPrice(v float32) {
	o.SellingPrice = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ProductjsonldWrite) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductjsonldWrite) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ProductjsonldWrite) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *ProductjsonldWrite) SetWeight(v float32) {
	o.Weight = &v
}

func (o ProductjsonldWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JsonContext != nil {
		toSerialize["json_context"] = o.JsonContext
	}
	if o.JsonId != nil {
		toSerialize["json_id"] = o.JsonId
	}
	if o.JsonType != nil {
		toSerialize["json_type"] = o.JsonType
	}
	if true {
		toSerialize["barcode"] = o.Barcode
	}
	if true {
		toSerialize["channelId"] = o.ChannelId
	}
	if true {
		toSerialize["channelReference"] = o.ChannelReference
	}
	if o.ChannelSku != nil {
		toSerialize["channelSku"] = o.ChannelSku
	}
	if o.CostPrice != nil {
		toSerialize["costPrice"] = o.CostPrice
	}
	if o.DimensionX != nil {
		toSerialize["dimensionX"] = o.DimensionX
	}
	if o.DimensionY != nil {
		toSerialize["dimensionY"] = o.DimensionY
	}
	if o.DimensionZ != nil {
		toSerialize["dimensionZ"] = o.DimensionZ
	}
	if o.HsCode != nil {
		toSerialize["hsCode"] = o.HsCode
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.SellingPrice != nil {
		toSerialize["sellingPrice"] = o.SellingPrice
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableProductjsonldWrite struct {
	value *ProductjsonldWrite
	isSet bool
}

func (v NullableProductjsonldWrite) Get() *ProductjsonldWrite {
	return v.value
}

func (v *NullableProductjsonldWrite) Set(val *ProductjsonldWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableProductjsonldWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableProductjsonldWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductjsonldWrite(val *ProductjsonldWrite) *NullableProductjsonldWrite {
	return &NullableProductjsonldWrite{value: val, isSet: true}
}

func (v NullableProductjsonldWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductjsonldWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


