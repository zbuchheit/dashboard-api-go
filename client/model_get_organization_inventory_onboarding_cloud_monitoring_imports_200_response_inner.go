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

// checks if the GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner{}

// GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner struct for GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
type GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner struct {
	// Database ID for the new entity entry.
	ImportId *string `json:"importId,omitempty"`
	Device *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice `json:"device,omitempty"`
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner instantiates a new GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner() *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner {
	this := GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner{}
	return &this
}

// NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerWithDefaults instantiates a new GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerWithDefaults() *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner {
	this := GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner{}
	return &this
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) GetImportId() string {
	if o == nil || IsNil(o.ImportId) {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) GetImportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImportId) {
		return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) HasImportId() bool {
	if o != nil && !IsNil(o.ImportId) {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) SetImportId(v string) {
	o.ImportId = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) GetDevice() GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice {
	if o == nil || IsNil(o.Device) {
		var ret GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) GetDeviceOk() (*GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice and assigns it to the Device field.
func (o *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) SetDevice(v GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInnerDevice) {
	o.Device = &v
}

func (o GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImportId) {
		toSerialize["importId"] = o.ImportId
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner struct {
	value *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) Get() *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) Set(val *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner(val *GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner {
	return &NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


