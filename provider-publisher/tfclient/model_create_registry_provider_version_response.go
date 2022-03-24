/*
Terraform Private Registry

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tfclient

import (
	"encoding/json"
)

// CreateRegistryProviderVersionResponse struct for CreateRegistryProviderVersionResponse
type CreateRegistryProviderVersionResponse struct {
	Data *RegistryProviderVersionResponseData `json:"data,omitempty"`
}

// NewCreateRegistryProviderVersionResponse instantiates a new CreateRegistryProviderVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRegistryProviderVersionResponse() *CreateRegistryProviderVersionResponse {
	this := CreateRegistryProviderVersionResponse{}
	return &this
}

// NewCreateRegistryProviderVersionResponseWithDefaults instantiates a new CreateRegistryProviderVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRegistryProviderVersionResponseWithDefaults() *CreateRegistryProviderVersionResponse {
	this := CreateRegistryProviderVersionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateRegistryProviderVersionResponse) GetData() RegistryProviderVersionResponseData {
	if o == nil || o.Data == nil {
		var ret RegistryProviderVersionResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRegistryProviderVersionResponse) GetDataOk() (*RegistryProviderVersionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateRegistryProviderVersionResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RegistryProviderVersionResponseData and assigns it to the Data field.
func (o *CreateRegistryProviderVersionResponse) SetData(v RegistryProviderVersionResponseData) {
	o.Data = &v
}

func (o CreateRegistryProviderVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRegistryProviderVersionResponse struct {
	value *CreateRegistryProviderVersionResponse
	isSet bool
}

func (v NullableCreateRegistryProviderVersionResponse) Get() *CreateRegistryProviderVersionResponse {
	return v.value
}

func (v *NullableCreateRegistryProviderVersionResponse) Set(val *CreateRegistryProviderVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRegistryProviderVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRegistryProviderVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRegistryProviderVersionResponse(val *CreateRegistryProviderVersionResponse) *NullableCreateRegistryProviderVersionResponse {
	return &NullableCreateRegistryProviderVersionResponse{value: val, isSet: true}
}

func (v NullableCreateRegistryProviderVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRegistryProviderVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
