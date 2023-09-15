/*
Mailchimp Marketing API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.55
Contact: apihelp@mailchimp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AddWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddWebhook{}

// AddWebhook Configure a webhook for the given list.
type AddWebhook struct {
	// A valid URL for the Webhook.
	Url *string `json:"url,omitempty"`
	Events *Events2 `json:"events,omitempty"`
	Sources *Sources1 `json:"sources,omitempty"`
}

// NewAddWebhook instantiates a new AddWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddWebhook() *AddWebhook {
	this := AddWebhook{}
	return &this
}

// NewAddWebhookWithDefaults instantiates a new AddWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddWebhookWithDefaults() *AddWebhook {
	this := AddWebhook{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AddWebhook) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWebhook) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AddWebhook) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AddWebhook) SetUrl(v string) {
	o.Url = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *AddWebhook) GetEvents() Events2 {
	if o == nil || IsNil(o.Events) {
		var ret Events2
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWebhook) GetEventsOk() (*Events2, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *AddWebhook) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given Events2 and assigns it to the Events field.
func (o *AddWebhook) SetEvents(v Events2) {
	o.Events = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *AddWebhook) GetSources() Sources1 {
	if o == nil || IsNil(o.Sources) {
		var ret Sources1
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddWebhook) GetSourcesOk() (*Sources1, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *AddWebhook) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given Sources1 and assigns it to the Sources field.
func (o *AddWebhook) SetSources(v Sources1) {
	o.Sources = &v
}

func (o AddWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	return toSerialize, nil
}

type NullableAddWebhook struct {
	value *AddWebhook
	isSet bool
}

func (v NullableAddWebhook) Get() *AddWebhook {
	return v.value
}

func (v *NullableAddWebhook) Set(val *AddWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableAddWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableAddWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddWebhook(val *AddWebhook) *NullableAddWebhook {
	return &NullableAddWebhook{value: val, isSet: true}
}

func (v NullableAddWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


