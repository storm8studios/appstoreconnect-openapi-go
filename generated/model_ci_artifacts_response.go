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

// checks if the CiArtifactsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiArtifactsResponse{}

// CiArtifactsResponse struct for CiArtifactsResponse
type CiArtifactsResponse struct {
	Data []CiArtifact `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewCiArtifactsResponse instantiates a new CiArtifactsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiArtifactsResponse(data []CiArtifact, links PagedDocumentLinks) *CiArtifactsResponse {
	this := CiArtifactsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiArtifactsResponseWithDefaults instantiates a new CiArtifactsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiArtifactsResponseWithDefaults() *CiArtifactsResponse {
	this := CiArtifactsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiArtifactsResponse) GetData() []CiArtifact {
	if o == nil {
		var ret []CiArtifact
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiArtifactsResponse) GetDataOk() ([]CiArtifact, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CiArtifactsResponse) SetData(v []CiArtifact) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *CiArtifactsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiArtifactsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiArtifactsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CiArtifactsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiArtifactsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CiArtifactsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *CiArtifactsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o CiArtifactsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiArtifactsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableCiArtifactsResponse struct {
	value *CiArtifactsResponse
	isSet bool
}

func (v NullableCiArtifactsResponse) Get() *CiArtifactsResponse {
	return v.value
}

func (v *NullableCiArtifactsResponse) Set(val *CiArtifactsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiArtifactsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiArtifactsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiArtifactsResponse(val *CiArtifactsResponse) *NullableCiArtifactsResponse {
	return &NullableCiArtifactsResponse{value: val, isSet: true}
}

func (v NullableCiArtifactsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiArtifactsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


