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

// checks if the SubscriptionPromotionalOfferRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferRelationships{}

// SubscriptionPromotionalOfferRelationships struct for SubscriptionPromotionalOfferRelationships
type SubscriptionPromotionalOfferRelationships struct {
	Subscription *PromotedPurchaseRelationshipsSubscription `json:"subscription,omitempty"`
	Prices *SubscriptionPromotionalOfferRelationshipsPrices `json:"prices,omitempty"`
}

// NewSubscriptionPromotionalOfferRelationships instantiates a new SubscriptionPromotionalOfferRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferRelationships() *SubscriptionPromotionalOfferRelationships {
	this := SubscriptionPromotionalOfferRelationships{}
	return &this
}

// NewSubscriptionPromotionalOfferRelationshipsWithDefaults instantiates a new SubscriptionPromotionalOfferRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferRelationshipsWithDefaults() *SubscriptionPromotionalOfferRelationships {
	this := SubscriptionPromotionalOfferRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferRelationships) GetSubscription() PromotedPurchaseRelationshipsSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret PromotedPurchaseRelationshipsSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationships) GetSubscriptionOk() (*PromotedPurchaseRelationshipsSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferRelationships) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given PromotedPurchaseRelationshipsSubscription and assigns it to the Subscription field.
func (o *SubscriptionPromotionalOfferRelationships) SetSubscription(v PromotedPurchaseRelationshipsSubscription) {
	o.Subscription = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferRelationships) GetPrices() SubscriptionPromotionalOfferRelationshipsPrices {
	if o == nil || IsNil(o.Prices) {
		var ret SubscriptionPromotionalOfferRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferRelationships) GetPricesOk() (*SubscriptionPromotionalOfferRelationshipsPrices, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferRelationships) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given SubscriptionPromotionalOfferRelationshipsPrices and assigns it to the Prices field.
func (o *SubscriptionPromotionalOfferRelationships) SetPrices(v SubscriptionPromotionalOfferRelationshipsPrices) {
	o.Prices = &v
}

func (o SubscriptionPromotionalOfferRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferRelationships struct {
	value *SubscriptionPromotionalOfferRelationships
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferRelationships) Get() *SubscriptionPromotionalOfferRelationships {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferRelationships) Set(val *SubscriptionPromotionalOfferRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferRelationships(val *SubscriptionPromotionalOfferRelationships) *NullableSubscriptionPromotionalOfferRelationships {
	return &NullableSubscriptionPromotionalOfferRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


