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

// checks if the CustomerReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReview{}

// CustomerReview struct for CustomerReview
type CustomerReview struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *CustomerReviewAttributes `json:"attributes,omitempty"`
	Relationships *CustomerReviewRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewCustomerReview instantiates a new CustomerReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReview(type_ string, id string, links ResourceLinks) *CustomerReview {
	this := CustomerReview{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewCustomerReviewWithDefaults instantiates a new CustomerReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewWithDefaults() *CustomerReview {
	this := CustomerReview{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerReview) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerReview) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CustomerReview) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerReview) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerReview) GetAttributes() CustomerReviewAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CustomerReviewAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetAttributesOk() (*CustomerReviewAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerReview) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CustomerReviewAttributes and assigns it to the Attributes field.
func (o *CustomerReview) SetAttributes(v CustomerReviewAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CustomerReview) GetRelationships() CustomerReviewRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CustomerReviewRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetRelationshipsOk() (*CustomerReviewRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CustomerReview) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CustomerReviewRelationships and assigns it to the Relationships field.
func (o *CustomerReview) SetRelationships(v CustomerReviewRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *CustomerReview) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CustomerReview) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CustomerReview) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o CustomerReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableCustomerReview struct {
	value *CustomerReview
	isSet bool
}

func (v NullableCustomerReview) Get() *CustomerReview {
	return v.value
}

func (v *NullableCustomerReview) Set(val *CustomerReview) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReview) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReview(val *CustomerReview) *NullableCustomerReview {
	return &NullableCustomerReview{value: val, isSet: true}
}

func (v NullableCustomerReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


