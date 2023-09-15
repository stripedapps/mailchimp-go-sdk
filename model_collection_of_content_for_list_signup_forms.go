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

// checks if the CollectionOfContentForListSignupForms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfContentForListSignupForms{}

// CollectionOfContentForListSignupForms Collection of Content for List Signup Forms.
type CollectionOfContentForListSignupForms struct {
	// The content section name.
	Section *string `json:"section,omitempty"`
	// The content section text.
	Value *string `json:"value,omitempty"`
}

// NewCollectionOfContentForListSignupForms instantiates a new CollectionOfContentForListSignupForms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfContentForListSignupForms() *CollectionOfContentForListSignupForms {
	this := CollectionOfContentForListSignupForms{}
	return &this
}

// NewCollectionOfContentForListSignupFormsWithDefaults instantiates a new CollectionOfContentForListSignupForms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfContentForListSignupFormsWithDefaults() *CollectionOfContentForListSignupForms {
	this := CollectionOfContentForListSignupForms{}
	return &this
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *CollectionOfContentForListSignupForms) GetSection() string {
	if o == nil || IsNil(o.Section) {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfContentForListSignupForms) GetSectionOk() (*string, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *CollectionOfContentForListSignupForms) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *CollectionOfContentForListSignupForms) SetSection(v string) {
	o.Section = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfContentForListSignupForms) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfContentForListSignupForms) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfContentForListSignupForms) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CollectionOfContentForListSignupForms) SetValue(v string) {
	o.Value = &v
}

func (o CollectionOfContentForListSignupForms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfContentForListSignupForms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCollectionOfContentForListSignupForms struct {
	value *CollectionOfContentForListSignupForms
	isSet bool
}

func (v NullableCollectionOfContentForListSignupForms) Get() *CollectionOfContentForListSignupForms {
	return v.value
}

func (v *NullableCollectionOfContentForListSignupForms) Set(val *CollectionOfContentForListSignupForms) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfContentForListSignupForms) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfContentForListSignupForms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfContentForListSignupForms(val *CollectionOfContentForListSignupForms) *NullableCollectionOfContentForListSignupForms {
	return &NullableCollectionOfContentForListSignupForms{value: val, isSet: true}
}

func (v NullableCollectionOfContentForListSignupForms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfContentForListSignupForms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


