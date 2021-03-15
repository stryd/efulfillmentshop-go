# SupplierjsonldWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
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

### NewSupplierjsonldWrite

`func NewSupplierjsonldWrite(name string, ) *SupplierjsonldWrite`

NewSupplierjsonldWrite instantiates a new SupplierjsonldWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplierjsonldWriteWithDefaults

`func NewSupplierjsonldWriteWithDefaults() *SupplierjsonldWrite`

NewSupplierjsonldWriteWithDefaults instantiates a new SupplierjsonldWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SupplierjsonldWrite) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SupplierjsonldWrite) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SupplierjsonldWrite) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *SupplierjsonldWrite) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *SupplierjsonldWrite) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SupplierjsonldWrite) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SupplierjsonldWrite) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SupplierjsonldWrite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *SupplierjsonldWrite) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupplierjsonldWrite) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupplierjsonldWrite) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SupplierjsonldWrite) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCity

`func (o *SupplierjsonldWrite) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SupplierjsonldWrite) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SupplierjsonldWrite) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SupplierjsonldWrite) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *SupplierjsonldWrite) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SupplierjsonldWrite) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SupplierjsonldWrite) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SupplierjsonldWrite) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEmail

`func (o *SupplierjsonldWrite) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SupplierjsonldWrite) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SupplierjsonldWrite) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SupplierjsonldWrite) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *SupplierjsonldWrite) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *SupplierjsonldWrite) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *SupplierjsonldWrite) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *SupplierjsonldWrite) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetName

`func (o *SupplierjsonldWrite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupplierjsonldWrite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupplierjsonldWrite) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *SupplierjsonldWrite) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SupplierjsonldWrite) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SupplierjsonldWrite) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SupplierjsonldWrite) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *SupplierjsonldWrite) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SupplierjsonldWrite) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SupplierjsonldWrite) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SupplierjsonldWrite) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetStreet2

`func (o *SupplierjsonldWrite) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *SupplierjsonldWrite) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *SupplierjsonldWrite) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *SupplierjsonldWrite) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *SupplierjsonldWrite) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *SupplierjsonldWrite) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *SupplierjsonldWrite) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *SupplierjsonldWrite) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


