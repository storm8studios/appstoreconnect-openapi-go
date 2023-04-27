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

// checks if the AppPreviewAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewAttributes{}

// AppPreviewAttributes struct for AppPreviewAttributes
type AppPreviewAttributes struct {
	FileSize *int32 `json:"fileSize,omitempty"`
	FileName *string `json:"fileName,omitempty"`
	SourceFileChecksum *string `json:"sourceFileChecksum,omitempty"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	MimeType *string `json:"mimeType,omitempty"`
	VideoUrl *string `json:"videoUrl,omitempty"`
	PreviewImage *ImageAsset `json:"previewImage,omitempty"`
	UploadOperations []UploadOperation `json:"uploadOperations,omitempty"`
	AssetDeliveryState *AppMediaAssetState `json:"assetDeliveryState,omitempty"`
}

// NewAppPreviewAttributes instantiates a new AppPreviewAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewAttributes() *AppPreviewAttributes {
	this := AppPreviewAttributes{}
	return &this
}

// NewAppPreviewAttributesWithDefaults instantiates a new AppPreviewAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewAttributesWithDefaults() *AppPreviewAttributes {
	this := AppPreviewAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *AppPreviewAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AppPreviewAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetSourceFileChecksum returns the SourceFileChecksum field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetSourceFileChecksum() string {
	if o == nil || IsNil(o.SourceFileChecksum) {
		var ret string
		return ret
	}
	return *o.SourceFileChecksum
}

// GetSourceFileChecksumOk returns a tuple with the SourceFileChecksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetSourceFileChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.SourceFileChecksum) {
		return nil, false
	}
	return o.SourceFileChecksum, true
}

// HasSourceFileChecksum returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasSourceFileChecksum() bool {
	if o != nil && !IsNil(o.SourceFileChecksum) {
		return true
	}

	return false
}

// SetSourceFileChecksum gets a reference to the given string and assigns it to the SourceFileChecksum field.
func (o *AppPreviewAttributes) SetSourceFileChecksum(v string) {
	o.SourceFileChecksum = &v
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && !IsNil(o.PreviewFrameTimeCode) {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppPreviewAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *AppPreviewAttributes) SetMimeType(v string) {
	o.MimeType = &v
}

// GetVideoUrl returns the VideoUrl field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetVideoUrl() string {
	if o == nil || IsNil(o.VideoUrl) {
		var ret string
		return ret
	}
	return *o.VideoUrl
}

// GetVideoUrlOk returns a tuple with the VideoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetVideoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VideoUrl) {
		return nil, false
	}
	return o.VideoUrl, true
}

// HasVideoUrl returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasVideoUrl() bool {
	if o != nil && !IsNil(o.VideoUrl) {
		return true
	}

	return false
}

// SetVideoUrl gets a reference to the given string and assigns it to the VideoUrl field.
func (o *AppPreviewAttributes) SetVideoUrl(v string) {
	o.VideoUrl = &v
}

// GetPreviewImage returns the PreviewImage field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetPreviewImage() ImageAsset {
	if o == nil || IsNil(o.PreviewImage) {
		var ret ImageAsset
		return ret
	}
	return *o.PreviewImage
}

// GetPreviewImageOk returns a tuple with the PreviewImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetPreviewImageOk() (*ImageAsset, bool) {
	if o == nil || IsNil(o.PreviewImage) {
		return nil, false
	}
	return o.PreviewImage, true
}

// HasPreviewImage returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasPreviewImage() bool {
	if o != nil && !IsNil(o.PreviewImage) {
		return true
	}

	return false
}

// SetPreviewImage gets a reference to the given ImageAsset and assigns it to the PreviewImage field.
func (o *AppPreviewAttributes) SetPreviewImage(v ImageAsset) {
	o.PreviewImage = &v
}

// GetUploadOperations returns the UploadOperations field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetUploadOperations() []UploadOperation {
	if o == nil || IsNil(o.UploadOperations) {
		var ret []UploadOperation
		return ret
	}
	return o.UploadOperations
}

// GetUploadOperationsOk returns a tuple with the UploadOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetUploadOperationsOk() ([]UploadOperation, bool) {
	if o == nil || IsNil(o.UploadOperations) {
		return nil, false
	}
	return o.UploadOperations, true
}

// HasUploadOperations returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasUploadOperations() bool {
	if o != nil && !IsNil(o.UploadOperations) {
		return true
	}

	return false
}

// SetUploadOperations gets a reference to the given []UploadOperation and assigns it to the UploadOperations field.
func (o *AppPreviewAttributes) SetUploadOperations(v []UploadOperation) {
	o.UploadOperations = v
}

// GetAssetDeliveryState returns the AssetDeliveryState field value if set, zero value otherwise.
func (o *AppPreviewAttributes) GetAssetDeliveryState() AppMediaAssetState {
	if o == nil || IsNil(o.AssetDeliveryState) {
		var ret AppMediaAssetState
		return ret
	}
	return *o.AssetDeliveryState
}

// GetAssetDeliveryStateOk returns a tuple with the AssetDeliveryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewAttributes) GetAssetDeliveryStateOk() (*AppMediaAssetState, bool) {
	if o == nil || IsNil(o.AssetDeliveryState) {
		return nil, false
	}
	return o.AssetDeliveryState, true
}

// HasAssetDeliveryState returns a boolean if a field has been set.
func (o *AppPreviewAttributes) HasAssetDeliveryState() bool {
	if o != nil && !IsNil(o.AssetDeliveryState) {
		return true
	}

	return false
}

// SetAssetDeliveryState gets a reference to the given AppMediaAssetState and assigns it to the AssetDeliveryState field.
func (o *AppPreviewAttributes) SetAssetDeliveryState(v AppMediaAssetState) {
	o.AssetDeliveryState = &v
}

func (o AppPreviewAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.SourceFileChecksum) {
		toSerialize["sourceFileChecksum"] = o.SourceFileChecksum
	}
	if !IsNil(o.PreviewFrameTimeCode) {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
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
	return toSerialize, nil
}

type NullableAppPreviewAttributes struct {
	value *AppPreviewAttributes
	isSet bool
}

func (v NullableAppPreviewAttributes) Get() *AppPreviewAttributes {
	return v.value
}

func (v *NullableAppPreviewAttributes) Set(val *AppPreviewAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewAttributes(val *AppPreviewAttributes) *NullableAppPreviewAttributes {
	return &NullableAppPreviewAttributes{value: val, isSet: true}
}

func (v NullableAppPreviewAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


