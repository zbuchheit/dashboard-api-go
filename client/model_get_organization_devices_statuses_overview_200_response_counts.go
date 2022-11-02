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

// GetOrganizationDevicesStatusesOverview200ResponseCounts counts
type GetOrganizationDevicesStatusesOverview200ResponseCounts struct {
	ByStatus *GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus `json:"byStatus,omitempty"`
}

// NewGetOrganizationDevicesStatusesOverview200ResponseCounts instantiates a new GetOrganizationDevicesStatusesOverview200ResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesStatusesOverview200ResponseCounts() *GetOrganizationDevicesStatusesOverview200ResponseCounts {
	this := GetOrganizationDevicesStatusesOverview200ResponseCounts{}
	return &this
}

// NewGetOrganizationDevicesStatusesOverview200ResponseCountsWithDefaults instantiates a new GetOrganizationDevicesStatusesOverview200ResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesStatusesOverview200ResponseCountsWithDefaults() *GetOrganizationDevicesStatusesOverview200ResponseCounts {
	this := GetOrganizationDevicesStatusesOverview200ResponseCounts{}
	return &this
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCounts) GetByStatus() GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCounts) GetByStatusOk() (*GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCounts) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus and assigns it to the ByStatus field.
func (o *GetOrganizationDevicesStatusesOverview200ResponseCounts) SetByStatus(v GetOrganizationDevicesStatusesOverview200ResponseCountsByStatus) {
	o.ByStatus = &v
}

func (o GetOrganizationDevicesStatusesOverview200ResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationDevicesStatusesOverview200ResponseCounts struct {
	value *GetOrganizationDevicesStatusesOverview200ResponseCounts
	isSet bool
}

func (v NullableGetOrganizationDevicesStatusesOverview200ResponseCounts) Get() *GetOrganizationDevicesStatusesOverview200ResponseCounts {
	return v.value
}

func (v *NullableGetOrganizationDevicesStatusesOverview200ResponseCounts) Set(val *GetOrganizationDevicesStatusesOverview200ResponseCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesStatusesOverview200ResponseCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesStatusesOverview200ResponseCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesStatusesOverview200ResponseCounts(val *GetOrganizationDevicesStatusesOverview200ResponseCounts) *NullableGetOrganizationDevicesStatusesOverview200ResponseCounts {
	return &NullableGetOrganizationDevicesStatusesOverview200ResponseCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesStatusesOverview200ResponseCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesStatusesOverview200ResponseCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


