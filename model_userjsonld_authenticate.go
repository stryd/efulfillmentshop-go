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

// UserjsonldAuthenticate struct for UserjsonldAuthenticate
type UserjsonldAuthenticate struct {
	JsonContext *string `json:"json_context,omitempty"`
	JsonId *string `json:"json_id,omitempty"`
	JsonType *string `json:"json_type,omitempty"`
	// The user email address
	Email string `json:"email"`
	// The user password
	Password string `json:"password"`
}

// NewUserjsonldAuthenticate instantiates a new UserjsonldAuthenticate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserjsonldAuthenticate(email string, password string, ) *UserjsonldAuthenticate {
	this := UserjsonldAuthenticate{}
	this.Email = email
	this.Password = password
	return &this
}

// NewUserjsonldAuthenticateWithDefaults instantiates a new UserjsonldAuthenticate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserjsonldAuthenticateWithDefaults() *UserjsonldAuthenticate {
	this := UserjsonldAuthenticate{}
	return &this
}

// GetJsonContext returns the JsonContext field value if set, zero value otherwise.
func (o *UserjsonldAuthenticate) GetJsonContext() string {
	if o == nil || o.JsonContext == nil {
		var ret string
		return ret
	}
	return *o.JsonContext
}

// GetJsonContextOk returns a tuple with the JsonContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserjsonldAuthenticate) GetJsonContextOk() (*string, bool) {
	if o == nil || o.JsonContext == nil {
		return nil, false
	}
	return o.JsonContext, true
}

// HasJsonContext returns a boolean if a field has been set.
func (o *UserjsonldAuthenticate) HasJsonContext() bool {
	if o != nil && o.JsonContext != nil {
		return true
	}

	return false
}

// SetJsonContext gets a reference to the given string and assigns it to the JsonContext field.
func (o *UserjsonldAuthenticate) SetJsonContext(v string) {
	o.JsonContext = &v
}

// GetJsonId returns the JsonId field value if set, zero value otherwise.
func (o *UserjsonldAuthenticate) GetJsonId() string {
	if o == nil || o.JsonId == nil {
		var ret string
		return ret
	}
	return *o.JsonId
}

// GetJsonIdOk returns a tuple with the JsonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserjsonldAuthenticate) GetJsonIdOk() (*string, bool) {
	if o == nil || o.JsonId == nil {
		return nil, false
	}
	return o.JsonId, true
}

// HasJsonId returns a boolean if a field has been set.
func (o *UserjsonldAuthenticate) HasJsonId() bool {
	if o != nil && o.JsonId != nil {
		return true
	}

	return false
}

// SetJsonId gets a reference to the given string and assigns it to the JsonId field.
func (o *UserjsonldAuthenticate) SetJsonId(v string) {
	o.JsonId = &v
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *UserjsonldAuthenticate) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserjsonldAuthenticate) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *UserjsonldAuthenticate) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *UserjsonldAuthenticate) SetJsonType(v string) {
	o.JsonType = &v
}

// GetEmail returns the Email field value
func (o *UserjsonldAuthenticate) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserjsonldAuthenticate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserjsonldAuthenticate) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *UserjsonldAuthenticate) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserjsonldAuthenticate) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserjsonldAuthenticate) SetPassword(v string) {
	o.Password = v
}

func (o UserjsonldAuthenticate) MarshalJSON() ([]byte, error) {
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
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserjsonldAuthenticate struct {
	value *UserjsonldAuthenticate
	isSet bool
}

func (v NullableUserjsonldAuthenticate) Get() *UserjsonldAuthenticate {
	return v.value
}

func (v *NullableUserjsonldAuthenticate) Set(val *UserjsonldAuthenticate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserjsonldAuthenticate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserjsonldAuthenticate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserjsonldAuthenticate(val *UserjsonldAuthenticate) *NullableUserjsonldAuthenticate {
	return &NullableUserjsonldAuthenticate{value: val, isSet: true}
}

func (v NullableUserjsonldAuthenticate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserjsonldAuthenticate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


