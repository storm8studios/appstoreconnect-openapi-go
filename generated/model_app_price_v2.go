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

// checks if the AppPriceV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceV2{}

// AppPriceV2 struct for AppPriceV2
type AppPriceV2 struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppPriceV2Attributes `json:"attributes,omitempty"`
	Relationships *AppPriceV2Relationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppPriceV2 instantiates a new AppPriceV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceV2(type_ string, id string, links ResourceLinks) *AppPriceV2 {
	this := AppPriceV2{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppPriceV2WithDefaults instantiates a new AppPriceV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceV2WithDefaults() *AppPriceV2 {
	this := AppPriceV2{}
	return &this
}

// GetType returns the Type field value
func (o *AppPriceV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPriceV2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPriceV2) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPriceV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPriceV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPriceV2) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppPriceV2) GetAttributes() AppPriceV2Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppPriceV2Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceV2) GetAttributesOk() (*AppPriceV2Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppPriceV2) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppPriceV2Attributes and assigns it to the Attributes field.
func (o *AppPriceV2) SetAttributes(v AppPriceV2Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppPriceV2) GetRelationships() AppPriceV2Relationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppPriceV2Relationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceV2) GetRelationshipsOk() (*AppPriceV2Relationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppPriceV2) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPriceV2Relationships and assigns it to the Relationships field.
func (o *AppPriceV2) SetRelationships(v AppPriceV2Relationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppPriceV2) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPriceV2) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPriceV2) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppPriceV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceV2) ToMap() (map[string]interface{}, error) {
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

type NullableAppPriceV2 struct {
	value *AppPriceV2
	isSet bool
}

func (v NullableAppPriceV2) Get() *AppPriceV2 {
	return v.value
}

func (v *NullableAppPriceV2) Set(val *AppPriceV2) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceV2) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceV2(val *AppPriceV2) *NullableAppPriceV2 {
	return &NullableAppPriceV2{value: val, isSet: true}
}

func (v NullableAppPriceV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


