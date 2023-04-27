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

// checks if the BetaAppReviewDetailUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppReviewDetailUpdateRequestData{}

// BetaAppReviewDetailUpdateRequestData struct for BetaAppReviewDetailUpdateRequestData
type BetaAppReviewDetailUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreReviewDetailAttributes `json:"attributes,omitempty"`
}

// NewBetaAppReviewDetailUpdateRequestData instantiates a new BetaAppReviewDetailUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppReviewDetailUpdateRequestData(type_ string, id string) *BetaAppReviewDetailUpdateRequestData {
	this := BetaAppReviewDetailUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaAppReviewDetailUpdateRequestDataWithDefaults instantiates a new BetaAppReviewDetailUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppReviewDetailUpdateRequestDataWithDefaults() *BetaAppReviewDetailUpdateRequestData {
	this := BetaAppReviewDetailUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaAppReviewDetailUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewDetailUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaAppReviewDetailUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaAppReviewDetailUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewDetailUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaAppReviewDetailUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaAppReviewDetailUpdateRequestData) GetAttributes() AppStoreReviewDetailAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreReviewDetailAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppReviewDetailUpdateRequestData) GetAttributesOk() (*AppStoreReviewDetailAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaAppReviewDetailUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreReviewDetailAttributes and assigns it to the Attributes field.
func (o *BetaAppReviewDetailUpdateRequestData) SetAttributes(v AppStoreReviewDetailAttributes) {
	o.Attributes = &v
}

func (o BetaAppReviewDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppReviewDetailUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableBetaAppReviewDetailUpdateRequestData struct {
	value *BetaAppReviewDetailUpdateRequestData
	isSet bool
}

func (v NullableBetaAppReviewDetailUpdateRequestData) Get() *BetaAppReviewDetailUpdateRequestData {
	return v.value
}

func (v *NullableBetaAppReviewDetailUpdateRequestData) Set(val *BetaAppReviewDetailUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppReviewDetailUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppReviewDetailUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppReviewDetailUpdateRequestData(val *BetaAppReviewDetailUpdateRequestData) *NullableBetaAppReviewDetailUpdateRequestData {
	return &NullableBetaAppReviewDetailUpdateRequestData{value: val, isSet: true}
}

func (v NullableBetaAppReviewDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppReviewDetailUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


