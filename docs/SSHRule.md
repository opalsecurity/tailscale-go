# SSHRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Src** | **[]string** |  | 
**Dst** | **[]string** |  | 
**Users** | **[]string** |  | 

## Methods

### NewSSHRule

`func NewSSHRule(action string, src []string, dst []string, users []string, ) *SSHRule`

NewSSHRule instantiates a new SSHRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHRuleWithDefaults

`func NewSSHRuleWithDefaults() *SSHRule`

NewSSHRuleWithDefaults instantiates a new SSHRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SSHRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SSHRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SSHRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetSrc

`func (o *SSHRule) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *SSHRule) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *SSHRule) SetSrc(v []string)`

SetSrc sets Src field to given value.


### GetDst

`func (o *SSHRule) GetDst() []string`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *SSHRule) GetDstOk() (*[]string, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *SSHRule) SetDst(v []string)`

SetDst sets Dst field to given value.


### GetUsers

`func (o *SSHRule) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SSHRule) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SSHRule) SetUsers(v []string)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


