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

// checks if the AppEventVideoClipAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventVideoClipAttributes{}

// AppEventVideoClipAttributes struct for AppEventVideoClipAttributes
type AppEventVideoClipAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	VideoUrl *string `json:"videoUrl,omitempty"`
	PreviewImage *ImageAsset `json:"previewImage,omitempty"`
	UploadOperations []UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
	AppEventAssetType *AppEventAssetType `json:"appEventAssetType,omitempty"`
}

// NewAppEventVideoClipAttributes instantiates a new AppEventVideoClipAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventVideoClipAttributes() *AppEventVideoClipAttributes {
	this := AppEventVideoClipAttributes{}
	return &this
}

// NewAppEventVideoClipAttributesWithDefaults instantiates a new AppEventVideoClipAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventVideoClipAttributesWithDefaults() *AppEventVideoClipAttributes {
	this := AppEventVideoClipAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AppEventVideoClipAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AppEventVideoClipAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && !IsNil(o.PreviewFrameTimeCode) {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppEventVideoClipAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl) {
		var ret string
		return ret
	}
	return *o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetVideoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VideoUrl) {
		return nil, false
	}
	return o.VideoUrl, true
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasVideoUrl() bool {
	if o != nil && !IsNil(o.VideoUrl) {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given string and assigns it to the VideoUrl field.
func (o *AppEventVideoClipAttributes) SetVideoUrl(v string) {
	o.VideoUrl = &v
}

// GetPreviewImage returns the PreviewImage field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetPreviewImage() ImageAsset {
	if o == nil || IsNil(o.PreviewImage) {
		var ret ImageAsset
		return ret
	}
	return *o.PreviewImage
}

// GetPreviewImageOk returns a tuple with the PreviewImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetPreviewImageOk() (*ImageAsset, bool) {
	if o == nil || IsNil(o.PreviewImage) {
		return nil, false
	}
	return o.PreviewImage, true
}

// HasPreviewImage returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasPreviewImage() bool {
	if o != nil && !IsNil(o.PreviewImage) {
		return true
	}

	return false
}

// SetPreviewImage gets a reference to the given ImageAsset and assigns it to the PreviewImage field.
func (o *AppEventVideoClipAttributes) SetPreviewImage(v ImageAsset) {
	o.PreviewImage = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || IsNil(o.UploadOperations) {
		var ret []UploadOperation
		return ret
	}
	return o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetUploadOperationsOk() ([]UploadOperation, bool) {
	if o == nil || IsNil(o.UploadOperations) {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasUploadOperations() bool {
	if o != nil && !IsNil(o.UploadOperations) {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *AppEventVideoClipAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || IsNil(o.AssetDeliveryState) {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || IsNil(o.AssetDeliveryState) {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasAssetDeliveryState() bool {
	if o != nil && !IsNil(o.AssetDeliveryState) {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *AppEventVideoClipAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

// GetAppEventAssetType returns the AppEventAssetType field value if set, zero value otherwise.
func (o *AppEventVideoClipAttributes) GetAppEventAssetType() AppEventAssetType {
	if o == nil || IsNil(o.AppEventAssetType) {
		var ret AppEventAssetType
		return ret
	}
	return *o.AppEventAssetType
}

// GetAppEventAssetTypeOk returns a tuple with the AppEventAssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipAttributes) GetAppEventAssetTypeOk() (*AppEventAssetType, bool) {
	if o == nil || IsNil(o.AppEventAssetType) {
		return nil, false
	}
	return o.AppEventAssetType, true
}

// HasAppEventAssetType returns a boolean if a field has been set.
func (o *AppEventVideoClipAttributes) HasAppEventAssetType() bool {
	if o != nil && !IsNil(o.AppEventAssetType) {
		return true
	}

	return false
}

// SetAppEventAssetType gets a reference to the given AppEventAssetType and assigns it to the AppEventAssetType field.
func (o *AppEventVideoClipAttributes) SetAppEventAssetType(v AppEventAssetType) {
	o.AppEventAssetType = &v
}

func (o AppEventVideoClipAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventVideoClipAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.PreviewFrameTimeCode) {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	if !IsNil(o.VideoUrl) {
		toSerialize["videoUrl"] = o.VideoUrl
	}
	if !IsNil(o.PreviewImage) {
		toSerialize["previewImage"] = o.PreviewImage
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

type NullableAppEventVideoClipAttributes struct {
	value *AppEventVideoClipAttributes
	isSet bool
}

func (v NullableAppEventVideoClipAttributes) Get() *AppEventVideoClipAttributes {
	return v.value
}

func (v *NullableAppEventVideoClipAttributes) Set(val *AppEventVideoClipAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventVideoClipAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventVideoClipAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventVideoClipAttributes(val *AppEventVideoClipAttributes) *NullableAppEventVideoClipAttributes {
	return &NullableAppEventVideoClipAttributes{value: val, isSet: true}
}

func (v NullableAppEventVideoClipAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventVideoClipAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


