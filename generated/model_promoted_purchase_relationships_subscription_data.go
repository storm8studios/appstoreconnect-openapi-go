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

// checks if the PromotedPurchaseRelationshipsSubscriptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseRelationshipsSubscriptionData{}

// PromotedPurchaseRelationshipsSubscriptionData struct for PromotedPurchaseRelationshipsSubscriptionData
type PromotedPurchaseRelationshipsSubscriptionData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewPromotedPurchaseRelationshipsSubscriptionData instantiates a new PromotedPurchaseRelationshipsSubscriptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseRelationshipsSubscriptionData(type_ string, id string) *PromotedPurchaseRelationshipsSubscriptionData {
	this := PromotedPurchaseRelationshipsSubscriptionData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewPromotedPurchaseRelationshipsSubscriptionDataWithDefaults instantiates a new PromotedPurchaseRelationshipsSubscriptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseRelationshipsSubscriptionDataWithDefaults() *PromotedPurchaseRelationshipsSubscriptionData {
	this := PromotedPurchaseRelationshipsSubscriptionData{}
	return &this
}

// GetType returns the Type field value
func (o *PromotedPurchaseRelationshipsSubscriptionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsSubscriptionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromotedPurchaseRelationshipsSubscriptionData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PromotedPurchaseRelationshipsSubscriptionData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseRelationshipsSubscriptionData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PromotedPurchaseRelationshipsSubscriptionData) SetId(v string) {
	o.Id = v
}

func (o PromotedPurchaseRelationshipsSubscriptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseRelationshipsSubscriptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullablePromotedPurchaseRelationshipsSubscriptionData struct {
	value *PromotedPurchaseRelationshipsSubscriptionData
	isSet bool
}

func (v NullablePromotedPurchaseRelationshipsSubscriptionData) Get() *PromotedPurchaseRelationshipsSubscriptionData {
	return v.value
}

func (v *NullablePromotedPurchaseRelationshipsSubscriptionData) Set(val *PromotedPurchaseRelationshipsSubscriptionData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseRelationshipsSubscriptionData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseRelationshipsSubscriptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseRelationshipsSubscriptionData(val *PromotedPurchaseRelationshipsSubscriptionData) *NullablePromotedPurchaseRelationshipsSubscriptionData {
	return &NullablePromotedPurchaseRelationshipsSubscriptionData{value: val, isSet: true}
}

func (v NullablePromotedPurchaseRelationshipsSubscriptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseRelationshipsSubscriptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


