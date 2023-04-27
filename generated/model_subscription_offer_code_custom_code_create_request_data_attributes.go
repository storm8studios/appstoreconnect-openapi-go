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

// checks if the SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes{}

// SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes struct for SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes
type SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes struct {
	CustomCode string `json:"customCode"`
	NumberOfCodes int32 `json:"numberOfCodes"`
	ExpirationDate *string `json:"expirationDate,omitempty"`
}

// NewSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes instantiates a new SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes(customCode string, numberOfCodes int32) *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes {
	this := SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes{}
	this.CustomCode = customCode
	this.NumberOfCodes = numberOfCodes
	return &this
}

// NewSubscriptionOfferCodeCustomCodeCreateRequestDataAttributesWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeCreateRequestDataAttributesWithDefaults() *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes {
	this := SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes{}
	return &this
}

// GetCustomCode returns the CustomCode field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) GetCustomCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomCode
}

// GetCustomCodeOk returns a tuple with the CustomCode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) GetCustomCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomCode, true
}

// SetCustomCode sets field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) SetCustomCode(v string) {
	o.CustomCode = v
}

// GetNumberOfCodes returns the NumberOfCodes field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) GetNumberOfCodes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfCodes
}

// GetNumberOfCodesOk returns a tuple with the NumberOfCodes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) GetNumberOfCodesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfCodes, true
}

// SetNumberOfCodes sets field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) SetNumberOfCodes(v int32) {
	o.NumberOfCodes = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

func (o SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customCode"] = o.CustomCode
	toSerialize["numberOfCodes"] = o.NumberOfCodes
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes struct {
	value *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) Get() *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) Set(val *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes(val *SubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes {
	return &NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


