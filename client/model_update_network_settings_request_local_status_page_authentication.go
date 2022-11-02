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

// UpdateNetworkSettingsRequestLocalStatusPageAuthentication A hash of Local Status page(s)' authentication options applied to the Network.
type UpdateNetworkSettingsRequestLocalStatusPageAuthentication struct {
	// Enables / disables the authentication on Local Status page(s).
	Enabled *bool `json:"enabled,omitempty"`
	// A password used for Local Status Page(s). Set this null to clear the password.
	Password *string `json:"password,omitempty"`
}

// NewUpdateNetworkSettingsRequestLocalStatusPageAuthentication instantiates a new UpdateNetworkSettingsRequestLocalStatusPageAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSettingsRequestLocalStatusPageAuthentication() *UpdateNetworkSettingsRequestLocalStatusPageAuthentication {
	this := UpdateNetworkSettingsRequestLocalStatusPageAuthentication{}
	return &this
}

// NewUpdateNetworkSettingsRequestLocalStatusPageAuthenticationWithDefaults instantiates a new UpdateNetworkSettingsRequestLocalStatusPageAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSettingsRequestLocalStatusPageAuthenticationWithDefaults() *UpdateNetworkSettingsRequestLocalStatusPageAuthentication {
	this := UpdateNetworkSettingsRequestLocalStatusPageAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateNetworkSettingsRequestLocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication struct {
	value *UpdateNetworkSettingsRequestLocalStatusPageAuthentication
	isSet bool
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication) Get() *UpdateNetworkSettingsRequestLocalStatusPageAuthentication {
	return v.value
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication) Set(val *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication(val *UpdateNetworkSettingsRequestLocalStatusPageAuthentication) *NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication {
	return &NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication{value: val, isSet: true}
}

func (v NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSettingsRequestLocalStatusPageAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


