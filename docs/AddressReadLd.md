# AddressReadLd

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
**Id** | Pointer to **int32** | The address ID | [optional] 
**Name** | **string** | The address name | 
**Phone** | Pointer to **string** | The address phone | [optional] 
**Street** | **string** | The address street | 
**Street2** | Pointer to **string** | The address street 2 | [optional] 
**Zip** | **string** | The address zip | 

## Methods

### NewAddressReadLd

`func NewAddressReadLd(city string, countryCode string, name string, street string, zip string, ) *AddressReadLd`

NewAddressReadLd instantiates a new AddressReadLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressReadLdWithDefaults

`func NewAddressReadLdWithDefaults() *AddressReadLd`

NewAddressReadLdWithDefaults instantiates a new AddressReadLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *AddressReadLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *AddressReadLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *AddressReadLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *AddressReadLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *AddressReadLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *AddressReadLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *AddressReadLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *AddressReadLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *AddressReadLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *AddressReadLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *AddressReadLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *AddressReadLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetCity

`func (o *AddressReadLd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressReadLd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressReadLd) SetCity(v string)`

SetCity sets City field to given value.


### GetCompany

`func (o *AddressReadLd) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressReadLd) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressReadLd) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressReadLd) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressReadLd) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressReadLd) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressReadLd) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressReadLd) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressReadLd) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressReadLd) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressReadLd) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *AddressReadLd) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AddressReadLd) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AddressReadLd) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *AddressReadLd) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetId

`func (o *AddressReadLd) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressReadLd) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressReadLd) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AddressReadLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddressReadLd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressReadLd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressReadLd) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *AddressReadLd) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressReadLd) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressReadLd) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressReadLd) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *AddressReadLd) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressReadLd) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressReadLd) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetStreet2

`func (o *AddressReadLd) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *AddressReadLd) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *AddressReadLd) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *AddressReadLd) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *AddressReadLd) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddressReadLd) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddressReadLd) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


