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

// checks if the SubscriptionIntroductoryOffersLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOffersLinkagesResponse{}

// SubscriptionIntroductoryOffersLinkagesResponse struct for SubscriptionIntroductoryOffersLinkagesResponse
type SubscriptionIntroductoryOffersLinkagesResponse struct {
	Data []SubscriptionRelationshipsIntroductoryOffersDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSubscriptionIntroductoryOffersLinkagesResponse instantiates a new SubscriptionIntroductoryOffersLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOffersLinkagesResponse(data []SubscriptionRelationshipsIntroductoryOffersDataInner, links PagedDocumentLinks) *SubscriptionIntroductoryOffersLinkagesResponse {
	this := SubscriptionIntroductoryOffersLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionIntroductoryOffersLinkagesResponseWithDefaults instantiates a new SubscriptionIntroductoryOffersLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOffersLinkagesResponseWithDefaults() *SubscriptionIntroductoryOffersLinkagesResponse {
	this := SubscriptionIntroductoryOffersLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionIntroductoryOffersLinkagesResponse) GetData() []SubscriptionRelationshipsIntroductoryOffersDataInner {
	if o == nil {
		var ret []SubscriptionRelationshipsIntroductoryOffersDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersLinkagesResponse) GetDataOk() ([]SubscriptionRelationshipsIntroductoryOffersDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionIntroductoryOffersLinkagesResponse) SetData(v []SubscriptionRelationshipsIntroductoryOffersDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *SubscriptionIntroductoryOffersLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionIntroductoryOffersLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOffersLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOffersLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionIntroductoryOffersLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionIntroductoryOffersLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOffersLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOffersLinkagesResponse struct {
	value *SubscriptionIntroductoryOffersLinkagesResponse
	isSet bool
}

func (v NullableSubscriptionIntroductoryOffersLinkagesResponse) Get() *SubscriptionIntroductoryOffersLinkagesResponse {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOffersLinkagesResponse) Set(val *SubscriptionIntroductoryOffersLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOffersLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOffersLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOffersLinkagesResponse(val *SubscriptionIntroductoryOffersLinkagesResponse) *NullableSubscriptionIntroductoryOffersLinkagesResponse {
	return &NullableSubscriptionIntroductoryOffersLinkagesResponse{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOffersLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOffersLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


