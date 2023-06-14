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

// CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X Quality and resolution for MV22X/MV72X camera models.
type CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.
	Resolution string `json:"resolution"`
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X(quality string, resolution string) *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72XWithDefaults instantiates a new CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72XWithDefaults() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X {
	this := CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X{}
	return &this
}

// GetQuality returns the Quality field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) SetResolution(v string) {
	o.Resolution = v
}

func (o CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X struct {
	value *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X
	isSet bool
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) Get() *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X {
	return v.value
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) Set(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X(val *CreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X {
	return &NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraQualityRetentionProfileRequestVideoSettingsMV22XMV72X) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


