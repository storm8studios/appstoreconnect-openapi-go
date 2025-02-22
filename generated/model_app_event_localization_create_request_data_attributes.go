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

// checks if the AppEventLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationCreateRequestDataAttributes{}

// AppEventLocalizationCreateRequestDataAttributes struct for AppEventLocalizationCreateRequestDataAttributes
type AppEventLocalizationCreateRequestDataAttributes struct {
	Locale string `json:"locale"`
	Name *string `json:"name,omitempty"`
	ShortDescription *string `json:"shortDescription,omitempty"`
	LongDescription *string `json:"longDescription,omitempty"`
}

// NewAppEventLocalizationCreateRequestDataAttributes instantiates a new AppEventLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationCreateRequestDataAttributes(locale string) *AppEventLocalizationCreateRequestDataAttributes {
	this := AppEventLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewAppEventLocalizationCreateRequestDataAttributesWithDefaults instantiates a new AppEventLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationCreateRequestDataAttributesWithDefaults() *AppEventLocalizationCreateRequestDataAttributes {
	this := AppEventLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *AppEventLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AppEventLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AppEventLocalizationCreateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *AppEventLocalizationCreateRequestDataAttributes) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetLongDescription returns the LongDescription field value if set, zero value otherwise.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetLongDescription() string {
	if o == nil || IsNil(o.LongDescription) {
		var ret string
		return ret
	}
	return *o.LongDescription
}

// GetLongDescriptionOk returns a tuple with the LongDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) GetLongDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.LongDescription) {
		return nil, false
	}
	return o.LongDescription, true
}

// HasLongDescription returns a boolean if a field has been set.
func (o *AppEventLocalizationCreateRequestDataAttributes) HasLongDescription() bool {
	if o != nil && !IsNil(o.LongDescription) {
		return true
	}

	return false
}

// SetLongDescription gets a reference to the given string and assigns it to the LongDescription field.
func (o *AppEventLocalizationCreateRequestDataAttributes) SetLongDescription(v string) {
	o.LongDescription = &v
}

func (o AppEventLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locale"] = o.Locale
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.LongDescription) {
		toSerialize["longDescription"] = o.LongDescription
	}
	return toSerialize, nil
}

type NullableAppEventLocalizationCreateRequestDataAttributes struct {
	value *AppEventLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppEventLocalizationCreateRequestDataAttributes) Get() *AppEventLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEventLocalizationCreateRequestDataAttributes) Set(val *AppEventLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationCreateRequestDataAttributes(val *AppEventLocalizationCreateRequestDataAttributes) *NullableAppEventLocalizationCreateRequestDataAttributes {
	return &NullableAppEventLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEventLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


