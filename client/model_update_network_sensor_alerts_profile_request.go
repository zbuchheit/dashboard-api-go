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

// checks if the UpdateNetworkSensorAlertsProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSensorAlertsProfileRequest{}

// UpdateNetworkSensorAlertsProfileRequest struct for UpdateNetworkSensorAlertsProfileRequest
type UpdateNetworkSensorAlertsProfileRequest struct {
	// Name of the sensor alert profile.
	Name *string `json:"name,omitempty"`
	Schedule *CreateNetworkSensorAlertsProfileRequestSchedule `json:"schedule,omitempty"`
	// List of conditions that will cause the profile to send an alert.
	Conditions []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner `json:"conditions,omitempty"`
	Recipients *GetNetworkSensorAlertsProfiles200ResponseInnerRecipients `json:"recipients,omitempty"`
	// List of device serials assigned to this sensor alert profile.
	Serials []string `json:"serials,omitempty"`
}

// NewUpdateNetworkSensorAlertsProfileRequest instantiates a new UpdateNetworkSensorAlertsProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSensorAlertsProfileRequest() *UpdateNetworkSensorAlertsProfileRequest {
	this := UpdateNetworkSensorAlertsProfileRequest{}
	return &this
}

// NewUpdateNetworkSensorAlertsProfileRequestWithDefaults instantiates a new UpdateNetworkSensorAlertsProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSensorAlertsProfileRequestWithDefaults() *UpdateNetworkSensorAlertsProfileRequest {
	this := UpdateNetworkSensorAlertsProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkSensorAlertsProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetSchedule() CreateNetworkSensorAlertsProfileRequestSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret CreateNetworkSensorAlertsProfileRequestSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetScheduleOk() (*CreateNetworkSensorAlertsProfileRequestSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CreateNetworkSensorAlertsProfileRequestSchedule and assigns it to the Schedule field.
func (o *UpdateNetworkSensorAlertsProfileRequest) SetSchedule(v CreateNetworkSensorAlertsProfileRequestSchedule) {
	o.Schedule = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetConditions() []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	if o == nil || IsNil(o.Conditions) {
		var ret []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetConditionsOk() ([]GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner and assigns it to the Conditions field.
func (o *UpdateNetworkSensorAlertsProfileRequest) SetConditions(v []GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) {
	o.Conditions = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetRecipients() GetNetworkSensorAlertsProfiles200ResponseInnerRecipients {
	if o == nil || IsNil(o.Recipients) {
		var ret GetNetworkSensorAlertsProfiles200ResponseInnerRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetRecipientsOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerRecipients, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given GetNetworkSensorAlertsProfiles200ResponseInnerRecipients and assigns it to the Recipients field.
func (o *UpdateNetworkSensorAlertsProfileRequest) SetRecipients(v GetNetworkSensorAlertsProfiles200ResponseInnerRecipients) {
	o.Recipients = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *UpdateNetworkSensorAlertsProfileRequest) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *UpdateNetworkSensorAlertsProfileRequest) SetSerials(v []string) {
	o.Serials = v
}

func (o UpdateNetworkSensorAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSensorAlertsProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSensorAlertsProfileRequest struct {
	value *UpdateNetworkSensorAlertsProfileRequest
	isSet bool
}

func (v NullableUpdateNetworkSensorAlertsProfileRequest) Get() *UpdateNetworkSensorAlertsProfileRequest {
	return v.value
}

func (v *NullableUpdateNetworkSensorAlertsProfileRequest) Set(val *UpdateNetworkSensorAlertsProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSensorAlertsProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSensorAlertsProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSensorAlertsProfileRequest(val *UpdateNetworkSensorAlertsProfileRequest) *NullableUpdateNetworkSensorAlertsProfileRequest {
	return &NullableUpdateNetworkSensorAlertsProfileRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSensorAlertsProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSensorAlertsProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


