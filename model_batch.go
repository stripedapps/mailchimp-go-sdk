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

// checks if the Batch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Batch{}

// Batch The status of a batch request
type Batch struct {
	// A string that uniquely identifies this batch request.
	Id *string `json:"id,omitempty"`
	// The status of the batch call. [Learn more](https://mailchimp.com/developer/marketing/guides/run-async-requests-batch-endpoint/#check-the-status-of-a-batch-operation) about the batch operation status.
	Status *string `json:"status,omitempty"`
	// The total number of operations to complete as part of this batch request. For GET requests requiring pagination, each page counts as a separate operation.
	TotalOperations *int32 `json:"total_operations,omitempty"`
	// The number of completed operations. This includes operations that returned an error.
	FinishedOperations *int32 `json:"finished_operations,omitempty"`
	// The number of completed operations that returned an error.
	ErroredOperations *int32 `json:"errored_operations,omitempty"`
	// The date and time when the server received the batch request in ISO 8601 format.
	SubmittedAt *time.Time `json:"submitted_at,omitempty"`
	// The date and time when all operations in the batch request completed in ISO 8601 format.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// The URL of the gzipped archive of the results of all the operations.
	ResponseBodyUrl *string `json:"response_body_url,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewBatch instantiates a new Batch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatch() *Batch {
	this := Batch{}
	return &this
}

// NewBatchWithDefaults instantiates a new Batch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchWithDefaults() *Batch {
	this := Batch{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Batch) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Batch) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Batch) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Batch) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Batch) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Batch) SetStatus(v string) {
	o.Status = &v
}

// GetTotalOperations returns the TotalOperations field value if set, zero value otherwise.
func (o *Batch) GetTotalOperations() int32 {
	if o == nil || IsNil(o.TotalOperations) {
		var ret int32
		return ret
	}
	return *o.TotalOperations
}

// GetTotalOperationsOk returns a tuple with the TotalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetTotalOperationsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalOperations) {
		return nil, false
	}
	return o.TotalOperations, true
}

// HasTotalOperations returns a boolean if a field has been set.
func (o *Batch) HasTotalOperations() bool {
	if o != nil && !IsNil(o.TotalOperations) {
		return true
	}

	return false
}

// SetTotalOperations gets a reference to the given int32 and assigns it to the TotalOperations field.
func (o *Batch) SetTotalOperations(v int32) {
	o.TotalOperations = &v
}

// GetFinishedOperations returns the FinishedOperations field value if set, zero value otherwise.
func (o *Batch) GetFinishedOperations() int32 {
	if o == nil || IsNil(o.FinishedOperations) {
		var ret int32
		return ret
	}
	return *o.FinishedOperations
}

// GetFinishedOperationsOk returns a tuple with the FinishedOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetFinishedOperationsOk() (*int32, bool) {
	if o == nil || IsNil(o.FinishedOperations) {
		return nil, false
	}
	return o.FinishedOperations, true
}

// HasFinishedOperations returns a boolean if a field has been set.
func (o *Batch) HasFinishedOperations() bool {
	if o != nil && !IsNil(o.FinishedOperations) {
		return true
	}

	return false
}

// SetFinishedOperations gets a reference to the given int32 and assigns it to the FinishedOperations field.
func (o *Batch) SetFinishedOperations(v int32) {
	o.FinishedOperations = &v
}

// GetErroredOperations returns the ErroredOperations field value if set, zero value otherwise.
func (o *Batch) GetErroredOperations() int32 {
	if o == nil || IsNil(o.ErroredOperations) {
		var ret int32
		return ret
	}
	return *o.ErroredOperations
}

// GetErroredOperationsOk returns a tuple with the ErroredOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetErroredOperationsOk() (*int32, bool) {
	if o == nil || IsNil(o.ErroredOperations) {
		return nil, false
	}
	return o.ErroredOperations, true
}

// HasErroredOperations returns a boolean if a field has been set.
func (o *Batch) HasErroredOperations() bool {
	if o != nil && !IsNil(o.ErroredOperations) {
		return true
	}

	return false
}

// SetErroredOperations gets a reference to the given int32 and assigns it to the ErroredOperations field.
func (o *Batch) SetErroredOperations(v int32) {
	o.ErroredOperations = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *Batch) GetSubmittedAt() time.Time {
	if o == nil || IsNil(o.SubmittedAt) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedAt) {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *Batch) HasSubmittedAt() bool {
	if o != nil && !IsNil(o.SubmittedAt) {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *Batch) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *Batch) GetCompletedAt() time.Time {
	if o == nil || IsNil(o.CompletedAt) {
		var ret time.Time
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *Batch) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given time.Time and assigns it to the CompletedAt field.
func (o *Batch) SetCompletedAt(v time.Time) {
	o.CompletedAt = &v
}

// GetResponseBodyUrl returns the ResponseBodyUrl field value if set, zero value otherwise.
func (o *Batch) GetResponseBodyUrl() string {
	if o == nil || IsNil(o.ResponseBodyUrl) {
		var ret string
		return ret
	}
	return *o.ResponseBodyUrl
}

// GetResponseBodyUrlOk returns a tuple with the ResponseBodyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetResponseBodyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseBodyUrl) {
		return nil, false
	}
	return o.ResponseBodyUrl, true
}

// HasResponseBodyUrl returns a boolean if a field has been set.
func (o *Batch) HasResponseBodyUrl() bool {
	if o != nil && !IsNil(o.ResponseBodyUrl) {
		return true
	}

	return false
}

// SetResponseBodyUrl gets a reference to the given string and assigns it to the ResponseBodyUrl field.
func (o *Batch) SetResponseBodyUrl(v string) {
	o.ResponseBodyUrl = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Batch) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Batch) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Batch) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *Batch) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o Batch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Batch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TotalOperations) {
		toSerialize["total_operations"] = o.TotalOperations
	}
	if !IsNil(o.FinishedOperations) {
		toSerialize["finished_operations"] = o.FinishedOperations
	}
	if !IsNil(o.ErroredOperations) {
		toSerialize["errored_operations"] = o.ErroredOperations
	}
	if !IsNil(o.SubmittedAt) {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if !IsNil(o.CompletedAt) {
		toSerialize["completed_at"] = o.CompletedAt
	}
	if !IsNil(o.ResponseBodyUrl) {
		toSerialize["response_body_url"] = o.ResponseBodyUrl
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBatch struct {
	value *Batch
	isSet bool
}

func (v NullableBatch) Get() *Batch {
	return v.value
}

func (v *NullableBatch) Set(val *Batch) {
	v.value = val
	v.isSet = true
}

func (v NullableBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatch(val *Batch) *NullableBatch {
	return &NullableBatch{value: val, isSet: true}
}

func (v NullableBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


