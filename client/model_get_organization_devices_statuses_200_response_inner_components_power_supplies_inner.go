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

// checks if the GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner{}

// GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner struct for GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner
type GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner struct {
	// Slot the power supply is in
	Slot *int32 `json:"slot,omitempty"`
	// Serial of the power supply
	Serial *string `json:"serial,omitempty"`
	// Model of the power supply
	Model *string `json:"model,omitempty"`
	// Status of the power supply
	Status *string `json:"status,omitempty"`
	Poe *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerPoe `json:"poe,omitempty"`
}

// NewGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner instantiates a new GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner() *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner {
	this := GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner{}
	return &this
}

// NewGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerWithDefaults instantiates a new GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerWithDefaults() *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner {
	this := GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner{}
	return &this
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetSlot() int32 {
	if o == nil || IsNil(o.Slot) {
		var ret int32
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetSlotOk() (*int32, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given int32 and assigns it to the Slot field.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) SetSlot(v int32) {
	o.Slot = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) SetModel(v string) {
	o.Model = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) SetStatus(v string) {
	o.Status = &v
}

// GetPoe returns the Poe field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetPoe() GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerPoe {
	if o == nil || IsNil(o.Poe) {
		var ret GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerPoe
		return ret
	}
	return *o.Poe
}

// GetPoeOk returns a tuple with the Poe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) GetPoeOk() (*GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerPoe, bool) {
	if o == nil || IsNil(o.Poe) {
		return nil, false
	}
	return o.Poe, true
}

// HasPoe returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) HasPoe() bool {
	if o != nil && !IsNil(o.Poe) {
		return true
	}

	return false
}

// SetPoe gets a reference to the given GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerPoe and assigns it to the Poe field.
func (o *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) SetPoe(v GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInnerPoe) {
	o.Poe = &v
}

func (o GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Poe) {
		toSerialize["poe"] = o.Poe
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner struct {
	value *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner
	isSet bool
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) Get() *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) Set(val *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner(val *GetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) *NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner {
	return &NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesStatuses200ResponseInnerComponentsPowerSuppliesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


