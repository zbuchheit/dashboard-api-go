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

// checks if the GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner{}

// GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner struct for GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner
type GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner struct {
	// The identifier of the switch port.
	PortId *string `json:"portId,omitempty"`
	// The name of the switch port.
	Name *string `json:"name,omitempty"`
	// The list of tags of the switch port.
	Tags []string `json:"tags,omitempty"`
	// The status of the switch port.
	Enabled *bool `json:"enabled,omitempty"`
	// The PoE status of the switch port.
	PoeEnabled *bool `json:"poeEnabled,omitempty"`
	// The type of the switch port ('trunk' or 'access').
	Type *string `json:"type,omitempty"`
	// The VLAN of the switch port. A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the switch port. Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the switch port. Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The rapid spanning tree protocol status.
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	StpGuard *string `json:"stpGuard,omitempty"`
	// The link speed for the switch port.
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowList []string `json:"stickyMacAllowList,omitempty"`
	// The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int32 `json:"stickyMacAllowListLimit,omitempty"`
}

// NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner() *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner {
	this := GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner{}
	return &this
}

// NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInnerWithDefaults instantiates a new GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInnerWithDefaults() *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner {
	this := GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner{}
	return &this
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetPortId(v string) {
	o.PortId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetTags(v []string) {
	o.Tags = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPoeEnabled returns the PoeEnabled field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPoeEnabled() bool {
	if o == nil || IsNil(o.PoeEnabled) {
		var ret bool
		return ret
	}
	return *o.PoeEnabled
}

// GetPoeEnabledOk returns a tuple with the PoeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetPoeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PoeEnabled) {
		return nil, false
	}
	return o.PoeEnabled, true
}

// HasPoeEnabled returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasPoeEnabled() bool {
	if o != nil && !IsNil(o.PoeEnabled) {
		return true
	}

	return false
}

// SetPoeEnabled gets a reference to the given bool and assigns it to the PoeEnabled field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetPoeEnabled(v bool) {
	o.PoeEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVoiceVlan() int32 {
	if o == nil || IsNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.VoiceVlan) {
		return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasVoiceVlan() bool {
	if o != nil && !IsNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAllowedVlans() string {
	if o == nil || IsNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAllowedVlansOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedVlans) {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasAllowedVlans() bool {
	if o != nil && !IsNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetRstpEnabled() bool {
	if o == nil || IsNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RstpEnabled) {
		return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasRstpEnabled() bool {
	if o != nil && !IsNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpGuard returns the StpGuard field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStpGuard() string {
	if o == nil || IsNil(o.StpGuard) {
		var ret string
		return ret
	}
	return *o.StpGuard
}

// GetStpGuardOk returns a tuple with the StpGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStpGuardOk() (*string, bool) {
	if o == nil || IsNil(o.StpGuard) {
		return nil, false
	}
	return o.StpGuard, true
}

// HasStpGuard returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasStpGuard() bool {
	if o != nil && !IsNil(o.StpGuard) {
		return true
	}

	return false
}

// SetStpGuard gets a reference to the given string and assigns it to the StpGuard field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetStpGuard(v string) {
	o.StpGuard = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetLinkNegotiation() string {
	if o == nil || IsNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || IsNil(o.LinkNegotiation) {
		return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasLinkNegotiation() bool {
	if o != nil && !IsNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAccessPolicyType() string {
	if o == nil || IsNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicyType) {
		return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasAccessPolicyType() bool {
	if o != nil && !IsNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetStickyMacAllowList returns the StickyMacAllowList field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowList() []string {
	if o == nil || IsNil(o.StickyMacAllowList) {
		var ret []string
		return ret
	}
	return o.StickyMacAllowList
}

// GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.StickyMacAllowList) {
		return nil, false
	}
	return o.StickyMacAllowList, true
}

// HasStickyMacAllowList returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasStickyMacAllowList() bool {
	if o != nil && !IsNil(o.StickyMacAllowList) {
		return true
	}

	return false
}

// SetStickyMacAllowList gets a reference to the given []string and assigns it to the StickyMacAllowList field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetStickyMacAllowList(v []string) {
	o.StickyMacAllowList = v
}

// GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowListLimit() int32 {
	if o == nil || IsNil(o.StickyMacAllowListLimit) {
		var ret int32
		return ret
	}
	return *o.StickyMacAllowListLimit
}

// GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) GetStickyMacAllowListLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.StickyMacAllowListLimit) {
		return nil, false
	}
	return o.StickyMacAllowListLimit, true
}

// HasStickyMacAllowListLimit returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) HasStickyMacAllowListLimit() bool {
	if o != nil && !IsNil(o.StickyMacAllowListLimit) {
		return true
	}

	return false
}

// SetStickyMacAllowListLimit gets a reference to the given int32 and assigns it to the StickyMacAllowListLimit field.
func (o *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) SetStickyMacAllowListLimit(v int32) {
	o.StickyMacAllowListLimit = &v
}

func (o GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.PoeEnabled) {
		toSerialize["poeEnabled"] = o.PoeEnabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.VoiceVlan) {
		toSerialize["voiceVlan"] = o.VoiceVlan
	}
	if !IsNil(o.AllowedVlans) {
		toSerialize["allowedVlans"] = o.AllowedVlans
	}
	if !IsNil(o.RstpEnabled) {
		toSerialize["rstpEnabled"] = o.RstpEnabled
	}
	if !IsNil(o.StpGuard) {
		toSerialize["stpGuard"] = o.StpGuard
	}
	if !IsNil(o.LinkNegotiation) {
		toSerialize["linkNegotiation"] = o.LinkNegotiation
	}
	if !IsNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !IsNil(o.StickyMacAllowList) {
		toSerialize["stickyMacAllowList"] = o.StickyMacAllowList
	}
	if !IsNil(o.StickyMacAllowListLimit) {
		toSerialize["stickyMacAllowListLimit"] = o.StickyMacAllowListLimit
	}
	return toSerialize, nil
}

type NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner struct {
	value *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner
	isSet bool
}

func (v NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) Get() *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner {
	return v.value
}

func (v *NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) Set(val *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner(val *GetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) *NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner {
	return &NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSwitchPortsBySwitch200ResponseInnerPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


