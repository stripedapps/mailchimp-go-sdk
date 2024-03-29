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

// checks if the AutomationTrigger1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutomationTrigger1{}

// AutomationTrigger1 Trigger settings for the Automation.
type AutomationTrigger1 struct {
	// The type of Automation workflow. Currently only supports 'abandonedCart'.
	WorkflowType string `json:"workflow_type"`
}

// NewAutomationTrigger1 instantiates a new AutomationTrigger1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTrigger1(workflowType string) *AutomationTrigger1 {
	this := AutomationTrigger1{}
	this.WorkflowType = workflowType
	return &this
}

// NewAutomationTrigger1WithDefaults instantiates a new AutomationTrigger1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTrigger1WithDefaults() *AutomationTrigger1 {
	this := AutomationTrigger1{}
	return &this
}

// GetWorkflowType returns the WorkflowType field value
func (o *AutomationTrigger1) GetWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value
// and a boolean to check if the value has been set.
func (o *AutomationTrigger1) GetWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowType, true
}

// SetWorkflowType sets field value
func (o *AutomationTrigger1) SetWorkflowType(v string) {
	o.WorkflowType = v
}

func (o AutomationTrigger1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomationTrigger1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflow_type"] = o.WorkflowType
	return toSerialize, nil
}

type NullableAutomationTrigger1 struct {
	value *AutomationTrigger1
	isSet bool
}

func (v NullableAutomationTrigger1) Get() *AutomationTrigger1 {
	return v.value
}

func (v *NullableAutomationTrigger1) Set(val *AutomationTrigger1) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationTrigger1) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationTrigger1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationTrigger1(val *AutomationTrigger1) *NullableAutomationTrigger1 {
	return &NullableAutomationTrigger1{value: val, isSet: true}
}

func (v NullableAutomationTrigger1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationTrigger1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


