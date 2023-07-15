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

// checks if the CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices{}

// CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices The devices and Switch Stacks assigned to the Group
type CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices struct {
	// Data Array of Devices containing the name and serial
	Devices []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner `json:"devices,omitempty"`
	// Data Array of Switch Stacks containing the name and id
	SwitchStacks []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner `json:"switchStacks,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices {
	this := CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesWithDefaults() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices {
	this := CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices{}
	return &this
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetDevices() []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner {
	if o == nil || IsNil(o.Devices) {
		var ret []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetDevicesOk() ([]CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner and assigns it to the Devices field.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) SetDevices(v []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesDevicesInner) {
	o.Devices = v
}

// GetSwitchStacks returns the SwitchStacks field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetSwitchStacks() []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner {
	if o == nil || IsNil(o.SwitchStacks) {
		var ret []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner
		return ret
	}
	return o.SwitchStacks
}

// GetSwitchStacksOk returns a tuple with the SwitchStacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) GetSwitchStacksOk() ([]CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner, bool) {
	if o == nil || IsNil(o.SwitchStacks) {
		return nil, false
	}
	return o.SwitchStacks, true
}

// HasSwitchStacks returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) HasSwitchStacks() bool {
	if o != nil && !IsNil(o.SwitchStacks) {
		return true
	}

	return false
}

// SetSwitchStacks gets a reference to the given []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner and assigns it to the SwitchStacks field.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) SetSwitchStacks(v []CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevicesSwitchStacksInner) {
	o.SwitchStacks = v
}

func (o CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.SwitchStacks) {
		toSerialize["switchStacks"] = o.SwitchStacks
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices struct {
	value *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) Get() *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) Set(val *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices(val *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices {
	return &NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


