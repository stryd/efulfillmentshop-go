# UserjsonldAuthenticate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsonContext** | Pointer to **string** |  | [optional] [readonly] 
**JsonId** | Pointer to **string** |  | [optional] [readonly] 
**JsonType** | Pointer to **string** |  | [optional] [readonly] 
**Email** | **string** | The user email address | 
**Password** | **string** | The user password | 

## Methods

### NewUserjsonldAuthenticate

`func NewUserjsonldAuthenticate(email string, password string, ) *UserjsonldAuthenticate`

NewUserjsonldAuthenticate instantiates a new UserjsonldAuthenticate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserjsonldAuthenticateWithDefaults

`func NewUserjsonldAuthenticateWithDefaults() *UserjsonldAuthenticate`

NewUserjsonldAuthenticateWithDefaults instantiates a new UserjsonldAuthenticate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsonContext

`func (o *UserjsonldAuthenticate) GetJsonContext() string`

GetJsonContext returns the JsonContext field if non-nil, zero value otherwise.

### GetJsonContextOk

`func (o *UserjsonldAuthenticate) GetJsonContextOk() (*string, bool)`

GetJsonContextOk returns a tuple with the JsonContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonContext

`func (o *UserjsonldAuthenticate) SetJsonContext(v string)`

SetJsonContext sets JsonContext field to given value.

### HasJsonContext

`func (o *UserjsonldAuthenticate) HasJsonContext() bool`

HasJsonContext returns a boolean if a field has been set.

### GetJsonId

`func (o *UserjsonldAuthenticate) GetJsonId() string`

GetJsonId returns the JsonId field if non-nil, zero value otherwise.

### GetJsonIdOk

`func (o *UserjsonldAuthenticate) GetJsonIdOk() (*string, bool)`

GetJsonIdOk returns a tuple with the JsonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonId

`func (o *UserjsonldAuthenticate) SetJsonId(v string)`

SetJsonId sets JsonId field to given value.

### HasJsonId

`func (o *UserjsonldAuthenticate) HasJsonId() bool`

HasJsonId returns a boolean if a field has been set.

### GetJsonType

`func (o *UserjsonldAuthenticate) GetJsonType() string`

GetJsonType returns the JsonType field if non-nil, zero value otherwise.

### GetJsonTypeOk

`func (o *UserjsonldAuthenticate) GetJsonTypeOk() (*string, bool)`

GetJsonTypeOk returns a tuple with the JsonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonType

`func (o *UserjsonldAuthenticate) SetJsonType(v string)`

SetJsonType sets JsonType field to given value.

### HasJsonType

`func (o *UserjsonldAuthenticate) HasJsonType() bool`

HasJsonType returns a boolean if a field has been set.

### GetEmail

`func (o *UserjsonldAuthenticate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserjsonldAuthenticate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserjsonldAuthenticate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UserjsonldAuthenticate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserjsonldAuthenticate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserjsonldAuthenticate) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


