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

// checks if the SubscriptionOfferCodeOneTimeUseCodeCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeOneTimeUseCodeCreateRequestData{}

// SubscriptionOfferCodeOneTimeUseCodeCreateRequestData struct for SubscriptionOfferCodeOneTimeUseCodeCreateRequestData
type SubscriptionOfferCodeOneTimeUseCodeCreateRequestData struct {
	Type string `json:"type"`
	Attributes SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes `json:"attributes"`
	Relationships SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships `json:"relationships"`
}

// NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestData instantiates a new SubscriptionOfferCodeOneTimeUseCodeCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestData(type_ string, attributes SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes, relationships SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData {
	this := SubscriptionOfferCodeOneTimeUseCodeCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataWithDefaults instantiates a new SubscriptionOfferCodeOneTimeUseCodeCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeOneTimeUseCodeCreateRequestDataWithDefaults() *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData {
	this := SubscriptionOfferCodeOneTimeUseCodeCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) GetAttributes() SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes {
	if o == nil {
		var ret SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) GetAttributesOk() (*SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) SetAttributes(v SubscriptionOfferCodeOneTimeUseCodeCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) GetRelationships() SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) GetRelationshipsOk() (*SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) SetRelationships(v SubscriptionOfferCodeCustomCodeCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData struct {
	value *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData
	isSet bool
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData) Get() *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData) Set(val *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData(val *SubscriptionOfferCodeOneTimeUseCodeCreateRequestData) *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData {
	return &NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


