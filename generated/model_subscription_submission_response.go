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

// checks if the SubscriptionSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionSubmissionResponse{}

// SubscriptionSubmissionResponse struct for SubscriptionSubmissionResponse
type SubscriptionSubmissionResponse struct {
	Data SubscriptionSubmission `json:"data"`
	Included []Subscription `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewSubscriptionSubmissionResponse instantiates a new SubscriptionSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionSubmissionResponse(data SubscriptionSubmission, links DocumentLinks) *SubscriptionSubmissionResponse {
	this := SubscriptionSubmissionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionSubmissionResponseWithDefaults instantiates a new SubscriptionSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionSubmissionResponseWithDefaults() *SubscriptionSubmissionResponse {
	this := SubscriptionSubmissionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionSubmissionResponse) GetData() SubscriptionSubmission {
	if o == nil {
		var ret SubscriptionSubmission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSubmissionResponse) GetDataOk() (*SubscriptionSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionSubmissionResponse) SetData(v SubscriptionSubmission) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionSubmissionResponse) GetIncluded() []Subscription {
	if o == nil || IsNil(o.Included) {
		var ret []Subscription
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionSubmissionResponse) GetIncludedOk() ([]Subscription, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionSubmissionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Subscription and assigns it to the Included field.
func (o *SubscriptionSubmissionResponse) SetIncluded(v []Subscription) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionSubmissionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSubmissionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionSubmissionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionSubmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionSubmissionResponse struct {
	value *SubscriptionSubmissionResponse
	isSet bool
}

func (v NullableSubscriptionSubmissionResponse) Get() *SubscriptionSubmissionResponse {
	return v.value
}

func (v *NullableSubscriptionSubmissionResponse) Set(val *SubscriptionSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionSubmissionResponse(val *SubscriptionSubmissionResponse) *NullableSubscriptionSubmissionResponse {
	return &NullableSubscriptionSubmissionResponse{value: val, isSet: true}
}

func (v NullableSubscriptionSubmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


