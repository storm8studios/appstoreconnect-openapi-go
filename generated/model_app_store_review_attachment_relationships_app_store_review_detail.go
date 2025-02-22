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

// checks if the AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{}

// AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail struct for AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail
type AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData `json:"data,omitempty"`
}

// NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail instantiates a new AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail() *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	this := AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{}
	return &this
}

// NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetailWithDefaults instantiates a new AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewAttachmentRelationshipsAppStoreReviewDetailWithDefaults() *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	this := AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetData() AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData {
	if o == nil || IsNil(o.Data) {
		var ret AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) GetDataOk() (*AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData and assigns it to the Data field.
func (o *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) SetData(v AppStoreReviewAttachmentRelationshipsAppStoreReviewDetailData) {
	o.Data = &v
}

func (o AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail struct {
	value *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail
	isSet bool
}

func (v NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) Get() *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	return v.value
}

func (v *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) Set(val *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail(val *AppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail {
	return &NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail{value: val, isSet: true}
}

func (v NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewAttachmentRelationshipsAppStoreReviewDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


