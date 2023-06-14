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

// GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner struct for GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner
type GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner struct {
	Network *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork `json:"network,omitempty"`
	// Name of the appliance
	Name *string `json:"name,omitempty"`
	// Mac address of the appliance
	Mac *string `json:"mac,omitempty"`
	// Serial number of the appliance
	Serial *string `json:"serial,omitempty"`
	// Model of the appliance
	Model *string `json:"model,omitempty"`
	Utilization *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization `json:"utilization,omitempty"`
}

// NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner instantiates a new GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner() *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner {
	this := GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerWithDefaults() *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner {
	this := GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetNetwork() GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork {
	if o == nil || isNil(o.Network) {
		var ret GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetNetworkOk() (*GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) SetNetwork(v GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerNetwork) {
	o.Network = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetUtilization() GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization {
	if o == nil || isNil(o.Utilization) {
		var ret GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) GetUtilizationOk() (*GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization, bool) {
	if o == nil || isNil(o.Utilization) {
    return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) HasUtilization() bool {
	if o != nil && !isNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization and assigns it to the Utilization field.
func (o *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) SetUtilization(v GetOrganizationSummaryTopAppliancesByUtilization200ResponseInnerUtilization) {
	o.Utilization = &v
}

func (o GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner struct {
	value *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) Get() *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) Set(val *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner(val *GetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner {
	return &NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopAppliancesByUtilization200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


