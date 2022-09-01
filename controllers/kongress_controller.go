/*
Copyright 2022.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	hepaerdacloudv1 "github.com/erda-project/hepa-operator/api/v1"
)

// KongressReconciler reconciles a Kongress object and is not implemented yet.
type KongressReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=kongresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=kongresses/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=kongresses/finalizers,verbs=update

//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get;update;patch

//+kubebuilder:rbac:groups={""},resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups={""},resources=services/status,verbs=get;update;patch

//+kubebuilder:rbac:groups={"", "events.k8s.io"},resources=events,verbs=get;list;watch;create;update;patch;delete

// Reconcile is not implemented yet
func (r *KongressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KongressReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hepaerdacloudv1.Kongress{}).
		Complete(r)
}
