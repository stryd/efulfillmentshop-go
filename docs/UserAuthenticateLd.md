# UserAuthenticateLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdContext** | Pointer to **string** |  | [optional] [readonly] 
**LdId** | Pointer to **string** |  | [optional] [readonly] 
**LdType** | Pointer to **string** |  | [optional] [readonly] 
**Email** | **string** | The user email address | 
**Password** | **string** | The user password | 

## Methods

### NewUserAuthenticateLd

`func NewUserAuthenticateLd(email string, password string, ) *UserAuthenticateLd`

NewUserAuthenticateLd instantiates a new UserAuthenticateLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAuthenticateLdWithDefaults

`func NewUserAuthenticateLdWithDefaults() *UserAuthenticateLd`

NewUserAuthenticateLdWithDefaults instantiates a new UserAuthenticateLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdContext

`func (o *UserAuthenticateLd) GetLdContext() string`

GetLdContext returns the LdContext field if non-nil, zero value otherwise.

### GetLdContextOk

`func (o *UserAuthenticateLd) GetLdContextOk() (*string, bool)`

GetLdContextOk returns a tuple with the LdContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdContext

`func (o *UserAuthenticateLd) SetLdContext(v string)`

SetLdContext sets LdContext field to given value.

### HasLdContext

`func (o *UserAuthenticateLd) HasLdContext() bool`

HasLdContext returns a boolean if a field has been set.

### GetLdId

`func (o *UserAuthenticateLd) GetLdId() string`

GetLdId returns the LdId field if non-nil, zero value otherwise.

### GetLdIdOk

`func (o *UserAuthenticateLd) GetLdIdOk() (*string, bool)`

GetLdIdOk returns a tuple with the LdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdId

`func (o *UserAuthenticateLd) SetLdId(v string)`

SetLdId sets LdId field to given value.

### HasLdId

`func (o *UserAuthenticateLd) HasLdId() bool`

HasLdId returns a boolean if a field has been set.

### GetLdType

`func (o *UserAuthenticateLd) GetLdType() string`

GetLdType returns the LdType field if non-nil, zero value otherwise.

### GetLdTypeOk

`func (o *UserAuthenticateLd) GetLdTypeOk() (*string, bool)`

GetLdTypeOk returns a tuple with the LdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdType

`func (o *UserAuthenticateLd) SetLdType(v string)`

SetLdType sets LdType field to given value.

### HasLdType

`func (o *UserAuthenticateLd) HasLdType() bool`

HasLdType returns a boolean if a field has been set.

### GetEmail

`func (o *UserAuthenticateLd) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAuthenticateLd) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAuthenticateLd) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UserAuthenticateLd) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserAuthenticateLd) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserAuthenticateLd) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


