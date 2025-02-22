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

// InAppPurchasePricesResponseIncludedInner - struct for InAppPurchasePricesResponseIncludedInner
type InAppPurchasePricesResponseIncludedInner struct {
	InAppPurchasePricePoint *InAppPurchasePricePoint
	Territory *Territory
}

// InAppPurchasePricePointAsInAppPurchasePricesResponseIncludedInner is a convenience function that returns InAppPurchasePricePoint wrapped in InAppPurchasePricesResponseIncludedInner
func InAppPurchasePricePointAsInAppPurchasePricesResponseIncludedInner(v *InAppPurchasePricePoint) InAppPurchasePricesResponseIncludedInner {
	return InAppPurchasePricesResponseIncludedInner{
		InAppPurchasePricePoint: v,
	}
}

// TerritoryAsInAppPurchasePricesResponseIncludedInner is a convenience function that returns Territory wrapped in InAppPurchasePricesResponseIncludedInner
func TerritoryAsInAppPurchasePricesResponseIncludedInner(v *Territory) InAppPurchasePricesResponseIncludedInner {
	return InAppPurchasePricesResponseIncludedInner{
		Territory: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InAppPurchasePricesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InAppPurchasePricePoint
	err = newStrictDecoder(data).Decode(&dst.InAppPurchasePricePoint)
	if err == nil {
		jsonInAppPurchasePricePoint, _ := json.Marshal(dst.InAppPurchasePricePoint)
		if string(jsonInAppPurchasePricePoint) == "{}" { // empty struct
			dst.InAppPurchasePricePoint = nil
		} else {
			match++
		}
	} else {
		dst.InAppPurchasePricePoint = nil
	}

	// try to unmarshal data into Territory
	err = newStrictDecoder(data).Decode(&dst.Territory)
	if err == nil {
		jsonTerritory, _ := json.Marshal(dst.Territory)
		if string(jsonTerritory) == "{}" { // empty struct
			dst.Territory = nil
		} else {
			match++
		}
	} else {
		dst.Territory = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InAppPurchasePricePoint = nil
		dst.Territory = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InAppPurchasePricesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InAppPurchasePricesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InAppPurchasePricesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.InAppPurchasePricePoint != nil {
		return json.Marshal(&src.InAppPurchasePricePoint)
	}

	if src.Territory != nil {
		return json.Marshal(&src.Territory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InAppPurchasePricesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InAppPurchasePricePoint != nil {
		return obj.InAppPurchasePricePoint
	}

	if obj.Territory != nil {
		return obj.Territory
	}

	// all schemas are nil
	return nil
}

type NullableInAppPurchasePricesResponseIncludedInner struct {
	value *InAppPurchasePricesResponseIncludedInner
	isSet bool
}

func (v NullableInAppPurchasePricesResponseIncludedInner) Get() *InAppPurchasePricesResponseIncludedInner {
	return v.value
}

func (v *NullableInAppPurchasePricesResponseIncludedInner) Set(val *InAppPurchasePricesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePricesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePricesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePricesResponseIncludedInner(val *InAppPurchasePricesResponseIncludedInner) *NullableInAppPurchasePricesResponseIncludedInner {
	return &NullableInAppPurchasePricesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableInAppPurchasePricesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePricesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


