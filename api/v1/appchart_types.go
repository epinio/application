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

	// ShortDescription of the chart. To be used in list displays
	ShortDescription string `json:"shortDescription,omitempty"`

	// Description of the chart. Long form to be used in detailed displays
	Description string `json:"description,omitempty"`

	// HelmRepo is the URL to the Helm repository where to fetch the helm
	// chart. This can be empty. In that case the HelmChart field has to
	// reference the chart as full URL instead of as a simple name.
	HelmRepo string `json:"helmRepo,omitempty"`

	// HelmChart is the name of the Helm chart used to deploy an application.
	HelmChart string `json:"helmChart,omitempty"`
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
