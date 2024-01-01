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

// checks if the GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings{}

// GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings Manual radio settings for 2.4 GHz
type GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings struct {
	// Manual channel for 2.4 GHz
	Channel *int32 `json:"channel,omitempty"`
	// Manual target power for 2.4 GHz
	TargetPower *int32 `json:"targetPower,omitempty"`
}

// NewGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings instantiates a new GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings() *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings {
	this := GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings{}
	return &this
}

// NewGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettingsWithDefaults instantiates a new GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettingsWithDefaults() *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings {
	this := GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) SetChannel(v int32) {
	o.Channel = &v
}

// GetTargetPower returns the TargetPower field value if set, zero value otherwise.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) GetTargetPower() int32 {
	if o == nil || IsNil(o.TargetPower) {
		var ret int32
		return ret
	}
	return *o.TargetPower
}

// GetTargetPowerOk returns a tuple with the TargetPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) GetTargetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetPower) {
		return nil, false
	}
	return o.TargetPower, true
}

// HasTargetPower returns a boolean if a field has been set.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) HasTargetPower() bool {
	if o != nil && !IsNil(o.TargetPower) {
		return true
	}

	return false
}

// SetTargetPower gets a reference to the given int32 and assigns it to the TargetPower field.
func (o *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) SetTargetPower(v int32) {
	o.TargetPower = &v
}

func (o GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.TargetPower) {
		toSerialize["targetPower"] = o.TargetPower
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings struct {
	value *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings
	isSet bool
}

func (v NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) Get() *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings {
	return v.value
}

func (v *NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) Set(val *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings(val *GetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) *NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings {
	return &NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceRadioSettings200ResponseTwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


