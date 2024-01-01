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

// checks if the CreateDeviceLiveToolsPingRequestCallback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsPingRequestCallback{}

// CreateDeviceLiveToolsPingRequestCallback Details for the callback. Please include either an httpServerId OR url and sharedSecret
type CreateDeviceLiveToolsPingRequestCallback struct {
	// The callback URL for the webhook target. If using this field, please also specify a sharedSecret.
	Url *string `json:"url,omitempty"`
	// A shared secret that will be included in the requests sent to the callback URL. It can be used to verify that the request was sent by Meraki. If using this field, please also specify an url.
	SharedSecret *string `json:"sharedSecret,omitempty"`
	HttpServer *CreateDeviceLiveToolsPingRequestCallbackHttpServer `json:"httpServer,omitempty"`
	PayloadTemplate *CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate `json:"payloadTemplate,omitempty"`
}

// NewCreateDeviceLiveToolsPingRequestCallback instantiates a new CreateDeviceLiveToolsPingRequestCallback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsPingRequestCallback() *CreateDeviceLiveToolsPingRequestCallback {
	this := CreateDeviceLiveToolsPingRequestCallback{}
	return &this
}

// NewCreateDeviceLiveToolsPingRequestCallbackWithDefaults instantiates a new CreateDeviceLiveToolsPingRequestCallback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsPingRequestCallbackWithDefaults() *CreateDeviceLiveToolsPingRequestCallback {
	this := CreateDeviceLiveToolsPingRequestCallback{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CreateDeviceLiveToolsPingRequestCallback) SetUrl(v string) {
	o.Url = &v
}

// GetSharedSecret returns the SharedSecret field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetSharedSecret() string {
	if o == nil || IsNil(o.SharedSecret) {
		var ret string
		return ret
	}
	return *o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetSharedSecretOk() (*string, bool) {
	if o == nil || IsNil(o.SharedSecret) {
		return nil, false
	}
	return o.SharedSecret, true
}

// HasSharedSecret returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) HasSharedSecret() bool {
	if o != nil && !IsNil(o.SharedSecret) {
		return true
	}

	return false
}

// SetSharedSecret gets a reference to the given string and assigns it to the SharedSecret field.
func (o *CreateDeviceLiveToolsPingRequestCallback) SetSharedSecret(v string) {
	o.SharedSecret = &v
}

// GetHttpServer returns the HttpServer field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetHttpServer() CreateDeviceLiveToolsPingRequestCallbackHttpServer {
	if o == nil || IsNil(o.HttpServer) {
		var ret CreateDeviceLiveToolsPingRequestCallbackHttpServer
		return ret
	}
	return *o.HttpServer
}

// GetHttpServerOk returns a tuple with the HttpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetHttpServerOk() (*CreateDeviceLiveToolsPingRequestCallbackHttpServer, bool) {
	if o == nil || IsNil(o.HttpServer) {
		return nil, false
	}
	return o.HttpServer, true
}

// HasHttpServer returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) HasHttpServer() bool {
	if o != nil && !IsNil(o.HttpServer) {
		return true
	}

	return false
}

// SetHttpServer gets a reference to the given CreateDeviceLiveToolsPingRequestCallbackHttpServer and assigns it to the HttpServer field.
func (o *CreateDeviceLiveToolsPingRequestCallback) SetHttpServer(v CreateDeviceLiveToolsPingRequestCallbackHttpServer) {
	o.HttpServer = &v
}

// GetPayloadTemplate returns the PayloadTemplate field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetPayloadTemplate() CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate {
	if o == nil || IsNil(o.PayloadTemplate) {
		var ret CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate
		return ret
	}
	return *o.PayloadTemplate
}

// GetPayloadTemplateOk returns a tuple with the PayloadTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) GetPayloadTemplateOk() (*CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate, bool) {
	if o == nil || IsNil(o.PayloadTemplate) {
		return nil, false
	}
	return o.PayloadTemplate, true
}

// HasPayloadTemplate returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsPingRequestCallback) HasPayloadTemplate() bool {
	if o != nil && !IsNil(o.PayloadTemplate) {
		return true
	}

	return false
}

// SetPayloadTemplate gets a reference to the given CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate and assigns it to the PayloadTemplate field.
func (o *CreateDeviceLiveToolsPingRequestCallback) SetPayloadTemplate(v CreateDeviceLiveToolsPingRequestCallbackPayloadTemplate) {
	o.PayloadTemplate = &v
}

func (o CreateDeviceLiveToolsPingRequestCallback) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsPingRequestCallback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.SharedSecret) {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	if !IsNil(o.HttpServer) {
		toSerialize["httpServer"] = o.HttpServer
	}
	if !IsNil(o.PayloadTemplate) {
		toSerialize["payloadTemplate"] = o.PayloadTemplate
	}
	return toSerialize, nil
}

type NullableCreateDeviceLiveToolsPingRequestCallback struct {
	value *CreateDeviceLiveToolsPingRequestCallback
	isSet bool
}

func (v NullableCreateDeviceLiveToolsPingRequestCallback) Get() *CreateDeviceLiveToolsPingRequestCallback {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsPingRequestCallback) Set(val *CreateDeviceLiveToolsPingRequestCallback) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsPingRequestCallback) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsPingRequestCallback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsPingRequestCallback(val *CreateDeviceLiveToolsPingRequestCallback) *NullableCreateDeviceLiveToolsPingRequestCallback {
	return &NullableCreateDeviceLiveToolsPingRequestCallback{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsPingRequestCallback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsPingRequestCallback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


