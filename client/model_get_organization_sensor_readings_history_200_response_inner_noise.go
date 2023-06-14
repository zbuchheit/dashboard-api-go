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

// GetOrganizationSensorReadingsHistory200ResponseInnerNoise Reading for the 'noise' metric. This will only be present if the 'metric' property equals 'noise'.
type GetOrganizationSensorReadingsHistory200ResponseInnerNoise struct {
	Ambient *GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient `json:"ambient,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNoise instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNoise object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNoise() *GetOrganizationSensorReadingsHistory200ResponseInnerNoise {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNoise{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNoiseWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNoise object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNoiseWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerNoise {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNoise{}
	return &this
}

// GetAmbient returns the Ambient field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoise) GetAmbient() GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient {
	if o == nil || isNil(o.Ambient) {
		var ret GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient
		return ret
	}
	return *o.Ambient
}

// GetAmbientOk returns a tuple with the Ambient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoise) GetAmbientOk() (*GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient, bool) {
	if o == nil || isNil(o.Ambient) {
    return nil, false
	}
	return o.Ambient, true
}

// HasAmbient returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoise) HasAmbient() bool {
	if o != nil && !isNil(o.Ambient) {
		return true
	}

	return false
}

// SetAmbient gets a reference to the given GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient and assigns it to the Ambient field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNoise) SetAmbient(v GetOrganizationSensorReadingsHistory200ResponseInnerNoiseAmbient) {
	o.Ambient = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerNoise) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ambient) {
		toSerialize["ambient"] = o.Ambient
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerNoise
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerNoise {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerNoise) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise(val *GetOrganizationSensorReadingsHistory200ResponseInnerNoise) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNoise) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


