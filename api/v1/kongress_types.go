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

// KongressSpec defines the desired state of Kongress
type KongressSpec struct {
	Route    KongRoute         `json:"route,omitempty"`
	Services []KongService     `json:"services,omitempty"`
	Plugins  map[string]string `json:"plugins,omitempty"`
}

// KongressStatus defines the observed state of Kongress
type KongressStatus struct {
	Route    KongRoute     `json:"route,omitempty"`
	Services []KongService `json:"services,omitempty"`
	//Plugins  map[string]string `json:"plugins,omitempty"`
}

type KongRoute struct {
	Name      string   `json:"name,omitempty"`
	Protocols []string `json:"protocols,omitempty"`
	Methods   []string `json:"methods,omitempty"`
	Paths     []string `json:"paths,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type KongService struct {
	Name     string   `json:"name,omitempty"`
	Protocol string   `json:"protocol,omitempty"`
	Host     string   `json:"host,omitempty"`
	Port     int      `json:"port,omitempty"`
	Path     string   `json:"path,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

type KongPlugin struct {
	Name      string   `json:"name,omitempty"`
	Config    string   `json:"config,omitempty"`
	Protocols []string `json:"protocols,omitempty"`
	Enabled   bool     `json:"enabled,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Kongress is the Schema for the kongresses API
type Kongress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KongressSpec   `json:"spec,omitempty"`
	Status KongressStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KongressList contains a list of Kongress
type KongressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kongress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kongress{}, &KongressList{})
}
