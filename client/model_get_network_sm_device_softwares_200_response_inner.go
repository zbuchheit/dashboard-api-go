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

// checks if the GetNetworkSmDeviceSoftwares200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmDeviceSoftwares200ResponseInner{}

// GetNetworkSmDeviceSoftwares200ResponseInner struct for GetNetworkSmDeviceSoftwares200ResponseInner
type GetNetworkSmDeviceSoftwares200ResponseInner struct {
	// The Meraki managed application Id for this record on a particular device.
	AppId *string `json:"appId,omitempty"`
	// The size of the software bundle.
	BundleSize *int32 `json:"bundleSize,omitempty"`
	// When the Meraki record for the software was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The Meraki managed device Id.
	DeviceId *string `json:"deviceId,omitempty"`
	// The size of the data stored in the application.
	DynamicSize *int32 `json:"dynamicSize,omitempty"`
	// The Meraki software Id.
	Id *string `json:"id,omitempty"`
	// Software bundle identifier.
	Identifier *string `json:"identifier,omitempty"`
	// When the Software was installed on the device.
	InstalledAt *string `json:"installedAt,omitempty"`
	// A boolean indicating this software record should be installed on the associated device.
	ToInstall *bool `json:"toInstall,omitempty"`
	// A boolean indicating whether or not an iOS redemption code was used.
	IosRedemptionCode *bool `json:"iosRedemptionCode,omitempty"`
	// A boolean indicating whether or not the software is managed by Meraki.
	IsManaged *bool `json:"isManaged,omitempty"`
	// The itunes numerical identifier.
	ItunesId *string `json:"itunesId,omitempty"`
	// The license key associated with this software installation.
	LicenseKey *string `json:"licenseKey,omitempty"`
	// The name of the software.
	Name *string `json:"name,omitempty"`
	// The path on the device where the software record is located.
	Path *string `json:"path,omitempty"`
	// The redemption code used for this software.
	RedemptionCode *int32 `json:"redemptionCode,omitempty"`
	// Short version notation for the software.
	ShortVersion *string `json:"shortVersion,omitempty"`
	// The management status of the software.
	Status *string `json:"status,omitempty"`
	// A boolean indicating this software record should be uninstalled on the associated device.
	ToUninstall *bool `json:"toUninstall,omitempty"`
	// When the record was uninstalled from the device.
	UninstalledAt *string `json:"uninstalledAt,omitempty"`
	// When the record was last updated by Meraki.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The vendor of the software.
	Vendor *string `json:"vendor,omitempty"`
	// Full version notation for the software.
	Version *string `json:"version,omitempty"`
}

// NewGetNetworkSmDeviceSoftwares200ResponseInner instantiates a new GetNetworkSmDeviceSoftwares200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceSoftwares200ResponseInner() *GetNetworkSmDeviceSoftwares200ResponseInner {
	this := GetNetworkSmDeviceSoftwares200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceSoftwares200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceSoftwares200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceSoftwares200ResponseInnerWithDefaults() *GetNetworkSmDeviceSoftwares200ResponseInner {
	this := GetNetworkSmDeviceSoftwares200ResponseInner{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetAppId(v string) {
	o.AppId = &v
}

// GetBundleSize returns the BundleSize field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetBundleSize() int32 {
	if o == nil || IsNil(o.BundleSize) {
		var ret int32
		return ret
	}
	return *o.BundleSize
}

// GetBundleSizeOk returns a tuple with the BundleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetBundleSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BundleSize) {
		return nil, false
	}
	return o.BundleSize, true
}

// HasBundleSize returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasBundleSize() bool {
	if o != nil && !IsNil(o.BundleSize) {
		return true
	}

	return false
}

// SetBundleSize gets a reference to the given int32 and assigns it to the BundleSize field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetBundleSize(v int32) {
	o.BundleSize = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDynamicSize returns the DynamicSize field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDynamicSize() int32 {
	if o == nil || IsNil(o.DynamicSize) {
		var ret int32
		return ret
	}
	return *o.DynamicSize
}

// GetDynamicSizeOk returns a tuple with the DynamicSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetDynamicSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DynamicSize) {
		return nil, false
	}
	return o.DynamicSize, true
}

// HasDynamicSize returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasDynamicSize() bool {
	if o != nil && !IsNil(o.DynamicSize) {
		return true
	}

	return false
}

// SetDynamicSize gets a reference to the given int32 and assigns it to the DynamicSize field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetDynamicSize(v int32) {
	o.DynamicSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetInstalledAt returns the InstalledAt field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetInstalledAt() string {
	if o == nil || IsNil(o.InstalledAt) {
		var ret string
		return ret
	}
	return *o.InstalledAt
}

