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

// checks if the AppEventScreenshotAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventScreenshotAttributes{}

// AppEventScreenshotAttributes struct for AppEventScreenshotAttributes
type AppEventScreenshotAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	ImageAsset *ImageAsset `json:"imageAsset,omitempty"`
	AssetToken *string `json:"assetToken,omitempty"`
	UploadOperations []UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	AppEventAssetType *AppEventAssetType `json:"appEventAssetType,omitempty"`
}

// NewAppEventScreenshotAttributes instantiates a new AppEventScreenshotAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventScreenshotAttributes() *AppEventScreenshotAttributes {
	this := AppEventScreenshotAttributes{}
	return &this
}

// NewAppEventScreenshotAttributesWithDefaults instantiates a new AppEventScreenshotAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventScreenshotAttributesWithDefaults() *AppEventScreenshotAttributes {
	this := AppEventScreenshotAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AppEventScreenshotAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AppEventScreenshotAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetImageAsset returns the ImageAsset field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetImageAsset() ImageAsset {
	if o == nil || IsNil(o.ImageAsset) {
		var ret ImageAsset
		return ret
	}
	return *o.ImageAsset
}

// GetImageAssetOk returns a tuple with the ImageAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetImageAssetOk() (*ImageAsset, bool) {
	if o == nil || IsNil(o.ImageAsset) {
		return nil, false
	}
	return o.ImageAsset, true
}

// HasImageAsset returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasImageAsset() bool {
	if o != nil && !IsNil(o.ImageAsset) {
		return true
	}

	return false
}

// SetImageAsset gets a reference to the given ImageAsset and assigns it to the ImageAsset field.
func (o *AppEventScreenshotAttributes) SetImageAsset(v ImageAsset) {
	o.ImageAsset = &v
}

// GetAssetToken returns the AssetToken field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetAssetToken() string {
	if o == nil || IsNil(o.AssetToken) {
		var ret string
		return ret
	}
	return *o.AssetToken
}

// GetAssetTokenOk returns a tuple with the AssetToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetAssetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AssetToken) {
		return nil, false
	}
	return o.AssetToken, true
}

// HasAssetToken returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasAssetToken() bool {
	if o != nil && !IsNil(o.AssetToken) {
		return true
	}

	return false
}

// SetAssetToken gets a reference to the given string and assigns it to the AssetToken field.
func (o *AppEventScreenshotAttributes) SetAssetToken(v string) {
	o.AssetToken = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || IsNil(o.UploadOperations) {
		var ret []UploadOperation
		return ret
	}
	return o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetUploadOperationsOk() ([]UploadOperation, bool) {
	if o == nil || IsNil(o.UploadOperations) {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasUploadOperations() bool {
	if o != nil && !IsNil(o.UploadOperations) {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *AppEventScreenshotAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || IsNil(o.AssetDeliveryState) {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || IsNil(o.AssetDeliveryState) {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasAssetDeliveryState() bool {
	if o != nil && !IsNil(o.AssetDeliveryState) {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *AppEventScreenshotAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

// GetAppEventAssetType returns the AppEventAssetType field value if set, zero value otherwise.
func (o *AppEventScreenshotAttributes) GetAppEventAssetType() AppEventAssetType {
	if o == nil || IsNil(o.AppEventAssetType) {
		var ret AppEventAssetType
		return ret
	}
	return *o.AppEventAssetType
}

// GetAppEventAssetTypeOk returns a tuple with the AppEventAssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventScreenshotAttributes) GetAppEventAssetTypeOk() (*AppEventAssetType, bool) {
	if o == nil || IsNil(o.AppEventAssetType) {
		return nil, false
	}
	return o.AppEventAssetType, true
}

// HasAppEventAssetType returns a boolean if a field has been set.
func (o *AppEventScreenshotAttributes) HasAppEventAssetType() bool {
	if o != nil && !IsNil(o.AppEventAssetType) {
		return true
	}

	return false
}

// SetAppEventAssetType gets a reference to the given AppEventAssetType and assigns it to the AppEventAssetType field.
func (o *AppEventScreenshotAttributes) SetAppEventAssetType(v AppEventAssetType) {
	o.AppEventAssetType = &v
}

func (o AppEventScreenshotAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventScreenshotAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.ImageAsset) {
		toSerialize["imageAsset"] = o.ImageAsset
	}
	if !IsNil(o.AssetToken) {
		toSerialize["assetToken"] = o.AssetToken
	}
	if !IsNil(o.UploadOperations) {
		toSerialize["uploadOperations"] = o.UploadOperations
	}
	if !IsNil(o.AssetDeliveryState) {
		toSerialize["assetDeliveryState"] = o.AssetDeliveryState
	}
	if !IsNil(o.AppEventAssetType) {
		toSerialize["appEventAssetType"] = o.AppEventAssetType
	}
	return toSerialize, nil
}

type NullableAppEventScreenshotAttributes struct {
	value *AppEventScreenshotAttributes
	isSet bool
}

func (v NullableAppEventScreenshotAttributes) Get() *AppEventScreenshotAttributes {
	return v.value
}

func (v *NullableAppEventScreenshotAttributes) Set(val *AppEventScreenshotAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventScreenshotAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventScreenshotAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventScreenshotAttributes(val *AppEventScreenshotAttributes) *NullableAppEventScreenshotAttributes {
	return &NullableAppEventScreenshotAttributes{value: val, isSet: true}
}

func (v NullableAppEventScreenshotAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventScreenshotAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


