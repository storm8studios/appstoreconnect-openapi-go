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

// checks if the PromotedPurchasesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchasesResponse{}

// PromotedPurchasesResponse struct for PromotedPurchasesResponse
type PromotedPurchasesResponse struct {
	Data []PromotedPurchase `json:"data"`
	Included []PromotedPurchasesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewPromotedPurchasesResponse instantiates a new PromotedPurchasesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchasesResponse(data []PromotedPurchase, links PagedDocumentLinks) *PromotedPurchasesResponse {
	this := PromotedPurchasesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPromotedPurchasesResponseWithDefaults instantiates a new PromotedPurchasesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchasesResponseWithDefaults() *PromotedPurchasesResponse {
	this := PromotedPurchasesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PromotedPurchasesResponse) GetData() []PromotedPurchase {
	if o == nil {
		var ret []PromotedPurchase
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchasesResponse) GetDataOk() ([]PromotedPurchase, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PromotedPurchasesResponse) SetData(v []PromotedPurchase) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *PromotedPurchasesResponse) GetIncluded() []PromotedPurchasesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []PromotedPurchasesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchasesResponse) GetIncludedOk() ([]PromotedPurchasesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *PromotedPurchasesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []PromotedPurchasesResponseIncludedInner and assigns it to the Included field.
func (o *PromotedPurchasesResponse) SetIncluded(v []PromotedPurchasesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *PromotedPurchasesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchasesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PromotedPurchasesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PromotedPurchasesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchasesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PromotedPurchasesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *PromotedPurchasesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o PromotedPurchasesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchasesResponse) ToMap() (map[string]interface{}, error) {
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

type NullablePromotedPurchasesResponse struct {
	value *PromotedPurchasesResponse
	isSet bool
}

func (v NullablePromotedPurchasesResponse) Get() *PromotedPurchasesResponse {
	return v.value
}

func (v *NullablePromotedPurchasesResponse) Set(val *PromotedPurchasesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchasesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchasesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchasesResponse(val *PromotedPurchasesResponse) *NullablePromotedPurchasesResponse {
	return &NullablePromotedPurchasesResponse{value: val, isSet: true}
}

func (v NullablePromotedPurchasesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchasesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


