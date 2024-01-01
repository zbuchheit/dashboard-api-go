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

// checks if the GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc{}

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc AC power details object
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc struct {
	// AC power connected
	IsConnected *bool `json:"isConnected,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAcWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAcWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc{}
	return &this
}

// GetIsConnected returns the IsConnected field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) GetIsConnected() bool {
	if o == nil || IsNil(o.IsConnected) {
		var ret bool
		return ret
	}
	return *o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) GetIsConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConnected) {
		return nil, false
	}
	return o.IsConnected, true
}

// HasIsConnected returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) HasIsConnected() bool {
	if o != nil && !IsNil(o.IsConnected) {
		return true
	}

	return false
}

// SetIsConnected gets a reference to the given bool and assigns it to the IsConnected field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) SetIsConnected(v bool) {
	o.IsConnected = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsConnected) {
		toSerialize["isConnected"] = o.IsConnected
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerPowerAc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


