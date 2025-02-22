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

// checks if the ScmPullRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmPullRequestResponse{}

// ScmPullRequestResponse struct for ScmPullRequestResponse
type ScmPullRequestResponse struct {
	Data ScmPullRequest `json:"data"`
	Included []ScmRepository `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewScmPullRequestResponse instantiates a new ScmPullRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmPullRequestResponse(data ScmPullRequest, links DocumentLinks) *ScmPullRequestResponse {
	this := ScmPullRequestResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewScmPullRequestResponseWithDefaults instantiates a new ScmPullRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmPullRequestResponseWithDefaults() *ScmPullRequestResponse {
	this := ScmPullRequestResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ScmPullRequestResponse) GetData() ScmPullRequest {
	if o == nil {
		var ret ScmPullRequest
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ScmPullRequestResponse) GetDataOk() (*ScmPullRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ScmPullRequestResponse) SetData(v ScmPullRequest) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ScmPullRequestResponse) GetIncluded() []ScmRepository {
	if o == nil || IsNil(o.Included) {
		var ret []ScmRepository
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmPullRequestResponse) GetIncludedOk() ([]ScmRepository, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ScmPullRequestResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ScmRepository and assigns it to the Included field.
func (o *ScmPullRequestResponse) SetIncluded(v []ScmRepository) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *ScmPullRequestResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ScmPullRequestResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ScmPullRequestResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o ScmPullRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmPullRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableScmPullRequestResponse struct {
	value *ScmPullRequestResponse
	isSet bool
}

func (v NullableScmPullRequestResponse) Get() *ScmPullRequestResponse {
	return v.value
}

func (v *NullableScmPullRequestResponse) Set(val *ScmPullRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScmPullRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScmPullRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmPullRequestResponse(val *ScmPullRequestResponse) *NullableScmPullRequestResponse {
	return &NullableScmPullRequestResponse{value: val, isSet: true}
}

func (v NullableScmPullRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmPullRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


