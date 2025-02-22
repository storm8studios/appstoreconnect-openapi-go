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

// checks if the SubscriptionAppStoreReviewScreenshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAppStoreReviewScreenshot{}

// SubscriptionAppStoreReviewScreenshot struct for SubscriptionAppStoreReviewScreenshot
type SubscriptionAppStoreReviewScreenshot struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppScreenshotAttributes `json:"attributes,omitempty"`
	Relationships *SubscriptionAppStoreReviewScreenshotRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewSubscriptionAppStoreReviewScreenshot instantiates a new SubscriptionAppStoreReviewScreenshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAppStoreReviewScreenshot(type_ string, id string, links ResourceLinks) *SubscriptionAppStoreReviewScreenshot {
	this := SubscriptionAppStoreReviewScreenshot{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewSubscriptionAppStoreReviewScreenshotWithDefaults instantiates a new SubscriptionAppStoreReviewScreenshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAppStoreReviewScreenshotWithDefaults() *SubscriptionAppStoreReviewScreenshot {
	this := SubscriptionAppStoreReviewScreenshot{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionAppStoreReviewScreenshot) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshot) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionAppStoreReviewScreenshot) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *SubscriptionAppStoreReviewScreenshot) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionAppStoreReviewScreenshot) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *SubscriptionAppStoreReviewScreenshot) GetAttributes() AppScreenshotAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppScreenshotAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshot) GetAttributesOk() (*AppScreenshotAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *SubscriptionAppStoreReviewScreenshot) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppScreenshotAttributes and assigns it to the Attributes field.
func (o *SubscriptionAppStoreReviewScreenshot) SetAttributes(v AppScreenshotAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *SubscriptionAppStoreReviewScreenshot) GetRelationships() SubscriptionAppStoreReviewScreenshotRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret SubscriptionAppStoreReviewScreenshotRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshot) GetRelationshipsOk() (*SubscriptionAppStoreReviewScreenshotRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *SubscriptionAppStoreReviewScreenshot) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given SubscriptionAppStoreReviewScreenshotRelationships and assigns it to the Relationships field.
func (o *SubscriptionAppStoreReviewScreenshot) SetRelationships(v SubscriptionAppStoreReviewScreenshotRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *SubscriptionAppStoreReviewScreenshot) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshot) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionAppStoreReviewScreenshot) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o SubscriptionAppStoreReviewScreenshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAppStoreReviewScreenshot) ToMap() (map[string]interface{}, error) {
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

type NullableSubscriptionAppStoreReviewScreenshot struct {
	value *SubscriptionAppStoreReviewScreenshot
	isSet bool
}

func (v NullableSubscriptionAppStoreReviewScreenshot) Get() *SubscriptionAppStoreReviewScreenshot {
	return v.value
}

func (v *NullableSubscriptionAppStoreReviewScreenshot) Set(val *SubscriptionAppStoreReviewScreenshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAppStoreReviewScreenshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAppStoreReviewScreenshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAppStoreReviewScreenshot(val *SubscriptionAppStoreReviewScreenshot) *NullableSubscriptionAppStoreReviewScreenshot {
	return &NullableSubscriptionAppStoreReviewScreenshot{value: val, isSet: true}
}

func (v NullableSubscriptionAppStoreReviewScreenshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAppStoreReviewScreenshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


