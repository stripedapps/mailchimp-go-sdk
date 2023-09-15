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

// checks if the MergeField1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeField1{}

// MergeField1 A [merge field](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.
type MergeField1 struct {
	// The merge tag used for Mailchimp campaigns and [adding contact information](https://mailchimp.com/developer/marketing/docs/merge-fields/#add-merge-data-to-contacts).
	Tag *string `json:"tag,omitempty"`
	// The name of the merge field (audience field).
	Name string `json:"name"`
	// The [type](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for the merge field.
	Type string `json:"type"`
	// Whether the merge field is required to import a contact.
	Required *bool `json:"required,omitempty"`
	// The default value for the merge field if `null`.
	DefaultValue *string `json:"default_value,omitempty"`
	// Whether the merge field is displayed on the signup form.
	Public *bool `json:"public,omitempty"`
	// The order that the merge field displays on the list signup form.
	DisplayOrder *int32 `json:"display_order,omitempty"`
	Options *MergeFieldOptions1 `json:"options,omitempty"`
	// Extra text to help the subscriber fill out the form.
	HelpText *string `json:"help_text,omitempty"`
}

// NewMergeField1 instantiates a new MergeField1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeField1(name string, type_ string) *MergeField1 {
	this := MergeField1{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewMergeField1WithDefaults instantiates a new MergeField1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeField1WithDefaults() *MergeField1 {
	this := MergeField1{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *MergeField1) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *MergeField1) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *MergeField1) SetTag(v string) {
	o.Tag = &v
}

// GetName returns the Name field value
func (o *MergeField1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MergeField1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MergeField1) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MergeField1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MergeField1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MergeField1) SetType(v string) {
	o.Type = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *MergeField1) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *MergeField1) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *MergeField1) SetRequired(v bool) {
	o.Required = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *MergeField1) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *MergeField1) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *MergeField1) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *MergeField1) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *MergeField1) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *MergeField1) SetPublic(v bool) {
	o.Public = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *MergeField1) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *MergeField1) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *MergeField1) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *MergeField1) GetOptions() MergeFieldOptions1 {
	if o == nil || IsNil(o.Options) {
		var ret MergeFieldOptions1
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetOptionsOk() (*MergeFieldOptions1, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MergeField1) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given MergeFieldOptions1 and assigns it to the Options field.
func (o *MergeField1) SetOptions(v MergeFieldOptions1) {
	o.Options = &v
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *MergeField1) GetHelpText() string {
	if o == nil || IsNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField1) GetHelpTextOk() (*string, bool) {
	if o == nil || IsNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *MergeField1) HasHelpText() bool {
	if o != nil && !IsNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *MergeField1) SetHelpText(v string) {
	o.HelpText = &v
}

func (o MergeField1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeField1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["default_value"] = o.DefaultValue
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["display_order"] = o.DisplayOrder
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.HelpText) {
		toSerialize["help_text"] = o.HelpText
	}
	return toSerialize, nil
}

type NullableMergeField1 struct {
	value *MergeField1
	isSet bool
}

func (v NullableMergeField1) Get() *MergeField1 {
	return v.value
}

func (v *NullableMergeField1) Set(val *MergeField1) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeField1) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeField1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeField1(val *MergeField1) *NullableMergeField1 {
	return &NullableMergeField1{value: val, isSet: true}
}

func (v NullableMergeField1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeField1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


