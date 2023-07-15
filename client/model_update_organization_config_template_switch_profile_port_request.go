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

// checks if the UpdateOrganizationConfigTemplateSwitchProfilePortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationConfigTemplateSwitchProfilePortRequest{}

// UpdateOrganizationConfigTemplateSwitchProfilePortRequest struct for UpdateOrganizationConfigTemplateSwitchProfilePortRequest
type UpdateOrganizationConfigTemplateSwitchProfilePortRequest struct {
	// The name of the switch template port.
	Name *string `json:"name,omitempty"`
	// The list of tags of the switch template port.
	Tags []string `json:"tags,omitempty"`
	// The status of the switch template port.
	Enabled *bool `json:"enabled,omitempty"`
	// The PoE status of the switch template port.
	PoeEnabled *bool `json:"poeEnabled,omitempty"`
	// The type of the switch template port ('trunk' or 'access').
	Type *string `json:"type,omitempty"`
	// The VLAN of the switch template port. A null value will clear the value set for trunk ports.
	Vlan *int32 `json:"vlan,omitempty"`
	// The voice VLAN of the switch template port. Only applicable to access ports.
	VoiceVlan *int32 `json:"voiceVlan,omitempty"`
	// The VLANs allowed on the switch template port. Only applicable to trunk ports.
	AllowedVlans *string `json:"allowedVlans,omitempty"`
	// The isolation status of the switch template port.
	IsolationEnabled *bool `json:"isolationEnabled,omitempty"`
	// The rapid spanning tree protocol status.
	RstpEnabled *bool `json:"rstpEnabled,omitempty"`
	// The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').
	StpGuard *string `json:"stpGuard,omitempty"`
	// The link speed for the switch template port.
	LinkNegotiation *string `json:"linkNegotiation,omitempty"`
	// The ID of the port schedule. A value of null will clear the port schedule.
	PortScheduleId *string `json:"portScheduleId,omitempty"`
	// The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.
	Udld *string `json:"udld,omitempty"`
	// The type of the access policy of the switch template port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.
	AccessPolicyType *string `json:"accessPolicyType,omitempty"`
	// The number of a custom access policy to configure on the switch template port. Only applicable when 'accessPolicyType' is 'Custom access policy'.
	AccessPolicyNumber *int32 `json:"accessPolicyNumber,omitempty"`
	// Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.
	MacAllowList []string `json:"macAllowList,omitempty"`
	// The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowList []string `json:"stickyMacAllowList,omitempty"`
	// The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.
	StickyMacAllowListLimit *int32 `json:"stickyMacAllowListLimit,omitempty"`
	// The storm control status of the switch template port.
	StormControlEnabled *bool `json:"stormControlEnabled,omitempty"`
	// For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.
	FlexibleStackingEnabled *bool `json:"flexibleStackingEnabled,omitempty"`
	// If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.
	DaiTrusted *bool `json:"daiTrusted,omitempty"`
	Profile *GetDeviceSwitchPorts200ResponseInnerProfile `json:"profile,omitempty"`
}

// NewUpdateOrganizationConfigTemplateSwitchProfilePortRequest instantiates a new UpdateOrganizationConfigTemplateSwitchProfilePortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationConfigTemplateSwitchProfilePortRequest() *UpdateOrganizationConfigTemplateSwitchProfilePortRequest {
	this := UpdateOrganizationConfigTemplateSwitchProfilePortRequest{}
	return &this
}

// NewUpdateOrganizationConfigTemplateSwitchProfilePortRequestWithDefaults instantiates a new UpdateOrganizationConfigTemplateSwitchProfilePortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationConfigTemplateSwitchProfilePortRequestWithDefaults() *UpdateOrganizationConfigTemplateSwitchProfilePortRequest {
	this := UpdateOrganizationConfigTemplateSwitchProfilePortRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetTags(v []string) {
	o.Tags = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPoeEnabled returns the PoeEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetPoeEnabled() bool {
	if o == nil || IsNil(o.PoeEnabled) {
		var ret bool
		return ret
	}
	return *o.PoeEnabled
}

// GetPoeEnabledOk returns a tuple with the PoeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetPoeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PoeEnabled) {
		return nil, false
	}
	return o.PoeEnabled, true
}

// HasPoeEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasPoeEnabled() bool {
	if o != nil && !IsNil(o.PoeEnabled) {
		return true
	}

	return false
}

// SetPoeEnabled gets a reference to the given bool and assigns it to the PoeEnabled field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetPoeEnabled(v bool) {
	o.PoeEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetType(v string) {
	o.Type = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetVlan() int32 {
	if o == nil || IsNil(o.Vlan) {
		var ret int32
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given int32 and assigns it to the Vlan field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetVlan(v int32) {
	o.Vlan = &v
}

// GetVoiceVlan returns the VoiceVlan field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetVoiceVlan() int32 {
	if o == nil || IsNil(o.VoiceVlan) {
		var ret int32
		return ret
	}
	return *o.VoiceVlan
}

// GetVoiceVlanOk returns a tuple with the VoiceVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetVoiceVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.VoiceVlan) {
		return nil, false
	}
	return o.VoiceVlan, true
}

// HasVoiceVlan returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasVoiceVlan() bool {
	if o != nil && !IsNil(o.VoiceVlan) {
		return true
	}

	return false
}

// SetVoiceVlan gets a reference to the given int32 and assigns it to the VoiceVlan field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetVoiceVlan(v int32) {
	o.VoiceVlan = &v
}

// GetAllowedVlans returns the AllowedVlans field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetAllowedVlans() string {
	if o == nil || IsNil(o.AllowedVlans) {
		var ret string
		return ret
	}
	return *o.AllowedVlans
}

// GetAllowedVlansOk returns a tuple with the AllowedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetAllowedVlansOk() (*string, bool) {
	if o == nil || IsNil(o.AllowedVlans) {
		return nil, false
	}
	return o.AllowedVlans, true
}

// HasAllowedVlans returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasAllowedVlans() bool {
	if o != nil && !IsNil(o.AllowedVlans) {
		return true
	}

	return false
}

// SetAllowedVlans gets a reference to the given string and assigns it to the AllowedVlans field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetAllowedVlans(v string) {
	o.AllowedVlans = &v
}

// GetIsolationEnabled returns the IsolationEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetIsolationEnabled() bool {
	if o == nil || IsNil(o.IsolationEnabled) {
		var ret bool
		return ret
	}
	return *o.IsolationEnabled
}

// GetIsolationEnabledOk returns a tuple with the IsolationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetIsolationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsolationEnabled) {
		return nil, false
	}
	return o.IsolationEnabled, true
}

// HasIsolationEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasIsolationEnabled() bool {
	if o != nil && !IsNil(o.IsolationEnabled) {
		return true
	}

	return false
}

// SetIsolationEnabled gets a reference to the given bool and assigns it to the IsolationEnabled field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetIsolationEnabled(v bool) {
	o.IsolationEnabled = &v
}

// GetRstpEnabled returns the RstpEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetRstpEnabled() bool {
	if o == nil || IsNil(o.RstpEnabled) {
		var ret bool
		return ret
	}
	return *o.RstpEnabled
}

// GetRstpEnabledOk returns a tuple with the RstpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetRstpEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RstpEnabled) {
		return nil, false
	}
	return o.RstpEnabled, true
}

// HasRstpEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasRstpEnabled() bool {
	if o != nil && !IsNil(o.RstpEnabled) {
		return true
	}

	return false
}

// SetRstpEnabled gets a reference to the given bool and assigns it to the RstpEnabled field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetRstpEnabled(v bool) {
	o.RstpEnabled = &v
}

// GetStpGuard returns the StpGuard field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStpGuard() string {
	if o == nil || IsNil(o.StpGuard) {
		var ret string
		return ret
	}
	return *o.StpGuard
}

