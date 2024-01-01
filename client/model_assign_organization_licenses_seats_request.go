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

// checks if the AssignOrganizationLicensesSeatsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignOrganizationLicensesSeatsRequest{}

// AssignOrganizationLicensesSeatsRequest struct for AssignOrganizationLicensesSeatsRequest
type AssignOrganizationLicensesSeatsRequest struct {
	// The ID of the SM license to assign seats from
	LicenseId string `json:"licenseId"`
	// The ID of the SM network to assign the seats to
	NetworkId string `json:"networkId"`
	// The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license
	SeatCount int32 `json:"seatCount"`
}

// NewAssignOrganizationLicensesSeatsRequest instantiates a new AssignOrganizationLicensesSeatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignOrganizationLicensesSeatsRequest(licenseId string, networkId string, seatCount int32) *AssignOrganizationLicensesSeatsRequest {
	this := AssignOrganizationLicensesSeatsRequest{}
	this.LicenseId = licenseId
	this.NetworkId = networkId
	this.SeatCount = seatCount
	return &this
}

// NewAssignOrganizationLicensesSeatsRequestWithDefaults instantiates a new AssignOrganizationLicensesSeatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignOrganizationLicensesSeatsRequestWithDefaults() *AssignOrganizationLicensesSeatsRequest {
	this := AssignOrganizationLicensesSeatsRequest{}
	return &this
}

// GetLicenseId returns the LicenseId field value
func (o *AssignOrganizationLicensesSeatsRequest) GetLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value
// and a boolean to check if the value has been set.
func (o *AssignOrganizationLicensesSeatsRequest) GetLicenseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseId, true
}

// SetLicenseId sets field value
func (o *AssignOrganizationLicensesSeatsRequest) SetLicenseId(v string) {
	o.LicenseId = v
}

// GetNetworkId returns the NetworkId field value
func (o *AssignOrganizationLicensesSeatsRequest) GetNetworkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value
// and a boolean to check if the value has been set.
func (o *AssignOrganizationLicensesSeatsRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkId, true
}

// SetNetworkId sets field value
func (o *AssignOrganizationLicensesSeatsRequest) SetNetworkId(v string) {
	o.NetworkId = v
}

// GetSeatCount returns the SeatCount field value
func (o *AssignOrganizationLicensesSeatsRequest) GetSeatCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SeatCount
}

// GetSeatCountOk returns a tuple with the SeatCount field value
// and a boolean to check if the value has been set.
func (o *AssignOrganizationLicensesSeatsRequest) GetSeatCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeatCount, true
}

// SetSeatCount sets field value
func (o *AssignOrganizationLicensesSeatsRequest) SetSeatCount(v int32) {
	o.SeatCount = v
}

func (o AssignOrganizationLicensesSeatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignOrganizationLicensesSeatsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["licenseId"] = o.LicenseId
	toSerialize["networkId"] = o.NetworkId
	toSerialize["seatCount"] = o.SeatCount
	return toSerialize, nil
}

type NullableAssignOrganizationLicensesSeatsRequest struct {
	value *AssignOrganizationLicensesSeatsRequest
	isSet bool
}

func (v NullableAssignOrganizationLicensesSeatsRequest) Get() *AssignOrganizationLicensesSeatsRequest {
	return v.value
}

func (v *NullableAssignOrganizationLicensesSeatsRequest) Set(val *AssignOrganizationLicensesSeatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignOrganizationLicensesSeatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignOrganizationLicensesSeatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignOrganizationLicensesSeatsRequest(val *AssignOrganizationLicensesSeatsRequest) *NullableAssignOrganizationLicensesSeatsRequest {
	return &NullableAssignOrganizationLicensesSeatsRequest{value: val, isSet: true}
}

func (v NullableAssignOrganizationLicensesSeatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignOrganizationLicensesSeatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


