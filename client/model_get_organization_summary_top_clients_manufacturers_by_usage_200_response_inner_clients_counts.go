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

// GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts Counts of clients
type GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts struct {
	// Total counts of clients
	Total *int32 `json:"total,omitempty"`
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts{}
	return &this
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCountsWithDefaults instantiates a new GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCountsWithDefaults() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts {
	this := GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) SetTotal(v int32) {
	o.Total = &v
}

func (o GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts struct {
	value *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts
	isSet bool
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) Get() *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) Set(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts(val *GetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts {
	return &NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopClientsManufacturersByUsage200ResponseInnerClientsCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


