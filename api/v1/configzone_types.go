/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	ConfigZoneLabelKey = "configZone"
)

// ConfigZoneSpec defines the desired state of ConfigZone
type ConfigZoneSpec struct {
	Scene  string `json:"scene,omitempty"`
	Hosts  Hosts  `json:"hosts,omitempty"`
	Policy Policy `json:"policy,omitempty"`
}

// ConfigZoneStatus defines the observed state of ConfigZone
type ConfigZoneStatus struct {
	Phase      StatusPhase    `json:"phase,omitempty"`
	HapisCount int            `json:"hapisCount,omitempty"`
	Hapis      []string       `json:"hapis,omitempty"`
	Policies   []string       `json:"policies,omitempty"`
	Spec       ConfigZoneSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConfigZone is the Schema for the configzones API
// +kubebuilder:resource:shortName={cz,czr}
// +kubebuilder:printcolumn:name="Scene",type=string,JSONPath=`.spec.scene`
// +kubebuilder:printcolumn:name="Hosts",type=string,JSONPath=`.spec.hosts`
// +kubebuilder:printcolumn:name="Hapi_Count",type=integer,JSONPath=`.status.hapisCount`
// +kubebuilder:printcolumn:name="Policies",type=string,JSONPath=`.status.policies`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
type ConfigZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigZoneSpec   `json:"spec,omitempty"`
	Status ConfigZoneStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConfigZoneList contains a list of ConfigZone
type ConfigZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigZone `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigZone{}, &ConfigZoneList{})
}
