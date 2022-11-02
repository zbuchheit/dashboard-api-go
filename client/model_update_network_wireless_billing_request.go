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

// UpdateNetworkWirelessBillingRequest struct for UpdateNetworkWirelessBillingRequest
type UpdateNetworkWirelessBillingRequest struct {
	// The currency code of this node group's billing plans
	Currency *string `json:"currency,omitempty"`
	// Array of billing plans in the node group. (Can configure a maximum of 5)
	Plans []UpdateNetworkWirelessBillingRequestPlansInner `json:"plans,omitempty"`
}

// NewUpdateNetworkWirelessBillingRequest instantiates a new UpdateNetworkWirelessBillingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessBillingRequest() *UpdateNetworkWirelessBillingRequest {
	this := UpdateNetworkWirelessBillingRequest{}
	return &this
}

// NewUpdateNetworkWirelessBillingRequestWithDefaults instantiates a new UpdateNetworkWirelessBillingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessBillingRequestWithDefaults() *UpdateNetworkWirelessBillingRequest {
	this := UpdateNetworkWirelessBillingRequest{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessBillingRequest) GetCurrency() string {
	if o == nil || isNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || isNil(o.Currency) {
    return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessBillingRequest) HasCurrency() bool {
	if o != nil && !isNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UpdateNetworkWirelessBillingRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessBillingRequest) GetPlans() []UpdateNetworkWirelessBillingRequestPlansInner {
	if o == nil || isNil(o.Plans) {
		var ret []UpdateNetworkWirelessBillingRequestPlansInner
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessBillingRequest) GetPlansOk() ([]UpdateNetworkWirelessBillingRequestPlansInner, bool) {
	if o == nil || isNil(o.Plans) {
    return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessBillingRequest) HasPlans() bool {
	if o != nil && !isNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []UpdateNetworkWirelessBillingRequestPlansInner and assigns it to the Plans field.
func (o *UpdateNetworkWirelessBillingRequest) SetPlans(v []UpdateNetworkWirelessBillingRequestPlansInner) {
	o.Plans = v
}

func (o UpdateNetworkWirelessBillingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !isNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessBillingRequest struct {
	value *UpdateNetworkWirelessBillingRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessBillingRequest) Get() *UpdateNetworkWirelessBillingRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessBillingRequest) Set(val *UpdateNetworkWirelessBillingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessBillingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessBillingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessBillingRequest(val *UpdateNetworkWirelessBillingRequest) *NullableUpdateNetworkWirelessBillingRequest {
	return &NullableUpdateNetworkWirelessBillingRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessBillingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessBillingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


