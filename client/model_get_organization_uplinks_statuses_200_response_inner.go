/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetOrganizationUplinksStatuses200ResponseInner struct for GetOrganizationUplinksStatuses200ResponseInner
type GetOrganizationUplinksStatuses200ResponseInner struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// The uplink serial
	Serial *string `json:"serial,omitempty"`
	// The uplink model
	Model *string `json:"model,omitempty"`
	// Last reported time for the device
	LastReportedAt *time.Time `json:"lastReportedAt,omitempty"`
	// Uplinks
	Uplinks []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner `json:"uplinks,omitempty"`
}

// NewGetOrganizationUplinksStatuses200ResponseInner instantiates a new GetOrganizationUplinksStatuses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationUplinksStatuses200ResponseInner() *GetOrganizationUplinksStatuses200ResponseInner {
	this := GetOrganizationUplinksStatuses200ResponseInner{}
	return &this
}

// NewGetOrganizationUplinksStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationUplinksStatuses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationUplinksStatuses200ResponseInnerWithDefaults() *GetOrganizationUplinksStatuses200ResponseInner {
	this := GetOrganizationUplinksStatuses200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationUplinksStatuses200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationUplinksStatuses200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationUplinksStatuses200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetLastReportedAt() time.Time {
	if o == nil || isNil(o.LastReportedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetLastReportedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastReportedAt) {
    return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) HasLastReportedAt() bool {
	if o != nil && !isNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given time.Time and assigns it to the LastReportedAt field.
func (o *GetOrganizationUplinksStatuses200ResponseInner) SetLastReportedAt(v time.Time) {
	o.LastReportedAt = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetUplinks() []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner {
	if o == nil || isNil(o.Uplinks) {
		var ret []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) GetUplinksOk() ([]GetOrganizationUplinksStatuses200ResponseInnerUplinksInner, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *GetOrganizationUplinksStatuses200ResponseInner) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner and assigns it to the Uplinks field.
func (o *GetOrganizationUplinksStatuses200ResponseInner) SetUplinks(v []GetOrganizationUplinksStatuses200ResponseInnerUplinksInner) {
	o.Uplinks = v
}

func (o GetOrganizationUplinksStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !isNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationUplinksStatuses200ResponseInner struct {
	value *GetOrganizationUplinksStatuses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationUplinksStatuses200ResponseInner) Get() *GetOrganizationUplinksStatuses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationUplinksStatuses200ResponseInner) Set(val *GetOrganizationUplinksStatuses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationUplinksStatuses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationUplinksStatuses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationUplinksStatuses200ResponseInner(val *GetOrganizationUplinksStatuses200ResponseInner) *NullableGetOrganizationUplinksStatuses200ResponseInner {
	return &NullableGetOrganizationUplinksStatuses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationUplinksStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationUplinksStatuses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


