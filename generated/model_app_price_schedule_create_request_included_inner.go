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

// AppPriceScheduleCreateRequestIncludedInner - struct for AppPriceScheduleCreateRequestIncludedInner
type AppPriceScheduleCreateRequestIncludedInner struct {
	AppPriceV2InlineCreate *AppPriceV2InlineCreate
	TerritoryInlineCreate *TerritoryInlineCreate
}

// AppPriceV2InlineCreateAsAppPriceScheduleCreateRequestIncludedInner is a convenience function that returns AppPriceV2InlineCreate wrapped in AppPriceScheduleCreateRequestIncludedInner
func AppPriceV2InlineCreateAsAppPriceScheduleCreateRequestIncludedInner(v *AppPriceV2InlineCreate) AppPriceScheduleCreateRequestIncludedInner {
	return AppPriceScheduleCreateRequestIncludedInner{
		AppPriceV2InlineCreate: v,
	}
}

// TerritoryInlineCreateAsAppPriceScheduleCreateRequestIncludedInner is a convenience function that returns TerritoryInlineCreate wrapped in AppPriceScheduleCreateRequestIncludedInner
func TerritoryInlineCreateAsAppPriceScheduleCreateRequestIncludedInner(v *TerritoryInlineCreate) AppPriceScheduleCreateRequestIncludedInner {
	return AppPriceScheduleCreateRequestIncludedInner{
		TerritoryInlineCreate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppPriceScheduleCreateRequestIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppPriceV2InlineCreate
	err = newStrictDecoder(data).Decode(&dst.AppPriceV2InlineCreate)
	if err == nil {
		jsonAppPriceV2InlineCreate, _ := json.Marshal(dst.AppPriceV2InlineCreate)
		if string(jsonAppPriceV2InlineCreate) == "{}" { // empty struct
			dst.AppPriceV2InlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppPriceV2InlineCreate = nil
	}

	// try to unmarshal data into TerritoryInlineCreate
	err = newStrictDecoder(data).Decode(&dst.TerritoryInlineCreate)
	if err == nil {
		jsonTerritoryInlineCreate, _ := json.Marshal(dst.TerritoryInlineCreate)
		if string(jsonTerritoryInlineCreate) == "{}" { // empty struct
			dst.TerritoryInlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.TerritoryInlineCreate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppPriceV2InlineCreate = nil
		dst.TerritoryInlineCreate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppPriceScheduleCreateRequestIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppPriceScheduleCreateRequestIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppPriceScheduleCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppPriceV2InlineCreate != nil {
		return json.Marshal(&src.AppPriceV2InlineCreate)
	}

	if src.TerritoryInlineCreate != nil {
		return json.Marshal(&src.TerritoryInlineCreate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppPriceScheduleCreateRequestIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppPriceV2InlineCreate != nil {
		return obj.AppPriceV2InlineCreate
	}

	if obj.TerritoryInlineCreate != nil {
		return obj.TerritoryInlineCreate
	}

	// all schemas are nil
	return nil
}

type NullableAppPriceScheduleCreateRequestIncludedInner struct {
	value *AppPriceScheduleCreateRequestIncludedInner
	isSet bool
}

func (v NullableAppPriceScheduleCreateRequestIncludedInner) Get() *AppPriceScheduleCreateRequestIncludedInner {
	return v.value
}

func (v *NullableAppPriceScheduleCreateRequestIncludedInner) Set(val *AppPriceScheduleCreateRequestIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceScheduleCreateRequestIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceScheduleCreateRequestIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceScheduleCreateRequestIncludedInner(val *AppPriceScheduleCreateRequestIncludedInner) *NullableAppPriceScheduleCreateRequestIncludedInner {
	return &NullableAppPriceScheduleCreateRequestIncludedInner{value: val, isSet: true}
}

func (v NullableAppPriceScheduleCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceScheduleCreateRequestIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


