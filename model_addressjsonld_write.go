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

// AddressjsonldWrite Address entity
type AddressjsonldWrite struct {
	JsonContext *string `json:"json_context,omitempty"`
	JsonId *string `json:"json_id,omitempty"`
	JsonType *string `json:"json_type,omitempty"`
	// The address city
	City string `json:"city"`
	// The address company
	Company *string `json:"company,omitempty"`
	// The address country code (ISO 3166-1 alpha-2 format)
	CountryCode string `json:"countryCode"`
	// The address email
	Email *string `json:"email,omitempty"`
	// The address house number
	HouseNumber *string `json:"houseNumber,omitempty"`
	// The address name
	Name string `json:"name"`
	// The address phone
	Phone *string `json:"phone,omitempty"`
	// The address street
	Street string `json:"street"`
	// The address street 2
	Street2 *string `json:"street2,omitempty"`
	// The address zip
	Zip string `json:"zip"`
}

// NewAddressjsonldWrite instantiates a new AddressjsonldWrite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressjsonldWrite(city string, countryCode string, name string, street string, zip string, ) *AddressjsonldWrite {
	this := AddressjsonldWrite{}
	this.City = city
	this.CountryCode = countryCode
	this.Name = name
	this.Street = street
	this.Zip = zip
	return &this
}

// NewAddressjsonldWriteWithDefaults instantiates a new AddressjsonldWrite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressjsonldWriteWithDefaults() *AddressjsonldWrite {
	this := AddressjsonldWrite{}
	return &this
}

// GetJsonContext returns the JsonContext field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetJsonContext() string {
	if o == nil || o.JsonContext == nil {
		var ret string
		return ret
	}
	return *o.JsonContext
}

// GetJsonContextOk returns a tuple with the JsonContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetJsonContextOk() (*string, bool) {
	if o == nil || o.JsonContext == nil {
		return nil, false
	}
	return o.JsonContext, true
}

// HasJsonContext returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasJsonContext() bool {
	if o != nil && o.JsonContext != nil {
		return true
	}

	return false
}

// SetJsonContext gets a reference to the given string and assigns it to the JsonContext field.
func (o *AddressjsonldWrite) SetJsonContext(v string) {
	o.JsonContext = &v
}

// GetJsonId returns the JsonId field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetJsonId() string {
	if o == nil || o.JsonId == nil {
		var ret string
		return ret
	}
	return *o.JsonId
}

// GetJsonIdOk returns a tuple with the JsonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetJsonIdOk() (*string, bool) {
	if o == nil || o.JsonId == nil {
		return nil, false
	}
	return o.JsonId, true
}

// HasJsonId returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasJsonId() bool {
	if o != nil && o.JsonId != nil {
		return true
	}

	return false
}

// SetJsonId gets a reference to the given string and assigns it to the JsonId field.
func (o *AddressjsonldWrite) SetJsonId(v string) {
	o.JsonId = &v
}

// GetJsonType returns the JsonType field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetJsonType() string {
	if o == nil || o.JsonType == nil {
		var ret string
		return ret
	}
	return *o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetJsonTypeOk() (*string, bool) {
	if o == nil || o.JsonType == nil {
		return nil, false
	}
	return o.JsonType, true
}

// HasJsonType returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasJsonType() bool {
	if o != nil && o.JsonType != nil {
		return true
	}

	return false
}

// SetJsonType gets a reference to the given string and assigns it to the JsonType field.
func (o *AddressjsonldWrite) SetJsonType(v string) {
	o.JsonType = &v
}

// GetCity returns the City field value
func (o *AddressjsonldWrite) GetCity() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AddressjsonldWrite) SetCity(v string) {
	o.City = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *AddressjsonldWrite) SetCompany(v string) {
	o.Company = &v
}

// GetCountryCode returns the CountryCode field value
func (o *AddressjsonldWrite) GetCountryCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *AddressjsonldWrite) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddressjsonldWrite) SetEmail(v string) {
	o.Email = &v
}

// GetHouseNumber returns the HouseNumber field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetHouseNumber() string {
	if o == nil || o.HouseNumber == nil {
		var ret string
		return ret
	}
	return *o.HouseNumber
}

// GetHouseNumberOk returns a tuple with the HouseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetHouseNumberOk() (*string, bool) {
	if o == nil || o.HouseNumber == nil {
		return nil, false
	}
	return o.HouseNumber, true
}

// HasHouseNumber returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasHouseNumber() bool {
	if o != nil && o.HouseNumber != nil {
		return true
	}

	return false
}

// SetHouseNumber gets a reference to the given string and assigns it to the HouseNumber field.
func (o *AddressjsonldWrite) SetHouseNumber(v string) {
	o.HouseNumber = &v
}

// GetName returns the Name field value
func (o *AddressjsonldWrite) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddressjsonldWrite) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *AddressjsonldWrite) SetPhone(v string) {
	o.Phone = &v
}

// GetStreet returns the Street field value
func (o *AddressjsonldWrite) GetStreet() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *AddressjsonldWrite) SetStreet(v string) {
	o.Street = v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *AddressjsonldWrite) GetStreet2() string {
	if o == nil || o.Street2 == nil {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetStreet2Ok() (*string, bool) {
	if o == nil || o.Street2 == nil {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *AddressjsonldWrite) HasStreet2() bool {
	if o != nil && o.Street2 != nil {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *AddressjsonldWrite) SetStreet2(v string) {
	o.Street2 = &v
}

// GetZip returns the Zip field value
func (o *AddressjsonldWrite) GetZip() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Zip
}

// GetZipOk returns a tuple with the Zip field value
// and a boolean to check if the value has been set.
func (o *AddressjsonldWrite) GetZipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Zip, true
}

// SetZip sets field value
func (o *AddressjsonldWrite) SetZip(v string) {
	o.Zip = v
}

func (o AddressjsonldWrite) MarshalJSON() ([]byte, error) {
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
		toSerialize["city"] = o.City
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if true {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.HouseNumber != nil {
		toSerialize["houseNumber"] = o.HouseNumber
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if o.Street2 != nil {
		toSerialize["street2"] = o.Street2
	}
	if true {
		toSerialize["zip"] = o.Zip
	}
	return json.Marshal(toSerialize)
}

type NullableAddressjsonldWrite struct {
	value *AddressjsonldWrite
	isSet bool
}

func (v NullableAddressjsonldWrite) Get() *AddressjsonldWrite {
	return v.value
}

func (v *NullableAddressjsonldWrite) Set(val *AddressjsonldWrite) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressjsonldWrite) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressjsonldWrite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressjsonldWrite(val *AddressjsonldWrite) *NullableAddressjsonldWrite {
	return &NullableAddressjsonldWrite{value: val, isSet: true}
}

func (v NullableAddressjsonldWrite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressjsonldWrite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


