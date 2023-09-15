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

// checks if the RSSOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RSSOptions{}

// RSSOptions [RSS](https://mailchimp.com/help/share-your-blog-posts-with-mailchimp/) options for a campaign.
type RSSOptions struct {
	// The URL for the RSS feed.
	FeedUrl *string `json:"feed_url,omitempty"`
	// The frequency of the RSS Campaign.
	Frequency *string `json:"frequency,omitempty"`
	Schedule *SendingSchedule `json:"schedule,omitempty"`
	// The date the campaign was last sent.
	LastSent *time.Time `json:"last_sent,omitempty"`
	// Whether to add CSS to images in the RSS feed to constrain their width in campaigns.
	ConstrainRssImg *bool `json:"constrain_rss_img,omitempty"`
}

// NewRSSOptions instantiates a new RSSOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRSSOptions() *RSSOptions {
	this := RSSOptions{}
	return &this
}

// NewRSSOptionsWithDefaults instantiates a new RSSOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRSSOptionsWithDefaults() *RSSOptions {
	this := RSSOptions{}
	return &this
}

// GetFeedUrl returns the FeedUrl field value if set, zero value otherwise.
func (o *RSSOptions) GetFeedUrl() string {
	if o == nil || IsNil(o.FeedUrl) {
		var ret string
		return ret
	}
	return *o.FeedUrl
}

// GetFeedUrlOk returns a tuple with the FeedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions) GetFeedUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FeedUrl) {
		return nil, false
	}
	return o.FeedUrl, true
}

// HasFeedUrl returns a boolean if a field has been set.
func (o *RSSOptions) HasFeedUrl() bool {
	if o != nil && !IsNil(o.FeedUrl) {
		return true
	}

	return false
}

// SetFeedUrl gets a reference to the given string and assigns it to the FeedUrl field.
func (o *RSSOptions) SetFeedUrl(v string) {
	o.FeedUrl = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *RSSOptions) GetFrequency() string {
	if o == nil || IsNil(o.Frequency) {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions) GetFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.Frequency) {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *RSSOptions) HasFrequency() bool {
	if o != nil && !IsNil(o.Frequency) {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *RSSOptions) SetFrequency(v string) {
	o.Frequency = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *RSSOptions) GetSchedule() SendingSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret SendingSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions) GetScheduleOk() (*SendingSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *RSSOptions) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given SendingSchedule and assigns it to the Schedule field.
func (o *RSSOptions) SetSchedule(v SendingSchedule) {
	o.Schedule = &v
}

// GetLastSent returns the LastSent field value if set, zero value otherwise.
func (o *RSSOptions) GetLastSent() time.Time {
	if o == nil || IsNil(o.LastSent) {
		var ret time.Time
		return ret
	}
	return *o.LastSent
}

// GetLastSentOk returns a tuple with the LastSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions) GetLastSentOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastSent) {
		return nil, false
	}
	return o.LastSent, true
}

// HasLastSent returns a boolean if a field has been set.
func (o *RSSOptions) HasLastSent() bool {
	if o != nil && !IsNil(o.LastSent) {
		return true
	}

	return false
}

// SetLastSent gets a reference to the given time.Time and assigns it to the LastSent field.
func (o *RSSOptions) SetLastSent(v time.Time) {
	o.LastSent = &v
}

// GetConstrainRssImg returns the ConstrainRssImg field value if set, zero value otherwise.
func (o *RSSOptions) GetConstrainRssImg() bool {
	if o == nil || IsNil(o.ConstrainRssImg) {
		var ret bool
		return ret
	}
	return *o.ConstrainRssImg
}

// GetConstrainRssImgOk returns a tuple with the ConstrainRssImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions) GetConstrainRssImgOk() (*bool, bool) {
	if o == nil || IsNil(o.ConstrainRssImg) {
		return nil, false
	}
	return o.ConstrainRssImg, true
}

// HasConstrainRssImg returns a boolean if a field has been set.
func (o *RSSOptions) HasConstrainRssImg() bool {
	if o != nil && !IsNil(o.ConstrainRssImg) {
		return true
	}

	return false
}

// SetConstrainRssImg gets a reference to the given bool and assigns it to the ConstrainRssImg field.
func (o *RSSOptions) SetConstrainRssImg(v bool) {
	o.ConstrainRssImg = &v
}

func (o RSSOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RSSOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeedUrl) {
		toSerialize["feed_url"] = o.FeedUrl
	}
	if !IsNil(o.Frequency) {
		toSerialize["frequency"] = o.Frequency
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.LastSent) {
		toSerialize["last_sent"] = o.LastSent
	}
	if !IsNil(o.ConstrainRssImg) {
		toSerialize["constrain_rss_img"] = o.ConstrainRssImg
	}
	return toSerialize, nil
}

type NullableRSSOptions struct {
	value *RSSOptions
	isSet bool
}

func (v NullableRSSOptions) Get() *RSSOptions {
	return v.value
}

func (v *NullableRSSOptions) Set(val *RSSOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRSSOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRSSOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRSSOptions(val *RSSOptions) *NullableRSSOptions {
	return &NullableRSSOptions{value: val, isSet: true}
}

func (v NullableRSSOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRSSOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

