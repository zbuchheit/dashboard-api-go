/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner{}

// GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner struct for GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
type GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner struct {
	// The start time of the access period
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the access period
	EndTs *time.Time `json:"endTs,omitempty"`
	// list of response codes and a count of how many requests had that code in the given time period
	Counts []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner `json:"counts,omitempty"`
}

// NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner instantiates a new GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner {
	this := GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner{}
	return &this
}

// NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerWithDefaults instantiates a new GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerWithDefaults() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner {
	this := GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetCounts() []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner {
	if o == nil || IsNil(o.Counts) {
		var ret []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner
		return ret
	}
	return o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) GetCountsOk() ([]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner, bool) {
	if o == nil || IsNil(o.Counts) {
		return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) HasCounts() bool {
	if o != nil && !IsNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner and assigns it to the Counts field.
func (o *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) SetCounts(v []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInnerCountsInner) {
	o.Counts = v
}

func (o GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return toSerialize, nil
}

type NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner struct {
	value *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) Get() *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) Set(val *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner(val *GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner {
	return &NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


