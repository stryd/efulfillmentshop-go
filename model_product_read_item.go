/*
E-fulfilment Shop API

The E-fulfilment Shop API allows you to integrate your service with our warehouse.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
)

// ProductReadItem Product entity
type ProductReadItem struct {
	// The product barcode
	Barcode string `json:"barcode"`
	// The product channel reference (DEPRECATED)
	ChannelReference *string `json:"channelReference,omitempty"`
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
	// The product ID
	Id *int32 `json:"id,omitempty"`
	// The product image (As a base64 encoded string)
	Image *string `json:"image,omitempty"`
	// The product name
	Name string `json:"name"`
	// The physical available product quantity
	QuantityAvailable *int32 `json:"quantityAvailable,omitempty"`
	// The incoming product quantity
	QuantityIncoming *int32 `json:"quantityIncoming,omitempty"`
	// The product quantity on hand
	QuantityOnHand *int32 `json:"quantityOnHand,omitempty"`
	// The outgoing product quantity
	QuantityOutgoing *int32 `json:"quantityOutgoing,omitempty"`
	// The sold product quantity
	QuantitySold *int32 `json:"quantitySold,omitempty"`
	// Your product reference (This could be your product ID)
	Reference *string `json:"reference,omitempty"`
	// The product selling price (Excluding taxes)
	SellingPrice *float32 `json:"sellingPrice,omitempty"`
	// The product warehouse SKU
	Sku *string `json:"sku,omitempty"`
	// The product volume in L (Calculated using dimensions)
	Volume *float32 `json:"volume,omitempty"`
	// The product weight in kg
	Weight *float32 `json:"weight,omitempty"`
}

// NewProductReadItem instantiates a new ProductReadItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductReadItem(barcode string, name string) *ProductReadItem {
	this := ProductReadItem{}
	this.Barcode = barcode
	this.Name = name
	return &this
}

// NewProductReadItemWithDefaults instantiates a new ProductReadItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductReadItemWithDefaults() *ProductReadItem {
	this := ProductReadItem{}
	return &this
}

// GetBarcode returns the Barcode field value
func (o *ProductReadItem) GetBarcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetBarcodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Barcode, true
}

// SetBarcode sets field value
func (o *ProductReadItem) SetBarcode(v string) {
	o.Barcode = v
}

// GetChannelReference returns the ChannelReference field value if set, zero value otherwise.
func (o *ProductReadItem) GetChannelReference() string {
	if o == nil || o.ChannelReference == nil {
		var ret string
		return ret
	}
	return *o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetChannelReferenceOk() (*string, bool) {
	if o == nil || o.ChannelReference == nil {
		return nil, false
	}
	return o.ChannelReference, true
}

// HasChannelReference returns a boolean if a field has been set.
func (o *ProductReadItem) HasChannelReference() bool {
	if o != nil && o.ChannelReference != nil {
		return true
	}

	return false
}

// SetChannelReference gets a reference to the given string and assigns it to the ChannelReference field.
func (o *ProductReadItem) SetChannelReference(v string) {
	o.ChannelReference = &v
}

// GetCostPrice returns the CostPrice field value if set, zero value otherwise.
func (o *ProductReadItem) GetCostPrice() float32 {
	if o == nil || o.CostPrice == nil {
		var ret float32
		return ret
	}
	return *o.CostPrice
}

// GetCostPriceOk returns a tuple with the CostPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetCostPriceOk() (*float32, bool) {
	if o == nil || o.CostPrice == nil {
		return nil, false
	}
	return o.CostPrice, true
}

// HasCostPrice returns a boolean if a field has been set.
func (o *ProductReadItem) HasCostPrice() bool {
	if o != nil && o.CostPrice != nil {
		return true
	}

	return false
}

// SetCostPrice gets a reference to the given float32 and assigns it to the CostPrice field.
func (o *ProductReadItem) SetCostPrice(v float32) {
	o.CostPrice = &v
}

// GetDimensionX returns the DimensionX field value if set, zero value otherwise.
func (o *ProductReadItem) GetDimensionX() float32 {
	if o == nil || o.DimensionX == nil {
		var ret float32
		return ret
	}
	return *o.DimensionX
}

// GetDimensionXOk returns a tuple with the DimensionX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetDimensionXOk() (*float32, bool) {
	if o == nil || o.DimensionX == nil {
		return nil, false
	}
	return o.DimensionX, true
}

// HasDimensionX returns a boolean if a field has been set.
func (o *ProductReadItem) HasDimensionX() bool {
	if o != nil && o.DimensionX != nil {
		return true
	}

	return false
}

// SetDimensionX gets a reference to the given float32 and assigns it to the DimensionX field.
func (o *ProductReadItem) SetDimensionX(v float32) {
	o.DimensionX = &v
}

// GetDimensionY returns the DimensionY field value if set, zero value otherwise.
func (o *ProductReadItem) GetDimensionY() float32 {
	if o == nil || o.DimensionY == nil {
		var ret float32
		return ret
	}
	return *o.DimensionY
}

// GetDimensionYOk returns a tuple with the DimensionY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetDimensionYOk() (*float32, bool) {
	if o == nil || o.DimensionY == nil {
		return nil, false
	}
	return o.DimensionY, true
}

// HasDimensionY returns a boolean if a field has been set.
func (o *ProductReadItem) HasDimensionY() bool {
	if o != nil && o.DimensionY != nil {
		return true
	}

	return false
}

// SetDimensionY gets a reference to the given float32 and assigns it to the DimensionY field.
func (o *ProductReadItem) SetDimensionY(v float32) {
	o.DimensionY = &v
}

// GetDimensionZ returns the DimensionZ field value if set, zero value otherwise.
func (o *ProductReadItem) GetDimensionZ() float32 {
	if o == nil || o.DimensionZ == nil {
		var ret float32
		return ret
	}
	return *o.DimensionZ
}

// GetDimensionZOk returns a tuple with the DimensionZ field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetDimensionZOk() (*float32, bool) {
	if o == nil || o.DimensionZ == nil {
		return nil, false
	}
	return o.DimensionZ, true
}

// HasDimensionZ returns a boolean if a field has been set.
func (o *ProductReadItem) HasDimensionZ() bool {
	if o != nil && o.DimensionZ != nil {
		return true
	}

	return false
}

// SetDimensionZ gets a reference to the given float32 and assigns it to the DimensionZ field.
func (o *ProductReadItem) SetDimensionZ(v float32) {
	o.DimensionZ = &v
}

// GetHsCode returns the HsCode field value if set, zero value otherwise.
func (o *ProductReadItem) GetHsCode() string {
	if o == nil || o.HsCode == nil {
		var ret string
		return ret
	}
	return *o.HsCode
}

// GetHsCodeOk returns a tuple with the HsCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetHsCodeOk() (*string, bool) {
	if o == nil || o.HsCode == nil {
		return nil, false
	}
	return o.HsCode, true
}

// HasHsCode returns a boolean if a field has been set.
func (o *ProductReadItem) HasHsCode() bool {
	if o != nil && o.HsCode != nil {
		return true
	}

	return false
}

// SetHsCode gets a reference to the given string and assigns it to the HsCode field.
func (o *ProductReadItem) SetHsCode(v string) {
	o.HsCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductReadItem) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductReadItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProductReadItem) SetId(v int32) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ProductReadItem) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ProductReadItem) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ProductReadItem) SetImage(v string) {
	o.Image = &v
}

// GetName returns the Name field value
func (o *ProductReadItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductReadItem) SetName(v string) {
	o.Name = v
}

// GetQuantityAvailable returns the QuantityAvailable field value if set, zero value otherwise.
func (o *ProductReadItem) GetQuantityAvailable() int32 {
	if o == nil || o.QuantityAvailable == nil {
		var ret int32
		return ret
	}
	return *o.QuantityAvailable
}

// GetQuantityAvailableOk returns a tuple with the QuantityAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetQuantityAvailableOk() (*int32, bool) {
	if o == nil || o.QuantityAvailable == nil {
		return nil, false
	}
	return o.QuantityAvailable, true
}

// HasQuantityAvailable returns a boolean if a field has been set.
func (o *ProductReadItem) HasQuantityAvailable() bool {
	if o != nil && o.QuantityAvailable != nil {
		return true
	}

	return false
}

// SetQuantityAvailable gets a reference to the given int32 and assigns it to the QuantityAvailable field.
func (o *ProductReadItem) SetQuantityAvailable(v int32) {
	o.QuantityAvailable = &v
}

// GetQuantityIncoming returns the QuantityIncoming field value if set, zero value otherwise.
func (o *ProductReadItem) GetQuantityIncoming() int32 {
	if o == nil || o.QuantityIncoming == nil {
		var ret int32
		return ret
	}
	return *o.QuantityIncoming
}

// GetQuantityIncomingOk returns a tuple with the QuantityIncoming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetQuantityIncomingOk() (*int32, bool) {
	if o == nil || o.QuantityIncoming == nil {
		return nil, false
	}
	return o.QuantityIncoming, true
}

// HasQuantityIncoming returns a boolean if a field has been set.
func (o *ProductReadItem) HasQuantityIncoming() bool {
	if o != nil && o.QuantityIncoming != nil {
		return true
	}

	return false
}

// SetQuantityIncoming gets a reference to the given int32 and assigns it to the QuantityIncoming field.
func (o *ProductReadItem) SetQuantityIncoming(v int32) {
	o.QuantityIncoming = &v
}

// GetQuantityOnHand returns the QuantityOnHand field value if set, zero value otherwise.
func (o *ProductReadItem) GetQuantityOnHand() int32 {
	if o == nil || o.QuantityOnHand == nil {
		var ret int32
		return ret
	}
	return *o.QuantityOnHand
}

// GetQuantityOnHandOk returns a tuple with the QuantityOnHand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetQuantityOnHandOk() (*int32, bool) {
	if o == nil || o.QuantityOnHand == nil {
		return nil, false
	}
	return o.QuantityOnHand, true
}

// HasQuantityOnHand returns a boolean if a field has been set.
func (o *ProductReadItem) HasQuantityOnHand() bool {
	if o != nil && o.QuantityOnHand != nil {
		return true
	}

	return false
}

// SetQuantityOnHand gets a reference to the given int32 and assigns it to the QuantityOnHand field.
func (o *ProductReadItem) SetQuantityOnHand(v int32) {
	o.QuantityOnHand = &v
}

// GetQuantityOutgoing returns the QuantityOutgoing field value if set, zero value otherwise.
func (o *ProductReadItem) GetQuantityOutgoing() int32 {
	if o == nil || o.QuantityOutgoing == nil {
		var ret int32
		return ret
	}
	return *o.QuantityOutgoing
}

// GetQuantityOutgoingOk returns a tuple with the QuantityOutgoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetQuantityOutgoingOk() (*int32, bool) {
	if o == nil || o.QuantityOutgoing == nil {
		return nil, false
	}
	return o.QuantityOutgoing, true
}

// HasQuantityOutgoing returns a boolean if a field has been set.
func (o *ProductReadItem) HasQuantityOutgoing() bool {
	if o != nil && o.QuantityOutgoing != nil {
		return true
	}

	return false
}

// SetQuantityOutgoing gets a reference to the given int32 and assigns it to the QuantityOutgoing field.
func (o *ProductReadItem) SetQuantityOutgoing(v int32) {
	o.QuantityOutgoing = &v
}

// GetQuantitySold returns the QuantitySold field value if set, zero value otherwise.
func (o *ProductReadItem) GetQuantitySold() int32 {
	if o == nil || o.QuantitySold == nil {
		var ret int32
		return ret
	}
	return *o.QuantitySold
}

// GetQuantitySoldOk returns a tuple with the QuantitySold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetQuantitySoldOk() (*int32, bool) {
	if o == nil || o.QuantitySold == nil {
		return nil, false
	}
	return o.QuantitySold, true
}

// HasQuantitySold returns a boolean if a field has been set.
func (o *ProductReadItem) HasQuantitySold() bool {
	if o != nil && o.QuantitySold != nil {
		return true
	}

	return false
}

// SetQuantitySold gets a reference to the given int32 and assigns it to the QuantitySold field.
func (o *ProductReadItem) SetQuantitySold(v int32) {
	o.QuantitySold = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ProductReadItem) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ProductReadItem) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ProductReadItem) SetReference(v string) {
	o.Reference = &v
}

// GetSellingPrice returns the SellingPrice field value if set, zero value otherwise.
func (o *ProductReadItem) GetSellingPrice() float32 {
	if o == nil || o.SellingPrice == nil {
		var ret float32
		return ret
	}
	return *o.SellingPrice
}

// GetSellingPriceOk returns a tuple with the SellingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetSellingPriceOk() (*float32, bool) {
	if o == nil || o.SellingPrice == nil {
		return nil, false
	}
	return o.SellingPrice, true
}

// HasSellingPrice returns a boolean if a field has been set.
func (o *ProductReadItem) HasSellingPrice() bool {
	if o != nil && o.SellingPrice != nil {
		return true
	}

	return false
}

// SetSellingPrice gets a reference to the given float32 and assigns it to the SellingPrice field.
func (o *ProductReadItem) SetSellingPrice(v float32) {
	o.SellingPrice = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductReadItem) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductReadItem) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductReadItem) SetSku(v string) {
	o.Sku = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ProductReadItem) GetVolume() float32 {
	if o == nil || o.Volume == nil {
		var ret float32
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetVolumeOk() (*float32, bool) {
	if o == nil || o.Volume == nil {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ProductReadItem) HasVolume() bool {
	if o != nil && o.Volume != nil {
		return true
	}

	return false
}

// SetVolume gets a reference to the given float32 and assigns it to the Volume field.
func (o *ProductReadItem) SetVolume(v float32) {
	o.Volume = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ProductReadItem) GetWeight() float32 {
	if o == nil || o.Weight == nil {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductReadItem) GetWeightOk() (*float32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ProductReadItem) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *ProductReadItem) SetWeight(v float32) {
	o.Weight = &v
}

func (o ProductReadItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["barcode"] = o.Barcode
	}
	if o.ChannelReference != nil {
		toSerialize["channelReference"] = o.ChannelReference
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.QuantityAvailable != nil {
		toSerialize["quantityAvailable"] = o.QuantityAvailable
	}
	if o.QuantityIncoming != nil {
		toSerialize["quantityIncoming"] = o.QuantityIncoming
	}
	if o.QuantityOnHand != nil {
		toSerialize["quantityOnHand"] = o.QuantityOnHand
	}
	if o.QuantityOutgoing != nil {
		toSerialize["quantityOutgoing"] = o.QuantityOutgoing
	}
	if o.QuantitySold != nil {
		toSerialize["quantitySold"] = o.QuantitySold
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.SellingPrice != nil {
		toSerialize["sellingPrice"] = o.SellingPrice
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.Volume != nil {
		toSerialize["volume"] = o.Volume
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableProductReadItem struct {
	value *ProductReadItem
	isSet bool
}

func (v NullableProductReadItem) Get() *ProductReadItem {
	return v.value
}

func (v *NullableProductReadItem) Set(val *ProductReadItem) {
	v.value = val
	v.isSet = true
}

func (v NullableProductReadItem) IsSet() bool {
	return v.isSet
}

func (v *NullableProductReadItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductReadItem(val *ProductReadItem) *NullableProductReadItem {
	return &NullableProductReadItem{value: val, isSet: true}
}

func (v NullableProductReadItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductReadItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


