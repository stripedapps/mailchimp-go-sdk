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

// checks if the List8 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &List8{}

// List8 Information about a specific list segment.
type List8 struct {
	// The name of the segment.
	Name string `json:"name"`
	// An array of emails to be used for a static segment. Any emails provided that are not present on the list will be ignored. Passing an empty array will create a static segment without any subscribers. This field cannot be provided with the options field.
	StaticSegment []string `json:"static_segment,omitempty"`
	Options *Conditions1 `json:"options,omitempty"`
}

// NewList8 instantiates a new List8 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewList8(name string) *List8 {
	this := List8{}
	this.Name = name
	return &this
}

// NewList8WithDefaults instantiates a new List8 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewList8WithDefaults() *List8 {
	this := List8{}
	return &this
}

// GetName returns the Name field value
func (o *List8) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *List8) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *List8) SetName(v string) {
	o.Name = v
}

// GetStaticSegment returns the StaticSegment field value if set, zero value otherwise.
func (o *List8) GetStaticSegment() []string {
	if o == nil || IsNil(o.StaticSegment) {
		var ret []string
		return ret
	}
	return o.StaticSegment
}

// GetStaticSegmentOk returns a tuple with the StaticSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List8) GetStaticSegmentOk() ([]string, bool) {
	if o == nil || IsNil(o.StaticSegment) {
		return nil, false
	}
	return o.StaticSegment, true
}

// HasStaticSegment returns a boolean if a field has been set.
func (o *List8) HasStaticSegment() bool {
	if o != nil && !IsNil(o.StaticSegment) {
		return true
	}

	return false
}

// SetStaticSegment gets a reference to the given []string and assigns it to the StaticSegment field.
func (o *List8) SetStaticSegment(v []string) {
	o.StaticSegment = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *List8) GetOptions() Conditions1 {
	if o == nil || IsNil(o.Options) {
		var ret Conditions1
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *List8) GetOptionsOk() (*Conditions1, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *List8) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Conditions1 and assigns it to the Options field.
func (o *List8) SetOptions(v Conditions1) {
	o.Options = &v
}

func (o List8) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o List8) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.StaticSegment) {
		toSerialize["static_segment"] = o.StaticSegment
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableList8 struct {
	value *List8
	isSet bool
}

func (v NullableList8) Get() *List8 {
	return v.value
}

func (v *NullableList8) Set(val *List8) {
	v.value = val
	v.isSet = true
}

func (v NullableList8) IsSet() bool {
	return v.isSet
}

func (v *NullableList8) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableList8(val *List8) *NullableList8 {
	return &NullableList8{value: val, isSet: true}
}

func (v NullableList8) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableList8) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


