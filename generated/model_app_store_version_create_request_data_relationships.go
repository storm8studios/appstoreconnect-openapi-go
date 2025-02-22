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

// checks if the AppStoreVersionCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionCreateRequestDataRelationships{}

// AppStoreVersionCreateRequestDataRelationships struct for AppStoreVersionCreateRequestDataRelationships
type AppStoreVersionCreateRequestDataRelationships struct {
	App AppAvailabilityCreateRequestDataRelationshipsApp `json:"app"`
	AppStoreVersionLocalizations *AppStoreVersionCreateRequestDataRelationshipsAppStoreVersionLocalizations `json:"appStoreVersionLocalizations,omitempty"`
	Build *AppStoreVersionCreateRequestDataRelationshipsBuild `json:"build,omitempty"`
}

// NewAppStoreVersionCreateRequestDataRelationships instantiates a new AppStoreVersionCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionCreateRequestDataRelationships(app AppAvailabilityCreateRequestDataRelationshipsApp) *AppStoreVersionCreateRequestDataRelationships {
	this := AppStoreVersionCreateRequestDataRelationships{}
	this.App = app
	return &this
}

// NewAppStoreVersionCreateRequestDataRelationshipsWithDefaults instantiates a new AppStoreVersionCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionCreateRequestDataRelationshipsWithDefaults() *AppStoreVersionCreateRequestDataRelationships {
	this := AppStoreVersionCreateRequestDataRelationships{}
	return &this
}

// GetApp returns the App field value
func (o *AppStoreVersionCreateRequestDataRelationships) GetApp() AppAvailabilityCreateRequestDataRelationshipsApp {
	if o == nil {
		var ret AppAvailabilityCreateRequestDataRelationshipsApp
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataRelationships) GetAppOk() (*AppAvailabilityCreateRequestDataRelationshipsApp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *AppStoreVersionCreateRequestDataRelationships) SetApp(v AppAvailabilityCreateRequestDataRelationshipsApp) {
	o.App = v
}

// GetAppStoreVersionLocalizations returns the AppStoreVersionLocalizations field value if set, zero value otherwise.
func (o *AppStoreVersionCreateRequestDataRelationships) GetAppStoreVersionLocalizations() AppStoreVersionCreateRequestDataRelationshipsAppStoreVersionLocalizations {
	if o == nil || IsNil(o.AppStoreVersionLocalizations) {
		var ret AppStoreVersionCreateRequestDataRelationshipsAppStoreVersionLocalizations
		return ret
	}
	return *o.AppStoreVersionLocalizations
}

// GetAppStoreVersionLocalizationsOk returns a tuple with the AppStoreVersionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataRelationships) GetAppStoreVersionLocalizationsOk() (*AppStoreVersionCreateRequestDataRelationshipsAppStoreVersionLocalizations, bool) {
	if o == nil || IsNil(o.AppStoreVersionLocalizations) {
		return nil, false
	}
	return o.AppStoreVersionLocalizations, true
}

// HasAppStoreVersionLocalizations returns a boolean if a field has been set.
func (o *AppStoreVersionCreateRequestDataRelationships) HasAppStoreVersionLocalizations() bool {
	if o != nil && !IsNil(o.AppStoreVersionLocalizations) {
		return true
	}

	return false
}

// SetAppStoreVersionLocalizations gets a reference to the given AppStoreVersionCreateRequestDataRelationshipsAppStoreVersionLocalizations and assigns it to the AppStoreVersionLocalizations field.
func (o *AppStoreVersionCreateRequestDataRelationships) SetAppStoreVersionLocalizations(v AppStoreVersionCreateRequestDataRelationshipsAppStoreVersionLocalizations) {
	o.AppStoreVersionLocalizations = &v
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *AppStoreVersionCreateRequestDataRelationships) GetBuild() AppStoreVersionCreateRequestDataRelationshipsBuild {
	if o == nil || IsNil(o.Build) {
		var ret AppStoreVersionCreateRequestDataRelationshipsBuild
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestDataRelationships) GetBuildOk() (*AppStoreVersionCreateRequestDataRelationshipsBuild, bool) {
	if o == nil || IsNil(o.Build) {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *AppStoreVersionCreateRequestDataRelationships) HasBuild() bool {
	if o != nil && !IsNil(o.Build) {
		return true
	}

	return false
}

// SetBuild gets a reference to the given AppStoreVersionCreateRequestDataRelationshipsBuild and assigns it to the Build field.
func (o *AppStoreVersionCreateRequestDataRelationships) SetBuild(v AppStoreVersionCreateRequestDataRelationshipsBuild) {
	o.Build = &v
}

func (o AppStoreVersionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	if !IsNil(o.AppStoreVersionLocalizations) {
		toSerialize["appStoreVersionLocalizations"] = o.AppStoreVersionLocalizations
	}
	if !IsNil(o.Build) {
		toSerialize["build"] = o.Build
	}
	return toSerialize, nil
}

type NullableAppStoreVersionCreateRequestDataRelationships struct {
	value *AppStoreVersionCreateRequestDataRelationships
	isSet bool
}

func (v NullableAppStoreVersionCreateRequestDataRelationships) Get() *AppStoreVersionCreateRequestDataRelationships {
	return v.value
}

func (v *NullableAppStoreVersionCreateRequestDataRelationships) Set(val *AppStoreVersionCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionCreateRequestDataRelationships(val *AppStoreVersionCreateRequestDataRelationships) *NullableAppStoreVersionCreateRequestDataRelationships {
	return &NullableAppStoreVersionCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


