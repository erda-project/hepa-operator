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

// HapiStatus defines the observed state of Hapi
type HapiStatus struct {
	Phase           StatusPhase               `json:"phase,omitempty"`
	Endpoint        string                    `json:"endpoint,omitempty"`
	RedirectTo      string                    `json:"redirectTo,omitempty"`
	Policies        []string                  `json:"policies,omitempty"`
	Spec            HapiSpec                  `json:"spec,omitempty"`
	ResourceVersion HapiStatusResourceVersion `json:"resourceVersion,omitempty"`
}

func (in *HapiStatus) SetServiceResourceVersion(version string) {
	in.ResourceVersion.Service = version
}

func (in *HapiStatus) SetIngressResourceVersion(version string) {
	in.ResourceVersion.Ingress = version
}

func (in *HapiStatus) SetConfigZoneResourceVersion(version string) {
	in.ResourceVersion.ConfigZone = version
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Hapi is the Schema for the hapis API
// +kubebuilder:printcolumn:name="Endpoint",type=string,JSONPath=`.status.endpoint`
// +kubebuilder:printcolumn:name="RedirectTo",type=string,JSONPath=`.status.redirectTo`
// +kubebuilder:printcolumn:name="Policies",type=string,JSONPath=`.status.policies`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
type Hapi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HapiSpec   `json:"spec,omitempty"`
	Status HapiStatus `json:"status,omitempty"`
}

func (in *Hapi) ReverseProxyRule(global *Policy) *ReverseProxyRule {
	return &ReverseProxyRule{
		hapi:   in,
		global: global,
	}
}

func (in *Hapi) GetExternalServiceName() string {
	return "external-" + in.GetName()
}

//+kubebuilder:object:root=true

// HapiList contains a list of Hapi
type HapiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hapi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Hapi{}, &HapiList{})
}
