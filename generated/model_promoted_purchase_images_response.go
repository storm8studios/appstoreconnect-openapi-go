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

// checks if the PromotedPurchaseImagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImagesResponse{}

// PromotedPurchaseImagesResponse struct for PromotedPurchaseImagesResponse
type PromotedPurchaseImagesResponse struct {
	Data []PromotedPurchaseImage `json:"data"`
	Included []PromotedPurchase `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewPromotedPurchaseImagesResponse instantiates a new PromotedPurchaseImagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImagesResponse(data []PromotedPurchaseImage, links PagedDocumentLinks) *PromotedPurchaseImagesResponse {
	this := PromotedPurchaseImagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewPromotedPurchaseImagesResponseWithDefaults instantiates a new PromotedPurchaseImagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImagesResponseWithDefaults() *PromotedPurchaseImagesResponse {
	this := PromotedPurchaseImagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *PromotedPurchaseImagesResponse) GetData() []PromotedPurchaseImage {
	if o == nil {
		var ret []PromotedPurchaseImage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImagesResponse) GetDataOk() ([]PromotedPurchaseImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *PromotedPurchaseImagesResponse) SetData(v []PromotedPurchaseImage) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *PromotedPurchaseImagesResponse) GetIncluded() []PromotedPurchase {
	if o == nil || IsNil(o.Included) {
		var ret []PromotedPurchase
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImagesResponse) GetIncludedOk() ([]PromotedPurchase, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *PromotedPurchaseImagesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []PromotedPurchase and assigns it to the Included field.
func (o *PromotedPurchaseImagesResponse) SetIncluded(v []PromotedPurchase) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *PromotedPurchaseImagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *PromotedPurchaseImagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PromotedPurchaseImagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PromotedPurchaseImagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *PromotedPurchaseImagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o PromotedPurchaseImagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImagesResponse) ToMap() (map[string]interface{}, error) {
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

type NullablePromotedPurchaseImagesResponse struct {
	value *PromotedPurchaseImagesResponse
	isSet bool
}

func (v NullablePromotedPurchaseImagesResponse) Get() *PromotedPurchaseImagesResponse {
	return v.value
}

func (v *NullablePromotedPurchaseImagesResponse) Set(val *PromotedPurchaseImagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImagesResponse(val *PromotedPurchaseImagesResponse) *NullablePromotedPurchaseImagesResponse {
	return &NullablePromotedPurchaseImagesResponse{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


