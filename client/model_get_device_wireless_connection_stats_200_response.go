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

// checks if the GetDeviceWirelessConnectionStats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceWirelessConnectionStats200Response{}

// GetDeviceWirelessConnectionStats200Response struct for GetDeviceWirelessConnectionStats200Response
type GetDeviceWirelessConnectionStats200Response struct {
	// The serial number for the device
	Serial *string `json:"serial,omitempty"`
	ConnectionStats *GetDeviceWirelessConnectionStats200ResponseConnectionStats `json:"connectionStats,omitempty"`
}

// NewGetDeviceWirelessConnectionStats200Response instantiates a new GetDeviceWirelessConnectionStats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceWirelessConnectionStats200Response() *GetDeviceWirelessConnectionStats200Response {
	this := GetDeviceWirelessConnectionStats200Response{}
	return &this
}

// NewGetDeviceWirelessConnectionStats200ResponseWithDefaults instantiates a new GetDeviceWirelessConnectionStats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceWirelessConnectionStats200ResponseWithDefaults() *GetDeviceWirelessConnectionStats200Response {
	this := GetDeviceWirelessConnectionStats200Response{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetDeviceWirelessConnectionStats200Response) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessConnectionStats200Response) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetDeviceWirelessConnectionStats200Response) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetDeviceWirelessConnectionStats200Response) SetSerial(v string) {
	o.Serial = &v
}

// GetConnectionStats returns the ConnectionStats field value if set, zero value otherwise.
func (o *GetDeviceWirelessConnectionStats200Response) GetConnectionStats() GetDeviceWirelessConnectionStats200ResponseConnectionStats {
	if o == nil || IsNil(o.ConnectionStats) {
		var ret GetDeviceWirelessConnectionStats200ResponseConnectionStats
		return ret
	}
	return *o.ConnectionStats
}

// GetConnectionStatsOk returns a tuple with the ConnectionStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessConnectionStats200Response) GetConnectionStatsOk() (*GetDeviceWirelessConnectionStats200ResponseConnectionStats, bool) {
	if o == nil || IsNil(o.ConnectionStats) {
		return nil, false
	}
	return o.ConnectionStats, true
}

// HasConnectionStats returns a boolean if a field has been set.
func (o *GetDeviceWirelessConnectionStats200Response) HasConnectionStats() bool {
	if o != nil && !IsNil(o.ConnectionStats) {
		return true
	}

	return false
}

// SetConnectionStats gets a reference to the given GetDeviceWirelessConnectionStats200ResponseConnectionStats and assigns it to the ConnectionStats field.
func (o *GetDeviceWirelessConnectionStats200Response) SetConnectionStats(v GetDeviceWirelessConnectionStats200ResponseConnectionStats) {
	o.ConnectionStats = &v
}

func (o GetDeviceWirelessConnectionStats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceWirelessConnectionStats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.ConnectionStats) {
		toSerialize["connectionStats"] = o.ConnectionStats
	}
	return toSerialize, nil
}

type NullableGetDeviceWirelessConnectionStats200Response struct {
	value *GetDeviceWirelessConnectionStats200Response
	isSet bool
}

func (v NullableGetDeviceWirelessConnectionStats200Response) Get() *GetDeviceWirelessConnectionStats200Response {
	return v.value
}

func (v *NullableGetDeviceWirelessConnectionStats200Response) Set(val *GetDeviceWirelessConnectionStats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceWirelessConnectionStats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceWirelessConnectionStats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceWirelessConnectionStats200Response(val *GetDeviceWirelessConnectionStats200Response) *NullableGetDeviceWirelessConnectionStats200Response {
	return &NullableGetDeviceWirelessConnectionStats200Response{value: val, isSet: true}
}

func (v NullableGetDeviceWirelessConnectionStats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceWirelessConnectionStats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


