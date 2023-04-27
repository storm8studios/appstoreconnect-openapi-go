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

// checks if the RoutingAppCoverageCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingAppCoverageCreateRequest{}

// RoutingAppCoverageCreateRequest struct for RoutingAppCoverageCreateRequest
type RoutingAppCoverageCreateRequest struct {
	Data RoutingAppCoverageCreateRequestData `json:"data"`
}

// NewRoutingAppCoverageCreateRequest instantiates a new RoutingAppCoverageCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAppCoverageCreateRequest(data RoutingAppCoverageCreateRequestData) *RoutingAppCoverageCreateRequest {
	this := RoutingAppCoverageCreateRequest{}
	this.Data = data
	return &this
}

// NewRoutingAppCoverageCreateRequestWithDefaults instantiates a new RoutingAppCoverageCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAppCoverageCreateRequestWithDefaults() *RoutingAppCoverageCreateRequest {
	this := RoutingAppCoverageCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *RoutingAppCoverageCreateRequest) GetData() RoutingAppCoverageCreateRequestData {
	if o == nil {
		var ret RoutingAppCoverageCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageCreateRequest) GetDataOk() (*RoutingAppCoverageCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RoutingAppCoverageCreateRequest) SetData(v RoutingAppCoverageCreateRequestData) {
	o.Data = v
}

func (o RoutingAppCoverageCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingAppCoverageCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableRoutingAppCoverageCreateRequest struct {
	value *RoutingAppCoverageCreateRequest
	isSet bool
}

func (v NullableRoutingAppCoverageCreateRequest) Get() *RoutingAppCoverageCreateRequest {
	return v.value
}

func (v *NullableRoutingAppCoverageCreateRequest) Set(val *RoutingAppCoverageCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAppCoverageCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAppCoverageCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAppCoverageCreateRequest(val *RoutingAppCoverageCreateRequest) *NullableRoutingAppCoverageCreateRequest {
	return &NullableRoutingAppCoverageCreateRequest{value: val, isSet: true}
}

func (v NullableRoutingAppCoverageCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAppCoverageCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


