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

// checks if the GetOrganizationWebhooksAlertTypes200ResponseExample type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWebhooksAlertTypes200ResponseExample{}

// GetOrganizationWebhooksAlertTypes200ResponseExample Example alert type
type GetOrganizationWebhooksAlertTypes200ResponseExample struct {
	// Version of the alert
	Version *string `json:"version,omitempty"`
	// Shared secret for the alert
	SharedSecret *string `json:"sharedSecret,omitempty"`
	// When the alert notification was sent, in ISO8601 format
	SentAt *time.Time `json:"sentAt,omitempty"`
	// ID for the alert instance
	AlertId *string `json:"alertId,omitempty"`
	// Severity level of the alert
	AlertLevel *string `json:"alertLevel,omitempty"`
	// When the alert occurred, in ISO8601 format
	OccurredAt *time.Time `json:"occurredAt,omitempty"`
	// Data for the specific alert. Contents depend on the type of the alert
	AlertData map[string]interface{} `json:"alertData,omitempty"`
	// ID of the organization associated with the alert
	OrganizationId *string `json:"organizationId,omitempty"`
	// Name of the organization associated with the alert
	OrganizationName *string `json:"organizationName,omitempty"`
	// URL of the organization associated with the alert
	OrganizationUrl *string `json:"organizationUrl,omitempty"`
	// Serial for the device associated with the alert
	DeviceSerial *string `json:"deviceSerial,omitempty"`
	// Mac address for the device associated with the alert
	DeviceMac *string `json:"deviceMac,omitempty"`
	// Name of the device associated with the alert
	DeviceName *string `json:"deviceName,omitempty"`
	// URL of the device associated with the alert
	DeviceUrl *string `json:"deviceUrl,omitempty"`
	// List of tags for the device associated with the alert
	DeviceTags []string `json:"deviceTags,omitempty"`
	// Model of the device associated with the alert
	DeviceModel *string `json:"deviceModel,omitempty"`
	// ID of the network associated with the alert
	NetworkId *string `json:"networkId,omitempty"`
	// Name of the network associated with the alert
	NetworkName *string `json:"networkName,omitempty"`
	// URL of the network associated with the alert
	NetworkUrl *string `json:"networkUrl,omitempty"`
	// Enrollment string of the network associated with the alert
	EnrollmentString *string `json:"enrollmentString,omitempty"`
	// Notes for the network associated with the alert
	Notes *string `json:"notes,omitempty"`
	// List of product types that are part of the network associated with the alert
	ProductTypes []string `json:"productTypes,omitempty"`
	// Encrypted ID of the network associated with the alert
	EncryptedId *string `json:"encryptedId,omitempty"`
}

// NewGetOrganizationWebhooksAlertTypes200ResponseExample instantiates a new GetOrganizationWebhooksAlertTypes200ResponseExample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWebhooksAlertTypes200ResponseExample() *GetOrganizationWebhooksAlertTypes200ResponseExample {
	this := GetOrganizationWebhooksAlertTypes200ResponseExample{}
	return &this
}

// NewGetOrganizationWebhooksAlertTypes200ResponseExampleWithDefaults instantiates a new GetOrganizationWebhooksAlertTypes200ResponseExample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWebhooksAlertTypes200ResponseExampleWithDefaults() *GetOrganizationWebhooksAlertTypes200ResponseExample {
	this := GetOrganizationWebhooksAlertTypes200ResponseExample{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetVersion(v string) {
	o.Version = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetSentAt returns the SentAt field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetSentAt() time.Time {
	if o == nil || IsNil(o.SentAt) {
		var ret time.Time
		return ret
	}
	return *o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetSentAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SentAt) {
		return nil, false
	}
	return o.SentAt, true
}

// HasSentAt returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasSentAt() bool {
	if o != nil && !IsNil(o.SentAt) {
		return true
	}

	return false
}

// SetSentAt gets a reference to the given time.Time and assigns it to the SentAt field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetSentAt(v time.Time) {
	o.SentAt = &v
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetAlertId() string {
	if o == nil || IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetAlertIdOk() (*string, bool) {
	if o == nil || IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasAlertId() bool {
	if o != nil && !IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetAlertId(v string) {
	o.AlertId = &v
}

// GetAlertLevel returns the AlertLevel field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetAlertLevel() string {
	if o == nil || IsNil(o.AlertLevel) {
		var ret string
		return ret
	}
	return *o.AlertLevel
}

// GetAlertLevelOk returns a tuple with the AlertLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetAlertLevelOk() (*string, bool) {
	if o == nil || IsNil(o.AlertLevel) {
		return nil, false
	}
	return o.AlertLevel, true
}

// HasAlertLevel returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasAlertLevel() bool {
	if o != nil && !IsNil(o.AlertLevel) {
		return true
	}

	return false
}

