package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Tenant `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TenantSpec   `json:"spec"`
	Status            TenantStatus `json:"status,omitempty"`
}

type TenantSpec struct {
	Namespace string `json:"namespace"`
	LimitRange  []v1.LimitRangeSpec `json:"limitrange,omitempty"`
	ResourceQuota *v1.ResourceQuota `json:"resourcequota"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}
type TenantStatus struct {
	Phase string `json:"phase"`
}
