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

// checks if the SubscriptionPricesLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPricesLinkagesResponse{}

// SubscriptionPricesLinkagesResponse struct for SubscriptionPricesLinkagesResponse
type SubscriptionPricesLinkagesResponse struct {
	Data []SubscriptionRelationshipsPricesDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSubscriptionPricesLinkagesResponse instantiates a new SubscriptionPricesLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPricesLinkagesResponse(data []SubscriptionRelationshipsPricesDataInner, links PagedDocumentLinks) *SubscriptionPricesLinkagesResponse {
	this := SubscriptionPricesLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionPricesLinkagesResponseWithDefaults instantiates a new SubscriptionPricesLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPricesLinkagesResponseWithDefaults() *SubscriptionPricesLinkagesResponse {
	this := SubscriptionPricesLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPricesLinkagesResponse) GetData() []SubscriptionRelationshipsPricesDataInner {
	if o == nil {
		var ret []SubscriptionRelationshipsPricesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPricesLinkagesResponse) GetDataOk() ([]SubscriptionRelationshipsPricesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionPricesLinkagesResponse) SetData(v []SubscriptionRelationshipsPricesDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *SubscriptionPricesLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPricesLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionPricesLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionPricesLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPricesLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionPricesLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionPricesLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionPricesLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPricesLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableSubscriptionPricesLinkagesResponse struct {
	value *SubscriptionPricesLinkagesResponse
	isSet bool
}

func (v NullableSubscriptionPricesLinkagesResponse) Get() *SubscriptionPricesLinkagesResponse {
	return v.value
}

func (v *NullableSubscriptionPricesLinkagesResponse) Set(val *SubscriptionPricesLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPricesLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPricesLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPricesLinkagesResponse(val *SubscriptionPricesLinkagesResponse) *NullableSubscriptionPricesLinkagesResponse {
	return &NullableSubscriptionPricesLinkagesResponse{value: val, isSet: true}
}

func (v NullableSubscriptionPricesLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPricesLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


