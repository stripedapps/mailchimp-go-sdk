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

// checks if the GetAllFacebookAds200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllFacebookAds200Response{}

// GetAllFacebookAds200Response Contains an array of facebook ads.
type GetAllFacebookAds200Response struct {
	FacebookAds []GetAllFacebookAds200ResponseFacebookAdsInner `json:"facebook_ads,omitempty"`
	// The total number of items matching the query regardless of pagination.
	TotalItems *int32 `json:"total_items,omitempty"`
	// A list of link types and descriptions for the API schema documents.
	Links []ResourceLink `json:"_links,omitempty"`
}

// NewGetAllFacebookAds200Response instantiates a new GetAllFacebookAds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllFacebookAds200Response() *GetAllFacebookAds200Response {
	this := GetAllFacebookAds200Response{}
	return &this
}

// NewGetAllFacebookAds200ResponseWithDefaults instantiates a new GetAllFacebookAds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllFacebookAds200ResponseWithDefaults() *GetAllFacebookAds200Response {
	this := GetAllFacebookAds200Response{}
	return &this
}

// GetFacebookAds returns the FacebookAds field value if set, zero value otherwise.
func (o *GetAllFacebookAds200Response) GetFacebookAds() []GetAllFacebookAds200ResponseFacebookAdsInner {
	if o == nil || IsNil(o.FacebookAds) {
		var ret []GetAllFacebookAds200ResponseFacebookAdsInner
		return ret
	}
	return o.FacebookAds
}

// GetFacebookAdsOk returns a tuple with the FacebookAds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200Response) GetFacebookAdsOk() ([]GetAllFacebookAds200ResponseFacebookAdsInner, bool) {
	if o == nil || IsNil(o.FacebookAds) {
		return nil, false
	}
	return o.FacebookAds, true
}

// HasFacebookAds returns a boolean if a field has been set.
func (o *GetAllFacebookAds200Response) HasFacebookAds() bool {
	if o != nil && !IsNil(o.FacebookAds) {
		return true
	}

	return false
}

// SetFacebookAds gets a reference to the given []GetAllFacebookAds200ResponseFacebookAdsInner and assigns it to the FacebookAds field.
func (o *GetAllFacebookAds200Response) SetFacebookAds(v []GetAllFacebookAds200ResponseFacebookAdsInner) {
	o.FacebookAds = v
}

// GetTotalItems returns the TotalItems field value if set, zero value otherwise.
func (o *GetAllFacebookAds200Response) GetTotalItems() int32 {
	if o == nil || IsNil(o.TotalItems) {
		var ret int32
		return ret
	}
	return *o.TotalItems
}

// GetTotalItemsOk returns a tuple with the TotalItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200Response) GetTotalItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalItems) {
		return nil, false
	}
	return o.TotalItems, true
}

// HasTotalItems returns a boolean if a field has been set.
func (o *GetAllFacebookAds200Response) HasTotalItems() bool {
	if o != nil && !IsNil(o.TotalItems) {
		return true
	}

	return false
}

// SetTotalItems gets a reference to the given int32 and assigns it to the TotalItems field.
func (o *GetAllFacebookAds200Response) SetTotalItems(v int32) {
	o.TotalItems = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetAllFacebookAds200Response) GetLinks() []ResourceLink {
	if o == nil || IsNil(o.Links) {
		var ret []ResourceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllFacebookAds200Response) GetLinksOk() ([]ResourceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetAllFacebookAds200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ResourceLink and assigns it to the Links field.
func (o *GetAllFacebookAds200Response) SetLinks(v []ResourceLink) {
	o.Links = v
}

func (o GetAllFacebookAds200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllFacebookAds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FacebookAds) {
		toSerialize["facebook_ads"] = o.FacebookAds
	}
	if !IsNil(o.TotalItems) {
		toSerialize["total_items"] = o.TotalItems
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetAllFacebookAds200Response struct {
	value *GetAllFacebookAds200Response
	isSet bool
}

func (v NullableGetAllFacebookAds200Response) Get() *GetAllFacebookAds200Response {
	return v.value
}

func (v *NullableGetAllFacebookAds200Response) Set(val *GetAllFacebookAds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllFacebookAds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllFacebookAds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllFacebookAds200Response(val *GetAllFacebookAds200Response) *NullableGetAllFacebookAds200Response {
	return &NullableGetAllFacebookAds200Response{value: val, isSet: true}
}

func (v NullableGetAllFacebookAds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllFacebookAds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


