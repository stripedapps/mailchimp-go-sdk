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

// checks if the RSSOptions1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RSSOptions1{}

// RSSOptions1 [RSS](https://mailchimp.com/help/share-your-blog-posts-with-mailchimp/) options, specific to an RSS campaign.
type RSSOptions1 struct {
	// The URL for the RSS feed.
	FeedUrl string `json:"feed_url"`
	// The frequency of the RSS Campaign.
	Frequency string `json:"frequency"`
	Schedule *SendingSchedule `json:"schedule,omitempty"`
	// Whether to add CSS to images in the RSS feed to constrain their width in campaigns.
	ConstrainRssImg *bool `json:"constrain_rss_img,omitempty"`
}

// NewRSSOptions1 instantiates a new RSSOptions1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRSSOptions1(feedUrl string, frequency string) *RSSOptions1 {
	this := RSSOptions1{}
	this.FeedUrl = feedUrl
	this.Frequency = frequency
	return &this
}

// NewRSSOptions1WithDefaults instantiates a new RSSOptions1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRSSOptions1WithDefaults() *RSSOptions1 {
	this := RSSOptions1{}
	return &this
}

// GetFeedUrl returns the FeedUrl field value
func (o *RSSOptions1) GetFeedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeedUrl
}

// GetFeedUrlOk returns a tuple with the FeedUrl field value
// and a boolean to check if the value has been set.
func (o *RSSOptions1) GetFeedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeedUrl, true
}

// SetFeedUrl sets field value
func (o *RSSOptions1) SetFeedUrl(v string) {
	o.FeedUrl = v
}

// GetFrequency returns the Frequency field value
func (o *RSSOptions1) GetFrequency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *RSSOptions1) GetFrequencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *RSSOptions1) SetFrequency(v string) {
	o.Frequency = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *RSSOptions1) GetSchedule() SendingSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret SendingSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions1) GetScheduleOk() (*SendingSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *RSSOptions1) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given SendingSchedule and assigns it to the Schedule field.
func (o *RSSOptions1) SetSchedule(v SendingSchedule) {
	o.Schedule = &v
}

// GetConstrainRssImg returns the ConstrainRssImg field value if set, zero value otherwise.
func (o *RSSOptions1) GetConstrainRssImg() bool {
	if o == nil || IsNil(o.ConstrainRssImg) {
		var ret bool
		return ret
	}
	return *o.ConstrainRssImg
}

// GetConstrainRssImgOk returns a tuple with the ConstrainRssImg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RSSOptions1) GetConstrainRssImgOk() (*bool, bool) {
	if o == nil || IsNil(o.ConstrainRssImg) {
		return nil, false
	}
	return o.ConstrainRssImg, true
}

// HasConstrainRssImg returns a boolean if a field has been set.
func (o *RSSOptions1) HasConstrainRssImg() bool {
	if o != nil && !IsNil(o.ConstrainRssImg) {
		return true
	}

	return false
}

// SetConstrainRssImg gets a reference to the given bool and assigns it to the ConstrainRssImg field.
func (o *RSSOptions1) SetConstrainRssImg(v bool) {
	o.ConstrainRssImg = &v
}

func (o RSSOptions1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RSSOptions1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feed_url"] = o.FeedUrl
	toSerialize["frequency"] = o.Frequency
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.ConstrainRssImg) {
		toSerialize["constrain_rss_img"] = o.ConstrainRssImg
	}
	return toSerialize, nil
}

type NullableRSSOptions1 struct {
	value *RSSOptions1
	isSet bool
}

func (v NullableRSSOptions1) Get() *RSSOptions1 {
	return v.value
}

func (v *NullableRSSOptions1) Set(val *RSSOptions1) {
	v.value = val
	v.isSet = true
}

func (v NullableRSSOptions1) IsSet() bool {
	return v.isSet
}

func (v *NullableRSSOptions1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRSSOptions1(val *RSSOptions1) *NullableRSSOptions1 {
	return &NullableRSSOptions1{value: val, isSet: true}
}

func (v NullableRSSOptions1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRSSOptions1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


