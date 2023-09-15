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

// checks if the ECommerceOrder1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceOrder1{}

// ECommerceOrder1 Information about a specific order.
type ECommerceOrder1 struct {
	// A unique identifier for the order.
	Id string `json:"id"`
	Customer ECommerceCustomer1 `json:"customer"`
	// A string that uniquely identifies the campaign for an order.
	CampaignId *string `json:"campaign_id,omitempty"`
	// The URL for the page where the buyer landed when entering the shop.
	LandingSite *string `json:"landing_site,omitempty"`
	// The order status. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FinancialStatus *string `json:"financial_status,omitempty"`
	// The fulfillment status for the order. Use this parameter to trigger [Order Notifications](https://mailchimp.com/developer/marketing/docs/e-commerce/#order-notifications).
	FulfillmentStatus *string `json:"fulfillment_status,omitempty"`
	// The three-letter ISO 4217 code for the currency that the store accepts.
	CurrencyCode string `json:"currency_code"`
	// The total for the order.
	OrderTotal float32 `json:"order_total"`
	// The URL for the order.
	OrderUrl *string `json:"order_url,omitempty"`
	// The total amount of the discounts to be applied to the price of the order.
	DiscountTotal *float32 `json:"discount_total,omitempty"`
	// The tax total for the order.
	TaxTotal *float32 `json:"tax_total,omitempty"`
	// The shipping total for the order.
	ShippingTotal *float32 `json:"shipping_total,omitempty"`
	// The Mailchimp tracking code for the order. Uses the 'mc_tc' parameter in E-Commerce tracking URLs.
	TrackingCode *string `json:"tracking_code,omitempty"`
	// The date and time the order was processed in ISO 8601 format.
	ProcessedAtForeign *time.Time `json:"processed_at_foreign,omitempty"`
	// The date and time the order was cancelled in ISO 8601 format. Note: passing a value for this parameter will cancel the order being created.
	CancelledAtForeign *time.Time `json:"cancelled_at_foreign,omitempty"`
	// The date and time the order was updated in ISO 8601 format.
	UpdatedAtForeign *time.Time `json:"updated_at_foreign,omitempty"`
	ShippingAddress *ShippingAddress1 `json:"shipping_address,omitempty"`
	BillingAddress *BillingAddress1 `json:"billing_address,omitempty"`
	// The promo codes applied on the order
	Promos []PromosInner1 `json:"promos,omitempty"`
	// An array of the order's line items.
	Lines []ECommerceOrderLineItem1 `json:"lines"`
	Outreach *Outreach1 `json:"outreach,omitempty"`
	// The tracking number associated with the order.
	TrackingNumber *string `json:"tracking_number,omitempty"`
	// The tracking carrier associated with the order.
	TrackingCarrier *string `json:"tracking_carrier,omitempty"`
	// The tracking URL associated with the order.
	TrackingUrl *string `json:"tracking_url,omitempty"`
}

// NewECommerceOrder1 instantiates a new ECommerceOrder1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceOrder1(id string, customer ECommerceCustomer1, currencyCode string, orderTotal float32, lines []ECommerceOrderLineItem1) *ECommerceOrder1 {
	this := ECommerceOrder1{}
	this.Id = id
	this.Customer = customer
	this.CurrencyCode = currencyCode
	this.OrderTotal = orderTotal
	this.Lines = lines
	return &this
}

// NewECommerceOrder1WithDefaults instantiates a new ECommerceOrder1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceOrder1WithDefaults() *ECommerceOrder1 {
	this := ECommerceOrder1{}
	return &this
}

// GetId returns the Id field value
func (o *ECommerceOrder1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ECommerceOrder1) SetId(v string) {
	o.Id = v
}

