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

// checks if the CreateNetworkWirelessRfProfileRequestPerSsidSettings12 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequestPerSsidSettings12{}

// CreateNetworkWirelessRfProfileRequestPerSsidSettings12 Settings for SSID 12
type CreateNetworkWirelessRfProfileRequestPerSsidSettings12 struct {
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *CreateNetworkWirelessRfProfileRequestApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings12 instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings12() *CreateNetworkWirelessRfProfileRequestPerSsidSettings12 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings12{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings12WithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings12WithDefaults() *CreateNetworkWirelessRfProfileRequestPerSsidSettings12 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings12{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetMinBitrate() float32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetMinBitrateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetBands() CreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	if o == nil || IsNil(o.Bands) {
		var ret CreateNetworkWirelessRfProfileRequestApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetBandsOk() (*CreateNetworkWirelessRfProfileRequestApBandSettingsBands, bool) {
	if o == nil || IsNil(o.Bands) {
		return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) HasBands() bool {
	if o != nil && !IsNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given CreateNetworkWirelessRfProfileRequestApBandSettingsBands and assigns it to the Bands field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) SetBands(v CreateNetworkWirelessRfProfileRequestApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings12) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings12) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12 struct {
	value *CreateNetworkWirelessRfProfileRequestPerSsidSettings12
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12) Get() *CreateNetworkWirelessRfProfileRequestPerSsidSettings12 {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12) Set(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings12) *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12 {
	return &NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


