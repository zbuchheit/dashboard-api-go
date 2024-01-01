/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 18 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.41.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the CreateNetworkFirmwareUpgradesRollbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesRollbackRequest{}

// CreateNetworkFirmwareUpgradesRollbackRequest struct for CreateNetworkFirmwareUpgradesRollbackRequest
type CreateNetworkFirmwareUpgradesRollbackRequest struct {
	// Product type to rollback (if the network is a combined network)
	Product *string `json:"product,omitempty"`
	// Scheduled time for the rollback
	Time *time.Time `json:"time,omitempty"`
	// Reasons for the rollback
	Reasons []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner `json:"reasons"`
	ToVersion *CreateNetworkFirmwareUpgradesRollbackRequestToVersion `json:"toVersion,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesRollbackRequest instantiates a new CreateNetworkFirmwareUpgradesRollbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesRollbackRequest(reasons []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) *CreateNetworkFirmwareUpgradesRollbackRequest {
	this := CreateNetworkFirmwareUpgradesRollbackRequest{}
	this.Reasons = reasons
	return &this
}

// NewCreateNetworkFirmwareUpgradesRollbackRequestWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesRollbackRequestWithDefaults() *CreateNetworkFirmwareUpgradesRollbackRequest {
	this := CreateNetworkFirmwareUpgradesRollbackRequest{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetProduct(v string) {
	o.Product = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetTime(v time.Time) {
	o.Time = &v
}

// GetReasons returns the Reasons field value
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetReasons() []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner {
	if o == nil {
		var ret []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner
		return ret
	}

	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetReasonsOk() ([]CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reasons, true
}

// SetReasons sets field value
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetReasons(v []CreateNetworkFirmwareUpgradesRollbackRequestReasonsInner) {
	o.Reasons = v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetToVersion() CreateNetworkFirmwareUpgradesRollbackRequestToVersion {
	if o == nil || IsNil(o.ToVersion) {
		var ret CreateNetworkFirmwareUpgradesRollbackRequestToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) GetToVersionOk() (*CreateNetworkFirmwareUpgradesRollbackRequestToVersion, bool) {
	if o == nil || IsNil(o.ToVersion) {
		return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) HasToVersion() bool {
	if o != nil && !IsNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given CreateNetworkFirmwareUpgradesRollbackRequestToVersion and assigns it to the ToVersion field.
func (o *CreateNetworkFirmwareUpgradesRollbackRequest) SetToVersion(v CreateNetworkFirmwareUpgradesRollbackRequestToVersion) {
	o.ToVersion = &v
}

func (o CreateNetworkFirmwareUpgradesRollbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesRollbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	toSerialize["reasons"] = o.Reasons
	if !IsNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesRollbackRequest struct {
	value *CreateNetworkFirmwareUpgradesRollbackRequest
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequest) Get() *CreateNetworkFirmwareUpgradesRollbackRequest {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequest) Set(val *CreateNetworkFirmwareUpgradesRollbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesRollbackRequest(val *CreateNetworkFirmwareUpgradesRollbackRequest) *NullableCreateNetworkFirmwareUpgradesRollbackRequest {
	return &NullableCreateNetworkFirmwareUpgradesRollbackRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesRollbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesRollbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


