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

// checks if the BuildIndividualTestersLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildIndividualTestersLinkagesResponse{}

// BuildIndividualTestersLinkagesResponse struct for BuildIndividualTestersLinkagesResponse
type BuildIndividualTestersLinkagesResponse struct {
	Data []BetaGroupRelationshipsBetaTestersDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewBuildIndividualTestersLinkagesResponse instantiates a new BuildIndividualTestersLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildIndividualTestersLinkagesResponse(data []BetaGroupRelationshipsBetaTestersDataInner, links PagedDocumentLinks) *BuildIndividualTestersLinkagesResponse {
	this := BuildIndividualTestersLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildIndividualTestersLinkagesResponseWithDefaults instantiates a new BuildIndividualTestersLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildIndividualTestersLinkagesResponseWithDefaults() *BuildIndividualTestersLinkagesResponse {
	this := BuildIndividualTestersLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildIndividualTestersLinkagesResponse) GetData() []BetaGroupRelationshipsBetaTestersDataInner {
	if o == nil {
		var ret []BetaGroupRelationshipsBetaTestersDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildIndividualTestersLinkagesResponse) GetDataOk() ([]BetaGroupRelationshipsBetaTestersDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BuildIndividualTestersLinkagesResponse) SetData(v []BetaGroupRelationshipsBetaTestersDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BuildIndividualTestersLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildIndividualTestersLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildIndividualTestersLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BuildIndividualTestersLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildIndividualTestersLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BuildIndividualTestersLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BuildIndividualTestersLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BuildIndividualTestersLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildIndividualTestersLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBuildIndividualTestersLinkagesResponse struct {
	value *BuildIndividualTestersLinkagesResponse
	isSet bool
}

func (v NullableBuildIndividualTestersLinkagesResponse) Get() *BuildIndividualTestersLinkagesResponse {
	return v.value
}

func (v *NullableBuildIndividualTestersLinkagesResponse) Set(val *BuildIndividualTestersLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildIndividualTestersLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildIndividualTestersLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildIndividualTestersLinkagesResponse(val *BuildIndividualTestersLinkagesResponse) *NullableBuildIndividualTestersLinkagesResponse {
	return &NullableBuildIndividualTestersLinkagesResponse{value: val, isSet: true}
}

func (v NullableBuildIndividualTestersLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildIndividualTestersLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


