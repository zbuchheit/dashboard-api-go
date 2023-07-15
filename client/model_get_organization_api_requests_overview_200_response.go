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

// checks if the GetOrganizationApiRequestsOverview200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApiRequestsOverview200Response{}

// GetOrganizationApiRequestsOverview200Response struct for GetOrganizationApiRequestsOverview200Response
type GetOrganizationApiRequestsOverview200Response struct {
	ResponseCodeCounts *GetOrganizationApiRequestsOverview200ResponseResponseCodeCounts `json:"responseCodeCounts,omitempty"`
}

// NewGetOrganizationApiRequestsOverview200Response instantiates a new GetOrganizationApiRequestsOverview200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApiRequestsOverview200Response() *GetOrganizationApiRequestsOverview200Response {
	this := GetOrganizationApiRequestsOverview200Response{}
	return &this
}

// NewGetOrganizationApiRequestsOverview200ResponseWithDefaults instantiates a new GetOrganizationApiRequestsOverview200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApiRequestsOverview200ResponseWithDefaults() *GetOrganizationApiRequestsOverview200Response {
	this := GetOrganizationApiRequestsOverview200Response{}
	return &this
}

// GetResponseCodeCounts returns the ResponseCodeCounts field value if set, zero value otherwise.
func (o *GetOrganizationApiRequestsOverview200Response) GetResponseCodeCounts() GetOrganizationApiRequestsOverview200ResponseResponseCodeCounts {
	if o == nil || IsNil(o.ResponseCodeCounts) {
		var ret GetOrganizationApiRequestsOverview200ResponseResponseCodeCounts
		return ret
	}
	return *o.ResponseCodeCounts
}

// GetResponseCodeCountsOk returns a tuple with the ResponseCodeCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApiRequestsOverview200Response) GetResponseCodeCountsOk() (*GetOrganizationApiRequestsOverview200ResponseResponseCodeCounts, bool) {
	if o == nil || IsNil(o.ResponseCodeCounts) {
		return nil, false
	}
	return o.ResponseCodeCounts, true
}

// HasResponseCodeCounts returns a boolean if a field has been set.
func (o *GetOrganizationApiRequestsOverview200Response) HasResponseCodeCounts() bool {
	if o != nil && !IsNil(o.ResponseCodeCounts) {
		return true
	}

	return false
}

// SetResponseCodeCounts gets a reference to the given GetOrganizationApiRequestsOverview200ResponseResponseCodeCounts and assigns it to the ResponseCodeCounts field.
func (o *GetOrganizationApiRequestsOverview200Response) SetResponseCodeCounts(v GetOrganizationApiRequestsOverview200ResponseResponseCodeCounts) {
	o.ResponseCodeCounts = &v
}

func (o GetOrganizationApiRequestsOverview200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApiRequestsOverview200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResponseCodeCounts) {
		toSerialize["responseCodeCounts"] = o.ResponseCodeCounts
	}
	return toSerialize, nil
}

type NullableGetOrganizationApiRequestsOverview200Response struct {
	value *GetOrganizationApiRequestsOverview200Response
	isSet bool
}

func (v NullableGetOrganizationApiRequestsOverview200Response) Get() *GetOrganizationApiRequestsOverview200Response {
	return v.value
}

func (v *NullableGetOrganizationApiRequestsOverview200Response) Set(val *GetOrganizationApiRequestsOverview200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApiRequestsOverview200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApiRequestsOverview200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApiRequestsOverview200Response(val *GetOrganizationApiRequestsOverview200Response) *NullableGetOrganizationApiRequestsOverview200Response {
	return &NullableGetOrganizationApiRequestsOverview200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationApiRequestsOverview200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApiRequestsOverview200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


