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

// checks if the CampaignDefaults1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignDefaults1{}

// CampaignDefaults1 [Default values for campaigns](https://mailchimp.com/help/edit-your-emails-subject-preview-text-from-name-or-from-email-address/) created for this list.
type CampaignDefaults1 struct {
	// The default from name for campaigns sent to this list.
	FromName string `json:"from_name"`
	// The default from email for campaigns sent to this list.
	FromEmail string `json:"from_email"`
	// The default subject line for campaigns sent to this list.
	Subject string `json:"subject"`
	// The default language for this lists's forms.
	Language string `json:"language"`
}

// NewCampaignDefaults1 instantiates a new CampaignDefaults1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDefaults1(fromName string, fromEmail string, subject string, language string) *CampaignDefaults1 {
	this := CampaignDefaults1{}
	this.FromName = fromName
	this.FromEmail = fromEmail
	this.Subject = subject
	this.Language = language
	return &this
}

// NewCampaignDefaults1WithDefaults instantiates a new CampaignDefaults1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDefaults1WithDefaults() *CampaignDefaults1 {
	this := CampaignDefaults1{}
	return &this
}

// GetFromName returns the FromName field value
func (o *CampaignDefaults1) GetFromName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromName
}

// GetFromNameOk returns a tuple with the FromName field value
// and a boolean to check if the value has been set.
func (o *CampaignDefaults1) GetFromNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromName, true
}

// SetFromName sets field value
func (o *CampaignDefaults1) SetFromName(v string) {
	o.FromName = v
}

// GetFromEmail returns the FromEmail field value
func (o *CampaignDefaults1) GetFromEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromEmail
}

// GetFromEmailOk returns a tuple with the FromEmail field value
// and a boolean to check if the value has been set.
func (o *CampaignDefaults1) GetFromEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromEmail, true
}

// SetFromEmail sets field value
func (o *CampaignDefaults1) SetFromEmail(v string) {
	o.FromEmail = v
}

// GetSubject returns the Subject field value
func (o *CampaignDefaults1) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *CampaignDefaults1) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *CampaignDefaults1) SetSubject(v string) {
	o.Subject = v
}

// GetLanguage returns the Language field value
func (o *CampaignDefaults1) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *CampaignDefaults1) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *CampaignDefaults1) SetLanguage(v string) {
	o.Language = v
}

func (o CampaignDefaults1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignDefaults1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_name"] = o.FromName
	toSerialize["from_email"] = o.FromEmail
	toSerialize["subject"] = o.Subject
	toSerialize["language"] = o.Language
	return toSerialize, nil
}

type NullableCampaignDefaults1 struct {
	value *CampaignDefaults1
	isSet bool
}

func (v NullableCampaignDefaults1) Get() *CampaignDefaults1 {
	return v.value
}

func (v *NullableCampaignDefaults1) Set(val *CampaignDefaults1) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDefaults1) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDefaults1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDefaults1(val *CampaignDefaults1) *NullableCampaignDefaults1 {
	return &NullableCampaignDefaults1{value: val, isSet: true}
}

func (v NullableCampaignDefaults1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDefaults1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


