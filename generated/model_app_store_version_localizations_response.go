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

// checks if the AppStoreVersionLocalizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationsResponse{}

// AppStoreVersionLocalizationsResponse struct for AppStoreVersionLocalizationsResponse
type AppStoreVersionLocalizationsResponse struct {
	Data []AppStoreVersionLocalization `json:"data"`
	Included []AppStoreVersionLocalizationsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppStoreVersionLocalizationsResponse instantiates a new AppStoreVersionLocalizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationsResponse(data []AppStoreVersionLocalization, links PagedDocumentLinks) *AppStoreVersionLocalizationsResponse {
	this := AppStoreVersionLocalizationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionLocalizationsResponseWithDefaults instantiates a new AppStoreVersionLocalizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationsResponseWithDefaults() *AppStoreVersionLocalizationsResponse {
	this := AppStoreVersionLocalizationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationsResponse) GetData() []AppStoreVersionLocalization {
	if o == nil {
		var ret []AppStoreVersionLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsResponse) GetDataOk() ([]AppStoreVersionLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationsResponse) SetData(v []AppStoreVersionLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationsResponse) GetIncluded() []AppStoreVersionLocalizationsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppStoreVersionLocalizationsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsResponse) GetIncludedOk() ([]AppStoreVersionLocalizationsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppStoreVersionLocalizationsResponseIncludedInner and assigns it to the Included field.
func (o *AppStoreVersionLocalizationsResponse) SetIncluded(v []AppStoreVersionLocalizationsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionLocalizationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionLocalizationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreVersionLocalizationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppStoreVersionLocalizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppStoreVersionLocalizationsResponse struct {
	value *AppStoreVersionLocalizationsResponse
	isSet bool
}

func (v NullableAppStoreVersionLocalizationsResponse) Get() *AppStoreVersionLocalizationsResponse {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationsResponse) Set(val *AppStoreVersionLocalizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationsResponse(val *AppStoreVersionLocalizationsResponse) *NullableAppStoreVersionLocalizationsResponse {
	return &NullableAppStoreVersionLocalizationsResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


