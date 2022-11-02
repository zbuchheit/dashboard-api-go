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

// UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits A mapping of uplinks to their bandwidth settings (be sure to check which uplinks are supported for your network)
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits struct {
	Wan1 *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 `json:"wan1,omitempty"`
	Wan2 *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 `json:"wan2,omitempty"`
	Cellular *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular `json:"cellular,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits{}
	return &this
}

// GetWan1 returns the Wan1 field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) GetWan1() UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 {
	if o == nil || isNil(o.Wan1) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1
		return ret
	}
	return *o.Wan1
}

// GetWan1Ok returns a tuple with the Wan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) GetWan1Ok() (*UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1, bool) {
	if o == nil || isNil(o.Wan1) {
    return nil, false
	}
	return o.Wan1, true
}

// HasWan1 returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) HasWan1() bool {
	if o != nil && !isNil(o.Wan1) {
		return true
	}

	return false
}

// SetWan1 gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1 and assigns it to the Wan1 field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) SetWan1(v UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan1) {
	o.Wan1 = &v
}

// GetWan2 returns the Wan2 field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) GetWan2() UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 {
	if o == nil || isNil(o.Wan2) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2
		return ret
	}
	return *o.Wan2
}

// GetWan2Ok returns a tuple with the Wan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) GetWan2Ok() (*UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2, bool) {
	if o == nil || isNil(o.Wan2) {
    return nil, false
	}
	return o.Wan2, true
}

// HasWan2 returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) HasWan2() bool {
	if o != nil && !isNil(o.Wan2) {
		return true
	}

	return false
}

// SetWan2 gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 and assigns it to the Wan2 field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) SetWan2(v UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) {
	o.Wan2 = &v
}

// GetCellular returns the Cellular field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) GetCellular() UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular {
	if o == nil || isNil(o.Cellular) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular
		return ret
	}
	return *o.Cellular
}

// GetCellularOk returns a tuple with the Cellular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) GetCellularOk() (*UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.Cellular) {
    return nil, false
	}
	return o.Cellular, true
}

// HasCellular returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) HasCellular() bool {
	if o != nil && !isNil(o.Cellular) {
		return true
	}

	return false
}

// SetCellular gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular and assigns it to the Cellular field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) SetCellular(v UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) {
	o.Cellular = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Wan1) {
		toSerialize["wan1"] = o.Wan1
	}
	if !isNil(o.Wan2) {
		toSerialize["wan2"] = o.Wan2
	}
	if !isNil(o.Cellular) {
		toSerialize["cellular"] = o.Cellular
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) Get() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) Set(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


