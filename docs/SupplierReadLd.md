# SupplierReadLd

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
**Id** | Pointer to **int32** | The supplier ID | [optional] 
**Name** | **string** | The supplier name | 
**Phone** | Pointer to **string** | The supplier phone | [optional] 
**Street** | Pointer to **string** | The supplier street | [optional] 
**Street2** | Pointer to **string** | The supplier street 2 | [optional] 
**Zip** | Pointer to **string** | The supplier zip | [optional] 

## Methods

### NewSupplierReadLd

`func NewSupplierReadLd(name string, ) *SupplierReadLd`

NewSupplierReadLd instantiates a new SupplierReadLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierReadLdWithDefaults

`func NewSupplierReadLdWithDefaults() *SupplierReadLd`

NewSupplierReadLdWithDefaults instantiates a new SupplierReadLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *SupplierReadLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *SupplierReadLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *SupplierReadLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *SupplierReadLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *SupplierReadLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *SupplierReadLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *SupplierReadLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *SupplierReadLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *SupplierReadLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *SupplierReadLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *SupplierReadLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *SupplierReadLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetCity

`func (o *SupplierReadLd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SupplierReadLd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SupplierReadLd) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SupplierReadLd) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *SupplierReadLd) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SupplierReadLd) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SupplierReadLd) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SupplierReadLd) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *SupplierReadLd) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupplierReadLd) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupplierReadLd) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupplierReadLd) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *SupplierReadLd) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *SupplierReadLd) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *SupplierReadLd) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *SupplierReadLd) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetId

`func (o *SupplierReadLd) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupplierReadLd) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupplierReadLd) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SupplierReadLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SupplierReadLd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplierReadLd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplierReadLd) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *SupplierReadLd) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SupplierReadLd) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SupplierReadLd) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SupplierReadLd) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *SupplierReadLd) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SupplierReadLd) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SupplierReadLd) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SupplierReadLd) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetStreet2

`func (o *SupplierReadLd) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *SupplierReadLd) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *SupplierReadLd) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *SupplierReadLd) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *SupplierReadLd) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *SupplierReadLd) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *SupplierReadLd) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *SupplierReadLd) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


