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

// checks if the BuildResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildResponse{}

// BuildResponse struct for BuildResponse
type BuildResponse struct {
	Data Build `json:"data"`
	Included []BuildsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewBuildResponse instantiates a new BuildResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildResponse(data Build, links DocumentLinks) *BuildResponse {
	this := BuildResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildResponseWithDefaults instantiates a new BuildResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildResponseWithDefaults() *BuildResponse {
	this := BuildResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildResponse) GetData() Build {
	if o == nil {
		var ret Build
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildResponse) GetDataOk() (*Build, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildResponse) SetData(v Build) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BuildResponse) GetIncluded() []BuildsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []BuildsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildResponse) GetIncludedOk() ([]BuildsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BuildResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BuildsResponseIncludedInner and assigns it to the Included field.
func (o *BuildResponse) SetIncluded(v []BuildsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BuildResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BuildResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBuildResponse struct {
	value *BuildResponse
	isSet bool
}

func (v NullableBuildResponse) Get() *BuildResponse {
	return v.value
}

func (v *NullableBuildResponse) Set(val *BuildResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildResponse(val *BuildResponse) *NullableBuildResponse {
	return &NullableBuildResponse{value: val, isSet: true}
}

func (v NullableBuildResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


