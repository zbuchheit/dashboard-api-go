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

// checks if the GetOrganizationDevicesStatuses200ResponseInnerComponents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesStatuses200ResponseInnerComponents{}

// GetOrganizationDevicesStatuses200ResponseInnerComponents Components
type GetOrganizationDevicesStatuses200ResponseInnerComponents struct {
	// Power Supplies
	PowerSupplies []GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner `json:"powerSupplies,omitempty"`
}

// NewGetOrganizationDevicesStatuses200ResponseInnerComponents instantiates a new GetOrganizationDevicesStatuses200ResponseInnerComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesStatuses200ResponseInnerComponents() *GetOrganizationDevicesStatuses200ResponseInnerComponents {
	this := GetOrganizationDevicesStatuses200ResponseInnerComponents{}
	return &this
}

// NewGetOrganizationDevicesStatuses200ResponseInnerComponentsWithDefaults instantiates a new GetOrganizationDevicesStatuses200ResponseInnerComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesStatuses200ResponseInnerComponentsWithDefaults() *GetOrganizationDevicesStatuses200ResponseInnerComponents {
	this := GetOrganizationDevicesStatuses200ResponseInnerComponents{}
	return &this
}

// GetPowerSupplies returns the PowerSupplies field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponents) GetPowerSupplies() []GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner {
	if o == nil || IsNil(o.PowerSupplies) {
		var ret []GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner
		return ret
	}
	return o.PowerSupplies
}

// GetPowerSuppliesOk returns a tuple with the PowerSupplies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponents) GetPowerSuppliesOk() ([]GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner, bool) {
	if o == nil || IsNil(o.PowerSupplies) {
		return nil, false
	}
	return o.PowerSupplies, true
}

// HasPowerSupplies returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponents) HasPowerSupplies() bool {
	if o != nil && !IsNil(o.PowerSupplies) {
		return true
	}

	return false
}

// SetPowerSupplies gets a reference to the given []GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner and assigns it to the PowerSupplies field.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponents) SetPowerSupplies(v []GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) {
	o.PowerSupplies = v
}

func (o GetOrganizationDevicesStatuses200ResponseInnerComponents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesStatuses200ResponseInnerComponents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PowerSupplies) {
		toSerialize["powerSupplies"] = o.PowerSupplies
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesStatuses200ResponseInnerComponents struct {
	value *GetOrganizationDevicesStatuses200ResponseInnerComponents
	isSet bool
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInnerComponents) Get() *GetOrganizationDevicesStatuses200ResponseInnerComponents {
	return v.value
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInnerComponents) Set(val *GetOrganizationDevicesStatuses200ResponseInnerComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInnerComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInnerComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesStatuses200ResponseInnerComponents(val *GetOrganizationDevicesStatuses200ResponseInnerComponents) *NullableGetOrganizationDevicesStatuses200ResponseInnerComponents {
	return &NullableGetOrganizationDevicesStatuses200ResponseInnerComponents{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInnerComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInnerComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


