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

// checks if the AppCustomProductPageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageResponse{}

// AppCustomProductPageResponse struct for AppCustomProductPageResponse
type AppCustomProductPageResponse struct {
	Data AppCustomProductPage `json:"data"`
	Included []AppCustomProductPagesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppCustomProductPageResponse instantiates a new AppCustomProductPageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageResponse(data AppCustomProductPage, links DocumentLinks) *AppCustomProductPageResponse {
	this := AppCustomProductPageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppCustomProductPageResponseWithDefaults instantiates a new AppCustomProductPageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageResponseWithDefaults() *AppCustomProductPageResponse {
	this := AppCustomProductPageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageResponse) GetData() AppCustomProductPage {
	if o == nil {
		var ret AppCustomProductPage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageResponse) GetDataOk() (*AppCustomProductPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageResponse) SetData(v AppCustomProductPage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppCustomProductPageResponse) GetIncluded() []AppCustomProductPagesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppCustomProductPagesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageResponse) GetIncludedOk() ([]AppCustomProductPagesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppCustomProductPageResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppCustomProductPagesResponseIncludedInner and assigns it to the Included field.
func (o *AppCustomProductPageResponse) SetIncluded(v []AppCustomProductPagesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppCustomProductPageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppCustomProductPageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppCustomProductPageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppCustomProductPageResponse struct {
	value *AppCustomProductPageResponse
	isSet bool
}

func (v NullableAppCustomProductPageResponse) Get() *AppCustomProductPageResponse {
	return v.value
}

func (v *NullableAppCustomProductPageResponse) Set(val *AppCustomProductPageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageResponse(val *AppCustomProductPageResponse) *NullableAppCustomProductPageResponse {
	return &NullableAppCustomProductPageResponse{value: val, isSet: true}
}

func (v NullableAppCustomProductPageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


