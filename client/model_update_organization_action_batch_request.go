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

// checks if the UpdateOrganizationActionBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationActionBatchRequest{}

// UpdateOrganizationActionBatchRequest struct for UpdateOrganizationActionBatchRequest
type UpdateOrganizationActionBatchRequest struct {
	// A boolean representing whether or not the batch has been confirmed. This property cannot be unset once it is true.
	Confirmed *bool `json:"confirmed,omitempty"`
	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch.
	Synchronous *bool `json:"synchronous,omitempty"`
}

// NewUpdateOrganizationActionBatchRequest instantiates a new UpdateOrganizationActionBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationActionBatchRequest() *UpdateOrganizationActionBatchRequest {
	this := UpdateOrganizationActionBatchRequest{}
	return &this
}

// NewUpdateOrganizationActionBatchRequestWithDefaults instantiates a new UpdateOrganizationActionBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationActionBatchRequestWithDefaults() *UpdateOrganizationActionBatchRequest {
	this := UpdateOrganizationActionBatchRequest{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *UpdateOrganizationActionBatchRequest) GetConfirmed() bool {
	if o == nil || IsNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationActionBatchRequest) GetConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmed) {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *UpdateOrganizationActionBatchRequest) HasConfirmed() bool {
	if o != nil && !IsNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *UpdateOrganizationActionBatchRequest) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *UpdateOrganizationActionBatchRequest) GetSynchronous() bool {
	if o == nil || IsNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationActionBatchRequest) GetSynchronousOk() (*bool, bool) {
	if o == nil || IsNil(o.Synchronous) {
		return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *UpdateOrganizationActionBatchRequest) HasSynchronous() bool {
	if o != nil && !IsNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *UpdateOrganizationActionBatchRequest) SetSynchronous(v bool) {
	o.Synchronous = &v
}

func (o UpdateOrganizationActionBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationActionBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !IsNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationActionBatchRequest struct {
	value *UpdateOrganizationActionBatchRequest
	isSet bool
}

func (v NullableUpdateOrganizationActionBatchRequest) Get() *UpdateOrganizationActionBatchRequest {
	return v.value
}

func (v *NullableUpdateOrganizationActionBatchRequest) Set(val *UpdateOrganizationActionBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationActionBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationActionBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationActionBatchRequest(val *UpdateOrganizationActionBatchRequest) *NullableUpdateOrganizationActionBatchRequest {
	return &NullableUpdateOrganizationActionBatchRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationActionBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationActionBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


