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

// checks if the ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest{}

// ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct for ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
type ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct {
	// The subscription's claim key
	ClaimKey string `json:"claimKey"`
}

// NewValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest instantiates a new ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(claimKey string) *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	this := ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest{}
	this.ClaimKey = claimKey
	return &this
}

// NewValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequestWithDefaults instantiates a new ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequestWithDefaults() *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	this := ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest{}
	return &this
}

// GetClaimKey returns the ClaimKey field value
func (o *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) GetClaimKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClaimKey
}

// GetClaimKeyOk returns a tuple with the ClaimKey field value
// and a boolean to check if the value has been set.
func (o *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) GetClaimKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClaimKey, true
}

// SetClaimKey sets field value
func (o *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) SetClaimKey(v string) {
	o.ClaimKey = v
}

func (o ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["claimKey"] = o.ClaimKey
	return toSerialize, nil
}

type NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct {
	value *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest
	isSet bool
}

func (v NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) Get() *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	return v.value
}

func (v *NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) Set(val *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(val *ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) *NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest {
	return &NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest{value: val, isSet: true}
}

func (v NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


