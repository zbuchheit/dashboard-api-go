/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion{}

// GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion Details of the version the device will upgrade to if it exists
type GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion struct {
	// Firmware version identifier
	Id *string `json:"id,omitempty"`
	// Name of the firmware version
	Firmware *string `json:"firmware,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
	// Release type of the firmware version
	ReleaseType *string `json:"releaseType,omitempty"`
	// Release date of the firmware version
	ReleaseDate *time.Time `json:"releaseDate,omitempty"`
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion{}
	return &this
}

// NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersionWithDefaults instantiates a new GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersionWithDefaults() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion {
	this := GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetFirmware returns the Firmware field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetFirmware() string {
	if o == nil || IsNil(o.Firmware) {
		var ret string
		return ret
	}
	return *o.Firmware
}

// GetFirmwareOk returns a tuple with the Firmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetFirmwareOk() (*string, bool) {
	if o == nil || IsNil(o.Firmware) {
		return nil, false
	}
	return o.Firmware, true
}

// HasFirmware returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) HasFirmware() bool {
	if o != nil && !IsNil(o.Firmware) {
		return true
	}

	return false
}

// SetFirmware gets a reference to the given string and assigns it to the Firmware field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) SetFirmware(v string) {
	o.Firmware = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetReleaseType() string {
	if o == nil || IsNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetReleaseDate() time.Time {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given time.Time and assigns it to the ReleaseDate field.
func (o *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) SetReleaseDate(v time.Time) {
	o.ReleaseDate = &v
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Firmware) {
		toSerialize["firmware"] = o.Firmware
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion struct {
	value *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) Get() *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) Set(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion(val *GetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion {
	return &NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgrades200ResponseProductsWirelessNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


