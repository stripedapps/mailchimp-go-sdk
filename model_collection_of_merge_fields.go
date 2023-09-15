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

// checks if the CollectionOfMergeFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionOfMergeFields{}

// CollectionOfMergeFields The [merge fields](https://mailchimp.com/developer/marketing/docs/merge-fields/) for an audience.
type CollectionOfMergeFields struct {
	// An array of objects, each representing a merge field resource.
	MergeFields []MergeField `json:"merge_fields,omitempty"`
	// The list id.
	ListId *string `json:"list_id,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewCollectionOfMergeFields instantiates a new CollectionOfMergeFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfMergeFields() *CollectionOfMergeFields {
	this := CollectionOfMergeFields{}
	return &this
}

// NewCollectionOfMergeFieldsWithDefaults instantiates a new CollectionOfMergeFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfMergeFieldsWithDefaults() *CollectionOfMergeFields {
	this := CollectionOfMergeFields{}
	return &this
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *CollectionOfMergeFields) GetMergeFields() []MergeField {
	if o == nil || IsNil(o.MergeFields) {
		var ret []MergeField
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMergeFields) GetMergeFieldsOk() ([]MergeField, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return nil, false
	}
	return o.MergeFields, true
}

// HasMergeFields returns a boolean if a field has been set.
func (o *CollectionOfMergeFields) HasMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given []MergeField and assigns it to the MergeFields field.
func (o *CollectionOfMergeFields) SetMergeFields(v []MergeField) {
	o.MergeFields = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *CollectionOfMergeFields) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMergeFields) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *CollectionOfMergeFields) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *CollectionOfMergeFields) SetListId(v string) {
	o.ListId = &v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *CollectionOfMergeFields) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMergeFields) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *CollectionOfMergeFields) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *CollectionOfMergeFields) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CollectionOfMergeFields) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfMergeFields) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CollectionOfMergeFields) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *CollectionOfMergeFields) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o CollectionOfMergeFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionOfMergeFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MergeFields) {
		toSerialize["merge_fields"] = o.MergeFields
	}
	if !IsNil(o.ListId) {
		toSerialize["list_id"] = o.ListId
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCollectionOfMergeFields struct {
	value *CollectionOfMergeFields
	isSet bool
}

func (v NullableCollectionOfMergeFields) Get() *CollectionOfMergeFields {
	return v.value
}

func (v *NullableCollectionOfMergeFields) Set(val *CollectionOfMergeFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfMergeFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfMergeFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfMergeFields(val *CollectionOfMergeFields) *NullableCollectionOfMergeFields {
	return &NullableCollectionOfMergeFields{value: val, isSet: true}
}

func (v NullableCollectionOfMergeFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfMergeFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


