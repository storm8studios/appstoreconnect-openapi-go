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

// checks if the SubscriptionOfferCodeCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCreateRequestDataRelationships{}

// SubscriptionOfferCodeCreateRequestDataRelationships struct for SubscriptionOfferCodeCreateRequestDataRelationships
type SubscriptionOfferCodeCreateRequestDataRelationships struct {
	Subscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription `json:"subscription"`
	Prices SubscriptionOfferCodeCreateRequestDataRelationshipsPrices `json:"prices"`
}

// NewSubscriptionOfferCodeCreateRequestDataRelationships instantiates a new SubscriptionOfferCodeCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCreateRequestDataRelationships(subscription SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription, prices SubscriptionOfferCodeCreateRequestDataRelationshipsPrices) *SubscriptionOfferCodeCreateRequestDataRelationships {
	this := SubscriptionOfferCodeCreateRequestDataRelationships{}
	this.Subscription = subscription
	this.Prices = prices
	return &this
}

// NewSubscriptionOfferCodeCreateRequestDataRelationshipsWithDefaults instantiates a new SubscriptionOfferCodeCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCreateRequestDataRelationshipsWithDefaults() *SubscriptionOfferCodeCreateRequestDataRelationships {
	this := SubscriptionOfferCodeCreateRequestDataRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value
func (o *SubscriptionOfferCodeCreateRequestDataRelationships) GetSubscription() SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription {
	if o == nil {
		var ret SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataRelationships) GetSubscriptionOk() (*SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *SubscriptionOfferCodeCreateRequestDataRelationships) SetSubscription(v SubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription) {
	o.Subscription = v
}

// GetPrices returns the Prices field value
func (o *SubscriptionOfferCodeCreateRequestDataRelationships) GetPrices() SubscriptionOfferCodeCreateRequestDataRelationshipsPrices {
	if o == nil {
		var ret SubscriptionOfferCodeCreateRequestDataRelationshipsPrices
		return ret
	}

	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCreateRequestDataRelationships) GetPricesOk() (*SubscriptionOfferCodeCreateRequestDataRelationshipsPrices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prices, true
}

// SetPrices sets field value
func (o *SubscriptionOfferCodeCreateRequestDataRelationships) SetPrices(v SubscriptionOfferCodeCreateRequestDataRelationshipsPrices) {
	o.Prices = v
}

func (o SubscriptionOfferCodeCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscription"] = o.Subscription
	toSerialize["prices"] = o.Prices
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCreateRequestDataRelationships struct {
	value *SubscriptionOfferCodeCreateRequestDataRelationships
	isSet bool
}

func (v NullableSubscriptionOfferCodeCreateRequestDataRelationships) Get() *SubscriptionOfferCodeCreateRequestDataRelationships {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCreateRequestDataRelationships) Set(val *SubscriptionOfferCodeCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCreateRequestDataRelationships(val *SubscriptionOfferCodeCreateRequestDataRelationships) *NullableSubscriptionOfferCodeCreateRequestDataRelationships {
	return &NullableSubscriptionOfferCodeCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


