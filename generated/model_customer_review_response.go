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

// checks if the CustomerReviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReviewResponse{}

// CustomerReviewResponse struct for CustomerReviewResponse
type CustomerReviewResponse struct {
	Data CustomerReview `json:"data"`
	Included []CustomerReviewResponseV1 `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewCustomerReviewResponse instantiates a new CustomerReviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReviewResponse(data CustomerReview, links DocumentLinks) *CustomerReviewResponse {
	this := CustomerReviewResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCustomerReviewResponseWithDefaults instantiates a new CustomerReviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewResponseWithDefaults() *CustomerReviewResponse {
	this := CustomerReviewResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerReviewResponse) GetData() CustomerReview {
	if o == nil {
		var ret CustomerReview
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponse) GetDataOk() (*CustomerReview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerReviewResponse) SetData(v CustomerReview) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CustomerReviewResponse) GetIncluded() []CustomerReviewResponseV1 {
	if o == nil || IsNil(o.Included) {
		var ret []CustomerReviewResponseV1
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponse) GetIncludedOk() ([]CustomerReviewResponseV1, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CustomerReviewResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CustomerReviewResponseV1 and assigns it to the Included field.
func (o *CustomerReviewResponse) SetIncluded(v []CustomerReviewResponseV1) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CustomerReviewResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CustomerReviewResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CustomerReviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerReviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableCustomerReviewResponse struct {
	value *CustomerReviewResponse
	isSet bool
}

func (v NullableCustomerReviewResponse) Get() *CustomerReviewResponse {
	return v.value
}

func (v *NullableCustomerReviewResponse) Set(val *CustomerReviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReviewResponse(val *CustomerReviewResponse) *NullableCustomerReviewResponse {
	return &NullableCustomerReviewResponse{value: val, isSet: true}
}

func (v NullableCustomerReviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerReviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


