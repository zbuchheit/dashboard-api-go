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

// GetOrganizationEarlyAccessFeatures200ResponseInner struct for GetOrganizationEarlyAccessFeatures200ResponseInner
type GetOrganizationEarlyAccessFeatures200ResponseInner struct {
	// Short name of the early access feature
	ShortName *string `json:"shortName,omitempty"`
	// Name of the early access feature
	Name *string `json:"name,omitempty"`
	Descriptions *GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions `json:"descriptions,omitempty"`
	// Topic of the early access feature
	Topic *string `json:"topic,omitempty"`
	// If this early access feature can only be opted in for the entire organization
	IsOrgScopedOnly *bool `json:"isOrgScopedOnly,omitempty"`
	// Link to the documentation of this early access feature
	DocumentationLink *string `json:"documentationLink,omitempty"`
	// Link to get support for this early access feature
	SupportLink *string `json:"supportLink,omitempty"`
}

// NewGetOrganizationEarlyAccessFeatures200ResponseInner instantiates a new GetOrganizationEarlyAccessFeatures200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationEarlyAccessFeatures200ResponseInner() *GetOrganizationEarlyAccessFeatures200ResponseInner {
	this := GetOrganizationEarlyAccessFeatures200ResponseInner{}
	return &this
}

// NewGetOrganizationEarlyAccessFeatures200ResponseInnerWithDefaults instantiates a new GetOrganizationEarlyAccessFeatures200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationEarlyAccessFeatures200ResponseInnerWithDefaults() *GetOrganizationEarlyAccessFeatures200ResponseInner {
	this := GetOrganizationEarlyAccessFeatures200ResponseInner{}
	return &this
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetShortName(v string) {
	o.ShortName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDescriptions() GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions {
	if o == nil || isNil(o.Descriptions) {
		var ret GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDescriptionsOk() (*GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions, bool) {
	if o == nil || isNil(o.Descriptions) {
    return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasDescriptions() bool {
	if o != nil && !isNil(o.Descriptions) {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions and assigns it to the Descriptions field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetDescriptions(v GetOrganizationEarlyAccessFeatures200ResponseInnerDescriptions) {
	o.Descriptions = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetTopic() string {
	if o == nil || isNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetTopicOk() (*string, bool) {
	if o == nil || isNil(o.Topic) {
    return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasTopic() bool {
	if o != nil && !isNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetTopic(v string) {
	o.Topic = &v
}

// GetIsOrgScopedOnly returns the IsOrgScopedOnly field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetIsOrgScopedOnly() bool {
	if o == nil || isNil(o.IsOrgScopedOnly) {
		var ret bool
		return ret
	}
	return *o.IsOrgScopedOnly
}

// GetIsOrgScopedOnlyOk returns a tuple with the IsOrgScopedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetIsOrgScopedOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.IsOrgScopedOnly) {
    return nil, false
	}
	return o.IsOrgScopedOnly, true
}

// HasIsOrgScopedOnly returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasIsOrgScopedOnly() bool {
	if o != nil && !isNil(o.IsOrgScopedOnly) {
		return true
	}

	return false
}

// SetIsOrgScopedOnly gets a reference to the given bool and assigns it to the IsOrgScopedOnly field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetIsOrgScopedOnly(v bool) {
	o.IsOrgScopedOnly = &v
}

// GetDocumentationLink returns the DocumentationLink field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDocumentationLink() string {
	if o == nil || isNil(o.DocumentationLink) {
		var ret string
		return ret
	}
	return *o.DocumentationLink
}

// GetDocumentationLinkOk returns a tuple with the DocumentationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetDocumentationLinkOk() (*string, bool) {
	if o == nil || isNil(o.DocumentationLink) {
    return nil, false
	}
	return o.DocumentationLink, true
}

// HasDocumentationLink returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasDocumentationLink() bool {
	if o != nil && !isNil(o.DocumentationLink) {
		return true
	}

	return false
}

// SetDocumentationLink gets a reference to the given string and assigns it to the DocumentationLink field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetDocumentationLink(v string) {
	o.DocumentationLink = &v
}

// GetSupportLink returns the SupportLink field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetSupportLink() string {
	if o == nil || isNil(o.SupportLink) {
		var ret string
		return ret
	}
	return *o.SupportLink
}

// GetSupportLinkOk returns a tuple with the SupportLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) GetSupportLinkOk() (*string, bool) {
	if o == nil || isNil(o.SupportLink) {
    return nil, false
	}
	return o.SupportLink, true
}

// HasSupportLink returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) HasSupportLink() bool {
	if o != nil && !isNil(o.SupportLink) {
		return true
	}

	return false
}

// SetSupportLink gets a reference to the given string and assigns it to the SupportLink field.
func (o *GetOrganizationEarlyAccessFeatures200ResponseInner) SetSupportLink(v string) {
	o.SupportLink = &v
}

func (o GetOrganizationEarlyAccessFeatures200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Descriptions) {
		toSerialize["descriptions"] = o.Descriptions
	}
	if !isNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !isNil(o.IsOrgScopedOnly) {
		toSerialize["isOrgScopedOnly"] = o.IsOrgScopedOnly
	}
	if !isNil(o.DocumentationLink) {
		toSerialize["documentationLink"] = o.DocumentationLink
	}
	if !isNil(o.SupportLink) {
		toSerialize["supportLink"] = o.SupportLink
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationEarlyAccessFeatures200ResponseInner struct {
	value *GetOrganizationEarlyAccessFeatures200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationEarlyAccessFeatures200ResponseInner) Get() *GetOrganizationEarlyAccessFeatures200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationEarlyAccessFeatures200ResponseInner) Set(val *GetOrganizationEarlyAccessFeatures200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationEarlyAccessFeatures200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationEarlyAccessFeatures200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationEarlyAccessFeatures200ResponseInner(val *GetOrganizationEarlyAccessFeatures200ResponseInner) *NullableGetOrganizationEarlyAccessFeatures200ResponseInner {
	return &NullableGetOrganizationEarlyAccessFeatures200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationEarlyAccessFeatures200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationEarlyAccessFeatures200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


