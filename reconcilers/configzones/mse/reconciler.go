// Copyright (c) 2022 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package mse

import (
	"context"
	"sort"
	"strings"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	hepav1 "github.com/erda-project/hepa-operator/api/v1"
	"github.com/erda-project/hepa-operator/pkg/interfaces"
)

// Reconciler is not implemented
type Reconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

func New(cli client.Client, scheme *runtime.Scheme, log logr.Logger) *Reconciler {
	var r = &Reconciler{
		Client: cli,
		Scheme: scheme,
		Log:    log.WithName("reconciler/mse"),
	}
	return r
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (ctrl.Result, error) {
	l := r.Log.WithValues("cz", req.NamespacedName)

	var instance hepav1.ConfigZone
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.WithValues("current phase", instance.Status.Phase).Info("get the ConfigZone")
	instance.Status.Phase = hepav1.OK
	if err := r.findTheReferencedHapis(&instance); err != nil {
		l.Error(err, "failed to findTheReferencedHapis")
		return ctrl.Result{}, err
	}
	instance.Status.HapisCount = len(instance.Status.Hapis)
	r.patchSwitchedOnPolicies(&instance)
	instance.Spec.DeepCopyInto(&instance.Status.Spec)

	return ctrl.Result{}, r.Status().Update(ctx, &instance)
}

func (r *Reconciler) findTheReferencedHapis(instance *hepav1.ConfigZone) error {
	var hapis hepav1.HapiList
	if err := r.List(context.Background(), &hapis, &client.ListOptions{
		LabelSelector: labels.SelectorFromSet(labels.Set{hepav1.ConfigZoneLabelKey: instance.GetName()}),
		Namespace:     instance.GetNamespace(),
	}); err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	instance.Status.Hapis = nil
	for _, item := range hapis.Items {
		instance.Status.Hapis = append(instance.Status.Hapis, item.GetName())
	}
	instance.Status.HapisCount = len(instance.Status.Hapis)
	sort.Strings(instance.Status.Hapis)
	return nil
}

func (r *Reconciler) patchSwitchedOnPolicies(instance *hepav1.ConfigZone) {
	instance.Status.Policies = nil
	polices := []interfaces.SwitchGetter{
		&instance.Spec.Policy.Auth,
		&instance.Spec.Policy.CORS,
		&instance.Spec.Policy.Metric,
		&instance.Spec.Policy.Proxy,
		&instance.Spec.Policy.SafetyCSRF,
		&instance.Spec.Policy.SafetyIP,
		&instance.Spec.Policy.SafetySBAC,
		&instance.Spec.Policy.ServerGuard,
	}
	policesNames := []string{
		"Auth",
		"CORS",
		"Metric",
		"Proxy",
		"SafetyCSRF",
		"SafetyIP",
		"SafetySBAC",
		"ServerGuard",
	}
	for i := 0; i < len(polices); i++ {
		if polices[i].GetSwitch() {
			instance.Status.Policies = append(instance.Status.Policies, strings.ToUpper(policesNames[i]))
		}
	}
	sort.Strings(instance.Status.Policies)
}
