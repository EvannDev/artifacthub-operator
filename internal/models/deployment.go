package models

import (
	corev1alpha1 "github.com/EvannDev/artifacthub-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func CreateDeployment(artifacthub corev1alpha1.ArtifactHub) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
