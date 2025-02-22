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

// checks if the ProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileResponse{}

// ProfileResponse struct for ProfileResponse
type ProfileResponse struct {
	Data Profile `json:"data"`
	Included []ProfilesResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewProfileResponse instantiates a new ProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileResponse(data Profile, links DocumentLinks) *ProfileResponse {
	this := ProfileResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewProfileResponseWithDefaults instantiates a new ProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileResponseWithDefaults() *ProfileResponse {
	this := ProfileResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ProfileResponse) GetData() Profile {
	if o == nil {
		var ret Profile
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProfileResponse) GetDataOk() (*Profile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ProfileResponse) SetData(v Profile) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *ProfileResponse) GetIncluded() []ProfilesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []ProfilesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileResponse) GetIncludedOk() ([]ProfilesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *ProfileResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []ProfilesResponseIncludedInner and assigns it to the Included field.
func (o *ProfileResponse) SetIncluded(v []ProfilesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *ProfileResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ProfileResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ProfileResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o ProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableProfileResponse struct {
	value *ProfileResponse
	isSet bool
}

func (v NullableProfileResponse) Get() *ProfileResponse {
	return v.value
}

func (v *NullableProfileResponse) Set(val *ProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileResponse(val *ProfileResponse) *NullableProfileResponse {
	return &NullableProfileResponse{value: val, isSet: true}
}

func (v NullableProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


