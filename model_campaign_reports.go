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

// checks if the CampaignReports type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignReports{}

// CampaignReports A summary of the comment feedback for a specific campaign.
type CampaignReports struct {
	// A collection of feedback items for a campaign.
	Feedback []CampaignFeedback `json:"feedback,omitempty"`
	// The unique id for the campaign.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCampaignReports instantiates a new CampaignReports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignReports() *CampaignReports {
	this := CampaignReports{}
	return &this
}

// NewCampaignReportsWithDefaults instantiates a new CampaignReports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignReportsWithDefaults() *CampaignReports {
	this := CampaignReports{}
	return &this
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *CampaignReports) GetFeedback() []CampaignFeedback {
	if o == nil || IsNil(o.Feedback) {
		var ret []CampaignFeedback
		return ret
	}
	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReports) GetFeedbackOk() ([]CampaignFeedback, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *CampaignReports) HasFeedback() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given []CampaignFeedback and assigns it to the Feedback field.
func (o *CampaignReports) SetFeedback(v []CampaignFeedback) {
	o.Feedback = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignReports) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReports) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignReports) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *CampaignReports) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CampaignReports) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReports) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CampaignReports) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CampaignReports) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CampaignReports) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignReports) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CampaignReports) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CampaignReports) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CampaignReports) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignReports) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCampaignReports struct {
	value *CampaignReports
	isSet bool
}

func (v NullableCampaignReports) Get() *CampaignReports {
	return v.value
}

func (v *NullableCampaignReports) Set(val *CampaignReports) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignReports) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignReports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignReports(val *CampaignReports) *NullableCampaignReports {
	return &NullableCampaignReports{value: val, isSet: true}
}

func (v NullableCampaignReports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignReports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


