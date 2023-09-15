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

// checks if the CampaignDeliveryStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignDeliveryStatus{}

// CampaignDeliveryStatus Updates on campaigns in the process of sending.
type CampaignDeliveryStatus struct {
	// Whether Campaign Delivery Status is enabled for this account and campaign.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether a campaign send can be canceled.
	CanCancel *bool `json:"can_cancel,omitempty"`
	// The current state of a campaign delivery.
	Status *string `json:"status,omitempty"`
	// The total number of emails confirmed sent for this campaign so far.
	EmailsSent *int32 `json:"emails_sent,omitempty"`
	// The total number of emails canceled for this campaign.
	EmailsCanceled *int32 `json:"emails_canceled,omitempty"`
}

// NewCampaignDeliveryStatus instantiates a new CampaignDeliveryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDeliveryStatus() *CampaignDeliveryStatus {
	this := CampaignDeliveryStatus{}
	return &this
}

// NewCampaignDeliveryStatusWithDefaults instantiates a new CampaignDeliveryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDeliveryStatusWithDefaults() *CampaignDeliveryStatus {
	this := CampaignDeliveryStatus{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CampaignDeliveryStatus) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeliveryStatus) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CampaignDeliveryStatus) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CampaignDeliveryStatus) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCanCancel returns the CanCancel field value if set, zero value otherwise.
func (o *CampaignDeliveryStatus) GetCanCancel() bool {
	if o == nil || IsNil(o.CanCancel) {
		var ret bool
		return ret
	}
	return *o.CanCancel
}

// GetCanCancelOk returns a tuple with the CanCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeliveryStatus) GetCanCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.CanCancel) {
		return nil, false
	}
	return o.CanCancel, true
}

// HasCanCancel returns a boolean if a field has been set.
func (o *CampaignDeliveryStatus) HasCanCancel() bool {
	if o != nil && !IsNil(o.CanCancel) {
		return true
	}

	return false
}

// SetCanCancel gets a reference to the given bool and assigns it to the CanCancel field.
func (o *CampaignDeliveryStatus) SetCanCancel(v bool) {
	o.CanCancel = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CampaignDeliveryStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeliveryStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CampaignDeliveryStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CampaignDeliveryStatus) SetStatus(v string) {
	o.Status = &v
}

// GetEmailsSent returns the EmailsSent field value if set, zero value otherwise.
func (o *CampaignDeliveryStatus) GetEmailsSent() int32 {
	if o == nil || IsNil(o.EmailsSent) {
		var ret int32
		return ret
	}
	return *o.EmailsSent
}

// GetEmailsSentOk returns a tuple with the EmailsSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeliveryStatus) GetEmailsSentOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsSent) {
		return nil, false
	}
	return o.EmailsSent, true
}

// HasEmailsSent returns a boolean if a field has been set.
func (o *CampaignDeliveryStatus) HasEmailsSent() bool {
	if o != nil && !IsNil(o.EmailsSent) {
		return true
	}

	return false
}

// SetEmailsSent gets a reference to the given int32 and assigns it to the EmailsSent field.
func (o *CampaignDeliveryStatus) SetEmailsSent(v int32) {
	o.EmailsSent = &v
}

// GetEmailsCanceled returns the EmailsCanceled field value if set, zero value otherwise.
func (o *CampaignDeliveryStatus) GetEmailsCanceled() int32 {
	if o == nil || IsNil(o.EmailsCanceled) {
		var ret int32
		return ret
	}
	return *o.EmailsCanceled
}

// GetEmailsCanceledOk returns a tuple with the EmailsCanceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignDeliveryStatus) GetEmailsCanceledOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailsCanceled) {
		return nil, false
	}
	return o.EmailsCanceled, true
}

// HasEmailsCanceled returns a boolean if a field has been set.
func (o *CampaignDeliveryStatus) HasEmailsCanceled() bool {
	if o != nil && !IsNil(o.EmailsCanceled) {
		return true
	}

	return false
}

// SetEmailsCanceled gets a reference to the given int32 and assigns it to the EmailsCanceled field.
func (o *CampaignDeliveryStatus) SetEmailsCanceled(v int32) {
	o.EmailsCanceled = &v
}

func (o CampaignDeliveryStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignDeliveryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.CanCancel) {
		toSerialize["can_cancel"] = o.CanCancel
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.EmailsSent) {
		toSerialize["emails_sent"] = o.EmailsSent
	}
	if !IsNil(o.EmailsCanceled) {
		toSerialize["emails_canceled"] = o.EmailsCanceled
	}
	return toSerialize, nil
}

type NullableCampaignDeliveryStatus struct {
	value *CampaignDeliveryStatus
	isSet bool
}

func (v NullableCampaignDeliveryStatus) Get() *CampaignDeliveryStatus {
	return v.value
}

func (v *NullableCampaignDeliveryStatus) Set(val *CampaignDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDeliveryStatus(val *CampaignDeliveryStatus) *NullableCampaignDeliveryStatus {
	return &NullableCampaignDeliveryStatus{value: val, isSet: true}
}

func (v NullableCampaignDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


