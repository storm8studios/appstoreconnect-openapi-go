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

// checks if the SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint{}

// SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint struct for SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint
type SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint struct {
	Links *AppAvailabilityRelationshipsAppLinks `json:"links,omitempty"`
	Data *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData `json:"data,omitempty"`
}

// NewSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint instantiates a new SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint() *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint {
	this := SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint{}
	return &this
}

// NewSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointWithDefaults instantiates a new SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointWithDefaults() *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint {
	this := SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) GetLinks() AppAvailabilityRelationshipsAppLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityRelationshipsAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) GetLinksOk() (*AppAvailabilityRelationshipsAppLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityRelationshipsAppLinks and assigns it to the Links field.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) SetLinks(v AppAvailabilityRelationshipsAppLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) GetData() SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData {
	if o == nil || IsNil(o.Data) {
		var ret SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) GetDataOk() (*SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData and assigns it to the Data field.
func (o *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) SetData(v SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData) {
	o.Data = &v
}

func (o SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint struct {
	value *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) Get() *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) Set(val *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint(val *SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) *NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint {
	return &NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferRelationshipsSubscriptionPricePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


