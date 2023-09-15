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

// checks if the SegmentOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SegmentOptions{}

// SegmentOptions An object representing all segmentation options.
type SegmentOptions struct {
	// The id for an existing saved segment.
	SavedSegmentId *int32 `json:"saved_segment_id,omitempty"`
	// Segment match type.
	Match *string `json:"match,omitempty"`
	// Segment match conditions. There are multiple possible types, see the [condition types documentation](https://mailchimp.com/developer/marketing/docs/alternative-schemas/#segment-condition-schemas).
	Conditions []map[string]interface{} `json:"conditions,omitempty"`
}

// NewSegmentOptions instantiates a new SegmentOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegmentOptions() *SegmentOptions {
	this := SegmentOptions{}
	return &this
}

// NewSegmentOptionsWithDefaults instantiates a new SegmentOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentOptionsWithDefaults() *SegmentOptions {
	this := SegmentOptions{}
	return &this
}

// GetSavedSegmentId returns the SavedSegmentId field value if set, zero value otherwise.
func (o *SegmentOptions) GetSavedSegmentId() int32 {
	if o == nil || IsNil(o.SavedSegmentId) {
		var ret int32
		return ret
	}
	return *o.SavedSegmentId
}

// GetSavedSegmentIdOk returns a tuple with the SavedSegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions) GetSavedSegmentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SavedSegmentId) {
		return nil, false
	}
	return o.SavedSegmentId, true
}

// HasSavedSegmentId returns a boolean if a field has been set.
func (o *SegmentOptions) HasSavedSegmentId() bool {
	if o != nil && !IsNil(o.SavedSegmentId) {
		return true
	}

	return false
}

// SetSavedSegmentId gets a reference to the given int32 and assigns it to the SavedSegmentId field.
func (o *SegmentOptions) SetSavedSegmentId(v int32) {
	o.SavedSegmentId = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *SegmentOptions) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *SegmentOptions) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *SegmentOptions) SetMatch(v string) {
	o.Match = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *SegmentOptions) GetConditions() []map[string]interface{} {
	if o == nil || IsNil(o.Conditions) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentOptions) GetConditionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *SegmentOptions) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []map[string]interface{} and assigns it to the Conditions field.
func (o *SegmentOptions) SetConditions(v []map[string]interface{}) {
	o.Conditions = v
}

func (o SegmentOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SegmentOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SavedSegmentId) {
		toSerialize["saved_segment_id"] = o.SavedSegmentId
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableSegmentOptions struct {
	value *SegmentOptions
	isSet bool
}

func (v NullableSegmentOptions) Get() *SegmentOptions {
	return v.value
}

func (v *NullableSegmentOptions) Set(val *SegmentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSegmentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSegmentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegmentOptions(val *SegmentOptions) *NullableSegmentOptions {
	return &NullableSegmentOptions{value: val, isSet: true}
}

func (v NullableSegmentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegmentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


