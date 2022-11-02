/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkSensorAlertsProfiles200ResponseInnerSchedule The sensor schedule to use with the alert profile.
type GetNetworkSensorAlertsProfiles200ResponseInnerSchedule struct {
	// ID of the sensor schedule to use with the alert profile. If not defined, the alert profile will be active at all times.
	Id *string `json:"id,omitempty"`
	// Name of the sensor schedule to use with the alert profile.
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerSchedule instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerSchedule() *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerSchedule{}
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerScheduleWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerScheduleWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule(val *GetNetworkSensorAlertsProfiles200ResponseInnerSchedule) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


