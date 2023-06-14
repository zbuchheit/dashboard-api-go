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

// GetNetworkSwitchMtu200ResponseOverridesInner struct for GetNetworkSwitchMtu200ResponseOverridesInner
type GetNetworkSwitchMtu200ResponseOverridesInner struct {
	// List of switch serials. Applicable only for switch network.
	Switches []string `json:"switches,omitempty"`
	// List of switch profile IDs. Applicable only for template network.
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// MTU size for the switches or switch profiles.
	MtuSize int32 `json:"mtuSize"`
}

// NewGetNetworkSwitchMtu200ResponseOverridesInner instantiates a new GetNetworkSwitchMtu200ResponseOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchMtu200ResponseOverridesInner(mtuSize int32) *GetNetworkSwitchMtu200ResponseOverridesInner {
	this := GetNetworkSwitchMtu200ResponseOverridesInner{}
	this.MtuSize = mtuSize
	return &this
}

// NewGetNetworkSwitchMtu200ResponseOverridesInnerWithDefaults instantiates a new GetNetworkSwitchMtu200ResponseOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchMtu200ResponseOverridesInnerWithDefaults() *GetNetworkSwitchMtu200ResponseOverridesInner {
	this := GetNetworkSwitchMtu200ResponseOverridesInner{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) SetSwitches(v []string) {
	o.Switches = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetMtuSize returns the MtuSize field value
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetMtuSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MtuSize
}

// GetMtuSizeOk returns a tuple with the MtuSize field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) GetMtuSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MtuSize, true
}

// SetMtuSize sets field value
func (o *GetNetworkSwitchMtu200ResponseOverridesInner) SetMtuSize(v int32) {
	o.MtuSize = v
}

func (o GetNetworkSwitchMtu200ResponseOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !isNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if true {
		toSerialize["mtuSize"] = o.MtuSize
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchMtu200ResponseOverridesInner struct {
	value *GetNetworkSwitchMtu200ResponseOverridesInner
	isSet bool
}

func (v NullableGetNetworkSwitchMtu200ResponseOverridesInner) Get() *GetNetworkSwitchMtu200ResponseOverridesInner {
	return v.value
}

func (v *NullableGetNetworkSwitchMtu200ResponseOverridesInner) Set(val *GetNetworkSwitchMtu200ResponseOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchMtu200ResponseOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchMtu200ResponseOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchMtu200ResponseOverridesInner(val *GetNetworkSwitchMtu200ResponseOverridesInner) *NullableGetNetworkSwitchMtu200ResponseOverridesInner {
	return &NullableGetNetworkSwitchMtu200ResponseOverridesInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchMtu200ResponseOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchMtu200ResponseOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


