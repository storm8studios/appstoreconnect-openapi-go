/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package appstoreopenapi

import (
	"encoding/json"
	"time"
)

// checks if the ScmRepositoryAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmRepositoryAttributes{}

// ScmRepositoryAttributes struct for ScmRepositoryAttributes
type ScmRepositoryAttributes struct {
	LastAccessedDate *time.Time `json:"lastAccessedDate,omitempty"`
	HttpCloneUrl *string `json:"httpCloneUrl,omitempty"`
	SshCloneUrl *string `json:"sshCloneUrl,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	RepositoryName *string `json:"repositoryName,omitempty"`
}

// NewScmRepositoryAttributes instantiates a new ScmRepositoryAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmRepositoryAttributes() *ScmRepositoryAttributes {
	this := ScmRepositoryAttributes{}
	return &this
}

// NewScmRepositoryAttributesWithDefaults instantiates a new ScmRepositoryAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmRepositoryAttributesWithDefaults() *ScmRepositoryAttributes {
	this := ScmRepositoryAttributes{}
	return &this
}

// GetLastAccessedDate returns the LastAccessedDate field value if set, zero value otherwise.
func (o *ScmRepositoryAttributes) GetLastAccessedDate() time.Time {
	if o == nil || IsNil(o.LastAccessedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastAccessedDate
}

// GetLastAccessedDateOk returns a tuple with the LastAccessedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryAttributes) GetLastAccessedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAccessedDate) {
		return nil, false
	}
	return o.LastAccessedDate, true
}

// HasLastAccessedDate returns a boolean if a field has been set.
func (o *ScmRepositoryAttributes) HasLastAccessedDate() bool {
	if o != nil && !IsNil(o.LastAccessedDate) {
		return true
	}

	return false
}

// SetLastAccessedDate gets a reference to the given time.Time and assigns it to the LastAccessedDate field.
func (o *ScmRepositoryAttributes) SetLastAccessedDate(v time.Time) {
	o.LastAccessedDate = &v
}

// GetHttpCloneUrl returns the HttpCloneUrl field value if set, zero value otherwise.
func (o *ScmRepositoryAttributes) GetHttpCloneUrl() string {
	if o == nil || IsNil(o.HttpCloneUrl) {
		var ret string
		return ret
	}
	return *o.HttpCloneUrl
}

// GetHttpCloneUrlOk returns a tuple with the HttpCloneUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryAttributes) GetHttpCloneUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HttpCloneUrl) {
		return nil, false
	}
	return o.HttpCloneUrl, true
}

// HasHttpCloneUrl returns a boolean if a field has been set.
func (o *ScmRepositoryAttributes) HasHttpCloneUrl() bool {
	if o != nil && !IsNil(o.HttpCloneUrl) {
		return true
	}

	return false
}

// SetHttpCloneUrl gets a reference to the given string and assigns it to the HttpCloneUrl field.
func (o *ScmRepositoryAttributes) SetHttpCloneUrl(v string) {
	o.HttpCloneUrl = &v
}

// GetSshCloneUrl returns the SshCloneUrl field value if set, zero value otherwise.
func (o *ScmRepositoryAttributes) GetSshCloneUrl() string {
	if o == nil || IsNil(o.SshCloneUrl) {
		var ret string
		return ret
	}
	return *o.SshCloneUrl
}

// GetSshCloneUrlOk returns a tuple with the SshCloneUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryAttributes) GetSshCloneUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SshCloneUrl) {
		return nil, false
	}
	return o.SshCloneUrl, true
}

// HasSshCloneUrl returns a boolean if a field has been set.
func (o *ScmRepositoryAttributes) HasSshCloneUrl() bool {
	if o != nil && !IsNil(o.SshCloneUrl) {
		return true
	}

	return false
}

// SetSshCloneUrl gets a reference to the given string and assigns it to the SshCloneUrl field.
func (o *ScmRepositoryAttributes) SetSshCloneUrl(v string) {
	o.SshCloneUrl = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *ScmRepositoryAttributes) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryAttributes) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *ScmRepositoryAttributes) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *ScmRepositoryAttributes) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetRepositoryName returns the RepositoryName field value if set, zero value otherwise.
func (o *ScmRepositoryAttributes) GetRepositoryName() string {
	if o == nil || IsNil(o.RepositoryName) {
		var ret string
		return ret
	}
	return *o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryAttributes) GetRepositoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryName) {
		return nil, false
	}
	return o.RepositoryName, true
}

// HasRepositoryName returns a boolean if a field has been set.
func (o *ScmRepositoryAttributes) HasRepositoryName() bool {
	if o != nil && !IsNil(o.RepositoryName) {
		return true
	}

	return false
}

// SetRepositoryName gets a reference to the given string and assigns it to the RepositoryName field.
func (o *ScmRepositoryAttributes) SetRepositoryName(v string) {
	o.RepositoryName = &v
}

func (o ScmRepositoryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmRepositoryAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastAccessedDate) {
		toSerialize["lastAccessedDate"] = o.LastAccessedDate
	}
	if !IsNil(o.HttpCloneUrl) {
		toSerialize["httpCloneUrl"] = o.HttpCloneUrl
	}
	if !IsNil(o.SshCloneUrl) {
		toSerialize["sshCloneUrl"] = o.SshCloneUrl
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.RepositoryName) {
		toSerialize["repositoryName"] = o.RepositoryName
	}
	return toSerialize, nil
}

type NullableScmRepositoryAttributes struct {
	value *ScmRepositoryAttributes
	isSet bool
}

func (v NullableScmRepositoryAttributes) Get() *ScmRepositoryAttributes {
	return v.value
}

func (v *NullableScmRepositoryAttributes) Set(val *ScmRepositoryAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepositoryAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepositoryAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepositoryAttributes(val *ScmRepositoryAttributes) *NullableScmRepositoryAttributes {
	return &NullableScmRepositoryAttributes{value: val, isSet: true}
}

func (v NullableScmRepositoryAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepositoryAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


