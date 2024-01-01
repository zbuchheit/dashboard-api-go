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

// checks if the GetOrganizationInsightApplications200ResponseInnerThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationInsightApplications200ResponseInnerThresholds{}

// GetOrganizationInsightApplications200ResponseInnerThresholds Thresholds defined by a user or Meraki models for each application
type GetOrganizationInsightApplications200ResponseInnerThresholds struct {
	// Threshold type (static or smart)
	Type *string `json:"type,omitempty"`
	// Threshold for each network
	ByNetwork []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner `json:"byNetwork,omitempty"`
}

// NewGetOrganizationInsightApplications200ResponseInnerThresholds instantiates a new GetOrganizationInsightApplications200ResponseInnerThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationInsightApplications200ResponseInnerThresholds() *GetOrganizationInsightApplications200ResponseInnerThresholds {
	this := GetOrganizationInsightApplications200ResponseInnerThresholds{}
	return &this
}

// NewGetOrganizationInsightApplications200ResponseInnerThresholdsWithDefaults instantiates a new GetOrganizationInsightApplications200ResponseInnerThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationInsightApplications200ResponseInnerThresholdsWithDefaults() *GetOrganizationInsightApplications200ResponseInnerThresholds {
	this := GetOrganizationInsightApplications200ResponseInnerThresholds{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) SetType(v string) {
	o.Type = &v
}

// GetByNetwork returns the ByNetwork field value if set, zero value otherwise.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetByNetwork() []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner {
	if o == nil || IsNil(o.ByNetwork) {
		var ret []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner
		return ret
	}
	return o.ByNetwork
}

// GetByNetworkOk returns a tuple with the ByNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) GetByNetworkOk() ([]GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner, bool) {
	if o == nil || IsNil(o.ByNetwork) {
		return nil, false
	}
	return o.ByNetwork, true
}

// HasByNetwork returns a boolean if a field has been set.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) HasByNetwork() bool {
	if o != nil && !IsNil(o.ByNetwork) {
		return true
	}

	return false
}

// SetByNetwork gets a reference to the given []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner and assigns it to the ByNetwork field.
func (o *GetOrganizationInsightApplications200ResponseInnerThresholds) SetByNetwork(v []GetOrganizationInsightApplications200ResponseInnerThresholdsByNetworkInner) {
	o.ByNetwork = v
}

func (o GetOrganizationInsightApplications200ResponseInnerThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationInsightApplications200ResponseInnerThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ByNetwork) {
		toSerialize["byNetwork"] = o.ByNetwork
	}
	return toSerialize, nil
}

type NullableGetOrganizationInsightApplications200ResponseInnerThresholds struct {
	value *GetOrganizationInsightApplications200ResponseInnerThresholds
	isSet bool
}

func (v NullableGetOrganizationInsightApplications200ResponseInnerThresholds) Get() *GetOrganizationInsightApplications200ResponseInnerThresholds {
	return v.value
}

func (v *NullableGetOrganizationInsightApplications200ResponseInnerThresholds) Set(val *GetOrganizationInsightApplications200ResponseInnerThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationInsightApplications200ResponseInnerThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationInsightApplications200ResponseInnerThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationInsightApplications200ResponseInnerThresholds(val *GetOrganizationInsightApplications200ResponseInnerThresholds) *NullableGetOrganizationInsightApplications200ResponseInnerThresholds {
	return &NullableGetOrganizationInsightApplications200ResponseInnerThresholds{value: val, isSet: true}
}

func (v NullableGetOrganizationInsightApplications200ResponseInnerThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationInsightApplications200ResponseInnerThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


