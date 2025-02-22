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

// checks if the CiBuildRunRelationshipsPullRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunRelationshipsPullRequestData{}

// CiBuildRunRelationshipsPullRequestData struct for CiBuildRunRelationshipsPullRequestData
type CiBuildRunRelationshipsPullRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewCiBuildRunRelationshipsPullRequestData instantiates a new CiBuildRunRelationshipsPullRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunRelationshipsPullRequestData(type_ string, id string) *CiBuildRunRelationshipsPullRequestData {
	this := CiBuildRunRelationshipsPullRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiBuildRunRelationshipsPullRequestDataWithDefaults instantiates a new CiBuildRunRelationshipsPullRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunRelationshipsPullRequestDataWithDefaults() *CiBuildRunRelationshipsPullRequestData {
	this := CiBuildRunRelationshipsPullRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *CiBuildRunRelationshipsPullRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationshipsPullRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiBuildRunRelationshipsPullRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiBuildRunRelationshipsPullRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiBuildRunRelationshipsPullRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiBuildRunRelationshipsPullRequestData) SetId(v string) {
	o.Id = v
}

func (o CiBuildRunRelationshipsPullRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunRelationshipsPullRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCiBuildRunRelationshipsPullRequestData struct {
	value *CiBuildRunRelationshipsPullRequestData
	isSet bool
}

func (v NullableCiBuildRunRelationshipsPullRequestData) Get() *CiBuildRunRelationshipsPullRequestData {
	return v.value
}

func (v *NullableCiBuildRunRelationshipsPullRequestData) Set(val *CiBuildRunRelationshipsPullRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunRelationshipsPullRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunRelationshipsPullRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunRelationshipsPullRequestData(val *CiBuildRunRelationshipsPullRequestData) *NullableCiBuildRunRelationshipsPullRequestData {
	return &NullableCiBuildRunRelationshipsPullRequestData{value: val, isSet: true}
}

func (v NullableCiBuildRunRelationshipsPullRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunRelationshipsPullRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


