/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
)

// checks if the ReviewSubmissionItemAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionItemAttributes{}

// ReviewSubmissionItemAttributes struct for ReviewSubmissionItemAttributes
type ReviewSubmissionItemAttributes struct {
	State *string `json:"state,omitempty"`
}

// NewReviewSubmissionItemAttributes instantiates a new ReviewSubmissionItemAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionItemAttributes() *ReviewSubmissionItemAttributes {
	this := ReviewSubmissionItemAttributes{}
	return &this
}

// NewReviewSubmissionItemAttributesWithDefaults instantiates a new ReviewSubmissionItemAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionItemAttributesWithDefaults() *ReviewSubmissionItemAttributes {
	this := ReviewSubmissionItemAttributes{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ReviewSubmissionItemAttributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemAttributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ReviewSubmissionItemAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ReviewSubmissionItemAttributes) SetState(v string) {
	o.State = &v
}

func (o ReviewSubmissionItemAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionItemAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableReviewSubmissionItemAttributes struct {
	value *ReviewSubmissionItemAttributes
	isSet bool
}

func (v NullableReviewSubmissionItemAttributes) Get() *ReviewSubmissionItemAttributes {
	return v.value
}

func (v *NullableReviewSubmissionItemAttributes) Set(val *ReviewSubmissionItemAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionItemAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionItemAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionItemAttributes(val *ReviewSubmissionItemAttributes) *NullableReviewSubmissionItemAttributes {
	return &NullableReviewSubmissionItemAttributes{value: val, isSet: true}
}

func (v NullableReviewSubmissionItemAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionItemAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


