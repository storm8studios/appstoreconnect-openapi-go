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

// checks if the AppEventScreenshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshot{}

// AppEventScreenshot struct for AppEventScreenshot
type AppEventScreenshot struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppEventScreenshotAttributes `json:"attributes,omitempty"`
	Relationships *AppEventScreenshotRelationships `json:"relationships,omitempty"`
	Links ResourceLinks `json:"links"`
}

// NewAppEventScreenshot instantiates a new AppEventScreenshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshot(type_ string, id string, links ResourceLinks) *AppEventScreenshot {
	this := AppEventScreenshot{}
	this.Type = type_
	this.Id = id
	this.Links = links
	return &this
}

// NewAppEventScreenshotWithDefaults instantiates a new AppEventScreenshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotWithDefaults() *AppEventScreenshot {
	this := AppEventScreenshot{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventScreenshot) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshot) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventScreenshot) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppEventScreenshot) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppEventScreenshot) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppEventScreenshot) GetAttributes() AppEventScreenshotAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppEventScreenshotAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshot) GetAttributesOk() (*AppEventScreenshotAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppEventScreenshot) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppEventScreenshotAttributes and assigns it to the Attributes field.
func (o *AppEventScreenshot) SetAttributes(v AppEventScreenshotAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppEventScreenshot) GetRelationships() AppEventScreenshotRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppEventScreenshotRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshot) GetRelationshipsOk() (*AppEventScreenshotRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppEventScreenshot) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppEventScreenshotRelationships and assigns it to the Relationships field.
func (o *AppEventScreenshot) SetRelationships(v AppEventScreenshotRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value
func (o *AppEventScreenshot) GetLinks() ResourceLinks {
	if o == nil {
		var ret ResourceLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppEventScreenshot) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppEventScreenshot) SetLinks(v ResourceLinks) {
	o.Links = v
}

func (o AppEventScreenshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshot) ToMap() (map[string]interface{}, error) {
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

type NullableAppEventScreenshot struct {
	value *AppEventScreenshot
	isSet bool
}

func (v NullableAppEventScreenshot) Get() *AppEventScreenshot {
	return v.value
}

func (v *NullableAppEventScreenshot) Set(val *AppEventScreenshot) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshot) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshot(val *AppEventScreenshot) *NullableAppEventScreenshot {
	return &NullableAppEventScreenshot{value: val, isSet: true}
}

func (v NullableAppEventScreenshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


