/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type IntegrationCloudwatch struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationCloudwatchSpec   `json:"spec,omitempty"`
	Status            IntegrationCloudwatchStatus `json:"status,omitempty"`
}

type IntegrationCloudwatchSpec struct {
	State *IntegrationCloudwatchSpecResource `json:"state,omitempty" tf:"-"`

	Resource IntegrationCloudwatchSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type IntegrationCloudwatchSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalTags *map[string]string `json:"additionalTags,omitempty" tf:"additional_tags"`
	ExternalID     *string            `json:"externalID" tf:"external_id"`
	// +optional
	ForceSave *bool `json:"forceSave,omitempty" tf:"force_save"`
	// +optional
	InstanceSelectionTags *map[string]string `json:"instanceSelectionTags,omitempty" tf:"instance_selection_tags"`
	// +optional
	MetricFilterRegex *string `json:"metricFilterRegex,omitempty" tf:"metric_filter_regex"`
	Name              *string `json:"name" tf:"name"`
	// +optional
	Namespaces []string `json:"namespaces,omitempty" tf:"namespaces"`
	// +optional
	PointTagFilterRegex *string `json:"pointTagFilterRegex,omitempty" tf:"point_tag_filter_regex"`
	RoleArn             *string `json:"roleArn" tf:"role_arn"`
	Service             *string `json:"service" tf:"service"`
	// +optional
	ServiceRefreshRateInMinutes *int64 `json:"serviceRefreshRateInMinutes,omitempty" tf:"service_refresh_rate_in_minutes"`
	// +optional
	VolumeSelectionTags *map[string]string `json:"volumeSelectionTags,omitempty" tf:"volume_selection_tags"`
}

type IntegrationCloudwatchStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IntegrationCloudwatchList is a list of IntegrationCloudwatchs
type IntegrationCloudwatchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IntegrationCloudwatch CRD objects
	Items []IntegrationCloudwatch `json:"items,omitempty"`
}
