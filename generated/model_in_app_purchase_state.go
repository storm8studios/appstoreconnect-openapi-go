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

// InAppPurchaseState the model 'InAppPurchaseState'
type InAppPurchaseState string

// List of InAppPurchaseState
const (
	INAPPPURCHASESTATE_MISSING_METADATA InAppPurchaseState = "MISSING_METADATA"
	INAPPPURCHASESTATE_WAITING_FOR_UPLOAD InAppPurchaseState = "WAITING_FOR_UPLOAD"
	INAPPPURCHASESTATE_PROCESSING_CONTENT InAppPurchaseState = "PROCESSING_CONTENT"
	INAPPPURCHASESTATE_READY_TO_SUBMIT InAppPurchaseState = "READY_TO_SUBMIT"
	INAPPPURCHASESTATE_WAITING_FOR_REVIEW InAppPurchaseState = "WAITING_FOR_REVIEW"
	INAPPPURCHASESTATE_IN_REVIEW InAppPurchaseState = "IN_REVIEW"
	INAPPPURCHASESTATE_DEVELOPER_ACTION_NEEDED InAppPurchaseState = "DEVELOPER_ACTION_NEEDED"
	INAPPPURCHASESTATE_PENDING_BINARY_APPROVAL InAppPurchaseState = "PENDING_BINARY_APPROVAL"
	INAPPPURCHASESTATE_APPROVED InAppPurchaseState = "APPROVED"
	INAPPPURCHASESTATE_DEVELOPER_REMOVED_FROM_SALE InAppPurchaseState = "DEVELOPER_REMOVED_FROM_SALE"
	INAPPPURCHASESTATE_REMOVED_FROM_SALE InAppPurchaseState = "REMOVED_FROM_SALE"
	INAPPPURCHASESTATE_REJECTED InAppPurchaseState = "REJECTED"
)

// All allowed values of InAppPurchaseState enum
var AllowedInAppPurchaseStateEnumValues = []InAppPurchaseState{
	"MISSING_METADATA",
	"WAITING_FOR_UPLOAD",
	"PROCESSING_CONTENT",
	"READY_TO_SUBMIT",
	"WAITING_FOR_REVIEW",
	"IN_REVIEW",
	"DEVELOPER_ACTION_NEEDED",
	"PENDING_BINARY_APPROVAL",
	"APPROVED",
	"DEVELOPER_REMOVED_FROM_SALE",
	"REMOVED_FROM_SALE",
	"REJECTED",
}

func (v *InAppPurchaseState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InAppPurchaseState(value)
	for _, existing := range AllowedInAppPurchaseStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InAppPurchaseState", value)
}

// NewInAppPurchaseStateFromValue returns a pointer to a valid InAppPurchaseState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInAppPurchaseStateFromValue(v string) (*InAppPurchaseState, error) {
	ev := InAppPurchaseState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InAppPurchaseState: valid values are %v", v, AllowedInAppPurchaseStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InAppPurchaseState) IsValid() bool {
	for _, existing := range AllowedInAppPurchaseStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InAppPurchaseState value
func (v InAppPurchaseState) Ptr() *InAppPurchaseState {
	return &v
}

type NullableInAppPurchaseState struct {
	value *InAppPurchaseState
	isSet bool
}

func (v NullableInAppPurchaseState) Get() *InAppPurchaseState {
	return v.value
}

func (v *NullableInAppPurchaseState) Set(val *InAppPurchaseState) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseState) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseState(val *InAppPurchaseState) *NullableInAppPurchaseState {
	return &NullableInAppPurchaseState{value: val, isSet: true}
}

func (v NullableInAppPurchaseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

