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

// checks if the GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner{}

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner struct for GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner struct {
	// The type of sensor metric that will be monitored for changes. Available metrics are door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water.
	Metric string `json:"metric"`
	Threshold GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold `json:"threshold"`
	// If 'above', an alert will be sent when a sensor reads above the threshold. If 'below', an alert will be sent when a sensor reads below the threshold. Only applicable for temperature and humidity thresholds.
	Direction *string `json:"direction,omitempty"`
	// Length of time in seconds that the triggering state must persist before an alert is sent. Available options are 0 seconds, 1 minute, 2 minutes, 3 minutes, 4 minutes, 5 minutes, 10 minutes, 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, and 8 hours. Default is 0.
	Duration *int32 `json:"duration,omitempty"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner(metric string, threshold GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner{}
	this.Metric = metric
	this.Threshold = threshold
	var duration int32 = 0
	this.Duration = &duration
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner{}
	var duration int32 = 0
	this.Duration = &duration
	return &this
}

// GetMetric returns the Metric field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetMetric(v string) {
	o.Metric = v
}

// GetThreshold returns the Threshold field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetThreshold() GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold {
	if o == nil {
		var ret GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetThresholdOk() (*GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetThreshold(v GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold) {
	o.Threshold = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetDirection(v string) {
	o.Direction = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) SetDuration(v int32) {
	o.Duration = &v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metric"] = o.Metric
	toSerialize["threshold"] = o.Threshold
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	return toSerialize, nil
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


