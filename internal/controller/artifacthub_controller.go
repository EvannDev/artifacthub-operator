/*
Copyright 2025.

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

package controller

import (
	"context"
	"google.golang.org/protobuf/proto"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	corev1alpha1 "github.com/EvannDev/artifacthub-operator/api/v1alpha1"
)

// ArtifactHubReconciler reconciles a ArtifactHub object
type ArtifactHubReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=core.artifacthub.io,resources=artifacthubs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core.artifacthub.io,resources=artifacthubs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=core.artifacthub.io,resources=artifacthubs/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=create;get;update;patch;watch;delete;list

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ArtifactHub object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.4/pkg/reconcile
func (r *ArtifactHubReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := logf.FromContext(ctx)

	var artifacthub corev1alpha1.ArtifactHub
	if err := r.Get(ctx, req.NamespacedName, &artifacthub); err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	var version = artifacthub.Spec.Version
	switch {
	case version == "":
		version = "latest"
	}
	image := "artifacthub/hub:" + version

	var deploy = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      artifacthub.Name + "-deployment",
			Namespace: artifacthub.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: proto.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": artifacthub.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": artifacthub.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "artifacthub",
						Image: image,
					}},
				},
			},
		},
	}
	var found appsv1.Deployment
	err := r.Get(ctx, client.ObjectKey{Name: deploy.Name, Namespace: deploy.Namespace}, &found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Creating Deployment", "Namespace", deploy.Namespace, "Name", deploy.Name)
		if err := r.Create(ctx, deploy); err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		return ctrl.Result{}, err
	}

	artifacthub.Status.Ready = found.Status.Conditions[1].Status
	artifacthub.Status.Message = found.Status.Conditions[1].Message
	if err := r.Status().Update(ctx, &artifacthub); err != nil {
		logger.Error(err, "unable to update ArtifactHub status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ArtifactHubReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1alpha1.ArtifactHub{}).
		Named("artifacthub").
		Complete(r)
}
