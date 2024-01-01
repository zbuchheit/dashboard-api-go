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

// checks if the UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request{}

// UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct for UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request
type UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct {
	// configured alternate management interface addresses
	Addresses []UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInner `json:"addresses,omitempty"`
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request{}
	return &this
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestWithDefaults instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestWithDefaults() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) GetAddresses() []UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInner {
	if o == nil || IsNil(o.Addresses) {
		var ret []UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInner
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) GetAddressesOk() ([]UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInner, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInner and assigns it to the Addresses field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) SetAddresses(v []UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInner) {
	o.Addresses = v
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return toSerialize, nil
}

type NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct {
	value *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request
	isSet bool
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) Get() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	return v.value
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) Set(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request {
	return &NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


