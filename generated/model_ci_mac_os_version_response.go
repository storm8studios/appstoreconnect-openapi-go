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

// checks if the CiMacOsVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiMacOsVersionResponse{}

// CiMacOsVersionResponse struct for CiMacOsVersionResponse
type CiMacOsVersionResponse struct {
	Data CiMacOsVersion `json:"data"`
	Included []CiXcodeVersion `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewCiMacOsVersionResponse instantiates a new CiMacOsVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiMacOsVersionResponse(data CiMacOsVersion, links DocumentLinks) *CiMacOsVersionResponse {
	this := CiMacOsVersionResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiMacOsVersionResponseWithDefaults instantiates a new CiMacOsVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiMacOsVersionResponseWithDefaults() *CiMacOsVersionResponse {
	this := CiMacOsVersionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiMacOsVersionResponse) GetData() CiMacOsVersion {
	if o == nil {
		var ret CiMacOsVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionResponse) GetDataOk() (*CiMacOsVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiMacOsVersionResponse) SetData(v CiMacOsVersion) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CiMacOsVersionResponse) GetIncluded() []CiXcodeVersion {
	if o == nil || IsNil(o.Included) {
		var ret []CiXcodeVersion
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionResponse) GetIncludedOk() ([]CiXcodeVersion, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CiMacOsVersionResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CiXcodeVersion and assigns it to the Included field.
func (o *CiMacOsVersionResponse) SetIncluded(v []CiXcodeVersion) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CiMacOsVersionResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiMacOsVersionResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CiMacOsVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiMacOsVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableCiMacOsVersionResponse struct {
	value *CiMacOsVersionResponse
	isSet bool
}

func (v NullableCiMacOsVersionResponse) Get() *CiMacOsVersionResponse {
	return v.value
}

func (v *NullableCiMacOsVersionResponse) Set(val *CiMacOsVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiMacOsVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiMacOsVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiMacOsVersionResponse(val *CiMacOsVersionResponse) *NullableCiMacOsVersionResponse {
	return &NullableCiMacOsVersionResponse{value: val, isSet: true}
}

func (v NullableCiMacOsVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiMacOsVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


