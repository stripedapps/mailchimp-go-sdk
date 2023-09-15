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

// checks if the OrderLines type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderLines{}

// OrderLines A collection of an order's line items.
type OrderLines struct {
	// The store id.
	StoreId *string `json:"store_id,omitempty"`
	// The order id.
	OrderId *string `json:"order_id,omitempty"`
	// An array of objects, each representing an order's line item.
	Lines []ECommerceOrderLineItem `json:"lines,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewOrderLines instantiates a new OrderLines object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderLines() *OrderLines {
	this := OrderLines{}
	return &this
}

// NewOrderLinesWithDefaults instantiates a new OrderLines object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderLinesWithDefaults() *OrderLines {
	this := OrderLines{}
	return &this
}

// GetStoreId returns the StoreId field value if set, zero value otherwise.
func (o *OrderLines) GetStoreId() string {
	if o == nil || IsNil(o.StoreId) {
		var ret string
		return ret
	}
	return *o.StoreId
}

// GetStoreIdOk returns a tuple with the StoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLines) GetStoreIdOk() (*string, bool) {
	if o == nil || IsNil(o.StoreId) {
		return nil, false
	}
	return o.StoreId, true
}

// HasStoreId returns a boolean if a field has been set.
func (o *OrderLines) HasStoreId() bool {
	if o != nil && !IsNil(o.StoreId) {
		return true
	}

	return false
}

// SetStoreId gets a reference to the given string and assigns it to the StoreId field.
func (o *OrderLines) SetStoreId(v string) {
	o.StoreId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderLines) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLines) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderLines) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderLines) SetOrderId(v string) {
	o.OrderId = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderLines) GetLines() []ECommerceOrderLineItem {
	if o == nil || IsNil(o.Lines) {
		var ret []ECommerceOrderLineItem
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLines) GetLinesOk() ([]ECommerceOrderLineItem, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderLines) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []ECommerceOrderLineItem and assigns it to the Lines field.
func (o *OrderLines) SetLines(v []ECommerceOrderLineItem) {
	o.Lines = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *OrderLines) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLines) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *OrderLines) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *OrderLines) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderLines) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLines) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderLines) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *OrderLines) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o OrderLines) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderLines) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StoreId) {
		toSerialize["store_id"] = o.StoreId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableOrderLines struct {
	value *OrderLines
	isSet bool
}

func (v NullableOrderLines) Get() *OrderLines {
	return v.value
}

func (v *NullableOrderLines) Set(val *OrderLines) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderLines) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderLines) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderLines(val *OrderLines) *NullableOrderLines {
	return &NullableOrderLines{value: val, isSet: true}
}

func (v NullableOrderLines) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderLines) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

