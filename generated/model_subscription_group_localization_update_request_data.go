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

// checks if the SubscriptionGroupLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationUpdateRequestData{}

// SubscriptionGroupLocalizationUpdateRequestData struct for SubscriptionGroupLocalizationUpdateRequestData
type SubscriptionGroupLocalizationUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *SubscriptionGroupLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewSubscriptionGroupLocalizationUpdateRequestData instantiates a new SubscriptionGroupLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationUpdateRequestData(type_ string, id string) *SubscriptionGroupLocalizationUpdateRequestData {
	this := SubscriptionGroupLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewSubscriptionGroupLocalizationUpdateRequestDataWithDefaults instantiates a new SubscriptionGroupLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationUpdateRequestDataWithDefaults() *SubscriptionGroupLocalizationUpdateRequestData {
	this := SubscriptionGroupLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionGroupLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionGroupLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionGroupLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionGroupLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalizationUpdateRequestData) GetAttributes() SubscriptionGroupLocalizationUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret SubscriptionGroupLocalizationUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationUpdateRequestData) GetAttributesOk() (*SubscriptionGroupLocalizationUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given SubscriptionGroupLocalizationUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *SubscriptionGroupLocalizationUpdateRequestData) SetAttributes(v SubscriptionGroupLocalizationUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o SubscriptionGroupLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableSubscriptionGroupLocalizationUpdateRequestData struct {
	value *SubscriptionGroupLocalizationUpdateRequestData
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationUpdateRequestData) Get() *SubscriptionGroupLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationUpdateRequestData) Set(val *SubscriptionGroupLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationUpdateRequestData(val *SubscriptionGroupLocalizationUpdateRequestData) *NullableSubscriptionGroupLocalizationUpdateRequestData {
	return &NullableSubscriptionGroupLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