// GetStpGuardOk returns a tuple with the StpGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStpGuardOk() (*string, bool) {
	if o == nil || IsNil(o.StpGuard) {
		return nil, false
	}
	return o.StpGuard, true
}

// HasStpGuard returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasStpGuard() bool {
	if o != nil && !IsNil(o.StpGuard) {
		return true
	}

	return false
}

// SetStpGuard gets a reference to the given string and assigns it to the StpGuard field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetStpGuard(v string) {
	o.StpGuard = &v
}

// GetLinkNegotiation returns the LinkNegotiation field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetLinkNegotiation() string {
	if o == nil || IsNil(o.LinkNegotiation) {
		var ret string
		return ret
	}
	return *o.LinkNegotiation
}

// GetLinkNegotiationOk returns a tuple with the LinkNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetLinkNegotiationOk() (*string, bool) {
	if o == nil || IsNil(o.LinkNegotiation) {
		return nil, false
	}
	return o.LinkNegotiation, true
}

// HasLinkNegotiation returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasLinkNegotiation() bool {
	if o != nil && !IsNil(o.LinkNegotiation) {
		return true
	}

	return false
}

// SetLinkNegotiation gets a reference to the given string and assigns it to the LinkNegotiation field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetLinkNegotiation(v string) {
	o.LinkNegotiation = &v
}

// GetPortScheduleId returns the PortScheduleId field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetPortScheduleId() string {
	if o == nil || IsNil(o.PortScheduleId) {
		var ret string
		return ret
	}
	return *o.PortScheduleId
}

// GetPortScheduleIdOk returns a tuple with the PortScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetPortScheduleIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortScheduleId) {
		return nil, false
	}
	return o.PortScheduleId, true
}

// HasPortScheduleId returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasPortScheduleId() bool {
	if o != nil && !IsNil(o.PortScheduleId) {
		return true
	}

	return false
}

// SetPortScheduleId gets a reference to the given string and assigns it to the PortScheduleId field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetPortScheduleId(v string) {
	o.PortScheduleId = &v
}

// GetUdld returns the Udld field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetUdld() string {
	if o == nil || IsNil(o.Udld) {
		var ret string
		return ret
	}
	return *o.Udld
}

// GetUdldOk returns a tuple with the Udld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetUdldOk() (*string, bool) {
	if o == nil || IsNil(o.Udld) {
		return nil, false
	}
	return o.Udld, true
}

// HasUdld returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasUdld() bool {
	if o != nil && !IsNil(o.Udld) {
		return true
	}

	return false
}

// SetUdld gets a reference to the given string and assigns it to the Udld field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetUdld(v string) {
	o.Udld = &v
}

// GetAccessPolicyType returns the AccessPolicyType field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetAccessPolicyType() string {
	if o == nil || IsNil(o.AccessPolicyType) {
		var ret string
		return ret
	}
	return *o.AccessPolicyType
}

// GetAccessPolicyTypeOk returns a tuple with the AccessPolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetAccessPolicyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPolicyType) {
		return nil, false
	}
	return o.AccessPolicyType, true
}

// HasAccessPolicyType returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasAccessPolicyType() bool {
	if o != nil && !IsNil(o.AccessPolicyType) {
		return true
	}

	return false
}

// SetAccessPolicyType gets a reference to the given string and assigns it to the AccessPolicyType field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetAccessPolicyType(v string) {
	o.AccessPolicyType = &v
}

// GetAccessPolicyNumber returns the AccessPolicyNumber field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetAccessPolicyNumber() int32 {
	if o == nil || IsNil(o.AccessPolicyNumber) {
		var ret int32
		return ret
	}
	return *o.AccessPolicyNumber
}

// GetAccessPolicyNumberOk returns a tuple with the AccessPolicyNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetAccessPolicyNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AccessPolicyNumber) {
		return nil, false
	}
	return o.AccessPolicyNumber, true
}

