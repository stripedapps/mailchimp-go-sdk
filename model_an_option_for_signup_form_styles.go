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

// checks if the AnOptionForSignupFormStyles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnOptionForSignupFormStyles{}

// AnOptionForSignupFormStyles An option for Signup Form Styles.
type AnOptionForSignupFormStyles struct {
	// A string that identifies the property.
	Property *string `json:"property,omitempty"`
	// A string that identifies value of the property.
	Value *string `json:"value,omitempty"`
}

// NewAnOptionForSignupFormStyles instantiates a new AnOptionForSignupFormStyles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnOptionForSignupFormStyles() *AnOptionForSignupFormStyles {
	this := AnOptionForSignupFormStyles{}
	return &this
}

// NewAnOptionForSignupFormStylesWithDefaults instantiates a new AnOptionForSignupFormStyles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnOptionForSignupFormStylesWithDefaults() *AnOptionForSignupFormStyles {
	this := AnOptionForSignupFormStyles{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *AnOptionForSignupFormStyles) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnOptionForSignupFormStyles) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *AnOptionForSignupFormStyles) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *AnOptionForSignupFormStyles) SetProperty(v string) {
	o.Property = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AnOptionForSignupFormStyles) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnOptionForSignupFormStyles) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AnOptionForSignupFormStyles) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AnOptionForSignupFormStyles) SetValue(v string) {
	o.Value = &v
}

func (o AnOptionForSignupFormStyles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnOptionForSignupFormStyles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAnOptionForSignupFormStyles struct {
	value *AnOptionForSignupFormStyles
	isSet bool
}

func (v NullableAnOptionForSignupFormStyles) Get() *AnOptionForSignupFormStyles {
	return v.value
}

func (v *NullableAnOptionForSignupFormStyles) Set(val *AnOptionForSignupFormStyles) {
	v.value = val
	v.isSet = true
}

func (v NullableAnOptionForSignupFormStyles) IsSet() bool {
	return v.isSet
}

func (v *NullableAnOptionForSignupFormStyles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnOptionForSignupFormStyles(val *AnOptionForSignupFormStyles) *NullableAnOptionForSignupFormStyles {
	return &NullableAnOptionForSignupFormStyles{value: val, isSet: true}
}

func (v NullableAnOptionForSignupFormStyles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnOptionForSignupFormStyles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


