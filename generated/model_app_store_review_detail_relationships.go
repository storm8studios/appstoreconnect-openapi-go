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

// checks if the AppStoreReviewDetailRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetailRelationships{}

// AppStoreReviewDetailRelationships struct for AppStoreReviewDetailRelationships
type AppStoreReviewDetailRelationships struct {
	AppStoreVersion *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersion,omitempty"`
	AppStoreReviewAttachments *AppStoreReviewDetailRelationshipsAppStoreReviewAttachments `json:"appStoreReviewAttachments,omitempty"`
}

// NewAppStoreReviewDetailRelationships instantiates a new AppStoreReviewDetailRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailRelationships() *AppStoreReviewDetailRelationships {
	this := AppStoreReviewDetailRelationships{}
	return &this
}

// NewAppStoreReviewDetailRelationshipsWithDefaults instantiates a new AppStoreReviewDetailRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailRelationshipsWithDefaults() *AppStoreReviewDetailRelationships {
	this := AppStoreReviewDetailRelationships{}
	return &this
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationships) GetAppStoreVersion() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersion) {
		var ret AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersion) {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationships) HasAppStoreVersion() bool {
	if o != nil && !IsNil(o.AppStoreVersion) {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *AppStoreReviewDetailRelationships) SetAppStoreVersion(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetAppStoreReviewAttachments returns the AppStoreReviewAttachments field value if set, zero value otherwise.
func (o *AppStoreReviewDetailRelationships) GetAppStoreReviewAttachments() AppStoreReviewDetailRelationshipsAppStoreReviewAttachments {
	if o == nil || IsNil(o.AppStoreReviewAttachments) {
		var ret AppStoreReviewDetailRelationshipsAppStoreReviewAttachments
		return ret
	}
	return *o.AppStoreReviewAttachments
}

// GetAppStoreReviewAttachmentsOk returns a tuple with the AppStoreReviewAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailRelationships) GetAppStoreReviewAttachmentsOk() (*AppStoreReviewDetailRelationshipsAppStoreReviewAttachments, bool) {
	if o == nil || IsNil(o.AppStoreReviewAttachments) {
		return nil, false
	}
	return o.AppStoreReviewAttachments, true
}

// HasAppStoreReviewAttachments returns a boolean if a field has been set.
func (o *AppStoreReviewDetailRelationships) HasAppStoreReviewAttachments() bool {
	if o != nil && !IsNil(o.AppStoreReviewAttachments) {
		return true
	}

	return false
}

// SetAppStoreReviewAttachments gets a reference to the given AppStoreReviewDetailRelationshipsAppStoreReviewAttachments and assigns it to the AppStoreReviewAttachments field.
func (o *AppStoreReviewDetailRelationships) SetAppStoreReviewAttachments(v AppStoreReviewDetailRelationshipsAppStoreReviewAttachments) {
	o.AppStoreReviewAttachments = &v
}

func (o AppStoreReviewDetailRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetailRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreVersion) {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	if !IsNil(o.AppStoreReviewAttachments) {
		toSerialize["appStoreReviewAttachments"] = o.AppStoreReviewAttachments
	}
	return toSerialize, nil
}

type NullableAppStoreReviewDetailRelationships struct {
	value *AppStoreReviewDetailRelationships
	isSet bool
}

func (v NullableAppStoreReviewDetailRelationships) Get() *AppStoreReviewDetailRelationships {
	return v.value
}

func (v *NullableAppStoreReviewDetailRelationships) Set(val *AppStoreReviewDetailRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailRelationships(val *AppStoreReviewDetailRelationships) *NullableAppStoreReviewDetailRelationships {
	return &NullableAppStoreReviewDetailRelationships{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


