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

// checks if the SubscriptionGroupLocalizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationsResponse{}

// SubscriptionGroupLocalizationsResponse struct for SubscriptionGroupLocalizationsResponse
type SubscriptionGroupLocalizationsResponse struct {
	Data []SubscriptionGroupLocalization `json:"data"`
	Included []SubscriptionGroup `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSubscriptionGroupLocalizationsResponse instantiates a new SubscriptionGroupLocalizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationsResponse(data []SubscriptionGroupLocalization, links PagedDocumentLinks) *SubscriptionGroupLocalizationsResponse {
	this := SubscriptionGroupLocalizationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionGroupLocalizationsResponseWithDefaults instantiates a new SubscriptionGroupLocalizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationsResponseWithDefaults() *SubscriptionGroupLocalizationsResponse {
	this := SubscriptionGroupLocalizationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupLocalizationsResponse) GetData() []SubscriptionGroupLocalization {
	if o == nil {
		var ret []SubscriptionGroupLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationsResponse) GetDataOk() ([]SubscriptionGroupLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupLocalizationsResponse) SetData(v []SubscriptionGroupLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalizationsResponse) GetIncluded() []SubscriptionGroup {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionGroup
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationsResponse) GetIncludedOk() ([]SubscriptionGroup, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalizationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionGroup and assigns it to the Included field.
func (o *SubscriptionGroupLocalizationsResponse) SetIncluded(v []SubscriptionGroup) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionGroupLocalizationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionGroupLocalizationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionGroupLocalizationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionGroupLocalizationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionGroupLocalizationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionGroupLocalizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableSubscriptionGroupLocalizationsResponse struct {
	value *SubscriptionGroupLocalizationsResponse
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationsResponse) Get() *SubscriptionGroupLocalizationsResponse {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationsResponse) Set(val *SubscriptionGroupLocalizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationsResponse(val *SubscriptionGroupLocalizationsResponse) *NullableSubscriptionGroupLocalizationsResponse {
	return &NullableSubscriptionGroupLocalizationsResponse{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


