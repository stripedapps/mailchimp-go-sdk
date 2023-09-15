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

// checks if the ECommerceReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECommerceReport{}

// ECommerceReport E-Commerce stats for a campaign.
type ECommerceReport struct {
	// The total orders for a campaign.
	TotalOrders *int32 `json:"total_orders,omitempty"`
	// The total spent for a campaign. Calculated as the sum of all order totals with no deductions.
	TotalSpent *float32 `json:"total_spent,omitempty"`
	// The total revenue for a campaign. Calculated as the sum of all order totals minus shipping and tax totals.
	TotalRevenue *float32 `json:"total_revenue,omitempty"`
}

// NewECommerceReport instantiates a new ECommerceReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECommerceReport() *ECommerceReport {
	this := ECommerceReport{}
	return &this
}

// NewECommerceReportWithDefaults instantiates a new ECommerceReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECommerceReportWithDefaults() *ECommerceReport {
	this := ECommerceReport{}
	return &this
}

// GetTotalOrders returns the TotalOrders field value if set, zero value otherwise.
func (o *ECommerceReport) GetTotalOrders() int32 {
	if o == nil || IsNil(o.TotalOrders) {
		var ret int32
		return ret
	}
	return *o.TotalOrders
}

// GetTotalOrdersOk returns a tuple with the TotalOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceReport) GetTotalOrdersOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalOrders) {
		return nil, false
	}
	return o.TotalOrders, true
}

// HasTotalOrders returns a boolean if a field has been set.
func (o *ECommerceReport) HasTotalOrders() bool {
	if o != nil && !IsNil(o.TotalOrders) {
		return true
	}

	return false
}

// SetTotalOrders gets a reference to the given int32 and assigns it to the TotalOrders field.
func (o *ECommerceReport) SetTotalOrders(v int32) {
	o.TotalOrders = &v
}

// GetTotalSpent returns the TotalSpent field value if set, zero value otherwise.
func (o *ECommerceReport) GetTotalSpent() float32 {
	if o == nil || IsNil(o.TotalSpent) {
		var ret float32
		return ret
	}
	return *o.TotalSpent
}

// GetTotalSpentOk returns a tuple with the TotalSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceReport) GetTotalSpentOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalSpent) {
		return nil, false
	}
	return o.TotalSpent, true
}

// HasTotalSpent returns a boolean if a field has been set.
func (o *ECommerceReport) HasTotalSpent() bool {
	if o != nil && !IsNil(o.TotalSpent) {
		return true
	}

	return false
}

// SetTotalSpent gets a reference to the given float32 and assigns it to the TotalSpent field.
func (o *ECommerceReport) SetTotalSpent(v float32) {
	o.TotalSpent = &v
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *ECommerceReport) GetTotalRevenue() float32 {
	if o == nil || IsNil(o.TotalRevenue) {
		var ret float32
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECommerceReport) GetTotalRevenueOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRevenue) {
		return nil, false
	}
	return o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *ECommerceReport) HasTotalRevenue() bool {
	if o != nil && !IsNil(o.TotalRevenue) {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given float32 and assigns it to the TotalRevenue field.
func (o *ECommerceReport) SetTotalRevenue(v float32) {
	o.TotalRevenue = &v
}

func (o ECommerceReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECommerceReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalOrders) {
		toSerialize["total_orders"] = o.TotalOrders
	}
	if !IsNil(o.TotalSpent) {
		toSerialize["total_spent"] = o.TotalSpent
	}
	if !IsNil(o.TotalRevenue) {
		toSerialize["total_revenue"] = o.TotalRevenue
	}
	return toSerialize, nil
}

type NullableECommerceReport struct {
	value *ECommerceReport
	isSet bool
}

func (v NullableECommerceReport) Get() *ECommerceReport {
	return v.value
}

func (v *NullableECommerceReport) Set(val *ECommerceReport) {
	v.value = val
	v.isSet = true
}

func (v NullableECommerceReport) IsSet() bool {
	return v.isSet
}

func (v *NullableECommerceReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECommerceReport(val *ECommerceReport) *NullableECommerceReport {
	return &NullableECommerceReport{value: val, isSet: true}
}

func (v NullableECommerceReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECommerceReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


