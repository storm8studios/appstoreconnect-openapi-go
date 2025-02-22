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

// checks if the AppClipDefaultExperienceCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDefaultExperienceCreateRequestDataRelationships{}

// AppClipDefaultExperienceCreateRequestDataRelationships struct for AppClipDefaultExperienceCreateRequestDataRelationships
type AppClipDefaultExperienceCreateRequestDataRelationships struct {
	AppClip AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip `json:"appClip"`
	ReleaseWithAppStoreVersion *AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion `json:"releaseWithAppStoreVersion,omitempty"`
	AppClipDefaultExperienceTemplate *AppClipDefaultExperienceCreateRequestDataRelationshipsAppClipDefaultExperienceTemplate `json:"appClipDefaultExperienceTemplate,omitempty"`
}

// NewAppClipDefaultExperienceCreateRequestDataRelationships instantiates a new AppClipDefaultExperienceCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDefaultExperienceCreateRequestDataRelationships(appClip AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) *AppClipDefaultExperienceCreateRequestDataRelationships {
	this := AppClipDefaultExperienceCreateRequestDataRelationships{}
	this.AppClip = appClip
	return &this
}

// NewAppClipDefaultExperienceCreateRequestDataRelationshipsWithDefaults instantiates a new AppClipDefaultExperienceCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDefaultExperienceCreateRequestDataRelationshipsWithDefaults() *AppClipDefaultExperienceCreateRequestDataRelationships {
	this := AppClipDefaultExperienceCreateRequestDataRelationships{}
	return &this
}

// GetAppClip returns the AppClip field value
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) GetAppClip() AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip {
	if o == nil {
		var ret AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip
		return ret
	}

	return o.AppClip
}

// GetAppClipOk returns a tuple with the AppClip field value
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) GetAppClipOk() (*AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppClip, true
}

// SetAppClip sets field value
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) SetAppClip(v AppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip) {
	o.AppClip = v
}

// GetReleaseWithAppStoreVersion returns the ReleaseWithAppStoreVersion field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) GetReleaseWithAppStoreVersion() AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.ReleaseWithAppStoreVersion) {
		var ret AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.ReleaseWithAppStoreVersion
}

// GetReleaseWithAppStoreVersionOk returns a tuple with the ReleaseWithAppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) GetReleaseWithAppStoreVersionOk() (*AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.ReleaseWithAppStoreVersion) {
		return nil, false
	}
	return o.ReleaseWithAppStoreVersion, true
}

// HasReleaseWithAppStoreVersion returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) HasReleaseWithAppStoreVersion() bool {
	if o != nil && !IsNil(o.ReleaseWithAppStoreVersion) {
		return true
	}

	return false
}

// SetReleaseWithAppStoreVersion gets a reference to the given AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion and assigns it to the ReleaseWithAppStoreVersion field.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) SetReleaseWithAppStoreVersion(v AppClipDefaultExperienceCreateRequestDataRelationshipsReleaseWithAppStoreVersion) {
	o.ReleaseWithAppStoreVersion = &v
}

// GetAppClipDefaultExperienceTemplate returns the AppClipDefaultExperienceTemplate field value if set, zero value otherwise.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) GetAppClipDefaultExperienceTemplate() AppClipDefaultExperienceCreateRequestDataRelationshipsAppClipDefaultExperienceTemplate {
	if o == nil || IsNil(o.AppClipDefaultExperienceTemplate) {
		var ret AppClipDefaultExperienceCreateRequestDataRelationshipsAppClipDefaultExperienceTemplate
		return ret
	}
	return *o.AppClipDefaultExperienceTemplate
}

// GetAppClipDefaultExperienceTemplateOk returns a tuple with the AppClipDefaultExperienceTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) GetAppClipDefaultExperienceTemplateOk() (*AppClipDefaultExperienceCreateRequestDataRelationshipsAppClipDefaultExperienceTemplate, bool) {
	if o == nil || IsNil(o.AppClipDefaultExperienceTemplate) {
		return nil, false
	}
	return o.AppClipDefaultExperienceTemplate, true
}

// HasAppClipDefaultExperienceTemplate returns a boolean if a field has been set.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) HasAppClipDefaultExperienceTemplate() bool {
	if o != nil && !IsNil(o.AppClipDefaultExperienceTemplate) {
		return true
	}

	return false
}

// SetAppClipDefaultExperienceTemplate gets a reference to the given AppClipDefaultExperienceCreateRequestDataRelationshipsAppClipDefaultExperienceTemplate and assigns it to the AppClipDefaultExperienceTemplate field.
func (o *AppClipDefaultExperienceCreateRequestDataRelationships) SetAppClipDefaultExperienceTemplate(v AppClipDefaultExperienceCreateRequestDataRelationshipsAppClipDefaultExperienceTemplate) {
	o.AppClipDefaultExperienceTemplate = &v
}

func (o AppClipDefaultExperienceCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDefaultExperienceCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appClip"] = o.AppClip
	if !IsNil(o.ReleaseWithAppStoreVersion) {
		toSerialize["releaseWithAppStoreVersion"] = o.ReleaseWithAppStoreVersion
	}
	if !IsNil(o.AppClipDefaultExperienceTemplate) {
		toSerialize["appClipDefaultExperienceTemplate"] = o.AppClipDefaultExperienceTemplate
	}
	return toSerialize, nil
}

type NullableAppClipDefaultExperienceCreateRequestDataRelationships struct {
	value *AppClipDefaultExperienceCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppClipDefaultExperienceCreateRequestDataRelationships) Get() *AppClipDefaultExperienceCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppClipDefaultExperienceCreateRequestDataRelationships) Set(val *AppClipDefaultExperienceCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperienceCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperienceCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperienceCreateRequestDataRelationships(val *AppClipDefaultExperienceCreateRequestDataRelationships) *NullableAppClipDefaultExperienceCreateRequestDataRelationships {
	return &NullableAppClipDefaultExperienceCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperienceCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperienceCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


