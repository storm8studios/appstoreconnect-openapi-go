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

// checks if the SubscriptionIntroductoryOffersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOffersResponse{}

// SubscriptionIntroductoryOffersResponse struct for SubscriptionIntroductoryOffersResponse
type SubscriptionIntroductoryOffersResponse struct {
	Data []SubscriptionIntroductoryOffer `json:"data"`
	Included []SubscriptionIntroductoryOffersResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSubscriptionIntroductoryOffersResponse instantiates a new SubscriptionIntroductoryOffersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOffersResponse(data []SubscriptionIntroductoryOffer, links PagedDocumentLinks) *SubscriptionIntroductoryOffersResponse {
	this := SubscriptionIntroductoryOffersResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionIntroductoryOffersResponseWithDefaults instantiates a new SubscriptionIntroductoryOffersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOffersResponseWithDefaults() *SubscriptionIntroductoryOffersResponse {
	this := SubscriptionIntroductoryOffersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionIntroductoryOffersResponse) GetData() []SubscriptionIntroductoryOffer {
	if o == nil {
		var ret []SubscriptionIntroductoryOffer
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersResponse) GetDataOk() ([]SubscriptionIntroductoryOffer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionIntroductoryOffersResponse) SetData(v []SubscriptionIntroductoryOffer) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOffersResponse) GetIncluded() []SubscriptionIntroductoryOffersResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionIntroductoryOffersResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersResponse) GetIncludedOk() ([]SubscriptionIntroductoryOffersResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOffersResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionIntroductoryOffersResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionIntroductoryOffersResponse) SetIncluded(v []SubscriptionIntroductoryOffersResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionIntroductoryOffersResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionIntroductoryOffersResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOffersResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOffersResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOffersResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionIntroductoryOffersResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionIntroductoryOffersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOffersResponse) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionIntroductoryOffersResponse struct {
	value *SubscriptionIntroductoryOffersResponse
	isSet bool
}

func (v NullableSubscriptionIntroductoryOffersResponse) Get() *SubscriptionIntroductoryOffersResponse {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOffersResponse) Set(val *SubscriptionIntroductoryOffersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOffersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOffersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOffersResponse(val *SubscriptionIntroductoryOffersResponse) *NullableSubscriptionIntroductoryOffersResponse {
	return &NullableSubscriptionIntroductoryOffersResponse{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOffersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOffersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


