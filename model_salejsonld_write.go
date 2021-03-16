/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfilmentshop

import (
	"encoding/json"
)

// SalejsonldWrite Sale entity
type SalejsonldWrite struct {
	JSONContext *string `json:"@context,omitempty"`
	JSONId *string `json:"@id,omitempty"`
	JSONType *string `json:"@type,omitempty"`
	// The sale carrier ID
	CarrierId *int32 `json:"carrierId,omitempty"`
	// The sale channel
	ChannelId int32 `json:"channelId"`
	// The sale channel reference
	ChannelReference string `json:"channelReference"`
	// The sale invoice address ID
	InvoiceAddressId int32 `json:"invoiceAddressId"`
	// The sale shipping address ID
	ShippingAddressId int32 `json:"shippingAddressId"`
}

// NewSalejsonldWrite instantiates a new SalejsonldWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSalejsonldWrite(channelId int32, channelReference string, invoiceAddressId int32, shippingAddressId int32, ) *SalejsonldWrite {
	this := SalejsonldWrite{}
	this.ChannelId = channelId
	this.ChannelReference = channelReference
	this.InvoiceAddressId = invoiceAddressId
	this.ShippingAddressId = shippingAddressId
	return &this
}

// NewSalejsonldWriteWithDefaults instantiates a new SalejsonldWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSalejsonldWriteWithDefaults() *SalejsonldWrite {
	this := SalejsonldWrite{}
	return &this
}

// GetJSONContext returns the JSONContext field value if set, zero value otherwise.
func (o *SalejsonldWrite) GetJSONContext() string {
	if o == nil || o.JSONContext == nil {
		var ret string
		return ret
	}
	return *o.JSONContext
}

// GetJSONContextOk returns a tuple with the JSONContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetJSONContextOk() (*string, bool) {
	if o == nil || o.JSONContext == nil {
		return nil, false
	}
	return o.JSONContext, true
}

// HasJSONContext returns a boolean if a field has been set.
func (o *SalejsonldWrite) HasJSONContext() bool {
	if o != nil && o.JSONContext != nil {
		return true
	}

	return false
}

// SetJSONContext gets a reference to the given string and assigns it to the JSONContext field.
func (o *SalejsonldWrite) SetJSONContext(v string) {
	o.JSONContext = &v
}

// GetJSONId returns the JSONId field value if set, zero value otherwise.
func (o *SalejsonldWrite) GetJSONId() string {
	if o == nil || o.JSONId == nil {
		var ret string
		return ret
	}
	return *o.JSONId
}

// GetJSONIdOk returns a tuple with the JSONId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetJSONIdOk() (*string, bool) {
	if o == nil || o.JSONId == nil {
		return nil, false
	}
	return o.JSONId, true
}

// HasJSONId returns a boolean if a field has been set.
func (o *SalejsonldWrite) HasJSONId() bool {
	if o != nil && o.JSONId != nil {
		return true
	}

	return false
}

// SetJSONId gets a reference to the given string and assigns it to the JSONId field.
func (o *SalejsonldWrite) SetJSONId(v string) {
	o.JSONId = &v
}

// GetJSONType returns the JSONType field value if set, zero value otherwise.
func (o *SalejsonldWrite) GetJSONType() string {
	if o == nil || o.JSONType == nil {
		var ret string
		return ret
	}
	return *o.JSONType
}

// GetJSONTypeOk returns a tuple with the JSONType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetJSONTypeOk() (*string, bool) {
	if o == nil || o.JSONType == nil {
		return nil, false
	}
	return o.JSONType, true
}

// HasJSONType returns a boolean if a field has been set.
func (o *SalejsonldWrite) HasJSONType() bool {
	if o != nil && o.JSONType != nil {
		return true
	}

	return false
}

// SetJSONType gets a reference to the given string and assigns it to the JSONType field.
func (o *SalejsonldWrite) SetJSONType(v string) {
	o.JSONType = &v
}

// GetCarrierId returns the CarrierId field value if set, zero value otherwise.
func (o *SalejsonldWrite) GetCarrierId() int32 {
	if o == nil || o.CarrierId == nil {
		var ret int32
		return ret
	}
	return *o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetCarrierIdOk() (*int32, bool) {
	if o == nil || o.CarrierId == nil {
		return nil, false
	}
	return o.CarrierId, true
}

// HasCarrierId returns a boolean if a field has been set.
func (o *SalejsonldWrite) HasCarrierId() bool {
	if o != nil && o.CarrierId != nil {
		return true
	}

	return false
}

// SetCarrierId gets a reference to the given int32 and assigns it to the CarrierId field.
func (o *SalejsonldWrite) SetCarrierId(v int32) {
	o.CarrierId = &v
}

// GetChannelId returns the ChannelId field value
func (o *SalejsonldWrite) GetChannelId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetChannelIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *SalejsonldWrite) SetChannelId(v int32) {
	o.ChannelId = v
}

// GetChannelReference returns the ChannelReference field value
func (o *SalejsonldWrite) GetChannelReference() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetChannelReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelReference, true
}

// SetChannelReference sets field value
func (o *SalejsonldWrite) SetChannelReference(v string) {
	o.ChannelReference = v
}

// GetInvoiceAddressId returns the InvoiceAddressId field value
func (o *SalejsonldWrite) GetInvoiceAddressId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.InvoiceAddressId
}

// GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field value
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetInvoiceAddressIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InvoiceAddressId, true
}

// SetInvoiceAddressId sets field value
func (o *SalejsonldWrite) SetInvoiceAddressId(v int32) {
	o.InvoiceAddressId = v
}

// GetShippingAddressId returns the ShippingAddressId field value
func (o *SalejsonldWrite) GetShippingAddressId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ShippingAddressId
}

// GetShippingAddressIdOk returns a tuple with the ShippingAddressId field value
// and a boolean to check if the value has been set.
func (o *SalejsonldWrite) GetShippingAddressIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ShippingAddressId, true
}

// SetShippingAddressId sets field value
func (o *SalejsonldWrite) SetShippingAddressId(v int32) {
	o.ShippingAddressId = v
}

func (o SalejsonldWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JSONContext != nil {
		toSerialize["@context"] = o.JSONContext
	}
	if o.JSONId != nil {
		toSerialize["@id"] = o.JSONId
	}
	if o.JSONType != nil {
		toSerialize["@type"] = o.JSONType
	}
	if o.CarrierId != nil {
		toSerialize["carrierId"] = o.CarrierId
	}
	if true {
		toSerialize["channelId"] = o.ChannelId
	}
	if true {
		toSerialize["channelReference"] = o.ChannelReference
	}
	if true {
		toSerialize["invoiceAddressId"] = o.InvoiceAddressId
	}
	if true {
		toSerialize["shippingAddressId"] = o.ShippingAddressId
	}
	return json.Marshal(toSerialize)
}

type NullableSalejsonldWrite struct {
	value *SalejsonldWrite
	isSet bool
}

func (v NullableSalejsonldWrite) Get() *SalejsonldWrite {
	return v.value
}

func (v *NullableSalejsonldWrite) Set(val *SalejsonldWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableSalejsonldWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableSalejsonldWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSalejsonldWrite(val *SalejsonldWrite) *NullableSalejsonldWrite {
	return &NullableSalejsonldWrite{value: val, isSet: true}
}

func (v NullableSalejsonldWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSalejsonldWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


