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

// UpdateDeviceCameraSenseRequestAudioDetection The details of the audio detection config.
type UpdateDeviceCameraSenseRequestAudioDetection struct {
	// Boolean indicating if audio detection is enabled(true) or disabled(false) on the camera
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateDeviceCameraSenseRequestAudioDetection instantiates a new UpdateDeviceCameraSenseRequestAudioDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCameraSenseRequestAudioDetection() *UpdateDeviceCameraSenseRequestAudioDetection {
	this := UpdateDeviceCameraSenseRequestAudioDetection{}
	return &this
}

// NewUpdateDeviceCameraSenseRequestAudioDetectionWithDefaults instantiates a new UpdateDeviceCameraSenseRequestAudioDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCameraSenseRequestAudioDetectionWithDefaults() *UpdateDeviceCameraSenseRequestAudioDetection {
	this := UpdateDeviceCameraSenseRequestAudioDetection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceCameraSenseRequestAudioDetection) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraSenseRequestAudioDetection) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceCameraSenseRequestAudioDetection) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceCameraSenseRequestAudioDetection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateDeviceCameraSenseRequestAudioDetection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceCameraSenseRequestAudioDetection struct {
	value *UpdateDeviceCameraSenseRequestAudioDetection
	isSet bool
}

func (v NullableUpdateDeviceCameraSenseRequestAudioDetection) Get() *UpdateDeviceCameraSenseRequestAudioDetection {
	return v.value
}

func (v *NullableUpdateDeviceCameraSenseRequestAudioDetection) Set(val *UpdateDeviceCameraSenseRequestAudioDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCameraSenseRequestAudioDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCameraSenseRequestAudioDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCameraSenseRequestAudioDetection(val *UpdateDeviceCameraSenseRequestAudioDetection) *NullableUpdateDeviceCameraSenseRequestAudioDetection {
	return &NullableUpdateDeviceCameraSenseRequestAudioDetection{value: val, isSet: true}
}

func (v NullableUpdateDeviceCameraSenseRequestAudioDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCameraSenseRequestAudioDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


