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

// SubscriptionOfferCodesResponseIncludedInner - struct for SubscriptionOfferCodesResponseIncludedInner
type SubscriptionOfferCodesResponseIncludedInner struct {
	Subscription *Subscription
	SubscriptionOfferCodeCustomCode *SubscriptionOfferCodeCustomCode
	SubscriptionOfferCodeOneTimeUseCode *SubscriptionOfferCodeOneTimeUseCode
	SubscriptionOfferCodePrice *SubscriptionOfferCodePrice
}

// SubscriptionAsSubscriptionOfferCodesResponseIncludedInner is a convenience function that returns Subscription wrapped in SubscriptionOfferCodesResponseIncludedInner
func SubscriptionAsSubscriptionOfferCodesResponseIncludedInner(v *Subscription) SubscriptionOfferCodesResponseIncludedInner {
	return SubscriptionOfferCodesResponseIncludedInner{
		Subscription: v,
	}
}

// SubscriptionOfferCodeCustomCodeAsSubscriptionOfferCodesResponseIncludedInner is a convenience function that returns SubscriptionOfferCodeCustomCode wrapped in SubscriptionOfferCodesResponseIncludedInner
func SubscriptionOfferCodeCustomCodeAsSubscriptionOfferCodesResponseIncludedInner(v *SubscriptionOfferCodeCustomCode) SubscriptionOfferCodesResponseIncludedInner {
	return SubscriptionOfferCodesResponseIncludedInner{
		SubscriptionOfferCodeCustomCode: v,
	}
}

// SubscriptionOfferCodeOneTimeUseCodeAsSubscriptionOfferCodesResponseIncludedInner is a convenience function that returns SubscriptionOfferCodeOneTimeUseCode wrapped in SubscriptionOfferCodesResponseIncludedInner
func SubscriptionOfferCodeOneTimeUseCodeAsSubscriptionOfferCodesResponseIncludedInner(v *SubscriptionOfferCodeOneTimeUseCode) SubscriptionOfferCodesResponseIncludedInner {
	return SubscriptionOfferCodesResponseIncludedInner{
		SubscriptionOfferCodeOneTimeUseCode: v,
	}
}

// SubscriptionOfferCodePriceAsSubscriptionOfferCodesResponseIncludedInner is a convenience function that returns SubscriptionOfferCodePrice wrapped in SubscriptionOfferCodesResponseIncludedInner
func SubscriptionOfferCodePriceAsSubscriptionOfferCodesResponseIncludedInner(v *SubscriptionOfferCodePrice) SubscriptionOfferCodesResponseIncludedInner {
	return SubscriptionOfferCodesResponseIncludedInner{
		SubscriptionOfferCodePrice: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionOfferCodesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Subscription
	err = newStrictDecoder(data).Decode(&dst.Subscription)
	if err == nil {
		jsonSubscription, _ := json.Marshal(dst.Subscription)
		if string(jsonSubscription) == "{}" { // empty struct
			dst.Subscription = nil
		} else {
			match++
		}
	} else {
		dst.Subscription = nil
	}

	// try to unmarshal data into SubscriptionOfferCodeCustomCode
	err = newStrictDecoder(data).Decode(&dst.SubscriptionOfferCodeCustomCode)
	if err == nil {
		jsonSubscriptionOfferCodeCustomCode, _ := json.Marshal(dst.SubscriptionOfferCodeCustomCode)
		if string(jsonSubscriptionOfferCodeCustomCode) == "{}" { // empty struct
			dst.SubscriptionOfferCodeCustomCode = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionOfferCodeCustomCode = nil
	}

	// try to unmarshal data into SubscriptionOfferCodeOneTimeUseCode
	err = newStrictDecoder(data).Decode(&dst.SubscriptionOfferCodeOneTimeUseCode)
	if err == nil {
		jsonSubscriptionOfferCodeOneTimeUseCode, _ := json.Marshal(dst.SubscriptionOfferCodeOneTimeUseCode)
		if string(jsonSubscriptionOfferCodeOneTimeUseCode) == "{}" { // empty struct
			dst.SubscriptionOfferCodeOneTimeUseCode = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionOfferCodeOneTimeUseCode = nil
	}

	// try to unmarshal data into SubscriptionOfferCodePrice
	err = newStrictDecoder(data).Decode(&dst.SubscriptionOfferCodePrice)
	if err == nil {
		jsonSubscriptionOfferCodePrice, _ := json.Marshal(dst.SubscriptionOfferCodePrice)
		if string(jsonSubscriptionOfferCodePrice) == "{}" { // empty struct
			dst.SubscriptionOfferCodePrice = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionOfferCodePrice = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Subscription = nil
		dst.SubscriptionOfferCodeCustomCode = nil
		dst.SubscriptionOfferCodeOneTimeUseCode = nil
		dst.SubscriptionOfferCodePrice = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionOfferCodesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionOfferCodesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionOfferCodesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.Subscription != nil {
		return json.Marshal(&src.Subscription)
	}

	if src.SubscriptionOfferCodeCustomCode != nil {
		return json.Marshal(&src.SubscriptionOfferCodeCustomCode)
	}

	if src.SubscriptionOfferCodeOneTimeUseCode != nil {
		return json.Marshal(&src.SubscriptionOfferCodeOneTimeUseCode)
	}

	if src.SubscriptionOfferCodePrice != nil {
		return json.Marshal(&src.SubscriptionOfferCodePrice)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionOfferCodesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Subscription != nil {
		return obj.Subscription
	}

	if obj.SubscriptionOfferCodeCustomCode != nil {
		return obj.SubscriptionOfferCodeCustomCode
	}

	if obj.SubscriptionOfferCodeOneTimeUseCode != nil {
		return obj.SubscriptionOfferCodeOneTimeUseCode
	}

	if obj.SubscriptionOfferCodePrice != nil {
		return obj.SubscriptionOfferCodePrice
	}

	// all schemas are nil
	return nil
}

type NullableSubscriptionOfferCodesResponseIncludedInner struct {
	value *SubscriptionOfferCodesResponseIncludedInner
	isSet bool
}

func (v NullableSubscriptionOfferCodesResponseIncludedInner) Get() *SubscriptionOfferCodesResponseIncludedInner {
	return v.value
}

func (v *NullableSubscriptionOfferCodesResponseIncludedInner) Set(val *SubscriptionOfferCodesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodesResponseIncludedInner(val *SubscriptionOfferCodesResponseIncludedInner) *NullableSubscriptionOfferCodesResponseIncludedInner {
	return &NullableSubscriptionOfferCodesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


