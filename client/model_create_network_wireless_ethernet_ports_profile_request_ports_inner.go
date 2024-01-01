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

// checks if the CreateNetworkWirelessEthernetPortsProfileRequestPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessEthernetPortsProfileRequestPortsInner{}

// CreateNetworkWirelessEthernetPortsProfileRequestPortsInner struct for CreateNetworkWirelessEthernetPortsProfileRequestPortsInner
type CreateNetworkWirelessEthernetPortsProfileRequestPortsInner struct {
	// AP port name
	Name string `json:"name"`
	// AP port enabled
	Enabled *bool `json:"enabled,omitempty"`
	// AP port ssid number
	Ssid *int32 `json:"ssid,omitempty"`
	// AP port PSK Group ID
	PskGroupId *string `json:"pskGroupId,omitempty"`
}

// NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner instantiates a new CreateNetworkWirelessEthernetPortsProfileRequestPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner(name string) *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner {
	this := CreateNetworkWirelessEthernetPortsProfileRequestPortsInner{}
	this.Name = name
	return &this
}

// NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults instantiates a new CreateNetworkWirelessEthernetPortsProfileRequestPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInnerWithDefaults() *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner {
	this := CreateNetworkWirelessEthernetPortsProfileRequestPortsInner{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetSsid() int32 {
	if o == nil || IsNil(o.Ssid) {
		var ret int32
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetSsidOk() (*int32, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int32 and assigns it to the Ssid field.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetSsid(v int32) {
	o.Ssid = &v
}

// GetPskGroupId returns the PskGroupId field value if set, zero value otherwise.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetPskGroupId() string {
	if o == nil || IsNil(o.PskGroupId) {
		var ret string
		return ret
	}
	return *o.PskGroupId
}

// GetPskGroupIdOk returns a tuple with the PskGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) GetPskGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PskGroupId) {
		return nil, false
	}
	return o.PskGroupId, true
}

// HasPskGroupId returns a boolean if a field has been set.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) HasPskGroupId() bool {
	if o != nil && !IsNil(o.PskGroupId) {
		return true
	}

	return false
}

// SetPskGroupId gets a reference to the given string and assigns it to the PskGroupId field.
func (o *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) SetPskGroupId(v string) {
	o.PskGroupId = &v
}

func (o CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.PskGroupId) {
		toSerialize["pskGroupId"] = o.PskGroupId
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner struct {
	value *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner
	isSet bool
}

func (v NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner) Get() *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner {
	return v.value
}

func (v *NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner) Set(val *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner(val *CreateNetworkWirelessEthernetPortsProfileRequestPortsInner) *NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner {
	return &NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessEthernetPortsProfileRequestPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


