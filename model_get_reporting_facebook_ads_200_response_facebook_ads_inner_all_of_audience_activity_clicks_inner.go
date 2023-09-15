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

// checks if the GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner{}

// GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner struct for GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner
type GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner struct {
	Date *string `json:"date,omitempty"`
	Clicks *int32 `json:"clicks,omitempty"`
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner{}
	return &this
}

// NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInnerWithDefaults instantiates a new GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInnerWithDefaults() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner {
	this := GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) SetDate(v string) {
	o.Date = &v
}

// GetClicks returns the Clicks field value if set, zero value otherwise.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) GetClicks() int32 {
	if o == nil || IsNil(o.Clicks) {
		var ret int32
		return ret
	}
	return *o.Clicks
}

// GetClicksOk returns a tuple with the Clicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) GetClicksOk() (*int32, bool) {
	if o == nil || IsNil(o.Clicks) {
		return nil, false
	}
	return o.Clicks, true
}

// HasClicks returns a boolean if a field has been set.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) HasClicks() bool {
	if o != nil && !IsNil(o.Clicks) {
		return true
	}

	return false
}

// SetClicks gets a reference to the given int32 and assigns it to the Clicks field.
func (o *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) SetClicks(v int32) {
	o.Clicks = &v
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Clicks) {
		toSerialize["clicks"] = o.Clicks
	}
	return toSerialize, nil
}

type NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner struct {
	value *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner
	isSet bool
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) Get() *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner {
	return v.value
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) Set(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner(val *GetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner {
	return &NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner{value: val, isSet: true}
}

func (v NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReportingFacebookAds200ResponseFacebookAdsInnerAllOfAudienceActivityClicksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

