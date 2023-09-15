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

// checks if the AddListMembers1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddListMembers1{}

// AddListMembers1 Individuals who are currently or have been previously subscribed to this list, including members who have bounced or unsubscribed.
type AddListMembers1 struct {
	// Email address for a subscriber.
	EmailAddress string `json:"email_address"`
	// Type of email this member asked to get ('html' or 'text').
	EmailType *string `json:"email_type,omitempty"`
	// Subscriber's current status.
	Status string `json:"status"`
	// A dictionary of merge fields where the keys are the merge tags. See the [Merge Fields documentation](https://mailchimp.com/developer/marketing/docs/merge-fields/#structure) for more about the structure.
	MergeFields map[string]map[string]interface{} `json:"merge_fields,omitempty"`
	// The key of this object's properties is the ID of the interest in question.
	Interests *map[string]bool `json:"interests,omitempty"`
	// If set/detected, the [subscriber's language](https://mailchimp.com/help/view-and-edit-contact-languages/).
	Language *string `json:"language,omitempty"`
	// [VIP status](https://mailchimp.com/help/designate-and-send-to-vip-contacts/) for subscriber.
	Vip *bool `json:"vip,omitempty"`
	Location *Location `json:"location,omitempty"`
	// The marketing permissions for the subscriber.
	MarketingPermissions []MarketingPermission1 `json:"marketing_permissions,omitempty"`
	// IP address the subscriber signed up from.
	IpSignup *string `json:"ip_signup,omitempty"`
	// The date and time the subscriber signed up for the list in ISO 8601 format.
	TimestampSignup *time.Time `json:"timestamp_signup,omitempty"`
	// The IP address the subscriber used to confirm their opt-in status.
	IpOpt *string `json:"ip_opt,omitempty"`
	// The date and time the subscriber confirmed their opt-in status in ISO 8601 format.
	TimestampOpt *time.Time `json:"timestamp_opt,omitempty"`
	// The tags that are associated with a member.
	Tags []string `json:"tags,omitempty"`
}

// NewAddListMembers1 instantiates a new AddListMembers1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddListMembers1(emailAddress string, status string) *AddListMembers1 {
	this := AddListMembers1{}
	this.EmailAddress = emailAddress
	this.Status = status
	return &this
}

// NewAddListMembers1WithDefaults instantiates a new AddListMembers1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddListMembers1WithDefaults() *AddListMembers1 {
	this := AddListMembers1{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *AddListMembers1) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *AddListMembers1) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *AddListMembers1) GetEmailType() string {
	if o == nil || IsNil(o.EmailType) {
		var ret string
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetEmailTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HasEmailType returns a boolean if a field has been set.
func (o *AddListMembers1) HasEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given string and assigns it to the EmailType field.
func (o *AddListMembers1) SetEmailType(v string) {
	o.EmailType = &v
}

// GetStatus returns the Status field value
func (o *AddListMembers1) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AddListMembers1) SetStatus(v string) {
	o.Status = v
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *AddListMembers1) GetMergeFields() map[string]map[string]interface{} {
	if o == nil || IsNil(o.MergeFields) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetMergeFieldsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return map[string]map[string]interface{}{}, false
	}
	return o.MergeFields, true
}

// HasMergeFields returns a boolean if a field has been set.
func (o *AddListMembers1) HasMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given map[string]map[string]interface{} and assigns it to the MergeFields field.
func (o *AddListMembers1) SetMergeFields(v map[string]map[string]interface{}) {
	o.MergeFields = v
}

// GetInterests returns the Interests field value if set, zero value otherwise.
func (o *AddListMembers1) GetInterests() map[string]bool {
	if o == nil || IsNil(o.Interests) {
		var ret map[string]bool
		return ret
	}
	return *o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetInterestsOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.Interests) {
		return nil, false
	}
	return o.Interests, true
}

// HasInterests returns a boolean if a field has been set.
func (o *AddListMembers1) HasInterests() bool {
	if o != nil && !IsNil(o.Interests) {
		return true
	}

	return false
}

// SetInterests gets a reference to the given map[string]bool and assigns it to the Interests field.
func (o *AddListMembers1) SetInterests(v map[string]bool) {
	o.Interests = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AddListMembers1) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AddListMembers1) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AddListMembers1) SetLanguage(v string) {
	o.Language = &v
}

// GetVip returns the Vip field value if set, zero value otherwise.
func (o *AddListMembers1) GetVip() bool {
	if o == nil || IsNil(o.Vip) {
		var ret bool
		return ret
	}
	return *o.Vip
}

// GetVipOk returns a tuple with the Vip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetVipOk() (*bool, bool) {
	if o == nil || IsNil(o.Vip) {
		return nil, false
	}
	return o.Vip, true
}

// HasVip returns a boolean if a field has been set.
func (o *AddListMembers1) HasVip() bool {
	if o != nil && !IsNil(o.Vip) {
		return true
	}

	return false
}

// SetVip gets a reference to the given bool and assigns it to the Vip field.
func (o *AddListMembers1) SetVip(v bool) {
	o.Vip = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AddListMembers1) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AddListMembers1) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *AddListMembers1) SetLocation(v Location) {
	o.Location = &v
}

// GetMarketingPermissions returns the MarketingPermissions field value if set, zero value otherwise.
func (o *AddListMembers1) GetMarketingPermissions() []MarketingPermission1 {
	if o == nil || IsNil(o.MarketingPermissions) {
		var ret []MarketingPermission1
		return ret
	}
	return o.MarketingPermissions
}

// GetMarketingPermissionsOk returns a tuple with the MarketingPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetMarketingPermissionsOk() ([]MarketingPermission1, bool) {
	if o == nil || IsNil(o.MarketingPermissions) {
		return nil, false
	}
	return o.MarketingPermissions, true
}