// GetCustomer returns the Customer field value
func (o *ECommerceOrder1) GetCustomer() ECommerceCustomer1 {
	if o == nil {
		var ret ECommerceCustomer1
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetCustomerOk() (*ECommerceCustomer1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *ECommerceOrder1) SetCustomer(v ECommerceCustomer1) {
	o.Customer = v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *ECommerceOrder1) SetCampaignId(v string) {
	o.CampaignId = &v
}

// GetLandingSite returns the LandingSite field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetLandingSite() string {
	if o == nil || IsNil(o.LandingSite) {
		var ret string
		return ret
	}
	return *o.LandingSite
}

// GetLandingSiteOk returns a tuple with the LandingSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetLandingSiteOk() (*string, bool) {
	if o == nil || IsNil(o.LandingSite) {
		return nil, false
	}
	return o.LandingSite, true
}

// HasLandingSite returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasLandingSite() bool {
	if o != nil && !IsNil(o.LandingSite) {
		return true
	}

	return false
}

// SetLandingSite gets a reference to the given string and assigns it to the LandingSite field.
func (o *ECommerceOrder1) SetLandingSite(v string) {
	o.LandingSite = &v
}

// GetFinancialStatus returns the FinancialStatus field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetFinancialStatus() string {
	if o == nil || IsNil(o.FinancialStatus) {
		var ret string
		return ret
	}
	return *o.FinancialStatus
}

// GetFinancialStatusOk returns a tuple with the FinancialStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetFinancialStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FinancialStatus) {
		return nil, false
	}
	return o.FinancialStatus, true
}

// HasFinancialStatus returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasFinancialStatus() bool {
	if o != nil && !IsNil(o.FinancialStatus) {
		return true
	}

	return false
}

// SetFinancialStatus gets a reference to the given string and assigns it to the FinancialStatus field.
func (o *ECommerceOrder1) SetFinancialStatus(v string) {
	o.FinancialStatus = &v
}

// GetFulfillmentStatus returns the FulfillmentStatus field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetFulfillmentStatus() string {
	if o == nil || IsNil(o.FulfillmentStatus) {
		var ret string
		return ret
	}
	return *o.FulfillmentStatus
}

// GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetFulfillmentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentStatus) {
		return nil, false
	}
	return o.FulfillmentStatus, true
}

// HasFulfillmentStatus returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasFulfillmentStatus() bool {
	if o != nil && !IsNil(o.FulfillmentStatus) {
		return true
	}

	return false
}

// SetFulfillmentStatus gets a reference to the given string and assigns it to the FulfillmentStatus field.
func (o *ECommerceOrder1) SetFulfillmentStatus(v string) {
	o.FulfillmentStatus = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *ECommerceOrder1) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *ECommerceOrder1) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetOrderTotal returns the OrderTotal field value
func (o *ECommerceOrder1) GetOrderTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetOrderTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderTotal, true
}

// SetOrderTotal sets field value
func (o *ECommerceOrder1) SetOrderTotal(v float32) {
	o.OrderTotal = v
}

// GetOrderUrl returns the OrderUrl field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetOrderUrl() string {
	if o == nil || IsNil(o.OrderUrl) {
		var ret string
		return ret
	}
	return *o.OrderUrl
}

// GetOrderUrlOk returns a tuple with the OrderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetOrderUrlOk() (*string, bool) {
	if o == nil || IsNil(o.OrderUrl) {
		return nil, false
	}
	return o.OrderUrl, true
}

// HasOrderUrl returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasOrderUrl() bool {
	if o != nil && !IsNil(o.OrderUrl) {
		return true
	}

	return false
}

// SetOrderUrl gets a reference to the given string and assigns it to the OrderUrl field.
func (o *ECommerceOrder1) SetOrderUrl(v string) {
	o.OrderUrl = &v
}

// GetDiscountTotal returns the DiscountTotal field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetDiscountTotal() float32 {
	if o == nil || IsNil(o.DiscountTotal) {
		var ret float32
		return ret
	}
	return *o.DiscountTotal
}

// GetDiscountTotalOk returns a tuple with the DiscountTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetDiscountTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountTotal) {
		return nil, false
	}
	return o.DiscountTotal, true
}

