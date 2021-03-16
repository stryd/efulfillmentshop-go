/*
 * E-fulfilment Shop API
 *
 * The E-fulfilment Shop API allows you to integrate your service with our warehouse.
 *
 * API version: 1.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package efulfillmentshop

import (
	"encoding/json"
)

// SaleRead Sale entity
type SaleRead struct {
	// The sale carrier ID
	CarrierId *int32 `json:"carrierId,omitempty"`
	// The sale channel
	ChannelId int32 `json:"channelId"`
	// The sale channel reference
	ChannelReference string `json:"channelReference"`
	// The sale ID
	Id *int32 `json:"id,omitempty"`
	// The sale invoice address ID
	InvoiceAddressId int32 `json:"invoiceAddressId"`
	// The sale line IDs
	LineIds *[]int32 `json:"lineIds,omitempty"`
	// The sale name
	Name *string `json:"name,omitempty"`
	// The sale shipment IDs
	ShipmentIds *[]int32 `json:"shipmentIds,omitempty"`
	// The sale shipping address ID
	ShippingAddressId int32 `json:"shippingAddressId"`
	// The sale warehouse status
	Status *string `json:"status,omitempty"`
}

// NewSaleRead instantiates a new SaleRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaleRead(channelId int32, channelReference string, invoiceAddressId int32, shippingAddressId int32, ) *SaleRead {
	this := SaleRead{}
	this.ChannelId = channelId
	this.ChannelReference = channelReference
	this.InvoiceAddressId = invoiceAddressId
	this.ShippingAddressId = shippingAddressId
	return &this
}

// NewSaleReadWithDefaults instantiates a new SaleRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaleReadWithDefaults() *SaleRead {
	this := SaleRead{}
	return &this
}

// GetCarrierId returns the CarrierId field value if set, zero value otherwise.
func (o *SaleRead) GetCarrierId() int32 {
	if o == nil || o.CarrierId == nil {
		var ret int32
		return ret
	}
	return *o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleRead) GetCarrierIdOk() (*int32, bool) {
	if o == nil || o.CarrierId == nil {
		return nil, false
	}
	return o.CarrierId, true
}

// HasCarrierId returns a boolean if a field has been set.
func (o *SaleRead) HasCarrierId() bool {
	if o != nil && o.CarrierId != nil {
		return true
	}

	return false
}

// SetCarrierId gets a reference to the given int32 and assigns it to the CarrierId field.
func (o *SaleRead) SetCarrierId(v int32) {
	o.CarrierId = &v
}

// GetChannelId returns the ChannelId field value
func (o *SaleRead) GetChannelId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *SaleRead) GetChannelIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *SaleRead) SetChannelId(v int32) {
	o.ChannelId = v
}

// GetChannelReference returns the ChannelReference field value
func (o *SaleRead) GetChannelReference() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ChannelReference
}

// GetChannelReferenceOk returns a tuple with the ChannelReference field value
// and a boolean to check if the value has been set.
func (o *SaleRead) GetChannelReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ChannelReference, true
}

// SetChannelReference sets field value
func (o *SaleRead) SetChannelReference(v string) {
	o.ChannelReference = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SaleRead) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleRead) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SaleRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SaleRead) SetId(v int32) {
	o.Id = &v
}

// GetInvoiceAddressId returns the InvoiceAddressId field value
func (o *SaleRead) GetInvoiceAddressId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.InvoiceAddressId
}

// GetInvoiceAddressIdOk returns a tuple with the InvoiceAddressId field value
// and a boolean to check if the value has been set.
func (o *SaleRead) GetInvoiceAddressIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InvoiceAddressId, true
}

// SetInvoiceAddressId sets field value
func (o *SaleRead) SetInvoiceAddressId(v int32) {
	o.InvoiceAddressId = v
}

// GetLineIds returns the LineIds field value if set, zero value otherwise.
func (o *SaleRead) GetLineIds() []int32 {
	if o == nil || o.LineIds == nil {
		var ret []int32
		return ret
	}
	return *o.LineIds
}

// GetLineIdsOk returns a tuple with the LineIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleRead) GetLineIdsOk() (*[]int32, bool) {
	if o == nil || o.LineIds == nil {
		return nil, false
	}
	return o.LineIds, true
}

// HasLineIds returns a boolean if a field has been set.
func (o *SaleRead) HasLineIds() bool {
	if o != nil && o.LineIds != nil {
		return true
	}

	return false
}

// SetLineIds gets a reference to the given []int32 and assigns it to the LineIds field.
func (o *SaleRead) SetLineIds(v []int32) {
	o.LineIds = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SaleRead) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleRead) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SaleRead) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SaleRead) SetName(v string) {
	o.Name = &v
}

// GetShipmentIds returns the ShipmentIds field value if set, zero value otherwise.
func (o *SaleRead) GetShipmentIds() []int32 {
	if o == nil || o.ShipmentIds == nil {
		var ret []int32
		return ret
	}
	return *o.ShipmentIds
}

// GetShipmentIdsOk returns a tuple with the ShipmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleRead) GetShipmentIdsOk() (*[]int32, bool) {
	if o == nil || o.ShipmentIds == nil {
		return nil, false
	}
	return o.ShipmentIds, true
}

// HasShipmentIds returns a boolean if a field has been set.
func (o *SaleRead) HasShipmentIds() bool {
	if o != nil && o.ShipmentIds != nil {
		return true
	}

	return false
}

// SetShipmentIds gets a reference to the given []int32 and assigns it to the ShipmentIds field.
func (o *SaleRead) SetShipmentIds(v []int32) {
	o.ShipmentIds = &v
}

// GetShippingAddressId returns the ShippingAddressId field value
func (o *SaleRead) GetShippingAddressId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ShippingAddressId
}

// GetShippingAddressIdOk returns a tuple with the ShippingAddressId field value
// and a boolean to check if the value has been set.
func (o *SaleRead) GetShippingAddressIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ShippingAddressId, true
}

// SetShippingAddressId sets field value
func (o *SaleRead) SetShippingAddressId(v int32) {
	o.ShippingAddressId = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SaleRead) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaleRead) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SaleRead) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SaleRead) SetStatus(v string) {
	o.Status = &v
}

func (o SaleRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CarrierId != nil {
		toSerialize["carrierId"] = o.CarrierId
	}
	if true {
		toSerialize["channelId"] = o.ChannelId
	}
	if true {
		toSerialize["channelReference"] = o.ChannelReference
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["invoiceAddressId"] = o.InvoiceAddressId
	}
	if o.LineIds != nil {
		toSerialize["lineIds"] = o.LineIds
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ShipmentIds != nil {
		toSerialize["shipmentIds"] = o.ShipmentIds
	}
	if true {
		toSerialize["shippingAddressId"] = o.ShippingAddressId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSaleRead struct {
	value *SaleRead
	isSet bool
}

func (v NullableSaleRead) Get() *SaleRead {
	return v.value
}

func (v *NullableSaleRead) Set(val *SaleRead) {
	v.value = val
	v.isSet = true
}

func (v NullableSaleRead) IsSet() bool {
	return v.isSet
}

func (v *NullableSaleRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaleRead(val *SaleRead) *NullableSaleRead {
	return &NullableSaleRead{value: val, isSet: true}
}

func (v NullableSaleRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaleRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


