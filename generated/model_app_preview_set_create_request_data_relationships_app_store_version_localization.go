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

// checks if the AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization{}

// AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization struct for AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization
type AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization struct {
	Data *AppPreviewSetRelationshipsAppStoreVersionLocalizationData `json:"data,omitempty"`
}

// NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization instantiates a new AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization() *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization {
	this := AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization{}
	return &this
}

// NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalizationWithDefaults instantiates a new AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalizationWithDefaults() *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization {
	this := AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) GetData() AppPreviewSetRelationshipsAppStoreVersionLocalizationData {
	if o == nil || IsNil(o.Data) {
		var ret AppPreviewSetRelationshipsAppStoreVersionLocalizationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) GetDataOk() (*AppPreviewSetRelationshipsAppStoreVersionLocalizationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPreviewSetRelationshipsAppStoreVersionLocalizationData and assigns it to the Data field.
func (o *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) SetData(v AppPreviewSetRelationshipsAppStoreVersionLocalizationData) {
	o.Data = &v
}

func (o AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization struct {
	value *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization
	isSet bool
}

func (v NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) Get() *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization {
	return v.value
}

func (v *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) Set(val *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization(val *AppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization {
	return &NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization{value: val, isSet: true}
}

func (v NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetCreateRequestDataRelationshipsAppStoreVersionLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


