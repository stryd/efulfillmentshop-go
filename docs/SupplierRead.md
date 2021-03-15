# SupplierRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The supplier city | [optional] 
**CountryCode** | Pointer to **string** | The supplier country code (ISO 3166-1 alpha-2 format) | [optional] 
**Email** | Pointer to **string** | The supplier email | [optional] 
**HouseNumber** | Pointer to **string** | The supplier house number | [optional] 
**Id** | Pointer to **int32** | The supplier ID | [optional] [readonly] 
**Name** | **string** | The supplier name | 
**Phone** | Pointer to **string** | The supplier phone | [optional] 
**Street** | Pointer to **string** | The supplier street | [optional] 
**Street2** | Pointer to **string** | The supplier street 2 | [optional] 
**Zip** | Pointer to **string** | The supplier zip | [optional] 

## Methods

### NewSupplierRead

`func NewSupplierRead(name string, ) *SupplierRead`

NewSupplierRead instantiates a new SupplierRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierReadWithDefaults

`func NewSupplierReadWithDefaults() *SupplierRead`

NewSupplierReadWithDefaults instantiates a new SupplierRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *SupplierRead) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SupplierRead) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SupplierRead) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SupplierRead) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *SupplierRead) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SupplierRead) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SupplierRead) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SupplierRead) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *SupplierRead) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupplierRead) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupplierRead) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupplierRead) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *SupplierRead) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *SupplierRead) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *SupplierRead) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *SupplierRead) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetId

`func (o *SupplierRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupplierRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupplierRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SupplierRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SupplierRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplierRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplierRead) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *SupplierRead) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SupplierRead) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SupplierRead) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SupplierRead) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *SupplierRead) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SupplierRead) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SupplierRead) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SupplierRead) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetStreet2

`func (o *SupplierRead) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *SupplierRead) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *SupplierRead) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *SupplierRead) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *SupplierRead) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *SupplierRead) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *SupplierRead) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *SupplierRead) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


