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

// BetaTestersResponseIncludedInner - struct for BetaTestersResponseIncludedInner
type BetaTestersResponseIncludedInner struct {
	App *App
	BetaGroup *BetaGroup
	Build *Build
}

// AppAsBetaTestersResponseIncludedInner is a convenience function that returns App wrapped in BetaTestersResponseIncludedInner
func AppAsBetaTestersResponseIncludedInner(v *App) BetaTestersResponseIncludedInner {
	return BetaTestersResponseIncludedInner{
		App: v,
	}
}

// BetaGroupAsBetaTestersResponseIncludedInner is a convenience function that returns BetaGroup wrapped in BetaTestersResponseIncludedInner
func BetaGroupAsBetaTestersResponseIncludedInner(v *BetaGroup) BetaTestersResponseIncludedInner {
	return BetaTestersResponseIncludedInner{
		BetaGroup: v,
	}
}

// BuildAsBetaTestersResponseIncludedInner is a convenience function that returns Build wrapped in BetaTestersResponseIncludedInner
func BuildAsBetaTestersResponseIncludedInner(v *Build) BetaTestersResponseIncludedInner {
	return BetaTestersResponseIncludedInner{
		Build: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BetaTestersResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into App
	err = newStrictDecoder(data).Decode(&dst.App)
	if err == nil {
		jsonApp, _ := json.Marshal(dst.App)
		if string(jsonApp) == "{}" { // empty struct
			dst.App = nil
		} else {
			match++
		}
	} else {
		dst.App = nil
	}

	// try to unmarshal data into BetaGroup
	err = newStrictDecoder(data).Decode(&dst.BetaGroup)
	if err == nil {
		jsonBetaGroup, _ := json.Marshal(dst.BetaGroup)
		if string(jsonBetaGroup) == "{}" { // empty struct
			dst.BetaGroup = nil
		} else {
			match++
		}
	} else {
		dst.BetaGroup = nil
	}

	// try to unmarshal data into Build
	err = newStrictDecoder(data).Decode(&dst.Build)
	if err == nil {
		jsonBuild, _ := json.Marshal(dst.Build)
		if string(jsonBuild) == "{}" { // empty struct
			dst.Build = nil
		} else {
			match++
		}
	} else {
		dst.Build = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.BetaGroup = nil
		dst.Build = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BetaTestersResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BetaTestersResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BetaTestersResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.BetaGroup != nil {
		return json.Marshal(&src.BetaGroup)
	}

	if src.Build != nil {
		return json.Marshal(&src.Build)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BetaTestersResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.BetaGroup != nil {
		return obj.BetaGroup
	}

	if obj.Build != nil {
		return obj.Build
	}

	// all schemas are nil
	return nil
}

type NullableBetaTestersResponseIncludedInner struct {
	value *BetaTestersResponseIncludedInner
	isSet bool
}

func (v NullableBetaTestersResponseIncludedInner) Get() *BetaTestersResponseIncludedInner {
	return v.value
}

func (v *NullableBetaTestersResponseIncludedInner) Set(val *BetaTestersResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTestersResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTestersResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTestersResponseIncludedInner(val *BetaTestersResponseIncludedInner) *NullableBetaTestersResponseIncludedInner {
	return &NullableBetaTestersResponseIncludedInner{value: val, isSet: true}
}

func (v NullableBetaTestersResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTestersResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


