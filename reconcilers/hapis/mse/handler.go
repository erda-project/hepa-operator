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
	"encoding/json"
	"reflect"
	"strconv"
	"strings"

	v1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	hepav1 "github.com/erda-project/hepa-operator/api/v1"
	k8sutil "github.com/erda-project/hepa-operator/pkg/k8s"
	"github.com/erda-project/hepa-operator/pkg/mse"
)

func (r *Reconciler) RegisterHandlers() {
	if len(r.handlers) == 0 {
		r.handlers = make(map[hepav1.StatusPhase]func(ctx context.Context, instance *hepav1.Hapi) error)
	}
	valueOf := reflect.ValueOf(r)
	typeOf := reflect.TypeOf(r)
	for i := 0; i < valueOf.NumMethod(); i++ {
		method := valueOf.Method(i)
		if f, ok := method.Interface().(func(ctx context.Context, instance *hepav1.Hapi) error); ok &&
			strings.HasPrefix(typeOf.Method(i).Name, "Handle") {
			r.handlers[hepav1.StatusPhase(strings.TrimPrefix(typeOf.Method(i).Name, "Handle"))] = f
		}
	}
	for k := range r.handlers {
		log.FromContext(context.Background()).WithValues("handler", k).Info("registered handler")
	}
}

func (r *Reconciler) HandleOK(ctx context.Context, instance *hepav1.Hapi) (err error) {
	l := r.Log.WithValues("hapi", client.ObjectKeyFromObject(instance))
	l.Info("HandleOK")
	reason, ok := r.needToReconcile(ctx, instance)
	if !ok {
		instance.Status.Phase = hepav1.OK
		return nil
	}
	r.Log.WithValues("reason", reason).Info("to reconcile")

	if instance.Spec.Backend.RedirectBy == hepav1.RedirectByService {
		instance.Status.Phase = hepav1.ReconcileBackendService
	} else {
		instance.Status.Phase = hepav1.ReconcileBackendUpstream
	}
	return nil
}

func (r *Reconciler) HandleReconcileBackendService(ctx context.Context, instance *hepav1.Hapi) (err error) {
	l := r.Log.WithValues(
		"hapi", client.ObjectKeyFromObject(instance),
		"current phase", instance.Status.Phase,
	)
	l.Info("HandleReconcileBackendService")
	l.Info("deleteExternalNameService")
	if err = r.deleteExternalNameService(ctx, instance); err != nil {
		l.Error(err, "failed to deleteExternalNameService")
		return err
	}
	l.Info("getConfigZone")
	cm, _, err := r.getConfigZone(ctx, instance)
	if err != nil {
		l.Error(err, "failed to getConfigZone")
		return err
	}
	l.Info("getGlobalPolicy")
	globalPolicy, _ := r.getGlobalPolicy(ctx, cm)
	if err != nil {
		l.Error(err, "failed to getGlobalPolicy")
		return err
	}
	l.Info("makeTheReverseProxyEffective")
	ingressVersion, err := r.makeTheReverseProxyEffective(ctx, instance.ReverseProxyRule(globalPolicy))
	if err != nil {
		l.Error(err, "failed to makeTheReverseProxyEffective")
		return err
	}

	// change status
	l.Info("modifyStatusToService")
	resourceVersion := hepav1.HapiStatusResourceVersion{Ingress: ingressVersion}
	resourceVersion.ConfigZone = k8sutil.GetResourceVersion(cm)
	r.modifyStatusToService(ctx, instance, resourceVersion)
	l.WithValues("next phase", instance.Status.Phase).Info("turn to next status")
	return nil
}

func (r *Reconciler) HandleReconcileBackendUpstream(ctx context.Context, instance *hepav1.Hapi) (err error) {
	l := r.Log.WithValues(
		"hapi", client.ObjectKeyFromObject(instance),
		"current phase", instance.Status.Phase,
	)
	l.Info("HandleReconcileBackendUpstream")
	l.Info("makeTheExternalNameServiceEffective")
	serviceVersion, err := r.makeTheExternalNameServiceEffective(ctx, instance)
	if err != nil {
		l.Error(err, "failed to makeTheExternalNameServiceEffective")
		return err
	}
	l.Info("getConfigZone")
	configZone, _, err := r.getConfigZone(ctx, instance)
	if err != nil {
		l.Error(err, "failed to getConfigZone")
		return err
	}
	l.Info("getGlobalPolicy")
	globalPolicy, _ := r.getGlobalPolicy(ctx, configZone)
	if err != nil {
		l.Error(err, "failed to getGlobalPolicy")
		return err
	}
	l.Info("makeTheReverseProxyEffective")
	ingressVersion, err := r.makeTheReverseProxyEffective(ctx, instance.ReverseProxyRule(globalPolicy))
	if err != nil {
		l.Error(err, "failed to makeTheReverseProxyEffective")
		return err
	}

	// change status
	l.Info("modifyStatusToUpstream")
	resourceVersion := hepav1.HapiStatusResourceVersion{Service: serviceVersion, Ingress: ingressVersion}
	resourceVersion.ConfigZone = k8sutil.GetResourceVersion(configZone)
	r.modifyStatusToUpstream(ctx, instance, resourceVersion)
	l.WithValues("next phase", instance.Status.Phase).Info("turn to next status")
	return nil
}

