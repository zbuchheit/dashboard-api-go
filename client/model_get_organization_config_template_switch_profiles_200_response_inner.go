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

// checks if the GetOrganizationConfigTemplateSwitchProfiles200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationConfigTemplateSwitchProfiles200ResponseInner{}

// GetOrganizationConfigTemplateSwitchProfiles200ResponseInner struct for GetOrganizationConfigTemplateSwitchProfiles200ResponseInner
type GetOrganizationConfigTemplateSwitchProfiles200ResponseInner struct {
	// Switch template id
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// Switch template name
	Name *string `json:"name,omitempty"`
	// Switch model
	Model *string `json:"model,omitempty"`
}

// NewGetOrganizationConfigTemplateSwitchProfiles200ResponseInner instantiates a new GetOrganizationConfigTemplateSwitchProfiles200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationConfigTemplateSwitchProfiles200ResponseInner() *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner {
	this := GetOrganizationConfigTemplateSwitchProfiles200ResponseInner{}
	return &this
}

// NewGetOrganizationConfigTemplateSwitchProfiles200ResponseInnerWithDefaults instantiates a new GetOrganizationConfigTemplateSwitchProfiles200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationConfigTemplateSwitchProfiles200ResponseInnerWithDefaults() *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner {
	this := GetOrganizationConfigTemplateSwitchProfiles200ResponseInner{}
	return &this
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) GetSwitchProfileId() string {
	if o == nil || IsNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwitchProfileId) {
		return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) HasSwitchProfileId() bool {
	if o != nil && !IsNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) SetModel(v string) {
	o.Model = &v
}

func (o GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SwitchProfileId) {
		toSerialize["switchProfileId"] = o.SwitchProfileId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	return toSerialize, nil
}

type NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner struct {
	value *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner) Get() *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner) Set(val *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner(val *GetOrganizationConfigTemplateSwitchProfiles200ResponseInner) *NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner {
	return &NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationConfigTemplateSwitchProfiles200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


