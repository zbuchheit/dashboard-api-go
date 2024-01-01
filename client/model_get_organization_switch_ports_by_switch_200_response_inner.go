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

// checks if the GetOrganizationSwitchPortsBySwitch200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSwitchPortsBySwitch200ResponseInner{}

// GetOrganizationSwitchPortsBySwitch200ResponseInner struct for GetOrganizationSwitchPortsBySwitch200ResponseInner
type GetOrganizationSwitchPortsBySwitch200ResponseInner struct {
	// The name of the switch.
	Name *string `json:"name,omitempty"`
	// The serial number of the switch.
	Serial *string `json:"serial,omitempty"`
	// The MAC address of the switch.
	Mac *string `json:"mac,omitempty"`
	Network *GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork `json:"network,omitempty"`
	// The model of the switch.
	Model *string `json:"model,omitempty"`
	// Ports belonging to the switch
	Ports []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner `json:"ports,omitempty"`
}

// NewGetOrganizationSwitchPortsBySwitch200ResponseInner instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSwitchPortsBySwitch200ResponseInner() *GetOrganizationSwitchPortsBySwitch200ResponseInner {
	this := GetOrganizationSwitchPortsBySwitch200ResponseInner{}
	return &this
}

// NewGetOrganizationSwitchPortsBySwitch200ResponseInnerWithDefaults instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSwitchPortsBySwitch200ResponseInnerWithDefaults() *GetOrganizationSwitchPortsBySwitch200ResponseInner {
	this := GetOrganizationSwitchPortsBySwitch200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetNetwork() GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetNetworkOk() (*GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetNetwork(v GetOrganizationSwitchPortsBySwitch200ResponseInnerNetwork) {
	o.Network = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetPorts() []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner {
	if o == nil || IsNil(o.Ports) {
		var ret []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) GetPortsOk() ([]GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner and assigns it to the Ports field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInner) SetPorts(v []GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) {
	o.Ports = v
}

func (o GetOrganizationSwitchPortsBySwitch200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSwitchPortsBySwitch200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableGetOrganizationSwitchPortsBySwitch200ResponseInner struct {
	value *GetOrganizationSwitchPortsBySwitch200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSwitchPortsBySwitch200ResponseInner) Get() *GetOrganizationSwitchPortsBySwitch200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSwitchPortsBySwitch200ResponseInner) Set(val *GetOrganizationSwitchPortsBySwitch200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSwitchPortsBySwitch200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSwitchPortsBySwitch200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSwitchPortsBySwitch200ResponseInner(val *GetOrganizationSwitchPortsBySwitch200ResponseInner) *NullableGetOrganizationSwitchPortsBySwitch200ResponseInner {
	return &NullableGetOrganizationSwitchPortsBySwitch200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSwitchPortsBySwitch200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSwitchPortsBySwitch200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


