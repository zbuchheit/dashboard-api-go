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

// checks if the GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule{}

// GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule Port schedule
type GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule struct {
	Monday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleMonday `json:"monday,omitempty"`
	Tuesday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleTuesday `json:"tuesday,omitempty"`
	Wednesday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday `json:"wednesday,omitempty"`
	Thursday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleThursday `json:"thursday,omitempty"`
	Friday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday `json:"friday,omitempty"`
	Saturday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSaturday `json:"saturday,omitempty"`
	Sunday *GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSunday `json:"sunday,omitempty"`
}

// NewGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule instantiates a new GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule() *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule {
	this := GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule{}
	return &this
}

// NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWithDefaults instantiates a new GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWithDefaults() *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule {
	this := GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule{}
	return &this
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetMonday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleMonday {
	if o == nil || IsNil(o.Monday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleMonday
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetMondayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleMonday, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleMonday and assigns it to the Monday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetMonday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleMonday) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetTuesday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleTuesday {
	if o == nil || IsNil(o.Tuesday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleTuesday
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetTuesdayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleTuesday, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleTuesday and assigns it to the Tuesday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetTuesday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleTuesday) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetWednesday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday {
	if o == nil || IsNil(o.Wednesday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetWednesdayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday and assigns it to the Wednesday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetWednesday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleWednesday) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetThursday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleThursday {
	if o == nil || IsNil(o.Thursday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleThursday
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetThursdayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleThursday, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleThursday and assigns it to the Thursday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetThursday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleThursday) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetFriday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday {
	if o == nil || IsNil(o.Friday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetFridayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday and assigns it to the Friday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetFriday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleFriday) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetSaturday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSaturday {
	if o == nil || IsNil(o.Saturday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSaturday
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetSaturdayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSaturday, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSaturday and assigns it to the Saturday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetSaturday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSaturday) {
	o.Saturday = &v
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetSunday() GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSunday {
	if o == nil || IsNil(o.Sunday) {
		var ret GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSunday
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) GetSundayOk() (*GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSunday, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSunday and assigns it to the Sunday field.
func (o *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) SetSunday(v GetNetworkSwitchPortSchedules200ResponseInnerPortScheduleSunday) {
	o.Sunday = &v
}

func (o GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Monday) {
		toSerialize["monday"] = o.Monday
	}
	if !IsNil(o.Tuesday) {
		toSerialize["tuesday"] = o.Tuesday
	}
	if !IsNil(o.Wednesday) {
		toSerialize["wednesday"] = o.Wednesday
	}
	if !IsNil(o.Thursday) {
		toSerialize["thursday"] = o.Thursday
	}
	if !IsNil(o.Friday) {
		toSerialize["friday"] = o.Friday
	}
	if !IsNil(o.Saturday) {
		toSerialize["saturday"] = o.Saturday
	}
	if !IsNil(o.Sunday) {
		toSerialize["sunday"] = o.Sunday
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule struct {
	value *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule
	isSet bool
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) Get() *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule {
	return v.value
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) Set(val *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule(val *GetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule {
	return &NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchPortSchedules200ResponseInnerPortSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


