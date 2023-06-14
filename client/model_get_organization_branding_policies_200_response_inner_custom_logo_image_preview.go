/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview Preview of the image
type GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview struct {
	// Url of the preview image
	Url *string `json:"url,omitempty"`
	// Timestamp of the preview image
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview instantiates a new GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview {
	this := GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview{}
	return &this
}

// NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreviewWithDefaults instantiates a new GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreviewWithDefaults() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview {
	this := GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) SetUrl(v string) {
	o.Url = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview struct {
	value *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview
	isSet bool
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) Get() *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview {
	return v.value
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) Set(val *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview(val *GetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview {
	return &NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview{value: val, isSet: true}
}

func (v NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationBrandingPolicies200ResponseInnerCustomLogoImagePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


