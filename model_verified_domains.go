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

// checks if the VerifiedDomains type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerifiedDomains{}

// VerifiedDomains The verified domains currently on the account.
type VerifiedDomains struct {
	// The name of this domain.
	Domain *string `json:"domain,omitempty"`
	// Whether the domain has been verified for sending.
	Verified *bool `json:"verified,omitempty"`
	// Whether domain authentication is enabled for this domain.
	Authenticated *bool `json:"authenticated,omitempty"`
	// The e-mail address receiving the two-factor challenge for this domain.
	VerificationEmail *string `json:"verification_email,omitempty"`
	// The date/time that the two-factor challenge was sent to the verification email.
	VerificationSent *time.Time `json:"verification_sent,omitempty"`
}

// NewVerifiedDomains instantiates a new VerifiedDomains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifiedDomains() *VerifiedDomains {
	this := VerifiedDomains{}
	return &this
}

// NewVerifiedDomainsWithDefaults instantiates a new VerifiedDomains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifiedDomainsWithDefaults() *VerifiedDomains {
	this := VerifiedDomains{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *VerifiedDomains) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *VerifiedDomains) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *VerifiedDomains) SetDomain(v string) {
	o.Domain = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *VerifiedDomains) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *VerifiedDomains) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *VerifiedDomains) SetVerified(v bool) {
	o.Verified = &v
}

// GetAuthenticated returns the Authenticated field value if set, zero value otherwise.
func (o *VerifiedDomains) GetAuthenticated() bool {
	if o == nil || IsNil(o.Authenticated) {
		var ret bool
		return ret
	}
	return *o.Authenticated
}

// GetAuthenticatedOk returns a tuple with the Authenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains) GetAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Authenticated) {
		return nil, false
	}
	return o.Authenticated, true
}

// HasAuthenticated returns a boolean if a field has been set.
func (o *VerifiedDomains) HasAuthenticated() bool {
	if o != nil && !IsNil(o.Authenticated) {
		return true
	}

	return false
}

// SetAuthenticated gets a reference to the given bool and assigns it to the Authenticated field.
func (o *VerifiedDomains) SetAuthenticated(v bool) {
	o.Authenticated = &v
}

// GetVerificationEmail returns the VerificationEmail field value if set, zero value otherwise.
func (o *VerifiedDomains) GetVerificationEmail() string {
	if o == nil || IsNil(o.VerificationEmail) {
		var ret string
		return ret
	}
	return *o.VerificationEmail
}

// GetVerificationEmailOk returns a tuple with the VerificationEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains) GetVerificationEmailOk() (*string, bool) {
	if o == nil || IsNil(o.VerificationEmail) {
		return nil, false
	}
	return o.VerificationEmail, true
}

// HasVerificationEmail returns a boolean if a field has been set.
func (o *VerifiedDomains) HasVerificationEmail() bool {
	if o != nil && !IsNil(o.VerificationEmail) {
		return true
	}

	return false
}

// SetVerificationEmail gets a reference to the given string and assigns it to the VerificationEmail field.
func (o *VerifiedDomains) SetVerificationEmail(v string) {
	o.VerificationEmail = &v
}

// GetVerificationSent returns the VerificationSent field value if set, zero value otherwise.
func (o *VerifiedDomains) GetVerificationSent() time.Time {
	if o == nil || IsNil(o.VerificationSent) {
		var ret time.Time
		return ret
	}
	return *o.VerificationSent
}

// GetVerificationSentOk returns a tuple with the VerificationSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifiedDomains) GetVerificationSentOk() (*time.Time, bool) {
	if o == nil || IsNil(o.VerificationSent) {
		return nil, false
	}
	return o.VerificationSent, true
}

// HasVerificationSent returns a boolean if a field has been set.
func (o *VerifiedDomains) HasVerificationSent() bool {
	if o != nil && !IsNil(o.VerificationSent) {
		return true
	}

	return false
}

// SetVerificationSent gets a reference to the given time.Time and assigns it to the VerificationSent field.
func (o *VerifiedDomains) SetVerificationSent(v time.Time) {
	o.VerificationSent = &v
}

func (o VerifiedDomains) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerifiedDomains) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !IsNil(o.Authenticated) {
		toSerialize["authenticated"] = o.Authenticated
	}
	if !IsNil(o.VerificationEmail) {
		toSerialize["verification_email"] = o.VerificationEmail
	}
	if !IsNil(o.VerificationSent) {
		toSerialize["verification_sent"] = o.VerificationSent
	}
	return toSerialize, nil
}

type NullableVerifiedDomains struct {
	value *VerifiedDomains
	isSet bool
}

func (v NullableVerifiedDomains) Get() *VerifiedDomains {
	return v.value
}

func (v *NullableVerifiedDomains) Set(val *VerifiedDomains) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifiedDomains) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifiedDomains) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifiedDomains(val *VerifiedDomains) *NullableVerifiedDomains {
	return &NullableVerifiedDomains{value: val, isSet: true}
}

func (v NullableVerifiedDomains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifiedDomains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


