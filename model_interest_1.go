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

// checks if the Interest1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Interest1{}

// Interest1 Assign subscribers to interests to group them together. Interests are referred to as 'group names' in the Mailchimp application.
type Interest1 struct {
	// The name of the interest. This can be shown publicly on a subscription form.
	Name string `json:"name"`
	// The display order for interests.
	DisplayOrder *int32 `json:"display_order,omitempty"`
}

// NewInterest1 instantiates a new Interest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterest1(name string) *Interest1 {
	this := Interest1{}
	this.Name = name
	return &this
}

// NewInterest1WithDefaults instantiates a new Interest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterest1WithDefaults() *Interest1 {
	this := Interest1{}
	return &this
}

// GetName returns the Name field value
func (o *Interest1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Interest1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Interest1) SetName(v string) {
	o.Name = v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *Interest1) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interest1) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *Interest1) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *Interest1) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

func (o Interest1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Interest1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DisplayOrder) {
		toSerialize["display_order"] = o.DisplayOrder
	}
	return toSerialize, nil
}

type NullableInterest1 struct {
	value *Interest1
	isSet bool
}

func (v NullableInterest1) Get() *Interest1 {
	return v.value
}

func (v *NullableInterest1) Set(val *Interest1) {
	v.value = val
	v.isSet = true
}

func (v NullableInterest1) IsSet() bool {
	return v.isSet
}

func (v *NullableInterest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterest1(val *Interest1) *NullableInterest1 {
	return &NullableInterest1{value: val, isSet: true}
}

func (v NullableInterest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


