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

// checks if the BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations{}

// BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations struct for BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations
type BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations struct {
	Data []BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner `json:"data"`
}

// NewBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations instantiates a new BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations(data []BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner) *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations {
	this := BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations{}
	this.Data = data
	return &this
}

// NewBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizationsWithDefaults instantiates a new BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizationsWithDefaults() *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations {
	this := BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) GetData() []BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner {
	if o == nil {
		var ret []BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) GetDataOk() ([]BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) SetData(v []BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner) {
	o.Data = v
}

func (o BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations struct {
	value *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations
	isSet bool
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) Get() *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations {
	return v.value
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) Set(val *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations(val *BetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations {
	return &NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


