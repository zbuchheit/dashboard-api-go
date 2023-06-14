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

// UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner struct for UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner
type UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner struct {
	// A descriptive name for the rule
	Name *string `json:"name,omitempty"`
	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	LanIp string `json:"lanIp"`
	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2' or 'both')
	Uplink *string `json:"uplink,omitempty"`
	// A port or port ranges that will be forwarded to the host on the LAN
	PublicPort string `json:"publicPort"`
	// A port or port ranges that will receive the forwarded traffic from the WAN
	LocalPort string `json:"localPort"`
	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)
	AllowedIps []string `json:"allowedIps"`
	// TCP or UDP
	Protocol string `json:"protocol"`
}

// NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner(lanIp string, publicPort string, localPort string, allowedIps []string, protocol string) *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{}
	this.LanIp = lanIp
	this.PublicPort = publicPort
	this.LocalPort = localPort
	this.AllowedIps = allowedIps
	this.Protocol = protocol
	return &this
}

// NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetName(v string) {
	o.Name = &v
}

// GetLanIp returns the LanIp field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLanIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LanIp
}

// GetLanIpOk returns a tuple with the LanIp field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLanIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LanIp, true
}

// SetLanIp sets field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetLanIp(v string) {
	o.LanIp = v
}

// GetUplink returns the Uplink field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetUplink() string {
	if o == nil || isNil(o.Uplink) {
		var ret string
		return ret
	}
	return *o.Uplink
}

// GetUplinkOk returns a tuple with the Uplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetUplinkOk() (*string, bool) {
	if o == nil || isNil(o.Uplink) {
    return nil, false
	}
	return o.Uplink, true
}

// HasUplink returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) HasUplink() bool {
	if o != nil && !isNil(o.Uplink) {
		return true
	}

	return false
}

// SetUplink gets a reference to the given string and assigns it to the Uplink field.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetUplink(v string) {
	o.Uplink = &v
}

// GetPublicPort returns the PublicPort field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetPublicPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicPort
}

// GetPublicPortOk returns a tuple with the PublicPort field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetPublicPortOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublicPort, true
}

// SetPublicPort sets field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetPublicPort(v string) {
	o.PublicPort = v
}

// GetLocalPort returns the LocalPort field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLocalPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocalPort
}

// GetLocalPortOk returns a tuple with the LocalPort field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetLocalPortOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LocalPort, true
}

// SetLocalPort sets field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetLocalPort(v string) {
	o.LocalPort = v
}

// GetAllowedIps returns the AllowedIps field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetAllowedIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedIps
}

// GetAllowedIpsOk returns a tuple with the AllowedIps field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetAllowedIpsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AllowedIps, true
}

// SetAllowedIps sets field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetAllowedIps(v []string) {
	o.AllowedIps = v
}

// GetProtocol returns the Protocol field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) GetProtocolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) SetProtocol(v string) {
	o.Protocol = v
}

func (o UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["lanIp"] = o.LanIp
	}
	if !isNil(o.Uplink) {
		toSerialize["uplink"] = o.Uplink
	}
	if true {
		toSerialize["publicPort"] = o.PublicPort
	}
	if true {
		toSerialize["localPort"] = o.LocalPort
	}
	if true {
		toSerialize["allowedIps"] = o.AllowedIps
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner struct {
	value *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) Get() *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) Set(val *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner(val *UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) *NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner {
	return &NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


