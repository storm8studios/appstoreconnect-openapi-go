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

// checks if the PromotedPurchaseUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseUpdateRequestDataAttributes{}

// PromotedPurchaseUpdateRequestDataAttributes struct for PromotedPurchaseUpdateRequestDataAttributes
type PromotedPurchaseUpdateRequestDataAttributes struct {
	VisibleForAllUsers *bool `json:"visibleForAllUsers,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// NewPromotedPurchaseUpdateRequestDataAttributes instantiates a new PromotedPurchaseUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseUpdateRequestDataAttributes() *PromotedPurchaseUpdateRequestDataAttributes {
	this := PromotedPurchaseUpdateRequestDataAttributes{}
	return &this
}

// NewPromotedPurchaseUpdateRequestDataAttributesWithDefaults instantiates a new PromotedPurchaseUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseUpdateRequestDataAttributesWithDefaults() *PromotedPurchaseUpdateRequestDataAttributes {
	this := PromotedPurchaseUpdateRequestDataAttributes{}
	return &this
}

// GetVisibleForAllUsers returns the VisibleForAllUsers field value if set, zero value otherwise.
func (o *PromotedPurchaseUpdateRequestDataAttributes) GetVisibleForAllUsers() bool {
	if o == nil || IsNil(o.VisibleForAllUsers) {
		var ret bool
		return ret
	}
	return *o.VisibleForAllUsers
}

// GetVisibleForAllUsersOk returns a tuple with the VisibleForAllUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseUpdateRequestDataAttributes) GetVisibleForAllUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.VisibleForAllUsers) {
		return nil, false
	}
	return o.VisibleForAllUsers, true
}

// HasVisibleForAllUsers returns a boolean if a field has been set.
func (o *PromotedPurchaseUpdateRequestDataAttributes) HasVisibleForAllUsers() bool {
	if o != nil && !IsNil(o.VisibleForAllUsers) {
		return true
	}

	return false
}

// SetVisibleForAllUsers gets a reference to the given bool and assigns it to the VisibleForAllUsers field.
func (o *PromotedPurchaseUpdateRequestDataAttributes) SetVisibleForAllUsers(v bool) {
	o.VisibleForAllUsers = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PromotedPurchaseUpdateRequestDataAttributes) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseUpdateRequestDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PromotedPurchaseUpdateRequestDataAttributes) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PromotedPurchaseUpdateRequestDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o PromotedPurchaseUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VisibleForAllUsers) {
		toSerialize["visibleForAllUsers"] = o.VisibleForAllUsers
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseUpdateRequestDataAttributes struct {
	value *PromotedPurchaseUpdateRequestDataAttributes
	isSet bool
}

func (v NullablePromotedPurchaseUpdateRequestDataAttributes) Get() *PromotedPurchaseUpdateRequestDataAttributes {
	return v.value
}

func (v *NullablePromotedPurchaseUpdateRequestDataAttributes) Set(val *PromotedPurchaseUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseUpdateRequestDataAttributes(val *PromotedPurchaseUpdateRequestDataAttributes) *NullablePromotedPurchaseUpdateRequestDataAttributes {
	return &NullablePromotedPurchaseUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePromotedPurchaseUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


