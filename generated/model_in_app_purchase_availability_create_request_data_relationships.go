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

// checks if the InAppPurchaseAvailabilityCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseAvailabilityCreateRequestDataRelationships{}

// InAppPurchaseAvailabilityCreateRequestDataRelationships struct for InAppPurchaseAvailabilityCreateRequestDataRelationships
type InAppPurchaseAvailabilityCreateRequestDataRelationships struct {
	InAppPurchase InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2 `json:"inAppPurchase"`
	AvailableTerritories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories `json:"availableTerritories"`
}

// NewInAppPurchaseAvailabilityCreateRequestDataRelationships instantiates a new InAppPurchaseAvailabilityCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseAvailabilityCreateRequestDataRelationships(inAppPurchase InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2, availableTerritories AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) *InAppPurchaseAvailabilityCreateRequestDataRelationships {
	this := InAppPurchaseAvailabilityCreateRequestDataRelationships{}
	this.InAppPurchase = inAppPurchase
	this.AvailableTerritories = availableTerritories
	return &this
}

// NewInAppPurchaseAvailabilityCreateRequestDataRelationshipsWithDefaults instantiates a new InAppPurchaseAvailabilityCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseAvailabilityCreateRequestDataRelationshipsWithDefaults() *InAppPurchaseAvailabilityCreateRequestDataRelationships {
	this := InAppPurchaseAvailabilityCreateRequestDataRelationships{}
	return &this
}

// GetInAppPurchase returns the InAppPurchase field value
func (o *InAppPurchaseAvailabilityCreateRequestDataRelationships) GetInAppPurchase() InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2 {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2
		return ret
	}

	return o.InAppPurchase
}

// GetInAppPurchaseOk returns a tuple with the InAppPurchase field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityCreateRequestDataRelationships) GetInAppPurchaseOk() (*InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InAppPurchase, true
}

// SetInAppPurchase sets field value
func (o *InAppPurchaseAvailabilityCreateRequestDataRelationships) SetInAppPurchase(v InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationshipsInAppPurchaseV2) {
	o.InAppPurchase = v
}

// GetAvailableTerritories returns the AvailableTerritories field value
func (o *InAppPurchaseAvailabilityCreateRequestDataRelationships) GetAvailableTerritories() AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories {
	if o == nil {
		var ret AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories
		return ret
	}

	return o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseAvailabilityCreateRequestDataRelationships) GetAvailableTerritoriesOk() (*AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableTerritories, true
}

// SetAvailableTerritories sets field value
func (o *InAppPurchaseAvailabilityCreateRequestDataRelationships) SetAvailableTerritories(v AppAvailabilityCreateRequestDataRelationshipsAvailableTerritories) {
	o.AvailableTerritories = v
}

func (o InAppPurchaseAvailabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseAvailabilityCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inAppPurchase"] = o.InAppPurchase
	toSerialize["availableTerritories"] = o.AvailableTerritories
	return toSerialize, nil
}

type NullableInAppPurchaseAvailabilityCreateRequestDataRelationships struct {
	value *InAppPurchaseAvailabilityCreateRequestDataRelationships
	isSet bool
}

func (v NullableInAppPurchaseAvailabilityCreateRequestDataRelationships) Get() *InAppPurchaseAvailabilityCreateRequestDataRelationships {
	return v.value
}

func (v *NullableInAppPurchaseAvailabilityCreateRequestDataRelationships) Set(val *InAppPurchaseAvailabilityCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseAvailabilityCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseAvailabilityCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseAvailabilityCreateRequestDataRelationships(val *InAppPurchaseAvailabilityCreateRequestDataRelationships) *NullableInAppPurchaseAvailabilityCreateRequestDataRelationships {
	return &NullableInAppPurchaseAvailabilityCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchaseAvailabilityCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseAvailabilityCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


