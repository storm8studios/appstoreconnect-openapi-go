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

// checks if the BetaLicenseAgreementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaLicenseAgreementResponse{}

// BetaLicenseAgreementResponse struct for BetaLicenseAgreementResponse
type BetaLicenseAgreementResponse struct {
	Data BetaLicenseAgreement `json:"data"`
	Included []App `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewBetaLicenseAgreementResponse instantiates a new BetaLicenseAgreementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreementResponse(data BetaLicenseAgreement, links DocumentLinks) *BetaLicenseAgreementResponse {
	this := BetaLicenseAgreementResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaLicenseAgreementResponseWithDefaults instantiates a new BetaLicenseAgreementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementResponseWithDefaults() *BetaLicenseAgreementResponse {
	this := BetaLicenseAgreementResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaLicenseAgreementResponse) GetData() BetaLicenseAgreement {
	if o == nil {
		var ret BetaLicenseAgreement
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementResponse) GetDataOk() (*BetaLicenseAgreement, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaLicenseAgreementResponse) SetData(v BetaLicenseAgreement) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaLicenseAgreementResponse) GetIncluded() []App {
	if o == nil || IsNil(o.Included) {
		var ret []App
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementResponse) GetIncludedOk() ([]App, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaLicenseAgreementResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []App and assigns it to the Included field.
func (o *BetaLicenseAgreementResponse) SetIncluded(v []App) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaLicenseAgreementResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreementResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaLicenseAgreementResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaLicenseAgreementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaLicenseAgreementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBetaLicenseAgreementResponse struct {
	value *BetaLicenseAgreementResponse
	isSet bool
}

func (v NullableBetaLicenseAgreementResponse) Get() *BetaLicenseAgreementResponse {
	return v.value
}

func (v *NullableBetaLicenseAgreementResponse) Set(val *BetaLicenseAgreementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreementResponse(val *BetaLicenseAgreementResponse) *NullableBetaLicenseAgreementResponse {
	return &NullableBetaLicenseAgreementResponse{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


