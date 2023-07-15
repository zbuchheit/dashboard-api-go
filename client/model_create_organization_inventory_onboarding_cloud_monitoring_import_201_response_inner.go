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

// checks if the CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner{}

// CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner struct for CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
type CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner struct {
	// Cloud monitor import status
	Status *string `json:"status,omitempty"`
	// Unique id associated with the import of the device
	ImportId *string `json:"importId,omitempty"`
	// Response method
	Message *string `json:"message,omitempty"`
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner() *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner{}
	return &this
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInnerWithDefaults instantiates a new CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInnerWithDefaults() *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner {
	this := CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) GetImportId() string {
	if o == nil || IsNil(o.ImportId) {
		var ret string
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) GetImportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImportId) {
		return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) HasImportId() bool {
	if o != nil && !IsNil(o.ImportId) {
		return true
	}

	return false
}

// SetImportId gets a reference to the given string and assigns it to the ImportId field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) SetImportId(v string) {
	o.ImportId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) SetMessage(v string) {
	o.Message = &v
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ImportId) {
		toSerialize["importId"] = o.ImportId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner struct {
	value *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
	isSet bool
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) Get() *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner {
	return v.value
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) Set(val *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner(val *CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner {
	return &NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


