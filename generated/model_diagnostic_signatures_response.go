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

// checks if the DiagnosticSignaturesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticSignaturesResponse{}

// DiagnosticSignaturesResponse struct for DiagnosticSignaturesResponse
type DiagnosticSignaturesResponse struct {
	Data []DiagnosticSignature `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewDiagnosticSignaturesResponse instantiates a new DiagnosticSignaturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticSignaturesResponse(data []DiagnosticSignature, links PagedDocumentLinks) *DiagnosticSignaturesResponse {
	this := DiagnosticSignaturesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewDiagnosticSignaturesResponseWithDefaults instantiates a new DiagnosticSignaturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticSignaturesResponseWithDefaults() *DiagnosticSignaturesResponse {
	this := DiagnosticSignaturesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DiagnosticSignaturesResponse) GetData() []DiagnosticSignature {
	if o == nil {
		var ret []DiagnosticSignature
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSignaturesResponse) GetDataOk() ([]DiagnosticSignature, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *DiagnosticSignaturesResponse) SetData(v []DiagnosticSignature) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *DiagnosticSignaturesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DiagnosticSignaturesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DiagnosticSignaturesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DiagnosticSignaturesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignaturesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DiagnosticSignaturesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *DiagnosticSignaturesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o DiagnosticSignaturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticSignaturesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableDiagnosticSignaturesResponse struct {
	value *DiagnosticSignaturesResponse
	isSet bool
}

func (v NullableDiagnosticSignaturesResponse) Get() *DiagnosticSignaturesResponse {
	return v.value
}

func (v *NullableDiagnosticSignaturesResponse) Set(val *DiagnosticSignaturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticSignaturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticSignaturesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticSignaturesResponse(val *DiagnosticSignaturesResponse) *NullableDiagnosticSignaturesResponse {
	return &NullableDiagnosticSignaturesResponse{value: val, isSet: true}
}

func (v NullableDiagnosticSignaturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticSignaturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


