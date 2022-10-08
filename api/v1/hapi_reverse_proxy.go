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

package v1

import (
	"context"
	"reflect"
	"sort"
	"strings"

	"github.com/pkg/errors"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/erda-project/hepa-operator/pkg/middleware"
	"github.com/erda-project/hepa-operator/pkg/mse"
)

var (
	IngressClassNameMSE   = "mse"
	IngressClassNameNginx = "nginx"
)

type ReverseProxyRule struct {
	hapi   *Hapi
	global *Policy
}

func NewReverseProxyRule(hapi *Hapi, global *Policy) *ReverseProxyRule {
	return &ReverseProxyRule{hapi: hapi, global: global}
}

func (in *ReverseProxyRule) GetHapi() *Hapi {
	return in.hapi
}

func (in *ReverseProxyRule) GetLocal() *Policy {
	return &in.hapi.Spec.Policy
}

func (in *ReverseProxyRule) GetGlobal() *Policy {
	return in.global
}

func (in *ReverseProxyRule) GetPolicies() []string {
	var (
		policies []string
		locals   = in.GetLocal().ListAll()
		globals  = in.GetGlobal().ListAll()
	)
	for name, item := range locals {
		if item.GetGlobal() {
			item = globals[name]
			name = strings.ToUpper(name)
		}
		if item.GetSwitch() {
			policies = append(policies, name)
		}
	}
	sort.Strings(policies)
	return policies
}

func (in *ReverseProxyRule) PatchPolicies(ctx context.Context, ingress *netv1.Ingress, mw middleware.Middleware) error {
	if len(ingress.Annotations) == 0 {
		ingress.Annotations = make(map[string]string)
	}
	locals := in.GetLocal().ListAll()
	globals := in.GetGlobal().ListAll()
	for name, item := range locals {
		if item.GetGlobal() {
			item = globals[name]
			name = strings.ToUpper(name)
		}
		if item != nil && !reflect.ValueOf(item).IsNil() && item.GetSwitch() {
			if err := item.SwitchOnForIngress(ctx, ingress); err != nil {
				return errors.Wrapf(err, "failed to %s.SwitchOnForIngress", name)
			}
			if err := item.SwitchOnForMiddleware(ctx, mw); err != nil {
				return errors.Wrapf(err, "failed to %s.SwitchOnForMiddleware", name)
			}
		} else {
			if err := item.SwitchOffForIngress(ctx, ingress); err != nil {
				return errors.Wrapf(err, "failed to %s.SwitchOffForIngress", name)
			}
			if err := item.SwitcherOffForMiddleware(ctx, mw); err != nil {
				return errors.Wrapf(err, "failed to %s.SwitcherOffForMiddleware", name)
			}
		}
	}

	return nil
}

func (in *ReverseProxyRule) Patch(ctx context.Context, ingress *netv1.Ingress, request *mse.PluginRequest) error {
	if in.hapi.Spec.Backend.RedirectBy == RedirectByService {
		return in.PatchBackend(ctx, ingress, request)
	}
	return in.PatchExternal(ctx, ingress, request)
}

// todo: ut
func (in *ReverseProxyRule) PatchExternal(ctx context.Context, ingress *netv1.Ingress, request *mse.PluginRequest) error {
	var (
		pathType        = netv1.PathTypeImplementationSpecific
		target          = intstr.FromString(in.hapi.Spec.Backend.GetUpstreamHostPort(80))
		httpIngressPath = netv1.HTTPIngressPath{
			Path:     in.hapi.Spec.Path,
			PathType: &pathType,
			Backend: netv1.IngressBackend{
				Service: &netv1.IngressServiceBackend{
					Name: in.hapi.GetExternalServiceName(),
					Port: netv1.ServiceBackendPort{Number: int32(target.IntValue())},
				},
				Resource: nil,
			},
		}
		ingressRuleValue = netv1.IngressRuleValue{
			HTTP: &netv1.HTTPIngressRuleValue{Paths: []netv1.HTTPIngressPath{httpIngressPath}},
		}
	)
	*ingress = netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      in.hapi.GetName(),
			Namespace: in.hapi.GetNamespace(),
			Annotations: map[string]string{
				AnnotationUpstreamVHost.String(): in.hapi.Spec.Backend.GetUpstreamHostName(),
				AnnotationRewriteTarget.String(): in.hapi.Spec.Backend.RewriteTarget,
			},
			Labels: in.hapi.GetLabels(),
		},
		Spec: netv1.IngressSpec{
			IngressClassName: &IngressClassNameMSE,
			TLS:              []netv1.IngressTLS{{Hosts: in.hapi.Spec.Hosts}},
			Rules:            nil,
		},
	}
	for i := 0; i < len(in.hapi.Spec.Hosts); i++ {
		ingress.Spec.Rules = append(ingress.Spec.Rules, netv1.IngressRule{
			Host:             in.hapi.Spec.Hosts[i],
			IngressRuleValue: ingressRuleValue,
		})
	}
	return in.PatchPolicies(ctx, ingress, request)
}

func (in *ReverseProxyRule) PatchBackend(ctx context.Context, ingress *netv1.Ingress, request *mse.PluginRequest) error {
	var (
		pathType        = netv1.PathTypeImplementationSpecific
		httpIngressPath = netv1.HTTPIngressPath{
			Path:     in.hapi.Spec.Path,
			PathType: &pathType,
			Backend: netv1.IngressBackend{
				Service: &netv1.IngressServiceBackend{
					Name: in.hapi.Spec.Backend.ServiceName,
					Port: netv1.ServiceBackendPort{Number: int32(in.hapi.Spec.Backend.ServicePort)},
				},
				Resource: nil,
			},
		}
		ingressRuleValue = netv1.IngressRuleValue{
			HTTP: &netv1.HTTPIngressRuleValue{Paths: []netv1.HTTPIngressPath{httpIngressPath}},
		}
	)
	*ingress = netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      in.hapi.GetName(),
			Namespace: in.hapi.GetNamespace(),
			Labels:    in.hapi.GetLabels(),
		},
		Spec: netv1.IngressSpec{
			IngressClassName: &IngressClassNameMSE,
			TLS:              []netv1.IngressTLS{{Hosts: in.hapi.Spec.Hosts}},
			Rules:            nil,
		},
	}
	for i := 0; i < len(in.hapi.Spec.Hosts); i++ {
		ingress.Spec.Rules = append(ingress.Spec.Rules, netv1.IngressRule{
			Host:             in.hapi.Spec.Hosts[i],
			IngressRuleValue: ingressRuleValue,
		})
	}

	return in.PatchPolicies(ctx, ingress, request)
}
