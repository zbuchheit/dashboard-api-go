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

// checks if the UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits{}

// UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.
type UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits struct {
	// The maximum upload limit (integer, in Kbps).
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps).
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimitsWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimitsWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) GetLimitUp() int32 {
	if o == nil || IsNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) GetLimitUpOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitUp) {
		return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) HasLimitUp() bool {
	if o != nil && !IsNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) GetLimitDown() int32 {
	if o == nil || IsNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) GetLimitDownOk() (*int32, bool) {
	if o == nil || IsNil(o.LimitDown) {
		return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) HasLimitDown() bool {
	if o != nil && !IsNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !IsNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits struct {
	value *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) Get() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) Set(val *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits(val *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits {
	return &NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


