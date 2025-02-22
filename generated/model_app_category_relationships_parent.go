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

// checks if the AppCategoryRelationshipsParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCategoryRelationshipsParent{}

// AppCategoryRelationshipsParent struct for AppCategoryRelationshipsParent
type AppCategoryRelationshipsParent struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppCategoryRelationshipsSubcategoriesDataInner `json:"data,omitempty"`
}

// NewAppCategoryRelationshipsParent instantiates a new AppCategoryRelationshipsParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoryRelationshipsParent() *AppCategoryRelationshipsParent {
	this := AppCategoryRelationshipsParent{}
	return &this
}

// NewAppCategoryRelationshipsParentWithDefaults instantiates a new AppCategoryRelationshipsParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoryRelationshipsParentWithDefaults() *AppCategoryRelationshipsParent {
	this := AppCategoryRelationshipsParent{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsParent) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsParent) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsParent) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppCategoryRelationshipsParent) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCategoryRelationshipsParent) GetData() AppCategoryRelationshipsSubcategoriesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppCategoryRelationshipsSubcategoriesDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoryRelationshipsParent) GetDataOk() (*AppCategoryRelationshipsSubcategoriesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCategoryRelationshipsParent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCategoryRelationshipsSubcategoriesDataInner and assigns it to the Data field.
func (o *AppCategoryRelationshipsParent) SetData(v AppCategoryRelationshipsSubcategoriesDataInner) {
	o.Data = &v
}

func (o AppCategoryRelationshipsParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCategoryRelationshipsParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCategoryRelationshipsParent struct {
	value *AppCategoryRelationshipsParent
	isSet bool
}

func (v NullableAppCategoryRelationshipsParent) Get() *AppCategoryRelationshipsParent {
	return v.value
}

func (v *NullableAppCategoryRelationshipsParent) Set(val *AppCategoryRelationshipsParent) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoryRelationshipsParent) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoryRelationshipsParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoryRelationshipsParent(val *AppCategoryRelationshipsParent) *NullableAppCategoryRelationshipsParent {
	return &NullableAppCategoryRelationshipsParent{value: val, isSet: true}
}

func (v NullableAppCategoryRelationshipsParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoryRelationshipsParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


