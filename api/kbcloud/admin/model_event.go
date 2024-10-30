// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd


package admin

import (
	"github.com/google/uuid"
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api"

)



// Event Event related to the view 
type Event struct {
	// Group is the API group of the event.
	Group string `json:"group"`
	// Kind is the type of the event.
	Kind string `json:"kind"`
	// Name is the name of the event.
	Name string `json:"name"`
	// Namespace is the namespace of the event.
	Namespace string `json:"namespace"`
	// Owners are the owners of the event.
	Owners []Owner `json:"owners,omitempty"`
	// Progress describes the progress of the event.
	Progress *string `json:"progress,omitempty"`
	// Message is the message contained in the event.
	Message *string `json:"message,omitempty"`
	// ResourceVersion is the version of the resource as seen by the system.
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// Type is the type of the event.
	Type string `json:"type"`
	// Version is the version of the resource that generated the event.
	Version *string `json:"version,omitempty"`
	// Time is the time of the resource that generated the event.
	Time *string `json:"time,omitempty"`
	// Level is the level the event.
	Level *string `json:"level,omitempty"`
	// Reason is the reason of the event.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}


// NewEvent instantiates a new Event object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEvent(group string, kind string, name string, namespace string, typeVar string) *Event {
	this := Event{}
	this.Group = group
	this.Kind = kind
	this.Name = name
	this.Namespace = namespace
	this.Type = typeVar
	return &this
}

// NewEventWithDefaults instantiates a new Event object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}
// GetGroup returns the Group field value.
func (o *Event) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *Event) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *Event) SetGroup(v string) {
	o.Group = v
}


// GetKind returns the Kind field value.
func (o *Event) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *Event) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value.
func (o *Event) SetKind(v string) {
	o.Kind = v
}


// GetName returns the Name field value.
func (o *Event) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Event) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Event) SetName(v string) {
	o.Name = v
}


// GetNamespace returns the Namespace field value.
func (o *Event) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *Event) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *Event) SetNamespace(v string) {
	o.Namespace = v
}


// GetOwners returns the Owners field value if set, zero value otherwise.
func (o *Event) GetOwners() []Owner {
	if o == nil || o.Owners == nil {
		var ret []Owner
		return ret
	}
	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetOwnersOk() (*[]Owner, bool) {
	if o == nil || o.Owners == nil {
		return nil, false
	}
	return &o.Owners, true
}

// HasOwners returns a boolean if a field has been set.
func (o *Event) HasOwners() bool {
	return o != nil && o.Owners != nil
}

// SetOwners gets a reference to the given []Owner and assigns it to the Owners field.
func (o *Event) SetOwners(v []Owner) {
	o.Owners = v
}


// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *Event) GetProgress() string {
	if o == nil || o.Progress == nil {
		var ret string
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetProgressOk() (*string, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *Event) HasProgress() bool {
	return o != nil && o.Progress != nil
}

// SetProgress gets a reference to the given string and assigns it to the Progress field.
func (o *Event) SetProgress(v string) {
	o.Progress = &v
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Event) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Event) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Event) SetMessage(v string) {
	o.Message = &v
}


// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *Event) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *Event) HasResourceVersion() bool {
	return o != nil && o.ResourceVersion != nil
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *Event) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}


// GetType returns the Type field value.
func (o *Event) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Event) SetType(v string) {
	o.Type = v
}


// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Event) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Event) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Event) SetVersion(v string) {
	o.Version = &v
}


// GetTime returns the Time field value if set, zero value otherwise.
func (o *Event) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Event) HasTime() bool {
	return o != nil && o.Time != nil
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *Event) SetTime(v string) {
	o.Time = &v
}


// GetLevel returns the Level field value if set, zero value otherwise.
func (o *Event) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *Event) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *Event) SetLevel(v string) {
	o.Level = &v
}


// GetReason returns the Reason field value if set, zero value otherwise.
func (o *Event) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *Event) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *Event) SetReason(v string) {
	o.Reason = &v
}



// MarshalJSON serializes the struct using spec logic.
func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["group"] = o.Group
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	if o.Owners != nil {
		toSerialize["owners"] = o.Owners
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ResourceVersion != nil {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	toSerialize["type"] = o.Type
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Event) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Group *string `json:"group"`
		Kind *string `json:"kind"`
		Name *string `json:"name"`
		Namespace *string `json:"namespace"`
		Owners []Owner `json:"owners,omitempty"`
		Progress *string `json:"progress,omitempty"`
		Message *string `json:"message,omitempty"`
		ResourceVersion *string `json:"resourceVersion,omitempty"`
		Type *string `json:"type"`
		Version *string `json:"version,omitempty"`
		Time *string `json:"time,omitempty"`
		Level *string `json:"level,omitempty"`
		Reason *string `json:"reason,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.Kind == nil {
		return fmt.Errorf("required field kind missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{ "group", "kind", "name", "namespace", "owners", "progress", "message", "resourceVersion", "type", "version", "time", "level", "reason",  })
	} else {
		return err
	}
	o.Group = *all.Group
	o.Kind = *all.Kind
	o.Name = *all.Name
	o.Namespace = *all.Namespace
	o.Owners = all.Owners
	o.Progress = all.Progress
	o.Message = all.Message
	o.ResourceVersion = all.ResourceVersion
	o.Type = *all.Type
	o.Version = all.Version
	o.Time = all.Time
	o.Level = all.Level
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
