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

// CreateOrganizationSamlRoleRequestNetworksInner struct for CreateOrganizationSamlRoleRequestNetworksInner
type CreateOrganizationSamlRoleRequestNetworksInner struct {
	// The network ID
	Id string `json:"id"`
	// The privilege of the SAML administrator on the network. Can be one of 'full', 'read-only', 'guest-ambassador', 'monitor-only' or 'ssid-admin'
	Access string `json:"access"`
}

// NewCreateOrganizationSamlRoleRequestNetworksInner instantiates a new CreateOrganizationSamlRoleRequestNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationSamlRoleRequestNetworksInner(id string, access string) *CreateOrganizationSamlRoleRequestNetworksInner {
	this := CreateOrganizationSamlRoleRequestNetworksInner{}
	this.Id = id
	this.Access = access
	return &this
}

// NewCreateOrganizationSamlRoleRequestNetworksInnerWithDefaults instantiates a new CreateOrganizationSamlRoleRequestNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationSamlRoleRequestNetworksInnerWithDefaults() *CreateOrganizationSamlRoleRequestNetworksInner {
	this := CreateOrganizationSamlRoleRequestNetworksInner{}
	return &this
}

// GetId returns the Id field value
func (o *CreateOrganizationSamlRoleRequestNetworksInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequestNetworksInner) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateOrganizationSamlRoleRequestNetworksInner) SetId(v string) {
	o.Id = v
}

// GetAccess returns the Access field value
func (o *CreateOrganizationSamlRoleRequestNetworksInner) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequestNetworksInner) GetAccessOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *CreateOrganizationSamlRoleRequestNetworksInner) SetAccess(v string) {
	o.Access = v
}

func (o CreateOrganizationSamlRoleRequestNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrganizationSamlRoleRequestNetworksInner struct {
	value *CreateOrganizationSamlRoleRequestNetworksInner
	isSet bool
}

func (v NullableCreateOrganizationSamlRoleRequestNetworksInner) Get() *CreateOrganizationSamlRoleRequestNetworksInner {
	return v.value
}

func (v *NullableCreateOrganizationSamlRoleRequestNetworksInner) Set(val *CreateOrganizationSamlRoleRequestNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationSamlRoleRequestNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationSamlRoleRequestNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationSamlRoleRequestNetworksInner(val *CreateOrganizationSamlRoleRequestNetworksInner) *NullableCreateOrganizationSamlRoleRequestNetworksInner {
	return &NullableCreateOrganizationSamlRoleRequestNetworksInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationSamlRoleRequestNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationSamlRoleRequestNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


