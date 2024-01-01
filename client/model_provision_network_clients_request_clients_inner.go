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

// checks if the ProvisionNetworkClientsRequestClientsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionNetworkClientsRequestClientsInner{}

// ProvisionNetworkClientsRequestClientsInner struct for ProvisionNetworkClientsRequestClientsInner
type ProvisionNetworkClientsRequestClientsInner struct {
	// The MAC address of the client. Required.
	Mac string `json:"mac"`
	// The display name for the client. Optional. Limited to 255 bytes.
	Name *string `json:"name,omitempty"`
}

// NewProvisionNetworkClientsRequestClientsInner instantiates a new ProvisionNetworkClientsRequestClientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionNetworkClientsRequestClientsInner(mac string) *ProvisionNetworkClientsRequestClientsInner {
	this := ProvisionNetworkClientsRequestClientsInner{}
	this.Mac = mac
	return &this
}

// NewProvisionNetworkClientsRequestClientsInnerWithDefaults instantiates a new ProvisionNetworkClientsRequestClientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionNetworkClientsRequestClientsInnerWithDefaults() *ProvisionNetworkClientsRequestClientsInner {
	this := ProvisionNetworkClientsRequestClientsInner{}
	return &this
}

// GetMac returns the Mac field value
func (o *ProvisionNetworkClientsRequestClientsInner) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestClientsInner) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *ProvisionNetworkClientsRequestClientsInner) SetMac(v string) {
	o.Mac = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProvisionNetworkClientsRequestClientsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionNetworkClientsRequestClientsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProvisionNetworkClientsRequestClientsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProvisionNetworkClientsRequestClientsInner) SetName(v string) {
	o.Name = &v
}

func (o ProvisionNetworkClientsRequestClientsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionNetworkClientsRequestClientsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mac"] = o.Mac
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableProvisionNetworkClientsRequestClientsInner struct {
	value *ProvisionNetworkClientsRequestClientsInner
	isSet bool
}

func (v NullableProvisionNetworkClientsRequestClientsInner) Get() *ProvisionNetworkClientsRequestClientsInner {
	return v.value
}

func (v *NullableProvisionNetworkClientsRequestClientsInner) Set(val *ProvisionNetworkClientsRequestClientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionNetworkClientsRequestClientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionNetworkClientsRequestClientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionNetworkClientsRequestClientsInner(val *ProvisionNetworkClientsRequestClientsInner) *NullableProvisionNetworkClientsRequestClientsInner {
	return &NullableProvisionNetworkClientsRequestClientsInner{value: val, isSet: true}
}

func (v NullableProvisionNetworkClientsRequestClientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionNetworkClientsRequestClientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