// HasAccessPolicyNumber returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasAccessPolicyNumber() bool {
	if o != nil && !IsNil(o.AccessPolicyNumber) {
		return true
	}

	return false
}

// SetAccessPolicyNumber gets a reference to the given int32 and assigns it to the AccessPolicyNumber field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetAccessPolicyNumber(v int32) {
	o.AccessPolicyNumber = &v
}

// GetMacAllowList returns the MacAllowList field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetMacAllowList() []string {
	if o == nil || IsNil(o.MacAllowList) {
		var ret []string
		return ret
	}
	return o.MacAllowList
}

// GetMacAllowListOk returns a tuple with the MacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetMacAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.MacAllowList) {
		return nil, false
	}
	return o.MacAllowList, true
}

// HasMacAllowList returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasMacAllowList() bool {
	if o != nil && !IsNil(o.MacAllowList) {
		return true
	}

	return false
}

// SetMacAllowList gets a reference to the given []string and assigns it to the MacAllowList field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetMacAllowList(v []string) {
	o.MacAllowList = v
}

// GetStickyMacAllowList returns the StickyMacAllowList field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStickyMacAllowList() []string {
	if o == nil || IsNil(o.StickyMacAllowList) {
		var ret []string
		return ret
	}
	return o.StickyMacAllowList
}

// GetStickyMacAllowListOk returns a tuple with the StickyMacAllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStickyMacAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.StickyMacAllowList) {
		return nil, false
	}
	return o.StickyMacAllowList, true
}

// HasStickyMacAllowList returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasStickyMacAllowList() bool {
	if o != nil && !IsNil(o.StickyMacAllowList) {
		return true
	}

	return false
}

// SetStickyMacAllowList gets a reference to the given []string and assigns it to the StickyMacAllowList field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetStickyMacAllowList(v []string) {
	o.StickyMacAllowList = v
}

// GetStickyMacAllowListLimit returns the StickyMacAllowListLimit field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStickyMacAllowListLimit() int32 {
	if o == nil || IsNil(o.StickyMacAllowListLimit) {
		var ret int32
		return ret
	}
	return *o.StickyMacAllowListLimit
}

// GetStickyMacAllowListLimitOk returns a tuple with the StickyMacAllowListLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStickyMacAllowListLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.StickyMacAllowListLimit) {
		return nil, false
	}
	return o.StickyMacAllowListLimit, true
}

// HasStickyMacAllowListLimit returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasStickyMacAllowListLimit() bool {
	if o != nil && !IsNil(o.StickyMacAllowListLimit) {
		return true
	}

	return false
}

// SetStickyMacAllowListLimit gets a reference to the given int32 and assigns it to the StickyMacAllowListLimit field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetStickyMacAllowListLimit(v int32) {
	o.StickyMacAllowListLimit = &v
}

// GetStormControlEnabled returns the StormControlEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStormControlEnabled() bool {
	if o == nil || IsNil(o.StormControlEnabled) {
		var ret bool
		return ret
	}
	return *o.StormControlEnabled
}

// GetStormControlEnabledOk returns a tuple with the StormControlEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetStormControlEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StormControlEnabled) {
		return nil, false
	}
	return o.StormControlEnabled, true
}

// HasStormControlEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasStormControlEnabled() bool {
	if o != nil && !IsNil(o.StormControlEnabled) {
		return true
	}

	return false
}

// SetStormControlEnabled gets a reference to the given bool and assigns it to the StormControlEnabled field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetStormControlEnabled(v bool) {
	o.StormControlEnabled = &v
}

// GetFlexibleStackingEnabled returns the FlexibleStackingEnabled field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetFlexibleStackingEnabled() bool {
	if o == nil || IsNil(o.FlexibleStackingEnabled) {
		var ret bool
		return ret
	}
	return *o.FlexibleStackingEnabled
}

