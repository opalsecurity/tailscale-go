/*
Tailscale API

Tailscale API spec

API version: 1.0
Contact: hello@opal.dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opal

import (
	"encoding/json"
)

// SSHRule struct for SSHRule
type SSHRule struct {
	Action string `json:"action"`
	Src []string `json:"src"`
	Dst []string `json:"dst"`
	Users []string `json:"users"`
}

// NewSSHRule instantiates a new SSHRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHRule(action string, src []string, dst []string, users []string) *SSHRule {
	this := SSHRule{}
	this.Action = action
	this.Src = src
	this.Dst = dst
	this.Users = users
	return &this
}

// NewSSHRuleWithDefaults instantiates a new SSHRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHRuleWithDefaults() *SSHRule {
	this := SSHRule{}
	return &this
}

// GetAction returns the Action field value
func (o *SSHRule) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *SSHRule) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *SSHRule) SetAction(v string) {
	o.Action = v
}

// GetSrc returns the Src field value
func (o *SSHRule) GetSrc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Src
}

// GetSrcOk returns a tuple with the Src field value
// and a boolean to check if the value has been set.
func (o *SSHRule) GetSrcOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Src, true
}

// SetSrc sets field value
func (o *SSHRule) SetSrc(v []string) {
	o.Src = v
}

// GetDst returns the Dst field value
func (o *SSHRule) GetDst() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dst
}

// GetDstOk returns a tuple with the Dst field value
// and a boolean to check if the value has been set.
func (o *SSHRule) GetDstOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dst, true
}

// SetDst sets field value
func (o *SSHRule) SetDst(v []string) {
	o.Dst = v
}

// GetUsers returns the Users field value
func (o *SSHRule) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SSHRule) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SSHRule) SetUsers(v []string) {
	o.Users = v
}

func (o SSHRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["src"] = o.Src
	}
	if true {
		toSerialize["dst"] = o.Dst
	}
	if true {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableSSHRule struct {
	value *SSHRule
	isSet bool
}

func (v NullableSSHRule) Get() *SSHRule {
	return v.value
}

func (v *NullableSSHRule) Set(val *SSHRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHRule(val *SSHRule) *NullableSSHRule {
	return &NullableSSHRule{value: val, isSet: true}
}

func (v NullableSSHRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


