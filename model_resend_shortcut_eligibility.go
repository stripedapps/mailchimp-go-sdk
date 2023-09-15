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

// checks if the ResendShortcutEligibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResendShortcutEligibility{}

// ResendShortcutEligibility Determines if the campaign qualifies for the Campaign Resend Shortcuts. Only included when query parameter `include_resend_shortcuts` is `true`.
type ResendShortcutEligibility struct {
	ToNonOpeners *ToNonOpeners `json:"to_non_openers,omitempty"`
	ToNewSubscribers *ToNewSubscribers `json:"to_new_subscribers,omitempty"`
	ToNonClickers *ToNonClickers `json:"to_non_clickers,omitempty"`
}

// NewResendShortcutEligibility instantiates a new ResendShortcutEligibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResendShortcutEligibility() *ResendShortcutEligibility {
	this := ResendShortcutEligibility{}
	return &this
}

// NewResendShortcutEligibilityWithDefaults instantiates a new ResendShortcutEligibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResendShortcutEligibilityWithDefaults() *ResendShortcutEligibility {
	this := ResendShortcutEligibility{}
	return &this
}

// GetToNonOpeners returns the ToNonOpeners field value if set, zero value otherwise.
func (o *ResendShortcutEligibility) GetToNonOpeners() ToNonOpeners {
	if o == nil || IsNil(o.ToNonOpeners) {
		var ret ToNonOpeners
		return ret
	}
	return *o.ToNonOpeners
}

// GetToNonOpenersOk returns a tuple with the ToNonOpeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResendShortcutEligibility) GetToNonOpenersOk() (*ToNonOpeners, bool) {
	if o == nil || IsNil(o.ToNonOpeners) {
		return nil, false
	}
	return o.ToNonOpeners, true
}

// HasToNonOpeners returns a boolean if a field has been set.
func (o *ResendShortcutEligibility) HasToNonOpeners() bool {
	if o != nil && !IsNil(o.ToNonOpeners) {
		return true
	}

	return false
}

// SetToNonOpeners gets a reference to the given ToNonOpeners and assigns it to the ToNonOpeners field.
func (o *ResendShortcutEligibility) SetToNonOpeners(v ToNonOpeners) {
	o.ToNonOpeners = &v
}

// GetToNewSubscribers returns the ToNewSubscribers field value if set, zero value otherwise.
func (o *ResendShortcutEligibility) GetToNewSubscribers() ToNewSubscribers {
	if o == nil || IsNil(o.ToNewSubscribers) {
		var ret ToNewSubscribers
		return ret
	}
	return *o.ToNewSubscribers
}

// GetToNewSubscribersOk returns a tuple with the ToNewSubscribers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResendShortcutEligibility) GetToNewSubscribersOk() (*ToNewSubscribers, bool) {
	if o == nil || IsNil(o.ToNewSubscribers) {
		return nil, false
	}
	return o.ToNewSubscribers, true
}

// HasToNewSubscribers returns a boolean if a field has been set.
func (o *ResendShortcutEligibility) HasToNewSubscribers() bool {
	if o != nil && !IsNil(o.ToNewSubscribers) {
		return true
	}

	return false
}

// SetToNewSubscribers gets a reference to the given ToNewSubscribers and assigns it to the ToNewSubscribers field.
func (o *ResendShortcutEligibility) SetToNewSubscribers(v ToNewSubscribers) {
	o.ToNewSubscribers = &v
}

// GetToNonClickers returns the ToNonClickers field value if set, zero value otherwise.
func (o *ResendShortcutEligibility) GetToNonClickers() ToNonClickers {
	if o == nil || IsNil(o.ToNonClickers) {
		var ret ToNonClickers
		return ret
	}
	return *o.ToNonClickers
}

// GetToNonClickersOk returns a tuple with the ToNonClickers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResendShortcutEligibility) GetToNonClickersOk() (*ToNonClickers, bool) {
	if o == nil || IsNil(o.ToNonClickers) {
		return nil, false
	}
	return o.ToNonClickers, true
}

// HasToNonClickers returns a boolean if a field has been set.
func (o *ResendShortcutEligibility) HasToNonClickers() bool {
	if o != nil && !IsNil(o.ToNonClickers) {
		return true
	}

	return false
}

// SetToNonClickers gets a reference to the given ToNonClickers and assigns it to the ToNonClickers field.
func (o *ResendShortcutEligibility) SetToNonClickers(v ToNonClickers) {
	o.ToNonClickers = &v
}

func (o ResendShortcutEligibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResendShortcutEligibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ToNonOpeners) {
		toSerialize["to_non_openers"] = o.ToNonOpeners
	}
	if !IsNil(o.ToNewSubscribers) {
		toSerialize["to_new_subscribers"] = o.ToNewSubscribers
	}
	if !IsNil(o.ToNonClickers) {
		toSerialize["to_non_clickers"] = o.ToNonClickers
	}
	return toSerialize, nil
}

type NullableResendShortcutEligibility struct {
	value *ResendShortcutEligibility
	isSet bool
}

func (v NullableResendShortcutEligibility) Get() *ResendShortcutEligibility {
	return v.value
}

func (v *NullableResendShortcutEligibility) Set(val *ResendShortcutEligibility) {
	v.value = val
	v.isSet = true
}

func (v NullableResendShortcutEligibility) IsSet() bool {
	return v.isSet
}

func (v *NullableResendShortcutEligibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResendShortcutEligibility(val *ResendShortcutEligibility) *NullableResendShortcutEligibility {
	return &NullableResendShortcutEligibility{value: val, isSet: true}
}

func (v NullableResendShortcutEligibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResendShortcutEligibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


