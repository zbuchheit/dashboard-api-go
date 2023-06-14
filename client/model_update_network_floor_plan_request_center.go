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

// UpdateNetworkFloorPlanRequestCenter The longitude and latitude of the center of your floor plan. If you want to change the geolocation data of your floor plan, either the 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.
type UpdateNetworkFloorPlanRequestCenter struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewUpdateNetworkFloorPlanRequestCenter instantiates a new UpdateNetworkFloorPlanRequestCenter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFloorPlanRequestCenter() *UpdateNetworkFloorPlanRequestCenter {
	this := UpdateNetworkFloorPlanRequestCenter{}
	return &this
}

// NewUpdateNetworkFloorPlanRequestCenterWithDefaults instantiates a new UpdateNetworkFloorPlanRequestCenter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFloorPlanRequestCenterWithDefaults() *UpdateNetworkFloorPlanRequestCenter {
	this := UpdateNetworkFloorPlanRequestCenter{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequestCenter) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequestCenter) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequestCenter) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *UpdateNetworkFloorPlanRequestCenter) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *UpdateNetworkFloorPlanRequestCenter) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFloorPlanRequestCenter) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *UpdateNetworkFloorPlanRequestCenter) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *UpdateNetworkFloorPlanRequestCenter) SetLng(v float32) {
	o.Lng = &v
}

func (o UpdateNetworkFloorPlanRequestCenter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkFloorPlanRequestCenter struct {
	value *UpdateNetworkFloorPlanRequestCenter
	isSet bool
}

func (v NullableUpdateNetworkFloorPlanRequestCenter) Get() *UpdateNetworkFloorPlanRequestCenter {
	return v.value
}

func (v *NullableUpdateNetworkFloorPlanRequestCenter) Set(val *UpdateNetworkFloorPlanRequestCenter) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFloorPlanRequestCenter) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFloorPlanRequestCenter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFloorPlanRequestCenter(val *UpdateNetworkFloorPlanRequestCenter) *NullableUpdateNetworkFloorPlanRequestCenter {
	return &NullableUpdateNetworkFloorPlanRequestCenter{value: val, isSet: true}
}

func (v NullableUpdateNetworkFloorPlanRequestCenter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFloorPlanRequestCenter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


