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

// UpdateNetworkAppliancePrefixesDelegatedStaticRequest struct for UpdateNetworkAppliancePrefixesDelegatedStaticRequest
type UpdateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	// A static IPv6 prefix
	Prefix *string `json:"prefix,omitempty"`
	Origin *CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin `json:"origin,omitempty"`
	// A name or description for the prefix
	Description *string `json:"description,omitempty"`
}

// NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest instantiates a new UpdateNetworkAppliancePrefixesDelegatedStaticRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest() *UpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	this := UpdateNetworkAppliancePrefixesDelegatedStaticRequest{}
	return &this
}

// NewUpdateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults instantiates a new UpdateNetworkAppliancePrefixesDelegatedStaticRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkAppliancePrefixesDelegatedStaticRequestWithDefaults() *UpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	this := UpdateNetworkAppliancePrefixesDelegatedStaticRequest{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) SetPrefix(v string) {
	o.Prefix = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin {
	if o == nil || isNil(o.Origin) {
		var ret CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool) {
	if o == nil || isNil(o.Origin) {
    return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) HasOrigin() bool {
	if o != nil && !isNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin and assigns it to the Origin field.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin) {
	o.Origin = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateNetworkAppliancePrefixesDelegatedStaticRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct {
	value *UpdateNetworkAppliancePrefixesDelegatedStaticRequest
	isSet bool
}

func (v NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest) Get() *UpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	return v.value
}

func (v *NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest) Set(val *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest(val *UpdateNetworkAppliancePrefixesDelegatedStaticRequest) *NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest {
	return &NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkAppliancePrefixesDelegatedStaticRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


