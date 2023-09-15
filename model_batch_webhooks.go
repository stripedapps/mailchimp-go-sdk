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

// checks if the BatchWebhooks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchWebhooks{}

// BatchWebhooks Manage webhooks for batch requests.
type BatchWebhooks struct {
	// An array of objects, each representing a Batch Webhook.
	Webhooks []BatchWebhook `json:"webhooks,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewBatchWebhooks instantiates a new BatchWebhooks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchWebhooks() *BatchWebhooks {
	this := BatchWebhooks{}
	return &this
}

// NewBatchWebhooksWithDefaults instantiates a new BatchWebhooks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchWebhooksWithDefaults() *BatchWebhooks {
	this := BatchWebhooks{}
	return &this
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *BatchWebhooks) GetWebhooks() []BatchWebhook {
	if o == nil || IsNil(o.Webhooks) {
		var ret []BatchWebhook
		return ret
	}
	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchWebhooks) GetWebhooksOk() ([]BatchWebhook, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *BatchWebhooks) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []BatchWebhook and assigns it to the Webhooks field.
func (o *BatchWebhooks) SetWebhooks(v []BatchWebhook) {
	o.Webhooks = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *BatchWebhooks) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchWebhooks) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *BatchWebhooks) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *BatchWebhooks) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchWebhooks) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchWebhooks) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchWebhooks) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *BatchWebhooks) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o BatchWebhooks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchWebhooks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBatchWebhooks struct {
	value *BatchWebhooks
	isSet bool
}

func (v NullableBatchWebhooks) Get() *BatchWebhooks {
	return v.value
}

func (v *NullableBatchWebhooks) Set(val *BatchWebhooks) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchWebhooks) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchWebhooks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchWebhooks(val *BatchWebhooks) *NullableBatchWebhooks {
	return &NullableBatchWebhooks{value: val, isSet: true}
}

func (v NullableBatchWebhooks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchWebhooks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

