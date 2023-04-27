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

// checks if the SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup{}

// SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup struct for SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup
type SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup struct {
	Data AppRelationshipsSubscriptionGroupsDataInner `json:"data"`
}

// NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup instantiates a new SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup(data AppRelationshipsSubscriptionGroupsDataInner) *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup {
	this := SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup{}
	this.Data = data
	return &this
}

// NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroupWithDefaults instantiates a new SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroupWithDefaults() *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup {
	this := SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) GetData() AppRelationshipsSubscriptionGroupsDataInner {
	if o == nil {
		var ret AppRelationshipsSubscriptionGroupsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) GetDataOk() (*AppRelationshipsSubscriptionGroupsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) SetData(v AppRelationshipsSubscriptionGroupsDataInner) {
	o.Data = v
}

func (o SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup struct {
	value *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) Get() *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) Set(val *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup(val *SubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) *NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup {
	return &NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


