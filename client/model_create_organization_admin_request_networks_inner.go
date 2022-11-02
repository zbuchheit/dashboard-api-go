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

// CreateOrganizationAdminRequestNetworksInner struct for CreateOrganizationAdminRequestNetworksInner
type CreateOrganizationAdminRequestNetworksInner struct {
	// The network ID
	Id string `json:"id"`
	// The privilege of the dashboard administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

// NewCreateOrganizationAdminRequestNetworksInner instantiates a new CreateOrganizationAdminRequestNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdminRequestNetworksInner(id string, access string) *CreateOrganizationAdminRequestNetworksInner {
	this := CreateOrganizationAdminRequestNetworksInner{}
	this.Id = id
	this.Access = access
	return &this
}

// NewCreateOrganizationAdminRequestNetworksInnerWithDefaults instantiates a new CreateOrganizationAdminRequestNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdminRequestNetworksInnerWithDefaults() *CreateOrganizationAdminRequestNetworksInner {
	this := CreateOrganizationAdminRequestNetworksInner{}
	return &this
}

// GetId returns the Id field value
func (o *CreateOrganizationAdminRequestNetworksInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequestNetworksInner) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateOrganizationAdminRequestNetworksInner) SetId(v string) {
	o.Id = v
}

// GetAccess returns the Access field value
func (o *CreateOrganizationAdminRequestNetworksInner) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequestNetworksInner) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *CreateOrganizationAdminRequestNetworksInner) SetAccess(v string) {
	o.Access = v
}

func (o CreateOrganizationAdminRequestNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationAdminRequestNetworksInner struct {
	value *CreateOrganizationAdminRequestNetworksInner
	isSet bool
}

func (v NullableCreateOrganizationAdminRequestNetworksInner) Get() *CreateOrganizationAdminRequestNetworksInner {
	return v.value
}

func (v *NullableCreateOrganizationAdminRequestNetworksInner) Set(val *CreateOrganizationAdminRequestNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdminRequestNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdminRequestNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdminRequestNetworksInner(val *CreateOrganizationAdminRequestNetworksInner) *NullableCreateOrganizationAdminRequestNetworksInner {
	return &NullableCreateOrganizationAdminRequestNetworksInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdminRequestNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdminRequestNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