func (r *Reconciler) HandleDefault(ctx context.Context, instance *hepav1.Hapi) (err error) {
	instance.Status.Phase = hepav1.OK
	return nil
}

func (r *Reconciler) needToReconcile(ctx context.Context, instance *hepav1.Hapi) (string, bool) {
	hapiStat, err := r.GetHapiStat(ctx, instance)
	if err != nil {
		return err.Error(), true
	}
	statusStat, err := r.GetHapiStatusStat(ctx, instance)
	if err != nil {
		return err.Error(), true
	}
	if fields, ok := hepav1.StatEqual(hapiStat, statusStat); !ok {
		return strings.Join(fields, ".") + " has been changed", true
	}
	return "", false
}

func (r *Reconciler) makeTheExternalNameServiceEffective(ctx context.Context, instance *hepav1.Hapi) (resourceVersion string, err error) {
	var (
		l = r.Log.WithValues("Hapi", client.ObjectKeyFromObject(instance),
			"current phase", instance.Status.Phase,
			"Service", types.NamespacedName{Namespace: instance.GetNamespace(), Name: instance.GetExternalServiceName()})
		service = v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      instance.GetExternalServiceName(),
				Namespace: instance.GetNamespace(),
				Labels:    make(map[string]string),
			},
			Spec: v1.ServiceSpec{
				Ports:        nil,
				Type:         v1.ServiceTypeExternalName,
				ExternalName: instance.Spec.Backend.GetUpstreamHostName(),
			},
		}
	)
	if portStr := instance.Spec.Backend.GetUpstreamHostPort(0); portStr == "" {
		service.Spec.Ports = []v1.ServicePort{
			{Protocol: "TCP", Port: 80, TargetPort: intstr.FromInt(80), Name: "target0"},
			{Protocol: "TCP", Port: 443, TargetPort: intstr.FromInt(443), Name: "target1"},
		}
	} else {
		targetPort := intstr.FromString(portStr)
		service.Spec.Ports = []v1.ServicePort{
			{Protocol: "TCP", Port: int32(targetPort.IntValue()), TargetPort: targetPort, Name: "target"},
		}
	}
	op, err := ctrl.CreateOrUpdate(ctx, r.Client, &service, func() error {
		(&service).Labels = instance.GetLabels()
		return r.setControllerReference(ctx, instance, &service)
	})
	if err != nil {
		l.Error(err, "failed to CreateOrUpdate the service")
		return "", err
	}
	r.EventRecorder.Eventf(instance, "Normal", string(op), "CreateOrUpdate Service %s", client.ObjectKeyFromObject(&service))

	return service.GetResourceVersion(), nil
}

func (r *Reconciler) makeTheReverseProxyEffective(ctx context.Context, rule *hepav1.ReverseProxyRule) (resourceVersion string, err error) {
	var (
		l = r.Log.WithValues(
			"hapi", client.ObjectKeyFromObject(rule.GetHapi()),
			"current phase", rule.GetHapi().Status.Phase,
		)
		ingress netv1.Ingress
		pr      mse.PluginRequest
	)
	if err = rule.Patch(ctx, &ingress, &pr); err != nil {
		return "", err
	}
	rule.GetHapi().Status.Policies = rule.GetPolicies()
	l = l.WithValues("ingress", client.ObjectKeyFromObject(&ingress))
	op, err := ctrl.CreateOrUpdate(ctx, r.Client, &ingress, func() error {
		(&ingress).Labels = rule.GetHapi().GetLabels()
		return r.setControllerReference(ctx, rule.GetHapi(), &ingress)
	})
	if err != nil {
		data, _ := json.Marshal(ingress)
		l.Error(err, "failed to CreateOrUpdate the ingress", "ingress", string(data))
		return "", err
	}
	l.Info("CreateOrUpdate ingress", " OperationResult", op)
	r.EventRecorder.Eventf(rule.GetHapi(), "Normal", string(op), "CreateOrUpdate Ingress %s", client.ObjectKeyFromObject(&ingress))

	return ingress.GetResourceVersion(), nil
}

