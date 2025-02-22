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

// checks if the InAppPurchaseSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseSubmissionResponse{}

// InAppPurchaseSubmissionResponse struct for InAppPurchaseSubmissionResponse
type InAppPurchaseSubmissionResponse struct {
	Data InAppPurchaseSubmission `json:"data"`
	Included []InAppPurchaseV2 `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewInAppPurchaseSubmissionResponse instantiates a new InAppPurchaseSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseSubmissionResponse(data InAppPurchaseSubmission, links DocumentLinks) *InAppPurchaseSubmissionResponse {
	this := InAppPurchaseSubmissionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewInAppPurchaseSubmissionResponseWithDefaults instantiates a new InAppPurchaseSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseSubmissionResponseWithDefaults() *InAppPurchaseSubmissionResponse {
	this := InAppPurchaseSubmissionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseSubmissionResponse) GetData() InAppPurchaseSubmission {
	if o == nil {
		var ret InAppPurchaseSubmission
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseSubmissionResponse) GetDataOk() (*InAppPurchaseSubmission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseSubmissionResponse) SetData(v InAppPurchaseSubmission) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *InAppPurchaseSubmissionResponse) GetIncluded() []InAppPurchaseV2 {
	if o == nil || IsNil(o.Included) {
		var ret []InAppPurchaseV2
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseSubmissionResponse) GetIncludedOk() ([]InAppPurchaseV2, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *InAppPurchaseSubmissionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []InAppPurchaseV2 and assigns it to the Included field.
func (o *InAppPurchaseSubmissionResponse) SetIncluded(v []InAppPurchaseV2) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *InAppPurchaseSubmissionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseSubmissionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *InAppPurchaseSubmissionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o InAppPurchaseSubmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableInAppPurchaseSubmissionResponse struct {
	value *InAppPurchaseSubmissionResponse
	isSet bool
}

func (v NullableInAppPurchaseSubmissionResponse) Get() *InAppPurchaseSubmissionResponse {
	return v.value
}

func (v *NullableInAppPurchaseSubmissionResponse) Set(val *InAppPurchaseSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseSubmissionResponse(val *InAppPurchaseSubmissionResponse) *NullableInAppPurchaseSubmissionResponse {
	return &NullableInAppPurchaseSubmissionResponse{value: val, isSet: true}
}

func (v NullableInAppPurchaseSubmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


