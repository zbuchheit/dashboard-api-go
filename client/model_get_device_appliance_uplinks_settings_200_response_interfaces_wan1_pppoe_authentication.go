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

// checks if the GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication{}

// GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication Settings for PPPoE Authentication.
type GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication struct {
	// Whether PPPoE authentication is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Username for PPPoE authentication.
	Username *string `json:"username,omitempty"`
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication{}
	return &this
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthenticationWithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthenticationWithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) SetUsername(v string) {
	o.Username = &v
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication struct {
	value *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication
	isSet bool
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) Get() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication {
	return v.value
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) Set(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication {
	return &NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1PppoeAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


