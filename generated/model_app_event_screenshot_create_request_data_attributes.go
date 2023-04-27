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

// checks if the AppEventScreenshotCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotCreateRequestDataAttributes{}

// AppEventScreenshotCreateRequestDataAttributes struct for AppEventScreenshotCreateRequestDataAttributes
type AppEventScreenshotCreateRequestDataAttributes struct {
	FileSize int32 `json:"fileSize"`
	FileName string `json:"fileName"`
	AppEventAssetType AppEventAssetType `json:"appEventAssetType"`
}

// NewAppEventScreenshotCreateRequestDataAttributes instantiates a new AppEventScreenshotCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotCreateRequestDataAttributes(fileSize int32, fileName string, appEventAssetType AppEventAssetType) *AppEventScreenshotCreateRequestDataAttributes {
	this := AppEventScreenshotCreateRequestDataAttributes{}
	this.FileSize = fileSize
	this.FileName = fileName
	this.AppEventAssetType = appEventAssetType
	return &this
}

// NewAppEventScreenshotCreateRequestDataAttributesWithDefaults instantiates a new AppEventScreenshotCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotCreateRequestDataAttributesWithDefaults() *AppEventScreenshotCreateRequestDataAttributes {
	this := AppEventScreenshotCreateRequestDataAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value
func (o *AppEventScreenshotCreateRequestDataAttributes) GetFileSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotCreateRequestDataAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *AppEventScreenshotCreateRequestDataAttributes) SetFileSize(v int32) {
	o.FileSize = v
}

// GetFileName returns the FileName field value
func (o *AppEventScreenshotCreateRequestDataAttributes) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotCreateRequestDataAttributes) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *AppEventScreenshotCreateRequestDataAttributes) SetFileName(v string) {
	o.FileName = v
}

// GetAppEventAssetType returns the AppEventAssetType field value
func (o *AppEventScreenshotCreateRequestDataAttributes) GetAppEventAssetType() AppEventAssetType {
	if o == nil {
		var ret AppEventAssetType
		return ret
	}

	return o.AppEventAssetType
}

// GetAppEventAssetTypeOk returns a tuple with the AppEventAssetType field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotCreateRequestDataAttributes) GetAppEventAssetTypeOk() (*AppEventAssetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppEventAssetType, true
}

// SetAppEventAssetType sets field value
func (o *AppEventScreenshotCreateRequestDataAttributes) SetAppEventAssetType(v AppEventAssetType) {
	o.AppEventAssetType = v
}

func (o AppEventScreenshotCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileSize"] = o.FileSize
	toSerialize["fileName"] = o.FileName
	toSerialize["appEventAssetType"] = o.AppEventAssetType
	return toSerialize, nil
}

type NullableAppEventScreenshotCreateRequestDataAttributes struct {
	value *AppEventScreenshotCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppEventScreenshotCreateRequestDataAttributes) Get() *AppEventScreenshotCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEventScreenshotCreateRequestDataAttributes) Set(val *AppEventScreenshotCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotCreateRequestDataAttributes(val *AppEventScreenshotCreateRequestDataAttributes) *NullableAppEventScreenshotCreateRequestDataAttributes {
	return &NullableAppEventScreenshotCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEventScreenshotCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


