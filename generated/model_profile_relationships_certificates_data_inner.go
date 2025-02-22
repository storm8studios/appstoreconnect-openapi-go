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

// checks if the ProfileRelationshipsCertificatesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileRelationshipsCertificatesDataInner{}

// ProfileRelationshipsCertificatesDataInner struct for ProfileRelationshipsCertificatesDataInner
type ProfileRelationshipsCertificatesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewProfileRelationshipsCertificatesDataInner instantiates a new ProfileRelationshipsCertificatesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileRelationshipsCertificatesDataInner(type_ string, id string) *ProfileRelationshipsCertificatesDataInner {
	this := ProfileRelationshipsCertificatesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewProfileRelationshipsCertificatesDataInnerWithDefaults instantiates a new ProfileRelationshipsCertificatesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileRelationshipsCertificatesDataInnerWithDefaults() *ProfileRelationshipsCertificatesDataInner {
	this := ProfileRelationshipsCertificatesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *ProfileRelationshipsCertificatesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProfileRelationshipsCertificatesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ProfileRelationshipsCertificatesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ProfileRelationshipsCertificatesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProfileRelationshipsCertificatesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProfileRelationshipsCertificatesDataInner) SetId(v string) {
	o.Id = v
}

func (o ProfileRelationshipsCertificatesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileRelationshipsCertificatesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableProfileRelationshipsCertificatesDataInner struct {
	value *ProfileRelationshipsCertificatesDataInner
	isSet bool
}

func (v NullableProfileRelationshipsCertificatesDataInner) Get() *ProfileRelationshipsCertificatesDataInner {
	return v.value
}

func (v *NullableProfileRelationshipsCertificatesDataInner) Set(val *ProfileRelationshipsCertificatesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileRelationshipsCertificatesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileRelationshipsCertificatesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileRelationshipsCertificatesDataInner(val *ProfileRelationshipsCertificatesDataInner) *NullableProfileRelationshipsCertificatesDataInner {
	return &NullableProfileRelationshipsCertificatesDataInner{value: val, isSet: true}
}

func (v NullableProfileRelationshipsCertificatesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileRelationshipsCertificatesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


