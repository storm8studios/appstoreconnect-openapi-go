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

// checks if the BetaGroupCreateRequestDataRelationshipsBetaTesters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupCreateRequestDataRelationshipsBetaTesters{}

// BetaGroupCreateRequestDataRelationshipsBetaTesters struct for BetaGroupCreateRequestDataRelationshipsBetaTesters
type BetaGroupCreateRequestDataRelationshipsBetaTesters struct {
	Data []BetaGroupRelationshipsBetaTestersDataInner `json:"data,omitempty"`
}

// NewBetaGroupCreateRequestDataRelationshipsBetaTesters instantiates a new BetaGroupCreateRequestDataRelationshipsBetaTesters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupCreateRequestDataRelationshipsBetaTesters() *BetaGroupCreateRequestDataRelationshipsBetaTesters {
	this := BetaGroupCreateRequestDataRelationshipsBetaTesters{}
	return &this
}

// NewBetaGroupCreateRequestDataRelationshipsBetaTestersWithDefaults instantiates a new BetaGroupCreateRequestDataRelationshipsBetaTesters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupCreateRequestDataRelationshipsBetaTestersWithDefaults() *BetaGroupCreateRequestDataRelationshipsBetaTesters {
	this := BetaGroupCreateRequestDataRelationshipsBetaTesters{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataRelationshipsBetaTesters) GetData() []BetaGroupRelationshipsBetaTestersDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []BetaGroupRelationshipsBetaTestersDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataRelationshipsBetaTesters) GetDataOk() ([]BetaGroupRelationshipsBetaTestersDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataRelationshipsBetaTesters) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BetaGroupRelationshipsBetaTestersDataInner and assigns it to the Data field.
func (o *BetaGroupCreateRequestDataRelationshipsBetaTesters) SetData(v []BetaGroupRelationshipsBetaTestersDataInner) {
	o.Data = v
}

func (o BetaGroupCreateRequestDataRelationshipsBetaTesters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupCreateRequestDataRelationshipsBetaTesters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBetaGroupCreateRequestDataRelationshipsBetaTesters struct {
	value *BetaGroupCreateRequestDataRelationshipsBetaTesters
	isSet bool
}

func (v NullableBetaGroupCreateRequestDataRelationshipsBetaTesters) Get() *BetaGroupCreateRequestDataRelationshipsBetaTesters {
	return v.value
}

func (v *NullableBetaGroupCreateRequestDataRelationshipsBetaTesters) Set(val *BetaGroupCreateRequestDataRelationshipsBetaTesters) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupCreateRequestDataRelationshipsBetaTesters) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupCreateRequestDataRelationshipsBetaTesters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupCreateRequestDataRelationshipsBetaTesters(val *BetaGroupCreateRequestDataRelationshipsBetaTesters) *NullableBetaGroupCreateRequestDataRelationshipsBetaTesters {
	return &NullableBetaGroupCreateRequestDataRelationshipsBetaTesters{value: val, isSet: true}
}

func (v NullableBetaGroupCreateRequestDataRelationshipsBetaTesters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupCreateRequestDataRelationshipsBetaTesters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


