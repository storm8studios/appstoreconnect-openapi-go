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

// checks if the AppResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppResponse{}

// AppResponse struct for AppResponse
type AppResponse struct {
	Data App `json:"data"`
	Included []AppsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppResponse instantiates a new AppResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppResponse(data App, links DocumentLinks) *AppResponse {
	this := AppResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppResponseWithDefaults instantiates a new AppResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppResponseWithDefaults() *AppResponse {
	this := AppResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppResponse) GetData() App {
	if o == nil {
		var ret App
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppResponse) GetDataOk() (*App, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppResponse) SetData(v App) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppResponse) GetIncluded() []AppsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppResponse) GetIncludedOk() ([]AppsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppsResponseIncludedInner and assigns it to the Included field.
func (o *AppResponse) SetIncluded(v []AppsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppResponse struct {
	value *AppResponse
	isSet bool
}

func (v NullableAppResponse) Get() *AppResponse {
	return v.value
}

func (v *NullableAppResponse) Set(val *AppResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppResponse(val *AppResponse) *NullableAppResponse {
	return &NullableAppResponse{value: val, isSet: true}
}

func (v NullableAppResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


