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

// checks if the AppStoreVersionExperimentTreatmentAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionExperimentTreatmentAttributes{}

// AppStoreVersionExperimentTreatmentAttributes struct for AppStoreVersionExperimentTreatmentAttributes
type AppStoreVersionExperimentTreatmentAttributes struct {
	Name *string `json:"name,omitempty"`
	AppIcon *ImageAsset `json:"appIcon,omitempty"`
	AppIconName *string `json:"appIconName,omitempty"`
	PromotedDate *time.Time `json:"promotedDate,omitempty"`
}

// NewAppStoreVersionExperimentTreatmentAttributes instantiates a new AppStoreVersionExperimentTreatmentAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionExperimentTreatmentAttributes() *AppStoreVersionExperimentTreatmentAttributes {
	this := AppStoreVersionExperimentTreatmentAttributes{}
	return &this
}

// NewAppStoreVersionExperimentTreatmentAttributesWithDefaults instantiates a new AppStoreVersionExperimentTreatmentAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionExperimentTreatmentAttributesWithDefaults() *AppStoreVersionExperimentTreatmentAttributes {
	this := AppStoreVersionExperimentTreatmentAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppStoreVersionExperimentTreatmentAttributes) SetName(v string) {
	o.Name = &v
}

// GetAppIcon returns the AppIcon field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetAppIcon() ImageAsset {
	if o == nil || IsNil(o.AppIcon) {
		var ret ImageAsset
		return ret
	}
	return *o.AppIcon
}

// GetAppIconOk returns a tuple with the AppIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetAppIconOk() (*ImageAsset, bool) {
	if o == nil || IsNil(o.AppIcon) {
		return nil, false
	}
	return o.AppIcon, true
}

// HasAppIcon returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) HasAppIcon() bool {
	if o != nil && !IsNil(o.AppIcon) {
		return true
	}

	return false
}

// SetAppIcon gets a reference to the given ImageAsset and assigns it to the AppIcon field.
func (o *AppStoreVersionExperimentTreatmentAttributes) SetAppIcon(v ImageAsset) {
	o.AppIcon = &v
}

// GetAppIconName returns the AppIconName field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetAppIconName() string {
	if o == nil || IsNil(o.AppIconName) {
		var ret string
		return ret
	}
	return *o.AppIconName
}

// GetAppIconNameOk returns a tuple with the AppIconName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetAppIconNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppIconName) {
		return nil, false
	}
	return o.AppIconName, true
}

// HasAppIconName returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) HasAppIconName() bool {
	if o != nil && !IsNil(o.AppIconName) {
		return true
	}

	return false
}

// SetAppIconName gets a reference to the given string and assigns it to the AppIconName field.
func (o *AppStoreVersionExperimentTreatmentAttributes) SetAppIconName(v string) {
	o.AppIconName = &v
}

// GetPromotedDate returns the PromotedDate field value if set, zero value otherwise.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetPromotedDate() time.Time {
	if o == nil || IsNil(o.PromotedDate) {
		var ret time.Time
		return ret
	}
	return *o.PromotedDate
}

// GetPromotedDateOk returns a tuple with the PromotedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) GetPromotedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PromotedDate) {
		return nil, false
	}
	return o.PromotedDate, true
}

// HasPromotedDate returns a boolean if a field has been set.
func (o *AppStoreVersionExperimentTreatmentAttributes) HasPromotedDate() bool {
	if o != nil && !IsNil(o.PromotedDate) {
		return true
	}

	return false
}

// SetPromotedDate gets a reference to the given time.Time and assigns it to the PromotedDate field.
func (o *AppStoreVersionExperimentTreatmentAttributes) SetPromotedDate(v time.Time) {
	o.PromotedDate = &v
}

func (o AppStoreVersionExperimentTreatmentAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionExperimentTreatmentAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AppIcon) {
		toSerialize["appIcon"] = o.AppIcon
	}
	if !IsNil(o.AppIconName) {
		toSerialize["appIconName"] = o.AppIconName
	}
	if !IsNil(o.PromotedDate) {
		toSerialize["promotedDate"] = o.PromotedDate
	}
	return toSerialize, nil
}

type NullableAppStoreVersionExperimentTreatmentAttributes struct {
	value *AppStoreVersionExperimentTreatmentAttributes
	isSet bool
}

func (v NullableAppStoreVersionExperimentTreatmentAttributes) Get() *AppStoreVersionExperimentTreatmentAttributes {
	return v.value
}

func (v *NullableAppStoreVersionExperimentTreatmentAttributes) Set(val *AppStoreVersionExperimentTreatmentAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentTreatmentAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentTreatmentAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentTreatmentAttributes(val *AppStoreVersionExperimentTreatmentAttributes) *NullableAppStoreVersionExperimentTreatmentAttributes {
	return &NullableAppStoreVersionExperimentTreatmentAttributes{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentTreatmentAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentTreatmentAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


