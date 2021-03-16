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

// ChanneljsonldRead Channel entity
type ChanneljsonldRead struct {
	JSONContext *string `json:"@context,omitempty"`
	JSONId *string `json:"@id,omitempty"`
	JSONType *string `json:"@type,omitempty"`
	// The channel ID
	Id *int32 `json:"id,omitempty"`
	// The channel name
	Name *string `json:"name,omitempty"`
}

// NewChanneljsonldRead instantiates a new ChanneljsonldRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChanneljsonldRead() *ChanneljsonldRead {
	this := ChanneljsonldRead{}
	return &this
}

// NewChanneljsonldReadWithDefaults instantiates a new ChanneljsonldRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChanneljsonldReadWithDefaults() *ChanneljsonldRead {
	this := ChanneljsonldRead{}
	return &this
}

// GetJSONContext returns the JSONContext field value if set, zero value otherwise.
func (o *ChanneljsonldRead) GetJSONContext() string {
	if o == nil || o.JSONContext == nil {
		var ret string
		return ret
	}
	return *o.JSONContext
}

// GetJSONContextOk returns a tuple with the JSONContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChanneljsonldRead) GetJSONContextOk() (*string, bool) {
	if o == nil || o.JSONContext == nil {
		return nil, false
	}
	return o.JSONContext, true
}

// HasJSONContext returns a boolean if a field has been set.
func (o *ChanneljsonldRead) HasJSONContext() bool {
	if o != nil && o.JSONContext != nil {
		return true
	}

	return false
}

// SetJSONContext gets a reference to the given string and assigns it to the JSONContext field.
func (o *ChanneljsonldRead) SetJSONContext(v string) {
	o.JSONContext = &v
}

// GetJSONId returns the JSONId field value if set, zero value otherwise.
func (o *ChanneljsonldRead) GetJSONId() string {
	if o == nil || o.JSONId == nil {
		var ret string
		return ret
	}
	return *o.JSONId
}

// GetJSONIdOk returns a tuple with the JSONId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChanneljsonldRead) GetJSONIdOk() (*string, bool) {
	if o == nil || o.JSONId == nil {
		return nil, false
	}
	return o.JSONId, true
}

// HasJSONId returns a boolean if a field has been set.
func (o *ChanneljsonldRead) HasJSONId() bool {
	if o != nil && o.JSONId != nil {
		return true
	}

	return false
}

// SetJSONId gets a reference to the given string and assigns it to the JSONId field.
func (o *ChanneljsonldRead) SetJSONId(v string) {
	o.JSONId = &v
}

// GetJSONType returns the JSONType field value if set, zero value otherwise.
func (o *ChanneljsonldRead) GetJSONType() string {
	if o == nil || o.JSONType == nil {
		var ret string
		return ret
	}
	return *o.JSONType
}

// GetJSONTypeOk returns a tuple with the JSONType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChanneljsonldRead) GetJSONTypeOk() (*string, bool) {
	if o == nil || o.JSONType == nil {
		return nil, false
	}
	return o.JSONType, true
}

// HasJSONType returns a boolean if a field has been set.
func (o *ChanneljsonldRead) HasJSONType() bool {
	if o != nil && o.JSONType != nil {
		return true
	}

	return false
}

// SetJSONType gets a reference to the given string and assigns it to the JSONType field.
func (o *ChanneljsonldRead) SetJSONType(v string) {
	o.JSONType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChanneljsonldRead) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChanneljsonldRead) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChanneljsonldRead) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ChanneljsonldRead) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChanneljsonldRead) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChanneljsonldRead) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChanneljsonldRead) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChanneljsonldRead) SetName(v string) {
	o.Name = &v
}

func (o ChanneljsonldRead) MarshalJSON() ([]byte, error) {
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableChanneljsonldRead struct {
	value *ChanneljsonldRead
	isSet bool
}

func (v NullableChanneljsonldRead) Get() *ChanneljsonldRead {
	return v.value
}

func (v *NullableChanneljsonldRead) Set(val *ChanneljsonldRead) {
	v.value = val
	v.isSet = true
}

func (v NullableChanneljsonldRead) IsSet() bool {
	return v.isSet
}

func (v *NullableChanneljsonldRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChanneljsonldRead(val *ChanneljsonldRead) *NullableChanneljsonldRead {
	return &NullableChanneljsonldRead{value: val, isSet: true}
}

func (v NullableChanneljsonldRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChanneljsonldRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