// HasDiscountTotal returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasDiscountTotal() bool {
	if o != nil && !IsNil(o.DiscountTotal) {
		return true
	}

	return false
}

// SetDiscountTotal gets a reference to the given float32 and assigns it to the DiscountTotal field.
func (o *ECommerceOrder1) SetDiscountTotal(v float32) {
	o.DiscountTotal = &v
}

// GetTaxTotal returns the TaxTotal field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetTaxTotal() float32 {
	if o == nil || IsNil(o.TaxTotal) {
		var ret float32
		return ret
	}
	return *o.TaxTotal
}

// GetTaxTotalOk returns a tuple with the TaxTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetTaxTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxTotal) {
		return nil, false
	}
	return o.TaxTotal, true
}

// HasTaxTotal returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasTaxTotal() bool {
	if o != nil && !IsNil(o.TaxTotal) {
		return true
	}

	return false
}

// SetTaxTotal gets a reference to the given float32 and assigns it to the TaxTotal field.
func (o *ECommerceOrder1) SetTaxTotal(v float32) {
	o.TaxTotal = &v
}

// GetShippingTotal returns the ShippingTotal field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetShippingTotal() float32 {
	if o == nil || IsNil(o.ShippingTotal) {
		var ret float32
		return ret
	}
	return *o.ShippingTotal
}

// GetShippingTotalOk returns a tuple with the ShippingTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetShippingTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippingTotal) {
		return nil, false
	}
	return o.ShippingTotal, true
}

// HasShippingTotal returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasShippingTotal() bool {
	if o != nil && !IsNil(o.ShippingTotal) {
		return true
	}

	return false
}

// SetShippingTotal gets a reference to the given float32 and assigns it to the ShippingTotal field.
func (o *ECommerceOrder1) SetShippingTotal(v float32) {
	o.ShippingTotal = &v
}

// GetTrackingCode returns the TrackingCode field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetTrackingCode() string {
	if o == nil || IsNil(o.TrackingCode) {
		var ret string
		return ret
	}
	return *o.TrackingCode
}

// GetTrackingCodeOk returns a tuple with the TrackingCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetTrackingCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingCode) {
		return nil, false
	}
	return o.TrackingCode, true
}

// HasTrackingCode returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasTrackingCode() bool {
	if o != nil && !IsNil(o.TrackingCode) {
		return true
	}

	return false
}

// SetTrackingCode gets a reference to the given string and assigns it to the TrackingCode field.
func (o *ECommerceOrder1) SetTrackingCode(v string) {
	o.TrackingCode = &v
}

// GetProcessedAtForeign returns the ProcessedAtForeign field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetProcessedAtForeign() time.Time {
	if o == nil || IsNil(o.ProcessedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.ProcessedAtForeign
}

// GetProcessedAtForeignOk returns a tuple with the ProcessedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetProcessedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ProcessedAtForeign) {
		return nil, false
	}
	return o.ProcessedAtForeign, true
}

// HasProcessedAtForeign returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasProcessedAtForeign() bool {
	if o != nil && !IsNil(o.ProcessedAtForeign) {
		return true
	}

	return false
}

// SetProcessedAtForeign gets a reference to the given time.Time and assigns it to the ProcessedAtForeign field.
func (o *ECommerceOrder1) SetProcessedAtForeign(v time.Time) {
	o.ProcessedAtForeign = &v
}

// GetCancelledAtForeign returns the CancelledAtForeign field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetCancelledAtForeign() time.Time {
	if o == nil || IsNil(o.CancelledAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.CancelledAtForeign
}

// GetCancelledAtForeignOk returns a tuple with the CancelledAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetCancelledAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CancelledAtForeign) {
		return nil, false
	}
	return o.CancelledAtForeign, true
}

// HasCancelledAtForeign returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasCancelledAtForeign() bool {
	if o != nil && !IsNil(o.CancelledAtForeign) {
		return true
	}

	return false
}

