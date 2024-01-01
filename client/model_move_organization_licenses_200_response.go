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

// checks if the MoveOrganizationLicenses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveOrganizationLicenses200Response{}

// MoveOrganizationLicenses200Response struct for MoveOrganizationLicenses200Response
type MoveOrganizationLicenses200Response struct {
	// The ID of the organization to move the licenses to
	DestOrganizationId *string `json:"destOrganizationId,omitempty"`
	// A list of IDs of licenses to move to the new organization
	LicenseIds []string `json:"licenseIds,omitempty"`
}

// NewMoveOrganizationLicenses200Response instantiates a new MoveOrganizationLicenses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveOrganizationLicenses200Response() *MoveOrganizationLicenses200Response {
	this := MoveOrganizationLicenses200Response{}
	return &this
}

// NewMoveOrganizationLicenses200ResponseWithDefaults instantiates a new MoveOrganizationLicenses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveOrganizationLicenses200ResponseWithDefaults() *MoveOrganizationLicenses200Response {
	this := MoveOrganizationLicenses200Response{}
	return &this
}

// GetDestOrganizationId returns the DestOrganizationId field value if set, zero value otherwise.
func (o *MoveOrganizationLicenses200Response) GetDestOrganizationId() string {
	if o == nil || IsNil(o.DestOrganizationId) {
		var ret string
		return ret
	}
	return *o.DestOrganizationId
}

// GetDestOrganizationIdOk returns a tuple with the DestOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicenses200Response) GetDestOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestOrganizationId) {
		return nil, false
	}
	return o.DestOrganizationId, true
}

// HasDestOrganizationId returns a boolean if a field has been set.
func (o *MoveOrganizationLicenses200Response) HasDestOrganizationId() bool {
	if o != nil && !IsNil(o.DestOrganizationId) {
		return true
	}

	return false
}

// SetDestOrganizationId gets a reference to the given string and assigns it to the DestOrganizationId field.
func (o *MoveOrganizationLicenses200Response) SetDestOrganizationId(v string) {
	o.DestOrganizationId = &v
}

// GetLicenseIds returns the LicenseIds field value if set, zero value otherwise.
func (o *MoveOrganizationLicenses200Response) GetLicenseIds() []string {
	if o == nil || IsNil(o.LicenseIds) {
		var ret []string
		return ret
	}
	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveOrganizationLicenses200Response) GetLicenseIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LicenseIds) {
		return nil, false
	}
	return o.LicenseIds, true
}

// HasLicenseIds returns a boolean if a field has been set.
func (o *MoveOrganizationLicenses200Response) HasLicenseIds() bool {
	if o != nil && !IsNil(o.LicenseIds) {
		return true
	}

	return false
}

// SetLicenseIds gets a reference to the given []string and assigns it to the LicenseIds field.
func (o *MoveOrganizationLicenses200Response) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

func (o MoveOrganizationLicenses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveOrganizationLicenses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DestOrganizationId) {
		toSerialize["destOrganizationId"] = o.DestOrganizationId
	}
	if !IsNil(o.LicenseIds) {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	return toSerialize, nil
}

type NullableMoveOrganizationLicenses200Response struct {
	value *MoveOrganizationLicenses200Response
	isSet bool
}

func (v NullableMoveOrganizationLicenses200Response) Get() *MoveOrganizationLicenses200Response {
	return v.value
}

func (v *NullableMoveOrganizationLicenses200Response) Set(val *MoveOrganizationLicenses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveOrganizationLicenses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveOrganizationLicenses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveOrganizationLicenses200Response(val *MoveOrganizationLicenses200Response) *NullableMoveOrganizationLicenses200Response {
	return &NullableMoveOrganizationLicenses200Response{value: val, isSet: true}
}

func (v NullableMoveOrganizationLicenses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveOrganizationLicenses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


