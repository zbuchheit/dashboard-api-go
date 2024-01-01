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

// checks if the CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings{}

// CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings The OSPF routing settings of the interface.
type CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings struct {
	// The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an existing OSPF area. Defaults to 'disabled'.
	Area *string `json:"area,omitempty"`
	// The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.
	Cost *int32 `json:"cost,omitempty"`
	// When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled *bool `json:"isPassiveEnabled,omitempty"`
}

// NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings instantiates a new CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings() *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings {
	this := CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings{}
	return &this
}

// NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettingsWithDefaults instantiates a new CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettingsWithDefaults() *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings {
	this := CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings{}
	return &this
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) SetArea(v string) {
	o.Area = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetCost() int32 {
	if o == nil || IsNil(o.Cost) {
		var ret int32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetCostOk() (*int32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given int32 and assigns it to the Cost field.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) SetCost(v int32) {
	o.Cost = &v
}

// GetIsPassiveEnabled returns the IsPassiveEnabled field value if set, zero value otherwise.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetIsPassiveEnabled() bool {
	if o == nil || IsNil(o.IsPassiveEnabled) {
		var ret bool
		return ret
	}
	return *o.IsPassiveEnabled
}

// GetIsPassiveEnabledOk returns a tuple with the IsPassiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) GetIsPassiveEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPassiveEnabled) {
		return nil, false
	}
	return o.IsPassiveEnabled, true
}

// HasIsPassiveEnabled returns a boolean if a field has been set.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) HasIsPassiveEnabled() bool {
	if o != nil && !IsNil(o.IsPassiveEnabled) {
		return true
	}

	return false
}

// SetIsPassiveEnabled gets a reference to the given bool and assigns it to the IsPassiveEnabled field.
func (o *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) SetIsPassiveEnabled(v bool) {
	o.IsPassiveEnabled = &v
}

func (o CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.IsPassiveEnabled) {
		toSerialize["isPassiveEnabled"] = o.IsPassiveEnabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings struct {
	value *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings
	isSet bool
}

func (v NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) Get() *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings {
	return v.value
}

func (v *NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) Set(val *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings(val *CreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) *NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings {
	return &NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchStackRoutingInterfaceRequestOspfSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


