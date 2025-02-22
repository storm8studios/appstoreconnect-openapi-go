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

// checks if the SubscriptionPromotionalOfferUpdateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferUpdateRequestDataRelationships{}

// SubscriptionPromotionalOfferUpdateRequestDataRelationships struct for SubscriptionPromotionalOfferUpdateRequestDataRelationships
type SubscriptionPromotionalOfferUpdateRequestDataRelationships struct {
	Prices *SubscriptionPromotionalOfferInlineCreateRelationshipsPrices `json:"prices,omitempty"`
}

// NewSubscriptionPromotionalOfferUpdateRequestDataRelationships instantiates a new SubscriptionPromotionalOfferUpdateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferUpdateRequestDataRelationships() *SubscriptionPromotionalOfferUpdateRequestDataRelationships {
	this := SubscriptionPromotionalOfferUpdateRequestDataRelationships{}
	return &this
}

// NewSubscriptionPromotionalOfferUpdateRequestDataRelationshipsWithDefaults instantiates a new SubscriptionPromotionalOfferUpdateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferUpdateRequestDataRelationshipsWithDefaults() *SubscriptionPromotionalOfferUpdateRequestDataRelationships {
	this := SubscriptionPromotionalOfferUpdateRequestDataRelationships{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferUpdateRequestDataRelationships) GetPrices() SubscriptionPromotionalOfferInlineCreateRelationshipsPrices {
	if o == nil || IsNil(o.Prices) {
		var ret SubscriptionPromotionalOfferInlineCreateRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferUpdateRequestDataRelationships) GetPricesOk() (*SubscriptionPromotionalOfferInlineCreateRelationshipsPrices, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferUpdateRequestDataRelationships) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given SubscriptionPromotionalOfferInlineCreateRelationshipsPrices and assigns it to the Prices field.
func (o *SubscriptionPromotionalOfferUpdateRequestDataRelationships) SetPrices(v SubscriptionPromotionalOfferInlineCreateRelationshipsPrices) {
	o.Prices = &v
}

func (o SubscriptionPromotionalOfferUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferUpdateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships struct {
	value *SubscriptionPromotionalOfferUpdateRequestDataRelationships
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships) Get() *SubscriptionPromotionalOfferUpdateRequestDataRelationships {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships) Set(val *SubscriptionPromotionalOfferUpdateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferUpdateRequestDataRelationships(val *SubscriptionPromotionalOfferUpdateRequestDataRelationships) *NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships {
	return &NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferUpdateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


