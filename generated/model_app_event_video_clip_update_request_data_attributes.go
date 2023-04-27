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

// checks if the AppEventVideoClipUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventVideoClipUpdateRequestDataAttributes{}

// AppEventVideoClipUpdateRequestDataAttributes struct for AppEventVideoClipUpdateRequestDataAttributes
type AppEventVideoClipUpdateRequestDataAttributes struct {
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	Uploaded *bool `json:"uploaded,omitempty"`
}

// NewAppEventVideoClipUpdateRequestDataAttributes instantiates a new AppEventVideoClipUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventVideoClipUpdateRequestDataAttributes() *AppEventVideoClipUpdateRequestDataAttributes {
	this := AppEventVideoClipUpdateRequestDataAttributes{}
	return &this
}

// NewAppEventVideoClipUpdateRequestDataAttributesWithDefaults instantiates a new AppEventVideoClipUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventVideoClipUpdateRequestDataAttributesWithDefaults() *AppEventVideoClipUpdateRequestDataAttributes {
	this := AppEventVideoClipUpdateRequestDataAttributes{}
	return &this
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppEventVideoClipUpdateRequestDataAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipUpdateRequestDataAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppEventVideoClipUpdateRequestDataAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && !IsNil(o.PreviewFrameTimeCode) {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppEventVideoClipUpdateRequestDataAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetUploaded returns the Uploaded field value if set, zero value otherwise.
func (o *AppEventVideoClipUpdateRequestDataAttributes) GetUploaded() bool {
	if o == nil || IsNil(o.Uploaded) {
		var ret bool
		return ret
	}
	return *o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipUpdateRequestDataAttributes) GetUploadedOk() (*bool, bool) {
	if o == nil || IsNil(o.Uploaded) {
		return nil, false
	}
	return o.Uploaded, true
}

// HasUploaded returns a boolean if a field has been set.
func (o *AppEventVideoClipUpdateRequestDataAttributes) HasUploaded() bool {
	if o != nil && !IsNil(o.Uploaded) {
		return true
	}

	return false
}

// SetUploaded gets a reference to the given bool and assigns it to the Uploaded field.
func (o *AppEventVideoClipUpdateRequestDataAttributes) SetUploaded(v bool) {
	o.Uploaded = &v
}

func (o AppEventVideoClipUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventVideoClipUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreviewFrameTimeCode) {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	if !IsNil(o.Uploaded) {
		toSerialize["uploaded"] = o.Uploaded
	}
	return toSerialize, nil
}

type NullableAppEventVideoClipUpdateRequestDataAttributes struct {
	value *AppEventVideoClipUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppEventVideoClipUpdateRequestDataAttributes) Get() *AppEventVideoClipUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEventVideoClipUpdateRequestDataAttributes) Set(val *AppEventVideoClipUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventVideoClipUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventVideoClipUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventVideoClipUpdateRequestDataAttributes(val *AppEventVideoClipUpdateRequestDataAttributes) *NullableAppEventVideoClipUpdateRequestDataAttributes {
	return &NullableAppEventVideoClipUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEventVideoClipUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventVideoClipUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


