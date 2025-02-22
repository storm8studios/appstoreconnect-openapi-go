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

// checks if the BetaBuildLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildLocalizationResponse{}

// BetaBuildLocalizationResponse struct for BetaBuildLocalizationResponse
type BetaBuildLocalizationResponse struct {
	Data BetaBuildLocalization `json:"data"`
	Included []Build `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewBetaBuildLocalizationResponse instantiates a new BetaBuildLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationResponse(data BetaBuildLocalization, links DocumentLinks) *BetaBuildLocalizationResponse {
	this := BetaBuildLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaBuildLocalizationResponseWithDefaults instantiates a new BetaBuildLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationResponseWithDefaults() *BetaBuildLocalizationResponse {
	this := BetaBuildLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaBuildLocalizationResponse) GetData() BetaBuildLocalization {
	if o == nil {
		var ret BetaBuildLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationResponse) GetDataOk() (*BetaBuildLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaBuildLocalizationResponse) SetData(v BetaBuildLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaBuildLocalizationResponse) GetIncluded() []Build {
	if o == nil || IsNil(o.Included) {
		var ret []Build
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationResponse) GetIncludedOk() ([]Build, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaBuildLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Build and assigns it to the Included field.
func (o *BetaBuildLocalizationResponse) SetIncluded(v []Build) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaBuildLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaBuildLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaBuildLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBetaBuildLocalizationResponse struct {
	value *BetaBuildLocalizationResponse
	isSet bool
}

func (v NullableBetaBuildLocalizationResponse) Get() *BetaBuildLocalizationResponse {
	return v.value
}

func (v *NullableBetaBuildLocalizationResponse) Set(val *BetaBuildLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationResponse(val *BetaBuildLocalizationResponse) *NullableBetaBuildLocalizationResponse {
	return &NullableBetaBuildLocalizationResponse{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


