# AddressWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | The address city | 
**Company** | Pointer to **string** | The address company | [optional] 
**CountryCode** | **string** | The address country code (ISO 3166-1 alpha-2 format) | 
**Email** | Pointer to **string** | The address email | [optional] 
**HouseNumber** | Pointer to **string** | The address house number | [optional] 
**Name** | **string** | The address name | 
**Phone** | Pointer to **string** | The address phone | [optional] 
**Street** | **string** | The address street | 
**Street2** | Pointer to **string** | The address street 2 | [optional] 
**Zip** | **string** | The address zip | 

## Methods

### NewAddressWrite

`func NewAddressWrite(city string, countryCode string, name string, street string, zip string, ) *AddressWrite`

NewAddressWrite instantiates a new AddressWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWriteWithDefaults

`func NewAddressWriteWithDefaults() *AddressWrite`

NewAddressWriteWithDefaults instantiates a new AddressWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *AddressWrite) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressWrite) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressWrite) SetCity(v string)`

SetCity sets City field to given value.


### GetCompany

`func (o *AddressWrite) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressWrite) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressWrite) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressWrite) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressWrite) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressWrite) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressWrite) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressWrite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressWrite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressWrite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressWrite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *AddressWrite) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AddressWrite) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AddressWrite) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *AddressWrite) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetName

`func (o *AddressWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressWrite) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *AddressWrite) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressWrite) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressWrite) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressWrite) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *AddressWrite) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressWrite) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressWrite) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetStreet2

`func (o *AddressWrite) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *AddressWrite) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *AddressWrite) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *AddressWrite) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *AddressWrite) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddressWrite) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddressWrite) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


