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
	"time"
)

// checks if the MemberActivity1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberActivity1{}

// MemberActivity1 A summary of the interaction with the campaign.
type MemberActivity1 struct {
	// The date and time recorded for the action in ISO 8601 format.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewMemberActivity1 instantiates a new MemberActivity1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberActivity1() *MemberActivity1 {
	this := MemberActivity1{}
	return &this
}

// NewMemberActivity1WithDefaults instantiates a new MemberActivity1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberActivity1WithDefaults() *MemberActivity1 {
	this := MemberActivity1{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MemberActivity1) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberActivity1) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MemberActivity1) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MemberActivity1) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o MemberActivity1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberActivity1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableMemberActivity1 struct {
	value *MemberActivity1
	isSet bool
}

func (v NullableMemberActivity1) Get() *MemberActivity1 {
	return v.value
}

func (v *NullableMemberActivity1) Set(val *MemberActivity1) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberActivity1) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberActivity1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberActivity1(val *MemberActivity1) *NullableMemberActivity1 {
	return &NullableMemberActivity1{value: val, isSet: true}
}

func (v NullableMemberActivity1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberActivity1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


