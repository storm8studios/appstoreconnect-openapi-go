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

// checks if the BetaGroupRelationshipsBetaTestersDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupRelationshipsBetaTestersDataInner{}

// BetaGroupRelationshipsBetaTestersDataInner struct for BetaGroupRelationshipsBetaTestersDataInner
type BetaGroupRelationshipsBetaTestersDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewBetaGroupRelationshipsBetaTestersDataInner instantiates a new BetaGroupRelationshipsBetaTestersDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupRelationshipsBetaTestersDataInner(type_ string, id string) *BetaGroupRelationshipsBetaTestersDataInner {
	this := BetaGroupRelationshipsBetaTestersDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaGroupRelationshipsBetaTestersDataInnerWithDefaults instantiates a new BetaGroupRelationshipsBetaTestersDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupRelationshipsBetaTestersDataInnerWithDefaults() *BetaGroupRelationshipsBetaTestersDataInner {
	this := BetaGroupRelationshipsBetaTestersDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *BetaGroupRelationshipsBetaTestersDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaGroupRelationshipsBetaTestersDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaGroupRelationshipsBetaTestersDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaGroupRelationshipsBetaTestersDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaGroupRelationshipsBetaTestersDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaGroupRelationshipsBetaTestersDataInner) SetId(v string) {
	o.Id = v
}

func (o BetaGroupRelationshipsBetaTestersDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupRelationshipsBetaTestersDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableBetaGroupRelationshipsBetaTestersDataInner struct {
	value *BetaGroupRelationshipsBetaTestersDataInner
	isSet bool
}

func (v NullableBetaGroupRelationshipsBetaTestersDataInner) Get() *BetaGroupRelationshipsBetaTestersDataInner {
	return v.value
}

func (v *NullableBetaGroupRelationshipsBetaTestersDataInner) Set(val *BetaGroupRelationshipsBetaTestersDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupRelationshipsBetaTestersDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupRelationshipsBetaTestersDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupRelationshipsBetaTestersDataInner(val *BetaGroupRelationshipsBetaTestersDataInner) *NullableBetaGroupRelationshipsBetaTestersDataInner {
	return &NullableBetaGroupRelationshipsBetaTestersDataInner{value: val, isSet: true}
}

func (v NullableBetaGroupRelationshipsBetaTestersDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupRelationshipsBetaTestersDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


