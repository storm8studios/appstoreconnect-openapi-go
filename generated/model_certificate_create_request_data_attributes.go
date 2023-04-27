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

// checks if the CertificateCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateCreateRequestDataAttributes{}

// CertificateCreateRequestDataAttributes struct for CertificateCreateRequestDataAttributes
type CertificateCreateRequestDataAttributes struct {
	CsrContent string `json:"csrContent"`
	CertificateType CertificateType `json:"certificateType"`
}

// NewCertificateCreateRequestDataAttributes instantiates a new CertificateCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateCreateRequestDataAttributes(csrContent string, certificateType CertificateType) *CertificateCreateRequestDataAttributes {
	this := CertificateCreateRequestDataAttributes{}
	this.CsrContent = csrContent
	this.CertificateType = certificateType
	return &this
}

// NewCertificateCreateRequestDataAttributesWithDefaults instantiates a new CertificateCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateCreateRequestDataAttributesWithDefaults() *CertificateCreateRequestDataAttributes {
	this := CertificateCreateRequestDataAttributes{}
	return &this
}

// GetCsrContent returns the CsrContent field value
func (o *CertificateCreateRequestDataAttributes) GetCsrContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CsrContent
}

// GetCsrContentOk returns a tuple with the CsrContent field value
// and a boolean to check if the value has been set.
func (o *CertificateCreateRequestDataAttributes) GetCsrContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CsrContent, true
}

// SetCsrContent sets field value
func (o *CertificateCreateRequestDataAttributes) SetCsrContent(v string) {
	o.CsrContent = v
}

// GetCertificateType returns the CertificateType field value
func (o *CertificateCreateRequestDataAttributes) GetCertificateType() CertificateType {
	if o == nil {
		var ret CertificateType
		return ret
	}

	return o.CertificateType
}

// GetCertificateTypeOk returns a tuple with the CertificateType field value
// and a boolean to check if the value has been set.
func (o *CertificateCreateRequestDataAttributes) GetCertificateTypeOk() (*CertificateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateType, true
}

// SetCertificateType sets field value
func (o *CertificateCreateRequestDataAttributes) SetCertificateType(v CertificateType) {
	o.CertificateType = v
}

func (o CertificateCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["csrContent"] = o.CsrContent
	toSerialize["certificateType"] = o.CertificateType
	return toSerialize, nil
}

type NullableCertificateCreateRequestDataAttributes struct {
	value *CertificateCreateRequestDataAttributes
	isSet bool
}

func (v NullableCertificateCreateRequestDataAttributes) Get() *CertificateCreateRequestDataAttributes {
	return v.value
}

func (v *NullableCertificateCreateRequestDataAttributes) Set(val *CertificateCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateCreateRequestDataAttributes(val *CertificateCreateRequestDataAttributes) *NullableCertificateCreateRequestDataAttributes {
	return &NullableCertificateCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableCertificateCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


