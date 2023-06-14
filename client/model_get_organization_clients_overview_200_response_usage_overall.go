/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationClientsOverview200ResponseUsageOverall Overall data usage of all clients across organization
type GetOrganizationClientsOverview200ResponseUsageOverall struct {
	// Total data usage (in kb) of all clients across organization
	Total *float32 `json:"total,omitempty"`
	// Downstream data usage (in kb) of all clients across organization
	Downstream *float32 `json:"downstream,omitempty"`
	// Upstream data usage (in kb) of all clients across organization
	Upstream *float32 `json:"upstream,omitempty"`
}

// NewGetOrganizationClientsOverview200ResponseUsageOverall instantiates a new GetOrganizationClientsOverview200ResponseUsageOverall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationClientsOverview200ResponseUsageOverall() *GetOrganizationClientsOverview200ResponseUsageOverall {
	this := GetOrganizationClientsOverview200ResponseUsageOverall{}
	return &this
}

// NewGetOrganizationClientsOverview200ResponseUsageOverallWithDefaults instantiates a new GetOrganizationClientsOverview200ResponseUsageOverall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationClientsOverview200ResponseUsageOverallWithDefaults() *GetOrganizationClientsOverview200ResponseUsageOverall {
	this := GetOrganizationClientsOverview200ResponseUsageOverall{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) GetTotal() float32 {
	if o == nil || isNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) GetTotalOk() (*float32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) SetTotal(v float32) {
	o.Total = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) GetDownstream() float32 {
	if o == nil || isNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) GetDownstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Downstream) {
    return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) HasDownstream() bool {
	if o != nil && !isNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) SetDownstream(v float32) {
	o.Downstream = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) GetUpstream() float32 {
	if o == nil || isNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) GetUpstreamOk() (*float32, bool) {
	if o == nil || isNil(o.Upstream) {
    return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) HasUpstream() bool {
	if o != nil && !isNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *GetOrganizationClientsOverview200ResponseUsageOverall) SetUpstream(v float32) {
	o.Upstream = &v
}

func (o GetOrganizationClientsOverview200ResponseUsageOverall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	if !isNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationClientsOverview200ResponseUsageOverall struct {
	value *GetOrganizationClientsOverview200ResponseUsageOverall
	isSet bool
}

func (v NullableGetOrganizationClientsOverview200ResponseUsageOverall) Get() *GetOrganizationClientsOverview200ResponseUsageOverall {
	return v.value
}

func (v *NullableGetOrganizationClientsOverview200ResponseUsageOverall) Set(val *GetOrganizationClientsOverview200ResponseUsageOverall) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationClientsOverview200ResponseUsageOverall) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationClientsOverview200ResponseUsageOverall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationClientsOverview200ResponseUsageOverall(val *GetOrganizationClientsOverview200ResponseUsageOverall) *NullableGetOrganizationClientsOverview200ResponseUsageOverall {
	return &NullableGetOrganizationClientsOverview200ResponseUsageOverall{value: val, isSet: true}
}

func (v NullableGetOrganizationClientsOverview200ResponseUsageOverall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationClientsOverview200ResponseUsageOverall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


