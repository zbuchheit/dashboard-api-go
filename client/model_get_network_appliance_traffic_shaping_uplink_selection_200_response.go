/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkApplianceTrafficShapingUplinkSelection200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceTrafficShapingUplinkSelection200Response{}

// GetNetworkApplianceTrafficShapingUplinkSelection200Response struct for GetNetworkApplianceTrafficShapingUplinkSelection200Response
type GetNetworkApplianceTrafficShapingUplinkSelection200Response struct {
	// Whether active-active AutoVPN is enabled
	ActiveActiveAutoVpnEnabled *bool `json:"activeActiveAutoVpnEnabled,omitempty"`
	// The default uplink. Must be one of: 'wan1' or 'wan2'
	DefaultUplink *string `json:"defaultUplink,omitempty"`
	// Whether load balancing is enabled
	LoadBalancingEnabled *bool `json:"loadBalancingEnabled,omitempty"`
	FailoverAndFailback *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback `json:"failoverAndFailback,omitempty"`
	// Uplink preference rules for WAN traffic
	WanTrafficUplinkPreferences []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner `json:"wanTrafficUplinkPreferences,omitempty"`
	// Uplink preference rules for VPN traffic
	VpnTrafficUplinkPreferences []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner `json:"vpnTrafficUplinkPreferences,omitempty"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200Response instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200Response() *GetNetworkApplianceTrafficShapingUplinkSelection200Response {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200Response{}
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200Response {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200Response{}
	return &this
}

// GetActiveActiveAutoVpnEnabled returns the ActiveActiveAutoVpnEnabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetActiveActiveAutoVpnEnabled() bool {
	if o == nil || IsNil(o.ActiveActiveAutoVpnEnabled) {
		var ret bool
		return ret
	}
	return *o.ActiveActiveAutoVpnEnabled
}

// GetActiveActiveAutoVpnEnabledOk returns a tuple with the ActiveActiveAutoVpnEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetActiveActiveAutoVpnEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveActiveAutoVpnEnabled) {
		return nil, false
	}
	return o.ActiveActiveAutoVpnEnabled, true
}

// HasActiveActiveAutoVpnEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasActiveActiveAutoVpnEnabled() bool {
	if o != nil && !IsNil(o.ActiveActiveAutoVpnEnabled) {
		return true
	}

	return false
}

// SetActiveActiveAutoVpnEnabled gets a reference to the given bool and assigns it to the ActiveActiveAutoVpnEnabled field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetActiveActiveAutoVpnEnabled(v bool) {
	o.ActiveActiveAutoVpnEnabled = &v
}

// GetDefaultUplink returns the DefaultUplink field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetDefaultUplink() string {
	if o == nil || IsNil(o.DefaultUplink) {
		var ret string
		return ret
	}
	return *o.DefaultUplink
}

// GetDefaultUplinkOk returns a tuple with the DefaultUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetDefaultUplinkOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultUplink) {
		return nil, false
	}
	return o.DefaultUplink, true
}

// HasDefaultUplink returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasDefaultUplink() bool {
	if o != nil && !IsNil(o.DefaultUplink) {
		return true
	}

	return false
}

// SetDefaultUplink gets a reference to the given string and assigns it to the DefaultUplink field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetDefaultUplink(v string) {
	o.DefaultUplink = &v
}

// GetLoadBalancingEnabled returns the LoadBalancingEnabled field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetLoadBalancingEnabled() bool {
	if o == nil || IsNil(o.LoadBalancingEnabled) {
		var ret bool
		return ret
	}
	return *o.LoadBalancingEnabled
}

// GetLoadBalancingEnabledOk returns a tuple with the LoadBalancingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetLoadBalancingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LoadBalancingEnabled) {
		return nil, false
	}
	return o.LoadBalancingEnabled, true
}

// HasLoadBalancingEnabled returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasLoadBalancingEnabled() bool {
	if o != nil && !IsNil(o.LoadBalancingEnabled) {
		return true
	}

	return false
}

// SetLoadBalancingEnabled gets a reference to the given bool and assigns it to the LoadBalancingEnabled field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetLoadBalancingEnabled(v bool) {
	o.LoadBalancingEnabled = &v
}

// GetFailoverAndFailback returns the FailoverAndFailback field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetFailoverAndFailback() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback {
	if o == nil || IsNil(o.FailoverAndFailback) {
		var ret GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback
		return ret
	}
	return *o.FailoverAndFailback
}

// GetFailoverAndFailbackOk returns a tuple with the FailoverAndFailback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetFailoverAndFailbackOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback, bool) {
	if o == nil || IsNil(o.FailoverAndFailback) {
		return nil, false
	}
	return o.FailoverAndFailback, true
}

// HasFailoverAndFailback returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasFailoverAndFailback() bool {
	if o != nil && !IsNil(o.FailoverAndFailback) {
		return true
	}

	return false
}

// SetFailoverAndFailback gets a reference to the given GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback and assigns it to the FailoverAndFailback field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetFailoverAndFailback(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseFailoverAndFailback) {
	o.FailoverAndFailback = &v
}

// GetWanTrafficUplinkPreferences returns the WanTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetWanTrafficUplinkPreferences() []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner {
	if o == nil || IsNil(o.WanTrafficUplinkPreferences) {
		var ret []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner
		return ret
	}
	return o.WanTrafficUplinkPreferences
}

// GetWanTrafficUplinkPreferencesOk returns a tuple with the WanTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetWanTrafficUplinkPreferencesOk() ([]GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner, bool) {
	if o == nil || IsNil(o.WanTrafficUplinkPreferences) {
		return nil, false
	}
	return o.WanTrafficUplinkPreferences, true
}

// HasWanTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasWanTrafficUplinkPreferences() bool {
	if o != nil && !IsNil(o.WanTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetWanTrafficUplinkPreferences gets a reference to the given []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner and assigns it to the WanTrafficUplinkPreferences field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetWanTrafficUplinkPreferences(v []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInner) {
	o.WanTrafficUplinkPreferences = v
}

// GetVpnTrafficUplinkPreferences returns the VpnTrafficUplinkPreferences field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetVpnTrafficUplinkPreferences() []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner {
	if o == nil || IsNil(o.VpnTrafficUplinkPreferences) {
		var ret []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner
		return ret
	}
	return o.VpnTrafficUplinkPreferences
}

// GetVpnTrafficUplinkPreferencesOk returns a tuple with the VpnTrafficUplinkPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) GetVpnTrafficUplinkPreferencesOk() ([]GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner, bool) {
	if o == nil || IsNil(o.VpnTrafficUplinkPreferences) {
		return nil, false
	}
	return o.VpnTrafficUplinkPreferences, true
}

// HasVpnTrafficUplinkPreferences returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) HasVpnTrafficUplinkPreferences() bool {
	if o != nil && !IsNil(o.VpnTrafficUplinkPreferences) {
		return true
	}

	return false
}

// SetVpnTrafficUplinkPreferences gets a reference to the given []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner and assigns it to the VpnTrafficUplinkPreferences field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200Response) SetVpnTrafficUplinkPreferences(v []GetNetworkApplianceTrafficShapingUplinkSelection200ResponseVpnTrafficUplinkPreferencesInner) {
	o.VpnTrafficUplinkPreferences = v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveActiveAutoVpnEnabled) {
		toSerialize["activeActiveAutoVpnEnabled"] = o.ActiveActiveAutoVpnEnabled
	}
	if !IsNil(o.DefaultUplink) {
		toSerialize["defaultUplink"] = o.DefaultUplink
	}
	if !IsNil(o.LoadBalancingEnabled) {
		toSerialize["loadBalancingEnabled"] = o.LoadBalancingEnabled
	}
	if !IsNil(o.FailoverAndFailback) {
		toSerialize["failoverAndFailback"] = o.FailoverAndFailback
	}
	if !IsNil(o.WanTrafficUplinkPreferences) {
		toSerialize["wanTrafficUplinkPreferences"] = o.WanTrafficUplinkPreferences
	}
	if !IsNil(o.VpnTrafficUplinkPreferences) {
		toSerialize["vpnTrafficUplinkPreferences"] = o.VpnTrafficUplinkPreferences
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200Response
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200Response(val *GetNetworkApplianceTrafficShapingUplinkSelection200Response) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


