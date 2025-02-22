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

// checks if the BundleIdResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdResponse{}

// BundleIdResponse struct for BundleIdResponse
type BundleIdResponse struct {
	Data BundleId `json:"data"`
	Included []BundleIdsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewBundleIdResponse instantiates a new BundleIdResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdResponse(data BundleId, links DocumentLinks) *BundleIdResponse {
	this := BundleIdResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBundleIdResponseWithDefaults instantiates a new BundleIdResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdResponseWithDefaults() *BundleIdResponse {
	this := BundleIdResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdResponse) GetData() BundleId {
	if o == nil {
		var ret BundleId
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdResponse) GetDataOk() (*BundleId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdResponse) SetData(v BundleId) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BundleIdResponse) GetIncluded() []BundleIdsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []BundleIdsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdResponse) GetIncludedOk() ([]BundleIdsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BundleIdResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BundleIdsResponseIncludedInner and assigns it to the Included field.
func (o *BundleIdResponse) SetIncluded(v []BundleIdsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BundleIdResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BundleIdResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BundleIdResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BundleIdResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBundleIdResponse struct {
	value *BundleIdResponse
	isSet bool
}

func (v NullableBundleIdResponse) Get() *BundleIdResponse {
	return v.value
}

func (v *NullableBundleIdResponse) Set(val *BundleIdResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdResponse(val *BundleIdResponse) *NullableBundleIdResponse {
	return &NullableBundleIdResponse{value: val, isSet: true}
}

func (v NullableBundleIdResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


