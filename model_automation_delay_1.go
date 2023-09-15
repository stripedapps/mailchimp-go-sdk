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

// checks if the AutomationDelay1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutomationDelay1{}

// AutomationDelay1 The delay settings for an automation email.
type AutomationDelay1 struct {
	// The delay amount for an automation email.
	Amount *int32 `json:"amount,omitempty"`
	// The type of delay for an automation email.
	Type *string `json:"type,omitempty"`
	// Whether the delay settings describe before or after the delay action of an automation email.
	Direction *string `json:"direction,omitempty"`
	// The action that triggers the delay of an automation emails.
	Action string `json:"action"`
}

// NewAutomationDelay1 instantiates a new AutomationDelay1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationDelay1(action string) *AutomationDelay1 {
	this := AutomationDelay1{}
	this.Action = action
	return &this
}

// NewAutomationDelay1WithDefaults instantiates a new AutomationDelay1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationDelay1WithDefaults() *AutomationDelay1 {
	this := AutomationDelay1{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AutomationDelay1) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationDelay1) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AutomationDelay1) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *AutomationDelay1) SetAmount(v int32) {
	o.Amount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AutomationDelay1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationDelay1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AutomationDelay1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AutomationDelay1) SetType(v string) {
	o.Type = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *AutomationDelay1) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationDelay1) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *AutomationDelay1) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *AutomationDelay1) SetDirection(v string) {
	o.Direction = &v
}

// GetAction returns the Action field value
func (o *AutomationDelay1) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *AutomationDelay1) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *AutomationDelay1) SetAction(v string) {
	o.Action = v
}

func (o AutomationDelay1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomationDelay1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableAutomationDelay1 struct {
	value *AutomationDelay1
	isSet bool
}

func (v NullableAutomationDelay1) Get() *AutomationDelay1 {
	return v.value
}

func (v *NullableAutomationDelay1) Set(val *AutomationDelay1) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationDelay1) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationDelay1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationDelay1(val *AutomationDelay1) *NullableAutomationDelay1 {
	return &NullableAutomationDelay1{value: val, isSet: true}
}

func (v NullableAutomationDelay1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationDelay1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


