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

// checks if the GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth{}

// GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth Critical auth settings for when authentication is rejected by the RADIUS server
type GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth struct {
	// VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	DataVlanId *int32 `json:"dataVlanId,omitempty"`
	// VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	VoiceVlanId *int32 `json:"voiceVlanId,omitempty"`
	// Enable to suspend port bounce when RADIUS servers are unreachable
	SuspendPortBounce *bool `json:"suspendPortBounce,omitempty"`
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth{}
	return &this
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuthWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuthWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth{}
	return &this
}

// GetDataVlanId returns the DataVlanId field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) GetDataVlanId() int32 {
	if o == nil || IsNil(o.DataVlanId) {
		var ret int32
		return ret
	}
	return *o.DataVlanId
}

// GetDataVlanIdOk returns a tuple with the DataVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) GetDataVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DataVlanId) {
		return nil, false
	}
	return o.DataVlanId, true
}

// HasDataVlanId returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) HasDataVlanId() bool {
	if o != nil && !IsNil(o.DataVlanId) {
		return true
	}

	return false
}

// SetDataVlanId gets a reference to the given int32 and assigns it to the DataVlanId field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) SetDataVlanId(v int32) {
	o.DataVlanId = &v
}

// GetVoiceVlanId returns the VoiceVlanId field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) GetVoiceVlanId() int32 {
	if o == nil || IsNil(o.VoiceVlanId) {
		var ret int32
		return ret
	}
	return *o.VoiceVlanId
}

// GetVoiceVlanIdOk returns a tuple with the VoiceVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) GetVoiceVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VoiceVlanId) {
		return nil, false
	}
	return o.VoiceVlanId, true
}

// HasVoiceVlanId returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) HasVoiceVlanId() bool {
	if o != nil && !IsNil(o.VoiceVlanId) {
		return true
	}

	return false
}

// SetVoiceVlanId gets a reference to the given int32 and assigns it to the VoiceVlanId field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) SetVoiceVlanId(v int32) {
	o.VoiceVlanId = &v
}

// GetSuspendPortBounce returns the SuspendPortBounce field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) GetSuspendPortBounce() bool {
	if o == nil || IsNil(o.SuspendPortBounce) {
		var ret bool
		return ret
	}
	return *o.SuspendPortBounce
}

// GetSuspendPortBounceOk returns a tuple with the SuspendPortBounce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) GetSuspendPortBounceOk() (*bool, bool) {
	if o == nil || IsNil(o.SuspendPortBounce) {
		return nil, false
	}
	return o.SuspendPortBounce, true
}

// HasSuspendPortBounce returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) HasSuspendPortBounce() bool {
	if o != nil && !IsNil(o.SuspendPortBounce) {
		return true
	}

	return false
}

// SetSuspendPortBounce gets a reference to the given bool and assigns it to the SuspendPortBounce field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) SetSuspendPortBounce(v bool) {
	o.SuspendPortBounce = &v
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataVlanId) {
		toSerialize["dataVlanId"] = o.DataVlanId
	}
	if !IsNil(o.VoiceVlanId) {
		toSerialize["voiceVlanId"] = o.VoiceVlanId
	}
	if !IsNil(o.SuspendPortBounce) {
		toSerialize["suspendPortBounce"] = o.SuspendPortBounce
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth struct {
	value *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth
	isSet bool
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) Get() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) Set(val *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth(val *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth {
	return &NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerRadiusCriticalAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


