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

// checks if the GetNetworkWirelessSsids200ResponseInnerRadiusServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsids200ResponseInnerRadiusServersInner{}

// GetNetworkWirelessSsids200ResponseInnerRadiusServersInner struct for GetNetworkWirelessSsids200ResponseInnerRadiusServersInner
type GetNetworkWirelessSsids200ResponseInnerRadiusServersInner struct {
	// IP address (or FQDN) of your RADIUS server
	Host *string `json:"host,omitempty"`
	// UDP port the RADIUS server listens on for Access-requests
	Port *int32 `json:"port,omitempty"`
	// The ID of the Openroaming Certificate attached to radius server
	OpenRoamingCertificateId *int32 `json:"openRoamingCertificateId,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewGetNetworkWirelessSsids200ResponseInnerRadiusServersInner instantiates a new GetNetworkWirelessSsids200ResponseInnerRadiusServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsids200ResponseInnerRadiusServersInner() *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner {
	this := GetNetworkWirelessSsids200ResponseInnerRadiusServersInner{}
	return &this
}

// NewGetNetworkWirelessSsids200ResponseInnerRadiusServersInnerWithDefaults instantiates a new GetNetworkWirelessSsids200ResponseInnerRadiusServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsids200ResponseInnerRadiusServersInnerWithDefaults() *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner {
	this := GetNetworkWirelessSsids200ResponseInnerRadiusServersInner{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetOpenRoamingCertificateId() int32 {
	if o == nil || IsNil(o.OpenRoamingCertificateId) {
		var ret int32
		return ret
	}
	return *o.OpenRoamingCertificateId
}

// GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetOpenRoamingCertificateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenRoamingCertificateId) {
		return nil, false
	}
	return o.OpenRoamingCertificateId, true
}

// HasOpenRoamingCertificateId returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) HasOpenRoamingCertificateId() bool {
	if o != nil && !IsNil(o.OpenRoamingCertificateId) {
		return true
	}

	return false
}

// SetOpenRoamingCertificateId gets a reference to the given int32 and assigns it to the OpenRoamingCertificateId field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) SetOpenRoamingCertificateId(v int32) {
	o.OpenRoamingCertificateId = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.OpenRoamingCertificateId) {
		toSerialize["openRoamingCertificateId"] = o.OpenRoamingCertificateId
	}
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner struct {
	value *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner
	isSet bool
}

func (v NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner) Get() *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner {
	return v.value
}

func (v *NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner) Set(val *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner(val *GetNetworkWirelessSsids200ResponseInnerRadiusServersInner) *NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner {
	return &NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsids200ResponseInnerRadiusServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


