/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkCellularGatewayUplinkRequest struct for UpdateNetworkCellularGatewayUplinkRequest
type UpdateNetworkCellularGatewayUplinkRequest struct {
	BandwidthLimits *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular `json:"bandwidthLimits,omitempty"`
}

// NewUpdateNetworkCellularGatewayUplinkRequest instantiates a new UpdateNetworkCellularGatewayUplinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkCellularGatewayUplinkRequest() *UpdateNetworkCellularGatewayUplinkRequest {
	this := UpdateNetworkCellularGatewayUplinkRequest{}
	return &this
}

// NewUpdateNetworkCellularGatewayUplinkRequestWithDefaults instantiates a new UpdateNetworkCellularGatewayUplinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkCellularGatewayUplinkRequestWithDefaults() *UpdateNetworkCellularGatewayUplinkRequest {
	this := UpdateNetworkCellularGatewayUplinkRequest{}
	return &this
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *UpdateNetworkCellularGatewayUplinkRequest) GetBandwidthLimits() UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkCellularGatewayUplinkRequest) GetBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *UpdateNetworkCellularGatewayUplinkRequest) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular and assigns it to the BandwidthLimits field.
func (o *UpdateNetworkCellularGatewayUplinkRequest) SetBandwidthLimits(v UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsCellular) {
	o.BandwidthLimits = &v
}

func (o UpdateNetworkCellularGatewayUplinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkCellularGatewayUplinkRequest struct {
	value *UpdateNetworkCellularGatewayUplinkRequest
	isSet bool
}

func (v NullableUpdateNetworkCellularGatewayUplinkRequest) Get() *UpdateNetworkCellularGatewayUplinkRequest {
	return v.value
}

func (v *NullableUpdateNetworkCellularGatewayUplinkRequest) Set(val *UpdateNetworkCellularGatewayUplinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkCellularGatewayUplinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkCellularGatewayUplinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkCellularGatewayUplinkRequest(val *UpdateNetworkCellularGatewayUplinkRequest) *NullableUpdateNetworkCellularGatewayUplinkRequest {
	return &NullableUpdateNetworkCellularGatewayUplinkRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkCellularGatewayUplinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkCellularGatewayUplinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


