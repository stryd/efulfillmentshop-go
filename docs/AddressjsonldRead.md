# AddressjsonldRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**City** | **string** | The address city | 
**Company** | Pointer to **string** | The address company | [optional] 
**CountryCode** | **string** | The address country code (ISO 3166-1 alpha-2 format) | 
**Email** | Pointer to **string** | The address email | [optional] 
**HouseNumber** | Pointer to **string** | The address house number | [optional] 
**Id** | Pointer to **int32** | The address ID | [optional] [readonly] 
**Name** | **string** | The address name | 
**Phone** | Pointer to **string** | The address phone | [optional] 
**Street** | **string** | The address street | 
**Street2** | Pointer to **string** | The address street 2 | [optional] 
**Zip** | **string** | The address zip | 

## Methods

### NewAddressjsonldRead

`func NewAddressjsonldRead(city string, countryCode string, name string, street string, zip string, ) *AddressjsonldRead`

NewAddressjsonldRead instantiates a new AddressjsonldRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressjsonldReadWithDefaults

`func NewAddressjsonldReadWithDefaults() *AddressjsonldRead`

NewAddressjsonldReadWithDefaults instantiates a new AddressjsonldRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AddressjsonldRead) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AddressjsonldRead) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AddressjsonldRead) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AddressjsonldRead) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *AddressjsonldRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressjsonldRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressjsonldRead) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddressjsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddressjsonldRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressjsonldRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressjsonldRead) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddressjsonldRead) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCity

`func (o *AddressjsonldRead) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressjsonldRead) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressjsonldRead) SetCity(v string)`

SetCity sets City field to given value.


### GetCompany

`func (o *AddressjsonldRead) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AddressjsonldRead) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AddressjsonldRead) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AddressjsonldRead) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountryCode

`func (o *AddressjsonldRead) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AddressjsonldRead) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AddressjsonldRead) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetEmail

`func (o *AddressjsonldRead) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddressjsonldRead) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddressjsonldRead) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AddressjsonldRead) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHouseNumber

`func (o *AddressjsonldRead) GetHouseNumber() string`

GetHouseNumber returns the HouseNumber field if non-nil, zero value otherwise.

### GetHouseNumberOk

`func (o *AddressjsonldRead) GetHouseNumberOk() (*string, bool)`

GetHouseNumberOk returns a tuple with the HouseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHouseNumber

`func (o *AddressjsonldRead) SetHouseNumber(v string)`

SetHouseNumber sets HouseNumber field to given value.

### HasHouseNumber

`func (o *AddressjsonldRead) HasHouseNumber() bool`

HasHouseNumber returns a boolean if a field has been set.

### GetId

`func (o *AddressjsonldRead) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressjsonldRead) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressjsonldRead) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AddressjsonldRead) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddressjsonldRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddressjsonldRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddressjsonldRead) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *AddressjsonldRead) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *AddressjsonldRead) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *AddressjsonldRead) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *AddressjsonldRead) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *AddressjsonldRead) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressjsonldRead) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressjsonldRead) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetStreet2

`func (o *AddressjsonldRead) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *AddressjsonldRead) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *AddressjsonldRead) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *AddressjsonldRead) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetZip

`func (o *AddressjsonldRead) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *AddressjsonldRead) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *AddressjsonldRead) SetZip(v string)`

SetZip sets Zip field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


