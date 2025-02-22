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

// checks if the CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag{}

// CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag struct for CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag
type CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag struct {
	Data *CiBuildRunRelationshipsSourceBranchOrTagData `json:"data,omitempty"`
}

// NewCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag instantiates a new CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag() *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag {
	this := CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag{}
	return &this
}

// NewCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTagWithDefaults instantiates a new CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTagWithDefaults() *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag {
	this := CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) GetData() CiBuildRunRelationshipsSourceBranchOrTagData {
	if o == nil || IsNil(o.Data) {
		var ret CiBuildRunRelationshipsSourceBranchOrTagData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) GetDataOk() (*CiBuildRunRelationshipsSourceBranchOrTagData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CiBuildRunRelationshipsSourceBranchOrTagData and assigns it to the Data field.
func (o *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) SetData(v CiBuildRunRelationshipsSourceBranchOrTagData) {
	o.Data = &v
}

func (o CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag struct {
	value *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag
	isSet bool
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) Get() *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag {
	return v.value
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) Set(val *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag(val *CiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) *NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag {
	return &NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag{value: val, isSet: true}
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsSourceBranchOrTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


