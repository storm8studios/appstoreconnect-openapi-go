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

// checks if the TerritoryAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoryAttributes{}

// TerritoryAttributes struct for TerritoryAttributes
type TerritoryAttributes struct {
	Currency *string `json:"currency,omitempty"`
}

// NewTerritoryAttributes instantiates a new TerritoryAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoryAttributes() *TerritoryAttributes {
	this := TerritoryAttributes{}
	return &this
}

// NewTerritoryAttributesWithDefaults instantiates a new TerritoryAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoryAttributesWithDefaults() *TerritoryAttributes {
	this := TerritoryAttributes{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *TerritoryAttributes) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAttributes) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *TerritoryAttributes) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *TerritoryAttributes) SetCurrency(v string) {
	o.Currency = &v
}

func (o TerritoryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoryAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableTerritoryAttributes struct {
	value *TerritoryAttributes
	isSet bool
}

func (v NullableTerritoryAttributes) Get() *TerritoryAttributes {
	return v.value
}

func (v *NullableTerritoryAttributes) Set(val *TerritoryAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoryAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoryAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoryAttributes(val *TerritoryAttributes) *NullableTerritoryAttributes {
	return &NullableTerritoryAttributes{value: val, isSet: true}
}

func (v NullableTerritoryAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoryAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


