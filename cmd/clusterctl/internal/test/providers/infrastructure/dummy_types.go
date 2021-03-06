/*
Copyright 2020 The Kubernetes Authors.

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

/*
package infrastructure defines the types for dummy infrastructure provider used for tests
*/
package infrastructure

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

type DummyInfrastructureCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// +kubebuilder:object:root=true

type DummyInfrastructureClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DummyInfrastructureCluster `json:"items"`
}

// +kubebuilder:object:root=true

type DummyInfrastructureMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// +kubebuilder:object:root=true

type DummyInfrastructureMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DummyInfrastructureMachine `json:"items"`
}

// +kubebuilder:object:root=true

type DummyInfrastructureMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
}

// +kubebuilder:object:root=true

type DummyInfrastructureMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DummyInfrastructureMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&DummyInfrastructureCluster{}, &DummyInfrastructureClusterList{},
		&DummyInfrastructureMachine{}, &DummyInfrastructureMachineList{},
		&DummyInfrastructureMachineTemplate{}, &DummyInfrastructureMachineTemplateList{},
	)
}
