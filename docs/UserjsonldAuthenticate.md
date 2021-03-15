# UserjsonldAuthenticate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] [readonly] 
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

### GetContext

`func (o *UserjsonldAuthenticate) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserjsonldAuthenticate) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserjsonldAuthenticate) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *UserjsonldAuthenticate) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *UserjsonldAuthenticate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserjsonldAuthenticate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserjsonldAuthenticate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserjsonldAuthenticate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserjsonldAuthenticate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserjsonldAuthenticate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserjsonldAuthenticate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserjsonldAuthenticate) HasType() bool`

HasType returns a boolean if a field has been set.

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