// SetCancelledAtForeign gets a reference to the given time.Time and assigns it to the CancelledAtForeign field.
func (o *ECommerceOrder1) SetCancelledAtForeign(v time.Time) {
	o.CancelledAtForeign = &v
}

// GetUpdatedAtForeign returns the UpdatedAtForeign field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetUpdatedAtForeign() time.Time {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtForeign
}

// GetUpdatedAtForeignOk returns a tuple with the UpdatedAtForeign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetUpdatedAtForeignOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAtForeign) {
		return nil, false
	}
	return o.UpdatedAtForeign, true
}

// HasUpdatedAtForeign returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasUpdatedAtForeign() bool {
	if o != nil && !IsNil(o.UpdatedAtForeign) {
		return true
	}

	return false
}

// SetUpdatedAtForeign gets a reference to the given time.Time and assigns it to the UpdatedAtForeign field.
func (o *ECommerceOrder1) SetUpdatedAtForeign(v time.Time) {
	o.UpdatedAtForeign = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetShippingAddress() ShippingAddress1 {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret ShippingAddress1
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetShippingAddressOk() (*ShippingAddress1, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given ShippingAddress1 and assigns it to the ShippingAddress field.
func (o *ECommerceOrder1) SetShippingAddress(v ShippingAddress1) {
	o.ShippingAddress = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetBillingAddress() BillingAddress1 {
	if o == nil || IsNil(o.BillingAddress) {
		var ret BillingAddress1
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetBillingAddressOk() (*BillingAddress1, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given BillingAddress1 and assigns it to the BillingAddress field.
func (o *ECommerceOrder1) SetBillingAddress(v BillingAddress1) {
	o.BillingAddress = &v
}

// GetPromos returns the Promos field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetPromos() []PromosInner1 {
	if o == nil || IsNil(o.Promos) {
		var ret []PromosInner1
		return ret
	}
	return o.Promos
}

// GetPromosOk returns a tuple with the Promos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetPromosOk() ([]PromosInner1, bool) {
	if o == nil || IsNil(o.Promos) {
		return nil, false
	}
	return o.Promos, true
}

// HasPromos returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasPromos() bool {
	if o != nil && !IsNil(o.Promos) {
		return true
	}

	return false
}

// SetPromos gets a reference to the given []PromosInner1 and assigns it to the Promos field.
func (o *ECommerceOrder1) SetPromos(v []PromosInner1) {
	o.Promos = v
}

// GetLines returns the Lines field value
func (o *ECommerceOrder1) GetLines() []ECommerceOrderLineItem1 {
	if o == nil {
		var ret []ECommerceOrderLineItem1
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetLinesOk() ([]ECommerceOrderLineItem1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *ECommerceOrder1) SetLines(v []ECommerceOrderLineItem1) {
	o.Lines = v
}

// GetOutreach returns the Outreach field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetOutreach() Outreach1 {
	if o == nil || IsNil(o.Outreach) {
		var ret Outreach1
		return ret
	}
	return *o.Outreach
}

// GetOutreachOk returns a tuple with the Outreach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetOutreachOk() (*Outreach1, bool) {
	if o == nil || IsNil(o.Outreach) {
		return nil, false
	}
	return o.Outreach, true
}

// HasOutreach returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasOutreach() bool {
	if o != nil && !IsNil(o.Outreach) {
		return true
	}

	return false
}

// SetOutreach gets a reference to the given Outreach1 and assigns it to the Outreach field.
func (o *ECommerceOrder1) SetOutreach(v Outreach1) {
	o.Outreach = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *ECommerceOrder1) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

// GetTrackingCarrier returns the TrackingCarrier field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetTrackingCarrier() string {
	if o == nil || IsNil(o.TrackingCarrier) {
		var ret string
		return ret
	}
	return *o.TrackingCarrier
}

// GetTrackingCarrierOk returns a tuple with the TrackingCarrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetTrackingCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingCarrier) {
		return nil, false
	}
	return o.TrackingCarrier, true
}

