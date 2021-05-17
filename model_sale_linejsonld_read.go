/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.3.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
)

// SaleLinejsonldRead Sale Line entity
type SaleLinejsonldRead struct {
	JsonContext *string `json:"json_context,omitempty"`
	JsonId *string `json:"json_id,omitempty"`
	JsonType *string `json:"json_type,omitempty"`
	// The sale line description
	Description string `json:"description"`
	// The sale line ID
	Id *int32 `json:"id,omitempty"`
	// The product ID
	ProductId int32 `json:"productId"`
	// The sale line quantity
	Quantity int32 `json:"quantity"`
	// The sale line quantity delivered
	QuantityDelivered *int32 `json:"quantityDelivered,omitempty"`
	// The sale ID
	SaleId int32 `json:"saleId"`
}

// NewSaleLinejsonldRead instantiates a new SaleLinejsonldRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaleLinejsonldRead(description string, productId int32, quantity int32, saleId int32, ) *SaleLinejsonldRead {
	this := SaleLinejsonldRead{}
	this.Description = description
	this.ProductId = productId
	this.Quantity = quantity
	this.SaleId = saleId
	return &this
}

// NewSaleLinejsonldReadWithDefaults instantiates a new SaleLinejsonldRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaleLinejsonldReadWithDefaults() *SaleLinejsonldRead {
	this := SaleLinejsonldRead{}
	return &this
}

// GetJsonContext returns the JsonContext field value if set, zero value otherwise.
func (o *SaleLinejsonldRead) GetJsonContext() string {
	if o == nil || o.JsonContext == nil {
		var ret string
		return ret
	}
	return *o.JsonContext
}

// GetJsonContextOk returns a tuple with the JsonContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetJsonContextOk() (*string, bool) {
	if o == nil || o.JsonContext == nil {
		return nil, false
	}
	return o.JsonContext, true
}

// HasJsonContext returns a boolean if a field has been set.
func (o *SaleLinejsonldRead) HasJsonContext() bool {
	if o != nil && o.JsonContext != nil {
		return true
	}

	return false
}

// SetJsonContext gets a reference to the given string and assigns it to the JsonContext field.
func (o *SaleLinejsonldRead) SetJsonContext(v string) {
	o.JsonContext = &v
}

// GetJsonId returns the JsonId field value if set, zero value otherwise.
func (o *SaleLinejsonldRead) GetJsonId() string {
	if o == nil || o.JsonId == nil {
		var ret string
		return ret
	}
	return *o.JsonId
}

// GetJsonIdOk returns a tuple with the JsonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetJsonIdOk() (*string, bool) {
	if o == nil || o.JsonId == nil {
		return nil, false
	}
	return o.JsonId, true
}

// HasJsonId returns a boolean if a field has been set.
func (o *SaleLinejsonldRead) HasJsonId() bool {
	if o != nil && o.JsonId != nil {
		return true
	}

	return false
}

// SetJsonId gets a reference to the given string and assigns it to the JsonId field.
func (o *SaleLinejsonldRead) SetJsonId(v string) {
	o.JsonId = &v
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *SaleLinejsonldRead) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *SaleLinejsonldRead) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *SaleLinejsonldRead) SetJsonType(v string) {
	o.JsonType = &v
}

// GetDescription returns the Description field value
func (o *SaleLinejsonldRead) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SaleLinejsonldRead) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SaleLinejsonldRead) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SaleLinejsonldRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SaleLinejsonldRead) SetId(v int32) {
	o.Id = &v
}

// GetProductId returns the ProductId field value
func (o *SaleLinejsonldRead) GetProductId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetProductIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *SaleLinejsonldRead) SetProductId(v int32) {
	o.ProductId = v
}

// GetQuantity returns the Quantity field value
func (o *SaleLinejsonldRead) GetQuantity() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetQuantityOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *SaleLinejsonldRead) SetQuantity(v int32) {
	o.Quantity = v
}

// GetQuantityDelivered returns the QuantityDelivered field value if set, zero value otherwise.
func (o *SaleLinejsonldRead) GetQuantityDelivered() int32 {
	if o == nil || o.QuantityDelivered == nil {
		var ret int32
		return ret
	}
	return *o.QuantityDelivered
}

// GetQuantityDeliveredOk returns a tuple with the QuantityDelivered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetQuantityDeliveredOk() (*int32, bool) {
	if o == nil || o.QuantityDelivered == nil {
		return nil, false
	}
	return o.QuantityDelivered, true
}

// HasQuantityDelivered returns a boolean if a field has been set.
func (o *SaleLinejsonldRead) HasQuantityDelivered() bool {
	if o != nil && o.QuantityDelivered != nil {
		return true
	}

	return false
}

// SetQuantityDelivered gets a reference to the given int32 and assigns it to the QuantityDelivered field.
func (o *SaleLinejsonldRead) SetQuantityDelivered(v int32) {
	o.QuantityDelivered = &v
}

// GetSaleId returns the SaleId field value
func (o *SaleLinejsonldRead) GetSaleId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.SaleId
}

// GetSaleIdOk returns a tuple with the SaleId field value
// and a boolean to check if the value has been set.
func (o *SaleLinejsonldRead) GetSaleIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SaleId, true
}

// SetSaleId sets field value
func (o *SaleLinejsonldRead) SetSaleId(v int32) {
	o.SaleId = v
}

func (o SaleLinejsonldRead) MarshalJSON() ([]byte, error) {
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
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["productId"] = o.ProductId
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if o.QuantityDelivered != nil {
		toSerialize["quantityDelivered"] = o.QuantityDelivered
	}
	if true {
		toSerialize["saleId"] = o.SaleId
	}
	return json.Marshal(toSerialize)
}

type NullableSaleLinejsonldRead struct {
	value *SaleLinejsonldRead
	isSet bool
}

func (v NullableSaleLinejsonldRead) Get() *SaleLinejsonldRead {
	return v.value
}

func (v *NullableSaleLinejsonldRead) Set(val *SaleLinejsonldRead) {
	v.value = val
	v.isSet = true
}

func (v NullableSaleLinejsonldRead) IsSet() bool {
	return v.isSet
}

func (v *NullableSaleLinejsonldRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaleLinejsonldRead(val *SaleLinejsonldRead) *NullableSaleLinejsonldRead {
	return &NullableSaleLinejsonldRead{value: val, isSet: true}
}

func (v NullableSaleLinejsonldRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaleLinejsonldRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


