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

// checks if the CiWorkflowCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowCreateRequestDataRelationships{}

// CiWorkflowCreateRequestDataRelationships struct for CiWorkflowCreateRequestDataRelationships
type CiWorkflowCreateRequestDataRelationships struct {
	Product CiWorkflowCreateRequestDataRelationshipsProduct `json:"product"`
	Repository CiWorkflowCreateRequestDataRelationshipsRepository `json:"repository"`
	XcodeVersion CiWorkflowCreateRequestDataRelationshipsXcodeVersion `json:"xcodeVersion"`
	MacOsVersion CiWorkflowCreateRequestDataRelationshipsMacOsVersion `json:"macOsVersion"`
}

// NewCiWorkflowCreateRequestDataRelationships instantiates a new CiWorkflowCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowCreateRequestDataRelationships(product CiWorkflowCreateRequestDataRelationshipsProduct, repository CiWorkflowCreateRequestDataRelationshipsRepository, xcodeVersion CiWorkflowCreateRequestDataRelationshipsXcodeVersion, macOsVersion CiWorkflowCreateRequestDataRelationshipsMacOsVersion) *CiWorkflowCreateRequestDataRelationships {
	this := CiWorkflowCreateRequestDataRelationships{}
	this.Product = product
	this.Repository = repository
	this.XcodeVersion = xcodeVersion
	this.MacOsVersion = macOsVersion
	return &this
}

// NewCiWorkflowCreateRequestDataRelationshipsWithDefaults instantiates a new CiWorkflowCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowCreateRequestDataRelationshipsWithDefaults() *CiWorkflowCreateRequestDataRelationships {
	this := CiWorkflowCreateRequestDataRelationships{}
	return &this
}

// GetProduct returns the Product field value
func (o *CiWorkflowCreateRequestDataRelationships) GetProduct() CiWorkflowCreateRequestDataRelationshipsProduct {
	if o == nil {
		var ret CiWorkflowCreateRequestDataRelationshipsProduct
		return ret
	}

	return o.Product
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataRelationships) GetProductOk() (*CiWorkflowCreateRequestDataRelationshipsProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Product, true
}

// SetProduct sets field value
func (o *CiWorkflowCreateRequestDataRelationships) SetProduct(v CiWorkflowCreateRequestDataRelationshipsProduct) {
	o.Product = v
}

// GetRepository returns the Repository field value
func (o *CiWorkflowCreateRequestDataRelationships) GetRepository() CiWorkflowCreateRequestDataRelationshipsRepository {
	if o == nil {
		var ret CiWorkflowCreateRequestDataRelationshipsRepository
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataRelationships) GetRepositoryOk() (*CiWorkflowCreateRequestDataRelationshipsRepository, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *CiWorkflowCreateRequestDataRelationships) SetRepository(v CiWorkflowCreateRequestDataRelationshipsRepository) {
	o.Repository = v
}

// GetXcodeVersion returns the XcodeVersion field value
func (o *CiWorkflowCreateRequestDataRelationships) GetXcodeVersion() CiWorkflowCreateRequestDataRelationshipsXcodeVersion {
	if o == nil {
		var ret CiWorkflowCreateRequestDataRelationshipsXcodeVersion
		return ret
	}

	return o.XcodeVersion
}

// GetXcodeVersionOk returns a tuple with the XcodeVersion field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataRelationships) GetXcodeVersionOk() (*CiWorkflowCreateRequestDataRelationshipsXcodeVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XcodeVersion, true
}

// SetXcodeVersion sets field value
func (o *CiWorkflowCreateRequestDataRelationships) SetXcodeVersion(v CiWorkflowCreateRequestDataRelationshipsXcodeVersion) {
	o.XcodeVersion = v
}

// GetMacOsVersion returns the MacOsVersion field value
func (o *CiWorkflowCreateRequestDataRelationships) GetMacOsVersion() CiWorkflowCreateRequestDataRelationshipsMacOsVersion {
	if o == nil {
		var ret CiWorkflowCreateRequestDataRelationshipsMacOsVersion
		return ret
	}

	return o.MacOsVersion
}

// GetMacOsVersionOk returns a tuple with the MacOsVersion field value
// and a boolean to check if the value has been set.
func (o *CiWorkflowCreateRequestDataRelationships) GetMacOsVersionOk() (*CiWorkflowCreateRequestDataRelationshipsMacOsVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacOsVersion, true
}

// SetMacOsVersion sets field value
func (o *CiWorkflowCreateRequestDataRelationships) SetMacOsVersion(v CiWorkflowCreateRequestDataRelationshipsMacOsVersion) {
	o.MacOsVersion = v
}

func (o CiWorkflowCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product"] = o.Product
	toSerialize["repository"] = o.Repository
	toSerialize["xcodeVersion"] = o.XcodeVersion
	toSerialize["macOsVersion"] = o.MacOsVersion
	return toSerialize, nil
}

type NullableCiWorkflowCreateRequestDataRelationships struct {
	value *CiWorkflowCreateRequestDataRelationships
	isSet bool
}

func (v NullableCiWorkflowCreateRequestDataRelationships) Get() *CiWorkflowCreateRequestDataRelationships {
	return v.value
}

func (v *NullableCiWorkflowCreateRequestDataRelationships) Set(val *CiWorkflowCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowCreateRequestDataRelationships(val *CiWorkflowCreateRequestDataRelationships) *NullableCiWorkflowCreateRequestDataRelationships {
	return &NullableCiWorkflowCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableCiWorkflowCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


