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

// checks if the GetAllFacebookAds200ResponseFacebookAdsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllFacebookAds200ResponseFacebookAdsInner{}

// GetAllFacebookAds200ResponseFacebookAdsInner struct for GetAllFacebookAds200ResponseFacebookAdsInner
type GetAllFacebookAds200ResponseFacebookAdsInner struct {
	// Unique ID of an Outreach.
	Id *string `json:"id,omitempty"`
	// The ID used in the Mailchimp web application. For example, for a `regular` outreach, you can view this campaign in your Mailchimp account at `https://{dc}.admin.mailchimp.com/campaigns/show/?id={web_id}`.
	WebId *int32 `json:"web_id,omitempty"`
	// Title or name of an Outreach.
	Name *string `json:"name,omitempty"`
	// The type of outreach this object is.
	Type *string `json:"type,omitempty"`
	// The status of this outreach.
	Status *string `json:"status,omitempty"`
	// Outreach report availability. Note: This property is hotly debated in what it _should_ convey. See [MCP-1371](https://jira.mailchimp.com/browse/MCP-1371) for more context.
	ShowReport *bool `json:"show_report,omitempty"`
	// The date and time the outreach was created in ISO 8601 format.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The date and time the outreach was started in ISO 8601 format.
	StartTime *time.Time `json:"start_time,omitempty"`
	// The date and time the outreach was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The date and time the outreach was canceled in ISO 8601 format.
	CanceledAt *time.Time `json:"canceled_at,omitempty"`
	// The date and time the outreach was (or will be) published in ISO 8601 format.
	PublishedTime *time.Time `json:"published_time,omitempty"`
	// If this outreach targets a segment of your audience.
	HasSegment *bool `json:"has_segment,omitempty"`
	ReportSummary *ReportSummary `json:"report_summary,omitempty"`
	Recipients *Recipients `json:"recipients,omitempty"`
	// The URL of the thumbnail for this outreach.
	Thumbnail *string `json:"thumbnail,omitempty"`
	EmailSourceName *string `json:"email_source_name,omitempty"`
	// The date and time the ad was paused in ISO 8601 format.
	PausedAt *time.Time `json:"paused_at,omitempty"`
	// The date and time the ad was ended in ISO 8601 format.
	EndTime *time.Time `json:"end_time,omitempty"`
	// If the ad has a problem and needs attention.
	NeedsAttention *bool `json:"needs_attention,omitempty"`
	WasCanceledByFacebook *bool `json:"was_canceled_by_facebook,omitempty"`
	// Check if this ad is connected to a facebook page
	IsConnected *bool `json:"is_connected,omitempty"`
	// Check if this ad has audience setup
	HasAudience *bool `json:"has_audience,omitempty"`
	// Check if this ad has content
	HasContent *bool `json:"has_content,omitempty"`
	Channel *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel `json:"channel,omitempty"`
	Feedback *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback `json:"feedback,omitempty"`
	Site *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite `json:"site,omitempty"`
	Audience *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience `json:"audience,omitempty"`
	Budget *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget `json:"budget,omitempty"`
	Content *GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent `json:"content,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewGetAllFacebookAds200ResponseFacebookAdsInner instantiates a new GetAllFacebookAds200ResponseFacebookAdsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllFacebookAds200ResponseFacebookAdsInner() *GetAllFacebookAds200ResponseFacebookAdsInner {
	this := GetAllFacebookAds200ResponseFacebookAdsInner{}
	return &this
}

// NewGetAllFacebookAds200ResponseFacebookAdsInnerWithDefaults instantiates a new GetAllFacebookAds200ResponseFacebookAdsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllFacebookAds200ResponseFacebookAdsInnerWithDefaults() *GetAllFacebookAds200ResponseFacebookAdsInner {
	this := GetAllFacebookAds200ResponseFacebookAdsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetId(v string) {
	o.Id = &v
}

// GetWebId returns the WebId field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetWebId() int32 {
	if o == nil || IsNil(o.WebId) {
		var ret int32
		return ret
	}
	return *o.WebId
}

// GetWebIdOk returns a tuple with the WebId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetWebIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebId) {
		return nil, false
	}
	return o.WebId, true
}

// HasWebId returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasWebId() bool {
	if o != nil && !IsNil(o.WebId) {
		return true
	}

	return false
}

// SetWebId gets a reference to the given int32 and assigns it to the WebId field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetWebId(v int32) {
	o.WebId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetStatus(v string) {
	o.Status = &v
}

// GetShowReport returns the ShowReport field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetShowReport() bool {
	if o == nil || IsNil(o.ShowReport) {
		var ret bool
		return ret
	}
	return *o.ShowReport
}

// GetShowReportOk returns a tuple with the ShowReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetShowReportOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowReport) {
		return nil, false
	}
	return o.ShowReport, true
}

// HasShowReport returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasShowReport() bool {
	if o != nil && !IsNil(o.ShowReport) {
		return true
	}

	return false
}

// SetShowReport gets a reference to the given bool and assigns it to the ShowReport field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetShowReport(v bool) {
	o.ShowReport = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCanceledAt returns the CanceledAt field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetCanceledAt() time.Time {
	if o == nil || IsNil(o.CanceledAt) {
		var ret time.Time
		return ret
	}
	return *o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetCanceledAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CanceledAt) {
		return nil, false
	}
	return o.CanceledAt, true
}

// HasCanceledAt returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasCanceledAt() bool {
	if o != nil && !IsNil(o.CanceledAt) {
		return true
	}

	return false
}

// SetCanceledAt gets a reference to the given time.Time and assigns it to the CanceledAt field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetCanceledAt(v time.Time) {
	o.CanceledAt = &v
}

// GetPublishedTime returns the PublishedTime field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetPublishedTime() time.Time {
	if o == nil || IsNil(o.PublishedTime) {
		var ret time.Time
		return ret
	}
	return *o.PublishedTime
}

// GetPublishedTimeOk returns a tuple with the PublishedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetPublishedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishedTime) {
		return nil, false
	}
	return o.PublishedTime, true
}

// HasPublishedTime returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasPublishedTime() bool {
	if o != nil && !IsNil(o.PublishedTime) {
		return true
	}

	return false
}

// SetPublishedTime gets a reference to the given time.Time and assigns it to the PublishedTime field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetPublishedTime(v time.Time) {
	o.PublishedTime = &v
}

// GetHasSegment returns the HasSegment field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetHasSegment() bool {
	if o == nil || IsNil(o.HasSegment) {
		var ret bool
		return ret
	}
	return *o.HasSegment
}

// GetHasSegmentOk returns a tuple with the HasSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetHasSegmentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasSegment) {
		return nil, false
	}
	return o.HasSegment, true
}

// HasHasSegment returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasHasSegment() bool {
	if o != nil && !IsNil(o.HasSegment) {
		return true
	}

	return false
}

// SetHasSegment gets a reference to the given bool and assigns it to the HasSegment field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetHasSegment(v bool) {
	o.HasSegment = &v
}

// GetReportSummary returns the ReportSummary field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetReportSummary() ReportSummary {
	if o == nil || IsNil(o.ReportSummary) {
		var ret ReportSummary
		return ret
	}
	return *o.ReportSummary
}

// GetReportSummaryOk returns a tuple with the ReportSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetReportSummaryOk() (*ReportSummary, bool) {
	if o == nil || IsNil(o.ReportSummary) {
		return nil, false
	}
	return o.ReportSummary, true
}

// HasReportSummary returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasReportSummary() bool {
	if o != nil && !IsNil(o.ReportSummary) {
		return true
	}

	return false
}

// SetReportSummary gets a reference to the given ReportSummary and assigns it to the ReportSummary field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetReportSummary(v ReportSummary) {
	o.ReportSummary = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetRecipients() Recipients {
	if o == nil || IsNil(o.Recipients) {
		var ret Recipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetRecipientsOk() (*Recipients, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given Recipients and assigns it to the Recipients field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetRecipients(v Recipients) {
	o.Recipients = &v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetThumbnail() string {
	if o == nil || IsNil(o.Thumbnail) {
		var ret string
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetThumbnailOk() (*string, bool) {
	if o == nil || IsNil(o.Thumbnail) {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasThumbnail() bool {
	if o != nil && !IsNil(o.Thumbnail) {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given string and assigns it to the Thumbnail field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetThumbnail(v string) {
	o.Thumbnail = &v
}

// GetEmailSourceName returns the EmailSourceName field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetEmailSourceName() string {
	if o == nil || IsNil(o.EmailSourceName) {
		var ret string
		return ret
	}
	return *o.EmailSourceName
}

// GetEmailSourceNameOk returns a tuple with the EmailSourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetEmailSourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSourceName) {
		return nil, false
	}
	return o.EmailSourceName, true
}

// HasEmailSourceName returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasEmailSourceName() bool {
	if o != nil && !IsNil(o.EmailSourceName) {
		return true
	}

	return false
}

// SetEmailSourceName gets a reference to the given string and assigns it to the EmailSourceName field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetEmailSourceName(v string) {
	o.EmailSourceName = &v
}

// GetPausedAt returns the PausedAt field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetPausedAt() time.Time {
	if o == nil || IsNil(o.PausedAt) {
		var ret time.Time
		return ret
	}
	return *o.PausedAt
}

// GetPausedAtOk returns a tuple with the PausedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetPausedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PausedAt) {
		return nil, false
	}
	return o.PausedAt, true
}

// HasPausedAt returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasPausedAt() bool {
	if o != nil && !IsNil(o.PausedAt) {
		return true
	}

	return false
}

// SetPausedAt gets a reference to the given time.Time and assigns it to the PausedAt field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetPausedAt(v time.Time) {
	o.PausedAt = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetNeedsAttention returns the NeedsAttention field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetNeedsAttention() bool {
	if o == nil || IsNil(o.NeedsAttention) {
		var ret bool
		return ret
	}
	return *o.NeedsAttention
}

// GetNeedsAttentionOk returns a tuple with the NeedsAttention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetNeedsAttentionOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsAttention) {
		return nil, false
	}
	return o.NeedsAttention, true
}

// HasNeedsAttention returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasNeedsAttention() bool {
	if o != nil && !IsNil(o.NeedsAttention) {
		return true
	}

	return false
}

// SetNeedsAttention gets a reference to the given bool and assigns it to the NeedsAttention field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetNeedsAttention(v bool) {
	o.NeedsAttention = &v
}

// GetWasCanceledByFacebook returns the WasCanceledByFacebook field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetWasCanceledByFacebook() bool {
	if o == nil || IsNil(o.WasCanceledByFacebook) {
		var ret bool
		return ret
	}
	return *o.WasCanceledByFacebook
}

// GetWasCanceledByFacebookOk returns a tuple with the WasCanceledByFacebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetWasCanceledByFacebookOk() (*bool, bool) {
	if o == nil || IsNil(o.WasCanceledByFacebook) {
		return nil, false
	}
	return o.WasCanceledByFacebook, true
}

// HasWasCanceledByFacebook returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasWasCanceledByFacebook() bool {
	if o != nil && !IsNil(o.WasCanceledByFacebook) {
		return true
	}

	return false
}

// SetWasCanceledByFacebook gets a reference to the given bool and assigns it to the WasCanceledByFacebook field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetWasCanceledByFacebook(v bool) {
	o.WasCanceledByFacebook = &v
}

// GetIsConnected returns the IsConnected field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetIsConnected() bool {
	if o == nil || IsNil(o.IsConnected) {
		var ret bool
		return ret
	}
	return *o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetIsConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsConnected) {
		return nil, false
	}
	return o.IsConnected, true
}

// HasIsConnected returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasIsConnected() bool {
	if o != nil && !IsNil(o.IsConnected) {
		return true
	}

	return false
}

// SetIsConnected gets a reference to the given bool and assigns it to the IsConnected field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetIsConnected(v bool) {
	o.IsConnected = &v
}

// GetHasAudience returns the HasAudience field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetHasAudience() bool {
	if o == nil || IsNil(o.HasAudience) {
		var ret bool
		return ret
	}
	return *o.HasAudience
}

// GetHasAudienceOk returns a tuple with the HasAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetHasAudienceOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAudience) {
		return nil, false
	}
	return o.HasAudience, true
}

// HasHasAudience returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasHasAudience() bool {
	if o != nil && !IsNil(o.HasAudience) {
		return true
	}

	return false
}

// SetHasAudience gets a reference to the given bool and assigns it to the HasAudience field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetHasAudience(v bool) {
	o.HasAudience = &v
}

// GetHasContent returns the HasContent field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetHasContent() bool {
	if o == nil || IsNil(o.HasContent) {
		var ret bool
		return ret
	}
	return *o.HasContent
}

// GetHasContentOk returns a tuple with the HasContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetHasContentOk() (*bool, bool) {
	if o == nil || IsNil(o.HasContent) {
		return nil, false
	}
	return o.HasContent, true
}

// HasHasContent returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasHasContent() bool {
	if o != nil && !IsNil(o.HasContent) {
		return true
	}

	return false
}

// SetHasContent gets a reference to the given bool and assigns it to the HasContent field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetHasContent(v bool) {
	o.HasContent = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetChannel() GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel {
	if o == nil || IsNil(o.Channel) {
		var ret GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetChannelOk() (*GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel and assigns it to the Channel field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetChannel(v GetAllFacebookAds200ResponseFacebookAdsInnerAllOfChannel) {
	o.Channel = &v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetFeedback() GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback {
	if o == nil || IsNil(o.Feedback) {
		var ret GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetFeedbackOk() (*GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasFeedback() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback and assigns it to the Feedback field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetFeedback(v GetAllFacebookAds200ResponseFacebookAdsInnerAllOfFeedback) {
	o.Feedback = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetSite() GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite {
	if o == nil || IsNil(o.Site) {
		var ret GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetSiteOk() (*GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite and assigns it to the Site field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetSite(v GetAllFacebookAds200ResponseFacebookAdsInnerAllOfSite) {
	o.Site = &v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetAudience() GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience {
	if o == nil || IsNil(o.Audience) {
		var ret GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetAudienceOk() (*GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience and assigns it to the Audience field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetAudience(v GetAllFacebookAds200ResponseFacebookAdsInnerAllOfAudience) {
	o.Audience = &v
}

// GetBudget returns the Budget field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetBudget() GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget {
	if o == nil || IsNil(o.Budget) {
		var ret GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget
		return ret
	}
	return *o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetBudgetOk() (*GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget, bool) {
	if o == nil || IsNil(o.Budget) {
		return nil, false
	}
	return o.Budget, true
}

// HasBudget returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasBudget() bool {
	if o != nil && !IsNil(o.Budget) {
		return true
	}

	return false
}

// SetBudget gets a reference to the given GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget and assigns it to the Budget field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetBudget(v GetAllFacebookAds200ResponseFacebookAdsInnerAllOfBudget) {
	o.Budget = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetContent() GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent {
	if o == nil || IsNil(o.Content) {
		var ret GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetContentOk() (*GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent and assigns it to the Content field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetContent(v GetAllFacebookAds200ResponseFacebookAdsInnerAllOfContent) {
	o.Content = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *GetAllFacebookAds200ResponseFacebookAdsInner) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o GetAllFacebookAds200ResponseFacebookAdsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllFacebookAds200ResponseFacebookAdsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.WebId) {
		toSerialize["web_id"] = o.WebId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ShowReport) {
		toSerialize["show_report"] = o.ShowReport
	}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CanceledAt) {
		toSerialize["canceled_at"] = o.CanceledAt
	}
	if !IsNil(o.PublishedTime) {
		toSerialize["published_time"] = o.PublishedTime
	}
	if !IsNil(o.HasSegment) {
		toSerialize["has_segment"] = o.HasSegment
	}
	if !IsNil(o.ReportSummary) {
		toSerialize["report_summary"] = o.ReportSummary
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.Thumbnail) {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if !IsNil(o.EmailSourceName) {
		toSerialize["email_source_name"] = o.EmailSourceName
	}
	if !IsNil(o.PausedAt) {
		toSerialize["paused_at"] = o.PausedAt
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.NeedsAttention) {
		toSerialize["needs_attention"] = o.NeedsAttention
	}
	if !IsNil(o.WasCanceledByFacebook) {
		toSerialize["was_canceled_by_facebook"] = o.WasCanceledByFacebook
	}
	if !IsNil(o.IsConnected) {
		toSerialize["is_connected"] = o.IsConnected
	}
	if !IsNil(o.HasAudience) {
		toSerialize["has_audience"] = o.HasAudience
	}
	if !IsNil(o.HasContent) {
		toSerialize["has_content"] = o.HasContent
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	if !IsNil(o.Budget) {
		toSerialize["budget"] = o.Budget
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetAllFacebookAds200ResponseFacebookAdsInner struct {
	value *GetAllFacebookAds200ResponseFacebookAdsInner
	isSet bool
}

func (v NullableGetAllFacebookAds200ResponseFacebookAdsInner) Get() *GetAllFacebookAds200ResponseFacebookAdsInner {
	return v.value
}

func (v *NullableGetAllFacebookAds200ResponseFacebookAdsInner) Set(val *GetAllFacebookAds200ResponseFacebookAdsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllFacebookAds200ResponseFacebookAdsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllFacebookAds200ResponseFacebookAdsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllFacebookAds200ResponseFacebookAdsInner(val *GetAllFacebookAds200ResponseFacebookAdsInner) *NullableGetAllFacebookAds200ResponseFacebookAdsInner {
	return &NullableGetAllFacebookAds200ResponseFacebookAdsInner{value: val, isSet: true}
}

func (v NullableGetAllFacebookAds200ResponseFacebookAdsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllFacebookAds200ResponseFacebookAdsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


