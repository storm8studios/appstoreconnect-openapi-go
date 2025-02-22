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

// checks if the SubscriptionAppStoreReviewScreenshotRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAppStoreReviewScreenshotRelationships{}

// SubscriptionAppStoreReviewScreenshotRelationships struct for SubscriptionAppStoreReviewScreenshotRelationships
type SubscriptionAppStoreReviewScreenshotRelationships struct {
	Subscription *PromotedPurchaseRelationshipsSubscription `json:"subscription,omitempty"`
}

// NewSubscriptionAppStoreReviewScreenshotRelationships instantiates a new SubscriptionAppStoreReviewScreenshotRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAppStoreReviewScreenshotRelationships() *SubscriptionAppStoreReviewScreenshotRelationships {
	this := SubscriptionAppStoreReviewScreenshotRelationships{}
	return &this
}

// NewSubscriptionAppStoreReviewScreenshotRelationshipsWithDefaults instantiates a new SubscriptionAppStoreReviewScreenshotRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAppStoreReviewScreenshotRelationshipsWithDefaults() *SubscriptionAppStoreReviewScreenshotRelationships {
	this := SubscriptionAppStoreReviewScreenshotRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SubscriptionAppStoreReviewScreenshotRelationships) GetSubscription() PromotedPurchaseRelationshipsSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret PromotedPurchaseRelationshipsSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshotRelationships) GetSubscriptionOk() (*PromotedPurchaseRelationshipsSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SubscriptionAppStoreReviewScreenshotRelationships) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given PromotedPurchaseRelationshipsSubscription and assigns it to the Subscription field.
func (o *SubscriptionAppStoreReviewScreenshotRelationships) SetSubscription(v PromotedPurchaseRelationshipsSubscription) {
	o.Subscription = &v
}

func (o SubscriptionAppStoreReviewScreenshotRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAppStoreReviewScreenshotRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableSubscriptionAppStoreReviewScreenshotRelationships struct {
	value *SubscriptionAppStoreReviewScreenshotRelationships
	isSet bool
}

func (v NullableSubscriptionAppStoreReviewScreenshotRelationships) Get() *SubscriptionAppStoreReviewScreenshotRelationships {
	return v.value
}

func (v *NullableSubscriptionAppStoreReviewScreenshotRelationships) Set(val *SubscriptionAppStoreReviewScreenshotRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAppStoreReviewScreenshotRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAppStoreReviewScreenshotRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAppStoreReviewScreenshotRelationships(val *SubscriptionAppStoreReviewScreenshotRelationships) *NullableSubscriptionAppStoreReviewScreenshotRelationships {
	return &NullableSubscriptionAppStoreReviewScreenshotRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionAppStoreReviewScreenshotRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAppStoreReviewScreenshotRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


