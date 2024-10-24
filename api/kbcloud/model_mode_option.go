// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2019-Present ApeCloud, Inc.

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ModeOption struct {
	// NODESCRIPTION Name
	Name string `json:"name"`
	// NODESCRIPTION Title
	Title LocalizedDescription `json:"title"`
	// NODESCRIPTION Description
	Description LocalizedDescription `json:"description"`
	// NODESCRIPTION Components
	Components []ModeComponent `json:"components"`
	// NODESCRIPTION Proxy
	Proxy *ModeOptionProxy `json:"proxy,omitempty"`
	// NODESCRIPTION Extra
	Extra map[string]interface{} `json:"extra,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModeOption instantiates a new ModeOption object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModeOption(name string, title LocalizedDescription, description LocalizedDescription, components []ModeComponent) *ModeOption {
	this := ModeOption{}
	this.Name = name
	this.Title = title
	this.Description = description
	this.Components = components
	return &this
}

// NewModeOptionWithDefaults instantiates a new ModeOption object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModeOptionWithDefaults() *ModeOption {
	this := ModeOption{}
	return &this
}

// GetName returns the Name field value.
func (o *ModeOption) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ModeOption) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value.
func (o *ModeOption) GetTitle() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetTitleOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ModeOption) SetTitle(v LocalizedDescription) {
	o.Title = v
}

// GetDescription returns the Description field value.
func (o *ModeOption) GetDescription() LocalizedDescription {
	if o == nil {
		var ret LocalizedDescription
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetDescriptionOk() (*LocalizedDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ModeOption) SetDescription(v LocalizedDescription) {
	o.Description = v
}

// GetComponents returns the Components field value.
func (o *ModeOption) GetComponents() []ModeComponent {
	if o == nil {
		var ret []ModeComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *ModeOption) GetComponentsOk() (*[]ModeComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *ModeOption) SetComponents(v []ModeComponent) {
	o.Components = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *ModeOption) GetProxy() ModeOptionProxy {
	if o == nil || o.Proxy == nil {
		var ret ModeOptionProxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetProxyOk() (*ModeOptionProxy, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *ModeOption) HasProxy() bool {
	return o != nil && o.Proxy != nil
}

// SetProxy gets a reference to the given ModeOptionProxy and assigns it to the Proxy field.
func (o *ModeOption) SetProxy(v ModeOptionProxy) {
	o.Proxy = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *ModeOption) GetExtra() map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModeOption) GetExtraOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return &o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *ModeOption) HasExtra() bool {
	return o != nil && o.Extra != nil
}

// SetExtra gets a reference to the given map[string]interface{} and assigns it to the Extra field.
func (o *ModeOption) SetExtra(v map[string]interface{}) {
	o.Extra = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	toSerialize["components"] = o.Components
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModeOption) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name        *string                `json:"name"`
		Title       *LocalizedDescription  `json:"title"`
		Description *LocalizedDescription  `json:"description"`
		Components  *[]ModeComponent       `json:"components"`
		Proxy       *ModeOptionProxy       `json:"proxy,omitempty"`
		Extra       map[string]interface{} `json:"extra,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "title", "description", "components", "proxy", "extra"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = *all.Title
	if all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = *all.Description
	o.Components = *all.Components
	if all.Proxy != nil && all.Proxy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Proxy = all.Proxy
	o.Extra = all.Extra

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
