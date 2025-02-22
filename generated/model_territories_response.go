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

// checks if the TerritoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoriesResponse{}

// TerritoriesResponse struct for TerritoriesResponse
type TerritoriesResponse struct {
	Data []Territory `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewTerritoriesResponse instantiates a new TerritoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoriesResponse(data []Territory, links PagedDocumentLinks) *TerritoriesResponse {
	this := TerritoriesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewTerritoriesResponseWithDefaults instantiates a new TerritoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoriesResponseWithDefaults() *TerritoriesResponse {
	this := TerritoriesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *TerritoriesResponse) GetData() []Territory {
	if o == nil {
		var ret []Territory
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TerritoriesResponse) GetDataOk() ([]Territory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *TerritoriesResponse) SetData(v []Territory) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *TerritoriesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *TerritoriesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *TerritoriesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TerritoriesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoriesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TerritoriesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *TerritoriesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o TerritoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableTerritoriesResponse struct {
	value *TerritoriesResponse
	isSet bool
}

func (v NullableTerritoriesResponse) Get() *TerritoriesResponse {
	return v.value
}

func (v *NullableTerritoriesResponse) Set(val *TerritoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoriesResponse(val *TerritoriesResponse) *NullableTerritoriesResponse {
	return &NullableTerritoriesResponse{value: val, isSet: true}
}

func (v NullableTerritoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


