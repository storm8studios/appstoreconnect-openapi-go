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

// checks if the SubscriptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsResponse{}

// SubscriptionsResponse struct for SubscriptionsResponse
type SubscriptionsResponse struct {
	Data []Subscription `json:"data"`
	Included []SubscriptionsResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSubscriptionsResponse instantiates a new SubscriptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsResponse(data []Subscription, links PagedDocumentLinks) *SubscriptionsResponse {
	this := SubscriptionsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionsResponseWithDefaults instantiates a new SubscriptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsResponseWithDefaults() *SubscriptionsResponse {
	this := SubscriptionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionsResponse) GetData() []Subscription {
	if o == nil {
		var ret []Subscription
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsResponse) GetDataOk() ([]Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionsResponse) SetData(v []Subscription) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionsResponse) GetIncluded() []SubscriptionsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsResponse) GetIncludedOk() ([]SubscriptionsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionsResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionsResponse) SetIncluded(v []SubscriptionsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionsResponse struct {
	value *SubscriptionsResponse
	isSet bool
}

func (v NullableSubscriptionsResponse) Get() *SubscriptionsResponse {
	return v.value
}

func (v *NullableSubscriptionsResponse) Set(val *SubscriptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsResponse(val *SubscriptionsResponse) *NullableSubscriptionsResponse {
	return &NullableSubscriptionsResponse{value: val, isSet: true}
}

func (v NullableSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


