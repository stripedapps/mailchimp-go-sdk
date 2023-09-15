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

// checks if the SubscriberList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberList{}

// SubscriberList Information about a specific list.
type SubscriberList struct {
	// A string that uniquely identifies this list.
	Id *string `json:"id,omitempty"`
	// The ID used in the Mailchimp web application. View this list in your Mailchimp account at `https://{dc}.admin.mailchimp.com/lists/members/?id={web_id}`.
	WebId *int32 `json:"web_id,omitempty"`
	// The name of the list.
	Name *string `json:"name,omitempty"`
	Contact *ListContact `json:"contact,omitempty"`
	// The [permission reminder](https://mailchimp.com/help/edit-the-permission-reminder/) for the list.
	PermissionReminder *string `json:"permission_reminder,omitempty"`
	// Whether campaigns for this list use the [Archive Bar](https://mailchimp.com/help/about-email-campaign-archives-and-pages/) in archives by default.
	UseArchiveBar *bool `json:"use_archive_bar,omitempty"`
	CampaignDefaults *CampaignDefaults `json:"campaign_defaults,omitempty"`
	// The email address to send [subscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
	NotifyOnSubscribe *string `json:"notify_on_subscribe,omitempty"`
	// The email address to send [unsubscribe notifications](https://mailchimp.com/help/change-subscribe-and-unsubscribe-notifications/) to.
	NotifyOnUnsubscribe *string `json:"notify_on_unsubscribe,omitempty"`
	// The date and time that this list was created in ISO 8601 format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// An auto-generated activity score for the list (0-5).
	ListRating *int32 `json:"list_rating,omitempty"`
	// Whether the list supports [multiple formats for emails](https://mailchimp.com/help/change-audience-name-defaults/). When set to `true`, subscribers can choose whether they want to receive HTML or plain-text emails. When set to `false`, subscribers will receive HTML emails, with a plain-text alternative backup.
	EmailTypeOption *bool `json:"email_type_option,omitempty"`
	// Our [url shortened](https://mailchimp.com/help/share-your-signup-form/) version of this list's subscribe form.
	SubscribeUrlShort *string `json:"subscribe_url_short,omitempty"`
	// The full version of this list's subscribe form (host will vary).
	SubscribeUrlLong *string `json:"subscribe_url_long,omitempty"`
	// The list's [Email Beamer](https://mailchimp.com/help/use-email-beamer-to-create-a-campaign/) address.
	BeamerAddress *string `json:"beamer_address,omitempty"`
	// Legacy - visibility settings are no longer used
	Visibility *string `json:"visibility,omitempty"`
	// Whether or not to require the subscriber to confirm subscription via email.
	DoubleOptin *bool `json:"double_optin,omitempty"`
	// Whether or not this list has a welcome automation connected. Welcome Automations: welcomeSeries, singleWelcome, emailFollowup.
	HasWelcome *bool `json:"has_welcome,omitempty"`
	// Whether or not the list has marketing permissions (eg. GDPR) enabled.
	MarketingPermissions *bool `json:"marketing_permissions,omitempty"`
	// Any list-specific modules installed for this list.
	Modules []string `json:"modules,omitempty"`
	Stats *Statistics `json:"stats,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewSubscriberList instantiates a new SubscriberList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberList() *SubscriberList {
	this := SubscriberList{}
	var useArchiveBar bool = false
	this.UseArchiveBar = &useArchiveBar
	var notifyOnSubscribe string = "false"
	this.NotifyOnSubscribe = &notifyOnSubscribe
	var notifyOnUnsubscribe string = "false"
	this.NotifyOnUnsubscribe = &notifyOnUnsubscribe
	var doubleOptin bool = false
	this.DoubleOptin = &doubleOptin
	var hasWelcome bool = false
	this.HasWelcome = &hasWelcome
	var marketingPermissions bool = false
	this.MarketingPermissions = &marketingPermissions
	return &this
}

// NewSubscriberListWithDefaults instantiates a new SubscriberList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberListWithDefaults() *SubscriberList {
	this := SubscriberList{}
	var useArchiveBar bool = false
	this.UseArchiveBar = &useArchiveBar
	var notifyOnSubscribe string = "false"
	this.NotifyOnSubscribe = &notifyOnSubscribe
	var notifyOnUnsubscribe string = "false"
	this.NotifyOnUnsubscribe = &notifyOnUnsubscribe
	var doubleOptin bool = false
	this.DoubleOptin = &doubleOptin
	var hasWelcome bool = false
	this.HasWelcome = &hasWelcome
	var marketingPermissions bool = false
	this.MarketingPermissions = &marketingPermissions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriberList) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriberList) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriberList) SetId(v string) {
	o.Id = &v
}

// GetWebId returns the WebId field value if set, zero value otherwise.
func (o *SubscriberList) GetWebId() int32 {
	if o == nil || IsNil(o.WebId) {
		var ret int32
		return ret
	}
	return *o.WebId
}

// GetWebIdOk returns a tuple with the WebId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetWebIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WebId) {
		return nil, false
	}
	return o.WebId, true
}

// HasWebId returns a boolean if a field has been set.
func (o *SubscriberList) HasWebId() bool {
	if o != nil && !IsNil(o.WebId) {
		return true
	}

	return false
}

// SetWebId gets a reference to the given int32 and assigns it to the WebId field.
func (o *SubscriberList) SetWebId(v int32) {
	o.WebId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriberList) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriberList) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriberList) SetName(v string) {
	o.Name = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *SubscriberList) GetContact() ListContact {
	if o == nil || IsNil(o.Contact) {
		var ret ListContact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetContactOk() (*ListContact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *SubscriberList) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ListContact and assigns it to the Contact field.
func (o *SubscriberList) SetContact(v ListContact) {
	o.Contact = &v
}

// GetPermissionReminder returns the PermissionReminder field value if set, zero value otherwise.
func (o *SubscriberList) GetPermissionReminder() string {
	if o == nil || IsNil(o.PermissionReminder) {
		var ret string
		return ret
	}
	return *o.PermissionReminder
}

// GetPermissionReminderOk returns a tuple with the PermissionReminder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetPermissionReminderOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionReminder) {
		return nil, false
	}
	return o.PermissionReminder, true
}

// HasPermissionReminder returns a boolean if a field has been set.
func (o *SubscriberList) HasPermissionReminder() bool {
	if o != nil && !IsNil(o.PermissionReminder) {
		return true
	}

	return false
}

// SetPermissionReminder gets a reference to the given string and assigns it to the PermissionReminder field.
func (o *SubscriberList) SetPermissionReminder(v string) {
	o.PermissionReminder = &v
}

// GetUseArchiveBar returns the UseArchiveBar field value if set, zero value otherwise.
func (o *SubscriberList) GetUseArchiveBar() bool {
	if o == nil || IsNil(o.UseArchiveBar) {
		var ret bool
		return ret
	}
	return *o.UseArchiveBar
}

// GetUseArchiveBarOk returns a tuple with the UseArchiveBar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetUseArchiveBarOk() (*bool, bool) {
	if o == nil || IsNil(o.UseArchiveBar) {
		return nil, false
	}
	return o.UseArchiveBar, true
}

// HasUseArchiveBar returns a boolean if a field has been set.
func (o *SubscriberList) HasUseArchiveBar() bool {
	if o != nil && !IsNil(o.UseArchiveBar) {
		return true
	}

	return false
}

// SetUseArchiveBar gets a reference to the given bool and assigns it to the UseArchiveBar field.
func (o *SubscriberList) SetUseArchiveBar(v bool) {
	o.UseArchiveBar = &v
}

// GetCampaignDefaults returns the CampaignDefaults field value if set, zero value otherwise.
func (o *SubscriberList) GetCampaignDefaults() CampaignDefaults {
	if o == nil || IsNil(o.CampaignDefaults) {
		var ret CampaignDefaults
		return ret
	}
	return *o.CampaignDefaults
}

// GetCampaignDefaultsOk returns a tuple with the CampaignDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetCampaignDefaultsOk() (*CampaignDefaults, bool) {
	if o == nil || IsNil(o.CampaignDefaults) {
		return nil, false
	}
	return o.CampaignDefaults, true
}

// HasCampaignDefaults returns a boolean if a field has been set.
func (o *SubscriberList) HasCampaignDefaults() bool {
	if o != nil && !IsNil(o.CampaignDefaults) {
		return true
	}

	return false
}

// SetCampaignDefaults gets a reference to the given CampaignDefaults and assigns it to the CampaignDefaults field.
func (o *SubscriberList) SetCampaignDefaults(v CampaignDefaults) {
	o.CampaignDefaults = &v
}

// GetNotifyOnSubscribe returns the NotifyOnSubscribe field value if set, zero value otherwise.
func (o *SubscriberList) GetNotifyOnSubscribe() string {
	if o == nil || IsNil(o.NotifyOnSubscribe) {
		var ret string
		return ret
	}
	return *o.NotifyOnSubscribe
}

// GetNotifyOnSubscribeOk returns a tuple with the NotifyOnSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetNotifyOnSubscribeOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyOnSubscribe) {
		return nil, false
	}
	return o.NotifyOnSubscribe, true
}

// HasNotifyOnSubscribe returns a boolean if a field has been set.
func (o *SubscriberList) HasNotifyOnSubscribe() bool {
	if o != nil && !IsNil(o.NotifyOnSubscribe) {
		return true
	}

	return false
}

// SetNotifyOnSubscribe gets a reference to the given string and assigns it to the NotifyOnSubscribe field.
func (o *SubscriberList) SetNotifyOnSubscribe(v string) {
	o.NotifyOnSubscribe = &v
}

// GetNotifyOnUnsubscribe returns the NotifyOnUnsubscribe field value if set, zero value otherwise.
func (o *SubscriberList) GetNotifyOnUnsubscribe() string {
	if o == nil || IsNil(o.NotifyOnUnsubscribe) {
		var ret string
		return ret
	}
	return *o.NotifyOnUnsubscribe
}

// GetNotifyOnUnsubscribeOk returns a tuple with the NotifyOnUnsubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetNotifyOnUnsubscribeOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyOnUnsubscribe) {
		return nil, false
	}
	return o.NotifyOnUnsubscribe, true
}

// HasNotifyOnUnsubscribe returns a boolean if a field has been set.
func (o *SubscriberList) HasNotifyOnUnsubscribe() bool {
	if o != nil && !IsNil(o.NotifyOnUnsubscribe) {
		return true
	}

	return false
}

// SetNotifyOnUnsubscribe gets a reference to the given string and assigns it to the NotifyOnUnsubscribe field.
func (o *SubscriberList) SetNotifyOnUnsubscribe(v string) {
	o.NotifyOnUnsubscribe = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SubscriberList) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SubscriberList) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *SubscriberList) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetListRating returns the ListRating field value if set, zero value otherwise.
func (o *SubscriberList) GetListRating() int32 {
	if o == nil || IsNil(o.ListRating) {
		var ret int32
		return ret
	}
	return *o.ListRating
}

// GetListRatingOk returns a tuple with the ListRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetListRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.ListRating) {
		return nil, false
	}
	return o.ListRating, true
}

// HasListRating returns a boolean if a field has been set.
func (o *SubscriberList) HasListRating() bool {
	if o != nil && !IsNil(o.ListRating) {
		return true
	}

	return false
}

// SetListRating gets a reference to the given int32 and assigns it to the ListRating field.
func (o *SubscriberList) SetListRating(v int32) {
	o.ListRating = &v
}

// GetEmailTypeOption returns the EmailTypeOption field value if set, zero value otherwise.
func (o *SubscriberList) GetEmailTypeOption() bool {
	if o == nil || IsNil(o.EmailTypeOption) {
		var ret bool
		return ret
	}
	return *o.EmailTypeOption
}

// GetEmailTypeOptionOk returns a tuple with the EmailTypeOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetEmailTypeOptionOk() (*bool, bool) {
	if o == nil || IsNil(o.EmailTypeOption) {
		return nil, false
	}
	return o.EmailTypeOption, true
}

// HasEmailTypeOption returns a boolean if a field has been set.
func (o *SubscriberList) HasEmailTypeOption() bool {
	if o != nil && !IsNil(o.EmailTypeOption) {
		return true
	}

	return false
}

// SetEmailTypeOption gets a reference to the given bool and assigns it to the EmailTypeOption field.
func (o *SubscriberList) SetEmailTypeOption(v bool) {
	o.EmailTypeOption = &v
}

// GetSubscribeUrlShort returns the SubscribeUrlShort field value if set, zero value otherwise.
func (o *SubscriberList) GetSubscribeUrlShort() string {
	if o == nil || IsNil(o.SubscribeUrlShort) {
		var ret string
		return ret
	}
	return *o.SubscribeUrlShort
}

// GetSubscribeUrlShortOk returns a tuple with the SubscribeUrlShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetSubscribeUrlShortOk() (*string, bool) {
	if o == nil || IsNil(o.SubscribeUrlShort) {
		return nil, false
	}
	return o.SubscribeUrlShort, true
}

// HasSubscribeUrlShort returns a boolean if a field has been set.
func (o *SubscriberList) HasSubscribeUrlShort() bool {
	if o != nil && !IsNil(o.SubscribeUrlShort) {
		return true
	}

	return false
}

// SetSubscribeUrlShort gets a reference to the given string and assigns it to the SubscribeUrlShort field.
func (o *SubscriberList) SetSubscribeUrlShort(v string) {
	o.SubscribeUrlShort = &v
}

// GetSubscribeUrlLong returns the SubscribeUrlLong field value if set, zero value otherwise.
func (o *SubscriberList) GetSubscribeUrlLong() string {
	if o == nil || IsNil(o.SubscribeUrlLong) {
		var ret string
		return ret
	}
	return *o.SubscribeUrlLong
}

// GetSubscribeUrlLongOk returns a tuple with the SubscribeUrlLong field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetSubscribeUrlLongOk() (*string, bool) {
	if o == nil || IsNil(o.SubscribeUrlLong) {
		return nil, false
	}
	return o.SubscribeUrlLong, true
}

// HasSubscribeUrlLong returns a boolean if a field has been set.
func (o *SubscriberList) HasSubscribeUrlLong() bool {
	if o != nil && !IsNil(o.SubscribeUrlLong) {
		return true
	}

	return false
}

// SetSubscribeUrlLong gets a reference to the given string and assigns it to the SubscribeUrlLong field.
func (o *SubscriberList) SetSubscribeUrlLong(v string) {
	o.SubscribeUrlLong = &v
}

// GetBeamerAddress returns the BeamerAddress field value if set, zero value otherwise.
func (o *SubscriberList) GetBeamerAddress() string {
	if o == nil || IsNil(o.BeamerAddress) {
		var ret string
		return ret
	}
	return *o.BeamerAddress
}

// GetBeamerAddressOk returns a tuple with the BeamerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetBeamerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.BeamerAddress) {
		return nil, false
	}
	return o.BeamerAddress, true
}

// HasBeamerAddress returns a boolean if a field has been set.
func (o *SubscriberList) HasBeamerAddress() bool {
	if o != nil && !IsNil(o.BeamerAddress) {
		return true
	}

	return false
}

// SetBeamerAddress gets a reference to the given string and assigns it to the BeamerAddress field.
func (o *SubscriberList) SetBeamerAddress(v string) {
	o.BeamerAddress = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *SubscriberList) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *SubscriberList) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *SubscriberList) SetVisibility(v string) {
	o.Visibility = &v
}

// GetDoubleOptin returns the DoubleOptin field value if set, zero value otherwise.
func (o *SubscriberList) GetDoubleOptin() bool {
	if o == nil || IsNil(o.DoubleOptin) {
		var ret bool
		return ret
	}
	return *o.DoubleOptin
}

// GetDoubleOptinOk returns a tuple with the DoubleOptin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetDoubleOptinOk() (*bool, bool) {
	if o == nil || IsNil(o.DoubleOptin) {
		return nil, false
	}
	return o.DoubleOptin, true
}

// HasDoubleOptin returns a boolean if a field has been set.
func (o *SubscriberList) HasDoubleOptin() bool {
	if o != nil && !IsNil(o.DoubleOptin) {
		return true
	}

	return false
}

// SetDoubleOptin gets a reference to the given bool and assigns it to the DoubleOptin field.
func (o *SubscriberList) SetDoubleOptin(v bool) {
	o.DoubleOptin = &v
}

// GetHasWelcome returns the HasWelcome field value if set, zero value otherwise.
func (o *SubscriberList) GetHasWelcome() bool {
	if o == nil || IsNil(o.HasWelcome) {
		var ret bool
		return ret
	}
	return *o.HasWelcome
}

// GetHasWelcomeOk returns a tuple with the HasWelcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetHasWelcomeOk() (*bool, bool) {
	if o == nil || IsNil(o.HasWelcome) {
		return nil, false
	}
	return o.HasWelcome, true
}

// HasHasWelcome returns a boolean if a field has been set.
func (o *SubscriberList) HasHasWelcome() bool {
	if o != nil && !IsNil(o.HasWelcome) {
		return true
	}

	return false
}

// SetHasWelcome gets a reference to the given bool and assigns it to the HasWelcome field.
func (o *SubscriberList) SetHasWelcome(v bool) {
	o.HasWelcome = &v
}

// GetMarketingPermissions returns the MarketingPermissions field value if set, zero value otherwise.
func (o *SubscriberList) GetMarketingPermissions() bool {
	if o == nil || IsNil(o.MarketingPermissions) {
		var ret bool
		return ret
	}
	return *o.MarketingPermissions
}

// GetMarketingPermissionsOk returns a tuple with the MarketingPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetMarketingPermissionsOk() (*bool, bool) {
	if o == nil || IsNil(o.MarketingPermissions) {
		return nil, false
	}
	return o.MarketingPermissions, true
}

// HasMarketingPermissions returns a boolean if a field has been set.
func (o *SubscriberList) HasMarketingPermissions() bool {
	if o != nil && !IsNil(o.MarketingPermissions) {
		return true
	}

	return false
}

// SetMarketingPermissions gets a reference to the given bool and assigns it to the MarketingPermissions field.
func (o *SubscriberList) SetMarketingPermissions(v bool) {
	o.MarketingPermissions = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *SubscriberList) GetModules() []string {
	if o == nil || IsNil(o.Modules) {
		var ret []string
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetModulesOk() ([]string, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *SubscriberList) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given []string and assigns it to the Modules field.
func (o *SubscriberList) SetModules(v []string) {
	o.Modules = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *SubscriberList) GetStats() Statistics {
	if o == nil || IsNil(o.Stats) {
		var ret Statistics
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetStatsOk() (*Statistics, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *SubscriberList) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given Statistics and assigns it to the Stats field.
func (o *SubscriberList) SetStats(v Statistics) {
	o.Stats = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriberList) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriberList) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriberList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *SubscriberList) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o SubscriberList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberList) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.PermissionReminder) {
		toSerialize["permission_reminder"] = o.PermissionReminder
	}
	if !IsNil(o.UseArchiveBar) {
		toSerialize["use_archive_bar"] = o.UseArchiveBar
	}
	if !IsNil(o.CampaignDefaults) {
		toSerialize["campaign_defaults"] = o.CampaignDefaults
	}
	if !IsNil(o.NotifyOnSubscribe) {
		toSerialize["notify_on_subscribe"] = o.NotifyOnSubscribe
	}
	if !IsNil(o.NotifyOnUnsubscribe) {
		toSerialize["notify_on_unsubscribe"] = o.NotifyOnUnsubscribe
	}
	if !IsNil(o.DateCreated) {
		toSerialize["date_created"] = o.DateCreated
	}
	if !IsNil(o.ListRating) {
		toSerialize["list_rating"] = o.ListRating
	}
	if !IsNil(o.EmailTypeOption) {
		toSerialize["email_type_option"] = o.EmailTypeOption
	}
	if !IsNil(o.SubscribeUrlShort) {
		toSerialize["subscribe_url_short"] = o.SubscribeUrlShort
	}
	if !IsNil(o.SubscribeUrlLong) {
		toSerialize["subscribe_url_long"] = o.SubscribeUrlLong
	}
	if !IsNil(o.BeamerAddress) {
		toSerialize["beamer_address"] = o.BeamerAddress
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.DoubleOptin) {
		toSerialize["double_optin"] = o.DoubleOptin
	}
	if !IsNil(o.HasWelcome) {
		toSerialize["has_welcome"] = o.HasWelcome
	}
	if !IsNil(o.MarketingPermissions) {
		toSerialize["marketing_permissions"] = o.MarketingPermissions
	}
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSubscriberList struct {
	value *SubscriberList
	isSet bool
}

func (v NullableSubscriberList) Get() *SubscriberList {
	return v.value
}

func (v *NullableSubscriberList) Set(val *SubscriberList) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberList) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberList(val *SubscriberList) *NullableSubscriberList {
	return &NullableSubscriberList{value: val, isSet: true}
}

func (v NullableSubscriberList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


