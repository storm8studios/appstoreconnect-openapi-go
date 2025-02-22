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

// checks if the AppEventLocalizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationsResponse{}

// AppEventLocalizationsResponse struct for AppEventLocalizationsResponse
type AppEventLocalizationsResponse struct {
	Data []AppEventLocalization `json:"data"`
	Included []AppEventLocalizationsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppEventLocalizationsResponse instantiates a new AppEventLocalizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationsResponse(data []AppEventLocalization, links PagedDocumentLinks) *AppEventLocalizationsResponse {
	this := AppEventLocalizationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppEventLocalizationsResponseWithDefaults instantiates a new AppEventLocalizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationsResponseWithDefaults() *AppEventLocalizationsResponse {
	this := AppEventLocalizationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventLocalizationsResponse) GetData() []AppEventLocalization {
	if o == nil {
		var ret []AppEventLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationsResponse) GetDataOk() ([]AppEventLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppEventLocalizationsResponse) SetData(v []AppEventLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppEventLocalizationsResponse) GetIncluded() []AppEventLocalizationsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppEventLocalizationsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationsResponse) GetIncludedOk() ([]AppEventLocalizationsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppEventLocalizationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppEventLocalizationsResponseIncludedInner and assigns it to the Included field.
func (o *AppEventLocalizationsResponse) SetIncluded(v []AppEventLocalizationsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppEventLocalizationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppEventLocalizationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppEventLocalizationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppEventLocalizationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppEventLocalizationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppEventLocalizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppEventLocalizationsResponse struct {
	value *AppEventLocalizationsResponse
	isSet bool
}

func (v NullableAppEventLocalizationsResponse) Get() *AppEventLocalizationsResponse {
	return v.value
}

func (v *NullableAppEventLocalizationsResponse) Set(val *AppEventLocalizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationsResponse(val *AppEventLocalizationsResponse) *NullableAppEventLocalizationsResponse {
	return &NullableAppEventLocalizationsResponse{value: val, isSet: true}
}

func (v NullableAppEventLocalizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


