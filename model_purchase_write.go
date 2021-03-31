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
	"time"
)

// PurchaseWrite Purchase entity
type PurchaseWrite struct {
	// The purchase channel
	ChannelId int32 `json:"channelId"`
	// The purchase channel reference
	ChannelReference string `json:"channelReference"`
	// The purchase planned delivery date
	PlannedDate time.Time `json:"plannedDate"`
	// The purchase supplier ID
	SupplierId int32 `json:"supplierId"`
}

// NewPurchaseWrite instantiates a new PurchaseWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseWrite(channelId int32, channelReference string, plannedDate time.Time, supplierId int32, ) *PurchaseWrite {
	this := PurchaseWrite{}
	this.ChannelId = channelId
	this.ChannelReference = channelReference
	this.PlannedDate = plannedDate
	this.SupplierId = supplierId
	return &this
}

// NewPurchaseWriteWithDefaults instantiates a new PurchaseWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseWriteWithDefaults() *PurchaseWrite {
	this := PurchaseWrite{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *PurchaseWrite) GetChannelId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *PurchaseWrite) GetChannelIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *PurchaseWrite) SetChannelId(v int32) {
	o.ChannelId = v
}

// GetChannelReference returns the ChannelReference field value
func (o *PurchaseWrite) GetChannelReference() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value
// and a boolean to check if the value has been set.
func (o *PurchaseWrite) GetChannelReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelReference, true
}

// SetChannelReference sets field value
func (o *PurchaseWrite) SetChannelReference(v string) {
	o.ChannelReference = v
}

// GetPlannedDate returns the PlannedDate field value
func (o *PurchaseWrite) GetPlannedDate() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.PlannedDate
}

// GetPlannedDateOk returns a tuple with the PlannedDate field value
// and a boolean to check if the value has been set.
func (o *PurchaseWrite) GetPlannedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlannedDate, true
}

// SetPlannedDate sets field value
func (o *PurchaseWrite) SetPlannedDate(v time.Time) {
	o.PlannedDate = v
}

// GetSupplierId returns the SupplierId field value
func (o *PurchaseWrite) GetSupplierId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.SupplierId
}

// GetSupplierIdOk returns a tuple with the SupplierId field value
// and a boolean to check if the value has been set.
func (o *PurchaseWrite) GetSupplierIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupplierId, true
}

// SetSupplierId sets field value
func (o *PurchaseWrite) SetSupplierId(v int32) {
	o.SupplierId = v
}

func (o PurchaseWrite) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["channelId"] = o.ChannelId
	}
	if true {
		toSerialize["channelReference"] = o.ChannelReference
	}
	if true {
		toSerialize["plannedDate"] = o.PlannedDate
	}
	if true {
		toSerialize["supplierId"] = o.SupplierId
	}
	return json.Marshal(toSerialize)
}

type NullablePurchaseWrite struct {
	value *PurchaseWrite
	isSet bool
}

func (v NullablePurchaseWrite) Get() *PurchaseWrite {
	return v.value
}

func (v *NullablePurchaseWrite) Set(val *PurchaseWrite) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseWrite) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseWrite(val *PurchaseWrite) *NullablePurchaseWrite {
	return &NullablePurchaseWrite{value: val, isSet: true}
}

func (v NullablePurchaseWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePurchaseWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


