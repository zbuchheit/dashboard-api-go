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

// checks if the UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner{}

// UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner struct for UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner
type UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner struct {
	// The network ID of the hub.
	HubId string `json:"hubId"`
	// Only valid in 'spoke' mode. Indicates whether default route traffic should be sent to this hub.
	UseDefaultRoute *bool `json:"useDefaultRoute,omitempty"`
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner(hubId string) *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner {
	this := UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner{}
	this.HubId = hubId
	return &this
}

// NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInnerWithDefaults instantiates a new UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInnerWithDefaults() *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner {
	this := UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner{}
	return &this
}

// GetHubId returns the HubId field value
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetHubId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetHubIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubId, true
}

// SetHubId sets field value
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) SetHubId(v string) {
	o.HubId = v
}

// GetUseDefaultRoute returns the UseDefaultRoute field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetUseDefaultRoute() bool {
	if o == nil || IsNil(o.UseDefaultRoute) {
		var ret bool
		return ret
	}
	return *o.UseDefaultRoute
}

// GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) GetUseDefaultRouteOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaultRoute) {
		return nil, false
	}
	return o.UseDefaultRoute, true
}

// HasUseDefaultRoute returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) HasUseDefaultRoute() bool {
	if o != nil && !IsNil(o.UseDefaultRoute) {
		return true
	}

	return false
}

// SetUseDefaultRoute gets a reference to the given bool and assigns it to the UseDefaultRoute field.
func (o *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) SetUseDefaultRoute(v bool) {
	o.UseDefaultRoute = &v
}

func (o UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hubId"] = o.HubId
	if !IsNil(o.UseDefaultRoute) {
		toSerialize["useDefaultRoute"] = o.UseDefaultRoute
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner struct {
	value *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) Get() *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) Set(val *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner(val *UpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner {
	return &NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnSiteToSiteVpnRequestHubsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


