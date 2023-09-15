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

// checks if the CampaignFeedback2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignFeedback2{}

// CampaignFeedback2 A specific feedback message from a specific campaign.
type CampaignFeedback2 struct {
	// The individual id for the feedback item.
	FeedbackId *int32 `json:"feedback_id,omitempty"`
	// If a reply, the id of the parent feedback item.
	ParentId *int32 `json:"parent_id,omitempty"`
	// The block id for the editable block that the feedback addresses.
	BlockId *int32 `json:"block_id,omitempty"`
	// The content of the feedback.
	Message *string `json:"message,omitempty"`
	// The status of feedback.
	IsComplete *bool `json:"is_complete,omitempty"`
	// The login name of the user who created the feedback.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time the feedback item was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time the feedback was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The source of the feedback.
	Source *string `json:"source,omitempty"`
	// The unique id for the campaign.
	CampaignId *string `json:"campaign_id,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCampaignFeedback2 instantiates a new CampaignFeedback2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignFeedback2() *CampaignFeedback2 {
	this := CampaignFeedback2{}
	return &this
}

// NewCampaignFeedback2WithDefaults instantiates a new CampaignFeedback2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignFeedback2WithDefaults() *CampaignFeedback2 {
	this := CampaignFeedback2{}
	return &this
}

// GetFeedbackId returns the FeedbackId field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetFeedbackId() int32 {
	if o == nil || IsNil(o.FeedbackId) {
		var ret int32
		return ret
	}
	return *o.FeedbackId
}

// GetFeedbackIdOk returns a tuple with the FeedbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetFeedbackIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FeedbackId) {
		return nil, false
	}
	return o.FeedbackId, true
}

// HasFeedbackId returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasFeedbackId() bool {
	if o != nil && !IsNil(o.FeedbackId) {
		return true
	}

	return false
}

// SetFeedbackId gets a reference to the given int32 and assigns it to the FeedbackId field.
func (o *CampaignFeedback2) SetFeedbackId(v int32) {
	o.FeedbackId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetParentId() int32 {
	if o == nil || IsNil(o.ParentId) {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetParentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *CampaignFeedback2) SetParentId(v int32) {
	o.ParentId = &v
}

// GetBlockId returns the BlockId field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetBlockId() int32 {
	if o == nil || IsNil(o.BlockId) {
		var ret int32
		return ret
	}
	return *o.BlockId
}

// GetBlockIdOk returns a tuple with the BlockId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetBlockIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockId) {
		return nil, false
	}
	return o.BlockId, true
}

// HasBlockId returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasBlockId() bool {
	if o != nil && !IsNil(o.BlockId) {
		return true
	}

	return false
}

// SetBlockId gets a reference to the given int32 and assigns it to the BlockId field.
func (o *CampaignFeedback2) SetBlockId(v int32) {
	o.BlockId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CampaignFeedback2) SetMessage(v string) {
	o.Message = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetIsComplete() bool {
	if o == nil || IsNil(o.IsComplete) {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetIsCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsComplete) {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasIsComplete() bool {
	if o != nil && !IsNil(o.IsComplete) {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *CampaignFeedback2) SetIsComplete(v bool) {
	o.IsComplete = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *CampaignFeedback2) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CampaignFeedback2) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CampaignFeedback2) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CampaignFeedback2) SetSource(v string) {
	o.Source = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *CampaignFeedback2) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CampaignFeedback2) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignFeedback2) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CampaignFeedback2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CampaignFeedback2) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CampaignFeedback2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignFeedback2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeedbackId) {
		toSerialize["feedback_id"] = o.FeedbackId
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.BlockId) {
		toSerialize["block_id"] = o.BlockId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.IsComplete) {
		toSerialize["is_complete"] = o.IsComplete
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCampaignFeedback2 struct {
	value *CampaignFeedback2
	isSet bool
}

func (v NullableCampaignFeedback2) Get() *CampaignFeedback2 {
	return v.value
}

func (v *NullableCampaignFeedback2) Set(val *CampaignFeedback2) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignFeedback2) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignFeedback2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignFeedback2(val *CampaignFeedback2) *NullableCampaignFeedback2 {
	return &NullableCampaignFeedback2{value: val, isSet: true}
}

func (v NullableCampaignFeedback2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignFeedback2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

