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

// AppStoreVersionExperimentsResponseIncludedInner - struct for AppStoreVersionExperimentsResponseIncludedInner
type AppStoreVersionExperimentsResponseIncludedInner struct {
	AppStoreVersion *AppStoreVersion
	AppStoreVersionExperimentTreatment *AppStoreVersionExperimentTreatment
}

// AppStoreVersionAsAppStoreVersionExperimentsResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in AppStoreVersionExperimentsResponseIncludedInner
func AppStoreVersionAsAppStoreVersionExperimentsResponseIncludedInner(v *AppStoreVersion) AppStoreVersionExperimentsResponseIncludedInner {
	return AppStoreVersionExperimentsResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// AppStoreVersionExperimentTreatmentAsAppStoreVersionExperimentsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentTreatment wrapped in AppStoreVersionExperimentsResponseIncludedInner
func AppStoreVersionExperimentTreatmentAsAppStoreVersionExperimentsResponseIncludedInner(v *AppStoreVersionExperimentTreatment) AppStoreVersionExperimentsResponseIncludedInner {
	return AppStoreVersionExperimentsResponseIncludedInner{
		AppStoreVersionExperimentTreatment: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppStoreVersionExperimentsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppStoreVersion
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersion)
	if err == nil {
		jsonAppStoreVersion, _ := json.Marshal(dst.AppStoreVersion)
		if string(jsonAppStoreVersion) == "{}" { // empty struct
			dst.AppStoreVersion = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersion = nil
	}

	// try to unmarshal data into AppStoreVersionExperimentTreatment
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentTreatment)
	if err == nil {
		jsonAppStoreVersionExperimentTreatment, _ := json.Marshal(dst.AppStoreVersionExperimentTreatment)
		if string(jsonAppStoreVersionExperimentTreatment) == "{}" { // empty struct
			dst.AppStoreVersionExperimentTreatment = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperimentTreatment = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppStoreVersion = nil
		dst.AppStoreVersionExperimentTreatment = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppStoreVersionExperimentsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppStoreVersionExperimentsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppStoreVersionExperimentsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	if src.AppStoreVersionExperimentTreatment != nil {
		return json.Marshal(&src.AppStoreVersionExperimentTreatment)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppStoreVersionExperimentsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	if obj.AppStoreVersionExperimentTreatment != nil {
		return obj.AppStoreVersionExperimentTreatment
	}

	// all schemas are nil
	return nil
}

type NullableAppStoreVersionExperimentsResponseIncludedInner struct {
	value *AppStoreVersionExperimentsResponseIncludedInner
	isSet bool
}

func (v NullableAppStoreVersionExperimentsResponseIncludedInner) Get() *AppStoreVersionExperimentsResponseIncludedInner {
	return v.value
}

func (v *NullableAppStoreVersionExperimentsResponseIncludedInner) Set(val *AppStoreVersionExperimentsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionExperimentsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionExperimentsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionExperimentsResponseIncludedInner(val *AppStoreVersionExperimentsResponseIncludedInner) *NullableAppStoreVersionExperimentsResponseIncludedInner {
	return &NullableAppStoreVersionExperimentsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppStoreVersionExperimentsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionExperimentsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


