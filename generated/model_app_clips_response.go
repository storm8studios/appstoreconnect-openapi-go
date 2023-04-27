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

// checks if the AppClipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipsResponse{}

// AppClipsResponse struct for AppClipsResponse
type AppClipsResponse struct {
	Data []AppClip `json:"data"`
	Included []AppClipsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppClipsResponse instantiates a new AppClipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipsResponse(data []AppClip, links PagedDocumentLinks) *AppClipsResponse {
	this := AppClipsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipsResponseWithDefaults instantiates a new AppClipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipsResponseWithDefaults() *AppClipsResponse {
	this := AppClipsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipsResponse) GetData() []AppClip {
	if o == nil {
		var ret []AppClip
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipsResponse) GetDataOk() ([]AppClip, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppClipsResponse) SetData(v []AppClip) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppClipsResponse) GetIncluded() []AppClipsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppClipsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipsResponse) GetIncludedOk() ([]AppClipsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppClipsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppClipsResponseIncludedInner and assigns it to the Included field.
func (o *AppClipsResponse) SetIncluded(v []AppClipsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppClipsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppClipsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppClipsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppClipsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppClipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableAppClipsResponse struct {
	value *AppClipsResponse
	isSet bool
}

func (v NullableAppClipsResponse) Get() *AppClipsResponse {
	return v.value
}

func (v *NullableAppClipsResponse) Set(val *AppClipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipsResponse(val *AppClipsResponse) *NullableAppClipsResponse {
	return &NullableAppClipsResponse{value: val, isSet: true}
}

func (v NullableAppClipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


