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

// SubscriptionOfferEligibility the model 'SubscriptionOfferEligibility'
type SubscriptionOfferEligibility string

// List of SubscriptionOfferEligibility
const (
	SUBSCRIPTIONOFFERELIGIBILITY_STACK_WITH_INTRO_OFFERS SubscriptionOfferEligibility = "STACK_WITH_INTRO_OFFERS"
	SUBSCRIPTIONOFFERELIGIBILITY_REPLACE_INTRO_OFFERS SubscriptionOfferEligibility = "REPLACE_INTRO_OFFERS"
)

// All allowed values of SubscriptionOfferEligibility enum
var AllowedSubscriptionOfferEligibilityEnumValues = []SubscriptionOfferEligibility{
	"STACK_WITH_INTRO_OFFERS",
	"REPLACE_INTRO_OFFERS",
}

func (v *SubscriptionOfferEligibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionOfferEligibility(value)
	for _, existing := range AllowedSubscriptionOfferEligibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionOfferEligibility", value)
}

// NewSubscriptionOfferEligibilityFromValue returns a pointer to a valid SubscriptionOfferEligibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionOfferEligibilityFromValue(v string) (*SubscriptionOfferEligibility, error) {
	ev := SubscriptionOfferEligibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionOfferEligibility: valid values are %v", v, AllowedSubscriptionOfferEligibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionOfferEligibility) IsValid() bool {
	for _, existing := range AllowedSubscriptionOfferEligibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionOfferEligibility value
func (v SubscriptionOfferEligibility) Ptr() *SubscriptionOfferEligibility {
	return &v
}

type NullableSubscriptionOfferEligibility struct {
	value *SubscriptionOfferEligibility
	isSet bool
}

func (v NullableSubscriptionOfferEligibility) Get() *SubscriptionOfferEligibility {
	return v.value
}

func (v *NullableSubscriptionOfferEligibility) Set(val *SubscriptionOfferEligibility) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferEligibility) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferEligibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferEligibility(val *SubscriptionOfferEligibility) *NullableSubscriptionOfferEligibility {
	return &NullableSubscriptionOfferEligibility{value: val, isSet: true}
}

func (v NullableSubscriptionOfferEligibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferEligibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

