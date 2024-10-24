// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// PermissionList PermissionList is a list of Permissions
// NODESCRIPTION PermissionList
//
// Deprecated: This model is deprecated.
type PermissionList struct {
	// Items is the list of Permission objects in the list
	Items []Permission `json:"items"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPermissionList instantiates a new PermissionList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPermissionList(items []Permission) *PermissionList {
	this := PermissionList{}
	this.Items = items
	return &this
}

// NewPermissionListWithDefaults instantiates a new PermissionList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPermissionListWithDefaults() *PermissionList {
	this := PermissionList{}
	return &this
}

// GetItems returns the Items field value.
func (o *PermissionList) GetItems() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PermissionList) GetItemsOk() (*[]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *PermissionList) SetItems(v []Permission) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PermissionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PermissionList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]Permission `json:"items"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items"})
	} else {
		return err
	}
	o.Items = *all.Items

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
