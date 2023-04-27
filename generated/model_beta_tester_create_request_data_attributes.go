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

// checks if the BetaTesterCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterCreateRequestDataAttributes{}

// BetaTesterCreateRequestDataAttributes struct for BetaTesterCreateRequestDataAttributes
type BetaTesterCreateRequestDataAttributes struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Email string `json:"email"`
}

// NewBetaTesterCreateRequestDataAttributes instantiates a new BetaTesterCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterCreateRequestDataAttributes(email string) *BetaTesterCreateRequestDataAttributes {
	this := BetaTesterCreateRequestDataAttributes{}
	this.Email = email
	return &this
}

// NewBetaTesterCreateRequestDataAttributesWithDefaults instantiates a new BetaTesterCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterCreateRequestDataAttributesWithDefaults() *BetaTesterCreateRequestDataAttributes {
	this := BetaTesterCreateRequestDataAttributes{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BetaTesterCreateRequestDataAttributes) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestDataAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BetaTesterCreateRequestDataAttributes) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BetaTesterCreateRequestDataAttributes) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BetaTesterCreateRequestDataAttributes) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestDataAttributes) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BetaTesterCreateRequestDataAttributes) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BetaTesterCreateRequestDataAttributes) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value
func (o *BetaTesterCreateRequestDataAttributes) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *BetaTesterCreateRequestDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *BetaTesterCreateRequestDataAttributes) SetEmail(v string) {
	o.Email = v
}

func (o BetaTesterCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

type NullableBetaTesterCreateRequestDataAttributes struct {
	value *BetaTesterCreateRequestDataAttributes
	isSet bool
}

func (v NullableBetaTesterCreateRequestDataAttributes) Get() *BetaTesterCreateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaTesterCreateRequestDataAttributes) Set(val *BetaTesterCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterCreateRequestDataAttributes(val *BetaTesterCreateRequestDataAttributes) *NullableBetaTesterCreateRequestDataAttributes {
	return &NullableBetaTesterCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaTesterCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


