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

// checks if the GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin{}

// GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin WAN1/WAN2/Independent prefix.
type GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin struct {
	// Origin type
	Type *string `json:"type,omitempty"`
	// Uplink provided or independent
	Interfaces []string `json:"interfaces,omitempty"`
}

// NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin instantiates a new GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin {
	this := GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin{}
	return &this
}

// NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOriginWithDefaults instantiates a new GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOriginWithDefaults() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin {
	this := GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) SetType(v string) {
	o.Type = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) GetInterfaces() []string {
	if o == nil || IsNil(o.Interfaces) {
		var ret []string
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) GetInterfacesOk() ([]string, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []string and assigns it to the Interfaces field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) SetInterfaces(v []string) {
	o.Interfaces = v
}

func (o GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	return toSerialize, nil
}

type NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin struct {
	value *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin
	isSet bool
}

func (v NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) Get() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin {
	return v.value
}

func (v *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) Set(val *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin(val *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin {
	return &NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin{value: val, isSet: true}
}

func (v NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


