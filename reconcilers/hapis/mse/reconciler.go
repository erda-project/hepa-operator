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
	"strings"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hepav1 "github.com/erda-project/hepa-operator/api/v1"
)

type Reconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	Log           logr.Logger
	EventRecorder record.EventRecorder

	handlers map[hepav1.StatusPhase]func(ctx context.Context, instance *hepav1.Hapi) (err error)
}

func New(cli client.Client, scheme *runtime.Scheme, logger logr.Logger, recorder record.EventRecorder) *Reconciler {
	var r = &Reconciler{
		Client:        cli,
		Scheme:        scheme,
		Log:           logger.WithName("reconciler/mse"),
		EventRecorder: recorder,
		handlers:      nil,
	}
	r.RegisterHandlers()
	return r
}

func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := r.Log.WithValues("hapi", req.NamespacedName)

	var instance hepav1.Hapi
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.WithValues("current phase", instance.Status.Phase).Info("get the instance")
	if instance.Status.Phase == "" {
		instance.Status.Phase = hepav1.OK
	}

	cpy := (&instance).DeepCopy()
	if err := r.getHandler(instance.Status.Phase)(ctx, cpy); err != nil {
		return ctrl.Result{}, err
	}
	if fields, ok := instance.Status.DeepEqual(cpy.Status); !ok {
		l.Info("status was changed, to .Status().Update(...)", "field", "."+strings.Join(fields, "."))
		if err := r.Client.Status().Update(ctx, cpy); err != nil {
			if apierrors.IsConflict(err) {
				l.Info("conflict to .Status().Update(...)", "error", err)
				return ctrl.Result{Requeue: true}, nil
			}
			l.Error(err, "failed to .Status().Update(...)")
			return ctrl.Result{}, err
		}
	} else {
		l.Info("status no change", "last phase", instance.Status.Phase, "this phase", cpy.Status.Phase)
	}
	requeue := cpy.Status.Phase != hepav1.OK
	if requeue {
		l.Info("requeue", "next phase", cpy.Status.Phase)
	}
	return ctrl.Result{Requeue: requeue}, nil
}
