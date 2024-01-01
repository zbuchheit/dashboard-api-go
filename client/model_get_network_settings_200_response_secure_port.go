/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkSettings200ResponseSecurePort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSettings200ResponseSecurePort{}

// GetNetworkSettings200ResponseSecurePort A hash of SecureConnect options applied to the Network.
type GetNetworkSettings200ResponseSecurePort struct {
	// Enables / disables SecureConnect on the network. Optional.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkSettings200ResponseSecurePort instantiates a new GetNetworkSettings200ResponseSecurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSettings200ResponseSecurePort() *GetNetworkSettings200ResponseSecurePort {
	this := GetNetworkSettings200ResponseSecurePort{}
	return &this
}

// NewGetNetworkSettings200ResponseSecurePortWithDefaults instantiates a new GetNetworkSettings200ResponseSecurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSettings200ResponseSecurePortWithDefaults() *GetNetworkSettings200ResponseSecurePort {
	this := GetNetworkSettings200ResponseSecurePort{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSettings200ResponseSecurePort) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSettings200ResponseSecurePort) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSettings200ResponseSecurePort) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSettings200ResponseSecurePort) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkSettings200ResponseSecurePort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSettings200ResponseSecurePort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkSettings200ResponseSecurePort struct {
	value *GetNetworkSettings200ResponseSecurePort
	isSet bool
}

func (v NullableGetNetworkSettings200ResponseSecurePort) Get() *GetNetworkSettings200ResponseSecurePort {
	return v.value
}

func (v *NullableGetNetworkSettings200ResponseSecurePort) Set(val *GetNetworkSettings200ResponseSecurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSettings200ResponseSecurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSettings200ResponseSecurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSettings200ResponseSecurePort(val *GetNetworkSettings200ResponseSecurePort) *NullableGetNetworkSettings200ResponseSecurePort {
	return &NullableGetNetworkSettings200ResponseSecurePort{value: val, isSet: true}
}

func (v NullableGetNetworkSettings200ResponseSecurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSettings200ResponseSecurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