// GetInstalledAtOk returns a tuple with the InstalledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetInstalledAtOk() (*string, bool) {
	if o == nil || IsNil(o.InstalledAt) {
		return nil, false
	}
	return o.InstalledAt, true
}

// HasInstalledAt returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasInstalledAt() bool {
	if o != nil && !IsNil(o.InstalledAt) {
		return true
	}

	return false
}

// SetInstalledAt gets a reference to the given string and assigns it to the InstalledAt field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetInstalledAt(v string) {
	o.InstalledAt = &v
}

// GetToInstall returns the ToInstall field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToInstall() bool {
	if o == nil || IsNil(o.ToInstall) {
		var ret bool
		return ret
	}
	return *o.ToInstall
}

// GetToInstallOk returns a tuple with the ToInstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToInstallOk() (*bool, bool) {
	if o == nil || IsNil(o.ToInstall) {
		return nil, false
	}
	return o.ToInstall, true
}

// HasToInstall returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasToInstall() bool {
	if o != nil && !IsNil(o.ToInstall) {
		return true
	}

	return false
}

// SetToInstall gets a reference to the given bool and assigns it to the ToInstall field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetToInstall(v bool) {
	o.ToInstall = &v
}

// GetIosRedemptionCode returns the IosRedemptionCode field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIosRedemptionCode() bool {
	if o == nil || IsNil(o.IosRedemptionCode) {
		var ret bool
		return ret
	}
	return *o.IosRedemptionCode
}

// GetIosRedemptionCodeOk returns a tuple with the IosRedemptionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIosRedemptionCodeOk() (*bool, bool) {
	if o == nil || IsNil(o.IosRedemptionCode) {
		return nil, false
	}
	return o.IosRedemptionCode, true
}

// HasIosRedemptionCode returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasIosRedemptionCode() bool {
	if o != nil && !IsNil(o.IosRedemptionCode) {
		return true
	}

	return false
}

// SetIosRedemptionCode gets a reference to the given bool and assigns it to the IosRedemptionCode field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetIosRedemptionCode(v bool) {
	o.IosRedemptionCode = &v
}

// GetIsManaged returns the IsManaged field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIsManaged() bool {
	if o == nil || IsNil(o.IsManaged) {
		var ret bool
		return ret
	}
	return *o.IsManaged
}

// GetIsManagedOk returns a tuple with the IsManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetIsManagedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManaged) {
		return nil, false
	}
	return o.IsManaged, true
}

// HasIsManaged returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasIsManaged() bool {
	if o != nil && !IsNil(o.IsManaged) {
		return true
	}

	return false
}

// SetIsManaged gets a reference to the given bool and assigns it to the IsManaged field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetIsManaged(v bool) {
	o.IsManaged = &v
}

// GetItunesId returns the ItunesId field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetItunesId() string {
	if o == nil || IsNil(o.ItunesId) {
		var ret string
		return ret
	}
	return *o.ItunesId
}

// GetItunesIdOk returns a tuple with the ItunesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetItunesIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItunesId) {
		return nil, false
	}
	return o.ItunesId, true
}

// HasItunesId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasItunesId() bool {
	if o != nil && !IsNil(o.ItunesId) {
		return true
	}

	return false
}

// SetItunesId gets a reference to the given string and assigns it to the ItunesId field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetItunesId(v string) {
	o.ItunesId = &v
}

// GetLicenseKey returns the LicenseKey field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetLicenseKey() string {
	if o == nil || IsNil(o.LicenseKey) {
		var ret string
		return ret
	}
	return *o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetLicenseKeyOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseKey) {
		return nil, false
	}
	return o.LicenseKey, true
}

// HasLicenseKey returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasLicenseKey() bool {
	if o != nil && !IsNil(o.LicenseKey) {
		return true
	}

	return false
}

// SetLicenseKey gets a reference to the given string and assigns it to the LicenseKey field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetLicenseKey(v string) {
	o.LicenseKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetPath(v string) {
	o.Path = &v
}

// GetRedemptionCode returns the RedemptionCode field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetRedemptionCode() int32 {
	if o == nil || IsNil(o.RedemptionCode) {
		var ret int32
		return ret
	}
	return *o.RedemptionCode
}

// GetRedemptionCodeOk returns a tuple with the RedemptionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetRedemptionCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.RedemptionCode) {
		return nil, false
	}
	return o.RedemptionCode, true
}

