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

// checks if the AppPricePointV3Relationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPricePointV3Relationships{}

// AppPricePointV3Relationships struct for AppPricePointV3Relationships
type AppPricePointV3Relationships struct {
	App *AppAvailabilityRelationshipsApp `json:"app,omitempty"`
	Territory *AppPricePointV2RelationshipsTerritory `json:"territory,omitempty"`
}

// NewAppPricePointV3Relationships instantiates a new AppPricePointV3Relationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointV3Relationships() *AppPricePointV3Relationships {
	this := AppPricePointV3Relationships{}
	return &this
}

// NewAppPricePointV3RelationshipsWithDefaults instantiates a new AppPricePointV3Relationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointV3RelationshipsWithDefaults() *AppPricePointV3Relationships {
	this := AppPricePointV3Relationships{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *AppPricePointV3Relationships) GetApp() AppAvailabilityRelationshipsApp {
	if o == nil || IsNil(o.App) {
		var ret AppAvailabilityRelationshipsApp
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointV3Relationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *AppPricePointV3Relationships) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppAvailabilityRelationshipsApp and assigns it to the App field.
func (o *AppPricePointV3Relationships) SetApp(v AppAvailabilityRelationshipsApp) {
	o.App = &v
}

// GetTerritory returns the Territory field value if set, zero value otherwise.
func (o *AppPricePointV3Relationships) GetTerritory() AppPricePointV2RelationshipsTerritory {
	if o == nil || IsNil(o.Territory) {
		var ret AppPricePointV2RelationshipsTerritory
		return ret
	}
	return *o.Territory
}

// GetTerritoryOk returns a tuple with the Territory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointV3Relationships) GetTerritoryOk() (*AppPricePointV2RelationshipsTerritory, bool) {
	if o == nil || IsNil(o.Territory) {
		return nil, false
	}
	return o.Territory, true
}

// HasTerritory returns a boolean if a field has been set.
func (o *AppPricePointV3Relationships) HasTerritory() bool {
	if o != nil && !IsNil(o.Territory) {
		return true
	}

	return false
}

// SetTerritory gets a reference to the given AppPricePointV2RelationshipsTerritory and assigns it to the Territory field.
func (o *AppPricePointV3Relationships) SetTerritory(v AppPricePointV2RelationshipsTerritory) {
	o.Territory = &v
}

func (o AppPricePointV3Relationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPricePointV3Relationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Territory) {
		toSerialize["territory"] = o.Territory
	}
	return toSerialize, nil
}

type NullableAppPricePointV3Relationships struct {
	value *AppPricePointV3Relationships
	isSet bool
}

func (v NullableAppPricePointV3Relationships) Get() *AppPricePointV3Relationships {
	return v.value
}

func (v *NullableAppPricePointV3Relationships) Set(val *AppPricePointV3Relationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointV3Relationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointV3Relationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointV3Relationships(val *AppPricePointV3Relationships) *NullableAppPricePointV3Relationships {
	return &NullableAppPricePointV3Relationships{value: val, isSet: true}
}

func (v NullableAppPricePointV3Relationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointV3Relationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