// HasMarketingPermissions returns a boolean if a field has been set.
func (o *AddListMembers1) HasMarketingPermissions() bool {
	if o != nil && !IsNil(o.MarketingPermissions) {
		return true
	}

	return false
}

// SetMarketingPermissions gets a reference to the given []MarketingPermission1 and assigns it to the MarketingPermissions field.
func (o *AddListMembers1) SetMarketingPermissions(v []MarketingPermission1) {
	o.MarketingPermissions = v
}

// GetIpSignup returns the IpSignup field value if set, zero value otherwise.
func (o *AddListMembers1) GetIpSignup() string {
	if o == nil || IsNil(o.IpSignup) {
		var ret string
		return ret
	}
	return *o.IpSignup
}

// GetIpSignupOk returns a tuple with the IpSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetIpSignupOk() (*string, bool) {
	if o == nil || IsNil(o.IpSignup) {
		return nil, false
	}
	return o.IpSignup, true
}

// HasIpSignup returns a boolean if a field has been set.
func (o *AddListMembers1) HasIpSignup() bool {
	if o != nil && !IsNil(o.IpSignup) {
		return true
	}

	return false
}

// SetIpSignup gets a reference to the given string and assigns it to the IpSignup field.
func (o *AddListMembers1) SetIpSignup(v string) {
	o.IpSignup = &v
}

// GetTimestampSignup returns the TimestampSignup field value if set, zero value otherwise.
func (o *AddListMembers1) GetTimestampSignup() time.Time {
	if o == nil || IsNil(o.TimestampSignup) {
		var ret time.Time
		return ret
	}
	return *o.TimestampSignup
}

// GetTimestampSignupOk returns a tuple with the TimestampSignup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetTimestampSignupOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampSignup) {
		return nil, false
	}
	return o.TimestampSignup, true
}

// HasTimestampSignup returns a boolean if a field has been set.
func (o *AddListMembers1) HasTimestampSignup() bool {
	if o != nil && !IsNil(o.TimestampSignup) {
		return true
	}

	return false
}

// SetTimestampSignup gets a reference to the given time.Time and assigns it to the TimestampSignup field.
func (o *AddListMembers1) SetTimestampSignup(v time.Time) {
	o.TimestampSignup = &v
}

// GetIpOpt returns the IpOpt field value if set, zero value otherwise.
func (o *AddListMembers1) GetIpOpt() string {
	if o == nil || IsNil(o.IpOpt) {
		var ret string
		return ret
	}
	return *o.IpOpt
}

// GetIpOptOk returns a tuple with the IpOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetIpOptOk() (*string, bool) {
	if o == nil || IsNil(o.IpOpt) {
		return nil, false
	}
	return o.IpOpt, true
}

// HasIpOpt returns a boolean if a field has been set.
func (o *AddListMembers1) HasIpOpt() bool {
	if o != nil && !IsNil(o.IpOpt) {
		return true
	}

	return false
}

// SetIpOpt gets a reference to the given string and assigns it to the IpOpt field.
func (o *AddListMembers1) SetIpOpt(v string) {
	o.IpOpt = &v
}

// GetTimestampOpt returns the TimestampOpt field value if set, zero value otherwise.
func (o *AddListMembers1) GetTimestampOpt() time.Time {
	if o == nil || IsNil(o.TimestampOpt) {
		var ret time.Time
		return ret
	}
	return *o.TimestampOpt
}

// GetTimestampOptOk returns a tuple with the TimestampOpt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetTimestampOptOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimestampOpt) {
		return nil, false
	}
	return o.TimestampOpt, true
}

// HasTimestampOpt returns a boolean if a field has been set.
func (o *AddListMembers1) HasTimestampOpt() bool {
	if o != nil && !IsNil(o.TimestampOpt) {
		return true
	}

	return false
}

// SetTimestampOpt gets a reference to the given time.Time and assigns it to the TimestampOpt field.
func (o *AddListMembers1) SetTimestampOpt(v time.Time) {
	o.TimestampOpt = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AddListMembers1) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddListMembers1) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AddListMembers1) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AddListMembers1) SetTags(v []string) {
	o.Tags = v
}

func (o AddListMembers1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddListMembers1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_address"] = o.EmailAddress
	if !IsNil(o.EmailType) {
		toSerialize["email_type"] = o.EmailType
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.MergeFields) {
		toSerialize["merge_fields"] = o.MergeFields
	}
	if !IsNil(o.Interests) {
		toSerialize["interests"] = o.Interests
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Vip) {
		toSerialize["vip"] = o.Vip
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.MarketingPermissions) {
		toSerialize["marketing_permissions"] = o.MarketingPermissions
	}
	if !IsNil(o.IpSignup) {
		toSerialize["ip_signup"] = o.IpSignup
	}
	if !IsNil(o.TimestampSignup) {
		toSerialize["timestamp_signup"] = o.TimestampSignup
	}
	if !IsNil(o.IpOpt) {
		toSerialize["ip_opt"] = o.IpOpt
	}
	if !IsNil(o.TimestampOpt) {
		toSerialize["timestamp_opt"] = o.TimestampOpt
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAddListMembers1 struct {
	value *AddListMembers1
	isSet bool
}

func (v NullableAddListMembers1) Get() *AddListMembers1 {
	return v.value
}

func (v *NullableAddListMembers1) Set(val *AddListMembers1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddListMembers1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddListMembers1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddListMembers1(val *AddListMembers1) *NullableAddListMembers1 {
	return &NullableAddListMembers1{value: val, isSet: true}
}

func (v NullableAddListMembers1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddListMembers1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


