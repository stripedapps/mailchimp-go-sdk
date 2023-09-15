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

// checks if the CollectionOfElementStyleForListSignupForms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfElementStyleForListSignupForms{}

// CollectionOfElementStyleForListSignupForms Collection of Element style for List Signup Forms.
type CollectionOfElementStyleForListSignupForms struct {
	// A string that identifies the element selector.
	Selector *string `json:"selector,omitempty"`
	// A collection of options for a selector.
	Options []AnOptionForSignupFormStyles `json:"options,omitempty"`
}

// NewCollectionOfElementStyleForListSignupForms instantiates a new CollectionOfElementStyleForListSignupForms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfElementStyleForListSignupForms() *CollectionOfElementStyleForListSignupForms {
	this := CollectionOfElementStyleForListSignupForms{}
	return &this
}

// NewCollectionOfElementStyleForListSignupFormsWithDefaults instantiates a new CollectionOfElementStyleForListSignupForms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfElementStyleForListSignupFormsWithDefaults() *CollectionOfElementStyleForListSignupForms {
	this := CollectionOfElementStyleForListSignupForms{}
	return &this
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *CollectionOfElementStyleForListSignupForms) GetSelector() string {
	if o == nil || IsNil(o.Selector) {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfElementStyleForListSignupForms) GetSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *CollectionOfElementStyleForListSignupForms) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *CollectionOfElementStyleForListSignupForms) SetSelector(v string) {
	o.Selector = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CollectionOfElementStyleForListSignupForms) GetOptions() []AnOptionForSignupFormStyles {
	if o == nil || IsNil(o.Options) {
		var ret []AnOptionForSignupFormStyles
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfElementStyleForListSignupForms) GetOptionsOk() ([]AnOptionForSignupFormStyles, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CollectionOfElementStyleForListSignupForms) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []AnOptionForSignupFormStyles and assigns it to the Options field.
func (o *CollectionOfElementStyleForListSignupForms) SetOptions(v []AnOptionForSignupFormStyles) {
	o.Options = v
}

func (o CollectionOfElementStyleForListSignupForms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfElementStyleForListSignupForms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableCollectionOfElementStyleForListSignupForms struct {
	value *CollectionOfElementStyleForListSignupForms
	isSet bool
}

func (v NullableCollectionOfElementStyleForListSignupForms) Get() *CollectionOfElementStyleForListSignupForms {
	return v.value
}

func (v *NullableCollectionOfElementStyleForListSignupForms) Set(val *CollectionOfElementStyleForListSignupForms) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfElementStyleForListSignupForms) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfElementStyleForListSignupForms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfElementStyleForListSignupForms(val *CollectionOfElementStyleForListSignupForms) *NullableCollectionOfElementStyleForListSignupForms {
	return &NullableCollectionOfElementStyleForListSignupForms{value: val, isSet: true}
}

func (v NullableCollectionOfElementStyleForListSignupForms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfElementStyleForListSignupForms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


