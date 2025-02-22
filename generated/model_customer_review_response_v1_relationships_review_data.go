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

// checks if the CustomerReviewResponseV1RelationshipsReviewData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReviewResponseV1RelationshipsReviewData{}

// CustomerReviewResponseV1RelationshipsReviewData struct for CustomerReviewResponseV1RelationshipsReviewData
type CustomerReviewResponseV1RelationshipsReviewData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewCustomerReviewResponseV1RelationshipsReviewData instantiates a new CustomerReviewResponseV1RelationshipsReviewData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReviewResponseV1RelationshipsReviewData(type_ string, id string) *CustomerReviewResponseV1RelationshipsReviewData {
	this := CustomerReviewResponseV1RelationshipsReviewData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCustomerReviewResponseV1RelationshipsReviewDataWithDefaults instantiates a new CustomerReviewResponseV1RelationshipsReviewData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewResponseV1RelationshipsReviewDataWithDefaults() *CustomerReviewResponseV1RelationshipsReviewData {
	this := CustomerReviewResponseV1RelationshipsReviewData{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerReviewResponseV1RelationshipsReviewData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1RelationshipsReviewData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerReviewResponseV1RelationshipsReviewData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerReviewResponseV1RelationshipsReviewData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1RelationshipsReviewData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerReviewResponseV1RelationshipsReviewData) SetId(v string) {
	o.Id = v
}

func (o CustomerReviewResponseV1RelationshipsReviewData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerReviewResponseV1RelationshipsReviewData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCustomerReviewResponseV1RelationshipsReviewData struct {
	value *CustomerReviewResponseV1RelationshipsReviewData
	isSet bool
}

func (v NullableCustomerReviewResponseV1RelationshipsReviewData) Get() *CustomerReviewResponseV1RelationshipsReviewData {
	return v.value
}

func (v *NullableCustomerReviewResponseV1RelationshipsReviewData) Set(val *CustomerReviewResponseV1RelationshipsReviewData) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReviewResponseV1RelationshipsReviewData) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReviewResponseV1RelationshipsReviewData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReviewResponseV1RelationshipsReviewData(val *CustomerReviewResponseV1RelationshipsReviewData) *NullableCustomerReviewResponseV1RelationshipsReviewData {
	return &NullableCustomerReviewResponseV1RelationshipsReviewData{value: val, isSet: true}
}

func (v NullableCustomerReviewResponseV1RelationshipsReviewData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerReviewResponseV1RelationshipsReviewData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


