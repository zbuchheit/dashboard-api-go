/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage{}

// GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage Clients usage
type GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage struct {
	// Total data usage by client
	Total *float32 `json:"total,omitempty"`
	// Upstream data usage by client
	Upstream *float32 `json:"upstream,omitempty"`
	// Downstream data usage by client
	Downstream *float32 `json:"downstream,omitempty"`
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage{}
	return &this
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsageWithDefaults instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsageWithDefaults() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) SetTotal(v float32) {
	o.Total = &v
}

// GetUpstream returns the Upstream field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) GetUpstream() float32 {
	if o == nil || IsNil(o.Upstream) {
		var ret float32
		return ret
	}
	return *o.Upstream
}

// GetUpstreamOk returns a tuple with the Upstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) GetUpstreamOk() (*float32, bool) {
	if o == nil || IsNil(o.Upstream) {
		return nil, false
	}
	return o.Upstream, true
}

// HasUpstream returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) HasUpstream() bool {
	if o != nil && !IsNil(o.Upstream) {
		return true
	}

	return false
}

// SetUpstream gets a reference to the given float32 and assigns it to the Upstream field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) SetUpstream(v float32) {
	o.Upstream = &v
}

// GetDownstream returns the Downstream field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) GetDownstream() float32 {
	if o == nil || IsNil(o.Downstream) {
		var ret float32
		return ret
	}
	return *o.Downstream
}

// GetDownstreamOk returns a tuple with the Downstream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) GetDownstreamOk() (*float32, bool) {
	if o == nil || IsNil(o.Downstream) {
		return nil, false
	}
	return o.Downstream, true
}

// HasDownstream returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) HasDownstream() bool {
	if o != nil && !IsNil(o.Downstream) {
		return true
	}

	return false
}

// SetDownstream gets a reference to the given float32 and assigns it to the Downstream field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) SetDownstream(v float32) {
	o.Downstream = &v
}

func (o GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Upstream) {
		toSerialize["upstream"] = o.Upstream
	}
	if !IsNil(o.Downstream) {
		toSerialize["downstream"] = o.Downstream
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage struct {
	value *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage
	isSet bool
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) Get() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) Set(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage {
	return &NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


