/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateDeviceSwitchRoutingStaticRouteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceSwitchRoutingStaticRouteRequest{}

// UpdateDeviceSwitchRoutingStaticRouteRequest struct for UpdateDeviceSwitchRoutingStaticRouteRequest
type UpdateDeviceSwitchRoutingStaticRouteRequest struct {
	// Name or description for layer 3 static route
	Name *string `json:"name,omitempty"`
	// The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet *string `json:"subnet,omitempty"`
	// IP address of the next hop device to which the device sends its traffic for the subnet
	NextHopIp *string `json:"nextHopIp,omitempty"`
	// Option to advertise static route via OSPF
	AdvertiseViaOspfEnabled *bool `json:"advertiseViaOspfEnabled,omitempty"`
	// Option to prefer static route over OSPF routes
	PreferOverOspfRoutesEnabled *bool `json:"preferOverOspfRoutesEnabled,omitempty"`
}

// NewUpdateDeviceSwitchRoutingStaticRouteRequest instantiates a new UpdateDeviceSwitchRoutingStaticRouteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSwitchRoutingStaticRouteRequest() *UpdateDeviceSwitchRoutingStaticRouteRequest {
	this := UpdateDeviceSwitchRoutingStaticRouteRequest{}
	return &this
}

// NewUpdateDeviceSwitchRoutingStaticRouteRequestWithDefaults instantiates a new UpdateDeviceSwitchRoutingStaticRouteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSwitchRoutingStaticRouteRequestWithDefaults() *UpdateDeviceSwitchRoutingStaticRouteRequest {
	this := UpdateDeviceSwitchRoutingStaticRouteRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetNextHopIp returns the NextHopIp field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetNextHopIp() string {
	if o == nil || IsNil(o.NextHopIp) {
		var ret string
		return ret
	}
	return *o.NextHopIp
}

// GetNextHopIpOk returns a tuple with the NextHopIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetNextHopIpOk() (*string, bool) {
	if o == nil || IsNil(o.NextHopIp) {
		return nil, false
	}
	return o.NextHopIp, true
}

// HasNextHopIp returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) HasNextHopIp() bool {
	if o != nil && !IsNil(o.NextHopIp) {
		return true
	}

	return false
}

// SetNextHopIp gets a reference to the given string and assigns it to the NextHopIp field.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) SetNextHopIp(v string) {
	o.NextHopIp = &v
}

// GetAdvertiseViaOspfEnabled returns the AdvertiseViaOspfEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetAdvertiseViaOspfEnabled() bool {
	if o == nil || IsNil(o.AdvertiseViaOspfEnabled) {
		var ret bool
		return ret
	}
	return *o.AdvertiseViaOspfEnabled
}

// GetAdvertiseViaOspfEnabledOk returns a tuple with the AdvertiseViaOspfEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetAdvertiseViaOspfEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdvertiseViaOspfEnabled) {
		return nil, false
	}
	return o.AdvertiseViaOspfEnabled, true
}

// HasAdvertiseViaOspfEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) HasAdvertiseViaOspfEnabled() bool {
	if o != nil && !IsNil(o.AdvertiseViaOspfEnabled) {
		return true
	}

	return false
}

// SetAdvertiseViaOspfEnabled gets a reference to the given bool and assigns it to the AdvertiseViaOspfEnabled field.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) SetAdvertiseViaOspfEnabled(v bool) {
	o.AdvertiseViaOspfEnabled = &v
}

// GetPreferOverOspfRoutesEnabled returns the PreferOverOspfRoutesEnabled field value if set, zero value otherwise.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetPreferOverOspfRoutesEnabled() bool {
	if o == nil || IsNil(o.PreferOverOspfRoutesEnabled) {
		var ret bool
		return ret
	}
	return *o.PreferOverOspfRoutesEnabled
}

// GetPreferOverOspfRoutesEnabledOk returns a tuple with the PreferOverOspfRoutesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) GetPreferOverOspfRoutesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferOverOspfRoutesEnabled) {
		return nil, false
	}
	return o.PreferOverOspfRoutesEnabled, true
}

// HasPreferOverOspfRoutesEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) HasPreferOverOspfRoutesEnabled() bool {
	if o != nil && !IsNil(o.PreferOverOspfRoutesEnabled) {
		return true
	}

	return false
}

// SetPreferOverOspfRoutesEnabled gets a reference to the given bool and assigns it to the PreferOverOspfRoutesEnabled field.
func (o *UpdateDeviceSwitchRoutingStaticRouteRequest) SetPreferOverOspfRoutesEnabled(v bool) {
	o.PreferOverOspfRoutesEnabled = &v
}

func (o UpdateDeviceSwitchRoutingStaticRouteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceSwitchRoutingStaticRouteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.NextHopIp) {
		toSerialize["nextHopIp"] = o.NextHopIp
	}
	if !IsNil(o.AdvertiseViaOspfEnabled) {
		toSerialize["advertiseViaOspfEnabled"] = o.AdvertiseViaOspfEnabled
	}
	if !IsNil(o.PreferOverOspfRoutesEnabled) {
		toSerialize["preferOverOspfRoutesEnabled"] = o.PreferOverOspfRoutesEnabled
	}
	return toSerialize, nil
}

type NullableUpdateDeviceSwitchRoutingStaticRouteRequest struct {
	value *UpdateDeviceSwitchRoutingStaticRouteRequest
	isSet bool
}

func (v NullableUpdateDeviceSwitchRoutingStaticRouteRequest) Get() *UpdateDeviceSwitchRoutingStaticRouteRequest {
	return v.value
}

func (v *NullableUpdateDeviceSwitchRoutingStaticRouteRequest) Set(val *UpdateDeviceSwitchRoutingStaticRouteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSwitchRoutingStaticRouteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSwitchRoutingStaticRouteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSwitchRoutingStaticRouteRequest(val *UpdateDeviceSwitchRoutingStaticRouteRequest) *NullableUpdateDeviceSwitchRoutingStaticRouteRequest {
	return &NullableUpdateDeviceSwitchRoutingStaticRouteRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceSwitchRoutingStaticRouteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSwitchRoutingStaticRouteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


