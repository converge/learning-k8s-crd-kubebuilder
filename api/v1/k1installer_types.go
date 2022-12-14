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

// K1InstallerSpec defines the desired state of K1Installer
type K1InstallerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of K1Installer. Edit k1installer_types.go to remove/update
	Color string `json:"color,omitempty"`
}

// K1InstallerStatus defines the observed state of K1Installer
type K1InstallerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="color",type=string,JSONPath=`.spec.color`
// +kubebuilder:printcolumn:name="status",type=string,JSONPath=`.status.status`

// K1Installer is the Schema for the k1installers API
type K1Installer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   K1InstallerSpec   `json:"spec,omitempty"`
	Status K1InstallerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// K1InstallerList contains a list of K1Installer
type K1InstallerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K1Installer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K1Installer{}, &K1InstallerList{})
}
