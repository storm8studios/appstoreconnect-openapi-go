/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
)

// checks if the InAppPurchaseAvailabilityRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAvailabilityRelationships{}

// InAppPurchaseAvailabilityRelationships struct for InAppPurchaseAvailabilityRelationships
type InAppPurchaseAvailabilityRelationships struct {
	AvailableTerritories *AppAvailabilityRelationshipsAvailableTerritories `json:"availableTerritories,omitempty"`
}

// NewInAppPurchaseAvailabilityRelationships instantiates a new InAppPurchaseAvailabilityRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAvailabilityRelationships() *InAppPurchaseAvailabilityRelationships {
	this := InAppPurchaseAvailabilityRelationships{}
	return &this
}

// NewInAppPurchaseAvailabilityRelationshipsWithDefaults instantiates a new InAppPurchaseAvailabilityRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAvailabilityRelationshipsWithDefaults() *InAppPurchaseAvailabilityRelationships {
	this := InAppPurchaseAvailabilityRelationships{}
	return &this
}

// GetAvailableTerritories returns the AvailableTerritories field value if set, zero value otherwise.
func (o *InAppPurchaseAvailabilityRelationships) GetAvailableTerritories() AppAvailabilityRelationshipsAvailableTerritories {
	if o == nil || IsNil(o.AvailableTerritories) {
		var ret AppAvailabilityRelationshipsAvailableTerritories
		return ret
	}
	return *o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityRelationships) GetAvailableTerritoriesOk() (*AppAvailabilityRelationshipsAvailableTerritories, bool) {
	if o == nil || IsNil(o.AvailableTerritories) {
		return nil, false
	}
	return o.AvailableTerritories, true
}

// HasAvailableTerritories returns a boolean if a field has been set.
func (o *InAppPurchaseAvailabilityRelationships) HasAvailableTerritories() bool {
	if o != nil && !IsNil(o.AvailableTerritories) {
		return true
	}

	return false
}

// SetAvailableTerritories gets a reference to the given AppAvailabilityRelationshipsAvailableTerritories and assigns it to the AvailableTerritories field.
func (o *InAppPurchaseAvailabilityRelationships) SetAvailableTerritories(v AppAvailabilityRelationshipsAvailableTerritories) {
	o.AvailableTerritories = &v
}

func (o InAppPurchaseAvailabilityRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAvailabilityRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailableTerritories) {
		toSerialize["availableTerritories"] = o.AvailableTerritories
	}
	return toSerialize, nil
}

type NullableInAppPurchaseAvailabilityRelationships struct {
	value *InAppPurchaseAvailabilityRelationships
	isSet bool
}

func (v NullableInAppPurchaseAvailabilityRelationships) Get() *InAppPurchaseAvailabilityRelationships {
	return v.value
}

func (v *NullableInAppPurchaseAvailabilityRelationships) Set(val *InAppPurchaseAvailabilityRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAvailabilityRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAvailabilityRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAvailabilityRelationships(val *InAppPurchaseAvailabilityRelationships) *NullableInAppPurchaseAvailabilityRelationships {
	return &NullableInAppPurchaseAvailabilityRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchaseAvailabilityRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAvailabilityRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


