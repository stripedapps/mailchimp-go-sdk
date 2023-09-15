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

// checks if the SurveyReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SurveyReport{}

// SurveyReport The report for a survey.
type SurveyReport struct {
	// A string that uniquely identifies this survey.
	Id *string `json:"id,omitempty"`
	// The ID used in the Mailchimp web application. View this survey report in your Mailchimp account at `https://{dc}.admin.mailchimp.com/lists/surveys/results?survey_id={web_id}`.
	WebId *int32 `json:"web_id,omitempty"`
	// The ID of the list connected to this survey.
	ListId *string `json:"list_id,omitempty"`
	// The name of the list connected to this survey.
	ListName *string `json:"list_name,omitempty"`
	// The title of the survey.
	Title *string `json:"title,omitempty"`
	// The URL for the survey.
	Url *string `json:"url,omitempty"`
	// The survey's status.
	Status *string `json:"status,omitempty"`
	// The date and time the survey was published in ISO 8601 format.
	PublishedAt *time.Time `json:"published_at,omitempty"`
	// The date and time the survey was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time the survey was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The total number of responses to this survey.
	TotalResponses *int32 `json:"total_responses,omitempty"`
}

// NewSurveyReport instantiates a new SurveyReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSurveyReport() *SurveyReport {
	this := SurveyReport{}
	return &this
}

// NewSurveyReportWithDefaults instantiates a new SurveyReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSurveyReportWithDefaults() *SurveyReport {
	this := SurveyReport{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SurveyReport) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SurveyReport) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SurveyReport) SetId(v string) {
	o.Id = &v
}

// GetWebId returns the WebId field value if set, zero value otherwise.
func (o *SurveyReport) GetWebId() int32 {
	if o == nil || IsNil(o.WebId) {
		var ret int32
		return ret
	}
	return *o.WebId
}

// GetWebIdOk returns a tuple with the WebId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetWebIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebId) {
		return nil, false
	}
	return o.WebId, true
}

// HasWebId returns a boolean if a field has been set.
func (o *SurveyReport) HasWebId() bool {
	if o != nil && !IsNil(o.WebId) {
		return true
	}

	return false
}

// SetWebId gets a reference to the given int32 and assigns it to the WebId field.
func (o *SurveyReport) SetWebId(v int32) {
	o.WebId = &v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *SurveyReport) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *SurveyReport) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *SurveyReport) SetListId(v string) {
	o.ListId = &v
}

// GetListName returns the ListName field value if set, zero value otherwise.
func (o *SurveyReport) GetListName() string {
	if o == nil || IsNil(o.ListName) {
		var ret string
		return ret
	}
	return *o.ListName
}

// GetListNameOk returns a tuple with the ListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetListNameOk() (*string, bool) {
	if o == nil || IsNil(o.ListName) {
		return nil, false
	}
	return o.ListName, true
}

// HasListName returns a boolean if a field has been set.
func (o *SurveyReport) HasListName() bool {
	if o != nil && !IsNil(o.ListName) {
		return true
	}

	return false
}

// SetListName gets a reference to the given string and assigns it to the ListName field.
func (o *SurveyReport) SetListName(v string) {
	o.ListName = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SurveyReport) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SurveyReport) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SurveyReport) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SurveyReport) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SurveyReport) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SurveyReport) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SurveyReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SurveyReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SurveyReport) SetStatus(v string) {
	o.Status = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *SurveyReport) GetPublishedAt() time.Time {
	if o == nil || IsNil(o.PublishedAt) {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedAt) {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *SurveyReport) HasPublishedAt() bool {
	if o != nil && !IsNil(o.PublishedAt) {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *SurveyReport) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SurveyReport) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SurveyReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SurveyReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SurveyReport) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SurveyReport) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SurveyReport) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetTotalResponses returns the TotalResponses field value if set, zero value otherwise.
func (o *SurveyReport) GetTotalResponses() int32 {
	if o == nil || IsNil(o.TotalResponses) {
		var ret int32
		return ret
	}
	return *o.TotalResponses
}

// GetTotalResponsesOk returns a tuple with the TotalResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SurveyReport) GetTotalResponsesOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResponses) {
		return nil, false
	}
	return o.TotalResponses, true
}

// HasTotalResponses returns a boolean if a field has been set.
func (o *SurveyReport) HasTotalResponses() bool {
	if o != nil && !IsNil(o.TotalResponses) {
		return true
	}

	return false
}

// SetTotalResponses gets a reference to the given int32 and assigns it to the TotalResponses field.
func (o *SurveyReport) SetTotalResponses(v int32) {
	o.TotalResponses = &v
}

func (o SurveyReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SurveyReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.WebId) {
		toSerialize["web_id"] = o.WebId
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.ListName) {
		toSerialize["list_name"] = o.ListName
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.PublishedAt) {
		toSerialize["published_at"] = o.PublishedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.TotalResponses) {
		toSerialize["total_responses"] = o.TotalResponses
	}
	return toSerialize, nil
}

type NullableSurveyReport struct {
	value *SurveyReport
	isSet bool
}

func (v NullableSurveyReport) Get() *SurveyReport {
	return v.value
}

func (v *NullableSurveyReport) Set(val *SurveyReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSurveyReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSurveyReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSurveyReport(val *SurveyReport) *NullableSurveyReport {
	return &NullableSurveyReport{value: val, isSet: true}
}

func (v NullableSurveyReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSurveyReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


