// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"
)

type SshConfigCheckRequest struct {
	// List of nodes to perform preflight checks
	Nodes      []SshConfigCheckSpec `json:"nodes"`
	DefaultSsh *SshConfig           `json:"defaultSSH,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSshConfigCheckRequest instantiates a new SshConfigCheckRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSshConfigCheckRequest(nodes []SshConfigCheckSpec) *SshConfigCheckRequest {
	this := SshConfigCheckRequest{}
	this.Nodes = nodes
	return &this
}

// NewSshConfigCheckRequestWithDefaults instantiates a new SshConfigCheckRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSshConfigCheckRequestWithDefaults() *SshConfigCheckRequest {
	this := SshConfigCheckRequest{}
	return &this
}

// GetNodes returns the Nodes field value.
func (o *SshConfigCheckRequest) GetNodes() []SshConfigCheckSpec {
	if o == nil {
		var ret []SshConfigCheckSpec
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *SshConfigCheckRequest) GetNodesOk() (*[]SshConfigCheckSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value.
func (o *SshConfigCheckRequest) SetNodes(v []SshConfigCheckSpec) {
	o.Nodes = v
}

// GetDefaultSsh returns the DefaultSsh field value if set, zero value otherwise.
func (o *SshConfigCheckRequest) GetDefaultSsh() SshConfig {
	if o == nil || o.DefaultSsh == nil {
		var ret SshConfig
		return ret
	}
	return *o.DefaultSsh
}

// GetDefaultSshOk returns a tuple with the DefaultSsh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshConfigCheckRequest) GetDefaultSshOk() (*SshConfig, bool) {
	if o == nil || o.DefaultSsh == nil {
		return nil, false
	}
	return o.DefaultSsh, true
}

// HasDefaultSsh returns a boolean if a field has been set.
func (o *SshConfigCheckRequest) HasDefaultSsh() bool {
	return o != nil && o.DefaultSsh != nil
}

// SetDefaultSsh gets a reference to the given SshConfig and assigns it to the DefaultSsh field.
func (o *SshConfigCheckRequest) SetDefaultSsh(v SshConfig) {
	o.DefaultSsh = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SshConfigCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["nodes"] = o.Nodes
	if o.DefaultSsh != nil {
		toSerialize["defaultSSH"] = o.DefaultSsh
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SshConfigCheckRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Nodes      *[]SshConfigCheckSpec `json:"nodes"`
		DefaultSsh *SshConfig            `json:"defaultSSH,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Nodes == nil {
		return fmt.Errorf("required field nodes missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodes", "defaultSSH"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Nodes = *all.Nodes
	if all.DefaultSsh != nil && all.DefaultSsh.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DefaultSsh = all.DefaultSsh

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
