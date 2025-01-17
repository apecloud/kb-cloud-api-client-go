// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// InstanceMetrics instance metrics
type InstanceMetrics struct {
	// the name of the instance
	InstanceName string `json:"instanceName"`
	// cpu with uint cores.
	CpuUsage string `json:"cpuUsage"`
	// memory with uint Gi.
	MemoryUsage string `json:"memoryUsage"`
	// disk usage items
	DiskUsageItems []InstanceDiskUsageItem `json:"diskUsageItems"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInstanceMetrics instantiates a new InstanceMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInstanceMetrics(instanceName string, cpuUsage string, memoryUsage string, diskUsageItems []InstanceDiskUsageItem) *InstanceMetrics {
	this := InstanceMetrics{}
	this.InstanceName = instanceName
	this.CpuUsage = cpuUsage
	this.MemoryUsage = memoryUsage
	this.DiskUsageItems = diskUsageItems
	return &this
}

// NewInstanceMetricsWithDefaults instantiates a new InstanceMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInstanceMetricsWithDefaults() *InstanceMetrics {
	this := InstanceMetrics{}
	return &this
}

// GetInstanceName returns the InstanceName field value.
func (o *InstanceMetrics) GetInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *InstanceMetrics) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceName, true
}

// SetInstanceName sets field value.
func (o *InstanceMetrics) SetInstanceName(v string) {
	o.InstanceName = v
}

// GetCpuUsage returns the CpuUsage field value.
func (o *InstanceMetrics) GetCpuUsage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value
// and a boolean to check if the value has been set.
func (o *InstanceMetrics) GetCpuUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUsage, true
}

// SetCpuUsage sets field value.
func (o *InstanceMetrics) SetCpuUsage(v string) {
	o.CpuUsage = v
}

// GetMemoryUsage returns the MemoryUsage field value.
func (o *InstanceMetrics) GetMemoryUsage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MemoryUsage
}

// GetMemoryUsageOk returns a tuple with the MemoryUsage field value
// and a boolean to check if the value has been set.
func (o *InstanceMetrics) GetMemoryUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MemoryUsage, true
}

// SetMemoryUsage sets field value.
func (o *InstanceMetrics) SetMemoryUsage(v string) {
	o.MemoryUsage = v
}

// GetDiskUsageItems returns the DiskUsageItems field value.
func (o *InstanceMetrics) GetDiskUsageItems() []InstanceDiskUsageItem {
	if o == nil {
		var ret []InstanceDiskUsageItem
		return ret
	}
	return o.DiskUsageItems
}

// GetDiskUsageItemsOk returns a tuple with the DiskUsageItems field value
// and a boolean to check if the value has been set.
func (o *InstanceMetrics) GetDiskUsageItemsOk() (*[]InstanceDiskUsageItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskUsageItems, true
}

// SetDiskUsageItems sets field value.
func (o *InstanceMetrics) SetDiskUsageItems(v []InstanceDiskUsageItem) {
	o.DiskUsageItems = v
}

// MarshalJSON serializes the struct using spec logic.
func (o InstanceMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["instanceName"] = o.InstanceName
	toSerialize["cpuUsage"] = o.CpuUsage
	toSerialize["memoryUsage"] = o.MemoryUsage
	toSerialize["diskUsageItems"] = o.DiskUsageItems

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InstanceMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName   *string                  `json:"instanceName"`
		CpuUsage       *string                  `json:"cpuUsage"`
		MemoryUsage    *string                  `json:"memoryUsage"`
		DiskUsageItems *[]InstanceDiskUsageItem `json:"diskUsageItems"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.InstanceName == nil {
		return fmt.Errorf("required field instanceName missing")
	}
	if all.CpuUsage == nil {
		return fmt.Errorf("required field cpuUsage missing")
	}
	if all.MemoryUsage == nil {
		return fmt.Errorf("required field memoryUsage missing")
	}
	if all.DiskUsageItems == nil {
		return fmt.Errorf("required field diskUsageItems missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"instanceName", "cpuUsage", "memoryUsage", "diskUsageItems"})
	} else {
		return err
	}
	o.InstanceName = *all.InstanceName
	o.CpuUsage = *all.CpuUsage
	o.MemoryUsage = *all.MemoryUsage
	o.DiskUsageItems = *all.DiskUsageItems

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
