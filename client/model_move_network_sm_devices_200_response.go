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

// checks if the MoveNetworkSmDevices200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveNetworkSmDevices200Response{}

// MoveNetworkSmDevices200Response struct for MoveNetworkSmDevices200Response
type MoveNetworkSmDevices200Response struct {
	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids,omitempty"`
	// The network to which the devices was moved.
	NewNetwork *string `json:"newNetwork,omitempty"`
}

// NewMoveNetworkSmDevices200Response instantiates a new MoveNetworkSmDevices200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveNetworkSmDevices200Response() *MoveNetworkSmDevices200Response {
	this := MoveNetworkSmDevices200Response{}
	return &this
}

// NewMoveNetworkSmDevices200ResponseWithDefaults instantiates a new MoveNetworkSmDevices200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveNetworkSmDevices200ResponseWithDefaults() *MoveNetworkSmDevices200Response {
	this := MoveNetworkSmDevices200Response{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *MoveNetworkSmDevices200Response) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevices200Response) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *MoveNetworkSmDevices200Response) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *MoveNetworkSmDevices200Response) SetIds(v []string) {
	o.Ids = v
}

// GetNewNetwork returns the NewNetwork field value if set, zero value otherwise.
func (o *MoveNetworkSmDevices200Response) GetNewNetwork() string {
	if o == nil || IsNil(o.NewNetwork) {
		var ret string
		return ret
	}
	return *o.NewNetwork
}

// GetNewNetworkOk returns a tuple with the NewNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveNetworkSmDevices200Response) GetNewNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.NewNetwork) {
		return nil, false
	}
	return o.NewNetwork, true
}

// HasNewNetwork returns a boolean if a field has been set.
func (o *MoveNetworkSmDevices200Response) HasNewNetwork() bool {
	if o != nil && !IsNil(o.NewNetwork) {
		return true
	}

	return false
}

// SetNewNetwork gets a reference to the given string and assigns it to the NewNetwork field.
func (o *MoveNetworkSmDevices200Response) SetNewNetwork(v string) {
	o.NewNetwork = &v
}

func (o MoveNetworkSmDevices200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveNetworkSmDevices200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.NewNetwork) {
		toSerialize["newNetwork"] = o.NewNetwork
	}
	return toSerialize, nil
}

type NullableMoveNetworkSmDevices200Response struct {
	value *MoveNetworkSmDevices200Response
	isSet bool
}

func (v NullableMoveNetworkSmDevices200Response) Get() *MoveNetworkSmDevices200Response {
	return v.value
}

func (v *NullableMoveNetworkSmDevices200Response) Set(val *MoveNetworkSmDevices200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveNetworkSmDevices200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveNetworkSmDevices200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveNetworkSmDevices200Response(val *MoveNetworkSmDevices200Response) *NullableMoveNetworkSmDevices200Response {
	return &NullableMoveNetworkSmDevices200Response{value: val, isSet: true}
}

func (v NullableMoveNetworkSmDevices200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveNetworkSmDevices200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


