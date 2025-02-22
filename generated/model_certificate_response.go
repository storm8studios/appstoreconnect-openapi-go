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

// checks if the CertificateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateResponse{}

// CertificateResponse struct for CertificateResponse
type CertificateResponse struct {
	Data Certificate `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewCertificateResponse instantiates a new CertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateResponse(data Certificate, links DocumentLinks) *CertificateResponse {
	this := CertificateResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCertificateResponseWithDefaults instantiates a new CertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateResponseWithDefaults() *CertificateResponse {
	this := CertificateResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CertificateResponse) GetData() Certificate {
	if o == nil {
		var ret Certificate
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetDataOk() (*Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CertificateResponse) SetData(v Certificate) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *CertificateResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CertificateResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableCertificateResponse struct {
	value *CertificateResponse
	isSet bool
}

func (v NullableCertificateResponse) Get() *CertificateResponse {
	return v.value
}

func (v *NullableCertificateResponse) Set(val *CertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateResponse(val *CertificateResponse) *NullableCertificateResponse {
	return &NullableCertificateResponse{value: val, isSet: true}
}

func (v NullableCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


