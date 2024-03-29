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

// checks if the OptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionsInner{}

// OptionsInner struct for OptionsInner
type OptionsInner struct {
	// The label for this survey question option.
	Label *string `json:"label,omitempty"`
	// The ID for this survey question option.
	Id *string `json:"id,omitempty"`
	// The count of responses that selected this survey question option.
	Count *int32 `json:"count,omitempty"`
}

// NewOptionsInner instantiates a new OptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionsInner() *OptionsInner {
	this := OptionsInner{}
	return &this
}

// NewOptionsInnerWithDefaults instantiates a new OptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionsInnerWithDefaults() *OptionsInner {
	this := OptionsInner{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OptionsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OptionsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OptionsInner) SetLabel(v string) {
	o.Label = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OptionsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OptionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OptionsInner) SetId(v string) {
	o.Id = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OptionsInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionsInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OptionsInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OptionsInner) SetCount(v int32) {
	o.Count = &v
}

func (o OptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableOptionsInner struct {
	value *OptionsInner
	isSet bool
}

func (v NullableOptionsInner) Get() *OptionsInner {
	return v.value
}

func (v *NullableOptionsInner) Set(val *OptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionsInner(val *OptionsInner) *NullableOptionsInner {
	return &NullableOptionsInner{value: val, isSet: true}
}

func (v NullableOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


