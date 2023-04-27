/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
	"fmt"
)

// AppPreviewSetsResponseIncludedInner - struct for AppPreviewSetsResponseIncludedInner
type AppPreviewSetsResponseIncludedInner struct {
	AppCustomProductPageLocalization *AppCustomProductPageLocalization
	AppPreview *AppPreview
	AppStoreVersionExperimentTreatmentLocalization *AppStoreVersionExperimentTreatmentLocalization
	AppStoreVersionLocalization *AppStoreVersionLocalization
}

// AppCustomProductPageLocalizationAsAppPreviewSetsResponseIncludedInner is a convenience function that returns AppCustomProductPageLocalization wrapped in AppPreviewSetsResponseIncludedInner
func AppCustomProductPageLocalizationAsAppPreviewSetsResponseIncludedInner(v *AppCustomProductPageLocalization) AppPreviewSetsResponseIncludedInner {
	return AppPreviewSetsResponseIncludedInner{
		AppCustomProductPageLocalization: v,
	}
}

// AppPreviewAsAppPreviewSetsResponseIncludedInner is a convenience function that returns AppPreview wrapped in AppPreviewSetsResponseIncludedInner
func AppPreviewAsAppPreviewSetsResponseIncludedInner(v *AppPreview) AppPreviewSetsResponseIncludedInner {
	return AppPreviewSetsResponseIncludedInner{
		AppPreview: v,
	}
}

// AppStoreVersionExperimentTreatmentLocalizationAsAppPreviewSetsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentTreatmentLocalization wrapped in AppPreviewSetsResponseIncludedInner
func AppStoreVersionExperimentTreatmentLocalizationAsAppPreviewSetsResponseIncludedInner(v *AppStoreVersionExperimentTreatmentLocalization) AppPreviewSetsResponseIncludedInner {
	return AppPreviewSetsResponseIncludedInner{
		AppStoreVersionExperimentTreatmentLocalization: v,
	}
}

// AppStoreVersionLocalizationAsAppPreviewSetsResponseIncludedInner is a convenience function that returns AppStoreVersionLocalization wrapped in AppPreviewSetsResponseIncludedInner
func AppStoreVersionLocalizationAsAppPreviewSetsResponseIncludedInner(v *AppStoreVersionLocalization) AppPreviewSetsResponseIncludedInner {
	return AppPreviewSetsResponseIncludedInner{
		AppStoreVersionLocalization: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppPreviewSetsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppCustomProductPageLocalization
	err = newStrictDecoder(data).Decode(&dst.AppCustomProductPageLocalization)
	if err == nil {
		jsonAppCustomProductPageLocalization, _ := json.Marshal(dst.AppCustomProductPageLocalization)
		if string(jsonAppCustomProductPageLocalization) == "{}" { // empty struct
			dst.AppCustomProductPageLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppCustomProductPageLocalization = nil
	}

	// try to unmarshal data into AppPreview
	err = newStrictDecoder(data).Decode(&dst.AppPreview)
	if err == nil {
		jsonAppPreview, _ := json.Marshal(dst.AppPreview)
		if string(jsonAppPreview) == "{}" { // empty struct
			dst.AppPreview = nil
		} else {
			match++
		}
	} else {
		dst.AppPreview = nil
	}

	// try to unmarshal data into AppStoreVersionExperimentTreatmentLocalization
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentTreatmentLocalization)
	if err == nil {
		jsonAppStoreVersionExperimentTreatmentLocalization, _ := json.Marshal(dst.AppStoreVersionExperimentTreatmentLocalization)
		if string(jsonAppStoreVersionExperimentTreatmentLocalization) == "{}" { // empty struct
			dst.AppStoreVersionExperimentTreatmentLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperimentTreatmentLocalization = nil
	}

	// try to unmarshal data into AppStoreVersionLocalization
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionLocalization)
	if err == nil {
		jsonAppStoreVersionLocalization, _ := json.Marshal(dst.AppStoreVersionLocalization)
		if string(jsonAppStoreVersionLocalization) == "{}" { // empty struct
			dst.AppStoreVersionLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionLocalization = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppCustomProductPageLocalization = nil
		dst.AppPreview = nil
		dst.AppStoreVersionExperimentTreatmentLocalization = nil
		dst.AppStoreVersionLocalization = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppPreviewSetsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppPreviewSetsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppPreviewSetsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppCustomProductPageLocalization != nil {
		return json.Marshal(&src.AppCustomProductPageLocalization)
	}

	if src.AppPreview != nil {
		return json.Marshal(&src.AppPreview)
	}

	if src.AppStoreVersionExperimentTreatmentLocalization != nil {
		return json.Marshal(&src.AppStoreVersionExperimentTreatmentLocalization)
	}

	if src.AppStoreVersionLocalization != nil {
		return json.Marshal(&src.AppStoreVersionLocalization)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppPreviewSetsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppCustomProductPageLocalization != nil {
		return obj.AppCustomProductPageLocalization
	}

	if obj.AppPreview != nil {
		return obj.AppPreview
	}

	if obj.AppStoreVersionExperimentTreatmentLocalization != nil {
		return obj.AppStoreVersionExperimentTreatmentLocalization
	}

	if obj.AppStoreVersionLocalization != nil {
		return obj.AppStoreVersionLocalization
	}

	// all schemas are nil
	return nil
}

type NullableAppPreviewSetsResponseIncludedInner struct {
	value *AppPreviewSetsResponseIncludedInner
	isSet bool
}

func (v NullableAppPreviewSetsResponseIncludedInner) Get() *AppPreviewSetsResponseIncludedInner {
	return v.value
}

func (v *NullableAppPreviewSetsResponseIncludedInner) Set(val *AppPreviewSetsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetsResponseIncludedInner(val *AppPreviewSetsResponseIncludedInner) *NullableAppPreviewSetsResponseIncludedInner {
	return &NullableAppPreviewSetsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppPreviewSetsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


