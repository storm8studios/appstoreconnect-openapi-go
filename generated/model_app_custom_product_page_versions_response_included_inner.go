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

// AppCustomProductPageVersionsResponseIncludedInner - struct for AppCustomProductPageVersionsResponseIncludedInner
type AppCustomProductPageVersionsResponseIncludedInner struct {
	AppCustomProductPage *AppCustomProductPage
	AppCustomProductPageLocalization *AppCustomProductPageLocalization
}

// AppCustomProductPageAsAppCustomProductPageVersionsResponseIncludedInner is a convenience function that returns AppCustomProductPage wrapped in AppCustomProductPageVersionsResponseIncludedInner
func AppCustomProductPageAsAppCustomProductPageVersionsResponseIncludedInner(v *AppCustomProductPage) AppCustomProductPageVersionsResponseIncludedInner {
	return AppCustomProductPageVersionsResponseIncludedInner{
		AppCustomProductPage: v,
	}
}

// AppCustomProductPageLocalizationAsAppCustomProductPageVersionsResponseIncludedInner is a convenience function that returns AppCustomProductPageLocalization wrapped in AppCustomProductPageVersionsResponseIncludedInner
func AppCustomProductPageLocalizationAsAppCustomProductPageVersionsResponseIncludedInner(v *AppCustomProductPageLocalization) AppCustomProductPageVersionsResponseIncludedInner {
	return AppCustomProductPageVersionsResponseIncludedInner{
		AppCustomProductPageLocalization: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppCustomProductPageVersionsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppCustomProductPage
	err = newStrictDecoder(data).Decode(&dst.AppCustomProductPage)
	if err == nil {
		jsonAppCustomProductPage, _ := json.Marshal(dst.AppCustomProductPage)
		if string(jsonAppCustomProductPage) == "{}" { // empty struct
			dst.AppCustomProductPage = nil
		} else {
			match++
		}
	} else {
		dst.AppCustomProductPage = nil
	}

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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppCustomProductPage = nil
		dst.AppCustomProductPageLocalization = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppCustomProductPageVersionsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppCustomProductPageVersionsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppCustomProductPageVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppCustomProductPage != nil {
		return json.Marshal(&src.AppCustomProductPage)
	}

	if src.AppCustomProductPageLocalization != nil {
		return json.Marshal(&src.AppCustomProductPageLocalization)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppCustomProductPageVersionsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppCustomProductPage != nil {
		return obj.AppCustomProductPage
	}

	if obj.AppCustomProductPageLocalization != nil {
		return obj.AppCustomProductPageLocalization
	}

	// all schemas are nil
	return nil
}

type NullableAppCustomProductPageVersionsResponseIncludedInner struct {
	value *AppCustomProductPageVersionsResponseIncludedInner
	isSet bool
}

func (v NullableAppCustomProductPageVersionsResponseIncludedInner) Get() *AppCustomProductPageVersionsResponseIncludedInner {
	return v.value
}

func (v *NullableAppCustomProductPageVersionsResponseIncludedInner) Set(val *AppCustomProductPageVersionsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionsResponseIncludedInner(val *AppCustomProductPageVersionsResponseIncludedInner) *NullableAppCustomProductPageVersionsResponseIncludedInner {
	return &NullableAppCustomProductPageVersionsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


