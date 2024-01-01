/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkWirelessSignalQualityHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSignalQualityHistory200ResponseInner{}

// GetNetworkWirelessSignalQualityHistory200ResponseInner struct for GetNetworkWirelessSignalQualityHistory200ResponseInner
type GetNetworkWirelessSignalQualityHistory200ResponseInner struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Signal to noise ratio
	Snr *int32 `json:"snr,omitempty"`
	// Received signal strength indicator
	Rssi *int32 `json:"rssi,omitempty"`
}

// NewGetNetworkWirelessSignalQualityHistory200ResponseInner instantiates a new GetNetworkWirelessSignalQualityHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSignalQualityHistory200ResponseInner() *GetNetworkWirelessSignalQualityHistory200ResponseInner {
	this := GetNetworkWirelessSignalQualityHistory200ResponseInner{}
	return &this
}

// NewGetNetworkWirelessSignalQualityHistory200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessSignalQualityHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSignalQualityHistory200ResponseInnerWithDefaults() *GetNetworkWirelessSignalQualityHistory200ResponseInner {
	this := GetNetworkWirelessSignalQualityHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetSnr returns the Snr field value if set, zero value otherwise.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetSnr() int32 {
	if o == nil || IsNil(o.Snr) {
		var ret int32
		return ret
	}
	return *o.Snr
}

// GetSnrOk returns a tuple with the Snr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetSnrOk() (*int32, bool) {
	if o == nil || IsNil(o.Snr) {
		return nil, false
	}
	return o.Snr, true
}

// HasSnr returns a boolean if a field has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasSnr() bool {
	if o != nil && !IsNil(o.Snr) {
		return true
	}

	return false
}

// SetSnr gets a reference to the given int32 and assigns it to the Snr field.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetSnr(v int32) {
	o.Snr = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetRssi() int32 {
	if o == nil || IsNil(o.Rssi) {
		var ret int32
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) GetRssiOk() (*int32, bool) {
	if o == nil || IsNil(o.Rssi) {
		return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) HasRssi() bool {
	if o != nil && !IsNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given int32 and assigns it to the Rssi field.
func (o *GetNetworkWirelessSignalQualityHistory200ResponseInner) SetRssi(v int32) {
	o.Rssi = &v
}

func (o GetNetworkWirelessSignalQualityHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSignalQualityHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.Snr) {
		toSerialize["snr"] = o.Snr
	}
	if !IsNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSignalQualityHistory200ResponseInner struct {
	value *GetNetworkWirelessSignalQualityHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkWirelessSignalQualityHistory200ResponseInner) Get() *GetNetworkWirelessSignalQualityHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkWirelessSignalQualityHistory200ResponseInner) Set(val *GetNetworkWirelessSignalQualityHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSignalQualityHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSignalQualityHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSignalQualityHistory200ResponseInner(val *GetNetworkWirelessSignalQualityHistory200ResponseInner) *NullableGetNetworkWirelessSignalQualityHistory200ResponseInner {
	return &NullableGetNetworkWirelessSignalQualityHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSignalQualityHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSignalQualityHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


