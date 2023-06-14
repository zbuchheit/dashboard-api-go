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

// GetNetworkSmDeviceDeviceProfiles200ResponseInner struct for GetNetworkSmDeviceDeviceProfiles200ResponseInner
type GetNetworkSmDeviceDeviceProfiles200ResponseInner struct {
	// The Meraki managed device Id.
	DeviceId *string `json:"deviceId,omitempty"`
	// The numerical Meraki Id of the profile.
	Id *string `json:"id,omitempty"`
	// A boolean indicating if the profile is encrypted.
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// Whether or not the profile is managed by Meraki.
	IsManaged *bool `json:"isManaged,omitempty"`
	// A string containing a JSON object with the profile data.
	ProfileData *string `json:"profileData,omitempty"`
	// The identifier of the profile.
	ProfileIdentifier *string `json:"profileIdentifier,omitempty"`
	// The name of the profile.
	Name *string `json:"name,omitempty"`
	// The verison of the profile.
	Version *string `json:"version,omitempty"`
}

// NewGetNetworkSmDeviceDeviceProfiles200ResponseInner instantiates a new GetNetworkSmDeviceDeviceProfiles200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceDeviceProfiles200ResponseInner() *GetNetworkSmDeviceDeviceProfiles200ResponseInner {
	this := GetNetworkSmDeviceDeviceProfiles200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceDeviceProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceDeviceProfiles200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceDeviceProfiles200ResponseInnerWithDefaults() *GetNetworkSmDeviceDeviceProfiles200ResponseInner {
	this := GetNetworkSmDeviceDeviceProfiles200ResponseInner{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
    return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetIsEncrypted returns the IsEncrypted field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsEncrypted() bool {
	if o == nil || isNil(o.IsEncrypted) {
		var ret bool
		return ret
	}
	return *o.IsEncrypted
}

// GetIsEncryptedOk returns a tuple with the IsEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsEncryptedOk() (*bool, bool) {
	if o == nil || isNil(o.IsEncrypted) {
    return nil, false
	}
	return o.IsEncrypted, true
}

// HasIsEncrypted returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasIsEncrypted() bool {
	if o != nil && !isNil(o.IsEncrypted) {
		return true
	}

	return false
}

// SetIsEncrypted gets a reference to the given bool and assigns it to the IsEncrypted field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetIsEncrypted(v bool) {
	o.IsEncrypted = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsManaged() bool {
	if o == nil || isNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetIsManagedOk() (*bool, bool) {
	if o == nil || isNil(o.IsManaged) {
    return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasIsManaged() bool {
	if o != nil && !isNil(o.IsManaged) {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetProfileData returns the ProfileData field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileData() string {
	if o == nil || isNil(o.ProfileData) {
		var ret string
		return ret
	}
	return *o.ProfileData
}

// GetProfileDataOk returns a tuple with the ProfileData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileDataOk() (*string, bool) {
	if o == nil || isNil(o.ProfileData) {
    return nil, false
	}
	return o.ProfileData, true
}

// HasProfileData returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasProfileData() bool {
	if o != nil && !isNil(o.ProfileData) {
		return true
	}

	return false
}

// SetProfileData gets a reference to the given string and assigns it to the ProfileData field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetProfileData(v string) {
	o.ProfileData = &v
}

// GetProfileIdentifier returns the ProfileIdentifier field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileIdentifier() string {
	if o == nil || isNil(o.ProfileIdentifier) {
		var ret string
		return ret
	}
	return *o.ProfileIdentifier
}

// GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetProfileIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.ProfileIdentifier) {
    return nil, false
	}
	return o.ProfileIdentifier, true
}

// HasProfileIdentifier returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasProfileIdentifier() bool {
	if o != nil && !isNil(o.ProfileIdentifier) {
		return true
	}

	return false
}

// SetProfileIdentifier gets a reference to the given string and assigns it to the ProfileIdentifier field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetProfileIdentifier(v string) {
	o.ProfileIdentifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetNetworkSmDeviceDeviceProfiles200ResponseInner) SetVersion(v string) {
	o.Version = &v
}

func (o GetNetworkSmDeviceDeviceProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IsEncrypted) {
		toSerialize["isEncrypted"] = o.IsEncrypted
	}
	if !isNil(o.IsManaged) {
		toSerialize["isManaged"] = o.IsManaged
	}
	if !isNil(o.ProfileData) {
		toSerialize["profileData"] = o.ProfileData
	}
	if !isNil(o.ProfileIdentifier) {
		toSerialize["profileIdentifier"] = o.ProfileIdentifier
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner struct {
	value *GetNetworkSmDeviceDeviceProfiles200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner) Get() *GetNetworkSmDeviceDeviceProfiles200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner) Set(val *GetNetworkSmDeviceDeviceProfiles200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceDeviceProfiles200ResponseInner(val *GetNetworkSmDeviceDeviceProfiles200ResponseInner) *NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner {
	return &NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceDeviceProfiles200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


