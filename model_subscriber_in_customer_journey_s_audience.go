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

// checks if the SubscriberInCustomerJourneySAudience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberInCustomerJourneySAudience{}

// SubscriberInCustomerJourneySAudience Information about subscribers in a Customer Journey's audience.
type SubscriberInCustomerJourneySAudience struct {
	// The list member's email address.
	EmailAddress string `json:"email_address"`
}

// NewSubscriberInCustomerJourneySAudience instantiates a new SubscriberInCustomerJourneySAudience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberInCustomerJourneySAudience(emailAddress string) *SubscriberInCustomerJourneySAudience {
	this := SubscriberInCustomerJourneySAudience{}
	this.EmailAddress = emailAddress
	return &this
}

// NewSubscriberInCustomerJourneySAudienceWithDefaults instantiates a new SubscriberInCustomerJourneySAudience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberInCustomerJourneySAudienceWithDefaults() *SubscriberInCustomerJourneySAudience {
	this := SubscriberInCustomerJourneySAudience{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *SubscriberInCustomerJourneySAudience) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *SubscriberInCustomerJourneySAudience) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *SubscriberInCustomerJourneySAudience) SetEmailAddress(v string) {
	o.EmailAddress = v
}

func (o SubscriberInCustomerJourneySAudience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberInCustomerJourneySAudience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_address"] = o.EmailAddress
	return toSerialize, nil
}

type NullableSubscriberInCustomerJourneySAudience struct {
	value *SubscriberInCustomerJourneySAudience
	isSet bool
}

func (v NullableSubscriberInCustomerJourneySAudience) Get() *SubscriberInCustomerJourneySAudience {
	return v.value
}

func (v *NullableSubscriberInCustomerJourneySAudience) Set(val *SubscriberInCustomerJourneySAudience) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberInCustomerJourneySAudience) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberInCustomerJourneySAudience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberInCustomerJourneySAudience(val *SubscriberInCustomerJourneySAudience) *NullableSubscriberInCustomerJourneySAudience {
	return &NullableSubscriberInCustomerJourneySAudience{value: val, isSet: true}
}

func (v NullableSubscriberInCustomerJourneySAudience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberInCustomerJourneySAudience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


