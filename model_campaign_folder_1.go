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

// checks if the CampaignFolder1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignFolder1{}

// CampaignFolder1 A folder used to organize campaigns.
type CampaignFolder1 struct {
	// Name to associate with the folder.
	Name string `json:"name"`
}

// NewCampaignFolder1 instantiates a new CampaignFolder1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignFolder1(name string) *CampaignFolder1 {
	this := CampaignFolder1{}
	this.Name = name
	return &this
}

// NewCampaignFolder1WithDefaults instantiates a new CampaignFolder1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignFolder1WithDefaults() *CampaignFolder1 {
	this := CampaignFolder1{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignFolder1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignFolder1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignFolder1) SetName(v string) {
	o.Name = v
}

func (o CampaignFolder1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignFolder1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableCampaignFolder1 struct {
	value *CampaignFolder1
	isSet bool
}

func (v NullableCampaignFolder1) Get() *CampaignFolder1 {
	return v.value
}

func (v *NullableCampaignFolder1) Set(val *CampaignFolder1) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignFolder1) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignFolder1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignFolder1(val *CampaignFolder1) *NullableCampaignFolder1 {
	return &NullableCampaignFolder1{value: val, isSet: true}
}

func (v NullableCampaignFolder1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignFolder1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

