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

// checks if the CreateNetworkWebhooksPayloadTemplateRequestHeadersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWebhooksPayloadTemplateRequestHeadersInner{}

// CreateNetworkWebhooksPayloadTemplateRequestHeadersInner struct for CreateNetworkWebhooksPayloadTemplateRequestHeadersInner
type CreateNetworkWebhooksPayloadTemplateRequestHeadersInner struct {
	// The name of the header template
	Name *string `json:"name,omitempty"`
	// The liquid template for the headers
	Template *string `json:"template,omitempty"`
}

// NewCreateNetworkWebhooksPayloadTemplateRequestHeadersInner instantiates a new CreateNetworkWebhooksPayloadTemplateRequestHeadersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWebhooksPayloadTemplateRequestHeadersInner() *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner {
	this := CreateNetworkWebhooksPayloadTemplateRequestHeadersInner{}
	return &this
}

// NewCreateNetworkWebhooksPayloadTemplateRequestHeadersInnerWithDefaults instantiates a new CreateNetworkWebhooksPayloadTemplateRequestHeadersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWebhooksPayloadTemplateRequestHeadersInnerWithDefaults() *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner {
	this := CreateNetworkWebhooksPayloadTemplateRequestHeadersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) SetName(v string) {
	o.Name = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) SetTemplate(v string) {
	o.Template = &v
}

func (o CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner struct {
	value *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner
	isSet bool
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner) Get() *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner {
	return v.value
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner) Set(val *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner(val *CreateNetworkWebhooksPayloadTemplateRequestHeadersInner) *NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner {
	return &NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner{value: val, isSet: true}
}

func (v NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWebhooksPayloadTemplateRequestHeadersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