// SetAlertLevel gets a reference to the given string and assigns it to the AlertLevel field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetAlertLevel(v string) {
	o.AlertLevel = &v
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOccurredAt() time.Time {
	if o == nil || IsNil(o.OccurredAt) {
		var ret time.Time
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OccurredAt) {
		return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasOccurredAt() bool {
	if o != nil && !IsNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given time.Time and assigns it to the OccurredAt field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetOccurredAt(v time.Time) {
	o.OccurredAt = &v
}

// GetAlertData returns the AlertData field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetAlertData() map[string]interface{} {
	if o == nil || IsNil(o.AlertData) {
		var ret map[string]interface{}
		return ret
	}
	return o.AlertData
}

// GetAlertDataOk returns a tuple with the AlertData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetAlertDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AlertData) {
		return map[string]interface{}{}, false
	}
	return o.AlertData, true
}

// HasAlertData returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasAlertData() bool {
	if o != nil && !IsNil(o.AlertData) {
		return true
	}

	return false
}

// SetAlertData gets a reference to the given map[string]interface{} and assigns it to the AlertData field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetAlertData(v map[string]interface{}) {
	o.AlertData = v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetOrganizationUrl returns the OrganizationUrl field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOrganizationUrl() string {
	if o == nil || IsNil(o.OrganizationUrl) {
		var ret string
		return ret
	}
	return *o.OrganizationUrl
}

// GetOrganizationUrlOk returns a tuple with the OrganizationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetOrganizationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationUrl) {
		return nil, false
	}
	return o.OrganizationUrl, true
}

// HasOrganizationUrl returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasOrganizationUrl() bool {
	if o != nil && !IsNil(o.OrganizationUrl) {
		return true
	}

	return false
}

// SetOrganizationUrl gets a reference to the given string and assigns it to the OrganizationUrl field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetOrganizationUrl(v string) {
	o.OrganizationUrl = &v
}

// GetDeviceSerial returns the DeviceSerial field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceSerial() string {
	if o == nil || IsNil(o.DeviceSerial) {
		var ret string
		return ret
	}
	return *o.DeviceSerial
}

// GetDeviceSerialOk returns a tuple with the DeviceSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceSerialOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceSerial) {
		return nil, false
	}
	return o.DeviceSerial, true
}

// HasDeviceSerial returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasDeviceSerial() bool {
	if o != nil && !IsNil(o.DeviceSerial) {
		return true
	}

	return false
}

// SetDeviceSerial gets a reference to the given string and assigns it to the DeviceSerial field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetDeviceSerial(v string) {
	o.DeviceSerial = &v
}

// GetDeviceMac returns the DeviceMac field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceMac() string {
	if o == nil || IsNil(o.DeviceMac) {
		var ret string
		return ret
	}
	return *o.DeviceMac
}

// GetDeviceMacOk returns a tuple with the DeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceMacOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceMac) {
		return nil, false
	}
	return o.DeviceMac, true
}

// HasDeviceMac returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasDeviceMac() bool {
	if o != nil && !IsNil(o.DeviceMac) {
		return true
	}

	return false
}

// SetDeviceMac gets a reference to the given string and assigns it to the DeviceMac field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetDeviceMac(v string) {
	o.DeviceMac = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceUrl returns the DeviceUrl field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceUrl() string {
	if o == nil || IsNil(o.DeviceUrl) {
		var ret string
		return ret
	}
	return *o.DeviceUrl
}

// GetDeviceUrlOk returns a tuple with the DeviceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceUrl) {
		return nil, false
	}
	return o.DeviceUrl, true
}

// HasDeviceUrl returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasDeviceUrl() bool {
	if o != nil && !IsNil(o.DeviceUrl) {
		return true
	}

	return false
}

// SetDeviceUrl gets a reference to the given string and assigns it to the DeviceUrl field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetDeviceUrl(v string) {
	o.DeviceUrl = &v
}

// GetDeviceTags returns the DeviceTags field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceTags() []string {
	if o == nil || IsNil(o.DeviceTags) {
		var ret []string
		return ret
	}
	return o.DeviceTags
}

// GetDeviceTagsOk returns a tuple with the DeviceTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.DeviceTags) {
		return nil, false
	}
	return o.DeviceTags, true
}

// HasDeviceTags returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasDeviceTags() bool {
	if o != nil && !IsNil(o.DeviceTags) {
		return true
	}

	return false
}

// SetDeviceTags gets a reference to the given []string and assigns it to the DeviceTags field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetDeviceTags(v []string) {
	o.DeviceTags = v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceModel() string {
	if o == nil || IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetDeviceModelOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasDeviceModel() bool {
	if o != nil && !IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetNetworkName returns the NetworkName field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNetworkName() string {
	if o == nil || IsNil(o.NetworkName) {
		var ret string
		return ret
	}
	return *o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkName) {
		return nil, false
	}
	return o.NetworkName, true
}

// HasNetworkName returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasNetworkName() bool {
	if o != nil && !IsNil(o.NetworkName) {
		return true
	}

	return false
}

