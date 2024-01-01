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

// checks if the UpdateOrganizationSamlRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationSamlRequest{}

// UpdateOrganizationSamlRequest struct for UpdateOrganizationSamlRequest
type UpdateOrganizationSamlRequest struct {
	// Boolean for updating SAML SSO enabled settings.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateOrganizationSamlRequest instantiates a new UpdateOrganizationSamlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationSamlRequest() *UpdateOrganizationSamlRequest {
	this := UpdateOrganizationSamlRequest{}
	return &this
}

// NewUpdateOrganizationSamlRequestWithDefaults instantiates a new UpdateOrganizationSamlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationSamlRequestWithDefaults() *UpdateOrganizationSamlRequest {
	this := UpdateOrganizationSamlRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateOrganizationSamlRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationSamlRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationSamlRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateOrganizationSamlRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateOrganizationSamlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationSamlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationSamlRequest struct {
	value *UpdateOrganizationSamlRequest
	isSet bool
}

func (v NullableUpdateOrganizationSamlRequest) Get() *UpdateOrganizationSamlRequest {
	return v.value
}

func (v *NullableUpdateOrganizationSamlRequest) Set(val *UpdateOrganizationSamlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationSamlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationSamlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationSamlRequest(val *UpdateOrganizationSamlRequest) *NullableUpdateOrganizationSamlRequest {
	return &NullableUpdateOrganizationSamlRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationSamlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationSamlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


