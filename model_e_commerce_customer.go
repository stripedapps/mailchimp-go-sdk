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

// checks if the ECommerceCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceCustomer{}

// ECommerceCustomer Information about a specific customer.
type ECommerceCustomer struct {
	// A unique identifier for the customer.
	Id *string `json:"id,omitempty"`
	// The customer's email address.
	EmailAddress *string `json:"email_address,omitempty"`
	// The customer's opt-in status. This value will never overwrite the opt-in status of a pre-existing Mailchimp list member, but will apply to list members that are added through the e-commerce API endpoints. Customers who don't opt in to your Mailchimp list [will be added as `Transactional` members](https://mailchimp.com/developer/marketing/docs/e-commerce/#customers).
	OptInStatus *bool `json:"opt_in_status,omitempty"`
	// The customer's company.
	Company *string `json:"company,omitempty"`
	// The customer's first name.
	FirstName *string `json:"first_name,omitempty"`
	// The customer's last name.
	LastName *string `json:"last_name,omitempty"`
	// The customer's total order count.
	OrdersCount *int32 `json:"orders_count,omitempty"`
	// The total amount the customer has spent.
	TotalSpent *float32 `json:"total_spent,omitempty"`
	Address *Address `json:"address,omitempty"`
	// The date and time the customer was created in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time the customer was last updated in ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewECommerceCustomer instantiates a new ECommerceCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceCustomer() *ECommerceCustomer {
	this := ECommerceCustomer{}
	return &this
}

// NewECommerceCustomerWithDefaults instantiates a new ECommerceCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceCustomerWithDefaults() *ECommerceCustomer {
	this := ECommerceCustomer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ECommerceCustomer) SetId(v string) {
	o.Id = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ECommerceCustomer) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetOptInStatus returns the OptInStatus field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetOptInStatus() bool {
	if o == nil || IsNil(o.OptInStatus) {
		var ret bool
		return ret
	}
	return *o.OptInStatus
}

// GetOptInStatusOk returns a tuple with the OptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetOptInStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.OptInStatus) {
		return nil, false
	}
	return o.OptInStatus, true
}

// HasOptInStatus returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasOptInStatus() bool {
	if o != nil && !IsNil(o.OptInStatus) {
		return true
	}

	return false
}

// SetOptInStatus gets a reference to the given bool and assigns it to the OptInStatus field.
func (o *ECommerceCustomer) SetOptInStatus(v bool) {
	o.OptInStatus = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ECommerceCustomer) SetCompany(v string) {
	o.Company = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ECommerceCustomer) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ECommerceCustomer) SetLastName(v string) {
	o.LastName = &v
}

// GetOrdersCount returns the OrdersCount field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetOrdersCount() int32 {
	if o == nil || IsNil(o.OrdersCount) {
		var ret int32
		return ret
	}
	return *o.OrdersCount
}

// GetOrdersCountOk returns a tuple with the OrdersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetOrdersCountOk() (*int32, bool) {
	if o == nil || IsNil(o.OrdersCount) {
		return nil, false
	}
	return o.OrdersCount, true
}

// HasOrdersCount returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasOrdersCount() bool {
	if o != nil && !IsNil(o.OrdersCount) {
		return true
	}

	return false
}

// SetOrdersCount gets a reference to the given int32 and assigns it to the OrdersCount field.
func (o *ECommerceCustomer) SetOrdersCount(v int32) {
	o.OrdersCount = &v
}

// GetTotalSpent returns the TotalSpent field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetTotalSpent() float32 {
	if o == nil || IsNil(o.TotalSpent) {
		var ret float32
		return ret
	}
	return *o.TotalSpent
}

// GetTotalSpentOk returns a tuple with the TotalSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetTotalSpentOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalSpent) {
		return nil, false
	}
	return o.TotalSpent, true
}

// HasTotalSpent returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasTotalSpent() bool {
	if o != nil && !IsNil(o.TotalSpent) {
		return true
	}

	return false
}

// SetTotalSpent gets a reference to the given float32 and assigns it to the TotalSpent field.
func (o *ECommerceCustomer) SetTotalSpent(v float32) {
	o.TotalSpent = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ECommerceCustomer) SetAddress(v Address) {
	o.Address = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ECommerceCustomer) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ECommerceCustomer) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ECommerceCustomer) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceCustomer) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ECommerceCustomer) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *ECommerceCustomer) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o ECommerceCustomer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	if !IsNil(o.OptInStatus) {
		toSerialize["opt_in_status"] = o.OptInStatus
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.OrdersCount) {
		toSerialize["orders_count"] = o.OrdersCount
	}
	if !IsNil(o.TotalSpent) {
		toSerialize["total_spent"] = o.TotalSpent
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableECommerceCustomer struct {
	value *ECommerceCustomer
	isSet bool
}

func (v NullableECommerceCustomer) Get() *ECommerceCustomer {
	return v.value
}

func (v *NullableECommerceCustomer) Set(val *ECommerceCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceCustomer(val *ECommerceCustomer) *NullableECommerceCustomer {
	return &NullableECommerceCustomer{value: val, isSet: true}
}

func (v NullableECommerceCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


