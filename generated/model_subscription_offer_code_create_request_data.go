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

// checks if the SubscriptionOfferCodeCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCreateRequestData{}

// SubscriptionOfferCodeCreateRequestData struct for SubscriptionOfferCodeCreateRequestData
type SubscriptionOfferCodeCreateRequestData struct {
	Type string `json:"type"`
	Attributes SubscriptionOfferCodeCreateRequestDataAttributes `json:"attributes"`
	Relationships SubscriptionOfferCodeCreateRequestDataRelationships `json:"relationships"`
}

// NewSubscriptionOfferCodeCreateRequestData instantiates a new SubscriptionOfferCodeCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCreateRequestData(type_ string, attributes SubscriptionOfferCodeCreateRequestDataAttributes, relationships SubscriptionOfferCodeCreateRequestDataRelationships) *SubscriptionOfferCodeCreateRequestData {
	this := SubscriptionOfferCodeCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewSubscriptionOfferCodeCreateRequestDataWithDefaults instantiates a new SubscriptionOfferCodeCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCreateRequestDataWithDefaults() *SubscriptionOfferCodeCreateRequestData {
	this := SubscriptionOfferCodeCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionOfferCodeCreateRequestData) GetAttributes() SubscriptionOfferCodeCreateRequestDataAttributes {
	if o == nil {
		var ret SubscriptionOfferCodeCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestData) GetAttributesOk() (*SubscriptionOfferCodeCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionOfferCodeCreateRequestData) SetAttributes(v SubscriptionOfferCodeCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionOfferCodeCreateRequestData) GetRelationships() SubscriptionOfferCodeCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionOfferCodeCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestData) GetRelationshipsOk() (*SubscriptionOfferCodeCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionOfferCodeCreateRequestData) SetRelationships(v SubscriptionOfferCodeCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionOfferCodeCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCreateRequestData struct {
	value *SubscriptionOfferCodeCreateRequestData
	isSet bool
}

func (v NullableSubscriptionOfferCodeCreateRequestData) Get() *SubscriptionOfferCodeCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCreateRequestData) Set(val *SubscriptionOfferCodeCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCreateRequestData(val *SubscriptionOfferCodeCreateRequestData) *NullableSubscriptionOfferCodeCreateRequestData {
	return &NullableSubscriptionOfferCodeCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


