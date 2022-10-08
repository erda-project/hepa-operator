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
	"strings"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	apiv1 "github.com/erda-project/hepa-operator/api/v1"
	czmse "github.com/erda-project/hepa-operator/reconcilers/configzones/mse"
	czkongv1 "github.com/erda-project/hepa-operator/reconcilers/configzones/nginx-kong/v1"
	czkongv2 "github.com/erda-project/hepa-operator/reconcilers/configzones/nginx-kong/v2"
	hapismse "github.com/erda-project/hepa-operator/reconcilers/hapis/mse"
	hapikongv1 "github.com/erda-project/hepa-operator/reconcilers/hapis/nginx-kong/v1"
	hapikongv2 "github.com/erda-project/hepa-operator/reconcilers/hapis/nginx-kong/v2"
)

const (
	GatewayArchMSE    = "mse"
	GatewayArchKongV1 = "nginx-kong.v1"
	GatewayArchKongV2 = "nginx-kong.v2"
)

// HapiReconciler reconciles a Hapi object
type HapiReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	Log           logr.Logger
	EventRecorder record.EventRecorder

	reconciler reconcile.Reconciler
}

func NewHapiReconciler(cli client.Client, scheme *runtime.Scheme, log logr.Logger, recorder record.EventRecorder, options ...ReconcilerOption) *HapiReconciler {
	var r = &HapiReconciler{
		Client:        cli,
		Scheme:        scheme,
		Log:           log,
		EventRecorder: recorder,
		reconciler:    nil,
	}
	for _, opt := range options {
		opt(r)
	}
	return r
}

//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=hapis,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=hapis/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=hapis/finalizers,verbs=update
//+kubebuilder:rbac:groups=hepa.erda.cloud,resources=events,verbs=create;patch

//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses/status,verbs=get;update;patch

//+kubebuilder:rbac:groups={""},resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups={""},resources=services/status,verbs=get;update;patch

//+kubebuilder:rbac:groups={"", "events.k8s.io"},resources=events,verbs=get;list;watch;create;update;patch;delete

func (r *HapiReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.reconciler.Reconcile(ctx, req)
}

// SetupWithManager sets up the controller with the Manager.
func (r *HapiReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.Hapi{}).
		//WithEventFilter(predicate.Funcs{
		//	CreateFunc: func(createEvent event.CreateEvent) bool {
		//		r.Log.Info("CreateEvent happened", "createEvent.Object.GetObjectKind()", createEvent.Object.GetObjectKind())
		//		return true
		//	},
		//	DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
		//		r.Log.Info("DeleteEvent happened")
		//		return true
		//	},
		//	UpdateFunc: func(updateEvent event.UpdateEvent) bool {
		//		r.Log.Info("UpdateEvent happened")
		//		return true
		//	},
		//	GenericFunc: func(genericEvent event.GenericEvent) bool {
		//		r.Log.Info("GenericEvent happened")
		//		return true
		//	},
		//}).
		Owns(&v1.Service{}).
		Owns(&netv1.Ingress{}).
		Watches(
			&source.Kind{Type: &apiv1.ConfigZone{}},
			handler.EnqueueRequestsFromMapFunc(r.findTheReferencedHapisFor),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}

func (r *HapiReconciler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var (
		list apiv1.HapiList
	)
	if err := r.List(req.Context(), &list, GetListOptions(req)...); err != nil && !apierrors.IsNotFound(err) {
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	data, _ := json.Marshal(list)
	_, _ = w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func (r *HapiReconciler) findTheReferencedHapisFor(configZone client.Object) []reconcile.Request {
	l := r.Log.WithName("Watches").WithValues("ConfigZone", client.ObjectKeyFromObject(configZone))
	l.Info("")
	var hapis apiv1.HapiList
	if err := r.List(context.Background(), &hapis, &client.ListOptions{
		LabelSelector: labels.SelectorFromSet(labels.Set{apiv1.ConfigZoneLabelKey: configZone.GetName()}),
		Namespace:     configZone.GetNamespace(),
	}); err != nil {
		if apierrors.IsNotFound(err) {
			l.Info("not found")
		}
		return nil
	}
	var requests = make([]reconcile.Request, len(hapis.Items))
	for i := 0; i < len(hapis.Items); i++ {
		requests[i] = reconcile.Request{
			NamespacedName: client.ObjectKeyFromObject(&hapis.Items[i]),
		}
	}
	l.Info("find the referenced hapis", "count", len(requests))
	return requests
}

type ReconcilerOption func(r interface{})

func WithGatewayArch(arch string) ReconcilerOption {
	switch arch {
	case GatewayArchMSE:
		return func(r interface{}) {
			switch t := r.(type) {
			case *ConfigZoneReconciler:
				t.Reconciler = czmse.New(t.Client, t.Scheme, t.Log)
			case *HapiReconciler:
				t.reconciler = hapismse.New(t.Client, t.Scheme, t.Log, t.EventRecorder)
			default:
				panic("invalid reconciler")
			}
		}
	case GatewayArchKongV1:
		return func(r interface{}) {
			switch t := r.(type) {
			case *ConfigZoneReconciler:
				t.Reconciler = czkongv1.New()
			case *HapiReconciler:
				t.reconciler = hapikongv1.New()
			default:
				panic("invalid reconciler")
			}
		}
	case GatewayArchKongV2:
		return func(r interface{}) {
			switch t := r.(type) {
			case *ConfigZoneReconciler:
				t.Reconciler = czkongv2.New(t.Client, t.Scheme, t.Log, t.EventRecorder)
			case *HapiReconciler:
				t.reconciler = hapikongv2.New(t.Client, t.Scheme, t.Log, t.EventRecorder)
			default:
				panic("invalid reconciler")
			}
		}
	default:
		panic("invalid gateway arch")
	}
}

func GetListOptions(req *http.Request) []client.ListOption {
	var (
		option *client.ListOptions
		sets   = make(labels.Set)
	)
	if labelsList := req.URL.Query()["label"]; len(labelsList) > 0 {
		for _, label := range labelsList {
			if index := strings.Index(label, "="); index > 0 && index < len(label)-1 {
				sets[label[:index]] = label[index+1:]
			}
		}
	}
	if len(sets) > 0 {
		if option == nil {
			option = new(client.ListOptions)
		}
		option.LabelSelector = labels.SelectorFromSet(sets)
	}
	if namespace := req.URL.Query().Get("namespace"); namespace != "" {
		if option == nil {
			option = new(client.ListOptions)
		}
		option.Namespace = namespace
	}
	if option != nil {
		return []client.ListOption{option}
	}
	return nil
}
