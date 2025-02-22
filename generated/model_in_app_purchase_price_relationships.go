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

// checks if the InAppPurchasePriceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceRelationships{}

// InAppPurchasePriceRelationships struct for InAppPurchasePriceRelationships
type InAppPurchasePriceRelationships struct {
	InAppPurchasePricePoint *InAppPurchasePriceRelationshipsInAppPurchasePricePoint `json:"inAppPurchasePricePoint,omitempty"`
	Territory *AppPricePointV2RelationshipsTerritory `json:"territory,omitempty"`
}

// NewInAppPurchasePriceRelationships instantiates a new InAppPurchasePriceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceRelationships() *InAppPurchasePriceRelationships {
	this := InAppPurchasePriceRelationships{}
	return &this
}

// NewInAppPurchasePriceRelationshipsWithDefaults instantiates a new InAppPurchasePriceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceRelationshipsWithDefaults() *InAppPurchasePriceRelationships {
	this := InAppPurchasePriceRelationships{}
	return &this
}

// GetInAppPurchasePricePoint returns the InAppPurchasePricePoint field value if set, zero value otherwise.
func (o *InAppPurchasePriceRelationships) GetInAppPurchasePricePoint() InAppPurchasePriceRelationshipsInAppPurchasePricePoint {
	if o == nil || IsNil(o.InAppPurchasePricePoint) {
		var ret InAppPurchasePriceRelationshipsInAppPurchasePricePoint
		return ret
	}
	return *o.InAppPurchasePricePoint
}

// GetInAppPurchasePricePointOk returns a tuple with the InAppPurchasePricePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceRelationships) GetInAppPurchasePricePointOk() (*InAppPurchasePriceRelationshipsInAppPurchasePricePoint, bool) {
	if o == nil || IsNil(o.InAppPurchasePricePoint) {
		return nil, false
	}
	return o.InAppPurchasePricePoint, true
}

// HasInAppPurchasePricePoint returns a boolean if a field has been set.
func (o *InAppPurchasePriceRelationships) HasInAppPurchasePricePoint() bool {
	if o != nil && !IsNil(o.InAppPurchasePricePoint) {
		return true
	}

	return false
}

// SetInAppPurchasePricePoint gets a reference to the given InAppPurchasePriceRelationshipsInAppPurchasePricePoint and assigns it to the InAppPurchasePricePoint field.
func (o *InAppPurchasePriceRelationships) SetInAppPurchasePricePoint(v InAppPurchasePriceRelationshipsInAppPurchasePricePoint) {
	o.InAppPurchasePricePoint = &v
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *InAppPurchasePriceRelationships) GetTerritory() AppPricePointV2RelationshipsTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret AppPricePointV2RelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceRelationships) GetTerritoryOk() (*AppPricePointV2RelationshipsTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *InAppPurchasePriceRelationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointV2RelationshipsTerritory and assigns it to the Territory field.
func (o *InAppPurchasePriceRelationships) SetTerritory(v AppPricePointV2RelationshipsTerritory) {
	o.Territory = &v
}

func (o InAppPurchasePriceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InAppPurchasePricePoint) {
		toSerialize["inAppPurchasePricePoint"] = o.InAppPurchasePricePoint
	}
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	return toSerialize, nil
}

type NullableInAppPurchasePriceRelationships struct {
	value *InAppPurchasePriceRelationships
	isSet bool
}

func (v NullableInAppPurchasePriceRelationships) Get() *InAppPurchasePriceRelationships {
	return v.value
}

func (v *NullableInAppPurchasePriceRelationships) Set(val *InAppPurchasePriceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceRelationships(val *InAppPurchasePriceRelationships) *NullableInAppPurchasePriceRelationships {
	return &NullableInAppPurchasePriceRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


