/*
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
	"k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CronJobSpec defines the desired state of CronJob
type CronJobSpec struct {
	// +kubebuilder:validation:Minlength=0
	// schedule in the ** cron format
	Schedule string `json:"schedule"`

	// +kubebuilder:validation:Minimum=0
	// +optional
	// deadline to start job if missed.
	DeadlineSec *int64 `json:"deadlineSec, omitempty"`

	// In case of concurrent executions, choose one of the following:
	// Allow, Forbid, Replace
	// +optional
	ConcurrencyPolicy ConcurrencyPolicy `json:"concurrecyPolicy, omitempty"`

	// Suspend current execution
	// +optional
	Suspend *bool `json:"suspend, omitempty"`

	// Job type to be run
	JobTemplate v1beta1.JobTemplateSpec `json:"jobTemplate"`

	// +kubebuilder:validation:Minimum=0
	// limit of number of jobs that were successful
	SuccessHistoryLimit *int32 `json:"successHistoryLimit, omitempty"`

	// +kubebuilder:validation:Minimum=0
	// Limit of number of jobs that have failed
	FailedHistoryLimit *int32 `json:"failedHistoryLimit, omitempty"`
}

// CronJobStatus defines the observed state of CronJob
type CronJobStatus struct {
	// +optional
	// List of active jobs running currently
	Active []corev1.LocalObjectReference `json:"active, omitempty"`

	// +optional
	// when the last successful schedule ran
	LastScheduleTime *metav1.Time `json:"lastScheduleTime, omitempty`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CronJob is the Schema for the cronjobs API
type CronJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronJobSpec   `json:"spec,omitempty"`
	Status CronJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CronJobList contains a list of CronJob
type CronJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronJob `json:"items"`
}

// +kubebuilder:validation:Enum=Allow;Forbid;Replace
type ConcurrencyPolicy string

const (
	AllowConcurrent   ConcurrencyPolicy = "Allow"
	ForbidConcurrent  ConcurrencyPolicy = "Forbid"
	ReplaceConcurrent ConcurrencyPolicy = "Replace"
)

func init() {
	SchemeBuilder.Register(&CronJob{}, &CronJobList{})
}
