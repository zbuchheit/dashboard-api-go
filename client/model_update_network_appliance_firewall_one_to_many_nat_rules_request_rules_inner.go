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

// UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner struct for UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner
type UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner struct {
	// The IP address that will be used to access the internal resource from the WAN
	PublicIp string `json:"publicIp"`
	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
	Uplink string `json:"uplink"`
	// An array of associated forwarding rules
	PortRules []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner `json:"portRules"`
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner(publicIp string, uplink string, portRules []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner{}
	this.PublicIp = publicIp
	this.Uplink = uplink
	this.PortRules = portRules
	return &this
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner{}
	return &this
}

// GetPublicIp returns the PublicIp field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) GetPublicIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) GetPublicIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublicIp, true
}

// SetPublicIp sets field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) SetPublicIp(v string) {
	o.PublicIp = v
}

// GetUplink returns the Uplink field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) GetUplink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) GetUplinkOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Uplink, true
}

// SetUplink sets field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) SetUplink(v string) {
	o.Uplink = v
}

// GetPortRules returns the PortRules field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) GetPortRules() []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner {
	if o == nil {
		var ret []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner
		return ret
	}

	return o.PortRules
}

// GetPortRulesOk returns a tuple with the PortRules field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) GetPortRulesOk() ([]UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.PortRules, true
}

// SetPortRules sets field value
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) SetPortRules(v []UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner) {
	o.PortRules = v
}

func (o UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publicIp"] = o.PublicIp
	}
	if true {
		toSerialize["uplink"] = o.Uplink
	}
	if true {
		toSerialize["portRules"] = o.PortRules
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner struct {
	value *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) Get() *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) Set(val *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner(val *UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner {
	return &NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


