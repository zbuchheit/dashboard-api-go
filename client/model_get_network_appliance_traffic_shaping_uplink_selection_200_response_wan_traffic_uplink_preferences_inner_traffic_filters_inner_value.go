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

// checks if the GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}

// GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue Value of traffic filter
type GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue struct {
	// Protocol of 'custom' type traffic filter. Must be one of: 'tcp', 'udp', 'icmp6' or 'any'
	Protocol *string `json:"protocol,omitempty"`
	Source GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource `json:"source"`
	Destination GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination `json:"destination"`
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue(source GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource, destination GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}
	this.Source = source
	this.Destination = destination
	return &this
}

// NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueWithDefaults instantiates a new GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueWithDefaults() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	this := GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSource returns the Source field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetSource() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource {
	if o == nil {
		var ret GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetSourceOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetSource(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueSource) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetDestination() GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination {
	if o == nil {
		var ret GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) GetDestinationOk() (*GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) SetDestination(v GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValueDestination) {
	o.Destination = v
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	toSerialize["source"] = o.Source
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

type NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue struct {
	value *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue
	isSet bool
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Get() *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	return v.value
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Set(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue(val *GetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue {
	return &NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceTrafficShapingUplinkSelection200ResponseWanTrafficUplinkPreferencesInnerTrafficFiltersInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


