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

// checks if the SubscriptionRelationshipsIntroductoryOffersDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionRelationshipsIntroductoryOffersDataInner{}

// SubscriptionRelationshipsIntroductoryOffersDataInner struct for SubscriptionRelationshipsIntroductoryOffersDataInner
type SubscriptionRelationshipsIntroductoryOffersDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewSubscriptionRelationshipsIntroductoryOffersDataInner instantiates a new SubscriptionRelationshipsIntroductoryOffersDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionRelationshipsIntroductoryOffersDataInner(type_ string, id string) *SubscriptionRelationshipsIntroductoryOffersDataInner {
	this := SubscriptionRelationshipsIntroductoryOffersDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionRelationshipsIntroductoryOffersDataInnerWithDefaults instantiates a new SubscriptionRelationshipsIntroductoryOffersDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionRelationshipsIntroductoryOffersDataInnerWithDefaults() *SubscriptionRelationshipsIntroductoryOffersDataInner {
	this := SubscriptionRelationshipsIntroductoryOffersDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionRelationshipsIntroductoryOffersDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsIntroductoryOffersDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionRelationshipsIntroductoryOffersDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionRelationshipsIntroductoryOffersDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionRelationshipsIntroductoryOffersDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionRelationshipsIntroductoryOffersDataInner) SetId(v string) {
	o.Id = v
}

func (o SubscriptionRelationshipsIntroductoryOffersDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionRelationshipsIntroductoryOffersDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSubscriptionRelationshipsIntroductoryOffersDataInner struct {
	value *SubscriptionRelationshipsIntroductoryOffersDataInner
	isSet bool
}

func (v NullableSubscriptionRelationshipsIntroductoryOffersDataInner) Get() *SubscriptionRelationshipsIntroductoryOffersDataInner {
	return v.value
}

func (v *NullableSubscriptionRelationshipsIntroductoryOffersDataInner) Set(val *SubscriptionRelationshipsIntroductoryOffersDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionRelationshipsIntroductoryOffersDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionRelationshipsIntroductoryOffersDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionRelationshipsIntroductoryOffersDataInner(val *SubscriptionRelationshipsIntroductoryOffersDataInner) *NullableSubscriptionRelationshipsIntroductoryOffersDataInner {
	return &NullableSubscriptionRelationshipsIntroductoryOffersDataInner{value: val, isSet: true}
}

func (v NullableSubscriptionRelationshipsIntroductoryOffersDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionRelationshipsIntroductoryOffersDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


