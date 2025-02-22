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

// checks if the AppInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoResponse{}

// AppInfoResponse struct for AppInfoResponse
type AppInfoResponse struct {
	Data AppInfo `json:"data"`
	Included []AppInfosResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppInfoResponse instantiates a new AppInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoResponse(data AppInfo, links DocumentLinks) *AppInfoResponse {
	this := AppInfoResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppInfoResponseWithDefaults instantiates a new AppInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoResponseWithDefaults() *AppInfoResponse {
	this := AppInfoResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppInfoResponse) GetData() AppInfo {
	if o == nil {
		var ret AppInfo
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppInfoResponse) GetDataOk() (*AppInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppInfoResponse) SetData(v AppInfo) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppInfoResponse) GetIncluded() []AppInfosResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppInfosResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoResponse) GetIncludedOk() ([]AppInfosResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppInfoResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppInfosResponseIncludedInner and assigns it to the Included field.
func (o *AppInfoResponse) SetIncluded(v []AppInfosResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppInfoResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppInfoResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppInfoResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppInfoResponse struct {
	value *AppInfoResponse
	isSet bool
}

func (v NullableAppInfoResponse) Get() *AppInfoResponse {
	return v.value
}

func (v *NullableAppInfoResponse) Set(val *AppInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoResponse(val *AppInfoResponse) *NullableAppInfoResponse {
	return &NullableAppInfoResponse{value: val, isSet: true}
}

func (v NullableAppInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


