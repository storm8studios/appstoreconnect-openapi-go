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

// checks if the SubscriptionIntroductoryOfferUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferUpdateRequestDataAttributes{}

// SubscriptionIntroductoryOfferUpdateRequestDataAttributes struct for SubscriptionIntroductoryOfferUpdateRequestDataAttributes
type SubscriptionIntroductoryOfferUpdateRequestDataAttributes struct {
	EndDate *string `json:"endDate,omitempty"`
}

// NewSubscriptionIntroductoryOfferUpdateRequestDataAttributes instantiates a new SubscriptionIntroductoryOfferUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferUpdateRequestDataAttributes() *SubscriptionIntroductoryOfferUpdateRequestDataAttributes {
	this := SubscriptionIntroductoryOfferUpdateRequestDataAttributes{}
	return &this
}

// NewSubscriptionIntroductoryOfferUpdateRequestDataAttributesWithDefaults instantiates a new SubscriptionIntroductoryOfferUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferUpdateRequestDataAttributesWithDefaults() *SubscriptionIntroductoryOfferUpdateRequestDataAttributes {
	this := SubscriptionIntroductoryOfferUpdateRequestDataAttributes{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferUpdateRequestDataAttributes) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferUpdateRequestDataAttributes) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferUpdateRequestDataAttributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SubscriptionIntroductoryOfferUpdateRequestDataAttributes) SetEndDate(v string) {
	o.EndDate = &v
}

func (o SubscriptionIntroductoryOfferUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes struct {
	value *SubscriptionIntroductoryOfferUpdateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes) Get() *SubscriptionIntroductoryOfferUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes) Set(val *SubscriptionIntroductoryOfferUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes(val *SubscriptionIntroductoryOfferUpdateRequestDataAttributes) *NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes {
	return &NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


