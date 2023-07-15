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

// checks if the UpdateNetworkWirelessSsidSchedulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSchedulesRequest{}

// UpdateNetworkWirelessSsidSchedulesRequest struct for UpdateNetworkWirelessSsidSchedulesRequest
type UpdateNetworkWirelessSsidSchedulesRequest struct {
	// If true, the SSID outage schedule is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.
	Ranges []UpdateNetworkWirelessSsidSchedulesRequestRangesInner `json:"ranges,omitempty"`
	// List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.
	RangesInSeconds []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner `json:"rangesInSeconds,omitempty"`
}

// NewUpdateNetworkWirelessSsidSchedulesRequest instantiates a new UpdateNetworkWirelessSsidSchedulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSchedulesRequest() *UpdateNetworkWirelessSsidSchedulesRequest {
	this := UpdateNetworkWirelessSsidSchedulesRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidSchedulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidSchedulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSchedulesRequestWithDefaults() *UpdateNetworkWirelessSsidSchedulesRequest {
	this := UpdateNetworkWirelessSsidSchedulesRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRanges() []UpdateNetworkWirelessSsidSchedulesRequestRangesInner {
	if o == nil || IsNil(o.Ranges) {
		var ret []UpdateNetworkWirelessSsidSchedulesRequestRangesInner
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRangesOk() ([]UpdateNetworkWirelessSsidSchedulesRequestRangesInner, bool) {
	if o == nil || IsNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) HasRanges() bool {
	if o != nil && !IsNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []UpdateNetworkWirelessSsidSchedulesRequestRangesInner and assigns it to the Ranges field.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) SetRanges(v []UpdateNetworkWirelessSsidSchedulesRequestRangesInner) {
	o.Ranges = v
}

// GetRangesInSeconds returns the RangesInSeconds field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRangesInSeconds() []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner {
	if o == nil || IsNil(o.RangesInSeconds) {
		var ret []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner
		return ret
	}
	return o.RangesInSeconds
}

// GetRangesInSecondsOk returns a tuple with the RangesInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) GetRangesInSecondsOk() ([]UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner, bool) {
	if o == nil || IsNil(o.RangesInSeconds) {
		return nil, false
	}
	return o.RangesInSeconds, true
}

// HasRangesInSeconds returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) HasRangesInSeconds() bool {
	if o != nil && !IsNil(o.RangesInSeconds) {
		return true
	}

	return false
}

// SetRangesInSeconds gets a reference to the given []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner and assigns it to the RangesInSeconds field.
func (o *UpdateNetworkWirelessSsidSchedulesRequest) SetRangesInSeconds(v []UpdateNetworkWirelessSsidSchedulesRequestRangesInSecondsInner) {
	o.RangesInSeconds = v
}

func (o UpdateNetworkWirelessSsidSchedulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSchedulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	if !IsNil(o.RangesInSeconds) {
		toSerialize["rangesInSeconds"] = o.RangesInSeconds
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSchedulesRequest struct {
	value *UpdateNetworkWirelessSsidSchedulesRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSchedulesRequest) Get() *UpdateNetworkWirelessSsidSchedulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSchedulesRequest) Set(val *UpdateNetworkWirelessSsidSchedulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSchedulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSchedulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSchedulesRequest(val *UpdateNetworkWirelessSsidSchedulesRequest) *NullableUpdateNetworkWirelessSsidSchedulesRequest {
	return &NullableUpdateNetworkWirelessSsidSchedulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSchedulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSchedulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