// HasRedemptionCode returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasRedemptionCode() bool {
	if o != nil && !IsNil(o.RedemptionCode) {
		return true
	}

	return false
}

// SetRedemptionCode gets a reference to the given int32 and assigns it to the RedemptionCode field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetRedemptionCode(v int32) {
	o.RedemptionCode = &v
}

// GetShortVersion returns the ShortVersion field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetShortVersion() string {
	if o == nil || IsNil(o.ShortVersion) {
		var ret string
		return ret
	}
	return *o.ShortVersion
}

// GetShortVersionOk returns a tuple with the ShortVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetShortVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortVersion) {
		return nil, false
	}
	return o.ShortVersion, true
}

// HasShortVersion returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasShortVersion() bool {
	if o != nil && !IsNil(o.ShortVersion) {
		return true
	}

	return false
}

// SetShortVersion gets a reference to the given string and assigns it to the ShortVersion field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetShortVersion(v string) {
	o.ShortVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetToUninstall returns the ToUninstall field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToUninstall() bool {
	if o == nil || IsNil(o.ToUninstall) {
		var ret bool
		return ret
	}
	return *o.ToUninstall
}

// GetToUninstallOk returns a tuple with the ToUninstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetToUninstallOk() (*bool, bool) {
	if o == nil || IsNil(o.ToUninstall) {
		return nil, false
	}
	return o.ToUninstall, true
}

// HasToUninstall returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasToUninstall() bool {
	if o != nil && !IsNil(o.ToUninstall) {
		return true
	}

	return false
}

// SetToUninstall gets a reference to the given bool and assigns it to the ToUninstall field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetToUninstall(v bool) {
	o.ToUninstall = &v
}

// GetUninstalledAt returns the UninstalledAt field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUninstalledAt() string {
	if o == nil || IsNil(o.UninstalledAt) {
		var ret string
		return ret
	}
	return *o.UninstalledAt
}

// GetUninstalledAtOk returns a tuple with the UninstalledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUninstalledAtOk() (*string, bool) {
	if o == nil || IsNil(o.UninstalledAt) {
		return nil, false
	}
	return o.UninstalledAt, true
}

// HasUninstalledAt returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasUninstalledAt() bool {
	if o != nil && !IsNil(o.UninstalledAt) {
		return true
	}

	return false
}

// SetUninstalledAt gets a reference to the given string and assigns it to the UninstalledAt field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetUninstalledAt(v string) {
	o.UninstalledAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetNetworkSmDeviceSoftwares200ResponseInner) SetVersion(v string) {
	o.Version = &v
}

func (o GetNetworkSmDeviceSoftwares200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmDeviceSoftwares200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.BundleSize) {
		toSerialize["bundleSize"] = o.BundleSize
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.DynamicSize) {
		toSerialize["dynamicSize"] = o.DynamicSize
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.InstalledAt) {
		toSerialize["installedAt"] = o.InstalledAt
	}
	if !IsNil(o.ToInstall) {
		toSerialize["toInstall"] = o.ToInstall
	}
	if !IsNil(o.IosRedemptionCode) {
		toSerialize["iosRedemptionCode"] = o.IosRedemptionCode
	}
	if !IsNil(o.IsManaged) {
		toSerialize["isManaged"] = o.IsManaged
	}
	if !IsNil(o.ItunesId) {
		toSerialize["itunesId"] = o.ItunesId
	}
	if !IsNil(o.LicenseKey) {
		toSerialize["licenseKey"] = o.LicenseKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.RedemptionCode) {
		toSerialize["redemptionCode"] = o.RedemptionCode
	}
	if !IsNil(o.ShortVersion) {
		toSerialize["shortVersion"] = o.ShortVersion
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ToUninstall) {
		toSerialize["toUninstall"] = o.ToUninstall
	}
	if !IsNil(o.UninstalledAt) {
		toSerialize["uninstalledAt"] = o.UninstalledAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableGetNetworkSmDeviceSoftwares200ResponseInner struct {
	value *GetNetworkSmDeviceSoftwares200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceSoftwares200ResponseInner) Get() *GetNetworkSmDeviceSoftwares200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceSoftwares200ResponseInner) Set(val *GetNetworkSmDeviceSoftwares200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceSoftwares200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceSoftwares200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceSoftwares200ResponseInner(val *GetNetworkSmDeviceSoftwares200ResponseInner) *NullableGetNetworkSmDeviceSoftwares200ResponseInner {
	return &NullableGetNetworkSmDeviceSoftwares200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceSoftwares200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceSoftwares200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