func (r *Reconciler) deleteExternalNameService(ctx context.Context, instance *hepav1.Hapi) (err error) {
	var service = &v1.Service{ObjectMeta: metav1.ObjectMeta{
		Name:      instance.GetExternalServiceName(),
		Namespace: instance.GetNamespace(),
	}}
	if err := client.IgnoreNotFound(r.Delete(ctx, service)); err != nil {
		r.EventRecorder.Eventf(instance, "Normal", "delete", "failed to Delete Service %s", client.ObjectKeyFromObject(service))
		return err
	}
	r.EventRecorder.Eventf(instance, "Normal", string("delete"), "Delete Service %s", client.ObjectKeyFromObject(service))
	return nil
}

func (r *Reconciler) modifyStatusToService(ctx context.Context, instance *hepav1.Hapi, resourceVersion hepav1.HapiStatusResourceVersion) {
	instance.Spec.DeepCopyInto(&instance.Status.Spec)
	instance.Status.Spec.Backend.UpstreamHost = ""
	instance.Status.Spec.Backend.RewriteTarget = ""

	instance.Status.Endpoint = instance.Status.Spec.Hosts[0] + instance.Status.Spec.Path
	if runes := []rune(instance.Status.Spec.Path); len(runes) > 40 {
		path := append(runes[:20], []rune("...")...)
		path = append(path, runes[len(runes)-15:]...)
		instance.Status.Endpoint = instance.Status.Spec.Hosts[0] + string(path)
	}
	instance.Status.RedirectTo = instance.Status.Spec.Backend.ServiceName + "." + instance.GetNamespace() + ":" +
		strconv.FormatInt(int64(instance.Spec.Backend.ServicePort), 10)
	if instance.Spec.Backend.ServicePort == 80 {
		instance.Status.RedirectTo = instance.Status.Spec.Backend.ServiceName + "." + instance.GetNamespace()
	}
	instance.Status.SetServiceResourceVersion(resourceVersion.Service)
	instance.Status.SetIngressResourceVersion(resourceVersion.Ingress)
	instance.Status.SetConfigZoneResourceVersion(resourceVersion.ConfigZone)
	instance.Status.Phase = hepav1.OK
}

func (r *Reconciler) modifyStatusToUpstream(ctx context.Context, instance *hepav1.Hapi, resourceVersion hepav1.HapiStatusResourceVersion) {
	instance.Spec.DeepCopyInto(&instance.Status.Spec)
	instance.Status.Spec.Backend.ServiceName = ""
	instance.Status.Spec.Backend.ServicePort = 0

	instance.Status.Endpoint = instance.Status.Spec.Hosts[0] + instance.Status.Spec.Path
	if runes := []rune(instance.Status.Spec.Path); len(runes) > 40 {
		path := append(runes[:20], []rune("...")...)
		path = append(path, runes[len(runes)-15:]...)
		instance.Status.Endpoint = instance.Status.Spec.Hosts[0] + string(path)
	}
	rewriteTarget := instance.Spec.Backend.RewriteTarget
	if runes := []rune(rewriteTarget); len(runes) > 40 {
		path := append(runes[:20], []rune("...")...)
		path = append(path, runes[len(runes)-15:]...)
		rewriteTarget = instance.Status.Spec.Hosts[0] + string(path)
	}
	instance.Status.RedirectTo = instance.Status.Spec.Backend.UpstreamHost + rewriteTarget
	if instance.Status.Spec.Backend.GetUpstreamHostPort(80) == "80" {
		instance.Status.RedirectTo = instance.Status.Spec.Backend.GetUpstreamHostName() + rewriteTarget
	}
	instance.Status.SetServiceResourceVersion(resourceVersion.Service)
	instance.Status.SetIngressResourceVersion(resourceVersion.Ingress)
	instance.Status.SetConfigZoneResourceVersion(resourceVersion.ConfigZone)
	instance.Status.Phase = hepav1.OK
}

