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
	"encoding/json"
	"net/http"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	apiv1 "github.com/erda-project/hepa-operator/api/v1"
)

// ConfigZoneReconciler reconciles a ConfigZone object
type ConfigZoneReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	Log           logr.Logger
	Reconciler    reconcile.Reconciler
	EventRecorder record.EventRecorder
}

func NewConfigZoneReconciler(cli client.Client, scheme *runtime.Scheme, log logr.Logger, eventRecord record.EventRecorder, options ...ReconcilerOption) *ConfigZoneReconciler {
	var r = &ConfigZoneReconciler{
		Client:     cli,
		Scheme:     scheme,
		Log:        log,
		Reconciler: nil,
	}
	for _, opt := range options {
		opt(r)
	}
	return r
}

//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=configzones,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=configzones/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=configzones/finalizers,verbs=update

//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get;update;patch

//+kubebuilder:rbac:groups={""},resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups={""},resources=services/status,verbs=get;update;patch

//+kubebuilder:rbac:groups={"", "events.k8s.io"},resources=events,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ConfigZone object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *ConfigZoneReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.Reconciler.Reconcile(ctx, req)
}

// SetupWithManager sets up the controller with the Manager.
func (r *ConfigZoneReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.ConfigZone{}).
		Complete(r)
}

func (r *ConfigZoneReconciler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		list apiv1.ConfigZoneList
	)
	if err := r.List(req.Context(), &list, GetListOptions(req)...); err != nil && !apierrors.IsNotFound(err) {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	data, _ := json.Marshal(list)
	_, _ = w.Write(data)
	w.WriteHeader(http.StatusOK)
}
