# SupplierWriteLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**City** | Pointer to **string** | The supplier city | [optional] 
**CountryCode** | Pointer to **string** | The supplier country code (ISO 3166-1 alpha-2 format) | [optional] 
**Email** | Pointer to **string** | The supplier email | [optional] 
**HouseNumber** | Pointer to **string** | The supplier house number | [optional] 
**Name** | **string** | The supplier name | 
**Phone** | Pointer to **string** | The supplier phone | [optional] 
**Street** | Pointer to **string** | The supplier street | [optional] 
**Street2** | Pointer to **string** | The supplier street 2 | [optional] 
**Zip** | Pointer to **string** | The supplier zip | [optional] 

## Methods

### NewSupplierWriteLd

`func NewSupplierWriteLd(name string, ) *SupplierWriteLd`

NewSupplierWriteLd instantiates a new SupplierWriteLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierWriteLdWithDefaults

`func NewSupplierWriteLdWithDefaults() *SupplierWriteLd`

NewSupplierWriteLdWithDefaults instantiates a new SupplierWriteLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *SupplierWriteLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *SupplierWriteLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *SupplierWriteLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *SupplierWriteLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *SupplierWriteLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *SupplierWriteLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *SupplierWriteLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *SupplierWriteLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *SupplierWriteLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *SupplierWriteLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *SupplierWriteLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *SupplierWriteLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetCity

`func (o *SupplierWriteLd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SupplierWriteLd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SupplierWriteLd) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SupplierWriteLd) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *SupplierWriteLd) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SupplierWriteLd) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SupplierWriteLd) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SupplierWriteLd) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *SupplierWriteLd) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupplierWriteLd) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupplierWriteLd) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupplierWriteLd) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *SupplierWriteLd) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *SupplierWriteLd) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *SupplierWriteLd) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *SupplierWriteLd) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetName

`func (o *SupplierWriteLd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplierWriteLd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplierWriteLd) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *SupplierWriteLd) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SupplierWriteLd) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SupplierWriteLd) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SupplierWriteLd) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *SupplierWriteLd) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SupplierWriteLd) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SupplierWriteLd) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SupplierWriteLd) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetStreet2

`func (o *SupplierWriteLd) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *SupplierWriteLd) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *SupplierWriteLd) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *SupplierWriteLd) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *SupplierWriteLd) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *SupplierWriteLd) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *SupplierWriteLd) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *SupplierWriteLd) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


