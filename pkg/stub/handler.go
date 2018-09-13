package stub

import (
	"context"

	"github.com/rekcah78/tenant-operator/pkg/apis/tenant/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.Tenant:
		err := sdk.Create(newNamespace(o))
		if err != nil && !errors.IsAlreadyExists(err) {
			logrus.Errorf("Failed to create namespace : %v", err)
			return err
		}
		err_rq := sdk.Create(newResourceQuota(o))
		if err_rq != nil && !errors.IsAlreadyExists(err_rq) {
			logrus.Errorf("Failed to create resource quota : %v", err_rq)
			return err_rq
		}
	}
	return nil
}

func newNamespace(cr *v1alpha1.Tenant) *corev1.Namespace {
	return &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Spec.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Tenant",
				}),
			},
		},
	}
}

func newResourceQuota(cr *v1alpha1.Tenant) *corev1.ResourceQuota {
	return &corev1.ResourceQuota{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ResourceQuota",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Spec.Namespace,
			Namespace: cr.Spec.Namespace,
		},
		Spec: cr.Spec.ResourceQuota.Spec,
	}
}
