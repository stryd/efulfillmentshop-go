# AddressWriteLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
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

### NewAddressWriteLd

`func NewAddressWriteLd(city string, countryCode string, name string, street string, zip string, ) *AddressWriteLd`

NewAddressWriteLd instantiates a new AddressWriteLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWriteLdWithDefaults

`func NewAddressWriteLdWithDefaults() *AddressWriteLd`

NewAddressWriteLdWithDefaults instantiates a new AddressWriteLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *AddressWriteLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *AddressWriteLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *AddressWriteLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *AddressWriteLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *AddressWriteLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *AddressWriteLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *AddressWriteLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *AddressWriteLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *AddressWriteLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *AddressWriteLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *AddressWriteLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *AddressWriteLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetCity

`func (o *AddressWriteLd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressWriteLd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressWriteLd) SetCity(v string)`

SetCity sets City field to given value.


### GetCompany

`func (o *AddressWriteLd) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressWriteLd) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressWriteLd) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressWriteLd) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressWriteLd) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressWriteLd) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressWriteLd) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressWriteLd) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressWriteLd) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressWriteLd) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressWriteLd) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *AddressWriteLd) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AddressWriteLd) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AddressWriteLd) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *AddressWriteLd) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetName

`func (o *AddressWriteLd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressWriteLd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressWriteLd) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *AddressWriteLd) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressWriteLd) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressWriteLd) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressWriteLd) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *AddressWriteLd) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressWriteLd) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressWriteLd) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetStreet2

`func (o *AddressWriteLd) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *AddressWriteLd) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *AddressWriteLd) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *AddressWriteLd) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *AddressWriteLd) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddressWriteLd) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddressWriteLd) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


