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

// checks if the UpdateNetworkWirelessSsidRequestLdapServerCaCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestLdapServerCaCertificate{}

// UpdateNetworkWirelessSsidRequestLdapServerCaCertificate The CA certificate used to sign the LDAP server's key.
type UpdateNetworkWirelessSsidRequestLdapServerCaCertificate struct {
	// The contents of the CA certificate. Must be in PEM or DER format.
	Contents *string `json:"contents,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestLdapServerCaCertificate instantiates a new UpdateNetworkWirelessSsidRequestLdapServerCaCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestLdapServerCaCertificate() *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate {
	this := UpdateNetworkWirelessSsidRequestLdapServerCaCertificate{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestLdapServerCaCertificateWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestLdapServerCaCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestLdapServerCaCertificateWithDefaults() *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate {
	this := UpdateNetworkWirelessSsidRequestLdapServerCaCertificate{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) SetContents(v string) {
	o.Contents = &v
}

func (o UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate struct {
	value *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate) Get() *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate) Set(val *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate(val *UpdateNetworkWirelessSsidRequestLdapServerCaCertificate) *NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate {
	return &NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestLdapServerCaCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


