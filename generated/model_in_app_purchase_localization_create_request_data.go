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

// checks if the InAppPurchaseLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseLocalizationCreateRequestData{}

// InAppPurchaseLocalizationCreateRequestData struct for InAppPurchaseLocalizationCreateRequestData
type InAppPurchaseLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes InAppPurchaseLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships `json:"relationships"`
}

// NewInAppPurchaseLocalizationCreateRequestData instantiates a new InAppPurchaseLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseLocalizationCreateRequestData(type_ string, attributes InAppPurchaseLocalizationCreateRequestDataAttributes, relationships InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships) *InAppPurchaseLocalizationCreateRequestData {
	this := InAppPurchaseLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewInAppPurchaseLocalizationCreateRequestDataWithDefaults instantiates a new InAppPurchaseLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseLocalizationCreateRequestDataWithDefaults() *InAppPurchaseLocalizationCreateRequestData {
	this := InAppPurchaseLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InAppPurchaseLocalizationCreateRequestData) GetAttributes() InAppPurchaseLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret InAppPurchaseLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationCreateRequestData) GetAttributesOk() (*InAppPurchaseLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InAppPurchaseLocalizationCreateRequestData) SetAttributes(v InAppPurchaseLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *InAppPurchaseLocalizationCreateRequestData) GetRelationships() InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseLocalizationCreateRequestData) GetRelationshipsOk() (*InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *InAppPurchaseLocalizationCreateRequestData) SetRelationships(v InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o InAppPurchaseLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableInAppPurchaseLocalizationCreateRequestData struct {
	value *InAppPurchaseLocalizationCreateRequestData
	isSet bool
}

func (v NullableInAppPurchaseLocalizationCreateRequestData) Get() *InAppPurchaseLocalizationCreateRequestData {
	return v.value
}

func (v *NullableInAppPurchaseLocalizationCreateRequestData) Set(val *InAppPurchaseLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseLocalizationCreateRequestData(val *InAppPurchaseLocalizationCreateRequestData) *NullableInAppPurchaseLocalizationCreateRequestData {
	return &NullableInAppPurchaseLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableInAppPurchaseLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


