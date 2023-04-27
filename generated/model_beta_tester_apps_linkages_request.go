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

// checks if the BetaTesterAppsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterAppsLinkagesRequest{}

// BetaTesterAppsLinkagesRequest struct for BetaTesterAppsLinkagesRequest
type BetaTesterAppsLinkagesRequest struct {
	Data []AppAvailabilityRelationshipsAppData `json:"data"`
}

// NewBetaTesterAppsLinkagesRequest instantiates a new BetaTesterAppsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterAppsLinkagesRequest(data []AppAvailabilityRelationshipsAppData) *BetaTesterAppsLinkagesRequest {
	this := BetaTesterAppsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBetaTesterAppsLinkagesRequestWithDefaults instantiates a new BetaTesterAppsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterAppsLinkagesRequestWithDefaults() *BetaTesterAppsLinkagesRequest {
	this := BetaTesterAppsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterAppsLinkagesRequest) GetData() []AppAvailabilityRelationshipsAppData {
	if o == nil {
		var ret []AppAvailabilityRelationshipsAppData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterAppsLinkagesRequest) GetDataOk() ([]AppAvailabilityRelationshipsAppData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaTesterAppsLinkagesRequest) SetData(v []AppAvailabilityRelationshipsAppData) {
	o.Data = v
}

func (o BetaTesterAppsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterAppsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaTesterAppsLinkagesRequest struct {
	value *BetaTesterAppsLinkagesRequest
	isSet bool
}

func (v NullableBetaTesterAppsLinkagesRequest) Get() *BetaTesterAppsLinkagesRequest {
	return v.value
}

func (v *NullableBetaTesterAppsLinkagesRequest) Set(val *BetaTesterAppsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterAppsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterAppsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterAppsLinkagesRequest(val *BetaTesterAppsLinkagesRequest) *NullableBetaTesterAppsLinkagesRequest {
	return &NullableBetaTesterAppsLinkagesRequest{value: val, isSet: true}
}

func (v NullableBetaTesterAppsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterAppsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


