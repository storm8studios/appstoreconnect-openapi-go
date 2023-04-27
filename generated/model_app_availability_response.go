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

// checks if the AppAvailabilityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityResponse{}

// AppAvailabilityResponse struct for AppAvailabilityResponse
type AppAvailabilityResponse struct {
	Data AppAvailability `json:"data"`
	Included []AppAvailabilityResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewAppAvailabilityResponse instantiates a new AppAvailabilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityResponse(data AppAvailability, links DocumentLinks) *AppAvailabilityResponse {
	this := AppAvailabilityResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppAvailabilityResponseWithDefaults instantiates a new AppAvailabilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityResponseWithDefaults() *AppAvailabilityResponse {
	this := AppAvailabilityResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppAvailabilityResponse) GetData() AppAvailability {
	if o == nil {
		var ret AppAvailability
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityResponse) GetDataOk() (*AppAvailability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppAvailabilityResponse) SetData(v AppAvailability) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppAvailabilityResponse) GetIncluded() []AppAvailabilityResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppAvailabilityResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityResponse) GetIncludedOk() ([]AppAvailabilityResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppAvailabilityResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppAvailabilityResponseIncludedInner and assigns it to the Included field.
func (o *AppAvailabilityResponse) SetIncluded(v []AppAvailabilityResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppAvailabilityResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppAvailabilityResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppAvailabilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppAvailabilityResponse struct {
	value *AppAvailabilityResponse
	isSet bool
}

func (v NullableAppAvailabilityResponse) Get() *AppAvailabilityResponse {
	return v.value
}

func (v *NullableAppAvailabilityResponse) Set(val *AppAvailabilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityResponse(val *AppAvailabilityResponse) *NullableAppAvailabilityResponse {
	return &NullableAppAvailabilityResponse{value: val, isSet: true}
}

func (v NullableAppAvailabilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