// HasTrackingCarrier returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasTrackingCarrier() bool {
	if o != nil && !IsNil(o.TrackingCarrier) {
		return true
	}

	return false
}

// SetTrackingCarrier gets a reference to the given string and assigns it to the TrackingCarrier field.
func (o *ECommerceOrder1) SetTrackingCarrier(v string) {
	o.TrackingCarrier = &v
}

// GetTrackingUrl returns the TrackingUrl field value if set, zero value otherwise.
func (o *ECommerceOrder1) GetTrackingUrl() string {
	if o == nil || IsNil(o.TrackingUrl) {
		var ret string
		return ret
	}
	return *o.TrackingUrl
}

// GetTrackingUrlOk returns a tuple with the TrackingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceOrder1) GetTrackingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingUrl) {
		return nil, false
	}
	return o.TrackingUrl, true
}

// HasTrackingUrl returns a boolean if a field has been set.
func (o *ECommerceOrder1) HasTrackingUrl() bool {
	if o != nil && !IsNil(o.TrackingUrl) {
		return true
	}

	return false
}

// SetTrackingUrl gets a reference to the given string and assigns it to the TrackingUrl field.
func (o *ECommerceOrder1) SetTrackingUrl(v string) {
	o.TrackingUrl = &v
}

func (o ECommerceOrder1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceOrder1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["customer"] = o.Customer
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	if !IsNil(o.LandingSite) {
		toSerialize["landing_site"] = o.LandingSite
	}
	if !IsNil(o.FinancialStatus) {
		toSerialize["financial_status"] = o.FinancialStatus
	}
	if !IsNil(o.FulfillmentStatus) {
		toSerialize["fulfillment_status"] = o.FulfillmentStatus
	}
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["order_total"] = o.OrderTotal
	if !IsNil(o.OrderUrl) {
		toSerialize["order_url"] = o.OrderUrl
	}
	if !IsNil(o.DiscountTotal) {
		toSerialize["discount_total"] = o.DiscountTotal
	}
	if !IsNil(o.TaxTotal) {
		toSerialize["tax_total"] = o.TaxTotal
	}
	if !IsNil(o.ShippingTotal) {
		toSerialize["shipping_total"] = o.ShippingTotal
	}
	if !IsNil(o.TrackingCode) {
		toSerialize["tracking_code"] = o.TrackingCode
	}
	if !IsNil(o.ProcessedAtForeign) {
		toSerialize["processed_at_foreign"] = o.ProcessedAtForeign
	}
	if !IsNil(o.CancelledAtForeign) {
		toSerialize["cancelled_at_foreign"] = o.CancelledAtForeign
	}
	if !IsNil(o.UpdatedAtForeign) {
		toSerialize["updated_at_foreign"] = o.UpdatedAtForeign
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shipping_address"] = o.ShippingAddress
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billing_address"] = o.BillingAddress
	}
	if !IsNil(o.Promos) {
		toSerialize["promos"] = o.Promos
	}
	toSerialize["lines"] = o.Lines
	if !IsNil(o.Outreach) {
		toSerialize["outreach"] = o.Outreach
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["tracking_number"] = o.TrackingNumber
	}
	if !IsNil(o.TrackingCarrier) {
		toSerialize["tracking_carrier"] = o.TrackingCarrier
	}
	if !IsNil(o.TrackingUrl) {
		toSerialize["tracking_url"] = o.TrackingUrl
	}
	return toSerialize, nil
}

type NullableECommerceOrder1 struct {
	value *ECommerceOrder1
	isSet bool
}

func (v NullableECommerceOrder1) Get() *ECommerceOrder1 {
	return v.value
}

func (v *NullableECommerceOrder1) Set(val *ECommerceOrder1) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceOrder1) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceOrder1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceOrder1(val *ECommerceOrder1) *NullableECommerceOrder1 {
	return &NullableECommerceOrder1{value: val, isSet: true}
}

func (v NullableECommerceOrder1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceOrder1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


