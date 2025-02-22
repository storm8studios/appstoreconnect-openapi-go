/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
	"time"
)

// checks if the AppStoreVersionCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionCreateRequestDataAttributes{}

// AppStoreVersionCreateRequestDataAttributes struct for AppStoreVersionCreateRequestDataAttributes
type AppStoreVersionCreateRequestDataAttributes struct {
	Platform Platform `json:"platform"`
	VersionString string `json:"versionString"`
	Copyright *string `json:"copyright,omitempty"`
	ReleaseType *string `json:"releaseType,omitempty"`
	EarliestReleaseDate *time.Time `json:"earliestReleaseDate,omitempty"`
}

// NewAppStoreVersionCreateRequestDataAttributes instantiates a new AppStoreVersionCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionCreateRequestDataAttributes(platform Platform, versionString string) *AppStoreVersionCreateRequestDataAttributes {
	this := AppStoreVersionCreateRequestDataAttributes{}
	this.Platform = platform
	this.VersionString = versionString
	return &this
}

// NewAppStoreVersionCreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionCreateRequestDataAttributesWithDefaults() *AppStoreVersionCreateRequestDataAttributes {
	this := AppStoreVersionCreateRequestDataAttributes{}
	return &this
}

// GetPlatform returns the Platform field value
func (o *AppStoreVersionCreateRequestDataAttributes) GetPlatform() Platform {
	if o == nil {
		var ret Platform
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) GetPlatformOk() (*Platform, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *AppStoreVersionCreateRequestDataAttributes) SetPlatform(v Platform) {
	o.Platform = v
}

// GetVersionString returns the VersionString field value
func (o *AppStoreVersionCreateRequestDataAttributes) GetVersionString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionString
}

// GetVersionStringOk returns a tuple with the VersionString field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) GetVersionStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionString, true
}

// SetVersionString sets field value
func (o *AppStoreVersionCreateRequestDataAttributes) SetVersionString(v string) {
	o.VersionString = v
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *AppStoreVersionCreateRequestDataAttributes) GetCopyright() string {
	if o == nil || IsNil(o.Copyright) {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) GetCopyrightOk() (*string, bool) {
	if o == nil || IsNil(o.Copyright) {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) HasCopyright() bool {
	if o != nil && !IsNil(o.Copyright) {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *AppStoreVersionCreateRequestDataAttributes) SetCopyright(v string) {
	o.Copyright = &v
}

// GetReleaseType returns the ReleaseType field value if set, zero value otherwise.
func (o *AppStoreVersionCreateRequestDataAttributes) GetReleaseType() string {
	if o == nil || IsNil(o.ReleaseType) {
		var ret string
		return ret
	}
	return *o.ReleaseType
}

// GetReleaseTypeOk returns a tuple with the ReleaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) GetReleaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseType) {
		return nil, false
	}
	return o.ReleaseType, true
}

// HasReleaseType returns a boolean if a field has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) HasReleaseType() bool {
	if o != nil && !IsNil(o.ReleaseType) {
		return true
	}

	return false
}

// SetReleaseType gets a reference to the given string and assigns it to the ReleaseType field.
func (o *AppStoreVersionCreateRequestDataAttributes) SetReleaseType(v string) {
	o.ReleaseType = &v
}

// GetEarliestReleaseDate returns the EarliestReleaseDate field value if set, zero value otherwise.
func (o *AppStoreVersionCreateRequestDataAttributes) GetEarliestReleaseDate() time.Time {
	if o == nil || IsNil(o.EarliestReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.EarliestReleaseDate
}

// GetEarliestReleaseDateOk returns a tuple with the EarliestReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) GetEarliestReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EarliestReleaseDate) {
		return nil, false
	}
	return o.EarliestReleaseDate, true
}

// HasEarliestReleaseDate returns a boolean if a field has been set.
func (o *AppStoreVersionCreateRequestDataAttributes) HasEarliestReleaseDate() bool {
	if o != nil && !IsNil(o.EarliestReleaseDate) {
		return true
	}

	return false
}

// SetEarliestReleaseDate gets a reference to the given time.Time and assigns it to the EarliestReleaseDate field.
func (o *AppStoreVersionCreateRequestDataAttributes) SetEarliestReleaseDate(v time.Time) {
	o.EarliestReleaseDate = &v
}

func (o AppStoreVersionCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["platform"] = o.Platform
	toSerialize["versionString"] = o.VersionString
	if !IsNil(o.Copyright) {
		toSerialize["copyright"] = o.Copyright
	}
	if !IsNil(o.ReleaseType) {
		toSerialize["releaseType"] = o.ReleaseType
	}
	if !IsNil(o.EarliestReleaseDate) {
		toSerialize["earliestReleaseDate"] = o.EarliestReleaseDate
	}
	return toSerialize, nil
}

type NullableAppStoreVersionCreateRequestDataAttributes struct {
	value *AppStoreVersionCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppStoreVersionCreateRequestDataAttributes) Get() *AppStoreVersionCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppStoreVersionCreateRequestDataAttributes) Set(val *AppStoreVersionCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionCreateRequestDataAttributes(val *AppStoreVersionCreateRequestDataAttributes) *NullableAppStoreVersionCreateRequestDataAttributes {
	return &NullableAppStoreVersionCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