func (r *Reconciler) getConfigZone(ctx context.Context, instance *hepav1.Hapi) (*hepav1.ConfigZone, bool, error) {
	if len(instance.GetLabels()) == 0 {
		return nil, false, nil
	}
	configZoneName, ok := instance.GetLabels()[hepav1.ConfigZoneLabelKey]
	if !ok {
		return nil, false, nil
	}
	var configZone hepav1.ConfigZone
	switch err := r.Get(ctx, types.NamespacedName{Namespace: instance.GetNamespace(), Name: configZoneName}, &configZone); {
	case err == nil:
		return &configZone, true, nil
	case apierrors.IsNotFound(err):
		return nil, false, nil
	default:
		return nil, false, err
	}
}

func (r *Reconciler) getGlobalPolicy(ctx context.Context, configZone *hepav1.ConfigZone) (*hepav1.Policy, bool) {
	if configZone == nil {
		return nil, false
	}
	return &configZone.Spec.Policy, true
}

func (r *Reconciler) setControllerReference(ctx context.Context, owner, controlled metav1.Object) error {
	err := ctrl.SetControllerReference(owner, controlled, r.Scheme)
	if err != nil && !strings.Contains(err.Error(), "cross-namespace owner references are disallowed") {
		return err
	}
	return nil
}

func (r *Reconciler) getHandler(phase hepav1.StatusPhase) func(ctx context.Context, instance *hepav1.Hapi) (err error) {
	if h, ok := r.handlers[phase]; ok {
		return h
	}
	return r.HandleDefault
}

func (r *Reconciler) GetHapiStat(ctx context.Context, instance *hepav1.Hapi) (interface{}, error) {
	var getResourceVersionFunc = func() (*hepav1.HapiStatusResourceVersion, error) {
		var (
			service    v1.Service
			ingress    netv1.Ingress
			configZone hepav1.ConfigZone
		)
		if instance.Spec.Backend.RedirectBy == hepav1.RedirectByUrl {
			if err := r.Get(ctx, types.NamespacedName{Namespace: instance.GetNamespace(), Name: instance.GetExternalServiceName()}, &service); err != nil {
				return nil, err
			}
		}
		if err := r.Get(ctx, types.NamespacedName{Namespace: instance.GetNamespace(), Name: instance.GetName()}, &ingress); err != nil {
			return nil, err
		}
		if err := r.Get(ctx, types.NamespacedName{Namespace: instance.GetNamespace(), Name: k8sutil.GetLabel(instance, hepav1.ConfigZoneLabelKey)}, &configZone); err != nil && !apierrors.IsNotFound(err) {
			return nil, err
		}
		rv := &hepav1.HapiStatusResourceVersion{
			Service:    service.GetResourceVersion(),
			Ingress:    ingress.GetResourceVersion(),
			ConfigZone: configZone.GetResourceVersion(),
		}
		return rv, nil
	}
	return r.getStatFunc(ctx, getResourceVersionFunc)(instance.Spec)
}

func (r *Reconciler) GetHapiStatusStat(ctx context.Context, instance *hepav1.Hapi) (interface{}, error) {
	var getResourceVersionFunc = func() (*hepav1.HapiStatusResourceVersion, error) {
		rv := &instance.Status.ResourceVersion
		return rv, nil
	}
	return r.getStatFunc(ctx, getResourceVersionFunc)(instance.Status.Spec)
}

func (r *Reconciler) getStatFunc(ctx context.Context, getResourceVersionFunc func() (*hepav1.HapiStatusResourceVersion, error)) func(spec hepav1.HapiSpec) (interface{}, error) {
	return func(spec hepav1.HapiSpec) (interface{}, error) {
		resourceVersion, err := getResourceVersionFunc()
		if err != nil {
			return nil, err
		}
		var baseStat = hepav1.BaseStat{
			Hosts:           spec.Hosts.String(),
			Path:            spec.Path,
			RedirectBy:      spec.Backend.RedirectBy,
			Policy:          spec.Policy,
			ResourceVersion: *resourceVersion,
		}
		if spec.Backend.RedirectBy == hepav1.RedirectByService {
			return hepav1.RedirectByServiceStat{
				BaseStat:    baseStat,
				ServiceName: spec.Backend.ServiceName,
				ServicePort: spec.Backend.ServicePort,
			}, nil
		}
		return hepav1.RedirectByUrlStat{
			BaseStat:      baseStat,
			UpstreamHost:  spec.Backend.UpstreamHost,
			RewriteTarget: spec.Backend.RewriteTarget,
		}, nil
	}
}
