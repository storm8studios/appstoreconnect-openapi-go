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

// checks if the InAppPurchasePriceInlineCreateRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceInlineCreateRelationships{}

// InAppPurchasePriceInlineCreateRelationships struct for InAppPurchasePriceInlineCreateRelationships
type InAppPurchasePriceInlineCreateRelationships struct {
	InAppPurchaseV2 *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 `json:"inAppPurchaseV2,omitempty"`
	InAppPurchasePricePoint *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint `json:"inAppPurchasePricePoint,omitempty"`
}

// NewInAppPurchasePriceInlineCreateRelationships instantiates a new InAppPurchasePriceInlineCreateRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceInlineCreateRelationships() *InAppPurchasePriceInlineCreateRelationships {
	this := InAppPurchasePriceInlineCreateRelationships{}
	return &this
}

// NewInAppPurchasePriceInlineCreateRelationshipsWithDefaults instantiates a new InAppPurchasePriceInlineCreateRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceInlineCreateRelationshipsWithDefaults() *InAppPurchasePriceInlineCreateRelationships {
	this := InAppPurchasePriceInlineCreateRelationships{}
	return &this
}

// GetInAppPurchaseV2 returns the InAppPurchaseV2 field value if set, zero value otherwise.
func (o *InAppPurchasePriceInlineCreateRelationships) GetInAppPurchaseV2() InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 {
	if o == nil || IsNil(o.InAppPurchaseV2) {
		var ret InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2
		return ret
	}
	return *o.InAppPurchaseV2
}

// GetInAppPurchaseV2Ok returns a tuple with the InAppPurchaseV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceInlineCreateRelationships) GetInAppPurchaseV2Ok() (*InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2, bool) {
	if o == nil || IsNil(o.InAppPurchaseV2) {
		return nil, false
	}
	return o.InAppPurchaseV2, true
}

// HasInAppPurchaseV2 returns a boolean if a field has been set.
func (o *InAppPurchasePriceInlineCreateRelationships) HasInAppPurchaseV2() bool {
	if o != nil && !IsNil(o.InAppPurchaseV2) {
		return true
	}

	return false
}

// SetInAppPurchaseV2 gets a reference to the given InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 and assigns it to the InAppPurchaseV2 field.
func (o *InAppPurchasePriceInlineCreateRelationships) SetInAppPurchaseV2(v InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) {
	o.InAppPurchaseV2 = &v
}

// GetInAppPurchasePricePoint returns the InAppPurchasePricePoint field value if set, zero value otherwise.
func (o *InAppPurchasePriceInlineCreateRelationships) GetInAppPurchasePricePoint() InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint {
	if o == nil || IsNil(o.InAppPurchasePricePoint) {
		var ret InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint
		return ret
	}
	return *o.InAppPurchasePricePoint
}

// GetInAppPurchasePricePointOk returns a tuple with the InAppPurchasePricePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceInlineCreateRelationships) GetInAppPurchasePricePointOk() (*InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint, bool) {
	if o == nil || IsNil(o.InAppPurchasePricePoint) {
		return nil, false
	}
	return o.InAppPurchasePricePoint, true
}

// HasInAppPurchasePricePoint returns a boolean if a field has been set.
func (o *InAppPurchasePriceInlineCreateRelationships) HasInAppPurchasePricePoint() bool {
	if o != nil && !IsNil(o.InAppPurchasePricePoint) {
		return true
	}

	return false
}

// SetInAppPurchasePricePoint gets a reference to the given InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint and assigns it to the InAppPurchasePricePoint field.
func (o *InAppPurchasePriceInlineCreateRelationships) SetInAppPurchasePricePoint(v InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) {
	o.InAppPurchasePricePoint = &v
}

func (o InAppPurchasePriceInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceInlineCreateRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InAppPurchaseV2) {
		toSerialize["inAppPurchaseV2"] = o.InAppPurchaseV2
	}
	if !IsNil(o.InAppPurchasePricePoint) {
		toSerialize["inAppPurchasePricePoint"] = o.InAppPurchasePricePoint
	}
	return toSerialize, nil
}

type NullableInAppPurchasePriceInlineCreateRelationships struct {
	value *InAppPurchasePriceInlineCreateRelationships
	isSet bool
}

func (v NullableInAppPurchasePriceInlineCreateRelationships) Get() *InAppPurchasePriceInlineCreateRelationships {
	return v.value
}

func (v *NullableInAppPurchasePriceInlineCreateRelationships) Set(val *InAppPurchasePriceInlineCreateRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceInlineCreateRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceInlineCreateRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceInlineCreateRelationships(val *InAppPurchasePriceInlineCreateRelationships) *NullableInAppPurchasePriceInlineCreateRelationships {
	return &NullableInAppPurchasePriceInlineCreateRelationships{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceInlineCreateRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceInlineCreateRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


