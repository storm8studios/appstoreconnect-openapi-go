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

// checks if the InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint{}

// InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint struct for InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint
type InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint struct {
	Data *InAppPurchasePriceRelationshipsInAppPurchasePricePointData `json:"data,omitempty"`
}

// NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint instantiates a new InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint() *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint {
	this := InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint{}
	return &this
}

// NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePointWithDefaults instantiates a new InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePointWithDefaults() *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint {
	this := InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) GetData() InAppPurchasePriceRelationshipsInAppPurchasePricePointData {
	if o == nil || IsNil(o.Data) {
		var ret InAppPurchasePriceRelationshipsInAppPurchasePricePointData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) GetDataOk() (*InAppPurchasePriceRelationshipsInAppPurchasePricePointData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given InAppPurchasePriceRelationshipsInAppPurchasePricePointData and assigns it to the Data field.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) SetData(v InAppPurchasePriceRelationshipsInAppPurchasePricePointData) {
	o.Data = &v
}

func (o InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint struct {
	value *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint
	isSet bool
}

func (v NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) Get() *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint {
	return v.value
}

func (v *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) Set(val *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint(val *InAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint {
	return &NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchasePricePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


