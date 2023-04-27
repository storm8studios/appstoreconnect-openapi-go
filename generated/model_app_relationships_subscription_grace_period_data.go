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

// checks if the AppRelationshipsSubscriptionGracePeriodData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsSubscriptionGracePeriodData{}

// AppRelationshipsSubscriptionGracePeriodData struct for AppRelationshipsSubscriptionGracePeriodData
type AppRelationshipsSubscriptionGracePeriodData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppRelationshipsSubscriptionGracePeriodData instantiates a new AppRelationshipsSubscriptionGracePeriodData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsSubscriptionGracePeriodData(type_ string, id string) *AppRelationshipsSubscriptionGracePeriodData {
	this := AppRelationshipsSubscriptionGracePeriodData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsSubscriptionGracePeriodDataWithDefaults instantiates a new AppRelationshipsSubscriptionGracePeriodData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsSubscriptionGracePeriodDataWithDefaults() *AppRelationshipsSubscriptionGracePeriodData {
	this := AppRelationshipsSubscriptionGracePeriodData{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsSubscriptionGracePeriodData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsSubscriptionGracePeriodData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsSubscriptionGracePeriodData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsSubscriptionGracePeriodData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsSubscriptionGracePeriodData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsSubscriptionGracePeriodData) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsSubscriptionGracePeriodData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsSubscriptionGracePeriodData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppRelationshipsSubscriptionGracePeriodData struct {
	value *AppRelationshipsSubscriptionGracePeriodData
	isSet bool
}

func (v NullableAppRelationshipsSubscriptionGracePeriodData) Get() *AppRelationshipsSubscriptionGracePeriodData {
	return v.value
}

func (v *NullableAppRelationshipsSubscriptionGracePeriodData) Set(val *AppRelationshipsSubscriptionGracePeriodData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsSubscriptionGracePeriodData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsSubscriptionGracePeriodData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsSubscriptionGracePeriodData(val *AppRelationshipsSubscriptionGracePeriodData) *NullableAppRelationshipsSubscriptionGracePeriodData {
	return &NullableAppRelationshipsSubscriptionGracePeriodData{value: val, isSet: true}
}

func (v NullableAppRelationshipsSubscriptionGracePeriodData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsSubscriptionGracePeriodData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


