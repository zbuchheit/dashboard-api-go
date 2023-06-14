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

// UpdateNetworkApplianceVpnSiteToSiteVpnRequest struct for UpdateNetworkApplianceVpnSiteToSiteVpnRequest
type UpdateNetworkApplianceVpnSiteToSiteVpnRequest struct {
	// The site-to-site VPN mode. Can be one of 'none', 'spoke' or 'hub'
	Mode string `json:"mode"`
	// The list of VPN hubs, in order of preference. In spoke mode, at least 1 hub is required.
	Hubs []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner `json:"hubs,omitempty"`
	// The list of subnets and their VPN presence.
	Subnets []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner `json:"subnets,omitempty"`
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest(mode string) *UpdateNetworkApplianceVpnSiteToSiteVpnRequest {
	this := UpdateNetworkApplianceVpnSiteToSiteVpnRequest{}
	this.Mode = mode
	return &this
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestWithDefaults instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestWithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnRequest {
	this := UpdateNetworkApplianceVpnSiteToSiteVpnRequest{}
	return &this
}

// GetMode returns the Mode field value
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetModeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SetMode(v string) {
	o.Mode = v
}

// GetHubs returns the Hubs field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetHubs() []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner {
	if o == nil || isNil(o.Hubs) {
		var ret []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner
		return ret
	}
	return o.Hubs
}

// GetHubsOk returns a tuple with the Hubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetHubsOk() ([]UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner, bool) {
	if o == nil || isNil(o.Hubs) {
    return nil, false
	}
	return o.Hubs, true
}

// HasHubs returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) HasHubs() bool {
	if o != nil && !isNil(o.Hubs) {
		return true
	}

	return false
}

// SetHubs gets a reference to the given []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner and assigns it to the Hubs field.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SetHubs(v []UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) {
	o.Hubs = v
}

// GetSubnets returns the Subnets field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetSubnets() []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner {
	if o == nil || isNil(o.Subnets) {
		var ret []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner
		return ret
	}
	return o.Subnets
}

// GetSubnetsOk returns a tuple with the Subnets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) GetSubnetsOk() ([]UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner, bool) {
	if o == nil || isNil(o.Subnets) {
    return nil, false
	}
	return o.Subnets, true
}

// HasSubnets returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) HasSubnets() bool {
	if o != nil && !isNil(o.Subnets) {
		return true
	}

	return false
}

// SetSubnets gets a reference to the given []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner and assigns it to the Subnets field.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) SetSubnets(v []UpdateNetworkApplianceVpnSiteToSiteVpnRequestSubnetsInner) {
	o.Subnets = v
}

func (o UpdateNetworkApplianceVpnSiteToSiteVpnRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.Hubs) {
		toSerialize["hubs"] = o.Hubs
	}
	if !isNil(o.Subnets) {
		toSerialize["subnets"] = o.Subnets
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct {
	value *UpdateNetworkApplianceVpnSiteToSiteVpnRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest) Get() *UpdateNetworkApplianceVpnSiteToSiteVpnRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest) Set(val *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest(val *UpdateNetworkApplianceVpnSiteToSiteVpnRequest) *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest {
	return &NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


