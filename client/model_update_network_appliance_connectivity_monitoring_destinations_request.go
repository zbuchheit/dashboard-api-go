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

// checks if the UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest{}

// UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct for UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest
type UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct {
	// The list of connectivity monitoring destinations
	Destinations []UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner `json:"destinations,omitempty"`
}

// NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest instantiates a new UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest() *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest {
	this := UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest{}
	return &this
}

// NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestWithDefaults instantiates a new UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequestWithDefaults() *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest {
	this := UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) GetDestinations() []UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner {
	if o == nil || IsNil(o.Destinations) {
		var ret []UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) GetDestinationsOk() ([]UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner, bool) {
	if o == nil || IsNil(o.Destinations) {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) HasDestinations() bool {
	if o != nil && !IsNil(o.Destinations) {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner and assigns it to the Destinations field.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) SetDestinations(v []UpdateNetworkApplianceConnectivityMonitoringDestinationsRequestDestinationsInner) {
	o.Destinations = v
}

func (o UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Destinations) {
		toSerialize["destinations"] = o.Destinations
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct {
	value *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) Get() *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) Set(val *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest(val *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest {
	return &NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


