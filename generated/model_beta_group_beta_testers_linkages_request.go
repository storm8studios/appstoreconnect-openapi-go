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

// checks if the BetaGroupBetaTestersLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupBetaTestersLinkagesRequest{}

// BetaGroupBetaTestersLinkagesRequest struct for BetaGroupBetaTestersLinkagesRequest
type BetaGroupBetaTestersLinkagesRequest struct {
	Data []BetaGroupRelationshipsBetaTestersDataInner `json:"data"`
}

// NewBetaGroupBetaTestersLinkagesRequest instantiates a new BetaGroupBetaTestersLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupBetaTestersLinkagesRequest(data []BetaGroupRelationshipsBetaTestersDataInner) *BetaGroupBetaTestersLinkagesRequest {
	this := BetaGroupBetaTestersLinkagesRequest{}
	this.Data = data
	return &this
}

// NewBetaGroupBetaTestersLinkagesRequestWithDefaults instantiates a new BetaGroupBetaTestersLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupBetaTestersLinkagesRequestWithDefaults() *BetaGroupBetaTestersLinkagesRequest {
	this := BetaGroupBetaTestersLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupBetaTestersLinkagesRequest) GetData() []BetaGroupRelationshipsBetaTestersDataInner {
	if o == nil {
		var ret []BetaGroupRelationshipsBetaTestersDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBetaTestersLinkagesRequest) GetDataOk() ([]BetaGroupRelationshipsBetaTestersDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaGroupBetaTestersLinkagesRequest) SetData(v []BetaGroupRelationshipsBetaTestersDataInner) {
	o.Data = v
}

func (o BetaGroupBetaTestersLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupBetaTestersLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaGroupBetaTestersLinkagesRequest struct {
	value *BetaGroupBetaTestersLinkagesRequest
	isSet bool
}

func (v NullableBetaGroupBetaTestersLinkagesRequest) Get() *BetaGroupBetaTestersLinkagesRequest {
	return v.value
}

func (v *NullableBetaGroupBetaTestersLinkagesRequest) Set(val *BetaGroupBetaTestersLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupBetaTestersLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupBetaTestersLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupBetaTestersLinkagesRequest(val *BetaGroupBetaTestersLinkagesRequest) *NullableBetaGroupBetaTestersLinkagesRequest {
	return &NullableBetaGroupBetaTestersLinkagesRequest{value: val, isSet: true}
}

func (v NullableBetaGroupBetaTestersLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupBetaTestersLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


