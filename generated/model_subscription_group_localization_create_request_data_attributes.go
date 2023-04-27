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

// checks if the SubscriptionGroupLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationCreateRequestDataAttributes{}

// SubscriptionGroupLocalizationCreateRequestDataAttributes struct for SubscriptionGroupLocalizationCreateRequestDataAttributes
type SubscriptionGroupLocalizationCreateRequestDataAttributes struct {
	Name string `json:"name"`
	CustomAppName *string `json:"customAppName,omitempty"`
	Locale string `json:"locale"`
}

// NewSubscriptionGroupLocalizationCreateRequestDataAttributes instantiates a new SubscriptionGroupLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationCreateRequestDataAttributes(name string, locale string) *SubscriptionGroupLocalizationCreateRequestDataAttributes {
	this := SubscriptionGroupLocalizationCreateRequestDataAttributes{}
	this.Name = name
	this.Locale = locale
	return &this
}

// NewSubscriptionGroupLocalizationCreateRequestDataAttributesWithDefaults instantiates a new SubscriptionGroupLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationCreateRequestDataAttributesWithDefaults() *SubscriptionGroupLocalizationCreateRequestDataAttributes {
	this := SubscriptionGroupLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCustomAppName returns the CustomAppName field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) GetCustomAppName() string {
	if o == nil || IsNil(o.CustomAppName) {
		var ret string
		return ret
	}
	return *o.CustomAppName
}

// GetCustomAppNameOk returns a tuple with the CustomAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) GetCustomAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomAppName) {
		return nil, false
	}
	return o.CustomAppName, true
}

// HasCustomAppName returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) HasCustomAppName() bool {
	if o != nil && !IsNil(o.CustomAppName) {
		return true
	}

	return false
}

// SetCustomAppName gets a reference to the given string and assigns it to the CustomAppName field.
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) SetCustomAppName(v string) {
	o.CustomAppName = &v
}

// GetLocale returns the Locale field value
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *SubscriptionGroupLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

func (o SubscriptionGroupLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.CustomAppName) {
		toSerialize["customAppName"] = o.CustomAppName
	}
	toSerialize["locale"] = o.Locale
	return toSerialize, nil
}

type NullableSubscriptionGroupLocalizationCreateRequestDataAttributes struct {
	value *SubscriptionGroupLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataAttributes) Get() *SubscriptionGroupLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataAttributes) Set(val *SubscriptionGroupLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationCreateRequestDataAttributes(val *SubscriptionGroupLocalizationCreateRequestDataAttributes) *NullableSubscriptionGroupLocalizationCreateRequestDataAttributes {
	return &NullableSubscriptionGroupLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