// GetFlexibleStackingEnabledOk returns a tuple with the FlexibleStackingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetFlexibleStackingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FlexibleStackingEnabled) {
		return nil, false
	}
	return o.FlexibleStackingEnabled, true
}

// HasFlexibleStackingEnabled returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasFlexibleStackingEnabled() bool {
	if o != nil && !IsNil(o.FlexibleStackingEnabled) {
		return true
	}

	return false
}

// SetFlexibleStackingEnabled gets a reference to the given bool and assigns it to the FlexibleStackingEnabled field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetFlexibleStackingEnabled(v bool) {
	o.FlexibleStackingEnabled = &v
}

// GetDaiTrusted returns the DaiTrusted field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetDaiTrusted() bool {
	if o == nil || IsNil(o.DaiTrusted) {
		var ret bool
		return ret
	}
	return *o.DaiTrusted
}

// GetDaiTrustedOk returns a tuple with the DaiTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetDaiTrustedOk() (*bool, bool) {
	if o == nil || IsNil(o.DaiTrusted) {
		return nil, false
	}
	return o.DaiTrusted, true
}

// HasDaiTrusted returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasDaiTrusted() bool {
	if o != nil && !IsNil(o.DaiTrusted) {
		return true
	}

	return false
}

// SetDaiTrusted gets a reference to the given bool and assigns it to the DaiTrusted field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetDaiTrusted(v bool) {
	o.DaiTrusted = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetProfile() GetDeviceSwitchPorts200ResponseInnerProfile {
	if o == nil || IsNil(o.Profile) {
		var ret GetDeviceSwitchPorts200ResponseInnerProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) GetProfileOk() (*GetDeviceSwitchPorts200ResponseInnerProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given GetDeviceSwitchPorts200ResponseInnerProfile and assigns it to the Profile field.
func (o *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) SetProfile(v GetDeviceSwitchPorts200ResponseInnerProfile) {
	o.Profile = &v
}

func (o UpdateOrganizationConfigTemplateSwitchProfilePortRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationConfigTemplateSwitchProfilePortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.IsolationEnabled) {
		toSerialize["isolationEnabled"] = o.IsolationEnabled
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
	if !IsNil(o.PortScheduleId) {
		toSerialize["portScheduleId"] = o.PortScheduleId
	}
	if !IsNil(o.Udld) {
		toSerialize["udld"] = o.Udld
	}
	if !IsNil(o.AccessPolicyType) {
		toSerialize["accessPolicyType"] = o.AccessPolicyType
	}
	if !IsNil(o.AccessPolicyNumber) {
		toSerialize["accessPolicyNumber"] = o.AccessPolicyNumber
	}
	if !IsNil(o.MacAllowList) {
		toSerialize["macAllowList"] = o.MacAllowList
	}
	if !IsNil(o.StickyMacAllowList) {
		toSerialize["stickyMacAllowList"] = o.StickyMacAllowList
	}
	if !IsNil(o.StickyMacAllowListLimit) {
		toSerialize["stickyMacAllowListLimit"] = o.StickyMacAllowListLimit
	}
	if !IsNil(o.StormControlEnabled) {
		toSerialize["stormControlEnabled"] = o.StormControlEnabled
	}
	if !IsNil(o.FlexibleStackingEnabled) {
		toSerialize["flexibleStackingEnabled"] = o.FlexibleStackingEnabled
	}
	if !IsNil(o.DaiTrusted) {
		toSerialize["daiTrusted"] = o.DaiTrusted
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest struct {
	value *UpdateOrganizationConfigTemplateSwitchProfilePortRequest
	isSet bool
}

func (v NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest) Get() *UpdateOrganizationConfigTemplateSwitchProfilePortRequest {
	return v.value
}

func (v *NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest) Set(val *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest(val *UpdateOrganizationConfigTemplateSwitchProfilePortRequest) *NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest {
	return &NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationConfigTemplateSwitchProfilePortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


