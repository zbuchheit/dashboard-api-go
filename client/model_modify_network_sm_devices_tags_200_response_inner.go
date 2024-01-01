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

// checks if the ModifyNetworkSmDevicesTags200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModifyNetworkSmDevicesTags200ResponseInner{}

// ModifyNetworkSmDevicesTags200ResponseInner struct for ModifyNetworkSmDevicesTags200ResponseInner
type ModifyNetworkSmDevicesTags200ResponseInner struct {
	// The Meraki Id of the device record.
	Id *string `json:"id,omitempty"`
	// An array of tags associated with the device.
	Tags []string `json:"tags,omitempty"`
	// The MAC of the device.
	WifiMac *string `json:"wifiMac,omitempty"`
	// The device serial.
	Serial *string `json:"serial,omitempty"`
}

// NewModifyNetworkSmDevicesTags200ResponseInner instantiates a new ModifyNetworkSmDevicesTags200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyNetworkSmDevicesTags200ResponseInner() *ModifyNetworkSmDevicesTags200ResponseInner {
	this := ModifyNetworkSmDevicesTags200ResponseInner{}
	return &this
}

// NewModifyNetworkSmDevicesTags200ResponseInnerWithDefaults instantiates a new ModifyNetworkSmDevicesTags200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyNetworkSmDevicesTags200ResponseInnerWithDefaults() *ModifyNetworkSmDevicesTags200ResponseInner {
	this := ModifyNetworkSmDevicesTags200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetWifiMac returns the WifiMac field value if set, zero value otherwise.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetWifiMac() string {
	if o == nil || IsNil(o.WifiMac) {
		var ret string
		return ret
	}
	return *o.WifiMac
}

// GetWifiMacOk returns a tuple with the WifiMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetWifiMacOk() (*string, bool) {
	if o == nil || IsNil(o.WifiMac) {
		return nil, false
	}
	return o.WifiMac, true
}

// HasWifiMac returns a boolean if a field has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) HasWifiMac() bool {
	if o != nil && !IsNil(o.WifiMac) {
		return true
	}

	return false
}

// SetWifiMac gets a reference to the given string and assigns it to the WifiMac field.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) SetWifiMac(v string) {
	o.WifiMac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *ModifyNetworkSmDevicesTags200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

func (o ModifyNetworkSmDevicesTags200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyNetworkSmDevicesTags200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.WifiMac) {
		toSerialize["wifiMac"] = o.WifiMac
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	return toSerialize, nil
}

type NullableModifyNetworkSmDevicesTags200ResponseInner struct {
	value *ModifyNetworkSmDevicesTags200ResponseInner
	isSet bool
}

func (v NullableModifyNetworkSmDevicesTags200ResponseInner) Get() *ModifyNetworkSmDevicesTags200ResponseInner {
	return v.value
}

func (v *NullableModifyNetworkSmDevicesTags200ResponseInner) Set(val *ModifyNetworkSmDevicesTags200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyNetworkSmDevicesTags200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyNetworkSmDevicesTags200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyNetworkSmDevicesTags200ResponseInner(val *ModifyNetworkSmDevicesTags200ResponseInner) *NullableModifyNetworkSmDevicesTags200ResponseInner {
	return &NullableModifyNetworkSmDevicesTags200ResponseInner{value: val, isSet: true}
}

func (v NullableModifyNetworkSmDevicesTags200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyNetworkSmDevicesTags200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


