/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// CreateNetworkGroupPolicyRequestBonjourForwarding The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.
type CreateNetworkGroupPolicyRequestBonjourForwarding struct {
	// How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	// A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.
	Rules []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner `json:"rules,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestBonjourForwarding instantiates a new CreateNetworkGroupPolicyRequestBonjourForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestBonjourForwarding() *CreateNetworkGroupPolicyRequestBonjourForwarding {
	this := CreateNetworkGroupPolicyRequestBonjourForwarding{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestBonjourForwardingWithDefaults instantiates a new CreateNetworkGroupPolicyRequestBonjourForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestBonjourForwardingWithDefaults() *CreateNetworkGroupPolicyRequestBonjourForwarding {
	this := CreateNetworkGroupPolicyRequestBonjourForwarding{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetSettings() string {
	if o == nil || isNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetSettingsOk() (*string, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) SetSettings(v string) {
	o.Settings = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetRules() []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner {
	if o == nil || isNil(o.Rules) {
		var ret []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) GetRulesOk() ([]CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner and assigns it to the Rules field.
func (o *CreateNetworkGroupPolicyRequestBonjourForwarding) SetRules(v []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) {
	o.Rules = v
}

func (o CreateNetworkGroupPolicyRequestBonjourForwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkGroupPolicyRequestBonjourForwarding struct {
	value *CreateNetworkGroupPolicyRequestBonjourForwarding
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestBonjourForwarding) Get() *CreateNetworkGroupPolicyRequestBonjourForwarding {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestBonjourForwarding) Set(val *CreateNetworkGroupPolicyRequestBonjourForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestBonjourForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestBonjourForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestBonjourForwarding(val *CreateNetworkGroupPolicyRequestBonjourForwarding) *NullableCreateNetworkGroupPolicyRequestBonjourForwarding {
	return &NullableCreateNetworkGroupPolicyRequestBonjourForwarding{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestBonjourForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestBonjourForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


