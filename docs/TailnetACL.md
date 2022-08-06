# TailnetACL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagOwners** | Pointer to **map[string]interface{}** |  | [optional] 
**Ssh** | Pointer to [**[]SSHRule**](SSHRule.md) |  | [optional] 
**Acls** | Pointer to [**[]ACLRule**](ACLRule.md) |  | [optional] 
**Tests** | Pointer to [**[]Test**](Test.md) |  | [optional] 
**Groups** | Pointer to **map[string]interface{}** |  | [optional] 
**Hosts** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewTailnetACL

`func NewTailnetACL() *TailnetACL`

NewTailnetACL instantiates a new TailnetACL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTailnetACLWithDefaults

`func NewTailnetACLWithDefaults() *TailnetACL`

NewTailnetACLWithDefaults instantiates a new TailnetACL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagOwners

`func (o *TailnetACL) GetTagOwners() map[string]interface{}`

GetTagOwners returns the TagOwners field if non-nil, zero value otherwise.

### GetTagOwnersOk

`func (o *TailnetACL) GetTagOwnersOk() (*map[string]interface{}, bool)`

GetTagOwnersOk returns a tuple with the TagOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagOwners

`func (o *TailnetACL) SetTagOwners(v map[string]interface{})`

SetTagOwners sets TagOwners field to given value.

### HasTagOwners

`func (o *TailnetACL) HasTagOwners() bool`

HasTagOwners returns a boolean if a field has been set.

### GetSsh

`func (o *TailnetACL) GetSsh() []SSHRule`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *TailnetACL) GetSshOk() (*[]SSHRule, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *TailnetACL) SetSsh(v []SSHRule)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *TailnetACL) HasSsh() bool`

HasSsh returns a boolean if a field has been set.

### GetAcls

`func (o *TailnetACL) GetAcls() []ACLRule`

GetAcls returns the Acls field if non-nil, zero value otherwise.

### GetAclsOk

`func (o *TailnetACL) GetAclsOk() (*[]ACLRule, bool)`

GetAclsOk returns a tuple with the Acls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcls

`func (o *TailnetACL) SetAcls(v []ACLRule)`

SetAcls sets Acls field to given value.

### HasAcls

`func (o *TailnetACL) HasAcls() bool`

HasAcls returns a boolean if a field has been set.

### GetTests

`func (o *TailnetACL) GetTests() []Test`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *TailnetACL) GetTestsOk() (*[]Test, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *TailnetACL) SetTests(v []Test)`

SetTests sets Tests field to given value.

### HasTests

`func (o *TailnetACL) HasTests() bool`

HasTests returns a boolean if a field has been set.

### GetGroups

`func (o *TailnetACL) GetGroups() map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TailnetACL) GetGroupsOk() (*map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TailnetACL) SetGroups(v map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TailnetACL) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHosts

`func (o *TailnetACL) GetHosts() map[string]interface{}`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *TailnetACL) GetHostsOk() (*map[string]interface{}, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *TailnetACL) SetHosts(v map[string]interface{})`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *TailnetACL) HasHosts() bool`

HasHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


