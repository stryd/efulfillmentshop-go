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

// UserAuthenticate struct for UserAuthenticate
type UserAuthenticate struct {
	// The user email address
	Email string `json:"email"`
	// The user password
	Password string `json:"password"`
}

// NewUserAuthenticate instantiates a new UserAuthenticate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuthenticate(email string, password string, ) *UserAuthenticate {
	this := UserAuthenticate{}
	this.Email = email
	this.Password = password
	return &this
}

// NewUserAuthenticateWithDefaults instantiates a new UserAuthenticate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthenticateWithDefaults() *UserAuthenticate {
	this := UserAuthenticate{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserAuthenticate) GetEmail() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserAuthenticate) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserAuthenticate) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *UserAuthenticate) GetPassword() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserAuthenticate) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserAuthenticate) SetPassword(v string) {
	o.Password = v
}

func (o UserAuthenticate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUserAuthenticate struct {
	value *UserAuthenticate
	isSet bool
}

func (v NullableUserAuthenticate) Get() *UserAuthenticate {
	return v.value
}

func (v *NullableUserAuthenticate) Set(val *UserAuthenticate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthenticate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthenticate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthenticate(val *UserAuthenticate) *NullableUserAuthenticate {
	return &NullableUserAuthenticate{value: val, isSet: true}
}

func (v NullableUserAuthenticate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthenticate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


