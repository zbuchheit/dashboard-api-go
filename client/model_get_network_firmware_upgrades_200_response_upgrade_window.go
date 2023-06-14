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

// GetNetworkFirmwareUpgrades200ResponseUpgradeWindow Upgrade window for devices in network
type GetNetworkFirmwareUpgrades200ResponseUpgradeWindow struct {
	// Day of the week
	DayOfWeek *string `json:"dayOfWeek,omitempty"`
	// Hour of the day
	HourOfDay *string `json:"hourOfDay,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200ResponseUpgradeWindow instantiates a new GetNetworkFirmwareUpgrades200ResponseUpgradeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200ResponseUpgradeWindow() *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow {
	this := GetNetworkFirmwareUpgrades200ResponseUpgradeWindow{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseUpgradeWindowWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseUpgradeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseUpgradeWindowWithDefaults() *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow {
	this := GetNetworkFirmwareUpgrades200ResponseUpgradeWindow{}
	return &this
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) GetDayOfWeek() string {
	if o == nil || isNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) GetDayOfWeekOk() (*string, bool) {
	if o == nil || isNil(o.DayOfWeek) {
    return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) HasDayOfWeek() bool {
	if o != nil && !isNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetHourOfDay returns the HourOfDay field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) GetHourOfDay() string {
	if o == nil || isNil(o.HourOfDay) {
		var ret string
		return ret
	}
	return *o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) GetHourOfDayOk() (*string, bool) {
	if o == nil || isNil(o.HourOfDay) {
    return nil, false
	}
	return o.HourOfDay, true
}

// HasHourOfDay returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) HasHourOfDay() bool {
	if o != nil && !isNil(o.HourOfDay) {
		return true
	}

	return false
}

// SetHourOfDay gets a reference to the given string and assigns it to the HourOfDay field.
func (o *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) SetHourOfDay(v string) {
	o.HourOfDay = &v
}

func (o GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DayOfWeek) {
		toSerialize["dayOfWeek"] = o.DayOfWeek
	}
	if !isNil(o.HourOfDay) {
		toSerialize["hourOfDay"] = o.HourOfDay
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow struct {
	value *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow) Get() *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow) Set(val *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow(val *GetNetworkFirmwareUpgrades200ResponseUpgradeWindow) *NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow {
	return &NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseUpgradeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


