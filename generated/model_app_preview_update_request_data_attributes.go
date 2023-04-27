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

// checks if the AppPreviewUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewUpdateRequestDataAttributes{}

// AppPreviewUpdateRequestDataAttributes struct for AppPreviewUpdateRequestDataAttributes
type AppPreviewUpdateRequestDataAttributes struct {
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	Uploaded *bool `json:"uploaded,omitempty"`
}

// NewAppPreviewUpdateRequestDataAttributes instantiates a new AppPreviewUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewUpdateRequestDataAttributes() *AppPreviewUpdateRequestDataAttributes {
	this := AppPreviewUpdateRequestDataAttributes{}
	return &this
}

// NewAppPreviewUpdateRequestDataAttributesWithDefaults instantiates a new AppPreviewUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewUpdateRequestDataAttributesWithDefaults() *AppPreviewUpdateRequestDataAttributes {
	this := AppPreviewUpdateRequestDataAttributes{}
	return &this
}

// GetSourceFileChecksum returns the SourceFileChecksum field value if set, zero value otherwise.
func (o *AppPreviewUpdateRequestDataAttributes) GetSourceFileChecksum() string {
	if o == nil || IsNil(o.SourceFileChecksum) {
		var ret string
		return ret
	}
	return *o.SourceFileChecksum
}

// GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequestDataAttributes) GetSourceFileChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.SourceFileChecksum) {
		return nil, false
	}
	return o.SourceFileChecksum, true
}

// HasSourceFileChecksum returns a boolean if a field has been set.
func (o *AppPreviewUpdateRequestDataAttributes) HasSourceFileChecksum() bool {
	if o != nil && !IsNil(o.SourceFileChecksum) {
		return true
	}

	return false
}

// SetSourceFileChecksum gets a reference to the given string and assigns it to the SourceFileChecksum field.
func (o *AppPreviewUpdateRequestDataAttributes) SetSourceFileChecksum(v string) {
	o.SourceFileChecksum = &v
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppPreviewUpdateRequestDataAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequestDataAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppPreviewUpdateRequestDataAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && !IsNil(o.PreviewFrameTimeCode) {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppPreviewUpdateRequestDataAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetUploaded returns the Uploaded field value if set, zero value otherwise.
func (o *AppPreviewUpdateRequestDataAttributes) GetUploaded() bool {
	if o == nil || IsNil(o.Uploaded) {
		var ret bool
		return ret
	}
	return *o.Uploaded
}

// GetUploadedOk returns a tuple with the Uploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewUpdateRequestDataAttributes) GetUploadedOk() (*bool, bool) {
	if o == nil || IsNil(o.Uploaded) {
		return nil, false
	}
	return o.Uploaded, true
}

// HasUploaded returns a boolean if a field has been set.
func (o *AppPreviewUpdateRequestDataAttributes) HasUploaded() bool {
	if o != nil && !IsNil(o.Uploaded) {
		return true
	}

	return false
}

// SetUploaded gets a reference to the given bool and assigns it to the Uploaded field.
func (o *AppPreviewUpdateRequestDataAttributes) SetUploaded(v bool) {
	o.Uploaded = &v
}

func (o AppPreviewUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceFileChecksum) {
		toSerialize["sourceFileChecksum"] = o.SourceFileChecksum
	}
	if !IsNil(o.PreviewFrameTimeCode) {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	if !IsNil(o.Uploaded) {
		toSerialize["uploaded"] = o.Uploaded
	}
	return toSerialize, nil
}

type NullableAppPreviewUpdateRequestDataAttributes struct {
	value *AppPreviewUpdateRequestDataAttributes
	isSet bool
}

func (v NullableAppPreviewUpdateRequestDataAttributes) Get() *AppPreviewUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableAppPreviewUpdateRequestDataAttributes) Set(val *AppPreviewUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewUpdateRequestDataAttributes(val *AppPreviewUpdateRequestDataAttributes) *NullableAppPreviewUpdateRequestDataAttributes {
	return &NullableAppPreviewUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppPreviewUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


