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

// checks if the BatchUpdateListMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchUpdateListMembers{}

// BatchUpdateListMembers Batch update list members.
type BatchUpdateListMembers struct {
	// An array of objects, each representing a new member that was added to the list.
	NewMembers []ListMembers `json:"new_members,omitempty"`
	// An array of objects, each representing an existing list member whose subscription status was updated.
	UpdatedMembers []ListMembers `json:"updated_members,omitempty"`
	// An array of objects, each representing an email address that could not be added to the list or updated and an error message providing more details.
	Errors []ErrorsInner `json:"errors,omitempty"`
	// The total number of items matching the query, irrespective of pagination.
	TotalCreated *int32 `json:"total_created,omitempty"`
	// The total number of items matching the query, irrespective of pagination.
	TotalUpdated *int32 `json:"total_updated,omitempty"`
	// The total number of items matching the query, irrespective of pagination.
	ErrorCount *int32 `json:"error_count,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewBatchUpdateListMembers instantiates a new BatchUpdateListMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchUpdateListMembers() *BatchUpdateListMembers {
	this := BatchUpdateListMembers{}
	return &this
}

// NewBatchUpdateListMembersWithDefaults instantiates a new BatchUpdateListMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchUpdateListMembersWithDefaults() *BatchUpdateListMembers {
	this := BatchUpdateListMembers{}
	return &this
}

// GetNewMembers returns the NewMembers field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetNewMembers() []ListMembers {
	if o == nil || IsNil(o.NewMembers) {
		var ret []ListMembers
		return ret
	}
	return o.NewMembers
}

// GetNewMembersOk returns a tuple with the NewMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetNewMembersOk() ([]ListMembers, bool) {
	if o == nil || IsNil(o.NewMembers) {
		return nil, false
	}
	return o.NewMembers, true
}

// HasNewMembers returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasNewMembers() bool {
	if o != nil && !IsNil(o.NewMembers) {
		return true
	}

	return false
}

// SetNewMembers gets a reference to the given []ListMembers and assigns it to the NewMembers field.
func (o *BatchUpdateListMembers) SetNewMembers(v []ListMembers) {
	o.NewMembers = v
}

// GetUpdatedMembers returns the UpdatedMembers field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetUpdatedMembers() []ListMembers {
	if o == nil || IsNil(o.UpdatedMembers) {
		var ret []ListMembers
		return ret
	}
	return o.UpdatedMembers
}

// GetUpdatedMembersOk returns a tuple with the UpdatedMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetUpdatedMembersOk() ([]ListMembers, bool) {
	if o == nil || IsNil(o.UpdatedMembers) {
		return nil, false
	}
	return o.UpdatedMembers, true
}

// HasUpdatedMembers returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasUpdatedMembers() bool {
	if o != nil && !IsNil(o.UpdatedMembers) {
		return true
	}

	return false
}

// SetUpdatedMembers gets a reference to the given []ListMembers and assigns it to the UpdatedMembers field.
func (o *BatchUpdateListMembers) SetUpdatedMembers(v []ListMembers) {
	o.UpdatedMembers = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetErrors() []ErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetErrorsOk() ([]ErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorsInner and assigns it to the Errors field.
func (o *BatchUpdateListMembers) SetErrors(v []ErrorsInner) {
	o.Errors = v
}

// GetTotalCreated returns the TotalCreated field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetTotalCreated() int32 {
	if o == nil || IsNil(o.TotalCreated) {
		var ret int32
		return ret
	}
	return *o.TotalCreated
}

// GetTotalCreatedOk returns a tuple with the TotalCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetTotalCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCreated) {
		return nil, false
	}
	return o.TotalCreated, true
}

// HasTotalCreated returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasTotalCreated() bool {
	if o != nil && !IsNil(o.TotalCreated) {
		return true
	}

	return false
}

// SetTotalCreated gets a reference to the given int32 and assigns it to the TotalCreated field.
func (o *BatchUpdateListMembers) SetTotalCreated(v int32) {
	o.TotalCreated = &v
}

// GetTotalUpdated returns the TotalUpdated field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetTotalUpdated() int32 {
	if o == nil || IsNil(o.TotalUpdated) {
		var ret int32
		return ret
	}
	return *o.TotalUpdated
}

// GetTotalUpdatedOk returns a tuple with the TotalUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetTotalUpdatedOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalUpdated) {
		return nil, false
	}
	return o.TotalUpdated, true
}

// HasTotalUpdated returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasTotalUpdated() bool {
	if o != nil && !IsNil(o.TotalUpdated) {
		return true
	}

	return false
}

// SetTotalUpdated gets a reference to the given int32 and assigns it to the TotalUpdated field.
func (o *BatchUpdateListMembers) SetTotalUpdated(v int32) {
	o.TotalUpdated = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetErrorCount() int32 {
	if o == nil || IsNil(o.ErrorCount) {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetErrorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasErrorCount() bool {
	if o != nil && !IsNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *BatchUpdateListMembers) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchUpdateListMembers) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchUpdateListMembers) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchUpdateListMembers) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *BatchUpdateListMembers) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o BatchUpdateListMembers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchUpdateListMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewMembers) {
		toSerialize["new_members"] = o.NewMembers
	}
	if !IsNil(o.UpdatedMembers) {
		toSerialize["updated_members"] = o.UpdatedMembers
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.TotalCreated) {
		toSerialize["total_created"] = o.TotalCreated
	}
	if !IsNil(o.TotalUpdated) {
		toSerialize["total_updated"] = o.TotalUpdated
	}
	if !IsNil(o.ErrorCount) {
		toSerialize["error_count"] = o.ErrorCount
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBatchUpdateListMembers struct {
	value *BatchUpdateListMembers
	isSet bool
}

func (v NullableBatchUpdateListMembers) Get() *BatchUpdateListMembers {
	return v.value
}

func (v *NullableBatchUpdateListMembers) Set(val *BatchUpdateListMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchUpdateListMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchUpdateListMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchUpdateListMembers(val *BatchUpdateListMembers) *NullableBatchUpdateListMembers {
	return &NullableBatchUpdateListMembers{value: val, isSet: true}
}

func (v NullableBatchUpdateListMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchUpdateListMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


