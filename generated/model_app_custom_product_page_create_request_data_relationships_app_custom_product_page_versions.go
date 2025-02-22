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

// checks if the AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions{}

// AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions struct for AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions
type AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions struct {
	Data []AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData `json:"data,omitempty"`
}

// NewAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions instantiates a new AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions() *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions {
	this := AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions{}
	return &this
}

// NewAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersionsWithDefaults instantiates a new AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersionsWithDefaults() *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions {
	this := AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) GetData() []AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	if o == nil || IsNil(o.Data) {
		var ret []AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) GetDataOk() ([]AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData and assigns it to the Data field.
func (o *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) SetData(v []AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) {
	o.Data = v
}

func (o AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions struct {
	value *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions
	isSet bool
}

func (v NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) Get() *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions {
	return v.value
}

func (v *NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) Set(val *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions(val *AppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) *NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions {
	return &NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions{value: val, isSet: true}
}

func (v NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageCreateRequestDataRelationshipsAppCustomProductPageVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


