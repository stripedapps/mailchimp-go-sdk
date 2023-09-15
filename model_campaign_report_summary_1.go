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

// checks if the CampaignReportSummary1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignReportSummary1{}

// CampaignReportSummary1 For sent campaigns, a summary of opens and clicks.
type CampaignReportSummary1 struct {
	// The total number of opens for a campaign.
	Opens *int32 `json:"opens,omitempty"`
	// The number of unique opens.
	UniqueOpens *int32 `json:"unique_opens,omitempty"`
	// The number of unique opens divided by the total number of successful deliveries.
	OpenRate *float32 `json:"open_rate,omitempty"`
	// The total number of clicks for an campaign.
	Clicks *int32 `json:"clicks,omitempty"`
	// The number of unique clicks.
	SubscriberClicks *int32 `json:"subscriber_clicks,omitempty"`
	// The number of unique clicks divided by the total number of successful deliveries.
	ClickRate *float32 `json:"click_rate,omitempty"`
}

// NewCampaignReportSummary1 instantiates a new CampaignReportSummary1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReportSummary1() *CampaignReportSummary1 {
	this := CampaignReportSummary1{}
	return &this
}

// NewCampaignReportSummary1WithDefaults instantiates a new CampaignReportSummary1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReportSummary1WithDefaults() *CampaignReportSummary1 {
	this := CampaignReportSummary1{}
	return &this
}

// GetOpens returns the Opens field value if set, zero value otherwise.
func (o *CampaignReportSummary1) GetOpens() int32 {
	if o == nil || IsNil(o.Opens) {
		var ret int32
		return ret
	}
	return *o.Opens
}

// GetOpensOk returns a tuple with the Opens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary1) GetOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.Opens) {
		return nil, false
	}
	return o.Opens, true
}

// HasOpens returns a boolean if a field has been set.
func (o *CampaignReportSummary1) HasOpens() bool {
	if o != nil && !IsNil(o.Opens) {
		return true
	}

	return false
}

// SetOpens gets a reference to the given int32 and assigns it to the Opens field.
func (o *CampaignReportSummary1) SetOpens(v int32) {
	o.Opens = &v
}

// GetUniqueOpens returns the UniqueOpens field value if set, zero value otherwise.
func (o *CampaignReportSummary1) GetUniqueOpens() int32 {
	if o == nil || IsNil(o.UniqueOpens) {
		var ret int32
		return ret
	}
	return *o.UniqueOpens
}

// GetUniqueOpensOk returns a tuple with the UniqueOpens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary1) GetUniqueOpensOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueOpens) {
		return nil, false
	}
	return o.UniqueOpens, true
}

// HasUniqueOpens returns a boolean if a field has been set.
func (o *CampaignReportSummary1) HasUniqueOpens() bool {
	if o != nil && !IsNil(o.UniqueOpens) {
		return true
	}

	return false
}

// SetUniqueOpens gets a reference to the given int32 and assigns it to the UniqueOpens field.
func (o *CampaignReportSummary1) SetUniqueOpens(v int32) {
	o.UniqueOpens = &v
}

// GetOpenRate returns the OpenRate field value if set, zero value otherwise.
func (o *CampaignReportSummary1) GetOpenRate() float32 {
	if o == nil || IsNil(o.OpenRate) {
		var ret float32
		return ret
	}
	return *o.OpenRate
}

// GetOpenRateOk returns a tuple with the OpenRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary1) GetOpenRateOk() (*float32, bool) {
	if o == nil || IsNil(o.OpenRate) {
		return nil, false
	}
	return o.OpenRate, true
}

// HasOpenRate returns a boolean if a field has been set.
func (o *CampaignReportSummary1) HasOpenRate() bool {
	if o != nil && !IsNil(o.OpenRate) {
		return true
	}

	return false
}

// SetOpenRate gets a reference to the given float32 and assigns it to the OpenRate field.
func (o *CampaignReportSummary1) SetOpenRate(v float32) {
	o.OpenRate = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *CampaignReportSummary1) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary1) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *CampaignReportSummary1) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *CampaignReportSummary1) SetClicks(v int32) {
	o.Clicks = &v
}

// GetSubscriberClicks returns the SubscriberClicks field value if set, zero value otherwise.
func (o *CampaignReportSummary1) GetSubscriberClicks() int32 {
	if o == nil || IsNil(o.SubscriberClicks) {
		var ret int32
		return ret
	}
	return *o.SubscriberClicks
}

// GetSubscriberClicksOk returns a tuple with the SubscriberClicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary1) GetSubscriberClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriberClicks) {
		return nil, false
	}
	return o.SubscriberClicks, true
}

// HasSubscriberClicks returns a boolean if a field has been set.
func (o *CampaignReportSummary1) HasSubscriberClicks() bool {
	if o != nil && !IsNil(o.SubscriberClicks) {
		return true
	}

	return false
}

// SetSubscriberClicks gets a reference to the given int32 and assigns it to the SubscriberClicks field.
func (o *CampaignReportSummary1) SetSubscriberClicks(v int32) {
	o.SubscriberClicks = &v
}

// GetClickRate returns the ClickRate field value if set, zero value otherwise.
func (o *CampaignReportSummary1) GetClickRate() float32 {
	if o == nil || IsNil(o.ClickRate) {
		var ret float32
		return ret
	}
	return *o.ClickRate
}

// GetClickRateOk returns a tuple with the ClickRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReportSummary1) GetClickRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ClickRate) {
		return nil, false
	}
	return o.ClickRate, true
}

// HasClickRate returns a boolean if a field has been set.
func (o *CampaignReportSummary1) HasClickRate() bool {
	if o != nil && !IsNil(o.ClickRate) {
		return true
	}

	return false
}

// SetClickRate gets a reference to the given float32 and assigns it to the ClickRate field.
func (o *CampaignReportSummary1) SetClickRate(v float32) {
	o.ClickRate = &v
}

func (o CampaignReportSummary1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignReportSummary1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Opens) {
		toSerialize["opens"] = o.Opens
	}
	if !IsNil(o.UniqueOpens) {
		toSerialize["unique_opens"] = o.UniqueOpens
	}
	if !IsNil(o.OpenRate) {
		toSerialize["open_rate"] = o.OpenRate
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	if !IsNil(o.SubscriberClicks) {
		toSerialize["subscriber_clicks"] = o.SubscriberClicks
	}
	if !IsNil(o.ClickRate) {
		toSerialize["click_rate"] = o.ClickRate
	}
	return toSerialize, nil
}

type NullableCampaignReportSummary1 struct {
	value *CampaignReportSummary1
	isSet bool
}

func (v NullableCampaignReportSummary1) Get() *CampaignReportSummary1 {
	return v.value
}

func (v *NullableCampaignReportSummary1) Set(val *CampaignReportSummary1) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReportSummary1) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReportSummary1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReportSummary1(val *CampaignReportSummary1) *NullableCampaignReportSummary1 {
	return &NullableCampaignReportSummary1{value: val, isSet: true}
}

func (v NullableCampaignReportSummary1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReportSummary1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


