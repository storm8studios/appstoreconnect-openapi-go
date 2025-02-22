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

// checks if the BundleIdUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdUpdateRequest{}

// BundleIdUpdateRequest struct for BundleIdUpdateRequest
type BundleIdUpdateRequest struct {
	Data BundleIdUpdateRequestData `json:"data"`
}

// NewBundleIdUpdateRequest instantiates a new BundleIdUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdUpdateRequest(data BundleIdUpdateRequestData) *BundleIdUpdateRequest {
	this := BundleIdUpdateRequest{}
	this.Data = data
	return &this
}

// NewBundleIdUpdateRequestWithDefaults instantiates a new BundleIdUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdUpdateRequestWithDefaults() *BundleIdUpdateRequest {
	this := BundleIdUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdUpdateRequest) GetData() BundleIdUpdateRequestData {
	if o == nil {
		var ret BundleIdUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdUpdateRequest) GetDataOk() (*BundleIdUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BundleIdUpdateRequest) SetData(v BundleIdUpdateRequestData) {
	o.Data = v
}

func (o BundleIdUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBundleIdUpdateRequest struct {
	value *BundleIdUpdateRequest
	isSet bool
}

func (v NullableBundleIdUpdateRequest) Get() *BundleIdUpdateRequest {
	return v.value
}

func (v *NullableBundleIdUpdateRequest) Set(val *BundleIdUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdUpdateRequest(val *BundleIdUpdateRequest) *NullableBundleIdUpdateRequest {
	return &NullableBundleIdUpdateRequest{value: val, isSet: true}
}

func (v NullableBundleIdUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


