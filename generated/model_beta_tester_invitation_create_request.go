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

// checks if the BetaTesterInvitationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterInvitationCreateRequest{}

// BetaTesterInvitationCreateRequest struct for BetaTesterInvitationCreateRequest
type BetaTesterInvitationCreateRequest struct {
	Data BetaTesterInvitationCreateRequestData `json:"data"`
}

// NewBetaTesterInvitationCreateRequest instantiates a new BetaTesterInvitationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterInvitationCreateRequest(data BetaTesterInvitationCreateRequestData) *BetaTesterInvitationCreateRequest {
	this := BetaTesterInvitationCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaTesterInvitationCreateRequestWithDefaults instantiates a new BetaTesterInvitationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterInvitationCreateRequestWithDefaults() *BetaTesterInvitationCreateRequest {
	this := BetaTesterInvitationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterInvitationCreateRequest) GetData() BetaTesterInvitationCreateRequestData {
	if o == nil {
		var ret BetaTesterInvitationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationCreateRequest) GetDataOk() (*BetaTesterInvitationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaTesterInvitationCreateRequest) SetData(v BetaTesterInvitationCreateRequestData) {
	o.Data = v
}

func (o BetaTesterInvitationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterInvitationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaTesterInvitationCreateRequest struct {
	value *BetaTesterInvitationCreateRequest
	isSet bool
}

func (v NullableBetaTesterInvitationCreateRequest) Get() *BetaTesterInvitationCreateRequest {
	return v.value
}

func (v *NullableBetaTesterInvitationCreateRequest) Set(val *BetaTesterInvitationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterInvitationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterInvitationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterInvitationCreateRequest(val *BetaTesterInvitationCreateRequest) *NullableBetaTesterInvitationCreateRequest {
	return &NullableBetaTesterInvitationCreateRequest{value: val, isSet: true}
}

func (v NullableBetaTesterInvitationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterInvitationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


