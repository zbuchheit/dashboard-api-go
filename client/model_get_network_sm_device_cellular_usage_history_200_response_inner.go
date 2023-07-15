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

// checks if the GetNetworkSmDeviceCellularUsageHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmDeviceCellularUsageHistory200ResponseInner{}

// GetNetworkSmDeviceCellularUsageHistory200ResponseInner struct for GetNetworkSmDeviceCellularUsageHistory200ResponseInner
type GetNetworkSmDeviceCellularUsageHistory200ResponseInner struct {
	// The amount of cellular data received by the device.
	Received *float32 `json:"received,omitempty"`
	// The amount of cellular sent received by the device.
	Sent *float32 `json:"sent,omitempty"`
	// When the cellular usage data was collected.
	Ts *string `json:"ts,omitempty"`
}

// NewGetNetworkSmDeviceCellularUsageHistory200ResponseInner instantiates a new GetNetworkSmDeviceCellularUsageHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceCellularUsageHistory200ResponseInner() *GetNetworkSmDeviceCellularUsageHistory200ResponseInner {
	this := GetNetworkSmDeviceCellularUsageHistory200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceCellularUsageHistory200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceCellularUsageHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceCellularUsageHistory200ResponseInnerWithDefaults() *GetNetworkSmDeviceCellularUsageHistory200ResponseInner {
	this := GetNetworkSmDeviceCellularUsageHistory200ResponseInner{}
	return &this
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) GetReceived() float32 {
	if o == nil || IsNil(o.Received) {
		var ret float32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) GetReceivedOk() (*float32, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given float32 and assigns it to the Received field.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) SetReceived(v float32) {
	o.Received = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) GetSent() float32 {
	if o == nil || IsNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) GetSentOk() (*float32, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) SetSent(v float32) {
	o.Sent = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) SetTs(v string) {
	o.Ts = &v
}

func (o GetNetworkSmDeviceCellularUsageHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmDeviceCellularUsageHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	return toSerialize, nil
}

type NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner struct {
	value *GetNetworkSmDeviceCellularUsageHistory200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner) Get() *GetNetworkSmDeviceCellularUsageHistory200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner) Set(val *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner(val *GetNetworkSmDeviceCellularUsageHistory200ResponseInner) *NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner {
	return &NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceCellularUsageHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


