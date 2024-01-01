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

// checks if the GetNetworkWirelessSettings200ResponseNamedVlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSettings200ResponseNamedVlans{}

// GetNetworkWirelessSettings200ResponseNamedVlans Named VLAN settings for wireless networks.
type GetNetworkWirelessSettings200ResponseNamedVlans struct {
	PoolDhcpMonitoring *GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring `json:"poolDhcpMonitoring,omitempty"`
}

// NewGetNetworkWirelessSettings200ResponseNamedVlans instantiates a new GetNetworkWirelessSettings200ResponseNamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSettings200ResponseNamedVlans() *GetNetworkWirelessSettings200ResponseNamedVlans {
	this := GetNetworkWirelessSettings200ResponseNamedVlans{}
	return &this
}

// NewGetNetworkWirelessSettings200ResponseNamedVlansWithDefaults instantiates a new GetNetworkWirelessSettings200ResponseNamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSettings200ResponseNamedVlansWithDefaults() *GetNetworkWirelessSettings200ResponseNamedVlans {
	this := GetNetworkWirelessSettings200ResponseNamedVlans{}
	return &this
}

// GetPoolDhcpMonitoring returns the PoolDhcpMonitoring field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200ResponseNamedVlans) GetPoolDhcpMonitoring() GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring {
	if o == nil || IsNil(o.PoolDhcpMonitoring) {
		var ret GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring
		return ret
	}
	return *o.PoolDhcpMonitoring
}

// GetPoolDhcpMonitoringOk returns a tuple with the PoolDhcpMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200ResponseNamedVlans) GetPoolDhcpMonitoringOk() (*GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring, bool) {
	if o == nil || IsNil(o.PoolDhcpMonitoring) {
		return nil, false
	}
	return o.PoolDhcpMonitoring, true
}

// HasPoolDhcpMonitoring returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200ResponseNamedVlans) HasPoolDhcpMonitoring() bool {
	if o != nil && !IsNil(o.PoolDhcpMonitoring) {
		return true
	}

	return false
}

// SetPoolDhcpMonitoring gets a reference to the given GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring and assigns it to the PoolDhcpMonitoring field.
func (o *GetNetworkWirelessSettings200ResponseNamedVlans) SetPoolDhcpMonitoring(v GetNetworkWirelessSettings200ResponseNamedVlansPoolDhcpMonitoring) {
	o.PoolDhcpMonitoring = &v
}

func (o GetNetworkWirelessSettings200ResponseNamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSettings200ResponseNamedVlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolDhcpMonitoring) {
		toSerialize["poolDhcpMonitoring"] = o.PoolDhcpMonitoring
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSettings200ResponseNamedVlans struct {
	value *GetNetworkWirelessSettings200ResponseNamedVlans
	isSet bool
}

func (v NullableGetNetworkWirelessSettings200ResponseNamedVlans) Get() *GetNetworkWirelessSettings200ResponseNamedVlans {
	return v.value
}

func (v *NullableGetNetworkWirelessSettings200ResponseNamedVlans) Set(val *GetNetworkWirelessSettings200ResponseNamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSettings200ResponseNamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSettings200ResponseNamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSettings200ResponseNamedVlans(val *GetNetworkWirelessSettings200ResponseNamedVlans) *NullableGetNetworkWirelessSettings200ResponseNamedVlans {
	return &NullableGetNetworkWirelessSettings200ResponseNamedVlans{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSettings200ResponseNamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSettings200ResponseNamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


