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

// checks if the BetaLicenseAgreementAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaLicenseAgreementAttributes{}

// BetaLicenseAgreementAttributes struct for BetaLicenseAgreementAttributes
type BetaLicenseAgreementAttributes struct {
	AgreementText *string `json:"agreementText,omitempty"`
}

// NewBetaLicenseAgreementAttributes instantiates a new BetaLicenseAgreementAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreementAttributes() *BetaLicenseAgreementAttributes {
	this := BetaLicenseAgreementAttributes{}
	return &this
}

// NewBetaLicenseAgreementAttributesWithDefaults instantiates a new BetaLicenseAgreementAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementAttributesWithDefaults() *BetaLicenseAgreementAttributes {
	this := BetaLicenseAgreementAttributes{}
	return &this
}

// GetAgreementText returns the AgreementText field value if set, zero value otherwise.
func (o *BetaLicenseAgreementAttributes) GetAgreementText() string {
	if o == nil || IsNil(o.AgreementText) {
		var ret string
		return ret
	}
	return *o.AgreementText
}

// GetAgreementTextOk returns a tuple with the AgreementText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementAttributes) GetAgreementTextOk() (*string, bool) {
	if o == nil || IsNil(o.AgreementText) {
		return nil, false
	}
	return o.AgreementText, true
}

// HasAgreementText returns a boolean if a field has been set.
func (o *BetaLicenseAgreementAttributes) HasAgreementText() bool {
	if o != nil && !IsNil(o.AgreementText) {
		return true
	}

	return false
}

// SetAgreementText gets a reference to the given string and assigns it to the AgreementText field.
func (o *BetaLicenseAgreementAttributes) SetAgreementText(v string) {
	o.AgreementText = &v
}

func (o BetaLicenseAgreementAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaLicenseAgreementAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgreementText) {
		toSerialize["agreementText"] = o.AgreementText
	}
	return toSerialize, nil
}

type NullableBetaLicenseAgreementAttributes struct {
	value *BetaLicenseAgreementAttributes
	isSet bool
}

func (v NullableBetaLicenseAgreementAttributes) Get() *BetaLicenseAgreementAttributes {
	return v.value
}

func (v *NullableBetaLicenseAgreementAttributes) Set(val *BetaLicenseAgreementAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreementAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreementAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreementAttributes(val *BetaLicenseAgreementAttributes) *NullableBetaLicenseAgreementAttributes {
	return &NullableBetaLicenseAgreementAttributes{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreementAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreementAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


