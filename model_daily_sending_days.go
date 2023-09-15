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

// checks if the DailySendingDays type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailySendingDays{}

// DailySendingDays The days of the week to send a daily RSS Campaign.
type DailySendingDays struct {
	// Sends the daily RSS Campaign on Sundays.
	Sunday *bool `json:"sunday,omitempty"`
	// Sends the daily RSS Campaign on Mondays.
	Monday *bool `json:"monday,omitempty"`
	// Sends the daily RSS Campaign on Tuesdays.
	Tuesday *bool `json:"tuesday,omitempty"`
	// Sends the daily RSS Campaign on Wednesdays.
	Wednesday *bool `json:"wednesday,omitempty"`
	// Sends the daily RSS Campaign on Thursdays.
	Thursday *bool `json:"thursday,omitempty"`
	// Sends the daily RSS Campaign on Fridays.
	Friday *bool `json:"friday,omitempty"`
	// Sends the daily RSS Campaign on Saturdays.
	Saturday *bool `json:"saturday,omitempty"`
}

// NewDailySendingDays instantiates a new DailySendingDays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailySendingDays() *DailySendingDays {
	this := DailySendingDays{}
	return &this
}

// NewDailySendingDaysWithDefaults instantiates a new DailySendingDays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailySendingDaysWithDefaults() *DailySendingDays {
	this := DailySendingDays{}
	return &this
}

// GetSunday returns the Sunday field value if set, zero value otherwise.
func (o *DailySendingDays) GetSunday() bool {
	if o == nil || IsNil(o.Sunday) {
		var ret bool
		return ret
	}
	return *o.Sunday
}

// GetSundayOk returns a tuple with the Sunday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetSundayOk() (*bool, bool) {
	if o == nil || IsNil(o.Sunday) {
		return nil, false
	}
	return o.Sunday, true
}

// HasSunday returns a boolean if a field has been set.
func (o *DailySendingDays) HasSunday() bool {
	if o != nil && !IsNil(o.Sunday) {
		return true
	}

	return false
}

// SetSunday gets a reference to the given bool and assigns it to the Sunday field.
func (o *DailySendingDays) SetSunday(v bool) {
	o.Sunday = &v
}

// GetMonday returns the Monday field value if set, zero value otherwise.
func (o *DailySendingDays) GetMonday() bool {
	if o == nil || IsNil(o.Monday) {
		var ret bool
		return ret
	}
	return *o.Monday
}

// GetMondayOk returns a tuple with the Monday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetMondayOk() (*bool, bool) {
	if o == nil || IsNil(o.Monday) {
		return nil, false
	}
	return o.Monday, true
}

// HasMonday returns a boolean if a field has been set.
func (o *DailySendingDays) HasMonday() bool {
	if o != nil && !IsNil(o.Monday) {
		return true
	}

	return false
}

// SetMonday gets a reference to the given bool and assigns it to the Monday field.
func (o *DailySendingDays) SetMonday(v bool) {
	o.Monday = &v
}

// GetTuesday returns the Tuesday field value if set, zero value otherwise.
func (o *DailySendingDays) GetTuesday() bool {
	if o == nil || IsNil(o.Tuesday) {
		var ret bool
		return ret
	}
	return *o.Tuesday
}

// GetTuesdayOk returns a tuple with the Tuesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetTuesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Tuesday) {
		return nil, false
	}
	return o.Tuesday, true
}

// HasTuesday returns a boolean if a field has been set.
func (o *DailySendingDays) HasTuesday() bool {
	if o != nil && !IsNil(o.Tuesday) {
		return true
	}

	return false
}

// SetTuesday gets a reference to the given bool and assigns it to the Tuesday field.
func (o *DailySendingDays) SetTuesday(v bool) {
	o.Tuesday = &v
}

// GetWednesday returns the Wednesday field value if set, zero value otherwise.
func (o *DailySendingDays) GetWednesday() bool {
	if o == nil || IsNil(o.Wednesday) {
		var ret bool
		return ret
	}
	return *o.Wednesday
}

// GetWednesdayOk returns a tuple with the Wednesday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetWednesdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Wednesday) {
		return nil, false
	}
	return o.Wednesday, true
}

