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

// InAppPurchasePriceScheduleCreateRequestIncludedInner - struct for InAppPurchasePriceScheduleCreateRequestIncludedInner
type InAppPurchasePriceScheduleCreateRequestIncludedInner struct {
	InAppPurchasePriceInlineCreate *InAppPurchasePriceInlineCreate
	TerritoryInlineCreate *TerritoryInlineCreate
}

// InAppPurchasePriceInlineCreateAsInAppPurchasePriceScheduleCreateRequestIncludedInner is a convenience function that returns InAppPurchasePriceInlineCreate wrapped in InAppPurchasePriceScheduleCreateRequestIncludedInner
func InAppPurchasePriceInlineCreateAsInAppPurchasePriceScheduleCreateRequestIncludedInner(v *InAppPurchasePriceInlineCreate) InAppPurchasePriceScheduleCreateRequestIncludedInner {
	return InAppPurchasePriceScheduleCreateRequestIncludedInner{
		InAppPurchasePriceInlineCreate: v,
	}
}

// TerritoryInlineCreateAsInAppPurchasePriceScheduleCreateRequestIncludedInner is a convenience function that returns TerritoryInlineCreate wrapped in InAppPurchasePriceScheduleCreateRequestIncludedInner
func TerritoryInlineCreateAsInAppPurchasePriceScheduleCreateRequestIncludedInner(v *TerritoryInlineCreate) InAppPurchasePriceScheduleCreateRequestIncludedInner {
	return InAppPurchasePriceScheduleCreateRequestIncludedInner{
		TerritoryInlineCreate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InAppPurchasePriceScheduleCreateRequestIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InAppPurchasePriceInlineCreate
	err = newStrictDecoder(data).Decode(&dst.InAppPurchasePriceInlineCreate)
	if err == nil {
		jsonInAppPurchasePriceInlineCreate, _ := json.Marshal(dst.InAppPurchasePriceInlineCreate)
		if string(jsonInAppPurchasePriceInlineCreate) == "{}" { // empty struct
			dst.InAppPurchasePriceInlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.InAppPurchasePriceInlineCreate = nil
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
		dst.InAppPurchasePriceInlineCreate = nil
		dst.TerritoryInlineCreate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InAppPurchasePriceScheduleCreateRequestIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InAppPurchasePriceScheduleCreateRequestIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InAppPurchasePriceScheduleCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	if src.InAppPurchasePriceInlineCreate != nil {
		return json.Marshal(&src.InAppPurchasePriceInlineCreate)
	}

	if src.TerritoryInlineCreate != nil {
		return json.Marshal(&src.TerritoryInlineCreate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InAppPurchasePriceScheduleCreateRequestIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InAppPurchasePriceInlineCreate != nil {
		return obj.InAppPurchasePriceInlineCreate
	}

	if obj.TerritoryInlineCreate != nil {
		return obj.TerritoryInlineCreate
	}

	// all schemas are nil
	return nil
}

type NullableInAppPurchasePriceScheduleCreateRequestIncludedInner struct {
	value *InAppPurchasePriceScheduleCreateRequestIncludedInner
	isSet bool
}

func (v NullableInAppPurchasePriceScheduleCreateRequestIncludedInner) Get() *InAppPurchasePriceScheduleCreateRequestIncludedInner {
	return v.value
}

func (v *NullableInAppPurchasePriceScheduleCreateRequestIncludedInner) Set(val *InAppPurchasePriceScheduleCreateRequestIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceScheduleCreateRequestIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceScheduleCreateRequestIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceScheduleCreateRequestIncludedInner(val *InAppPurchasePriceScheduleCreateRequestIncludedInner) *NullableInAppPurchasePriceScheduleCreateRequestIncludedInner {
	return &NullableInAppPurchasePriceScheduleCreateRequestIncludedInner{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceScheduleCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceScheduleCreateRequestIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


