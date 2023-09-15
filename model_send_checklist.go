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

// checks if the SendChecklist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendChecklist{}

// SendChecklist The send checklist for the campaign.
type SendChecklist struct {
	// Whether the campaign is ready to send.
	IsReady *bool `json:"is_ready,omitempty"`
	// A list of feedback items to review before sending your campaign.
	Items []ItemsInner `json:"items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewSendChecklist instantiates a new SendChecklist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendChecklist() *SendChecklist {
	this := SendChecklist{}
	return &this
}

// NewSendChecklistWithDefaults instantiates a new SendChecklist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendChecklistWithDefaults() *SendChecklist {
	this := SendChecklist{}
	return &this
}

// GetIsReady returns the IsReady field value if set, zero value otherwise.
func (o *SendChecklist) GetIsReady() bool {
	if o == nil || IsNil(o.IsReady) {
		var ret bool
		return ret
	}
	return *o.IsReady
}

// GetIsReadyOk returns a tuple with the IsReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChecklist) GetIsReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReady) {
		return nil, false
	}
	return o.IsReady, true
}

// HasIsReady returns a boolean if a field has been set.
func (o *SendChecklist) HasIsReady() bool {
	if o != nil && !IsNil(o.IsReady) {
		return true
	}

	return false
}

// SetIsReady gets a reference to the given bool and assigns it to the IsReady field.
func (o *SendChecklist) SetIsReady(v bool) {
	o.IsReady = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SendChecklist) GetItems() []ItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []ItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChecklist) GetItemsOk() ([]ItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SendChecklist) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ItemsInner and assigns it to the Items field.
func (o *SendChecklist) SetItems(v []ItemsInner) {
	o.Items = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SendChecklist) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendChecklist) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SendChecklist) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *SendChecklist) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o SendChecklist) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendChecklist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsReady) {
		toSerialize["is_ready"] = o.IsReady
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSendChecklist struct {
	value *SendChecklist
	isSet bool
}

func (v NullableSendChecklist) Get() *SendChecklist {
	return v.value
}

func (v *NullableSendChecklist) Set(val *SendChecklist) {
	v.value = val
	v.isSet = true
}

func (v NullableSendChecklist) IsSet() bool {
	return v.isSet
}

func (v *NullableSendChecklist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendChecklist(val *SendChecklist) *NullableSendChecklist {
	return &NullableSendChecklist{value: val, isSet: true}
}

func (v NullableSendChecklist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendChecklist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