// HasWednesday returns a boolean if a field has been set.
func (o *DailySendingDays) HasWednesday() bool {
	if o != nil && !IsNil(o.Wednesday) {
		return true
	}

	return false
}

// SetWednesday gets a reference to the given bool and assigns it to the Wednesday field.
func (o *DailySendingDays) SetWednesday(v bool) {
	o.Wednesday = &v
}

// GetThursday returns the Thursday field value if set, zero value otherwise.
func (o *DailySendingDays) GetThursday() bool {
	if o == nil || IsNil(o.Thursday) {
		var ret bool
		return ret
	}
	return *o.Thursday
}

// GetThursdayOk returns a tuple with the Thursday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetThursdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Thursday) {
		return nil, false
	}
	return o.Thursday, true
}

// HasThursday returns a boolean if a field has been set.
func (o *DailySendingDays) HasThursday() bool {
	if o != nil && !IsNil(o.Thursday) {
		return true
	}

	return false
}

// SetThursday gets a reference to the given bool and assigns it to the Thursday field.
func (o *DailySendingDays) SetThursday(v bool) {
	o.Thursday = &v
}

// GetFriday returns the Friday field value if set, zero value otherwise.
func (o *DailySendingDays) GetFriday() bool {
	if o == nil || IsNil(o.Friday) {
		var ret bool
		return ret
	}
	return *o.Friday
}

// GetFridayOk returns a tuple with the Friday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetFridayOk() (*bool, bool) {
	if o == nil || IsNil(o.Friday) {
		return nil, false
	}
	return o.Friday, true
}

// HasFriday returns a boolean if a field has been set.
func (o *DailySendingDays) HasFriday() bool {
	if o != nil && !IsNil(o.Friday) {
		return true
	}

	return false
}

// SetFriday gets a reference to the given bool and assigns it to the Friday field.
func (o *DailySendingDays) SetFriday(v bool) {
	o.Friday = &v
}

// GetSaturday returns the Saturday field value if set, zero value otherwise.
func (o *DailySendingDays) GetSaturday() bool {
	if o == nil || IsNil(o.Saturday) {
		var ret bool
		return ret
	}
	return *o.Saturday
}

// GetSaturdayOk returns a tuple with the Saturday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySendingDays) GetSaturdayOk() (*bool, bool) {
	if o == nil || IsNil(o.Saturday) {
		return nil, false
	}
	return o.Saturday, true
}

// HasSaturday returns a boolean if a field has been set.
func (o *DailySendingDays) HasSaturday() bool {
	if o != nil && !IsNil(o.Saturday) {
		return true
	}

	return false
}

// SetSaturday gets a reference to the given bool and assigns it to the Saturday field.
func (o *DailySendingDays) SetSaturday(v bool) {
	o.Saturday = &v
}

func (o DailySendingDays) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailySendingDays) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sunday) {
		toSerialize["sunday"] = o.Sunday
	}
	if !IsNil(o.Monday) {
		toSerialize["monday"] = o.Monday
	}
	if !IsNil(o.Tuesday) {
		toSerialize["tuesday"] = o.Tuesday
	}
	if !IsNil(o.Wednesday) {
		toSerialize["wednesday"] = o.Wednesday
	}
	if !IsNil(o.Thursday) {
		toSerialize["thursday"] = o.Thursday
	}
	if !IsNil(o.Friday) {
		toSerialize["friday"] = o.Friday
	}
	if !IsNil(o.Saturday) {
		toSerialize["saturday"] = o.Saturday
	}
	return toSerialize, nil
}

type NullableDailySendingDays struct {
	value *DailySendingDays
	isSet bool
}

func (v NullableDailySendingDays) Get() *DailySendingDays {
	return v.value
}

func (v *NullableDailySendingDays) Set(val *DailySendingDays) {
	v.value = val
	v.isSet = true
}

func (v NullableDailySendingDays) IsSet() bool {
	return v.isSet
}

func (v *NullableDailySendingDays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailySendingDays(val *DailySendingDays) *NullableDailySendingDays {
	return &NullableDailySendingDays{value: val, isSet: true}
}

func (v NullableDailySendingDays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailySendingDays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


