/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct for UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest
type UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct {
	// An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule)
	Rules []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner `json:"rules,omitempty"`
	// Allow wireless client access to local LAN (boolean value - true allows access and false denies access) (optional)
	AllowLanAccess *bool `json:"allowLanAccess,omitempty"`
}

// NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest instantiates a new UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	this := UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestWithDefaults() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	this := UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetRules() []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner {
	if o == nil || isNil(o.Rules) {
		var ret []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetRulesOk() ([]UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner, bool) {
	if o == nil || isNil(o.Rules) {
    return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) HasRules() bool {
	if o != nil && !isNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) SetRules(v []UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequestRulesInner) {
	o.Rules = v
}

// GetAllowLanAccess returns the AllowLanAccess field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetAllowLanAccess() bool {
	if o == nil || isNil(o.AllowLanAccess) {
		var ret bool
		return ret
	}
	return *o.AllowLanAccess
}

// GetAllowLanAccessOk returns a tuple with the AllowLanAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) GetAllowLanAccessOk() (*bool, bool) {
	if o == nil || isNil(o.AllowLanAccess) {
    return nil, false
	}
	return o.AllowLanAccess, true
}

// HasAllowLanAccess returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) HasAllowLanAccess() bool {
	if o != nil && !isNil(o.AllowLanAccess) {
		return true
	}

	return false
}

// SetAllowLanAccess gets a reference to the given bool and assigns it to the AllowLanAccess field.
func (o *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) SetAllowLanAccess(v bool) {
	o.AllowLanAccess = &v
}

func (o UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !isNil(o.AllowLanAccess) {
		toSerialize["allowLanAccess"] = o.AllowLanAccess
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct {
	value *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) Get() *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) Set(val *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest(val *UpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest {
	return &NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


