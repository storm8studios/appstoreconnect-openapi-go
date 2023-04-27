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

// checks if the ReviewSubmissionRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionRelationships{}

// ReviewSubmissionRelationships struct for ReviewSubmissionRelationships
type ReviewSubmissionRelationships struct {
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
	Items *ReviewSubmissionRelationshipsItems `json:"items,omitempty"`
	AppStoreVersionForReview *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersionForReview,omitempty"`
}

// NewReviewSubmissionRelationships instantiates a new ReviewSubmissionRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionRelationships() *ReviewSubmissionRelationships {
	this := ReviewSubmissionRelationships{}
	return &this
}

// NewReviewSubmissionRelationshipsWithDefaults instantiates a new ReviewSubmissionRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionRelationshipsWithDefaults() *ReviewSubmissionRelationships {
	this := ReviewSubmissionRelationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *ReviewSubmissionRelationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *ReviewSubmissionRelationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *ReviewSubmissionRelationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ReviewSubmissionRelationships) GetItems() ReviewSubmissionRelationshipsItems {
	if o == nil || IsNil(o.Items) {
		var ret ReviewSubmissionRelationshipsItems
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationships) GetItemsOk() (*ReviewSubmissionRelationshipsItems, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ReviewSubmissionRelationships) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given ReviewSubmissionRelationshipsItems and assigns it to the Items field.
func (o *ReviewSubmissionRelationships) SetItems(v ReviewSubmissionRelationshipsItems) {
	o.Items = &v
}

// GetAppStoreVersionForReview returns the AppStoreVersionForReview field value if set, zero value otherwise.
func (o *ReviewSubmissionRelationships) GetAppStoreVersionForReview() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersionForReview) {
		var ret AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersionForReview
}

// GetAppStoreVersionForReviewOk returns a tuple with the AppStoreVersionForReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionRelationships) GetAppStoreVersionForReviewOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersionForReview) {
		return nil, false
	}
	return o.AppStoreVersionForReview, true
}

// HasAppStoreVersionForReview returns a boolean if a field has been set.
func (o *ReviewSubmissionRelationships) HasAppStoreVersionForReview() bool {
	if o != nil && !IsNil(o.AppStoreVersionForReview) {
		return true
	}

	return false
}

// SetAppStoreVersionForReview gets a reference to the given AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersionForReview field.
func (o *ReviewSubmissionRelationships) SetAppStoreVersionForReview(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersionForReview = &v
}

func (o ReviewSubmissionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.AppStoreVersionForReview) {
		toSerialize["appStoreVersionForReview"] = o.AppStoreVersionForReview
	}
	return toSerialize, nil
}

type NullableReviewSubmissionRelationships struct {
	value *ReviewSubmissionRelationships
	isSet bool
}

func (v NullableReviewSubmissionRelationships) Get() *ReviewSubmissionRelationships {
	return v.value
}

func (v *NullableReviewSubmissionRelationships) Set(val *ReviewSubmissionRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionRelationships(val *ReviewSubmissionRelationships) *NullableReviewSubmissionRelationships {
	return &NullableReviewSubmissionRelationships{value: val, isSet: true}
}

func (v NullableReviewSubmissionRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


