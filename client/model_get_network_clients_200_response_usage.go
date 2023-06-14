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

// GetNetworkClients200ResponseUsage Usage, sent and received
type GetNetworkClients200ResponseUsage struct {
	// Usage sent by the client
	Sent *float32 `json:"sent,omitempty"`
	// Usage received by the client
	Recv *float32 `json:"recv,omitempty"`
}

// NewGetNetworkClients200ResponseUsage instantiates a new GetNetworkClients200ResponseUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkClients200ResponseUsage() *GetNetworkClients200ResponseUsage {
	this := GetNetworkClients200ResponseUsage{}
	return &this
}

// NewGetNetworkClients200ResponseUsageWithDefaults instantiates a new GetNetworkClients200ResponseUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkClients200ResponseUsageWithDefaults() *GetNetworkClients200ResponseUsage {
	this := GetNetworkClients200ResponseUsage{}
	return &this
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *GetNetworkClients200ResponseUsage) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200ResponseUsage) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *GetNetworkClients200ResponseUsage) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *GetNetworkClients200ResponseUsage) SetSent(v float32) {
	o.Sent = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *GetNetworkClients200ResponseUsage) GetRecv() float32 {
	if o == nil || isNil(o.Recv) {
		var ret float32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkClients200ResponseUsage) GetRecvOk() (*float32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *GetNetworkClients200ResponseUsage) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given float32 and assigns it to the Recv field.
func (o *GetNetworkClients200ResponseUsage) SetRecv(v float32) {
	o.Recv = &v
}

func (o GetNetworkClients200ResponseUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkClients200ResponseUsage struct {
	value *GetNetworkClients200ResponseUsage
	isSet bool
}

func (v NullableGetNetworkClients200ResponseUsage) Get() *GetNetworkClients200ResponseUsage {
	return v.value
}

func (v *NullableGetNetworkClients200ResponseUsage) Set(val *GetNetworkClients200ResponseUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkClients200ResponseUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkClients200ResponseUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkClients200ResponseUsage(val *GetNetworkClients200ResponseUsage) *NullableGetNetworkClients200ResponseUsage {
	return &NullableGetNetworkClients200ResponseUsage{value: val, isSet: true}
}

func (v NullableGetNetworkClients200ResponseUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkClients200ResponseUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


