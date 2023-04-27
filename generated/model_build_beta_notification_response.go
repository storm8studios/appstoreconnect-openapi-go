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

// checks if the BuildBetaNotificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaNotificationResponse{}

// BuildBetaNotificationResponse struct for BuildBetaNotificationResponse
type BuildBetaNotificationResponse struct {
	Data BuildBetaNotification `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewBuildBetaNotificationResponse instantiates a new BuildBetaNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaNotificationResponse(data BuildBetaNotification, links DocumentLinks) *BuildBetaNotificationResponse {
	this := BuildBetaNotificationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildBetaNotificationResponseWithDefaults instantiates a new BuildBetaNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaNotificationResponseWithDefaults() *BuildBetaNotificationResponse {
	this := BuildBetaNotificationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaNotificationResponse) GetData() BuildBetaNotification {
	if o == nil {
		var ret BuildBetaNotification
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaNotificationResponse) GetDataOk() (*BuildBetaNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildBetaNotificationResponse) SetData(v BuildBetaNotification) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BuildBetaNotificationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildBetaNotificationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildBetaNotificationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BuildBetaNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBuildBetaNotificationResponse struct {
	value *BuildBetaNotificationResponse
	isSet bool
}

func (v NullableBuildBetaNotificationResponse) Get() *BuildBetaNotificationResponse {
	return v.value
}

func (v *NullableBuildBetaNotificationResponse) Set(val *BuildBetaNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaNotificationResponse(val *BuildBetaNotificationResponse) *NullableBuildBetaNotificationResponse {
	return &NullableBuildBetaNotificationResponse{value: val, isSet: true}
}

func (v NullableBuildBetaNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


