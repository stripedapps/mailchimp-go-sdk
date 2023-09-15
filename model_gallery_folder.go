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

// checks if the GalleryFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GalleryFolder{}

// GalleryFolder An individual folder listed in the File Manager.
type GalleryFolder struct {
	// The unique id for the folder.
	Id *int32 `json:"id,omitempty"`
	// The name of the folder.
	Name *string `json:"name,omitempty"`
	// The number of files in the folder.
	FileCount *int32 `json:"file_count,omitempty"`
	// The date and time a file was added to the File Manager in ISO 8601 format.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The username of the profile that created the folder.
	CreatedBy *string `json:"created_by,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewGalleryFolder instantiates a new GalleryFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGalleryFolder() *GalleryFolder {
	this := GalleryFolder{}
	return &this
}

// NewGalleryFolderWithDefaults instantiates a new GalleryFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGalleryFolderWithDefaults() *GalleryFolder {
	this := GalleryFolder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GalleryFolder) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFolder) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GalleryFolder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GalleryFolder) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GalleryFolder) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFolder) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GalleryFolder) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GalleryFolder) SetName(v string) {
	o.Name = &v
}

// GetFileCount returns the FileCount field value if set, zero value otherwise.
func (o *GalleryFolder) GetFileCount() int32 {
	if o == nil || IsNil(o.FileCount) {
		var ret int32
		return ret
	}
	return *o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFolder) GetFileCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FileCount) {
		return nil, false
	}
	return o.FileCount, true
}

// HasFileCount returns a boolean if a field has been set.
func (o *GalleryFolder) HasFileCount() bool {
	if o != nil && !IsNil(o.FileCount) {
		return true
	}

	return false
}

// SetFileCount gets a reference to the given int32 and assigns it to the FileCount field.
func (o *GalleryFolder) SetFileCount(v int32) {
	o.FileCount = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GalleryFolder) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFolder) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GalleryFolder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GalleryFolder) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *GalleryFolder) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFolder) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *GalleryFolder) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *GalleryFolder) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GalleryFolder) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalleryFolder) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GalleryFolder) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *GalleryFolder) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o GalleryFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GalleryFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.FileCount) {
		toSerialize["file_count"] = o.FileCount
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["created_by"] = o.CreatedBy
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGalleryFolder struct {
	value *GalleryFolder
	isSet bool
}

func (v NullableGalleryFolder) Get() *GalleryFolder {
	return v.value
}

func (v *NullableGalleryFolder) Set(val *GalleryFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableGalleryFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableGalleryFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGalleryFolder(val *GalleryFolder) *NullableGalleryFolder {
	return &NullableGalleryFolder{value: val, isSet: true}
}

func (v NullableGalleryFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGalleryFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


