/* -*- fill-column: 80 -*-
Copyright 2021.

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

// AppChartSpec defines the desired state of AppChart
type AppChartSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// This is similar to the ServiceSpec. Main difference is that this does
	// not support custom values.

	// Name of the app support chart (i.e. standard)
	Name string `json:"name,omitempty"`

	// ShortDescription of the chart to be used in lists
	ShortDescription string `json:"shortDescription,omitempty"`

	// Description of the chart, long form for detailed display
	Description string `json:"description,omitempty"`

	// HelmRepo is the Helm repository where to fetch the helm chart. This
	// can be empty. In that case the HelmChart field has to reference the
	// chart as full URL instead of as a simple name.
	HelmRepo AppHelmRepo `json:"helmRepo,omitempty"`

	// HelmChart is the name of the Helm chart used to deploy the service
	HelmChart string `json:"chart,omitempty"`
}

// AppHelmRepo is the Helm repository where to fetch the app helm chart
type AppHelmRepo struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	// See service_types.go for a structurally identical definition
	// (HelmRepo) we cannot reuse, lest the code generator fails to create
	// the necessary DeepCopy support.
}

// AppChartStatus defines the observed state of AppChart
type AppChartStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AppChart is the Schema for the appcharts API
type AppChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppChartSpec   `json:"spec,omitempty"`
	Status AppChartStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppChartList contains a list of AppChart
type AppChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppChart `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppChart{}, &AppChartList{})
}
