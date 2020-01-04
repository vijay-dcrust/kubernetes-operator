/*

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

// GenericDaemonSpec defines the desired state of GenericDaemon
type GenericDaemonSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GenericDaemon. Edit GenericDaemon_types.go to remove/update
	Label string `json:"label"`
	Image string `json:"image"`
}

// GenericDaemonStatus defines the observed state of GenericDaemon
type GenericDaemonStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Count int32 `json:"count"`
}

// +kubebuilder:object:root=true

// GenericDaemon is the Schema for the genericdaemons API
type GenericDaemon struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GenericDaemonSpec   `json:"spec,omitempty"`
	Status GenericDaemonStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GenericDaemonList contains a list of GenericDaemon
type GenericDaemonList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GenericDaemon `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GenericDaemon{}, &GenericDaemonList{})
}
