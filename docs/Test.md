# Test

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Src** | Pointer to **string** |  | [optional] 
**Accept** | Pointer to **[]string** |  | [optional] 
**Deny** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTest

`func NewTest() *Test`

NewTest instantiates a new Test object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestWithDefaults

`func NewTestWithDefaults() *Test`

NewTestWithDefaults instantiates a new Test object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrc

`func (o *Test) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Test) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Test) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Test) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetAccept

`func (o *Test) GetAccept() []string`

GetAccept returns the Accept field if non-nil, zero value otherwise.

### GetAcceptOk

`func (o *Test) GetAcceptOk() (*[]string, bool)`

GetAcceptOk returns a tuple with the Accept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccept

`func (o *Test) SetAccept(v []string)`

SetAccept sets Accept field to given value.

### HasAccept

`func (o *Test) HasAccept() bool`

HasAccept returns a boolean if a field has been set.

### GetDeny

`func (o *Test) GetDeny() []string`

GetDeny returns the Deny field if non-nil, zero value otherwise.

### GetDenyOk

`func (o *Test) GetDenyOk() (*[]string, bool)`

GetDenyOk returns a tuple with the Deny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeny

`func (o *Test) SetDeny(v []string)`

SetDeny sets Deny field to given value.

### HasDeny

`func (o *Test) HasDeny() bool`

HasDeny returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