// SetNetworkName gets a reference to the given string and assigns it to the NetworkName field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetNetworkName(v string) {
	o.NetworkName = &v
}

// GetNetworkUrl returns the NetworkUrl field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNetworkUrl() string {
	if o == nil || IsNil(o.NetworkUrl) {
		var ret string
		return ret
	}
	return *o.NetworkUrl
}

// GetNetworkUrlOk returns a tuple with the NetworkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNetworkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkUrl) {
		return nil, false
	}
	return o.NetworkUrl, true
}

// HasNetworkUrl returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasNetworkUrl() bool {
	if o != nil && !IsNil(o.NetworkUrl) {
		return true
	}

	return false
}

// SetNetworkUrl gets a reference to the given string and assigns it to the NetworkUrl field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetNetworkUrl(v string) {
	o.NetworkUrl = &v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetEnrollmentString() string {
	if o == nil || IsNil(o.EnrollmentString) {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollmentString) {
		return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasEnrollmentString() bool {
	if o != nil && !IsNil(o.EnrollmentString) {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetNotes(v string) {
	o.Notes = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetProductTypes() []string {
	if o == nil || IsNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetProductTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypes) {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasProductTypes() bool {
	if o != nil && !IsNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetEncryptedId returns the EncryptedId field value if set, zero value otherwise.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetEncryptedId() string {
	if o == nil || IsNil(o.EncryptedId) {
		var ret string
		return ret
	}
	return *o.EncryptedId
}

// GetEncryptedIdOk returns a tuple with the EncryptedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) GetEncryptedIdOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedId) {
		return nil, false
	}
	return o.EncryptedId, true
}

// HasEncryptedId returns a boolean if a field has been set.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) HasEncryptedId() bool {
	if o != nil && !IsNil(o.EncryptedId) {
		return true
	}

	return false
}

// SetEncryptedId gets a reference to the given string and assigns it to the EncryptedId field.
func (o *GetOrganizationWebhooksAlertTypes200ResponseExample) SetEncryptedId(v string) {
	o.EncryptedId = &v
}

func (o GetOrganizationWebhooksAlertTypes200ResponseExample) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWebhooksAlertTypes200ResponseExample) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !IsNil(o.SentAt) {
		toSerialize["sentAt"] = o.SentAt
	}
	if !IsNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !IsNil(o.AlertLevel) {
		toSerialize["alertLevel"] = o.AlertLevel
	}
	if !IsNil(o.OccurredAt) {
		toSerialize["occurredAt"] = o.OccurredAt
	}
	if !IsNil(o.AlertData) {
		toSerialize["alertData"] = o.AlertData
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.OrganizationUrl) {
		toSerialize["organizationUrl"] = o.OrganizationUrl
	}
	if !IsNil(o.DeviceSerial) {
		toSerialize["deviceSerial"] = o.DeviceSerial
	}
	if !IsNil(o.DeviceMac) {
		toSerialize["deviceMac"] = o.DeviceMac
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.DeviceUrl) {
		toSerialize["deviceUrl"] = o.DeviceUrl
	}
	if !IsNil(o.DeviceTags) {
		toSerialize["deviceTags"] = o.DeviceTags
	}
	if !IsNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.NetworkName) {
		toSerialize["networkName"] = o.NetworkName
	}
	if !IsNil(o.NetworkUrl) {
		toSerialize["networkUrl"] = o.NetworkUrl
	}
	if !IsNil(o.EnrollmentString) {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !IsNil(o.EncryptedId) {
		toSerialize["encryptedId"] = o.EncryptedId
	}
	return toSerialize, nil
}

type NullableGetOrganizationWebhooksAlertTypes200ResponseExample struct {
	value *GetOrganizationWebhooksAlertTypes200ResponseExample
	isSet bool
}

func (v NullableGetOrganizationWebhooksAlertTypes200ResponseExample) Get() *GetOrganizationWebhooksAlertTypes200ResponseExample {
	return v.value
}

func (v *NullableGetOrganizationWebhooksAlertTypes200ResponseExample) Set(val *GetOrganizationWebhooksAlertTypes200ResponseExample) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWebhooksAlertTypes200ResponseExample) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWebhooksAlertTypes200ResponseExample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWebhooksAlertTypes200ResponseExample(val *GetOrganizationWebhooksAlertTypes200ResponseExample) *NullableGetOrganizationWebhooksAlertTypes200ResponseExample {
	return &NullableGetOrganizationWebhooksAlertTypes200ResponseExample{value: val, isSet: true}
}

func (v NullableGetOrganizationWebhooksAlertTypes200ResponseExample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWebhooksAlertTypes200ResponseExample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


