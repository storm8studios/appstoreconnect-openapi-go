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

// checks if the AppPriceV2RelationshipsAppPricePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceV2RelationshipsAppPricePoint{}

// AppPriceV2RelationshipsAppPricePoint struct for AppPriceV2RelationshipsAppPricePoint
type AppPriceV2RelationshipsAppPricePoint struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *AppPriceTierRelationshipsPricePointsDataInner `json:"data,omitempty"`
}

// NewAppPriceV2RelationshipsAppPricePoint instantiates a new AppPriceV2RelationshipsAppPricePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceV2RelationshipsAppPricePoint() *AppPriceV2RelationshipsAppPricePoint {
	this := AppPriceV2RelationshipsAppPricePoint{}
	return &this
}

// NewAppPriceV2RelationshipsAppPricePointWithDefaults instantiates a new AppPriceV2RelationshipsAppPricePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceV2RelationshipsAppPricePointWithDefaults() *AppPriceV2RelationshipsAppPricePoint {
	this := AppPriceV2RelationshipsAppPricePoint{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppPriceV2RelationshipsAppPricePoint) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceV2RelationshipsAppPricePoint) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppPriceV2RelationshipsAppPricePoint) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *AppPriceV2RelationshipsAppPricePoint) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppPriceV2RelationshipsAppPricePoint) GetData() AppPriceTierRelationshipsPricePointsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppPriceTierRelationshipsPricePointsDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceV2RelationshipsAppPricePoint) GetDataOk() (*AppPriceTierRelationshipsPricePointsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppPriceV2RelationshipsAppPricePoint) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppPriceTierRelationshipsPricePointsDataInner and assigns it to the Data field.
func (o *AppPriceV2RelationshipsAppPricePoint) SetData(v AppPriceTierRelationshipsPricePointsDataInner) {
	o.Data = &v
}

func (o AppPriceV2RelationshipsAppPricePoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceV2RelationshipsAppPricePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppPriceV2RelationshipsAppPricePoint struct {
	value *AppPriceV2RelationshipsAppPricePoint
	isSet bool
}

func (v NullableAppPriceV2RelationshipsAppPricePoint) Get() *AppPriceV2RelationshipsAppPricePoint {
	return v.value
}

func (v *NullableAppPriceV2RelationshipsAppPricePoint) Set(val *AppPriceV2RelationshipsAppPricePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceV2RelationshipsAppPricePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceV2RelationshipsAppPricePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceV2RelationshipsAppPricePoint(val *AppPriceV2RelationshipsAppPricePoint) *NullableAppPriceV2RelationshipsAppPricePoint {
	return &NullableAppPriceV2RelationshipsAppPricePoint{value: val, isSet: true}
}

func (v NullableAppPriceV2RelationshipsAppPricePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceV2RelationshipsAppPricePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


