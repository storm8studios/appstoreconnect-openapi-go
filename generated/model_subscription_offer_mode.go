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

// SubscriptionOfferMode the model 'SubscriptionOfferMode'
type SubscriptionOfferMode string

// List of SubscriptionOfferMode
const (
	SUBSCRIPTIONOFFERMODE_PAY_AS_YOU_GO SubscriptionOfferMode = "PAY_AS_YOU_GO"
	SUBSCRIPTIONOFFERMODE_PAY_UP_FRONT SubscriptionOfferMode = "PAY_UP_FRONT"
	SUBSCRIPTIONOFFERMODE_FREE_TRIAL SubscriptionOfferMode = "FREE_TRIAL"
)

// All allowed values of SubscriptionOfferMode enum
var AllowedSubscriptionOfferModeEnumValues = []SubscriptionOfferMode{
	"PAY_AS_YOU_GO",
	"PAY_UP_FRONT",
	"FREE_TRIAL",
}

func (v *SubscriptionOfferMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionOfferMode(value)
	for _, existing := range AllowedSubscriptionOfferModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionOfferMode", value)
}

// NewSubscriptionOfferModeFromValue returns a pointer to a valid SubscriptionOfferMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionOfferModeFromValue(v string) (*SubscriptionOfferMode, error) {
	ev := SubscriptionOfferMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionOfferMode: valid values are %v", v, AllowedSubscriptionOfferModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionOfferMode) IsValid() bool {
	for _, existing := range AllowedSubscriptionOfferModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionOfferMode value
func (v SubscriptionOfferMode) Ptr() *SubscriptionOfferMode {
	return &v
}

type NullableSubscriptionOfferMode struct {
	value *SubscriptionOfferMode
	isSet bool
}

func (v NullableSubscriptionOfferMode) Get() *SubscriptionOfferMode {
	return v.value
}

func (v *NullableSubscriptionOfferMode) Set(val *SubscriptionOfferMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferMode(val *SubscriptionOfferMode) *NullableSubscriptionOfferMode {
	return &NullableSubscriptionOfferMode{value: val, isSet: true}
}

func (v NullableSubscriptionOfferMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

