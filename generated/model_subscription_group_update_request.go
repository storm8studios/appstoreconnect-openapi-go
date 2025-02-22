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

// checks if the SubscriptionGroupUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupUpdateRequest{}

// SubscriptionGroupUpdateRequest struct for SubscriptionGroupUpdateRequest
type SubscriptionGroupUpdateRequest struct {
	Data SubscriptionGroupUpdateRequestData `json:"data"`
}

// NewSubscriptionGroupUpdateRequest instantiates a new SubscriptionGroupUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupUpdateRequest(data SubscriptionGroupUpdateRequestData) *SubscriptionGroupUpdateRequest {
	this := SubscriptionGroupUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionGroupUpdateRequestWithDefaults instantiates a new SubscriptionGroupUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupUpdateRequestWithDefaults() *SubscriptionGroupUpdateRequest {
	this := SubscriptionGroupUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupUpdateRequest) GetData() SubscriptionGroupUpdateRequestData {
	if o == nil {
		var ret SubscriptionGroupUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupUpdateRequest) GetDataOk() (*SubscriptionGroupUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupUpdateRequest) SetData(v SubscriptionGroupUpdateRequestData) {
	o.Data = v
}

func (o SubscriptionGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionGroupUpdateRequest struct {
	value *SubscriptionGroupUpdateRequest
	isSet bool
}

func (v NullableSubscriptionGroupUpdateRequest) Get() *SubscriptionGroupUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionGroupUpdateRequest) Set(val *SubscriptionGroupUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupUpdateRequest(val *SubscriptionGroupUpdateRequest) *NullableSubscriptionGroupUpdateRequest {
	return &NullableSubscriptionGroupUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


