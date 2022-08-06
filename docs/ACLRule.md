# ACLRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Src** | **[]string** |  | 
**Dst** | **[]string** |  | 

## Methods

### NewACLRule

`func NewACLRule(action string, src []string, dst []string, ) *ACLRule`

NewACLRule instantiates a new ACLRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLRuleWithDefaults

`func NewACLRuleWithDefaults() *ACLRule`

NewACLRuleWithDefaults instantiates a new ACLRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ACLRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ACLRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ACLRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetSrc

`func (o *ACLRule) GetSrc() []string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *ACLRule) GetSrcOk() (*[]string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *ACLRule) SetSrc(v []string)`

SetSrc sets Src field to given value.


### GetDst

`func (o *ACLRule) GetDst() []string`

GetDst returns the Dst field if non-nil, zero value otherwise.

### GetDstOk

`func (o *ACLRule) GetDstOk() (*[]string, bool)`

GetDstOk returns a tuple with the Dst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDst

`func (o *ACLRule) SetDst(v []string)`

SetDst sets Dst field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


